package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 uint16 = 32767
    var v2 int32 = -932
    v3 := []int{2147483647}
    v4 := GenS1{F0: 165}
    v5 := GenS1{F0: -9223372036854775808}
    v6 := GenS2{F0: 65535, F1: 42}
    var v7 uint = 4294967295
    var v8 uint32 = 2147483648
    if true {
        var v9 int64 = -370
        _ = v9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        v3 = append(v3, fi10)
        if true {
            var v11 int = fi10
            for fi12 := 0; fi12 < 17; fi12++ {
                switch -720 {
                case 9223372036854775807:
                    v13 := GenS3{F0: 4294967295}
                    _ = v13
                case -2147483648:
                    var v14 int32 = -2147483648
                    _ = v14
                case -9223372036854775808:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                default:
                    v15 := GenS1{F0: -514}
                    _ = v15
                }
                var v16 float32 = ((-2.25 - 0.0) / 1.0)
                var v17 bool = true
                _ = fi12
                _ = v16
                _ = v17
            }
            if true {
                _ = len(v3)
                v1 = (^(v1))
                v3 = append(v3, (v11 * (fi10 | fi10)))
            }
            _ = v11
        }
        var v18 int = fi10
        _ = fi10
        _ = v18
    }
    fmt.Println(v2, v8, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
