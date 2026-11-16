package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 int8; F2 uint16 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS5 struct { F0 uint16; F1 int16; F2 int }
type GenS6 struct { F0 uint }
func genFn0() int32 {
    var v4 int16 = 42
    v4 = (v4 / (((v4 * (v4 + (v4 << ((4294967295) & 63))))) | 1))
    _ = v4
    return 2147483647
}

func main() {
    var v1 int32 = -759
    var v2 int64 = -419
    var v3 uint = 2
    switch ((("\n\t" + "a") + "hello") + "hello") {
    case "\n\t":
        v5 := GenS1{F0: 1}
        _ = v5
    case "hello":
        var v6 uint8 = 255
        _ = v6
    case "a":
        if (-22 == -27) {
            if (!(1 <= int16(v1))) {
                var v7 string = "a"
                v1 = v1
                _ = v7
            } else {
                v2 = (^(v2))
                switch 9223372036854775807 {
                case 2147483647:
                    for fi8 := 0; fi8 < 3; fi8++ {
                        v9 := GenS2{F0: ((((("foo bar" + ((("hello" + ("\n\t" + "hello")) + ("foo bar" + "hello")) + "hello")) + ("\n\t" + "hello")) + "") + ("\n\t" + "foo bar")) + "hello"), F1: -128, F2: 65535}
                        var v10 string = "\n\t"
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                case 1:
                    var v11 float32 = 0.0
                    _ = v11
                }
                var v12 uint = (v3 & (v3 | 568))
                _ = v12
            }
            v13 := GenS3{F0: 2, F1: (v2 & v2)}
            if ((1e10 + float64(-2.25)) >= 0.1) {
                for fi14 := 0; fi14 < 5; fi14++ {
                    switch "\n\t" {
                    case "foo bar":
                        v15 := GenS2{F0: "", F1: -1, F2: 848}
                        _ = v15
                    }
                    _ = fi14
                }
                v16 := GenS4{F0: 32767, F1: 127, F2: 749}
                _ = v16
            } else {
                var v17 float64 = 1e10
                _ = v17
            }
            _ = v13
        } else {
            fmt.Println(v3, v1)
            var v18 float32 = 1.5
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        if false {
            v3 += 65535
            var v20 int32 = (-(v1))
            _ = v20
        }
        if false {
            for fi21 := 0; fi21 < 7; fi21++ {
                var v22 int = (fi21 ^ fi19)
                _ = fi21
                _ = v22
            }
            for fi23 := 0; fi23 < 1; fi23++ {
                v3 -= v3
                _ = fi23
            }
            v24 := GenS5{F0: 944, F1: 754, F2: fi19}
            _ = v24
        } else {
            switch "hello" {
            case "\n\t":
                v25 := []int{(fi19 + -1), (fi19 + 938), fi19}
                _ = v25
            }
            switch 1 {
            case 127:
                var v26 float64 = 1e10
                _ = v26
            }
            v27 := []int{(fi19 % ((0) | 1)), (fi19 * fi19)}
            _ = v27
        }
        for fi28 := 0; fi28 < 12; fi28++ {
            fmt.Println(v3, fi28, fi19)
            v2 = (v2 * v2)
            _ = fi28
        }
        _ = fi19
    }
    v29 := []int{9223372036854775807, 0, 256, -783}
    v30 := GenS6{F0: (v3 | (^(v3)))}
    v1 -= (v1 - v1)
    if (!(!false)) {
        if len(v29) > 0 {
            v29[0] = -630
        }
    } else {
        v3 = (v3 | 841)
    }
    if true {
        v31 := GenS3{F0: 2, F1: v2}
        var v32 bool = true
        _ = v31
        _ = v32
    } else {
        var v33 int64 = v2
        v1 -= (-(v1))
        var v34 string = "foo bar"
        _ = v33
        _ = v34
    }
    if true {
        if len(v29) >= 2 {
            v29 = v29[1:len(v29)]
        }
    }
    v3 += v3
    if false {
        var v35 bool = false
        fmt.Println(v2)
        _ = v35
    }
    fmt.Println(v2, v3)
    if len(v29) >= 2 {
        v29 = v29[1:len(v29)]
    }
    v36 := GenS4{F0: 1, F1: 247, F2: uint16(v3)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v36
}
