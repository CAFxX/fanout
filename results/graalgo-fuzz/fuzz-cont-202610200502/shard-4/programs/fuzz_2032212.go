package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 string }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 uint }
func genFn0() uint64 {
    var v5 string = "foo bar"
    v6 := GenS1{F0: 2147483648, F1: 8, F2: v5}
    _ = v5
    _ = v6
    return 327
}
func genFn1(p0 string, p1 int, p2 uint) int8 {
    v7 := []int{}
    for fi8 := 0; fi8 < 8; fi8++ {
        fmt.Println(p2, p0, p2)
        _ = fi8
    }
    var v9 int = p1
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := GenS1{F0: 285, F1: -2147483648, F2: p0}
        if false {
            for fi12 := 0; fi12 < 5; fi12++ {
                v7 = append(v7, v9)
                p0 = ((p0 + "hello") + (p0 + p0))
                fmt.Println(v9)
                _ = fi12
            }
            for fi13 := 0; fi13 < 5; fi13++ {
                p1 += 0
                switch (fi10 | -1) {
                case -1:
                    var v14 int32 = 1
                    _ = v14
                case -551:
                    fmt.Println(p0, p2, v9)
                case -823:
                    switch (v9 | (p1 * (p1 & p1))) {
                    case 32767:
                        if (101 > 1) {
                            var v15 uint8 = 192
                            v15 -= (v15 & v15)
                            var v16 uint = (p2 | (p2 | p2))
                            var v17 float64 = 1.5
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                    case -447:
                        _ = len(v7)
                    case 4294967295:
                        if false {
                            fmt.Println(v9, p2, p2)
                            v18 := GenS2{F0: 32, F1: p2}
                            v9 += 9223372036854775807
                            var v19 int8 = -128
                            _ = v18
                            _ = v19
                        }
                    }
                }
                var v20 string = (p0 + (p0 + p0))
                if false {
                    v9 -= 127
                    var v21 uint = (^(p2))
                    var v22 uint8 = 73
                    v20 = "\n\t"
                    var v23 uint32 = 0
                    _ = v21
                    _ = v22
                    _ = v23
                }
                var v24 uint8 = 116
                v9 = v9
                _ = fi13
                _ = v20
                _ = v24
            }
        } else {
            switch "" {
            case "foo bar":
                if len(v7) > 0 {
                    v7[0] = fi10
                }
            }
            if (0 != 1) {
                var v25 int8 = 97
                _ = v25
            }
            var v26 uint64 = genFn0()
            _ = v26
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v9
    return 100
}

func main() {
    var v1 int = 127
    var v2 int64 = 0
    v3 := make(map[int]uint32)
    var v4 int32 = -1
    _ = len(v3)
    if false {
        var v27 uint8 = 150
        v27 += v27
        _ = v27
    }
    if ((!(v1 > (v1 / ((0) | 1)))) || (!false)) {
        v28 := GenS3{F0: 241, F1: 18446744073709551615}
        _ = v28
    }
    if ((3.14159 / (float64(0.1) / 3.14159)) != 0.1) {
        var v29 float64 = (((3.14159 * 1.5) - (1.5 - 0.1)) * ((float64((float32(3.14159) / (((1e10 + 0.1) - (3.14159 / 1.0)) - 0.0))) + -2.25) + (1e10 + ((0.0 + float64(0.0)) * 3.14159))))
        var v30 uint = 65535
        _ = v29
        _ = v30
    } else {
        switch (("hello" + (("hello" + (("a" + "foo bar") + ("a" + (("hello" + ("hello" + ((("hello" + "") + ("" + ("foo bar" + "foo bar"))) + ""))) + "\n\t")))) + ("" + "hello"))) + (("\n\t" + (("" + "a") + ("\n\t" + "a"))) + ((("\n\t" + ("" + ("a" + "hello"))) + "a") + ("\n\t" + "\n\t")))) {
        case "foo bar":
            for fi31 := 0; fi31 < 20; fi31++ {
                v32 := GenS4{F0: 0}
                _ = fi31
                _ = v32
            }
        }
    }
    v4 -= (v4 - v4)
    if true {
        fmt.Println(v4, v2, v2)
        var v33 int = (v1 + 966)
        _ = v33
    }
    var v34 int = 1
    switch (v1 ^ -560) {
    case -771:
        v35 := GenS4{F0: 65535}
        _ = v35
    default:
        v36 := GenS4{F0: 1}
        _ = v36
    }
    fmt.Println(v1, v1, v2)
    v4 -= (v4 - v4)
    for fi37 := 0; fi37 < 2; fi37++ {
        fmt.Println(v2, v1, v4)
        if (301 >= genFn0()) {
            v34 = (fi37 ^ fi37)
            fmt.Println(v1, v2)
            v38 := v3[(fi37 - 32767)]
            _ = v38
        }
        _ = fi37
    }
    for fi39 := 0; fi39 < 3; fi39++ {
        v34 += v34
        var v40 int = -2147483648
        var v41 int8 = int8(v1)
        _ = fi39
        _ = v40
        _ = v41
    }
    v4 = 42
    v2 -= (v2 | 375)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
