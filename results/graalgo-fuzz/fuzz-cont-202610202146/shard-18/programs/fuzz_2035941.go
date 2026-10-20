package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int64; F1 int8 }
func genFn0() int32 {
    if (18446744073709551615 != 0) {
        var v6 int8 = 0
        v6 += (v6 / ((v6) | 1))
        fmt.Println(v6)
        _ = v6
    }
    var v7 string = (("a" + "") + ("" + "foo bar"))
    for fi8 := 0; fi8 < 13; fi8++ {
        switch "foo bar" {
        case "foo bar":
            var v9 bool = (((42 != -32768) && ((0 == 65535) && ((fi8 | (fi8 & (fi8 << ((980) & 63)))) >= (fi8 + 128)))) && false)
            _ = v9
        case "hello":
            var v10 uint64 = 438
            _ = v10
        }
        v11 := []int{fi8, (fi8 - 127)}
        _ = fi8
        _ = v11
    }
    v7 = ""
    _ = v7
    return -190
}

func main() {
    v1 := []string{"", ("" + "hello"), "\n\t", ("\n\t" + "\n\t")}
    v2 := []string{}
    var v3 int = 127
    v4 := make(map[string]bool)
    v5 := make(map[uint32]string)
    for fi12 := 0; fi12 < 6; fi12++ {
        v3 = v3
        switch "foo bar" {
        case "hello":
            if false {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if false {
                    v3 = (fi12 % ((1) | 1))
                } else {
                    v3 -= v3
                }
                for fi13 := 0; fi13 < 12; fi13++ {
                    v3 += -170
                    _ = fi13
                }
            }
        }
        switch ("hello" + "a") {
        case "hello":
            v3 += 127
        default:
            v3 = fi12
        }
        _ = fi12
    }
    if ((1 != -32768) && (-2.25 <= (0.0 / (1.5 + 1e10)))) {
        _ = len(v4)
        fmt.Println(v3)
        switch 32767 {
        case 1:
            fmt.Println(v3)
        case -592:
            fmt.Println(v3)
        case -261:
            if (1e10 != 1e10) {
                v14 := GenS1{F0: true, F1: genFn0()}
                v1 = append(v1, "hello")
                _ = len(v5)
                _ = v14
            }
        default:
            for fi15 := 0; fi15 < 3; fi15++ {
                var v16 uint8 = 1
                v17 := GenS1{F0: (18446744073709551615 <= 2), F1: -2147483648}
                v3 = (v3 & (fi15 * fi15))
                _ = fi15
                _ = v16
                _ = v17
            }
        }
    }
    switch "" {
    case "hello":
        if ((float64((float32(float64(1.5)) - float32(1e10))) * 1.5) <= -2.25) {
            fmt.Println(v3)
        }
    case "\n\t":
        v3 = (v3 ^ v3)
    }
    var v18 float64 = 3.14159
    v18 += v18
    v5[0] = "hello"
    for fi19 := 0; fi19 < 13; fi19++ {
        for fi20 := 0; fi20 < 16; fi20++ {
            if false {
                var v21 string = "a"
                v22 := GenS1{F0: (((v21 + ("hello" + ("hello" + (v21 + (v21 + ("\n\t" + (v21 + "foo bar"))))))) + ("\n\t" + "")) < v21), F1: -198}
                _ = v21
                _ = v22
            } else {
                var v23 uint16 = 0
                fmt.Println(v23, v23, v23)
                _ = v23
            }
            _ = fi20
        }
        if true {
            fmt.Println(v18, v18)
            var v24 string = ""
            _ = v24
        }
        _ = fi19
    }
    var v25 float64 = v18
    switch (-(v3)) {
    case -2147483648:
        var v26 int8 = 0
        _ = v26
    case 581:
        v27 := GenS2{F0: -9223372036854775808, F1: 100}
        _ = v27
    }
    if (false && ((0 < -32768) && true)) {
        switch v3 {
        case 2147483647:
            v25 -= (v18 * v25)
        case 403:
            var v28 uint32 = 2147483647
            _ = v28
        case 0:
            v29 := GenS1{F0: false, F1: int32(v3)}
            _ = v29
        }
        if true {
            v30 := GenS1{F0: (!false), F1: 1}
            _ = v30
        } else {
            switch "" {
            case "\n\t":
                v25 -= (((v18 + float64(0.1)) + float64((-2.25 * (((3.14159 + (0.1 * -2.25)) / 1.5) * float32(v25))))) / v18)
            case "":
                for fi31 := 0; fi31 < 18; fi31++ {
                    fmt.Println(v18)
                    if (3.14159 != v18) {
                        v25 -= 0.0
                    }
                    fmt.Println(v18, v25)
                    _ = fi31
                }
            default:
                v18 = 1.5
            }
        }
        var v32 int = -2147483648
        _ = v32
    } else {
        var v33 float32 = 1e10
        _ = v33
    }
    v25 += (((3.14159 + float64(0.0)) + 1e10) * v25)
    v25 = (((0.0 / v18) - v25) + 0.0)
    v34 := GenS1{F0: true, F1: 42}
    v35 := v4["\n\t"]
    switch 640 {
    case 9223372036854775807:
        v5[1] = "\n\t"
    case 127:
        var v36 int64 = 561
        _ = v36
    default:
        fmt.Println(v18, v18, v25)
    }
    for fi37 := 0; fi37 < 19; fi37++ {
        v3 -= v3
        _ = fi37
    }
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
