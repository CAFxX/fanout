package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int8 }
type GenS2 struct { F0 string; F1 uint16; F2 uint }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint8; F1 int8 }
func genFn0(p0 int, p1 bool) uint32 {
    p1 = (128 <= 1)
    if p1 {
        p1 = (32768 >= 1)
        if false {
            v4 := []int{p0, p0, (p0 & (p0 ^ (p0 + -259))), p0}
            _ = v4
        }
        v5 := GenS1{F0: 127, F1: 706, F2: -1}
        _ = v5
    }
    for fi6 := 0; fi6 < 8; fi6++ {
        if (("foo bar" + "\n\t") > "hello") {
            for fi7 := 0; fi7 < 20; fi7++ {
                v8 := GenS1{F0: -1, F1: 2147483647, F2: -52}
                var v9 bool = ((p1 || p1) && ((!(-85 > 0)) || (!(!(float64((0.1 + 1.5)) >= 0.1)))))
                _ = fi7
                _ = v8
                _ = v9
            }
        }
        var v10 string = "\n\t"
        _ = fi6
        _ = v10
    }
    v11 := GenS2{F0: ("a" + "foo bar"), F1: 255, F2: 4294967295}
    _ = p0
    _ = p1
    _ = v11
    return 4294967295
}
func genFn1(p0 uint16, p1 string, p2 int) bool {
    switch p2 {
    case -847:
        v12 := []int{-649, (p2 - (p2 ^ (p2 | 9223372036854775807)))}
        _ = v12
    case 128:
        fmt.Println(p2, p1, p2)
    case 360:
        var v13 string = ((p1 + "a") + (((p1 + ("hello" + p1)) + "foo bar") + p1))
        _ = v13
    default:
        var v14 int = (p2 << ((18446744073709551615) & 63))
        _ = v14
    }
    v15 := GenS3{F0: false, F1: 1}
    switch 0 {
    case 4294967295:
        v16 := GenS4{F0: 128, F1: -100}
        _ = v16
    case -1:
        v17 := []int{(p2 & (p2 - (p2 >> ((642) & 63))))}
        _ = v17
    case -687:
        if (42 > 9223372036854775807) {
            var v18 int32 = 1
            _ = v18
        }
    default:
        p2 = p2
    }
    p1 = "\n\t"
    if true {
        v19 := []int{(p2 * p2)}
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return true
}

func main() {
    v1 := []uint32{957}
    v2 := make(map[uint32]int)
    var v3 uint32 = 2147483648
    for fi20 := 0; fi20 < 19; fi20++ {
        v3 = (v3 / (((v3 << ((2) & 63))) | 1))
        _ = fi20
    }
    switch "a" {
    case "\n\t":
        var v21 int32 = 577
        _ = v21
    default:
        var v22 int8 = -128
        _ = v22
    }
    delete(v2, v3)
    var v23 int8 = -1
    var v24 string = "hello"
    v24 = "foo bar"
    v3 = v3
    for fi25 := 0; fi25 < 11; fi25++ {
        v24 = v24
        for fi26 := 0; fi26 < 7; fi26++ {
            var v27 string = v24
            _ = len(v2)
            _ = fi26
            _ = v27
        }
        _ = fi25
    }
    _ = len(v2)
    for fi28 := 0; fi28 < 2; fi28++ {
        v29 := v2[v3]
        switch (v29 - fi28) {
        case -31:
            var v30 int64 = -9223372036854775808
            _ = v30
        case -339:
            if (genFn0(127, false) == (v3 & genFn0((v29 + (v29 ^ (^(v29)))), (true && false)))) {
                var v31 int8 = v23
                _ = v31
            }
        }
        for fi32 := 0; fi32 < 15; fi32++ {
            var v33 string = "hello"
            _ = fi32
            _ = v33
        }
        _ = fi28
        _ = v29
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v1 = append(v1, 1)
    switch 42 {
    case 128:
        var v34 uint = 1
        _ = v34
    case -38:
        _ = len(v1)
    case 32767:
        v23 -= v23
    default:
        for fi35 := 0; fi35 < 15; fi35++ {
            var v36 uint64 = 18446744073709551615
            fmt.Println(v3, fi35)
            if true {
                v37 := GenS4{F0: 1, F1: v23}
                v36 = (v36 * (v36 << ((2) & 63)))
                _ = v37
            } else {
                for fi38 := 0; fi38 < 7; fi38++ {
                    var v39 int32 = 428
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi38
                    _ = v39
                }
            }
            _ = fi35
            _ = v36
        }
    }
    for fi40 := 0; fi40 < 4; fi40++ {
        fmt.Println(v24)
        var v41 uint64 = 1
        switch v24 {
        case "a":
            if false {
                fmt.Println(v41, v24)
                var v42 uint8 = 255
                _ = v42
            } else {
                var v43 int64 = -603
                _ = v43
            }
        case "foo bar":
            var v44 uint64 = v41
            _ = v44
        case "\n\t":
            _ = len(v2)
        default:
            v24 = v24
        }
        _ = fi40
        _ = v41
    }
    if len(v1) > 0 {
        v1[0] = (v3 * (v3 * v3))
    }
    var v45 int = 2147483647
    fmt.Println(v3, v24)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v45
}
