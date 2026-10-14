package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS3 struct { F0 bool; F1 uint }
func genFn0(p0 bool) int {
    p0 = (1 == 1)
    p0 = false
    v6 := []int{}
    if (!p0) {
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 int64 = 42
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 10; fi9++ {
            for fi10 := 0; fi10 < 11; fi10++ {
                var v11 string = "\n\t"
                _ = fi10
                _ = v11
            }
            p0 = (false && (!("hello" < "hello")))
            v12 := GenS1{F0: 131, F1: ""}
            _ = fi9
            _ = v12
        }
        switch -683 {
        case 1:
            v13 := GenS1{F0: 2, F1: (("a" + "hello") + "")}
            _ = v13
        case -2147483648:
            p0 = (!(!p0))
        default:
            v6 = append(v6, 205)
        }
    }
    _ = p0
    _ = v6
    return 327
}
func genFn1(p0 int16, p1 bool, p2 uint64) int32 {
    fmt.Println(p1, p1)
    v14 := []int{-125, 9223372036854775807, 1, 32768}
    for fi15 := 0; fi15 < 10; fi15++ {
        var v16 int16 = (p0 | 42)
        fmt.Println(fi15, p0, p1)
        _ = fi15
        _ = v16
    }
    var v17 int64 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v17
    return 42
}

func main() {
    var v1 bool = (1 >= -260)
    var v2 int32 = 0
    var v3 int8 = 42
    var v4 string = "\n\t"
    var v5 uint8 = 1
    for fi18 := 0; fi18 < 19; fi18++ {
        fmt.Println(v1)
        var v19 uint16 = 1
        for fi20 := 0; fi20 < 9; fi20++ {
            v1 = v1
            if (0.0 >= 1e10) {
                var v21 bool = (!true)
                fmt.Println(fi20, fi20, fi20)
                _ = v21
            } else {
                if v1 {
                    var v22 uint32 = uint32(v5)
                    _ = v22
                } else {
                    var v23 int = 9223372036854775807
                    _ = v23
                }
                var v24 uint32 = 65535
                if (0.1 == (3.14159 * float32((3.14159 * (3.14159 - 1.5))))) {
                    v25 := GenS2{F0: (v24 - (v24 + (v24 - (^(v24))))), F1: (v5 + v5), F2: v19}
                    _ = v25
                }
                _ = v24
            }
            _ = fi20
        }
        _ = fi18
        _ = v19
    }
    v2 = v2
    v4 = v4
    v3 += (-(v3))
    v26 := GenS3{F0: false, F1: 4294967295}
    for fi27 := 0; fi27 < 3; fi27++ {
        for fi28 := 0; fi28 < 5; fi28++ {
            var v29 float32 = 1e10
            if true {
                v30 := []int{}
                if (-36 == 42) {
                    var v31 string = ("foo bar" + ("" + v4))
                    if (849 > 942) {
                        if false {
                            v32 := GenS2{F0: 1, F1: 2, F2: 443}
                            var v33 uint8 = (^(v5))
                            v34 := GenS3{F0: (false && v1), F1: 1}
                            _ = v32
                            _ = v33
                            _ = v34
                        } else {
                            switch genFn0(v1) {
                            case -1:
                                var v35 uint8 = (v5 - 2)
                                _ = v35
                            case 32768:
                                if (!v1) {
                                    var v36 int32 = v2
                                    v5 = v5
                                    v37 := GenS3{F0: true, F1: 2}
                                    _ = v36
                                    _ = v37
                                }
                            default:
                                var v38 int = fi28
                                _ = v38
                            }
                            v29 = ((0.0 / float32(((0.1 + 1e10) / (1.5 / 1e10)))) - 0.1)
                            switch v31 {
                            case "foo bar":
                                var v39 float64 = (3.14159 / (float64((v29 + v29)) * 3.14159))
                                _ = v39
                            }
                        }
                    }
                    v40 := GenS1{F0: 0, F1: (v31 + "a")}
                    v30 = append(v30, fi27)
                    _ = v31
                    _ = v40
                } else {
                    v30 = append(v30, 270)
                    var v41 uint32 = 42
                    _ = v41
                }
                switch "foo bar" {
                case "a":
                    v1 = v1
                default:
                    switch 9223372036854775807 {
                    case -974:
                        var v42 uint16 = 42
                        _ = v42
                    default:
                        var v43 int32 = (v2 >> ((18446744073709551615) & 63))
                        _ = v43
                    }
                }
                _ = v30
            }
            _ = fi28
            _ = v29
        }
        if ((("foo bar" + ((((v4 + ("" + v4)) + v4) + ((v4 + v4) + "hello")) + "\n\t")) < (((v4 + v4) + ("hello" + "")) + v4)) || false) {
            v5 = (v5 * 42)
        } else {
            var v44 uint64 = 4294967295
            _ = v44
        }
        v4 = v4
        _ = fi27
    }
    var v45 int32 = 42
    if v1 {
        var v46 int8 = v3
        var v47 int32 = 2147483647
        if (!v1) {
            for fi48 := 0; fi48 < 19; fi48++ {
                v46 -= 100
                var v49 uint16 = 369
                var v50 string = (v4 + v4)
                _ = fi48
                _ = v49
                _ = v50
            }
        }
        _ = v46
        _ = v47
    }
    switch 32768 {
    case -898:
        v45 = 0
    }
    v51 := []int{}
    var v52 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v51
    _ = v52
}
