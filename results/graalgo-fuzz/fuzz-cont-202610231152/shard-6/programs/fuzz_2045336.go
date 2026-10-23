package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 float32 = (0.0 * ((0.0 / ((1.5 - ((1e10 - 3.14159) - ((float32((-2.25 - (3.14159 + float64(0.1)))) + 1.5) / ((float32(-2.25) * (float32(0.0) - -2.25)) + 3.14159)))) * (float32((((float64(1.5) / 1e10) - 1.5) - ((float64(3.14159) / ((0.1 + (3.14159 * 1.5)) - (0.0 + -2.25))) * (0.0 * -2.25)))) - 3.14159))) - -2.25))
    var v2 uint8 = 42
    v3 := make(map[string]uint32)
    var v4 uint16 = uint16(v2)
    v5 := GenS1{F0: 143, F1: (v4 >> ((4294967295) & 63)), F2: 2}
    var v6 int16 = 0
    switch 32767 {
    case -770:
        v1 += 3.14159
    case 42:
        fmt.Println(v4)
    default:
        switch ("" + (("" + ("\n\t" + ("foo bar" + "hello"))) + ("hello" + "\n\t"))) {
        case "\n\t":
            if true {
                v7 := GenS1{F0: 42, F1: v4, F2: 0}
                if false {
                    v8 := GenS1{F0: 4294967295, F1: v4, F2: 1}
                    _ = v8
                } else {
                    if (-246 > 42) {
                        v9 := GenS2{F0: 440}
                        v10 := []int{0, 32767, 4294967295, 2147483647}
                        _ = v10
                        _ = v9
                    }
                    fmt.Println(v4, v6, v6)
                    var v11 uint8 = v2
                    _ = v11
                }
                if ((false && (18446744073709551615 < 18446744073709551615)) || (!true)) {
                    v12 := GenS1{F0: -821, F1: (v4 & v4), F2: 18446744073709551615}
                    v6 += (v6 - 966)
                    var v13 float32 = v1
                    _ = v12
                    _ = v13
                } else {
                    if (((v4 / ((v4) | 1)) < v4) && (true && (!((float32(1.5) / v1) <= float32(((1e10 + 1e10) * (0.0 * float64(1e10)))))))) {
                        for fi14 := 0; fi14 < 16; fi14++ {
                            var v15 int64 = 1
                            var v16 int32 = 42
                            v2 -= (v2 + (^(v2)))
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    }
                    var v17 uint16 = (v4 - v4)
                    _ = v17
                }
                _ = v7
            }
        default:
            for fi18 := 0; fi18 < 10; fi18++ {
                v19 := GenS2{F0: (v6 ^ v6)}
                for fi20 := 0; fi20 < 3; fi20++ {
                    var v21 int = 9223372036854775807
                    _ = fi20
                    _ = v21
                }
                var v22 float32 = 1e10
                _ = fi18
                _ = v19
                _ = v22
            }
        }
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        if ("foo bar" > "") {
            fmt.Println(v1)
        } else {
            if true {
                var v24 float32 = v1
                var v25 int16 = 1
                _ = v24
                _ = v25
            }
            v26 := GenS1{F0: (^(fi23)), F1: v4, F2: 2}
            _ = v26
        }
        v1 = ((v1 / 0.1) / v1)
        _ = fi23
    }
    for fi27 := 0; fi27 < 4; fi27++ {
        var v28 float64 = (float64(v1) - (1e10 * (float64(-2.25) + 0.0)))
        _ = len(v3)
        _ = fi27
        _ = v28
    }
    v6 += (v6 - 32767)
    v4 = (v4 * v4)
    var v29 float64 = 3.14159
    var v30 int64 = -1
    v31 := GenS2{F0: 816}
    _ = len(v3)
    var v32 int8 = -90
    v33 := GenS3{F0: (v32 / (((v32 / ((-128) | 1))) | 1))}
    var v34 uint16 = 255
    var v35 uint64 = 42
    if ("a" == "\n\t") {
        v4 -= v34
        v3["a"] = 772
    } else {
        switch ("hello" + (("foo bar" + "hello") + (("a" + "") + "hello"))) {
        case "\n\t":
            for fi36 := 0; fi36 < 11; fi36++ {
                fmt.Println(fi36, fi36, v32)
                var v37 uint = uint(v34)
                _ = fi36
                _ = v37
            }
        }
        if (!((-2.25 / v1) == v1)) {
            fmt.Println(v2, v35, v32)
        }
        var v38 int64 = (^(v30))
        _ = v38
    }
    for fi39 := 0; fi39 < 7; fi39++ {
        if (!false) {
            fmt.Println(v29, v29, v6)
            var v40 uint64 = (v35 + (v35 & v35))
            _ = v40
        } else {
            var v41 uint8 = 2
            _ = v41
        }
        switch int(v35) {
        case -1:
            var v42 uint8 = (^(v2))
            _ = v42
        }
        fmt.Println(v34, v1)
        _ = fi39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
