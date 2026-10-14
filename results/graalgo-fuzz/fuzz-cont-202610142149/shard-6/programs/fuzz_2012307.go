package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint64 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint64; F1 int; F2 int64 }
func genFn0(p0 uint, p1 bool, p2 int16) uint64 {
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 int64 = 0
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        switch (("" + "") + "a") {
        case "hello":
            v8 := []int{fi7, fi7, fi7}
            _ = v8
        case "a":
            for fi9 := 0; fi9 < 14; fi9++ {
                var v10 uint8 = 28
                if (42 != 374) {
                    switch "hello" {
                    case "":
                        if p1 {
                            fmt.Println(fi7)
                            if p1 {
                                var v11 uint64 = 2
                                v11 = (v11 << ((p0) & 63))
                                _ = v11
                            }
                            v12 := []int{}
                            _ = v12
                        }
                        v10 = 2
                    case "foo bar":
                        p2 = (p2 | int16(p0))
                    case "a":
                        v10 = v10
                    default:
                        var v13 int64 = -528
                        _ = v13
                    }
                    if true {
                        var v14 uint64 = 0
                        fmt.Println(v14, p1)
                        _ = v14
                    }
                }
                p1 = ((v10 - 247) >= v10)
                _ = fi9
                _ = v10
            }
        default:
            p0 -= (p0 << ((2) & 63))
        }
        _ = fi7
    }
    p0 += (p0 ^ (p0 * p0))
    v15 := []int{673, 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 340
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    var v3 int = 63
    var v4 uint = 65535
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 int8 = 100
        var v18 uint64 = 820
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 int = 0
    switch "\n\t" {
    case "foo bar":
        v4 += 65535
    case "hello":
        v20 := GenS1{F0: 2, F1: "hello", F2: 42}
        _ = v20
    default:
        for fi21 := 0; fi21 < 11; fi21++ {
            var v22 float64 = ((3.14159 + 3.14159) / 0.1)
            var v23 uint32 = 131
            v22 = (((v22 / (((3.14159 + 0.0) * (3.14159 + v22)) + (0.1 - v22))) + ((float64(3.14159) * (v22 + 1e10)) * v22)) * ((-2.25 + (float64(-2.25) - (v22 + float64(3.14159)))) + v22))
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        v25 := GenS2{F0: "foo bar", F1: 255}
        _ = fi24
        _ = v25
    }
    v3 -= v19
    v4 = v4
    switch 9223372036854775807 {
    case 9223372036854775807:
        v4 = (v4 - (v4 * (v4 >> ((v4) & 63))))
    case 32767:
        for fi26 := 0; fi26 < 1; fi26++ {
            var v27 uint32 = 2
            v28 := GenS2{F0: ("\n\t" + "foo bar"), F1: 1}
            _ = fi26
            _ = v27
            _ = v28
        }
    case 127:
        v29 := GenS1{F0: 843, F1: "hello", F2: 18446744073709551615}
        _ = v29
    }
    fmt.Println(v19, v4)
    _ = len(v2)
    v30 := v2[(("hello" + ((("\n\t" + "") + "a") + "")) + "hello")]
    for fi31 := 0; fi31 < 10; fi31++ {
        switch (("hello" + "\n\t") + "\n\t") {
        case "foo bar":
            delete(v1, 4294967295)
        default:
            for fi32 := 0; fi32 < 5; fi32++ {
                fmt.Println(v4, v3, fi31)
                v19 -= (v3 << (((v4 / (((v4 / ((v4) | 1))) | 1))) & 63))
                delete(v1, 2147483648)
                _ = fi32
            }
        }
        _ = fi31
    }
    switch v19 {
    case -9223372036854775808:
        v30 = v30
    case 9223372036854775807:
        for fi33 := 0; fi33 < 4; fi33++ {
            delete(v1, 4294967295)
            _ = fi33
        }
    case 127:
        fmt.Println(v3, v19)
    default:
        for fi34 := 0; fi34 < 12; fi34++ {
            var v35 int = 2147483647
            fmt.Println(fi34, v4, v30)
            v19 -= v3
            _ = fi34
            _ = v35
        }
    }
    v36 := GenS3{F0: 0, F1: (v3 | (-(v19))), F2: 9223372036854775807}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v36
    _ = v4
}
