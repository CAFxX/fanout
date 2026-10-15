package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string }
func genFn0(p0 uint32, p1 uint16) string {
    p1 = 2
    p1 -= 273
    var v4 float64 = 1.5
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := []int{-9223372036854775808, (fi5 | (fi5 + (fi5 | (fi5 | fi5)))), fi5}
        p1 = p1
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 float32 = (-2.25 / 0.1)
            v9 := GenS1{F0: 2}
            var v10 float32 = v8
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        if true {
            for fi12 := 0; fi12 < 18; fi12++ {
                v13 := GenS1{F0: 4294967295}
                var v14 uint16 = p1
                _ = fi12
                _ = v13
                _ = v14
            }
            if false {
                var v15 float32 = 0.1
                if ((0.0 != (((v15 / v15) / 1e10) + v15)) && (42 != 915)) {
                    var v16 int16 = -959
                    v17 := GenS1{F0: 2}
                    v18 := GenS2{F0: (fi11 & 766), F1: 32768}
                    _ = v16
                    _ = v17
                    _ = v18
                }
                _ = v15
            } else {
                v4 = 0.1
            }
            switch (fi11 & fi11) {
            case 417:
                var v19 int64 = -93
                _ = v19
            case -108:
                var v20 float64 = v4
                _ = v20
            }
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v4
    return "a"
}

func main() {
    var v1 int8 = 100
    var v2 int8 = v1
    v3 := []int{}
    v21 := GenS2{F0: 1, F1: 4294967295}
    var v22 uint16 = 42
    v23 := GenS3{F0: v2}
    if (((42 >= 1) && false) && (797 < 1)) {
        var v24 int32 = -541
        _ = v24
    }
    switch -9223372036854775808 {
    case 32767:
        for fi25 := 0; fi25 < 9; fi25++ {
            var v26 int32 = 2147483647
            var v27 uint32 = 65535
            fmt.Println(v2, v26)
            _ = fi25
            _ = v26
            _ = v27
        }
    case 752:
        for fi28 := 0; fi28 < 1; fi28++ {
            v29 := GenS2{F0: (fi28 + (fi28 >> ((600) & 63))), F1: fi28}
            _ = fi28
            _ = v29
        }
    default:
        if (false || false) {
            if true {
                if (!(!true)) {
                    var v30 int16 = 0
                    var v31 float64 = 3.14159
                    _ = v30
                    _ = v31
                }
            }
            for fi32 := 0; fi32 < 11; fi32++ {
                var v33 float64 = 0.1
                v22 = (v22 * (^(v22)))
                _ = fi32
                _ = v33
            }
        }
    }
    for fi34 := 0; fi34 < 4; fi34++ {
        fmt.Println(v22, v2, v2)
        v1 = v1
        v1 = (v1 ^ (v1 * (v1 | v1)))
        _ = fi34
    }
    for fi35 := 0; fi35 < 10; fi35++ {
        if ((4294967295 == 1) || true) {
            var v36 int8 = -62
            _ = v36
        } else {
            v2 = v1
            _ = len(v3)
        }
        fmt.Println(fi35)
        _ = fi35
    }
    v37 := GenS1{F0: 2}
    if len(v3) > 0 {
        v3[0] = 32768
    }
    if true {
        var v38 uint = 65535
        _ = v38
    }
    for fi39 := 0; fi39 < 4; fi39++ {
        if true {
            var v40 int64 = 1
            v22 -= v22
            for fi41 := 0; fi41 < 17; fi41++ {
                var v42 float32 = 1e10
                _ = fi41
                _ = v42
            }
            _ = v40
        } else {
            var v43 uint16 = (v22 * v22)
            v22 += v22
            v44 := GenS4{F0: "foo bar"}
            _ = v43
            _ = v44
        }
        for fi45 := 0; fi45 < 14; fi45++ {
            switch fi39 {
            case 250:
                if (!true) {
                    var v46 uint16 = v22
                    if (180 > 127) {
                        fmt.Println(v1, v46, fi39)
                        v2 -= (v1 >> ((18446744073709551615) & 63))
                        if (true && true) {
                            switch "" {
                            case "hello":
                                v22 = (v22 >> ((405) & 63))
                            case "foo bar":
                                var v47 int8 = 42
                                _ = v47
                            case "a":
                                var v48 uint32 = 65535
                                fmt.Println(v22, v48)
                                _ = v48
                            default:
                                v49 := GenS3{F0: (v1 << ((65535) & 63))}
                                _ = v49
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v50 string = "a"
                            _ = v50
                        } else {
                            var v51 uint16 = (v22 << ((4294967295) & 63))
                            var v52 int8 = (v1 / (((v2 & (-(v2)))) | 1))
                            _ = v51
                            _ = v52
                        }
                    } else {
                        if true {
                            var v53 int32 = -202
                            _ = v53
                        }
                        if (true || (!(!false))) {
                            v54 := GenS4{F0: "hello"}
                            if false {
                                v55 := GenS3{F0: -1}
                                v2 += (^(v2))
                                v2 += -128
                                _ = v55
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = fi39
                                }
                                if len(v3) > 0 {
                                    v3[0] = fi39
                                }
                                v56 := GenS4{F0: "hello"}
                                _ = v56
                            }
                            v1 = 82
                            _ = v54
                        }
                        fmt.Println(fi39, v2)
                    }
                    _ = v46
                }
            default:
                v1 += -128
            }
            var v57 string = "\n\t"
            _ = fi45
            _ = v57
        }
        _ = fi39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v37
}
