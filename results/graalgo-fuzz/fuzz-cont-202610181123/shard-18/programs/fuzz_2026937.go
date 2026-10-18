package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int32 }
func genFn0(p0 uint, p1 int64) uint32 {
    var v6 float32 = -2.25
    p1 = (-(p1))
    if (1 < 127) {
        p1 = p1
        p1 -= p1
    } else {
        for fi7 := 0; fi7 < 19; fi7++ {
            fmt.Println(v6, fi7, v6)
            v8 := GenS1{F0: (fi7 >> ((p0) & 63)), F1: false, F2: 0}
            _ = fi7
            _ = v8
        }
        p0 -= (p0 << ((4294967295) & 63))
    }
    fmt.Println(v6, v6)
    _ = p0
    _ = p1
    _ = v6
    return 2147483648
}

func main() {
    var v1 uint64 = 1
    v2 := []string{"", "hello"}
    v3 := make(map[string]bool)
    v4 := []string{"\n\t", "hello"}
    v5 := make(map[uint32]bool)
    switch 1 {
    case 662:
        var v9 int64 = 9223372036854775807
        _ = v9
    }
    v1 -= 979
    if true {
        var v10 int64 = -180
        var v11 bool = false
        _ = v10
        _ = v11
    } else {
        var v12 float32 = 0.1
        _ = v12
    }
    var v13 int8 = 100
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v13 = v13
    switch "a" {
    case "hello":
        for fi14 := 0; fi14 < 18; fi14++ {
            v15 := v5[253]
            var v16 uint16 = 964
            _ = fi14
            _ = v15
            _ = v16
        }
    case "\n\t":
        for fi17 := 0; fi17 < 16; fi17++ {
            v18 := v5[42]
            _ = fi17
            _ = v18
        }
    case "foo bar":
        for fi19 := 0; fi19 < 20; fi19++ {
            delete(v5, 51)
            _ = fi19
        }
    }
    if len(v4) > 0 {
        v4[0] = "a"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
