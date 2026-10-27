package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 string; F2 uint16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8; F1 string; F2 int8 }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 int32; F1 uint; F2 bool }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int; F1 uint16 }
func genFn0(p0 uint16, p1 int16, p2 int32) bool {
    for fi4 := 0; fi4 < 9; fi4++ {
        p1 = (-(p1))
        if false {
            var v5 int = (fi4 | (fi4 ^ fi4))
            _ = v5
        } else {
            switch "" {
            case "a":
                v6 := GenS1{F0: (("a" + ("foo bar" + "hello")) + "hello"), F1: "hello"}
                _ = v6
            case "hello":
                v7 := GenS1{F0: "\n\t", F1: "\n\t"}
                _ = v7
            case "":
                var v8 int16 = 42
                _ = v8
            }
            var v9 uint16 = 2
            _ = v9
        }
        _ = fi4
    }
    v10 := []int{}
    v10 = append(v10, 9223372036854775807)
    var v11 float32 = (float32(0.1) / 1e10)
    v12 := GenS1{F0: "hello", F1: ""}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return (false || (!(255 == 255)))
}
func genFn1(p0 uint16) int32 {
    p0 += (^(p0))
    switch "" {
    case "\n\t":
        var v13 int8 = -94
        _ = v13
    case "a":
        v14 := []int{137, 4294967295, -834, 32768}
        _ = v14
    }
    v15 := GenS1{F0: "hello", F1: ("foo bar" + "foo bar")}
    var v16 uint = 2
    v17 := []int{-359}
    _ = p0
    _ = v15
    _ = v16
    _ = v17
    return -2147483648
}

func main() {
    var v1 bool = (false && true)
    var v2 string = "hello"
    v3 := make(map[string]uint32)
    if true {
        _ = len(v3)
        var v18 string = v2
        var v19 uint64 = 452
        _ = v18
        _ = v19
    }
    if v1 {
        v20 := []int{647}
        v21 := GenS2{F0: 1}
        _ = v20
        _ = v21
    }
    v22 := GenS3{F0: 42, F1: "\n\t", F2: 255}
    v23 := GenS4{F0: 2}
    v24 := GenS4{F0: 4294967295}
    for fi25 := 0; fi25 < 6; fi25++ {
        v26 := GenS1{F0: "foo bar", F1: v2}
        _ = fi25
        _ = v26
    }
    if true {
        v27 := GenS5{F0: 140, F1: (((v2 + v2) + ((v2 + "\n\t") + v2)) + v2), F2: 42}
        _ = v27
    }
    v28 := GenS6{F0: 42, F1: 65535}
    var v29 uint8 = 70
    var v30 uint8 = (v29 >> ((1) & 63))
    if true {
        for fi31 := 0; fi31 < 17; fi31++ {
            switch (-(fi31)) {
            case -9223372036854775808:
                fmt.Println(v2, fi31, v29)
            default:
                v29 = 127
            }
            var v32 int32 = 17
            _ = fi31
            _ = v32
        }
        for fi33 := 0; fi33 < 20; fi33++ {
            v34 := []int{}
            if v1 {
                fmt.Println(v30, v30, v29)
                for fi35 := 0; fi35 < 16; fi35++ {
                    if true {
                        var v36 string = "a"
                        v29 = (v29 + v29)
                        var v37 uint64 = 815
                        _ = v36
                        _ = v37
                    }
                    v2 = v2
                    switch (fi33 >> ((255) & 63)) {
                    case 2147483647:
                        v30 = (v30 ^ (v29 & (v30 & (v29 ^ v30))))
                    case 1:
                        v38 := GenS2{F0: 1}
                        _ = v38
                    default:
                        delete(v3, v2)
                    }
                    _ = fi35
                }
            } else {
                var v39 uint32 = 2147483647
                for fi40 := 0; fi40 < 15; fi40++ {
                    if v1 {
                        var v41 int32 = 903
                        var v42 bool = false
                        _ = v41
                        _ = v42
                    } else {
                        v43 := GenS4{F0: v39}
                        _ = v43
                    }
                    switch (fi40 % ((fi33) | 1)) {
                    case -9223372036854775808:
                        var v44 int16 = 1000
                        var v45 float32 = (0.1 + 3.14159)
                        v46 := v3["\n\t"]
                        _ = v44
                        _ = v45
                        _ = v46
                    case 138:
                        v47 := GenS6{F0: 1, F1: 42}
                        _ = v47
                    default:
                        var v48 string = (v2 + v2)
                        _ = v48
                    }
                    var v49 uint8 = (v29 - 105)
                    _ = fi40
                    _ = v49
                }
                v29 += (v29 % ((v30) | 1))
                _ = v39
            }
            switch (fi33 / (((-(fi33))) | 1)) {
            case -540:
                var v50 uint64 = 958
                _ = v50
            case 2147483647:
                var v51 uint = 4294967295
                _ = v51
            }
            _ = fi33
            _ = v34
        }
        for fi52 := 0; fi52 < 3; fi52++ {
            _ = len(v3)
            v3[((v2 + "") + v2)] = 65535
            _ = fi52
        }
    }
    var v53 uint32 = 65535
    for fi54 := 0; fi54 < 6; fi54++ {
        v2 = v2
        var v55 int = fi54
        _ = fi54
        _ = v55
    }
    if (32767 == 253) {
        if (!v1) {
            var v56 uint8 = (v30 >> ((4294967295) & 63))
            _ = v56
        }
    } else {
        fmt.Println(v2, v30)
    }
    if (v1 || (111 <= 457)) {
        for fi57 := 0; fi57 < 15; fi57++ {
            v58 := GenS7{F0: -506, F1: 18446744073709551615, F2: (v1 && false)}
            for fi59 := 0; fi59 < 1; fi59++ {
                v60 := GenS7{F0: 1, F1: 65535, F2: v1}
                if false {
                    v61 := GenS6{F0: 1, F1: v53}
                    _ = v61
                }
                _ = len(v3)
                _ = fi59
                _ = v60
            }
            _ = fi57
            _ = v58
        }
        for fi62 := 0; fi62 < 9; fi62++ {
            for fi63 := 0; fi63 < 11; fi63++ {
                v53 = v53
                _ = fi63
            }
            var v64 float64 = 1.5
            _ = fi62
            _ = v64
        }
        v30 -= 127
    } else {
        for fi65 := 0; fi65 < 18; fi65++ {
            var v66 uint32 = (^(v53))
            if v1 {
                v67 := GenS8{F0: ("a" + v2)}
                for fi68 := 0; fi68 < 4; fi68++ {
                    _ = len(v3)
                    _ = fi68
                }
                for fi69 := 0; fi69 < 6; fi69++ {
                    v29 = 2
                    if false {
                        if (v66 <= 2) {
                            v70 := GenS9{F0: fi69, F1: 32767}
                            var v71 uint16 = 65535
                            _ = v70
                            _ = v71
                        } else {
                            var v72 string = ("foo bar" + "\n\t")
                            _ = v72
                        }
                    }
                    _ = fi69
                }
                _ = v67
            } else {
                fmt.Println(v53, v66, v66)
                v73 := []int{fi65, (fi65 ^ 128)}
                _ = v73
            }
            _ = fi65
            _ = v66
        }
    }
    var v74 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v53:", v53)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v53
    _ = v74
}
