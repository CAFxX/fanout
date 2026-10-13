package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 int32 }

func main() {
    var v1 uint8 = 176
    var v2 uint16 = 329
    v3 := []uint32{}
    v4 := make(map[uint32]bool)
    _ = len(v4)
    var v5 uint32 = 2
    switch "foo bar" {
    case "hello":
        _ = len(v4)
    }
    if ((18446744073709551615 != 18446744073709551615) && ((v5 / (((v5 ^ v5)) | 1)) != (v5 % (((^(v5))) | 1)))) {
        var v6 uint = 0
        _ = v6
    } else {
        v7 := GenS1{F0: -949}
        v8 := GenS2{F0: 841}
        if true {
            if (65535 < 907) {
                v5 = v5
            }
            var v9 uint64 = uint64(v2)
            _ = v9
        } else {
            v10 := GenS2{F0: (v5 >> ((42) & 63))}
            v1 = (v1 >> ((255) & 63))
            _ = v10
        }
        _ = v7
        _ = v8
    }
    if (!false) {
        v11 := GenS3{F0: -1, F1: 1}
        v2 -= (v2 << ((255) & 63))
        _ = v11
    }
    v3 = append(v3, uint32(v1))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v5, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
