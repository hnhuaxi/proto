
sync:
	@cp -r ../hnhuaxi/platform/proto .
	@rg -t go -l "github.com/hnhuaxi/platform/proto" | xargs -I {} sed -i '' 's|github.com/hnhuaxi/platform/proto|github.com/hnhuaxi/proto/proto|g' {}