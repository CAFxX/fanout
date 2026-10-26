package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int8) uint8 {
    v6 := GenS1{F0: 138}
    var v7 float64 = 3.14159
    p0 = -35
    _ = p0
    _ = v6
    _ = v7
    return 128
}

func main() {
    v1 := make(map[int]string)
    var v2 int32 = 141
    v3 := make(map[string]bool)
    var v4 int16 = 0
    var v5 int8 = 0
    v4 = (v4 | (-(v4)))
    switch "foo bar" {
    case "\n\t":
        var v8 float64 = float64(-2.25)
        _ = v8
    case "":
        if true {
            if ((v5 + (v5 ^ (v5 / ((v5) | 1)))) >= v5) {
                for fi9 := 0; fi9 < 14; fi9++ {
                    v4 = (v4 ^ (v4 >> ((280) & 63)))
                    v5 = v5
                    _ = fi9
                }
                for fi10 := 0; fi10 < 20; fi10++ {
                    var v11 int8 = v5
                    _ = fi10
                    _ = v11
                }
                v2 = 930
            }
        }
    case "a":
        v12 := GenS1{F0: 255}
        _ = v12
    }
    delete(v3, "hello")
    var v13 float32 = ((1e10 - 3.14159) + 3.14159)
    v4 = (v4 - v4)
    v5 = v5
    v4 = 766
    v5 += -44
    delete(v3, ("\n\t" + "hello"))
    var v14 int8 = v5
    if (float64(-2.25) > (1e10 - (1.5 - 1e10))) {
        if false {
            var v15 uint64 = 2
            v14 = 127
            _ = v15
        } else {
            v16 := []int{0, -1}
            _ = v16
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
