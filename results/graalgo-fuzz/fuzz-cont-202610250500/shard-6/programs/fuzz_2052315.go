package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 int8; F1 uint64 }
func genFn0(p0 string) bool {
    for fi6 := 0; fi6 < 19; fi6++ {
        p0 = p0
        var v7 uint64 = 1
        for fi8 := 0; fi8 < 14; fi8++ {
            if ((((0 <= 0) || (1 < 1)) || true) && ((true && (!(!((v7 % (((v7 - 264)) | 1)) != (^(v7)))))) && ((fi8 >> ((42) & 63)) == fi8))) {
                v9 := GenS1{F0: "a", F1: p0, F2: -1}
                _ = v9
            } else {
                var v10 int16 = 1000
                var v11 bool = true
                _ = v10
                _ = v11
            }
            var v12 int64 = -551
            switch (fi6 ^ (fi8 * (^(fi6)))) {
            case 539:
                v13 := GenS2{F0: 2147483648}
                _ = v13
            case 685:
                if (((3.14159 + 1e10) - 0.1) <= 1.5) {
                    var v14 string = (p0 + p0)
                    var v15 int = fi8
                    var v16 float64 = 0.0
                    _ = v14
                    _ = v15
                    _ = v16
                }
            }
            _ = fi8
            _ = v12
        }
        _ = fi6
        _ = v7
    }
    fmt.Println(p0)
    _ = p0
    return true
}
func genFn1() int16 {
    var v17 uint8 = 127
    var v18 int = 127
    v18 -= (v18 ^ v18)
    v19 := []int{v18, (v18 - -2147483648), (v18 * v18), -927}
    _ = v17
    _ = v18
    _ = v19
    return 1
}

func main() {
    var v1 uint8 = 205
    var v2 uint16 = 32767
    var v3 uint = 2
    var v4 int32 = 2147483647
    var v5 int = 127
    for fi20 := 0; fi20 < 16; fi20++ {
        switch 247 {
        case 32768:
            for fi21 := 0; fi21 < 6; fi21++ {
                fmt.Println(v5, v5, v1)
                if true {
                    v5 = fi20
                    var v22 uint32 = 40
                    switch "hello" {
                    case "foo bar":
                        v3 -= (v3 + (v3 | (v3 << ((623) & 63))))
                    case "hello":
                        var v23 uint8 = (^(v1))
                        _ = v23
                    case "":
                        switch 9223372036854775807 {
                        case 1:
                            switch (fi21 & (fi20 >> ((1) & 63))) {
                            case 127:
                                v24 := GenS3{F0: "a", F1: v2, F2: v22}
                                _ = v24
                            case -499:
                                var v25 int8 = -64
                                _ = v25
                            case -500:
                                var v26 int16 = 389
                                _ = v26
                            default:
                                var v27 int16 = -263
                                _ = v27
                            }
                        }
                    }
                    _ = v22
                }
                if ((v2 | (v2 / ((1) | 1))) != 1) {
                    v1 -= (^(v1))
                }
                _ = fi21
            }
        case -1:
            fmt.Println(fi20)
        case 2147483647:
            fmt.Println(v1)
        default:
            v28 := []int{fi20}
            _ = v28
        }
        _ = fi20
    }
    for fi29 := 0; fi29 < 7; fi29++ {
        v5 += (-(v5))
        fmt.Println(v4)
        v4 -= (v4 - 0)
        _ = fi29
    }
    v30 := GenS4{F0: 42, F1: int64(v1)}
    for fi31 := 0; fi31 < 12; fi31++ {
        var v32 int64 = 42
        if (!true) {
            var v33 uint8 = 1
            if false {
                fmt.Println(v4, v5, v4)
            } else {
                for fi34 := 0; fi34 < 3; fi34++ {
                    var v35 int64 = -246
                    v36 := GenS2{F0: 553}
                    var v37 int32 = -2147483648
                    v5 -= (fi31 ^ 4294967295)
                    _ = fi34
                    _ = v35
                    _ = v36
                    _ = v37
                }
                if (!false) {
                    v38 := GenS1{F0: ("hello" + ("a" + "\n\t")), F1: "\n\t", F2: -128}
                    v33 = (v33 / ((v33) | 1))
                    for fi39 := 0; fi39 < 10; fi39++ {
                        v40 := []int{fi31, 127}
                        _ = fi39
                        _ = v40
                    }
                    _ = v38
                }
            }
            _ = v33
        } else {
            if (!(817 < 107)) {
                v1 = v1
                var v41 uint8 = (v1 | 1)
                var v42 uint8 = v41
                _ = v41
                _ = v42
            }
            v43 := []int{}
            switch (fi31 & int(v32)) {
            case 4294967295:
                var v44 string = ""
                _ = v44
            case 32767:
                for fi45 := 0; fi45 < 13; fi45++ {
                    _ = len(v43)
                    v46 := GenS2{F0: 756}
                    var v47 string = "foo bar"
                    _ = fi45
                    _ = v46
                    _ = v47
                }
            }
            _ = v43
        }
        v48 := []int{496, (-(fi31)), (v5 | (v5 << ((0) & 63)))}
        _ = fi31
        _ = v32
        _ = v48
    }
    switch ("foo bar" + "a") {
    case "\n\t":
        if (1.5 <= 1e10) {
            switch v5 {
            case 127:
                var v49 string = ""
                _ = v49
            case 227:
                v50 := []int{128}
                _ = v50
            default:
                var v51 string = ""
                _ = v51
            }
            fmt.Println(v2)
            var v52 string = ""
            _ = v52
        } else {
            switch "hello" {
            case "hello":
                v53 := GenS5{F0: -40, F1: uint64(v2)}
                _ = v53
            case "a":
                v54 := []int{v5, v5}
                _ = v54
            case "\n\t":
                v3 = v3
            }
            if false {
                v55 := []int{0, -9223372036854775808}
                var v56 int = v5
                fmt.Println(v5)
                _ = v55
                _ = v56
            }
            switch v5 {
            case -2147483648:
                var v57 string = ("foo bar" + "")
                _ = v57
            case -641:
                v3 -= (^(v3))
            default:
                for fi58 := 0; fi58 < 12; fi58++ {
                    v2 = (^(v2))
                    v59 := []int{fi58}
                    _ = fi58
                    _ = v59
                }
            }
        }
    case "hello":
        switch -339 {
        case -2147483648:
            if false {
                var v60 bool = true
                _ = v60
            } else {
                v61 := []int{}
                _ = v61
            }
        case 127:
            var v62 bool = ((!false) && (((v1 | v1) > v1) || (2147483648 > 1)))
            _ = v62
        case 0:
            if (-604 >= -390) {
                var v63 bool = (-1 < -693)
                var v64 string = "foo bar"
                _ = v63
                _ = v64
            }
        }
    case "":
        if (true || false) {
            v65 := GenS5{F0: int8(v3), F1: 2}
            _ = v65
        }
    default:
        v5 = 2147483647
    }
    var v66 float64 = 0.0
    fmt.Println(v66)
    v67 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v66:", v66)
    fmt.Println("slicelen v67:", len(v67), v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v66
    _ = v67
}
