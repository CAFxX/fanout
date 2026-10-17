package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 uint; F2 int8 }
type GenS4 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string; F1 int }
type GenS7 struct { F0 int; F1 string; F2 string }
type GenS8 struct { F0 uint16; F1 int64; F2 int64 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 bool; F1 uint16 }
func genFn0() uint32 {
    var v7 float32 = 3.14159
    v7 += 1.5
    switch "" {
    case "a":
        var v8 int = -207
        _ = v8
    case "foo bar":
        v7 += (1e10 * 0.1)
    case "\n\t":
        v9 := GenS1{F0: 2147483647}
        _ = v9
    }
    v7 = v7
    var v10 int64 = 840
    _ = v10
    _ = v7
    return 237
}

func main() {
    var v1 int8 = -128
    var v2 uint8 = 177
    var v3 int32 = -2147483648
    var v4 int32 = v3
    var v5 int = 1
    v6 := []uint32{436, 973, 2147483647, 42}
    var v11 uint32 = 434
    var v12 string = "\n\t"
    for fi13 := 0; fi13 < 12; fi13++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v12 = v12
        _ = fi13
    }
    v14 := GenS1{F0: 2147483647}
    var v15 uint64 = 77
    for fi16 := 0; fi16 < 19; fi16++ {
        v6 = append(v6, (v11 - genFn0()))
        fmt.Println(fi16, v4)
        if true {
            var v17 int = fi16
            var v18 uint = 255
            if (false || (255 == v18)) {
                var v19 int = v17
                _ = v19
            }
            _ = v17
            _ = v18
        }
        _ = fi16
    }
    var v20 float64 = -2.25
    switch 32767 {
    case 2147483647:
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 uint8 = v2
            switch "a" {
            case "":
                var v23 uint64 = 18446744073709551615
                _ = v23
            case "foo bar":
                fmt.Println(v4, v4, v2)
            case "a":
                fmt.Println(v5, v22, v11)
            default:
                v6 = append(v6, 1)
            }
            v24 := GenS2{F0: -9223372036854775808}
            _ = fi21
            _ = v22
            _ = v24
        }
    default:
        var v25 float64 = (v20 - ((v20 - (v20 * ((1.5 * v20) + (float64(float32(v20)) + -2.25)))) + v20))
        _ = v25
    }
    if len(v6) > 0 {
        v6[0] = genFn0()
    }
    var v26 uint16 = 65535
    var v27 int32 = v4
    for fi28 := 0; fi28 < 2; fi28++ {
        switch "a" {
        case "\n\t":
            switch (v12 + v12) {
            case "\n\t":
                var v29 int = (-(fi28))
                _ = v29
            case "":
                v26 += (v26 | v26)
            case "a":
                v11 = (v11 ^ (v11 | (v11 + (v11 & 42))))
            }
        case "hello":
            switch (("\n\t" + v12) + v12) {
            case "hello":
                v2 += (v2 % ((237) | 1))
            default:
                for fi30 := 0; fi30 < 7; fi30++ {
                    if false {
                        v31 := GenS3{F0: "\n\t", F1: 42, F2: v1}
                        if true {
                            var v32 uint64 = 869
                            fmt.Println(v26, v5)
                            fmt.Println(v1, v11, v12)
                            if true {
                                _ = len(v6)
                                var v33 int16 = 1
                                v12 = v12
                                var v34 int64 = -1
                                var v35 int64 = v34
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                            var v36 string = ("a" + (("\n\t" + v12) + ""))
                            _ = v32
                            _ = v36
                        } else {
                            if false {
                                v37 := GenS4{F0: (v26 + (v26 - (^(v26)))), F1: v11, F2: (v1 / ((int8(v11)) | 1))}
                                _ = v37
                            } else {
                                v38 := GenS5{F0: 42}
                                v11 = v11
                                switch v12 {
                                case "hello":
                                    if true {
                                        v39 := GenS3{F0: v12, F1: 438, F2: v1}
                                        switch int(v11) {
                                        case 4294967295:
                                            var v40 int = -409
                                            _ = v40
                                        default:
                                            v41 := GenS6{F0: "", F1: v5}
                                            _ = v41
                                        }
                                        fmt.Println(fi30, v15, v20)
                                        switch ((v12 + v12) + "a") {
                                        case "hello":
                                            if true {
                                                v6 = append(v6, (v11 - v11))
                                                if (((v27 / (((v27 % (((v4 >> ((18446744073709551615) & 63))) | 1))) | 1)) < v3) && (float64(1.5) <= 1.5)) {
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                }
                                                fmt.Println(v4, v3)
                                            } else {
                                                var v42 bool = (!((v2 ^ 42) > (v2 + v2)))
                                                v43 := GenS2{F0: int64(fi28)}
                                                v26 = (v26 >> ((851) & 63))
                                                _ = v42
                                                _ = v43
                                            }
                                        }
                                        _ = v39
                                    } else {
                                        var v44 uint64 = v15
                                        _ = v44
                                    }
                                case "a":
                                    fmt.Println(v27)
                                case "foo bar":
                                    var v45 uint32 = uint32(v2)
                                    _ = v45
                                }
                                _ = v38
                            }
                            switch v12 {
                            case "a":
                                fmt.Println(v11, v20)
                            case "foo bar":
                                v46 := GenS6{F0: v12, F1: (fi30 * 127)}
                                _ = v46
                            }
                        }
                        switch (v5 | fi28) {
                        case -9223372036854775808:
                            _ = len(v6)
                        case -433:
                            fmt.Println(fi28, v11)
                        case -365:
                            var v47 int16 = 0
                            _ = v47
                        default:
                            var v48 uint = 1
                            _ = v48
                        }
                        _ = v31
                    } else {
                        if false {
                            if (!false) {
                                v12 = (v12 + v12)
                            }
                            if len(v6) > 0 {
                                v6[0] = (v11 - v11)
                            }
                            var v49 float32 = ((0.0 * 0.1) / 1e10)
                            v50 := GenS3{F0: v12, F1: 1, F2: 100}
                            if true {
                                if (v26 <= (v26 >> ((1) & 63))) {
                                    fmt.Println(v12, v12)
                                    v51 := GenS7{F0: 4294967295, F1: v12, F2: "\n\t"}
                                    fmt.Println(fi30, v2)
                                    fmt.Println(v1, fi28)
                                    _ = v51
                                }
                                var v52 int = v5
                                v15 = v15
                                var v53 float32 = (0.0 * v49)
                                v54 := GenS8{F0: v26, F1: -1, F2: -846}
                                v55 := GenS3{F0: ("hello" + "a"), F1: 1, F2: v1}
                                v56 := GenS9{F0: 479}
                                _ = v52
                                _ = v53
                                _ = v54
                                _ = v55
                                _ = v56
                            } else {
                                if true {
                                    if (219 > 65535) {
                                        v5 = (fi28 << ((18446744073709551615) & 63))
                                        var v57 uint8 = v2
                                        _ = v57
                                    } else {
                                        v49 = 1e10
                                        v58 := GenS3{F0: "a", F1: 1, F2: (v1 - -128)}
                                        _ = v58
                                    }
                                    var v59 uint32 = uint32(v2)
                                    if len(v6) > 0 {
                                        v6[0] = 42
                                    }
                                    _ = v59
                                } else {
                                    v60 := GenS10{F0: false, F1: (v26 * v26)}
                                    v1 = (v1 * v1)
                                    _ = v60
                                }
                            }
                            _ = v49
                            _ = v50
                        } else {
                            var v61 bool = true
                            if (-751 > 700) {
                                var v62 string = "foo bar"
                                v3 -= 1
                                _ = v62
                            }
                            _ = v61
                        }
                    }
                    v2 = 128
                    fmt.Println(v27, v1)
                    _ = fi30
                }
            }
        case "foo bar":
            var v63 float32 = (((1e10 * ((1.5 * (1.5 + 1.5)) + -2.25)) + (0.0 * 0.0)) - 1e10)
            _ = v63
        default:
            v1 += v1
        }
        fmt.Println(v11, v12, v12)
        if true {
            v12 = "hello"
            var v64 float32 = 0.0
            _ = v64
        } else {
            for fi65 := 0; fi65 < 7; fi65++ {
                v6 = append(v6, 65535)
                var v66 int = v5
                _ = fi65
                _ = v66
            }
            var v67 float64 = v20
            switch 104 {
            case -2147483648:
                for fi68 := 0; fi68 < 9; fi68++ {
                    var v69 uint64 = (v15 - (v15 * (v15 >> ((267) & 63))))
                    var v70 uint64 = (v69 ^ (^(v15)))
                    _ = fi68
                    _ = v69
                    _ = v70
                }
            }
            _ = v67
        }
        _ = fi28
    }
    switch -2147483648 {
    case -2147483648:
        for fi71 := 0; fi71 < 20; fi71++ {
            fmt.Println(v12)
            _ = fi71
        }
    case 931:
        var v72 int32 = v3
        _ = v72
    case -305:
        v20 = (v20 * (v20 + ((v20 * (((float64(3.14159) * float64((0.1 - (1.5 + 1.5)))) - 0.0) + 1e10)) * 0.0)))
    default:
        var v73 uint32 = (v11 % ((v11) | 1))
        _ = v73
    }
    v74 := GenS9{F0: 65535}
    var v75 int32 = v27
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v74
    _ = v75
}
