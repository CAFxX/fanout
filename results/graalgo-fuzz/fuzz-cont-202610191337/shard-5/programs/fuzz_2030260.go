package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 bool; F1 int64; F2 int8 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int32; F1 int64; F2 uint16 }
func genFn0(p0 int) int32 {
    v5 := []int{}
    switch 0 {
    case 32768:
        p0 = 1
    }
    var v6 int32 = 0
    p0 = -89
    _ = p0
    _ = v5
    _ = v6
    return (v6 - (v6 ^ -2147483648))
}
func genFn1(p0 uint64, p1 bool, p2 int16) uint8 {
    var v7 int64 = 9223372036854775807
    switch (("hello" + "") + "hello") {
    case "":
        if (false && p1) {
            v8 := []int{-59, -823, 1, 127}
            fmt.Println(p0, p2)
            _ = v8
        } else {
            v9 := GenS1{F0: -1, F1: p0}
            _ = v9
        }
    case "a":
        p0 -= (^(p0))
    case "\n\t":
        if (2 == 0) {
            if (4294967295 != p0) {
                for fi10 := 0; fi10 < 13; fi10++ {
                    v11 := GenS1{F0: fi10, F1: 4294967295}
                    var v12 int64 = v7
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                p1 = (!false)
            } else {
                v13 := []int{}
                for fi14 := 0; fi14 < 13; fi14++ {
                    if len(v13) > 0 {
                        v13[0] = (fi14 + (fi14 & (fi14 | (fi14 >> ((18446744073709551615) & 63)))))
                    }
                    fmt.Println(p2, p0, fi14)
                    v15 := GenS2{F0: p1, F1: -527, F2: 127}
                    _ = fi14
                    _ = v15
                }
                _ = v13
            }
            var v16 int32 = -1
            var v17 int = 0
            _ = v16
            _ = v17
        } else {
            var v18 float32 = 0.0
            _ = v18
        }
    default:
        p2 = p2
    }
    v19 := GenS3{F0: 42, F1: v7}
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v7
    return 42
}

func main() {
    var v1 uint = 953
    v2 := make(map[int]bool)
    v3 := make(map[int]int)
    var v4 int16 = -911
    var v20 int = 32767
    var v21 uint64 = 18446744073709551615
    var v22 string = "a"
    fmt.Println(v1, v21)
    switch v20 {
    case 824:
        switch "foo bar" {
        case "foo bar":
            var v23 float32 = 1e10
            _ = v23
        case "a":
            v22 = (((v22 + v22) + ("\n\t" + ((v22 + "\n\t") + (v22 + ("hello" + (v22 + (v22 + v22))))))) + v22)
        }
    default:
        fmt.Println(v21)
    }
    var v24 string = (((("foo bar" + "\n\t") + "\n\t") + v22) + "")
    v25 := GenS4{F0: (v1 << ((946) & 63)), F1: 14}
    if (((818 == int(v4)) || true) || true) {
        var v26 uint16 = 65535
        for fi27 := 0; fi27 < 9; fi27++ {
            var v28 string = ((v24 + ((((v22 + ("a" + ("foo bar" + v24))) + v22) + "") + v22)) + (("a" + v22) + v22))
            for fi29 := 0; fi29 < 3; fi29++ {
                v22 = (v24 + "hello")
                fmt.Println(fi29, v20)
                v30 := []int{(fi29 * fi27)}
                v31 := GenS5{F0: 0}
                var v32 int16 = 1
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
            }
            _ = fi27
            _ = v28
        }
        _ = v26
    } else {
        for fi33 := 0; fi33 < 16; fi33++ {
            v34 := GenS3{F0: 0, F1: -9223372036854775808}
            _ = fi33
            _ = v34
        }
        for fi35 := 0; fi35 < 4; fi35++ {
            v4 = (v4 | v4)
            _ = fi35
        }
    }
    var v36 uint64 = (v21 | (^(v21)))
    v36 = v21
    v37 := GenS6{F0: 42, F1: -450, F2: 7}
    v24 = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
