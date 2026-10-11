package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
func genFn0(p0 string, p1 uint8, p2 uint8) uint8 {
    v6 := []int{701, 1, -424, 721}
    switch 1 {
    case 128:
        fmt.Println(p2)
    default:
        switch "\n\t" {
        case "foo bar":
            fmt.Println(p1)
        }
    }
    var v7 int64 = 9223372036854775807
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p1
}

func main() {
    var v1 float64 = ((3.14159 / (0.1 / 0.1)) * (-2.25 / 1e10))
    v2 := make(map[string]bool)
    v3 := make(map[int]uint32)
    v4 := []string{"\n\t", "hello", "foo bar"}
    v5 := make(map[string]string)
    var v8 int32 = -1
    var v9 uint = 42
    delete(v3, -366)
    v10 := v5["a"]
    v8 = (-(v8))
    if false {
        v8 -= 1
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        v8 = (v8 / ((v8) | 1))
        _ = fi11
    }
    fmt.Println(v8)
    var v12 int16 = 1000
    v8 = 2147483647
    var v13 float64 = (1.5 - (((1e10 * (1e10 - v1)) + v1) + (v1 * (0.0 - ((1.5 + (v1 - float64((-2.25 * -2.25)))) - ((v1 - (v1 + (v1 / (v1 * -2.25)))) + 1.5))))))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v14 int64 = 1
    for fi15 := 0; fi15 < 18; fi15++ {
        v5[v10] = ((v10 + "\n\t") + "foo bar")
        if false {
            fmt.Println(v14, fi15)
        } else {
            var v16 int64 = 9223372036854775807
            v17 := GenS1{F0: (v16 + -1), F1: 556}
            _ = v16
            _ = v17
        }
        _ = fi15
    }
    if (!false) {
        switch (("a" + ((v10 + "a") + v10)) + "\n\t") {
        case "":
            var v18 uint32 = 761
            _ = v18
        }
    } else {
        switch -889 {
        case 181:
            v19 := GenS1{F0: (v14 % ((0) | 1)), F1: 630}
            _ = v19
        default:
            for fi20 := 0; fi20 < 19; fi20++ {
                v13 = (float64(float32((((float64(0.0) + (v13 + (v1 / (3.14159 - ((((v1 + 3.14159) * (float64(3.14159) + v1)) - float64((float32(0.0) + 1.5))) + 3.14159))))) - ((1e10 - float64(0.1)) - float64(0.0))) * v13))) * v1)
                _ = fi20
            }
        }
        var v21 int16 = (v12 * (v12 << ((v9) & 63)))
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
