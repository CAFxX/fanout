package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int64; F1 int64; F2 int32 }

func main() {
    var v1 float64 = 3.14159
    var v2 float32 = 0.1
    var v3 float64 = 1e10
    var v4 int16 = 0
    var v5 int8 = 100
    switch "\n\t" {
    case "hello":
        v6 := GenS1{F0: ("hello" + "\n\t"), F1: -692}
        _ = v6
    case "\n\t":
        if false {
            v7 := []int{-588, 196, 576}
            v2 = 3.14159
            _ = v7
        }
    }
    if (2 <= 255) {
        for fi8 := 0; fi8 < 7; fi8++ {
            v9 := GenS2{F0: 924, F1: 0, F2: 42}
            switch fi8 {
            case 32768:
                var v10 int8 = (^(v5))
                _ = v10
            case 42:
                var v11 bool = (-637 == -1)
                _ = v11
            default:
                for fi12 := 0; fi12 < 20; fi12++ {
                    v3 = v3
                    _ = fi12
                }
            }
            _ = fi8
            _ = v9
        }
    }
    v2 = (float32(1e10) - 1e10)
    var v13 uint = 255
    var v14 int64 = -1
    if ((!(!true)) && ((v13 | v13) != 4294967295)) {
        var v15 uint = (v13 & (v13 >> ((v13) & 63)))
        var v16 uint8 = 132
        if ((v4 - 1000) <= (v4 ^ (^(v4)))) {
            switch ("hello" + "\n\t") {
            case "":
                v17 := GenS1{F0: ("foo bar" + (("foo bar" + ("foo bar" + "\n\t")) + "hello")), F1: 1}
                _ = v17
            case "hello":
                var v18 int = 0
                _ = v18
            case "foo bar":
                switch -1 {
                case 1:
                    fmt.Println(v5, v5)
                case 9223372036854775807:
                    v2 = v2
                }
            default:
                for fi19 := 0; fi19 < 7; fi19++ {
                    fmt.Println(fi19, v5, v16)
                    _ = fi19
                }
            }
        }
        _ = v15
        _ = v16
    }
    v20 := GenS2{F0: v14, F1: v14, F2: -1}
    v21 := GenS2{F0: v14, F1: (-(v14)), F2: 0}
    v4 -= v4
    v4 = (v4 | (v4 + (v4 * (v4 / (((v4 - (v4 - -38))) | 1)))))
    var v22 int64 = (v14 | 9223372036854775807)
    fmt.Println(v14, v2)
    fmt.Println(v2, v3, v14)
    switch 32767 {
    case -2147483648:
        v23 := GenS1{F0: ("hello" + "foo bar"), F1: 0}
        _ = v23
    case -1:
        var v24 uint = (v13 - v13)
        _ = v24
    default:
        v25 := GenS2{F0: (v22 >> ((v13) & 63)), F1: v22, F2: 521}
        _ = v25
    }
    v4 -= 42
    switch ("a" + "foo bar") {
    case "foo bar":
        var v26 int16 = (v4 + v4)
        _ = v26
    case "\n\t":
        switch 9223372036854775807 {
        case 128:
            v27 := []int{32768, 4294967295, 42, 32768}
            _ = v27
        case -999:
            for fi28 := 0; fi28 < 20; fi28++ {
                switch (fi28 ^ 752) {
                case 42:
                    var v29 int = (fi28 + (fi28 ^ int(v5)))
                    _ = v29
                default:
                    if false {
                        v22 -= v14
                    } else {
                        var v30 uint64 = 818
                        _ = v30
                    }
                }
                _ = fi28
            }
        }
    default:
        v22 = v22
    }
    for fi31 := 0; fi31 < 6; fi31++ {
        var v32 bool = (true && false)
        _ = fi31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
