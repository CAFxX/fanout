package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int16; F1 int64 }
func genFn0() uint8 {
    v4 := GenS1{F0: "hello"}
    if ((false || false) || true) {
        var v5 uint32 = 1
        switch 4294967295 {
        case 4294967295:
            for fi6 := 0; fi6 < 6; fi6++ {
                var v7 uint8 = 2
                for fi8 := 0; fi8 < 11; fi8++ {
                    if true {
                        v7 = (v7 & v7)
                        var v9 int = 128
                        v10 := GenS1{F0: "a"}
                        _ = v10
                        _ = v9
                    } else {
                        v11 := GenS1{F0: ("" + "")}
                        v12 := GenS1{F0: ("" + ("" + "foo bar"))}
                        _ = v11
                        _ = v12
                    }
                    v5 += 4294967295
                    v13 := GenS1{F0: ((("" + "hello") + "foo bar") + "hello")}
                    var v14 int8 = -57
                    _ = fi8
                    _ = v13
                    _ = v14
                }
                var v15 int64 = -1
                _ = fi6
                _ = v15
                _ = v7
            }
        case 2147483647:
            v16 := GenS1{F0: "hello"}
            _ = v16
        }
        var v17 uint32 = 504
        _ = v17
        _ = v5
    } else {
        var v18 uint16 = 32767
        _ = v18
    }
    var v19 int16 = -32768
    _ = v19
    _ = v4
    return 1
}
func genFn1(p0 uint8, p1 int64) uint64 {
    v20 := GenS1{F0: "\n\t"}
    var v21 uint = 392
    if false {
        var v22 int = 128
        switch "foo bar" {
        case "hello":
            var v23 bool = false
            _ = v23
        case "foo bar":
            v24 := []int{v22, (v22 - v22), v22}
            _ = v24
        case "":
            switch "" {
            case "hello":
                v22 -= v22
            case "\n\t":
                if false {
                    if false {
                        v25 := GenS1{F0: ("\n\t" + ("" + ("\n\t" + "\n\t")))}
                        v22 -= (v22 - (^(v22)))
                        _ = v25
                    }
                    switch 4294967295 {
                    case 42:
                        for fi26 := 0; fi26 < 17; fi26++ {
                            fmt.Println(fi26, p0)
                            _ = fi26
                        }
                    case -9223372036854775808:
                        v22 = v22
                    case -1:
                        var v27 int64 = (p1 % ((42) | 1))
                        _ = v27
                    default:
                        fmt.Println(p1, v22)
                    }
                    var v28 uint = (^(v21))
                    _ = v28
                }
            default:
                var v29 float32 = -2.25
                _ = v29
            }
        }
        _ = v22
    } else {
        var v30 uint32 = 0
        _ = v30
    }
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    return 2
}

func main() {
    var v1 uint8 = 78
    v2 := make(map[uint32]bool)
    v3 := []uint32{uint32(v1), 337, 448}
    for fi31 := 0; fi31 < 5; fi31++ {
        var v32 uint16 = 2
        _ = fi31
        _ = v32
    }
    if (!false) {
        var v33 uint64 = 228
        _ = v33
    }
    var v34 int16 = 0
    v35 := GenS2{F0: 2, F1: 2}
    v36 := GenS3{F0: (^(v34)), F1: 1}
    v37 := GenS2{F0: 4294967295, F1: 4294967295}
    v1 -= (v1 - genFn0())
    v34 = (v34 & (v34 - v34))
    fmt.Println(v1, v34)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
}
