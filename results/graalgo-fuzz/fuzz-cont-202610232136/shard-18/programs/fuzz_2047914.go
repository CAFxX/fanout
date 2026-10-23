package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
func genFn0(p0 bool) int {
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := []int{fi7, -2147483648, fi7}
        p0 = p0
        _ = fi7
        _ = v8
    }
    p0 = (0 <= 9223372036854775807)
    _ = p0
    return 0
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint32 = 1
    var v3 float32 = 0.0
    var v4 int64 = 351
    var v5 string = "foo bar"
    var v6 int32 = 42
    var v9 uint64 = uint64(v1)
    fmt.Println(v2)
    switch v5 {
    case "\n\t":
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := GenS1{F0: true, F1: v2}
            var v12 uint16 = 255
            _ = fi10
            _ = v11
            _ = v12
        }
    case "":
        v13 := GenS1{F0: ((false || false) || ((true && ((0.0 / -2.25) != 3.14159)) && (((0.1 / 1.0) - 0.1) != ((float64(((float32(float64(1.5)) / float32((1e10 + (0.1 + -2.25)))) + ((v3 * float32(-2.25)) / v3))) * (0.1 + ((0.1 + (3.14159 - ((1.5 * -2.25) / 1.0))) - (1e10 * 0.0)))) / (0.0 / ((-2.25 + float64(v3)) - 3.14159)))))), F1: v2}
        _ = v13
    case "hello":
        v4 = (^(v4))
    }
    fmt.Println(v4, v5)
    var v14 int8 = 1
    for fi15 := 0; fi15 < 12; fi15++ {
        v16 := []int{genFn0(((^(v4)) > 1)), (-(fi15)), fi15}
        v17 := GenS1{F0: false, F1: v1}
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        var v19 int64 = 42
        v20 := GenS1{F0: (((v9 * (v9 | (^(v9)))) <= 273) && (v6 == (v6 - (-(v6))))), F1: (v1 + v1)}
        _ = fi18
        _ = v19
        _ = v20
    }
    if true {
        v5 = v5
        switch genFn0((((v3 == 3.14159) && (false && (-1 == -27))) && true)) {
        case 398:
            switch "a" {
            case "\n\t":
                fmt.Println(v2, v4, v1)
            case "":
                v21 := []int{}
                _ = v21
            default:
                v22 := []int{}
                _ = v22
            }
        }
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        fmt.Println(v9)
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
