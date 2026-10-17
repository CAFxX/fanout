package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int8, p1 uint64, p2 uint) int64 {
    var v6 int32 = 1
    var v7 int64 = -664
    v8 := GenS1{F0: 1, F1: (-(p0))}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return v7
}

func main() {
    var v1 uint = 117
    v2 := make(map[int]bool)
    v3 := []string{"\n\t", (("\n\t" + "foo bar") + ("a" + "a")), ("foo bar" + ("\n\t" + "a"))}
    var v4 int8 = -128
    var v5 bool = (-121 != (v4 / ((v4) | 1)))
    v2[2147483647] = (42 >= -2147483648)
    var v9 uint16 = 447
    var v10 uint16 = (v9 * v9)
    if true {
        var v11 float64 = 1e10
        _ = v11
    }
    fmt.Println(v1, v4, v1)
    var v12 int16 = 42
    var v13 int = 875
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 int8 = -1
        for fi16 := 0; fi16 < 18; fi16++ {
            switch (fi16 | 42) {
            case 4294967295:
                var v17 float64 = (3.14159 * 0.0)
                _ = v17
            case 42:
                if ((3.14159 + 3.14159) > 1e10) {
                    fmt.Println(fi16, v9)
                }
            }
            _ = fi16
        }
        for fi18 := 0; fi18 < 19; fi18++ {
            v1 = v1
            v19 := GenS2{F0: 2147483647, F1: v10}
            _ = fi18
            _ = v19
        }
        _ = fi14
        _ = v15
    }
    switch (("foo bar" + "") + ("a" + (("" + "hello") + "foo bar"))) {
    case "":
        v3 = append(v3, "")
    case "a":
        v4 += -29
    default:
        v20 := GenS3{F0: 1}
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
