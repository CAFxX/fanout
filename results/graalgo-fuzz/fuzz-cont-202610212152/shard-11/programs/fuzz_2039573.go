package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int; F2 uint }
type GenS4 struct { F0 int64; F1 uint32 }
func genFn0(p0 int64) uint32 {
    p0 = -490
    var v7 uint = 18446744073709551615
    switch -1 {
    case -653:
        v7 -= (v7 * 4294967295)
    case -918:
        var v8 int8 = 127
        _ = v8
    default:
        if false {
            var v9 uint = (v7 ^ v7)
            _ = v9
        }
    }
    if (2 <= 2) {
        if (-2147483648 > 127) {
            v10 := []int{42, -9223372036854775808}
            v11 := GenS1{F0: 754, F1: 141}
            _ = v10
            _ = v11
        } else {
            v12 := []int{}
            _ = v12
        }
    } else {
        if true {
            var v13 int32 = 2147483647
            _ = v13
        }
        if false {
            var v14 int16 = 876
            var v15 uint32 = 42
            var v16 int64 = (p0 | 1)
            _ = v14
            _ = v15
            _ = v16
        }
    }
    _ = p0
    _ = v7
    return 2
}
func genFn1(p0 int32, p1 uint16, p2 int16) int {
    v17 := []int{128, 1, -474}
    if true {
        for fi18 := 0; fi18 < 19; fi18++ {
            for fi19 := 0; fi19 < 16; fi19++ {
                var v20 bool = false
                if v20 {
                    if v20 {
                        v21 := GenS2{F0: (fi19 > fi18)}
                        _ = v21
                    }
                } else {
                    if len(v17) > 0 {
                        v17[0] = -2147483648
                    }
                    if true {
                        v20 = ((fi19 % ((fi18) | 1)) != fi18)
                        if (((42 == 1) || v20) || false) {
                            var v22 uint32 = 99
                            switch int(p2) {
                            case 128:
                                fmt.Println(v22)
                            case 201:
                                v22 -= (v22 / ((v22) | 1))
                            default:
                                var v23 int64 = 0
                                _ = v23
                            }
                            var v24 uint16 = (p1 << ((2) & 63))
                            v20 = true
                            _ = v22
                            _ = v24
                        }
                        if false {
                            var v25 uint = 655
                            fmt.Println(fi18)
                            _ = v25
                        }
                    } else {
                        _ = len(v17)
                    }
                    p1 -= (p1 - 677)
                    if ((!((p1 + 1) != 42)) && false) {
                        _ = len(v17)
                        var v26 int = fi18
                        v27 := GenS3{F0: 4294967295, F1: fi18, F2: 704}
                        _ = v26
                        _ = v27
                    }
                    p1 -= p1
                }
                _ = len(v17)
                _ = fi19
                _ = v20
            }
            for fi28 := 0; fi28 < 9; fi28++ {
                p0 -= p0
                var v29 string = ("\n\t" + "")
                var v30 int32 = 91
                p0 = (v30 / (((p0 - 976)) | 1))
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = fi18
        }
        var v31 float64 = (1.5 + 0.1)
        p0 += 42
        _ = v31
    } else {
        if ((p1 % ((765) | 1)) <= (^(p1))) {
            if true {
                var v32 int64 = -483
                v33 := GenS2{F0: true}
                p0 = p0
                _ = v32
                _ = v33
            }
            p2 -= (-(p2))
        }
        for fi34 := 0; fi34 < 7; fi34++ {
            p0 = (^(p0))
            _ = fi34
        }
        p2 = p2
    }
    p0 += (p0 | 1)
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return -2147483648
}

func main() {
    var v1 bool = (!false)
    var v2 uint8 = 1
    var v3 bool = (!(!(1.5 == 1e10)))
    v4 := make(map[string]string)
    var v5 uint32 = 1
    var v6 uint32 = (v5 >> ((18446744073709551615) & 63))
    v5 = (v5 + 65535)
    v3 = (0.1 == 0.1)
    switch "foo bar" {
    case "":
        v35 := GenS3{F0: v5, F1: 1, F2: 2}
        _ = v35
    }
    v36 := GenS3{F0: (v5 * v5), F1: 183, F2: 768}
    v37 := []int{9223372036854775807, 32768, 9223372036854775807}
    v3 = (v2 <= v2)
    _ = len(v37)
    fmt.Println(v3, v2)
    v38 := GenS4{F0: -1, F1: 2147483647}
    v39 := GenS2{F0: (!(42 >= -15))}
    var v40 int = 127
    if (genFn0(9223372036854775807) != 4294967295) {
        for fi41 := 0; fi41 < 17; fi41++ {
            v42 := GenS4{F0: -380, F1: v6}
            _ = fi41
            _ = v42
        }
        switch "a" {
        case "\n\t":
            var v43 bool = (0.0 == 1e10)
            _ = v43
        }
        var v44 float32 = -2.25
        _ = v44
    } else {
        switch "foo bar" {
        case "\n\t":
            if false {
                fmt.Println(v40, v1)
                var v45 float64 = 0.0
                switch genFn1(-1, 255, -1) {
                case 42:
                    var v46 uint16 = 1
                    _ = v46
                }
                _ = v45
            } else {
                var v47 uint32 = genFn0(-1)
                v5 = (v6 - v6)
                var v48 int64 = 0
                _ = v47
                _ = v48
            }
        default:
            delete(v4, "")
        }
    }
    v49 := GenS2{F0: ((!((4294967295 == 1) || (1 == v2))) && (uint(v2) != 209))}
    v6 += v5
    v5 = 1
    v6 -= (v6 - (v5 % ((genFn0(847)) | 1)))
    switch ("a" + (("\n\t" + ("\n\t" + "hello")) + ((("foo bar" + "a") + "a") + ("hello" + (("hello" + "a") + "hello"))))) {
    case "":
        var v50 int32 = 1
        _ = v50
    }
    v2 = 175
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v49
    _ = v5
    _ = v6
}
