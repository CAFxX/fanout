package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int32 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64; F1 int64 }
func genFn0() int {
    v5 := GenS1{F0: (1 >= 127), F1: 2147483648, F2: -2147483648}
    var v6 uint16 = 32767
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint32 = 906
        if true {
            v9 := GenS2{F0: -562, F1: 602}
            _ = v9
        }
        v6 -= v6
        _ = fi7
        _ = v8
    }
    var v10 int32 = 2147483647
    fmt.Println(v10)
    _ = v10
    _ = v5
    _ = v6
    return -2147483648
}
func genFn1() bool {
    switch "a" {
    case "\n\t":
        if false {
            if false {
                switch (("foo bar" + "\n\t") + ("foo bar" + "foo bar")) {
                case "hello":
                    var v11 string = "a"
                    _ = v11
                case "":
                    if true {
                        var v12 int16 = 417
                        _ = v12
                    } else {
                        switch -1 {
                        case -1:
                            var v13 uint32 = 4294967295
                            _ = v13
                        case 32768:
                            if false {
                                var v14 int64 = -9223372036854775808
                                var v15 uint = 255
                                v16 := []int{127, -2147483648, 885, 42}
                                _ = v14
                                _ = v15
                                _ = v16
                            }
                        case -473:
                            var v17 string = "hello"
                            _ = v17
                        default:
                            for fi18 := 0; fi18 < 16; fi18++ {
                                for fi19 := 0; fi19 < 4; fi19++ {
                                    if false {
                                        switch -2147483648 {
                                        case 32768:
                                            var v20 float64 = (((1.5 / 0.1) - 1.5) + (1e10 + (((-2.25 + ((float64(0.0) / 0.1) / 1e10)) / 1.0) * 3.14159)))
                                            _ = v20
                                        }
                                    } else {
                                        var v21 float64 = 1.5
                                        fmt.Println(fi19, fi19)
                                        v21 = v21
                                        _ = v21
                                    }
                                    if ((!true) && (476 == 0)) {
                                        var v22 float32 = (0.1 / 1.0)
                                        _ = v22
                                    } else {
                                        var v23 int64 = 1
                                        var v24 uint64 = 4294967295
                                        v23 -= v23
                                        v25 := []int{fi19, (^(fi19)), fi18}
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                    }
                                    v26 := GenS3{F0: 607}
                                    _ = fi19
                                    _ = v26
                                }
                                var v27 int32 = 627
                                _ = fi18
                                _ = v27
                            }
                        }
                        v28 := []int{9223372036854775807}
                        switch "foo bar" {
                        case "":
                            var v29 uint16 = 255
                            _ = v29
                        case "foo bar":
                            if false {
                                v28 = append(v28, 32768)
                            }
                        }
                        _ = v28
                    }
                case "a":
                    v30 := []int{}
                    _ = v30
                }
                var v31 uint64 = 291
                switch 9223372036854775807 {
                case 802:
                    if false {
                        v31 = v31
                        v31 += (^(v31))
                    } else {
                        v32 := GenS4{F0: 42}
                        switch 1 {
                        case 4294967295:
                            if false {
                                if true {
                                    var v33 int16 = 337
                                    _ = v33
                                }
                                for fi34 := 0; fi34 < 16; fi34++ {
                                    fmt.Println(v31, v31)
                                    v31 = (v31 ^ (^(v31)))
                                    _ = fi34
                                }
                                var v35 float64 = -2.25
                                _ = v35
                            }
                        case -2147483648:
                            var v36 bool = true
                            _ = v36
                        }
                        switch 979 {
                        case 0:
                            fmt.Println(v31)
                        case 430:
                            switch 985 {
                            case -749:
                                fmt.Println(v31)
                            case 128:
                                switch ("a" + "a") {
                                case "hello":
                                    v31 = (v31 - v31)
                                case "":
                                    v31 = v31
                                }
                            }
                        case 531:
                            v31 = 2
                        }
                        _ = v32
                    }
                case -1:
                    v37 := GenS5{F0: 630}
                    _ = v37
                case 418:
                    var v38 int = genFn0()
                    _ = v38
                }
                _ = v31
            }
        }
    case "hello":
        for fi39 := 0; fi39 < 5; fi39++ {
            v40 := GenS1{F0: (!(!false)), F1: 761, F2: -1}
            _ = fi39
            _ = v40
        }
    default:
        for fi41 := 0; fi41 < 15; fi41++ {
            var v42 int32 = 0
            if (515 >= 42) {
                if true {
                    v43 := []int{fi41, fi41}
                    v42 = v42
                    var v44 int64 = 0
                    _ = v43
                    _ = v44
                }
                fmt.Println(fi41, v42)
                v45 := GenS5{F0: 1}
                _ = v45
            }
            _ = fi41
            _ = v42
        }
    }
    if (65535 != 42) {
        var v46 bool = (-196 < -1)
        v46 = v46
        fmt.Println(v46)
        _ = v46
    } else {
        if (false || true) {
            var v47 uint8 = 0
            _ = v47
        }
    }
    if false {
        var v48 uint64 = 18446744073709551615
        var v49 uint64 = v48
        _ = v48
        _ = v49
    } else {
        if false {
            var v50 uint64 = 18446744073709551615
            for fi51 := 0; fi51 < 16; fi51++ {
                v50 += 1
                for fi52 := 0; fi52 < 13; fi52++ {
                    v50 = (v50 ^ 4294967295)
                    v50 = (v50 ^ v50)
                    v53 := GenS6{F0: 124, F1: 0}
                    _ = fi52
                    _ = v53
                }
                _ = fi51
            }
            for fi54 := 0; fi54 < 1; fi54++ {
                if ((!true) || ("a" != ("foo bar" + "foo bar"))) {
                    v50 -= (^(v50))
                    v50 = v50
                } else {
                    if (0 < 26) {
                        var v55 uint32 = 42
                        if (!true) {
                            var v56 int64 = -511
                            _ = v56
                        } else {
                            fmt.Println(v50, v50)
                        }
                        var v57 bool = (!(-422 <= 42))
                        _ = v55
                        _ = v57
                    }
                    v50 += v50
                }
                var v58 uint16 = 255
                for fi59 := 0; fi59 < 2; fi59++ {
                    v60 := []int{}
                    var v61 float64 = 1e10
                    v62 := GenS3{F0: 2}
                    _ = fi59
                    _ = v60
                    _ = v61
                    _ = v62
                }
                _ = fi54
                _ = v58
            }
            _ = v50
        }
    }
    return (0 != 918)
}

func main() {
    var v1 float32 = float32(1.5)
    var v2 uint16 = 52
    v3 := make(map[string]bool)
    var v4 uint16 = (v2 / ((v2) | 1))
    fmt.Println(v1)
    v1 += v1
    _ = len(v3)
    v1 += float32(float64(v1))
    var v63 int16 = 1000
    var v64 uint64 = 205
    _ = len(v3)
    v63 -= (v63 >> ((1) & 63))
    v65 := GenS2{F0: 2147483647, F1: -2147483648}
    v63 = v63
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v63
    _ = v64
    _ = v65
}
