package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int16 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 string; F1 int16; F2 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 bool }
func genFn0() int {
    v7 := GenS1{F0: true, F1: -94, F2: 292}
    for fi8 := 0; fi8 < 17; fi8++ {
        fmt.Println(fi8)
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := []int{}
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    var v11 string = ""
    for fi12 := 0; fi12 < 15; fi12++ {
        if (0 >= 2) {
            var v13 uint = 4294967295
            _ = v13
        }
        var v14 float64 = -2.25
        _ = fi12
        _ = v14
    }
    _ = v11
    _ = v7
    return 2147483647
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[uint32]bool)
    var v3 uint = 255
    var v4 float64 = 3.14159
    v5 := make(map[int]int)
    var v6 uint = (^(v3))
    v1 = (v1 + (v1 << ((uint(v1)) & 63)))
    v15 := []int{-1, 1}
    var v16 bool = (-128 >= 41)
    var v17 uint8 = 44
    for fi18 := 0; fi18 < 14; fi18++ {
        for fi19 := 0; fi19 < 12; fi19++ {
            if true {
                v20 := GenS2{F0: 0, F1: "a"}
                _ = v20
            } else {
                v4 = v4
                v21 := GenS2{F0: 4294967295, F1: "\n\t"}
                _ = v21
            }
            _ = fi19
        }
        if v16 {
            if len(v15) > 0 {
                v15[0] = 4294967295
            }
            switch "hello" {
            case "foo bar":
                v4 = v4
            case "a":
                v22 := GenS1{F0: ((v6 >> ((1) & 63)) > v3), F1: -109, F2: 42}
                _ = v22
            default:
                for fi23 := 0; fi23 < 9; fi23++ {
                    v4 -= 3.14159
                    v3 = v6
                    v24 := GenS3{F0: ("" + ("foo bar" + "a")), F1: 1, F2: ("hello" + "\n\t")}
                    var v25 bool = (v16 || (v6 < v6))
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
        }
        _ = fi18
    }
    fmt.Println(v6, v3)
    if len(v15) > 0 {
        v15[0] = 2147483647
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        var v27 uint = v6
        _ = fi26
        _ = v27
    }
    v28 := GenS4{F0: (!false)}
    for fi29 := 0; fi29 < 15; fi29++ {
        var v30 int32 = 1
        var v31 bool = (42 >= (v30 | 338))
        fmt.Println(fi29)
        _ = fi29
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 13; fi32++ {
        fmt.Println(v1)
        _ = fi32
    }
    switch "a" {
    case "":
        switch 9223372036854775807 {
        case -729:
            for fi33 := 0; fi33 < 6; fi33++ {
                v6 = (v6 - (^(v3)))
                _ = fi33
            }
        case 32767:
            fmt.Println(v1, v3, v6)
        case 909:
            switch 2147483647 {
            case -9223372036854775808:
                var v34 int16 = 741
                _ = v34
            default:
                var v35 uint16 = 477
                _ = v35
            }
        }
    default:
        v36 := GenS1{F0: (3.14159 < float32((float64(-2.25) / v4))), F1: 90, F2: -421}
        _ = v36
    }
    delete(v2, 2147483648)
    if (42 >= 1) {
        v6 = (v6 & v3)
        delete(v2, 0)
        for fi37 := 0; fi37 < 7; fi37++ {
            if (((32767 != 2) && (0.1 < 1.5)) && false) {
                if (127 != 127) {
                    for fi38 := 0; fi38 < 10; fi38++ {
                        var v39 string = ("a" + ("a" + ""))
                        if len(v15) > 0 {
                            v15[0] = fi38
                        }
                        v6 = v3
                        v17 = v17
                        _ = fi38
                        _ = v39
                    }
                } else {
                    fmt.Println(fi37, v17, v16)
                    v3 -= uint(v17)
                    for fi40 := 0; fi40 < 14; fi40++ {
                        v1 = 2147483648
                        v41 := GenS4{F0: (("a" < "foo bar") || (!v16))}
                        v42 := GenS5{F0: 4294967295, F1: false}
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                }
                if (v4 < (v4 / 1.5)) {
                    var v43 float64 = v4
                    _ = v43
                }
            } else {
                v44 := GenS2{F0: 4294967295, F1: "foo bar"}
                var v45 uint = uint(v1)
                _ = v44
                _ = v45
            }
            v46 := GenS3{F0: "", F1: int16(fi37), F2: ("foo bar" + ("" + "foo bar"))}
            _ = fi37
            _ = v46
        }
    }
    for fi47 := 0; fi47 < 20; fi47++ {
        v48 := v2[(v1 + uint32(v17))]
        _ = fi47
        _ = v48
    }
    var v49 float32 = 3.14159
    fmt.Println(v16, v3)
    fmt.Println(v3, v49)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
