package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 bool; F2 uint8 }
func genFn0() bool {
    var v6 uint8 = 2
    v6 += 2
    switch "" {
    case "foo bar":
        v7 := GenS1{F0: 9223372036854775807, F1: 255}
        _ = v7
    case "hello":
        v8 := []int{}
        _ = v8
    default:
        fmt.Println(v6)
    }
    if (2147483648 != 939) {
        switch (("a" + "\n\t") + "foo bar") {
        case "hello":
            for fi9 := 0; fi9 < 16; fi9++ {
                v6 = 115
                _ = fi9
            }
        }
    } else {
        var v10 int = 128
        if (false && (!((107 <= 255) || true))) {
            fmt.Println(v10, v10)
            var v11 int64 = -1
            _ = v11
        } else {
            v6 = (v6 - (v6 - (^(v6))))
            v12 := GenS2{F0: 628, F1: 42}
            _ = v12
        }
        if (1 != 9223372036854775807) {
            var v13 uint = 581
            var v14 uint = 65535
            var v15 uint16 = 685
            _ = v13
            _ = v14
            _ = v15
        }
        _ = v10
    }
    _ = v6
    return false
}
func genFn1(p0 uint, p1 int8) int32 {
    var v16 bool = false
    fmt.Println(p0, v16)
    _ = p0
    _ = p1
    _ = v16
    return 2147483647
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = -9223372036854775808
    v3 := []int{}
    var v4 uint16 = 303
    var v5 string = "hello"
    switch (^(v2)) {
    case 0:
        for fi17 := 0; fi17 < 13; fi17++ {
            var v18 int32 = 1
            _ = fi17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 float64 = -2.25
        _ = fi19
        _ = v20
    }
    var v21 int = (^(v2))
    if false {
        for fi22 := 0; fi22 < 19; fi22++ {
            var v23 uint = 4294967295
            v24 := GenS2{F0: 831, F1: 9223372036854775807}
            _ = fi22
            _ = v23
            _ = v24
        }
        switch v2 {
        case 4294967295:
            var v25 string = ((v5 + "a") + v5)
            _ = v25
        case -117:
            v26 := GenS1{F0: 0, F1: (v4 | (v4 * v4))}
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        v28 := GenS3{F0: 396}
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        var v30 float32 = 3.14159
        _ = fi29
        _ = v30
    }
    if (1 != 924) {
        _ = len(v1)
    }
    v2 = v21
    v5 = v5
    v1[(-(v21))] = 1
    var v31 uint = uint(v4)
    for fi32 := 0; fi32 < 10; fi32++ {
        v33 := GenS4{F0: false, F1: true, F2: uint8(v31)}
        if true {
            for fi34 := 0; fi34 < 10; fi34++ {
                var v35 int8 = -67
                var v36 int64 = 1
                _ = fi34
                _ = v35
                _ = v36
            }
            var v37 uint8 = 127
            for fi38 := 0; fi38 < 4; fi38++ {
                v2 -= (^(v21))
                _ = fi38
            }
            _ = v37
        }
        _ = fi32
        _ = v33
    }
    var v39 uint = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v39
    _ = v4
    _ = v5
}
