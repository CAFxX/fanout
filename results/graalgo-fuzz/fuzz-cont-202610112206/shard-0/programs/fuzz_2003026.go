package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint; F1 bool; F2 uint32 }
type GenS4 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int; F1 int32 }
type GenS9 struct { F0 int8; F1 int64; F2 int32 }
type GenS10 struct { F0 bool }
func genFn0(p0 uint) uint64 {
    p0 = (p0 >> ((p0) & 63))
    switch 42 {
    case 4294967295:
        v4 := GenS1{F0: 470}
        _ = v4
    default:
        for fi5 := 0; fi5 < 2; fi5++ {
            for fi6 := 0; fi6 < 9; fi6++ {
                if true {
                    switch "a" {
                    case "\n\t":
                        var v7 int16 = -104
                        _ = v7
                    default:
                        v8 := []int{(fi5 / ((fi5) | 1))}
                        _ = v8
                    }
                    v9 := GenS1{F0: 1}
                    _ = v9
                }
                p0 = (^(p0))
                _ = fi6
            }
            _ = fi5
        }
    }
    v10 := GenS2{F0: 0, F1: 60}
    v11 := []int{-743, 2147483647, 32768}
    _ = p0
    _ = v10
    _ = v11
    return 0
}
func genFn1(p0 uint, p1 string, p2 int32) uint8 {
    if false {
        if (("a" + ("foo bar" + p1)) != p1) {
            p2 += p2
        }
        var v12 uint = p0
        v12 += v12
        _ = v12
    }
    v13 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 128
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{0, 230}
    v3 := []string{"\n\t", ("\n\t" + "hello"), "", ("\n\t" + "")}
    if true {
        if ("a" != "") {
            var v14 uint8 = genFn1(255, ((("\n\t" + "") + "\n\t") + (("a" + "hello") + "")), -309)
            var v15 uint32 = 2
            switch "" {
            case "hello":
                v16 := GenS3{F0: 9, F1: false, F2: (v15 << ((767) & 63))}
                _ = v16
            case "foo bar":
                v17 := GenS4{F0: 666, F1: 42, F2: v14}
                _ = v17
            default:
                var v18 uint = 65535
                _ = v18
            }
            _ = v14
            _ = v15
        } else {
            if true {
                v19 := GenS1{F0: 213}
                for fi20 := 0; fi20 < 10; fi20++ {
                    v21 := v1[615]
                    _ = fi20
                    _ = v21
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v19
            } else {
                var v22 bool = (false && true)
                _ = v22
            }
            var v23 uint16 = 1
            _ = v23
        }
        var v24 bool = true
        _ = v24
    }
    v25 := GenS3{F0: 965, F1: (-1 >= -60), F2: 606}
    var v26 uint64 = 1
    v27 := GenS5{F0: 215}
    for fi28 := 0; fi28 < 6; fi28++ {
        fmt.Println(fi28)
        v29 := GenS1{F0: 42}
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 10; fi30++ {
        for fi31 := 0; fi31 < 3; fi31++ {
            var v32 uint16 = 2
            _ = fi31
            _ = v32
        }
        switch (("foo bar" + "a") + "a") {
        case "hello":
            if true {
                var v33 uint16 = 65535
                _ = len(v3)
                for fi34 := 0; fi34 < 12; fi34++ {
                    var v35 uint = 858
                    v26 += v26
                    v2 = append(v2, 241)
                    _ = fi34
                    _ = v35
                }
                _ = v33
            } else {
                var v36 int64 = -19
                var v37 uint = 24
                if true {
                    for fi38 := 0; fi38 < 16; fi38++ {
                        fmt.Println(fi30)
                        var v39 bool = (175 >= -32768)
                        _ = fi38
                        _ = v39
                    }
                    var v40 uint32 = 2147483647
                    _ = v40
                } else {
                    v26 += v26
                    var v41 int32 = -2147483648
                    _ = v41
                }
                _ = v36
                _ = v37
            }
        }
        _ = fi30
    }
    v42 := GenS2{F0: 1, F1: 1000}
    switch "a" {
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        v43 := GenS6{F0: 155}
        _ = v43
    }
    for fi44 := 0; fi44 < 1; fi44++ {
        v26 = v26
        _ = fi44
    }
    if (0 <= 1) {
        switch "foo bar" {
        case "foo bar":
            var v45 float64 = 0.0
            _ = v45
        default:
            v46 := GenS5{F0: 1}
            _ = v46
        }
        v47 := GenS3{F0: 147, F1: true, F2: 4294967295}
        v26 -= (v26 << ((18446744073709551615) & 63))
        _ = v47
    }
    if false {
        for fi48 := 0; fi48 < 13; fi48++ {
            v26 = genFn0(18446744073709551615)
            switch ("a" + "") {
            case "hello":
                v49 := GenS7{F0: 42}
                _ = v49
            case "":
                v50 := GenS8{F0: 903, F1: -1}
                _ = v50
            case "\n\t":
                v51 := GenS2{F0: 42, F1: 723}
                _ = v51
            }
            if (!(false && (74 <= 0))) {
                v52 := GenS1{F0: (v26 & 42)}
                for fi53 := 0; fi53 < 16; fi53++ {
                    var v54 string = ""
                    _ = fi53
                    _ = v54
                }
                v26 -= v26
                _ = v52
            }
            _ = fi48
        }
        delete(v1, uint32(v26))
        v26 -= (v26 ^ (^(v26)))
    } else {
        v26 += (^(v26))
    }
    var v55 float32 = (1.5 + (-2.25 + -2.25))
    var v56 uint32 = 462
    var v57 int16 = 1000
    if true {
        for fi58 := 0; fi58 < 9; fi58++ {
            v56 = 0
            _ = fi58
        }
    }
    v56 = 30
    v59 := GenS8{F0: 218, F1: 0}
    for fi60 := 0; fi60 < 13; fi60++ {
        v61 := GenS9{F0: -128, F1: -51, F2: 42}
        v62 := GenS1{F0: v26}
        switch "hello" {
        case "a":
            fmt.Println(v26, v55, fi60)
        case "hello":
            v63 := GenS10{F0: (-448 > 42)}
            _ = v63
        case "\n\t":
            if ((v57 % (((v57 / (((v57 >> ((42) & 63))) | 1))) | 1)) > (v57 & v57)) {
                v56 += 4294967295
                if len(v2) > 0 {
                    v2[0] = v56
                }
            }
        default:
            for fi64 := 0; fi64 < 3; fi64++ {
                var v65 int64 = 42
                var v66 int64 = (-(v65))
                _ = fi64
                _ = v65
                _ = v66
            }
        }
        _ = fi60
        _ = v61
        _ = v62
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v42
    _ = v55
    _ = v56
    _ = v57
    _ = v59
}
