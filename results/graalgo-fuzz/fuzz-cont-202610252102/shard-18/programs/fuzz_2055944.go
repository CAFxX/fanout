package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint }
type GenS2 struct { F0 bool; F1 uint16; F2 int }
type GenS3 struct { F0 int64; F1 string }
func genFn0(p0 int64, p1 bool) uint64 {
    v7 := GenS1{F0: -58, F1: 65535, F2: 4294967295}
    p0 = p0
    p0 -= p0
    _ = p0
    _ = p1
    _ = v7
    return 901
}
func genFn1() uint16 {
    v8 := []int{}
    v9 := GenS2{F0: (1 < 0), F1: 32767, F2: 42}
    for fi10 := 0; fi10 < 9; fi10++ {
        if (true || false) {
            if false {
                var v11 uint16 = 32767
                switch "" {
                case "\n\t":
                    v11 += v11
                case "a":
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                case "hello":
                    var v12 float32 = 3.14159
                    _ = v12
                }
                v11 += (v11 >> ((42) & 63))
                _ = v11
            } else {
                _ = len(v8)
            }
            var v13 float32 = 1e10
            var v14 float64 = (3.14159 * 1e10)
            _ = v13
            _ = v14
        }
        _ = fi10
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 uint8 = 128
        v16 -= 0
        _ = fi15
        _ = v16
    }
    if true {
        if true {
            if true {
                var v17 uint16 = 65535
                v8 = append(v8, 32768)
                _ = v17
            } else {
                for fi18 := 0; fi18 < 18; fi18++ {
                    for fi19 := 0; fi19 < 19; fi19++ {
                        fmt.Println(fi19)
                        switch -445 {
                        case 1:
                            var v20 int64 = 787
                            _ = v20
                        default:
                            var v21 float64 = 1e10
                            _ = v21
                        }
                        _ = fi19
                    }
                    v22 := GenS1{F0: 0, F1: 223, F2: 2}
                    _ = fi18
                    _ = v22
                }
                var v23 int = 128
                _ = v23
            }
        } else {
            for fi24 := 0; fi24 < 3; fi24++ {
                v25 := GenS1{F0: 1, F1: 2147483647, F2: 2}
                var v26 int8 = -1
                _ = fi24
                _ = v25
                _ = v26
            }
        }
    }
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := []uint32{}
    var v2 uint16 = 267
    v3 := make(map[string]bool)
    v4 := make(map[uint32]uint32)
    var v5 uint8 = 177
    var v6 uint64 = 4294967295
    switch "\n\t" {
    case "a":
        v27 := GenS3{F0: -935, F1: ""}
        _ = v27
    case "foo bar":
        fmt.Println(v5, v6)
    default:
        v6 = 280
    }
    v5 = (v5 * (v5 & (^(v5))))
    var v28 float32 = float32((float64((1e10 / (((0.0 / 1e10) * (1.5 - (1e10 - (0.0 / 3.14159)))) - 1e10))) / 0.1))
    var v29 uint64 = v6
    for fi30 := 0; fi30 < 11; fi30++ {
        delete(v4, 2147483647)
        v1 = append(v1, uint32(v2))
        _ = fi30
    }
    fmt.Println(v28, v6)
    switch 758 {
    case 0:
        var v31 bool = (v29 >= v29)
        _ = v31
    }
    v32 := GenS3{F0: -1, F1: "\n\t"}
    var v33 uint16 = genFn1()
    if false {
        v2 += uint16(v6)
        var v34 uint32 = 65535
        switch "foo bar" {
        case "hello":
            switch -715 {
            case 128:
                v5 = 160
            case -9223372036854775808:
                for fi35 := 0; fi35 < 2; fi35++ {
                    if false {
                        v1 = append(v1, 1)
                        switch (("\n\t" + "\n\t") + ("" + "hello")) {
                        case "hello":
                            v36 := v3["foo bar"]
                            _ = v36
                        case "foo bar":
                            v34 += v34
                        case "\n\t":
                            v37 := GenS1{F0: 47, F1: 2147483648, F2: 1}
                            _ = v37
                        }
                    }
                    _ = fi35
                }
            }
        case "\n\t":
            fmt.Println(v6, v33)
        }
        _ = v34
    } else {
        v29 = v29
        v38 := GenS2{F0: false, F1: (v33 | v33), F2: 128}
        _ = v38
    }
    fmt.Println(v6, v2, v6)
    var v39 uint32 = 42
    var v40 int16 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
