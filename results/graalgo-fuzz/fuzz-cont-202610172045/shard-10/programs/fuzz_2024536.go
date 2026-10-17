package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 uint32; F1 int64; F2 uint8 }
func genFn0() int8 {
    var v6 string = ("a" + "hello")
    if true {
        var v7 int16 = 1
        v6 = (v6 + v6)
        v7 = 1
        _ = v7
    } else {
        var v8 uint32 = 2147483647
        v9 := GenS1{F0: 42}
        _ = v8
        _ = v9
    }
    _ = v6
    return 42
}

func main() {
    var v1 uint = 399
    var v2 int32 = 477
    var v3 uint = v1
    v4 := []string{("" + "a"), "\n\t", "", ("\n\t" + ("foo bar" + "a"))}
    var v5 uint64 = 24
    var v10 bool = (false || ((float32((-2.25 + 1e10)) - 0.1) < 1.5))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (!false) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(v10, v5, v3)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v4 = append(v4, ("a" + "a"))
    for fi11 := 0; fi11 < 11; fi11++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if false {
            v12 := GenS1{F0: 972}
            if false {
                for fi13 := 0; fi13 < 7; fi13++ {
                    v14 := GenS2{F0: 2, F1: v5}
                    _ = fi13
                    _ = v14
                }
            }
            _ = v12
        }
        _ = fi11
    }
    v5 = (v5 << ((v3) & 63))
    if (4294967295 > 582) {
        fmt.Println(v3, v10, v10)
        v15 := GenS3{F0: 206, F1: -1, F2: 255}
        _ = v15
    } else {
        var v16 uint8 = 2
        fmt.Println(v5, v16)
        switch "" {
        case "":
            v2 -= v2
        }
        _ = v16
    }
    if v10 {
        for fi17 := 0; fi17 < 13; fi17++ {
            switch "a" {
            case "a":
                fmt.Println(v10)
            default:
                var v18 int64 = -1
                _ = v18
            }
            v5 = (v5 & 925)
            _ = fi17
        }
        v5 -= (v5 * uint64(v3))
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        v1 += (v1 + 18446744073709551615)
        var v20 float64 = 1.5
        _ = fi19
        _ = v20
    }
    var v21 float64 = -2.25
    v5 = v5
    var v22 int = 4294967295
    _ = len(v4)
    var v23 uint32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
