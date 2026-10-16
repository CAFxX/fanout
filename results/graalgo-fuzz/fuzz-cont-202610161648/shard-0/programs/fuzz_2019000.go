package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int8; F1 uint64; F2 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint; F1 uint64 }
type GenS6 struct { F0 string; F1 int32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int; F1 uint32; F2 string }
func genFn0(p0 int8, p1 int32, p2 int64) bool {
    var v5 uint64 = 0
    if (v5 == 18446744073709551615) {
        if true {
            var v6 uint32 = 0
            var v7 uint8 = 2
            if ((!true) && (520 > uint(v6))) {
                v8 := []int{127, 1, 691}
                var v9 string = "\n\t"
                if (!false) {
                    if true {
                        if false {
                            v10 := GenS1{F0: v5}
                            v11 := GenS1{F0: v5}
                            _ = v10
                            _ = v11
                        } else {
                            for fi12 := 0; fi12 < 13; fi12++ {
                                v13 := GenS2{F0: (fi12 % ((128) | 1)), F1: 65535}
                                _ = fi12
                                _ = v13
                            }
                            for fi14 := 0; fi14 < 8; fi14++ {
                                v15 := GenS3{F0: (^(p0)), F1: 42, F2: true}
                                _ = fi14
                                _ = v15
                            }
                        }
                    }
                    p0 -= p0
                }
                _ = v8
                _ = v9
            } else {
                var v16 int64 = 242
                switch 32768 {
                case -428:
                    switch ("foo bar" + "\n\t") {
                    case "a":
                        v17 := []int{-903, -932, 834}
                        _ = v17
                    case "foo bar":
                        var v18 int16 = 1
                        _ = v18
                    case "":
                        v19 := GenS4{F0: 2}
                        _ = v19
                    default:
                        var v20 float32 = (0.1 / float32(0.1))
                        _ = v20
                    }
                case 554:
                    for fi21 := 0; fi21 < 4; fi21++ {
                        v22 := GenS4{F0: 65535}
                        _ = fi21
                        _ = v22
                    }
                default:
                    if false {
                        if (p2 < (p2 + v16)) {
                            var v23 int32 = 0
                            _ = v23
                        } else {
                            for fi24 := 0; fi24 < 14; fi24++ {
                                var v25 uint = 618
                                switch "hello" {
                                case "\n\t":
                                    v25 -= v25
                                case "":
                                    v5 = 2
                                }
                                _ = fi24
                                _ = v25
                            }
                        }
                        v26 := GenS1{F0: v5}
                        _ = v26
                    } else {
                        for fi27 := 0; fi27 < 20; fi27++ {
                            var v28 bool = ((((uint16(v5) < 991) || ((fi27 >> ((0) & 63)) != -743)) && true) && (!(903 < 318)))
                            _ = fi27
                            _ = v28
                        }
                        v29 := GenS2{F0: 127, F1: uint16(v6)}
                        fmt.Println(p0)
                        _ = v29
                    }
                }
                var v30 uint16 = 0
                _ = v16
                _ = v30
            }
            _ = v6
            _ = v7
        }
        switch -2147483648 {
        case 2147483647:
            for fi31 := 0; fi31 < 18; fi31++ {
                v5 = (v5 - 18446744073709551615)
                if ((^(p1)) != p1) {
                    if true {
                        v5 = (^(v5))
                        v32 := GenS2{F0: fi31, F1: 1}
                        p1 += p1
                        _ = v32
                    } else {
                        p2 = 9223372036854775807
                        switch ("a" + ("foo bar" + ("\n\t" + "foo bar"))) {
                        case "a":
                            for fi33 := 0; fi33 < 4; fi33++ {
                                p1 -= -1
                                _ = fi33
                            }
                        case "\n\t":
                            v5 -= (v5 | (v5 / (((^(v5))) | 1)))
                        default:
                            for fi34 := 0; fi34 < 5; fi34++ {
                                p1 -= p1
                                var v35 uint64 = (v5 % ((42) | 1))
                                _ = fi34
                                _ = v35
                            }
                        }
                        var v36 uint64 = v5
                        _ = v36
                    }
                    p1 = p1
                    v37 := []int{fi31}
                    _ = v37
                } else {
                    v38 := GenS4{F0: 42}
                    if (false && false) {
                        fmt.Println(p1, p0, p0)
                        v39 := []int{fi31, fi31, fi31}
                        _ = v39
                    }
                    _ = v38
                }
                switch ("a" + (("foo bar" + "hello") + ("foo bar" + "\n\t"))) {
                case "hello":
                    if true {
                        var v40 int32 = p1
                        v41 := []int{(^(fi31)), fi31, 0}
                        _ = v40
                        _ = v41
                    } else {
                        fmt.Println(fi31)
                        for fi42 := 0; fi42 < 8; fi42++ {
                            p2 = (p2 | -9223372036854775808)
                            fmt.Println(fi31, p0)
                            v5 = 18446744073709551615
                            _ = fi42
                        }
                    }
                case "foo bar":
                    for fi43 := 0; fi43 < 8; fi43++ {
                        var v44 int8 = (p0 >> ((18446744073709551615) & 63))
                        v45 := GenS4{F0: 340}
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                default:
                    p2 += 1
                }
                _ = fi31
            }
        case 127:
            fmt.Println(v5, p2)
        }
    }
    var v46 string = ("\n\t" + (("a" + "\n\t") + "\n\t"))
    v47 := GenS5{F0: 65535, F1: (v5 << ((2) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v46
    _ = v47
    _ = v5
    return (156 > 42)
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{0, 1}
    var v3 uint32 = 477
    var v4 int8 = -1
    if (!false) {
        v48 := GenS6{F0: "hello", F1: 42}
        _ = v48
    } else {
        var v49 int16 = -979
        switch -1 {
        case 4294967295:
            v4 -= (v4 & v4)
        case -541:
            v2 = append(v2, v3)
        }
        _ = v49
    }
    if true {
        v50 := GenS6{F0: "foo bar", F1: -582}
        _ = v50
    }
    if true {
        var v51 int16 = -32768
        v52 := GenS4{F0: 32767}
        _ = v51
        _ = v52
    } else {
        v53 := GenS7{F0: 2}
        switch "foo bar" {
        case "":
            v54 := GenS5{F0: 4294967295, F1: 988}
            _ = v54
        case "foo bar":
            var v55 uint = 0
            _ = v55
        }
        v3 = 2147483648
        _ = v53
    }
    var v56 bool = (true || false)
    var v57 uint = 1
    var v58 float64 = 0.1
    for fi59 := 0; fi59 < 8; fi59++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v60 int64 = 9223372036854775807
        v61 := GenS8{F0: 2147483647, F1: v3, F2: ""}
        _ = fi59
        _ = v60
        _ = v61
    }
    var v62 string = "\n\t"
    if len(v2) > 0 {
        v2[0] = (^(v3))
    }
    for fi63 := 0; fi63 < 16; fi63++ {
        fmt.Println(fi63, v58)
        _ = fi63
    }
    v1[42] = (341 <= 65535)
    v56 = v56
    fmt.Println(v58)
    fmt.Println(v57, v4)
    for fi64 := 0; fi64 < 20; fi64++ {
        if v56 {
            v65 := GenS5{F0: 4294967295, F1: 1}
            _ = v65
        }
        for fi66 := 0; fi66 < 13; fi66++ {
            v3 = v3
            v4 -= v4
            _ = fi66
        }
        _ = fi64
    }
    v57 = v57
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v56
    _ = v57
    _ = v58
    _ = v62
}
