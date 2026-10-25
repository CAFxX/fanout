package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint) int32 {
    if true {
        switch -639 {
        case -94:
            p0 = 42
        case 128:
            if (127 != -43) {
                var v7 string = ("" + "\n\t")
                v7 = "a"
                _ = v7
            }
        case 32767:
            for fi8 := 0; fi8 < 3; fi8++ {
                p0 = 255
                fmt.Println(p0, p0)
                _ = fi8
            }
        }
    }
    var v9 bool = false
    _ = p0
    _ = v9
    return 35
}

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 42
    v3 := make(map[string]int)
    v4 := []int{}
    var v5 uint64 = 262
    var v6 uint = 4294967295
    var v10 int64 = 680
    var v11 int8 = -128
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 int64 = 406
        if ("a" >= "hello") {
            fmt.Println(v2, v2, v11)
        }
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: int64(v2)}
    v6 += (v6 / (((v6 % (((v6 ^ (^(v6)))) | 1))) | 1))
    fmt.Println(v10, v6)
    switch "\n\t" {
    case "hello":
        _ = len(v4)
    case "":
        switch -1 {
        case -9223372036854775808:
            if false {
                v6 -= (v6 ^ v6)
            } else {
                switch ("hello" + (("foo bar" + ("foo bar" + ("hello" + ("a" + "")))) + "hello")) {
                case "":
                    for fi15 := 0; fi15 < 7; fi15++ {
                        v6 += 2
                        if len(v4) > 0 {
                            v4[0] = -301
                        }
                        var v16 float32 = float32(3.14159)
                        _ = fi15
                        _ = v16
                    }
                }
                fmt.Println(v5)
                v17 := GenS2{F0: v11}
                _ = v17
            }
        case -134:
            v18 := GenS2{F0: (v11 & v11)}
            _ = v18
        case 33:
            if false {
                switch "" {
                case "":
                    v10 = v10
                case "hello":
                    var v19 string = ""
                    _ = v19
                case "foo bar":
                    var v20 bool = false
                    _ = v20
                default:
                    v21 := GenS2{F0: v11}
                    _ = v21
                }
            }
        default:
            var v22 int = -172
            _ = v22
        }
    case "\n\t":
        var v23 int8 = (^(v11))
        _ = v23
    default:
        v5 += (v5 << (((v6 % ((v6) | 1))) & 63))
    }
    var v24 float32 = (0.1 * 0.0)
    for fi25 := 0; fi25 < 19; fi25++ {
        var v26 string = ("" + (("hello" + "a") + ("a" + "\n\t")))
        _ = fi25
        _ = v26
    }
    var v27 int32 = 924
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
