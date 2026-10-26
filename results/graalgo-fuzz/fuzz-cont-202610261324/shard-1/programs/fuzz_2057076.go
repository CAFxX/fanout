package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8 }
func genFn0(p0 uint8, p1 string, p2 int16) int32 {
    v7 := []int{}
    p1 = ("a" + "hello")
    var v8 int = 32768
    p1 = p1
    p0 -= 42
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return int32(p0)
}
func genFn1(p0 string, p1 uint32, p2 uint64) uint {
    switch 1 {
    case 2147483647:
        v9 := GenS1{F0: p0, F1: p1, F2: 54}
        _ = v9
    case 1:
        if (genFn0(42, (p0 + p0), 42) >= -765) {
            var v10 bool = true
            _ = v10
        } else {
            var v11 int64 = 42
            _ = v11
        }
    case 73:
        if ((((((0.1 / -2.25) / 1.5) - ((0.0 * float32(0.0)) - (1.5 + ((float32(1.5) / float32(-2.25)) / float32(((((1.5 / ((1.5 * 0.1) * (0.0 + 1.5))) - ((1e10 * (0.0 - ((1.5 - 0.0) + (1.5 * 1e10)))) + float64((-2.25 + 0.0)))) - (1e10 * float64(float32(float64(((((0.0 * 0.1) - (0.1 - -2.25)) + 0.0) + -2.25)))))) * 1.5)))))) + 0.1) / 1.0) < 3.14159) {
            p2 = p2
            if (-628 < 0) {
                v12 := GenS1{F0: "hello", F1: p1, F2: 0}
                _ = v12
            }
            v13 := []int{}
            _ = v13
        } else {
            v14 := []int{}
            v14 = append(v14, 0)
            for fi15 := 0; fi15 < 14; fi15++ {
                v16 := GenS1{F0: p0, F1: p1, F2: 128}
                if false {
                    if (-1 == -32768) {
                        p0 = p0
                        fmt.Println(p0)
                    }
                    p1 = 2
                    fmt.Println(fi15, fi15, p0)
                } else {
                    p2 = (p2 + (^(p2)))
                }
                _ = fi15
                _ = v16
            }
            _ = v14
        }
    default:
        var v17 string = p0
        _ = v17
    }
    var v18 uint16 = 65535
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 uint32 = 583
        p1 = v20
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        for fi22 := 0; fi22 < 16; fi22++ {
            var v23 int8 = 21
            _ = fi22
            _ = v23
        }
        if true {
            var v24 float32 = (1.5 - 3.14159)
            _ = v24
        }
        _ = fi21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return 135
}

func main() {
    var v1 int32 = 712
    var v2 uint16 = 32767
    v3 := []string{"a"}
    var v4 uint16 = (v2 % ((v2) | 1))
    var v5 string = (("foo bar" + "a") + ((("" + "\n\t") + ("hello" + "hello")) + "\n\t"))
    v6 := make(map[uint32]uint32)
    _ = len(v3)
    v4 += v4
    v25 := GenS2{F0: genFn0(128, v5, 32767), F1: 32767}
    for fi26 := 0; fi26 < 15; fi26++ {
        _ = len(v3)
        for fi27 := 0; fi27 < 9; fi27++ {
            switch v5 {
            case "":
                v5 = "hello"
            case "foo bar":
                switch "\n\t" {
                case "hello":
                    var v28 float32 = float32(3.14159)
                    _ = v28
                case "\n\t":
                    v5 = (v5 + (v5 + v5))
                case "foo bar":
                    var v29 uint32 = 42
                    _ = v29
                default:
                    var v30 int = (-(fi26))
                    _ = v30
                }
            case "a":
                v5 = v5
            default:
                if (!true) {
                    fmt.Println(v1, fi27, v2)
                }
            }
            _ = fi27
        }
        _ = fi26
    }
    v31 := GenS3{F0: ("foo bar" + v5)}
    v32 := GenS4{F0: 2}
    v4 = (v4 + (^(v4)))
    if true {
        fmt.Println(v2)
    } else {
        var v33 uint8 = 103
        if (!((v1 <= (v1 << ((1) & 63))) || (false && true))) {
            switch 999 {
            case 44:
                for fi34 := 0; fi34 < 7; fi34++ {
                    switch v5 {
                    case "":
                        var v35 int8 = 127
                        _ = v35
                    default:
                        v2 += v4
                    }
                    _ = fi34
                }
            default:
                var v36 uint32 = 912
                _ = v36
            }
        }
        _ = v33
    }
    v6[569] = 2147483648
    delete(v6, 2)
    v2 += 32767
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
