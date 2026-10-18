package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint }
type GenS2 struct { F0 uint32; F1 uint16 }
func genFn0(p0 int8) uint32 {
    if true {
        v7 := GenS1{F0: p0, F1: 1, F2: 2}
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 uint = 42
            _ = fi8
            _ = v9
        }
        _ = v7
    } else {
        var v10 uint = 1
        _ = v10
    }
    if (!(true && (!true))) {
        if true {
            switch ("a" + "") {
            case "foo bar":
                if true {
                    v11 := GenS1{F0: (p0 | (p0 | (p0 ^ p0))), F1: -1, F2: 886}
                    _ = v11
                }
            }
        }
    } else {
        switch 577 {
        case 414:
            switch -56 {
            case 128:
                var v12 bool = ("a" == "hello")
                _ = v12
            }
        case -1:
            var v13 uint = 2
            _ = v13
        default:
            v14 := GenS1{F0: (p0 / (((p0 | p0)) | 1)), F1: -2147483648, F2: 18446744073709551615}
            _ = v14
        }
        var v15 uint16 = 581
        p0 += 0
        _ = v15
    }
    var v16 float64 = 1e10
    v16 -= (float64(-2.25) - v16)
    _ = p0
    _ = v16
    return 2147483648
}

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = 1
    v3 := make(map[int]string)
    v4 := []int{int(v2), 289}
    var v5 int = 32768
    v6 := make(map[uint32]int)
    _ = len(v4)
    v5 -= (v5 << ((65535) & 63))
    _ = len(v4)
    v4 = append(v4, v5)
    v17 := GenS1{F0: (v2 % ((v2) | 1)), F1: 42, F2: 132}
    fmt.Println(v5, v5)
    for fi18 := 0; fi18 < 19; fi18++ {
        v2 = (v2 << ((678) & 63))
        if ("" != ("\n\t" + "a")) {
            fmt.Println(fi18)
            for fi19 := 0; fi19 < 16; fi19++ {
                var v20 uint8 = 89
                v4 = append(v4, v5)
                _ = fi19
                _ = v20
            }
            switch ("a" + "foo bar") {
            case "":
                for fi21 := 0; fi21 < 17; fi21++ {
                    var v22 int16 = 1000
                    _ = fi21
                    _ = v22
                }
            case "hello":
                v4 = append(v4, v5)
            case "a":
                v2 = v2
            default:
                v2 -= (v2 - v2)
            }
        } else {
            v23 := GenS1{F0: v2, F1: -1, F2: 381}
            _ = v23
        }
        for fi24 := 0; fi24 < 1; fi24++ {
            var v25 uint32 = 65535
            var v26 uint8 = 42
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = fi18
    }
    v2 -= v2
    switch 4294967295 {
    case 32768:
        v27 := GenS2{F0: 139, F1: 0}
        _ = v27
    case -9223372036854775808:
        if len(v4) > 0 {
            v4[0] = v5
        }
    default:
        var v28 uint64 = 4294967295
        _ = v28
    }
    if ((749 < v5) || false) {
        v3[v5] = "foo bar"
    } else {
        v5 = v5
        fmt.Println(v5, v2)
        for fi29 := 0; fi29 < 19; fi29++ {
            if false {
                for fi30 := 0; fi30 < 2; fi30++ {
                    var v31 uint = 1
                    var v32 uint8 = 2
                    _ = len(v4)
                    v33 := GenS1{F0: int8(v32), F1: 172, F2: v31}
                    var v34 uint = 1
                    _ = fi30
                    _ = v31
                    _ = v32
                    _ = v33
                    _ = v34
                }
                v35 := GenS2{F0: 2, F1: 255}
                _ = len(v6)
                _ = v35
            } else {
                var v36 uint32 = genFn0(v2)
                v2 = v2
                _ = v36
            }
            fmt.Println(v5)
            _ = fi29
        }
    }
    v4 = append(v4, v5)
    var v37 uint64 = 18446744073709551615
    for fi38 := 0; fi38 < 11; fi38++ {
        v5 = (v5 * -1)
        v5 += (-(fi38))
        var v39 uint16 = 414
        _ = fi38
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
