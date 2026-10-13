package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int8; F1 int32; F2 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint }
func genFn0(p0 uint8) int8 {
    for fi5 := 0; fi5 < 15; fi5++ {
        p0 -= p0
        var v6 int = (fi5 ^ (^(fi5)))
        _ = fi5
        _ = v6
    }
    if false {
        switch ("hello" + ((("foo bar" + "foo bar") + ("foo bar" + ("\n\t" + "hello"))) + "")) {
        case "foo bar":
            fmt.Println(p0)
        case "a":
            p0 += p0
        }
    }
    if (42 > -19) {
        switch -9223372036854775808 {
        case 2147483647:
            v7 := []int{}
            _ = v7
        case -242:
            v8 := GenS1{F0: 255, F1: (p0 << ((1) & 63))}
            _ = v8
        case 0:
            switch 1 {
            case 0:
                v9 := GenS2{F0: 1, F1: 1}
                _ = v9
            case 4294967295:
                v10 := []int{32767}
                _ = v10
            case -1:
                fmt.Println(p0)
            }
        }
        v11 := GenS3{F0: 0, F1: 42, F2: 65535}
        _ = v11
    } else {
        switch "\n\t" {
        case "a":
            fmt.Println(p0)
        case "foo bar":
            for fi12 := 0; fi12 < 18; fi12++ {
                fmt.Println(p0)
                p0 -= p0
                _ = fi12
            }
        case "\n\t":
            p0 += (p0 % (((p0 | (p0 * p0))) | 1))
        }
        var v13 uint8 = p0
        _ = v13
    }
    _ = p0
    return -25
}
func genFn1(p0 int16, p1 bool, p2 int64) uint32 {
    v14 := GenS3{F0: -78, F1: -2147483648, F2: 2147483647}
    v15 := []int{-1, 32768, 509, 32768}
    p2 += (p2 % ((p2) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return 2147483648
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{42, 1}
    var v3 int = 42
    v4 := make(map[string]bool)
    v16 := GenS4{F0: 32767}
    fmt.Println(v3)
    v3 -= (v3 / (((v3 >> ((273) & 63))) | 1))
    fmt.Println(v3)
    v3 -= (v3 ^ 32767)
    v2 = append(v2, 615)
    var v17 int64 = -9223372036854775808
    for fi18 := 0; fi18 < 16; fi18++ {
        _ = len(v4)
        var v19 string = "foo bar"
        _ = fi18
        _ = v19
    }
    if false {
        if len(v2) > 0 {
            v2[0] = 2147483648
        }
        if (genFn1(-969, false, 9223372036854775807) >= 2147483647) {
            fmt.Println(v3, v3)
        } else {
            switch "hello" {
            case "a":
                switch v3 {
                case 359:
                    _ = len(v2)
                case 32767:
                    v17 += (v17 ^ (^(v17)))
                default:
                    if (((float32(3.14159) + (((0.0 * float32(3.14159)) + (0.1 / float32(-2.25))) + -2.25)) / -2.25) == 0.0) {
                        v3 = 128
                        v17 += v17
                        if false {
                            v3 = 32767
                        }
                    } else {
                        if true {
                            var v20 uint = 367
                            _ = v20
                        }
                        _ = len(v2)
                    }
                }
            case "hello":
                var v21 int64 = v17
                _ = v21
            }
            _ = len(v2)
        }
        v22 := GenS5{F0: 389}
        _ = v22
    }
    v23 := GenS3{F0: -121, F1: int32(v17), F2: 0}
    v3 -= (v3 >> ((351) & 63))
    v17 = v17
    v24 := GenS1{F0: 65535, F1: 42}
    var v25 float32 = 1.5
    var v26 string = "hello"
    _ = len(v2)
    switch "a" {
    case "":
        delete(v1, (v26 + v26))
    case "foo bar":
        var v27 float32 = v25
        _ = v27
    default:
        if false {
            var v28 uint32 = 2147483648
            if true {
                v29 := GenS4{F0: 379}
                v30 := GenS4{F0: -32768}
                var v31 int8 = 0
                _ = v29
                _ = v30
                _ = v31
            } else {
                if (42 == 466) {
                    var v32 bool = false
                    _ = v32
                } else {
                    if (false || true) {
                        var v33 int32 = -1
                        var v34 int32 = (v33 / (((v33 % (((v33 << ((4294967295) & 63))) | 1))) | 1))
                        _ = v33
                        _ = v34
                    }
                    v35 := GenS2{F0: 0, F1: 2}
                    var v36 int16 = -32768
                    _ = v35
                    _ = v36
                }
            }
            v25 -= 0.1
            _ = v28
        } else {
            v25 = (1e10 - v25)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
