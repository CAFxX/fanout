package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 1
    v3 := make(map[string]uint32)
    switch "foo bar" {
    case "hello":
        v4 := []int{32767}
        _ = v4
    case "\n\t":
        fmt.Println(v2)
    }
    var v5 uint = 18446744073709551615
    v2 -= -875
    if (3.14159 >= -2.25) {
        var v6 string = "foo bar"
        if (v6 > (v6 + v6)) {
            switch -1 {
            case 32767:
                var v7 float32 = (float32(((((1e10 * 3.14159) / 1.0) - (-2.25 + -2.25)) / 3.14159)) * 1e10)
                _ = v7
            case 4294967295:
                v5 = 0
            }
        } else {
            var v8 float32 = 1e10
            if (((float64(3.14159) + -2.25) * float64(-2.25)) != (0.1 + 0.0)) {
                var v9 uint32 = 949
                var v10 uint = (v5 >> ((2) & 63))
                var v11 uint64 = 4294967295
                _ = v10
                _ = v11
                _ = v9
            }
            v6 = (((v6 + ("hello" + v6)) + v6) + "hello")
            _ = v8
        }
        _ = v6
    } else {
        var v12 uint64 = 177
        var v13 float64 = -2.25
        for fi14 := 0; fi14 < 2; fi14++ {
            delete(v1, 42)
            _ = fi14
        }
        _ = v12
        _ = v13
    }
    v15 := []int{}
    v5 += v5
    v2 -= v2
    v2 += (v2 / ((v2) | 1))
    v5 = v5
    v2 += v2
    v16 := v1[199]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v5
}
