package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int32; F1 string; F2 uint }
type GenS3 struct { F0 int64; F1 string }
func genFn0(p0 int32, p1 uint, p2 int32) uint8 {
    fmt.Println(p0, p1, p0)
    var v6 uint32 = 0
    v7 := GenS1{F0: -32768, F1: 2}
    for fi8 := 0; fi8 < 5; fi8++ {
        if false {
            var v9 int64 = 9223372036854775807
            _ = v9
        } else {
            if ((true || false) && ((v6 % (((v6 * (v6 >> (((p1 + 4294967295)) & 63)))) | 1)) >= (v6 >> ((p1) & 63)))) {
                v10 := GenS2{F0: (p2 >> ((493) & 63)), F1: "hello", F2: (p1 | (p1 << (((^(p1))) & 63)))}
                if (!(!true)) {
                    switch ((("foo bar" + "foo bar") + "foo bar") + "a") {
                    case "hello":
                        v11 := []int{}
                        _ = v11
                    case "\n\t":
                        v12 := GenS2{F0: -2147483648, F1: (("foo bar" + "") + ("hello" + "a")), F2: 18446744073709551615}
                        _ = v12
                    default:
                        if (42 <= 1) {
                            v13 := GenS2{F0: p0, F1: "", F2: p1}
                            v14 := GenS2{F0: (^(p2)), F1: "a", F2: (p1 >> ((p1) & 63))}
                            var v15 bool = (false || (!false))
                            _ = v13
                            _ = v14
                            _ = v15
                        } else {
                            var v16 uint32 = 2147483647
                            _ = v16
                        }
                    }
                    v17 := []int{}
                    _ = len(v17)
                    _ = v17
                }
                _ = v10
            } else {
                switch 0 {
                case 128:
                    p0 -= 1
                case 9223372036854775807:
                    for fi18 := 0; fi18 < 3; fi18++ {
                        var v19 uint64 = 4294967295
                        _ = fi18
                        _ = v19
                    }
                default:
                    var v20 uint = (p1 | (p1 + (p1 | (p1 & 0))))
                    _ = v20
                }
                if (!(0 >= 255)) {
                    v21 := []int{(fi8 * fi8), fi8}
                    for fi22 := 0; fi22 < 8; fi22++ {
                        if true {
                            v23 := GenS3{F0: 354, F1: "hello"}
                            _ = v23
                        }
                        p1 = p1
                        _ = fi22
                    }
                    _ = v21
                }
            }
            p0 += p0
        }
        var v24 uint16 = 32767
        _ = fi8
        _ = v24
    }
    if (!false) {
        if true {
            for fi25 := 0; fi25 < 11; fi25++ {
                for fi26 := 0; fi26 < 3; fi26++ {
                    v27 := GenS3{F0: 42, F1: ("hello" + (("\n\t" + "hello") + "foo bar"))}
                    v28 := []int{fi25, (fi26 & (fi25 ^ fi25)), fi25}
                    var v29 uint64 = 0
                    _ = fi26
                    _ = v27
                    _ = v28
                    _ = v29
                }
                var v30 int16 = int16(fi25)
                var v31 int = (-(fi25))
                _ = fi25
                _ = v30
                _ = v31
            }
            for fi32 := 0; fi32 < 8; fi32++ {
                if false {
                    v33 := []int{}
                    if true {
                        for fi34 := 0; fi34 < 14; fi34++ {
                            v35 := GenS3{F0: -9223372036854775808, F1: ("" + "")}
                            var v36 uint64 = 4294967295
                            v33 = append(v33, fi32)
                            var v37 int64 = -1
                            if false {
                                if len(v33) >= 2 {
                                    v33 = v33[1:len(v33)]
                                }
                                p2 = p0
                            }
                            v38 := GenS3{F0: (v37 - v37), F1: "foo bar"}
                            fmt.Println(fi32, p1)
                            var v39 bool = (450 >= 42)
                            v40 := GenS2{F0: p2, F1: ("\n\t" + "a"), F2: (p1 / (((p1 * (^(p1)))) | 1))}
                            var v41 float32 = 0.1
                            var v42 int8 = -1
                            _ = fi34
                            _ = v35
                            _ = v36
                            _ = v37
                            _ = v38
                            _ = v39
                            _ = v40
                            _ = v41
                            _ = v42
                        }
                    }
                    if len(v33) > 0 {
                        v33[0] = (fi32 % (((fi32 % (((fi32 % ((fi32) | 1))) | 1))) | 1))
                    }
                    _ = v33
                } else {
                    var v43 uint64 = 190
                    _ = v43
                }
                _ = fi32
            }
            for fi44 := 0; fi44 < 7; fi44++ {
                v45 := []int{(fi44 >> (((p1 >> ((1) & 63))) & 63)), fi44, fi44}
                _ = fi44
                _ = v45
            }
        } else {
            v6 = (^(v6))
            v46 := []int{9223372036854775807, 129, 4294967295}
            _ = v46
        }
    } else {
        v47 := []int{-1, 1, -534, 676}
        _ = v47
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint, p1 string, p2 int16) int16 {
    fmt.Println(p2, p1)
    v48 := GenS3{F0: 333, F1: (((("" + p1) + p1) + p1) + p1)}
    p1 = ("" + (("foo bar" + p1) + p1))
    if (1.5 <= 0.1) {
        var v49 uint32 = 4294967295
        v50 := []int{}
        _ = v49
        _ = v50
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v48
    return 1
}

func main() {
    var v1 int16 = 42
    v2 := []uint32{2147483648, 29, 2}
    v3 := make(map[uint32]int)
    var v4 uint8 = 1
    var v5 int = 0
    v2 = append(v2, 2147483647)
    delete(v3, 4294967295)
    if false {
        var v51 float64 = 1e10
        _ = v51
    }
    for fi52 := 0; fi52 < 1; fi52++ {
        var v53 uint = 711
        _ = fi52
        _ = v53
    }
    v54 := GenS3{F0: -9223372036854775808, F1: "\n\t"}
    v55 := GenS1{F0: -1, F1: 255}
    for fi56 := 0; fi56 < 14; fi56++ {
        for fi57 := 0; fi57 < 9; fi57++ {
            switch (-(v5)) {
            case 42:
                v5 = (-(fi56))
            case 32767:
                v1 = genFn1(255, (("a" + ("\n\t" + "a")) + ("\n\t" + "\n\t")), v1)
            }
            _ = fi57
        }
        fmt.Println(fi56)
        switch v5 {
        case -409:
            for fi58 := 0; fi58 < 9; fi58++ {
                var v59 int16 = genFn1(65535, "a", (v1 - v1))
                if len(v2) > 0 {
                    v2[0] = 1
                }
                _ = fi58
                _ = v59
            }
        case -914:
            for fi60 := 0; fi60 < 19; fi60++ {
                var v61 uint32 = 42
                var v62 int64 = 89
                v62 = int64(v5)
                _ = fi60
                _ = v61
                _ = v62
            }
        }
        _ = fi56
    }
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        fmt.Println(v4, v1, v5)
    }
    if ((!false) || true) {
        var v63 string = (("\n\t" + ("hello" + "foo bar")) + "a")
        v5 = 4294967295
        _ = v63
    } else {
        for fi64 := 0; fi64 < 5; fi64++ {
            var v65 string = (("a" + (("hello" + (("" + (("hello" + "foo bar") + ("" + "a"))) + ("foo bar" + ((("\n\t" + "") + "hello") + "foo bar")))) + ("hello" + ("a" + "foo bar")))) + "hello")
            var v66 int16 = (v1 >> ((4294967295) & 63))
            _ = fi64
            _ = v65
            _ = v66
        }
        fmt.Println(v5, v4)
        v67 := GenS2{F0: 0, F1: ("foo bar" + ""), F2: 65535}
        _ = v67
    }
    v68 := GenS2{F0: 1, F1: ("\n\t" + "\n\t"), F2: 18446744073709551615}
    fmt.Println(v1, v4)
    var v69 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v68
    _ = v69
}
