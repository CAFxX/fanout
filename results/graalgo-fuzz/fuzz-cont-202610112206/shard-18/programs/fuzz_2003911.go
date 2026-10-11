package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }

func main() {
    v1 := []string{"a", ((("a" + ("foo bar" + "\n\t")) + "") + "\n\t")}
    var v2 uint32 = 909
    var v3 float64 = 3.14159
    var v4 bool = false
    var v5 int16 = -1
    for fi6 := 0; fi6 < 3; fi6++ {
        fmt.Println(v2, v2)
        v7 := GenS1{F0: 42, F1: 1}
        if ("" > "a") {
            if true {
                var v8 float64 = ((v3 + -2.25) + v3)
                if true {
                    var v9 int8 = -20
                    _ = v9
                }
                _ = v8
            }
        } else {
            if ((v4 || v4) && (true || (!v4))) {
                v3 -= v3
                for fi10 := 0; fi10 < 10; fi10++ {
                    v1 = append(v1, "a")
                    v2 -= (v2 % (((v2 | (v2 >> ((255) & 63)))) | 1))
                    _ = fi10
                }
            } else {
                _ = len(v1)
            }
        }
        _ = fi6
        _ = v7
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        fmt.Println(v3, v2, v5)
        var v12 bool = v4
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        v4 = (158 >= -9223372036854775808)
        _ = fi13
    }
    var v14 float64 = (0.1 - (0.0 - v3))
    v2 += (v2 ^ (v2 >> ((42) & 63)))
    switch 377 {
    case 570:
        v14 -= (((v3 * (float64(0.1) * (v3 + float64(1.5)))) - (float64(-2.25) + (v3 - ((((v3 - 0.0) * float64((1e10 - 1.5))) + ((-2.25 + ((v14 - ((v14 + 1e10) * -2.25)) + v3)) + (((float64((float32(3.14159) * ((1e10 - -2.25) - (float32(3.14159) * 1.5)))) / 0.1) / 1.0) + (1e10 + v14)))) * 0.0)))) - v3)
    case 0:
        var v15 bool = (602 < 32768)
        _ = v15
    default:
        var v16 string = "hello"
        _ = v16
    }
    if false {
        v2 += v2
        var v17 uint64 = 0
        v4 = ((0 <= 255) || v4)
        _ = v17
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        v2 += v2
        v2 = (v2 & v2)
        _ = fi18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
