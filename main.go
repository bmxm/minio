// Copyright (c) 2015-2021 MinIO, Inc.
//
// This file is part of MinIO Object Storage stack
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main // import "github.com/minio/minio"

import (
	"os"

	// MUST be first import.
	_ "github.com/minio/minio/internal/init"

	minio "github.com/minio/minio/cmd"
)

// oss 对象存储，是一种海量、安全、低成本、高可靠的云存储服务，适合存放任意类型的文件。
// 容量和处理能力弹性拓展，多种存储类型供选择，全面优化存储成本。

// 单机调试： minio server /data
// 需要跟两个参数：server 和 路径
//
// 默认用户名密码：minioadmin
func main() {
	minio.Main(os.Args)
}
