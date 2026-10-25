package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int32; F1 uint64 }
func genFn0(p0 int16) uint16 {
    p0 -= p0
    if false {
        var v5 int = 128
        var v6 string = "foo bar"
        _ = v5
        _ = v6
    }
    switch "" {
    case "":
        p0 += p0
    case "a":
        v7 := GenS1{F0: 924}
        _ = v7
    default:
        for fi8 := 0; fi8 < 9; fi8++ {
            fmt.Println(fi8, p0)
            var v9 uint = 18446744073709551615
            if (0 <= 0) {
                switch ("a" + "foo bar") {
                case "":
                    var v10 float32 = 1e10
                    _ = v10
                case "foo bar":
                    v11 := GenS1{F0: 0}
                    _ = v11
                case "a":
                    switch ((("a" + "\n\t") + "foo bar") + "a") {
                    case "\n\t":
                        fmt.Println(v9)
                    default:
                        fmt.Println(v9)
                    }
                }
                p0 = p0
                switch "hello" {
                case "\n\t":
                    for fi12 := 0; fi12 < 1; fi12++ {
                        fmt.Println(p0, fi12, fi12)
                        _ = fi12
                    }
                default:
                    v13 := GenS1{F0: 1}
                    _ = v13
                }
            }
            _ = fi8
            _ = v9
        }
    }
    p0 -= p0
    v14 := []int{346, 96}
    _ = p0
    _ = v14
    return 0
}
func genFn1(p0 int16) int64 {
    switch -101 {
    case 0:
        switch "a" {
        case "":
            p0 += p0
        case "foo bar":
            var v15 int = -2147483648
            _ = v15
        case "a":
            v16 := []int{2147483647, -2147483648, 128, 0}
            _ = v16
        }
    case 9223372036854775807:
        p0 += p0
    case -593:
        v17 := []int{9223372036854775807, -849}
        _ = v17
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        fmt.Println(fi18, p0)
        _ = fi18
    }
    var v19 uint32 = 42
    v20 := []int{}
    _ = p0
    _ = v19
    _ = v20
    return 0
}

func main() {
    var v1 uint16 = 32767
    var v2 int = 9223372036854775807
    v3 := make(map[string]uint32)
    var v4 int32 = 0
    fmt.Println(v1, v4)
    var v21 bool = true
    v2 = v2
    for fi22 := 0; fi22 < 17; fi22++ {
        var v23 uint8 = 127
        for fi24 := 0; fi24 < 17; fi24++ {
            v4 -= (v4 / (((v4 % (((v4 >> ((401) & 63))) | 1))) | 1))
            v25 := GenS2{F0: 62, F1: 1}
            _ = fi24
            _ = v25
        }
        var v26 uint16 = v1
        _ = fi22
        _ = v23
        _ = v26
    }
    if ((0.0 / 1e10) >= (0.1 - 3.14159)) {
        var v27 uint16 = (v1 << ((65535) & 63))
        v28 := GenS3{F0: v4, F1: uint64(v27)}
        _ = len(v3)
        _ = v27
        _ = v28
    } else {
        v2 -= v2
    }
    var v29 uint16 = 255
    v4 = 2147483647
    fmt.Println(v2)
    v30 := []int{32767, v2}
    var v31 int8 = -33
    if v21 {
        v2 = (v2 + (v2 / ((v2) | 1)))
    }
    var v32 string = "foo bar"
    for fi33 := 0; fi33 < 4; fi33++ {
        var v34 bool = (!v21)
        for fi35 := 0; fi35 < 4; fi35++ {
            switch (fi35 << ((969) & 63)) {
            case -394:
                v32 = ""
            default:
                var v36 uint = 0
                _ = v36
            }
            if len(v30) > 0 {
                v30[0] = (fi33 >> ((2) & 63))
            }
            _ = fi35
        }
        for fi37 := 0; fi37 < 2; fi37++ {
            var v38 int8 = v31
            _ = fi37
            _ = v38
        }
        _ = fi33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
