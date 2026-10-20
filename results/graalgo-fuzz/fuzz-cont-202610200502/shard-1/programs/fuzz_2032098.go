package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint; F2 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 uint32; F1 uint; F2 int }
type GenS6 struct { F0 uint8; F1 string; F2 uint32 }
type GenS7 struct { F0 int8; F1 int16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint16 }
func genFn0(p0 int64) int32 {
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 bool = ((!true) && false)
            fmt.Println(p0, p0)
            var v10 float64 = -2.25
            if (!true) {
                v11 := GenS1{F0: p0}
                v12 := []int{fi8}
                _ = v11
                _ = v12
            }
            v13 := GenS2{F0: 18446744073709551615, F1: 250, F2: (!((!(361 >= 0)) && v9))}
            var v14 uint32 = 2
            v15 := []int{fi7, 127, fi8, 9223372036854775807}
            v10 += 1.5
            if v9 {
                if false {
                    v10 += ((((v10 - v10) / 3.14159) * v10) + v10)
                    v14 = v14
                    switch fi7 {
                    case 852:
                        v10 = v10
                    case -2147483648:
                        v16 := GenS2{F0: 65535, F1: 0, F2: (-1 > -2147483648)}
                        _ = v16
                    }
                } else {
                    v17 := GenS3{F0: 70}
                    v18 := GenS3{F0: 126}
                    _ = v17
                    _ = v18
                }
                switch (^(fi7)) {
                case 531:
                    fmt.Println(v14, fi7)
                case 467:
                    if v9 {
                        fmt.Println(v14, p0)
                        var v19 int32 = -1
                        var v20 bool = ((^(p0)) >= (p0 / (((p0 + p0)) | 1)))
                        _ = v19
                        _ = v20
                    }
                case -2147483648:
                    if true {
                        var v21 bool = (v9 && v9)
                        _ = v21
                    } else {
                        var v22 uint32 = (^(v14))
                        var v23 uint8 = 1
                        v15 = append(v15, int(v23))
                        if ((p0 * p0) != (p0 + (p0 / ((int64(v23)) | 1)))) {
                            switch fi8 {
                            case 128:
                                switch (fi7 % (((fi8 << ((65535) & 63))) | 1)) {
                                case 42:
                                    v15 = append(v15, 483)
                                case 128:
                                    var v24 uint = 0
                                    _ = v24
                                case -501:
                                    _ = len(v15)
                                default:
                                    v25 := GenS2{F0: 42, F1: 2, F2: (!(1 == 2))}
                                    _ = v25
                                }
                            case 1:
                                v26 := GenS1{F0: (-(p0))}
                                _ = v26
                            default:
                                if ("a" <= ("\n\t" + "a")) {
                                    v15 = append(v15, (-(fi7)))
                                    var v27 int64 = (-(p0))
                                    v10 = float64((0.1 * (float32(v10) - (1.5 * ((((-2.25 + -2.25) - 0.0) * -2.25) / 3.14159)))))
                                    _ = v27
                                }
                            }
                            v23 -= v23
                            v23 -= v23
                            v9 = (v9 && v9)
                        } else {
                            switch fi7 {
                            case 128:
                                if len(v15) >= 2 {
                                    v15 = v15[1:len(v15)]
                                }
                            default:
                                var v28 int64 = (-(p0))
                                _ = v28
                            }
                            v29 := GenS4{F0: ("foo bar" + (((("foo bar" + ((("hello" + "foo bar") + "hello") + "")) + "") + "foo bar") + "\n\t")), F1: 4294967295}
                            var v30 int = fi7
                            _ = v29
                            _ = v30
                        }
                        v14 -= (v14 << ((255) & 63))
                        _ = v22
                        _ = v23
                    }
                }
            }
            v31 := GenS5{F0: (v14 << ((2) & 63)), F1: 1, F2: 128}
            _ = fi8
            _ = v10
            _ = v13
            _ = v14
            _ = v15
            _ = v31
            _ = v9
        }
        p0 = int64(fi7)
        var v32 uint16 = 2
        _ = fi7
        _ = v32
    }
    var v33 uint = 1
    for fi34 := 0; fi34 < 8; fi34++ {
        p0 += (p0 << ((18446744073709551615) & 63))
        v33 = v33
        for fi35 := 0; fi35 < 2; fi35++ {
            var v36 uint16 = 32767
            switch (fi35 >> (((v33 ^ v33)) & 63)) {
            case -497:
                if true {
                    v36 = (v36 ^ (v36 / ((v36) | 1)))
                    p0 -= 992
                } else {
                    v37 := []int{32768, fi34, (fi34 & (fi35 / (((fi35 | fi35)) | 1)))}
                    v38 := GenS6{F0: 127, F1: "\n\t", F2: 751}
                    if true {
                        var v39 int16 = 42
                        _ = v39
                    }
                    _ = v37
                    _ = v38
                }
            }
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    _ = p0
    _ = v33
    return 46
}
func genFn1(p0 bool, p1 uint64) int16 {
    var v40 int8 = 1
    var v41 int = 4294967295
    fmt.Println(v41, p1, v40)
    _ = p0
    _ = p1
    _ = v40
    _ = v41
    return -998
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint8 = 128
    v3 := make(map[uint32]int)
    var v4 int = 0
    v5 := []string{("\n\t" + "\n\t"), "\n\t"}
    v6 := []string{"a"}
    switch (v4 << ((18446744073709551615) & 63)) {
    case 553:
        fmt.Println(v1, v1)
    case 432:
        if (851 >= genFn0(v1)) {
            fmt.Println(v4)
            var v42 int32 = 0
            _ = v42
        } else {
            if true {
                v43 := GenS2{F0: 148, F1: 4294967295, F2: ((v2 << ((976) & 63)) < (v2 / ((v2) | 1)))}
                fmt.Println(v4, v1)
                v5 = append(v5, "")
                _ = v43
            } else {
                fmt.Println(v2)
                fmt.Println(v4)
                v44 := GenS7{F0: -120, F1: 0}
                _ = v44
            }
        }
    default:
        v5 = append(v5, "\n\t")
    }
    v5 = append(v5, ("hello" + "hello"))
    delete(v3, 245)
    var v45 string = "a"
    v45 = v45
    v1 = v1
    var v46 int64 = v1
    if len(v5) > 0 {
        v5[0] = "foo bar"
    }
    v47 := GenS1{F0: v46}
    if (!(!((false && true) || ((v4 - v4) >= v4)))) {
        v48 := GenS6{F0: 255, F1: v45, F2: 42}
        v45 = v45
        _ = v48
    }
    fmt.Println(v4, v45, v45)
    v49 := GenS8{F0: -1}
    if len(v6) > 0 {
        v6[0] = ((v45 + v45) + v45)
    }
    v46 = (v46 << ((0) & 63))
    if ((3.14159 * 1.5) >= 1e10) {
        v50 := GenS3{F0: -1}
        v51 := GenS9{F0: 900}
        var v52 int64 = (v46 % (((^(v46))) | 1))
        _ = v50
        _ = v51
        _ = v52
    } else {
        v6 = append(v6, v45)
        v6 = append(v6, "foo bar")
        var v53 int16 = genFn1(((1.5 - float32(1e10)) <= 1e10), 975)
        _ = v53
    }
    var v54 uint64 = 510
    var v55 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v49
    _ = v5
    _ = v54
    _ = v55
    _ = v6
}
