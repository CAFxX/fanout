package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint; F1 int; F2 uint }
func genFn0(p0 uint8, p1 int8, p2 bool) int {
    var v5 uint = 0
    p2 = p2
    for fi6 := 0; fi6 < 9; fi6++ {
        v5 = (^(v5))
        p0 = p0
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 32767
}

func main() {
    v1 := []int{-9223372036854775808, 840, -2147483648}
    v2 := make(map[int]int)
    var v3 uint16 = 979
    v4 := make(map[string]uint32)
    var v7 int64 = -9223372036854775808
    var v8 int64 = v7
    v9 := GenS1{F0: -57, F1: 32767, F2: 618}
    switch ("\n\t" + "") {
    case "foo bar":
        switch "\n\t" {
        case "hello":
            switch 4294967295 {
            case 499:
                v10 := GenS1{F0: 0, F1: -32768, F2: v3}
                _ = v10
            case 128:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case 302:
                var v11 bool = true
                _ = v11
            }
        case "a":
            fmt.Println(v7)
        case "foo bar":
            for fi12 := 0; fi12 < 8; fi12++ {
                v13 := v4[""]
                _ = fi12
                _ = v13
            }
        }
    case "a":
        v14 := GenS2{F0: 2, F1: 128, F2: 255}
        _ = v14
    default:
        var v15 uint32 = 2147483647
        _ = v15
    }
    v3 += (^(v3))
    v16 := GenS1{F0: 127, F1: 1000, F2: (^(v3))}
    if (((-2.25 - (1e10 - 1.5)) / 1.0) == (0.0 * (-2.25 + float64(-2.25)))) {
        switch "" {
        case "foo bar":
            switch 32768 {
            case 2147483647:
                if true {
                    v17 := v2[127]
                    _ = v17
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v18 int = 127
                    _ = v18
                }
            case 42:
                for fi19 := 0; fi19 < 9; fi19++ {
                    var v20 uint8 = 1
                    if false {
                        fmt.Println(v7)
                        var v21 uint = 750
                        _ = v21
                    }
                    delete(v4, "foo bar")
                    _ = fi19
                    _ = v20
                }
            }
        }
    } else {
        var v22 int32 = 2147483647
        var v23 float32 = 0.1
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v22
        _ = v23
    }
    delete(v4, "foo bar")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
