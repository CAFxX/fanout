package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int64, p1 int64, p2 int16) int8 {
    v7 := GenS1{F0: 127}
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 127
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []int{32768, -933, 0}
    v3 := make(map[int]int)
    var v4 bool = false
    v5 := make(map[uint32]bool)
    var v6 int64 = -1
    if ((0.0 / 1e10) != (((0.0 / (((1e10 - float32(3.14159)) - 0.1) - (float32(((3.14159 * 1.5) / 3.14159)) + ((1e10 * -2.25) + 0.0)))) / (0.1 / ((3.14159 - 3.14159) + float32((((((1.5 / -2.25) / 1.0) / (1e10 / ((float64(3.14159) + 1e10) + float64((1e10 * float32(-2.25)))))) + 3.14159) + float64(0.0)))))) + (0.0 - float32(1.5)))) {
        var v8 int32 = -2147483648
        _ = v8
    } else {
        var v9 float32 = (0.0 + float32(1.5))
        v4 = true
        _ = v9
    }
    var v10 string = "hello"
    _ = len(v2)
    var v11 uint64 = 653
    v5[1] = false
    var v12 int8 = -52
    for fi13 := 0; fi13 < 14; fi13++ {
        if false {
            for fi14 := 0; fi14 < 14; fi14++ {
                v11 = (v11 * (v11 >> ((110) & 63)))
                v12 = v12
                var v15 uint = 2
                _ = fi14
                _ = v15
            }
        }
        v4 = (!((fi13 % (((fi13 % ((fi13) | 1))) | 1)) > (-(fi13))))
        if false {
            v16 := GenS1{F0: fi13}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi17 := 0; fi17 < 14; fi17++ {
                switch (v10 + (("a" + v10) + v10)) {
                case "foo bar":
                    var v18 uint16 = 479
                    _ = v18
                }
                fmt.Println(v1, v10, v4)
                _ = fi17
            }
            _ = v16
        }
        _ = fi13
    }
    var v19 float64 = (-2.25 * (0.0 / -2.25))
    fmt.Println(v1, v12, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
