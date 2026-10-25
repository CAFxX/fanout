package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int16; F2 string }
type GenS3 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS4 struct { F0 int32; F1 bool; F2 uint64 }
type GenS5 struct { F0 int8; F1 uint32; F2 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    var v3 uint64 = 2
    v4 := make(map[string]int)
    var v5 bool = true
    var v6 uint32 = 468
    fmt.Println(v5, v6)
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint64 = 1
        _ = fi7
        _ = v8
    }
    if v5 {
        v5 = v5
        if (1e10 >= 0.1) {
            var v9 float64 = 0.1
            switch int(v3) {
            case 9223372036854775807:
                if false {
                    if (v5 || true) {
                        fmt.Println(v6, v5, v3)
                        var v10 uint8 = 82
                        var v11 uint64 = (v3 + (v3 * v3))
                        _ = v10
                        _ = v11
                    } else {
                        v12 := v1[(v6 << ((65535) & 63))]
                        _ = v12
                    }
                    if (("a" < "a") || true) {
                        var v13 float32 = (0.0 / 1.0)
                        _ = v13
                    }
                    v14 := v1[0]
                    _ = v14
                }
            case 42:
                if (!true) {
                    var v15 uint64 = (v3 >> ((710) & 63))
                    _ = v15
                }
            }
            for fi16 := 0; fi16 < 2; fi16++ {
                _ = len(v1)
                v17 := GenS1{F0: (-(fi16))}
                for fi18 := 0; fi18 < 17; fi18++ {
                    var v19 string = "\n\t"
                    delete(v2, v6)
                    if (32767 <= int(v3)) {
                        v3 -= 18446744073709551615
                        var v20 uint8 = 174
                        if (42 < -1) {
                            var v21 int = -9223372036854775808
                            _ = v21
                        }
                        v6 = 65535
                        _ = v20
                    } else {
                        v1[v6] = (v6 + v6)
                    }
                    _ = fi18
                    _ = v19
                }
                _ = fi16
                _ = v17
            }
            _ = v9
        }
    } else {
        var v22 int = 1
        _ = v22
    }
    if ((-626 >= 9223372036854775807) && (!(100 == 127))) {
        v6 = v6
        v23 := GenS2{F0: 4294967295, F1: -498, F2: "a"}
        _ = v23
    }
    v24 := []int{}
    fmt.Println(v6)
    switch -372 {
    case 2147483647:
        v25 := GenS1{F0: 0}
        _ = v25
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        if (v5 && v5) {
            var v27 float32 = 0.1
            _ = v27
        } else {
            for fi28 := 0; fi28 < 19; fi28++ {
                var v29 string = "foo bar"
                _ = fi28
                _ = v29
            }
        }
        _ = fi26
    }
    for fi30 := 0; fi30 < 15; fi30++ {
        var v31 uint32 = 874
        if len(v24) > 0 {
            v24[0] = -2147483648
        }
        var v32 string = ""
        _ = fi30
        _ = v31
        _ = v32
    }
    var v33 int32 = 1
    if true {
        if v5 {
            v34 := v4[(("foo bar" + "a") + "a")]
            _ = v34
        } else {
            v35 := GenS1{F0: -1}
            if v5 {
                v3 = 2
                if v5 {
                    v5 = true
                    v6 -= 2147483647
                } else {
                    v3 -= 42
                    var v36 float32 = 1.5
                    _ = v36
                }
                v6 += v6
            }
            _ = v35
        }
    } else {
        if false {
            var v37 string = "foo bar"
            if v5 {
                for fi38 := 0; fi38 < 7; fi38++ {
                    for fi39 := 0; fi39 < 20; fi39++ {
                        fmt.Println(v33, fi38)
                        v40 := v1[13]
                        var v41 float32 = 3.14159
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                    var v42 float64 = -2.25
                    v43 := GenS3{F0: v6, F1: 1, F2: (v33 - v33)}
                    _ = fi38
                    _ = v42
                    _ = v43
                }
                for fi44 := 0; fi44 < 2; fi44++ {
                    switch fi44 {
                    case 723:
                        if true {
                            v3 += v3
                            _ = len(v2)
                        }
                    case 844:
                        switch ((v37 + (v37 + v37)) + "foo bar") {
                        case "a":
                            for fi45 := 0; fi45 < 20; fi45++ {
                                var v46 float32 = (0.0 - (0.0 / -2.25))
                                var v47 uint16 = 0
                                v5 = (!(v3 == (v3 >> ((4294967295) & 63))))
                                _ = fi45
                                _ = v46
                                _ = v47
                            }
                        case "\n\t":
                            switch "hello" {
                            case "hello":
                                var v48 string = "hello"
                                _ = v48
                            case "":
                                if len(v24) >= 2 {
                                    v24 = v24[1:len(v24)]
                                }
                            }
                        case "":
                            switch v37 {
                            case "":
                                var v49 uint = 42
                                _ = v49
                            case "a":
                                for fi50 := 0; fi50 < 8; fi50++ {
                                    var v51 float64 = 0.0
                                    _ = fi50
                                    _ = v51
                                }
                            default:
                                var v52 int64 = -9223372036854775808
                                _ = v52
                            }
                        }
                    default:
                        v53 := GenS4{F0: v33, F1: v5, F2: v3}
                        _ = v53
                    }
                    v5 = (((0.0 + 0.0) * (-2.25 - 0.0)) == float32(0.1))
                    var v54 int8 = 100
                    _ = fi44
                    _ = v54
                }
                v55 := GenS2{F0: 840, F1: -1, F2: v37}
                _ = v55
            }
            switch v37 {
            case "foo bar":
                for fi56 := 0; fi56 < 15; fi56++ {
                    _ = len(v24)
                    var v57 int8 = 127
                    var v58 float64 = 1e10
                    _ = fi56
                    _ = v57
                    _ = v58
                }
            case "a":
                fmt.Println(v6)
            case "hello":
                fmt.Println(v37, v33)
            }
            _ = v37
        } else {
            if v5 {
                for fi59 := 0; fi59 < 19; fi59++ {
                    var v60 int16 = -797
                    _ = fi59
                    _ = v60
                }
                v6 = 0
            } else {
                fmt.Println(v3)
                if ("hello" != "") {
                    var v61 int8 = 127
                    _ = v61
                }
                var v62 string = "foo bar"
                _ = v62
            }
            if (355 == 1) {
                fmt.Println(v33, v6, v5)
                var v63 bool = v5
                var v64 float64 = 0.1
                _ = v63
                _ = v64
            }
            delete(v2, v6)
        }
    }
    switch 622 {
    case 0:
        var v65 int64 = 1
        _ = v65
    default:
        v66 := GenS5{F0: 100, F1: (v6 * 638), F2: 42}
        _ = v66
    }
    for fi67 := 0; fi67 < 10; fi67++ {
        fmt.Println(v6, fi67)
        var v68 uint64 = 42
        _ = len(v24)
        _ = fi67
        _ = v68
    }
    for fi69 := 0; fi69 < 6; fi69++ {
        v24 = append(v24, 689)
        v6 += (v6 / (((v6 >> ((65535) & 63))) | 1))
        _ = fi69
    }
    v33 -= (^(v33))
    var v70 int8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v70
}
