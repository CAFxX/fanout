package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 int8; F2 uint8 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 uint16; F1 bool; F2 int8 }
func genFn0(p0 int32) int16 {
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 7; fi6++ {
            p0 = p0
            var v7 bool = true
            v8 := []int{}
            _ = fi6
            _ = v7
            _ = v8
        }
        switch fi5 {
        case 42:
            v9 := []int{}
            _ = v9
        }
        _ = fi5
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        if true {
            var v11 float64 = 1.5
            _ = v11
        }
        _ = fi10
    }
    _ = p0
    return 0
}
func genFn1(p0 int64, p1 uint8, p2 uint) uint32 {
    var v12 string = ("" + ("hello" + "a"))
    v12 = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 10
}

func main() {
    v1 := []string{((("\n\t" + "a") + ("hello" + "hello")) + "foo bar"), "hello", ("hello" + "foo bar"), ("a" + "hello")}
    var v2 uint64 = 4294967295
    v3 := []int{32767, 328, 9223372036854775807}
    var v4 uint = 4294967295
    for fi13 := 0; fi13 < 20; fi13++ {
        if ("" != "foo bar") {
            v2 += (^(v2))
            if (-1 < 0) {
                switch "foo bar" {
                case "\n\t":
                    _ = len(v1)
                case "foo bar":
                    var v14 uint16 = 32767
                    _ = v14
                case "a":
                    var v15 uint8 = 2
                    _ = v15
                }
                v16 := GenS1{F0: 42}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v16
            } else {
                v17 := GenS2{F0: -9223372036854775808, F1: -1, F2: uint8(v2)}
                _ = v17
            }
            var v18 bool = true
            _ = v18
        }
        if (!(87 == 128)) {
            for fi19 := 0; fi19 < 18; fi19++ {
                if false {
                    v4 -= v4
                    var v20 bool = (((float32(0.1) / 1.0) == 1.5) || true)
                    if v20 {
                        if len(v1) > 0 {
                            v1[0] = "a"
                        }
                    } else {
                        var v21 bool = (v4 > (v4 >> (((v4 | v4)) & 63)))
                        if v21 {
                            v21 = ("hello" != (("a" + ("hello" + "foo bar")) + "a"))
                            if (128 != 2) {
                                if (v4 > v4) {
                                    var v22 bool = false
                                    var v23 uint = (v4 | v4)
                                    fmt.Println(v23, v20, v2)
                                    if (!false) {
                                        if v22 {
                                            fmt.Println(v21)
                                            var v24 int = (fi19 >> ((2) & 63))
                                            var v25 uint64 = v2
                                            _ = v24
                                            _ = v25
                                        }
                                        fmt.Println(v23, v2)
                                        fmt.Println(v20, fi13)
                                    } else {
                                        v21 = v20
                                    }
                                    _ = v22
                                    _ = v23
                                } else {
                                    var v26 uint = v4
                                    var v27 float64 = (0.0 + 0.0)
                                    var v28 uint8 = 159
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                            }
                        }
                        _ = v21
                    }
                    var v29 uint = 4294967295
                    _ = v20
                    _ = v29
                } else {
                    v4 -= 4294967295
                    switch -1 {
                    case 127:
                        var v30 uint8 = 2
                        _ = v30
                    case 32767:
                        var v31 float32 = float32(0.0)
                        _ = v31
                    case 32768:
                        v2 = (v2 & v2)
                    }
                    v32 := GenS3{F0: "hello", F1: 81}
                    var v33 float32 = 0.0
                    if true {
                        fmt.Println(v2, fi19)
                        fmt.Println(fi19)
                        if true {
                            if true {
                                fmt.Println(fi13)
                                v33 += v33
                                v33 += v33
                                switch "hello" {
                                case "a":
                                    var v34 int8 = 127
                                    _ = v34
                                case "\n\t":
                                    fmt.Println(fi13, v33, v33)
                                }
                            }
                        } else {
                            v33 = (float32((((0.0 + 3.14159) + 0.1) * ((-2.25 / 3.14159) - (3.14159 / ((float64(v33) + -2.25) + 3.14159))))) * -2.25)
                        }
                    }
                    _ = v32
                    _ = v33
                }
                _ = fi19
            }
            if true {
                v35 := GenS1{F0: 798}
                v36 := GenS1{F0: 9223372036854775807}
                _ = v35
                _ = v36
            }
            for fi37 := 0; fi37 < 17; fi37++ {
                v4 -= (v4 << (((v4 << ((v4) & 63))) & 63))
                _ = fi37
            }
        }
        switch (fi13 ^ fi13) {
        case 127:
            var v38 int16 = 1
            _ = v38
        case 32767:
            fmt.Println(v2, v4)
        }
        _ = fi13
    }
    for fi39 := 0; fi39 < 14; fi39++ {
        var v40 uint16 = 782
        _ = fi39
        _ = v40
    }
    v2 += 4294967295
    if (-1 < -414) {
        _ = len(v1)
        var v41 int32 = 1
        var v42 uint8 = 155
        _ = v41
        _ = v42
    } else {
        v43 := GenS1{F0: 9223372036854775807}
        switch 0 {
        case 0:
            for fi44 := 0; fi44 < 11; fi44++ {
                v4 -= 227
                if false {
                    var v45 bool = (v4 != v4)
                    v46 := GenS3{F0: "foo bar", F1: 255}
                    var v47 int = (fi44 >> ((v4) & 63))
                    _ = v45
                    _ = v46
                    _ = v47
                }
                v4 += (v4 << (((v4 + (v4 << ((v4) & 63)))) & 63))
                _ = fi44
            }
        }
        for fi48 := 0; fi48 < 7; fi48++ {
            for fi49 := 0; fi49 < 9; fi49++ {
                fmt.Println(fi48, fi48)
                var v50 float64 = float64((0.0 + (3.14159 * 1e10)))
                v51 := GenS3{F0: "", F1: 255}
                if (!(true || false)) {
                    switch "hello" {
                    case "\n\t":
                        v50 += v50
                    }
                } else {
                    var v52 string = ""
                    _ = v52
                }
                var v53 string = "a"
                fmt.Println(v4, v50, v50)
                _ = fi49
                _ = v50
                _ = v51
                _ = v53
            }
            switch 128 {
            case -1:
                for fi54 := 0; fi54 < 14; fi54++ {
                    v2 += 861
                    var v55 uint16 = 404
                    if true {
                        var v56 float32 = -2.25
                        _ = v56
                    } else {
                        var v57 string = (("foo bar" + "") + "a")
                        if (0.0 > 0.1) {
                            v4 = 4294967295
                            switch ("foo bar" + v57) {
                            case "a":
                                v58 := GenS1{F0: int64(v2)}
                                _ = v58
                            case "":
                                fmt.Println(v57)
                            case "\n\t":
                                var v59 uint8 = 0
                                _ = v59
                            }
                        } else {
                            v60 := GenS4{F0: 2, F1: false, F2: -128}
                            fmt.Println(v2)
                            _ = v60
                        }
                        _ = v57
                    }
                    _ = fi54
                    _ = v55
                }
            case -376:
                for fi61 := 0; fi61 < 13; fi61++ {
                    if false {
                        var v62 float64 = 3.14159
                        var v63 uint = 2
                        _ = v62
                        _ = v63
                    }
                    var v64 int16 = -32768
                    _ = fi61
                    _ = v64
                }
            default:
                if true {
                    var v65 int8 = 1
                    fmt.Println(fi48, v4, fi48)
                    _ = v65
                }
            }
            if false {
                var v66 int8 = -116
                v67 := GenS4{F0: 255, F1: false, F2: 126}
                v68 := GenS1{F0: -434}
                _ = v66
                _ = v67
                _ = v68
            } else {
                if len(v3) > 0 {
                    v3[0] = (fi48 * fi48)
                }
                var v69 uint64 = v2
                if true {
                    v2 = 0
                    v70 := GenS1{F0: 0}
                    _ = v70
                }
                _ = v69
            }
            _ = fi48
        }
        _ = v43
    }
    v2 = 0
    _ = len(v1)
    v4 = (v4 << ((v4) & 63))
    _ = len(v3)
    v71 := GenS2{F0: -663, F1: int8(v4), F2: 153}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v71
}
