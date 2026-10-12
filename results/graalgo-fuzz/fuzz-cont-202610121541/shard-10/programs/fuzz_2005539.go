package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 string; F1 int64; F2 uint64 }
type GenS3 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16; F1 uint16 }
func genFn0(p0 int64, p1 uint, p2 uint32) string {
    for fi7 := 0; fi7 < 8; fi7++ {
        fmt.Println(p2, p2, p2)
        if (!(p0 == p0)) {
            v8 := []int{fi7, (fi7 / (((fi7 % ((fi7) | 1))) | 1)), fi7}
            for fi9 := 0; fi9 < 8; fi9++ {
                switch int(p2) {
                case 32767:
                    if true {
                        var v10 int32 = -433
                        _ = len(v8)
                        fmt.Println(p0, fi9, v10)
                        if (p1 == (p1 ^ uint(p2))) {
                            switch -9223372036854775808 {
                            case 32767:
                                var v11 int16 = 1000
                                _ = v11
                            case 4294967295:
                                p2 = p2
                            case 0:
                                switch "foo bar" {
                                case "\n\t":
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                case "a":
                                    fmt.Println(p0, fi9, fi7)
                                case "":
                                    v12 := GenS1{F0: -965, F1: 0}
                                    _ = v12
                                default:
                                    var v13 bool = false
                                    _ = v13
                                }
                            }
                            p1 -= (p1 ^ (p1 & p1))
                        }
                        _ = v10
                    } else {
                        v8 = append(v8, (fi9 - (fi9 % ((fi9) | 1))))
                        var v14 string = "foo bar"
                        _ = v14
                    }
                case -730:
                    var v15 int64 = p0
                    _ = v15
                default:
                    p0 -= p0
                }
                var v16 string = ("hello" + ("a" + "\n\t"))
                _ = fi9
                _ = v16
            }
            for fi17 := 0; fi17 < 20; fi17++ {
                p1 += p1
                fmt.Println(p1, fi7)
                fmt.Println(p1, fi17)
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                switch "a" {
                case "foo bar":
                    if (!false) {
                        _ = len(v8)
                        if (p2 < (^(p2))) {
                            v18 := GenS2{F0: "", F1: p0, F2: 1}
                            _ = v18
                        } else {
                            v19 := GenS2{F0: (("" + "hello") + "a"), F1: p0, F2: 0}
                            fmt.Println(fi17, p0, fi7)
                            if (!false) {
                                var v20 int16 = -32768
                                if true {
                                    p2 += p2
                                } else {
                                    p2 -= 2147483648
                                    p2 = p2
                                }
                                _ = v20
                            } else {
                                var v21 string = "hello"
                                _ = v21
                            }
                            _ = v19
                        }
                    } else {
                        v22 := GenS3{F0: (p1 == 0), F1: (p2 / ((625) | 1)), F2: 18446744073709551615}
                        var v23 bool = true
                        _ = v22
                        _ = v23
                    }
                case "hello":
                    v8 = append(v8, -1)
                case "a":
                    var v24 int16 = 32767
                    _ = v24
                default:
                    var v25 int = fi17
                    _ = v25
                }
                _ = fi17
            }
            _ = v8
        } else {
            var v26 uint = 65535
            var v27 uint8 = 137
            v27 -= 1
            _ = v26
            _ = v27
        }
        switch fi7 {
        case 4294967295:
            if false {
                fmt.Println(fi7, p1)
                for fi28 := 0; fi28 < 16; fi28++ {
                    p2 += (p2 ^ (p2 << ((1) & 63)))
                    _ = fi28
                }
            }
        case 896:
            var v29 bool = ((p1 < p1) && true)
            _ = v29
        }
        _ = fi7
    }
    var v30 uint64 = 715
    v31 := GenS1{F0: 0, F1: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    _ = v31
    return ("hello" + ("foo bar" + "foo bar"))
}

func main() {
    var v1 int = 2147483647
    v2 := []uint32{42, 466}
    v3 := make(map[int]bool)
    var v4 uint32 = 0
    var v5 int32 = 1
    var v6 int8 = -1
    delete(v3, (v1 ^ v1))
    v32 := GenS4{F0: 65535}
    v5 -= v5
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch (("a" + ("" + "a")) + "") {
    case "hello":
        v2 = append(v2, (v4 >> ((42) & 63)))
    case "\n\t":
        fmt.Println(v5)
    case "a":
        v3[(v1 % ((v1) | 1))] = true
    }
    v4 = v4
    for fi33 := 0; fi33 < 9; fi33++ {
        v34 := GenS4{F0: 4294967295}
        switch "hello" {
        case "a":
            fmt.Println(v6, v1)
        case "":
            switch fi33 {
            case 52:
                var v35 uint32 = 42
                _ = v35
            case 0:
                v36 := GenS3{F0: false, F1: (^(v4)), F2: 475}
                _ = v36
            default:
                v6 = 100
            }
        case "hello":
            var v37 int64 = -9223372036854775808
            _ = v37
        }
        _ = fi33
        _ = v34
    }
    if len(v2) > 0 {
        v2[0] = v4
    }
    fmt.Println(v1, v6, v4)
    v38 := GenS5{F0: 42, F1: 65535}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        _ = len(v2)
    } else {
        v5 += (v5 & v5)
        v39 := GenS2{F0: "a", F1: 0, F2: 896}
        delete(v3, v1)
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
