package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int16; F2 int32 }
func genFn0() int {
    switch 151 {
    case 9223372036854775807:
        var v7 int = -9223372036854775808
        _ = v7
    default:
        switch 958 {
        case 42:
            v8 := []int{0}
            _ = v8
        case 128:
            v9 := []int{}
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            if false {
                var v12 float64 = 1.5
                _ = v12
            }
            switch fi10 {
            case -2147483648:
                var v13 int32 = 1
                _ = v13
            default:
                if false {
                    v14 := GenS1{F0: -114, F1: 0, F2: 2}
                    var v15 bool = (false || false)
                    fmt.Println(v15)
                    _ = v14
                    _ = v15
                }
            }
            _ = fi11
        }
        _ = fi10
    }
    var v16 string = ((("" + "a") + ("foo bar" + ("hello" + (("hello" + "hello") + ("\n\t" + "a"))))) + "hello")
    _ = v16
    return -44
}

func main() {
    var v1 int = -1
    var v2 float64 = 0.0
    v3 := make(map[string]bool)
    var v4 bool = (-566 < 32767)
    var v5 uint8 = 19
    var v6 string = (("hello" + "foo bar") + "foo bar")
    v17 := GenS1{F0: 1, F1: 0, F2: v5}
    v18 := GenS2{F0: 442, F1: 32767, F2: -2147483648}
    var v19 uint16 = 401
    v20 := GenS1{F0: 42, F1: 0, F2: (v5 + 255)}
    fmt.Println(v6)
    delete(v3, "foo bar")
    var v21 string = ("hello" + (v6 + v6))
    v2 += (v2 + v2)
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 int16 = 42
        _ = fi22
        _ = v23
    }
    v2 -= float64(((-2.25 / 1e10) + 0.1))
    fmt.Println(v2, v1, v6)
    v5 += v5
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
