package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int64; F2 int32 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 string; F1 int; F2 string }
type GenS5 struct { F0 bool; F1 string }
type GenS6 struct { F0 uint32 }
func genFn0(p0 bool) uint8 {
    v7 := GenS1{F0: (757 < 877)}
    v8 := []int{}
    switch ("foo bar" + (("\n\t" + ("foo bar" + "hello")) + "")) {
    case "hello":
        if (false && (!p0)) {
            for fi9 := 0; fi9 < 7; fi9++ {
                v10 := GenS1{F0: p0}
                switch 0 {
                case -1:
                    var v11 int64 = 1
                    _ = v11
                default:
                    fmt.Println(p0, p0)
                }
                for fi12 := 0; fi12 < 17; fi12++ {
                    v13 := GenS1{F0: ((32767 <= 255) || false)}
                    var v14 uint32 = 65535
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                _ = fi9
                _ = v10
            }
        }
    case "":
        if (42 != 1) {
            fmt.Println(p0)
            var v15 bool = (0.1 < (-2.25 + (float64(1.5) / (3.14159 * 1.5))))
            v15 = (!v15)
            _ = v15
        } else {
            for fi16 := 0; fi16 < 6; fi16++ {
                if false {
                    if (1e10 != 1e10) {
                        p0 = (((true && true) || (false || (true && (fi16 < (fi16 ^ fi16))))) || p0)
                        switch "\n\t" {
                        case "a":
                            var v17 float64 = ((0.0 / 1e10) + 1e10)
                            _ = v17
                        case "\n\t":
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                        case "hello":
                            p0 = (p0 || (fi16 == 2147483647))
                        default:
                            var v18 bool = p0
                            _ = v18
                        }
                        fmt.Println(fi16, p0)
                    } else {
                        switch -9223372036854775808 {
                        case 42:
                            for fi19 := 0; fi19 < 10; fi19++ {
                                switch 32768 {
                                case 42:
                                    switch "foo bar" {
                                    case "hello":
                                        var v20 bool = p0
                                        _ = v20
                                    case "":
                                        var v21 int8 = 127
                                        _ = v21
                                    case "\n\t":
                                        v22 := GenS1{F0: p0}
                                        _ = v22
                                    default:
                                        p0 = p0
                                        v23 := GenS1{F0: (-1 >= -1)}
                                        _ = v23
                                    }
                                case 9223372036854775807:
                                    if (1 > 202) {
                                        if true {
                                            switch ("foo bar" + ("foo bar" + "\n\t")) {
                                            case "hello":
                                                p0 = (((1e10 - 3.14159) / 3.14159) != 1.5)
                                            }
                                            p0 = true
                                            var v24 uint32 = 2
                                            _ = v24
                                        }
                                        var v25 int32 = 42
                                        _ = v25
                                    } else {
                                        var v26 uint8 = 76
                                        v26 = v26
                                        fmt.Println(v26, v26, fi19)
                                        _ = v26
                                    }
                                default:
                                    p0 = p0
                                }
                                if true {
                                    switch "hello" {
                                    case "foo bar":
                                        fmt.Println(fi19)
                                    case "a":
                                        v27 := GenS1{F0: false}
                                        _ = v27
                                    }
                                    if len(v8) > 0 {
                                        v8[0] = fi19
                                    }
                                    v28 := GenS1{F0: (true && (42 < 127))}
                                    _ = v28
                                } else {
                                    var v29 int = (fi16 ^ (fi19 ^ fi19))
                                    fmt.Println(p0, fi16, p0)
                                    v30 := GenS1{F0: (19 == 136)}
                                    _ = v29
                                    _ = v30
                                }
                                p0 = p0
                                _ = fi19
                            }
                        case -187:
                            switch (^(fi16)) {
                            case 4294967295:
                                for fi31 := 0; fi31 < 3; fi31++ {
                                    if (false || (486 < -334)) {
                                        v8 = append(v8, 701)
                                        if true {
                                            var v32 bool = ("hello" <= "foo bar")
                                            _ = v32
                                        } else {
                                            v33 := GenS2{F0: 1, F1: 1, F2: -341}
                                            _ = v33
                                        }
                                        var v34 uint32 = 4294967295
                                        _ = v34
                                    }
                                    v35 := GenS1{F0: (!p0)}
                                    _ = fi31
                                    _ = v35
                                }
                            default:
                                if len(v8) > 0 {
                                    v8[0] = fi16
                                }
                            }
                        }
                    }
                } else {
                    var v36 uint = 0
                    _ = v36
                }
                fmt.Println(fi16)
                v37 := GenS3{F0: 1, F1: "\n\t"}
                _ = fi16
                _ = v37
            }
            if (1 >= -352) {
                p0 = (842 > 176)
            }
            var v38 bool = (!(0 >= 255))
            _ = v38
        }
    default:
        var v39 int16 = 1
        _ = v39
    }
    fmt.Println(p0)
    if (2147483647 >= 42) {
        var v40 string = ""
        v41 := GenS4{F0: "a", F1: -2147483648, F2: v40}
        v40 = ""
        _ = v40
        _ = v41
    }
    _ = p0
    _ = v7
    _ = v8
    return 128
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = 1
    var v3 int32 = 2147483647
    v4 := make(map[string]int)
    var v5 int32 = 42
    var v6 bool = true
    switch "foo bar" {
    case "hello":
        delete(v1, 261)
    }
    var v42 int8 = 121
    if (42 > 492) {
        v43 := GenS5{F0: (((v2 * 1000) < v2) && false), F1: "hello"}
        var v44 int = 128
        _ = v43
        _ = v44
    }
    v3 = v5
    if (4294967295 <= 42) {
        if false {
            var v45 int32 = 1
            var v46 int8 = v42
            _ = v45
            _ = v46
        }
        v47 := []int{32768, -801}
        for fi48 := 0; fi48 < 6; fi48++ {
            delete(v1, 695)
            var v49 string = "a"
            _ = fi48
            _ = v49
        }
        _ = v47
    } else {
        v50 := GenS6{F0: 42}
        fmt.Println(v2)
        _ = v50
    }
    var v51 uint = 0
    v2 = v2
    var v52 int64 = -967
    var v53 uint64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v6
}
