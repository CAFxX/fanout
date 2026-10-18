package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 int64; F1 uint; F2 bool }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int64; F1 int32 }
func genFn0(p0 int64, p1 string) int {
    for fi4 := 0; fi4 < 16; fi4++ {
        for fi5 := 0; fi5 < 14; fi5++ {
            var v6 uint32 = 313
            var v7 int32 = -1
            _ = fi5
            _ = v6
            _ = v7
        }
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := []int{(fi8 + 0), 128}
            var v10 string = p1
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi4
    }
    p0 = (p0 | 1)
    v11 := []int{}
    switch 127 {
    case 127:
        fmt.Println(p0, p0)
    case 4294967295:
        v12 := GenS1{F0: -73, F1: 42, F2: 2147483648}
        _ = v12
    case 128:
        for fi13 := 0; fi13 < 9; fi13++ {
            p1 = p1
            _ = fi13
        }
    default:
        switch p1 {
        case "hello":
            if true {
                v14 := GenS2{F0: 1, F1: 273}
                v15 := GenS3{F0: 176, F1: 182}
                p0 -= (p0 % ((p0) | 1))
                _ = v14
                _ = v15
            }
        case "\n\t":
            if false {
                v16 := GenS4{F0: 0, F1: 100}
                p1 = "\n\t"
                v11 = append(v11, 4294967295)
                _ = v16
            } else {
                switch (("" + p1) + "hello") {
                case "hello":
                    var v17 int = 32768
                    _ = v17
                }
            }
        case "foo bar":
            for fi18 := 0; fi18 < 14; fi18++ {
                p1 = p1
                _ = fi18
            }
        default:
            v19 := GenS5{F0: p0, F1: 65535, F2: true}
            _ = v19
        }
    }
    _ = p0
    _ = p1
    _ = v11
    return 4294967295
}

func main() {
    var v1 string = (("hello" + "a") + ("hello" + "foo bar"))
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]uint32)
    var v20 string = "foo bar"
    var v21 float32 = 1e10
    v20 = v1
    v21 -= 0.0
    v22 := GenS6{F0: "\n\t"}
    var v23 bool = true
    if v23 {
        for fi24 := 0; fi24 < 19; fi24++ {
            v25 := v2[2147483648]
            v26 := GenS7{F0: 471}
            v23 = (65535 <= 255)
            _ = fi24
            _ = v25
            _ = v26
        }
    } else {
        v27 := []int{}
        var v28 int16 = 1
        _ = v27
        _ = v28
    }
    var v29 string = ("a" + "\n\t")
    v30 := GenS6{F0: v29}
    var v31 uint16 = 0
    v23 = v23
    v32 := GenS4{F0: 9223372036854775807, F1: -92}
    delete(v3, 737)
    for fi33 := 0; fi33 < 11; fi33++ {
        if (!(!v23)) {
            v34 := GenS2{F0: 2147483648, F1: 1}
            v20 = v29
            v21 -= v21
            _ = v34
        } else {
            fmt.Println(fi33, v20, v29)
            v1 = ("a" + v29)
        }
        v35 := GenS8{F0: 42, F1: 2147483647}
        v1 = v20
        _ = fi33
        _ = v35
    }
    var v36 int = int(v31)
    for fi37 := 0; fi37 < 3; fi37++ {
        v36 += v36
        _ = fi37
    }
    v3[42] = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v36
}
