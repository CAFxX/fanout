package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int8; F1 int16; F2 int16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 string = ((("hello" + "\n\t") + ("\n\t" + "")) + "")
    v3 := make(map[string]uint32)
    var v4 int32 = -1
    fmt.Println(v4)
    v2 = v2
    v2 = v2
    var v5 int16 = -1
    delete(v1, v2)
    _ = len(v1)
    if true {
        v6 := GenS1{F0: 768, F1: v2, F2: 255}
        var v7 uint16 = 0
        var v8 bool = (525 < 42)
        _ = v6
        _ = v7
        _ = v8
    }
    if false {
        for fi9 := 0; fi9 < 16; fi9++ {
            v10 := v3[v2]
            if true {
                if false {
                    v11 := GenS2{F0: (v5 << ((375) & 63))}
                    v12 := GenS1{F0: 4294967295, F1: v2, F2: 1}
                    _ = v11
                    _ = v12
                }
                v4 = (^(v4))
            } else {
                v13 := GenS2{F0: 764}
                v2 = v2
                var v14 uint16 = 42
                _ = v13
                _ = v14
            }
            switch v2 {
            case "foo bar":
                for fi15 := 0; fi15 < 11; fi15++ {
                    var v16 float64 = (0.1 * 0.0)
                    v10 -= 65535
                    var v17 int32 = (v4 / ((v4) | 1))
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            case "hello":
                if (-2147483648 > 1) {
                    var v18 uint8 = 255
                    var v19 bool = false
                    _ = v18
                    _ = v19
                } else {
                    v20 := GenS3{F0: v2}
                    v5 -= v5
                    _ = v20
                }
            default:
                _ = len(v3)
            }
            _ = fi9
            _ = v10
        }
    }
    var v21 float64 = 1e10
    v5 += (v5 % (((v5 * (v5 % ((-270) | 1)))) | 1))
    switch "hello" {
    case "hello":
        var v22 int32 = v4
        _ = v22
    case "foo bar":
        v4 = -739
    case "a":
        v23 := GenS4{F0: 2147483648}
        _ = v23
    default:
        var v24 bool = (v5 == (v5 ^ v5))
        _ = v24
    }
    if ((1 != int64(v4)) && ((0 >= 4294967295) && true)) {
        if false {
            v25 := GenS3{F0: "hello"}
            _ = v25
        }
    }
    fmt.Println(v5)
    switch "foo bar" {
    case "\n\t":
        var v26 uint16 = 530
        _ = v26
    case "hello":
        v27 := GenS4{F0: 280}
        _ = v27
    case "foo bar":
        v2 = "foo bar"
    }
    v2 = "a"
    for fi28 := 0; fi28 < 10; fi28++ {
        switch (fi28 >> ((388) & 63)) {
        case 2147483647:
            var v29 uint8 = 127
            _ = v29
        }
        if false {
            if false {
                fmt.Println(v5, v5, v5)
                for fi30 := 0; fi30 < 1; fi30++ {
                    var v31 int = 42
                    v21 = ((float64(1e10) / 1.0) * float64(((-2.25 + 0.1) - (((1e10 - (0.1 / 0.1)) - 1.5) / 1.5))))
                    v21 += ((float64(float32((v21 + (float64(((0.0 + 3.14159) - 1.5)) - (float64(3.14159) + (v21 / 1.5)))))) + float64(3.14159)) * ((v21 + ((v21 * (((float64((0.0 / float32((1.5 - (v21 * 0.0))))) / v21) * (float64(((1.5 / 1.0) / float32(3.14159))) * v21)) * 0.0)) + ((3.14159 / (float64(1.5) + (float64(0.1) + v21))) + 0.0))) - (0.1 / 1.5)))
                    _ = fi30
                    _ = v31
                }
            }
        } else {
            if true {
                v32 := []int{(fi28 - (fi28 | (fi28 << ((0) & 63)))), fi28}
                if false {
                    if false {
                        v21 -= 1.5
                        if (!true) {
                            v5 = v5
                        } else {
                            var v33 string = ("hello" + v2)
                            v34 := GenS5{F0: 4294967295, F1: (v33 + "")}
                            _ = v33
                            _ = v34
                        }
                        switch v2 {
                        case "a":
                            for fi35 := 0; fi35 < 8; fi35++ {
                                if (true || (!(0 == 4294967295))) {
                                    var v36 uint8 = 184
                                    _ = v36
                                }
                                _ = fi35
                            }
                        case "foo bar":
                            var v37 uint32 = 0
                            _ = v37
                        case "hello":
                            switch v2 {
                            case "foo bar":
                                var v38 int16 = v5
                                _ = v38
                            case "hello":
                                var v39 uint64 = 2
                                _ = v39
                            case "\n\t":
                                var v40 float32 = -2.25
                                _ = v40
                            default:
                                var v41 int16 = v5
                                _ = v41
                            }
                        default:
                            if true {
                                var v42 int32 = v4
                                v42 = (-(v4))
                                if false {
                                    if false {
                                        var v43 int32 = v42
                                        if true {
                                            var v44 float64 = v21
                                            _ = v44
                                        }
                                        _ = v43
                                    }
                                    v2 = ((v2 + v2) + v2)
                                    v42 -= int32(fi28)
                                }
                                _ = v42
                            }
                        }
                    }
                    if (true || true) {
                        v4 -= (v4 << ((18446744073709551615) & 63))
                        if (4294967295 < 4294967295) {
                            for fi45 := 0; fi45 < 12; fi45++ {
                                if ((!(!true)) && true) {
                                    fmt.Println(v2, fi45, fi28)
                                }
                                v32 = append(v32, int(v5))
                                var v46 int8 = 127
                                _ = fi45
                                _ = v46
                            }
                            v5 -= v5
                            var v47 bool = true
                            _ = v47
                        }
                        var v48 string = "hello"
                        _ = v48
                    } else {
                        v49 := GenS6{F0: (v4 | (v4 - -310))}
                        fmt.Println(v4)
                        fmt.Println(v4, fi28)
                        _ = v49
                    }
                    var v50 int64 = -9223372036854775808
                    _ = v50
                } else {
                    var v51 uint32 = 491
                    if (((fi28 ^ fi28) == (^(fi28))) && true) {
                        for fi52 := 0; fi52 < 12; fi52++ {
                            if (true || false) {
                                fmt.Println(v21)
                                v53 := GenS1{F0: 0, F1: ((v2 + v2) + v2), F2: 42}
                                var v54 uint64 = 2
                                _ = len(v1)
                                var v55 uint16 = 65535
                                v4 -= 0
                                _ = v53
                                _ = v54
                                _ = v55
                            } else {
                                var v56 int8 = -128
                                _ = v56
                            }
                            v5 -= (-(v5))
                            var v57 int8 = 42
                            var v58 int64 = 9223372036854775807
                            v21 = v21
                            _ = fi52
                            _ = v57
                            _ = v58
                        }
                        var v59 uint = 255
                        _ = v59
                    } else {
                        var v60 uint = 0
                        for fi61 := 0; fi61 < 8; fi61++ {
                            if len(v32) > 0 {
                                v32[0] = fi61
                            }
                            switch v2 {
                            case "":
                                if false {
                                    if true {
                                        v62 := GenS5{F0: 0, F1: ((v2 + v2) + "hello")}
                                        var v63 uint16 = 2
                                        v21 = 0.0
                                        _ = v62
                                        _ = v63
                                    }
                                } else {
                                    v2 = "a"
                                    var v64 uint64 = 72
                                    v51 = (v51 ^ (v51 % ((v51) | 1)))
                                    _ = v64
                                }
                            case "hello":
                                var v65 uint8 = 121
                                _ = v65
                            case "a":
                                v66 := GenS4{F0: 407}
                                _ = v66
                            }
                            _ = fi61
                        }
                        v67 := GenS6{F0: -2147483648}
                        _ = v60
                        _ = v67
                    }
                    _ = v51
                }
                _ = v32
            }
        }
        _ = fi28
    }
    var v68 int32 = (v4 >> ((2) & 63))
    switch (v2 + v2) {
    case "":
        if ("hello" > v2) {
            switch "a" {
            case "a":
                v69 := GenS5{F0: 2, F1: ""}
                _ = v69
            case "":
                var v70 uint64 = 4294967295
                _ = v70
            case "hello":
                var v71 bool = (0 > 255)
                _ = v71
            }
            v72 := GenS7{F0: 1, F1: v5, F2: 378}
            var v73 int8 = 42
            _ = v72
            _ = v73
        }
    case "a":
        var v74 float32 = -2.25
        _ = v74
    case "\n\t":
        var v75 int16 = -774
        _ = v75
    default:
        v2 = (v2 + "foo bar")
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v68
}
