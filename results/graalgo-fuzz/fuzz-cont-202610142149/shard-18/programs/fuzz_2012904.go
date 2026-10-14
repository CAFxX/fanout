package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int; F1 int; F2 uint32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int16; F1 bool; F2 int8 }
type GenS8 struct { F0 int8; F1 uint32 }
type GenS9 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS10 struct { F0 uint64; F1 int16 }
type GenS11 struct { F0 int; F1 uint64 }
type GenS12 struct { F0 uint; F1 uint; F2 uint8 }
type GenS13 struct { F0 uint8; F1 uint64 }
func genFn0(p0 int8) int {
    v5 := []int{0}
    if true {
        if true {
            fmt.Println(p0)
            if false {
                p0 += (p0 >> ((65535) & 63))
                if (false && false) {
                    v6 := GenS1{F0: 2}
                    _ = v6
                } else {
                    var v7 string = (("a" + ("a" + ("hello" + "hello"))) + "a")
                    switch v7 {
                    case "a":
                        p0 = 100
                    default:
                        v7 = v7
                    }
                    _ = v7
                }
            } else {
                var v8 uint8 = 2
                var v9 float32 = 1.5
                _ = v8
                _ = v9
            }
        }
    }
    _ = p0
    _ = v5
    return 127
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]bool)
    var v3 int16 = 0
    var v4 uint16 = 0
    var v10 int8 = 127
    var v11 uint16 = v4
    fmt.Println(v10, v11)
    v10 += 100
    switch -921 {
    case 35:
        if true {
            if false {
                fmt.Println(v4, v4)
            } else {
                var v12 uint8 = 0
                switch "" {
                case "hello":
                    v13 := GenS2{F0: 643, F1: (^(v12))}
                    _ = v13
                default:
                    fmt.Println(v10, v3, v3)
                }
                _ = v12
            }
            v14 := GenS3{F0: 32767, F1: 9223372036854775807}
            if ((4294967295 < 1) && true) {
                var v15 int16 = (v3 & (v3 + (-(v3))))
                v15 += (v15 | v3)
                v16 := []int{}
                _ = v15
                _ = v16
            } else {
                if (1.5 != (-2.25 - float64(1e10))) {
                    if true {
                        var v17 uint = 376
                        v11 += 32767
                        fmt.Println(v11, v10)
                        _ = v17
                    }
                } else {
                    v4 = v4
                    v18 := GenS2{F0: 445, F1: 127}
                    fmt.Println(v3, v11, v4)
                    _ = v18
                }
                for fi19 := 0; fi19 < 9; fi19++ {
                    if (false && (true && false)) {
                        v20 := []int{}
                        v20 = append(v20, (fi19 & fi19))
                        var v21 int32 = -2147483648
                        _ = v20
                        _ = v21
                    } else {
                        v10 = v10
                    }
                    var v22 uint64 = 42
                    _ = fi19
                    _ = v22
                }
            }
            _ = v14
        } else {
            var v23 uint = uint(v11)
            _ = v23
        }
    case 32768:
        var v24 uint64 = 105
        _ = v24
    default:
        v25 := GenS4{F0: 198}
        _ = v25
    }
    if true {
        if (2147483648 >= 726) {
            v26 := GenS5{F0: int(v4), F1: -2147483648, F2: 1}
            if false {
                v27 := []int{981, genFn0((v10 & -42))}
                _ = v27
            } else {
                v28 := GenS1{F0: (^(v4))}
                _ = v28
            }
            _ = v26
        }
    }
    v29 := v1[65535]
    for fi30 := 0; fi30 < 15; fi30++ {
        var v31 int = fi30
        for fi32 := 0; fi32 < 20; fi32++ {
            var v33 int32 = 2147483647
            switch ("foo bar" + "") {
            case "\n\t":
                v1[1] = v29
            case "foo bar":
                var v34 float64 = -2.25
                _ = v34
            case "hello":
                v4 -= (v4 * v4)
            default:
                var v35 uint32 = v29
                _ = v35
            }
            if (!(1 >= 4294967295)) {
                switch genFn0((v10 * -119)) {
                case -102:
                    var v36 int32 = (v33 * v33)
                    _ = v36
                }
                v31 = fi30
                v31 += genFn0((v10 & v10))
            } else {
                fmt.Println(v3, fi32)
                v37 := []int{v31, (v31 % ((genFn0(v10)) | 1)), genFn0(v10)}
                v38 := GenS1{F0: (v4 - 353)}
                var v39 float64 = (-2.25 - 3.14159)
                _ = v37
                _ = v38
                _ = v39
            }
            v40 := []int{fi30, (fi30 << ((871) & 63))}
            if false {
                var v41 int32 = v33
                _ = v41
            } else {
                fmt.Println(v3)
                v42 := GenS6{F0: uint16(v29)}
                _ = v42
            }
            _ = fi32
            _ = v33
            _ = v40
        }
        if false {
            var v43 uint16 = (v4 + 421)
            _ = v43
        } else {
            var v44 uint16 = v4
            fmt.Println(v11, v4)
            _ = v44
        }
        _ = fi30
        _ = v31
    }
    if (0 != -700) {
        if true {
            switch int(v29) {
            case -641:
                var v45 uint32 = v29
                _ = v45
            case 32768:
                v46 := GenS7{F0: -1, F1: (true || (false || true)), F2: int8(v4)}
                _ = v46
            default:
                switch ("foo bar" + ("a" + "hello")) {
                case "a":
                    v11 += (v4 + 1)
                }
            }
        }
        for fi47 := 0; fi47 < 6; fi47++ {
            v48 := GenS8{F0: v10, F1: v29}
            v49 := GenS6{F0: v11}
            _ = fi47
            _ = v48
            _ = v49
        }
        v29 -= v29
    }
    v50 := []int{genFn0(v10), 4294967295, -2147483648}
    switch "a" {
    case "foo bar":
        v51 := GenS9{F0: 42, F1: true, F2: (^(v29))}
        _ = v51
    case "":
        if false {
            if ((!(((259 == 255) && (-959 > 2147483647)) && true)) || true) {
                switch "" {
                case "":
                    var v52 uint64 = 18446744073709551615
                    _ = v52
                }
                if true {
                    v53 := GenS4{F0: 311}
                    _ = v53
                } else {
                    var v54 uint = 18446744073709551615
                    _ = v54
                }
            }
        } else {
            v3 = (v3 * v3)
            v11 = 1
            if false {
                v4 -= (v4 >> ((516) & 63))
                var v55 int8 = (^(v10))
                var v56 int32 = 2147483647
                _ = v55
                _ = v56
            } else {
                if false {
                    switch 0 {
                    case 4294967295:
                        v57 := GenS10{F0: 4294967295, F1: 366}
                        _ = v57
                    case 445:
                        var v58 int32 = 520
                        _ = v58
                    case 604:
                        v50 = append(v50, genFn0((-(v10))))
                    default:
                        v59 := GenS2{F0: 355, F1: 42}
                        _ = v59
                    }
                    for fi60 := 0; fi60 < 19; fi60++ {
                        var v61 float64 = ((0.0 * ((0.0 + (float64(1.5) * 3.14159)) + 1e10)) + 1e10)
                        _ = fi60
                        _ = v61
                    }
                } else {
                    v62 := GenS11{F0: int(v10), F1: 781}
                    for fi63 := 0; fi63 < 5; fi63++ {
                        _ = len(v2)
                        fmt.Println(v29, v4, v11)
                        _ = fi63
                    }
                    _ = v62
                }
                if (v4 == v4) {
                    fmt.Println(v29)
                }
                for fi64 := 0; fi64 < 4; fi64++ {
                    v65 := GenS6{F0: 540}
                    _ = fi64
                    _ = v65
                }
            }
        }
    case "hello":
        switch -9223372036854775808 {
        case -317:
            if false {
                for fi66 := 0; fi66 < 1; fi66++ {
                    var v67 int32 = 42
                    _ = fi66
                    _ = v67
                }
                switch -2147483648 {
                case 2147483647:
                    v68 := GenS12{F0: 42, F1: 1, F2: 127}
                    _ = v68
                case -970:
                    v69 := GenS13{F0: 1, F1: 0}
                    _ = v69
                }
                var v70 int32 = 1
                _ = v70
            }
        }
    }
    var v71 int16 = (^(v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v50
    _ = v71
}
