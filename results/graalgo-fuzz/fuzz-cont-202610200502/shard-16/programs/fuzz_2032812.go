package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int16 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := []int{}
    var v2 int64 = -9223372036854775808
    v3 := make(map[uint32]int)
    var v4 uint = 1
    var v5 int32 = 0
    fmt.Println(v5)
    v2 = v2
    for fi6 := 0; fi6 < 20; fi6++ {
        delete(v3, 2147483647)
        v5 += (v5 - v5)
        _ = fi6
    }
    var v7 float64 = 0.0
    v3[4294967295] = 32767
    v8 := GenS1{F0: 2147483648, F1: v2, F2: 706}
    if true {
        var v9 int64 = 42
        fmt.Println(v7, v9, v9)
        if false {
            v9 = v2
            for fi10 := 0; fi10 < 10; fi10++ {
                v7 += (1.5 / v7)
                _ = fi10
            }
            fmt.Println(v9)
        }
        _ = v9
    }
    var v11 float64 = float64(3.14159)
    _ = len(v3)
    switch ("\n\t" + ("foo bar" + "\n\t")) {
    case "foo bar":
        if true {
            fmt.Println(v4)
        }
    case "hello":
        switch 128 {
        case -541:
            var v12 int8 = -1
            _ = v12
        default:
            fmt.Println(v2, v2)
        }
    case "\n\t":
        v2 = v2
    }
    v4 -= v4
    var v13 int64 = (v2 - v2)
    if (!((-(v5)) == (v5 ^ v5))) {
        switch 42 {
        case 128:
            for fi14 := 0; fi14 < 19; fi14++ {
                fmt.Println(v5)
                fmt.Println(v13)
                _ = fi14
            }
        case 127:
            v3[725] = -228
        case -689:
            for fi15 := 0; fi15 < 15; fi15++ {
                var v16 uint16 = 381
                switch ("" + "\n\t") {
                case "\n\t":
                    for fi17 := 0; fi17 < 13; fi17++ {
                        fmt.Println(v11)
                        _ = fi17
                    }
                case "foo bar":
                    if true {
                        _ = len(v1)
                        if (!false) {
                            var v18 uint16 = (v16 % (((v16 >> ((v4) & 63))) | 1))
                            _ = v18
                        }
                    }
                }
                _ = fi15
                _ = v16
            }
        default:
            v19 := GenS2{F0: (v13 & (^(v13)))}
            _ = v19
        }
    }
    fmt.Println(v11)
    v13 += 42
    if (!false) {
        var v20 uint16 = 298
        var v21 int16 = 1000
        v20 += v20
        _ = v20
        _ = v21
    }
    var v22 float64 = v11
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
