package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 int }

func main() {
    var v1 float32 = 1.5
    var v2 uint16 = 65535
    var v3 int16 = 32767
    var v4 uint16 = (v2 - v2)
    var v5 int8 = 127
    if ("" >= (("a" + "foo bar") + "foo bar")) {
        v5 += (v5 >> ((445) & 63))
        v5 += (v5 / ((v5) | 1))
    } else {
        if true {
            v4 = v2
        }
        v2 = v4
        fmt.Println(v3)
    }
    v6 := []int{32767, 128, -9223372036854775808, 4294967295}
    var v7 int8 = v5
    for fi8 := 0; fi8 < 7; fi8++ {
        if false {
            v9 := GenS1{F0: 1, F1: fi8}
            var v10 string = ("hello" + "\n\t")
            _ = v10
            _ = v9
        }
        for fi11 := 0; fi11 < 14; fi11++ {
            if (true || (false && (0 <= 230))) {
                var v12 float32 = v1
                var v13 int = 32768
                if ((v4 ^ v2) != v4) {
                    v7 += v7
                    if true {
                        var v14 uint8 = 31
                        if false {
                            var v15 string = "foo bar"
                            fmt.Println(v1, v13)
                            var v16 int64 = -433
                            _ = v15
                            _ = v16
                        }
                        v2 += v2
                        var v17 int = -9223372036854775808
                        _ = v14
                        _ = v17
                    } else {
                        var v18 float64 = (0.0 / 1e10)
                        v19 := GenS2{F0: (v13 / ((fi8) | 1))}
                        _ = v18
                        _ = v19
                    }
                    v13 += -2147483648
                }
                var v20 uint = 244
                v21 := GenS1{F0: 11, F1: (-(fi11))}
                _ = v12
                _ = v13
                _ = v20
                _ = v21
            } else {
                v2 = (v4 / ((378) | 1))
            }
            _ = fi11
        }
        fmt.Println(v2, fi8)
        _ = fi8
    }
    v22 := GenS2{F0: -281}
    var v23 float32 = (v1 + v1)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v23 = ((v23 * v1) - v23)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
