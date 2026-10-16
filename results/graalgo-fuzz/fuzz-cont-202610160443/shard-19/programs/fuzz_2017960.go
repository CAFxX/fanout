package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int }
func genFn0(p0 bool, p1 int64) int16 {
    p1 -= (p1 >> ((0) & 63))
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 float32 = 3.14159
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: p1, F1: 107}
    v10 := []int{128, 1}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return -32768
}

func main() {
    v1 := make(map[uint32]int)
    var v2 string = "hello"
    var v3 float32 = 3.14159
    v4 := make(map[uint32]uint32)
    var v5 int32 = 869
    v6 := make(map[string]uint32)
    v11 := GenS2{F0: 0}
    for fi12 := 0; fi12 < 16; fi12++ {
        v5 += v5
        var v13 float32 = v3
        switch (fi12 << ((2) & 63)) {
        case 4294967295:
            var v14 int8 = 0
            _ = v14
        }
        _ = fi12
        _ = v13
    }
    if false {
        v15 := []int{-2147483648, 1, -896}
        v16 := GenS2{F0: 9223372036854775807}
        _ = v15
        _ = v16
    }
    fmt.Println(v3)
    v5 = -1
    fmt.Println(v5)
    if true {
        if ((100 > -76) || true) {
            switch (("hello" + v2) + (v2 + ("hello" + "a"))) {
            case "":
                v5 = v5
            case "\n\t":
                for fi17 := 0; fi17 < 7; fi17++ {
                    v2 = "foo bar"
                    _ = fi17
                }
            case "foo bar":
                for fi18 := 0; fi18 < 13; fi18++ {
                    for fi19 := 0; fi19 < 11; fi19++ {
                        v20 := GenS2{F0: (fi19 + fi19)}
                        v2 = "a"
                        _ = len(v1)
                        _ = fi19
                        _ = v20
                    }
                    _ = fi18
                }
            default:
                v3 = (float32(-2.25) * (-2.25 - ((1.5 + ((v3 / v3) + v3)) / v3)))
            }
        }
        v21 := []int{}
        if ((true || false) || (42 != 426)) {
            if len(v21) >= 2 {
                v21 = v21[1:len(v21)]
            }
        }
        _ = v21
    } else {
        switch 994 {
        case -1:
            for fi22 := 0; fi22 < 16; fi22++ {
                _ = len(v4)
                _ = fi22
            }
        case 128:
            var v23 uint32 = 1
            _ = v23
        case 786:
            v24 := GenS1{F0: 42, F1: 1}
            _ = v24
        }
    }
    var v25 bool = (-1 > -9223372036854775808)
    var v26 int64 = 9223372036854775807
    v2 = v2
    v5 = v5
    v27 := []int{-946, -9223372036854775808, 32767}
    delete(v6, "hello")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
