package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 0
    var v3 uint32 = 0
    var v4 int32 = 2147483647
    v5 := make(map[string]int)
    v6 := []uint32{v3, (v3 % (((^(v3))) | 1)), 65535}
    if (9223372036854775807 != -2147483648) {
        for fi7 := 0; fi7 < 9; fi7++ {
            v2 = (v2 * v2)
            v2 = v2
            switch ("a" + "\n\t") {
            case "\n\t":
                switch (^(fi7)) {
                case -1:
                    for fi8 := 0; fi8 < 4; fi8++ {
                        v9 := GenS1{F0: (fi7 | (fi7 | (fi7 << ((uint(v3)) & 63))))}
                        v2 = v2
                        _ = fi8
                        _ = v9
                    }
                case -2147483648:
                    for fi10 := 0; fi10 < 3; fi10++ {
                        v2 -= (v2 + 244)
                        v2 = 127
                        _ = fi10
                    }
                default:
                    var v11 uint8 = (v2 << ((42) & 63))
                    _ = v11
                }
            default:
                for fi12 := 0; fi12 < 18; fi12++ {
                    var v13 bool = (false && (!false))
                    v5[(("\n\t" + ("hello" + ("\n\t" + "hello"))) + "hello")] = (fi12 >> ((145) & 63))
                    _ = fi12
                    _ = v13
                }
            }
            _ = fi7
        }
    }
    _ = len(v5)
    fmt.Println(v2)
    var v14 string = "\n\t"
    if false {
        if (817 <= 42) {
            v3 -= v3
        } else {
            var v15 int64 = -367
            var v16 int = 127
            for fi17 := 0; fi17 < 9; fi17++ {
                var v18 string = v14
                fmt.Println(v15, v2)
                _ = fi17
                _ = v18
            }
            _ = v15
            _ = v16
        }
        v2 += (v2 >> ((0) & 63))
        var v19 uint64 = 45
        _ = v19
    } else {
        v14 = ((v14 + ((("foo bar" + v14) + ((("a" + v14) + (v14 + "\n\t")) + v14)) + v14)) + (v14 + v14))
    }
    v20 := GenS1{F0: 128}
    v21 := GenS1{F0: 691}
    var v22 uint32 = (v3 & v3)
    fmt.Println(v3, v3, v4)
    if true {
        delete(v1, v14)
        if true {
            _ = len(v1)
        } else {
            for fi23 := 0; fi23 < 10; fi23++ {
                fmt.Println(v14, v14, v2)
                v2 = (v2 + (v2 - 2))
                var v24 string = v14
                _ = fi23
                _ = v24
            }
        }
    }
    delete(v5, ((("hello" + "hello") + ("" + "")) + (v14 + "a")))
    if false {
        fmt.Println(v4, v3)
        if len(v6) > 0 {
            v6[0] = v22
        }
    } else {
        var v25 float32 = float32(0.1)
        _ = v25
    }
    fmt.Println(v14, v2, v2)
    _ = len(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
