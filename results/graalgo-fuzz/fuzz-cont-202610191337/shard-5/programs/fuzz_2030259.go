package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint16; F1 uint64; F2 uint16 }
func genFn0(p0 bool, p1 int16) bool {
    var v7 bool = (-1 > 9223372036854775807)
    var v8 int8 = -1
    v9 := []int{37, -2147483648, 9223372036854775807}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return v7
}
func genFn1(p0 uint16, p1 bool, p2 uint8) uint64 {
    var v10 int16 = 944
    v11 := GenS1{F0: 74}
    var v12 uint = 1
    fmt.Println(p2, p2)
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 int16 = -307
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return 1
}

func main() {
    v1 := []uint32{2147483647, 42, 65535, 955}
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]bool)
    var v4 uint = 18446744073709551615
    var v5 uint32 = 4294967295
    var v6 int16 = 32767
    switch "hello" {
    case "a":
        if false {
            v5 -= v5
            var v15 float64 = 3.14159
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v15
        } else {
            v4 -= v4
            v5 -= 283
            var v16 bool = true
            _ = v16
        }
    case "foo bar":
        if false {
            if (!true) {
                for fi17 := 0; fi17 < 10; fi17++ {
                    if false {
                        var v18 float32 = 0.1
                        var v19 uint32 = 404
                        _ = v18
                        _ = v19
                    }
                    _ = fi17
                }
                for fi20 := 0; fi20 < 3; fi20++ {
                    fmt.Println(v4, v6)
                    v4 = v4
                    if false {
                        switch "hello" {
                        case "foo bar":
                            fmt.Println(v4, v5, v5)
                        case "\n\t":
                            v21 := GenS1{F0: v5}
                            _ = v21
                        case "":
                            fmt.Println(fi20, v6)
                        default:
                            v4 = (^(v4))
                        }
                    }
                    _ = fi20
                }
            } else {
                var v22 int16 = 42
                for fi23 := 0; fi23 < 10; fi23++ {
                    v6 -= v22
                    _ = fi23
                }
                _ = v22
            }
            v6 += v6
        }
    case "\n\t":
        if true {
            for fi24 := 0; fi24 < 9; fi24++ {
                fmt.Println(v5, v6, v5)
                if (42 <= 0) {
                    for fi25 := 0; fi25 < 17; fi25++ {
                        v26 := GenS2{F0: true, F1: 0, F2: 65535}
                        fmt.Println(v5, fi24)
                        var v27 string = "foo bar"
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                }
                _ = fi24
            }
        } else {
            v4 = 0
            var v28 string = ("a" + (("a" + "a") + ""))
            _ = v28
        }
    default:
        switch 343 {
        case 0:
            switch ("hello" + (("foo bar" + "hello") + "hello")) {
            case "hello":
                _ = len(v2)
            case "\n\t":
                if ((v6 < 659) || true) {
                    if true {
                        v29 := GenS3{F0: 755, F1: 594, F2: 620}
                        v4 = v4
                        if false {
                            _ = len(v3)
                            fmt.Println(v5, v4, v5)
                        } else {
                            switch 700 {
                            case 32768:
                                var v30 uint = 385
                                _ = v30
                            case 42:
                                var v31 uint32 = v5
                                _ = v31
                            default:
                                v6 = v6
                            }
                        }
                        _ = v29
                    } else {
                        _ = len(v3)
                    }
                    v2[65535] = (v5 % (((^(v5))) | 1))
                } else {
                    v4 += (v4 / (((v4 % (((v4 + v4)) | 1))) | 1))
                    v5 += (v5 >> (((v4 | 18446744073709551615)) & 63))
                }
            default:
                v32 := GenS2{F0: (v5 == v5), F1: 42, F2: v5}
                _ = v32
            }
        case -973:
            fmt.Println(v4, v6, v5)
        }
    }
    v5 += v5
    delete(v2, v5)
    if true {
        var v33 uint = v4
        var v34 uint = v33
        v5 -= (^(v5))
        _ = v33
        _ = v34
    }
    v5 += (^(v5))
    var v35 uint8 = 0
    var v36 uint32 = v5
    var v37 uint64 = 2
    if (453 <= (v5 << ((255) & 63))) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v35 = (^(v35))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
