package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint, p1 int64) string {
    v5 := GenS1{F0: (^(p1))}
    p0 = (p0 & 42)
    _ = p0
    _ = p1
    _ = v5
    return ("foo bar" + "\n\t")
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = ((true && (true || (!((!(!false)) && false)))) || (-1 == 138))
    var v3 uint32 = 2
    var v4 int8 = 1
    v2 = (!(1e10 > 3.14159))
    v6 := []int{32767, int(v3)}
    var v7 int8 = v4
    v4 = -1
    var v8 float32 = ((((0.1 + (3.14159 - 3.14159)) + 1e10) / -2.25) / (-2.25 - 1e10))
    v6 = append(v6, -892)
    for fi9 := 0; fi9 < 20; fi9++ {
        fmt.Println(v3, v8, fi9)
        _ = fi9
    }
    v8 = v8
    v7 -= 127
    if (2 != 590) {
        if v2 {
            v4 = (v4 % (((v7 >> ((460) & 63))) | 1))
        } else {
            fmt.Println(v4, v3, v2)
        }
        v8 = float32((float64((-2.25 / v8)) / (float64(v8) / 1.5)))
        var v10 uint16 = 255
        _ = v10
    }
    v3 = (v3 - (v3 & v3))
    var v11 float32 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
