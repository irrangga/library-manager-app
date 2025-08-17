package usecase

import (
	"backend/pkg/constant"
	"backend/src/url/entity"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_urlUsecase_CleanupURL(t *testing.T) {
	type args struct {
		ctx   context.Context
		input entity.URL
	}
	tests := []struct {
		name    string
		uc      *urlUsecase
		args    args
		want    entity.URL
		wantErr error
	}{
		{
			name: "CleanupURL redirection without https:// on scheme and www. on host",
			args: args{
				ctx: context.Background(),
				input: entity.URL{
					InitialURL: "BYFOOD.com/food-EXPeriences?query=abc#fragment",
					Operation:  constant.URLOperationRedirection,
				},
			},
			want: entity.URL{
				ProcessedURL: "https://www.byfood.com/food-experiences?query=abc#fragment",
			},
			wantErr: nil,
		},
		{
			name: "CleanupURL redirection without https:// on scheme",
			args: args{
				ctx: context.Background(),
				input: entity.URL{
					InitialURL: "www.BYFOOD.com/food-EXPeriences?query=abc/",
					Operation:  constant.URLOperationRedirection,
				},
			},
			want: entity.URL{
				ProcessedURL: "https://www.byfood.com/food-experiences?query=abc/",
			},
			wantErr: nil,
		},
		{
			name: "CleanupURL redirection without www. on host",
			args: args{
				ctx: context.Background(),
				input: entity.URL{
					InitialURL: "https://BYFOOD.com/food-EXPeriences?query=abc/",
					Operation:  constant.URLOperationRedirection,
				},
			},
			want: entity.URL{
				ProcessedURL: "https://www.byfood.com/food-experiences?query=abc/",
			},
			wantErr: nil,
		},
		{
			name: "CleanupURL canonical with query",
			args: args{
				ctx: context.Background(),
				input: entity.URL{
					InitialURL: "https://BYFOOD.com/food-EXPeriences?query=abc/",
					Operation:  constant.URLOperationCanonical,
				},
			},
			want: entity.URL{
				ProcessedURL: "https://BYFOOD.com/food-EXPeriences",
			},
			wantErr: nil,
		},
		{
			name: "CleanupURL canonical with trailing slash",
			args: args{
				ctx: context.Background(),
				input: entity.URL{
					InitialURL: "https://BYFOOD.com/food-EXPeriences/",
					Operation:  constant.URLOperationCanonical,
				},
			},
			want: entity.URL{
				ProcessedURL: "https://BYFOOD.com/food-EXPeriences",
			},
			wantErr: nil,
		},
		{
			name: "CleanupURL all",
			args: args{
				ctx: context.Background(),
				input: entity.URL{
					InitialURL: "https://BYFOOD.com/food-EXPeriences?query=abc/",
					Operation:  constant.URLOperationAll,
				},
			},
			want:    entity.URL{ProcessedURL: "https://www.byfood.com/food-experiences"},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &urlUsecase{}
			got, err := uc.CleanupURL(tt.args.ctx, tt.args.input)
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, tt.wantErr, err)
		})
	}
}
