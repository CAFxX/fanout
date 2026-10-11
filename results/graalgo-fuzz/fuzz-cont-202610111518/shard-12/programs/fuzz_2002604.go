package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 int8; F1 uint8; F2 uint32 }
func genFn0() uint {
    var v4 int16 = 1
    var v5 string = ("hello" + "\n\t")
    var v6 float64 = -2.25
    switch 659 {
    case 32767:
        fmt.Println(v5, v4, v6)
    default:
        v7 := GenS1{F0: 93, F1: 2, F2: -782}
        _ = v7
    }
    v4 = (v4 - -32768)
    _ = v4
    _ = v5
    _ = v6
    return 707
}
func genFn1(p0 uint32, p1 uint32) uint32 {
    fmt.Println(p1)
    p0 = 2147483647
    _ = p0
    _ = p1
    return p0
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{-2147483648, 4294967295}
    v3 := make(map[int]bool)
    var v8 int64 = 42
    var v9 bool = (false && (!false))
    for fi10 := 0; fi10 < 15; fi10++ {
        v8 = 429
        _ = fi10
    }
    v11 := GenS2{F0: "\n\t", F1: 9223372036854775807}
    v8 += -9223372036854775808
    delete(v3, 2147483647)
    v2 = append(v2, 100)
    v9 = v9
    if (!true) {
        delete(v3, 2147483647)
        v12 := GenS2{F0: "a", F1: 520}
        _ = v12
    }
    if (("hello" < "") && (!v9)) {
        v13 := GenS3{F0: 2, F1: 0}
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 uint64 = 216
            switch (fi14 << ((0) & 63)) {
            case 2147483647:
                var v16 int16 = 1
                _ = v16
            case 32767:
                if (0 < -1) {
                    v9 = v9
                }
            default:
                switch ("a" + "foo bar") {
                case "hello":
                    v8 = (-(v8))
                default:
                    if (349 != v8) {
                        v17 := GenS1{F0: 4294967295, F1: (v15 ^ 130), F2: (-(fi14))}
                        v9 = false
                        v9 = (!v9)
                        _ = v17
                    }
                }
            }
            _ = fi14
            _ = v15
        }
        _ = v13
    } else {
        v9 = v9
        if v9 {
            var v18 uint = 2
            _ = v18
        } else {
            for fi19 := 0; fi19 < 13; fi19++ {
                v8 = 42
                _ = fi19
            }
        }
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        v21 := GenS3{F0: 336, F1: 4294967295}
        var v22 int64 = v8
        _ = fi20
        _ = v21
        _ = v22
    }
    if v9 {
        v23 := GenS4{F0: 100, F1: 255, F2: 23}
        v8 += (-(v8))
        _ = v23
    } else {
        var v24 bool = (v9 && v9)
        v8 = v8
        for fi25 := 0; fi25 < 4; fi25++ {
            v24 = false
            v24 = true
            delete(v3, 4294967295)
            _ = fi25
        }
        _ = v24
    }
    var v26 int8 = -108
    if v9 {
        for fi27 := 0; fi27 < 14; fi27++ {
            for fi28 := 0; fi28 < 2; fi28++ {
                v3[(fi27 ^ fi27)] = ("" >= "hello")
                var v29 int8 = v26
                v2 = append(v2, (fi28 + fi27))
                _ = fi28
                _ = v29
            }
            var v30 float64 = (0.0 - float64(-2.25))
            var v31 uint8 = 157
            _ = fi27
            _ = v30
            _ = v31
        }
    } else {
        v8 = (v8 | (-(v8)))
        if v9 {
            for fi32 := 0; fi32 < 7; fi32++ {
                var v33 float32 = 0.1
                v34 := GenS3{F0: 572, F1: genFn0()}
                v35 := GenS4{F0: 100, F1: 2, F2: 2}
                _ = fi32
                _ = v33
                _ = v34
                _ = v35
            }
            v36 := GenS4{F0: (v26 / (((v26 * v26)) | 1)), F1: 84, F2: genFn1(genFn1(128, 430), 1)}
            fmt.Println(v26, v9)
            _ = v36
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        for fi37 := 0; fi37 < 20; fi37++ {
            var v38 uint16 = 1
            _ = fi37
            _ = v38
        }
    }
    fmt.Println(v8, v8)
    v2 = append(v2, 43)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v26
    _ = v3
    _ = v8
    _ = v9
}
