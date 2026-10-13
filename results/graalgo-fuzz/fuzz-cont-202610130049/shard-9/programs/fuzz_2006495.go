package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 uint8 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 2147483648
    v4 := make(map[int]int)
    switch "hello" {
    case "":
        switch 0 {
        case 797:
            for fi5 := 0; fi5 < 15; fi5++ {
                v3 += v3
                _ = fi5
            }
        default:
            if ((!false) || ((0.1 != -2.25) && true)) {
                v6 := []int{-888}
                v3 -= v3
                var v7 int32 = -196
                _ = v6
                _ = v7
            }
        }
    case "hello":
        var v8 int16 = 0
        _ = v8
    case "\n\t":
        var v9 uint64 = 416
        _ = v9
    default:
        fmt.Println(v3)
    }
    fmt.Println(v3)
    var v10 int8 = -106
    if (true || true) {
        v11 := GenS1{F0: "a", F1: 234}
        v3 = (^(v3))
        switch 32767 {
        case 4294967295:
            v12 := GenS2{F0: 1000, F1: v3}
            _ = v12
        case -937:
            v3 -= v3
        case -629:
            for fi13 := 0; fi13 < 17; fi13++ {
                v14 := []int{-438, fi13, fi13}
                delete(v1, (fi13 % ((fi13) | 1)))
                _ = fi13
                _ = v14
            }
        }
        _ = v11
    }
    var v15 uint16 = 947
    for fi16 := 0; fi16 < 15; fi16++ {
        v17 := []int{(fi16 & -1)}
        var v18 uint16 = v15
        fmt.Println(v3, fi16)
        _ = fi16
        _ = v17
        _ = v18
    }
    if (549 != -917) {
        switch 1 {
        case -9223372036854775808:
            fmt.Println(v10)
        case 1:
            v15 -= (v15 / (((v15 - v15)) | 1))
        case 4294967295:
            v19 := GenS3{F0: (((("" + ("a" + "foo bar")) + ("" + "foo bar")) + ("hello" + "\n\t")) + "foo bar")}
            _ = v19
        default:
            var v20 int16 = 32767
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        if true {
            for fi22 := 0; fi22 < 17; fi22++ {
                var v23 string = (((("foo bar" + "") + "") + "") + "foo bar")
                v10 = 42
                _ = fi22
                _ = v23
            }
        } else {
            for fi24 := 0; fi24 < 16; fi24++ {
                v25 := GenS4{F0: false, F1: 13}
                var v26 float64 = 1.5
                v27 := v4[(^(fi21))]
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
            }
            v3 -= (v3 % ((v3) | 1))
            for fi28 := 0; fi28 < 7; fi28++ {
                v10 += 27
                v29 := []int{(fi21 % ((0) | 1)), fi21, (fi21 / ((4294967295) | 1))}
                v30 := GenS4{F0: (false || true), F1: 127}
                _ = fi28
                _ = v29
                _ = v30
            }
        }
        _ = fi21
    }
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
