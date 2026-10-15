package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint; F2 int32 }
type GenS3 struct { F0 int32; F1 int16; F2 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 int16 }
func genFn0(p0 uint16, p1 int64) uint64 {
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 bool = false
        _ = fi5
        _ = v6
    }
    var v7 int32 = 42
    _ = p0
    _ = p1
    _ = v7
    return 0
}

func main() {
    var v1 uint16 = 42
    var v2 float64 = -2.25
    var v3 uint64 = 18446744073709551615
    var v4 int64 = 1
    v3 -= v3
    var v8 int8 = 0
    v9 := GenS1{F0: "foo bar"}
    v10 := []int{289, 390, 53}
    switch "hello" {
    case "a":
        v1 -= (v1 ^ v1)
    case "\n\t":
        v11 := GenS1{F0: "foo bar"}
        _ = v11
    case "foo bar":
        if true {
            if true {
                switch (("foo bar" + ("foo bar" + "foo bar")) + "hello") {
                case "":
                    fmt.Println(v8, v8, v4)
                }
                switch 42 {
                case 632:
                    v3 = v3
                case -446:
                    v12 := GenS1{F0: "a"}
                    _ = v12
                default:
                    for fi13 := 0; fi13 < 5; fi13++ {
                        fmt.Println(v1, v4)
                        switch -83 {
                        case 128:
                            v14 := GenS2{F0: -128, F1: 2, F2: -634}
                            _ = v14
                        case 485:
                            if true {
                                var v15 float32 = -2.25
                                _ = v15
                            }
                        }
                        if false {
                            v2 -= v2
                            for fi16 := 0; fi16 < 6; fi16++ {
                                fmt.Println(v8, v4)
                                _ = fi16
                            }
                            var v17 string = ((("\n\t" + "\n\t") + "foo bar") + "foo bar")
                            _ = v17
                        }
                        _ = fi13
                    }
                }
                if true {
                    v8 = (v8 + v8)
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    if false {
                        _ = len(v10)
                    } else {
                        fmt.Println(v4)
                    }
                } else {
                    var v18 int16 = 1000
                    _ = v18
                }
            }
            v1 += v1
            var v19 int16 = 559
            _ = v19
        }
    }
    fmt.Println(v1, v4, v3)
    var v20 uint16 = 65535
    for fi21 := 0; fi21 < 4; fi21++ {
        for fi22 := 0; fi22 < 5; fi22++ {
            v23 := GenS3{F0: -145, F1: -32768, F2: 255}
            v8 = v8
            var v24 float64 = 1.5
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi21
    }
    fmt.Println(v20)
    var v25 uint = 18446744073709551615
    var v26 uint8 = 2
    v27 := GenS4{F0: (false && false)}
    for fi28 := 0; fi28 < 2; fi28++ {
        if (!true) {
            v29 := GenS4{F0: true}
            for fi30 := 0; fi30 < 18; fi30++ {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                if true {
                    v31 := GenS3{F0: 2147483647, F1: 204, F2: (^(v26))}
                    v20 += 1
                    if (!false) {
                        var v32 uint16 = v20
                        _ = v32
                    } else {
                        if (!false) {
                            v33 := GenS3{F0: 1, F1: 1, F2: v26}
                            fmt.Println(fi28, fi28, fi28)
                            if len(v10) > 0 {
                                v10[0] = 1
                            }
                            _ = v33
                        } else {
                            if false {
                                var v34 int64 = (v4 >> ((2) & 63))
                                fmt.Println(v34)
                                v2 = (-2.25 * v2)
                                _ = v34
                            } else {
                                v26 += (^(v26))
                                var v35 uint = v25
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                                _ = v35
                            }
                            if false {
                                switch 9223372036854775807 {
                                case -709:
                                    var v36 int32 = 0
                                    _ = v36
                                default:
                                    v25 += v25
                                }
                                switch fi28 {
                                case 613:
                                    var v37 float32 = ((0.0 / (0.1 * 1e10)) / 0.1)
                                    _ = v37
                                case 0:
                                    v26 = v26
                                case 680:
                                    if false {
                                        fmt.Println(v8, v26)
                                    }
                                }
                                if false {
                                    _ = len(v10)
                                    v38 := GenS3{F0: 0, F1: 32767, F2: v26}
                                    v4 -= v4
                                    _ = v38
                                } else {
                                    var v39 uint = 531
                                    v20 = (v1 * v20)
                                    if (!(756 != -345)) {
                                        var v40 string = ("a" + "a")
                                        var v41 uint32 = 2147483648
                                        if false {
                                            if false {
                                                var v42 int8 = (v8 / ((int8(v26)) | 1))
                                                var v43 uint64 = 715
                                                switch "hello" {
                                                case "\n\t":
                                                    var v44 int64 = (v4 / (((v4 - 1)) | 1))
                                                    _ = v44
                                                }
                                                if false {
                                                    v45 := GenS3{F0: -1, F1: 0, F2: uint8(v20)}
                                                    _ = v45
                                                }
                                                _ = v42
                                                _ = v43
                                            }
                                            if len(v10) >= 2 {
                                                v10 = v10[1:len(v10)]
                                            }
                                        } else {
                                            var v46 bool = ((^(v8)) < (v8 & (v8 ^ v8)))
                                            _ = len(v10)
                                            _ = v46
                                        }
                                        _ = v40
                                        _ = v41
                                    }
                                    switch (fi28 ^ 874) {
                                    case 128:
                                        switch fi30 {
                                        case -2147483648:
                                            fmt.Println(v1)
                                        default:
                                            v47 := GenS4{F0: (!(!false))}
                                            _ = v47
                                        }
                                    case 9223372036854775807:
                                        switch ((("\n\t" + "foo bar") + ("\n\t" + "hello")) + "a") {
                                        case "hello":
                                            if ((^(v26)) > (v26 - 1)) {
                                                v20 += (v20 * v20)
                                                var v48 string = (("a" + "foo bar") + "hello")
                                                v3 = genFn0(65535, v4)
                                                v2 -= float64(-2.25)
                                                _ = v48
                                            }
                                        case "foo bar":
                                            v49 := GenS5{F0: 42, F1: 42}
                                            if (true && (!((v1 + v1) != v20))) {
                                                if len(v10) > 0 {
                                                    v10[0] = int(v4)
                                                }
                                                v50 := GenS2{F0: (v8 / (((v8 * v8)) | 1)), F1: v25, F2: -1}
                                                var v51 int64 = v4
                                                _ = v50
                                                _ = v51
                                            }
                                            _ = v49
                                        case "":
                                            if (!(0 != 1000)) {
                                                fmt.Println(fi28)
                                                v52 := GenS2{F0: (v8 << ((v25) & 63)), F1: 0, F2: 0}
                                                var v53 uint = (v25 * v39)
                                                _ = v52
                                                _ = v53
                                            }
                                        default:
                                            v3 -= (^(v3))
                                        }
                                    }
                                    _ = v39
                                }
                            } else {
                                var v54 int16 = 661
                                v1 += (v20 * v20)
                                _ = v54
                            }
                            var v55 int = (^(fi30))
                            _ = v55
                        }
                        if ((-(fi30)) > 764) {
                            var v56 uint32 = 88
                            _ = v56
                        }
                    }
                    _ = v31
                } else {
                    var v57 bool = true
                    if true {
                        fmt.Println(v25, v1)
                        v58 := GenS2{F0: v8, F1: v25, F2: -440}
                        _ = v58
                    }
                    if (!false) {
                        v2 = ((3.14159 + v2) / v2)
                        v2 += 1e10
                        if len(v10) > 0 {
                            v10[0] = 32767
                        }
                    } else {
                        v26 -= uint8(v3)
                    }
                    _ = v57
                }
                _ = fi30
            }
            switch "hello" {
            case "hello":
                fmt.Println(v20, v3, v1)
            case "foo bar":
                v2 = v2
            }
            _ = v29
        }
        var v59 float32 = -2.25
        _ = fi28
        _ = v59
    }
    if (1 >= 1000) {
        if true {
            v10 = append(v10, 128)
        }
    } else {
        var v60 float32 = 0.1
        _ = v60
    }
    for fi61 := 0; fi61 < 15; fi61++ {
        var v62 float64 = v2
        _ = fi61
        _ = v62
    }
    fmt.Println(v20, v1)
    v63 := GenS5{F0: (^(v4)), F1: -32768}
    var v64 uint8 = (^(v26))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v64:", v64)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v63
    _ = v64
    _ = v8
    _ = v9
}
