package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint64; F2 string }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 bool }
func genFn0(p0 int64) uint64 {
    v7 := GenS1{F0: 147}
    var v8 int64 = p0
    var v9 int = 2147483647
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 uint64 = 2
    var v2 uint16 = 255
    var v3 uint64 = (^(v1))
    var v4 string = (("" + "hello") + "")
    v5 := make(map[string]bool)
    v6 := []uint32{0, 2147483647}
    var v10 string = v4
    switch -9223372036854775808 {
    case 32767:
        v11 := GenS2{F0: 0}
        _ = v11
    }
    var v12 int8 = 124
    for fi13 := 0; fi13 < 11; fi13++ {
        switch fi13 {
        case 0:
            var v14 uint32 = 0
            _ = v14
        default:
            if (!true) {
                var v15 float32 = 1.5
                v12 = (v12 + v12)
                if false {
                    v16 := GenS3{F0: v12, F1: (v3 >> ((42) & 63)), F2: v4}
                    _ = v16
                } else {
                    if false {
                        if (!(true || true)) {
                            switch "\n\t" {
                            case "\n\t":
                                _ = len(v6)
                            case "foo bar":
                                v17 := GenS4{F0: 4294967295, F1: (v3 > v3)}
                                _ = v17
                            case "hello":
                                switch fi13 {
                                case 812:
                                    var v18 int8 = int8(v2)
                                    _ = v18
                                case 261:
                                    v5[(v4 + v4)] = false
                                case 1:
                                    v5[v4] = true
                                }
                            default:
                                _ = len(v5)
                            }
                            fmt.Println(v2)
                        } else {
                            v10 = v4
                            for fi19 := 0; fi19 < 1; fi19++ {
                                fmt.Println(v2)
                                var v20 float64 = 3.14159
                                _ = fi19
                                _ = v20
                            }
                        }
                        v21 := GenS4{F0: 4294967295, F1: (2 == 2147483648)}
                        var v22 uint16 = 698
                        _ = v21
                        _ = v22
                    } else {
                        v12 = v12
                        switch fi13 {
                        case 32768:
                            var v23 string = "a"
                            _ = v23
                        case 128:
                            v1 = v1
                        case 916:
                            var v24 int64 = 9223372036854775807
                            _ = v24
                        }
                    }
                }
                _ = v15
            }
        }
        _ = fi13
    }
    var v25 int16 = 341
    switch -530 {
    case -327:
        for fi26 := 0; fi26 < 14; fi26++ {
            v27 := GenS5{F0: (v1 | genFn0(54)), F1: (v12 * (-(v12)))}
            _ = fi26
            _ = v27
        }
    default:
        v28 := GenS4{F0: 153, F1: ((float64((((((1.5 + 1e10) + (1.5 + 1e10)) / 0.1) + (-2.25 * float32(0.0))) - (3.14159 - 3.14159))) / (-2.25 + 3.14159)) != -2.25)}
        _ = v28
    }
    if (false || true) {
        for fi29 := 0; fi29 < 1; fi29++ {
            for fi30 := 0; fi30 < 7; fi30++ {
                var v31 int32 = 1
                switch v10 {
                case "foo bar":
                    v32 := GenS6{F0: (^(v1))}
                    v33 := GenS2{F0: (^(fi30))}
                    if true {
                        fmt.Println(v10)
                    } else {
                        v34 := GenS3{F0: (-(v12)), F1: 1, F2: "foo bar"}
                        var v35 float32 = ((1e10 / 1.0) + 3.14159)
                        v35 = v35
                        _ = v34
                        _ = v35
                    }
                    _ = v32
                    _ = v33
                case "hello":
                    if false {
                        v36 := GenS7{F0: (-(v12))}
                        var v37 float64 = 3.14159
                        var v38 string = v10
                        _ = v36
                        _ = v37
                        _ = v38
                    } else {
                        v39 := GenS6{F0: (v1 + (v1 << ((65535) & 63)))}
                        v2 = 44
                        v40 := GenS8{F0: (v31 << ((2) & 63))}
                        _ = v39
                        _ = v40
                    }
                }
                var v41 uint = 150
                _ = fi30
                _ = v31
                _ = v41
            }
            var v42 float32 = (1.5 / 1e10)
            if false {
                _ = len(v5)
            }
            _ = fi29
            _ = v42
        }
        v43 := GenS9{F0: int8(v25)}
        _ = v43
    } else {
        v44 := GenS8{F0: -2147483648}
        if false {
            fmt.Println(v25)
        } else {
            var v45 int32 = 862
            v4 = ((v4 + v4) + (v4 + v4))
            fmt.Println(v10, v1, v1)
            _ = v45
        }
        _ = v44
    }
    fmt.Println(v3)
    switch 127 {
    case 127:
        for fi46 := 0; fi46 < 6; fi46++ {
            v47 := GenS2{F0: fi46}
            v48 := GenS6{F0: (v1 / ((986) | 1))}
            if true {
                switch ("" + v4) {
                case "a":
                    if false {
                        fmt.Println(v25, v3)
                    }
                case "\n\t":
                    var v49 float32 = 1.5
                    _ = v49
                case "foo bar":
                    if true {
                        for fi50 := 0; fi50 < 7; fi50++ {
                            var v51 int = (-(fi46))
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            _ = fi50
                            _ = v51
                        }
                        v25 += v25
                    } else {
                        delete(v5, v10)
                    }
                default:
                    var v52 int64 = -727
                    _ = v52
                }
                v53 := GenS10{F0: true}
                _ = v53
            }
            _ = fi46
            _ = v47
            _ = v48
        }
    default:
        var v54 string = v10
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
