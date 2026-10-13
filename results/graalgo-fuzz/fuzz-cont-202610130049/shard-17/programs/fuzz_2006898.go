package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int64; F1 uint16; F2 uint }
func genFn0() int32 {
    var v7 uint32 = 65535
    if (((true || (2 < 128)) || ((1e10 + float32(((3.14159 - -2.25) * (1.5 / 1.0)))) <= (1.5 - 3.14159))) && ((!(false || (!(18446744073709551615 >= 288)))) && true)) {
        for fi8 := 0; fi8 < 2; fi8++ {
            v7 = (v7 & v7)
            _ = fi8
        }
    }
    _ = v7
    return 42
}

func main() {
    var v1 float32 = -2.25
    var v2 int = -120
    v3 := []int{}
    var v4 uint16 = 2
    var v5 uint8 = 239
    v6 := make(map[string]uint32)
    var v9 int8 = 100
    for fi10 := 0; fi10 < 16; fi10++ {
        v6[("\n\t" + "hello")] = 730
        var v11 uint32 = 391
        if false {
            for fi12 := 0; fi12 < 8; fi12++ {
                switch "a" {
                case "a":
                    v13 := GenS1{F0: 0, F1: "hello"}
                    _ = v13
                default:
                    v1 -= 1e10
                }
                var v14 int32 = 0
                v2 += v2
                _ = fi12
                _ = v14
            }
        }
        _ = fi10
        _ = v11
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        v1 += v1
        _ = fi15
    }
    v2 = -9223372036854775808
    var v16 string = (("foo bar" + "\n\t") + "a")
    var v17 float32 = 1e10
    var v18 uint8 = 128
    var v19 uint64 = 18446744073709551615
    for fi20 := 0; fi20 < 18; fi20++ {
        var v21 uint8 = v18
        _ = fi20
        _ = v21
    }
    if ("hello" <= v16) {
        v22 := GenS1{F0: 290, F1: v16}
        _ = v22
    } else {
        if false {
            fmt.Println(v2, v9, v4)
            v5 = v18
            switch v2 {
            case 4294967295:
                for fi23 := 0; fi23 < 16; fi23++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v24 int64 = -1
                    _ = fi23
                    _ = v24
                }
            }
        } else {
            v25 := GenS2{F0: 42, F1: (v4 >> ((1000) & 63)), F2: 582}
            _ = v25
        }
    }
    fmt.Println(v2)
    if len(v3) > 0 {
        v3[0] = v2
    }
    var v26 int32 = 211
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
