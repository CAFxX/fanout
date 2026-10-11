package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int; F1 string; F2 int32 }
type GenS3 struct { F0 int64; F1 int8 }
func genFn0() uint8 {
    var v7 uint32 = 2147483648
    v8 := GenS1{F0: "foo bar", F1: -1}
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 string = "\n\t"
        for fi11 := 0; fi11 < 17; fi11++ {
            switch "a" {
            case "":
                if true {
                    v12 := GenS1{F0: (("hello" + "a") + v10), F1: -2147483648}
                    v10 = ((v10 + v10) + ("foo bar" + "a"))
                    v10 = "\n\t"
                    _ = v12
                }
                v13 := []int{(fi9 + fi9), (fi9 & fi11), (fi11 << ((4294967295) & 63)), (^(fi11))}
                _ = v13
            case "\n\t":
                switch v10 {
                case "\n\t":
                    v10 = v10
                    var v14 uint64 = 42
                    if false {
                        var v15 uint16 = 65535
                        if false {
                            v16 := []int{fi11, fi9}
                            var v17 int32 = -2147483648
                            _ = v16
                            _ = v17
                        } else {
                            if (2147483647 > fi9) {
                                var v18 float64 = -2.25
                                if ((fi11 | (fi9 | fi11)) >= fi9) {
                                    var v19 int16 = 940
                                    var v20 string = "hello"
                                    var v21 uint16 = 506
                                    var v22 uint8 = 2
                                    var v23 float64 = v18
                                    var v24 float64 = ((v23 + -2.25) * float64((1.5 / 0.1)))
                                    var v25 string = (v10 + "\n\t")
                                    var v26 uint = 913
                                    var v27 uint8 = (v22 ^ v22)
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                } else {
                                    var v28 uint = 255
                                    _ = v28
                                }
                                _ = v18
                            }
                            v15 = 65535
                            var v29 uint32 = 4294967295
                            _ = v29
                        }
                        fmt.Println(v14)
                        v7 = (v7 << ((346) & 63))
                        switch fi11 {
                        case 32768:
                            var v30 string = (v10 + "foo bar")
                            _ = v30
                        default:
                            var v31 int64 = -9223372036854775808
                            _ = v31
                        }
                        _ = v15
                    } else {
                        if true {
                            v32 := []int{(fi11 & (fi9 % ((fi11) | 1))), (fi11 << ((65535) & 63)), fi11, fi9}
                            _ = v32
                        }
                        v10 = ("hello" + "")
                    }
                    _ = v14
                case "foo bar":
                    v10 = ("a" + "hello")
                case "hello":
                    var v33 float32 = (0.0 - 0.0)
                    fmt.Println(v7, fi11, fi9)
                    fmt.Println(fi9, v7)
                    _ = v33
                }
            case "a":
                v34 := GenS2{F0: fi9, F1: v10, F2: -1}
                _ = v34
            default:
                var v35 float32 = 0.1
                _ = v35
            }
            v7 = (v7 & (v7 >> ((18446744073709551615) & 63)))
            _ = fi11
        }
        if false {
            var v36 uint16 = 2
            fmt.Println(v7, v7, v10)
            _ = v36
        } else {
            for fi37 := 0; fi37 < 8; fi37++ {
                v10 = v10
                var v38 uint = 2
                if (v10 > "a") {
                    v39 := []int{}
                    var v40 uint16 = 32767
                    _ = v39
                    _ = v40
                } else {
                    var v41 uint = v38
                    v42 := GenS1{F0: v10, F1: 0}
                    _ = v41
                    _ = v42
                }
                _ = fi37
                _ = v38
            }
            fmt.Println(v7)
        }
        _ = fi9
        _ = v10
    }
    _ = v7
    _ = v8
    return uint8(v7)
}

func main() {
    var v1 uint32 = 2147483648
    var v2 int16 = 42
    var v3 uint16 = 1
    var v4 uint16 = 65535
    var v5 float64 = -2.25
    v6 := []string{"a"}
    v2 = (-(v2))
    _ = len(v6)
    var v43 int64 = 42
    var v44 uint32 = (^(v1))
    v44 = v44
    switch "\n\t" {
    case "a":
        v45 := GenS3{F0: -1, F1: 26}
        _ = v45
    }
    fmt.Println(v4)
    if (18446744073709551615 < 0) {
        v3 = (v4 ^ (v4 | v4))
        v46 := GenS3{F0: v43, F1: 1}
        switch ("" + "foo bar") {
        case "\n\t":
            v43 = (v43 / (((v43 / ((v43) | 1))) | 1))
        case "":
            switch 2147483647 {
            case -287:
                v3 += v3
            case -2147483648:
                if false {
                    switch 9223372036854775807 {
                    case 751:
                        for fi47 := 0; fi47 < 1; fi47++ {
                            v43 = (v43 >> ((0) & 63))
                            for fi48 := 0; fi48 < 8; fi48++ {
                                v44 -= v1
                                fmt.Println(fi47, v43)
                                _ = fi48
                            }
                            _ = fi47
                        }
                    case 996:
                        _ = len(v6)
                    default:
                        fmt.Println(v2)
                    }
                } else {
                    _ = len(v6)
                    fmt.Println(v1, v4, v44)
                }
            case 128:
                var v49 int = 42
                _ = v49
            default:
                _ = len(v6)
            }
        }
        _ = v46
    }
    if true {
        var v50 int8 = -85
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        switch 268 {
        case 752:
            v3 += v4
        }
        _ = v50
    } else {
        var v51 float32 = float32(v5)
        if true {
            v51 = 1.5
            if false {
                var v52 bool = false
                for fi53 := 0; fi53 < 3; fi53++ {
                    var v54 uint = 255
                    _ = fi53
                    _ = v54
                }
                _ = v52
            } else {
                v5 -= v5
                v2 += v2
                switch -6 {
                case 197:
                    v5 = (v5 + 1e10)
                default:
                    v2 -= -1
                }
            }
            var v55 int8 = 1
            _ = v55
        }
        v6 = append(v6, "hello")
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
