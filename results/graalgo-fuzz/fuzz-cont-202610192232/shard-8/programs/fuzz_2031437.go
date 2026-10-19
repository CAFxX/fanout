package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint16; F1 string; F2 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 uint8; F2 int8 }
type GenS5 struct { F0 uint8; F1 uint32; F2 bool }
type GenS6 struct { F0 uint64; F1 uint8 }
type GenS7 struct { F0 string; F1 string; F2 int32 }
type GenS8 struct { F0 int8; F1 int32; F2 int64 }

func main() {
    v1 := []string{""}
    v2 := make(map[uint32]string)
    var v3 int = 42
    var v4 float32 = (0.0 * (0.1 - (1.5 / 1.0)))
    v5 := make(map[uint32]string)
    fmt.Println(v4)
    v6 := GenS1{F0: 32768, F1: false}
    if (false || false) {
        if (true || true) {
            _ = len(v1)
        }
        switch (v3 - 700) {
        case 128:
            var v7 string = "hello"
            _ = v7
        case 937:
            fmt.Println(v4, v3)
        case 2147483647:
            if (92 > 128) {
                v8 := GenS1{F0: v3, F1: true}
                v3 = v3
                if false {
                    v9 := GenS1{F0: 32767, F1: false}
                    _ = v9
                }
                _ = v8
            }
        }
        v4 = (((((v4 * -2.25) * (0.0 + (0.0 / 1.0))) - v4) / (v4 + 0.1)) + 0.1)
    } else {
        var v10 float64 = -2.25
        if false {
            v3 -= (v3 - -1)
            v11 := GenS1{F0: (v3 / ((15) | 1)), F1: (false && (-612 <= (^(v3))))}
            v4 = (v4 + -2.25)
            _ = v11
        } else {
            switch 42 {
            case 32767:
                for fi12 := 0; fi12 < 18; fi12++ {
                    var v13 uint64 = 1
                    if (127 <= 75) {
                        var v14 int32 = -1
                        switch "\n\t" {
                        case "\n\t":
                            v15 := GenS1{F0: (-(v3)), F1: true}
                            _ = v15
                        case "a":
                            v16 := GenS2{F0: 255, F1: "hello", F2: (v13 / (((v13 / ((421) | 1))) | 1))}
                            _ = v16
                        }
                        _ = v14
                    }
                    _ = fi12
                    _ = v13
                }
            default:
                switch "hello" {
                case "hello":
                    switch (v3 & 2147483647) {
                    case 1:
                        for fi17 := 0; fi17 < 13; fi17++ {
                            v10 = (0.0 * float64(3.14159))
                            for fi18 := 0; fi18 < 8; fi18++ {
                                var v19 uint32 = 1
                                v10 = (float64((v4 - v4)) - v10)
                                var v20 int8 = 127
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                            _ = fi17
                        }
                    case 2147483647:
                        for fi21 := 0; fi21 < 5; fi21++ {
                            fmt.Println(fi21)
                            _ = fi21
                        }
                    }
                case "foo bar":
                    for fi22 := 0; fi22 < 4; fi22++ {
                        if true {
                            v23 := GenS3{F0: 1}
                            if true {
                                switch 302 {
                                case 0:
                                    v24 := GenS4{F0: 1, F1: 1, F2: 87}
                                    _ = v24
                                case 1:
                                    if false {
                                        fmt.Println(v10, v10, v10)
                                    }
                                }
                                if (-816 != -389) {
                                    if false {
                                        for fi25 := 0; fi25 < 9; fi25++ {
                                            v10 = (v10 - v10)
                                            var v26 uint32 = 482
                                            _ = fi25
                                            _ = v26
                                        }
                                        v27 := GenS1{F0: fi22, F1: true}
                                        v3 = 1
                                        _ = v27
                                    }
                                } else {
                                    if false {
                                        fmt.Println(v4, v3)
                                    }
                                }
                                if (0 < 1) {
                                    fmt.Println(v4)
                                    v4 += 0.1
                                } else {
                                    if true {
                                        var v28 uint = 983
                                        _ = v28
                                    } else {
                                        var v29 uint16 = 32767
                                        _ = v29
                                    }
                                }
                            } else {
                                fmt.Println(v4, fi22)
                            }
                            _ = v23
                        } else {
                            v30 := GenS5{F0: 67, F1: 0, F2: true}
                            var v31 int32 = 42
                            _ = v30
                            _ = v31
                        }
                        var v32 int64 = int64(v3)
                        var v33 int = 2147483647
                        _ = fi22
                        _ = v32
                        _ = v33
                    }
                case "a":
                    v34 := GenS6{F0: 788, F1: 96}
                    _ = v34
                }
            }
            if (!(true && (true || true))) {
                var v35 uint64 = 1
                _ = v35
            }
            if false {
                switch (v3 | (v3 - v3)) {
                case 32768:
                    delete(v2, 4294967295)
                case -1:
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(v3, v4, v3)
                    }
                default:
                    v4 = v4
                }
                var v36 uint = 255
                _ = v36
            } else {
                var v37 uint = 0
                if true {
                    var v38 uint = v37
                    _ = v38
                }
                for fi39 := 0; fi39 < 1; fi39++ {
                    delete(v2, uint32(v37))
                    _ = fi39
                }
                _ = v37
            }
        }
        _ = v10
    }
    switch "" {
    case "":
        v3 -= 42
    default:
        if false {
            switch "\n\t" {
            case "\n\t":
                v40 := GenS6{F0: 253, F1: 255}
                _ = v40
            case "":
                var v41 float32 = v4
                _ = v41
            case "hello":
                var v42 float64 = ((1.5 - 0.0) / 3.14159)
                _ = v42
            default:
                v43 := GenS3{F0: 100}
                _ = v43
            }
            v1 = append(v1, "hello")
        } else {
            switch (v3 * v3) {
            case 954:
                v44 := GenS7{F0: "foo bar", F1: "\n\t", F2: 42}
                _ = v44
            case -2147483648:
                if false {
                    var v45 uint64 = 2
                    var v46 uint64 = (v45 ^ (v45 - v45))
                    _ = v45
                    _ = v46
                }
            case 32767:
                _ = len(v1)
            default:
                var v47 int32 = -1
                _ = v47
            }
            v3 = v3
        }
    }
    if (false || (1 > -63)) {
        fmt.Println(v3)
    } else {
        v48 := GenS5{F0: 188, F1: 889, F2: false}
        _ = v48
    }
    var v49 int16 = -32768
    v49 = v49
    for fi50 := 0; fi50 < 2; fi50++ {
        v4 = 1e10
        var v51 int = v3
        if true {
            v49 += -801
        } else {
            fmt.Println(fi50)
            var v52 bool = false
            switch ("\n\t" + "foo bar") {
            case "hello":
                for fi53 := 0; fi53 < 9; fi53++ {
                    v2[42] = "hello"
                    if (!v52) {
                        if v52 {
                            fmt.Println(fi50, v4, v4)
                        }
                        if v52 {
                            v4 = v4
                            v52 = (65535 > 65535)
                            var v54 bool = v52
                            _ = v54
                        } else {
                            fmt.Println(fi50, fi53, fi53)
                            var v55 int32 = 0
                            _ = v55
                        }
                    }
                    _ = fi53
                }
            default:
                fmt.Println(fi50, v51, fi50)
            }
            _ = v52
        }
        _ = fi50
        _ = v51
    }
    v3 += 2147483647
    v56 := GenS8{F0: 39, F1: 264, F2: 1}
    for fi57 := 0; fi57 < 14; fi57++ {
        var v58 int64 = 9223372036854775807
        v59 := GenS4{F0: 102, F1: 1, F2: -128}
        _ = fi57
        _ = v58
        _ = v59
    }
    v49 -= v49
    fmt.Println(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v56
    _ = v6
}
