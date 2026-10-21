package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 int; F2 int32 }
type GenS3 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8; F1 int8 }
func genFn0(p0 int, p1 int) int64 {
    fmt.Println(p1)
    if false {
        var v6 uint8 = 202
        for fi7 := 0; fi7 < 19; fi7++ {
            p1 -= (fi7 & (^(p0)))
            var v8 int16 = 0
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    var v9 uint64 = 132
    switch (p1 + p0) {
    case 1:
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 uint32 = 65535
            v9 += (v9 + v9)
            _ = fi10
            _ = v11
        }
    }
    v12 := GenS1{F0: 2}
    _ = p0
    _ = p1
    _ = v12
    _ = v9
    return 42
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[string]string)
    var v3 uint = 725
    v4 := make(map[uint32]string)
    v5 := []string{"foo bar", ("" + "foo bar")}
    if true {
        var v13 uint64 = 2
        var v14 uint64 = (v13 * v13)
        _ = v13
        _ = v14
    } else {
        v15 := GenS2{F0: ((v1 * float32(((-2.25 + (0.0 + 1.5)) * 1e10))) > v1), F1: -1, F2: 612}
        _ = v15
    }
    if (4294967295 != 2147483648) {
        v1 += v1
    } else {
        var v16 uint32 = 0
        switch ("a" + "hello") {
        case "foo bar":
            v3 = v3
        case "a":
            v3 = (^(v3))
        }
        var v17 int32 = -1
        _ = v16
        _ = v17
    }
    if true {
        switch -728 {
        case 32768:
            v3 -= v3
        }
    } else {
        var v18 int8 = 66
        if true {
            if len(v5) > 0 {
                v5[0] = (("a" + "") + "")
            }
            for fi19 := 0; fi19 < 20; fi19++ {
                for fi20 := 0; fi20 < 8; fi20++ {
                    v3 -= (v3 * (v3 << ((v3) & 63)))
                    v5 = append(v5, ("foo bar" + ("hello" + "foo bar")))
                    _ = fi20
                }
                v21 := GenS3{F0: 95, F1: 128, F2: (v18 << (((v3 * v3)) & 63))}
                _ = fi19
                _ = v21
            }
            for fi22 := 0; fi22 < 13; fi22++ {
                v1 += v1
                _ = fi22
            }
        } else {
            switch ("a" + "") {
            case "":
                v18 = 62
            }
        }
        _ = v18
    }
    if (!true) {
        for fi23 := 0; fi23 < 5; fi23++ {
            delete(v4, 395)
            var v24 uint64 = 18446744073709551615
            _ = fi23
            _ = v24
        }
        v25 := GenS4{F0: 117}
        var v26 string = "\n\t"
        _ = v25
        _ = v26
    } else {
        v27 := GenS2{F0: true, F1: 127, F2: -1}
        fmt.Println(v1, v1)
        _ = len(v5)
        _ = v27
    }
    for fi28 := 0; fi28 < 11; fi28++ {
        var v29 uint32 = uint32(v3)
        _ = fi28
        _ = v29
    }
    switch "" {
    case "foo bar":
        fmt.Println(v3)
    case "\n\t":
        v30 := GenS5{F0: 255, F1: -128}
        _ = v30
    case "hello":
        var v31 string = "\n\t"
        _ = v31
    default:
        _ = len(v4)
    }
    var v32 int32 = -1
    fmt.Println(v32, v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
