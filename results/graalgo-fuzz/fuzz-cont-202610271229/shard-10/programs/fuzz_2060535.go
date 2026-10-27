package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
func genFn0(p0 bool, p1 int16, p2 int64) int {
    p2 += -44
    fmt.Println(p0, p1)
    v5 := GenS1{F0: "", F1: (p2 ^ p2)}
    switch "\n\t" {
    case "":
        v6 := GenS1{F0: "", F1: 42}
        _ = v6
    case "foo bar":
        fmt.Println(p2, p2)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}

func main() {
    var v1 uint8 = 1
    var v2 uint16 = 734
    var v3 int64 = 747
    var v4 int32 = -411
    if (0.0 >= 3.14159) {
        if true {
            fmt.Println(v2, v4)
            var v7 uint = 0
            _ = v7
        } else {
            var v8 string = "hello"
            _ = v8
        }
    }
    v9 := []int{128, -529}
    v3 = (v3 ^ v3)
    fmt.Println(v1)
    var v10 int8 = 30
    v11 := GenS1{F0: "foo bar", F1: 840}
    v2 -= v2
    fmt.Println(v10)
    var v12 string = "a"
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 uint16 = (v2 | v2)
        _ = fi13
        _ = v14
    }
    var v15 int = 524
    v1 = v1
    if (42 <= (v1 | (v1 - v1))) {
        var v16 string = v12
        v2 = v2
        _ = v16
    } else {
        if false {
            for fi17 := 0; fi17 < 5; fi17++ {
                fmt.Println(v15, v4)
                for fi18 := 0; fi18 < 14; fi18++ {
                    var v19 int8 = (v10 | (v10 * v10))
                    var v20 int16 = -32768
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                var v21 int = (v15 / ((v15) | 1))
                _ = fi17
                _ = v21
            }
            var v22 int8 = v10
            v3 += (v3 >> ((255) & 63))
            _ = v22
        }
        for fi23 := 0; fi23 < 19; fi23++ {
            v1 = v1
            var v24 float64 = -2.25
            _ = fi23
            _ = v24
        }
    }
    var v25 uint16 = v2
    if false {
        if len(v9) > 0 {
            v9[0] = v15
        }
        v15 = genFn0((58 < 127), 1000, v3)
        if len(v9) > 0 {
            v9[0] = (v15 >> ((255) & 63))
        }
    } else {
        v9 = append(v9, (v15 | (v15 & v15)))
    }
    var v26 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v9
}
