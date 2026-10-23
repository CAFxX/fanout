package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 bool; F2 int32 }
type GenS3 struct { F0 int8; F1 uint16; F2 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 int64 }
func genFn0(p0 uint16, p1 string) uint {
    var v4 bool = true
    v5 := GenS1{F0: 18446744073709551615}
    v6 := GenS2{F0: 0, F1: (42 > 2147483647), F2: 2147483647}
    for fi7 := 0; fi7 < 1; fi7++ {
        v4 = true
        var v8 float32 = 3.14159
        v9 := GenS3{F0: 65, F1: (p0 / (((p0 * p0)) | 1)), F2: 42}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 2
}
func genFn1(p0 int16) uint32 {
    p0 -= -878
    v10 := GenS3{F0: 1, F1: 32767, F2: p0}
    if (false && (((587 <= 172) || true) && (!false))) {
        p0 -= p0
    }
    _ = p0
    _ = v10
    return 768
}

func main() {
    v1 := []uint32{4294967295}
    v2 := make(map[int]string)
    var v3 int = -2147483648
    if ((1 == 42) && true) {
        switch "a" {
        case "foo bar":
            switch (v3 ^ (v3 | (v3 % (((^(v3))) | 1)))) {
            case 32767:
                fmt.Println(v3)
            case 0:
                if false {
                    for fi11 := 0; fi11 < 15; fi11++ {
                        v3 -= v3
                        var v12 int16 = 42
                        var v13 string = ("a" + "foo bar")
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    v14 := GenS3{F0: -128, F1: 42, F2: 447}
                    _ = v14
                } else {
                    v15 := GenS3{F0: -99, F1: 1, F2: 0}
                    _ = v15
                }
            case -9223372036854775808:
                for fi16 := 0; fi16 < 3; fi16++ {
                    var v17 float64 = 1e10
                    switch "foo bar" {
                    case "a":
                        fmt.Println(v17, fi16, fi16)
                    case "foo bar":
                        v17 += (v17 / (0.0 / v17))
                    }
                    _ = fi16
                    _ = v17
                }
            }
        case "hello":
            if false {
                for fi18 := 0; fi18 < 5; fi18++ {
                    v3 = 32767
                    for fi19 := 0; fi19 < 3; fi19++ {
                        if (-119 == 0) {
                            var v20 string = "a"
                            v21 := v2[fi18]
                            v22 := GenS4{F0: "hello"}
                            _ = v20
                            _ = v21
                            _ = v22
                        } else {
                            var v23 int = (-(v3))
                            var v24 int32 = 947
                            _ = v23
                            _ = v24
                        }
                        v1 = append(v1, 0)
                        fmt.Println(fi19, fi18, fi19)
                        _ = fi19
                    }
                    _ = fi18
                }
                v25 := GenS3{F0: 81, F1: 867, F2: 429}
                _ = v25
            }
        case "a":
            if false {
                for fi26 := 0; fi26 < 14; fi26++ {
                    var v27 int = fi26
                    for fi28 := 0; fi28 < 9; fi28++ {
                        v29 := v2[fi28]
                        _ = fi28
                        _ = v29
                    }
                    var v30 float32 = ((((0.0 + 0.0) + (0.1 + 0.0)) - (3.14159 - float32(1.5))) + 1e10)
                    _ = fi26
                    _ = v27
                    _ = v30
                }
                v2[v3] = ""
                for fi31 := 0; fi31 < 4; fi31++ {
                    for fi32 := 0; fi32 < 16; fi32++ {
                        var v33 int = 127
                        var v34 int = v3
                        fmt.Println(fi32, fi32, v34)
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                    _ = fi31
                }
            }
        }
        var v35 int16 = -491
        delete(v2, (v3 << ((171) & 63)))
        _ = v35
    }
    for fi36 := 0; fi36 < 9; fi36++ {
        var v37 uint64 = 2
        _ = fi36
        _ = v37
    }
    if true {
        var v38 float32 = 3.14159
        if false {
            for fi39 := 0; fi39 < 9; fi39++ {
                v40 := GenS4{F0: ((((("" + "a") + ("a" + ("" + ("" + "foo bar")))) + ("foo bar" + ("\n\t" + "foo bar"))) + "a") + "a")}
                v3 += (-(v3))
                _ = fi39
                _ = v40
            }
            if false {
                delete(v2, v3)
                var v41 int64 = -9223372036854775808
                _ = v41
            } else {
                var v42 float64 = 0.1
                var v43 int32 = -738
                for fi44 := 0; fi44 < 2; fi44++ {
                    v2[2147483647] = ""
                    _ = fi44
                }
                _ = v42
                _ = v43
            }
        } else {
            _ = len(v1)
        }
        v3 -= v3
        _ = v38
    } else {
        for fi45 := 0; fi45 < 7; fi45++ {
            if true {
                if false {
                    v3 = (v3 * -9223372036854775808)
                    var v46 int8 = 1
                    _ = v46
                } else {
                    var v47 uint32 = 65535
                    switch -1 {
                    case 32767:
                        for fi48 := 0; fi48 < 20; fi48++ {
                            v47 = v47
                            v49 := GenS1{F0: 2}
                            var v50 string = "a"
                            _ = fi48
                            _ = v49
                            _ = v50
                        }
                    case 128:
                        fmt.Println(v3, v47)
                    }
                    v3 = v3
                    _ = v47
                }
                var v51 int32 = -789
                _ = v51
            } else {
                if (0 > 805) {
                    v3 -= (fi45 - v3)
                    var v52 int = v3
                    v52 = fi45
                    _ = v52
                } else {
                    fmt.Println(v3)
                }
            }
            _ = fi45
        }
        for fi53 := 0; fi53 < 13; fi53++ {
            v54 := GenS5{F0: -32768, F1: 298}
            if (-246 <= 881) {
                fmt.Println(v3, v3)
                _ = len(v1)
                v2[fi53] = "hello"
            }
            fmt.Println(v3)
            _ = fi53
            _ = v54
        }
        v3 = v3
    }
    v2[v3] = ("" + ("" + "\n\t"))
    v2[(^(v3))] = ""
    var v55 float32 = 1e10
    switch "hello" {
    case "a":
        var v56 uint32 = 65535
        _ = v56
    case "\n\t":
        if false {
            var v57 int8 = 127
            v55 += float32(((3.14159 - 1e10) - 0.1))
            v58 := GenS6{F0: 1}
            _ = v57
            _ = v58
        }
    case "hello":
        fmt.Println(v55, v55)
    }
    fmt.Println(v3)
    var v59 int64 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v55:", v55)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v55
    _ = v59
}
