package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8; F1 uint16; F2 string }

func main() {
    v1 := []int{-9223372036854775808, 2147483647, -9223372036854775808}
    var v2 uint32 = 95
    v3 := make(map[int]string)
    v4 := make(map[uint32]bool)
    v5 := make(map[int]int)
    v6 := make(map[uint32]string)
    if true {
        v5[32767] = 524
    } else {
        v2 = v2
        v2 = v2
        var v7 uint16 = 32767
        _ = v7
    }
    if false {
        var v8 uint64 = 481
        var v9 float64 = 1.5
        _ = v8
        _ = v9
    } else {
        v2 = v2
        v2 += v2
        v2 -= (v2 | 4294967295)
    }
    v10 := GenS1{F0: 32767, F1: 171, F2: "foo bar"}
    if true {
        if true {
            fmt.Println(v2)
            v11 := GenS1{F0: 32767, F1: 249, F2: "\n\t"}
            _ = v11
        } else {
            v12 := GenS1{F0: 2, F1: 42, F2: "hello"}
            var v13 int16 = -1
            _ = v12
            _ = v13
        }
    }
    fmt.Println(v2)
    v2 = (v2 | v2)
    v1 = append(v1, -752)
    v14 := GenS2{F0: 247}
    switch "a" {
    case "foo bar":
        v15 := GenS2{F0: 2}
        _ = v15
    case "a":
        v16 := GenS1{F0: 198, F1: 0, F2: ("hello" + "foo bar")}
        _ = v16
    }
    v17 := GenS3{F0: -1, F1: true}
    v2 = 0
    v18 := GenS4{F0: 32767, F1: 42}
    if (32767 >= 0) {
        if false {
            fmt.Println(v2)
            v2 += 828
            v19 := GenS5{F0: 217}
            _ = v19
        }
        v20 := GenS6{F0: 1, F1: 2, F2: ("a" + (("foo bar" + ("foo bar" + "\n\t")) + ""))}
        var v21 int8 = -1
        _ = v20
        _ = v21
    } else {
        for fi22 := 0; fi22 < 5; fi22++ {
            switch -2147483648 {
            case 42:
                var v23 string = "\n\t"
                _ = v23
            default:
                v24 := v3[(fi22 - fi22)]
                _ = v24
            }
            _ = fi22
        }
        var v25 int32 = -1
        _ = len(v1)
        _ = v25
    }
    for fi26 := 0; fi26 < 19; fi26++ {
        if false {
            switch 119 {
            case 9223372036854775807:
                v2 += (v2 ^ (v2 * 1))
            case -89:
                var v27 float64 = (1.5 - 1e10)
                _ = v27
            default:
                fmt.Println(v2, v2)
            }
            switch (fi26 - (fi26 * -9223372036854775808)) {
            case 127:
                delete(v5, fi26)
            case -1:
                var v28 uint16 = uint16(v2)
                _ = v28
            case 4294967295:
                v2 = (v2 % (((v2 & 65535)) | 1))
            default:
                v5[(^(fi26))] = fi26
            }
            v2 += v2
        }
        _ = fi26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
