package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int16 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 bool; F1 string; F2 int32 }
type GenS4 struct { F0 string; F1 string; F2 int8 }
type GenS5 struct { F0 int32; F1 bool }
func genFn0(p0 int8) uint64 {
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 bool = (-1 >= -32768)
        v8 := GenS1{F0: fi6, F1: 1, F2: 42}
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 string = ""
        var v11 uint64 = 18446744073709551615
        _ = fi9
        _ = v10
        _ = v11
    }
    if (1.5 > 3.14159) {
        if false {
            var v12 uint64 = 18446744073709551615
            v13 := []int{-79, 127, 1, 32768}
            _ = v12
            _ = v13
        } else {
            p0 -= -1
            for fi14 := 0; fi14 < 7; fi14++ {
                p0 = (p0 * p0)
                _ = fi14
            }
            var v15 uint64 = 418
            _ = v15
        }
    }
    switch ("hello" + "") {
    case "a":
        var v16 int32 = -2147483648
        _ = v16
    case "\n\t":
        p0 = (-(p0))
    case "":
        v17 := []int{}
        _ = v17
    }
    switch 42 {
    case -63:
        v18 := GenS2{F0: ((809 > 1000) && (42 <= 307)), F1: 42}
        _ = v18
    case 834:
        v19 := GenS3{F0: true, F1: ("hello" + "hello"), F2: -667}
        _ = v19
    default:
        switch "hello" {
        case "a":
            v20 := GenS1{F0: 127, F1: 0, F2: -1}
            _ = v20
        case "\n\t":
            switch -1 {
            case 227:
                var v21 float64 = 0.1
                _ = v21
            default:
                v22 := GenS4{F0: "hello", F1: "foo bar", F2: (p0 & p0)}
                _ = v22
            }
        }
    }
    _ = p0
    return 42
}
func genFn1(p0 int32) int16 {
    var v23 uint8 = 1
    v24 := GenS1{F0: 2147483647, F1: 559, F2: 42}
    _ = p0
    _ = v23
    _ = v24
    return 0
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[string]string)
    v3 := make(map[uint32]bool)
    v4 := []string{"hello", ("" + "\n\t"), "hello"}
    v5 := []uint32{2147483648, 399, 42}
    var v25 bool = ((((0.1 / float64(-2.25)) + 1e10) > 1.5) || ((!false) || true))
    for fi26 := 0; fi26 < 4; fi26++ {
        v27 := GenS3{F0: v25, F1: ("a" + ""), F2: (v1 | -887)}
        _ = fi26
        _ = v27
    }
    v28 := GenS2{F0: (127 <= 255), F1: 42}
    v25 = (6 > -127)
    for fi29 := 0; fi29 < 13; fi29++ {
        if v25 {
            var v30 int16 = 1
            v5 = append(v5, 2147483647)
            _ = v30
        }
        _ = fi29
    }
    v25 = v25
    delete(v3, 42)
    switch ("\n\t" + "a") {
    case "a":
        v1 = -693
    case "\n\t":
        switch "a" {
        case "a":
            if (!v25) {
                v25 = v25
                var v31 string = ("" + "foo bar")
                var v32 float64 = -2.25
                _ = v31
                _ = v32
            }
        case "\n\t":
            var v33 uint = 375
            _ = v33
        }
    case "hello":
        for fi34 := 0; fi34 < 19; fi34++ {
            v35 := GenS5{F0: v1, F1: (1.5 != 0.0)}
            for fi36 := 0; fi36 < 1; fi36++ {
                v37 := GenS1{F0: (-(fi34)), F1: 255, F2: 0}
                _ = fi36
                _ = v37
            }
            _ = fi34
            _ = v35
        }
    }
    var v38 uint8 = 1
    var v39 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
