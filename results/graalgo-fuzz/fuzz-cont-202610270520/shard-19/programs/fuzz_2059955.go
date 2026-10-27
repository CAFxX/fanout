package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 uint64; F2 uint8 }
func genFn0(p0 uint, p1 int64) uint8 {
    if false {
        p0 -= p0
        switch "foo bar" {
        case "hello":
            for fi4 := 0; fi4 < 8; fi4++ {
                p0 = p0
                var v5 float64 = (0.0 - 1e10)
                _ = fi4
                _ = v5
            }
        default:
            if (-32768 > 1000) {
                v6 := GenS1{F0: p1}
                v7 := GenS1{F0: p1}
                _ = v6
                _ = v7
            }
        }
        v8 := GenS2{F0: 32767, F1: ("foo bar" + "")}
        _ = v8
    }
    var v9 int32 = 645
    var v10 uint = (p0 << ((p0) & 63))
    v9 = -434
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 81
}

func main() {
    v1 := []int{-9223372036854775808, 127, 42, 42}
    var v2 float64 = 0.1
    var v3 uint8 = 2
    v11 := GenS3{F0: 4294967295}
    if true {
        fmt.Println(v3, v3)
    }
    v3 = v3
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v12 int = 222
    var v13 int8 = 42
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 uint64 = 0
        switch (("hello" + "\n\t") + "hello") {
        case "foo bar":
            v16 := GenS2{F0: 970, F1: "hello"}
            _ = v16
        case "":
            for fi17 := 0; fi17 < 6; fi17++ {
                var v18 bool = true
                _ = fi17
                _ = v18
            }
        }
        _ = fi14
        _ = v15
    }
    var v19 string = "foo bar"
    if ((v13 % ((-1) | 1)) != v13) {
        v19 = v19
        v19 = "foo bar"
    } else {
        fmt.Println(v13, v13, v12)
        switch 127 {
        case -297:
            var v20 uint16 = 0
            _ = v20
        default:
            var v21 int = v12
            _ = v21
        }
    }
    fmt.Println(v2, v19)
    var v22 uint64 = 18446744073709551615
    for fi23 := 0; fi23 < 13; fi23++ {
        for fi24 := 0; fi24 < 16; fi24++ {
            v25 := GenS4{F0: 1, F1: (v22 | (v22 % ((v22) | 1))), F2: (v3 | 219)}
            v13 = v13
            v22 = v22
            _ = fi24
            _ = v25
        }
        var v26 uint32 = 1
        for fi27 := 0; fi27 < 15; fi27++ {
            var v28 uint8 = 255
            _ = fi27
            _ = v28
        }
        _ = fi23
        _ = v26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v22
    _ = v3
}
