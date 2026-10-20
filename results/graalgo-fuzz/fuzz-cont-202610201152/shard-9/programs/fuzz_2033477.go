package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 bool; F1 uint16 }
type GenS6 struct { F0 bool }
func genFn0(p0 string, p1 int16) uint16 {
    v5 := []int{}
    v6 := GenS1{F0: p1, F1: 42}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 uint16, p1 uint, p2 uint32) int32 {
    if false {
        p1 = uint(p2)
        fmt.Println(p2, p0)
        v7 := []int{2147483647, 32767, 128, 127}
        _ = v7
    }
    var v8 int = 127
    if false {
        switch "a" {
        case "foo bar":
            for fi9 := 0; fi9 < 8; fi9++ {
                p1 = 2
                _ = fi9
            }
        }
        switch ("\n\t" + "foo bar") {
        case "":
            p0 = p0
        case "a":
            p0 += p0
        case "hello":
            switch (v8 >> ((p1) & 63)) {
            case 127:
                if true {
                    var v10 uint8 = 230
                    _ = v10
                } else {
                    var v11 int64 = 970
                    switch v8 {
                    case 32767:
                        v12 := GenS2{F0: p1}
                        _ = v12
                    case -946:
                        switch "\n\t" {
                        case "":
                            switch "hello" {
                            case "hello":
                                p2 -= 42
                            default:
                                for fi13 := 0; fi13 < 17; fi13++ {
                                    switch 32768 {
                                    case 42:
                                        var v14 int16 = 142
                                        _ = v14
                                    case 32768:
                                        var v15 uint32 = (p2 | (p2 + 839))
                                        _ = v15
                                    default:
                                        var v16 float32 = 1e10
                                        _ = v16
                                    }
                                    var v17 int16 = -32768
                                    _ = fi13
                                    _ = v17
                                }
                            }
                        }
                    }
                    var v18 uint32 = (^(p2))
                    _ = v11
                    _ = v18
                }
            default:
                var v19 float32 = 1e10
                _ = v19
            }
        default:
            fmt.Println(v8, v8, v8)
        }
    } else {
        fmt.Println(v8)
        var v20 string = ("a" + ("\n\t" + "a"))
        _ = v20
    }
    v21 := GenS1{F0: -32768, F1: 27}
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v8
    return 1
}

func main() {
    var v1 uint = 1
    var v2 int32 = 1
    var v3 int8 = 42
    v4 := make(map[uint32]uint32)
    v22 := GenS1{F0: 42, F1: 127}
    v23 := GenS1{F0: 0, F1: 0}
    for fi24 := 0; fi24 < 10; fi24++ {
        v3 = -128
        var v25 int8 = -128
        _ = fi24
        _ = v25
    }
    v26 := GenS2{F0: v1}
    switch 0 {
    case -2147483648:
        var v27 uint16 = genFn0("hello", 916)
        _ = v27
    case 128:
        var v28 uint64 = 18446744073709551615
        _ = v28
    }
    var v29 int32 = (v2 | 1)
    v1 += 828
    v30 := GenS1{F0: -390, F1: 27}
    if true {
        fmt.Println(v1)
        v31 := GenS1{F0: -1, F1: 21}
        _ = v31
    }
    fmt.Println(v2)
    if false {
        delete(v4, 258)
        for fi32 := 0; fi32 < 20; fi32++ {
            fmt.Println(v29)
            _ = fi32
        }
    } else {
        var v33 int64 = 42
        for fi34 := 0; fi34 < 19; fi34++ {
            for fi35 := 0; fi35 < 3; fi35++ {
                if false {
                    var v36 uint32 = 42
                    v37 := []int{fi35, (fi35 - fi34), -291, 32768}
                    v38 := GenS3{F0: 65535, F1: (^(v33))}
                    v39 := GenS2{F0: 798}
                    var v40 float32 = -2.25
                    _ = v36
                    _ = v37
                    _ = v38
                    _ = v39
                    _ = v40
                }
                _ = fi35
            }
            if true {
                v41 := GenS4{F0: "", F1: 4294967295}
                for fi42 := 0; fi42 < 17; fi42++ {
                    if false {
                        if true {
                            if (-951 <= 42) {
                                v43 := GenS5{F0: (!true), F1: 917}
                                _ = v43
                            }
                        } else {
                            v44 := []int{fi42, fi42}
                            _ = v44
                        }
                    }
                    _ = fi42
                }
                _ = v41
            }
            v45 := GenS6{F0: true}
            _ = fi34
            _ = v45
        }
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
