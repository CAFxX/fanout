package main

import (
    "fmt"
)


func main() {
    var v1 string = "a"
    v2 := make(map[string]uint32)
    var v3 uint16 = 325
    v4 := []uint32{59}
    for fi5 := 0; fi5 < 5; fi5++ {
        _ = len(v2)
        _ = fi5
    }
    var v6 uint64 = 614
    if (float64(float32(1e10)) != (-2.25 / 1e10)) {
        if (-738 == -96) {
            var v7 float64 = (float64(float32(0.1)) / (3.14159 / 1.0))
            _ = v7
        }
        fmt.Println(v3, v6)
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 uint64 = v6
            if false {
                if (1 == -1) {
                    if (1 >= 128) {
                        if true {
                            v2["a"] = 65535
                            fmt.Println(v9, v9, v3)
                        }
                    }
                }
                v10 := v2[("foo bar" + (v1 + v1))]
                var v11 uint32 = v10
                _ = v10
                _ = v11
            } else {
                var v12 uint8 = 255
                if false {
                    var v13 uint64 = v9
                    v14 := v2["a"]
                    var v15 int = (-(fi8))
                    _ = v13
                    _ = v14
                    _ = v15
                }
                switch "foo bar" {
                case "hello":
                    switch 886 {
                    case 128:
                        fmt.Println(v12)
                    case -175:
                        switch fi8 {
                        case -1:
                            v12 += (v12 % ((128) | 1))
                        case 0:
                            if len(v4) > 0 {
                                v4[0] = 2147483647
                            }
                        case -897:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        default:
                            var v16 int = 0
                            _ = v16
                        }
                    }
                case "foo bar":
                    v6 -= v9
                default:
                    v1 = v1
                }
                _ = v12
            }
            _ = fi8
            _ = v9
        }
    }
    v6 += v6
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 bool = (-1 == -9223372036854775808)
        v18 = (v6 > 2)
        fmt.Println(v18)
        _ = fi17
        _ = v18
    }
    fmt.Println(v6, v6, v3)
    var v19 uint = 18446744073709551615
    switch ("a" + "") {
    case "\n\t":
        var v20 string = (((v1 + "hello") + ((v1 + "a") + (v1 + ((("" + "") + (v1 + (v1 + v1))) + "hello")))) + (("a" + "") + v1))
        _ = v20
    case "a":
        switch v1 {
        case "hello":
            var v21 float32 = 3.14159
            _ = v21
        default:
            _ = len(v4)
        }
    }
    v3 = v3
    for fi22 := 0; fi22 < 14; fi22++ {
        fmt.Println(v6, v3)
        for fi23 := 0; fi23 < 16; fi23++ {
            var v24 int64 = -1
            v19 -= v19
            _ = fi23
            _ = v24
        }
        var v25 int8 = int8(fi22)
        _ = fi22
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
