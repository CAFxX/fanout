package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS6 struct { F0 int32 }
func genFn0(p0 int64, p1 uint32, p2 bool) uint16 {
    if true {
        v4 := GenS1{F0: p0, F1: 42, F2: 4294967295}
        var v5 uint64 = 0
        for fi6 := 0; fi6 < 16; fi6++ {
            if true {
                for fi7 := 0; fi7 < 16; fi7++ {
                    v8 := GenS1{F0: p0, F1: 32767, F2: 4294967295}
                    var v9 int = 127
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            }
            _ = fi6
        }
        _ = v4
        _ = v5
    } else {
        v10 := []int{127, 4294967295, 2147483647, 1}
        p1 = (p1 >> ((215) & 63))
        p1 = 250
        _ = v10
    }
    switch -241 {
    case 2147483647:
        v11 := GenS2{F0: 0}
        _ = v11
    case 128:
        var v12 uint8 = 128
        _ = v12
    default:
        if p2 {
            for fi13 := 0; fi13 < 4; fi13++ {
                p0 = (-(p0))
                _ = fi13
            }
            if true {
                for fi14 := 0; fi14 < 13; fi14++ {
                    p1 += 4294967295
                    v15 := GenS1{F0: 1, F1: 65535, F2: (p1 | 792)}
                    for fi16 := 0; fi16 < 13; fi16++ {
                        v17 := []int{fi14}
                        p2 = p2
                        _ = fi16
                        _ = v17
                    }
                    _ = fi14
                    _ = v15
                }
            } else {
                var v18 string = ""
                v19 := GenS3{F0: p1}
                var v20 uint64 = 0
                _ = v18
                _ = v19
                _ = v20
            }
        }
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        v22 := []int{fi21}
        _ = fi21
        _ = v22
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1(p0 int8, p1 uint32) int8 {
    switch -1 {
    case 9223372036854775807:
        p0 -= p0
    case 4294967295:
        v23 := GenS4{F0: -1}
        _ = v23
    }
    var v24 int8 = 42
    switch -230 {
    case -2147483648:
        var v25 string = "hello"
        _ = v25
    case -132:
        switch "foo bar" {
        case "foo bar":
            for fi26 := 0; fi26 < 1; fi26++ {
                v27 := GenS5{F0: 65535, F1: 1, F2: 234}
                _ = fi26
                _ = v27
            }
        }
    }
    for fi28 := 0; fi28 < 14; fi28++ {
        v29 := []int{-448, fi28, fi28, (fi28 & (-(fi28)))}
        _ = fi28
        _ = v29
    }
    _ = p0
    _ = p1
    _ = v24
    return p0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 444
    var v3 int16 = 42
    if true {
        v30 := GenS2{F0: (-(v3))}
        if (!(((0 > 42) && true) && (!((genFn1(-128, 4294967295) > 127) || true)))) {
            for fi31 := 0; fi31 < 9; fi31++ {
                fmt.Println(v3, v3, fi31)
                var v32 uint = 2
                _ = fi31
                _ = v32
            }
            fmt.Println(v3)
            v3 -= v3
        }
        _ = v30
    }
    fmt.Println(v2, v2)
    var v33 uint64 = (^(v2))
    for fi34 := 0; fi34 < 11; fi34++ {
        var v35 int16 = v3
        v36 := GenS3{F0: 2147483647}
        _ = fi34
        _ = v35
        _ = v36
    }
    v3 = (v3 * v3)
    v33 += (v33 / ((42) | 1))
    if (!(!true)) {
        v3 -= (v3 * (^(v3)))
    } else {
        for fi37 := 0; fi37 < 13; fi37++ {
            fmt.Println(v3)
            for fi38 := 0; fi38 < 18; fi38++ {
                v39 := v1[2147483648]
                var v40 int8 = 100
                if false {
                    v41 := GenS1{F0: 42, F1: 65535, F2: 65535}
                    if (331 <= 0) {
                        v42 := []int{128, -9223372036854775808, (fi38 | (fi37 / (((fi37 / ((fi38) | 1))) | 1)))}
                        _ = v42
                    } else {
                        v43 := v1[0]
                        if v43 {
                            v44 := GenS6{F0: 1}
                            fmt.Println(fi38, v33, fi38)
                            _ = v44
                        }
                        v45 := GenS1{F0: -143, F1: 42, F2: 0}
                        _ = v43
                        _ = v45
                    }
                    _ = v41
                }
                _ = fi38
                _ = v39
                _ = v40
            }
            _ = fi37
        }
        v2 -= 2
        v46 := GenS6{F0: -840}
        _ = v46
    }
    if (v2 == v2) {
        for fi47 := 0; fi47 < 3; fi47++ {
            v3 = 52
            if false {
                v48 := []int{fi47, (fi47 & fi47), fi47}
                var v49 uint16 = 255
                var v50 uint32 = uint32(v2)
                _ = v48
                _ = v49
                _ = v50
            } else {
                switch 128 {
                case 128:
                    v51 := []int{fi47, 4294967295}
                    _ = v51
                case 32767:
                    fmt.Println(v33)
                default:
                    v52 := GenS1{F0: 1, F1: 789, F2: 42}
                    _ = v52
                }
            }
            if false {
                switch "a" {
                case "a":
                    for fi53 := 0; fi53 < 15; fi53++ {
                        v2 -= (v2 % (((v2 % (((v2 ^ v33)) | 1))) | 1))
                        var v54 int32 = 0
                        v2 -= 417
                        _ = fi53
                        _ = v54
                    }
                case "\n\t":
                    var v55 int32 = 0
                    _ = v55
                default:
                    v2 -= (v33 << ((42) & 63))
                }
            } else {
                v56 := GenS4{F0: 42}
                _ = v56
            }
            _ = fi47
        }
    } else {
        v2 = 1
        v2 = v2
        var v57 int64 = 42
        _ = v57
    }
    if ("" > "hello") {
        switch "hello" {
        case "hello":
            var v58 int32 = 0
            _ = v58
        }
        var v59 uint16 = uint16(v33)
        v2 = v33
        _ = v59
    }
    var v60 bool = (!true)
    var v61 uint32 = 1
    if v60 {
        fmt.Println(v2, v60)
        var v62 uint8 = 128
        _ = v62
    }
    var v63 uint16 = 543
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v60
    _ = v61
    _ = v63
}
