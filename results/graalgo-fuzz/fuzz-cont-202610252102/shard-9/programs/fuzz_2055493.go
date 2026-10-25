package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 uint; F2 uint8 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 int64; F1 uint16 }
func genFn0(p0 bool, p1 uint16, p2 uint) string {
    var v5 int32 = 559
    if true {
        var v6 uint64 = 862
        v7 := GenS1{F0: (p1 * p1)}
        fmt.Println(v6, v5, p1)
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 uint32 = 914
        switch fi8 {
        case 835:
            fmt.Println(p0, v9, v5)
        case -3:
            var v10 float32 = 0.0
            _ = v10
        case 413:
            var v11 float64 = (1e10 - ((1e10 + 0.1) - 3.14159))
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    if (-9223372036854775808 >= 42) {
        var v12 float32 = ((1.5 - 0.1) * (3.14159 - (((0.0 + 1e10) - (-2.25 * (-2.25 + float32(0.1)))) * 0.1)))
        _ = v12
    } else {
        p1 -= 0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (("hello" + "") + "foo bar")
}
func genFn1(p0 int64, p1 int16) int8 {
    fmt.Println(p0, p0)
    p1 -= 1
    _ = p0
    _ = p1
    return 127
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]int)
    var v3 bool = false
    var v4 int32 = -157
    if (!((9223372036854775807 != 4294967295) && true)) {
        var v13 uint32 = 1
        _ = v13
    } else {
        var v14 bool = (-757 != 42)
        v4 -= (v4 << ((727) & 63))
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        v16 := []int{}
        var v17 int8 = 100
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 string = "hello"
    if true {
        var v19 uint8 = 1
        _ = v19
    } else {
        var v20 float64 = 0.1
        if (!false) {
            v4 += 403
            v3 = (!(42 < 42))
            for fi21 := 0; fi21 < 15; fi21++ {
                if v3 {
                    var v22 int32 = 0
                    v23 := GenS2{F0: false, F1: 228, F2: 153}
                    _ = v22
                    _ = v23
                } else {
                    v3 = v3
                    fmt.Println(fi21, v20, v4)
                }
                for fi24 := 0; fi24 < 18; fi24++ {
                    if false {
                        switch fi24 {
                        case 32767:
                            switch ((v18 + "") + ("foo bar" + v18)) {
                            case "":
                                v3 = false
                            default:
                                v25 := GenS3{F0: 32768, F1: true}
                                _ = v25
                            }
                        case 128:
                            v18 = "\n\t"
                        }
                        v3 = (1 == 42)
                        v26 := GenS1{F0: 1}
                        _ = v26
                    }
                    _ = fi24
                }
                var v27 float64 = (v20 * 0.0)
                _ = fi21
                _ = v27
            }
        } else {
            for fi28 := 0; fi28 < 12; fi28++ {
                v20 = v20
                delete(v1, fi28)
                var v29 bool = v3
                _ = fi28
                _ = v29
            }
            switch v18 {
            case "":
                if (!v3) {
                    v18 = v18
                }
            }
            delete(v2, 1)
        }
        v3 = (4294967295 != 255)
        _ = v20
    }
    var v30 int32 = v4
    if (-120 > -9223372036854775808) {
        fmt.Println(v4, v18)
        fmt.Println(v30, v30)
        v4 = (v4 << ((0) & 63))
    }
    v31 := GenS4{F0: -2, F1: 860}
    v32 := []int{}
    v4 = 42
    fmt.Println(v30)
    for fi33 := 0; fi33 < 17; fi33++ {
        v4 += 42
        _ = fi33
    }
    for fi34 := 0; fi34 < 3; fi34++ {
        v4 = v4
        _ = fi34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
