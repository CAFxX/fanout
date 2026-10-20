package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int16 }

func main() {
    v1 := []int{-947, -358}
    v2 := []string{("foo bar" + ("\n\t" + "foo bar"))}
    var v3 uint = 42
    v4 := make(map[uint32]int)
    v5 := make(map[int]string)
    v3 = v3
    switch "" {
    case "foo bar":
        if (-198 == 1) {
            for fi6 := 0; fi6 < 13; fi6++ {
                for fi7 := 0; fi7 < 20; fi7++ {
                    _ = len(v2)
                    _ = fi7
                }
                var v8 int64 = 42
                var v9 uint64 = 926
                _ = fi6
                _ = v8
                _ = v9
            }
        }
    }
    var v10 int32 = 273
    v2 = append(v2, (("hello" + "") + "foo bar"))
    for fi11 := 0; fi11 < 2; fi11++ {
        if (818 >= 288) {
            for fi12 := 0; fi12 < 10; fi12++ {
                v13 := GenS1{F0: fi12, F1: 2, F2: 0}
                v3 += (v3 / (((^(v3))) | 1))
                _ = fi12
                _ = v13
            }
            var v14 float64 = 0.1
            var v15 uint = (v3 * (^(v3)))
            _ = v14
            _ = v15
        }
        _ = fi11
    }
    var v16 string = ""
    if (2 >= 128) {
        var v17 float32 = (-2.25 * float32(((-2.25 + 0.0) - 3.14159)))
        switch 127 {
        case 594:
            fmt.Println(v17, v16)
        case 128:
            var v18 int8 = 42
            _ = v18
        default:
            v16 = v16
        }
        _ = v17
    } else {
        var v19 uint = 2
        v19 = v19
        switch 99 {
        case 0:
            fmt.Println(v10, v10, v16)
        case 32768:
            fmt.Println(v10, v10)
        default:
            v20 := v5[42]
            _ = v20
        }
        _ = v19
    }
    v3 = 18446744073709551615
    if false {
        fmt.Println(v16)
    } else {
        v3 = v3
    }
    var v21 uint32 = 0
    _ = len(v5)
    switch v16 {
    case "":
        for fi22 := 0; fi22 < 15; fi22++ {
            v4[2] = (fi22 << ((4294967295) & 63))
            switch (^(fi22)) {
            case 127:
                v23 := v4[v21]
                _ = v23
            case -2147483648:
                var v24 uint32 = 2
                _ = v24
            }
            _ = fi22
        }
    case "hello":
        switch "\n\t" {
        case "a":
            switch -2147483648 {
            case -2147483648:
                v25 := GenS1{F0: 1, F1: v3, F2: 1}
                _ = v25
            }
        case "":
            var v26 int16 = 1000
            _ = v26
        }
    }
    v3 = (^(v3))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
