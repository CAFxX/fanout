package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 uint8 = 150
    var v2 uint8 = v1
    var v3 int32 = 2147483647
    v1 -= v2
    if true {
        for fi4 := 0; fi4 < 1; fi4++ {
            v2 = v1
            v2 = v1
            _ = fi4
        }
        v2 = (v2 << ((2) & 63))
    }
    if (v1 != v1) {
        fmt.Println(v3, v2, v1)
        v5 := GenS1{F0: (v1 | v2), F1: 0, F2: 600}
        _ = v5
    }
    v6 := GenS1{F0: (v1 * (v1 / ((v2) | 1))), F1: 65535, F2: 1}
    v2 = (v2 + (v2 - v1))
    var v7 uint8 = (v1 & v1)
    fmt.Println(v2)
    if true {
        v8 := GenS1{F0: 214, F1: 1, F2: 1}
        switch 66 {
        case -2147483648:
            fmt.Println(v7, v2, v2)
        }
        if true {
            for fi9 := 0; fi9 < 19; fi9++ {
                v10 := GenS1{F0: v2, F1: 255, F2: 0}
                _ = fi9
                _ = v10
            }
            var v11 float64 = float64(1e10)
            var v12 uint32 = 1
            _ = v11
            _ = v12
        } else {
            fmt.Println(v3)
        }
        _ = v8
    } else {
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 float32 = (((-2.25 - 0.0) - 1e10) + ((0.1 / float32(3.14159)) / float32(float64((-2.25 - 0.0)))))
            var v15 int = fi13
            v1 += 255
            _ = fi13
            _ = v14
            _ = v15
        }
        if ((!(1e10 <= 0.1)) && (9 < 47)) {
            fmt.Println(v3)
            switch -55 {
            case 42:
                var v16 uint16 = 65535
                _ = v16
            case -9223372036854775808:
                for fi17 := 0; fi17 < 11; fi17++ {
                    v18 := GenS1{F0: 2, F1: 32767, F2: 1}
                    _ = fi17
                    _ = v18
                }
            default:
                v19 := []int{1, 2147483647}
                _ = v19
            }
            for fi20 := 0; fi20 < 4; fi20++ {
                fmt.Println(v3, v1)
                for fi21 := 0; fi21 < 9; fi21++ {
                    v2 -= (v7 ^ 39)
                    v22 := GenS1{F0: (v2 ^ (v1 & 127)), F1: 255, F2: 18446744073709551615}
                    _ = fi21
                    _ = v22
                }
                _ = fi20
            }
        }
    }
    if (0.0 >= 1.5) {
        if (!("" >= "")) {
            v7 = 128
            v23 := GenS1{F0: (v7 | (v7 | v2)), F1: 107, F2: 0}
            switch ("hello" + ("" + "\n\t")) {
            case "a":
                var v24 float32 = (3.14159 + 0.0)
                _ = v24
            default:
                var v25 int64 = 9223372036854775807
                _ = v25
            }
            _ = v23
        } else {
            for fi26 := 0; fi26 < 15; fi26++ {
                if true {
                    if (-1 < 1) {
                        for fi27 := 0; fi27 < 14; fi27++ {
                            v28 := GenS1{F0: v2, F1: 2, F2: 4294967295}
                            v29 := GenS1{F0: v2, F1: 0, F2: 0}
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                    } else {
                        v30 := GenS2{F0: (fi26 % (((fi26 << ((17) & 63))) | 1))}
                        var v31 uint8 = (^(v2))
                        _ = v30
                        _ = v31
                    }
                    var v32 int32 = 2147483647
                    _ = v32
                }
                if true {
                    if true {
                        v7 += (^(v2))
                        v3 -= 643
                    } else {
                        v33 := []int{(fi26 >> ((600) & 63)), (fi26 | fi26), (fi26 + 4294967295)}
                        _ = v33
                    }
                    if false {
                        for fi34 := 0; fi34 < 14; fi34++ {
                            var v35 uint8 = v2
                            var v36 uint16 = 255
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                        var v37 int32 = (v3 | v3)
                        v38 := []int{(-(fi26)), 42, (fi26 ^ (fi26 >> ((0) & 63)))}
                        _ = v37
                        _ = v38
                    }
                } else {
                    for fi39 := 0; fi39 < 17; fi39++ {
                        if false {
                            if (!true) {
                                v7 = v1
                            } else {
                                v1 = v2
                                v40 := []int{(fi39 - (fi26 & (fi39 / ((fi39) | 1)))), fi26}
                                if (1.5 > 1e10) {
                                    if (("\n\t" + "hello") == "\n\t") {
                                        var v41 string = "hello"
                                        v42 := GenS1{F0: (v1 << ((908) & 63)), F1: 943, F2: 860}
                                        var v43 uint8 = 255
                                        _ = v41
                                        _ = v42
                                        _ = v43
                                    } else {
                                        var v44 int64 = 1
                                        if len(v40) >= 2 {
                                            v40 = v40[1:len(v40)]
                                        }
                                        _ = len(v40)
                                        v40 = append(v40, fi39)
                                        _ = v44
                                    }
                                }
                                _ = v40
                            }
                        }
                        _ = fi39
                    }
                }
                _ = fi26
            }
            v1 += 40
            for fi45 := 0; fi45 < 4; fi45++ {
                v46 := GenS3{F0: (v1 << ((42) & 63))}
                v47 := []int{1, 181, (fi45 | (^(fi45)))}
                v1 = v1
                _ = fi45
                _ = v46
                _ = v47
            }
        }
    } else {
        if false {
            v48 := []int{9223372036854775807, -862, -1}
            _ = v48
        } else {
            var v49 uint16 = 1
            _ = v49
        }
    }
    for fi50 := 0; fi50 < 8; fi50++ {
        var v51 bool = true
        v52 := []int{fi50, -632, (fi50 & fi50)}
        _ = fi50
        _ = v51
        _ = v52
    }
    fmt.Println(v1, v7, v7)
    var v53 float64 = 3.14159
    if (false && (1000 <= 1000)) {
        v53 = ((v53 / (0.1 / 1.5)) * v53)
    } else {
        var v54 int32 = v3
        v53 = float64(3.14159)
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v53:", v53)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v53
    _ = v6
    _ = v7
}
