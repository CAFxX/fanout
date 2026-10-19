package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 int }
type GenS5 struct { F0 string; F1 int; F2 bool }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint; F1 string; F2 int8 }
type GenS8 struct { F0 int32; F1 uint }
func genFn0(p0 uint, p1 string, p2 int) int64 {
    p0 -= p0
    v6 := GenS1{F0: p1, F1: "\n\t", F2: 2}
    switch "foo bar" {
    case "":
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 string = p1
            v9 := []int{}
            switch -1 {
            case 42:
                fmt.Println(p2, p0, p1)
            case 128:
                v10 := GenS2{F0: 38}
                _ = v10
            default:
                if false {
                    var v11 int8 = -1
                    _ = v11
                } else {
                    switch (("foo bar" + (v8 + "hello")) + v8) {
                    case "a":
                        _ = len(v9)
                    }
                }
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    case "hello":
        v12 := []int{p2, 9223372036854775807, (p2 + -601), p2}
        _ = v12
    default:
        p0 += p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 180
}
func genFn1() uint32 {
    var v13 uint64 = 18446744073709551615
    fmt.Println(v13)
    v13 += v13
    _ = v13
    return 236
}

func main() {
    v1 := []uint32{4294967295}
    var v2 int32 = 2147483647
    var v3 uint = 42
    var v4 uint64 = 601
    v5 := make(map[string]int)
    v14 := GenS1{F0: "", F1: ("hello" + "\n\t"), F2: (v4 & (v4 | 18446744073709551615))}
    var v15 uint16 = 414
    var v16 uint64 = 0
    switch -24 {
    case 42:
        for fi17 := 0; fi17 < 16; fi17++ {
            v18 := GenS2{F0: (v4 * (v16 | (^(v16))))}
            v19 := GenS2{F0: v16}
            _ = fi17
            _ = v18
            _ = v19
        }
    default:
        v20 := GenS1{F0: "", F1: (("foo bar" + "foo bar") + ""), F2: (v4 * 18446744073709551615)}
        _ = v20
    }
    if true {
        var v21 uint64 = v4
        v22 := v5["foo bar"]
        for fi23 := 0; fi23 < 19; fi23++ {
            fmt.Println(v21, v2)
            _ = fi23
        }
        _ = v21
        _ = v22
    } else {
        var v24 uint16 = uint16(v16)
        _ = v24
    }
    v25 := GenS3{F0: v15}
    v4 -= v4
    for fi26 := 0; fi26 < 9; fi26++ {
        fmt.Println(v15, v2, v3)
        if true {
            if (v4 >= (v16 << (((v3 + uint(v15))) & 63))) {
                var v27 uint = (v3 % ((0) | 1))
                v4 = 18446744073709551615
                for fi28 := 0; fi28 < 4; fi28++ {
                    if false {
                        _ = len(v5)
                        if (-9223372036854775808 != genFn0(v3, "a", (fi26 ^ 9223372036854775807))) {
                            v29 := GenS3{F0: uint16(v4)}
                            var v30 int8 = int8(v27)
                            _ = v29
                            _ = v30
                        }
                        var v31 int8 = -83
                        _ = v31
                    } else {
                        v32 := GenS4{F0: 238, F1: -1}
                        if len(v1) > 0 {
                            v1[0] = 0
                        }
                        v33 := GenS5{F0: "a", F1: -6, F2: (true && true)}
                        if false {
                            if (v4 != v16) {
                                v34 := GenS4{F0: 4294967295, F1: fi26}
                                v35 := GenS2{F0: (v16 >> (((^(v27))) & 63))}
                                _ = v34
                                _ = v35
                            } else {
                                v36 := GenS6{F0: (v2 & (^(v2)))}
                                v16 += v4
                                switch ("hello" + ("" + ("\n\t" + ""))) {
                                case "foo bar":
                                    v37 := GenS7{F0: (v27 & 1), F1: "hello", F2: int8(v2)}
                                    _ = v37
                                case "a":
                                    var v38 uint8 = 127
                                    if true {
                                        v5[""] = fi28
                                        var v39 string = ("hello" + "a")
                                        v39 = ((("a" + v39) + "foo bar") + (v39 + v39))
                                        _ = v39
                                    }
                                    fmt.Println(fi28)
                                    _ = v38
                                }
                                v3 = 549
                                _ = v36
                            }
                            v15 = v15
                            fmt.Println(fi26)
                            if true {
                                if len(v1) > 0 {
                                    v1[0] = 2147483647
                                }
                                v4 = (v4 - 641)
                                switch (fi26 >> ((v3) & 63)) {
                                case 42:
                                    v40 := GenS3{F0: (v15 << ((v27) & 63))}
                                    if false {
                                        var v41 uint64 = (v4 & (v16 ^ v16))
                                        v42 := GenS7{F0: v27, F1: "foo bar", F2: 1}
                                        _ = v41
                                        _ = v42
                                    } else {
                                        fmt.Println(v2, v27, v16)
                                        var v43 float32 = (((0.1 - (0.1 - (0.1 * (0.1 * 1e10)))) / 1.0) * 0.0)
                                        _ = v43
                                    }
                                    if ((^(fi28)) >= fi28) {
                                        if false {
                                            var v44 bool = (!(float32(0.1) >= ((1e10 / 0.1) - 1.5)))
                                            _ = v44
                                        } else {
                                            v15 = uint16(v3)
                                            v3 = v3
                                            _ = len(v1)
                                            switch fi28 {
                                            case 980:
                                                var v45 int16 = 1000
                                                _ = v45
                                            case 4294967295:
                                                var v46 int = fi26
                                                _ = v46
                                            case 2147483647:
                                                v3 = v3
                                            default:
                                                switch -835 {
                                                case 771:
                                                    v47 := GenS7{F0: (v27 % (((v3 * (v27 / (((^(v27))) | 1)))) | 1)), F1: "", F2: 42}
                                                    _ = v47
                                                case 42:
                                                    if false {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        v2 += v2
                                                        var v48 int8 = int8(fi26)
                                                        _ = v48
                                                    }
                                                }
                                            }
                                        }
                                    } else {
                                        var v49 int8 = 1
                                        var v50 float32 = (1e10 + 0.0)
                                        switch "" {
                                        case "a":
                                            var v51 float32 = v50
                                            _ = v51
                                        }
                                        _ = v49
                                        _ = v50
                                    }
                                    fmt.Println(v4, v2, v16)
                                    v2 += (v2 % ((v2) | 1))
                                    v52 := GenS2{F0: (v4 ^ (v16 & (v16 + (v16 ^ v4))))}
                                    v53 := GenS8{F0: v2, F1: 4294967295}
                                    _ = v40
                                    _ = v52
                                    _ = v53
                                }
                                switch "a" {
                                case "foo bar":
                                    v1 = append(v1, 729)
                                }
                                v15 -= uint16(v4)
                            }
                        } else {
                            v16 += 1
                        }
                        v3 = (v3 * (v27 - v3))
                        _ = v32
                        _ = v33
                    }
                    _ = fi28
                }
                _ = v27
            }
        }
        var v54 int64 = genFn0(1, "a", 42)
        _ = fi26
        _ = v54
    }
    switch "foo bar" {
    case "":
        fmt.Println(v4, v3, v16)
    case "foo bar":
        var v55 uint64 = v4
        _ = v55
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
