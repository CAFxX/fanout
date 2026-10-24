package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int8 }
func genFn0(p0 int8, p1 uint8, p2 int16) uint {
    var v6 int = 4294967295
    var v7 int32 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}
func genFn1() uint64 {
    switch 127 {
    case -328:
        switch 28 {
        case 500:
            for fi8 := 0; fi8 < 13; fi8++ {
                var v9 uint32 = 4294967295
                _ = fi8
                _ = v9
            }
        case -397:
            if (0 != 1) {
                var v10 int64 = -31
                var v11 uint8 = 1
                _ = v10
                _ = v11
            } else {
                var v12 int16 = -1
                _ = v12
            }
        }
    }
    var v13 int16 = 42
    v14 := []int{}
    switch -1 {
    case 398:
        var v15 uint32 = 400
        _ = v15
    }
    _ = v13
    _ = v14
    return 18446744073709551615
}

func main() {
    v1 := []uint32{963, 42}
    v2 := []uint32{42, 314, 4294967295, 42}
    var v3 string = (("foo bar" + "foo bar") + "a")
    v4 := []int{127, 0, -791}
    v5 := make(map[int]string)
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 bool = (((497 > 32767) || true) && (9223372036854775807 != 1))
        var v18 bool = (!(2 >= 168))
        for fi19 := 0; fi19 < 11; fi19++ {
            v17 = v18
            _ = fi19
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    for fi20 := 0; fi20 < 11; fi20++ {
        for fi21 := 0; fi21 < 16; fi21++ {
            v3 = v3
            _ = fi21
        }
        v22 := GenS1{F0: 1, F1: 42, F2: 255}
        var v23 uint = 2
        _ = fi20
        _ = v22
        _ = v23
    }
    var v24 uint64 = 1
    v3 = ((v3 + ("" + v3)) + v3)
    switch ((("\n\t" + v3) + ("\n\t" + (("" + v3) + ""))) + (v3 + v3)) {
    case "foo bar":
        v25 := GenS2{F0: -1, F1: 255}
        _ = v25
    default:
        fmt.Println(v24)
    }
    var v26 float64 = 0.1
    switch 0 {
    case 56:
        v27 := GenS3{F0: 44}
        _ = v27
    }
    switch 0 {
    case -805:
        v26 -= 1.5
    default:
        if (!false) {
            var v28 int16 = 32767
            var v29 uint64 = (v24 ^ genFn1())
            _ = v28
            _ = v29
        } else {
            fmt.Println(v26)
            v30 := GenS2{F0: 2147483647, F1: 195}
            _ = v30
        }
    }
    fmt.Println(v26, v26, v3)
    v24 += genFn1()
    for fi31 := 0; fi31 < 15; fi31++ {
        var v32 int = (fi31 & (^(fi31)))
        if false {
            var v33 uint8 = 255
            if true {
                for fi34 := 0; fi34 < 3; fi34++ {
                    v32 = 4294967295
                    v35 := GenS3{F0: 107}
                    _ = fi34
                    _ = v35
                }
                if false {
                    var v36 uint32 = 168
                    fmt.Println(fi31)
                    _ = v36
                } else {
                    if len(v1) > 0 {
                        v1[0] = 0
                    }
                }
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v33
        } else {
            var v37 string = v3
            _ = v37
        }
        _ = fi31
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
