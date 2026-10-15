package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 bool; F1 int8; F2 int16 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 int64; F1 int; F2 uint16 }
func genFn0(p0 int16, p1 uint32) uint32 {
    switch "foo bar" {
    case "foo bar":
        v4 := []int{127, 32768, 128}
        _ = v4
    default:
        var v5 float32 = 0.0
        _ = v5
    }
    if (!(42 < 30)) {
        switch "\n\t" {
        case "":
            p1 = (p1 * p1)
        default:
            for fi6 := 0; fi6 < 10; fi6++ {
                var v7 bool = true
                v8 := []int{fi6, (^(fi6))}
                v9 := GenS1{F0: (-(p0)), F1: 127}
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
        }
        v10 := GenS2{F0: (18446744073709551615 >= 412), F1: -36, F2: p0}
        _ = v10
    }
    switch "\n\t" {
    case "foo bar":
        if true {
            for fi11 := 0; fi11 < 4; fi11++ {
                p0 = p0
                _ = fi11
            }
        }
    }
    _ = p0
    _ = p1
    return (p1 * p1)
}

func main() {
    v1 := []string{"foo bar", "\n\t", ("\n\t" + "foo bar"), "\n\t"}
    var v2 float64 = 0.0
    v3 := make(map[int]uint32)
    var v12 uint64 = 526
    v13 := GenS3{F0: 32767, F1: ""}
    fmt.Println(v12)
    if len(v1) > 0 {
        v1[0] = ((("hello" + (("\n\t" + ((((("foo bar" + "hello") + "hello") + "hello") + ("hello" + "\n\t")) + "foo bar")) + "foo bar")) + "hello") + "a")
    }
    v2 = 1.5
    if false {
        var v14 bool = false
        var v15 uint = 1
        v16 := GenS3{F0: 32767, F1: "hello"}
        _ = v14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        v12 = 1
        for fi18 := 0; fi18 < 6; fi18++ {
            var v19 uint = 9
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        fmt.Println(v12, v12)
        v2 = v2
        _ = fi20
    }
    v12 -= 0
    v21 := GenS2{F0: ("a" <= ""), F1: -49, F2: 0}
    v22 := GenS4{F0: 9223372036854775807, F1: 846, F2: 42}
    for fi23 := 0; fi23 < 3; fi23++ {
        fmt.Println(fi23, fi23, v12)
        var v24 uint8 = 18
        _ = fi23
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
