package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 int64; F2 int }
func genFn0(p0 int16) uint {
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 string = ("\n\t" + "hello")
        for fi9 := 0; fi9 < 1; fi9++ {
            v10 := []int{9223372036854775807}
            v11 := GenS1{F0: 65535}
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    if (true || false) {
        v12 := GenS1{F0: 643}
        _ = v12
    }
    _ = p0
    return 0
}

func main() {
    var v1 uint = 65535
    v2 := make(map[int]bool)
    var v3 uint16 = 0
    var v4 bool = false
    v5 := make(map[int]int)
    var v6 uint8 = 42
    var v13 int8 = 126
    v3 += 255
    if false {
        for fi14 := 0; fi14 < 3; fi14++ {
            v15 := []int{fi14}
            _ = fi14
            _ = v15
        }
        if (v3 > 0) {
            v16 := GenS2{F0: 9223372036854775807, F1: 42, F2: -1}
            var v17 uint64 = uint64(v1)
            _ = v16
            _ = v17
        }
    } else {
        if (-1 >= 42) {
            var v18 int16 = -1
            fmt.Println(v1)
            _ = v18
        } else {
            switch 128 {
            case 32768:
                var v19 uint16 = 631
                _ = v19
            case -789:
                v20 := []int{127, 695}
                _ = v20
            default:
                v21 := GenS2{F0: -1, F1: 1, F2: -524}
                _ = v21
            }
            v1 -= 1
        }
    }
    v13 += -128
    switch "a" {
    case "hello":
        switch "foo bar" {
        case "\n\t":
            v1 = uint(v3)
        case "hello":
            for fi22 := 0; fi22 < 7; fi22++ {
                fmt.Println(v4, fi22)
                _ = fi22
            }
        case "foo bar":
            fmt.Println(v3, v4)
        default:
            switch ("" + "a") {
            case "a":
                for fi23 := 0; fi23 < 6; fi23++ {
                    v6 = (v6 * (v6 % (((v6 & v6)) | 1)))
                    v24 := GenS2{F0: -740, F1: 9223372036854775807, F2: -9223372036854775808}
                    _ = fi23
                    _ = v24
                }
            }
        }
    case "a":
        v13 = v13
    default:
        v4 = v4
    }
    var v25 bool = (513 <= 0)
    var v26 float64 = -2.25
    v4 = (-2147483648 <= -2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
