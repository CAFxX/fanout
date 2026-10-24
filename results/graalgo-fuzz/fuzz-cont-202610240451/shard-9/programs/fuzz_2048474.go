package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int32 }
func genFn0(p0 int8, p1 uint64, p2 string) uint16 {
    switch p2 {
    case "foo bar":
        for fi5 := 0; fi5 < 2; fi5++ {
            fmt.Println(p2)
            _ = fi5
        }
    case "hello":
        if (false && false) {
            v6 := []int{-1, -228}
            for fi7 := 0; fi7 < 1; fi7++ {
                var v8 int32 = 42
                var v9 int64 = 42
                _ = fi7
                _ = v8
                _ = v9
            }
            _ = v6
        } else {
            for fi10 := 0; fi10 < 3; fi10++ {
                var v11 int32 = -2147483648
                var v12 int16 = 0
                _ = fi10
                _ = v11
                _ = v12
            }
            if (!(false || (!(42 > 0)))) {
                switch "" {
                case "foo bar":
                    p0 = p0
                }
                p2 = "hello"
                v13 := []int{2147483647, 32768, 198, 401}
                _ = v13
            }
            var v14 uint8 = 255
            _ = v14
        }
    default:
        var v15 uint64 = 0
        _ = v15
    }
    var v16 uint32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 42
}

func main() {
    var v1 int64 = 1
    var v2 uint32 = 2147483647
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]int)
    if (!false) {
        v2 = (^(v2))
        v17 := GenS1{F0: 65535, F1: 2, F2: 1}
        v18 := v3[v2]
        _ = v17
        _ = v18
    } else {
        if false {
            var v19 uint32 = 631
            _ = v19
        }
        fmt.Println(v2)
    }
    delete(v4, v2)
    v2 -= v2
    for fi20 := 0; fi20 < 8; fi20++ {
        v21 := GenS1{F0: 2, F1: 255, F2: -2147483648}
        v22 := GenS1{F0: genFn0(-1, 1, "hello"), F1: 0, F2: 0}
        _ = fi20
        _ = v21
        _ = v22
    }
    v2 = v2
    for fi23 := 0; fi23 < 20; fi23++ {
        v2 = 65535
        var v24 bool = false
        fmt.Println(fi23, v24)
        _ = fi23
        _ = v24
    }
    v25 := []int{int(v2), 127, 2147483647}
    switch "" {
    case "foo bar":
        var v26 uint32 = v2
        _ = v26
    }
    v2 += 764
    v2 += (v2 + (v2 | v2))
    var v27 int8 = 6
    var v28 int32 = 0
    switch (((("foo bar" + ("a" + "a")) + "\n\t") + "foo bar") + "") {
    case "\n\t":
        fmt.Println(v1, v2, v28)
    case "hello":
        fmt.Println(v2, v1, v2)
    default:
        v2 += (v2 & 2147483647)
    }
    v29 := GenS1{F0: 976, F1: 42, F2: 42}
    var v30 float32 = 1.5
    fmt.Println(v30, v28)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
