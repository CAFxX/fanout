package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint64 }

func main() {
    var v1 int64 = 42
    var v2 float32 = 1e10
    var v3 int = -81
    var v4 uint32 = 65535
    var v5 int64 = -1
    var v6 uint32 = v4
    var v7 bool = (!true)
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 int8 = -84
        v10 := GenS1{F0: v7, F1: (^(v6))}
        _ = fi8
        _ = v10
        _ = v9
    }
    if true {
        if (127 >= 31) {
            if false {
                fmt.Println(v3, v2)
                v2 -= (v2 / -2.25)
            }
            v11 := []int{v3, -117, (v3 | (v3 * (v3 * 32768)))}
            if (1e10 > ((-2.25 * (0.1 + 3.14159)) + float64(3.14159))) {
                var v12 int8 = 66
                fmt.Println(v1, v7)
                v13 := GenS2{F0: v12, F1: 42}
                _ = v12
                _ = v13
            } else {
                v1 -= -9223372036854775808
                var v14 bool = v7
                var v15 string = "foo bar"
                _ = v14
                _ = v15
            }
            _ = v11
        }
    } else {
        v16 := []int{128, v3, (v3 & (^(v3)))}
        for fi17 := 0; fi17 < 20; fi17++ {
            if (18446744073709551615 == 569) {
                v5 = -1
                for fi18 := 0; fi18 < 5; fi18++ {
                    var v19 int64 = (v1 | -9223372036854775808)
                    _ = fi18
                    _ = v19
                }
            }
            var v20 int8 = -128
            switch ("" + "hello") {
            case "":
                var v21 uint8 = 1
                _ = v21
            default:
                var v22 int16 = 1000
                _ = v22
            }
            _ = fi17
            _ = v20
        }
        _ = v16
    }
    var v23 uint = 255
    switch v3 {
    case 127:
        v24 := GenS2{F0: 127, F1: 2}
        _ = v24
    case -9223372036854775808:
        var v25 float32 = (v2 * (v2 - (float32(float64(-2.25)) + v2)))
        _ = v25
    case 32768:
        if v7 {
            v3 = 2147483647
            v1 = -1
            for fi26 := 0; fi26 < 6; fi26++ {
                v4 = (^(v6))
                v23 += v23
                _ = fi26
            }
        } else {
            fmt.Println(v2, v3)
            var v27 int = (v3 + 1)
            _ = v27
        }
    default:
        v28 := []int{(v3 / (((v3 / (((v3 ^ (^(v3)))) | 1))) | 1)), v3, 127, 32767}
        _ = v28
    }
    var v29 float32 = v2
    v30 := GenS2{F0: -128, F1: 1}
    switch 2147483647 {
    case -213:
        switch ("\n\t" + "") {
        case "hello":
            v1 += 9223372036854775807
        case "a":
            fmt.Println(v4, v6, v7)
        case "":
            var v31 string = (("a" + "foo bar") + "a")
            _ = v31
        }
    case -998:
        for fi32 := 0; fi32 < 8; fi32++ {
            for fi33 := 0; fi33 < 15; fi33++ {
                var v34 bool = (("a" + (("hello" + "hello") + ("hello" + ""))) >= "hello")
                var v35 int = (fi32 & (-(fi33)))
                _ = fi33
                _ = v34
                _ = v35
            }
            var v36 float32 = 1e10
            _ = fi32
            _ = v36
        }
    }
    var v37 uint = v23
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
