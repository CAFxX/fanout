package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 int8; F1 uint16; F2 string }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int8; F1 uint; F2 int }
func genFn0(p0 uint8, p1 int8) uint32 {
    var v5 float64 = 3.14159
    p1 -= (p1 ^ 42)
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := []int{fi6}
        if true {
            v8 := GenS1{F0: "\n\t", F1: uint(p0), F2: 667}
            for fi9 := 0; fi9 < 10; fi9++ {
                v10 := GenS2{F0: ("\n\t" + "")}
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                var v11 string = "hello"
                _ = fi9
                _ = v10
                _ = v11
            }
            switch fi6 {
            case -1:
                var v12 bool = (748 == 32767)
                _ = v12
            default:
                for fi13 := 0; fi13 < 20; fi13++ {
                    p1 -= (p1 ^ 100)
                    p1 = p1
                    _ = fi13
                }
            }
            _ = v8
        } else {
            for fi14 := 0; fi14 < 19; fi14++ {
                v15 := GenS3{F0: -1, F1: 18446744073709551615}
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                if (float32(v5) <= float32(v5)) {
                    var v16 uint64 = 18446744073709551615
                    _ = v16
                }
                _ = fi14
                _ = v15
            }
            switch "\n\t" {
            case "":
                v7 = append(v7, (fi6 / ((640) | 1)))
            case "foo bar":
                if ((p0 >> ((255) & 63)) != 1) {
                    var v17 float64 = v5
                    _ = v17
                }
            default:
                var v18 int64 = -273
                _ = v18
            }
            v5 -= float64(0.1)
        }
        switch "\n\t" {
        case "a":
            for fi19 := 0; fi19 < 6; fi19++ {
                if len(v7) > 0 {
                    v7[0] = (fi19 >> ((65535) & 63))
                }
                _ = fi19
            }
        case "\n\t":
            v20 := GenS3{F0: -32768, F1: 1}
            _ = v20
        }
        _ = fi6
        _ = v7
    }
    var v21 string = "hello"
    if ((!false) && (-2.25 >= float32(1e10))) {
        fmt.Println(p1, p1)
    } else {
        fmt.Println(p1, v21, p1)
        p0 -= p0
    }
    _ = p0
    _ = p1
    _ = v21
    _ = v5
    return 65535
}
func genFn1(p0 bool, p1 int32, p2 uint16) int {
    var v22 uint64 = 2
    var v23 bool = (false || (100 <= 127))
    p2 -= p2
    var v24 uint64 = v22
    switch -2147483648 {
    case 42:
        for fi25 := 0; fi25 < 2; fi25++ {
            var v26 uint = 65535
            var v27 int = -1
            switch -2147483648 {
            case -86:
                var v28 float64 = 0.1
                _ = v28
            }
            _ = fi25
            _ = v26
            _ = v27
        }
    default:
        if true {
            switch "\n\t" {
            case "a":
                for fi29 := 0; fi29 < 16; fi29++ {
                    var v30 int16 = -582
                    p0 = false
                    v31 := []int{fi29}
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            default:
                p0 = (220 == uint8(p2))
            }
        } else {
            p2 = p2
            var v32 uint16 = p2
            p1 = p1
            _ = v32
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    _ = v24
    return 32767
}

func main() {
    var v1 uint32 = 2147483648
    var v2 bool = (false && false)
    var v3 uint = 2
    var v4 uint8 = 1
    if v2 {
        v33 := []int{128, 32767}
        if len(v33) >= 2 {
            v33 = v33[1:len(v33)]
        }
        var v34 string = "hello"
        _ = v33
        _ = v34
    } else {
        switch 625 {
        case 1:
            v35 := []int{-951, 32767, -1, -557}
            _ = v35
        case 42:
            for fi36 := 0; fi36 < 3; fi36++ {
                v37 := GenS4{F0: -96, F1: v2}
                _ = fi36
                _ = v37
            }
        default:
            v38 := []int{}
            _ = v38
        }
    }
    var v39 int64 = 9223372036854775807
    var v40 int8 = 127
    var v41 int64 = (v39 & (-(v39)))
    if (1 < v1) {
        v41 = -230
        v42 := GenS4{F0: 1, F1: ((!(65535 <= v3)) || v2)}
        _ = v42
    } else {
        for fi43 := 0; fi43 < 20; fi43++ {
            v40 += v40
            var v44 int16 = 1000
            _ = fi43
            _ = v44
        }
        if v2 {
            var v45 uint16 = 2
            fmt.Println(v40)
            var v46 int64 = (v41 << ((18446744073709551615) & 63))
            _ = v45
            _ = v46
        } else {
            var v47 bool = ((v41 ^ (v41 >> ((v3) & 63))) < (v39 << ((uint(v4)) & 63)))
            _ = v47
        }
        var v48 uint32 = v1
        _ = v48
    }
    var v49 float32 = float32(3.14159)
    v40 = v40
    for fi50 := 0; fi50 < 13; fi50++ {
        v51 := []int{-386, genFn1(("\n\t" >= ""), -2147483648, 448)}
        _ = len(v51)
        switch genFn1(v2, -179, 2) {
        case -2147483648:
            if (0 > -1) {
                for fi52 := 0; fi52 < 9; fi52++ {
                    var v53 uint = (v3 & (v3 + (v3 | v3)))
                    if (v2 || ((!true) && true)) {
                        v54 := GenS5{F0: 7, F1: 65535, F2: "foo bar"}
                        _ = v54
                    } else {
                        v55 := GenS6{F0: ("\n\t" + "a")}
                        _ = v55
                    }
                    var v56 int8 = (v40 - (v40 ^ v40))
                    _ = fi52
                    _ = v53
                    _ = v56
                }
                v57 := GenS7{F0: -128, F1: (v3 << (((^(v3))) & 63)), F2: fi50}
                _ = v57
            }
        }
        _ = fi50
        _ = v51
    }
    switch 974 {
    case -9223372036854775808:
        var v58 bool = true
        _ = v58
    case 127:
        switch 242 {
        case 2147483647:
            v59 := []int{166, 32768}
            _ = v59
        }
    default:
        if v2 {
            var v60 bool = true
            v60 = v2
            v60 = (-630 <= (v41 % ((1) | 1)))
            _ = v60
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v49
}
