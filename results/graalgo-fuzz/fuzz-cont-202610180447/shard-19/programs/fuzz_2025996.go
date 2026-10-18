package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint }
type GenS2 struct { F0 bool; F1 int }
func genFn0(p0 int64, p1 string, p2 int) int64 {
    var v7 uint16 = 2
    for fi8 := 0; fi8 < 1; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 bool = ((-289 < 1) && true)
            _ = fi9
            _ = v10
        }
        if (3.14159 <= 0.0) {
            if false {
                v7 = 65535
            }
            switch p1 {
            case "":
                p1 = "a"
            default:
                v11 := GenS1{F0: 8, F1: (-(fi8)), F2: 65535}
                _ = v11
            }
            switch fi8 {
            case -2147483648:
                var v12 int8 = 12
                _ = v12
            case 356:
                switch p2 {
                case 128:
                    p1 = ("foo bar" + p1)
                case -1:
                    var v13 uint = 42
                    _ = v13
                }
            case 127:
                var v14 string = (p1 + p1)
                _ = v14
            }
        }
        _ = fi8
    }
    v15 := GenS1{F0: 2, F1: p2, F2: 255}
    p1 = (p1 + p1)
    switch (^(p2)) {
    case 0:
        v16 := []int{p2}
        _ = v16
    case -383:
        var v17 uint64 = 42
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v7
    return p0
}
func genFn1() uint16 {
    for fi18 := 0; fi18 < 14; fi18++ {
        var v19 string = "foo bar"
        for fi20 := 0; fi20 < 13; fi20++ {
            v21 := []int{fi18, (fi18 & 128), fi18}
            var v22 int8 = 1
            var v23 float64 = 0.0
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
        _ = fi18
        _ = v19
    }
    var v24 bool = (!(true || false))
    switch "\n\t" {
    case "hello":
        v25 := GenS1{F0: 1, F1: 1, F2: 0}
        _ = v25
    case "a":
        v24 = (-1 == 1000)
    case "":
        v26 := []int{775, 127, 761}
        _ = v26
    default:
        var v27 bool = (true && (v24 && (!(("foo bar" + "a") <= "hello"))))
        _ = v27
    }
    v24 = v24
    var v28 uint8 = 1
    _ = v24
    _ = v28
    return 0
}

func main() {
    var v1 int32 = 2147483647
    var v2 int32 = (v1 / (((-(v1))) | 1))
    var v3 uint8 = 104
    var v4 float32 = -2.25
    var v5 int8 = 42
    var v6 uint = 42
    if false {
        for fi29 := 0; fi29 < 7; fi29++ {
            v30 := []int{128, (fi29 + fi29), (^(fi29)), 128}
            _ = fi29
            _ = v30
        }
    }
    var v31 uint16 = 2
    for fi32 := 0; fi32 < 7; fi32++ {
        var v33 int8 = (v5 / ((42) | 1))
        var v34 float32 = v4
        if (!(!false)) {
            v2 = (v2 & 42)
            for fi35 := 0; fi35 < 3; fi35++ {
                var v36 int64 = -1
                v31 = 42
                _ = fi35
                _ = v36
            }
            for fi37 := 0; fi37 < 15; fi37++ {
                v38 := []int{}
                _ = fi37
                _ = v38
            }
        }
        _ = fi32
        _ = v33
        _ = v34
    }
    v3 = v3
    v6 = v6
    if (-32768 != int16(v6)) {
        var v39 uint16 = v31
        v40 := []int{2147483647, -346}
        _ = v39
        _ = v40
    } else {
        v3 -= v3
        v6 += (v6 + 18446744073709551615)
    }
    v41 := GenS2{F0: false, F1: 2147483647}
    v1 = (v2 / ((v1) | 1))
    v42 := []int{}
    if len(v42) >= 2 {
        v42 = v42[1:len(v42)]
    }
    if (uint64(v31) > 4294967295) {
        _ = len(v42)
        v1 = int32(v31)
        if false {
            if (-779 < 4294967295) {
                var v43 int16 = 1
                _ = v43
            } else {
                switch -915 {
                case 0:
                    for fi44 := 0; fi44 < 11; fi44++ {
                        if (false && true) {
                            if (int8(v31) == v5) {
                                var v45 bool = true
                                _ = v45
                            }
                            v6 = v6
                        } else {
                            var v46 uint32 = 4294967295
                            _ = v46
                        }
                        if (!(!((v3 >= v3) && true))) {
                            v47 := GenS2{F0: false, F1: fi44}
                            var v48 uint32 = 128
                            v3 = (v3 ^ (v3 - 162))
                            _ = v47
                            _ = v48
                        }
                        v42 = append(v42, -2147483648)
                        _ = fi44
                    }
                case 4294967295:
                    v31 -= (v31 >> ((42) & 63))
                case 32767:
                    var v49 uint16 = genFn1()
                    _ = v49
                }
                if false {
                    v50 := GenS2{F0: true, F1: 447}
                    fmt.Println(v4, v5)
                    v31 = (v31 | 42)
                    _ = v50
                }
            }
            v3 += v3
            if true {
                fmt.Println(v3, v6, v31)
                var v51 int16 = -145
                v5 += 42
                _ = v51
            }
        } else {
            fmt.Println(v31, v1, v4)
        }
    }
    v31 += uint16(v3)
    if false {
        var v52 uint = (v6 ^ (v6 & uint(v3)))
        _ = v52
    } else {
        if ("\n\t" <= "") {
            var v53 float64 = (-2.25 - 1e10)
            v2 = 42
            _ = v53
        }
        var v54 uint8 = 42
        v31 -= 32767
        _ = v54
    }
    v55 := GenS1{F0: 128, F1: -2147483648, F2: v6}
    if false {
        _ = len(v42)
        var v56 uint64 = 18446744073709551615
        var v57 int32 = (v1 ^ (v1 ^ v2))
        _ = v56
        _ = v57
    }
    fmt.Println(v2)
    if len(v42) > 0 {
        v42[0] = 32768
    }
    if len(v42) >= 2 {
        v42 = v42[1:len(v42)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v55
    _ = v6
}
