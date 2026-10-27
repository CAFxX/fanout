package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64; F1 uint; F2 bool }
type GenS5 struct { F0 string; F1 string }

func main() {
    v1 := []int{2147483647}
    var v2 uint8 = 255
    v3 := []int{0, -2147483648}
    if true {
        var v4 uint = 893
        v5 := GenS1{F0: 18446744073709551615, F1: 414, F2: -54}
        _ = v4
        _ = v5
    }
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 int8 = 127
        var v8 uint64 = 779
        _ = fi6
        _ = v7
        _ = v8
    }
    v2 = 8
    var v9 string = ""
    if false {
        if true {
            v1 = append(v1, -9223372036854775808)
            var v10 float64 = -2.25
            _ = v10
        }
    }
    var v11 string = ("a" + (v9 + "hello"))
    for fi12 := 0; fi12 < 2; fi12++ {
        v13 := GenS2{F0: 471, F1: (fi12 / (((fi12 + fi12)) | 1))}
        v11 = v9
        var v14 string = "\n\t"
        _ = fi12
        _ = v13
        _ = v14
    }
    v15 := GenS3{F0: 4294967295}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v16 int = -1
    _ = len(v3)
    var v17 int16 = 0
    for fi18 := 0; fi18 < 18; fi18++ {
        v16 = (fi18 << ((63) & 63))
        var v19 uint = 53
        v9 = v11
        _ = fi18
        _ = v19
    }
    v1 = append(v1, v16)
    switch (-(v16)) {
    case -717:
        v9 = "a"
    case 128:
        for fi20 := 0; fi20 < 12; fi20++ {
            v11 = "foo bar"
            for fi21 := 0; fi21 < 6; fi21++ {
                v22 := GenS3{F0: 2}
                var v23 uint32 = 65535
                v17 -= (v17 * (v17 ^ v17))
                _ = fi21
                _ = v22
                _ = v23
            }
            if true {
                v24 := GenS3{F0: 0}
                v25 := GenS4{F0: -9223372036854775808, F1: 848, F2: (false || true)}
                var v26 float64 = 1.5
                _ = v24
                _ = v25
                _ = v26
            } else {
                var v27 int64 = 220
                _ = v27
            }
            _ = fi20
        }
    case -2147483648:
        for fi28 := 0; fi28 < 15; fi28++ {
            fmt.Println(v9, fi28)
            _ = fi28
        }
    default:
        v29 := GenS5{F0: "a", F1: v11}
        _ = v29
    }
    var v30 string = ("hello" + "\n\t")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v9
}
