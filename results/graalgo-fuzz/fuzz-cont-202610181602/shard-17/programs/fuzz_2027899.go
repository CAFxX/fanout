package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint; F2 bool }
type GenS3 struct { F0 string; F1 string; F2 uint64 }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 uint16; F1 uint32; F2 string }
type GenS6 struct { F0 int; F1 int64; F2 uint8 }
type GenS7 struct { F0 uint64; F1 uint32 }
func genFn0(p0 uint64, p1 int) int64 {
    if false {
        for fi6 := 0; fi6 < 11; fi6++ {
            v7 := []int{p1}
            _ = fi6
            _ = v7
        }
        v8 := GenS1{F0: -785, F1: 128}
        _ = v8
    }
    var v9 uint16 = 32
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := GenS2{F0: 138, F1: 0, F2: false}
        v12 := GenS2{F0: 202, F1: 65535, F2: (true && true)}
        var v13 string = (("\n\t" + "a") + ((("hello" + ((("a" + (("\n\t" + "\n\t") + (("foo bar" + (((("\n\t" + (("foo bar" + "foo bar") + "foo bar")) + "foo bar") + ("foo bar" + ((("a" + "a") + "") + "\n\t"))) + ("foo bar" + "hello"))) + "hello"))) + (("foo bar" + "foo bar") + "hello")) + "\n\t")) + ("foo bar" + "")) + "hello"))
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    v14 := GenS3{F0: ("foo bar" + "foo bar"), F1: "\n\t", F2: (p0 + (p0 | (p0 * p0)))}
    p1 = p1
    _ = p0
    _ = p1
    _ = v14
    _ = v9
    return -1
}
func genFn1(p0 uint8, p1 uint) int {
    var v15 uint16 = 2
    if false {
        var v16 uint16 = (v15 + v15)
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 string = ""
            _ = fi17
            _ = v18
        }
        p0 = p0
        _ = v16
    }
    var v19 float64 = (float64(1e10) + (1.5 + 3.14159))
    var v20 int = 537
    _ = p0
    _ = p1
    _ = v15
    _ = v19
    _ = v20
    return 37
}

func main() {
    var v1 uint16 = 65535
    var v2 float32 = 0.0
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]string)
    var v5 int16 = -1
    v2 = 3.14159
    switch ("foo bar" + "foo bar") {
    case "a":
        var v21 uint64 = 18446744073709551615
        _ = v21
    default:
        var v22 float64 = ((0.1 + 3.14159) / -2.25)
        _ = v22
    }
    v23 := GenS2{F0: 32767, F1: 4294967295, F2: false}
    var v24 string = "a"
    v5 -= (^(v5))
    var v25 int16 = (v5 & (v5 >> ((18446744073709551615) & 63)))
    v1 = v1
    v4[4294967295] = "a"
    v26 := []int{127}
    for fi27 := 0; fi27 < 2; fi27++ {
        v1 = 1
        _ = fi27
    }
    switch 32767 {
    case 882:
        switch 4294967295 {
        case 398:
            v1 = (v1 ^ (v1 - 512))
        case 4294967295:
            for fi28 := 0; fi28 < 19; fi28++ {
                fmt.Println(v2, fi28, fi28)
                _ = fi28
            }
        }
    case 640:
        for fi29 := 0; fi29 < 4; fi29++ {
            for fi30 := 0; fi30 < 15; fi30++ {
                var v31 uint32 = 1
                if ((511 <= 42) && true) {
                    v2 = v2
                    if len(v26) > 0 {
                        v26[0] = 32767
                    }
                    switch v24 {
                    case "hello":
                        if false {
                            v1 -= (v1 << ((42) & 63))
                            v3[genFn1(127, 0)] = uint32(v1)
                            if (18446744073709551615 == 18446744073709551615) {
                                var v32 uint16 = 2
                                var v33 uint = 0
                                v1 = v1
                                _ = v32
                                _ = v33
                            }
                        }
                    case "\n\t":
                        var v34 string = v24
                        _ = v34
                    case "a":
                        delete(v3, -1)
                    default:
                        var v35 float32 = 1.5
                        _ = v35
                    }
                } else {
                    if true {
                        v4[(v31 ^ 1)] = (v24 + "\n\t")
                        v31 = v31
                    }
                }
                _ = fi30
                _ = v31
            }
            for fi36 := 0; fi36 < 3; fi36++ {
                fmt.Println(fi36)
                v37 := v3[genFn1(38, 18446744073709551615)]
                v2 = v2
                _ = fi36
                _ = v37
            }
            var v38 bool = true
            _ = fi29
            _ = v38
        }
    case 1:
        switch genFn1(42, 2) {
        case 4294967295:
            switch (v24 + "hello") {
            case "\n\t":
                var v39 float64 = (0.0 * (1e10 * ((-2.25 - 3.14159) * (float64(0.0) + (1e10 + 3.14159)))))
                _ = v39
            }
        case 1:
            var v40 int16 = (v5 + (v25 >> ((18446744073709551615) & 63)))
            _ = v40
        }
    default:
        var v41 int64 = 790
        _ = v41
    }
    var v42 int8 = 127
    for fi43 := 0; fi43 < 11; fi43++ {
        v42 += (^(v42))
        fmt.Println(v1)
        _ = fi43
    }
    switch 127 {
    case 127:
        for fi44 := 0; fi44 < 2; fi44++ {
            switch v24 {
            case "foo bar":
                fmt.Println(v5)
            case "hello":
                fmt.Println(v1, v2, v42)
            default:
                delete(v4, 0)
            }
            _ = fi44
        }
    case 910:
        v45 := v4[0]
        _ = v45
    default:
        switch v24 {
        case "hello":
            for fi46 := 0; fi46 < 16; fi46++ {
                delete(v4, uint32(v1))
                v47 := GenS4{F0: (false || false), F1: v24}
                v5 = (v25 >> ((362) & 63))
                _ = fi46
                _ = v47
            }
        case "foo bar":
            v48 := GenS4{F0: ((!true) || (-2147483648 < 4294967295)), F1: v24}
            _ = v48
        default:
            if false {
                v49 := GenS5{F0: (^(v1)), F1: 42, F2: ((("foo bar" + v24) + (v24 + v24)) + v24)}
                for fi50 := 0; fi50 < 6; fi50++ {
                    fmt.Println(v24, v1, v1)
                    v2 = v2
                    _ = fi50
                }
                switch v24 {
                case "\n\t":
                    if (((898 <= 2) && true) || (!true)) {
                        var v51 string = "a"
                        var v52 int16 = (v5 | v5)
                        for fi53 := 0; fi53 < 8; fi53++ {
                            var v54 string = (v24 + v51)
                            var v55 int = -707
                            _ = fi53
                            _ = v54
                            _ = v55
                        }
                        _ = v51
                        _ = v52
                    } else {
                        v56 := GenS1{F0: 1000, F1: 32}
                        _ = v56
                    }
                default:
                    for fi57 := 0; fi57 < 17; fi57++ {
                        var v58 uint16 = (^(v1))
                        if (841 == 1) {
                            for fi59 := 0; fi59 < 15; fi59++ {
                                var v60 uint32 = 356
                                if ((fi59 % ((fi59) | 1)) == fi59) {
                                    var v61 uint32 = (v60 << ((85) & 63))
                                    if len(v26) > 0 {
                                        v26[0] = fi59
                                    }
                                    _ = v61
                                } else {
                                    if false {
                                        v24 = v24
                                        var v62 uint8 = 0
                                        _ = v62
                                    }
                                }
                                var v63 uint64 = 18446744073709551615
                                _ = fi59
                                _ = v60
                                _ = v63
                            }
                            v1 = v58
                            for fi64 := 0; fi64 < 3; fi64++ {
                                v65 := v4[0]
                                var v66 uint8 = 67
                                _ = fi64
                                _ = v65
                                _ = v66
                            }
                        } else {
                            v67 := GenS6{F0: genFn1(0, 65535), F1: 25, F2: 31}
                            _ = v67
                        }
                        _ = len(v4)
                        _ = fi57
                        _ = v58
                    }
                }
                _ = v49
            }
        }
    }
    v68 := GenS7{F0: 18446744073709551615, F1: 2147483648}
    v4[0] = ""
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v68
}
