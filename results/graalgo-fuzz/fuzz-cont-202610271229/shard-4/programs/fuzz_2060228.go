package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 bool }
func genFn0() int {
    var v4 int32 = 2147483647
    v5 := GenS1{F0: 127}
    v4 -= (v4 | -716)
    v6 := []int{}
    _ = len(v6)
    _ = v4
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 int64) uint16 {
    if true {
        var v7 uint = 42
        _ = v7
    } else {
        v8 := GenS1{F0: 0}
        _ = v8
    }
    p0 = p0
    var v9 uint = 18446744073709551615
    _ = p0
    _ = v9
    return 0
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[string]string)
    v3 := []uint32{395, 670}
    v1 = ""
    var v10 string = "a"
    if true {
        for fi11 := 0; fi11 < 3; fi11++ {
            v1 = (v10 + ((v10 + "") + v1))
            var v12 float32 = 3.14159
            switch -2147483648 {
            case 237:
                v13 := v2[v10]
                _ = v13
            case -297:
                v10 = ""
            default:
                v12 += v12
            }
            _ = fi11
            _ = v12
        }
        switch -1 {
        case 869:
            v14 := v2[v1]
            _ = v14
        case 127:
            v1 = ("a" + (v10 + v10))
        default:
            for fi15 := 0; fi15 < 12; fi15++ {
                if (!false) {
                    v10 = (v1 + "foo bar")
                } else {
                    for fi16 := 0; fi16 < 4; fi16++ {
                        v17 := GenS2{F0: 0, F1: 296, F2: 336}
                        var v18 uint = 1
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    v10 = ("a" + v10)
                    var v19 uint32 = 280
                    _ = v19
                }
                for fi20 := 0; fi20 < 10; fi20++ {
                    v10 = v10
                    _ = fi20
                }
                _ = fi15
            }
        }
        for fi21 := 0; fi21 < 4; fi21++ {
            if true {
                v1 = v1
                for fi22 := 0; fi22 < 3; fi22++ {
                    v1 = (v1 + v10)
                    v23 := GenS2{F0: 0, F1: 873, F2: 18446744073709551615}
                    fmt.Println(fi22, fi21)
                    _ = fi22
                    _ = v23
                }
            } else {
                v24 := v2[((v10 + ("hello" + (("\n\t" + v1) + v1))) + (v1 + v10))]
                _ = v24
            }
            switch (v1 + "foo bar") {
            case "foo bar":
                v25 := GenS1{F0: 127}
                _ = v25
            default:
                switch v1 {
                case "a":
                    var v26 int = fi21
                    _ = v26
                case "foo bar":
                    var v27 uint8 = 2
                    _ = v27
                }
            }
            _ = fi21
        }
    } else {
        var v28 uint8 = 42
        _ = v28
    }
    v1 = ("a" + (v1 + ""))
    if false {
        v29 := GenS2{F0: 442, F1: 432, F2: 816}
        switch genFn0() {
        case 127:
            fmt.Println(v1)
        case 9223372036854775807:
            fmt.Println(v10, v1)
        case -38:
            v30 := GenS2{F0: 2147483647, F1: 4294967295, F2: 0}
            _ = v30
        default:
            v31 := GenS2{F0: -2147483648, F1: 1, F2: 1}
            _ = v31
        }
        _ = v29
    } else {
        for fi32 := 0; fi32 < 18; fi32++ {
            var v33 float32 = 0.0
            switch (v10 + "foo bar") {
            case "":
                if (42 > -27) {
                    var v34 uint64 = 814
                    _ = v34
                } else {
                    fmt.Println(v10)
                    v1 = v10
                }
            default:
                switch fi32 {
                case -495:
                    fmt.Println(v33, v1, v33)
                default:
                    if (!(false && (!(false && true)))) {
                        v1 = v10
                        var v35 bool = false
                        _ = v35
                    } else {
                        switch fi32 {
                        case 2147483647:
                            var v36 uint16 = 1
                            _ = v36
                        default:
                            v1 = ("" + (v10 + (v10 + "\n\t")))
                        }
                        if (-32768 <= 984) {
                            var v37 int16 = 1
                            var v38 bool = ((2 == genFn1(-1)) && (!false))
                            _ = v37
                            _ = v38
                        }
                        for fi39 := 0; fi39 < 17; fi39++ {
                            var v40 string = ("" + "foo bar")
                            _ = len(v2)
                            fmt.Println(v10, fi32, v10)
                            _ = fi39
                            _ = v40
                        }
                    }
                }
            }
            _ = fi32
            _ = v33
        }
        fmt.Println(v10, v1)
        switch -980 {
        case 497:
            var v41 int16 = 368
            _ = v41
        case 127:
            var v42 int32 = 42
            _ = v42
        case -480:
            switch 127 {
            case 32768:
                fmt.Println(v10, v10)
            case 32767:
                v43 := GenS3{F0: true, F1: 0}
                _ = v43
            case -1:
                fmt.Println(v1)
            default:
                switch v10 {
                case "\n\t":
                    for fi44 := 0; fi44 < 2; fi44++ {
                        var v45 uint16 = 0
                        v2[(v10 + ((v10 + v10) + v1))] = "hello"
                        v46 := GenS4{F0: false}
                        _ = fi44
                        _ = v45
                        _ = v46
                    }
                }
            }
        default:
            if (-32768 > 42) {
                v10 = v10
            }
        }
    }
    _ = len(v3)
    switch 2147483647 {
    case -2147483648:
        var v47 int = -608
        _ = v47
    default:
        switch 127 {
        case -9223372036854775808:
            for fi48 := 0; fi48 < 7; fi48++ {
                if len(v3) > 0 {
                    v3[0] = 0
                }
                v1 = "hello"
                for fi49 := 0; fi49 < 4; fi49++ {
                    var v50 string = ("\n\t" + ("" + v1))
                    _ = fi49
                    _ = v50
                }
                _ = fi48
            }
        case 1:
            for fi51 := 0; fi51 < 5; fi51++ {
                for fi52 := 0; fi52 < 13; fi52++ {
                    v53 := GenS4{F0: (-1 <= 1000)}
                    switch fi51 {
                    case 32768:
                        v54 := GenS2{F0: 0, F1: 4294967295, F2: 42}
                        _ = v54
                    }
                    _ = fi52
                    _ = v53
                }
                v1 = "foo bar"
                _ = fi51
            }
        default:
            var v55 uint64 = 18446744073709551615
            _ = v55
        }
    }
    var v56 int64 = 492
    delete(v2, v10)
    v2[(v1 + "foo bar")] = v1
    var v57 int16 = -32768
    v1 = (v1 + "foo bar")
    var v58 uint32 = 2147483647
    v56 += v56
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v56
    _ = v57
    _ = v58
}
