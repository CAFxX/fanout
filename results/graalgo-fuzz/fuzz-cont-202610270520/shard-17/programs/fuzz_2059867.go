package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 string }
type GenS2 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS3 struct { F0 int; F1 string; F2 uint64 }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 string }
func genFn0(p0 int8, p1 int16, p2 int16) int64 {
    p0 = p0
    v6 := GenS1{F0: 65535, F1: 0, F2: "foo bar"}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 567
}
func genFn1(p0 uint64, p1 uint64, p2 uint8) int {
    fmt.Println(p0)
    var v7 string = "foo bar"
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 float64 = 0.0
            v7 = ("" + "")
            v7 = v7
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: 42, F1: (p2 % ((p2) | 1)), F2: v7}
        _ = fi8
        _ = v11
    }
    v12 := []int{2147483647, 1}
    var v13 float32 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v7
    return 2147483647
}

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 401
    var v3 int32 = 0
    v4 := make(map[int]bool)
    var v5 int32 = v3
    v14 := GenS1{F0: 717, F1: 213, F2: ""}
    switch 128 {
    case 42:
        var v15 uint8 = 127
        _ = v15
    case 32767:
        if true {
            delete(v4, 42)
            switch "hello" {
            case "\n\t":
                if false {
                    if false {
                        fmt.Println(v3, v2)
                    } else {
                        if false {
                            v5 = 2147483647
                            if true {
                                fmt.Println(v3)
                                for fi16 := 0; fi16 < 13; fi16++ {
                                    fmt.Println(fi16)
                                    _ = fi16
                                }
                                v5 -= (v5 - (v3 << ((42) & 63)))
                            }
                        }
                        for fi17 := 0; fi17 < 15; fi17++ {
                            fmt.Println(v3, v2)
                            var v18 int = 969
                            _ = fi17
                            _ = v18
                        }
                        v5 = v5
                    }
                }
            case "a":
                v5 -= 1
            default:
                switch "\n\t" {
                case "\n\t":
                    switch genFn1(18446744073709551615, v2, 127) {
                    case -1:
                        var v19 uint = 255
                        _ = v19
                    }
                default:
                    v20 := GenS1{F0: 560, F1: 0, F2: "a"}
                    _ = v20
                }
            }
        }
    case 924:
        if true {
            for fi21 := 0; fi21 < 8; fi21++ {
                for fi22 := 0; fi22 < 10; fi22++ {
                    var v23 int8 = 1
                    fmt.Println(v2, v2)
                    _ = fi22
                    _ = v23
                }
                _ = fi21
            }
            for fi24 := 0; fi24 < 2; fi24++ {
                v25 := GenS1{F0: 42, F1: 2, F2: ("\n\t" + "")}
                for fi26 := 0; fi26 < 5; fi26++ {
                    v4[fi24] = (!false)
                    v27 := []int{genFn1(69, (v2 & (^(v2))), 127), (-(fi24)), 7}
                    _ = fi26
                    _ = v27
                }
                _ = fi24
                _ = v25
            }
            fmt.Println(v3)
        } else {
            v5 = (v5 % (((v3 ^ v5)) | 1))
            v3 += v3
        }
    }
    v3 += v3
    v5 = v3
    v28 := []int{}
    v29 := GenS1{F0: 1, F1: 1, F2: "hello"}
    v2 -= 2
    if false {
        if (-366 != 2147483647) {
            v4[-9223372036854775808] = false
        } else {
            if true {
                v30 := GenS1{F0: 101, F1: 172, F2: ("\n\t" + "hello")}
                var v31 uint16 = 1
                v3 += v3
                _ = v30
                _ = v31
            } else {
                v3 -= v3
            }
            for fi32 := 0; fi32 < 18; fi32++ {
                for fi33 := 0; fi33 < 4; fi33++ {
                    switch genFn1(1, v2, 38) {
                    case -2147483648:
                        if (62 == 65535) {
                            if true {
                                v2 -= (v2 ^ v2)
                                var v34 string = (("hello" + "a") + ("" + "\n\t"))
                                var v35 int16 = 32767
                                _ = v34
                                _ = v35
                            } else {
                                var v36 uint16 = 1
                                var v37 float32 = (-2.25 * (3.14159 + ((3.14159 / 1.5) / 1e10)))
                                v38 := GenS2{F0: 255, F1: 4294967295, F2: (v5 << ((65535) & 63))}
                                var v39 uint64 = v2
                                _ = v36
                                _ = v37
                                _ = v38
                                _ = v39
                            }
                            v28 = append(v28, (fi32 << ((65535) & 63)))
                            var v40 uint64 = 4294967295
                            var v41 uint = 0
                            _ = v40
                            _ = v41
                        }
                    case 1:
                        if true {
                            v42 := GenS3{F0: (^(fi32)), F1: ("a" + ("a" + "a")), F2: (v2 - v2)}
                            _ = len(v28)
                            _ = v42
                        }
                    }
                    if true {
                        fmt.Println(v5)
                        fmt.Println(v2, v3, fi32)
                        v5 -= -247
                    }
                    switch "hello" {
                    case "a":
                        var v43 uint16 = 1
                        _ = v43
                    case "hello":
                        if len(v28) >= 2 {
                            v28 = v28[1:len(v28)]
                        }
                    }
                    _ = fi33
                }
                _ = fi32
            }
            v44 := GenS4{F0: 2, F1: 42}
            _ = v44
        }
        var v45 int = -602
        _ = v45
    } else {
        var v46 int8 = -57
        if true {
            v2 = v2
        } else {
            v4[-80] = (!true)
            if false {
                if false {
                    v2 = v2
                    v46 = (v46 - 127)
                } else {
                    fmt.Println(v2, v5, v3)
                    if (!(v2 < (v2 - 2))) {
                        v28 = append(v28, -133)
                    }
                    v47 := GenS2{F0: 1, F1: 18446744073709551615, F2: (v3 % ((v3) | 1))}
                    _ = v47
                }
            }
        }
        if true {
            var v48 int16 = -1
            _ = v48
        } else {
            for fi49 := 0; fi49 < 1; fi49++ {
                var v50 bool = (v2 != (v2 * v2))
                if true {
                    fmt.Println(v50, v50)
                    _ = len(v1)
                } else {
                    if (42 < -888) {
                        fmt.Println(v5, v46, v46)
                        var v51 uint8 = 1
                        _ = v51
                    }
                    v28 = append(v28, 32767)
                    var v52 int16 = -620
                    _ = v52
                }
                _ = fi49
                _ = v50
            }
        }
        _ = v46
    }
    for fi53 := 0; fi53 < 8; fi53++ {
        v2 = (v2 << ((65535) & 63))
        v54 := GenS2{F0: 42, F1: v2, F2: v5}
        var v55 float64 = 3.14159
        _ = fi53
        _ = v54
        _ = v55
    }
    if true {
        v5 += (v3 | v3)
        v2 = (v2 << ((597) & 63))
    }
    v56 := GenS5{F0: "a"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v56
}
