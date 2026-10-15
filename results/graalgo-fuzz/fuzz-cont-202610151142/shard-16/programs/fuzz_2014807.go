package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint; F2 string }
type GenS4 struct { F0 uint16; F1 string; F2 int64 }
type GenS5 struct { F0 uint8; F1 int64 }
type GenS6 struct { F0 bool }
func genFn0(p0 int, p1 uint, p2 int16) int {
    switch (("foo bar" + "foo bar") + "a") {
    case "hello":
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 float32 = (((float32(3.14159) + -2.25) + 1e10) / 1.5)
            var v9 int64 = 815
            _ = fi7
            _ = v8
            _ = v9
        }
    default:
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 uint = (p1 ^ 2)
            p2 += (-(p2))
            v12 := GenS1{F0: (fi10 / ((-700) | 1))}
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    v13 := GenS1{F0: (^(p0))}
    v14 := []int{p0, (p0 - (p0 - p0)), 32767, p0}
    var v15 uint64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    return (p0 | (p0 + (p0 % ((p0) | 1))))
}
func genFn1(p0 int16) int32 {
    for fi16 := 0; fi16 < 3; fi16++ {
        p0 += -32768
        _ = fi16
    }
    v17 := []int{9223372036854775807, 2147483647}
    var v18 string = ""
    v19 := GenS1{F0: -208}
    for fi20 := 0; fi20 < 15; fi20++ {
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        v17 = append(v17, (fi20 - (-(fi20))))
        v18 = (v18 + "")
        _ = fi20
    }
    _ = p0
    _ = v17
    _ = v18
    _ = v19
    return -767
}

func main() {
    v1 := make(map[int]int)
    var v2 int = -9223372036854775808
    v3 := []int{(-(v2)), (v2 >> ((2) & 63)), 127, -9223372036854775808}
    var v4 uint = 0
    v5 := make(map[uint32]int)
    var v6 float64 = (1e10 / 1e10)
    var v21 int64 = -1
    v3 = append(v3, 956)
    var v22 uint16 = 255
    var v23 bool = (false || (!true))
    v4 = 983
    switch "foo bar" {
    case "hello":
        v23 = ("foo bar" == "hello")
    default:
        v24 := v5[2147483648]
        _ = v24
    }
    fmt.Println(v2, v2)
    fmt.Println(v4)
    fmt.Println(v4)
    v25 := GenS2{F0: (v22 % ((uint16(v4)) | 1))}
    v26 := v5[2]
    if v23 {
        switch (("\n\t" + "hello") + "\n\t") {
        case "foo bar":
            v27 := GenS2{F0: v22}
            _ = v27
        case "a":
            v28 := GenS1{F0: v2}
            _ = v28
        case "":
            _ = len(v5)
        }
        for fi29 := 0; fi29 < 3; fi29++ {
            if true {
                fmt.Println(v4, v4)
                for fi30 := 0; fi30 < 5; fi30++ {
                    v31 := GenS2{F0: (v22 << ((v4) & 63))}
                    _ = fi30
                    _ = v31
                }
                var v32 uint64 = 18446744073709551615
                _ = v32
            }
            _ = fi29
        }
        v21 = v21
    } else {
        var v33 int64 = (-(v21))
        if (-62 > 100) {
            v6 = float64(1.5)
        } else {
            switch v2 {
            case 0:
                for fi34 := 0; fi34 < 16; fi34++ {
                    var v35 uint8 = 121
                    _ = fi34
                    _ = v35
                }
            }
            v33 = (-(v21))
        }
        _ = v33
    }
    if true {
        if v23 {
            var v36 int64 = (v21 - v21)
            var v37 int16 = 32767
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v36
            _ = v37
        } else {
            for fi38 := 0; fi38 < 4; fi38++ {
                if (false || v23) {
                    if ((-50 == 42) || v23) {
                        v39 := GenS3{F0: v26, F1: (v4 % ((65535) | 1)), F2: ""}
                        _ = len(v3)
                        _ = v39
                    } else {
                        if v23 {
                            for fi40 := 0; fi40 < 13; fi40++ {
                                v22 = v22
                                _ = fi40
                            }
                            if ((v21 | v21) < (v21 ^ (v21 << ((v4) & 63)))) {
                                var v41 int32 = 42
                                v22 = v22
                                v41 = genFn1(-32768)
                                _ = v41
                            }
                            var v42 uint16 = (v22 & 387)
                            _ = v42
                        }
                        switch (-(v2)) {
                        case 4294967295:
                            var v43 string = (("foo bar" + "hello") + "hello")
                            _ = v43
                        case -2147483648:
                            fmt.Println(v22, v26, v22)
                        case -396:
                            v2 = 32767
                        default:
                            var v44 uint16 = (v22 << ((924) & 63))
                            _ = v44
                        }
                    }
                    if (!(1.5 <= (1e10 / ((-2.25 / 1e10) + 1e10)))) {
                        v45 := GenS4{F0: 2, F1: "", F2: (v21 + (v21 >> ((v4) & 63)))}
                        _ = v45
                    } else {
                        v23 = true
                        v46 := GenS4{F0: (v22 ^ (v22 % (((v22 >> (((v4 | v4)) & 63))) | 1))), F1: "foo bar", F2: (^(v21))}
                        _ = v46
                    }
                }
                _ = fi38
            }
        }
        if v23 {
            if (!v23) {
                v47 := GenS5{F0: 0, F1: 0}
                var v48 int16 = -258
                _ = v47
                _ = v48
            } else {
                var v49 uint64 = 2
                fmt.Println(v6, v2, v49)
                _ = v49
            }
            if len(v3) > 0 {
                v3[0] = v26
            }
            fmt.Println(v2)
        } else {
            var v50 int32 = 2147483647
            _ = v50
        }
        v4 = 255
    } else {
        v22 = 643
        v5[2] = v2
    }
    v51 := GenS6{F0: (((v4 - v4) != (v4 / ((4294967295) | 1))) && (false || v23))}
    var v52 string = (("\n\t" + "") + ((("foo bar" + "foo bar") + (("hello" + (("\n\t" + ("\n\t" + "foo bar")) + "")) + ("a" + "a"))) + "hello"))
    var v53 float64 = float64(0.0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v6
}
