package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int; F1 uint64 }
func genFn0() int16 {
    for fi7 := 0; fi7 < 2; fi7++ {
        switch "\n\t" {
        case "a":
            var v8 string = ""
            _ = v8
        }
        _ = fi7
    }
    var v9 int16 = -1
    _ = v9
    return (v9 >> ((42) & 63))
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = float64((1e10 * float32(3.14159)))
    var v3 bool = false
    v4 := make(map[uint32]bool)
    var v5 int32 = 876
    v6 := []int{-9223372036854775808}
    for fi10 := 0; fi10 < 9; fi10++ {
        switch "hello" {
        case "\n\t":
            var v11 int8 = -7
            _ = v11
        case "hello":
            v12 := GenS1{F0: 32767, F1: 94}
            _ = v12
        default:
            var v13 int16 = -72
            _ = v13
        }
        _ = fi10
    }
    v2 = ((1e10 * v2) / v2)
    for fi14 := 0; fi14 < 9; fi14++ {
        v6 = append(v6, (fi14 - fi14))
        var v15 int32 = (^(v5))
        _ = fi14
        _ = v15
    }
    var v16 float32 = 1e10
    fmt.Println(v5, v16)
    for fi17 := 0; fi17 < 17; fi17++ {
        switch "" {
        case "\n\t":
            var v18 int64 = 1
            _ = v18
        }
        var v19 uint8 = 42
        if (627 != 774) {
            v20 := GenS1{F0: 42, F1: 2}
            _ = v20
        }
        _ = fi17
        _ = v19
    }
    switch "a" {
    case "hello":
        _ = len(v1)
    default:
        v2 = ((v2 * (v2 - v2)) * v2)
    }
    if v3 {
        var v21 int = 32768
        for fi22 := 0; fi22 < 16; fi22++ {
            v16 = v16
            v23 := GenS1{F0: 0, F1: 82}
            v3 = (v5 <= (v5 % ((v5) | 1)))
            _ = fi22
            _ = v23
        }
        _ = v21
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        var v25 float32 = 0.0
        for fi26 := 0; fi26 < 12; fi26++ {
            if (!false) {
                v16 = v16
                v27 := GenS1{F0: 255, F1: 1}
                var v28 int32 = 42
                v29 := GenS2{F0: 4294967295, F1: 2147483648}
                _ = v27
                _ = v28
                _ = v29
            }
            if (!((("hello" + "foo bar") + "\n\t") < "foo bar")) {
                if v3 {
                    if true {
                        v25 -= v16
                    } else {
                        switch "a" {
                        case "\n\t":
                            _ = len(v6)
                        case "":
                            v3 = true
                        default:
                            v30 := GenS2{F0: 65535, F1: 4294967295}
                            _ = v30
                        }
                        v16 += (v16 / ((1.5 / v16) - 1.5))
                        v16 -= ((v25 + v16) + (v16 - (0.0 - v16)))
                    }
                    v3 = false
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    v5 += (-(v5))
                    v5 = v5
                }
                fmt.Println(v16, v25, v25)
                if true {
                    v5 = (v5 / (((-(v5))) | 1))
                    v31 := GenS3{F0: (v5 << ((42) & 63))}
                    _ = v31
                }
            }
            var v32 int32 = v5
            _ = fi26
            _ = v32
        }
        var v33 int16 = 0
        _ = fi24
        _ = v25
        _ = v33
    }
    var v34 int16 = 1
    for fi35 := 0; fi35 < 10; fi35++ {
        v36 := GenS2{F0: 153, F1: 2147483647}
        v34 = genFn0()
        _ = fi35
        _ = v36
    }
    v37 := GenS4{F0: -2147483648, F1: 4294967295}
    delete(v4, 0)
    _ = len(v6)
    v38 := GenS2{F0: 2, F1: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v34
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
