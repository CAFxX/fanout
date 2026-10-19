package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int }
type GenS2 struct { F0 string; F1 int32; F2 int }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 int16; F2 int }
func genFn0(p0 int, p1 int8, p2 int64) uint8 {
    var v4 uint32 = 42
    var v5 uint32 = 14
    p0 -= -498
    if (true || false) {
        if (32767 == 42) {
            v6 := GenS1{F0: int(v4), F1: p1, F2: p0}
            _ = v6
        }
        v7 := []int{(-(p0)), p0}
        fmt.Println(p2, p1)
        _ = v7
    }
    switch "\n\t" {
    case "":
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := []int{(-(fi8))}
            _ = fi8
            _ = v9
        }
    case "\n\t":
        p0 += p0
    case "foo bar":
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 2
}
func genFn1(p0 bool, p1 bool) bool {
    p0 = p1
    var v10 int64 = 0
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 int64 = (-(v10))
        _ = fi11
        _ = v12
    }
    v13 := GenS2{F0: "foo bar", F1: 42, F2: 32767}
    fmt.Println(p0, p0, p1)
    _ = p0
    _ = p1
    _ = v10
    _ = v13
    return false
}

func main() {
    var v1 int = -1
    var v2 int8 = 42
    v3 := []string{"", "foo bar", "foo bar"}
    v14 := GenS2{F0: ("hello" + ("a" + "\n\t")), F1: 1, F2: (^(v1))}
    v15 := GenS3{F0: ((v1 >= (v1 - (v1 ^ v1))) && (2147483648 > 65535)), F1: -1}
    if (42 <= 371) {
        var v16 int = v1
        _ = v16
    } else {
        for fi17 := 0; fi17 < 3; fi17++ {
            if ("\n\t" >= ("foo bar" + "hello")) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v1 -= (fi17 ^ fi17)
            } else {
                v1 -= (^(fi17))
                v18 := GenS3{F0: true, F1: 32767}
                for fi19 := 0; fi19 < 2; fi19++ {
                    var v20 uint32 = 493
                    if (!((fi17 << ((65535) & 63)) == fi19)) {
                        v20 -= (^(v20))
                    }
                    if (!false) {
                        v20 = v20
                    }
                    _ = fi19
                    _ = v20
                }
                _ = v18
            }
            switch ("hello" + "") {
            case "\n\t":
                v21 := GenS4{F0: 647}
                _ = v21
            case "a":
                for fi22 := 0; fi22 < 15; fi22++ {
                    v1 -= 1
                    if (!(!(2 >= genFn0((-(v1)), 127, 1)))) {
                        switch "" {
                        case "a":
                            v2 = v2
                        case "hello":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        default:
                            fmt.Println(fi22, fi17, v1)
                        }
                    }
                    _ = fi22
                }
            default:
                v23 := GenS2{F0: (((("a" + ((("\n\t" + "") + "") + "\n\t")) + ("" + "hello")) + "") + "a"), F1: 42, F2: v1}
                _ = v23
            }
            _ = fi17
        }
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        v25 := GenS3{F0: (!(1 <= (fi24 >> ((432) & 63)))), F1: -1}
        v2 += (v2 * (v2 * int8(fi24)))
        _ = fi24
        _ = v25
    }
    v1 -= v1
    v1 += int(v2)
    var v26 uint = 255
    v27 := GenS4{F0: v26}
    v28 := GenS3{F0: true, F1: 32767}
    if (("foo bar" + "a") >= "foo bar") {
        if true {
            var v29 int64 = 1
            switch "\n\t" {
            case "\n\t":
                v29 -= v29
            default:
                for fi30 := 0; fi30 < 3; fi30++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v31 := GenS2{F0: ("hello" + "\n\t"), F1: 0, F2: v1}
                    _ = fi30
                    _ = v31
                }
            }
            if true {
                switch "foo bar" {
                case "hello":
                    switch "foo bar" {
                    case "":
                        if (2147483647 > 65535) {
                            for fi32 := 0; fi32 < 20; fi32++ {
                                fmt.Println(v29, fi32, v29)
                                if (!true) {
                                    v1 = v1
                                    var v33 uint = (v26 << ((4294967295) & 63))
                                    _ = v33
                                }
                                _ = fi32
                            }
                        }
                    }
                case "":
                    _ = len(v3)
                case "a":
                    fmt.Println(v2, v26, v1)
                }
                fmt.Println(v2, v26, v2)
                if true {
                    for fi34 := 0; fi34 < 3; fi34++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v1 -= (fi34 & v1)
                        if false {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        _ = fi34
                    }
                    switch ("\n\t" + "foo bar") {
                    case "":
                        v35 := GenS1{F0: int(v26), F1: (v2 >> (((v26 | (v26 | (v26 % ((v26) | 1))))) & 63)), F2: (v1 % (((v1 | v1)) | 1))}
                        _ = v35
                    default:
                        v1 += v1
                    }
                } else {
                    for fi36 := 0; fi36 < 7; fi36++ {
                        v37 := GenS5{F0: -128, F1: -32768, F2: fi36}
                        for fi38 := 0; fi38 < 20; fi38++ {
                            var v39 int32 = -906
                            _ = len(v3)
                            _ = fi38
                            _ = v39
                        }
                        v40 := GenS1{F0: 128, F1: v2, F2: v1}
                        _ = fi36
                        _ = v37
                        _ = v40
                    }
                    _ = len(v3)
                    var v41 string = "foo bar"
                    _ = v41
                }
            } else {
                v26 -= v26
            }
            _ = v29
        }
    }
    v1 += 32768
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
