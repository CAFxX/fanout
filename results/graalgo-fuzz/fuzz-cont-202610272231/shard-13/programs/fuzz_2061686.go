package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int32 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS4 struct { F0 int32 }

func main() {
    var v1 uint64 = 0
    var v2 int16 = -146
    var v3 int8 = 1
    var v4 uint16 = 0
    v5 := make(map[uint32]int)
    v6 := make(map[int]int)
    switch ("a" + "\n\t") {
    case "\n\t":
        v6[4294967295] = 0
    case "hello":
        v7 := GenS1{F0: 1, F1: -1, F2: 42}
        _ = v7
    case "":
        for fi8 := 0; fi8 < 17; fi8++ {
            switch fi8 {
            case -1:
                v4 -= (v4 & (v4 ^ (v4 / ((1) | 1))))
            case 940:
                fmt.Println(v3, v2, v2)
            case -723:
                switch "hello" {
                case "":
                    var v9 int64 = -9223372036854775808
                    _ = v9
                case "hello":
                    v10 := GenS1{F0: 311, F1: -697, F2: int32(fi8)}
                    _ = v10
                case "a":
                    var v11 int16 = v2
                    _ = v11
                }
            default:
                v4 = v4
            }
            _ = fi8
        }
    default:
        var v12 int64 = int64(v1)
        _ = v12
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := []int{42}
        v15 := GenS1{F0: 42, F1: -2147483648, F2: 37}
        fmt.Println(v1)
        _ = fi13
        _ = v14
        _ = v15
    }
    if (1e10 != 1e10) {
        for fi16 := 0; fi16 < 2; fi16++ {
            if (true || true) {
                var v17 int16 = 411
                v18 := GenS1{F0: 0, F1: 2147483647, F2: 11}
                _ = v17
                _ = v18
            }
            fmt.Println(v4, v3, v2)
            v19 := []int{int(v4), fi16}
            _ = fi16
            _ = v19
        }
        var v20 uint = 42
        fmt.Println(v1, v20)
        _ = v20
    } else {
        v4 -= (v4 * 2)
        v21 := GenS1{F0: 930, F1: 42, F2: 2147483647}
        _ = v21
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        v3 += v3
        _ = fi22
    }
    if true {
        for fi23 := 0; fi23 < 17; fi23++ {
            var v24 uint = 378
            _ = fi23
            _ = v24
        }
        if (514 > 1) {
            v25 := v5[1]
            if true {
                if false {
                    v5[646] = 1
                }
            } else {
                v5[uint32(v1)] = (v25 / ((int(v2)) | 1))
                fmt.Println(v1, v4, v1)
                var v26 float32 = 0.0
                _ = v26
            }
            _ = v25
        }
        v27 := GenS2{F0: false, F1: (v4 & v4)}
        _ = v27
    }
    var v28 uint16 = 42
    if ("a" != (("a" + "hello") + "")) {
        fmt.Println(v28)
        fmt.Println(v1, v3)
        fmt.Println(v1, v1, v3)
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        switch (fi29 - (fi29 | 378)) {
        case 127:
            var v30 uint16 = 42
            _ = v30
        case 679:
            for fi31 := 0; fi31 < 17; fi31++ {
                v4 += v4
                v32 := GenS1{F0: 19, F1: 1, F2: 2147483647}
                v1 -= v1
                v33 := GenS2{F0: false, F1: 42}
                _ = fi31
                _ = v32
                _ = v33
            }
        }
        _ = fi29
    }
    v34 := []int{}
    if (392 != 1) {
        if (0 < -9223372036854775808) {
            if false {
                v35 := GenS2{F0: true, F1: v4}
                v28 -= v28
                v36 := GenS3{F0: 255, F1: -45, F2: v4}
                _ = v35
                _ = v36
            } else {
                v37 := GenS4{F0: -2147483648}
                var v38 string = "\n\t"
                v1 = 209
                _ = v37
                _ = v38
            }
            fmt.Println(v28, v28)
            _ = len(v34)
        }
        v28 -= v4
        var v39 bool = (!(!(!false)))
        _ = v39
    } else {
        var v40 float64 = 0.1
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
