package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int32 }
func genFn0(p0 bool, p1 int, p2 bool) string {
    p1 = p1
    switch (("\n\t" + "foo bar") + ("\n\t" + "a")) {
    case "\n\t":
        if (p0 || true) {
            v7 := []int{-21}
            _ = v7
        } else {
            var v8 int16 = -538
            v9 := []int{671, p1, p1}
            _ = v8
            _ = v9
        }
    case "foo bar":
        p0 = false
    case "hello":
        var v10 string = "a"
        _ = v10
    }
    if (!(2 >= 156)) {
        switch -644 {
        case -9223372036854775808:
            p0 = (p0 || false)
        default:
            var v11 string = "hello"
            _ = v11
        }
        switch "foo bar" {
        case "\n\t":
            v12 := []int{406, 1}
            _ = v12
        case "foo bar":
            var v13 int32 = 995
            _ = v13
        }
    } else {
        p1 = 9223372036854775807
        var v14 uint8 = 85
        _ = v14
    }
    v15 := GenS1{F0: (p1 & p1)}
    var v16 string = ("" + "")
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return ("a" + "foo bar")
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 1e10
    v3 := []string{("hello" + "\n\t"), "foo bar"}
    v4 := make(map[int]int)
    var v5 string = "hello"
    var v6 uint64 = 0
    fmt.Println(v6)
    var v17 bool = false
    v18 := GenS2{F0: -345, F1: (^(v6)), F2: 65535}
    v19 := v1[715]
    v3 = append(v3, (("foo bar" + v5) + v5))
    fmt.Println(v19)
    v6 = (v6 >> ((858) & 63))
    if len(v3) > 0 {
        v3[0] = v5
    }
    v2 += (1.5 - v2)
    _ = len(v3)
    fmt.Println(v17, v2)
    if v19 {
        v19 = (!(65535 < 140))
    }
    var v20 int = 42
    if v19 {
        v3 = append(v3, "a")
        v21 := GenS3{F0: -132}
        _ = v21
    }
    var v22 int64 = -1
    v5 = v5
    var v23 uint = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
