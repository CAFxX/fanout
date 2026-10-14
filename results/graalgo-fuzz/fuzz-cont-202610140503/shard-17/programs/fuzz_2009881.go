package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int64; F2 int8 }
func genFn0(p0 uint8) uint8 {
    fmt.Println(p0)
    p0 = p0
    _ = p0
    return (p0 % (((p0 % ((p0) | 1))) | 1))
}
func genFn1(p0 uint32, p1 int64) uint32 {
    if (79 >= 1) {
        v5 := GenS1{F0: 32767, F1: ((("\n\t" + (("a" + "a") + "a")) + ("hello" + ((("" + "") + "\n\t") + ""))) + ((((((("" + "a") + "") + (("\n\t" + "a") + ((("" + "hello") + ("\n\t" + ("foo bar" + "\n\t"))) + "\n\t"))) + "\n\t") + "a") + "foo bar") + "foo bar")), F2: 32768}
        v6 := GenS1{F0: 1, F1: ("hello" + "foo bar"), F2: 9223372036854775807}
        fmt.Println(p1, p0)
        _ = v5
        _ = v6
    } else {
        if (516 > 42) {
            p1 += 447
        } else {
            v7 := []int{128, 2147483647, 127, int(p1)}
            _ = v7
        }
    }
    p1 = (-(p1))
    if false {
        p0 -= (p0 ^ p0)
        p0 = 1
        if true {
            v8 := []int{}
            _ = v8
        }
    } else {
        for fi9 := 0; fi9 < 5; fi9++ {
            for fi10 := 0; fi10 < 15; fi10++ {
                fmt.Println(fi9)
                if (!true) {
                    var v11 int = (fi10 << ((42) & 63))
                    p0 -= p0
                    _ = v11
                }
                var v12 int16 = -32768
                var v13 string = "hello"
                _ = fi10
                _ = v12
                _ = v13
            }
            fmt.Println(p0, p0)
            v14 := GenS1{F0: (fi9 % (((fi9 | (-(fi9)))) | 1)), F1: "\n\t", F2: (fi9 | (fi9 ^ fi9))}
            _ = fi9
            _ = v14
        }
        v15 := GenS2{F0: p1, F1: (p0 | 1)}
        for fi16 := 0; fi16 < 19; fi16++ {
            v17 := []int{}
            var v18 uint64 = 0
            p1 += (p1 * (p1 | p1))
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = v15
    }
    if true {
        for fi19 := 0; fi19 < 11; fi19++ {
            p1 = -9223372036854775808
            _ = fi19
        }
    } else {
        p1 -= -9223372036854775808
        v20 := []int{32768, 533, -733}
        if false {
            if true {
                switch 2147483647 {
                case -1:
                    var v21 int8 = -1
                    _ = v21
                case 0:
                    var v22 uint32 = p0
                    _ = v22
                }
                v20 = append(v20, -2147483648)
            }
            switch 9223372036854775807 {
            case 42:
                var v23 float64 = 0.0
                _ = v23
            }
            v24 := GenS2{F0: p1, F1: 65535}
            _ = v24
        } else {
            var v25 int64 = p1
            v20 = append(v20, 938)
            p1 += (^(p1))
            _ = v25
        }
        _ = v20
    }
    _ = p0
    _ = p1
    return 524
}

func main() {
    v1 := []int{}
    var v2 float64 = float64(float32(1e10))
    v3 := make(map[uint32]bool)
    var v4 uint32 = 2147483647
    var v26 float32 = (0.1 - 1.5)
    v4 = 2
    v26 = v26
    fmt.Println(v2, v4)
    if true {
        if (92 < 42) {
            for fi27 := 0; fi27 < 7; fi27++ {
                var v28 uint8 = 155
                var v29 uint8 = v28
                _ = fi27
                _ = v28
                _ = v29
            }
            fmt.Println(v4, v4, v4)
        } else {
            v4 = 4294967295
        }
        switch 270 {
        case -9223372036854775808:
            var v30 bool = (626 != 1)
            _ = v30
        case 32768:
            if (32767 >= 1) {
                fmt.Println(v26, v2, v4)
                if ((!(255 > 243)) || false) {
                    var v31 int32 = -227
                    v32 := GenS1{F0: 305, F1: "a", F2: -308}
                    v2 -= v2
                    _ = v31
                    _ = v32
                } else {
                    v2 -= (v2 * 1e10)
                    fmt.Println(v4)
                    var v33 uint = 42
                    _ = v33
                }
                v34 := GenS3{F0: 101, F1: 0, F2: int8(v4)}
                _ = v34
            }
        }
    }
    var v35 uint8 = 128
    var v36 uint = 702
    v4 = genFn1(204, 42)
    for fi37 := 0; fi37 < 14; fi37++ {
        v1 = append(v1, (fi37 << ((v36) & 63)))
        if true {
            switch fi37 {
            case -9223372036854775808:
                v4 = 948
            default:
                v3[(^(v4))] = false
            }
        } else {
            v4 -= v4
        }
        if true {
            switch "a" {
            case "foo bar":
                if len(v1) > 0 {
                    v1[0] = fi37
                }
            case "a":
                var v38 int32 = -2147483648
                _ = v38
            }
            var v39 float64 = v2
            var v40 uint64 = 18446744073709551615
            _ = v39
            _ = v40
        } else {
            v3[uint32(v36)] = false
            v26 -= v26
        }
        _ = fi37
    }
    var v41 float32 = 0.1
    for fi42 := 0; fi42 < 1; fi42++ {
        for fi43 := 0; fi43 < 20; fi43++ {
            fmt.Println(v2)
            _ = fi43
        }
        var v44 int64 = 9223372036854775807
        var v45 int64 = (v44 ^ 42)
        _ = fi42
        _ = v44
        _ = v45
    }
    if true {
        v46 := v3[(v4 << ((255) & 63))]
        _ = v46
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v41
}
