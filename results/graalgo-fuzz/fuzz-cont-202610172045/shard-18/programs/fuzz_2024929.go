package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 int16; F2 uint }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 uint16 }

func main() {
    var v1 int32 = 162
    v2 := []int{-149}
    v3 := make(map[int]int)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v4 := GenS1{F0: (0.1 < 1.5), F1: "hello"}
    _ = len(v2)
    var v5 uint64 = 4294967295
    switch 4294967295 {
    case 1:
        v3[0] = 127
    }
    v5 -= 18446744073709551615
    v5 = (v5 * (^(v5)))
    v5 += (v5 & 260)
    if false {
        var v6 float64 = (1e10 / 0.1)
        fmt.Println(v5, v1, v1)
        _ = len(v2)
        _ = v6
    } else {
        var v7 uint16 = 32767
        _ = v7
    }
    v1 = v1
    switch -2147483648 {
    case 95:
        if (!false) {
            v8 := v3[194]
            _ = v8
        } else {
            var v9 string = "a"
            _ = v9
        }
    case 0:
        var v10 float64 = (1e10 / 0.1)
        _ = v10
    case 4294967295:
        var v11 int32 = 1
        _ = v11
    }
    if true {
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if true {
                switch 669 {
                case 1:
                    for fi12 := 0; fi12 < 8; fi12++ {
                        fmt.Println(v1)
                        v1 -= 2147483647
                        v13 := GenS2{F0: 1}
                        _ = fi12
                        _ = v13
                    }
                case 0:
                    fmt.Println(v5)
                case 42:
                    if (435 < 288) {
                        var v14 int16 = -860
                        v14 = v14
                        switch ("hello" + "foo bar") {
                        case "a":
                            var v15 float64 = 1e10
                            _ = v15
                        }
                        _ = v14
                    }
                default:
                    v5 = (v5 & v5)
                }
                var v16 uint16 = 65535
                _ = v16
            } else {
                if false {
                    for fi17 := 0; fi17 < 3; fi17++ {
                        if true {
                            v5 -= v5
                            v5 += v5
                        } else {
                            v18 := GenS3{F0: 65535, F1: 1, F2: 2}
                            for fi19 := 0; fi19 < 1; fi19++ {
                                v5 = (v5 % ((v5) | 1))
                                var v20 bool = true
                                switch "a" {
                                case "foo bar":
                                    if true {
                                        delete(v3, -301)
                                    } else {
                                        switch "" {
                                        case "a":
                                            switch (fi17 + int(v1)) {
                                            case 139:
                                                fmt.Println(fi19)
                                            case 32767:
                                                switch "" {
                                                case "":
                                                    var v21 uint16 = 186
                                                    _ = v21
                                                case "foo bar":
                                                    var v22 int16 = 32767
                                                    _ = v22
                                                case "\n\t":
                                                    v5 = v5
                                                }
                                                var v23 int8 = 0
                                                v23 = (v23 >> ((42) & 63))
                                                _ = v23
                                            default:
                                                var v24 string = ("foo bar" + "foo bar")
                                                _ = v24
                                            }
                                        case "":
                                            var v25 bool = true
                                            _ = len(v2)
                                            _ = v25
                                        default:
                                            var v26 float64 = 0.0
                                            _ = v26
                                        }
                                        v27 := GenS4{F0: 1, F1: 764}
                                        _ = v27
                                    }
                                case "a":
                                    v20 = (!v20)
                                default:
                                    var v28 uint64 = v5
                                    _ = v28
                                }
                                _ = fi19
                                _ = v20
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v18
                        }
                        var v29 bool = ((true && (false && ((!true) || true))) && (18446744073709551615 >= v5))
                        _ = fi17
                        _ = v29
                    }
                    switch "hello" {
                    case "hello":
                        v30 := GenS3{F0: 42, F1: -444, F2: 4294967295}
                        _ = v30
                    case "foo bar":
                        var v31 uint16 = 0
                        _ = v31
                    case "a":
                        v5 += v5
                    default:
                        for fi32 := 0; fi32 < 14; fi32++ {
                            if (2147483648 >= 0) {
                                fmt.Println(v1, v5, fi32)
                            } else {
                                var v33 uint8 = 1
                                _ = v33
                            }
                            v34 := GenS3{F0: 42, F1: -32768, F2: 614}
                            _ = fi32
                            _ = v34
                        }
                    }
                    var v35 float64 = (1e10 + ((0.1 - 0.0) / 3.14159))
                    _ = v35
                } else {
                    var v36 float32 = -2.25
                    v5 = v5
                    if (((false && (v1 == v1)) && false) || true) {
                        v36 = (float32(-2.25) - v36)
                        if true {
                            fmt.Println(v36, v5, v5)
                            if true {
                                var v37 uint32 = 900
                                _ = v37
                            }
                            for fi38 := 0; fi38 < 12; fi38++ {
                                for fi39 := 0; fi39 < 7; fi39++ {
                                    var v40 int32 = v1
                                    fmt.Println(v1)
                                    _ = fi39
                                    _ = v40
                                }
                                switch 127 {
                                case 1:
                                    v1 -= v1
                                }
                                _ = fi38
                            }
                        } else {
                            v41 := GenS4{F0: 127, F1: 684}
                            for fi42 := 0; fi42 < 10; fi42++ {
                                if true {
                                    v1 = -493
                                    switch (fi42 + fi42) {
                                    case 42:
                                        if ((127 == 1) && false) {
                                            var v43 int32 = -1
                                            v44 := v3[(fi42 / (((^(fi42))) | 1))]
                                            fmt.Println(v36)
                                            _ = v43
                                            _ = v44
                                        }
                                    case 4294967295:
                                        v1 = (v1 - int32(v5))
                                    default:
                                        for fi45 := 0; fi45 < 4; fi45++ {
                                            v5 = (^(v5))
                                            var v46 int8 = 100
                                            _ = fi45
                                            _ = v46
                                        }
                                    }
                                    fmt.Println(fi42, v1, v1)
                                } else {
                                    var v47 uint32 = 65535
                                    _ = v47
                                }
                                _ = fi42
                            }
                            _ = v41
                        }
                    }
                    _ = v36
                }
                fmt.Println(v5, v5)
            }
            fmt.Println(v1, v5)
        }
        v48 := GenS5{F0: 255}
        if ((4294967295 >= 0) && (false && (!((v5 * 419) != v5)))) {
            fmt.Println(v1)
        }
        _ = v48
    } else {
        var v49 int8 = 1
        _ = v49
    }
    if (188 == 2) {
        v5 = v5
        v50 := v3[-2147483648]
        _ = v50
    }
    v5 = (v5 >> ((267) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
