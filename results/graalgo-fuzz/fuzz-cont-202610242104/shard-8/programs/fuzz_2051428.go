package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint32; F1 uint; F2 uint }
type GenS4 struct { F0 uint16; F1 uint16; F2 string }
func genFn0(p0 bool) int16 {
    for fi6 := 0; fi6 < 15; fi6++ {
        switch (^(fi6)) {
        case 464:
            v7 := []int{fi6, fi6}
            _ = v7
        case 128:
            v8 := []int{fi6, 185, (fi6 + fi6), fi6}
            _ = v8
        case 32768:
            p0 = p0
        default:
            var v9 uint8 = 255
            _ = v9
        }
        if (!false) {
            var v10 int32 = 2147483647
            _ = v10
        }
        _ = fi6
    }
    v11 := GenS1{F0: 100, F1: 2, F2: 32768}
    _ = p0
    _ = v11
    return -398
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 4294967295
    var v3 uint64 = 1
    var v4 uint32 = v2
    var v5 int32 = -994
    var v12 uint = 0
    v13 := GenS2{F0: v5, F1: "a"}
    v14 := v1[("hello" + ("a" + "\n\t"))]
    var v15 int64 = 0
    for fi16 := 0; fi16 < 12; fi16++ {
        if true {
            for fi17 := 0; fi17 < 16; fi17++ {
                var v18 int8 = 118
                v19 := GenS2{F0: v5, F1: ("foo bar" + ("foo bar" + "\n\t"))}
                if (2 == 489) {
                    v18 += v18
                }
                if v14 {
                    fmt.Println(v5)
                    v20 := GenS3{F0: (v4 | (v2 / (((v4 + v4)) | 1))), F1: v12, F2: (v12 & (v12 - v12))}
                    switch 32768 {
                    case -214:
                        v5 = (v5 * v5)
                    case -111:
                        var v21 uint = (v12 * v12)
                        _ = v21
                    case -2147483648:
                        v22 := GenS4{F0: 2, F1: 736, F2: "\n\t"}
                        _ = v22
                    }
                    _ = v20
                }
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        var v23 int32 = v5
        _ = fi16
        _ = v23
    }
    delete(v1, ("a" + "foo bar"))
    fmt.Println(v2, v2, v3)
    v14 = (false && ((false && true) && v14))
    v24 := []int{-66, 128, 1, -804}
    if len(v24) > 0 {
        v24[0] = int(v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
