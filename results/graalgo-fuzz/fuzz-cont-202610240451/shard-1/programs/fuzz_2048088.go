package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 bool; F2 int8 }
func genFn0() bool {
    v5 := GenS1{F0: 18446744073709551615, F1: 2147483647, F2: 1}
    var v6 int64 = 995
    fmt.Println(v6)
    _ = v5
    _ = v6
    return (true && (!(!(!true))))
}
func genFn1() int64 {
    if false {
        var v7 uint8 = 2
        v7 -= v7
        _ = v7
    }
    var v8 int = 128
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 string = "foo bar"
        v11 := GenS1{F0: 1, F1: 751, F2: 18446744073709551615}
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = v8
    return -9223372036854775808
}

func main() {
    var v1 float64 = 3.14159
    var v2 bool = ((-828 <= 0) || true)
    v3 := make(map[uint32]uint32)
    v4 := []int{0}
    if (!(!(false && ((!(!(!v2))) || v2)))) {
        v2 = v2
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        if (0 < 1) {
            v13 := v3[1]
            switch "foo bar" {
            case "hello":
                for fi14 := 0; fi14 < 13; fi14++ {
                    switch 2147483647 {
                    case -2147483648:
                        v15 := GenS1{F0: 1, F1: 507, F2: 4294967295}
                        _ = v15
                    case 1:
                        var v16 bool = (-128 <= -111)
                        _ = v16
                    }
                    fmt.Println(v2, v13)
                    var v17 string = ("hello" + "a")
                    v18 := GenS2{F0: 0}
                    _ = fi14
                    _ = v17
                    _ = v18
                }
            case "a":
                var v19 string = "a"
                _ = v19
            }
            _ = v13
        }
        for fi20 := 0; fi20 < 13; fi20++ {
            v2 = (!((((1e10 * 0.1) * 0.0) - ((((1.5 / -2.25) - -2.25) / ((1e10 - 0.1) + -2.25)) - ((((-2.25 / ((0.1 + float32(((v1 * 1e10) / (((((0.1 - 0.0) * (float64(1e10) + (v1 + v1))) * v1) + float64((0.0 - 1.5))) * v1)))) + (1.5 + float32(v1)))) * 3.14159) - 0.0) / 1.0))) <= 1.5))
            v21 := GenS3{F0: -128, F1: 32767}
            _ = fi20
            _ = v21
        }
        if ((1000 == 0) && (((v1 + v1) / (v1 - float64(1e10))) >= ((((v1 * ((((((((v1 + v1) / v1) * 0.1) * v1) + -2.25) / -2.25) * ((v1 - 3.14159) * v1)) / (v1 - (v1 - v1)))) * (((0.1 + (v1 * v1)) * 0.0) + float64(3.14159))) * (((v1 - (0.0 + (v1 + v1))) + 1e10) / v1)) * 3.14159))) {
            if true {
                fmt.Println(fi12, v2, v1)
                _ = len(v3)
            } else {
                if false {
                    v22 := GenS4{F0: 1}
                    _ = len(v3)
                    v2 = (v2 && v2)
                    _ = v22
                } else {
                    if (v2 || (18446744073709551615 == 65535)) {
                        for fi23 := 0; fi23 < 6; fi23++ {
                            v4 = append(v4, (fi12 & 1))
                            v2 = false
                            _ = fi23
                        }
                        if (!(!(!true))) {
                            for fi24 := 0; fi24 < 16; fi24++ {
                                v2 = true
                                v2 = (true || true)
                                _ = fi24
                            }
                        } else {
                            if v2 {
                                v25 := GenS3{F0: 0, F1: 42}
                                _ = v25
                            } else {
                                v2 = v2
                                var v26 string = ("hello" + "foo bar")
                                v2 = v2
                                _ = v26
                            }
                        }
                    }
                }
            }
        }
        _ = fi12
    }
    v3[2147483648] = 2
    if v2 {
        v27 := GenS3{F0: 1, F1: 1000}
        for fi28 := 0; fi28 < 7; fi28++ {
            switch "a" {
            case "\n\t":
                v1 = 1.5
            case "foo bar":
                if false {
                    v2 = v2
                }
            }
            var v29 bool = (0 > 918)
            _ = fi28
            _ = v29
        }
        _ = v27
    }
    var v30 uint8 = 1
    switch "hello" {
    case "hello":
        if true {
            var v31 int16 = 121
            if false {
                fmt.Println(v2, v31, v30)
                v32 := GenS5{F0: genFn1(), F1: v2, F2: 1}
                _ = v32
            }
            _ = v31
        }
    case "foo bar":
        fmt.Println(v30, v1, v2)
    case "a":
        if (true || v2) {
            v33 := GenS2{F0: 42}
            _ = v33
        }
    default:
        switch "a" {
        case "foo bar":
            var v34 int = 32767
            _ = v34
        }
    }
    v2 = v2
    fmt.Println(v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
