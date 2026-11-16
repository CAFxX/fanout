package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 int; F2 string }
type GenS4 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS5 struct { F0 bool; F1 uint32 }
type GenS6 struct { F0 uint16; F1 uint8; F2 uint }
func genFn0() uint64 {
    v5 := []int{4294967295, -911, 9223372036854775807, 9223372036854775807}
    var v6 string = (("\n\t" + "a") + "foo bar")
    var v7 int64 = -1
    for fi8 := 0; fi8 < 5; fi8++ {
        for fi9 := 0; fi9 < 2; fi9++ {
            v10 := GenS1{F0: 128, F1: 42}
            if (!(4294967295 >= 2)) {
                fmt.Println(fi8, fi8)
            } else {
                var v11 uint64 = 18446744073709551615
                v12 := GenS2{F0: (576 >= 4294967295)}
                var v13 uint32 = 456
                v14 := GenS3{F0: 981, F1: 1, F2: v6}
                switch (^(fi9)) {
                case 9223372036854775807:
                    var v15 uint8 = 128
                    v13 += (v13 / (((v13 ^ v13)) | 1))
                    if ((v13 ^ 42) >= 4294967295) {
                        var v16 float64 = 1e10
                        _ = v16
                    } else {
                        v11 += (v11 + uint64(v13))
                    }
                    _ = v15
                case -52:
                    v17 := GenS1{F0: -447, F1: 42}
                    _ = v17
                default:
                    v13 = 817
                }
                _ = v11
                _ = v12
                _ = v13
                _ = v14
            }
            _ = fi9
            _ = v10
        }
        var v18 int16 = -32768
        _ = fi8
        _ = v18
    }
    v6 = v6
    _ = v5
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := []string{"\n\t", "hello"}
    var v2 string = "hello"
    v3 := make(map[string]int)
    var v4 int16 = 1
    _ = len(v3)
    var v19 uint32 = 616
    v20 := GenS1{F0: 32768, F1: (^(v4))}
    _ = len(v3)
    fmt.Println(v4, v2, v2)
    for fi21 := 0; fi21 < 14; fi21++ {
        var v22 int32 = 42
        var v23 uint64 = 845
        _ = fi21
        _ = v22
        _ = v23
    }
    var v24 int16 = (v4 & v4)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = ((v2 + (v2 + v2)) + "")
    var v25 int32 = 46
    if true {
        for fi26 := 0; fi26 < 15; fi26++ {
            for fi27 := 0; fi27 < 14; fi27++ {
                v28 := GenS4{F0: 42, F1: 2, F2: -128}
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
    } else {
        var v29 uint64 = 524
        var v30 bool = true
        switch 341 {
        case 4294967295:
            v25 += -74
        case -643:
            var v31 int32 = v25
            _ = v31
        }
        _ = v29
        _ = v30
    }
    _ = len(v3)
    var v32 uint8 = 127
    for fi33 := 0; fi33 < 13; fi33++ {
        v34 := GenS5{F0: true, F1: 4294967295}
        for fi35 := 0; fi35 < 11; fi35++ {
            var v36 int8 = -128
            var v37 int32 = v25
            v4 = (v24 / ((-819) | 1))
            var v38 float64 = 0.1
            _ = fi35
            _ = v36
            _ = v37
            _ = v38
        }
        v39 := GenS3{F0: 0, F1: (fi33 >> ((42) & 63)), F2: v2}
        _ = fi33
        _ = v34
        _ = v39
    }
    if (554 >= 449) {
        for fi40 := 0; fi40 < 13; fi40++ {
            v25 = 0
            v41 := GenS6{F0: 65535, F1: v32, F2: 1}
            _ = fi40
            _ = v41
        }
        switch "a" {
        case "\n\t":
            v42 := GenS1{F0: 1, F1: (v4 << ((0) & 63))}
            _ = v42
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v32
    _ = v4
}
