package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint8; F1 bool; F2 int }
type GenS3 struct { F0 string; F1 uint8 }

func main() {
    var v1 string = "hello"
    var v2 uint16 = 2
    var v3 int16 = 1000
    var v4 float64 = 0.1
    v3 = v3
    v1 = ("foo bar" + v1)
    var v5 int32 = 2147483647
    fmt.Println(v2)
    if true {
        v2 = 255
    }
    switch v1 {
    case "":
        v4 = 0.1
    default:
        v3 = (v3 | v3)
    }
    switch v1 {
    case "a":
        for fi6 := 0; fi6 < 1; fi6++ {
            v5 += (v5 | (v5 << ((255) & 63)))
            _ = fi6
        }
    default:
        fmt.Println(v4, v3)
    }
    switch 4294967295 {
    case 128:
        if true {
            fmt.Println(v3, v1)
        }
    case 726:
        if (106 != 1) {
            if false {
                v7 := GenS1{F0: 1, F1: v3}
                v8 := GenS2{F0: 1, F1: true, F2: 4294967295}
                _ = v7
                _ = v8
            } else {
                v9 := []int{4294967295}
                v10 := GenS2{F0: 255, F1: false, F2: -1}
                _ = v10
                _ = v9
            }
            v11 := []int{32767, 1, 4294967295, 4294967295}
            v12 := GenS3{F0: v1, F1: 56}
            _ = v11
            _ = v12
        } else {
            var v13 uint8 = 128
            v14 := GenS3{F0: ((v1 + v1) + (v1 + v1)), F1: v13}
            switch -1 {
            case 898:
                v1 = v1
            case 42:
                var v15 uint8 = (^(v13))
                _ = v15
            case -9223372036854775808:
                v16 := []int{}
                _ = v16
            default:
                switch 32768 {
                case -2147483648:
                    v17 := []int{128, 32767, 645, 128}
                    _ = v17
                default:
                    var v18 bool = false
                    _ = v18
                }
            }
            _ = v13
            _ = v14
        }
    case -196:
        switch 419 {
        case 32768:
            if (!((v3 / (((v3 % (((v3 % (((^(v3))) | 1))) | 1))) | 1)) > v3)) {
                var v19 string = v1
                for fi20 := 0; fi20 < 4; fi20++ {
                    if true {
                        v4 += v4
                    }
                    v21 := GenS2{F0: 171, F1: (!(v5 == -1)), F2: 4294967295}
                    _ = fi20
                    _ = v21
                }
                v2 -= v2
                _ = v19
            }
        default:
            v22 := []int{}
            _ = v22
        }
    default:
        fmt.Println(v2, v4, v4)
    }
    v4 = -2.25
    var v23 bool = (-32768 <= v3)
    for fi24 := 0; fi24 < 7; fi24++ {
        v3 -= (v3 & v3)
        var v25 bool = (((v1 + ((((v1 + (((("foo bar" + v1) + ("\n\t" + v1)) + v1) + v1)) + (v1 + v1)) + v1) + ((v1 + "") + ("a" + "\n\t")))) + v1) < v1)
        v26 := GenS3{F0: v1, F1: 1}
        _ = fi24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 14; fi27++ {
        fmt.Println(v1, v3)
        switch v1 {
        case "a":
            var v28 int = (fi27 >> ((255) & 63))
            _ = v28
        }
        _ = fi27
    }
    if true {
        fmt.Println(v1, v5, v2)
        var v29 int64 = 394
        v5 = 2147483647
        _ = v29
    } else {
        var v30 uint32 = 662
        v2 = v2
        v31 := []int{1}
        _ = v30
        _ = v31
    }
    var v32 int = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
