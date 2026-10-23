package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint64; F2 int64 }
type GenS4 struct { F0 string; F1 int8 }
func genFn0(p0 bool) uint32 {
    for fi7 := 0; fi7 < 3; fi7++ {
        switch "\n\t" {
        case "a":
            p0 = p0
        case "":
            switch fi7 {
            case 127:
                for fi8 := 0; fi8 < 7; fi8++ {
                    v9 := []int{(fi7 - fi8), fi8, fi8, fi7}
                    _ = fi8
                    _ = v9
                }
            default:
                if p0 {
                    v10 := []int{fi7, 9223372036854775807, (fi7 - (-(fi7)))}
                    p0 = true
                    var v11 string = "foo bar"
                    _ = v10
                    _ = v11
                } else {
                    v12 := GenS1{F0: 1}
                    _ = v12
                }
            }
        case "foo bar":
            v13 := GenS1{F0: 32767}
            _ = v13
        }
        var v14 uint = 65535
        v15 := []int{fi7, -698}
        _ = fi7
        _ = v14
        _ = v15
    }
    p0 = false
    switch ("" + "a") {
    case "hello":
        p0 = p0
    case "\n\t":
        var v16 int = 0
        _ = v16
    default:
        p0 = true
    }
    switch 439 {
    case 263:
        v17 := GenS1{F0: 871}
        _ = v17
    default:
        fmt.Println(p0)
    }
    if true {
        var v18 int32 = 1
        v19 := []int{-2147483648, 128, 0}
        if (!false) {
            var v20 uint16 = 255
            _ = v20
        } else {
            v21 := GenS2{F0: 4294967295}
            v22 := GenS2{F0: 63}
            _ = v21
            _ = v22
        }
        _ = v18
        _ = v19
    } else {
        var v23 int16 = -32768
        _ = v23
    }
    _ = p0
    return 1
}
func genFn1(p0 uint, p1 int32, p2 uint64) int64 {
    p0 = (^(p0))
    for fi24 := 0; fi24 < 13; fi24++ {
        var v25 uint32 = 2147483647
        var v26 int8 = -1
        if true {
            var v27 uint64 = p2
            _ = v27
        } else {
            v28 := []int{fi24}
            _ = v28
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    var v29 uint8 = 233
    for fi30 := 0; fi30 < 20; fi30++ {
        var v31 int32 = p1
        var v32 int64 = 1
        _ = fi30
        _ = v31
        _ = v32
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    return int64(v29)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 32767
    v3 := make(map[string]bool)
    v4 := []int{32767, -2147483648, 0, 80}
    v5 := make(map[uint32]uint32)
    var v6 int16 = 0
    var v33 uint64 = 42
    v2 += (v2 - (v2 ^ (v2 % (((v2 | (v2 >> ((840) & 63)))) | 1))))
    v33 += (v33 & 4294967295)
    switch -2147483648 {
    case 2147483647:
        delete(v5, 1)
    }
    var v34 uint = 687
    switch 2147483647 {
    case -1:
        if true {
            if (2147483648 <= uint32(v34)) {
                v33 += (v33 / (((v33 / ((v33) | 1))) | 1))
                if (42 <= 128) {
                    v35 := GenS3{F0: -61, F1: 137, F2: int64(v33)}
                    var v36 int8 = -128
                    var v37 uint32 = 2147483648
                    _ = v35
                    _ = v36
                    _ = v37
                } else {
                    for fi38 := 0; fi38 < 5; fi38++ {
                        for fi39 := 0; fi39 < 20; fi39++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v40 := GenS3{F0: 0, F1: (v33 ^ v33), F2: -626}
                            _ = fi39
                            _ = v40
                        }
                        var v41 string = "a"
                        _ = fi38
                        _ = v41
                    }
                }
                if true {
                    for fi42 := 0; fi42 < 3; fi42++ {
                        v43 := GenS4{F0: "", F1: 1}
                        v33 = v33
                        _ = fi42
                        _ = v43
                    }
                    fmt.Println(v33, v33)
                } else {
                    var v44 int64 = -471
                    if (((false || (0 < 2)) || (!(0 != (v44 | -9223372036854775808)))) && (((v2 % ((v2) | 1)) > (v2 ^ (v2 >> ((274) & 63)))) || false)) {
                        if (!false) {
                            delete(v3, "a")
                            v45 := GenS1{F0: 1}
                            var v46 string = "hello"
                            _ = v45
                            _ = v46
                        }
                    } else {
                        var v47 string = "foo bar"
                        _ = v47
                    }
                    _ = v44
                }
            } else {
                var v48 int16 = -1
                v49 := GenS4{F0: ("hello" + "a"), F1: -1}
                fmt.Println(v48)
                _ = v48
                _ = v49
            }
            v33 -= v33
        }
    case 127:
        v33 += (v33 & v33)
    case -2147483648:
        _ = len(v4)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v50 := GenS2{F0: uint64(v2)}
    v33 += (v33 & v33)
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        v6 += int16(v33)
    case "":
        var v51 int8 = 16
        _ = v51
    default:
        var v52 string = "foo bar"
        _ = v52
    }
    v4 = append(v4, 32768)
    v53 := GenS4{F0: (("" + "foo bar") + "\n\t"), F1: -1}
    fmt.Println(v6)
    for fi54 := 0; fi54 < 20; fi54++ {
        delete(v1, fi54)
        var v55 uint = 42
        if (true || false) {
            var v56 uint16 = 32767
            _ = v56
        } else {
            for fi57 := 0; fi57 < 11; fi57++ {
                fmt.Println(v55, v34, v55)
                switch ("" + ("" + (("" + (("foo bar" + ("hello" + ("foo bar" + ""))) + ("foo bar" + (("\n\t" + "hello") + "\n\t")))) + ((("\n\t" + ("hello" + "hello")) + ((("a" + "foo bar") + "hello") + "foo bar")) + ("hello" + "foo bar"))))) {
                case "foo bar":
                    var v58 int = (fi57 % ((fi54) | 1))
                    _ = v58
                case "\n\t":
                    v34 += (v34 - 42)
                }
                _ = fi57
            }
        }
        _ = fi54
        _ = v55
    }
    if false {
        _ = len(v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v50
    _ = v53
    _ = v6
}
