package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int; F1 int64; F2 uint8 }
func genFn0(p0 int8, p1 int8) uint32 {
    var v5 int32 = -792
    p1 = -96
    if (!true) {
        var v6 uint16 = 1
        v7 := GenS1{F0: ("foo bar" + "\n\t"), F1: ("a" + ("a" + "a"))}
        var v8 bool = ((!false) || (true || true))
        _ = v6
        _ = v7
        _ = v8
    }
    p1 -= p1
    _ = p0
    _ = p1
    _ = v5
    return 42
}
func genFn1(p0 int, p1 string, p2 bool) int8 {
    var v9 bool = (42 != 42)
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 uint16 = 32767
        for fi12 := 0; fi12 < 20; fi12++ {
            v9 = (((false || ((-2.25 * 1.5) != (0.1 - -2.25))) || v9) && ((p2 && p2) || (595 > 1)))
            _ = fi12
        }
        v13 := GenS1{F0: "foo bar", F1: p1}
        _ = fi10
        _ = v11
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -128
}

func main() {
    var v1 int = 42
    v2 := []string{"hello", "", "hello"}
    v3 := make(map[string]int)
    v4 := make(map[int]string)
    var v14 int = 530
    fmt.Println(v1, v1)
    v15 := GenS2{F0: (^(v1)), F1: 0, F2: 177}
    var v16 string = "\n\t"
    fmt.Println(v1, v1)
    var v17 uint = 18446744073709551615
    _ = len(v2)
    for fi18 := 0; fi18 < 14; fi18++ {
        v17 += v17
        var v19 string = ("foo bar" + "a")
        fmt.Println(fi18, v16, v16)
        _ = fi18
        _ = v19
    }
    switch v1 {
    case 9223372036854775807:
        v17 = v17
    case 489:
        var v20 string = v16
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
