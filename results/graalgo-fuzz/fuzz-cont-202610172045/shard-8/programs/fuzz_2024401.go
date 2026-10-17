package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint32; F2 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 string; F2 int16 }
type GenS5 struct { F0 string; F1 int8; F2 int }
type GenS6 struct { F0 uint64; F1 uint; F2 int16 }
type GenS7 struct { F0 int16; F1 string; F2 int64 }
type GenS8 struct { F0 int16; F1 uint }

func main() {
    var v1 bool = (!(0 >= 1))
    v2 := make(map[int]bool)
    var v3 uint32 = 65535
    v4 := []int{128, -219, 2147483647, 2147483647}
    v3 += (v3 >> ((148) & 63))
    v1 = true
    v5 := GenS1{F0: (v3 / ((1) | 1))}
    for fi6 := 0; fi6 < 9; fi6++ {
        _ = len(v2)
        v1 = (((((((1e10 - 1.5) * 1.5) > -2.25) || true) || (v1 && v1)) || (42 == 0)) && (42 == -1))
        if true {
            if true {
                for fi7 := 0; fi7 < 12; fi7++ {
                    var v8 int8 = 101
                    switch "a" {
                    case "":
                        fmt.Println(v1)
                    default:
                        _ = len(v2)
                    }
                    fmt.Println(v3)
                    _ = fi7
                    _ = v8
                }
            } else {
                var v9 int = (fi6 * 1)
                _ = v9
            }
            fmt.Println(fi6, v3)
            fmt.Println(fi6, fi6)
        } else {
            if v1 {
                v1 = (42 <= 194)
            } else {
                for fi10 := 0; fi10 < 4; fi10++ {
                    var v11 int16 = 32767
                    if (v1 || (1 <= -40)) {
                        fmt.Println(v1, fi10, fi6)
                    }
                    v12 := GenS1{F0: 201}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                if true {
                    fmt.Println(v1, v1, fi6)
                    switch 32768 {
                    case 128:
                        v13 := GenS2{F0: 0, F1: (v3 * v3), F2: 1}
                        _ = v13
                    case 42:
                        var v14 uint32 = (v3 + v3)
                        _ = v14
                    }
                    v3 -= v3
                } else {
                    var v15 int16 = 32767
                    v16 := GenS3{F0: 1}
                    for fi17 := 0; fi17 < 2; fi17++ {
                        v2[(^(fi6))] = v1
                        var v18 uint = uint(v3)
                        v19 := GenS4{F0: (v18 << (((^(v18))) & 63)), F1: ("foo bar" + ""), F2: (v15 + v15)}
                        if v1 {
                            if v1 {
                                v20 := GenS5{F0: "a", F1: 1, F2: fi6}
                                var v21 int = (-(fi17))
                                _ = v20
                                _ = v21
                            } else {
                                v22 := GenS6{F0: 1, F1: 65535, F2: -32768}
                                _ = v22
                            }
                        } else {
                            var v23 uint16 = 2
                            _ = v23
                        }
                        var v24 float64 = 1.5
                        _ = fi17
                        _ = v18
                        _ = v19
                        _ = v24
                    }
                    _ = v15
                    _ = v16
                }
                if (92 <= 1) {
                    v25 := GenS7{F0: -1, F1: ("foo bar" + "hello"), F2: -129}
                    for fi26 := 0; fi26 < 8; fi26++ {
                        var v27 uint16 = 0
                        _ = fi26
                        _ = v27
                    }
                    _ = v25
                }
            }
            var v28 string = "hello"
            var v29 int = (fi6 | 42)
            _ = v28
            _ = v29
        }
        _ = fi6
    }
    fmt.Println(v3)
    var v30 int32 = -2147483648
    var v31 int64 = 42
    var v32 float64 = (-2.25 - 1.5)
    switch 0 {
    case 1:
        v33 := GenS1{F0: (v3 - (v3 << ((20) & 63)))}
        _ = v33
    case 128:
        if true {
            _ = len(v2)
        } else {
            var v34 int = 127
            if (!v1) {
                v35 := GenS8{F0: -32768, F1: 786}
                _ = v35
            } else {
                for fi36 := 0; fi36 < 3; fi36++ {
                    v2[(v34 - v34)] = (991 <= 10)
                    v2[v34] = ((v30 >> ((4294967295) & 63)) == (v30 * v30))
                    _ = fi36
                }
                delete(v2, 1)
                for fi37 := 0; fi37 < 12; fi37++ {
                    for fi38 := 0; fi38 < 1; fi38++ {
                        v31 += 9223372036854775807
                        v34 = -9223372036854775808
                        _ = fi38
                    }
                    v30 += v30
                    var v39 string = "hello"
                    _ = fi37
                    _ = v39
                }
            }
            v30 = v30
            _ = v34
        }
    case 32767:
        var v40 uint8 = 209
        _ = v40
    default:
        var v41 float32 = (((((float32((((v32 - v32) / v32) - 1.5)) / 1.0) / -2.25) - 1.5) * float32(v32)) * 3.14159)
        _ = v41
    }
    v3 -= (^(v3))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
