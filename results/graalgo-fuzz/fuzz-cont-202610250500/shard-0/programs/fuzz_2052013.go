package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint32; F1 int16; F2 string }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 int32 }
func genFn0(p0 bool) int8 {
    var v4 uint = 184
    v5 := GenS1{F0: 2, F1: -674}
    _ = p0
    _ = v4
    _ = v5
    return 100
}

func main() {
    var v1 uint16 = 2
    v2 := []int{32768, 32768, 2147483647, -131}
    v3 := make(map[int]bool)
    fmt.Println(v1)
    fmt.Println(v1)
    switch ("" + "a") {
    case "a":
        switch "hello" {
        case "hello":
            var v6 uint64 = 2
            _ = v6
        case "":
            v1 += 32767
        default:
            v1 += (v1 % (((v1 & 862)) | 1))
        }
    case "hello":
        switch ("" + "\n\t") {
        case "foo bar":
            for fi7 := 0; fi7 < 13; fi7++ {
                delete(v3, fi7)
                switch (fi7 + fi7) {
                case 97:
                    var v8 uint16 = 32767
                    _ = v8
                case 128:
                    v9 := GenS1{F0: 416, F1: (fi7 % ((fi7) | 1))}
                    _ = v9
                case 636:
                    if (1e10 < 1e10) {
                        var v10 int64 = 9223372036854775807
                        v1 = (v1 | (v1 << ((0) & 63)))
                        var v11 int8 = -1
                        _ = v10
                        _ = v11
                    }
                default:
                    var v12 uint32 = 2
                    _ = v12
                }
                _ = fi7
            }
        }
    }
    v1 -= 42
    v1 = v1
    if (-32768 <= -1) {
        for fi13 := 0; fi13 < 11; fi13++ {
            fmt.Println(v1, v1)
            v1 = (v1 - 65535)
            _ = fi13
        }
    }
    v14 := GenS2{F0: 2147483648, F1: 1, F2: ""}
    v1 += v1
    v15 := v3[-9223372036854775808]
    var v16 uint8 = 218
    for fi17 := 0; fi17 < 1; fi17++ {
        var v18 float32 = 1e10
        _ = fi17
        _ = v18
    }
    fmt.Println(v16)
    v19 := GenS3{F0: 4294967295, F1: 100}
    for fi20 := 0; fi20 < 10; fi20++ {
        v16 -= v16
        _ = fi20
    }
    var v21 uint32 = 0
    v22 := GenS4{F0: 1}
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
