package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint32; F1 bool; F2 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 uint; F2 uint8 }
type GenS6 struct { F0 bool; F1 uint64; F2 bool }

func main() {
    var v1 int32 = 42
    var v2 float32 = 0.0
    v3 := make(map[int]bool)
    var v4 bool = false
    v5 := make(map[uint32]bool)
    var v6 uint64 = 42
    fmt.Println(v4, v6)
    if v4 {
        switch "a" {
        case "\n\t":
            if v4 {
                var v7 float64 = 3.14159
                var v8 uint8 = 1
                v9 := []int{int(v6), 127, 32767, 0}
                _ = v7
                _ = v8
                _ = v9
            }
        case "foo bar":
            v10 := v5[42]
            _ = v10
        case "":
            v4 = false
        }
        v11 := GenS1{F0: (v1 * v1)}
        _ = v11
    }
    switch "a" {
    case "":
        v2 -= ((v2 / (((0.0 - v2) - 1.5) * float32(0.1))) * v2)
    case "\n\t":
        if ((!(v4 && (!(0.0 == 3.14159)))) || (((0.0 - 1e10) != 0.1) || (!v4))) {
            fmt.Println(v1, v6, v4)
            switch 42 {
            case 128:
                if ((!(915 != 32768)) || (127 < 78)) {
                    var v12 uint32 = 828
                    _ = v12
                }
            default:
                var v13 int16 = 201
                _ = v13
            }
            var v14 float64 = float64(((v2 - 3.14159) / v2))
            _ = v14
        } else {
            var v15 int64 = 0
            switch 0 {
            case 4294967295:
                var v16 string = ("" + (("foo bar" + "a") + ""))
                _ = v16
            case -1:
                v6 -= 0
            case 32767:
                switch ("hello" + "foo bar") {
                case "foo bar":
                    if (2 == 2) {
                        v4 = v4
                        v5[65535] = v4
                    }
                case "hello":
                    switch 641 {
                    case 209:
                        v17 := GenS1{F0: (v1 - int32(v15))}
                        _ = v17
                    default:
                        if v4 {
                            v18 := GenS2{F0: (("hello" != "hello") || (v2 != v2)), F1: 425}
                            v1 += -478
                            if (!v4) {
                                v6 = (v6 | v6)
                            }
                            _ = v18
                        } else {
                            var v19 uint64 = v6
                            _ = v19
                        }
                    }
                }
            default:
                var v20 uint = 128
                _ = v20
            }
            _ = v15
        }
    }
    switch 750 {
    case 696:
        var v21 uint = 685
        _ = v21
    case 128:
        for fi22 := 0; fi22 < 11; fi22++ {
            v2 = (-2.25 + (0.0 * v2))
            v23 := []int{}
            _ = fi22
            _ = v23
        }
    case 4294967295:
        if (!(132 <= 973)) {
            v24 := GenS3{F0: 267, F1: (!((v6 & v6) < 1)), F2: 42}
            v25 := GenS4{F0: (v1 + (v1 ^ (v1 * 264)))}
            _ = v24
            _ = v25
        } else {
            fmt.Println(v6, v2, v2)
            v26 := GenS5{F0: 42, F1: 42, F2: 127}
            v27 := GenS3{F0: 2147483648, F1: (v4 || (!true)), F2: -1}
            _ = v26
            _ = v27
        }
    }
    if true {
        for fi28 := 0; fi28 < 3; fi28++ {
            switch -9223372036854775808 {
            case -1:
                v29 := []int{fi28, fi28}
                _ = v29
            case 0:
                var v30 uint64 = v6
                _ = v30
            }
            v6 += 18446744073709551615
            v31 := GenS3{F0: 0, F1: v4, F2: 609}
            _ = fi28
            _ = v31
        }
    }
    v2 -= -2.25
    var v32 uint8 = 128
    v33 := GenS6{F0: (true && v4), F1: v6, F2: (!v4)}
    v34 := []int{9223372036854775807, 2147483647, 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
