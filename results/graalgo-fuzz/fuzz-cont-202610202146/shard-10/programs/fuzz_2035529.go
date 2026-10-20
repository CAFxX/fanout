package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 uint64 }
func genFn0(p0 uint) uint64 {
    var v4 bool = false
    if v4 {
        for fi5 := 0; fi5 < 19; fi5++ {
            for fi6 := 0; fi6 < 14; fi6++ {
                v7 := []int{(fi5 - (fi5 * (fi5 + (fi5 << (((p0 | p0)) & 63))))), (^(fi6)), 32768}
                _ = fi6
                _ = v7
            }
            v8 := GenS1{F0: (!(-32768 == 0))}
            _ = fi5
            _ = v8
        }
    }
    v9 := []int{2147483647}
    p0 += p0
    var v10 float64 = 1e10
    _ = p0
    _ = v10
    _ = v4
    _ = v9
    return 0
}

func main() {
    var v1 uint32 = 481
    var v2 uint = 255
    v3 := make(map[string]bool)
    var v11 int32 = 0
    v11 -= (v11 ^ 42)
    if ((!false) || (-9223372036854775808 <= 1)) {
        var v12 int16 = 32767
        _ = v12
    } else {
        var v13 int8 = 103
        _ = v13
    }
    var v14 uint32 = (v1 % (((^(v1))) | 1))
    var v15 uint = (^(v2))
    v16 := GenS1{F0: false}
    if true {
        for fi17 := 0; fi17 < 15; fi17++ {
            switch "\n\t" {
            case "\n\t":
                v18 := v3[(("hello" + "a") + ("foo bar" + ("" + "\n\t")))]
                _ = v18
            case "foo bar":
                if false {
                    v15 = (v15 & (v2 + (^(v2))))
                    v19 := GenS1{F0: (0 <= 255)}
                    switch "" {
                    case "hello":
                        v20 := GenS2{F0: 127, F1: 18446744073709551615}
                        _ = v20
                    case "foo bar":
                        var v21 string = (((("" + ("\n\t" + "hello")) + "\n\t") + "hello") + "foo bar")
                        _ = v21
                    }
                    _ = v19
                }
            case "hello":
                var v22 bool = true
                _ = v22
            default:
                var v23 float32 = (3.14159 / 1.0)
                _ = v23
            }
            fmt.Println(v14, fi17, v11)
            _ = fi17
        }
        v2 = 727
        if (1 > 42) {
            v24 := []int{-997, -1}
            for fi25 := 0; fi25 < 7; fi25++ {
                switch "foo bar" {
                case "a":
                    var v26 uint8 = 127
                    _ = v26
                }
                var v27 int64 = 9223372036854775807
                var v28 uint32 = v14
                _ = fi25
                _ = v27
                _ = v28
            }
            switch "foo bar" {
            case "foo bar":
                v3["a"] = ((!(((v1 / ((v1) | 1)) <= 504) && false)) && true)
            case "hello":
                v29 := GenS1{F0: (v2 <= (v2 * uint(v1)))}
                _ = v29
            }
            _ = v24
        } else {
            v2 = 65535
        }
    }
    v30 := []int{-9223372036854775808}
    v2 = 255
    v31 := v3["foo bar"]
    var v32 int8 = 127
    for fi33 := 0; fi33 < 12; fi33++ {
        var v34 uint64 = genFn0((v2 * v2))
        fmt.Println(fi33)
        _ = fi33
        _ = v34
    }
    var v35 bool = (1 < -1)
    v14 = (v1 + (v1 * 272))
    v15 = (v15 << (((v15 | (^(v2)))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v35
}
