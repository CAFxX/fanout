package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int8 }
type GenS2 struct { F0 uint; F1 string; F2 uint }
type GenS3 struct { F0 int; F1 int32; F2 uint8 }
type GenS4 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint64 }
func genFn0(p0 uint16, p1 int32) string {
    var v5 uint8 = 127
    p1 = p1
    _ = p0
    _ = p1
    _ = v5
    return ""
}
func genFn1(p0 uint16) uint16 {
    p0 -= p0
    for fi6 := 0; fi6 < 3; fi6++ {
        if true {
            p0 = 1
            p0 -= p0
            for fi7 := 0; fi7 < 14; fi7++ {
                var v8 uint = 4294967295
                _ = fi7
                _ = v8
            }
        } else {
            v9 := GenS1{F0: -969, F1: 127, F2: 75}
            for fi10 := 0; fi10 < 6; fi10++ {
                var v11 uint64 = 0
                switch fi10 {
                case 127:
                    var v12 string = ("foo bar" + "a")
                    _ = v12
                }
                v13 := []int{(fi10 & (^(fi10))), fi6, -941, fi6}
                _ = fi10
                _ = v11
                _ = v13
            }
            p0 = p0
            _ = v9
        }
        var v14 uint8 = 1
        _ = fi6
        _ = v14
    }
    p0 = (p0 - p0)
    var v15 int32 = -1
    if false {
        if false {
            v16 := []int{}
            _ = v16
        }
    }
    _ = p0
    _ = v15
    return p0
}

func main() {
    var v1 int16 = 32767
    v2 := []uint32{42, 893}
    v3 := make(map[int]bool)
    var v4 int64 = 0
    switch 128 {
    case 2147483647:
        v17 := GenS2{F0: 56, F1: "foo bar", F2: 1}
        _ = v17
    case 758:
        var v18 uint32 = 22
        _ = v18
    case -9223372036854775808:
        for fi19 := 0; fi19 < 9; fi19++ {
            for fi20 := 0; fi20 < 12; fi20++ {
                var v21 int64 = v4
                v22 := GenS1{F0: (v1 >> ((29) & 63)), F1: 1, F2: 42}
                var v23 float64 = (0.0 * 0.0)
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
            switch ("\n\t" + "hello") {
            case "foo bar":
                v24 := GenS3{F0: 32768, F1: -2147483648, F2: 25}
                _ = v24
            }
            for fi25 := 0; fi25 < 10; fi25++ {
                fmt.Println(v1)
                v26 := GenS4{F0: 842, F1: 1, F2: 562}
                v4 -= -1
                _ = fi25
                _ = v26
            }
            _ = fi19
        }
    }
    v27 := GenS4{F0: 65535, F1: 127, F2: 4294967295}
    v28 := GenS2{F0: 4294967295, F1: "foo bar", F2: 0}
    if false {
        var v29 uint16 = 998
        v4 -= (v4 << ((255) & 63))
        fmt.Println(v1)
        _ = v29
    }
    switch ("hello" + "\n\t") {
    case "\n\t":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "":
        v30 := v3[-2147483648]
        _ = v30
    default:
        for fi31 := 0; fi31 < 18; fi31++ {
            switch (fi31 / ((1) | 1)) {
            case 128:
                var v32 int = fi31
                _ = v32
            case 255:
                switch fi31 {
                case 421:
                    switch "foo bar" {
                    case "hello":
                        v1 += v1
                    case "":
                        v1 += (v1 | v1)
                    case "foo bar":
                        v33 := GenS5{F0: 1}
                        _ = v33
                    }
                case -2147483648:
                    v4 += 0
                case 227:
                    v4 += -9223372036854775808
                }
            }
            _ = fi31
        }
    }
    var v34 float32 = 1.5
    v4 = -9223372036854775808
    v3[889] = true
    if len(v2) > 0 {
        v2[0] = 366
    }
    v35 := GenS6{F0: 65535}
    fmt.Println(v4, v34, v1)
    fmt.Println(v4)
    fmt.Println(v1, v4, v1)
    var v36 int64 = (-(v4))
    v37 := GenS7{F0: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
