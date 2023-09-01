// Code generated by vfsgen; DO NOT EDIT.

// +build bindata

package migration

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			
			modTime: time.Date(2023, 9, 1, 8, 4, 8, 524443174, time.UTC),
			
		},
		"/issue.json": &vfsgen۰CompressedFileInfo{
			name:             "issue.json",
			
			modTime:          time.Date(2023, 9, 1, 8, 4, 8, 524443174, time.UTC),
			
			uncompressedSize: 2525,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xac\x56\xc1\x6e\xe4\x36\x0c\x3d\xdb\x5f\x41\xb8\x01\xa6\x05\x9c\x99\x1e\x7a\xca\xad\x28\x8a\x45\x81\xa0\x2d\xd0\xf6\xb4\x58\x2c\x38\x16\x3d\xe6\x56\x96\xbc\x12\x9d\xd4\x08\xf2\xef\x85\x24\xdb\x75\x3c\xe3\xdd\x49\x91\xa3\x49\x8a\xef\x89\x7a\x24\xfd\x94\x03\x00\x14\xc2\xa2\xa9\xb8\x83\xe2\x17\xef\x7b\x2a\xca\x64\x55\xe4\x2b\xc7\x9d\xb0\x35\xb3\xcf\x03\x7a\x6f\x2b\x46\x21\x05\x62\x01\xc1\x51\x67\x3d\x8b\x75\x03\x3c\xb2\x34\x6c\x00\xa1\xb6\xee\x44\xf0\xed\x3b\x16\xc2\x12\xde\xb1\xdc\xe3\xb1\x04\x92\x6a\xff\xdd\xbe\x28\xf3\x11\x74\xe8\x22\x26\x3a\x87\xc3\x84\xc9\x42\xad\x2f\xee\xe0\x29\xcf\xe6\x00\x7b\xfc\x44\x95\x14\x65\x9e\x15\xa8\x14\x07\x3e\xa8\x7f\x77\xb6\x23\x27\x4c\x21\xba\x46\xed\x29\xf8\xbb\xa5\xf5\x29\xcf\x62\x4e\xd3\xb7\x47\x72\xc9\x90\xad\x6f\xf5\x97\xe1\xcf\x3d\x01\x2b\x32\xc2\x35\x93\x2b\xc1\x91\x46\xe1\x07\x0a\xf7\x93\x86\x16\x37\x0c\xe4\xb3\xff\x88\x8d\x89\x13\xcc\x73\x39\xc2\x75\xd6\x0b\xb9\x8f\xac\xae\x45\x04\x5b\x47\x9c\xde\x93\x83\xc7\xc6\x02\xf6\xd2\x58\x17\x0a\xdc\x10\x70\x28\xfb\x6b\x80\x0d\xb6\xb4\x01\xfd\x2b\xb6\xf4\x7a\x34\x2f\x8e\xcd\x69\x03\x8d\x5a\x64\xbd\x01\xf7\x73\xf0\xbd\x15\xde\xa4\xd1\x4b\x40\x7f\x34\xd6\x09\x2c\x8c\xa0\xd8\x77\x1a\x07\x52\x80\x3e\x02\xc5\xe3\x57\x01\x55\xd6\x08\x19\xd9\x80\xba\xb7\xe6\x54\x42\xdb\x6b\x61\xcd\x86\xca\x25\xea\x55\xe9\x1d\xd5\x1b\xa9\xff\x44\x77\x22\x81\xa3\x43\x53\x35\xc0\xe6\xcb\xe2\xbb\x9c\xbd\x65\x4d\x5e\xac\xb9\x46\x01\x73\xec\x55\x99\xbd\xa0\x6c\x65\xfd\x11\x76\x95\xb6\x9e\xd4\x2e\x3d\x28\x3c\xb2\xd6\x60\xac\x80\x27\x02\x34\x03\x60\x25\xfc\xc0\x32\x4c\xd7\xaa\x7b\xe9\x1d\x95\x60\xa5\x21\xf7\xc8\x9e\x80\x05\xd8\xc3\xce\x76\x64\x76\x23\x21\x32\x7d\x5b\xdc\xc1\xfb\x3c\x1b\x5f\x26\x62\x44\x5f\xfc\x0e\xb1\x45\x9e\x65\x1f\x56\x54\xd9\x7f\xd4\xb6\xfa\x9b\xd4\x26\xdd\xe4\x1e\xc9\x56\x68\xc0\x1a\x3d\xc0\x91\xa0\xb5\x2a\xf4\xa4\x82\xe3\x00\x9d\xe3\x07\xd6\x74\x22\x15\xc5\xeb\x57\x65\x3a\x5a\xab\x09\xcd\x99\x7c\x1c\x85\xe1\xb8\x01\xfd\x53\xf0\x06\x81\x0a\xb7\x1b\x75\x8f\xb6\xda\xba\x16\x83\x06\x0b\x85\x42\xb7\x21\x7a\x0d\xd4\x77\xea\x0b\x40\xf7\xe8\x05\x52\xc8\x1b\x60\x8d\x95\xdf\xd0\x6d\x43\xa0\x03\x5c\x38\x09\xbb\x28\x94\x1d\x54\x0d\x9a\x53\x5a\x12\x93\x38\x46\x0e\x68\x86\xdf\xea\xc5\xbb\x86\x9c\x17\x99\x6d\x51\x9b\xb9\x5d\x38\x6e\x7a\xad\xe7\x90\x0b\xda\xd0\x78\x24\xed\xb7\x8a\xc6\x5e\x42\x77\xa4\xa0\x55\xc9\xa6\x3d\x95\x65\xcb\x2d\x95\xb2\xde\xa4\xae\x4e\xd9\xf7\x9f\xbc\x8d\xba\x7c\x3e\x6b\xfd\xd0\x06\xd6\x7c\x0d\x7e\x8e\xfb\x1f\x0c\xa6\xb3\x9b\x24\xd0\x7b\x3e\x19\xa2\xaf\x91\x98\xe3\x5e\x41\x62\x63\xd8\x60\x1a\xfe\x63\x46\x35\xed\xd5\xc5\xec\x7f\xf1\x43\x30\xcf\x9f\x99\x7b\x9e\x05\xfa\x85\xa3\xcf\x3d\xbb\xa8\xc3\xf7\x2f\x77\xfb\xf9\xee\xbd\xb4\x14\x5f\xae\x92\xf5\xbc\x7f\x39\xe8\xce\x67\xc9\xba\xc1\xd7\x7d\x98\x67\x1f\xf2\xb1\xd2\xe9\x5f\xe6\xc6\x57\x0d\xb5\x18\xae\xd4\x88\x74\x77\x87\x43\x78\x93\xdb\x64\xdd\x5b\x77\x3a\x28\x87\xb5\xdc\x7e\xff\xc3\x21\xd9\xbe\x99\xfe\x82\x6e\xe2\xcf\xc3\x74\x8a\xfe\xc1\xb6\xd3\xb4\xaf\x6c\x7b\x48\x35\x8b\x6f\x3b\xc5\xde\x48\xdc\x1b\xe1\xc0\xc2\x9b\x3f\xe7\xff\x06\x00\x00\xff\xff\x9d\xc3\xf0\xca\xdd\x09\x00\x00"),
		},
		"/label.json": &vfsgen۰CompressedFileInfo{
			name:             "label.json",
			
			modTime:          time.Date(2023, 9, 1, 8, 4, 8, 524443174, time.UTC),
			
			uncompressedSize: 600,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x84\x91\xb1\x6e\xc2\x30\x10\x86\xe7\xf8\x29\x4e\x2e\x63\x20\x1d\x3a\x65\xed\x8a\xaa\xee\x55\x07\x13\x1f\xc9\x21\xc7\x67\xec\x8b\x2a\x84\x78\xf7\xca\x09\x41\xa4\x2a\xed\x66\xdd\x67\x7f\xf7\xfb\xee\xac\x00\x00\xb4\x90\x38\xd4\x35\xe8\xad\xd9\xa1\xd3\xe5\x54\xb5\x98\x9a\x48\x41\x88\xfd\x8d\x81\x49\x89\x1b\x32\x82\x16\x84\xc1\x78\xa0\x94\x06\xdc\xe8\x52\x5d\x55\xa7\x30\x9a\x78\x77\xc0\x46\x66\x95\xb1\x96\xb2\xc7\xb8\xf7\xc8\x01\xa3\x10\x26\x5d\xc3\xde\xb8\x84\xd7\x2b\xe1\x1e\x9c\x55\xa1\xbd\xe9\x71\x3a\xfe\x96\xe6\xcd\xf4\x08\xbc\x07\xe9\x10\x5c\x4e\x56\xc2\xe0\xe9\x38\x20\x7c\x91\x74\xe4\x47\x10\x31\x70\x22\xe1\x78\xca\x01\x8b\x45\xc0\x24\x91\x7c\xab\x55\x71\x29\x55\xa1\x1b\x76\x1c\x1f\x77\x7b\xcd\x18\x1a\xb6\xcb\x9e\xff\x59\x97\x96\x07\xee\x2d\xfb\xb6\x84\x7e\x70\x42\x8e\x3c\x96\x70\xc7\xff\x6a\x30\x8e\xed\x72\x9d\x5e\xc4\xe3\x40\x11\xad\xae\xe1\x63\x9e\xdd\x48\x3e\xe7\xc5\xac\x52\xd3\x61\x6f\xb2\xa3\x13\x09\x75\x55\x1d\x12\xfb\xf5\x54\xdd\x70\x6c\x2b\x1b\xcd\x5e\xd6\xcf\x2f\xd5\x54\x7b\x9a\x97\xb7\xa2\xac\xd5\xd3\x8f\xf3\xa3\x1b\x58\x89\x89\x2d\xca\x0f\xaa\x2e\xea\x3b\x00\x00\xff\xff\xec\x1f\xb4\x2f\x58\x02\x00\x00"),
		},
		"/milestone.json": &vfsgen۰CompressedFileInfo{
			name:             "milestone.json",
			
			modTime:          time.Date(2023, 9, 1, 8, 4, 8, 524443174, time.UTC),
			
			uncompressedSize: 1394,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xac\x54\x41\x8f\xd3\x4c\x0c\x3d\x67\x7e\x85\x35\x5f\xa5\x5e\xd2\xf6\x3b\x70\xea\x0d\xc1\x71\x11\x48\x70\x5b\xed\xc1\x24\x4e\xe2\xd5\x64\x26\x78\x9c\x2d\x51\xb5\xff\x1d\x4d\xd2\x86\x36\xb4\x1c\x80\xab\xfd\xfc\xde\xb3\xc7\x9e\xa3\x01\x00\xb0\xca\xea\xc8\xee\xc1\x7e\x60\x47\x51\x83\x27\x9b\x4f\x99\x92\x62\x21\xdc\x29\x07\x7f\x95\x07\x8c\x31\x14\x8c\x4a\x25\x68\x00\x04\xa1\x2e\x44\xd6\x20\x03\x1c\x58\x1b\xf6\x80\x50\x05\xa9\x69\x6b\x73\x73\x52\x19\xba\x51\x04\x45\x70\x38\x0b\xb0\x52\x1b\xed\x1e\x8e\x26\x9b\x01\xe1\xeb\x33\x15\x6a\x73\x93\x59\x2c\x4b\x4e\xe2\xe8\x3e\x49\xe8\x48\x94\x29\xa1\x2b\x74\x91\x52\xbe\xbb\x8c\x1e\x4d\x76\xd5\xce\xd1\x64\xd9\xb2\x83\xcf\x4d\x10\x85\x8b\x60\xf2\x97\xfd\xd4\x8e\x2a\xec\x6b\x3b\x31\xbd\xe6\x27\xc6\x6b\x92\x5b\xbc\x0f\xc1\xd7\x39\xb4\xbd\x53\x76\xec\x29\xff\x03\x09\x2c\x53\xe5\x1d\xfe\xf7\xa7\x34\x60\xa5\x24\x70\x68\xb8\x68\x40\x1b\x82\x76\x7e\x12\x8e\x10\x5e\x48\xca\x9e\x6e\x0b\x8e\xb1\x2a\x48\x8b\x9a\xa2\x25\x2a\x6d\x94\x5b\x5a\x3a\x29\x84\xd2\xbb\xde\x31\xf2\x2e\x65\x39\x78\x48\xa5\x7f\x27\xd4\x77\xe5\x6f\x84\x1e\x30\x2a\x4c\x90\x7f\xa0\x55\xb8\x10\xef\x4a\x7d\x69\x08\x5c\x92\x4b\x95\xb0\x8e\x8a\x4a\x6b\x28\x1a\xf4\xf5\xb4\xdf\xeb\xa9\x7c\x7d\xf2\x80\x7e\xf8\x58\xd9\x3d\x3c\x9a\x6c\x64\x4f\x9c\x37\x9d\xdd\xb3\x36\x7b\xbb\x51\xee\x7b\xe7\x66\x88\xc9\xb2\xa7\x45\x27\xa3\xbb\x3b\x8d\xbc\x9d\x9d\x02\xc7\xd8\x13\x1c\xd8\x39\xf0\x41\x21\x12\x01\xfa\x01\xb0\x50\x7e\x61\x1d\x80\xfd\xb8\x3e\x55\xaf\xbd\x50\x0e\x41\x1b\x92\x03\x47\x02\xd6\xb4\x49\xeb\xd0\x91\x3f\xf7\x4b\xbe\x6f\x2f\xda\x3d\x0f\xf3\xec\xdf\x26\xac\xbd\xb0\x6a\xb2\xe4\xd6\x0a\x7d\xeb\x59\xc6\xa9\x3f\x5e\x9d\xe6\xad\xb3\xfa\xe5\x0e\x96\xeb\xb8\xdc\x9a\xc5\xcb\x5e\x8f\xc7\x64\x4f\xe6\x34\xb5\xe9\x9f\x59\xc5\xa2\xa1\x16\xd3\x90\x1a\xd5\x6e\xbf\xdb\x3d\xc7\xe0\x37\x53\x74\x1b\xa4\xde\x95\x82\x95\x6e\xfe\x7f\xb3\x9b\x62\xff\x9d\x7f\xa8\x15\x97\x17\x55\xf4\x1d\xdb\xce\xd1\xb6\x08\xed\x6e\x3e\xbd\x6d\xe2\x9a\xf1\x2b\x45\xa9\x69\x7c\xf3\x05\xc2\xbc\x9a\x1f\x01\x00\x00\xff\xff\xc6\xa6\xf2\x9e\x72\x05\x00\x00"),
		},
		"/reaction.json": &vfsgen۰CompressedFileInfo{
			name:             "reaction.json",
			
			modTime:          time.Date(2023, 9, 1, 8, 4, 8, 524443174, time.UTC),
			
			uncompressedSize: 690,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x9c\x51\xcb\x8e\xe2\x30\x10\x3c\x93\xaf\x68\x79\x39\x02\xd9\xc3\x9e\xf8\x88\xd5\x0a\x69\x4f\xab\xd5\xa8\x71\x2a\xc4\x28\xb6\x43\xbb\xa3\x99\x11\xe2\xdf\x47\xce\x83\x09\xd2\x70\x99\xa3\xbb\xab\xab\xca\x55\xd7\x82\x88\xc8\xa8\xd3\x16\x66\x4f\xe6\x00\xb6\xea\x62\x30\x9b\x71\x51\x21\x59\x71\xdd\x30\x5a\xac\x89\x53\x8a\xd6\xb1\xa2\x22\x8d\xc4\x81\x5c\x4a\x3d\x28\x0a\x31\xd9\xe8\x3d\x82\xee\xcc\xa6\x98\xd8\xdf\xbb\x81\x3c\x1e\xcf\xb0\x3a\x53\x73\x55\xb9\xcc\xc5\xed\x1f\x89\x1d\x44\x1d\x92\xd9\x53\xcd\x6d\xc2\x04\xe9\x96\x8b\x6b\xb1\x32\x7d\x82\xbc\xb8\x6a\x7c\x7d\x65\xf0\x6f\x70\x97\x1e\xe4\x2a\x04\x75\xb5\x83\x50\xac\x49\x1b\x50\xbe\xa4\xd7\x26\x12\xf7\xda\x44\xc9\xbe\x1b\x90\x4c\xff\xc9\x5e\x57\x0f\x5e\x43\xef\x8f\x10\x53\xac\x6e\x9b\x59\x37\xb0\xc7\x73\xe5\xdf\xec\xf1\x5d\xb1\xa4\xe2\xc2\x69\x16\xb3\x31\x28\x82\x3e\x97\x3a\xa0\x13\x24\x04\xe5\xa1\x8b\x49\x54\x3e\xab\x7b\xc6\x3e\xa4\x7a\x9b\xc2\x15\x5c\x7a\x27\xc8\x61\xfe\x5b\x44\xbb\x74\x30\x00\xff\xcf\x35\xae\x93\x6d\xe0\x39\x53\x36\xaa\xdd\xbe\x2c\xcf\x29\x86\xed\x38\xdd\x45\x39\x95\x95\x70\xad\xdb\x9f\xbf\xca\x71\xf6\x63\xae\x7a\x3d\x54\x36\x5f\xe1\x8d\x7d\xd7\x62\x67\xa3\x2f\xef\x99\x64\xaa\x3b\x7c\xad\x2c\x27\xe4\x04\xcc\x23\xa0\xb8\x15\x1f\x01\x00\x00\xff\xff\x3d\x35\x68\x23\xb2\x02\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/issue.json"].(os.FileInfo),
		fs["/label.json"].(os.FileInfo),
		fs["/milestone.json"].(os.FileInfo),
		fs["/reaction.json"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
