package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 bool; F2 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 bool; F1 uint64; F2 int }
type GenS6 struct { F0 int32 }
func genFn0() int8 {
    var v4 string = "a"
    v4 = v4
    if true {
        v4 = v4
        var v5 uint32 = 486
        _ = v5
    } else {
        switch -9223372036854775808 {
        case 1:
            var v6 uint16 = 505
            _ = v6
        case 127:
            if false {
                if (1 < 1) {
                    v7 := []int{1}
                    if (65535 > 2) {
                        v4 = v4
                    } else {
                        if false {
                            for fi8 := 0; fi8 < 6; fi8++ {
                                v4 = ("" + ((((v4 + v4) + "foo bar") + v4) + ("foo bar" + ("a" + (v4 + ("hello" + ""))))))
                                var v9 int16 = -948
                                var v10 int8 = 100
                                _ = fi8
                                _ = v10
                                _ = v9
                            }
                            if ((1 < 2) || (2 < 42)) {
                                for fi11 := 0; fi11 < 9; fi11++ {
                                    if true {
                                        var v12 float64 = 1.5
                                        for fi13 := 0; fi13 < 7; fi13++ {
                                            v12 -= float64(1.5)
                                            v14 := GenS1{F0: (fi11 << ((187) & 63))}
                                            var v15 int8 = 42
                                            _ = fi13
                                            _ = v14
                                            _ = v15
                                        }
                                        if len(v7) > 0 {
                                            v7[0] = (fi11 & fi11)
                                        }
                                        _ = v12
                                    }
                                    _ = fi11
                                }
                            } else {
                                if false {
                                    v16 := GenS2{F0: -1, F1: true, F2: (v4 + v4)}
                                    _ = v16
                                } else {
                                    for fi17 := 0; fi17 < 12; fi17++ {
                                        switch v4 {
                                        case "foo bar":
                                            var v18 uint16 = 42
                                            _ = v18
                                        }
                                        v19 := GenS3{F0: 491}
                                        _ = fi17
                                        _ = v19
                                    }
                                    v4 = v4
                                }
                                var v20 uint16 = 2
                                v21 := GenS4{F0: 607, F1: -32768}
                                _ = v20
                                _ = v21
                            }
                            switch -110 {
                            case 32767:
                                for fi22 := 0; fi22 < 3; fi22++ {
                                    v7 = append(v7, (^(fi22)))
                                    _ = fi22
                                }
                            case -1:
                                var v23 int = -9223372036854775808
                                _ = v23
                            case 2:
                                v4 = v4
                            }
                        }
                        var v24 float64 = 1e10
                        _ = v24
                    }
                    _ = v7
                } else {
                    var v25 int64 = -1
                    _ = v25
                }
                v26 := GenS3{F0: 909}
                switch v4 {
                case "\n\t":
                    var v27 uint = 18446744073709551615
                    _ = v27
                case "foo bar":
                    var v28 int16 = -1
                    _ = v28
                case "a":
                    for fi29 := 0; fi29 < 14; fi29++ {
                        switch v4 {
                        case "hello":
                            fmt.Println(v4, v4)
                        case "a":
                            fmt.Println(v4, fi29)
                        default:
                            if (true || (!false)) {
                                if false {
                                    fmt.Println(v4, v4)
                                    if false {
                                        fmt.Println(fi29, v4)
                                        v30 := []int{fi29, (^(fi29)), (fi29 & 128)}
                                        _ = v30
                                    } else {
                                        var v31 float64 = 1.5
                                        var v32 int32 = int32(fi29)
                                        _ = v31
                                        _ = v32
                                    }
                                } else {
                                    v4 = v4
                                }
                            } else {
                                v4 = "hello"
                                fmt.Println(fi29, v4)
                            }
                        }
                        v33 := GenS2{F0: 0, F1: (!true), F2: "hello"}
                        _ = fi29
                        _ = v33
                    }
                default:
                    v34 := GenS5{F0: (2147483648 >= 2147483647), F1: 42, F2: 9223372036854775807}
                    _ = v34
                }
                _ = v26
            }
        }
    }
    fmt.Println(v4)
    v35 := GenS2{F0: 100, F1: (float64(1.5) != float64((float32((1e10 - float64(3.14159))) - (((1.5 - float32(0.1)) - -2.25) - (0.1 + 1e10))))), F2: "a"}
    _ = v35
    _ = v4
    return 127
}

func main() {
    var v1 int32 = 42
    var v2 int32 = -807
    var v3 bool = (42 <= 2147483647)
    for fi36 := 0; fi36 < 6; fi36++ {
        for fi37 := 0; fi37 < 13; fi37++ {
            var v38 float64 = 0.1
            _ = fi37
            _ = v38
        }
        _ = fi36
    }
    var v39 float64 = float64(float32((float64((0.0 - 0.0)) - 1.5)))
    fmt.Println(v3)
    var v40 int8 = 42
    var v41 int8 = v40
    if v3 {
        v42 := []int{-1, int(v40), 1, 32767}
        fmt.Println(v41, v3)
        _ = v42
    }
    switch "foo bar" {
    case "foo bar":
        var v43 uint = 690
        _ = v43
    }
    var v44 float32 = 0.1
    v44 += v44
    v3 = true
    fmt.Println(v44, v40)
    if ((127 >= 128) && (v2 == v2)) {
        v2 = v2
        if true {
            for fi45 := 0; fi45 < 18; fi45++ {
                v46 := []int{fi45}
                _ = fi45
                _ = v46
            }
            v47 := GenS2{F0: (^(v41)), F1: v3, F2: "a"}
            _ = v47
        }
    } else {
        var v48 uint64 = 4294967295
        v49 := GenS3{F0: 929}
        fmt.Println(v39, v40, v1)
        _ = v48
        _ = v49
    }
    var v50 int = -655
    var v51 int = -1
    v52 := GenS3{F0: int64(v50)}
    v53 := GenS6{F0: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v41
    _ = v44
    _ = v50
    _ = v51
    _ = v52
    _ = v53
}
