package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int; F1 uint; F2 int8 }
type GenS4 struct { F0 uint32; F1 string; F2 uint }
type GenS5 struct { F0 int8; F1 string }
func genFn0(p0 uint32) uint64 {
    switch ("\n\t" + "hello") {
    case "a":
        for fi7 := 0; fi7 < 8; fi7++ {
            p0 -= (p0 >> ((0) & 63))
            _ = fi7
        }
    case "foo bar":
        v8 := []int{1, 46, 32768}
        _ = v8
    case "hello":
        var v9 uint64 = 2
        _ = v9
    default:
        switch -1 {
        case 1:
            if (4294967295 < 2) {
                p0 = 740
            } else {
                if true {
                    v10 := []int{2147483647}
                    _ = len(v10)
                    _ = v10
                }
                p0 += (p0 << ((255) & 63))
                v11 := GenS1{F0: -296, F1: true}
                _ = v11
            }
        case 128:
            v12 := GenS1{F0: -9223372036854775808, F1: false}
            _ = v12
        }
    }
    if true {
        if (42 == -9223372036854775808) {
            for fi13 := 0; fi13 < 17; fi13++ {
                p0 = (p0 ^ 65535)
                if (31 >= 1) {
                    for fi14 := 0; fi14 < 2; fi14++ {
                        var v15 bool = true
                        var v16 uint32 = (p0 - p0)
                        v15 = v15
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                } else {
                    v17 := []int{(-(fi13))}
                    _ = v17
                }
                p0 -= (p0 & p0)
                _ = fi13
            }
        }
    } else {
        v18 := []int{-2147483648, 9223372036854775807, 42, 1}
        fmt.Println(p0)
        v19 := GenS2{F0: true, F1: (0 >= 4294967295)}
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        p0 += 2147483647
        switch fi20 {
        case 573:
            for fi21 := 0; fi21 < 14; fi21++ {
                fmt.Println(fi20, p0)
                v22 := []int{127}
                _ = fi21
                _ = v22
            }
        case 127:
            if true {
                if true {
                    p0 += (p0 >> ((946) & 63))
                    v23 := []int{-245, fi20, 42}
                    _ = v23
                } else {
                    if (!(664 != 18446744073709551615)) {
                        fmt.Println(fi20, p0)
                        var v24 uint64 = 1
                        v25 := GenS1{F0: 4294967295, F1: (25 >= 2)}
                        _ = v24
                        _ = v25
                    } else {
                        v26 := GenS2{F0: (42 < 128), F1: true}
                        var v27 bool = false
                        v28 := []int{}
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                    var v29 float32 = 0.1
                    var v30 int16 = 32767
                    _ = v29
                    _ = v30
                }
                var v31 int8 = -1
                _ = v31
            } else {
                for fi32 := 0; fi32 < 17; fi32++ {
                    var v33 bool = false
                    p0 -= 1
                    v33 = v33
                    _ = fi32
                    _ = v33
                }
            }
        }
        var v34 int16 = 608
        _ = fi20
        _ = v34
    }
    v35 := GenS2{F0: false, F1: (255 < 1)}
    _ = p0
    _ = v35
    return 42
}
func genFn1() int16 {
    for fi36 := 0; fi36 < 14; fi36++ {
        if (!((3.14159 - (1e10 + 3.14159)) > 1e10)) {
            switch "a" {
            case "foo bar":
                v37 := []int{}
                _ = v37
            }
        }
        _ = fi36
    }
    if (((("a" + "a") + ("hello" + ("foo bar" + ("foo bar" + ("" + ""))))) >= "a") && ((!(!false)) && false)) {
        var v38 uint16 = 42
        var v39 int = 9223372036854775807
        _ = v38
        _ = v39
    }
    v40 := []int{}
    _ = v40
    return 629
}

func main() {
    var v1 int8 = -128
    var v2 string = ((("a" + ("a" + "foo bar")) + (("hello" + "foo bar") + "a")) + "")
    var v3 int16 = 1000
    v4 := make(map[string]string)
    var v5 uint8 = 2
    var v6 int64 = -1
    var v41 uint8 = v5
    var v42 string = v2
    var v43 int8 = 101
    switch 472 {
    case 0:
        v5 -= v41
    default:
        fmt.Println(v3)
    }
    switch v2 {
    case "":
        switch ("hello" + "\n\t") {
        case "\n\t":
            fmt.Println(v43, v5, v3)
        case "":
            var v44 string = "a"
            _ = v44
        default:
            v45 := GenS1{F0: 32767, F1: (4294967295 <= 0)}
            _ = v45
        }
    }
    var v46 uint16 = uint16(v5)
    var v47 uint32 = uint32(v46)
    if (uint8(v47) == 255) {
        var v48 uint8 = (v41 & uint8(v47))
        switch 1 {
        case 2147483647:
            v47 = (^(v47))
        }
        v2 = v42
        _ = v48
    } else {
        var v49 uint8 = 17
        v50 := []int{128, 4294967295, 9223372036854775807}
        v51 := GenS3{F0: 954, F1: uint(v46), F2: (v43 | v1)}
        _ = v49
        _ = v50
        _ = v51
    }
    v52 := v4[v42]
    v53 := GenS3{F0: int(v6), F1: 255, F2: 127}
    v46 += uint16(v5)
    fmt.Println(v41)
    if false {
        if ((v47 << ((923) & 63)) == 2147483648) {
            v54 := GenS3{F0: 2147483647, F1: 0, F2: v1}
            v55 := GenS4{F0: v47, F1: "foo bar", F2: 18446744073709551615}
            v56 := GenS3{F0: -631, F1: uint(v41), F2: (v1 << ((2) & 63))}
            _ = v54
            _ = v55
            _ = v56
        } else {
            v57 := GenS5{F0: 42, F1: v52}
            for fi58 := 0; fi58 < 17; fi58++ {
                if true {
                    switch fi58 {
                    case -736:
                        v52 = v52
                    default:
                        fmt.Println(v2)
                    }
                    var v59 string = (v52 + v42)
                    _ = v59
                }
                for fi60 := 0; fi60 < 19; fi60++ {
                    var v61 int8 = (v43 << ((990) & 63))
                    _ = fi60
                    _ = v61
                }
                _ = fi58
            }
            _ = v57
        }
        var v62 int16 = (v3 ^ v3)
        for fi63 := 0; fi63 < 16; fi63++ {
            var v64 uint64 = genFn0(v47)
            v3 = (^(v3))
            fmt.Println(v6, v41)
            _ = fi63
            _ = v64
        }
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v46
    _ = v47
    _ = v5
    _ = v52
    _ = v53
    _ = v6
}
