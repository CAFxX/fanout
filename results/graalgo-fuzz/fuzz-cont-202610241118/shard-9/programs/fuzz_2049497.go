package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint }
func genFn0() int64 {
    var v6 uint16 = 1
    fmt.Println(v6)
    if true {
        for fi7 := 0; fi7 < 10; fi7++ {
            switch "hello" {
            case "a":
                var v8 uint32 = 4294967295
                _ = v8
            case "hello":
                for fi9 := 0; fi9 < 2; fi9++ {
                    fmt.Println(v6)
                    _ = fi9
                }
            default:
                var v10 bool = true
                _ = v10
            }
            _ = fi7
        }
    } else {
        fmt.Println(v6)
        var v11 int = 240
        for fi12 := 0; fi12 < 5; fi12++ {
            v11 = fi12
            _ = fi12
        }
        _ = v11
    }
    v6 += 32767
    _ = v6
    return 394
}

func main() {
    var v1 int8 = -1
    var v2 bool = (true || (42 == 21))
    var v3 float32 = (0.1 - (1e10 + -2.25))
    var v4 int8 = v1
    var v5 int8 = (v1 + (^(v1)))
    fmt.Println(v5, v1, v2)
    v13 := GenS1{F0: 2, F1: -9223372036854775808}
    var v14 string = "\n\t"
    for fi15 := 0; fi15 < 14; fi15++ {
        v16 := []int{(fi15 << ((2) & 63)), 9223372036854775807}
        for fi17 := 0; fi17 < 7; fi17++ {
            if (-989 < -1) {
                v5 -= -128
            }
            _ = fi17
        }
        if (((3.14159 / 1e10) - float64((v3 + -2.25))) > -2.25) {
            fmt.Println(v3)
            v4 -= v5
            _ = len(v16)
        }
        _ = fi15
        _ = v16
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        v19 := []int{fi18, -257}
        switch v14 {
        case "a":
            var v20 int64 = 0
            _ = v20
        case "":
            fmt.Println(v3, v5, fi18)
        case "\n\t":
            v21 := GenS2{F0: 500, F1: 255}
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    fmt.Println(v2)
    var v22 string = (((v14 + (("a" + "foo bar") + "foo bar")) + v14) + v14)
    v23 := GenS2{F0: 2, F1: 823}
    fmt.Println("v1:", v1)
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
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
