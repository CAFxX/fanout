package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS4 struct { F0 int; F1 int64; F2 uint64 }
type GenS5 struct { F0 int32; F1 int; F2 int16 }
type GenS6 struct { F0 int8; F1 uint8 }
type GenS7 struct { F0 int64 }
func genFn0(p0 uint32, p1 int64, p2 uint8) int32 {
    if true {
        switch 42 {
        case 0:
            v6 := GenS1{F0: (p0 & p0), F1: "a", F2: 0}
            _ = v6
        case 98:
            var v7 int8 = -128
            _ = v7
        }
    } else {
        for fi8 := 0; fi8 < 18; fi8++ {
            for fi9 := 0; fi9 < 7; fi9++ {
                v10 := []int{fi8, fi8, fi9, fi8}
                p2 -= 0
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
        v11 := GenS1{F0: p0, F1: "foo bar", F2: (p0 * p0)}
        _ = v11
    }
    p0 += p0
    if false {
        fmt.Println(p2)
        fmt.Println(p2, p0)
        v12 := GenS1{F0: p0, F1: (((("a" + "") + "\n\t") + "") + ("foo bar" + "\n\t")), F2: p0}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return -445
}
func genFn1() string {
    if true {
        v13 := GenS2{F0: 1}
        v14 := GenS2{F0: 2}
        var v15 uint = 1
        _ = v13
        _ = v14
        _ = v15
    }
    var v16 int16 = -1
    if (true || true) {
        v17 := GenS3{F0: 255, F1: 0, F2: 65535}
        _ = v17
    } else {
        v18 := GenS4{F0: -446, F1: int64(v16), F2: 42}
        _ = v18
    }
    _ = v16
    return ("foo bar" + ("foo bar" + ("hello" + "a")))
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 bool = false
    var v3 uint8 = 42
    v4 := make(map[string]bool)
    v5 := make(map[int]int)
    var v19 int64 = (v1 * (v1 + (v1 / (((v1 & (v1 + v1))) | 1))))
    var v20 int32 = 0
    v21 := []int{0, 2147483647, 2147483647, 531}
    v2 = ((v20 + v20) > (v20 / ((-1) | 1)))
    var v22 int = 4294967295
    for fi23 := 0; fi23 < 15; fi23++ {
        v21 = append(v21, (v22 % (((^(v22))) | 1)))
        if v2 {
            v24 := GenS2{F0: 243}
            for fi25 := 0; fi25 < 11; fi25++ {
                if ((fi25 & -700) == (v22 >> ((53) & 63))) {
                    v22 += (v22 << ((512) & 63))
                    var v26 uint16 = 42
                    v27 := GenS5{F0: (v20 * v20), F1: fi25, F2: 42}
                    var v28 uint32 = 340
                    _ = v26
                    _ = v27
                    _ = v28
                } else {
                    v20 = v20
                    var v29 uint8 = (v3 | 255)
                    _ = v29
                }
                _ = fi25
            }
            _ = v24
        } else {
            delete(v4, "foo bar")
            var v30 uint8 = v3
            _ = v30
        }
        _ = fi23
    }
    v3 = (v3 >> ((4294967295) & 63))
    if (42 > 1) {
        v31 := GenS5{F0: 0, F1: (v22 | (v22 + (v22 / ((v22) | 1)))), F2: 42}
        _ = v31
    }
    v32 := GenS4{F0: (v22 & v22), F1: (^(v19)), F2: 215}
    if ((32767 > 1) || (true || (!(-32 < 41)))) {
        fmt.Println(v20, v2, v1)
    } else {
        for fi33 := 0; fi33 < 10; fi33++ {
            delete(v4, "hello")
            v34 := GenS6{F0: 1, F1: 1}
            var v35 float32 = 0.1
            _ = fi33
            _ = v34
            _ = v35
        }
    }
    var v36 int8 = -128
    if false {
        var v37 int64 = v19
        _ = v37
    }
    v38 := GenS7{F0: (v1 * (-(v1)))}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v32
    _ = v36
    _ = v38
    _ = v4
    _ = v5
}
