package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 uint64; F2 int8 }
type GenS4 struct { F0 uint; F1 int8; F2 bool }
type GenS5 struct { F0 bool }
func genFn0(p0 uint, p1 string) int {
    v5 := GenS1{F0: -1, F1: 0, F2: ("foo bar" + p1)}
    v6 := []int{499, 434, 2147483647}
    var v7 string = p1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 32768
}
func genFn1() uint8 {
    var v8 uint64 = 1
    if true {
        var v9 string = ("\n\t" + "a")
        for fi10 := 0; fi10 < 11; fi10++ {
            for fi11 := 0; fi11 < 5; fi11++ {
                var v12 string = v9
                _ = fi11
                _ = v12
            }
            var v13 float32 = 1e10
            _ = fi10
            _ = v13
        }
        _ = v9
    }
    fmt.Println(v8)
    _ = v8
    return 127
}

func main() {
    var v1 uint16 = 0
    var v2 int8 = 0
    var v3 int64 = -1
    var v4 int32 = 1
    switch 583 {
    case 0:
        v1 = v1
    case 32767:
        v14 := GenS2{F0: 42}
        _ = v14
    case 42:
        v15 := GenS3{F0: (v4 / (((v4 / (((v4 / ((v4) | 1))) | 1))) | 1)), F1: 18446744073709551615, F2: (v2 >> ((2) & 63))}
        _ = v15
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        v17 := []int{}
        v18 := GenS1{F0: (v4 << ((2) & 63)), F1: (fi16 % ((-750) | 1)), F2: (("\n\t" + ("a" + ("foo bar" + (("\n\t" + "hello") + "")))) + "")}
        v19 := GenS3{F0: (v4 >> ((uint(v1)) & 63)), F1: 18446744073709551615, F2: (v2 + v2)}
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    v20 := GenS3{F0: v4, F1: 151, F2: v2}
    if true {
        var v21 uint = 2
        _ = v21
    }
    fmt.Println(v1, v1, v2)
    for fi22 := 0; fi22 < 5; fi22++ {
        v23 := GenS1{F0: v4, F1: int(v4), F2: "\n\t"}
        if true {
            for fi24 := 0; fi24 < 20; fi24++ {
                v3 = v3
                _ = fi24
            }
        } else {
            var v25 int8 = v2
            v3 = v3
            _ = v25
        }
        _ = fi22
        _ = v23
    }
    var v26 uint32 = 0
    v26 += v26
    if (672 != -1) {
        switch 127 {
        case 1:
            var v27 uint32 = v26
            _ = v27
        default:
            if (3.14159 != 3.14159) {
                v4 = (v4 * v4)
                if false {
                    var v28 uint = 0
                    fmt.Println(v4, v4, v4)
                    switch 42 {
                    case 1:
                        v26 = (v26 | (^(v26)))
                    case -9223372036854775808:
                        switch "foo bar" {
                        case "a":
                            v1 += (v1 * 65535)
                        case "hello":
                            v29 := GenS4{F0: v28, F1: (v2 * (v2 ^ v2)), F2: true}
                            _ = v29
                        case "\n\t":
                            v2 = v2
                        }
                    case 32768:
                        var v30 int32 = v4
                        _ = v30
                    }
                    _ = v28
                }
                var v31 float32 = 3.14159
                _ = v31
            } else {
                switch 42 {
                case 1:
                    var v32 float64 = 1.5
                    _ = v32
                case -581:
                    v33 := []int{}
                    _ = v33
                case -228:
                    fmt.Println(v3, v26, v4)
                default:
                    for fi34 := 0; fi34 < 8; fi34++ {
                        switch "foo bar" {
                        case "":
                            if (!true) {
                                switch fi34 {
                                case -9223372036854775808:
                                    var v35 int = fi34
                                    _ = v35
                                default:
                                    for fi36 := 0; fi36 < 4; fi36++ {
                                        var v37 int = fi36
                                        v1 -= uint16(v26)
                                        v38 := GenS3{F0: int32(v3), F1: 42, F2: (^(v2))}
                                        v39 := GenS3{F0: v4, F1: 0, F2: int8(v1)}
                                        _ = fi36
                                        _ = v37
                                        _ = v38
                                        _ = v39
                                    }
                                }
                                switch ("hello" + "hello") {
                                case "hello":
                                    fmt.Println(v26, v2, v26)
                                }
                            }
                        case "a":
                            if true {
                                v1 = v1
                                if true {
                                    switch genFn0(4294967295, "\n\t") {
                                    case 4294967295:
                                        v4 = 2147483647
                                    case 700:
                                        v4 -= (v4 | 2147483647)
                                    }
                                    v40 := []int{(^(fi34)), genFn0(42, ("foo bar" + "hello"))}
                                    _ = v40
                                } else {
                                    if ("foo bar" <= "") {
                                        for fi41 := 0; fi41 < 20; fi41++ {
                                            fmt.Println(v1, v4, v3)
                                            fmt.Println(v1, v1)
                                            fmt.Println(v3, fi41)
                                            _ = fi41
                                        }
                                        v3 = v3
                                    }
                                    v42 := GenS3{F0: v4, F1: uint64(v26), F2: (v2 >> ((18446744073709551615) & 63))}
                                    _ = v42
                                }
                                v43 := []int{(fi34 ^ fi34), (-(fi34))}
                                _ = v43
                            } else {
                                v44 := []int{}
                                var v45 uint8 = 99
                                if true {
                                    if (genFn1() >= v45) {
                                        v3 = v3
                                    } else {
                                        v44 = append(v44, 1)
                                        var v46 string = (("\n\t" + "a") + ("" + ((("" + "a") + "") + "\n\t")))
                                        var v47 int64 = v3
                                        _ = v46
                                        _ = v47
                                    }
                                } else {
                                    if (229 < 802) {
                                        for fi48 := 0; fi48 < 8; fi48++ {
                                            var v49 bool = (42 >= (fi48 * fi34))
                                            _ = fi48
                                            _ = v49
                                        }
                                        fmt.Println(v2)
                                        for fi50 := 0; fi50 < 9; fi50++ {
                                            v44 = append(v44, fi50)
                                            v51 := GenS2{F0: 18446744073709551615}
                                            _ = fi50
                                            _ = v51
                                        }
                                    } else {
                                        v3 -= v3
                                        v2 -= (v2 | (v2 / ((v2) | 1)))
                                    }
                                    for fi52 := 0; fi52 < 9; fi52++ {
                                        var v53 bool = false
                                        var v54 int32 = (v4 | v4)
                                        switch ("\n\t" + ("foo bar" + "a")) {
                                        case "foo bar":
                                            switch "foo bar" {
                                            case "\n\t":
                                                switch "" {
                                                case "foo bar":
                                                    if false {
                                                        if len(v44) >= 2 {
                                                            v44 = v44[1:len(v44)]
                                                        }
                                                        _ = len(v44)
                                                        v55 := GenS3{F0: (v54 >> ((42) & 63)), F1: 42, F2: v2}
                                                        _ = v55
                                                    } else {
                                                        fmt.Println(fi52)
                                                    }
                                                }
                                                v56 := GenS2{F0: 42}
                                                v53 = (!v53)
                                                v57 := GenS5{F0: v53}
                                                _ = v56
                                                _ = v57
                                            }
                                        case "":
                                            if (-32768 != 32767) {
                                                v58 := GenS3{F0: v54, F1: 661, F2: v2}
                                                fmt.Println(fi52, fi52)
                                                v53 = v53
                                                v59 := GenS5{F0: v53}
                                                _ = v58
                                                _ = v59
                                            }
                                        default:
                                            v3 += -273
                                        }
                                        _ = fi52
                                        _ = v53
                                        _ = v54
                                    }
                                }
                                _ = v44
                                _ = v45
                            }
                        }
                        _ = fi34
                    }
                }
                v60 := []int{0, 9223372036854775807, genFn0(2, ("foo bar" + "")), 0}
                _ = v60
            }
        }
    }
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
}
