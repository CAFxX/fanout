package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 uint; F2 uint32 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 uint32; F1 int16 }
type GenS6 struct { F0 bool; F1 int8; F2 uint16 }
func genFn0(p0 uint64, p1 int, p2 bool) bool {
    v6 := GenS1{F0: 42, F1: 261}
    if false {
        var v7 int = 127
        _ = v7
    } else {
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := []int{int(p0), (fi8 * 115), 32768, p1}
            switch int(p0) {
            case -2147483648:
                if len(v9) > 0 {
                    v9[0] = fi8
                }
            case 0:
                if p2 {
                    for fi10 := 0; fi10 < 11; fi10++ {
                        switch "\n\t" {
                        case "foo bar":
                            v9 = append(v9, fi8)
                        }
                        p1 += (fi10 >> ((0) & 63))
                        if false {
                            p2 = (p2 && ("foo bar" != "a"))
                            v11 := GenS2{F0: p0}
                            p2 = false
                            _ = v11
                        } else {
                            if p2 {
                                var v12 uint8 = 246
                                _ = v12
                            } else {
                                if p2 {
                                    var v13 bool = false
                                    if p2 {
                                        v13 = ((1 <= 145) || v13)
                                        fmt.Println(p0, p1)
                                        var v14 float32 = 0.0
                                        _ = v14
                                    }
                                    if p2 {
                                        var v15 string = "\n\t"
                                        var v16 int = (fi10 + p1)
                                        var v17 bool = (v13 || (!(true || false)))
                                        _ = v15
                                        _ = v16
                                        _ = v17
                                    }
                                    _ = v13
                                }
                                if p2 {
                                    fmt.Println(fi10)
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                } else {
                                    if len(v9) > 0 {
                                        v9[0] = p1
                                    }
                                }
                            }
                            fmt.Println(p0)
                        }
                        _ = fi10
                    }
                    for fi18 := 0; fi18 < 13; fi18++ {
                        v19 := GenS1{F0: -1, F1: -1}
                        _ = fi18
                        _ = v19
                    }
                    p2 = (54 <= 192)
                } else {
                    var v20 uint64 = (p0 * (p0 << ((8) & 63)))
                    p0 = (^(p0))
                    var v21 uint64 = 1
                    _ = v20
                    _ = v21
                }
            case 127:
                if len(v9) > 0 {
                    v9[0] = (p1 - (p1 << ((1) & 63)))
                }
            default:
                p2 = p2
            }
            switch p1 {
            case 1:
                v22 := GenS3{F0: p0, F1: 696, F2: 42}
                _ = v22
            }
            _ = fi8
            _ = v9
        }
    }
    if p2 {
        fmt.Println(p1)
        var v23 float32 = 0.1
        _ = v23
    } else {
        p1 += (p1 + p1)
        var v24 int32 = -54
        fmt.Println(p2, p2, p0)
        _ = v24
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return true
}
func genFn1(p0 string, p1 int32, p2 string) uint16 {
    switch "hello" {
    case "hello":
        fmt.Println(p1, p0, p0)
    case "foo bar":
        for fi25 := 0; fi25 < 8; fi25++ {
            fmt.Println(p2)
            _ = fi25
        }
    default:
        var v26 int8 = 93
        _ = v26
    }
    if (0.1 == (0.0 / 1.5)) {
        var v27 uint = 4294967295
        _ = v27
    }
    v28 := []int{-276, 4294967295, 32768, 0}
    switch "" {
    case "a":
        var v29 bool = true
        _ = v29
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    return 870
}

func main() {
    var v1 bool = false
    v2 := make(map[int]uint32)
    var v3 float64 = 1.5
    var v4 float32 = ((1e10 + -2.25) - 0.1)
    v5 := make(map[string]uint32)
    var v30 int8 = 42
    v1 = true
    if true {
        v4 = v4
    }
    v4 += v4
    var v31 string = "foo bar"
    for fi32 := 0; fi32 < 14; fi32++ {
        for fi33 := 0; fi33 < 1; fi33++ {
            _ = len(v2)
            _ = fi33
        }
        v5[("" + "hello")] = 2147483647
        _ = fi32
    }
    var v34 int8 = (-(v30))
    v4 += v4
    v30 += (v34 ^ v34)
    for fi35 := 0; fi35 < 10; fi35++ {
        for fi36 := 0; fi36 < 12; fi36++ {
            v37 := []int{}
            var v38 int8 = (v30 & v30)
            var v39 uint64 = 1
            _ = fi36
            _ = v37
            _ = v38
            _ = v39
        }
        for fi40 := 0; fi40 < 10; fi40++ {
            fmt.Println(v31, v30, v31)
            _ = fi40
        }
        v1 = v1
        _ = fi35
    }
    var v41 int64 = -643
    v42 := GenS4{F0: v41, F1: 32767}
    for fi43 := 0; fi43 < 20; fi43++ {
        v44 := GenS5{F0: 4294967295, F1: -874}
        if ((!true) || v1) {
            if v1 {
                switch (fi43 / ((fi43) | 1)) {
                case 9223372036854775807:
                    if false {
                        if (!(475 != 32767)) {
                            v45 := GenS1{F0: v30, F1: 868}
                            v46 := []int{fi43, 0, (fi43 % ((-1) | 1)), (fi43 | fi43)}
                            _ = v45
                            _ = v46
                        } else {
                            fmt.Println(v34)
                        }
                        v47 := []int{(fi43 / ((fi43) | 1)), fi43, 1}
                        v4 -= (0.1 * v4)
                        _ = v47
                    } else {
                        if true {
                            var v48 uint = 18446744073709551615
                            v49 := GenS6{F0: v1, F1: 42, F2: 42}
                            _ = v48
                            _ = v49
                        }
                    }
                }
                fmt.Println(v41)
            } else {
                v31 = v31
                switch fi43 {
                case 2147483647:
                    for fi50 := 0; fi50 < 19; fi50++ {
                        v51 := GenS1{F0: (v34 >> ((809) & 63)), F1: 42}
                        var v52 string = v31
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                default:
                    switch ("" + (v31 + (v31 + "a"))) {
                    case "\n\t":
                        v30 -= 1
                    }
                }
                v1 = (!("\n\t" <= ("foo bar" + v31)))
            }
        } else {
            v41 = (-(v41))
        }
        _ = fi43
        _ = v44
    }
    v31 = ("hello" + ("hello" + v31))
    if (!false) {
        for fi53 := 0; fi53 < 9; fi53++ {
            for fi54 := 0; fi54 < 18; fi54++ {
                var v55 float32 = v4
                v4 += -2.25
                var v56 bool = ("foo bar" > (v31 + ((v31 + ("hello" + v31)) + v31)))
                _ = fi54
                _ = v55
                _ = v56
            }
            _ = fi53
        }
    }
    var v57 int = 32767
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v57
}
