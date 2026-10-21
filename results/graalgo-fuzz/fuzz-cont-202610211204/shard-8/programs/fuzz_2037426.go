package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 bool; F2 string }
type GenS4 struct { F0 uint32; F1 uint }
func genFn0(p0 bool) bool {
    for fi4 := 0; fi4 < 3; fi4++ {
        p0 = (9 != 42)
        if (-9223372036854775808 < 42) {
            v5 := []int{32768}
            var v6 int32 = -860
            _ = v5
            _ = v6
        } else {
            fmt.Println(fi4)
            p0 = ((!p0) || (468 == 58))
            p0 = (!(p0 && p0))
        }
        _ = fi4
    }
    switch "" {
    case "\n\t":
        if (!(!(-9223372036854775808 <= 32767))) {
            if p0 {
                fmt.Println(p0)
                p0 = false
                if true {
                    switch "foo bar" {
                    case "a":
                        v7 := []int{-1, 128}
                        _ = v7
                    case "foo bar":
                        if (1 > 65535) {
                            p0 = true
                        }
                    default:
                        p0 = p0
                    }
                } else {
                    p0 = (-2147483648 >= 510)
                    if ((127 == 800) && p0) {
                        if (788 != -563) {
                            var v8 int16 = 0
                            _ = v8
                        } else {
                            switch "" {
                            case "":
                                p0 = (p0 && true)
                            case "a":
                                fmt.Println(p0)
                            default:
                                var v9 string = "hello"
                                _ = v9
                            }
                        }
                    }
                }
            } else {
                p0 = (97 >= 2)
                if false {
                    if p0 {
                        switch "foo bar" {
                        case "\n\t":
                            var v10 int16 = -895
                            _ = v10
                        case "a":
                            if p0 {
                                if (-155 <= 42) {
                                    var v11 int8 = 100
                                    _ = v11
                                }
                            } else {
                                var v12 uint32 = 42
                                _ = v12
                            }
                        default:
                            var v13 float32 = 3.14159
                            _ = v13
                        }
                        var v14 int8 = 42
                        _ = v14
                    } else {
                        for fi15 := 0; fi15 < 11; fi15++ {
                            fmt.Println(fi15, fi15)
                            _ = fi15
                        }
                        v16 := GenS1{F0: 2, F1: "hello", F2: 576}
                        if (68 > 70) {
                            v17 := []int{-143}
                            _ = v17
                        }
                        _ = v16
                    }
                    v18 := GenS2{F0: 2}
                    _ = v18
                }
                fmt.Println(p0)
            }
        }
    default:
        v19 := []int{32768}
        _ = v19
    }
    _ = p0
    return p0
}
func genFn1() uint32 {
    var v20 int16 = 1
    for fi21 := 0; fi21 < 16; fi21++ {
        fmt.Println(fi21, fi21)
        v20 = v20
        for fi22 := 0; fi22 < 9; fi22++ {
            v20 = v20
            var v23 float64 = 3.14159
            if (2 <= 626) {
                if (!(2147483647 <= -2147483648)) {
                    v23 = float64(-2.25)
                    v24 := GenS3{F0: 65535, F1: false, F2: "foo bar"}
                    _ = v24
                }
                v25 := GenS3{F0: 928, F1: false, F2: (("\n\t" + "foo bar") + "a")}
                _ = v25
            }
            if true {
                var v26 uint64 = 329
                fmt.Println(v20, fi22, fi21)
                _ = v26
            }
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    switch -1 {
    case -2147483648:
        v20 = v20
    case 742:
        var v27 string = "hello"
        _ = v27
    default:
        if true {
            switch 149 {
            case 0:
                var v28 uint8 = 0
                _ = v28
            case 32767:
                var v29 bool = (-14 >= 100)
                _ = v29
            default:
                for fi30 := 0; fi30 < 20; fi30++ {
                    var v31 uint16 = 1
                    fmt.Println(v20, v31)
                    _ = fi30
                    _ = v31
                }
            }
        } else {
            v20 = -32768
            v32 := GenS4{F0: 428, F1: 4294967295}
            _ = v32
        }
    }
    _ = v20
    return 65535
}

func main() {
    var v1 int = 32767
    v2 := make(map[int]int)
    var v3 uint64 = 42
    var v33 uint = 1
    delete(v2, 9223372036854775807)
    v1 = (v1 & v1)
    for fi34 := 0; fi34 < 13; fi34++ {
        v1 = fi34
        _ = fi34
    }
    v35 := GenS3{F0: 255, F1: (false && (!false)), F2: "a"}
    for fi36 := 0; fi36 < 18; fi36++ {
        switch -2147483648 {
        case -788:
            var v37 int = fi36
            _ = v37
        }
        switch (v1 / ((42) | 1)) {
        case 946:
            v38 := []int{42, v1, (fi36 >> ((v33) & 63)), (fi36 - fi36)}
            _ = v38
        case 401:
            v39 := v2[fi36]
            _ = v39
        case 32767:
            switch "\n\t" {
            case "\n\t":
                v40 := v2[fi36]
                _ = v40
            case "foo bar":
                v33 = (v33 % ((v33) | 1))
            }
        default:
            v41 := GenS3{F0: 0, F1: false, F2: "\n\t"}
            _ = v41
        }
        _ = fi36
    }
    for fi42 := 0; fi42 < 11; fi42++ {
        var v43 uint32 = 2147483647
        fmt.Println(v1, v3, v1)
        v43 -= v43
        _ = fi42
        _ = v43
    }
    for fi44 := 0; fi44 < 19; fi44++ {
        var v45 uint = 255
        _ = fi44
        _ = v45
    }
    v33 = v33
    v33 -= v33
    var v46 int32 = 425
    fmt.Println(v33, v46)
    v47 := GenS3{F0: 0, F1: false, F2: "a"}
    fmt.Println(v46, v33, v33)
    var v48 uint16 = 91
    v33 = 629
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v35
    _ = v46
    _ = v47
    _ = v48
}
