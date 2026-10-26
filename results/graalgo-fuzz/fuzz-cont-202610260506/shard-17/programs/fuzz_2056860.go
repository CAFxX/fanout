package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int8; F1 int8; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 uint }
func genFn0() int {
    var v4 float64 = 0.1
    switch "" {
    case "a":
        for fi5 := 0; fi5 < 18; fi5++ {
            var v6 float64 = v4
            var v7 int = (fi5 * (fi5 % ((-1) | 1)))
            _ = fi5
            _ = v6
            _ = v7
        }
    case "\n\t":
        if (false && (!(!((1.5 != 0.1) || true)))) {
            fmt.Println(v4)
            var v8 uint8 = 183
            _ = v8
        }
    }
    fmt.Println(v4)
    _ = v4
    return -1
}

func main() {
    var v1 int32 = -929
    v2 := make(map[uint32]string)
    v3 := make(map[int]int)
    var v9 bool = (((-275 <= 0) || true) && true)
    v10 := GenS1{F0: 18446744073709551615, F1: v1}
    delete(v2, 214)
    var v11 uint32 = 2147483647
    fmt.Println(v1, v11, v9)
    v11 = v11
    v12 := GenS1{F0: 65535, F1: (v1 - 42)}
    v13 := GenS2{F0: -1, F1: -1, F2: 2}
    v1 += v1
    delete(v3, -9223372036854775808)
    var v14 uint32 = 2147483647
    v1 -= 299
    v9 = (417 <= (-(v1)))
    var v15 uint16 = 32767
    var v16 int = 122
    if ((v9 || v9) && (0 > -9223372036854775808)) {
        v15 += v15
    } else {
        var v17 float64 = 0.0
        var v18 int64 = 0
        v1 += (v1 | (-(v1)))
        _ = v17
        _ = v18
    }
    if (42 != -531) {
        if false {
            v19 := GenS1{F0: 255, F1: v1}
            var v20 int32 = (-(v1))
            _ = v19
            _ = v20
        } else {
            for fi21 := 0; fi21 < 16; fi21++ {
                v3[v16] = (v16 / ((genFn0()) | 1))
                _ = fi21
            }
        }
        v22 := GenS3{F0: 360}
        _ = v22
    } else {
        var v23 string = "hello"
        _ = v23
    }
    if (v1 == (v1 % ((v1) | 1))) {
        switch ("hello" + ("foo bar" + "hello")) {
        case "\n\t":
            var v24 int32 = v1
            _ = v24
        case "a":
            var v25 int8 = 36
            _ = v25
        default:
            v11 = v11
        }
        switch v16 {
        case 127:
            for fi26 := 0; fi26 < 19; fi26++ {
                var v27 float32 = 3.14159
                _ = fi26
                _ = v27
            }
        case 923:
            fmt.Println(v15)
        default:
            var v28 float64 = -2.25
            _ = v28
        }
        v15 -= v15
    } else {
        switch (v16 % ((v16) | 1)) {
        case 1:
            if (4294967295 > 4294967295) {
                var v29 float32 = 1e10
                v30 := GenS4{F0: ("" + ("foo bar" + "\n\t")), F1: 0}
                _ = v29
                _ = v30
            } else {
                switch (("foo bar" + ("a" + ("" + ""))) + "") {
                case "a":
                    if v9 {
                        if v9 {
                            var v31 int64 = -9223372036854775808
                            _ = v31
                        }
                        for fi32 := 0; fi32 < 11; fi32++ {
                            v9 = v9
                            v2[780] = ("hello" + "hello")
                            v33 := v3[(fi32 << ((1) & 63))]
                            _ = fi32
                            _ = v33
                        }
                    } else {
                        var v34 int32 = v1
                        v35 := GenS4{F0: ("" + "foo bar"), F1: 68}
                        _ = v34
                        _ = v35
                    }
                }
            }
        case 127:
            v36 := []int{127, (v16 ^ (v16 & (v16 & (-(v16))))), v16}
            _ = v36
        }
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v9
}
