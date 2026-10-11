package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int64 }
type GenS2 struct { F0 int16; F1 int; F2 string }
type GenS3 struct { F0 int8; F1 int16; F2 int8 }
func genFn0(p0 int8, p1 uint32, p2 string) int16 {
    v7 := []int{-1, -9223372036854775808, 42}
    if (true || (true && false)) {
        v8 := GenS1{F0: 2, F1: (p0 * 0), F2: 9223372036854775807}
        _ = v8
    } else {
        v9 := GenS1{F0: 2, F1: (p0 & (p0 | (p0 | 0))), F2: 0}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 76
}
func genFn1(p0 uint16, p1 uint16, p2 uint64) int32 {
    switch 42 {
    case 1:
        var v10 uint32 = 2
        _ = v10
    }
    switch "hello" {
    case "hello":
        var v11 int32 = 2147483647
        _ = v11
    case "":
        p2 += 2
    case "a":
        for fi12 := 0; fi12 < 3; fi12++ {
            p1 = (p1 | 0)
            var v13 int32 = -1
            _ = fi12
            _ = v13
        }
    default:
        fmt.Println(p2)
    }
    _ = p0
    _ = p1
    _ = p2
    return 606
}

func main() {
    v1 := []string{"\n\t", "hello", "foo bar", (("foo bar" + "hello") + ("a" + "foo bar"))}
    var v2 uint64 = 1
    v3 := make(map[string]int)
    var v4 float32 = float32((3.14159 + 3.14159))
    v5 := make(map[uint32]int)
    v6 := make(map[uint32]int)
    v14 := v5[0]
    for fi15 := 0; fi15 < 17; fi15++ {
        _ = len(v5)
        var v16 int8 = 1
        v1 = append(v1, "foo bar")
        _ = fi15
        _ = v16
    }
    var v17 int32 = 2147483647
    var v18 int = 1
    if ((!false) || true) {
        if false {
            fmt.Println(v17)
        } else {
            var v19 float32 = -2.25
            _ = v19
        }
    } else {
        delete(v6, 0)
        var v20 int64 = -9223372036854775808
        var v21 int64 = 951
        _ = v20
        _ = v21
    }
    v22 := v5[2]
    v2 = (v2 >> ((65535) & 63))
    var v23 uint16 = 975
    var v24 uint16 = v23
    for fi25 := 0; fi25 < 16; fi25++ {
        v18 -= (fi25 << ((599) & 63))
        v26 := GenS2{F0: 1000, F1: (^(fi25)), F2: (("\n\t" + ("\n\t" + "\n\t")) + "a")}
        _ = fi25
        _ = v26
    }
    v27 := GenS3{F0: 1, F1: -479, F2: -128}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
