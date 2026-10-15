package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
func genFn0(p0 uint32, p1 int32) int8 {
    if (true && (0 >= 32767)) {
        fmt.Println(p0)
    }
    switch "" {
    case "":
        p1 = 0
    case "foo bar":
        for fi4 := 0; fi4 < 14; fi4++ {
            var v5 bool = (4294967295 <= 1)
            _ = fi4
            _ = v5
        }
    }
    v6 := []int{}
    var v7 int16 = 1000
    if len(v6) > 0 {
        v6[0] = int(v7)
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 127
}

func main() {
    var v1 int16 = 1
    var v2 float64 = 3.14159
    v3 := make(map[string]int)
    if (9223372036854775807 == 1) {
        if (!true) {
            v2 -= (float64((1e10 + 3.14159)) + v2)
        }
    }
    delete(v3, "\n\t")
    v2 -= (v2 * float64(float32((1.5 * 0.1))))
    _ = len(v3)
    var v8 int = 2147483647
    var v9 int8 = -58
    var v10 uint = 255
    switch "foo bar" {
    case "":
        v11 := GenS1{F0: genFn0(0, -1), F1: v9}
        _ = v11
    }
    var v12 uint16 = 0
    for fi13 := 0; fi13 < 2; fi13++ {
        fmt.Println(v1, v10)
        switch int(v10) {
        case -366:
            var v14 string = "\n\t"
            _ = v14
        }
        for fi15 := 0; fi15 < 15; fi15++ {
            v16 := GenS1{F0: (v9 - (v9 / ((v9) | 1))), F1: genFn0(42, int32(v8))}
            _ = fi15
            _ = v16
        }
        _ = fi13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
