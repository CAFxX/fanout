package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint16 }
func genFn0(p0 int8, p1 string) uint {
    fmt.Println(p1)
    p1 = (p1 + "hello")
    _ = p0
    _ = p1
    return 65535
}

func main() {
    var v1 string = "\n\t"
    var v2 int32 = -1
    var v3 float64 = 1.5
    v4 := []int{-940, 753, -2}
    for fi5 := 0; fi5 < 19; fi5++ {
        v1 = v1
        v2 -= -560
        v1 = ""
        _ = fi5
    }
    v1 = "a"
    if false {
        v3 += ((v3 * (-2.25 + ((((v3 - v3) + 1e10) / float64(0.1)) / ((v3 - (((((v3 - v3) / (((1e10 * ((v3 / 0.1) + v3)) + v3) - -2.25)) - v3) * v3) + float64(0.0))) + -2.25)))) - 0.0)
    }
    v2 = v2
    fmt.Println(v2)
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 int16 = -32768
        _ = fi6
        _ = v7
    }
    if false {
        if (!true) {
            var v8 int64 = 0
            if true {
                fmt.Println(v2, v3, v2)
                for fi9 := 0; fi9 < 16; fi9++ {
                    var v10 int32 = (^(v2))
                    _ = fi9
                    _ = v10
                }
            }
            var v11 uint16 = 751
            _ = v11
            _ = v8
        }
        var v12 string = ((("" + (((v1 + ("a" + "")) + "") + "\n\t")) + v1) + v1)
        fmt.Println(v3)
        _ = v12
    }
    fmt.Println(v1, v2)
    var v13 bool = true
    v14 := GenS1{F0: 0, F1: 492, F2: 434}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
