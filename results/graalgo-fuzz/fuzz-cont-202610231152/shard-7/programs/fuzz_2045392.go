package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int; F1 int8; F2 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int32; F1 int32; F2 uint16 }
type GenS8 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS9 struct { F0 int16 }
func genFn0(p0 uint16, p1 int8) int {
    if false {
        var v7 int = -895
        fmt.Println(v7)
        _ = v7
    } else {
        var v8 string = "a"
        if ("foo bar" < (v8 + "hello")) {
            fmt.Println(p1)
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 int8 = p1
                if (v8 <= "a") {
                    fmt.Println(v8, p1)
                    v10 = (v10 / (((p1 & (p1 * (v10 / (((p1 | (p1 % ((10) | 1)))) | 1))))) | 1))
                    v10 = p1
                } else {
                    if (0 == 42) {
                        if false {
                            var v11 int64 = -440
                            if false {
                                v12 := GenS1{F0: 42, F1: 42}
                                _ = v12
                            }
                            p0 = 602
                            _ = v11
                        }
                    } else {
                        var v13 int = -2147483648
                        _ = v13
                    }
                }
                _ = fi9
                _ = v10
            }
        } else {
            var v14 bool = (true || true)
            var v15 int16 = -32768
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 10; fi16++ {
            if (1.5 == float32(((0.0 / 1e10) * 0.0))) {
                for fi17 := 0; fi17 < 8; fi17++ {
                    v18 := GenS1{F0: 0, F1: 312}
                    var v19 float64 = 0.0
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                if (18446744073709551615 > 0) {
                    v20 := GenS1{F0: 183, F1: -2147483648}
                    p0 = 420
                    _ = v20
                } else {
                    v8 = v8
                }
            }
            var v21 int16 = -32768
            _ = fi16
            _ = v21
        }
        _ = v8
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        for fi23 := 0; fi23 < 7; fi23++ {
            v24 := []int{fi22}
            v25 := GenS2{F0: (fi23 | -502), F1: (p1 >> ((4294967295) & 63)), F2: fi23}
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi22
    }
    if (42 == 0) {
        var v26 uint32 = 42
        var v27 uint8 = 2
        for fi28 := 0; fi28 < 19; fi28++ {
            var v29 float32 = (-2.25 / 1e10)
            var v30 int = (fi28 >> ((uint(v27)) & 63))
            for fi31 := 0; fi31 < 4; fi31++ {
                fmt.Println(p0)
                if true {
                    v32 := GenS3{F0: (p0 % (((p0 - p0)) | 1))}
                    fmt.Println(p1)
                    switch fi31 {
                    case -9223372036854775808:
                        var v33 int = (fi28 - (fi28 | 9223372036854775807))
                        _ = v33
                    case 2147483647:
                        switch ("a" + "") {
                        case "hello":
                            var v34 uint8 = v27
                            _ = v34
                        case "\n\t":
                            var v35 uint16 = 259
                            _ = v35
                        }
                    case 127:
                        if true {
                            fmt.Println(v29)
                            if true {
                                v36 := []int{127, fi28, (fi28 % ((fi31) | 1)), 114}
                                _ = v36
                            } else {
                                var v37 bool = true
                                _ = v37
                            }
                        } else {
                            v38 := GenS4{F0: -879, F1: 1000}
                            _ = v38
                        }
                    }
                    var v39 string = "hello"
                    var v40 string = v39
                    _ = v32
                    _ = v39
                    _ = v40
                }
                v41 := GenS5{F0: 18446744073709551615}
                _ = fi31
                _ = v41
            }
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = v26
        _ = v27
    }
    p0 = p0
    _ = p0
    _ = p1
    return 972
}
func genFn1() uint64 {
    var v42 float64 = 0.1
    v43 := []int{}
    var v44 int8 = 0
    if true {
        var v45 int16 = 0
        _ = v45
    } else {
        v44 -= v44
        for fi46 := 0; fi46 < 15; fi46++ {
            fmt.Println(v44)
            var v47 string = "foo bar"
            _ = fi46
            _ = v47
        }
    }
    for fi48 := 0; fi48 < 7; fi48++ {
        var v49 float64 = v42
        var v50 uint8 = 127
        switch (fi48 & (fi48 * 838)) {
        case 0:
            fmt.Println(v50)
        case 4294967295:
            v51 := GenS6{F0: 42}
            _ = v51
        case 2147483647:
            v49 = 0.1
        }
        _ = fi48
        _ = v49
        _ = v50
    }
    _ = v42
    _ = v43
    _ = v44
    return 44
}

func main() {
    var v1 uint = 257
    var v2 bool = (false || ((222 >= 32768) || (false && true)))
    var v3 uint = 65535
    v4 := make(map[uint32]int)
    var v5 int8 = 98
    var v6 float64 = ((0.1 * (1.5 + (0.1 / 1.0))) * 0.1)
    v3 += (v1 * (v1 + v1))
    v52 := GenS7{F0: 0, F1: -1, F2: 42}
    if (42 >= 0) {
        var v53 uint32 = 0
        if (((v53 + (^(v53))) != (v53 & v53)) && (!((2147483647 != 1) && false))) {
            var v54 string = ("a" + "foo bar")
            _ = v54
        } else {
            var v55 float32 = 0.0
            if (true && (32767 != 44)) {
                v56 := []int{4294967295, 2147483647, -817, 32768}
                if false {
                    _ = len(v56)
                    fmt.Println(v6, v53)
                }
                var v57 int8 = (v5 / (((^(v5))) | 1))
                _ = v56
                _ = v57
            }
            v58 := GenS8{F0: 963, F1: 719, F2: 2}
            _ = v55
            _ = v58
        }
        for fi59 := 0; fi59 < 10; fi59++ {
            v4[v53] = (fi59 / ((genFn0(1, v5)) | 1))
            _ = fi59
        }
        _ = v53
    }
    v6 -= (1.5 - ((((v6 - ((v6 - v6) / 1.5)) * 0.1) - v6) * -2.25))
    for fi60 := 0; fi60 < 14; fi60++ {
        v61 := GenS9{F0: 48}
        v62 := []int{genFn0(255, v5), (fi60 * (fi60 + (^(fi60)))), (fi60 << ((498) & 63)), 1}
        _ = fi60
        _ = v61
        _ = v62
    }
    v3 = (v3 ^ (v3 ^ v3))
    var v63 float64 = 1e10
    fmt.Println(v3, v63, v5)
    var v64 uint16 = 1
    var v65 bool = (2147483647 > -1)
    fmt.Println(v2, v64, v6)
    v65 = false
    v66 := []int{}
    for fi67 := 0; fi67 < 9; fi67++ {
        for fi68 := 0; fi68 < 20; fi68++ {
            v6 += v63
            if true {
                if true {
                    fmt.Println(v6, fi68)
                } else {
                    v64 = (^(v64))
                    var v69 int16 = 42
                    _ = v69
                }
            } else {
                v66 = append(v66, fi68)
            }
            _ = fi68
        }
        v5 -= -43
        switch "hello" {
        case "a":
            v4[344] = 32767
        case "foo bar":
            fmt.Println(v65)
        }
        _ = fi67
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("slicelen v66:", len(v66), v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
    _ = v63
    _ = v64
    _ = v65
    _ = v66
}
