package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 int16 }

func main() {
    v1 := []uint32{}
    v2 := []int{1, 200, -1}
    v3 := []string{}
    var v4 uint8 = 81
    v1 = append(v1, 65535)
    if len(v2) > 0 {
        v2[0] = 973
    }
    v5 := GenS1{F0: 42, F1: 821, F2: 2147483648}
    var v6 uint = 255
    v1 = append(v1, uint32(v4))
    v4 -= v4
    v1 = append(v1, 596)
    switch 32767 {
    case 32768:
        v7 := GenS2{F0: -2147483648}
        _ = v7
    }
    switch -962 {
    case 9223372036854775807:
        for fi8 := 0; fi8 < 15; fi8++ {
            if true {
                v6 -= 1
            } else {
                for fi9 := 0; fi9 < 17; fi9++ {
                    if (false || false) {
                        var v10 string = ("a" + "hello")
                        _ = v10
                    }
                    if (3.14159 != 0.0) {
                        _ = len(v1)
                        v6 = v6
                    }
                    _ = fi9
                }
                var v11 uint64 = 4294967295
                _ = v11
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v12 int8 = -1
            _ = fi8
            _ = v12
        }
    case 128:
        var v13 int = 4294967295
        _ = v13
    case 752:
        v14 := GenS2{F0: 0}
        _ = v14
    default:
        fmt.Println(v4)
    }
    v15 := GenS2{F0: 45}
    for fi16 := 0; fi16 < 8; fi16++ {
        v4 = uint8(v6)
        v4 = (v4 << (((v6 % (((v6 | 1)) | 1))) & 63))
        v17 := GenS3{F0: ("\n\t" + "a"), F1: 711}
        _ = fi16
        _ = v17
    }
    var v18 uint8 = (v4 - 119)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
