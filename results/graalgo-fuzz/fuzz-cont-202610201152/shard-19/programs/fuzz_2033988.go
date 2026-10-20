package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint16; F1 uint32 }
func genFn0() uint16 {
    var v4 int32 = -355
    var v5 uint = 811
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 uint64) int8 {
    fmt.Println(p0)
    v6 := GenS1{F0: (p0 * (^(p0))), F1: -288, F2: 275}
    if false {
        fmt.Println(p0)
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 int = fi7
            _ = fi7
            _ = v8
        }
        if (p0 > p0) {
            p0 = p0
        } else {
            v9 := GenS2{F0: 255, F1: 0}
            p0 += (p0 * 42)
            _ = v9
        }
    }
    _ = p0
    _ = v6
    return -17
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[int]string)
    var v3 int8 = 127
    var v10 uint8 = 116
    v3 -= (v3 & (v3 | (-(v3))))
    var v11 uint = 0
    v1 = 0
    fmt.Println(v10)
    switch "a" {
    case "foo bar":
        fmt.Println(v1, v11)
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        _ = len(v2)
        _ = fi12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        v2[(-(fi13))] = "a"
        fmt.Println(v10, v1, v1)
        _ = len(v2)
        _ = fi13
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        switch fi14 {
        case 0:
            v15 := GenS1{F0: uint64(v11), F1: 1, F2: 42}
            _ = v15
        case -1:
            v10 += v10
        case 42:
            v16 := GenS1{F0: 2, F1: 9223372036854775807, F2: 4294967295}
            _ = v16
        }
        _ = fi14
    }
    fmt.Println(v11, v1)
    if true {
        if true {
            if (1e10 <= 3.14159) {
                v11 -= (v11 + v11)
            }
        } else {
            v17 := GenS1{F0: 42, F1: -9223372036854775808, F2: 133}
            _ = v17
        }
        fmt.Println(v11)
    }
    fmt.Println(v10)
    var v18 uint8 = 255
    switch 128 {
    case 4294967295:
        for fi19 := 0; fi19 < 8; fi19++ {
            switch "foo bar" {
            case "foo bar":
                switch "hello" {
                case "":
                    v3 += -102
                case "foo bar":
                    v18 = v10
                case "a":
                    fmt.Println(v1)
                default:
                    switch fi19 {
                    case 42:
                        var v20 float32 = 1e10
                        _ = v20
                    case 127:
                        var v21 uint16 = v1
                        _ = v21
                    }
                }
            case "\n\t":
                for fi22 := 0; fi22 < 1; fi22++ {
                    v23 := GenS2{F0: genFn0(), F1: 134}
                    v1 += (v1 + v1)
                    _ = fi22
                    _ = v23
                }
            case "a":
                fmt.Println(fi19, v1, v3)
            default:
                for fi24 := 0; fi24 < 9; fi24++ {
                    var v25 float32 = ((3.14159 * (1e10 - ((-2.25 / 0.1) / float32(1e10)))) + 3.14159)
                    switch ("hello" + "") {
                    case "hello":
                        var v26 int8 = v3
                        _ = v26
                    case "foo bar":
                        v27 := []int{fi24, 42, fi19}
                        _ = v27
                    }
                    fmt.Println(v3, v11)
                    _ = fi24
                    _ = v25
                }
            }
            fmt.Println(v10)
            _ = fi19
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v3
}
