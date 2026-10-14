package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 string }
func genFn0(p0 uint) uint16 {
    var v6 float32 = (0.1 / (1.5 - -2.25))
    v7 := []int{32767, -9223372036854775808}
    p0 -= p0
    _ = p0
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 bool, p1 int16) bool {
    v8 := GenS1{F0: 570, F1: (1.5 <= 0.1), F2: ""}
    for fi9 := 0; fi9 < 8; fi9++ {
        p1 = -32768
        fmt.Println(p0)
        p0 = (337 > 985)
        _ = fi9
    }
    if true {
        v10 := []int{32767, -2147483648, -532, -2147483648}
        p0 = p0
        _ = v10
    }
    var v11 int16 = p1
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := []int{(fi12 & 225), (fi12 ^ (^(fi12))), 32767, (fi12 % ((fi12) | 1))}
        fmt.Println(p0, v11)
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v8
    return ((p0 && (-1 <= -9223372036854775808)) && p0)
}

func main() {
    var v1 bool = true
    v2 := make(map[string]bool)
    v3 := make(map[uint32]int)
    var v4 int = -106
    var v5 string = "hello"
    var v14 int32 = -1
    v5 = "foo bar"
    var v15 uint64 = 18446744073709551615
    v5 = ("foo bar" + v5)
    v16 := []int{(v4 * (v4 * (v4 ^ v4))), (v4 ^ v4), (v4 ^ v4)}
    fmt.Println(v4)
    switch ((("\n\t" + ("\n\t" + v5)) + (("foo bar" + v5) + v5)) + ("hello" + "")) {
    case "":
        v17 := GenS1{F0: 2, F1: ((v5 < v5) || (0.1 == 0.0)), F2: ("hello" + v5)}
        _ = v17
    case "\n\t":
        if ((^(v14)) > v14) {
            var v18 int8 = 0
            var v19 int16 = 887
            v4 = 42
            _ = v18
            _ = v19
        }
    case "a":
        switch (v4 >> ((255) & 63)) {
        case 4294967295:
            v20 := v2[(v5 + v5)]
            _ = v20
        }
    }
    if v1 {
        var v21 uint8 = 0
        _ = v21
    }
    var v22 uint64 = (v15 - 227)
    v15 = v15
    for fi23 := 0; fi23 < 3; fi23++ {
        var v24 float32 = 0.1
        var v25 uint16 = 2
        _ = fi23
        _ = v24
        _ = v25
    }
    fmt.Println(v14)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
