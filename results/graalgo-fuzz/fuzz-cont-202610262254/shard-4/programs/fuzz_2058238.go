package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 uint }
func genFn0(p0 int16, p1 bool) int32 {
    v7 := GenS1{F0: "\n\t"}
    p1 = (!false)
    _ = p0
    _ = p1
    _ = v7
    return 313
}
func genFn1(p0 string, p1 int, p2 uint16) bool {
    v8 := GenS1{F0: p0}
    if false {
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 string = (("a" + ("foo bar" + p0)) + p0)
            switch v10 {
            case "foo bar":
                v11 := []int{p1, 32768, p1, p1}
                _ = v11
            }
            _ = fi9
            _ = v10
        }
        for fi12 := 0; fi12 < 4; fi12++ {
            v13 := []int{(p1 & 205), (fi12 ^ fi12), (fi12 / ((fi12) | 1))}
            var v14 int64 = -904
            if (!(!false)) {
                v13 = append(v13, fi12)
            } else {
                var v15 float32 = (0.1 - 1.5)
                var v16 uint64 = 18446744073709551615
                v17 := GenS1{F0: p0}
                _ = v15
                _ = v16
                _ = v17
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        p0 = p0
    }
    p2 = p2
    if true {
        switch p1 {
        case -55:
            p2 = p2
        case -9223372036854775808:
            switch (p0 + "a") {
            case "foo bar":
                v18 := GenS1{F0: (p0 + "a")}
                _ = v18
            default:
                if false {
                    p2 = p2
                } else {
                    v19 := GenS1{F0: ((p0 + p0) + (p0 + "hello"))}
                    p2 = (^(p2))
                    _ = v19
                }
            }
        }
        var v20 uint32 = 1
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return true
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]uint32)
    v3 := []string{"\n\t", "\n\t", "foo bar"}
    var v4 bool = false
    v5 := []int{}
    var v6 uint16 = 270
    for fi21 := 0; fi21 < 18; fi21++ {
        v4 = true
        _ = fi21
    }
    switch ("\n\t" + "") {
    case "a":
        switch 4294967295 {
        case 127:
            if (!(true && ((!false) && (v4 || true)))) {
                fmt.Println(v4, v4)
            } else {
                var v22 int32 = genFn0(42, v4)
                v4 = (!v4)
                _ = v22
            }
        }
    case "hello":
        v23 := GenS1{F0: "a"}
        _ = v23
    case "foo bar":
        var v24 int32 = 42
        _ = v24
    default:
        var v25 uint8 = 255
        _ = v25
    }
    var v26 int64 = 42
    switch "hello" {
    case "foo bar":
        fmt.Println(v26, v6, v4)
    case "hello":
        v27 := GenS1{F0: (("foo bar" + "\n\t") + "hello")}
        _ = v27
    default:
        var v28 int16 = -1
        _ = v28
    }
    v29 := GenS2{F0: ("foo bar" + "hello"), F1: 1}
    v4 = v4
    var v30 int8 = 0
    v26 += 42
    for fi31 := 0; fi31 < 5; fi31++ {
        v32 := GenS1{F0: "\n\t"}
        _ = fi31
        _ = v32
    }
    v26 -= (-(v26))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
