/*
 * Copyright 2023 Flmelody
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"flag"
	"fmt"
	"image"
	"image-converter/converter"
	"path/filepath"
	"strings"
)

var (
	supportedFormat = map[string]struct{}{converter.BMP: {}, converter.GIF: {}, converter.JPEG: {}, converter.PNG: {}, converter.TIFF: {}, converter.WEBP: {}}
	source          = ""
	sourceFormat    = ""
	target          = ""
	targetFormat    = ""
)

func init() {
	flag.StringVar(&source, "i", "", "source image")
	flag.StringVar(&target, "o", "", "optional target image,if omitted,targetFormat must exist")
	flag.StringVar(&targetFormat, "f", "", "optional target format,if omitted,will be guessed via target.supported format bmp|gif|jpeg|png|tiff|webp")
}
func main() {
	flag.Parse()
	// 源文件
	if len(source) == 0 {
		fmt.Println("source image is empty!")
		return
	} else {
		sourceFormat = filepath.Ext(source)
		if len(sourceFormat) == 0 {
			fmt.Println("source format must be valid!")
			return
		}
		sourceFormat = strings.Split(sourceFormat, ".")[1]
		_, isPresent := supportedFormat[sourceFormat]
		if !isPresent {
			fmt.Println("source format must be valid!")
			return
		}

	}
	// 目标文件格式
	if len(targetFormat) == 0 {
		if len(target) == 0 {
			fmt.Println("target image is empty!")
			return
		}
		targetFormat = filepath.Ext(target)
		if len(targetFormat) == 0 {
			fmt.Println("target format must be valid!")
			return
		}
	}
	split := strings.Split(targetFormat, ".")
	targetFormat = split[len(split)-1]
	_, isPresent := supportedFormat[targetFormat]
	if !isPresent {
		fmt.Println("target format must be valid!")
		return
	}
	// 目标文件
	if len(target) == 0 {
		index := strings.LastIndex(source, sourceFormat)
		target = source[:index] + targetFormat
	}

	fmt.Println(source, sourceFormat, target, targetFormat)
	var metaData image.Image
	var err error
	switch sourceFormat {
	case converter.BMP:
		metaData, err = new(converter.BmpConverter).Decode(source)
	case converter.GIF:
		metaData, err = new(converter.GifConverter).Decode(source)
	case converter.JPEG:
		metaData, err = new(converter.JpegConverter).Decode(source)
	case converter.PNG:
		metaData, err = new(converter.PngConverter).Decode(source)
	case converter.TIFF:
		metaData, err = new(converter.TiffConverter).Decode(source)
	case converter.WEBP:
		metaData, err = new(converter.WebpConverter).Decode(source)
	default:
		_ = fmt.Errorf("unsupported source format")
		return
	}

	switch targetFormat {
	case converter.BMP:
		err = new(converter.BmpConverter).Encode(target, metaData)
	case converter.GIF:
		err = new(converter.BmpConverter).Encode(target, metaData)
	case converter.JPEG:
		err = new(converter.BmpConverter).Encode(target, metaData)
	case converter.PNG:
		err = new(converter.BmpConverter).Encode(target, metaData)
	case converter.TIFF:
		err = new(converter.BmpConverter).Encode(target, metaData)
	case converter.WEBP:
		err = new(converter.BmpConverter).Encode(target, metaData)
	default:
		err = fmt.Errorf("unsupported target format")
		return
	}
	if err != nil {
		fmt.Println("some error occurred", err)
	}
}
