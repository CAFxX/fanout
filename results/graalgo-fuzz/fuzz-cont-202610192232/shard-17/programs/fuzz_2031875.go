package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int; F2 int16 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 uint16 = 154
    var v2 int32 = 305
    v3 := make(map[string]uint32)
    var v4 int32 = v2
    var v5 uint = 255
    for fi6 := 0; fi6 < 2; fi6++ {
        if false {
            var v7 int = (fi6 ^ -1)
            _ = v7
        }
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 int64 = 42
            v1 -= 663
            _ = fi8
            _ = v9
        }
        var v10 int16 = 1000
        _ = fi6
        _ = v10
    }
    var v11 uint = 1
    for fi12 := 0; fi12 < 16; fi12++ {
        delete(v3, (("foo bar" + "foo bar") + "foo bar"))
        _ = fi12
    }
    v4 = v4
    fmt.Println(v5, v5)
    fmt.Println(v4)
    if (!true) {
        v11 += v11
        if (!((v4 >> ((uint(v1)) & 63)) < v2)) {
            var v13 int8 = -128
            if true {
                var v14 uint32 = 2147483648
                _ = v14
            }
            v15 := GenS1{F0: 212}
            _ = v13
            _ = v15
        } else {
            var v16 uint32 = 2
            for fi17 := 0; fi17 < 4; fi17++ {
                switch "" {
                case "foo bar":
                    v1 = v1
                default:
                    v18 := []int{fi17, (fi17 << (((v11 / ((uint(v16)) | 1))) & 63)), (fi17 + fi17), (fi17 * (fi17 | 127))}
                    _ = v18
                }
                var v19 uint64 = uint64(v11)
                _ = fi17
                _ = v19
            }
            v3["hello"] = 2147483648
            _ = v16
        }
        v20 := GenS1{F0: 1}
        _ = v20
    }
    v21 := []int{1, 128}
    var v22 uint32 = 42
    if (true || (!true)) {
        _ = len(v21)
    }
    var v23 uint = (v5 * (v5 / (((^(v11))) | 1)))
    v24 := GenS1{F0: 752}
    for fi25 := 0; fi25 < 3; fi25++ {
        v26 := GenS2{F0: -1, F1: fi25, F2: 32767}
        v11 = uint(v22)
        _ = fi25
        _ = v26
    }
    if len(v21) > 0 {
        v21[0] = 128
    }
    for fi27 := 0; fi27 < 9; fi27++ {
        v28 := GenS3{F0: uint64(v22)}
        _ = fi27
        _ = v28
    }
    v29 := v3["\n\t"]
    fmt.Println(v29, v22, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
