package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint; F1 bool; F2 uint8 }
type GenS3 struct { F0 bool }

func main() {
    var v1 float64 = ((0.0 + (-2.25 / 0.1)) / -2.25)
    var v2 int = 677
    var v3 float64 = 1e10
    v3 = (3.14159 + (v1 - v3))
    var v4 string = ("hello" + "a")
    v4 = v4
    if false {
        if (65535 > 536) {
            v4 = v4
            v3 += 1e10
        }
        v2 -= (v2 - 1)
    }
    v5 := GenS1{F0: ("foo bar" + "\n\t"), F1: 711}
    for fi6 := 0; fi6 < 16; fi6++ {
        fmt.Println(fi6, fi6, v3)
        _ = fi6
    }
    if true {
        v7 := []int{(v2 / ((9223372036854775807) | 1))}
        v7 = append(v7, (v2 / ((0) | 1)))
        fmt.Println(v3)
        _ = v7
    } else {
        for fi8 := 0; fi8 < 18; fi8++ {
            v2 = (fi8 << ((1) & 63))
            var v9 int32 = 145
            var v10 bool = true
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    switch 757 {
    case 774:
        switch "" {
        case "foo bar":
            v1 -= 0.1
        case "":
            v2 -= 128
        case "a":
            v11 := []int{}
            _ = v11
        default:
            for fi12 := 0; fi12 < 8; fi12++ {
                v2 = (v2 / (((^(fi12))) | 1))
                fmt.Println(v3, v3)
                fmt.Println(v1)
                _ = fi12
            }
        }
    case -282:
        v13 := []int{127, 127}
        _ = v13
    case -2147483648:
        var v14 uint16 = 2
        _ = v14
    default:
        if false {
            v15 := GenS1{F0: v4, F1: 269}
            _ = v15
        } else {
            v16 := GenS1{F0: v4, F1: 4294967295}
            _ = v16
        }
    }
    var v17 int16 = -619
    v2 = (v2 / (((v2 << ((42) & 63))) | 1))
    fmt.Println(v17)
    if (42 <= 0) {
        if false {
            if false {
                switch (v4 + (v4 + "foo bar")) {
                case "hello":
                    v18 := GenS1{F0: ("\n\t" + v4), F1: 65535}
                    _ = v18
                case "\n\t":
                    v19 := GenS2{F0: 65535, F1: (false && ((!(false && (v3 != (v1 / ((v1 - (v3 + v1)) - v1))))) || true)), F2: 2}
                    _ = v19
                case "":
                    if (2 < 127) {
                        for fi20 := 0; fi20 < 16; fi20++ {
                            for fi21 := 0; fi21 < 17; fi21++ {
                                fmt.Println(fi20, v17, v1)
                                v22 := []int{}
                                _ = fi21
                                _ = v22
                            }
                            var v23 bool = true
                            v1 = float64((0.0 + float32(((-2.25 / v3) + v1))))
                            _ = fi20
                            _ = v23
                        }
                    }
                default:
                    if true {
                        v24 := GenS3{F0: true}
                        v17 += v17
                        _ = v24
                    }
                }
            }
            if false {
                var v25 int32 = -1
                v26 := []int{(-(v2))}
                _ = v25
                _ = v26
            }
        } else {
            var v27 string = "a"
            v27 = "foo bar"
            var v28 float32 = float32(((v1 - v1) - 1.5))
            _ = v27
            _ = v28
        }
        fmt.Println(v2, v1, v3)
    } else {
        if false {
            fmt.Println(v3, v1, v1)
        }
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        var v30 string = v4
        _ = fi29
        _ = v30
    }
    var v31 int = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
