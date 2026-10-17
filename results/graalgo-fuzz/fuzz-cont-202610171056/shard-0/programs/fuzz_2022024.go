package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint8; F1 int16; F2 bool }
type GenS4 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS5 struct { F0 int32; F1 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32; F1 int32 }
type GenS8 struct { F0 int32; F1 uint8; F2 uint16 }
func genFn0(p0 uint16) uint {
    switch (("a" + ("a" + "a")) + "foo bar") {
    case "foo bar":
        var v4 int64 = 9223372036854775807
        _ = v4
    case "a":
        if false {
            p0 -= p0
            v5 := GenS1{F0: (!(true || false))}
            var v6 int16 = 0
            _ = v5
            _ = v6
        }
    }
    switch 32768 {
    case -2147483648:
        v7 := GenS2{F0: true, F1: 255}
        _ = v7
    case -1:
        v8 := []int{-137, 9223372036854775807, 128, 32767}
        _ = v8
    default:
        var v9 int64 = -1
        _ = v9
    }
    _ = p0
    return 1
}
func genFn1(p0 uint32, p1 uint16) uint {
    var v10 bool = false
    p1 = (p1 >> ((820) & 63))
    if true {
        var v11 int16 = 1000
        _ = v11
    } else {
        if v10 {
            switch 127 {
            case 32767:
                switch 32767 {
                case -9223372036854775808:
                    switch "hello" {
                    case "foo bar":
                        var v12 int32 = 0
                        _ = v12
                    case "hello":
                        var v13 bool = ((false || (float64(-2.25) <= 0.1)) || (!(false || (!v10))))
                        _ = v13
                    }
                case 0:
                    v14 := []int{32767, 128, 4294967295, -533}
                    _ = v14
                }
            case 127:
                v15 := []int{2147483647, 32768, -2147483648, -1}
                _ = v15
            case 1:
                switch "hello" {
                case "\n\t":
                    p0 -= p0
                case "hello":
                    var v16 uint16 = (p1 & (^(p1)))
                    _ = v16
                case "a":
                    v17 := GenS3{F0: 42, F1: -963, F2: (659 < (p0 | (p0 + 361)))}
                    _ = v17
                default:
                    fmt.Println(p1)
                }
            default:
                p0 += p0
            }
        }
        v18 := GenS4{F0: p0, F1: 65535, F2: 1}
        if ((^(p0)) < (p0 / (((p0 | 2147483647)) | 1))) {
            p0 += 81
        }
        _ = v18
    }
    p1 -= 1
    _ = p0
    _ = p1
    _ = v10
    return 42
}

func main() {
    var v1 uint8 = 128
    var v2 float64 = 1.5
    var v3 float32 = -2.25
    for fi19 := 0; fi19 < 16; fi19++ {
        if false {
            v20 := []int{(fi19 << ((uint(v1)) & 63)), -323, fi19, (fi19 | (-(fi19)))}
            v1 -= (^(v1))
            v1 += 127
            _ = v20
        } else {
            var v21 int = (^(fi19))
            for fi22 := 0; fi22 < 15; fi22++ {
                if true {
                    v23 := []int{}
                    _ = v23
                }
                var v24 uint = 0
                if true {
                    if (0 <= 1) {
                        var v25 float64 = (1.5 * (0.0 * 1e10))
                        switch ("foo bar" + "\n\t") {
                        case "":
                            var v26 uint64 = uint64(v1)
                            _ = v26
                        case "foo bar":
                            v2 += float64(v3)
                        default:
                            fmt.Println(v1)
                        }
                        _ = v25
                    } else {
                        v27 := GenS3{F0: 207, F1: 1, F2: true}
                        v28 := []int{v21, 2147483647, fi19, fi22}
                        _ = v27
                        _ = v28
                    }
                    fmt.Println(fi19, v21, v3)
                    v29 := GenS1{F0: false}
                    _ = v29
                }
                if false {
                    var v30 string = ""
                    _ = v30
                } else {
                    if false {
                        if (!(v24 != 255)) {
                            fmt.Println(v3)
                            v31 := []int{(^(fi22)), fi22}
                            var v32 uint64 = 349
                            _ = v31
                            _ = v32
                        } else {
                            if ((v3 >= v3) && (true || false)) {
                                v33 := GenS5{F0: 267, F1: v24}
                                v34 := GenS6{F0: 1}
                                v24 += (v24 + 65535)
                                _ = v33
                                _ = v34
                            } else {
                                var v35 int64 = 0
                                _ = v35
                            }
                            var v36 int64 = int64(fi19)
                            var v37 float32 = (v3 + (v3 - v3))
                            _ = v36
                            _ = v37
                        }
                        fmt.Println(v3)
                        if true {
                            fmt.Println(fi19)
                            v3 = v3
                        } else {
                            v38 := GenS4{F0: 42, F1: 4294967295, F2: uint8(v24)}
                            v39 := []int{fi22}
                            var v40 uint8 = (^(v1))
                            v21 -= 0
                            _ = v38
                            _ = v39
                            _ = v40
                        }
                    }
                }
                var v41 float64 = v2
                v42 := GenS7{F0: 902, F1: int32(v21)}
                _ = fi22
                _ = v24
                _ = v41
                _ = v42
            }
            _ = v21
        }
        _ = fi19
    }
    v2 = v2
    var v43 float32 = (((v3 / 3.14159) - v3) - v3)
    fmt.Println(v43)
    v44 := []int{}
    for fi45 := 0; fi45 < 19; fi45++ {
        v43 = (float32(v2) * float32((v2 + v2)))
        _ = fi45
    }
    v3 += v43
    for fi46 := 0; fi46 < 1; fi46++ {
        v47 := GenS8{F0: 42, F1: (^(v1)), F2: 665}
        var v48 int8 = -128
        fmt.Println(v1, v43)
        _ = fi46
        _ = v47
        _ = v48
    }
    for fi49 := 0; fi49 < 7; fi49++ {
        v43 = 1e10
        _ = fi49
    }
    switch "hello" {
    case "hello":
        v2 += ((((v2 + -2.25) - v2) - v2) * (0.1 + (v2 / (v2 - ((v2 - float64(v43)) * (v2 / v2))))))
    }
    if len(v44) >= 2 {
        v44 = v44[1:len(v44)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v44:", len(v44), v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v44
}
