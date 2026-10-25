package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int8; F1 int32; F2 uint8 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 int64; F1 int32 }
func genFn0(p0 uint64) uint16 {
    var v6 float64 = 1e10
    v7 := []int{128, 32767}
    fmt.Println(p0, p0)
    fmt.Println(p0, p0)
    v7 = append(v7, -734)
    _ = p0
    _ = v6
    _ = v7
    return 255
}

func main() {
    var v1 float64 = 3.14159
    v2 := []string{"", (("hello" + "hello") + ((("foo bar" + "a") + ((("" + ("\n\t" + "hello")) + "foo bar") + "\n\t")) + "a"))}
    var v3 int32 = 42
    var v4 uint16 = 255
    var v5 uint64 = 4294967295
    var v8 string = "\n\t"
    var v9 int32 = 42
    var v10 int64 = 9223372036854775807
    v4 -= genFn0(v5)
    v11 := GenS1{F0: (!(!false))}
    v1 -= v1
    if false {
        v8 = "\n\t"
    }
    v3 += v3
    var v12 float64 = (-2.25 - 0.0)
    for fi13 := 0; fi13 < 1; fi13++ {
        for fi14 := 0; fi14 < 8; fi14++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v15 := GenS1{F0: false}
            _ = fi14
            _ = v15
        }
        var v16 int = (fi13 & -1)
        _ = fi13
        _ = v16
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        switch ((v8 + "") + "hello") {
        case "":
            for fi18 := 0; fi18 < 20; fi18++ {
                v19 := GenS2{F0: (1 != 2147483648), F1: 0, F2: 0}
                var v20 int64 = v10
                _ = fi18
                _ = v19
                _ = v20
            }
        case "a":
            if (((((v8 + v8) + v8) + v8) == ((v8 + v8) + v8)) || (!true)) {
                if ((0.1 - 1.5) == 1e10) {
                    var v21 uint32 = 2
                    for fi22 := 0; fi22 < 6; fi22++ {
                        var v23 int32 = -2147483648
                        v5 = v5
                        _ = fi22
                        _ = v23
                    }
                    _ = v21
                } else {
                    for fi24 := 0; fi24 < 15; fi24++ {
                        v9 += (v9 * 0)
                        var v25 int32 = 2147483647
                        var v26 int = -168
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                fmt.Println(v12, v9, v9)
            } else {
                var v27 string = v8
                v28 := GenS3{F0: 1, F1: 0, F2: 49}
                _ = v27
                _ = v28
            }
        default:
            v29 := GenS1{F0: (!false)}
            _ = v29
        }
        _ = fi17
    }
    for fi30 := 0; fi30 < 13; fi30++ {
        v31 := GenS4{F0: (fi30 | fi30), F1: (v9 / (((^(v9))) | 1))}
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 8; fi32++ {
        switch (fi32 >> ((2) & 63)) {
        case 617:
            v8 = v8
        case 32767:
            _ = len(v2)
        case 804:
            var v33 int32 = (-(v3))
            _ = v33
        default:
            for fi34 := 0; fi34 < 18; fi34++ {
                v35 := GenS5{F0: (v10 ^ (v10 - -724)), F1: (-(v9))}
                v5 += v5
                fmt.Println(v9, v10)
                var v36 uint16 = (v4 << ((639) & 63))
                _ = fi34
                _ = v35
                _ = v36
            }
        }
        v37 := GenS3{F0: 70, F1: (v3 % ((-2147483648) | 1)), F2: 153}
        _ = fi32
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
