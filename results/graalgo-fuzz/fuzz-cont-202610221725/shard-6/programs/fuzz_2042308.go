package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int16; F2 int }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string }
func genFn0(p0 int16, p1 int32, p2 uint64) uint64 {
    v6 := GenS1{F0: 2, F1: 65535}
    switch ("foo bar" + "\n\t") {
    case "":
        p1 += p1
    case "a":
        p1 = (p1 - p1)
    case "\n\t":
        switch ("" + "foo bar") {
        case "hello":
            v7 := []int{-9223372036854775808}
            _ = v7
        case "\n\t":
            p0 = 1000
        default:
            p1 += -65
        }
    default:
        var v8 uint32 = 4294967295
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p2
}
func genFn1(p0 int16, p1 uint8, p2 int8) string {
    v9 := GenS2{F0: 4294967295, F1: p0, F2: 0}
    if true {
        v10 := GenS3{F0: 0, F1: 65535}
        _ = v10
    }
    switch "\n\t" {
    case "a":
        p2 += p2
    }
    if false {
        if ((2 < 42) || (-890 == 9223372036854775807)) {
            v11 := []int{32768, -2147483648, 4294967295}
            var v12 uint16 = 1
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 5; fi13++ {
            v14 := []int{86}
            switch fi13 {
            case 127:
                var v15 uint16 = 253
                _ = v15
            }
            _ = fi13
            _ = v14
        }
    }
    fmt.Println(p0, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return "\n\t"
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]bool)
    var v3 uint8 = 146
    v4 := make(map[uint32]int)
    var v5 float32 = (0.1 / (float32(-2.25) / 1.0))
    v16 := GenS2{F0: 190, F1: -1, F2: -169}
    v5 -= (float32(float64(v5)) * v5)
    v17 := GenS4{F0: "hello"}
    if false {
        var v18 uint = uint(v3)
        var v19 uint64 = 2
        v20 := GenS3{F0: -536, F1: 42}
        _ = v18
        _ = v19
        _ = v20
    }
    var v21 float64 = 0.0
    switch 4 {
    case 32768:
        var v22 uint16 = 43
        _ = v22
    case 32767:
        v21 += 3.14159
    case -2147483648:
        v23 := []int{42, 4294967295, -693, 32767}
        _ = v23
    }
    if (!false) {
        var v24 int = 128
        _ = v24
    } else {
        var v25 int16 = 1000
        switch "a" {
        case "":
            v26 := GenS3{F0: 313, F1: 65535}
            _ = v26
        }
        if true {
            switch -153 {
            case 4294967295:
                v27 := GenS5{F0: ""}
                _ = v27
            case 2147483647:
                fmt.Println(v3, v25, v5)
            case 239:
                var v28 float32 = v5
                _ = v28
            }
            var v29 int64 = -1
            for fi30 := 0; fi30 < 5; fi30++ {
                for fi31 := 0; fi31 < 10; fi31++ {
                    v5 = v5
                    _ = fi31
                }
                _ = fi30
            }
            _ = v29
        } else {
            v32 := GenS3{F0: 1, F1: 65535}
            if (1 == 128) {
                fmt.Println(v25)
                var v33 int64 = -1
                _ = v33
            } else {
                var v34 int32 = -1
                _ = v34
            }
            for fi35 := 0; fi35 < 12; fi35++ {
                fmt.Println(v5, v21)
                v36 := []int{(-(fi35))}
                _ = fi35
                _ = v36
            }
            _ = v32
        }
        _ = v25
    }
    v37 := v4[2147483647]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
