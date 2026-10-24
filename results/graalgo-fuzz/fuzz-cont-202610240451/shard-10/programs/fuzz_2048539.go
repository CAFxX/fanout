package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint32 }
type GenS2 struct { F0 int16; F1 int64; F2 int8 }
func genFn0(p0 string, p1 int8, p2 uint) uint32 {
    var v6 bool = false
    v7 := []int{128}
    v8 := GenS1{F0: (((float32(1.5) != -2.25) && ((-2.25 > float64((1e10 + 0.1))) && true)) || ("foo bar" <= p0)), F1: (-(p1)), F2: 629}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 2147483648
}

func main() {
    var v1 float32 = float32(float64(-2.25))
    v2 := make(map[string]int)
    var v3 float64 = (((0.1 - (-2.25 / 0.1)) - float64(((v1 / (0.1 + (-2.25 - -2.25))) - v1))) - (3.14159 - -2.25))
    var v4 int16 = 32767
    var v5 int64 = int64(v4)
    if true {
        if (!(!(!(19 > -49)))) {
            var v9 uint16 = 981
            _ = v9
        }
        v10 := GenS1{F0: (false && true), F1: 100, F2: 739}
        v11 := []int{128}
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: (!(!(!false))), F1: 0, F2: 2147483648}
    var v13 uint8 = 0
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 float64 = (((float64(((float32(v3) / v1) / v1)) / ((float64(float32((v3 * v3))) * v3) - -2.25)) + v3) - 3.14159)
        var v16 uint64 = 33
        var v17 uint8 = v13
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    var v18 int8 = -128
    switch ((("hello" + ("\n\t" + "foo bar")) + "a") + (("\n\t" + (("" + "a") + "")) + (("" + ((("\n\t" + "a") + "") + "\n\t")) + "\n\t"))) {
    case "foo bar":
        var v19 uint32 = 65535
        _ = v19
    default:
        var v20 float64 = (((v3 + 1.5) / v3) * ((v3 * v3) - (0.1 + 1.5)))
        _ = v20
    }
    if (v1 >= float32((v3 / v3))) {
        for fi21 := 0; fi21 < 5; fi21++ {
            for fi22 := 0; fi22 < 6; fi22++ {
                v23 := []int{(-(fi22)), (fi21 - (fi21 | (^(fi21)))), fi21, 0}
                _ = fi22
                _ = v23
            }
            v5 = (^(v5))
            _ = fi21
        }
    }
    v24 := GenS2{F0: (-(v4)), F1: (v5 * v5), F2: (v18 ^ v18)}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
