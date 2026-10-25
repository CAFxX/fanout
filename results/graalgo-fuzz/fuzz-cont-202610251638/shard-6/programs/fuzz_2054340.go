package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int16; F1 int32; F2 uint8 }
func genFn0() uint64 {
    if (-128 < 114) {
        var v5 uint32 = 65535
        var v6 bool = false
        switch 4294967295 {
        case 161:
            var v7 int16 = -1
            _ = v7
        case 42:
            switch 42 {
            case 2147483647:
                fmt.Println(v5, v5)
            default:
                if ((!v6) && true) {
                    fmt.Println(v6)
                    v8 := []int{}
                    _ = v8
                }
            }
        case -810:
            for fi9 := 0; fi9 < 13; fi9++ {
                v10 := GenS1{F0: 42}
                _ = fi9
                _ = v10
            }
        }
        _ = v5
        _ = v6
    }
    var v11 int16 = 1000
    if (!true) {
        v12 := []int{-2147483648, 2147483647, 9223372036854775807, 0}
        _ = v12
    } else {
        for fi13 := 0; fi13 < 1; fi13++ {
            switch ((("" + "") + "hello") + "\n\t") {
            case "a":
                var v14 float64 = (0.1 * 3.14159)
                _ = v14
            case "":
                v11 += v11
            }
            var v15 uint16 = 255
            _ = fi13
            _ = v15
        }
        var v16 int16 = -483
        _ = v16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        v11 = 1000
        _ = fi17
    }
    if false {
        if ((1.5 * 1.5) > 1.5) {
            var v18 int64 = -9223372036854775808
            for fi19 := 0; fi19 < 18; fi19++ {
                for fi20 := 0; fi20 < 10; fi20++ {
                    var v21 int64 = 42
                    v21 = v18
                    _ = fi20
                    _ = v21
                }
                fmt.Println(fi19, fi19)
                _ = fi19
            }
            _ = v18
        }
        fmt.Println(v11)
        v11 += v11
    } else {
        var v22 int = 2147483647
        _ = v22
    }
    _ = v11
    return 51
}

func main() {
    v1 := []uint32{2147483647, 42, 2147483647, 1}
    var v2 uint = 702
    var v3 uint32 = 2147483647
    var v4 int64 = -1
    v2 = v2
    v1 = append(v1, v3)
    v3 -= (v3 / (((v3 % ((v3) | 1))) | 1))
    var v23 uint = v2
    v24 := GenS2{F0: 234, F1: 1, F2: v3}
    var v25 uint16 = 1
    v26 := GenS3{F0: int16(v3), F1: -1, F2: 42}
    switch "a" {
    case "foo bar":
        var v27 uint = (v23 - (v23 + (v23 & 190)))
        _ = v27
    case "":
        if true {
            v28 := GenS1{F0: v25}
            _ = v28
        }
    case "a":
        if (0 != 42) {
            for fi29 := 0; fi29 < 11; fi29++ {
                switch (("\n\t" + "\n\t") + "foo bar") {
                case "":
                    var v30 float64 = ((3.14159 / 1.5) + 0.0)
                    _ = v30
                default:
                    var v31 int64 = v4
                    _ = v31
                }
                _ = fi29
            }
            var v32 uint64 = genFn0()
            _ = v32
        }
    default:
        fmt.Println(v4, v2)
    }
    fmt.Println(v2)
    switch (("" + ((("\n\t" + "a") + "\n\t") + "a")) + ("" + ("" + ("" + "")))) {
    case "\n\t":
        if false {
            fmt.Println(v3, v2)
            v1 = append(v1, v3)
            var v33 float64 = (1e10 / 1e10)
            _ = v33
        } else {
            if (true && (!((^(v4)) >= (v4 + (v4 * (v4 + v4)))))) {
                fmt.Println(v25, v2)
            } else {
                var v34 float32 = 3.14159
                v2 -= 810
                if (!false) {
                    switch 637 {
                    case 42:
                        v4 -= v4
                    }
                    switch "hello" {
                    case "\n\t":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    var v35 int64 = 9223372036854775807
                    _ = v35
                }
                _ = v34
            }
            v4 = (v4 | -1)
        }
    default:
        var v36 int8 = -1
        _ = v36
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
