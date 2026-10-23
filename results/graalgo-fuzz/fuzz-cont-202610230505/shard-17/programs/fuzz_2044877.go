package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 int32; F1 uint16; F2 int16 }
type GenS5 struct { F0 int64; F1 bool; F2 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int; F1 bool; F2 uint }
type GenS8 struct { F0 uint8; F1 string }
func genFn0(p0 uint8, p1 uint32) bool {
    v6 := GenS1{F0: 1, F1: 200}
    v7 := GenS2{F0: 0}
    switch 358 {
    case -1:
        switch ("" + "\n\t") {
        case "\n\t":
            var v8 bool = true
            _ = v8
        case "":
            fmt.Println(p0)
        case "a":
            var v9 uint8 = p0
            _ = v9
        default:
            if true {
                v10 := []int{-537, -1}
                _ = v10
            }
        }
    case 738:
        var v11 bool = false
        _ = v11
    case 32767:
        if ((1 > p1) || (0.0 > 1.5)) {
            for fi12 := 0; fi12 < 14; fi12++ {
                var v13 int16 = -1
                p0 = uint8(p1)
                if ((-2147483648 >= 42) && true) {
                    v14 := GenS1{F0: 32767, F1: 227}
                    _ = v14
                } else {
                    v13 = (^(v13))
                    v15 := []int{}
                    _ = v15
                }
                _ = fi12
                _ = v13
            }
            var v16 bool = (!false)
            if ((true || (!v16)) || v16) {
                v17 := GenS3{F0: 42, F1: 253}
                var v18 uint8 = (p0 & 170)
                _ = v17
                _ = v18
            } else {
                v16 = (423 == 1)
                p1 += p1
                var v19 int = 127
                _ = v19
            }
            _ = v16
        } else {
            v20 := GenS4{F0: 42, F1: 2, F2: -599}
            var v21 float64 = (float64(((0.0 / 1.5) + (((((-2.25 + 1.5) / 1.0) / 0.1) - (-2.25 - ((0.1 + 0.0) - 1e10))) + (1.5 * 3.14159)))) * 0.0)
            switch int(p0) {
            case 2147483647:
                for fi22 := 0; fi22 < 12; fi22++ {
                    var v23 uint16 = 184
                    fmt.Println(p1)
                    _ = fi22
                    _ = v23
                }
            case 816:
                p1 += p1
            case 32768:
                var v24 int = -1
                _ = v24
            default:
                var v25 uint32 = (^(p1))
                _ = v25
            }
            _ = v20
            _ = v21
        }
    default:
        p0 = uint8(p1)
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (!false)
}
func genFn1(p0 int32, p1 uint16) int64 {
    p0 = (-(p0))
    switch "\n\t" {
    case "hello":
        var v26 string = ("\n\t" + "")
        _ = v26
    default:
        v27 := []int{-26, 32768, 4294967295, -194}
        _ = v27
    }
    fmt.Println(p1, p1)
    switch 9223372036854775807 {
    case 32767:
        var v28 int8 = 127
        _ = v28
    case 2147483647:
        var v29 float64 = ((((3.14159 + (3.14159 - 3.14159)) * 0.1) - 1.5) + -2.25)
        _ = v29
    }
    var v30 int = 42
    _ = p0
    _ = p1
    _ = v30
    return 0
}

func main() {
    var v1 string = ""
    var v2 uint8 = 2
    var v3 uint = 1
    var v4 float64 = 1.5
    var v5 uint8 = 255
    v31 := []int{429, 42, 1, 2147483647}
    v32 := GenS5{F0: -9223372036854775808, F1: false, F2: v1}
    v33 := GenS5{F0: 1, F1: (!(((false && true) && (!true)) && false)), F2: v1}
    var v34 int64 = -621
    for fi35 := 0; fi35 < 12; fi35++ {
        v36 := GenS6{F0: 4294967295}
        _ = fi35
        _ = v36
    }
    if true {
        if len(v31) > 0 {
            v31[0] = -1
        }
        if len(v31) > 0 {
            v31[0] = 0
        }
    } else {
        v37 := GenS7{F0: -338, F1: (!true), F2: 42}
        _ = v37
    }
    v38 := GenS8{F0: 255, F1: v1}
    var v39 int32 = 0
    if (744 >= 1) {
        v5 = 100
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
