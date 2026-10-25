package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint16, p1 uint64) string {
    switch -11 {
    case -363:
        v5 := []int{508, 4294967295, 731}
        _ = v5
    case 95:
        var v6 int64 = -9223372036854775808
        _ = v6
    case 32768:
        v7 := []int{32768, -2147483648, 9223372036854775807}
        _ = v7
    default:
        p0 = (p0 | p0)
    }
    var v8 uint16 = p0
    p0 -= (p0 % ((0) | 1))
    for fi9 := 0; fi9 < 8; fi9++ {
        fmt.Println(fi9, fi9, p0)
        _ = fi9
    }
    v10 := []int{128}
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    return "\n\t"
}
func genFn1(p0 uint) int16 {
    v11 := []int{}
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 7; fi13++ {
            p0 -= 65535
            _ = fi13
        }
        if ((("\n\t" + "") + "\n\t") <= "hello") {
            var v14 uint64 = 18446744073709551615
            var v15 int = (fi12 ^ fi12)
            _ = v14
            _ = v15
        }
        var v16 int8 = 1
        _ = fi12
        _ = v16
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        var v18 uint64 = 0
        fmt.Println(v18, fi17, p0)
        for fi19 := 0; fi19 < 5; fi19++ {
            v11 = append(v11, fi17)
            _ = fi19
        }
        _ = fi17
        _ = v18
    }
    _ = p0
    _ = v11
    return -1
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 1
    v3 := make(map[uint32]string)
    var v4 int32 = (^(v2))
    for fi20 := 0; fi20 < 16; fi20++ {
        switch "foo bar" {
        case "":
            var v21 uint8 = 255
            _ = v21
        default:
            if ((v4 / ((v2) | 1)) != v2) {
                for fi22 := 0; fi22 < 15; fi22++ {
                    switch fi22 {
                    case 42:
                        _ = len(v3)
                    }
                    var v23 int32 = (^(v4))
                    fmt.Println(v4, v4, fi22)
                    _ = fi22
                    _ = v23
                }
            }
        }
        _ = fi20
    }
    v4 = -2147483648
    v2 = v2
    var v24 float64 = -2.25
    var v25 int = 188
    var v26 int64 = 353
    v27 := GenS1{F0: "foo bar"}
    if ((!(!(true || false))) || false) {
        v28 := GenS1{F0: ""}
        var v29 float64 = float64(1.5)
        if true {
            fmt.Println(v25, v24)
        }
        _ = v28
        _ = v29
    }
    var v30 int = v25
    switch ("" + "") {
    case "hello":
        v31 := []int{v25, v25}
        _ = v31
    }
    for fi32 := 0; fi32 < 18; fi32++ {
        var v33 int16 = -944
        v34 := []int{}
        _ = fi32
        _ = v33
        _ = v34
    }
    v24 = ((v24 + v24) * 1e10)
    for fi35 := 0; fi35 < 12; fi35++ {
        if (!(!true)) {
            v25 = v25
            for fi36 := 0; fi36 < 16; fi36++ {
                var v37 bool = true
                var v38 float64 = (float64(1e10) * 1.5)
                v4 = 2147483647
                _ = fi36
                _ = v37
                _ = v38
            }
        }
        _ = fi35
    }
    var v39 uint32 = 0
    v39 -= (v39 - 4294967295)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v39
    _ = v4
}
