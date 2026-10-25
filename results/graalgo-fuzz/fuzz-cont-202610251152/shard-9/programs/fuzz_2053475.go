package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int32; F1 uint32; F2 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 string; F1 uint16; F2 bool }
type GenS6 struct { F0 int }

func main() {
    var v1 uint = 315
    v2 := []string{"foo bar", ""}
    v3 := []int{32767}
    var v4 int32 = -563
    var v5 int8 = 1
    var v6 bool = (((!true) || true) || (190 < 0))
    if false {
        for fi7 := 0; fi7 < 7; fi7++ {
            v4 = (v4 << ((0) & 63))
            var v8 float64 = float64(3.14159)
            fmt.Println(fi7, v5)
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: -32768, F1: -1}
        for fi10 := 0; fi10 < 14; fi10++ {
            if v6 {
                var v11 int64 = -40
                var v12 int16 = 1000
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v11
                _ = v12
            } else {
                var v13 bool = (-515 != 0)
                if (!(65535 >= 2)) {
                    for fi14 := 0; fi14 < 16; fi14++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi14
                    }
                    v15 := GenS2{F0: (v4 ^ -1), F1: 119, F2: (v4 ^ (^(v4)))}
                    var v16 bool = v6
                    _ = v15
                    _ = v16
                } else {
                    switch (fi10 + 127) {
                    case -321:
                        v6 = v13
                    case 1:
                        fmt.Println(v4, v13, v1)
                    default:
                        v5 -= v5
                    }
                    var v17 int16 = 32767
                    var v18 float32 = float32((0.0 + 1e10))
                    _ = v17
                    _ = v18
                }
                _ = v13
            }
            v19 := GenS2{F0: v4, F1: 731, F2: 0}
            _ = fi10
            _ = v19
        }
        _ = v9
    } else {
        for fi20 := 0; fi20 < 1; fi20++ {
            v6 = (true || (v6 && (uint8(v1) > 255)))
            for fi21 := 0; fi21 < 9; fi21++ {
                var v22 uint8 = 127
                v5 -= (v5 | int8(v22))
                v23 := GenS3{F0: v6}
                var v24 string = "a"
                fmt.Println(v1)
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            fmt.Println(fi20, v4)
            _ = fi20
        }
        var v25 float64 = (0.0 * 0.1)
        if false {
            for fi26 := 0; fi26 < 9; fi26++ {
                v4 += v4
                _ = fi26
            }
            var v27 bool = ((0.0 / (1e10 + (1.5 * 0.0))) > v25)
            _ = v27
        } else {
            fmt.Println(v5)
            v4 -= 2147483647
            if (!(!v6)) {
                for fi28 := 0; fi28 < 14; fi28++ {
                    if v6 {
                        v29 := GenS4{F0: (v5 & v5), F1: "foo bar"}
                        _ = v29
                    }
                    var v30 uint8 = 42
                    _ = fi28
                    _ = v30
                }
                v4 = v4
                _ = len(v3)
            }
        }
        _ = v25
    }
    for fi31 := 0; fi31 < 7; fi31++ {
        if v6 {
            if v6 {
                v6 = (1 >= 65535)
                var v32 uint16 = 0
                var v33 float64 = 0.0
                _ = v32
                _ = v33
            } else {
                fmt.Println(v5)
                if (!(-32768 < 0)) {
                    if ((fi31 % ((fi31) | 1)) != 383) {
                        for fi34 := 0; fi34 < 14; fi34++ {
                            fmt.Println(fi34, v4, fi34)
                            v6 = false
                            _ = fi34
                        }
                    } else {
                        v35 := GenS3{F0: false}
                        var v36 string = ("foo bar" + "foo bar")
                        v36 = (v36 + v36)
                        _ = v35
                        _ = v36
                    }
                    if v6 {
                        if len(v3) > 0 {
                            v3[0] = 498
                        }
                    }
                }
                for fi37 := 0; fi37 < 6; fi37++ {
                    v3 = append(v3, int(v1))
                    _ = fi37
                }
            }
            v1 += (v1 % ((v1) | 1))
        }
        var v38 uint64 = 2
        v6 = v6
        _ = fi31
        _ = v38
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v39 int16 = 42
    v4 += -599
    v40 := GenS3{F0: false}
    v3 = append(v3, -59)
    var v41 uint32 = 299
    switch -613 {
    case 1:
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
    case -478:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println(v5, v41)
    switch "a" {
    case "foo bar":
        if v6 {
            var v42 int8 = v5
            if (v41 == 2) {
                v2 = append(v2, "foo bar")
            } else {
                var v43 int64 = 0
                _ = v43
            }
            v39 = (v39 << (((v1 % ((v1) | 1))) & 63))
            _ = v42
        } else {
            fmt.Println(v6, v6)
            if (!v6) {
                v44 := GenS4{F0: v5, F1: "\n\t"}
                var v45 bool = v6
                _ = v44
                _ = v45
            } else {
                v41 -= 752
            }
        }
    case "":
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            v46 := GenS1{F0: (v39 >> (((v1 - (v1 | (v1 * (v1 & (^(v1))))))) & 63)), F1: (-(v39))}
            v47 := GenS5{F0: ((("" + "") + ((("a" + ("\n\t" + "foo bar")) + ("\n\t" + ("\n\t" + ""))) + ("" + ("foo bar" + ("foo bar" + "foo bar"))))) + ("" + ("" + ("a" + (((("\n\t" + "\n\t") + "a") + "") + ("hello" + ("hello" + "foo bar"))))))), F1: 255, F2: (!v6)}
            fmt.Println(v6)
            _ = v46
            _ = v47
        }
    }
    v48 := GenS6{F0: -2147483648}
    fmt.Println(v5, v6, v1)
    for fi49 := 0; fi49 < 5; fi49++ {
        v6 = (v6 || v6)
        _ = fi49
    }
    var v50 uint = 0
    v5 -= (-(v5))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v48
    _ = v5
    _ = v50
    _ = v6
}
