package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := []string{"hello", "hello", "foo bar", (("" + (((((("\n\t" + "foo bar") + "a") + "foo bar") + "hello") + "foo bar") + ("foo bar" + "foo bar"))) + "hello")}
    var v2 int = 32768
    v3 := make(map[string]int)
    v4 := make(map[string]int)
    var v5 float64 = 1e10
    fmt.Println(v5, v2)
    var v6 uint16 = 42
    var v7 uint16 = (v6 + (v6 | (v6 % (((v6 % ((v6) | 1))) | 1))))
    fmt.Println(v7, v7)
    v2 = (v2 * (v2 + int(v7)))
    v5 = 1e10
    var v8 float64 = v5
    switch "\n\t" {
    case "foo bar":
        if false {
            var v9 string = "foo bar"
            v10 := GenS1{F0: -18, F1: 0, F2: -96}
            _ = v10
            _ = v9
        } else {
            switch v2 {
            case 32767:
                for fi11 := 0; fi11 < 13; fi11++ {
                    if ((0 >= 4294967295) && false) {
                        v12 := GenS1{F0: 100, F1: 42, F2: -128}
                        fmt.Println(v2, fi11)
                        _ = v12
                    } else {
                        var v13 uint = 65535
                        fmt.Println(fi11, v6, v5)
                        var v14 float64 = (0.0 / (v5 - (v8 / 0.1)))
                        _ = v13
                        _ = v14
                    }
                    v5 = v8
                    v15 := GenS1{F0: 100, F1: 9223372036854775807, F2: int8(v6)}
                    _ = fi11
                    _ = v15
                }
            case 968:
                fmt.Println(v7, v6, v2)
            case 101:
                v16 := GenS1{F0: -95, F1: 9223372036854775807, F2: -128}
                _ = v16
            }
            switch (v2 % ((1) | 1)) {
            case 4294967295:
                for fi17 := 0; fi17 < 19; fi17++ {
                    v7 = v6
                    _ = fi17
                }
            default:
                v18 := GenS2{F0: 689}
                _ = v18
            }
            for fi19 := 0; fi19 < 7; fi19++ {
                v20 := GenS3{F0: 42}
                v5 += 1.5
                var v21 float64 = v8
                _ = fi19
                _ = v20
                _ = v21
            }
        }
    }
    if ((1000 > 42) && ((225 >= 0) && ((v2 << ((141) & 63)) < (v2 & (v2 | (v2 >> ((799) & 63))))))) {
        var v22 int32 = -1
        var v23 int16 = -637
        for fi24 := 0; fi24 < 6; fi24++ {
            if true {
                v25 := GenS2{F0: 2}
                delete(v3, "foo bar")
                for fi26 := 0; fi26 < 16; fi26++ {
                    v6 += (v7 | (v7 - (v6 & 255)))
                    _ = fi26
                }
                _ = v25
            } else {
                v4["a"] = fi24
                v4[("a" + "")] = (v2 + fi24)
            }
            _ = fi24
        }
        _ = v22
        _ = v23
    }
    v5 += (((v5 / (v5 + 1e10)) * 3.14159) / (v5 - float64((1.5 - (1e10 + 1.5)))))
    switch 454 {
    case 689:
        for fi27 := 0; fi27 < 10; fi27++ {
            var v28 float32 = (0.0 + 0.0)
            var v29 int8 = -1
            _ = fi27
            _ = v28
            _ = v29
        }
    default:
        var v30 uint64 = 1
        _ = v30
    }
    var v31 uint8 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
