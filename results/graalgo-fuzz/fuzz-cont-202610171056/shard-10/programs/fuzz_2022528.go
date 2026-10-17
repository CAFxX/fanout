package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint64 }

func main() {
    var v1 uint64 = 751
    var v2 int8 = 110
    v3 := make(map[string]uint32)
    var v4 uint32 = 630
    switch 490 {
    case 128:
        switch 4294967295 {
        case 32767:
            switch 4294967295 {
            case -468:
                fmt.Println(v1, v1)
            case -488:
                var v5 bool = (32767 >= 1)
                _ = v5
            }
        case 4294967295:
            v1 = 18446744073709551615
        case 9223372036854775807:
            v1 = v1
        }
    case -9223372036854775808:
        fmt.Println(v1)
    default:
        _ = len(v3)
    }
    if (!true) {
        for fi6 := 0; fi6 < 4; fi6++ {
            v7 := []int{}
            _ = fi6
            _ = v7
        }
        fmt.Println(v1)
    } else {
        v1 += (^(v1))
        var v8 uint32 = (v4 + uint32(v1))
        _ = v8
    }
    var v9 string = "hello"
    if (true || true) {
        if (!true) {
            fmt.Println(v4, v1)
            v4 = 704
            var v10 int64 = -765
            _ = v10
        }
    }
    if (255 > uint8(v4)) {
        for fi11 := 0; fi11 < 20; fi11++ {
            switch "" {
            case "foo bar":
                for fi12 := 0; fi12 < 4; fi12++ {
                    fmt.Println(fi12, v4, v1)
                    var v13 int = 0
                    var v14 uint64 = (v1 ^ (v1 * v1))
                    var v15 uint8 = 42
                    if true {
                        var v16 float32 = (1e10 + 1.5)
                        var v17 uint16 = 42
                        switch ((((v9 + v9) + ("hello" + (("foo bar" + v9) + (("" + "\n\t") + ("hello" + "a"))))) + ("a" + v9)) + v9) {
                        case "\n\t":
                            var v18 bool = (2 <= 18446744073709551615)
                            _ = v18
                        default:
                            fmt.Println(v13, fi11, v14)
                        }
                        _ = v16
                        _ = v17
                    }
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
            case "\n\t":
                var v19 int32 = -1
                _ = v19
            }
            switch (("foo bar" + v9) + (v9 + v9)) {
            case "a":
                v9 = ""
            default:
                var v20 int8 = (v2 / (((v2 & v2)) | 1))
                _ = v20
            }
            _ = fi11
        }
    }
    var v21 uint8 = 176
    v22 := GenS1{F0: v2, F1: 1, F2: 0}
    v23 := []int{9223372036854775807, 4294967295, -94, -1}
    v21 -= v21
    v4 += 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v9
}
