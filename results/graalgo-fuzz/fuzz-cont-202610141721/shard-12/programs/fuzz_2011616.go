package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 int8; F1 uint }
func genFn0() bool {
    var v6 uint16 = 0
    fmt.Println(v6)
    v7 := []int{-651, 127}
    var v8 uint16 = 32767
    _ = v6
    _ = v7
    _ = v8
    return (127 != 0)
}

func main() {
    v1 := []string{"a"}
    var v2 uint = 255
    v3 := []string{"foo bar", ("" + "")}
    var v4 int8 = 42
    var v5 int64 = -9223372036854775808
    var v9 int16 = 42
    v5 -= (v5 | v5)
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    _ = len(v3)
    var v10 bool = (771 != 764)
    v5 = (-(v5))
    v11 := GenS1{F0: 2147483647}
    if true {
        if false {
            v12 := GenS1{F0: 2147483647}
            for fi13 := 0; fi13 < 14; fi13++ {
                if ((128 != 128) && v10) {
                    var v14 bool = false
                    for fi15 := 0; fi15 < 19; fi15++ {
                        switch (^(fi13)) {
                        case 472:
                            fmt.Println(v14)
                        default:
                            v9 -= -1
                        }
                        _ = fi15
                    }
                    switch (-(fi13)) {
                    case -192:
                        var v16 uint32 = uint32(v2)
                        _ = v16
                    }
                    _ = v14
                }
                _ = fi13
            }
            _ = v12
        }
        for fi17 := 0; fi17 < 5; fi17++ {
            fmt.Println(fi17, v9, v2)
            _ = fi17
        }
    } else {
        fmt.Println(v5)
        v18 := GenS1{F0: 2}
        var v19 uint8 = 84
        _ = v18
        _ = v19
    }
    if v10 {
        _ = len(v1)
        var v20 bool = v10
        _ = v20
    }
    v4 = (v4 % ((v4) | 1))
    for fi21 := 0; fi21 < 2; fi21++ {
        if ((fi21 | (fi21 << (((^(v2))) & 63))) < (fi21 - (^(fi21)))) {
            switch (("hello" + "hello") + "hello") {
            case "hello":
                fmt.Println(v10, v4, v5)
            case "\n\t":
                for fi22 := 0; fi22 < 8; fi22++ {
                    v2 = v2
                    v23 := GenS1{F0: 0}
                    _ = fi22
                    _ = v23
                }
            default:
                if (-414 > (fi21 - (^(fi21)))) {
                    var v24 uint8 = 255
                    v25 := GenS2{F0: 32767, F1: (-(v9)), F2: (v24 ^ (v24 + 122))}
                    _ = v24
                    _ = v25
                } else {
                    v9 = int16(fi21)
                    var v26 string = "\n\t"
                    _ = v26
                }
            }
            fmt.Println(fi21)
            if (v10 || (1 != 44)) {
                var v27 uint = 2
                v10 = (v5 <= -310)
                v28 := GenS3{F0: v10, F1: -23}
                _ = v27
                _ = v28
            } else {
                v9 = v9
            }
        }
        var v29 float64 = float64((0.1 + 3.14159))
        _ = len(v1)
        _ = fi21
        _ = v29
    }
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    _ = len(v3)
    v30 := GenS4{F0: 1, F1: 519}
    var v31 uint64 = 2
    v4 = v4
    v32 := GenS5{F0: (-(v4)), F1: (v2 + (v2 | (v2 + 255)))}
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v9
}
