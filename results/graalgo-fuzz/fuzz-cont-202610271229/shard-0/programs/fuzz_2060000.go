package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int32; F1 string; F2 int }
type GenS3 struct { F0 int; F1 string; F2 uint64 }
type GenS4 struct { F0 int32; F1 uint }
func genFn0(p0 uint64, p1 int8) int64 {
    if false {
        v4 := GenS1{F0: p1, F1: -454}
        for fi5 := 0; fi5 < 13; fi5++ {
            p1 -= p1
            var v6 uint16 = 2
            v7 := []int{(-(fi5)), fi5, int(p1)}
            _ = fi5
            _ = v6
            _ = v7
        }
        _ = v4
    }
    fmt.Println(p1, p0)
    p0 -= (p0 * 376)
    v8 := GenS2{F0: -1, F1: ("hello" + "foo bar"), F2: 42}
    var v9 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := []uint32{}
    var v2 uint32 = 4294967295
    var v3 int64 = 42
    var v10 uint16 = 2
    if false {
        fmt.Println(v2, v2)
    }
    v2 = (v2 >> ((0) & 63))
    for fi11 := 0; fi11 < 2; fi11++ {
        if true {
            if false {
                for fi12 := 0; fi12 < 15; fi12++ {
                    v3 = (v3 / ((v3) | 1))
                    if (((!true) || false) && (!(-849 == 42))) {
                        v3 -= genFn0(42, 100)
                    }
                    var v13 uint64 = 42
                    _ = fi12
                    _ = v13
                }
                v3 = genFn0(42, 0)
                v14 := GenS1{F0: 34, F1: v3}
                _ = v14
            }
        } else {
            v15 := GenS3{F0: (fi11 | fi11), F1: "a", F2: 0}
            v1 = append(v1, 327)
            v16 := GenS1{F0: int8(v10), F1: (v3 % ((v3) | 1))}
            _ = v15
            _ = v16
        }
        v2 -= 42
        for fi17 := 0; fi17 < 2; fi17++ {
            v3 = -233
            _ = fi17
        }
        _ = fi11
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        switch "" {
        case "":
            v10 = 537
        case "\n\t":
            for fi19 := 0; fi19 < 4; fi19++ {
                v10 -= v10
                var v20 float32 = -2.25
                _ = fi19
                _ = v20
            }
        default:
            var v21 bool = (508 < 2)
            _ = v21
        }
        if true {
            if false {
                var v22 uint = 1
                for fi23 := 0; fi23 < 16; fi23++ {
                    _ = len(v1)
                    if false {
                        fmt.Println(v3, v2, v22)
                        var v24 uint32 = v2
                        v3 += v3
                        _ = v24
                    }
                    _ = fi23
                }
                _ = v22
            }
            if len(v1) > 0 {
                v1[0] = 884
            }
            var v25 uint16 = 0
            _ = v25
        }
        _ = fi18
    }
    v26 := GenS4{F0: 1, F1: 4294967295}
    v2 = 2147483647
    if (0 != 224) {
        _ = len(v1)
        for fi27 := 0; fi27 < 15; fi27++ {
            fmt.Println(v2, v3, fi27)
            _ = fi27
        }
    }
    v3 = (v3 >> ((65535) & 63))
    v28 := GenS4{F0: -1, F1: 466}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v28
    _ = v3
}
