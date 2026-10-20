package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint64 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS5 struct { F0 bool; F1 uint }
func genFn0(p0 uint8, p1 int16, p2 int16) uint16 {
    switch "a" {
    case "foo bar":
        v7 := []int{}
        _ = v7
    case "\n\t":
        p1 = p1
    case "hello":
        var v8 uint64 = 2
        _ = v8
    }
    p1 -= -14
    var v9 float32 = 1.5
    var v10 int16 = 1000
    v9 -= 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 521
}
func genFn1(p0 bool, p1 bool) uint16 {
    for fi11 := 0; fi11 < 17; fi11++ {
        v12 := []int{(fi11 >> ((4294967295) & 63)), (-(fi11))}
        _ = len(v12)
        _ = fi11
        _ = v12
    }
    if (18446744073709551615 >= 2) {
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 bool = true
            var v15 bool = p1
            _ = fi13
            _ = v14
            _ = v15
        }
        if p0 {
            var v16 uint64 = 2
            v17 := GenS1{F0: 255, F1: -2147483648}
            v16 = (^(v16))
            _ = v16
            _ = v17
        }
        if p1 {
            v18 := GenS2{F0: 42, F1: 2, F2: 1}
            v19 := []int{817, -1}
            _ = v18
            _ = v19
        } else {
            v20 := []int{32767, -9223372036854775808}
            _ = v20
        }
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 32767
}

func main() {
    v1 := []int{127, -889}
    v2 := make(map[string]uint32)
    v3 := []uint32{}
    v4 := make(map[uint32]bool)
    var v5 int64 = -435
    var v6 bool = (42 <= 1000)
    var v21 uint64 = 2
    if len(v1) > 0 {
        v1[0] = -1
    }
    v5 = v5
    _ = len(v2)
    var v22 int = -1
    for fi23 := 0; fi23 < 15; fi23++ {
        v6 = ((uint32(v21) > 2147483647) && (1 <= 1))
        _ = fi23
    }
    delete(v2, (("\n\t" + "\n\t") + ""))
    for fi24 := 0; fi24 < 7; fi24++ {
        var v25 string = (((("foo bar" + "\n\t") + "") + "hello") + (("" + "") + "foo bar"))
        _ = fi24
        _ = v25
    }
    if (!(v5 != v5)) {
        var v26 int64 = (v5 >> ((1) & 63))
        if v6 {
            for fi27 := 0; fi27 < 4; fi27++ {
                var v28 float64 = (1.5 * float64(3.14159))
                _ = fi27
                _ = v28
            }
            var v29 int64 = (v5 / (((-(v26))) | 1))
            _ = v29
        } else {
            switch ("foo bar" + "\n\t") {
            case "\n\t":
                if v6 {
                    v21 = v21
                    if (914 < -1) {
                        if false {
                            v21 += v21
                        }
                        if (1 <= 67) {
                            for fi30 := 0; fi30 < 16; fi30++ {
                                if (true && (100 < 42)) {
                                    v31 := GenS2{F0: 2, F1: 2, F2: (^(v21))}
                                    for fi32 := 0; fi32 < 17; fi32++ {
                                        v33 := GenS3{F0: ("hello" + "\n\t"), F1: 1}
                                        _ = fi32
                                        _ = v33
                                    }
                                    if (false && v6) {
                                        v3 = append(v3, 65)
                                    }
                                    _ = v31
                                } else {
                                    switch "hello" {
                                    case "":
                                        for fi34 := 0; fi34 < 8; fi34++ {
                                            var v35 int = fi30
                                            delete(v4, 2147483647)
                                            _ = fi34
                                            _ = v35
                                        }
                                    case "\n\t":
                                        if v6 {
                                            fmt.Println(v21)
                                            var v36 string = "\n\t"
                                            _ = v36
                                        }
                                    }
                                    fmt.Println(v6, v21, v22)
                                    if (!v6) {
                                        if (true && false) {
                                            var v37 uint16 = 32767
                                            v1 = append(v1, -1)
                                            _ = v37
                                        }
                                        var v38 uint32 = 65535
                                        _ = v38
                                    } else {
                                        if (int32(v26) < -1) {
                                            switch v22 {
                                            case 32767:
                                                var v39 int = (fi30 - 9223372036854775807)
                                                _ = v39
                                            case -2147483648:
                                                var v40 bool = true
                                                _ = v40
                                            case -967:
                                                for fi41 := 0; fi41 < 9; fi41++ {
                                                    v42 := GenS1{F0: 328, F1: v22}
                                                    v6 = (v6 && ((42 > 18446744073709551615) || (281 < genFn1(v6, ((127 <= 1) || (!(0 < 770)))))))
                                                    _ = fi41
                                                    _ = v42
                                                }
                                            default:
                                                switch ("foo bar" + ((("foo bar" + "\n\t") + "a") + "")) {
                                                case "":
                                                    v43 := GenS1{F0: 0, F1: (v22 | fi30)}
                                                    _ = v43
                                                case "foo bar":
                                                    for fi44 := 0; fi44 < 7; fi44++ {
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            var v45 uint = 18446744073709551615
                                                            _ = v45
                                                        case "hello":
                                                            var v46 int8 = -30
                                                            _ = v46
                                                        case "foo bar":
                                                            if len(v3) > 0 {
                                                                v3[0] = 4294967295
                                                            }
                                                        }
                                                        _ = fi44
                                                    }
                                                }
                                            }
                                            var v47 uint64 = v21
                                            v48 := GenS3{F0: ("a" + (("" + "a") + "\n\t")), F1: 42}
                                            _ = v47
                                            _ = v48
                                        } else {
                                            var v49 string = ("\n\t" + ("hello" + "\n\t"))
                                            for fi50 := 0; fi50 < 3; fi50++ {
                                                v5 = 9223372036854775807
                                                var v51 uint = 286
                                                _ = fi50
                                                _ = v51
                                            }
                                            var v52 bool = true
                                            _ = v49
                                            _ = v52
                                        }
                                        v21 = 400
                                    }
                                }
                                v26 += int64(v21)
                                _ = fi30
                            }
                            fmt.Println(v26, v6, v26)
                        }
                        v6 = ("foo bar" == ((("" + "foo bar") + ((("foo bar" + "foo bar") + "foo bar") + "a")) + "\n\t"))
                    }
                }
            case "foo bar":
                fmt.Println(v5)
            }
            if v6 {
                switch (v22 / ((v22) | 1)) {
                case 128:
                    v22 -= v22
                default:
                    for fi53 := 0; fi53 < 12; fi53++ {
                        for fi54 := 0; fi54 < 7; fi54++ {
                            var v55 uint16 = 1
                            v21 -= v21
                            fmt.Println(fi54, v21)
                            _ = fi54
                            _ = v55
                        }
                        _ = fi53
                    }
                }
                var v56 int32 = 1
                var v57 bool = v6
                _ = v56
                _ = v57
            } else {
                for fi58 := 0; fi58 < 2; fi58++ {
                    if (!false) {
                        fmt.Println(v6, v5)
                        var v59 int16 = 1000
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v59
                    }
                    _ = fi58
                }
            }
        }
        _ = v26
    }
    var v60 int32 = 730
    if (0 > 1000) {
        v60 += (^(v60))
        if true {
            var v61 bool = v6
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v61
        } else {
            v62 := GenS4{F0: v21, F1: 42, F2: 0}
            if v6 {
                for fi63 := 0; fi63 < 1; fi63++ {
                    for fi64 := 0; fi64 < 19; fi64++ {
                        var v65 bool = false
                        var v66 uint = 2
                        _ = fi64
                        _ = v65
                        _ = v66
                    }
                    _ = fi63
                }
                fmt.Println(v21)
                var v67 uint = 18446744073709551615
                _ = v67
            }
            fmt.Println(v5)
            _ = v62
        }
        switch "\n\t" {
        case "\n\t":
            fmt.Println(v21)
        case "a":
            v68 := GenS5{F0: (v6 && (786 > v60)), F1: 778}
            _ = v68
        case "foo bar":
            v5 = (v5 % ((v5) | 1))
        default:
            v5 = v5
        }
    }
    fmt.Println(v21, v5, v6)
    for fi69 := 0; fi69 < 9; fi69++ {
        if (100 <= -1) {
            v60 = v60
        }
        _ = fi69
    }
    switch "\n\t" {
    case "hello":
        v5 = v5
    default:
        for fi70 := 0; fi70 < 7; fi70++ {
            for fi71 := 0; fi71 < 17; fi71++ {
                v22 += int(v21)
                delete(v2, "a")
                fmt.Println(v21, v21)
                _ = fi71
            }
            fmt.Println(v21, v22, v5)
            _ = fi70
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
}
