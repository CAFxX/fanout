package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int16; F1 int; F2 bool }
type GenS6 struct { F0 uint64 }
func genFn0(p0 string, p1 int8) int8 {
    fmt.Println(p0, p0)
    for fi5 := 0; fi5 < 6; fi5++ {
        p0 = p0
        var v6 string = ((("a" + ("" + "foo bar")) + (p0 + "hello")) + p0)
        v7 := []int{}
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    return p1
}
func genFn1(p0 bool, p1 int) int64 {
    if p0 {
        var v8 int64 = 1
        _ = v8
    }
    p0 = p0
    var v9 uint32 = 2
    for fi10 := 0; fi10 < 9; fi10++ {
        fmt.Println(fi10, v9)
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 string = "a"
    var v2 int64 = 0
    var v3 float32 = (float32(1.5) * (3.14159 * (float32(0.0) + 3.14159)))
    var v4 int16 = -1
    v11 := GenS1{F0: 42, F1: 312}
    v12 := GenS1{F0: 42, F1: 4294967295}
    v13 := GenS2{F0: 997, F1: 0}
    var v14 int8 = 127
    var v15 int = 187
    for fi16 := 0; fi16 < 15; fi16++ {
        v17 := []int{}
        v18 := GenS1{F0: 0, F1: 4294967295}
        fmt.Println(v14, v4)
        _ = fi16
        _ = v17
        _ = v18
    }
    if true {
        for fi19 := 0; fi19 < 13; fi19++ {
            if (!false) {
                fmt.Println(v15, v4, v4)
                for fi20 := 0; fi20 < 1; fi20++ {
                    if ((v2 + (v2 + genFn1(false, int(v14)))) == -9223372036854775808) {
                        v3 = (((v3 * v3) * (1e10 / v3)) * float32((-2.25 - (float64(v3) + 1.5))))
                        v3 += 0.0
                        v21 := []int{fi20, v15, (fi19 ^ 32767)}
                        var v22 string = "\n\t"
                        _ = v21
                        _ = v22
                    } else {
                        v23 := GenS2{F0: 4294967295, F1: 2}
                        var v24 float64 = 1.5
                        switch (-(fi19)) {
                        case -1:
                            var v25 int8 = genFn0(("hello" + ("foo bar" + "foo bar")), -128)
                            _ = v25
                        case 589:
                            var v26 float32 = -2.25
                            _ = v26
                        case 127:
                            switch fi20 {
                            case 4294967295:
                                switch (v1 + v1) {
                                case "\n\t":
                                    v15 = (fi20 << ((65535) & 63))
                                case "a":
                                    if ((1 != 2) || (int16(v14) != (v4 >> ((4294967295) & 63)))) {
                                        fmt.Println(fi20, fi20, v14)
                                    }
                                case "foo bar":
                                    v27 := GenS3{F0: v1}
                                    _ = v27
                                default:
                                    switch "" {
                                    case "\n\t":
                                        var v28 uint = 1
                                        _ = v28
                                    }
                                }
                            case 2147483647:
                                var v29 int = (v15 ^ v15)
                                _ = v29
                            case 32768:
                                v30 := []int{fi20}
                                _ = v30
                            }
                        }
                        v31 := []int{(v15 & (fi20 | (fi20 / (((^(fi20))) | 1)))), 388, (v15 - 1), (-(fi19))}
                        var v32 int32 = 42
                        _ = v23
                        _ = v24
                        _ = v31
                        _ = v32
                    }
                    var v33 uint64 = 2
                    v33 = (v33 * 2)
                    var v34 uint = 18446744073709551615
                    v1 = (("foo bar" + v1) + v1)
                    switch (fi19 % ((fi20) | 1)) {
                    case 894:
                        v2 = v2
                    case 32767:
                        var v35 uint8 = 2
                        _ = v35
                    case 893:
                        switch (fi20 + fi20) {
                        case 396:
                            if (v4 >= -32768) {
                                v14 = (-(v14))
                                v36 := []int{}
                                _ = v36
                            } else {
                                v37 := []int{(v15 - -9223372036854775808), (fi20 / (((fi20 >> (((v34 | uint(v33))) & 63))) | 1))}
                                var v38 bool = (float32(((((1e10 + (float64(-2.25) + 0.1)) - float64((v3 * 1.5))) / 0.1) * (float64((float32(0.0) * v3)) - float64(v3)))) >= ((v3 * float32(1.5)) / float32(-2.25)))
                                v34 -= 65535
                                _ = v37
                                _ = v38
                            }
                            v39 := []int{v15, (v15 + (v15 + fi19))}
                            _ = v39
                        case 1:
                            v40 := GenS3{F0: (("a" + v1) + v1)}
                            v3 += 1e10
                            v2 = (v2 ^ -9223372036854775808)
                            v41 := GenS4{F0: 0}
                            fmt.Println(v3, v1, v34)
                            _ = v40
                            _ = v41
                        case 127:
                            v15 -= v15
                            v33 += 489
                            v2 += v2
                        }
                    default:
                        v34 += v34
                    }
                    _ = fi20
                    _ = v33
                    _ = v34
                }
                var v42 float32 = v3
                _ = v42
            } else {
                v43 := []int{fi19, (fi19 << ((433) & 63)), (fi19 >> ((42) & 63)), v15}
                _ = v43
            }
            if true {
                v44 := GenS5{F0: v4, F1: fi19, F2: (!(!true))}
                v3 -= float32(-2.25)
                for fi45 := 0; fi45 < 17; fi45++ {
                    var v46 int64 = int64(fi45)
                    v47 := []int{v15, 4294967295, (-(v15)), 178}
                    _ = fi45
                    _ = v46
                    _ = v47
                }
                _ = v44
            } else {
                v4 += -1
            }
            v14 -= (v14 % ((v14) | 1))
            _ = fi19
        }
        v48 := GenS4{F0: 1}
        v2 = 42
        _ = v48
    }
    for fi49 := 0; fi49 < 17; fi49++ {
        v15 = 138
        _ = fi49
    }
    var v50 uint16 = 32767
    fmt.Println(v14, v14)
    for fi51 := 0; fi51 < 12; fi51++ {
        v52 := []int{fi51, (v15 & (v15 - int(v50)))}
        if true {
            switch (-(fi51)) {
            case 258:
                for fi53 := 0; fi53 < 10; fi53++ {
                    v54 := GenS2{F0: uint(v50), F1: 785}
                    v2 = (^(v2))
                    fmt.Println(v15, v15)
                    fmt.Println(v1, v1, fi51)
                    _ = fi53
                    _ = v54
                }
            case 9223372036854775807:
                v14 -= 100
            default:
                var v55 int8 = genFn0("foo bar", int8(v2))
                _ = v55
            }
        } else {
            if ((v4 <= v4) || ((871 >= (v50 & v50)) && (!((-2.25 <= 3.14159) && false)))) {
                v15 = -1
                v56 := GenS6{F0: 42}
                v3 -= v3
                _ = v56
            } else {
                var v57 int16 = v4
                if true {
                    _ = len(v52)
                }
                _ = v57
            }
            var v58 int = 128
            _ = v58
        }
        _ = fi51
        _ = v52
    }
    v50 = (v50 * 42)
    for fi59 := 0; fi59 < 7; fi59++ {
        if true {
            v60 := []int{4294967295}
            _ = v60
        }
        _ = fi59
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v50
}
