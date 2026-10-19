package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 bool }
func genFn0(p0 uint64, p1 uint64, p2 int32) string {
    switch "foo bar" {
    case "foo bar":
        p1 += p0
    default:
        var v5 bool = ("foo bar" <= "\n\t")
        _ = v5
    }
    var v6 uint32 = 1
    var v7 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return "hello"
}
func genFn1(p0 uint32, p1 int) uint32 {
    v8 := []int{127, p1}
    v9 := GenS1{F0: 0, F1: 0}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return (p0 % (((p0 + 759)) | 1))
}

func main() {
    var v1 uint8 = 242
    var v2 float64 = 3.14159
    v3 := make(map[string]int)
    var v4 float64 = 0.0
    var v10 float32 = float32((-2.25 - v2))
    v4 = ((3.14159 - v2) - v2)
    v10 = float32(v4)
    if false {
        if (true || false) {
            v11 := []int{}
            v4 += v2
            if (18446744073709551615 > 4294967295) {
                for fi12 := 0; fi12 < 20; fi12++ {
                    switch (fi12 + fi12) {
                    case -2147483648:
                        v13 := GenS2{F0: 1, F1: fi12}
                        _ = v13
                    case 733:
                        for fi14 := 0; fi14 < 15; fi14++ {
                            if (2 >= 623) {
                                if len(v11) > 0 {
                                    v11[0] = fi14
                                }
                                fmt.Println(v4, v2)
                                fmt.Println(v10, v4)
                            }
                            v2 = 1.5
                            _ = fi14
                        }
                    }
                    if false {
                        _ = len(v11)
                    }
                    _ = fi12
                }
            }
            _ = v11
        } else {
            var v15 int64 = 703
            _ = v15
        }
        fmt.Println(v4, v2, v10)
    } else {
        if false {
            v4 -= 0.1
        } else {
            for fi16 := 0; fi16 < 13; fi16++ {
                var v17 float64 = v2
                v18 := GenS3{F0: "", F1: 42}
                _ = fi16
                _ = v17
                _ = v18
            }
            switch 801 {
            case 9223372036854775807:
                var v19 float32 = v10
                _ = v19
            case 32767:
                fmt.Println(v2, v2, v2)
            default:
                if true {
                    for fi20 := 0; fi20 < 4; fi20++ {
                        v10 -= ((1e10 - v10) * 0.1)
                        v1 = v1
                        var v21 bool = true
                        _ = fi20
                        _ = v21
                    }
                    if false {
                        var v22 int64 = 0
                        v2 -= -2.25
                        v23 := GenS3{F0: "hello", F1: 42}
                        _ = v22
                        _ = v23
                    }
                    v24 := GenS1{F0: 28, F1: -1}
                    _ = v24
                }
            }
        }
        v1 -= (v1 ^ 224)
    }
    var v25 int8 = 42
    if (!(562 != 797)) {
        v26 := GenS1{F0: 255, F1: 2147483647}
        _ = v26
    } else {
        switch "hello" {
        case "":
            _ = len(v3)
        case "a":
            v4 += ((v2 * ((1.5 * (v2 / (float64((v10 - (-2.25 + float32(0.1)))) + 1e10))) + 0.0)) / 1.5)
        case "\n\t":
            v27 := []int{}
            _ = v27
        default:
            v28 := []int{288, 9223372036854775807, 32767}
            _ = v28
        }
        var v29 int32 = -373
        v30 := GenS1{F0: 255, F1: v29}
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 10; fi31++ {
        v1 = 23
        _ = fi31
    }
    var v32 uint32 = 0
    v33 := GenS3{F0: "a", F1: (v25 << ((255) & 63))}
    if ((v10 != 3.14159) && true) {
        fmt.Println(v10, v32, v1)
        for fi34 := 0; fi34 < 17; fi34++ {
            fmt.Println(v32)
            v35 := GenS4{F0: 1, F1: 0, F2: v32}
            v36 := v3["\n\t"]
            _ = fi34
            _ = v35
            _ = v36
        }
    } else {
        fmt.Println(v10, v1, v32)
        v37 := GenS5{F0: 9223372036854775807}
        v2 += v4
        _ = v37
    }
    v38 := GenS6{F0: (383 <= 0)}
    v3[("foo bar" + ("a" + "\n\t"))] = 127
    delete(v3, ("hello" + "hello"))
    v39 := GenS6{F0: true}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v33
    _ = v38
    _ = v39
    _ = v4
}
