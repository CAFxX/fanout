package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int64, p1 int8) int16 {
    v6 := []int{}
    p1 += 42
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 string = "\n\t"
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        fmt.Println(fi9)
        _ = fi9
    }
    switch 2147483647 {
    case 4294967295:
        p1 = p1
    case -9223372036854775808:
        switch "\n\t" {
        case "foo bar":
            if len(v6) > 0 {
                v6[0] = 4294967295
            }
        }
    case -331:
        v10 := GenS1{F0: -1}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v6
    return 1000
}

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 489
    var v3 uint16 = (v2 | 136)
    var v4 uint16 = 660
    var v5 float64 = 1.5
    v11 := GenS2{F0: v2}
    var v12 float64 = ((-2.25 * v5) / ((((3.14159 - (v5 * float64(0.0))) + 0.1) * v5) / 0.1))
    for fi13 := 0; fi13 < 3; fi13++ {
        switch (fi13 ^ (fi13 * (fi13 + fi13))) {
        case 9223372036854775807:
            for fi14 := 0; fi14 < 10; fi14++ {
                var v15 float32 = -2.25
                var v16 int = (fi13 >> ((42) & 63))
                var v17 int32 = -1
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
        }
        if true {
            v18 := []int{}
            _ = v18
        } else {
            var v19 float64 = v12
            fmt.Println(fi13, v19, v2)
            v20 := []int{fi13, 127}
            _ = v19
            _ = v20
        }
        v5 = 0.1
        _ = fi13
    }
    fmt.Println(v3, v4, v5)
    v5 = (1.5 + 1e10)
    v3 = v4
    for fi21 := 0; fi21 < 10; fi21++ {
        for fi22 := 0; fi22 < 10; fi22++ {
            fmt.Println(v5)
            v1[(((("hello" + "") + "a") + "foo bar") + ("hello" + ""))] = 1
            fmt.Println(v4, fi21)
            _ = fi22
        }
        switch "a" {
        case "":
            v23 := []int{(fi21 * (-(fi21))), (fi21 ^ (fi21 - -980))}
            _ = v23
        default:
            v5 = (v5 * v5)
        }
        _ = fi21
    }
    v24 := []int{32768}
    var v25 int64 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
