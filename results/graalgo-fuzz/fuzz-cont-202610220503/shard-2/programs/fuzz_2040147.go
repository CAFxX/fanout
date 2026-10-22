package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int8; F2 int32 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 string; F1 int }
type GenS6 struct { F0 bool; F1 uint8; F2 int8 }
func genFn0(p0 int64, p1 uint16) int64 {
    var v5 int8 = 39
    v6 := []int{32768, 42, 127}
    switch -1 {
    case 42:
        p1 -= 255
    case 32767:
        var v7 int64 = int64(p1)
        _ = v7
    case 32768:
        fmt.Println(p0, v5, v5)
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS1{F0: 1}
        var v10 uint64 = 42
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (p0 << ((2) & 63))
}
func genFn1(p0 uint16, p1 bool, p2 string) string {
    v11 := GenS2{F0: 2147483647, F1: 42, F2: -892}
    p2 = "a"
    var v12 string = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return p2
}

func main() {
    var v1 int16 = 32767
    var v2 uint32 = 2147483647
    var v3 uint64 = 658
    v4 := make(map[int]bool)
    v2 = (^(v2))
    var v13 uint32 = (^(v2))
    v14 := GenS2{F0: 32767, F1: 0, F2: -647}
    v15 := []int{4294967295}
    if (!false) {
        var v16 uint16 = 127
        v4[-2147483648] = (0 > 18446744073709551615)
        if true {
            if true {
                for fi17 := 0; fi17 < 18; fi17++ {
                    v15 = append(v15, fi17)
                    _ = fi17
                }
                fmt.Println(v3, v1, v1)
                v2 -= v2
            } else {
                if true {
                    var v18 float32 = 3.14159
                    _ = v18
                }
                v19 := GenS1{F0: 2}
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                _ = v19
            }
            fmt.Println(v16, v13)
            var v20 int64 = -1
            _ = v20
        } else {
            v21 := GenS2{F0: -744, F1: 100, F2: 42}
            v22 := v4[127]
            _ = v21
            _ = v22
        }
        _ = v16
    }
    v2 -= (v2 % ((v2) | 1))
    var v23 float64 = 3.14159
    v24 := GenS1{F0: (v13 - (v13 / (((v13 << ((2) & 63))) | 1)))}
    v25 := GenS1{F0: (v13 << ((42) & 63))}
    v26 := GenS1{F0: (v13 + v2)}
    _ = len(v15)
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v2, v23)
    default:
        var v27 uint32 = (^(v2))
        _ = v27
    }
    v13 = 2147483648
    if (!true) {
        v3 -= uint64(v13)
        v2 = (^(v13))
    }
    for fi28 := 0; fi28 < 15; fi28++ {
        var v29 uint32 = 0
        if true {
            var v30 int32 = -2147483648
            if false {
                var v31 uint16 = 29
                switch "a" {
                case "\n\t":
                    switch "\n\t" {
                    case "\n\t":
                        for fi32 := 0; fi32 < 5; fi32++ {
                            v2 = v29
                            _ = fi32
                        }
                    case "a":
                        var v33 int = fi28
                        _ = v33
                    case "":
                        for fi34 := 0; fi34 < 7; fi34++ {
                            fmt.Println(fi34, v29, v3)
                            var v35 string = "\n\t"
                            switch "foo bar" {
                            case "foo bar":
                                switch v35 {
                                case "\n\t":
                                    v2 = (v13 << ((0) & 63))
                                }
                            default:
                                fmt.Println(v23, v30)
                            }
                            _ = fi34
                            _ = v35
                        }
                    default:
                        var v36 int16 = 0
                        _ = v36
                    }
                case "a":
                    if false {
                        v3 = (v3 - 42)
                    }
                default:
                    switch "hello" {
                    case "a":
                        _ = len(v4)
                    case "hello":
                        v37 := GenS2{F0: fi28, F1: -128, F2: (v30 % ((v30) | 1))}
                        _ = v37
                    default:
                        _ = len(v15)
                    }
                }
                _ = v31
            } else {
                v29 -= (v13 % ((2147483647) | 1))
            }
            _ = v30
        } else {
            v3 = 42
            for fi38 := 0; fi38 < 7; fi38++ {
                if (!(-128 < 42)) {
                    v13 = v2
                    v39 := GenS3{F0: 345, F1: 127}
                    _ = v39
                }
                v40 := GenS4{F0: v2, F1: 2147483647}
                if false {
                    v41 := GenS5{F0: "hello", F1: (fi28 ^ int(v3))}
                    if (!(uint16(v13) == 65535)) {
                        var v42 uint = 4294967295
                        v43 := GenS4{F0: 2, F1: (^(v29))}
                        v44 := GenS3{F0: (v42 + (v42 ^ v42)), F1: (fi28 - fi38)}
                        _ = v42
                        _ = v43
                        _ = v44
                    } else {
                        v45 := GenS5{F0: ("" + "\n\t"), F1: fi28}
                        var v46 uint64 = v3
                        fmt.Println(v46, fi38)
                        _ = v45
                        _ = v46
                    }
                    _ = v41
                } else {
                    var v47 float64 = 1e10
                    _ = v47
                }
                _ = fi38
                _ = v40
            }
            v2 = (v29 & (^(v29)))
        }
        _ = fi28
        _ = v29
    }
    delete(v4, 591)
    var v48 uint8 = 128
    v49 := GenS6{F0: (false || false), F1: 1, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v48
    _ = v49
}
