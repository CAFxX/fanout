package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int16; F1 int8 }
func genFn0(p0 string, p1 uint16) int16 {
    p1 = (p1 >> ((945) & 63))
    v7 := []int{323}
    _ = p0
    _ = p1
    _ = v7
    return -1
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[string]uint32)
    var v3 int = 0
    var v4 float64 = (1e10 * -2.25)
    var v5 uint16 = (v1 * 415)
    var v6 int8 = 100
    if true {
        var v8 bool = true
        var v9 bool = (!false)
        v2["hello"] = 716
        _ = v8
        _ = v9
    } else {
        if (-1 != 856) {
            if (false && (!(355 <= 359))) {
                var v10 uint32 = 65535
                if (0 >= 77) {
                    v6 += v6
                    switch "a" {
                    case "\n\t":
                        v11 := GenS1{F0: 127, F1: 60}
                        _ = v11
                    default:
                        var v12 bool = (!false)
                        _ = v12
                    }
                }
                v10 = (v10 >> ((255) & 63))
                _ = v10
            }
            var v13 int64 = 0
            _ = v13
        }
    }
    var v14 uint16 = (v5 % ((v5) | 1))
    v1 = v14
    fmt.Println(v14, v14, v14)
    var v15 bool = false
    if (v15 && (42 == 531)) {
        v5 = 876
    }
    if ((!false) || (v4 >= ((0.1 * v4) + (v4 * v4)))) {
        v16 := []int{878, v3, -9223372036854775808}
        v17 := GenS1{F0: 1, F1: 100}
        _ = v16
        _ = v17
    }
    var v18 float32 = 0.0
    if (326 > 777) {
        var v19 int = v3
        for fi20 := 0; fi20 < 9; fi20++ {
            v21 := []int{}
            _ = fi20
            _ = v21
        }
        for fi22 := 0; fi22 < 15; fi22++ {
            v4 = v4
            delete(v2, "a")
            v23 := GenS2{F0: genFn0(("" + "a"), (^(v5))), F1: (v6 << ((694) & 63))}
            _ = fi22
            _ = v23
        }
        _ = v19
    }
    var v24 float64 = v4
    _ = len(v2)
    var v25 bool = v15
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
