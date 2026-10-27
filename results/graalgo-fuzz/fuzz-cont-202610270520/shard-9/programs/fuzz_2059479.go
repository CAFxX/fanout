package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 bool; F1 int32; F2 int32 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 string; F1 int8 }
func genFn0(p0 uint64) uint64 {
    var v6 uint64 = p0
    switch "foo bar" {
    case "a":
        if true {
            for fi7 := 0; fi7 < 9; fi7++ {
                v8 := GenS1{F0: (p0 & p0), F1: 127}
                _ = fi7
                _ = v8
            }
        }
    }
    fmt.Println(v6, p0)
    if false {
        fmt.Println(p0)
        var v9 int16 = -845
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 int8 = 127
            _ = fi10
            _ = v11
        }
        _ = v9
    } else {
        var v12 uint32 = 168
        for fi13 := 0; fi13 < 11; fi13++ {
            switch fi13 {
            case -1:
                for fi14 := 0; fi14 < 11; fi14++ {
                    v15 := GenS2{F0: false, F1: 254, F2: 323}
                    _ = fi14
                    _ = v15
                }
            case -2147483648:
                p0 -= 0
            case -852:
                v16 := GenS3{F0: 0, F1: (v12 & v12)}
                _ = v16
            default:
                v12 -= 65535
            }
            _ = fi13
        }
        _ = v12
    }
    _ = p0
    _ = v6
    return (p0 << ((1) & 63))
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{0, 128, 128}
    v3 := []string{"\n\t", ("" + "\n\t")}
    v4 := make(map[string]int)
    v5 := make(map[int]uint32)
    for fi17 := 0; fi17 < 11; fi17++ {
        for fi18 := 0; fi18 < 4; fi18++ {
            var v19 int16 = 1000
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    var v20 int64 = -351
    for fi21 := 0; fi21 < 7; fi21++ {
        if false {
            if true {
                v20 = -9223372036854775808
            } else {
                v20 += (-(v20))
                for fi22 := 0; fi22 < 20; fi22++ {
                    v20 += v20
                    _ = fi22
                }
                var v23 float32 = 3.14159
                _ = v23
            }
        }
        if ((1000 != 0) || (-250 <= -2147483648)) {
            v3 = append(v3, "hello")
            v24 := GenS4{F0: (("hello" + "foo bar") + "hello"), F1: 119}
            switch "" {
            case "":
                v20 += (v20 & (v20 & (v20 + (v20 | 817))))
            case "\n\t":
                switch 234 {
                case -2147483648:
                    var v25 bool = (106 != -51)
                    _ = v25
                default:
                    v5[(fi21 << ((255) & 63))] = 2147483648
                }
            case "hello":
                if (42 < 2) {
                    var v26 int16 = 32767
                    v20 -= v20
                    _ = v26
                }
            default:
                var v27 int32 = 0
                _ = v27
            }
            _ = v24
        }
        _ = fi21
    }
    v20 = (v20 ^ 0)
    v28 := GenS3{F0: 100, F1: 2147483648}
    _ = len(v4)
    switch "\n\t" {
    case "\n\t":
        for fi29 := 0; fi29 < 19; fi29++ {
            v20 += 626
            v20 += (v20 ^ v20)
            var v30 uint64 = genFn0(2)
            _ = fi29
            _ = v30
        }
    case "hello":
        var v31 int = 32767
        _ = v31
    case "foo bar":
        var v32 int = -622
        _ = v32
    }
    fmt.Println(v20)
    for fi33 := 0; fi33 < 4; fi33++ {
        v20 = (v20 << ((711) & 63))
        v20 -= (v20 & 972)
        _ = fi33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
