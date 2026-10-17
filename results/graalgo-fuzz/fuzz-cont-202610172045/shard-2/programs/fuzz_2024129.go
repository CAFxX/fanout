package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int16 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 int8; F1 int32; F2 string }
type GenS5 struct { F0 bool; F1 uint16 }
type GenS6 struct { F0 bool; F1 uint16; F2 string }
type GenS7 struct { F0 bool; F1 bool; F2 int64 }
func genFn0(p0 uint16, p1 string) uint32 {
    fmt.Println(p0, p1)
    v6 := GenS1{F0: 32767, F1: 426, F2: 1}
    _ = p0
    _ = p1
    _ = v6
    return 2147483647
}
func genFn1(p0 int, p1 uint64, p2 uint8) uint32 {
    v7 := []int{(p0 & (-(p0))), (-(p0)), 127}
    fmt.Println(p2, p1)
    p1 -= (p1 << ((2) & 63))
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := GenS1{F0: 32767, F1: 0, F2: 0}
        p0 = p0
        for fi10 := 0; fi10 < 10; fi10++ {
            switch "a" {
            case "foo bar":
                var v11 uint64 = p1
                _ = v11
            }
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    var v1 int64 = 42
    v2 := make(map[string]int)
    v3 := []uint32{1}
    var v4 float32 = 1.5
    var v5 int = 2147483647
    if len(v3) > 0 {
        v3[0] = genFn0(1, ("hello" + "a"))
    }
    var v12 int64 = 9223372036854775807
    v13 := GenS2{F0: 255, F1: 1000, F2: 42}
    var v14 float64 = 0.1
    var v15 uint8 = 116
    if false {
        _ = len(v3)
        var v16 string = ("hello" + "")
        for fi17 := 0; fi17 < 16; fi17++ {
            var v18 uint8 = v15
            if false {
                for fi19 := 0; fi19 < 1; fi19++ {
                    if len(v3) > 0 {
                        v3[0] = 4294967295
                    }
                    v18 = (v15 % ((v15) | 1))
                    fmt.Println(v4, fi19)
                    _ = fi19
                }
                v4 = v4
                fmt.Println(v1, v15, fi17)
            } else {
                fmt.Println(v12, v15, v12)
                for fi20 := 0; fi20 < 6; fi20++ {
                    v5 -= v5
                    var v21 int32 = 42
                    _ = fi20
                    _ = v21
                }
                v22 := GenS1{F0: 32767, F1: 42, F2: 0}
                _ = v22
            }
            _ = fi17
            _ = v18
        }
        _ = v16
    } else {
        v23 := GenS3{F0: -32768, F1: 2}
        _ = v23
    }
    v24 := GenS4{F0: -128, F1: -73, F2: ("a" + "foo bar")}
    var v25 int32 = -2147483648
    var v26 uint64 = 0
    v4 += (v4 + (v4 / 0.1))
    if (v25 == (v25 * v25)) {
        v27 := GenS5{F0: (!false), F1: 1}
        v28 := GenS4{F0: -49, F1: (v25 >> ((1) & 63)), F2: "hello"}
        _ = v27
        _ = v28
    } else {
        v3 = append(v3, 2147483648)
    }
    v14 += (1.5 / ((v14 + ((v14 + 3.14159) - 1e10)) - v14))
    switch ("hello" + "\n\t") {
    case "foo bar":
        _ = len(v2)
    case "hello":
        fmt.Println(v12, v1, v14)
    case "":
        for fi29 := 0; fi29 < 20; fi29++ {
            for fi30 := 0; fi30 < 3; fi30++ {
                if (true || ((v1 >> ((520) & 63)) > (v1 - (v12 * -465)))) {
                    fmt.Println(v12, v12)
                    v14 = (v14 - (float64(v4) / v14))
                }
                if false {
                    fmt.Println(v4)
                    v31 := GenS6{F0: true, F1: 42, F2: "\n\t"}
                    v32 := GenS7{F0: ((false && (!(-2147483648 < -2147483648))) && false), F1: true, F2: v1}
                    _ = v31
                    _ = v32
                } else {
                    v25 = -1
                    var v33 bool = (0 >= (v15 * v15))
                    _ = v33
                }
                v3 = append(v3, 186)
                _ = fi30
            }
            v25 = 1
            _ = fi29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
