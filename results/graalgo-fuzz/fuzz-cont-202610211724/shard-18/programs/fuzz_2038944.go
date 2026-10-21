package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 int64; F1 int; F2 uint64 }
type GenS4 struct { F0 int; F1 uint64; F2 uint }
type GenS5 struct { F0 uint16; F1 uint }
func genFn0(p0 string) uint8 {
    switch (((("foo bar" + ((p0 + ("\n\t" + (p0 + ("" + ("\n\t" + "a"))))) + "foo bar")) + (p0 + p0)) + (p0 + "\n\t")) + (p0 + p0)) {
    case "hello":
        v4 := GenS1{F0: 42, F1: 65535, F2: 85}
        _ = v4
    case "":
        v5 := []int{-2147483648, 2147483647, 32767}
        _ = v5
    case "a":
        p0 = "foo bar"
    default:
        v6 := []int{128}
        _ = v6
    }
    var v7 float64 = 0.0
    fmt.Println(v7)
    switch "hello" {
    case "hello":
        v7 -= 0.0
    case "foo bar":
        switch "" {
        case "hello":
            v7 += 0.1
        case "\n\t":
            for fi8 := 0; fi8 < 18; fi8++ {
                var v9 bool = (p0 >= ("a" + "foo bar"))
                var v10 float64 = (v7 - (3.14159 + (1e10 - float64((-2.25 + float32((1e10 + 0.0)))))))
                _ = fi8
                _ = v10
                _ = v9
            }
        case "":
            p0 = "foo bar"
        default:
            for fi11 := 0; fi11 < 18; fi11++ {
                if (-2.25 != 0.0) {
                    var v12 float32 = (3.14159 + -2.25)
                    for fi13 := 0; fi13 < 5; fi13++ {
                        if false {
                            p0 = p0
                            p0 = p0
                            v12 = (v12 + 3.14159)
                        } else {
                            p0 = ((p0 + p0) + (p0 + (((p0 + p0) + "") + (p0 + p0))))
                            var v14 int8 = -128
                            _ = v14
                        }
                        if false {
                            v15 := GenS2{F0: 65535, F1: -1}
                            switch fi11 {
                            case 9223372036854775807:
                                var v16 uint32 = 65535
                                _ = v16
                            case 128:
                                var v17 string = p0
                                _ = v17
                            default:
                                var v18 int = 32768
                                switch -298 {
                                case 291:
                                    var v19 int8 = 42
                                    _ = v19
                                default:
                                    if true {
                                        v20 := []int{32768}
                                        _ = v20
                                    }
                                }
                                fmt.Println(fi13, v7)
                                _ = v18
                            }
                            _ = v15
                        }
                        _ = fi13
                    }
                    _ = v12
                }
                _ = fi11
            }
        }
    case "\n\t":
        v21 := GenS1{F0: 0, F1: 1, F2: 305}
        _ = v21
    }
    _ = p0
    _ = v7
    return 2
}
func genFn1(p0 uint8, p1 int8) int16 {
    var v22 bool = (-2147483648 > -689)
    fmt.Println(p0, p0)
    var v23 int8 = 0
    switch ("a" + "") {
    case "a":
        var v24 float64 = -2.25
        _ = v24
    case "":
        p0 = p0
    default:
        v25 := []int{}
        _ = v25
    }
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    return 32767
}

func main() {
    var v1 uint8 = 2
    var v2 uint16 = 65535
    v3 := make(map[int]string)
    fmt.Println(v2)
    var v26 string = "\n\t"
    v27 := GenS2{F0: (v2 * 0), F1: -2147483648}
    if false {
        switch -1 {
        case 32768:
            if false {
                if (1 <= (v1 & genFn0(v26))) {
                    for fi28 := 0; fi28 < 13; fi28++ {
                        for fi29 := 0; fi29 < 5; fi29++ {
                            if false {
                                var v30 int32 = 174
                                v31 := GenS3{F0: 577, F1: fi29, F2: 464}
                                v32 := GenS3{F0: 1, F1: (fi28 + 127), F2: 779}
                                if ((!false) && (v30 <= -597)) {
                                    v33 := []int{fi28, fi28, -1}
                                    _ = v33
                                } else {
                                    var v34 int64 = 1
                                    v35 := GenS2{F0: v2, F1: (v30 - v30)}
                                    v36 := GenS1{F0: 2, F1: 2147483648, F2: 110}
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                }
                                _ = v30
                                _ = v31
                                _ = v32
                            }
                            _ = fi29
                        }
                        for fi37 := 0; fi37 < 10; fi37++ {
                            if ((v2 - (v2 | v2)) > (v2 << ((702) & 63))) {
                                v38 := GenS4{F0: (^(fi28)), F1: 1, F2: 954}
                                _ = v38
                            }
                            fmt.Println(v2, v26)
                            _ = fi37
                        }
                        _ = fi28
                    }
                }
            }
        case -1:
            _ = len(v3)
        case -2147483648:
            fmt.Println(v1, v2)
        default:
            for fi39 := 0; fi39 < 13; fi39++ {
                v40 := v3[(fi39 ^ fi39)]
                _ = fi39
                _ = v40
            }
        }
        var v41 uint8 = 255
        _ = v41
    } else {
        var v42 uint16 = 0
        _ = v42
    }
    v2 -= (v2 & v2)
    for fi43 := 0; fi43 < 11; fi43++ {
        fmt.Println(v26)
        if true {
            v3[(fi43 / (((fi43 / ((fi43) | 1))) | 1))] = ((v26 + (("\n\t" + v26) + "\n\t")) + v26)
            for fi44 := 0; fi44 < 3; fi44++ {
                if (((false || false) || false) && true) {
                    v1 -= 1
                    var v45 int = (fi44 + 128)
                    _ = v45
                } else {
                    if false {
                        v46 := GenS5{F0: 32767, F1: 402}
                        _ = v46
                    } else {
                        v26 = (v26 + ((((v26 + v26) + ((v26 + v26) + (v26 + v26))) + v26) + v26))
                        v3[(fi44 ^ int(v1))] = "\n\t"
                        v47 := GenS4{F0: fi44, F1: 1, F2: 42}
                        v2 -= (v2 / ((v2) | 1))
                        var v48 float64 = -2.25
                        fmt.Println(v1)
                        _ = v47
                        _ = v48
                    }
                    var v49 int32 = 0
                    fmt.Println(v26)
                    _ = v49
                }
                fmt.Println(v2)
                _ = fi44
            }
            var v50 bool = (1 < -1)
            _ = v50
        }
        v51 := GenS2{F0: v2, F1: 1}
        _ = fi43
        _ = v51
    }
    var v52 uint = 518
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v52
}
