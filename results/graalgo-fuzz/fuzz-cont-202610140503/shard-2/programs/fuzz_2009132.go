package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 int32; F1 int64; F2 int16 }
type GenS5 struct { F0 int16; F1 int }
func genFn0(p0 uint16) int8 {
    var v4 float32 = 0.1
    if true {
        if false {
            v5 := GenS1{F0: p0, F1: 2147483648, F2: (0 <= 127)}
            _ = v5
        }
        var v6 bool = false
        fmt.Println(p0, v4, p0)
        _ = v6
    }
    var v7 uint64 = 85
    v8 := []int{9223372036854775807, 32767}
    _ = p0
    _ = v4
    _ = v7
    _ = v8
    return -128
}
func genFn1(p0 uint64) bool {
    for fi9 := 0; fi9 < 18; fi9++ {
        if true {
            v10 := GenS1{F0: 255, F1: 2147483648, F2: (255 > 32767)}
            var v11 uint8 = 42
            _ = v10
            _ = v11
        }
        if ((1.5 >= (0.1 * 1e10)) && (0 < 255)) {
            for fi12 := 0; fi12 < 17; fi12++ {
                var v13 float32 = ((float32(float64(0.0)) + 0.0) - (3.14159 + (1.5 + 3.14159)))
                switch "\n\t" {
                case "\n\t":
                    var v14 uint64 = (p0 * p0)
                    _ = v14
                case "hello":
                    p0 = (p0 / ((42) | 1))
                default:
                    var v15 int = 32767
                    _ = v15
                }
                _ = fi12
                _ = v13
            }
            if true {
                v16 := GenS1{F0: 0, F1: 0, F2: (131 < 2)}
                p0 -= 1
                _ = v16
            } else {
                var v17 uint64 = p0
                fmt.Println(v17, fi9)
                _ = v17
            }
        } else {
            v18 := []int{(fi9 ^ (fi9 & 42)), (fi9 - (fi9 | (fi9 - fi9))), fi9, 9223372036854775807}
            _ = v18
        }
        _ = fi9
    }
    v19 := []int{2147483647, 0}
    for fi20 := 0; fi20 < 6; fi20++ {
        for fi21 := 0; fi21 < 13; fi21++ {
            if (4294967295 > 2147483648) {
                var v22 uint = 941
                var v23 uint16 = 998
                _ = v22
                _ = v23
            } else {
                switch "\n\t" {
                case "foo bar":
                    v24 := GenS1{F0: 0, F1: 4294967295, F2: ((false || (255 < 15)) || ((false && true) && (897 > 0)))}
                    _ = v24
                case "":
                    var v25 int = (fi20 ^ fi21)
                    _ = v25
                case "a":
                    var v26 int = (fi21 & (fi20 | -833))
                    _ = v26
                default:
                    if true {
                        p0 -= (p0 >> ((984) & 63))
                        p0 = 0
                    } else {
                        fmt.Println(fi20, fi21)
                        p0 = p0
                        fmt.Println(fi20, fi21)
                    }
                }
                fmt.Println(fi21, p0, fi20)
                if (true || ("a" > ("a" + "\n\t"))) {
                    fmt.Println(fi20)
                    if true {
                        p0 = (p0 / (((^(p0))) | 1))
                        switch (-(fi21)) {
                        case 936:
                            v27 := GenS2{F0: 32767}
                            _ = v27
                        case -9223372036854775808:
                            if (fi21 == fi21) {
                                p0 -= (p0 >> ((18446744073709551615) & 63))
                                if true {
                                    v28 := GenS3{F0: (("\n\t" + ("\n\t" + "a")) + (("" + ("" + "foo bar")) + "hello")), F1: 0}
                                    var v29 string = (("\n\t" + "hello") + ("\n\t" + "a"))
                                    fmt.Println(v29, fi20)
                                    _ = v28
                                    _ = v29
                                }
                                if (!(42 != 18446744073709551615)) {
                                    p0 += 2
                                    var v30 uint16 = 255
                                    var v31 float64 = 1e10
                                    fmt.Println(fi21)
                                    _ = v30
                                    _ = v31
                                }
                            } else {
                                p0 += p0
                                p0 += (p0 | (p0 % ((0) | 1)))
                                if true {
                                    p0 += (^(p0))
                                    p0 += p0
                                    v32 := GenS4{F0: -574, F1: 42, F2: -963}
                                    fmt.Println(fi20, p0, p0)
                                    if (fi20 != fi20) {
                                        if len(v19) >= 2 {
                                            v19 = v19[1:len(v19)]
                                        }
                                        v19 = append(v19, -651)
                                        p0 -= 125
                                    } else {
                                        v33 := GenS5{F0: 1000, F1: fi20}
                                        var v34 uint = 65535
                                        _ = v33
                                        _ = v34
                                    }
                                    fmt.Println(fi21, fi20)
                                    v35 := GenS4{F0: 42, F1: -9223372036854775808, F2: -404}
                                    _ = v32
                                    _ = v35
                                }
                            }
                        case -341:
                            if len(v19) > 0 {
                                v19[0] = fi21
                            }
                        }
                        p0 = p0
                    }
                    var v36 uint = 696
                    _ = v36
                } else {
                    _ = len(v19)
                    var v37 uint16 = 848
                    fmt.Println(fi20, v37)
                    _ = v37
                }
            }
            _ = fi21
        }
        _ = fi20
    }
    v38 := GenS1{F0: 32767, F1: 2147483647, F2: (true || (1 < 0))}
    p0 += (p0 - (p0 % (((p0 - 317)) | 1)))
    _ = p0
    _ = v19
    _ = v38
    return (32767 >= 32768)
}

func main() {
    var v1 int16 = 91
    var v2 uint16 = 255
    var v3 int = 0
    var v39 float32 = 0.1
    var v40 uint = 2
    for fi41 := 0; fi41 < 13; fi41++ {
        var v42 int16 = v1
        v43 := []int{fi41}
        var v44 uint16 = 1
        _ = fi41
        _ = v42
        _ = v43
        _ = v44
    }
    switch (^(v3)) {
    case -2147483648:
        v45 := GenS3{F0: ("" + ("a" + "")), F1: -566}
        _ = v45
    case 886:
        v46 := []int{v3, v3, v3}
        _ = v46
    case 0:
        var v47 int16 = 0
        _ = v47
    }
    var v48 int = v3
    var v49 int = v48
    for fi50 := 0; fi50 < 5; fi50++ {
        v40 = v40
        var v51 float32 = v39
        _ = fi50
        _ = v51
    }
    if false {
        v52 := []int{(v49 ^ v48), 2, (v3 - (v3 / (((v3 - (v3 % (((v49 / ((v48) | 1))) | 1)))) | 1))), int(v2)}
        v39 += (1e10 * v39)
        _ = v52
    }
    var v53 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v48
    _ = v49
    _ = v53
}
