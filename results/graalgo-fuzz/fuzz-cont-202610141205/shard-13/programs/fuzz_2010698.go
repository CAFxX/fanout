package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint32 = 636
    v2 := make(map[int]int)
    v3 := []uint32{v1, v1, v1}
    var v4 int16 = 32767
    var v5 uint8 = 255
    v5 -= (v5 / ((v5) | 1))
    for fi6 := 0; fi6 < 14; fi6++ {
        v1 -= v1
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 float32 = 0.1
            switch (("a" + "a") + ("\n\t" + ("" + "hello"))) {
            case "foo bar":
                v9 := GenS1{F0: 2}
                _ = v9
            default:
                switch ("\n\t" + "a") {
                case "":
                    var v10 int = fi7
                    _ = v10
                default:
                    var v11 uint8 = (v5 ^ (v5 % (((v5 & v5)) | 1)))
                    _ = v11
                }
            }
            if (v1 == 0) {
                var v12 uint8 = v5
                var v13 uint64 = 2
                switch (fi6 * 128) {
                case 840:
                    var v14 uint = 2
                    _ = v14
                case 337:
                    v4 += (v4 % ((0) | 1))
                case -63:
                    v13 = 2
                }
                _ = v12
                _ = v13
            } else {
                if false {
                    v8 += v8
                    var v15 uint64 = 42
                    _ = v15
                }
                var v16 int64 = -1
                fmt.Println(v1, v4, v8)
                _ = v16
            }
            v5 = (v5 / ((255) | 1))
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    v3 = append(v3, (v1 << ((1) & 63)))
    if false {
        v4 = (v4 * (v4 * v4))
        v1 = 2147483647
    } else {
        var v17 int8 = -1
        _ = v17
    }
    fmt.Println(v5)
    var v18 int = 4294967295
    fmt.Println(v5)
    switch 127 {
    case 127:
        v18 -= v18
    case -2147483648:
        v18 = -933
    default:
        if false {
            var v19 int32 = 42
            _ = v19
        }
    }
    if true {
        var v20 string = ""
        _ = v20
    }
    fmt.Println(v4, v1, v4)
    v21 := GenS1{F0: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
