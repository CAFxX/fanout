package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint8 }
func genFn0(p0 int16, p1 uint, p2 int) int8 {
    v6 := GenS1{F0: true, F1: "", F2: 255}
    var v7 uint16 = uint16(p1)
    if true {
        switch (p2 | 9223372036854775807) {
        case 722:
            if (0 < -2147483648) {
                var v8 float64 = 1e10
                _ = v8
            }
        case 1:
            v9 := []int{p2, p2, (^(p2))}
            _ = v9
        case 446:
            fmt.Println(p2, p0)
        default:
            var v10 float32 = 3.14159
            _ = v10
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := []string{"a", ((("\n\t" + (("foo bar" + "a") + "")) + "a") + ("a" + "\n\t")), ("" + ("\n\t" + "\n\t")), ""}
    var v2 float32 = (0.1 - 1.5)
    v3 := make(map[int]string)
    v4 := make(map[string]uint32)
    var v5 int16 = 42
    var v11 bool = (4294967295 < 415)
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 uint = 65535
        v13 -= v13
        var v14 float32 = v2
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := GenS1{F0: true, F1: (("hello" + "\n\t") + "\n\t"), F2: 2}
        fmt.Println(v5, fi15)
        v17 := GenS1{F0: (v11 && (v2 >= (((0.0 * 0.1) + ((0.0 - 1.5) + (1.5 - 0.0))) * 0.1))), F1: "foo bar", F2: 2}
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        v19 := GenS1{F0: false, F1: "", F2: 255}
        _ = fi18
        _ = v19
    }
    var v20 uint64 = 18446744073709551615
    v21 := GenS1{F0: (0 != 0), F1: "\n\t", F2: uint8(v20)}
    var v22 uint64 = 4294967295
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
