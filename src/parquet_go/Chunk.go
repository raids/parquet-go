package parquet_go

import (
//	"bytes"
//	"git.apache.org/thrift.git/lib/go/thrift"
	"log"
	"parquet"
)

type Chunk struct {
	Pages []*Page
	ChunkHeader *parquet.ColumnChunk
}

func PagesToChunk (pages []*Page) *Chunk {
	ln := len(pages)
	var numValues int64= 0
	var totalUncompressedSize int64 = 0
	var totalCompressedSize int64 = 0
	
	for i:=0; i<ln; i++ {
		numValues += int64(pages[i].Header.DataPageHeader.NumValues)
		totalUncompressedSize += int64(pages[i].Header.UncompressedPageSize) + int64(len(pages[i].RawData)) - int64(pages[i].Header.CompressedPageSize)
		totalCompressedSize += int64(len(pages[i].RawData))
	}
	
	chunk := new(Chunk)
	chunk.Pages = pages
	chunk.ChunkHeader = parquet.NewColumnChunk()
	metaData := parquet.NewColumnMetaData()
	metaData.Type = pages[0].DataType
	metaData.Encodings = append(metaData.Encodings, parquet.Encoding_RLE)
	metaData.Encodings = append(metaData.Encodings, parquet.Encoding_BIT_PACKED)
	metaData.Encodings = append(metaData.Encodings, parquet.Encoding_PLAIN)
	metaData.Codec = pages[0].CompressType
	metaData.NumValues = numValues
	metaData.TotalCompressedSize = totalCompressedSize
	metaData.TotalUncompressedSize = totalUncompressedSize
	metaData.PathInSchema = pages[0].DataTable.Path[1:]
	
	chunk.ChunkHeader.MetaData = metaData

	log.Println("PagesToChunk Finished")
	return chunk
}


