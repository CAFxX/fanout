package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS5 struct { F0 string; F1 uint32; F2 int }
func genFn0(p0 int32) string {
    var v5 uint = 4294967295
    v5 = v5
    _ = p0
    _ = v5
    return "a"
}

func main() {
    var v1 int64 = 1
    var v2 int = 782
    var v3 int32 = 0
    var v4 int16 = -105
    v6 := GenS1{F0: 32767, F1: v2}
    fmt.Println(v4, v3, v2)
    for fi7 := 0; fi7 < 12; fi7++ {
        fmt.Println(v2)
        v4 += 32767
        v8 := GenS1{F0: 2, F1: (fi7 / ((fi7) | 1))}
        _ = fi7
        _ = v8
    }
    var v9 uint8 = 0
    var v10 uint = 42
    for fi11 := 0; fi11 < 18; fi11++ {
        v12 := GenS2{F0: int8(v4)}
        _ = fi11
        _ = v12
    }
    var v13 string = "a"
    v14 := GenS2{F0: -93}
    v3 += 977
    v1 -= v1
    var v15 float64 = 1.5
    if (255 >= (v9 / ((v9) | 1))) {
        if true {
            if false {
                if (500 >= 47) {
                    fmt.Println(v3)
                }
                v16 := GenS3{F0: 2}
                for fi17 := 0; fi17 < 3; fi17++ {
                    v18 := GenS4{F0: 992, F1: 2147483647, F2: 65535}
                    var v19 float64 = 1.5
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = v16
            }
        }
        for fi20 := 0; fi20 < 20; fi20++ {
            v10 = (v10 % ((65535) | 1))
            if (true || false) {
                switch "hello" {
                case "a":
                    for fi21 := 0; fi21 < 7; fi21++ {
                        var v22 string = ""
                        fmt.Println(v3, v13)
                        v23 := []int{32768, v2, 32768}
                        v24 := GenS4{F0: 4294967295, F1: 65535, F2: 0}
                        _ = fi21
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                case "hello":
                    v15 += v15
                default:
                    var v25 uint64 = uint64(v10)
                    _ = v25
                }
                if false {
                    var v26 int64 = (v1 * v1)
                    if (v10 >= (v10 + 18446744073709551615)) {
                        v27 := GenS2{F0: -1}
                        var v28 int8 = 100
                        v2 += int(v10)
                        _ = v27
                        _ = v28
                    } else {
                        switch "\n\t" {
                        case "\n\t":
                            var v29 uint32 = 2
                            _ = v29
                        default:
                            v1 -= v26
                        }
                        v3 += 265
                    }
                    _ = v26
                } else {
                    v30 := []int{(^(fi20)), (-(v2))}
                    v1 -= v1
                    v31 := GenS5{F0: v13, F1: 42, F2: (v2 * fi20)}
                    _ = v30
                    _ = v31
                }
                var v32 int64 = v1
                _ = v32
            } else {
                var v33 int16 = 1000
                v33 = (v33 + v4)
                fmt.Println(v15, v1)
                _ = v33
            }
            _ = fi20
        }
    }
    if true {
        v3 -= (v3 | (v3 << ((v10) & 63)))
        var v34 int32 = -878
        v35 := GenS2{F0: -1}
        _ = v34
        _ = v35
    }
    var v36 uint = (v10 ^ v10)
    var v37 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v6
    _ = v9
}
