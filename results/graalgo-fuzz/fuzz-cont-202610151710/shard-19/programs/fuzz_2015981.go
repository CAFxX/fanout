package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint8 }

func main() {
    var v1 int64 = 0
    v2 := make(map[uint32]string)
    var v3 float64 = 1.5
    v4 := []int{-648, 95}
    switch ("hello" + "hello") {
    case "hello":
        if true {
            var v5 float64 = 1e10
            switch "" {
            case "foo bar":
                var v6 uint = 381
                _ = v6
            case "hello":
                v5 -= float64((1.5 / 0.1))
            default:
                switch ("hello" + "\n\t") {
                case "\n\t":
                    v5 -= v3
                case "":
                    v1 = 1
                default:
                    v7 := GenS1{F0: 2147483647, F1: 796}
                    _ = v7
                }
            }
            _ = v5
        }
    }
    v1 -= -1
    var v8 uint8 = 128
    v9 := GenS1{F0: -2147483648, F1: 765}
    for fi10 := 0; fi10 < 10; fi10++ {
        if true {
            var v11 uint32 = 65535
            _ = v11
        } else {
            v12 := GenS2{F0: 42, F1: 18446744073709551615, F2: 0}
            var v13 uint16 = 255
            _ = v12
            _ = v13
        }
        v14 := GenS1{F0: 42, F1: 42}
        for fi15 := 0; fi15 < 17; fi15++ {
            _ = len(v4)
            _ = fi15
        }
        _ = fi10
        _ = v14
    }
    if true {
        fmt.Println(v3)
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        v8 -= v8
        _ = fi16
    }
    _ = len(v2)
    for fi17 := 0; fi17 < 4; fi17++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v3 = v3
        v8 = 70
        _ = fi17
    }
    v3 -= -2.25
    v2[834] = ("a" + "\n\t")
    if ((!((27 == 65535) || (0 <= 0))) && ((((!(25 < -128)) || (268 < 42)) && (true && false)) && (v3 > 1.5))) {
        if (!true) {
            for fi18 := 0; fi18 < 18; fi18++ {
                for fi19 := 0; fi19 < 14; fi19++ {
                    if (false || ((v1 << ((2) & 63)) == (v1 * v1))) {
                        v8 = v8
                        var v20 float64 = v3
                        _ = v20
                    } else {
                        var v21 bool = false
                        _ = v21
                    }
                    v8 += 0
                    _ = fi19
                }
                v8 += (v8 | v8)
                v22 := GenS2{F0: 1, F1: 0, F2: v8}
                _ = fi18
                _ = v22
            }
        } else {
            var v23 uint = 497
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = len(v4)
            _ = v23
        }
    }
    switch -2147483648 {
    case 138:
        var v24 int8 = 54
        _ = v24
    }
    switch ("\n\t" + (("" + "\n\t") + "")) {
    case "\n\t":
        fmt.Println(v3, v8, v1)
    default:
        var v25 float32 = float32((((0.0 / v3) * v3) + v3))
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
