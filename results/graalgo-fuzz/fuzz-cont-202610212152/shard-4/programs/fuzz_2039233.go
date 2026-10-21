package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint }
type GenS2 struct { F0 uint8; F1 int64; F2 int32 }
type GenS3 struct { F0 int64; F1 uint; F2 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8; F1 int32 }
func genFn0(p0 int32, p1 int64, p2 int) uint {
    switch "hello" {
    case "\n\t":
        switch (p2 + p2) {
        case -1:
            for fi4 := 0; fi4 < 14; fi4++ {
                v5 := []int{(fi4 * 127), fi4}
                _ = fi4
                _ = v5
            }
        case -943:
            v6 := GenS1{F0: 2147483648, F1: 115, F2: 755}
            _ = v6
        }
    case "":
        if true {
            fmt.Println(p1, p0, p2)
            var v7 int8 = -42
            var v8 float64 = 0.0
            _ = v7
            _ = v8
        } else {
            var v9 uint8 = 85
            _ = v9
        }
    case "foo bar":
        if true {
            p2 = (p2 >> ((18446744073709551615) & 63))
        } else {
            p0 = (p0 | 999)
        }
    }
    var v10 uint64 = 621
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 1
}
func genFn1(p0 uint8, p1 int64) int8 {
    var v11 int = -614
    if false {
        v12 := GenS2{F0: (p0 - p0), F1: 0, F2: -1}
        _ = v12
    } else {
        if false {
            for fi13 := 0; fi13 < 2; fi13++ {
                p0 -= p0
                fmt.Println(p0, fi13)
                _ = fi13
            }
            var v14 int16 = 1
            _ = v14
        }
        if true {
            v15 := GenS2{F0: (p0 << ((18446744073709551615) & 63)), F1: p1, F2: 0}
            var v16 uint8 = 1
            _ = v15
            _ = v16
        }
    }
    v17 := []int{}
    v11 -= (v11 | (v11 & v11))
    _ = p0
    _ = p1
    _ = v11
    _ = v17
    return 71
}

func main() {
    var v1 int = -9223372036854775808
    var v2 int8 = -97
    v3 := make(map[string]uint32)
    switch (v1 & (v1 | -9223372036854775808)) {
    case 128:
        v18 := GenS3{F0: 42, F1: 65535, F2: (^(v1))}
        _ = v18
    case -1:
        fmt.Println(v2)
    case 0:
        switch v1 {
        case 875:
            var v19 int = v1
            _ = v19
        case -9223372036854775808:
            _ = len(v3)
        default:
            fmt.Println(v2, v2)
        }
    }
    var v20 uint = 18446744073709551615
    for fi21 := 0; fi21 < 11; fi21++ {
        fmt.Println(v20, v1)
        v22 := []int{}
        _ = fi21
        _ = v22
    }
    var v23 uint64 = 42
    v24 := []int{555}
    if true {
        if true {
            v1 = v1
            v25 := GenS4{F0: (v23 % ((v23) | 1))}
            fmt.Println(v23, v2)
            _ = v25
        }
        v23 = v23
        var v26 int8 = (v2 ^ (v2 + v2))
        _ = v26
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        var v28 float64 = 1e10
        for fi29 := 0; fi29 < 19; fi29++ {
            var v30 uint8 = 42
            var v31 float32 = 3.14159
            if (-375 != -1) {
                var v32 string = "foo bar"
                _ = v32
            } else {
                switch ("a" + "foo bar") {
                case "\n\t":
                    v28 = 0.1
                case "foo bar":
                    var v33 int8 = (v2 - v2)
                    _ = v33
                case "a":
                    var v34 int = (^(v1))
                    _ = v34
                }
                v2 -= (^(v2))
                v31 = 3.14159
            }
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = len(v3)
        _ = fi27
        _ = v28
    }
    for fi35 := 0; fi35 < 16; fi35++ {
        fmt.Println(v2, v1)
        if false {
            var v36 uint64 = v23
            if (424 <= 375) {
                fmt.Println(v20)
            } else {
                var v37 string = "foo bar"
                v38 := GenS4{F0: v23}
                _ = v37
                _ = v38
            }
            var v39 uint16 = 2
            _ = v36
            _ = v39
        } else {
            for fi40 := 0; fi40 < 4; fi40++ {
                switch "" {
                case "foo bar":
                    v41 := GenS5{F0: 42, F1: -2147483648}
                    _ = v41
                case "\n\t":
                    var v42 uint = (^(v20))
                    _ = v42
                case "hello":
                    fmt.Println(fi40, v2, fi40)
                }
                v43 := GenS4{F0: v23}
                _ = fi40
                _ = v43
            }
        }
        _ = fi35
    }
    fmt.Println(v1, v20, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
}
