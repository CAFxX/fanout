package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 bool }
func genFn0(p0 int16) int {
    for fi7 := 0; fi7 < 4; fi7++ {
        switch (("" + "") + "\n\t") {
        case "foo bar":
            var v8 int16 = (p0 - 0)
            _ = v8
        default:
            var v9 string = ("a" + "a")
            _ = v9
        }
        for fi10 := 0; fi10 < 11; fi10++ {
            v11 := GenS1{F0: 1, F1: -1, F2: false}
            var v12 int64 = 42
            _ = fi10
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 10; fi13++ {
            v14 := []int{fi7}
            _ = fi13
            _ = v14
        }
        _ = fi7
    }
    var v15 uint16 = 115
    p0 -= p0
    var v16 uint16 = v15
    fmt.Println(p0, p0, v15)
    _ = p0
    _ = v15
    _ = v16
    return 32768
}

func main() {
    v1 := make(map[int]string)
    var v2 int = -1
    v3 := make(map[string]bool)
    var v4 int64 = 9223372036854775807
    var v5 float32 = (-2.25 - 0.0)
    var v6 uint64 = 42
    if true {
        v5 = v5
    }
    var v17 int64 = 42
    v18 := GenS1{F0: 151, F1: int16(v6), F2: (1 > 331)}
    var v19 int8 = -128
    v20 := GenS1{F0: 0, F1: -458, F2: (v6 < (^(v6)))}
    delete(v3, "a")
    v17 -= v17
    v4 += 9223372036854775807
    v21 := []int{v2, (v2 << ((18446744073709551615) & 63)), v2}
    v17 += v17
    v22 := GenS1{F0: 38, F1: 1, F2: (true || (false && ((v17 >> ((65535) & 63)) != v4)))}
    var v23 uint64 = v6
    fmt.Println(v17, v6)
    v6 -= v6
    switch (v2 & 9223372036854775807) {
    case -375:
        for fi24 := 0; fi24 < 15; fi24++ {
            v21 = append(v21, v2)
            _ = fi24
        }
    }
    fmt.Println(v19, v5, v5)
    var v25 uint16 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
