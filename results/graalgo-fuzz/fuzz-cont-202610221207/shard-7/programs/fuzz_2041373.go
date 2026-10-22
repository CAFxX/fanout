package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint32 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 uint32; F2 bool }
type GenS5 struct { F0 int64; F1 int32; F2 uint32 }
type GenS6 struct { F0 int32; F1 string; F2 string }
type GenS7 struct { F0 uint64; F1 uint64; F2 uint }
func genFn0(p0 int32, p1 int16, p2 int32) int32 {
    v4 := []int{0, 780, 385}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    p2 += (p2 >> ((0) & 63))
    fmt.Println(p0, p0)
    fmt.Println(p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p0
}

func main() {
    var v1 uint64 = 358
    v2 := []string{}
    v3 := []int{-1}
    var v5 int8 = -102
    v6 := GenS1{F0: 42, F1: 0, F2: 2147483647}
    v7 := GenS2{F0: false, F1: -1}
    v1 = 4294967295
    fmt.Println(v5)
    v8 := GenS2{F0: false, F1: 386}
    var v9 uint64 = 4294967295
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 int = fi11
            var v13 string = "hello"
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    v14 := GenS3{F0: 1}
    switch 32768 {
    case 4294967295:
        fmt.Println(v5, v5, v1)
    case 32768:
        switch 128 {
        case 42:
            fmt.Println(v1, v5)
        case -478:
            var v15 string = "\n\t"
            _ = v15
        case -43:
            v5 -= (-(v5))
        }
    case -703:
        for fi16 := 0; fi16 < 14; fi16++ {
            v1 = v1
            _ = fi16
        }
    default:
        v1 += v9
    }
    if (1 >= 1000) {
        v9 += v1
        if (!(true || (true && false))) {
            v9 += v9
            if true {
                v5 += -128
            } else {
                v5 += 18
                v5 += v5
                var v17 bool = true
                _ = v17
            }
            switch 127 {
            case 42:
                v5 -= (v5 << ((742) & 63))
            default:
                var v18 uint64 = v1
                _ = v18
            }
        } else {
            switch ("" + ("hello" + ("foo bar" + ""))) {
            case "hello":
                for fi19 := 0; fi19 < 9; fi19++ {
                    if (42 == 2147483648) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v5 = (v5 - (v5 % ((v5) | 1)))
                    }
                    switch fi19 {
                    case -549:
                        if len(v3) > 0 {
                            v3[0] = (fi19 >> ((1) & 63))
                        }
                    case 0:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case -2147483648:
                        fmt.Println(fi19, v1, v5)
                    default:
                        if (65535 != 65535) {
                            switch (("" + "") + "hello") {
                            case "hello":
                                if false {
                                    var v20 uint16 = 255
                                    if (42 != (^(fi19))) {
                                        var v21 uint64 = v1
                                        switch ("\n\t" + "hello") {
                                        case "a":
                                            for fi22 := 0; fi22 < 9; fi22++ {
                                                var v23 int32 = 1000
                                                var v24 int64 = int64(v5)
                                                _ = fi22
                                                _ = v23
                                                _ = v24
                                            }
                                        case "hello":
                                            if ("" <= ("hello" + ("hello" + "a"))) {
                                                v20 = (v20 / (((^(v20))) | 1))
                                                switch 9223372036854775807 {
                                                case 127:
                                                    fmt.Println(v20, v1)
                                                case 4294967295:
                                                    switch "\n\t" {
                                                    case "a":
                                                        fmt.Println(v1)
                                                    }
                                                case 9223372036854775807:
                                                    v1 = 0
                                                default:
                                                    var v25 int16 = 929
                                                    _ = v25
                                                }
                                                var v26 string = ("\n\t" + "a")
                                                _ = v26
                                            }
                                        case "\n\t":
                                            var v27 int = (fi19 - (^(fi19)))
                                            _ = v27
                                        default:
                                            v28 := GenS2{F0: false, F1: -1}
                                            _ = v28
                                        }
                                        switch fi19 {
                                        case -502:
                                            var v29 float32 = 1e10
                                            _ = v29
                                        case 32767:
                                            for fi30 := 0; fi30 < 10; fi30++ {
                                                fmt.Println(v5, v20)
                                                v31 := GenS4{F0: 42, F1: 2, F2: (1 < 42)}
                                                v20 -= v20
                                                _ = fi30
                                                _ = v31
                                            }
                                        case -466:
                                            if false {
                                                var v32 float64 = 1.5
                                                if len(v3) > 0 {
                                                    v3[0] = 42
                                                }
                                                _ = v32
                                            } else {
                                                switch fi19 {
                                                case -1:
                                                    if (v20 <= (^(v20))) {
                                                        if ((false || true) && true) {
                                                            v33 := GenS1{F0: -775, F1: 32767, F2: 143}
                                                            _ = v33
                                                        }
                                                    } else {
                                                        v34 := GenS5{F0: 42, F1: genFn0(0, 42, 0), F2: 2}
                                                        for fi35 := 0; fi35 < 12; fi35++ {
                                                            var v36 float64 = (0.0 / (3.14159 / 0.1))
                                                            v36 = v36
                                                            _ = fi35
                                                            _ = v36
                                                        }
                                                        v21 -= v1
                                                        _ = v34
                                                    }
                                                }
                                                v20 = (^(v20))
                                            }
                                        }
                                        _ = v21
                                    }
                                    v37 := GenS5{F0: -957, F1: 0, F2: 0}
                                    _ = v20
                                    _ = v37
                                }
                            case "foo bar":
                                if true {
                                    for fi38 := 0; fi38 < 5; fi38++ {
                                        v3 = append(v3, fi19)
                                        var v39 float32 = 1.5
                                        var v40 uint64 = 42
                                        v41 := GenS6{F0: 242, F1: ("\n\t" + ((((("" + (("\n\t" + "foo bar") + "")) + ((((("hello" + "") + "a") + "") + ("" + "")) + "a")) + "\n\t") + ("\n\t" + "a")) + "foo bar")), F2: "hello"}
                                        _ = fi38
                                        _ = v39
                                        _ = v40
                                        _ = v41
                                    }
                                    v42 := GenS7{F0: (v9 << ((255) & 63)), F1: v9, F2: 4294967295}
                                    _ = v42
                                }
                            default:
                                v1 -= (v9 / (((^(v9))) | 1))
                            }
                        }
                    }
                    switch ("foo bar" + ("foo bar" + "a")) {
                    case "\n\t":
                        var v43 uint = 711
                        _ = v43
                    default:
                        for fi44 := 0; fi44 < 18; fi44++ {
                            v5 -= (v5 + -128)
                            _ = fi44
                        }
                    }
                    _ = fi19
                }
            case "foo bar":
                for fi45 := 0; fi45 < 3; fi45++ {
                    var v46 uint64 = v9
                    v47 := GenS6{F0: int32(v5), F1: "\n\t", F2: ""}
                    _ = fi45
                    _ = v46
                    _ = v47
                }
            case "a":
                _ = len(v3)
            default:
                v5 -= v5
            }
            var v48 float64 = float64(1e10)
            _ = v48
        }
        var v49 bool = false
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
