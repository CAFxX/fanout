package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int32 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 uint32; F1 int32; F2 uint }
type GenS5 struct { F0 bool; F1 string; F2 uint }
func genFn0(p0 int32) int16 {
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := GenS1{F0: 108, F1: (fi5 & fi5), F2: (p0 >> ((42) & 63))}
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        p0 = 1
        if ((18446744073709551615 >= 1) && (307 < -32768)) {
            if true {
                v8 := GenS1{F0: 127, F1: fi7, F2: -255}
                for fi9 := 0; fi9 < 17; fi9++ {
                    v10 := GenS1{F0: 100, F1: 513, F2: 968}
                    fmt.Println(fi9)
                    p0 = p0
                    _ = fi9
                    _ = v10
                }
                for fi11 := 0; fi11 < 14; fi11++ {
                    p0 -= 1
                    v12 := GenS2{F0: 42, F1: 161}
                    _ = fi11
                    _ = v12
                }
                _ = v8
            }
            fmt.Println(p0)
        } else {
            if ((0.1 + 0.0) <= (float32(0.0) * 1.5)) {
                var v13 bool = false
                if v13 {
                    var v14 int8 = 42
                    _ = v14
                }
                _ = v13
            } else {
                v15 := GenS3{F0: 2, F1: 2}
                for fi16 := 0; fi16 < 10; fi16++ {
                    v17 := GenS4{F0: 65535, F1: (^(p0)), F2: 42}
                    p0 = (p0 * (p0 / ((p0) | 1)))
                    v18 := GenS5{F0: (fi7 > (fi16 * 4294967295)), F1: ("foo bar" + (("" + "") + "foo bar")), F2: 4294967295}
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                _ = v15
            }
        }
        for fi19 := 0; fi19 < 18; fi19++ {
            var v20 int16 = -32768
            if false {
                v21 := []int{}
                _ = v21
            }
            v20 = (^(v20))
            _ = fi19
            _ = v20
        }
        _ = fi7
    }
    var v22 int16 = 42
    v22 = -558
    _ = p0
    _ = v22
    return (^(v22))
}

func main() {
    var v1 uint32 = 42
    var v2 int8 = 42
    var v3 uint64 = 42
    var v4 int = -741
    var v23 int8 = (-(v2))
    v24 := []int{(v4 % ((348) | 1)), 800, (v4 ^ (v4 / ((v4) | 1))), (v4 * 42)}
    v24 = append(v24, 4294967295)
    var v25 int64 = -397
    var v26 uint = 18446744073709551615
    v27 := GenS5{F0: (82 > v23), F1: "", F2: uint(v1)}
    v26 = v26
    var v28 int16 = genFn0(-1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
