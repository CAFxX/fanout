package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int32; F1 uint8 }
func genFn0(p0 int64) bool {
    var v4 float32 = 0.1
    if true {
        var v5 uint8 = 43
        _ = v5
    } else {
        v4 = float32((1.5 / -2.25))
    }
    v6 := GenS1{F0: 42, F1: 1, F2: 1}
    _ = p0
    _ = v4
    _ = v6
    return (!false)
}
func genFn1() uint64 {
    v7 := GenS1{F0: 65535, F1: 2, F2: 2147483647}
    var v8 float64 = 1e10
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 uint8 = 0
            fmt.Println(fi10)
            fmt.Println(v11)
            _ = fi10
            _ = v11
        }
        fmt.Println(fi9)
        _ = fi9
    }
    var v12 int = 9223372036854775807
    _ = v12
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = -222
    var v3 float32 = 0.0
    v2 = (v2 & (v2 / (((-(v2))) | 1)))
    if (true && false) {
        var v13 int64 = 9223372036854775807
        if true {
            for fi14 := 0; fi14 < 5; fi14++ {
                _ = len(v1)
                fmt.Println(v13)
                if (v2 == v2) {
                    var v15 uint16 = 65535
                    v13 = 1
                    _ = v15
                }
                _ = fi14
            }
            v16 := GenS2{F0: 65535, F1: 281}
            v17 := GenS1{F0: 32767, F1: 857, F2: (v2 + v2)}
            _ = v16
            _ = v17
        }
        switch 32767 {
        case -9223372036854775808:
            if (65535 != 385) {
                v18 := GenS2{F0: 2, F1: 32767}
                _ = v18
            }
        }
        _ = v13
    }
    switch 42 {
    case 9223372036854775807:
        var v19 int = 127
        _ = v19
    case 1:
        v20 := []int{-9223372036854775808, 32767, -9223372036854775808, 32768}
        _ = v20
    }
    switch "hello" {
    case "foo bar":
        v2 -= (v2 * v2)
    case "a":
        for fi21 := 0; fi21 < 6; fi21++ {
            for fi22 := 0; fi22 < 1; fi22++ {
                var v23 int = (^(fi21))
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        if true {
            var v25 uint8 = 42
            v3 = float32((float64(1e10) - 1e10))
            v26 := v1[(("" + "foo bar") + "a")]
            _ = v25
            _ = v26
        }
        if true {
            var v27 uint8 = 1
            _ = v27
        }
        _ = fi24
    }
    for fi28 := 0; fi28 < 20; fi28++ {
        v3 -= (v3 + 1e10)
        v3 = 0.0
        for fi29 := 0; fi29 < 10; fi29++ {
            var v30 int8 = -42
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    v31 := GenS2{F0: 39, F1: 32767}
    var v32 int16 = -230
    switch "\n\t" {
    case "hello":
        v32 = (v32 << ((18446744073709551615) & 63))
    case "foo bar":
        var v33 string = "a"
        _ = v33
    case "\n\t":
        var v34 bool = true
        _ = v34
    }
    v35 := GenS1{F0: 1, F1: 617, F2: (^(v2))}
    v36 := []int{-9223372036854775808, 1, -2147483648}
    v32 = v32
    if (true || (!(163 == 2))) {
        fmt.Println(v3)
        fmt.Println(v3)
    } else {
        var v37 int32 = v2
        var v38 int16 = (v32 + -1)
        var v39 int8 = 1
        _ = v37
        _ = v38
        _ = v39
    }
    if true {
        v32 = v32
        var v40 string = ""
        _ = v40
    }
    var v41 uint = 1
    if true {
        v42 := GenS3{F0: (v2 & -2147483648), F1: 0}
        v2 = 374
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v35
    _ = v36
    _ = v41
}
