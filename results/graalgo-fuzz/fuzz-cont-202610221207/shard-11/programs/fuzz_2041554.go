package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int16; F2 int32 }
func genFn0(p0 int32, p1 uint64, p2 uint64) string {
    var v5 int32 = p0
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return ""
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]int)
    var v3 uint16 = 659
    v4 := make(map[uint32]uint32)
    var v6 string = (("\n\t" + "hello") + "")
    v7 := []int{-9223372036854775808, 32768, 60}
    if false {
        var v8 int32 = -2147483648
        v6 = "hello"
        v9 := GenS1{F0: 127}
        _ = v8
        _ = v9
    } else {
        v3 = v3
        var v10 uint32 = 2147483648
        _ = v10
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        if true {
            switch ("" + v6) {
            case "\n\t":
                switch "hello" {
                case "hello":
                    fmt.Println(fi11, v3, fi11)
                default:
                    switch (("" + "hello") + "\n\t") {
                    case "\n\t":
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    case "":
                        v3 += (v3 * (v3 ^ (v3 - 0)))
                    }
                }
            case "a":
                switch "foo bar" {
                case "a":
                    v6 = (v6 + (v6 + ""))
                }
            }
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 bool = ((!false) && (-118 != 127))
                var v14 int64 = 42
                fmt.Println(fi11, v6, v13)
                _ = fi12
                _ = v13
                _ = v14
            }
            fmt.Println(v3, fi11)
        } else {
            var v15 int32 = 895
            _ = v15
        }
        _ = fi11
    }
    if false {
        switch "" {
        case "hello":
            v6 = "\n\t"
        default:
            v3 = 2
        }
        fmt.Println(v6, v6)
        var v16 uint16 = (v3 - v3)
        _ = v16
    }
    var v17 uint32 = 588
    var v18 float64 = 0.0
    var v19 int64 = -1
    var v20 int64 = v19
    if ((100 <= 127) || false) {
        for fi21 := 0; fi21 < 3; fi21++ {
            v6 = (v6 + "\n\t")
            for fi22 := 0; fi22 < 4; fi22++ {
                v20 += v20
                v3 = v3
                var v23 int = (fi21 - fi22)
                v17 = v17
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
    }
    v24 := GenS2{F0: 4294967295, F1: int16(v17), F2: -2147483648}
    switch v6 {
    case "":
        v25 := GenS2{F0: 255, F1: 0, F2: 513}
        _ = v25
    default:
        v6 = v6
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
