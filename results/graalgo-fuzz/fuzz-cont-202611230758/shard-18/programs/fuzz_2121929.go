package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 string; F1 int32 }
func genFn0(p0 uint32, p1 int) string {
    switch (("" + ("hello" + "a")) + "\n\t") {
    case "\n\t":
        for fi4 := 0; fi4 < 19; fi4++ {
            p0 = p0
            p1 = p1
            _ = fi4
        }
    }
    v5 := []int{}
    _ = p0
    _ = p1
    _ = v5
    return "hello"
}
func genFn1(p0 string, p1 int64) int16 {
    p1 = p1
    var v6 uint16 = 40
    var v7 uint = 252
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 32767
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]string)
    var v3 uint = 186
    var v8 float64 = (((1e10 + ((0.1 * 1e10) + 0.0)) + 0.1) * 1e10)
    v8 -= v8
    fmt.Println(v3, v8)
    for fi9 := 0; fi9 < 2; fi9++ {
        v3 -= v3
        _ = len(v1)
        var v10 int32 = 2147483647
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 uint16 = 42
        var v13 int16 = 0
        switch "foo bar" {
        case "":
            for fi14 := 0; fi14 < 7; fi14++ {
                v8 += (v8 * v8)
                _ = fi14
            }
        default:
            v3 = 4294967295
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    var v15 int16 = 1000
    var v16 uint64 = 901
    v17 := GenS1{F0: 128, F1: v3, F2: 42}
    if false {
        if (false && (42 > 2147483647)) {
            v18 := GenS2{F0: 0, F1: (false && false)}
            switch 32767 {
            case -9223372036854775808:
                if (true || (255 <= v3)) {
                    for fi19 := 0; fi19 < 18; fi19++ {
                        var v20 uint64 = v16
                        var v21 float32 = -2.25
                        v16 -= (v20 % (((v16 & 13)) | 1))
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 18; fi22++ {
                        v15 = v15
                        v16 -= (v16 << ((v3) & 63))
                        var v23 bool = ((v8 < ((v8 * v8) + -2.25)) && true)
                        _ = fi22
                        _ = v23
                    }
                    fmt.Println(v3, v3, v16)
                }
            case 9223372036854775807:
                switch 1 {
                case 4294967295:
                    v15 = genFn1("a", 42)
                }
            default:
                v24 := GenS3{F0: ((("foo bar" + "a") + "foo bar") + ("hello" + "hello")), F1: -1}
                _ = v24
            }
            _ = v18
        }
        var v25 bool = ((false || true) || ((^(v3)) < (v3 - uint(v16))))
        v15 = 1000
        _ = v25
    }
    if true {
        var v26 bool = (false || true)
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v8
}
