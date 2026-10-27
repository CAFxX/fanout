package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int32; F1 bool; F2 int }
func genFn0(p0 int64, p1 uint) string {
    if true {
        var v7 int = -310
        _ = v7
    } else {
        p0 = p0
    }
    v8 := GenS1{F0: 65535}
    p1 = p1
    v9 := GenS2{F0: ("hello" + ("" + "\n\t")), F1: ((-(p0)) != p0)}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return "\n\t"
}
func genFn1() uint16 {
    var v10 int = 4294967295
    var v11 uint32 = 42
    if true {
        v11 = v11
        v12 := GenS3{F0: 194, F1: "foo bar"}
        _ = v12
    }
    _ = v10
    _ = v11
    return 0
}

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 1
    v3 := make(map[uint32]uint32)
    v4 := []uint32{534, 65535, 2, 2}
    var v5 float64 = (-2.25 + -2.25)
    var v6 int32 = -1
    if (genFn1() <= 1) {
        v13 := GenS2{F0: "a", F1: true}
        _ = v13
    }
    fmt.Println(v5)
    v5 -= v5
    v5 = (v5 - v5)
    var v14 uint64 = 18446744073709551615
    v6 = 356
    v6 = (v6 + (v6 - -161))
    var v15 float32 = 1.5
    switch ("foo bar" + "hello") {
    case "":
        fmt.Println(v15)
    }
    var v16 uint64 = 4294967295
    v2 += v2
    var v17 string = "hello"
    switch "a" {
    case "hello":
        v5 -= v5
    case "foo bar":
        for fi18 := 0; fi18 < 12; fi18++ {
            if true {
                var v19 int16 = 42
                _ = v19
            } else {
                var v20 int = (^(fi18))
                switch (v17 + "foo bar") {
                case "hello":
                    var v21 bool = ((v5 != float64(v15)) && true)
                    _ = v21
                case "":
                    v22 := GenS4{F0: -128}
                    _ = v22
                default:
                    var v23 int64 = 9223372036854775807
                    _ = v23
                }
                v24 := GenS5{F0: v17}
                _ = v20
                _ = v24
            }
            v25 := v1[fi18]
            var v26 string = v17
            _ = fi18
            _ = v25
            _ = v26
        }
    default:
        if false {
            for fi27 := 0; fi27 < 18; fi27++ {
                v15 = 1.5
                for fi28 := 0; fi28 < 6; fi28++ {
                    var v29 int64 = v2
                    fmt.Println(v6, fi28)
                    _ = fi28
                    _ = v29
                }
                _ = fi27
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch 32768 {
            case 0:
                var v30 uint16 = genFn1()
                _ = v30
            case -635:
                v31 := GenS6{F0: (v6 >> ((402) & 63)), F1: (false && true), F2: 773}
                _ = v31
            default:
                _ = len(v4)
            }
        }
    }
    var v32 int16 = -32768
    v14 = 4294967295
    v17 = v17
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
