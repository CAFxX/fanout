package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS5 struct { F0 uint64; F1 bool; F2 int64 }
func genFn0(p0 uint16, p1 int) uint16 {
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := []int{}
        fmt.Println(p0)
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        for fi9 := 0; fi9 < 5; fi9++ {
            fmt.Println(p0, p1, fi8)
            if (!true) {
                v10 := GenS1{F0: 125, F1: "foo bar"}
                if true {
                    p0 += p0
                } else {
                    var v11 bool = true
                    _ = v11
                }
                p1 = p1
                _ = v10
            } else {
                switch "\n\t" {
                case "hello":
                    switch "foo bar" {
                    case "a":
                        p0 = p0
                    case "\n\t":
                        if true {
                            if true {
                                fmt.Println(p0, p1, p1)
                                v12 := GenS1{F0: 1, F1: ("" + "")}
                                _ = v12
                            }
                            var v13 uint16 = 65535
                            if false {
                                v14 := GenS2{F0: p0}
                                p1 = (p1 ^ 42)
                                if false {
                                    v13 += v13
                                    var v15 uint64 = 18446744073709551615
                                    _ = v15
                                }
                                var v16 float32 = 3.14159
                                _ = v14
                                _ = v16
                            }
                            v17 := []int{fi8}
                            _ = v13
                            _ = v17
                        }
                    case "foo bar":
                        var v18 float32 = (0.0 + 3.14159)
                        _ = v18
                    }
                }
                p0 -= 843
            }
            _ = fi9
        }
        fmt.Println(p0, fi8, p1)
        _ = fi8
    }
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 int, p1 int32) int {
    p0 += p0
    fmt.Println(p1)
    if (0.1 > float64((-2.25 - 1e10))) {
        switch "hello" {
        case "foo bar":
            v19 := []int{}
            _ = v19
        }
    } else {
        for fi20 := 0; fi20 < 7; fi20++ {
            fmt.Println(p1)
            var v21 int64 = 1
            _ = fi20
            _ = v21
        }
    }
    v22 := GenS2{F0: 255}
    var v23 int = (p0 - (^(p0)))
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    return v23
}

func main() {
    v1 := make(map[string]int)
    v2 := []string{}
    var v3 int = -2147483648
    v4 := []int{127}
    var v5 uint8 = 42
    for fi24 := 0; fi24 < 3; fi24++ {
        v3 -= -525
        fmt.Println(v3, fi24, fi24)
        _ = fi24
    }
    v3 = -683
    v25 := GenS3{F0: -89}
    switch (-(v3)) {
    case 0:
        for fi26 := 0; fi26 < 19; fi26++ {
            if true {
                v27 := GenS2{F0: 429}
                fmt.Println(v3, v5)
                _ = v27
            }
            if false {
                var v28 uint8 = (^(v5))
                var v29 string = ("hello" + "\n\t")
                _ = v28
                _ = v29
            } else {
                switch genFn1(32768, 1) {
                case -9223372036854775808:
                    v5 += v5
                case 1:
                    if (true && (32767 < 1)) {
                        for fi30 := 0; fi30 < 10; fi30++ {
                            delete(v1, ("" + ""))
                            _ = fi30
                        }
                    }
                }
                for fi31 := 0; fi31 < 18; fi31++ {
                    v5 = 255
                    _ = fi31
                }
                for fi32 := 0; fi32 < 14; fi32++ {
                    var v33 uint32 = 1
                    v33 = v33
                    switch "" {
                    case "":
                        v1[""] = genFn1(fi26, 432)
                    case "\n\t":
                        v34 := GenS4{F0: 1, F1: 922, F2: 2}
                        _ = v34
                    default:
                        v1["foo bar"] = 127
                    }
                    _ = fi32
                    _ = v33
                }
            }
            var v35 int16 = -32768
            _ = fi26
            _ = v35
        }
    case 1:
        var v36 int32 = 42
        _ = v36
    }
    v3 += (v3 & 208)
    switch genFn1(v3, 0) {
    case -2147483648:
        var v37 uint32 = 2147483647
        _ = v37
    }
    v5 = 255
    if (!(true || true)) {
        for fi38 := 0; fi38 < 18; fi38++ {
            v5 = v5
            switch ("hello" + ("foo bar" + "")) {
            case "hello":
                fmt.Println(v5, v3)
            case "a":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                v3 += 32767
            }
            if (42 < 34) {
                if (!false) {
                    var v39 int8 = 89
                    if false {
                        fmt.Println(v5, v5, v39)
                        v39 += v39
                        fmt.Println(v39, v39)
                    } else {
                        var v40 float32 = float32(0.1)
                        var v41 uint16 = genFn0(42, v3)
                        v42 := GenS5{F0: 4294967295, F1: ((0.1 != (0.1 + 1.5)) || false), F2: 0}
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                    v5 -= (^(v5))
                    _ = v39
                }
                var v43 int16 = 1
                v44 := GenS5{F0: 42, F1: (true || true), F2: 0}
                _ = v43
                _ = v44
            }
            _ = fi38
        }
        if true {
            switch 32767 {
            case -2147483648:
                for fi45 := 0; fi45 < 1; fi45++ {
                    v3 = fi45
                    v46 := GenS1{F0: 2147483647, F1: "hello"}
                    _ = fi45
                    _ = v46
                }
            case 1:
                if true {
                    for fi47 := 0; fi47 < 15; fi47++ {
                        v3 += -9223372036854775808
                        _ = fi47
                    }
                    v5 += v5
                    switch genFn1(v3, 0) {
                    case 2147483647:
                        var v48 int16 = 42
                        _ = v48
                    case -9223372036854775808:
                        var v49 int = (v3 >> ((uint(v5)) & 63))
                        _ = v49
                    case 42:
                        for fi50 := 0; fi50 < 1; fi50++ {
                            fmt.Println(v3)
                            _ = fi50
                        }
                    }
                }
            default:
                v5 = (^(v5))
            }
            var v51 uint8 = (v5 * 255)
            fmt.Println(v5, v3)
            _ = v51
        } else {
            v5 += (v5 >> ((2) & 63))
        }
    }
    v52 := GenS1{F0: 42, F1: "a"}
    v3 = (^(v3))
    if true {
        var v53 bool = (true && false)
        _ = v53
    }
    switch (^(v3)) {
    case 127:
        for fi54 := 0; fi54 < 7; fi54++ {
            v5 = (v5 + v5)
            _ = fi54
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v52
}
