package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 int8; F2 int32 }
func genFn0(p0 int8, p1 uint32, p2 int16) int32 {
    if (127 > 42) {
        v5 := []int{}
        _ = len(v5)
        _ = v5
    }
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 int = (fi6 - (fi6 >> ((42) & 63)))
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 228
}
func genFn1() uint {
    var v8 uint64 = 18446744073709551615
    v9 := GenS1{F0: 2}
    v10 := []int{127, 32767, 4294967295, -159}
    for fi11 := 0; fi11 < 10; fi11++ {
        switch 32767 {
        case 32767:
            switch fi11 {
            case 4294967295:
                var v12 int64 = int64(fi11)
                _ = v12
            case -273:
                for fi13 := 0; fi13 < 20; fi13++ {
                    v14 := GenS1{F0: 2}
                    _ = fi13
                    _ = v14
                }
            }
        case -260:
            v8 = (v8 * v8)
        default:
            var v15 uint64 = (^(v8))
            _ = v15
        }
        _ = fi11
    }
    _ = v10
    _ = v8
    _ = v9
    return 129
}

func main() {
    var v1 int = -2147483648
    var v2 uint8 = 0
    var v3 uint16 = 32767
    v4 := make(map[int]uint32)
    v3 += (^(v3))
    var v16 uint16 = (v3 + v3)
    v17 := GenS1{F0: v16}
    _ = len(v4)
    var v18 uint8 = v2
    switch (-(v1)) {
    case 32768:
        if (!true) {
            for fi19 := 0; fi19 < 16; fi19++ {
                for fi20 := 0; fi20 < 16; fi20++ {
                    var v21 float64 = 0.0
                    _ = fi20
                    _ = v21
                }
                _ = len(v4)
                _ = fi19
            }
        } else {
            fmt.Println(v1)
        }
    case 9223372036854775807:
        if ((v1 >> ((697) & 63)) >= 32768) {
            v22 := GenS1{F0: (^(v3))}
            var v23 int = 903
            _ = v22
            _ = v23
        }
    case -438:
        for fi24 := 0; fi24 < 9; fi24++ {
            fmt.Println(v2, fi24)
            _ = fi24
        }
    default:
        v3 += 46
    }
    var v25 uint64 = 302
    var v26 uint8 = (v18 & v2)
    var v27 int16 = 32767
    if ((v2 % ((v18) | 1)) < (v18 - (v18 << ((18446744073709551615) & 63)))) {
        var v28 string = ("\n\t" + "foo bar")
        _ = len(v4)
        var v29 int32 = -1
        _ = v28
        _ = v29
    }
    v30 := GenS2{F0: (v25 % (((v25 + (^(v25)))) | 1)), F1: 63, F2: 1}
    switch "\n\t" {
    case "a":
        v27 = v27
    case "foo bar":
        fmt.Println(v3, v27)
    default:
        for fi31 := 0; fi31 < 14; fi31++ {
            var v32 int8 = 100
            var v33 float64 = 1.5
            for fi34 := 0; fi34 < 20; fi34++ {
                v35 := GenS2{F0: v25, F1: v32, F2: 2147483647}
                var v36 uint16 = (v3 - 497)
                var v37 uint16 = (v36 + v16)
                _ = fi34
                _ = v35
                _ = v36
                _ = v37
            }
            _ = fi31
            _ = v32
            _ = v33
        }
    }
    for fi38 := 0; fi38 < 13; fi38++ {
        fmt.Println(v27, v2)
        var v39 float64 = float64(float32(1e10))
        fmt.Println(v18, v1, v18)
        _ = fi38
        _ = v39
    }
    fmt.Println(v25)
    v3 += 65535
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
}
