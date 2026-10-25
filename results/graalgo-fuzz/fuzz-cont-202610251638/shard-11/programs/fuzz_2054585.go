package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int16 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int16 }
func genFn0(p0 string, p1 uint16, p2 uint8) int16 {
    p1 = (p1 | p1)
    var v4 bool = (2 < 936)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -1
}
func genFn1(p0 int8, p1 string) bool {
    switch "a" {
    case "a":
        fmt.Println(p0, p0)
    case "\n\t":
        fmt.Println(p1)
    case "":
        switch 2147483647 {
        case -9223372036854775808:
            p1 = p1
        case 326:
            p0 -= p0
        case -2147483648:
            if true {
                p0 = (^(p0))
                if true {
                    var v5 float64 = (0.0 + 0.1)
                    v6 := []int{257, -2147483648}
                    var v7 float64 = v5
                    _ = v5
                    _ = v6
                    _ = v7
                } else {
                    p1 = "a"
                }
            } else {
                v8 := []int{128, -1}
                if true {
                    p0 = (p0 >> ((2) & 63))
                }
                if false {
                    p0 = 42
                    fmt.Println(p1)
                    p0 -= p0
                } else {
                    fmt.Println(p1)
                    p0 = (p0 >> ((651) & 63))
                }
                _ = v8
            }
        default:
            if true {
                v9 := GenS1{F0: -209, F1: 36, F2: 32767}
                for fi10 := 0; fi10 < 6; fi10++ {
                    fmt.Println(fi10)
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(fi10, p1, p0)
                    default:
                        var v11 int16 = 0
                        _ = v11
                    }
                    _ = fi10
                }
                _ = v9
            }
        }
    }
    var v12 int = 788
    if true {
        fmt.Println(v12, v12, p1)
    } else {
        v13 := GenS2{F0: 9223372036854775807, F1: 234}
        v12 = v12
        _ = v13
    }
    v14 := []int{(v12 << ((2) & 63)), (v12 & v12), v12}
    _ = p0
    _ = p1
    _ = v12
    _ = v14
    return true
}

func main() {
    var v1 uint16 = 27
    v2 := []uint32{}
    var v3 uint = 65535
    for fi15 := 0; fi15 < 7; fi15++ {
        v16 := GenS2{F0: (fi15 * (fi15 << ((v3) & 63))), F1: 172}
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        v1 = v1
        v18 := GenS2{F0: (fi17 & (fi17 ^ 42)), F1: 1}
        var v19 uint32 = 783
        _ = fi17
        _ = v18
        _ = v19
    }
    v3 -= v3
    var v20 int16 = genFn0("hello", 0, 2)
    v21 := GenS3{F0: 159, F1: (v1 >> ((v3) & 63)), F2: (v20 + v20)}
    var v22 uint64 = 42
    if true {
        if false {
            fmt.Println(v1)
            if (false || true) {
                fmt.Println(v1, v1, v3)
                var v23 int32 = int32(v22)
                _ = v23
            } else {
                var v24 uint8 = 1
                var v25 bool = true
                switch "" {
                case "hello":
                    v24 = (v24 >> ((v3) & 63))
                case "":
                    if (!("hello" != "foo bar")) {
                        var v26 int32 = -2147483648
                        var v27 uint8 = v24
                        _ = v26
                        _ = v27
                    } else {
                        switch -2147483648 {
                        case -669:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case 550:
                            var v28 int64 = 42
                            _ = v28
                        case 4294967295:
                            if false {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        }
                    }
                case "foo bar":
                    var v29 int32 = -247
                    _ = v29
                }
                _ = v24
                _ = v25
            }
        }
    } else {
        v22 += (v22 | v22)
    }
    for fi30 := 0; fi30 < 7; fi30++ {
        v31 := GenS1{F0: int(v1), F1: 2, F2: (v20 + int16(v22))}
        for fi32 := 0; fi32 < 4; fi32++ {
            v22 -= v22
            if (!true) {
                var v33 int8 = 77
                _ = v33
            }
            _ = fi32
        }
        _ = fi30
        _ = v31
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    _ = len(v2)
    for fi34 := 0; fi34 < 11; fi34++ {
        switch fi34 {
        case -2147483648:
            if true {
                _ = len(v2)
                fmt.Println(v3, v3)
            }
        case 548:
            var v35 uint = v3
            _ = v35
        case -1:
            v22 = v22
        default:
            switch fi34 {
            case -667:
                v20 += (v20 ^ 42)
            case 32767:
                if false {
                    v3 += 422
                    var v36 int32 = -2147483648
                    _ = v36
                }
            }
        }
        _ = fi34
    }
    var v37 int = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v37
}
