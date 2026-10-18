package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int64; F1 int; F2 int8 }
type GenS3 struct { F0 bool; F1 uint8; F2 bool }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 int; F1 int64; F2 int64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int }
func genFn0() uint64 {
    switch ("a" + "\n\t") {
    case "foo bar":
        var v5 string = (("hello" + "foo bar") + "a")
        _ = v5
    }
    switch "a" {
    case "":
        var v6 uint32 = 2147483648
        _ = v6
    default:
        var v7 int16 = 0
        _ = v7
    }
    var v8 bool = true
    if (97 > 42) {
        v9 := []int{9223372036854775807, 285, 4294967295}
        _ = v9
    } else {
        var v10 uint64 = 1
        v10 = 0
        _ = v10
    }
    v11 := GenS1{F0: 1, F1: 128}
    _ = v11
    _ = v8
    return 933
}

func main() {
    var v1 int = 42
    var v2 float32 = float32(0.1)
    var v3 int = v1
    var v4 int32 = 206
    if (-6 >= -128) {
        fmt.Println(v4, v3, v1)
        if ("a" <= ("\n\t" + "")) {
            var v12 int32 = v4
            _ = v12
        }
        v13 := GenS1{F0: -1, F1: 255}
        _ = v13
    } else {
        if (false && false) {
            var v14 int32 = (v4 << ((302) & 63))
            if (!true) {
                v1 = 323
            }
            for fi15 := 0; fi15 < 15; fi15++ {
                var v16 uint64 = 264
                _ = fi15
                _ = v16
            }
            _ = v14
        }
        v4 = -1
        v17 := []int{v1}
        _ = v17
    }
    v1 = (v3 & v3)
    if ((!false) || true) {
        v18 := []int{}
        _ = v18
    }
    var v19 uint8 = 0
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 bool = false
        _ = fi20
        _ = v21
    }
    if false {
        v4 -= v4
        switch (^(v3)) {
        case 1:
            v19 += 128
        default:
            var v22 float64 = 0.0
            _ = v22
        }
        if (1 < 0) {
            if (-1 > 0) {
                v1 += v3
            }
            fmt.Println(v1, v4, v1)
        } else {
            fmt.Println(v3, v4, v2)
            v23 := []int{}
            _ = v23
        }
    } else {
        for fi24 := 0; fi24 < 19; fi24++ {
            v1 -= (v1 - 1)
            _ = fi24
        }
        if false {
            if false {
                switch "" {
                case "foo bar":
                    fmt.Println(v19, v2)
                case "hello":
                    v25 := GenS2{F0: -353, F1: int(v19), F2: -64}
                    _ = v25
                case "\n\t":
                    fmt.Println(v4, v19)
                }
                var v26 bool = (("hello" < "\n\t") || ((v4 ^ int32(v19)) <= -1))
                _ = v26
            }
        }
        v1 = (-(v3))
    }
    var v27 uint64 = genFn0()
    switch "foo bar" {
    case "foo bar":
        for fi28 := 0; fi28 < 8; fi28++ {
            if (!(57 >= -1)) {
                v2 = (v2 - v2)
                if (32767 != 45) {
                    v2 = -2.25
                } else {
                    if (-1 > 0) {
                        if (true || false) {
                            if (!false) {
                                switch (v1 + int(v19)) {
                                case 127:
                                    var v29 uint16 = 0
                                    _ = v29
                                }
                                if false {
                                    v30 := []int{(v3 / ((674) | 1)), v3, (fi28 >> ((427) & 63)), (^(v3))}
                                    if len(v30) > 0 {
                                        v30[0] = (v1 + (v3 << ((4294967295) & 63)))
                                    }
                                    _ = v30
                                } else {
                                    for fi31 := 0; fi31 < 15; fi31++ {
                                        v32 := GenS3{F0: (!(42 <= (v4 & 541))), F1: v19, F2: (0.1 < -2.25)}
                                        _ = fi31
                                        _ = v32
                                    }
                                    var v33 float64 = -2.25
                                    if (1 > 1000) {
                                        fmt.Println(v27, v1, v1)
                                        v19 = (v19 / ((uint8(v27)) | 1))
                                    } else {
                                        v34 := []int{}
                                        _ = v34
                                    }
                                    _ = v33
                                }
                            } else {
                                var v35 bool = false
                                _ = v35
                            }
                            v36 := GenS2{F0: -424, F1: v3, F2: 42}
                            _ = v36
                        }
                        v37 := GenS4{F0: 44, F1: "foo bar"}
                        for fi38 := 0; fi38 < 6; fi38++ {
                            var v39 uint64 = (v27 % (((^(v27))) | 1))
                            v2 += (v2 / v2)
                            _ = fi38
                            _ = v39
                        }
                        _ = v37
                    }
                    v40 := []int{v3, 915}
                    _ = v40
                }
            } else {
                var v41 bool = ((1000 < 42) && (1 < 65535))
                _ = v41
            }
            v42 := []int{v1}
            for fi43 := 0; fi43 < 13; fi43++ {
                if (false && true) {
                    v44 := GenS5{F0: 0, F1: 835}
                    var v45 uint32 = 162
                    if false {
                        if len(v42) >= 2 {
                            v42 = v42[1:len(v42)]
                        }
                        fmt.Println(v1, v3, v19)
                        var v46 int32 = v4
                        fmt.Println(v27, v2)
                        _ = len(v42)
                        _ = v46
                    } else {
                        switch "\n\t" {
                        case "hello":
                            if true {
                                var v47 int8 = 102
                                _ = v47
                            }
                        case "":
                            if len(v42) > 0 {
                                v42[0] = (fi28 >> ((2) & 63))
                            }
                        case "foo bar":
                            var v48 string = ("hello" + "a")
                            _ = v48
                        }
                        v19 += (^(v19))
                        fmt.Println(v45, v1)
                        fmt.Println(fi28, v45)
                    }
                    _ = v44
                    _ = v45
                } else {
                    var v49 int8 = 127
                    var v50 uint = 194
                    var v51 string = "\n\t"
                    _ = v49
                    _ = v50
                    _ = v51
                }
                _ = fi43
            }
            _ = fi28
            _ = v42
        }
    case "":
        v52 := GenS1{F0: 0, F1: v19}
        _ = v52
    case "hello":
        v1 = -537
    default:
        v53 := []int{(v1 + -409), (v1 | int(v19))}
        _ = v53
    }
    v54 := GenS6{F0: 2147483647, F1: 479, F2: 0}
    v4 += v4
    v27 += v27
    v1 = v1
    v55 := GenS3{F0: (true || true), F1: (^(v19)), F2: true}
    var v56 uint8 = v19
    for fi57 := 0; fi57 < 13; fi57++ {
        v58 := GenS7{F0: 299}
        v59 := GenS8{F0: (^(v3))}
        _ = fi57
        _ = v58
        _ = v59
    }
    v1 = int(v4)
    v4 -= (-(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v54
    _ = v55
    _ = v56
}
