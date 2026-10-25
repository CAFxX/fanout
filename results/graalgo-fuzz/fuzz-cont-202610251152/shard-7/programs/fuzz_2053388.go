package main

import (
    "fmt"
)

func genFn0(p0 uint16, p1 int64, p2 uint64) int {
    p2 -= (p2 & p2)
    v6 := []int{}
    p1 -= -240
    var v7 uint8 = 134
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 128
}

func main() {
    var v1 float64 = 1e10
    var v2 int = 32767
    v3 := []string{"a"}
    var v4 float64 = v1
    v5 := make(map[string]bool)
    fmt.Println(v1)
    var v8 float32 = 1.5
    v4 += (v1 + 0.0)
    fmt.Println(v4, v8)
    v1 = 1.5
    v5["\n\t"] = (0 < 2)
    var v9 uint = 2
    if len(v3) > 0 {
        v3[0] = "a"
    }
    v4 = 0.1
    var v10 float32 = (v8 * 1.5)
    var v11 int = genFn0(0, 433, 2)
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 8; fi13++ {
            fmt.Println(v8)
            switch 10 {
            case -390:
                v14 := v5[("foo bar" + ((("foo bar" + "hello") + ("foo bar" + "")) + "hello"))]
                _ = v14
            default:
                var v15 uint64 = 235
                _ = v15
            }
            _ = fi13
        }
        v9 = 265
        _ = fi12
    }
    v9 = 42
    v11 = v11
    if true {
        if (!false) {
            for fi16 := 0; fi16 < 9; fi16++ {
                var v17 float64 = v1
                _ = fi16
                _ = v17
            }
        } else {
            var v18 bool = (!(!false))
            v5[""] = true
            _ = v18
        }
        v5[("foo bar" + (("foo bar" + "\n\t") + (("a" + ("" + "foo bar")) + ("hello" + "a"))))] = (false && (false && (786 < -1)))
    }
    switch "hello" {
    case "\n\t":
        v1 += (float64(0.0) + ((float64((1e10 + ((((v10 * v8) * (1e10 * ((-2.25 * (3.14159 / (float32(0.1) * -2.25))) / 1.5))) + (v10 + ((0.1 - v10) / -2.25))) + float32(float64((1e10 * 0.0)))))) / v1) * -2.25))
    case "":
        v11 += 0
    default:
        if ((!false) && true) {
            v9 = 290
            fmt.Println(v11, v9)
        } else {
            v3 = append(v3, "hello")
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    var v19 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
