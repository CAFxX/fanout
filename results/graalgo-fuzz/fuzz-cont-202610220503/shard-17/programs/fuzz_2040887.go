package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 uint8 = 198
    var v2 uint64 = 582
    v3 := make(map[uint32]string)
    v4 := make(map[string]uint32)
    v5 := []int{-2147483648}
    if (255 >= 2) {
        v1 = 42
        var v6 int16 = 1000
        var v7 int16 = v6
        _ = v6
        _ = v7
    }
    if (!true) {
        if (1.5 != 1e10) {
            var v8 int16 = 401
            _ = v8
        }
    }
    fmt.Println(v2)
    for fi9 := 0; fi9 < 18; fi9++ {
        if true {
            v2 = v2
            var v10 float32 = ((float32((0.0 / 1.0)) * 0.0) / (1.5 * (-2.25 + (-2.25 + ((1e10 - 0.0) * ((0.1 * 1.5) + 0.1))))))
            _ = v10
        }
        v1 = v1
        var v11 uint = 18446744073709551615
        _ = fi9
        _ = v11
    }
    var v12 uint32 = 2147483648
    var v13 int32 = -2147483648
    v2 = v2
    if ((v13 << ((1) & 63)) < 798) {
        v14 := GenS1{F0: (v2 | 341)}
        switch 566 {
        case 46:
            var v15 uint64 = 18446744073709551615
            _ = v15
        case 185:
            if true {
                var v16 uint32 = 1
                _ = v16
            } else {
                v5 = append(v5, 127)
                var v17 int8 = 100
                v3[v12] = "\n\t"
                _ = v17
            }
        }
        _ = v14
    } else {
        var v18 string = "foo bar"
        _ = v18
    }
    if len(v5) > 0 {
        v5[0] = 4294967295
    }
    var v19 uint32 = uint32(v2)
    if true {
        for fi20 := 0; fi20 < 12; fi20++ {
            var v21 string = ("foo bar" + "foo bar")
            _ = fi20
            _ = v21
        }
        v22 := GenS1{F0: 405}
        _ = v22
    } else {
        switch 410 {
        case 42:
            var v23 float64 = float64(((0.1 * 1.5) / (1.5 - (0.0 * (0.1 + 1e10)))))
            _ = v23
        case 127:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case 9223372036854775807:
            if len(v5) > 0 {
                v5[0] = -489
            }
        }
        var v24 uint = 815
        v4[""] = uint32(v1)
        _ = v24
    }
    fmt.Println(v12)
    var v25 uint = 18446744073709551615
    delete(v4, "a")
    v1 = 227
    var v26 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
