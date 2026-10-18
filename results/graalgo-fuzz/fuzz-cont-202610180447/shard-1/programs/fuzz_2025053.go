package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint16; F1 uint64 }
func genFn0() uint64 {
    v7 := []int{2147483647, -2147483648, 0}
    var v8 uint32 = 2
    for fi9 := 0; fi9 < 20; fi9++ {
        v8 = v8
        v8 = 212
        _ = fi9
    }
    v8 -= (v8 / ((v8) | 1))
    if (!true) {
        v8 += (v8 + v8)
    }
    _ = v7
    _ = v8
    return 2
}
func genFn1() uint32 {
    var v10 uint16 = 42
    fmt.Println(v10)
    var v11 uint16 = v10
    _ = v10
    _ = v11
    return 65535
}

func main() {
    v1 := []string{("" + ("a" + "a")), "\n\t", "\n\t"}
    v2 := make(map[string]string)
    var v3 int = 127
    v4 := make(map[uint32]int)
    var v5 float64 = float64(1.5)
    var v6 int = (^(v3))
    var v12 uint16 = 0
    switch (v3 % ((v6) | 1)) {
    case -1:
        var v13 uint16 = v12
        _ = v13
    case 4294967295:
        for fi14 := 0; fi14 < 2; fi14++ {
            v15 := GenS1{F0: true}
            _ = fi14
            _ = v15
        }
    default:
        if ((true && (true && (1000 != 1000))) || false) {
            if true {
                v16 := GenS1{F0: true}
                _ = v16
            } else {
                var v17 int8 = 42
                v18 := GenS2{F0: "\n\t", F1: v17}
                _ = v17
                _ = v18
            }
        } else {
            if true {
                var v19 uint32 = 881
                if false {
                    v6 = v3
                    for fi20 := 0; fi20 < 4; fi20++ {
                        _ = len(v1)
                        switch "\n\t" {
                        case "\n\t":
                            v2["\n\t"] = "foo bar"
                        default:
                            for fi21 := 0; fi21 < 3; fi21++ {
                                fmt.Println(v12, v6)
                                _ = fi21
                            }
                        }
                        _ = fi20
                    }
                    v22 := GenS1{F0: true}
                    _ = v22
                }
                _ = v19
            }
        }
    }
    v23 := GenS2{F0: "foo bar", F1: 89}
    if len(v1) > 0 {
        v1[0] = "a"
    }
    if true {
        v24 := GenS1{F0: false}
        _ = v24
    } else {
        _ = len(v1)
        var v25 uint32 = 1
        _ = v25
    }
    if false {
        if (-2.25 < ((0.1 / (0.1 / 1e10)) - -2.25)) {
            fmt.Println(v3, v12, v6)
        } else {
            fmt.Println(v5, v3, v6)
        }
        if true {
            fmt.Println(v5, v3)
            var v26 int = (v3 - v3)
            v27 := GenS2{F0: "", F1: 100}
            _ = v26
            _ = v27
        } else {
            for fi28 := 0; fi28 < 6; fi28++ {
                var v29 uint8 = 65
                var v30 bool = (42 > 18446744073709551615)
                _ = fi28
                _ = v29
                _ = v30
            }
            for fi31 := 0; fi31 < 9; fi31++ {
                if true {
                    for fi32 := 0; fi32 < 1; fi32++ {
                        v33 := GenS1{F0: false}
                        _ = fi32
                        _ = v33
                    }
                    _ = len(v1)
                } else {
                    v12 = v12
                    var v34 uint16 = (v12 | v12)
                    v35 := GenS1{F0: (100 >= -16)}
                    _ = v34
                    _ = v35
                }
                if false {
                    if (!true) {
                        switch 2147483647 {
                        case 127:
                            for fi36 := 0; fi36 < 13; fi36++ {
                                switch "hello" {
                                case "":
                                    if ((false && ((-964 >= 231) && (false && false))) && true) {
                                        var v37 int = (-(fi31))
                                        fmt.Println(v3, v3, v6)
                                        var v38 bool = true
                                        _ = v37
                                        _ = v38
                                    }
                                case "\n\t":
                                    var v39 uint16 = (v12 ^ 65535)
                                    _ = v39
                                default:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                _ = fi36
                            }
                        }
                        v40 := GenS1{F0: ((true && false) && true)}
                        _ = v40
                    }
                    var v41 float32 = 3.14159
                    _ = v41
                } else {
                    var v42 bool = false
                    var v43 string = ("foo bar" + "foo bar")
                    var v44 float64 = ((3.14159 + ((float64(0.0) + float64(((1e10 - (1e10 * 3.14159)) - (0.1 - 0.0)))) - ((0.0 + v5) + v5))) - (v5 - (v5 * v5)))
                    _ = v42
                    _ = v43
                    _ = v44
                }
                _ = fi31
            }
        }
    }
    v1 = append(v1, "a")
    v3 = v3
    var v45 float64 = (v5 / v5)
    if true {
        switch v6 {
        case -2147483648:
            fmt.Println(v45, v5)
        case 372:
            var v46 uint = 18446744073709551615
            _ = v46
        case 560:
            if true {
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
                for fi47 := 0; fi47 < 13; fi47++ {
                    v48 := GenS2{F0: "a", F1: -128}
                    v49 := GenS1{F0: false}
                    v50 := GenS2{F0: (("\n\t" + ("hello" + "a")) + "\n\t"), F1: 42}
                    _ = fi47
                    _ = v48
                    _ = v49
                    _ = v50
                }
                if (!(30 > 1)) {
                    for fi51 := 0; fi51 < 8; fi51++ {
                        v5 += 0.1
                        var v52 string = ""
                        _ = fi51
                        _ = v52
                    }
                    for fi53 := 0; fi53 < 6; fi53++ {
                        v3 += (v3 % ((fi53) | 1))
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi53
                    }
                }
            } else {
                var v54 float64 = v45
                _ = v54
            }
        default:
            fmt.Println(v45, v12)
        }
    }
    for fi55 := 0; fi55 < 9; fi55++ {
        switch (v3 * 32767) {
        case 2147483647:
            if len(v1) > 0 {
                v1[0] = "hello"
            }
        case 11:
            for fi56 := 0; fi56 < 13; fi56++ {
                var v57 uint16 = (^(v12))
                _ = fi56
                _ = v57
            }
        case 1:
            v6 += v3
        default:
            var v58 uint32 = 577
            _ = v58
        }
        for fi59 := 0; fi59 < 7; fi59++ {
            fmt.Println(v45, v6)
            var v60 float64 = v45
            _ = fi59
            _ = v60
        }
        delete(v2, "a")
        _ = fi55
    }
    v61 := GenS3{F0: (^(v12)), F1: 0}
    v5 += ((v5 + v45) * 3.14159)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v61
}
