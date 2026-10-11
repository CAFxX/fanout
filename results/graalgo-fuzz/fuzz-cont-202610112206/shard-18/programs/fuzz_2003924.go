package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint8; F1 int; F2 bool }
type GenS3 struct { F0 int32 }
func genFn0(p0 int64, p1 int16, p2 int8) uint32 {
    v4 := []int{}
    v5 := GenS1{F0: false, F1: p0}
    var v6 uint32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 4294967295
}
func genFn1(p0 bool, p1 uint16, p2 uint64) int64 {
    if ((false || p0) && p0) {
        p1 -= p1
        p1 -= p1
    } else {
        if p0 {
            v7 := GenS1{F0: false, F1: 1}
            v8 := []int{-2147483648, 0, -792}
            _ = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 4; fi9++ {
            fmt.Println(p0, fi9)
            var v10 uint8 = 127
            v11 := GenS1{F0: p0, F1: -9223372036854775808}
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    var v12 bool = (((3.14159 / (0.1 - ((((1.5 / 1.5) + 0.1) * float64(0.1)) - 1e10))) - (((1e10 - 0.1) - 1.5) * ((0.0 * 0.1) - 0.1))) > (-2.25 * float64((((0.0 / 0.1) * 3.14159) * -2.25))))
    switch 9223372036854775807 {
    case 48:
        var v13 uint64 = (^(p2))
        _ = v13
    case 2147483647:
        if (0 >= 716) {
            var v14 uint64 = (p2 >> ((65535) & 63))
            p0 = p0
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"\n\t", "\n\t", "\n\t"}
    v3 := make(map[uint32]bool)
    var v15 uint64 = 82
    v16 := GenS2{F0: 42, F1: 0, F2: true}
    v15 += v15
    fmt.Println(v15)
    var v17 uint = 18446744073709551615
    v18 := GenS2{F0: 2, F1: 666, F2: true}
    var v19 float64 = -2.25
    v15 = (v15 & v15)
    v17 -= (v17 - (v17 * v17))
    v19 -= float64(0.1)
    v20 := v1[0]
    for fi21 := 0; fi21 < 20; fi21++ {
        switch (v20 + v20) {
        case "":
            for fi22 := 0; fi22 < 9; fi22++ {
                switch ((v20 + ("hello" + "")) + "foo bar") {
                case "a":
                    var v23 int64 = 0
                    _ = v23
                default:
                    v3[genFn0(genFn1(((42 == 245) || true), 1, uint64(v17)), -32768, 100)] = (!(595 == 0))
                    v17 = (v17 & (v17 >> (((v17 + (v17 >> ((v17) & 63)))) & 63)))
                    v1[2147483647] = (v20 + ("" + ((((v20 + (v20 + ("" + "\n\t"))) + (v20 + "foo bar")) + ("a" + (v20 + "hello"))) + ((v20 + (v20 + (v20 + "foo bar"))) + ("hello" + "foo bar")))))
                    delete(v3, uint32(v15))
                }
                _ = fi22
            }
        case "\n\t":
            var v24 uint = (v17 ^ (v17 * 751))
            _ = v24
        case "foo bar":
            v25 := GenS3{F0: -2147483648}
            _ = v25
        default:
            switch (v20 + v20) {
            case "":
                v15 = v15
            }
        }
        _ = fi21
    }
    switch "hello" {
    case "":
        var v26 int = -1
        _ = v26
    default:
        var v27 int8 = 42
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
