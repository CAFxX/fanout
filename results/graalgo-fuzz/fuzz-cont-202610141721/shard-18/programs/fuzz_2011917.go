package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint64; F1 int8; F2 uint }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int16; F1 uint; F2 int8 }
type GenS5 struct { F0 int64 }
func genFn0(p0 int32, p1 int16, p2 uint8) uint32 {
    var v6 uint16 = 934
    switch 2147483647 {
    case 1:
        p0 = (p0 | -2147483648)
    case -337:
        p2 = p2
    }
    v7 := []int{-2147483648, 0, 0, 9223372036854775807}
    v6 += 42
    for fi8 := 0; fi8 < 1; fi8++ {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        if (true || true) {
            for fi9 := 0; fi9 < 2; fi9++ {
                if true {
                    switch fi8 {
                    case 954:
                        var v10 uint8 = p2
                        _ = v10
                    }
                    if (p1 <= 42) {
                        v11 := GenS1{F0: 2, F1: 773, F2: 1}
                        if len(v7) > 0 {
                            v7[0] = fi9
                        }
                        _ = v11
                    }
                }
                switch ("hello" + "\n\t") {
                case "foo bar":
                    v12 := GenS1{F0: 1, F1: 255, F2: 0}
                    p1 += int16(p2)
                    _ = v12
                default:
                    v13 := GenS2{F0: 0, F1: -63, F2: 42}
                    _ = v13
                }
                _ = fi9
            }
            var v14 uint64 = 496
            _ = v14
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{("\n\t" + "hello"), "a"}
    var v3 int64 = 1
    v4 := make(map[string]int)
    v5 := make(map[uint32]int)
    var v15 string = ""
    v4["hello"] = 1
    for fi16 := 0; fi16 < 16; fi16++ {
        delete(v5, 151)
        _ = fi16
    }
    if true {
        switch v15 {
        case "":
            v3 = v3
        }
        if (100 < 20) {
            for fi17 := 0; fi17 < 4; fi17++ {
                var v18 uint16 = 1
                if (true && (true && false)) {
                    v15 = (v15 + (v15 + "a"))
                } else {
                    if (v15 != (v15 + "")) {
                        for fi19 := 0; fi19 < 12; fi19++ {
                            v15 = v15
                            _ = fi19
                        }
                    }
                    for fi20 := 0; fi20 < 17; fi20++ {
                        var v21 uint = 18446744073709551615
                        _ = fi20
                        _ = v21
                    }
                    fmt.Println(v15, v15)
                }
                _ = fi17
                _ = v18
            }
            v22 := GenS3{F0: 127, F1: 29}
            var v23 int = 305
            _ = v22
            _ = v23
        }
        v15 = ((v15 + v15) + v15)
    }
    if (!true) {
        if false {
            if (1 > -469) {
                v3 = (v3 >> ((292) & 63))
                if (2 != 0) {
                    var v24 uint = 2
                    if false {
                        _ = len(v2)
                        fmt.Println(v3, v3)
                        v24 += (v24 ^ (v24 % (((v24 + v24)) | 1)))
                    } else {
                        if true {
                            v15 = "a"
                            v3 += v3
                            v25 := GenS4{F0: -890, F1: v24, F2: 42}
                            _ = v25
                        } else {
                            var v26 uint = 2
                            _ = v26
                        }
                        switch v15 {
                        case "foo bar":
                            v27 := GenS5{F0: v3}
                            _ = v27
                        case "\n\t":
                            var v28 uint16 = 1
                            _ = v28
                        case "":
                            v24 = 65535
                        default:
                            fmt.Println(v24, v15, v3)
                        }
                        delete(v5, 2147483647)
                    }
                    switch v15 {
                    case "foo bar":
                        v3 -= 1
                    case "hello":
                        v29 := GenS2{F0: 1, F1: -128, F2: v24}
                        _ = v29
                    default:
                        delete(v1, 179)
                    }
                    _ = v24
                } else {
                    _ = len(v2)
                    v15 = "a"
                }
            } else {
                _ = len(v5)
                var v30 uint8 = 255
                switch "a" {
                case "\n\t":
                    v31 := GenS3{F0: 1, F1: 184}
                    _ = v31
                case "hello":
                    var v32 uint8 = v30
                    _ = v32
                }
                _ = v30
            }
            _ = len(v2)
        }
    }
    var v33 string = ""
    if (255 >= 1) {
        for fi34 := 0; fi34 < 1; fi34++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v33 = v15
            _ = fi34
        }
    }
    v35 := GenS5{F0: (v3 ^ (^(v3)))}
    var v36 uint32 = 42
    var v37 bool = ((v3 - v3) < (v3 & (v3 | (v3 * 705))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
