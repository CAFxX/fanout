package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int }
type GenS2 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16; F1 int64; F2 uint32 }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 string; F1 uint8 }
type GenS8 struct { F0 int8; F1 int32; F2 uint32 }
func genFn0(p0 uint16, p1 int16) uint16 {
    p1 -= (^(p1))
    v5 := GenS1{F0: p0, F1: 200, F2: -1}
    _ = p0
    _ = p1
    _ = v5
    return p0
}
func genFn1(p0 uint) uint32 {
    fmt.Println(p0)
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS2{F0: 1, F1: 2, F2: 460}
        _ = fi6
        _ = v7
    }
    _ = p0
    return 2147483647
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    v3 := []int{222}
    var v4 uint32 = 42
    if (103 <= 1) {
        if true {
            if true {
                if (3.14159 >= 0.0) {
                    v4 = v4
                    v4 -= v4
                    v8 := v2[42]
                    _ = v8
                }
                v9 := GenS2{F0: 0, F1: 1, F2: 731}
                var v10 uint8 = 0
                _ = v10
                _ = v9
            }
            var v11 int8 = -1
            v12 := GenS1{F0: 120, F1: 2, F2: -507}
            _ = v11
            _ = v12
        } else {
            v4 += 2
            fmt.Println(v4)
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v13 uint64 = 2
    v14 := GenS3{F0: 128, F1: ""}
    for fi15 := 0; fi15 < 2; fi15++ {
        _ = len(v1)
        for fi16 := 0; fi16 < 20; fi16++ {
            v13 -= v13
            _ = fi16
        }
        fmt.Println(fi15, fi15)
        _ = fi15
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        if (3.14159 >= 3.14159) {
            for fi18 := 0; fi18 < 20; fi18++ {
                if (18446744073709551615 == 0) {
                    v13 = (v13 / (((v13 >> ((65535) & 63))) | 1))
                } else {
                    v19 := GenS1{F0: 65535, F1: 2, F2: (-(fi18))}
                    _ = v19
                }
                switch "hello" {
                case "a":
                    var v20 uint16 = 0
                    _ = v20
                }
                var v21 bool = (v13 <= v13)
                var v22 int = -1
                _ = fi18
                _ = v21
                _ = v22
            }
            var v23 bool = true
            _ = v23
        } else {
            v24 := GenS3{F0: 127, F1: (("hello" + (((("\n\t" + "hello") + "a") + "a") + "foo bar")) + "\n\t")}
            var v25 int16 = 1000
            v13 -= (v13 << ((0) & 63))
            _ = v24
            _ = v25
        }
        _ = fi17
    }
    var v26 int = 4294967295
    v2[int(v13)] = (v26 & v26)
    if ("a" != "foo bar") {
        for fi27 := 0; fi27 < 20; fi27++ {
            if false {
                for fi28 := 0; fi28 < 11; fi28++ {
                    var v29 int32 = 42
                    v30 := GenS3{F0: (fi27 + 32767), F1: (("\n\t" + "a") + ("foo bar" + "foo bar"))}
                    switch (fi28 >> ((732) & 63)) {
                    case -973:
                        _ = len(v2)
                    case 42:
                        switch ("hello" + ("hello" + ("hello" + ("foo bar" + (("hello" + ("\n\t" + ("hello" + "foo bar"))) + ("foo bar" + (("a" + "a") + "\n\t"))))))) {
                        case "":
                            var v31 uint16 = 2
                            _ = v31
                        case "\n\t":
                            fmt.Println(v29, v4)
                        }
                    default:
                        v13 -= v13
                    }
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                v13 -= (v13 << ((uint(v4)) & 63))
                if (false || true) {
                    v13 = 2
                }
            } else {
                v4 -= v4
                if true {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
            var v32 int16 = -438
            v33 := GenS4{F0: 1}
            _ = fi27
            _ = v32
            _ = v33
        }
        v26 -= v26
    } else {
        v34 := GenS5{F0: 42, F1: 0, F2: v4}
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        if true {
            _ = len(v2)
            fmt.Println(v4)
            v3 = append(v3, (v26 / (((v26 * (v26 >> ((732) & 63)))) | 1)))
        } else {
            if (false && (false && false)) {
                v35 := GenS2{F0: 93, F1: 32767, F2: -1}
                v36 := GenS6{F0: 42, F1: 611}
                v13 += (v13 + 18446744073709551615)
                _ = v35
                _ = v36
            } else {
                if true {
                    if len(v3) > 0 {
                        v3[0] = (v26 << ((589) & 63))
                    }
                } else {
                    switch (-(v26)) {
                    case -514:
                        if false {
                            for fi37 := 0; fi37 < 9; fi37++ {
                                for fi38 := 0; fi38 < 12; fi38++ {
                                    var v39 float64 = ((((float64(3.14159) / ((3.14159 + 0.1) * float64((float32(-2.25) / float32((-2.25 / 1e10)))))) - 1e10) / 1e10) / -2.25)
                                    if false {
                                        if ((v4 * 42) > (v4 ^ v4)) {
                                            v4 -= 0
                                        }
                                    } else {
                                        v13 += 631
                                        var v40 int64 = 42
                                        v41 := GenS2{F0: 54, F1: 1, F2: (v40 | v40)}
                                        _ = v40
                                        _ = v41
                                    }
                                    _ = fi38
                                    _ = v39
                                }
                                v42 := GenS5{F0: 1, F1: 1, F2: 65535}
                                v13 -= (v13 * (v13 >> ((472) & 63)))
                                _ = fi37
                                _ = v42
                            }
                        } else {
                            var v43 int64 = 9223372036854775807
                            _ = v43
                        }
                    case 128:
                        v4 += v4
                    }
                }
                for fi44 := 0; fi44 < 5; fi44++ {
                    v26 = (v26 << ((32) & 63))
                    var v45 bool = (0 <= 126)
                    _ = fi44
                    _ = v45
                }
            }
            v46 := GenS7{F0: "hello", F1: 128}
            var v47 int32 = -1
            _ = v46
            _ = v47
        }
        _ = v34
    }
    v13 = 1
    var v48 float32 = 1.5
    switch "hello" {
    case "foo bar":
        v4 = v4
    }
    if (!(v13 > 407)) {
        _ = len(v1)
        v49 := GenS8{F0: -1, F1: 350, F2: (v4 << ((0) & 63))}
        _ = v49
    } else {
        var v50 uint8 = 153
        var v51 float64 = (3.14159 + 1e10)
        var v52 uint32 = (v4 ^ 1)
        _ = v50
        _ = v51
        _ = v52
    }
    v48 = float32(float64(((v48 - float32((((3.14159 / ((-2.25 * 1e10) * 0.1)) - ((1.5 * 0.0) / ((float64(v48) + 1e10) * 1.5))) - (0.1 + (3.14159 / 1e10))))) + v48)))
    v48 = (v48 / float32(((-2.25 + (((3.14159 * 1.5) - 3.14159) - 3.14159)) + 1.5)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v48
}
