package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint8 }
type GenS2 struct { F0 int32; F1 int64; F2 uint32 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 uint8; F1 uint64; F2 int16 }
func genFn0() string {
    var v4 int8 = 0
    switch 42 {
    case 975:
        v4 += (v4 + 30)
    case 4294967295:
        var v5 uint32 = 0
        _ = v5
    case 32767:
        var v6 int16 = 1
        _ = v6
    }
    _ = v4
    return "\n\t"
}
func genFn1(p0 int16, p1 int16) int64 {
    var v7 int64 = 9223372036854775807
    if false {
        var v8 int64 = v7
        _ = v8
    } else {
        switch 32767 {
        case -723:
            var v9 uint64 = 18446744073709551615
            _ = v9
        case 2147483647:
            var v10 int = 2147483647
            _ = v10
        case 0:
            var v11 float32 = (-2.25 + 0.0)
            _ = v11
        }
        switch "a" {
        case "a":
            var v12 uint64 = 10
            _ = v12
        default:
            for fi13 := 0; fi13 < 9; fi13++ {
                v14 := []int{(fi13 + (-(fi13))), fi13, fi13}
                v15 := GenS1{F0: 2, F1: "\n\t", F2: 255}
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        var v16 uint = 2
        _ = v16
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        var v18 uint16 = 65535
        _ = fi17
        _ = v18
    }
    var v19 int64 = (v7 & v7)
    _ = p0
    _ = p1
    _ = v19
    _ = v7
    return -463
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{("foo bar" + "hello"), "\n\t", ("hello" + ""), ((("foo bar" + "\n\t") + "hello") + ("hello" + "foo bar"))}
    v3 := make(map[uint32]bool)
    v20 := GenS2{F0: 0, F1: 9223372036854775807, F2: 595}
    var v21 string = ""
    var v22 uint8 = 114
    v23 := GenS3{F0: 1, F1: (v21 + (v21 + v21))}
    var v24 bool = ("hello" >= v21)
    fmt.Println(v22)
    var v25 int8 = 90
    v26 := GenS4{F0: 0, F1: 947, F2: 0}
    var v27 uint16 = 255
    for fi28 := 0; fi28 < 6; fi28++ {
        if (122 < 42) {
            for fi29 := 0; fi29 < 8; fi29++ {
                v24 = (v27 >= 2)
                v25 = (v25 ^ v25)
                _ = fi29
            }
        } else {
            if ((1.5 <= (3.14159 - float64(-2.25))) && (v27 != (v27 & v27))) {
                if (v24 && v24) {
                    switch (v21 + v21) {
                    case "\n\t":
                        var v30 int = (fi28 * fi28)
                        _ = v30
                    case "foo bar":
                        v22 = 255
                    default:
                        if v24 {
                            var v31 int8 = int8(fi28)
                            v24 = (!(((2147483647 < -1) && (-333 == fi28)) && (fi28 > 32768)))
                            var v32 int32 = 146
                            _ = v31
                            _ = v32
                        }
                    }
                    if (!true) {
                        v25 = (-(v25))
                    } else {
                        v33 := GenS1{F0: v27, F1: (v21 + v21), F2: 2}
                        var v34 uint64 = 2
                        _ = v33
                        _ = v34
                    }
                }
                _ = len(v2)
                if len(v2) > 0 {
                    v2[0] = v21
                }
            }
            var v35 uint64 = 262
            _ = v35
        }
        var v36 int8 = v25
        if (!((v21 + "foo bar") == v21)) {
            var v37 uint8 = (v22 * v22)
            v36 -= 120
            _ = v37
        }
        _ = fi28
        _ = v36
    }
    v21 = v21
    v2 = append(v2, "hello")
    _ = len(v1)
    _ = len(v2)
    var v38 float32 = -2.25
    if (2 != 541) {
        if false {
            var v39 uint16 = 0
            v2 = append(v2, v21)
            var v40 int16 = 1
            _ = v39
            _ = v40
        }
        var v41 uint64 = 0
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v38
}
