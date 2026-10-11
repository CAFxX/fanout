package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int8 }
type GenS2 struct { F0 bool; F1 uint; F2 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 int; F1 uint8; F2 int16 }
type GenS6 struct { F0 uint32; F1 int }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int32; F1 int16; F2 int16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int32; F1 int8; F2 bool }
func genFn0() uint32 {
    switch -46 {
    case 32768:
        for fi5 := 0; fi5 < 9; fi5++ {
            var v6 uint64 = 163
            _ = fi5
            _ = v6
        }
    case 128:
        var v7 float32 = 0.0
        _ = v7
    case 845:
        var v8 uint64 = 0
        _ = v8
    }
    v9 := GenS1{F0: 128, F1: (!(!(true || false))), F2: 0}
    _ = v9
    return 0
}
func genFn1(p0 uint64) uint16 {
    if (!(true && false)) {
        var v10 uint16 = 1
        _ = v10
    }
    v11 := GenS1{F0: -937, F1: false, F2: 127}
    _ = p0
    _ = v11
    return 2
}

func main() {
    var v1 bool = true
    var v2 uint64 = 0
    var v3 int = 32768
    var v4 int = (v3 & (v3 | (-(v3))))
    fmt.Println(v3, v2)
    if v1 {
        v12 := []int{(v3 ^ (-(v4)))}
        v13 := GenS2{F0: (("a" + "") < ("hello" + "hello")), F1: 301, F2: ""}
        if (!true) {
            var v14 uint = 606
            if false {
                for fi15 := 0; fi15 < 12; fi15++ {
                    if true {
                        v3 = -142
                        var v16 uint64 = (v2 ^ (^(v2)))
                        _ = v16
                    } else {
                        fmt.Println(v14, fi15)
                    }
                    if ((42 >= 42) || v1) {
                        if (130 != genFn0()) {
                            if len(v12) > 0 {
                                v12[0] = v4
                            }
                            var v17 float32 = (1.5 + 3.14159)
                            _ = v17
                        } else {
                            if v1 {
                                fmt.Println(v1, v2)
                                for fi18 := 0; fi18 < 12; fi18++ {
                                    v19 := GenS3{F0: 580}
                                    v4 -= (fi18 * (v4 ^ (fi18 | (v3 >> (((v14 / (((v14 % ((v14) | 1))) | 1))) & 63)))))
                                    v2 += uint64(v14)
                                    _ = fi18
                                    _ = v19
                                }
                                var v20 uint8 = 128
                                _ = v20
                            } else {
                                v21 := GenS4{F0: 2147483648, F1: 32767}
                                if (-128 > 1) {
                                    if v1 {
                                        switch v3 {
                                        case 1:
                                            v22 := GenS4{F0: 0, F1: -32768}
                                            _ = v22
                                        case -943:
                                            for fi23 := 0; fi23 < 14; fi23++ {
                                                var v24 int8 = 42
                                                v25 := GenS5{F0: -1, F1: 0, F2: -32768}
                                                _ = fi23
                                                _ = v24
                                                _ = v25
                                            }
                                        }
                                        var v26 uint64 = (v2 | (v2 << (((v14 >> ((65535) & 63))) & 63)))
                                        _ = v26
                                    }
                                    var v27 uint16 = 1
                                    v27 += v27
                                    _ = v27
                                } else {
                                    v14 -= v14
                                }
                                fmt.Println(v14)
                                _ = v21
                            }
                        }
                        if (2147483648 > 0) {
                            for fi28 := 0; fi28 < 5; fi28++ {
                                switch ("hello" + "\n\t") {
                                case "hello":
                                    v29 := GenS2{F0: ("a" < "\n\t"), F1: 65535, F2: ("hello" + "hello")}
                                    _ = v29
                                default:
                                    fmt.Println(v2, v14, v3)
                                }
                                var v30 uint32 = 1
                                v31 := GenS6{F0: (v30 << (((v14 * (^(v14)))) & 63)), F1: (v3 & int(v2))}
                                _ = fi28
                                _ = v30
                                _ = v31
                            }
                            _ = len(v12)
                        }
                        var v32 uint = (v14 - v14)
                        _ = v32
                    }
                    switch "hello" {
                    case "a":
                        v33 := GenS4{F0: 731, F1: 32767}
                        _ = v33
                    case "\n\t":
                        switch 128 {
                        case 42:
                            for fi34 := 0; fi34 < 10; fi34++ {
                                switch ("a" + "foo bar") {
                                case "\n\t":
                                    v1 = false
                                case "a":
                                    if (!v1) {
                                        if (127 != int8(v2)) {
                                            fmt.Println(fi34, v1)
                                        }
                                        v35 := GenS7{F0: 1}
                                        _ = v35
                                    } else {
                                        var v36 uint64 = 93
                                        _ = v36
                                    }
                                }
                                v37 := GenS5{F0: (fi15 / (((v4 >> ((v14) & 63))) | 1)), F1: 255, F2: 32767}
                                var v38 uint16 = 32767
                                _ = fi34
                                _ = v37
                                _ = v38
                            }
                        }
                    }
                    _ = fi15
                }
                v39 := GenS8{F0: 2147483647, F1: 381, F2: 32767}
                _ = v39
            }
            for fi40 := 0; fi40 < 1; fi40++ {
                var v41 int64 = 42
                _ = fi40
                _ = v41
            }
            _ = v14
        }
        _ = v12
        _ = v13
    }
    v1 = true
    if v1 {
        v3 -= 32767
    }
    v42 := GenS9{F0: ((v1 && (2147483648 == 65535)) && (false && v1))}
    switch v4 {
    case 270:
        var v43 bool = v1
        _ = v43
    case 4294967295:
        var v44 uint64 = 148
        _ = v44
    case -937:
        v3 = (v4 + (v4 * 825))
    default:
        v2 = (v2 / (((v2 / ((v2) | 1))) | 1))
    }
    fmt.Println(v2, v1)
    var v45 uint8 = 186
    v46 := []int{-1, (^(v3))}
    v47 := GenS10{F0: -1, F1: 0, F2: v1}
    v46 = append(v46, (v4 << ((4294967295) & 63)))
    fmt.Println(v4, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v46:", len(v46), v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v45
    _ = v46
    _ = v47
}
