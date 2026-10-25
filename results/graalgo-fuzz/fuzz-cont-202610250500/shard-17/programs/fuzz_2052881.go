package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 string }
type GenS2 struct { F0 int; F1 bool; F2 uint64 }
func genFn0() uint8 {
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            if false {
                var v9 int8 = 60
                var v10 uint = 1
                _ = v10
                _ = v9
            } else {
                var v11 int = -773
                v11 -= -1
                _ = v11
            }
            switch ("foo bar" + ("" + ("" + "hello"))) {
            case "foo bar":
                v12 := []int{404}
                var v13 int8 = 1
                _ = v12
                _ = v13
            case "":
                v14 := GenS1{F0: 1, F1: 127, F2: (("" + "foo bar") + "hello")}
                _ = v14
            }
            _ = fi8
        }
        fmt.Println(fi7)
        _ = fi7
    }
    var v15 int64 = 371
    if true {
        if true {
            if false {
                var v16 float32 = (1.5 * (-2.25 * 1e10))
                var v17 uint64 = 4294967295
                _ = v16
                _ = v17
            }
            switch -228 {
            case 32768:
                var v18 string = "\n\t"
                _ = v18
            case 1:
                v15 -= v15
            case 9223372036854775807:
                v15 -= (v15 + -9223372036854775808)
            default:
                switch -455 {
                case 9223372036854775807:
                    for fi19 := 0; fi19 < 13; fi19++ {
                        v15 = v15
                        _ = fi19
                    }
                }
            }
            v15 -= -544
        }
        fmt.Println(v15)
    }
    _ = v15
    return 128
}
func genFn1(p0 int16, p1 uint8, p2 uint8) int16 {
    v20 := []int{32768, int(p2)}
    var v21 int32 = -2147483648
    fmt.Println(p1)
    p1 -= (p2 & p2)
    var v22 int32 = v21
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    _ = v22
    return (-(p0))
}

func main() {
    v1 := []int{-160, 42, 739}
    var v2 float32 = (float32(0.0) / -2.25)
    var v3 int = 128
    var v4 int8 = -1
    v5 := []int{(v3 * 42), v3, 2147483647, (v3 | (v3 & v3))}
    var v6 int = (^(v3))
    if true {
        var v23 uint = 4294967295
        fmt.Println(v6, v3, v3)
        var v24 uint32 = 42
        _ = v23
        _ = v24
    }
    if (!(-24 < -9223372036854775808)) {
        v2 -= v2
    }
    var v25 uint32 = 65535
    switch (v6 | 32768) {
    case -290:
        v3 -= v3
    case -581:
        var v26 int = v6
        _ = v26
    }
    v6 = 32767
    v27 := GenS1{F0: 2147483647, F1: 255, F2: (((("\n\t" + ("hello" + (("" + ("hello" + ("\n\t" + "a"))) + "foo bar"))) + "\n\t") + "foo bar") + "foo bar")}
    for fi28 := 0; fi28 < 13; fi28++ {
        var v29 int64 = 9223372036854775807
        _ = fi28
        _ = v29
    }
    if (!false) {
        v6 = int(v4)
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        v6 -= -880
        v6 -= v6
        _ = fi30
    }
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        switch "a" {
        case "a":
            v32 := GenS2{F0: fi31, F1: true, F2: 917}
            _ = v32
        default:
            switch (v3 / (((^(fi31))) | 1)) {
            case 455:
                fmt.Println(v4)
            case 1:
                v25 -= (^(v25))
            }
        }
        var v33 int32 = 2147483647
        v34 := GenS2{F0: (fi31 % (((-(v6))) | 1)), F1: false, F2: 0}
        _ = fi31
        _ = v33
        _ = v34
    }
    fmt.Println(v6, v2, v4)
    var v35 uint16 = 0
    var v36 uint8 = 2
    var v37 string = ("hello" + "a")
    var v38 float32 = -2.25
    if false {
        v4 += v4
        v25 = 2147483647
        if false {
            v2 -= ((v2 / v2) - 0.0)
            v39 := GenS2{F0: -2147483648, F1: true, F2: 2}
            if len(v5) > 0 {
                v5[0] = 2147483647
            }
            _ = v39
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
