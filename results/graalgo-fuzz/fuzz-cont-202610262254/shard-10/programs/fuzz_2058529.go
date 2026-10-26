package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 int8 }
func genFn0() uint64 {
    v5 := GenS1{F0: 209, F1: 65535}
    var v6 int8 = 127
    var v7 uint64 = 2
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 string = ("" + "foo bar")
        v10 := GenS1{F0: 722, F1: 2147483648}
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = v5
    _ = v6
    _ = v7
    return v7
}
func genFn1(p0 uint8) uint32 {
    v11 := GenS1{F0: 395, F1: 2147483648}
    fmt.Println(p0)
    v12 := GenS2{F0: 127, F1: 255}
    _ = p0
    _ = v11
    _ = v12
    return 520
}

func main() {
    var v1 float32 = ((1e10 - float32(0.0)) - -2.25)
    var v2 uint64 = 42
    var v3 bool = true
    v4 := make(map[uint32]uint32)
    switch 4294967295 {
    case 128:
        fmt.Println(v2, v1)
    case -961:
        fmt.Println(v2, v2, v3)
    case -210:
        switch (("foo bar" + "") + "foo bar") {
        case "hello":
            v13 := []int{9223372036854775807, 32767}
            _ = v13
        case "":
            fmt.Println(v3, v3)
        case "a":
            v14 := GenS1{F0: 2, F1: 2147483648}
            _ = v14
        }
    default:
        v2 = 0
    }
    v1 = (v1 * (1e10 - v1))
    var v15 uint8 = 2
    for fi16 := 0; fi16 < 9; fi16++ {
        switch 128 {
        case 518:
            if true {
                var v17 int8 = -1
                _ = v17
            }
        case 1:
            v1 = v1
        case 0:
            for fi18 := 0; fi18 < 8; fi18++ {
                v1 = 1e10
                _ = fi18
            }
        default:
            for fi19 := 0; fi19 < 18; fi19++ {
                v20 := []int{(fi19 ^ 0)}
                _ = fi19
                _ = v20
            }
        }
        v15 -= v15
        _ = fi16
    }
    if (570 < 32767) {
        v21 := GenS1{F0: 0, F1: 0}
        for fi22 := 0; fi22 < 18; fi22++ {
            v2 += v2
            var v23 uint8 = 1
            fmt.Println(fi22)
            _ = fi22
            _ = v23
        }
        if ((0.0 - float32(-2.25)) <= ((v1 / float32(1.5)) / v1)) {
            if ((-414 <= 42) || v3) {
                v24 := GenS3{F0: 0}
                v25 := []int{1}
                var v26 uint16 = 90
                _ = v24
                _ = v25
                _ = v26
            }
            v27 := v4[65535]
            _ = v27
        }
        _ = v21
    } else {
        v15 = v15
        v15 = v15
    }
    var v28 bool = (!(0 < 255))
    if v3 {
        fmt.Println(v1)
        v29 := []int{int(v2), 1, 9223372036854775807, 246}
        _ = v29
    }
    var v30 uint16 = 32767
    for fi31 := 0; fi31 < 18; fi31++ {
        for fi32 := 0; fi32 < 2; fi32++ {
            v28 = v28
            fmt.Println(v15)
            _ = fi32
        }
        var v33 int16 = 1
        _ = fi31
        _ = v33
    }
    switch ((("a" + "") + "hello") + "a") {
    case "hello":
        switch "a" {
        case "foo bar":
            v30 -= (v30 | v30)
        case "a":
            var v34 uint8 = 127
            _ = v34
        case "hello":
            fmt.Println(v1, v1)
        }
    case "foo bar":
        v1 += (((v1 - (v1 / 1e10)) / -2.25) / ((v1 * (v1 * v1)) - v1))
    default:
        v35 := []int{32768, 32768, -671, int(v30)}
        _ = v35
    }
    v36 := v4[65535]
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v36
    _ = v4
}
