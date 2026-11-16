package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
func genFn0(p0 uint, p1 int16, p2 uint64) bool {
    switch "a" {
    case "\n\t":
        for fi6 := 0; fi6 < 15; fi6++ {
            if false {
                switch (fi6 >> ((p0) & 63)) {
                case -9223372036854775808:
                    v7 := []int{(fi6 / (((^(fi6))) | 1))}
                    _ = v7
                }
            }
            v8 := GenS1{F0: -1, F1: p0}
            _ = fi6
            _ = v8
        }
    case "foo bar":
        var v9 int8 = -128
        _ = v9
    }
    if ((p0 << ((p0) & 63)) <= p0) {
        v10 := GenS1{F0: int64(p0), F1: 670}
        fmt.Println(p2)
        _ = v10
    } else {
        fmt.Println(p0, p2, p0)
    }
    switch ("a" + "a") {
    case "":
        for fi11 := 0; fi11 < 15; fi11++ {
            p2 = (p2 - p2)
            p1 -= (p1 >> ((65535) & 63))
            _ = fi11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return (p1 >= (p1 << ((p0) & 63)))
}

func main() {
    v1 := []string{"a", "", "a", "a"}
    v2 := make(map[int]bool)
    var v3 int32 = 42
    v4 := make(map[uint32]bool)
    v5 := make(map[string]string)
    var v12 uint32 = 65535
    var v13 uint64 = 915
    var v14 string = "\n\t"
    if false {
        v15 := GenS1{F0: -66, F1: 2}
        _ = v15
    }
    v16 := GenS1{F0: 42, F1: 0}
    if len(v1) > 0 {
        v1[0] = ("foo bar" + v14)
    }
    if false {
        v3 = 324
    }
    if len(v1) > 0 {
        v1[0] = ("foo bar" + "hello")
    }
    var v17 string = (v14 + ((v14 + "foo bar") + v14))
    if false {
        for fi18 := 0; fi18 < 3; fi18++ {
            var v19 float64 = ((-2.25 + 3.14159) / -2.25)
            for fi20 := 0; fi20 < 5; fi20++ {
                _ = len(v1)
                if true {
                    v3 = int32(v13)
                }
                _ = fi20
            }
            var v21 int16 = 1
            _ = fi18
            _ = v19
            _ = v21
        }
    } else {
        v13 -= v13
        fmt.Println(v13)
    }
    if (((float64(1e10) - 3.14159) * (0.1 * (0.0 * 0.0))) < ((0.0 - (-2.25 - 0.1)) / (0.1 + 0.1))) {
        fmt.Println(v3)
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        _ = len(v5)
        _ = fi22
    }
    if len(v1) > 0 {
        v1[0] = ("foo bar" + v14)
    }
    if false {
        for fi23 := 0; fi23 < 13; fi23++ {
            var v24 int64 = 0
            var v25 float64 = (1e10 - 1e10)
            fmt.Println(v12, v14)
            _ = fi23
            _ = v24
            _ = v25
        }
        if true {
            if (2 <= 209) {
                for fi26 := 0; fi26 < 4; fi26++ {
                    fmt.Println(fi26, v3, v12)
                    v14 = ("foo bar" + "")
                    for fi27 := 0; fi27 < 16; fi27++ {
                        v28 := v4[(v12 ^ v12)]
                        v17 = v14
                        _ = fi27
                        _ = v28
                    }
                    _ = fi26
                }
                v29 := GenS1{F0: 42, F1: 255}
                _ = v29
            }
        } else {
            var v30 int = -519
            delete(v5, v14)
            var v31 bool = (!false)
            _ = v30
            _ = v31
        }
    }
    var v32 int8 = 43
    v32 -= (v32 + 42)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
