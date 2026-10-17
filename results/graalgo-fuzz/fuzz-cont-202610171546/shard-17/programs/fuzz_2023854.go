package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 bool; F2 uint }
type GenS6 struct { F0 uint8; F1 uint; F2 uint16 }
func genFn0() uint64 {
    var v7 int16 = 1
    if false {
        var v8 int16 = (v7 >> ((978) & 63))
        _ = v8
    } else {
        switch 128 {
        case 127:
            if (true || true) {
                var v9 int64 = -374
                v9 = v9
                fmt.Println(v9)
                _ = v9
            }
        }
        v10 := GenS1{F0: 18446744073709551615, F1: 128}
        _ = v10
    }
    var v11 int8 = 1
    _ = v11
    _ = v7
    return 0
}
func genFn1(p0 uint, p1 uint64, p2 int16) int {
    for fi12 := 0; fi12 < 7; fi12++ {
        p0 -= p0
        v13 := []int{(fi12 & 650), 42}
        var v14 uint16 = 627
        _ = fi12
        _ = v13
        _ = v14
    }
    fmt.Println(p2, p1)
    if false {
        v15 := GenS2{F0: 1, F1: p1}
        var v16 int16 = -32768
        switch 4294967295 {
        case -459:
            p1 = 2
        case 229:
            var v17 int = 32767
            _ = v17
        case 2147483647:
            p0 = uint(p1)
        default:
            v18 := GenS3{F0: p0}
            _ = v18
        }
        _ = v15
        _ = v16
    } else {
        switch "foo bar" {
        case "a":
            v19 := GenS2{F0: 690, F1: p1}
            _ = v19
        }
    }
    p0 -= uint(p1)
    _ = p0
    _ = p1
    _ = p2
    return -9223372036854775808
}

func main() {
    var v1 uint16 = 2
    v2 := []string{}
    var v3 int8 = -65
    var v4 float64 = 0.1
    v5 := make(map[int]int)
    v6 := make(map[int]bool)
    v3 = (v3 ^ v3)
    v3 += int8(v1)
    switch "hello" {
    case "\n\t":
        var v20 bool = true
        _ = v20
    default:
        var v21 int = 128
        _ = v21
    }
    v5[42] = genFn1(827, 1, -32768)
    for fi22 := 0; fi22 < 5; fi22++ {
        var v23 int32 = 0
        v2 = append(v2, "")
        _ = fi22
        _ = v23
    }
    if (!(!true)) {
        v1 -= (v1 + (v1 * v1))
        for fi24 := 0; fi24 < 6; fi24++ {
            v25 := GenS4{F0: ("foo bar" + (("a" + ("" + (("\n\t" + "\n\t") + ""))) + ("hello" + "a")))}
            v1 = v1
            _ = fi24
            _ = v25
        }
        v4 -= (v4 * float64((-2.25 * 1e10)))
    } else {
        switch ("hello" + "hello") {
        case "\n\t":
            v3 = 0
        }
        var v26 int64 = 42
        v27 := v5[32768]
        _ = v26
        _ = v27
    }
    fmt.Println(v1)
    fmt.Println(v4, v4)
    v28 := GenS5{F0: 657, F1: true, F2: 65535}
    fmt.Println(v3)
    switch 563 {
    case 42:
        var v29 uint32 = 4294967295
        _ = v29
    case -972:
        for fi30 := 0; fi30 < 15; fi30++ {
            v3 = v3
            var v31 int16 = 70
            fmt.Println(v1)
            _ = fi30
            _ = v31
        }
    }
    switch -9223372036854775808 {
    case 1:
        v32 := GenS6{F0: 31, F1: 255, F2: v1}
        _ = v32
    case 0:
        switch -576 {
        case 1:
            var v33 uint8 = 42
            _ = v33
        case 42:
            var v34 bool = false
            _ = v34
        case 32768:
            var v35 int16 = -32768
            _ = v35
        default:
            var v36 string = ""
            _ = v36
        }
    case 127:
        fmt.Println(v3, v4, v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
