package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int }
func genFn0(p0 string, p1 int) bool {
    v7 := []int{p1, p1}
    for fi8 := 0; fi8 < 3; fi8++ {
        switch (p0 + p0) {
        case "\n\t":
            v9 := GenS1{F0: (-(p1)), F1: -53}
            _ = v9
        case "a":
            for fi10 := 0; fi10 < 1; fi10++ {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                v7 = append(v7, (fi8 & fi10))
                var v11 bool = true
                _ = fi10
                _ = v11
            }
        }
        var v12 uint32 = 42
        if len(v7) > 0 {
            v7[0] = (fi8 | -591)
        }
        _ = fi8
        _ = v12
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v13 uint16 = 65535
    switch p0 {
    case "\n\t":
        var v14 int64 = -9223372036854775808
        _ = v14
    case "foo bar":
        for fi15 := 0; fi15 < 3; fi15++ {
            var v16 bool = (100 > 1)
            _ = fi15
            _ = v16
        }
    case "":
        for fi17 := 0; fi17 < 13; fi17++ {
            for fi18 := 0; fi18 < 2; fi18++ {
                if true {
                    p1 += p1
                    var v19 int32 = 2147483647
                    if len(v7) > 0 {
                        v7[0] = (^(p1))
                    }
                    v7 = append(v7, 0)
                    _ = v19
                }
                var v20 int32 = 1
                _ = fi18
                _ = v20
            }
            _ = fi17
        }
    default:
        for fi21 := 0; fi21 < 13; fi21++ {
            if ((fi21 + fi21) < fi21) {
                p0 = p0
                switch fi21 {
                case 1:
                    fmt.Println(v13, p1, v13)
                }
            } else {
                var v22 int8 = 87
                var v23 bool = (637 <= 316)
                _ = v22
                _ = v23
            }
            _ = fi21
        }
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v7
    return true
}

func main() {
    var v1 int64 = -138
    var v2 bool = true
    var v3 string = ""
    v4 := []uint32{296, 895, 42, 553}
    v5 := []uint32{2147483648, 1, 0}
    v6 := []string{("hello" + "hello"), v3, ("hello" + (((v3 + ("\n\t" + v3)) + "\n\t") + "foo bar"))}
    var v24 int16 = 209
    if v2 {
        v25 := GenS2{F0: "a", F1: 65535}
        v2 = false
        _ = v25
    } else {
        var v26 uint = 42
        var v27 int16 = v24
        _ = v26
        _ = v27
    }
    switch (v3 + v3) {
    case "":
        var v28 int64 = 0
        _ = v28
    case "hello":
        var v29 int64 = (v1 ^ 0)
        _ = v29
    }
    switch "foo bar" {
    case "foo bar":
        var v30 float64 = 0.1
        _ = v30
    default:
        v31 := GenS2{F0: v3, F1: 32767}
        _ = v31
    }
    var v32 uint32 = 42
    switch "hello" {
    case "foo bar":
        switch "foo bar" {
        case "a":
            v24 -= 1000
        case "\n\t":
            v32 = (v32 / (((v32 / (((v32 - (v32 & v32))) | 1))) | 1))
        case "":
            v32 = 1
        default:
            var v33 uint64 = 4294967295
            _ = v33
        }
    case "":
        switch "\n\t" {
        case "\n\t":
            var v34 uint64 = 2
            _ = v34
        case "a":
            if len(v5) > 0 {
                v5[0] = v32
            }
        default:
            if v2 {
                var v35 int32 = 1
                var v36 int8 = 1
                _ = v35
                _ = v36
            } else {
                if ((v32 - v32) < v32) {
                    v32 = (v32 << ((532) & 63))
                    if ((!v2) || v2) {
                        if len(v4) > 0 {
                            v4[0] = v32
                        }
                        v37 := GenS2{F0: v3, F1: 27}
                        _ = v37
                    }
                    if (float32(-2.25) <= 0.0) {
                        switch v3 {
                        case "":
                            _ = len(v5)
                        case "a":
                            switch "" {
                            case "a":
                                v32 += (v32 * 847)
                            case "hello":
                                if len(v5) > 0 {
                                    v5[0] = v32
                                }
                            }
                        case "foo bar":
                            var v38 float32 = 0.1
                            _ = v38
                        default:
                            v1 += v1
                        }
                    } else {
                        for fi39 := 0; fi39 < 2; fi39++ {
                            for fi40 := 0; fi40 < 19; fi40++ {
                                if false {
                                    var v41 int = 1
                                    var v42 int = fi40
                                    _ = v41
                                    _ = v42
                                }
                                _ = fi40
                            }
                            v2 = v2
                            v4 = append(v4, v32)
                            _ = fi39
                        }
                    }
                } else {
                    _ = len(v5)
                    var v43 float64 = 0.0
                    v24 = (v24 << ((4294967295) & 63))
                    _ = v43
                }
                var v44 int64 = v1
                _ = v44
            }
        }
    default:
        switch v3 {
        case "hello":
            for fi45 := 0; fi45 < 8; fi45++ {
                v4 = append(v4, v32)
                _ = fi45
            }
        case "foo bar":
            fmt.Println(v1)
        }
    }
    for fi46 := 0; fi46 < 3; fi46++ {
        var v47 uint8 = 242
        var v48 uint32 = v32
        _ = fi46
        _ = v47
        _ = v48
    }
    var v49 bool = (v24 == v24)
    if len(v5) > 0 {
        v5[0] = v32
    }
    v3 = "\n\t"
    v2 = (v49 || (2 > 0))
    if v2 {
        v1 -= v1
        switch "hello" {
        case "a":
            if true {
                if v49 {
                    _ = len(v5)
                    switch 1 {
                    case 1:
                        if v49 {
                            v5 = append(v5, v32)
                            v49 = v49
                        } else {
                            switch 4294967295 {
                            case 32768:
                                fmt.Println(v3, v2)
                            case 491:
                                fmt.Println(v1)
                            case 343:
                                fmt.Println(v1, v3)
                            }
                            v50 := GenS3{F0: 0}
                            _ = len(v5)
                            _ = v50
                        }
                    case -9223372036854775808:
                        var v51 int32 = int32(v1)
                        _ = v51
                    case 32767:
                        if len(v4) > 0 {
                            v4[0] = (^(v32))
                        }
                    }
                }
            } else {
                var v52 uint64 = 42
                v3 = ((v3 + (v3 + "")) + v3)
                switch 127 {
                case 190:
                    fmt.Println(v49)
                case 0:
                    v1 -= 1
                default:
                    var v53 float64 = ((((1e10 + -2.25) + 0.1) + 1e10) * 1.5)
                    _ = v53
                }
                _ = v52
            }
        default:
            v3 = "\n\t"
        }
        v6 = append(v6, "\n\t")
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v32
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
