package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 bool; F1 uint }
func genFn0(p0 uint8, p1 uint8) uint16 {
    var v7 int16 = 42
    var v8 bool = (!false)
    var v9 uint16 = 2
    for fi10 := 0; fi10 < 5; fi10++ {
        p1 += (p1 + p1)
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return v9
}
func genFn1(p0 int32, p1 uint64) int32 {
    if true {
        p1 = (^(p1))
    } else {
        fmt.Println(p0)
    }
    v11 := []int{}
    var v12 uint16 = 375
    v13 := GenS1{F0: 13, F1: 457}
    for fi14 := 0; fi14 < 20; fi14++ {
        p0 += p0
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        for fi15 := 0; fi15 < 17; fi15++ {
            fmt.Println(p0, p0)
            if true {
                var v16 int = fi14
                var v17 int = (fi14 | (fi14 ^ fi14))
                _ = len(v11)
                _ = v16
                _ = v17
            }
            v12 += v12
            switch 185 {
            case 502:
                v18 := GenS2{F0: 0, F1: genFn0(127, 1)}
                _ = v18
            case 162:
                var v19 uint8 = 237
                _ = v19
            case 896:
                switch (fi15 - fi15) {
                case 9223372036854775807:
                    switch (("\n\t" + "foo bar") + "a") {
                    case "":
                        v12 = v12
                    }
                }
            }
            _ = fi15
        }
        _ = fi14
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    return (p0 + (^(p0)))
}

func main() {
    var v1 int8 = 26
    var v2 uint = 4294967295
    v3 := make(map[int]int)
    var v4 float64 = (((1.5 - float64((0.0 + (3.14159 + 0.0)))) - 1e10) * 3.14159)
    v5 := make(map[int]bool)
    v6 := make(map[uint32]string)
    v1 += v1
    v2 += v2
    for fi20 := 0; fi20 < 20; fi20++ {
        v21 := GenS3{F0: false, F1: 18446744073709551615}
        _ = fi20
        _ = v21
    }
    var v22 bool = (!true)
    var v23 float64 = v4
    for fi24 := 0; fi24 < 7; fi24++ {
        var v25 int32 = 42
        _ = fi24
        _ = v25
    }
    v26 := GenS3{F0: (!("\n\t" >= (("a" + "a") + "\n\t"))), F1: 698}
    fmt.Println(v23, v2, v1)
    var v27 int8 = 20
    switch "hello" {
    case "foo bar":
        var v28 int8 = v1
        _ = v28
    }
    var v29 float64 = ((float64((3.14159 - float32(v4))) / (3.14159 / 3.14159)) + (v4 * (float64((((3.14159 / 1.0) + 3.14159) / 1e10)) * v23)))
    switch "a" {
    case "a":
        for fi30 := 0; fi30 < 15; fi30++ {
            v4 += 3.14159
            v31 := GenS1{F0: 1, F1: -32768}
            _ = fi30
            _ = v31
        }
    case "hello":
        for fi32 := 0; fi32 < 19; fi32++ {
            v33 := []int{(fi32 + (fi32 & fi32)), (fi32 ^ (fi32 & fi32)), (fi32 + fi32)}
            if len(v33) >= 2 {
                v33 = v33[1:len(v33)]
            }
            _ = fi32
            _ = v33
        }
    case "foo bar":
        if true {
            for fi34 := 0; fi34 < 14; fi34++ {
                v35 := GenS1{F0: 0, F1: 0}
                v36 := GenS2{F0: (v2 ^ 716), F1: 310}
                v1 = v27
                _ = fi34
                _ = v35
                _ = v36
            }
        } else {
            var v37 int64 = -9223372036854775808
            _ = v37
        }
    default:
        v27 -= v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
