package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int8; F1 bool; F2 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 uint64; F2 bool }
func genFn0(p0 bool) uint8 {
    v7 := []int{42, -2147483648, 128}
    if p0 {
        fmt.Println(p0)
        switch 4294967295 {
        case -736:
            v8 := GenS1{F0: -245, F1: 3, F2: 167}
            _ = v8
        case -1:
            fmt.Println(p0)
        default:
            v9 := GenS2{F0: 127, F1: (2 >= 255), F2: 1}
            _ = v9
        }
        if p0 {
            if true {
                switch ("foo bar" + "foo bar") {
                case "\n\t":
                    v10 := GenS3{F0: 98}
                    _ = v10
                case "":
                    p0 = false
                case "foo bar":
                    var v11 int16 = 895
                    _ = v11
                }
            }
            var v12 uint8 = 2
            v12 += v12
            _ = v12
        }
    } else {
        _ = len(v7)
        p0 = p0
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        if (false || ((false || (127 < 1)) && p0)) {
            p0 = p0
            if true {
                v14 := GenS2{F0: 42, F1: true, F2: 45}
                if (0 >= 671) {
                    if ((p0 && (!(!false))) && (42 != 9223372036854775807)) {
                        v7 = append(v7, (fi13 + 42))
                        var v15 int8 = -79
                        var v16 int32 = 42
                        _ = v15
                        _ = v16
                    } else {
                        var v17 uint32 = 2147483648
                        if len(v7) > 0 {
                            v7[0] = (fi13 | (^(fi13)))
                        }
                        v18 := GenS2{F0: 7, F1: (18446744073709551615 >= 18446744073709551615), F2: 1}
                        _ = v17
                        _ = v18
                    }
                    v19 := GenS1{F0: 362, F1: 255, F2: 2}
                    v20 := GenS1{F0: -9223372036854775808, F1: 0, F2: 1}
                    _ = v19
                    _ = v20
                }
                _ = v14
            } else {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                var v21 int8 = 127
                for fi22 := 0; fi22 < 16; fi22++ {
                    v23 := GenS1{F0: 42, F1: 128, F2: 4294967295}
                    fmt.Println(fi13, p0)
                    var v24 uint16 = 2
                    var v25 int32 = 989
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                }
                _ = v21
            }
        }
        for fi26 := 0; fi26 < 4; fi26++ {
            p0 = (42 >= 127)
            _ = fi26
        }
        _ = fi13
    }
    _ = p0
    _ = v7
    return 160
}

func main() {
    var v1 uint = 331
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]string)
    v4 := []int{996, 0}
    var v5 uint32 = 2147483648
    var v6 float32 = 3.14159
    v27 := GenS4{F0: true, F1: uint64(v5), F2: true}
    for fi28 := 0; fi28 < 18; fi28++ {
        v1 = v1
        for fi29 := 0; fi29 < 3; fi29++ {
            switch (fi28 - -1) {
            case -2147483648:
                var v30 float64 = 1e10
                _ = v30
            case 127:
                var v31 uint8 = 1
                _ = v31
            default:
                var v32 uint16 = uint16(v5)
                _ = v32
            }
            var v33 int = 32
            _ = fi29
            _ = v33
        }
        _ = fi28
    }
    v34 := GenS4{F0: ((("foo bar" + "foo bar") + "\n\t") > ("\n\t" + "\n\t")), F1: 1, F2: false}
    v5 = (v5 ^ (v5 >> (((v1 & (^(v1)))) & 63)))
    var v35 int = 4294967295
    v35 = v35
    for fi36 := 0; fi36 < 8; fi36++ {
        fmt.Println(v1, v5, v5)
        _ = fi36
    }
    v1 = v1
    var v37 uint = (v1 % ((v1) | 1))
    v35 = (v35 ^ 9223372036854775807)
    for fi38 := 0; fi38 < 18; fi38++ {
        switch "\n\t" {
        case "foo bar":
            var v39 uint32 = v5
            _ = v39
        case "\n\t":
            if len(v4) > 0 {
                v4[0] = (fi38 * (v35 + 603))
            }
        }
        if false {
            v3[v5] = "hello"
            var v40 bool = (false && (0 > -1))
            _ = v40
        } else {
            var v41 uint32 = 2
            v42 := GenS4{F0: false, F1: 18446744073709551615, F2: (true && true)}
            v35 -= v35
            _ = v41
            _ = v42
        }
        _ = fi38
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
