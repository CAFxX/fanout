package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int64; F2 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 uint8; F2 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8 }
func genFn0(p0 uint) int8 {
    fmt.Println(p0)
    var v5 uint32 = 803
    v6 := []int{2147483647, 1}
    _ = p0
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 int64) int16 {
    v7 := GenS1{F0: (42 <= 925)}
    switch "hello" {
    case "foo bar":
        var v8 int64 = 0
        _ = v8
    case "a":
        switch "\n\t" {
        case "\n\t":
            v9 := GenS2{F0: true, F1: p0, F2: 784}
            _ = v9
        case "a":
            var v10 int32 = 1
            _ = v10
        case "foo bar":
            v11 := GenS3{F0: -128}
            _ = v11
        default:
            var v12 float32 = float32(0.1)
            _ = v12
        }
    case "":
        var v13 bool = (!(0 > 81))
        _ = v13
    default:
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 float32 = (1.5 - 3.14159)
            _ = fi14
            _ = v15
        }
    }
    p0 += p0
    v16 := []int{}
    v17 := GenS4{F0: true, F1: 157, F2: 0}
    _ = p0
    _ = v16
    _ = v17
    _ = v7
    return 1
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint8 = 127
    var v3 uint16 = 42
    var v4 int32 = 2147483647
    var v18 uint8 = v2
    v19 := GenS3{F0: 100}
    for fi20 := 0; fi20 < 7; fi20++ {
        if (!false) {
            if (-66 < genFn0((^(v1)))) {
                fmt.Println(v18, v4, v3)
                for fi21 := 0; fi21 < 1; fi21++ {
                    var v22 float64 = ((-2.25 / -2.25) - 1.5)
                    var v23 uint32 = 2147483647
                    v1 = v1
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                for fi24 := 0; fi24 < 8; fi24++ {
                    var v25 float64 = (((float64(-2.25) * 0.0) * 0.0) * (0.1 + 3.14159))
                    var v26 uint16 = v3
                    var v27 uint64 = 4294967295
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                }
            } else {
                v2 -= (v2 % (((^(v2))) | 1))
                var v28 uint32 = 65535
                v29 := []int{}
                _ = v28
                _ = v29
            }
        }
        switch "" {
        case "a":
            var v30 uint16 = v3
            _ = v30
        default:
            switch (fi20 & (fi20 & fi20)) {
            case 2147483647:
                var v31 float32 = (0.1 * 0.1)
                _ = v31
            }
        }
        _ = fi20
    }
    v2 = 75
    if false {
        v32 := GenS4{F0: false, F1: (v18 >> ((255) & 63)), F2: 386}
        v3 = (v3 - uint16(v2))
        v4 = (v4 & v4)
        _ = v32
    } else {
        v33 := []int{}
        v2 += v18
        _ = v33
    }
    v34 := []int{}
    var v35 uint = (v1 << ((0) & 63))
    if ((v3 * v3) < 32767) {
        _ = len(v34)
        v36 := GenS5{F0: v18}
        _ = v36
    }
    v37 := GenS1{F0: ((!false) && (true && false))}
    v1 = 368
    switch 1 {
    case -176:
        fmt.Println(v4, v18)
    case 42:
        fmt.Println(v35, v4)
    case -2147483648:
        v4 = v4
    }
    fmt.Println(v2, v3, v4)
    v2 = (^(v18))
    var v38 int = -958
    v38 = v38
    v3 = v3
    switch (v38 ^ -1) {
    case 42:
        for fi39 := 0; fi39 < 10; fi39++ {
            fmt.Println(v4, v4)
            if true {
                v40 := GenS6{F0: 100}
                var v41 int = 1
                _ = v40
                _ = v41
            } else {
                v4 = (v4 & int32(v2))
            }
            var v42 int64 = 859
            _ = fi39
            _ = v42
        }
    case 0:
        v38 -= v38
    default:
        var v43 uint64 = 2
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v37
    _ = v38
    _ = v4
}
