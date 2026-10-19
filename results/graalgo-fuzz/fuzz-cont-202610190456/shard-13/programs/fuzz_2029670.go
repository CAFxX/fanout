package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint8; F1 int16; F2 bool }
type GenS3 struct { F0 string; F1 int32 }
func genFn0() uint64 {
    var v5 int8 = 0
    v5 += 42
    v5 = v5
    v6 := []int{-348, -2147483648}
    _ = v5
    _ = v6
    return 780
}

func main() {
    var v1 bool = (!true)
    var v2 bool = false
    var v3 int16 = 1
    var v4 uint32 = 2
    if (743 == 0) {
        var v7 float32 = ((((1.5 * (1e10 - 0.1)) * 3.14159) / (1e10 - (1.5 + (0.0 - 3.14159)))) / -2.25)
        for fi8 := 0; fi8 < 13; fi8++ {
            if false {
                if (!v2) {
                    v9 := GenS1{F0: "hello", F1: v4, F2: 613}
                    switch "foo bar" {
                    case "foo bar":
                        v10 := []int{}
                        _ = v10
                    default:
                        v11 := GenS1{F0: ("foo bar" + ("foo bar" + "")), F1: 0, F2: v4}
                        _ = v11
                    }
                    switch "foo bar" {
                    case "\n\t":
                        v2 = (0 == 42)
                    case "":
                        if false {
                            var v12 float32 = -2.25
                            v12 -= (v7 + v7)
                            _ = v12
                        }
                    case "foo bar":
                        v7 = v7
                    }
                    _ = v9
                } else {
                    switch fi8 {
                    case 128:
                        v13 := GenS1{F0: "hello", F1: (v4 * (v4 << ((255) & 63))), F2: 2147483648}
                        _ = v13
                    default:
                        var v14 string = ("\n\t" + "\n\t")
                        _ = v14
                    }
                }
            }
            _ = fi8
        }
        _ = v7
    }
    var v15 int64 = 121
    if (!v2) {
        switch ("hello" + "a") {
        case "\n\t":
            v16 := GenS1{F0: ("foo bar" + "foo bar"), F1: 2147483647, F2: 0}
            _ = v16
        case "foo bar":
            v2 = (!false)
        case "hello":
            var v17 string = ""
            _ = v17
        default:
            if v1 {
                v18 := GenS2{F0: 2, F1: (v3 & v3), F2: true}
                switch 767 {
                case -149:
                    v19 := []int{-1, -1, 127}
                    _ = v19
                case 20:
                    var v20 uint64 = 463
                    _ = v20
                }
                if (216 != -419) {
                    for fi21 := 0; fi21 < 17; fi21++ {
                        if (0 > 815) {
                            v4 = 1
                        }
                        _ = fi21
                    }
                }
                _ = v18
            } else {
                switch 32768 {
                case 9223372036854775807:
                    v22 := []int{-14, int(v15), -9223372036854775808, 32767}
                    _ = v22
                case -2147483648:
                    if v1 {
                        v3 -= 599
                    } else {
                        for fi23 := 0; fi23 < 1; fi23++ {
                            fmt.Println(v15, v2, v3)
                            _ = fi23
                        }
                        var v24 int64 = v15
                        v1 = v2
                        _ = v24
                    }
                case 4294967295:
                    v25 := GenS3{F0: "\n\t", F1: -937}
                    _ = v25
                default:
                    v15 += (-(v15))
                }
            }
        }
        for fi26 := 0; fi26 < 9; fi26++ {
            for fi27 := 0; fi27 < 2; fi27++ {
                fmt.Println(v2)
                _ = fi27
            }
            for fi28 := 0; fi28 < 16; fi28++ {
                var v29 uint64 = 333
                _ = fi28
                _ = v29
            }
            v30 := []int{}
            _ = fi26
            _ = v30
        }
    }
    v4 = 0
    for fi31 := 0; fi31 < 17; fi31++ {
        var v32 int32 = -19
        _ = fi31
        _ = v32
    }
    var v33 uint = 4294967295
    var v34 uint32 = (v4 % ((734) | 1))
    switch 42 {
    case -989:
        v35 := GenS3{F0: "\n\t", F1: int32(v33)}
        _ = v35
    default:
        var v36 bool = v1
        _ = v36
    }
    v37 := GenS1{F0: "a", F1: 2147483647, F2: (v34 >> ((v33) & 63))}
    var v38 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v37
    _ = v38
    _ = v4
}
