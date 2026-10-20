package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 int64; F2 string }
func genFn0(p0 int, p1 uint, p2 int64) uint {
    var v7 int16 = 42
    var v8 string = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p1
}
func genFn1(p0 uint64, p1 uint32) string {
    var v9 bool = true
    v10 := []int{416, -1, 2147483647, -772}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return "foo bar"
}

func main() {
    var v1 int = -473
    v2 := make(map[string]bool)
    v3 := make(map[int]uint32)
    var v4 uint64 = 4294967295
    var v5 uint16 = 1
    var v6 uint16 = v5
    v11 := v2[""]
    fmt.Println(v1, v1, v4)
    fmt.Println(v5)
    v12 := v2[("" + "")]
    if (142 >= 0) {
        v12 = ((("" + "hello") == ("" + (("foo bar" + "hello") + "hello"))) && false)
        var v13 int64 = 9223372036854775807
        _ = v13
    } else {
        v6 += v6
        var v14 uint8 = 42
        v15 := GenS1{F0: "\n\t"}
        _ = v14
        _ = v15
    }
    fmt.Println(v5, v6)
    if (59 >= 127) {
        var v16 int64 = 782
        var v17 string = ("foo bar" + "foo bar")
        _ = v16
        _ = v17
    } else {
        v3[0] = 65535
        v18 := v3[v1]
        _ = v18
    }
    v12 = true
    for fi19 := 0; fi19 < 13; fi19++ {
        v20 := []int{(fi19 ^ 0)}
        _ = fi19
        _ = v20
    }
    switch "" {
    case "foo bar":
        var v21 int16 = -627
        _ = v21
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        for fi23 := 0; fi23 < 7; fi23++ {
            v24 := GenS2{F0: 0, F1: -1, F2: "foo bar"}
            v25 := GenS1{F0: ("foo bar" + "\n\t")}
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi22
    }
    var v26 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
