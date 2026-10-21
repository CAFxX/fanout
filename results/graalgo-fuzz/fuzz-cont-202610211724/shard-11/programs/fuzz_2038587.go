package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int8; F1 int8; F2 bool }
func genFn0(p0 string, p1 int64) string {
    var v6 float32 = 1e10
    switch p0 {
    case "foo bar":
        if true {
            if false {
                var v7 int64 = p1
                for fi8 := 0; fi8 < 15; fi8++ {
                    if (true || (!((209 == 0) || (false || false)))) {
                        var v9 bool = (true && (384 != -448))
                        for fi10 := 0; fi10 < 17; fi10++ {
                            p1 -= (p1 | (v7 * (v7 ^ v7)))
                            v11 := []int{(-(fi10)), (fi10 * fi10)}
                            _ = fi10
                            _ = v11
                        }
                        switch fi8 {
                        case 713:
                            var v12 int = 4294967295
                            _ = v12
                        case 42:
                            for fi13 := 0; fi13 < 2; fi13++ {
                                v14 := []int{}
                                if true {
                                    var v15 bool = v9
                                    _ = v15
                                } else {
                                    fmt.Println(v7)
                                    switch p0 {
                                    case "foo bar":
                                        if len(v14) > 0 {
                                            v14[0] = -9223372036854775808
                                        }
                                    case "\n\t":
                                        var v16 int = (fi8 / ((fi8) | 1))
                                        _ = v16
                                    case "a":
                                        v6 = v6
                                    }
                                }
                                _ = fi13
                                _ = v14
                            }
                        case 1:
                            for fi17 := 0; fi17 < 12; fi17++ {
                                v6 -= v6
                                v18 := []int{fi8, fi17}
                                v19 := GenS1{F0: -390, F1: 0, F2: 4294967295}
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                        }
                        _ = v9
                    }
                    var v20 int16 = 0
                    _ = fi8
                    _ = v20
                }
                _ = v7
            } else {
                var v21 uint64 = 2
                if ("hello" > p0) {
                    var v22 int32 = 0
                    _ = v22
                }
                _ = v21
            }
        }
    }
    if true {
        switch (p0 + ("a" + p0)) {
        case "foo bar":
            for fi23 := 0; fi23 < 9; fi23++ {
                fmt.Println(p0, p1, v6)
                _ = fi23
            }
        case "hello":
            v24 := []int{}
            _ = v24
        case "a":
            if true {
                v25 := GenS1{F0: 0, F1: 42, F2: 4294967295}
                switch p0 {
                case "\n\t":
                    var v26 uint8 = 58
                    _ = v26
                }
                if true {
                    v27 := GenS1{F0: 352, F1: -2147483648, F2: 2}
                    for fi28 := 0; fi28 < 16; fi28++ {
                        var v29 float64 = -2.25
                        v30 := []int{(fi28 - -588)}
                        v31 := GenS2{F0: 32767, F1: 65535}
                        _ = fi28
                        _ = v29
                        _ = v30
                        _ = v31
                    }
                    _ = v27
                } else {
                    p1 -= (^(p1))
                }
                _ = v25
            }
        default:
            if (-485 >= 2147483647) {
                var v32 float64 = 3.14159
                v6 += ((v6 / v6) - ((v6 / (float32(((((((float64(v6) * v32) - 3.14159) - float64(0.0)) / float64(0.1)) + (v32 + 3.14159)) * 0.1)) - 1.5)) - v6))
                if (!(!true)) {
                    p0 = (p0 + "\n\t")
                }
                _ = v32
            } else {
                var v33 string = "\n\t"
                for fi34 := 0; fi34 < 12; fi34++ {
                    var v35 int16 = 1000
                    v33 = ""
                    _ = fi34
                    _ = v35
                }
                switch v33 {
                case "a":
                    var v36 int32 = 0
                    _ = v36
                default:
                    var v37 uint32 = 485
                    _ = v37
                }
                _ = v33
            }
        }
        v38 := []int{0}
        if false {
            switch -1 {
            case 0:
                v39 := GenS3{F0: 104, F1: 127, F2: (true || false)}
                _ = v39
            case 32768:
                switch 357 {
                case -935:
                    switch "a" {
                    case "hello":
                        _ = len(v38)
                    default:
                        var v40 float64 = -2.25
                        _ = v40
                    }
                }
            case 32767:
                for fi41 := 0; fi41 < 4; fi41++ {
                    var v42 float64 = 3.14159
                    var v43 float32 = 1e10
                    _ = fi41
                    _ = v42
                    _ = v43
                }
            }
        } else {
            switch -439 {
            case 2147483647:
                fmt.Println(p0)
            case 9223372036854775807:
                if (v6 < (1.5 + 1.5)) {
                    _ = len(v38)
                }
            case 0:
                for fi44 := 0; fi44 < 17; fi44++ {
                    p1 -= p1
                    _ = fi44
                }
            }
            v45 := GenS3{F0: 127, F1: 100, F2: false}
            _ = v45
        }
        _ = v38
    }
    _ = p0
    _ = p1
    _ = v6
    return p0
}
func genFn1(p0 int64) bool {
    p0 += p0
    p0 -= p0
    p0 -= p0
    v46 := GenS3{F0: -1, F1: 116, F2: (!false)}
    var v47 int = 0
    _ = p0
    _ = v46
    _ = v47
    return (!false)
}

func main() {
    var v1 uint16 = 1
    var v2 float32 = 0.0
    var v3 uint8 = 128
    var v4 int64 = 751
    var v5 float32 = (float32(-2.25) + float32(0.1))
    var v48 bool = (!true)
    if true {
        for fi49 := 0; fi49 < 4; fi49++ {
            fmt.Println(v2, v48)
            fmt.Println(v2, fi49, v5)
            _ = fi49
        }
        v50 := []int{523}
        v2 -= 0.0
        _ = v50
    }
    v51 := []int{-644, 4294967295}
    for fi52 := 0; fi52 < 6; fi52++ {
        if len(v51) > 0 {
            v51[0] = -1
        }
        v1 -= v1
        _ = fi52
    }
    for fi53 := 0; fi53 < 1; fi53++ {
        for fi54 := 0; fi54 < 2; fi54++ {
            v55 := GenS1{F0: -32768, F1: 2147483647, F2: 2}
            v56 := GenS1{F0: 1000, F1: -2147483648, F2: 65535}
            _ = fi54
            _ = v55
            _ = v56
        }
        _ = fi53
    }
    switch ((("hello" + ("hello" + ((("a" + "") + "\n\t") + ""))) + "\n\t") + "") {
    case "hello":
        v57 := GenS3{F0: -12, F1: 100, F2: v48}
        _ = v57
    case "":
        switch 0 {
        case -448:
            v5 -= (v2 * float32(-2.25))
        default:
            var v58 int8 = 1
            _ = v58
        }
    }
    v59 := GenS2{F0: (v1 << ((1) & 63)), F1: 0}
    if (9 == -1) {
        if len(v51) > 0 {
            v51[0] = -572
        }
        if len(v51) >= 2 {
            v51 = v51[1:len(v51)]
        }
    } else {
        switch "foo bar" {
        case "foo bar":
            for fi60 := 0; fi60 < 20; fi60++ {
                var v61 string = "foo bar"
                _ = fi60
                _ = v61
            }
        default:
            v48 = (159 == 65535)
        }
        var v62 uint = 31
        _ = v62
    }
    switch -1 {
    case 2147483647:
        var v63 int = 9223372036854775807
        _ = v63
    case 1:
        if len(v51) >= 2 {
            v51 = v51[1:len(v51)]
        }
    default:
        var v64 int8 = 127
        _ = v64
    }
    for fi65 := 0; fi65 < 1; fi65++ {
        v66 := GenS1{F0: -1, F1: -1, F2: 2}
        _ = fi65
        _ = v66
    }
    if len(v51) >= 2 {
        v51 = v51[1:len(v51)]
    }
    var v67 uint8 = (^(v3))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v51
    _ = v59
    _ = v67
}
