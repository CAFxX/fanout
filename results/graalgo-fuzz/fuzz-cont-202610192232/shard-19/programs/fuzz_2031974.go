package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int32; F1 int64; F2 uint }
func genFn0(p0 string, p1 uint16, p2 uint) uint64 {
    fmt.Println(p0, p0, p1)
    v6 := GenS1{F0: ""}
    if false {
        p2 = (p2 % ((18446744073709551615) | 1))
        switch (p0 + (("" + p0) + "\n\t")) {
        case "":
            if true {
                for fi7 := 0; fi7 < 8; fi7++ {
                    v8 := []int{493, -285}
                    p1 -= (p1 ^ p1)
                    if (p0 != p0) {
                        fmt.Println(p2, p2, fi7)
                    }
                    _ = fi7
                    _ = v8
                }
                v9 := GenS1{F0: ((p0 + (((((("hello" + p0) + ((p0 + p0) + p0)) + ((p0 + p0) + p0)) + p0) + p0) + p0)) + p0)}
                _ = v9
            } else {
                v10 := []int{}
                _ = v10
            }
        }
        fmt.Println(p1, p0, p2)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}

func main() {
    var v1 float32 = 0.1
    var v2 bool = true
    var v3 int32 = -2147483648
    v4 := make(map[int]uint32)
    var v5 uint64 = 18446744073709551615
    v11 := GenS2{F0: 58, F1: 1}
    if ((!(true && (!true))) && (v2 || v2)) {
        v12 := []int{}
        _ = v12
    } else {
        v13 := []int{42, 561, 363, -549}
        v2 = v2
        v2 = v2
        _ = v13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        if false {
            switch "hello" {
            case "":
                v5 = (^(v5))
            case "a":
                switch 1 {
                case -2147483648:
                    var v15 uint = 4294967295
                    _ = v15
                case -527:
                    v5 += 4294967295
                case 211:
                    v3 -= 1
                default:
                    v2 = true
                }
            }
            var v16 int8 = -5
            _ = v16
        }
        if v2 {
            v1 -= ((1.5 * (v1 * v1)) - v1)
            _ = len(v4)
            fmt.Println(v1)
        }
        v2 = v2
        _ = fi14
    }
    var v17 uint64 = v5
    v2 = (!v2)
    switch (("" + "foo bar") + "foo bar") {
    case "\n\t":
        fmt.Println(v1)
    case "a":
        var v18 uint16 = 206
        _ = v18
    }
    v4[792] = 0
    for fi19 := 0; fi19 < 19; fi19++ {
        var v20 int32 = v3
        v21 := []int{0, fi19, (fi19 + (fi19 << ((42) & 63)))}
        v22 := GenS3{F0: (v3 >> ((0) & 63)), F1: -1, F2: 0}
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
