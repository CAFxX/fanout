package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
func genFn0() string {
    v4 := GenS1{F0: 2, F1: "\n\t"}
    var v5 int8 = 42
    v5 += (v5 * -128)
    _ = v4
    _ = v5
    return "foo bar"
}

func main() {
    var v1 float32 = (float32((3.14159 * 3.14159)) * float32(((0.1 - 3.14159) * float64(float32(float64((((1.5 - 3.14159) / 3.14159) - (-2.25 + 0.0))))))))
    var v2 int8 = -20
    var v3 string = "hello"
    var v6 bool = ((!false) || true)
    fmt.Println(v1, v2, v1)
    var v7 float32 = 0.1
    switch "" {
    case "\n\t":
        for fi8 := 0; fi8 < 3; fi8++ {
            v7 = v7
            var v9 bool = v6
            _ = fi8
            _ = v9
        }
    }
    v3 = v3
    if v6 {
        var v10 string = "foo bar"
        if true {
            for fi11 := 0; fi11 < 9; fi11++ {
                v1 = v7
                if false {
                    v12 := []int{fi11, (fi11 * 2147483647)}
                    _ = v12
                }
                switch fi11 {
                case 127:
                    v6 = (!v6)
                case 32768:
                    for fi13 := 0; fi13 < 13; fi13++ {
                        fmt.Println(v7, v2, v1)
                        fmt.Println(v7, v6)
                        v2 -= 0
                        v14 := []int{fi11}
                        _ = fi13
                        _ = v14
                    }
                case -998:
                    v10 = ""
                default:
                    var v15 uint64 = 42
                    _ = v15
                }
                _ = fi11
            }
            v6 = ("" <= ("a" + v10))
        }
        _ = v10
    }
    v16 := GenS1{F0: 42, F1: "foo bar"}
    v17 := GenS1{F0: 18446744073709551615, F1: v3}
    var v18 float64 = (0.0 + 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
