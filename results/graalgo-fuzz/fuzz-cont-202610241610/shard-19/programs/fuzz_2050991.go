package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int; F1 uint64 }
func genFn0(p0 string) int8 {
    var v4 int64 = 0
    for fi5 := 0; fi5 < 6; fi5++ {
        for fi6 := 0; fi6 < 12; fi6++ {
            var v7 int = 128
            _ = fi6
            _ = v7
        }
        v8 := GenS1{F0: 68, F1: 212, F2: 454}
        _ = fi5
        _ = v8
    }
    if true {
        var v9 float64 = -2.25
        _ = v9
    } else {
        var v10 uint8 = 0
        v11 := []int{127, 396}
        for fi12 := 0; fi12 < 6; fi12++ {
            v10 += (v10 & 58)
            v4 = (v4 ^ -9223372036854775808)
            var v13 int64 = v4
            _ = fi12
            _ = v13
        }
        _ = v10
        _ = v11
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        p0 = p0
        v15 := GenS1{F0: 0, F1: 1, F2: 1}
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = v4
    return -1
}
func genFn1(p0 uint, p1 bool) int64 {
    v16 := GenS1{F0: 42, F1: 411, F2: p0}
    if true {
        var v17 uint64 = 756
        _ = v17
    }
    v18 := GenS1{F0: 439, F1: 2147483648, F2: (^(p0))}
    for fi19 := 0; fi19 < 1; fi19++ {
        p1 = p1
        _ = fi19
    }
    var v20 string = ("hello" + ("hello" + ("a" + "hello")))
    _ = p0
    _ = p1
    _ = v16
    _ = v18
    _ = v20
    return 584
}

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[string]int)
    var v3 uint32 = 283
    var v21 int8 = -128
    var v22 uint = 18446744073709551615
    for fi23 := 0; fi23 < 2; fi23++ {
        fmt.Println(v1)
        _ = fi23
    }
    v24 := GenS1{F0: (v3 | v3), F1: v3, F2: v22}
    switch ("\n\t" + "a") {
    case "hello":
        if (v3 > (^(v3))) {
            if true {
                v25 := GenS2{F0: 95, F1: (^(v1))}
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v21, v22, v3)
                }
                v21 += (v21 - 0)
                _ = v25
            } else {
                v26 := GenS2{F0: 255, F1: v1}
                _ = v26
            }
        } else {
            var v27 float32 = 0.1
            _ = v27
        }
    case "\n\t":
        v28 := GenS2{F0: uint8(v3), F1: v1}
        _ = v28
    case "foo bar":
        v29 := v2[""]
        _ = v29
    }
    var v30 uint32 = v3
    for fi31 := 0; fi31 < 7; fi31++ {
        delete(v2, "foo bar")
        _ = fi31
    }
    if (true || false) {
        v22 += (v22 << ((132) & 63))
    } else {
        if (!true) {
            for fi32 := 0; fi32 < 9; fi32++ {
                v33 := GenS3{F0: fi32, F1: 4294967295}
                for fi34 := 0; fi34 < 4; fi34++ {
                    var v35 bool = (!true)
                    v36 := v2[""]
                    _ = fi34
                    _ = v35
                    _ = v36
                }
                v1 += int(v21)
                _ = fi32
                _ = v33
            }
            if true {
                v37 := GenS3{F0: v1, F1: 1}
                _ = v37
            }
        }
        v30 = 0
        for fi38 := 0; fi38 < 3; fi38++ {
            var v39 int8 = (-(v21))
            v22 = v22
            _ = fi38
            _ = v39
        }
    }
    switch v1 {
    case 956:
        v21 = (v21 / (((v21 % ((v21) | 1))) | 1))
    default:
        var v40 bool = true
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v30
}
