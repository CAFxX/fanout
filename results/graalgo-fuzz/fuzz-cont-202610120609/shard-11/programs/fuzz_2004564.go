package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 int16; F2 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64; F1 uint8 }
type GenS7 struct { F0 uint64; F1 int; F2 uint }
type GenS8 struct { F0 int64; F1 uint8; F2 bool }
func genFn0(p0 uint64) string {
    if true {
        v5 := GenS1{F0: -1, F1: false}
        p0 -= p0
        _ = v5
    }
    switch 9223372036854775807 {
    case -2147483648:
        for fi6 := 0; fi6 < 14; fi6++ {
            v7 := GenS2{F0: fi6}
            p0 += p0
            _ = fi6
            _ = v7
        }
    case 32768:
        for fi8 := 0; fi8 < 4; fi8++ {
            for fi9 := 0; fi9 < 16; fi9++ {
                if true {
                    var v10 float32 = (-2.25 + -2.25)
                    _ = v10
                } else {
                    if true {
                        v11 := GenS2{F0: (fi8 + fi9)}
                        p0 = p0
                        _ = v11
                    } else {
                        var v12 uint = 4294967295
                        _ = v12
                    }
                    switch fi9 {
                    case 127:
                        var v13 int32 = 1
                        _ = v13
                    case 4294967295:
                        var v14 int32 = 109
                        _ = v14
                    case 32767:
                        p0 -= p0
                    }
                }
                _ = fi9
            }
            fmt.Println(p0, p0)
            _ = fi8
        }
    default:
        v15 := []int{}
        _ = v15
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        p0 += (p0 % (((^(p0))) | 1))
        v17 := GenS1{F0: (fi16 >> ((939) & 63)), F1: (true || (-886 >= 113))}
        _ = fi16
        _ = v17
    }
    _ = p0
    return ""
}
func genFn1() uint64 {
    if false {
        for fi18 := 0; fi18 < 15; fi18++ {
            v19 := GenS3{F0: 42}
            var v20 float64 = 0.1
            _ = fi18
            _ = v19
            _ = v20
        }
        var v21 int16 = 32767
        _ = v21
    }
    v22 := GenS4{F0: 4294967295, F1: 1, F2: "a"}
    for fi23 := 0; fi23 < 4; fi23++ {
        var v24 uint8 = 0
        v25 := []int{fi23, fi23, fi23}
        v26 := GenS2{F0: fi23}
        _ = fi23
        _ = v24
        _ = v25
        _ = v26
    }
    var v27 uint16 = 42
    _ = v22
    _ = v27
    return 2
}

func main() {
    v1 := []string{("\n\t" + "\n\t"), "\n\t", ("a" + "\n\t")}
    var v2 bool = ((false || true) && true)
    var v3 int32 = -196
    var v4 uint8 = 42
    v4 -= (v4 + 0)
    v2 = true
    for fi28 := 0; fi28 < 4; fi28++ {
        var v29 int16 = 370
        v4 = v4
        v3 += (v3 >> ((0) & 63))
        _ = fi28
        _ = v29
    }
    var v30 int64 = -113
    v31 := GenS1{F0: 0, F1: true}
    _ = len(v1)
    if (("" + "") >= "a") {
        var v32 int64 = 0
        if ((((3.14159 + 1e10) * 0.1) - 3.14159) < (1e10 - float32(-2.25))) {
            var v33 int64 = 42
            v4 += 128
            var v34 uint16 = 65535
            _ = v33
            _ = v34
        } else {
            for fi35 := 0; fi35 < 4; fi35++ {
                v2 = v2
                if (!((^(v30)) == v32)) {
                    v36 := GenS3{F0: -9223372036854775808}
                    _ = v36
                } else {
                    var v37 uint = 646
                    v38 := GenS5{F0: -32768}
                    var v39 int32 = (v3 | v3)
                    _ = v37
                    _ = v38
                    _ = v39
                }
                for fi40 := 0; fi40 < 14; fi40++ {
                    var v41 int = fi40
                    var v42 float64 = -2.25
                    switch fi40 {
                    case 678:
                        v43 := GenS4{F0: 4294967295, F1: -1, F2: ((("a" + "\n\t") + (("\n\t" + ("\n\t" + "a")) + "hello")) + ("hello" + "foo bar"))}
                        _ = v43
                    default:
                        v2 = v2
                    }
                    _ = fi40
                    _ = v41
                    _ = v42
                }
                _ = fi35
            }
            var v44 int16 = -32768
            _ = v44
        }
        switch 373 {
        case 128:
            for fi45 := 0; fi45 < 5; fi45++ {
                fmt.Println(fi45, fi45, v4)
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
                _ = fi45
            }
        case 4294967295:
            v4 += (v4 % ((v4) | 1))
        case 0:
            v46 := GenS6{F0: (v30 << ((18446744073709551615) & 63)), F1: (^(v4))}
            _ = v46
        }
        _ = v32
    }
    var v47 uint8 = v4
    v47 -= v47
    var v48 bool = (574 >= 32767)
    v49 := GenS7{F0: 18446744073709551615, F1: -1, F2: 445}
    var v50 uint64 = 251
    for fi51 := 0; fi51 < 18; fi51++ {
        var v52 int64 = int64(v4)
        fmt.Println(fi51, v4, v3)
        fmt.Println(v30, v47)
        _ = fi51
        _ = v52
    }
    switch "foo bar" {
    case "foo bar":
        var v53 uint8 = (v47 - (v4 & 62))
        _ = v53
    default:
        fmt.Println(v48, v4, v50)
    }
    switch 32768 {
    case 0:
        v2 = (!((("" + "\n\t") + "a") < "hello"))
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if v48 {
        if (-639 <= v3) {
            if true {
                var v54 int32 = 1
                _ = v54
            } else {
                switch "a" {
                case "a":
                    fmt.Println(v47, v3)
                case "":
                    fmt.Println(v50, v30)
                }
                v55 := GenS8{F0: v30, F1: (v4 & (^(v47))), F2: (!true)}
                _ = v55
            }
        }
    } else {
        v56 := GenS7{F0: (v50 + v50), F1: -992, F2: 18446744073709551615}
        v47 -= (v4 % ((v4) | 1))
        v1 = append(v1, "")
        _ = v56
    }
    v47 = (v4 >> ((65535) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v50
}
