package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint32; F1 string; F2 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 int32 }
func genFn0(p0 uint16, p1 string, p2 int16) uint8 {
    switch 2147483647 {
    case 4294967295:
        v5 := []int{1, 9223372036854775807, -819, 522}
        _ = v5
    case 0:
        var v6 int16 = -712
        _ = v6
    case 128:
        v7 := GenS1{F0: 42, F1: 81}
        _ = v7
    default:
        p2 = p2
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        p1 = p1
        _ = fi8
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1(p0 int16) int64 {
    var v9 float64 = float64(((0.0 / 0.1) - -2.25))
    var v10 uint64 = 0
    switch 32767 {
    case 42:
        switch -338 {
        case 1:
            for fi11 := 0; fi11 < 17; fi11++ {
                v9 = ((3.14159 / (v9 / 1e10)) - 1.5)
                for fi12 := 0; fi12 < 15; fi12++ {
                    if true {
                        var v13 int16 = p0
                        v10 += v10
                        _ = v13
                    }
                    _ = fi12
                }
                _ = fi11
            }
        case 113:
            p0 -= (p0 & 0)
        default:
            v14 := GenS1{F0: 18446744073709551615, F1: -62}
            _ = v14
        }
    case 690:
        var v15 int = 32767
        _ = v15
    default:
        v10 -= 0
    }
    switch ("\n\t" + "") {
    case "\n\t":
        var v16 string = "hello"
        _ = v16
    case "hello":
        for fi17 := 0; fi17 < 7; fi17++ {
            var v18 uint = 1
            v19 := []int{(fi17 / ((fi17) | 1)), 42}
            _ = fi17
            _ = v18
            _ = v19
        }
    case "a":
        v10 = (v10 ^ 469)
    }
    _ = p0
    _ = v10
    _ = v9
    return 42
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]uint32)
    var v3 int16 = -737
    var v4 uint32 = 507
    v3 += 599
    v20 := GenS2{F0: v4, F1: "hello", F2: 4294967295}
    v4 += v4
    switch ("foo bar" + (("" + (("" + ((("a" + ("a" + "")) + "a") + "hello")) + "hello")) + "")) {
    case "\n\t":
        v4 = (v4 | v4)
    default:
        v21 := GenS3{F0: (("\n\t" + "\n\t") + "a")}
        _ = v21
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 uint32 = 2147483647
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        for fi25 := 0; fi25 < 2; fi25++ {
            v26 := GenS4{F0: 733, F1: 2147483647}
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    var v27 uint = 65535
    v28 := GenS1{F0: 891, F1: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
