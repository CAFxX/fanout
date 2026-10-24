package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int16 }
func genFn0(p0 uint64, p1 uint) int64 {
    var v6 uint16 = 992
    if (true && false) {
        fmt.Println(p0, v6, v6)
        if (!false) {
            var v7 uint16 = (v6 - v6)
            _ = v7
        } else {
            fmt.Println(v6, v6, v6)
            var v8 uint64 = (p0 >> (((p1 & (p1 - 1))) & 63))
            var v9 string = "\n\t"
            _ = v8
            _ = v9
        }
        var v10 uint32 = 722
        _ = v10
    }
    if (48 > 54) {
        var v11 uint64 = uint64(v6)
        _ = v11
    } else {
        var v12 uint16 = v6
        _ = v12
    }
    p1 += 18446744073709551615
    var v13 int16 = 42
    _ = p0
    _ = p1
    _ = v13
    _ = v6
    return 0
}

func main() {
    var v1 int64 = -682
    var v2 uint = 10
    var v3 uint = 18446744073709551615
    v4 := make(map[uint32]bool)
    var v5 uint32 = 65535
    var v14 bool = false
    var v15 uint64 = uint64(v5)
    var v16 float32 = float32(1e10)
    var v17 uint16 = 0
    v18 := []int{1}
    var v19 string = ("a" + (("hello" + "a") + ("\n\t" + ("a" + ("a" + "a")))))
    if (56 <= 1) {
        var v20 string = v19
        fmt.Println(v20)
        v21 := GenS1{F0: 336, F1: 4294967295, F2: 32767}
        _ = v20
        _ = v21
    }
    var v22 bool = (0 == 42)
    var v23 uint16 = v17
    for fi24 := 0; fi24 < 19; fi24++ {
        var v25 int32 = 913
        switch (fi24 - fi24) {
        case 4294967295:
            switch (v19 + v19) {
            case "\n\t":
                v1 += (v1 + v1)
            case "hello":
                delete(v4, (^(v5)))
            }
        case 32768:
            var v26 uint8 = 2
            _ = v26
        case 9223372036854775807:
            v4[(v5 % ((v5) | 1))] = v22
        default:
            v27 := GenS1{F0: v5, F1: fi24, F2: -1}
            _ = v27
        }
        v15 += 453
        _ = fi24
        _ = v25
    }
    v1 -= genFn0(35, 42)
    _ = len(v18)
    var v28 int32 = -709
    switch v19 {
    case "\n\t":
        if (18446744073709551615 != (v15 + v15)) {
            var v29 bool = v22
            _ = v29
        }
    case "":
        delete(v4, (v5 ^ 277))
    default:
        var v30 float64 = (((0.1 - 3.14159) * ((-2.25 + float64(float32(3.14159))) + 3.14159)) / (1.5 / (1.5 + 3.14159)))
        _ = v30
    }
    v31 := GenS1{F0: (v5 / ((v5) | 1)), F1: 0, F2: 540}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
