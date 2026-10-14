package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 string; F2 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 int16; F2 int }
func genFn0(p0 int64) int16 {
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := []int{(fi7 % ((fi7) | 1)), (fi7 * fi7), fi7, fi7}
        _ = fi7
        _ = v8
    }
    if true {
        var v9 uint32 = 4294967295
        fmt.Println(v9, p0)
        var v10 uint64 = 18446744073709551615
        _ = v10
        _ = v9
    }
    var v11 int64 = p0
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 uint64 = 1
        var v14 uint = 65535
        _ = fi12
        _ = v13
        _ = v14
    }
    v11 -= (-(v11))
    _ = p0
    _ = v11
    return 393
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 146
    v3 := []int{-815, -194}
    var v4 float64 = (1.5 / -2.25)
    var v5 uint16 = 0
    var v6 int32 = 2147483647
    switch "a" {
    case "a":
        var v15 float32 = float32(v4)
        _ = v15
    case "":
        var v16 uint8 = 42
        _ = v16
    }
    var v17 int16 = 1000
    if false {
        var v18 uint64 = 80
        v2 = v2
        for fi19 := 0; fi19 < 13; fi19++ {
            _ = len(v3)
            _ = fi19
        }
        _ = v18
    }
    delete(v1, 0)
    for fi20 := 0; fi20 < 14; fi20++ {
        var v21 int32 = (v6 / (((v6 & (v6 & 1000))) | 1))
        _ = fi20
        _ = v21
    }
    _ = len(v1)
    if (340 >= 0) {
        switch "foo bar" {
        case "hello":
            v22 := GenS1{F0: 1}
            _ = v22
        case "":
            for fi23 := 0; fi23 < 3; fi23++ {
                v24 := GenS2{F0: 42, F1: "hello", F2: "a"}
                var v25 int32 = (v6 * v6)
                fmt.Println(v17, v17)
                _ = fi23
                _ = v24
                _ = v25
            }
        case "\n\t":
            for fi26 := 0; fi26 < 3; fi26++ {
                fmt.Println(v17)
                fmt.Println(v4, v4, v4)
                fmt.Println(fi26, v6)
                _ = fi26
            }
        }
    }
    if true {
        v4 += 0.1
        v27 := GenS1{F0: 127}
        for fi28 := 0; fi28 < 14; fi28++ {
            fmt.Println(v6, fi28)
            v29 := GenS1{F0: 128}
            _ = fi28
            _ = v29
        }
        _ = v27
    } else {
        v30 := v1[2]
        _ = v30
    }
    switch 453 {
    case -784:
        switch "hello" {
        case "\n\t":
            var v31 string = ""
            _ = v31
        case "hello":
            if true {
                if (984 <= 4294967295) {
                    var v32 int8 = 83
                    var v33 bool = ((true || false) || false)
                    _ = v32
                    _ = v33
                } else {
                    fmt.Println(v2, v5)
                }
                for fi34 := 0; fi34 < 15; fi34++ {
                    v5 = 524
                    _ = fi34
                }
            } else {
                v35 := GenS2{F0: 2, F1: ("\n\t" + "a"), F2: ""}
                _ = v35
            }
        default:
            var v36 uint64 = 736
            _ = v36
        }
    case 32768:
        v37 := GenS3{F0: 2}
        _ = v37
    }
    for fi38 := 0; fi38 < 2; fi38++ {
        v17 -= v17
        v39 := GenS1{F0: (fi38 >> (((v2 - (v2 - v2))) & 63))}
        v6 = -1
        _ = fi38
        _ = v39
    }
    for fi40 := 0; fi40 < 14; fi40++ {
        for fi41 := 0; fi41 < 18; fi41++ {
            v42 := GenS4{F0: 34, F1: v17, F2: -2147483648}
            switch fi40 {
            case 327:
                var v43 int64 = -1
                _ = v43
            case 4294967295:
                v2 += v2
            case 614:
                v4 += (3.14159 * v4)
            }
            _ = fi41
            _ = v42
        }
        var v44 uint32 = 1
        _ = fi40
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
