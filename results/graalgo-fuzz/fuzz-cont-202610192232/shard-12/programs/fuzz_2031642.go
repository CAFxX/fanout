package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64; F1 uint16; F2 int }
type GenS6 struct { F0 uint64; F1 int; F2 uint16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 18446744073709551615
    var v3 int16 = -503
    var v4 uint32 = 2147483647
    delete(v1, 810)
    v5 := GenS1{F0: 0}
    var v6 int32 = 1
    v4 = 270
    v6 += 2147483647
    if (-79 == 4294967295) {
        if false {
            v7 := GenS2{F0: 128, F1: -1}
            _ = v7
        } else {
            for fi8 := 0; fi8 < 5; fi8++ {
                for fi9 := 0; fi9 < 15; fi9++ {
                    if (!(!false)) {
                        v10 := GenS1{F0: 0}
                        var v11 uint8 = 1
                        v12 := GenS2{F0: fi8, F1: -1}
                        _ = v10
                        _ = v11
                        _ = v12
                    } else {
                        fmt.Println(v6)
                        if true {
                            _ = len(v1)
                            v13 := GenS3{F0: 0, F1: v6}
                            _ = v13
                        }
                        v4 = (^(v4))
                        if true {
                            if false {
                                fmt.Println(v3)
                            }
                            fmt.Println(v2, v4, fi8)
                        }
                        if false {
                            v14 := []int{fi9, int(v3), (fi9 | int(v3))}
                            _ = v14
                        } else {
                            v6 = int32(v3)
                        }
                        v6 = v6
                        v2 = v2
                        var v15 uint32 = (^(v4))
                        _ = v15
                    }
                    v3 += 1000
                    _ = fi9
                }
                for fi16 := 0; fi16 < 18; fi16++ {
                    v17 := []int{fi8, (-(fi8)), (^(fi8)), fi16}
                    fmt.Println(v6, v4, fi16)
                    _ = fi16
                    _ = v17
                }
                v18 := GenS4{F0: 0}
                _ = fi8
                _ = v18
            }
            switch -834 {
            case 32768:
                v19 := []int{820, -661}
                _ = v19
            case 42:
                v1[(v4 | v4)] = -2147483648
            case 0:
                v20 := []int{-2147483648, -2147483648}
                _ = v20
            default:
                var v21 string = "foo bar"
                _ = v21
            }
            for fi22 := 0; fi22 < 18; fi22++ {
                v3 -= v3
                _ = len(v1)
                _ = fi22
            }
        }
        var v23 bool = (2 <= 42)
        for fi24 := 0; fi24 < 7; fi24++ {
            var v25 float64 = (3.14159 - 0.1)
            v4 = (v4 + v4)
            var v26 int16 = v3
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = v23
    } else {
        v27 := []int{32767, int(v4), -413, 128}
        _ = v27
    }
    delete(v1, v4)
    for fi28 := 0; fi28 < 19; fi28++ {
        fmt.Println(v4, fi28)
        v2 -= (v2 + 235)
        _ = fi28
    }
    v6 = -1
    _ = len(v1)
    v29 := []int{0, -458, 4294967295}
    if (1000 >= (v3 << ((v2) & 63))) {
        fmt.Println(v3, v2)
        for fi30 := 0; fi30 < 11; fi30++ {
            var v31 uint64 = 0
            _ = fi30
            _ = v31
        }
        if true {
            for fi32 := 0; fi32 < 17; fi32++ {
                switch "\n\t" {
                case "foo bar":
                    var v33 uint = (v2 - v2)
                    _ = v33
                default:
                    for fi34 := 0; fi34 < 17; fi34++ {
                        if len(v29) >= 2 {
                            v29 = v29[1:len(v29)]
                        }
                        if false {
                            v35 := GenS5{F0: 0, F1: uint16(v2), F2: fi34}
                            var v36 uint8 = 2
                            _ = v35
                            _ = v36
                        }
                        _ = fi34
                    }
                }
                _ = fi32
            }
            switch 2147483647 {
            case 127:
                v37 := GenS6{F0: 42, F1: 503, F2: 1}
                _ = v37
            }
            switch "a" {
            case "a":
                v2 = 2
            case "\n\t":
                var v38 uint64 = 78
                _ = v38
            }
        }
    }
    delete(v1, v4)
    v1[(v4 & v4)] = 4294967295
    for fi39 := 0; fi39 < 20; fi39++ {
        v3 += (v3 % ((v3) | 1))
        _ = fi39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
