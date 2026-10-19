package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 uint64 }
func genFn0(p0 int8, p1 uint8, p2 uint32) uint16 {
    p1 -= (p1 >> ((18446744073709551615) & 63))
    v6 := GenS1{F0: "", F1: (^(p1)), F2: 734}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 544
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 2147483648
    var v3 float64 = 1e10
    v4 := []string{"hello", "foo bar", "\n\t"}
    v5 := make(map[uint32]uint32)
    v2 = 2147483648
    var v7 uint16 = genFn0(42, uint8(v2), 239)
    var v8 uint32 = (v2 * (v2 ^ (v2 | 119)))
    var v9 int32 = -2147483648
    v10 := GenS1{F0: "foo bar", F1: 67, F2: 18446744073709551615}
    v9 -= 1
    v8 = 307
    if false {
        switch ("foo bar" + ("\n\t" + "a")) {
        case "hello":
            if true {
                v11 := GenS1{F0: "foo bar", F1: 0, F2: 65535}
                _ = v11
            }
        case "":
            var v12 uint64 = 18446744073709551615
            _ = v12
        case "a":
            v13 := v1[118]
            _ = v13
        default:
            fmt.Println(v9, v2, v2)
        }
        v14 := GenS1{F0: ((("a" + "\n\t") + "foo bar") + "a"), F1: 1, F2: 974}
        var v15 bool = false
        _ = v14
        _ = v15
    } else {
        v1[9223372036854775807] = false
    }
    v7 += v7
    switch "foo bar" {
    case "foo bar":
        v8 += v2
    case "":
        for fi16 := 0; fi16 < 15; fi16++ {
            var v17 uint64 = 1
            v18 := GenS1{F0: ("a" + ""), F1: 2, F2: 0}
            v17 = 0
            _ = fi16
            _ = v17
            _ = v18
        }
    case "a":
        var v19 string = ("hello" + "a")
        _ = v19
    }
    if false {
        for fi20 := 0; fi20 < 18; fi20++ {
            var v21 uint16 = (^(v7))
            _ = fi20
            _ = v21
        }
        v2 = (v8 % ((v8) | 1))
        v3 = v3
    }
    var v22 uint8 = 0
    var v23 int8 = 0
    if false {
        switch "foo bar" {
        case "hello":
            v23 = (^(v23))
        case "a":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "\n\t":
            if (2147483647 >= -1) {
                var v24 float32 = (3.14159 / 1.0)
                v23 = v23
                _ = v24
            } else {
                var v25 int16 = -1
                for fi26 := 0; fi26 < 17; fi26++ {
                    fmt.Println(fi26)
                    if (((v22 + v22) >= v22) || (true || (167 <= 9223372036854775807))) {
                        switch "" {
                        case "a":
                            v27 := GenS2{F0: -9223372036854775808}
                            _ = v27
                        case "hello":
                            var v28 uint32 = 65535
                            _ = v28
                        default:
                            var v29 float32 = (-2.25 - (3.14159 + -2.25))
                            _ = v29
                        }
                        for fi30 := 0; fi30 < 9; fi30++ {
                            v7 += (v7 / ((v7) | 1))
                            _ = fi30
                        }
                    } else {
                        for fi31 := 0; fi31 < 19; fi31++ {
                            fmt.Println(fi31, v3, v2)
                            _ = fi31
                        }
                        if true {
                            var v32 int16 = v25
                            v33 := GenS2{F0: 505}
                            var v34 int16 = 32767
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                    }
                    _ = fi26
                }
                _ = v25
            }
        }
        for fi35 := 0; fi35 < 20; fi35++ {
            for fi36 := 0; fi36 < 1; fi36++ {
                switch "" {
                case "hello":
                    v37 := GenS3{F0: (^(v23)), F1: 726}
                    _ = v37
                }
                _ = fi36
            }
            fmt.Println(v7, v7)
            _ = fi35
        }
        var v38 uint32 = 1
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
