package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 uint16; F1 int16; F2 int64 }
func genFn0() uint8 {
    switch -9223372036854775808 {
    case 781:
        if false {
            var v7 float32 = (-2.25 + 3.14159)
            v7 += ((1.5 / v7) - 0.0)
            var v8 uint32 = 4294967295
            _ = v7
            _ = v8
        } else {
            switch ((("a" + ("a" + ("foo bar" + "foo bar"))) + "\n\t") + "") {
            case "foo bar":
                var v9 string = (("" + ("hello" + "a")) + "a")
                _ = v9
            default:
                for fi10 := 0; fi10 < 4; fi10++ {
                    fmt.Println(fi10)
                    var v11 uint32 = 539
                    _ = fi10
                    _ = v11
                }
            }
            v12 := GenS1{F0: 42, F1: "hello"}
            for fi13 := 0; fi13 < 12; fi13++ {
                v14 := GenS2{F0: 255}
                var v15 uint16 = 42
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = v12
        }
    case 2147483647:
        v16 := GenS3{F0: 19, F1: (true || (((1000 <= 1000) || true) && false))}
        _ = v16
    case 0:
        var v17 uint16 = 255
        _ = v17
    default:
        var v18 uint64 = 80
        _ = v18
    }
    if false {
        var v19 int8 = -83
        v19 = v19
        _ = v19
    }
    return 122
}
func genFn1() uint {
    if (127 < -8) {
        if (32767 >= -175) {
            for fi20 := 0; fi20 < 2; fi20++ {
                v21 := []int{823, (fi20 / ((fi20) | 1)), 9223372036854775807}
                var v22 bool = false
                var v23 uint64 = 0
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
        }
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        for fi25 := 0; fi25 < 5; fi25++ {
            var v26 int16 = 202
            v26 -= (v26 & v26)
            _ = fi25
            _ = v26
        }
        v27 := GenS1{F0: 809, F1: (("" + "foo bar") + "\n\t")}
        var v28 string = ("\n\t" + "foo bar")
        _ = fi24
        _ = v27
        _ = v28
    }
    var v29 bool = (-100 < 100)
    fmt.Println(v29)
    _ = v29
    return 353
}

func main() {
    var v1 bool = (!(((-713 != -11) || true) || false))
    var v2 float32 = float32(float64(-2.25))
    var v3 uint8 = 1
    var v4 int8 = -60
    var v5 float64 = 3.14159
    var v6 int64 = -9223372036854775808
    switch ("" + "hello") {
    case "a":
        for fi30 := 0; fi30 < 16; fi30++ {
            v31 := []int{(fi30 / ((-1) | 1))}
            _ = fi30
            _ = v31
        }
    case "":
        v1 = ((32767 > 1) || v1)
    default:
        v32 := GenS1{F0: 2147483647, F1: ((((("" + "foo bar") + ("" + "hello")) + "a") + "hello") + "hello")}
        _ = v32
    }
    for fi33 := 0; fi33 < 5; fi33++ {
        v34 := []int{-324, (fi33 << ((1) & 63)), fi33}
        _ = fi33
        _ = v34
    }
    v35 := GenS4{F0: v4, F1: v1}
    var v36 float64 = (3.14159 * 3.14159)
    var v37 int32 = -1
    var v38 int64 = v6
    fmt.Println(v36, v38)
    v39 := []int{32768}
    v3 = (v3 | (v3 / ((v3) | 1)))
    if v1 {
        for fi40 := 0; fi40 < 2; fi40++ {
            v1 = (!v1)
            var v41 uint16 = 1
            v41 -= (^(v41))
            _ = fi40
            _ = v41
        }
        for fi42 := 0; fi42 < 18; fi42++ {
            v43 := GenS5{F0: 32767, F1: 1000, F2: (v38 - v38)}
            fmt.Println(v5, v38, v5)
            _ = fi42
            _ = v43
        }
    } else {
        for fi44 := 0; fi44 < 20; fi44++ {
            v45 := GenS4{F0: v4, F1: true}
            v3 -= 127
            _ = fi44
            _ = v45
        }
        var v46 uint8 = 127
        _ = v46
    }
    if len(v39) >= 2 {
        v39 = v39[1:len(v39)]
    }
    var v47 float64 = (0.1 - (v36 - ((0.1 * (3.14159 / v5)) / (float64(v2) * float64(3.14159)))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
