package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint; F2 string }
type GenS3 struct { F0 bool }

func main() {
    v1 := []int{-262, -994, 32768}
    var v2 int64 = 334
    v3 := make(map[int]int)
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]int)
    var v6 uint64 = 42
    var v7 string = "\n\t"
    if len(v1) > 0 {
        v1[0] = 553
    }
    if (true && true) {
        var v8 int16 = 1
        var v9 string = "\n\t"
        for fi10 := 0; fi10 < 19; fi10++ {
            v11 := GenS1{F0: 42}
            _ = fi10
            _ = v11
        }
        _ = v8
        _ = v9
    }
    if (0 != 2147483648) {
        var v12 string = (v7 + (v7 + ((("foo bar" + v7) + v7) + v7)))
        v1 = append(v1, -2147483648)
        fmt.Println(v7, v6, v7)
        _ = v12
    } else {
        fmt.Println(v7, v2)
        v13 := GenS1{F0: 1}
        var v14 int16 = -1
        _ = v13
        _ = v14
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v15 int64 = (v2 / ((-890) | 1))
    v16 := v3[383]
    fmt.Println(v7)
    v2 = v15
    switch 9223372036854775807 {
    case 0:
        if true {
            v2 = v15
            for fi17 := 0; fi17 < 6; fi17++ {
                v18 := GenS1{F0: -2147483648}
                _ = fi17
                _ = v18
            }
        }
    case 127:
        v16 = v16
    }
    _ = len(v1)
    _ = len(v1)
    for fi19 := 0; fi19 < 9; fi19++ {
        fmt.Println(fi19)
        for fi20 := 0; fi20 < 20; fi20++ {
            v21 := GenS1{F0: 42}
            if false {
                _ = len(v1)
                var v22 int32 = 0
                if ((!false) || (816 > 1000)) {
                    var v23 bool = true
                    var v24 float32 = (3.14159 * float32((1e10 - 0.1)))
                    v7 = v7
                    _ = v23
                    _ = v24
                }
                _ = v22
            } else {
                if (!(("" + "\n\t") <= v7)) {
                    var v25 uint = 1
                    v26 := GenS1{F0: 0}
                    _ = v25
                    _ = v26
                }
            }
            fmt.Println(v15)
            switch v7 {
            case "":
                delete(v4, 2147483647)
            case "foo bar":
                switch v7 {
                case "a":
                    v27 := GenS2{F0: 32767, F1: 2, F2: ""}
                    _ = v27
                case "\n\t":
                    v6 += v6
                default:
                    v28 := GenS3{F0: (!true)}
                    _ = v28
                }
                switch ("foo bar" + "hello") {
                case "":
                    var v29 string = "foo bar"
                    _ = v29
                case "foo bar":
                    fmt.Println(v15, fi19)
                case "a":
                    var v30 float64 = 0.0
                    _ = v30
                default:
                    fmt.Println(fi19)
                }
                switch 614 {
                case -2147483648:
                    v6 = (v6 * v6)
                }
                _ = len(v1)
            }
            _ = fi20
            _ = v21
        }
        var v31 float64 = ((0.0 * (1.5 + 0.1)) - (-2.25 + 3.14159))
        _ = fi19
        _ = v31
    }
    if (uint16(v6) < 255) {
        for fi32 := 0; fi32 < 5; fi32++ {
            v1 = append(v1, 32767)
            v2 -= (v15 & (v2 | (v2 & (v15 / ((-927) | 1)))))
            for fi33 := 0; fi33 < 3; fi33++ {
                v34 := GenS1{F0: -1}
                fmt.Println(v15, v7, v6)
                if false {
                    v15 -= (v15 << ((4294967295) & 63))
                }
                _ = fi33
                _ = v34
            }
            _ = fi32
        }
    } else {
        v2 -= (v15 * v2)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
