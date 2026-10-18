package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 string }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 string; F1 uint; F2 uint32 }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 18446744073709551615
    v3 := make(map[int]uint32)
    v4 := []int{4294967295, -316}
    v5 := v1[999]
    fmt.Println(v5, v5)
    v2 -= 42
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 int8 = 1
        if false {
            fmt.Println(v7, v5)
        } else {
            v2 -= 20
        }
        _ = fi6
        _ = v7
    }
    _ = len(v4)
    switch v5 {
    case "\n\t":
        for fi8 := 0; fi8 < 3; fi8++ {
            if true {
                v9 := v3[-221]
                var v10 uint16 = 947
                var v11 uint32 = 0
                _ = v10
                _ = v11
                _ = v9
            } else {
                v5 = ("hello" + (v5 + ""))
                if true {
                    var v12 uint = 18446744073709551615
                    v5 = ""
                    _ = v12
                }
                for fi13 := 0; fi13 < 9; fi13++ {
                    v14 := GenS1{F0: 65535, F1: 65535, F2: (v5 + ((((v5 + (v5 + "foo bar")) + v5) + v5) + ("a" + v5)))}
                    switch "hello" {
                    case "\n\t":
                        var v15 uint32 = 63
                        _ = v15
                    case "foo bar":
                        switch "\n\t" {
                        case "":
                            v5 = v5
                        default:
                            switch (((v5 + v5) + ((v5 + v5) + v5)) + ("" + ("" + v5))) {
                            case "a":
                                switch (v5 + "hello") {
                                case "":
                                    _ = len(v3)
                                case "foo bar":
                                    var v16 float64 = 0.0
                                    _ = v16
                                case "a":
                                    switch (v5 + (v5 + v5)) {
                                    case "foo bar":
                                        var v17 int = fi8
                                        _ = v17
                                    case "a":
                                        if false {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        } else {
                                            fmt.Println(fi13)
                                            var v18 uint32 = 2
                                            _ = v18
                                        }
                                    case "\n\t":
                                        var v19 int = fi13
                                        _ = v19
                                    }
                                }
                            case "\n\t":
                                switch fi8 {
                                case 32767:
                                    v20 := GenS1{F0: 744, F1: 142, F2: v5}
                                    _ = v20
                                default:
                                    var v21 int16 = 32767
                                    _ = v21
                                }
                                switch v5 {
                                case "":
                                    v5 = ""
                                case "foo bar":
                                    v22 := GenS1{F0: 309, F1: 598, F2: v5}
                                    _ = v22
                                }
                                if true {
                                    var v23 int = fi13
                                    delete(v3, -759)
                                    v24 := v3[(^(fi13))]
                                    _ = v23
                                    _ = v24
                                } else {
                                    if (0 <= 255) {
                                        var v25 string = (((v5 + v5) + v5) + (("" + v5) + "hello"))
                                        _ = v25
                                    } else {
                                        var v26 bool = true
                                        var v27 float64 = (3.14159 + (3.14159 / -2.25))
                                        _ = v26
                                        _ = v27
                                    }
                                    v2 += 42
                                    v28 := v1[1]
                                    v29 := GenS1{F0: 65535, F1: 2, F2: (v28 + v28)}
                                    _ = v28
                                    _ = v29
                                }
                            case "hello":
                                var v30 int16 = 760
                                _ = v30
                            default:
                                v2 = 2
                            }
                        }
                    }
                    _ = fi13
                    _ = v14
                }
            }
            _ = fi8
        }
    case "foo bar":
        if (false && true) {
            fmt.Println(v5)
        } else {
            var v31 int = 127
            var v32 int64 = 9223372036854775807
            _ = v31
            _ = v32
        }
    case "a":
        v33 := GenS1{F0: 42, F1: uint16(v2), F2: v5}
        _ = v33
    }
    v1[2147483648] = v5
    switch -860 {
    case 1:
        var v34 int16 = 42
        _ = v34
    case 0:
        switch ("\n\t" + "hello") {
        case "":
            var v35 int8 = -1
            _ = v35
        case "foo bar":
            if ((!(42 <= -191)) && true) {
                if true {
                    v36 := GenS1{F0: 377, F1: 1, F2: "hello"}
                    _ = v36
                }
                v37 := GenS1{F0: 42, F1: 65535, F2: ""}
                if len(v4) > 0 {
                    v4[0] = 42
                }
                _ = v37
            }
        }
    case -2147483648:
        for fi38 := 0; fi38 < 9; fi38++ {
            if false {
                v2 = v2
            }
            _ = fi38
        }
    }
    v39 := GenS2{F0: 666, F1: v5}
    var v40 string = v5
    for fi41 := 0; fi41 < 8; fi41++ {
        if true {
            v42 := GenS1{F0: 0, F1: 2, F2: v5}
            _ = v42
        }
        switch (fi41 & fi41) {
        case -9223372036854775808:
            var v43 uint16 = 32767
            _ = v43
        case 0:
            if (255 < 65535) {
                if (42 == 42) {
                    _ = len(v4)
                    _ = len(v1)
                    if (false || ((3.14159 - 1e10) != 1.5)) {
                        if len(v4) > 0 {
                            v4[0] = fi41
                        }
                        v3[fi41] = 366
                        fmt.Println(fi41)
                    } else {
                        v44 := v1[684]
                        v2 += 683
                        _ = v44
                    }
                }
            } else {
                if (true && false) {
                    v5 = "foo bar"
                    v45 := GenS3{F0: ((v5 + v5) + v40), F1: 4294967295, F2: 218}
                    v2 += (v2 * (v2 ^ 0))
                    _ = v45
                }
                v46 := GenS4{F0: "\n\t", F1: true}
                _ = v46
            }
        case 625:
            var v47 float64 = 1e10
            _ = v47
        default:
            switch fi41 {
            case 883:
                v40 = v5
            case -384:
                if false {
                    _ = len(v4)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    fmt.Println(v40, v40, v40)
                    _ = len(v1)
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(v2, v5)
                    case "hello":
                        v48 := GenS2{F0: 18446744073709551615, F1: "a"}
                        _ = v48
                    default:
                        v40 = ""
                    }
                }
            case 404:
                fmt.Println(v2, v40, v40)
            default:
                if (134 > 0) {
                    var v49 float32 = 1.5
                    _ = len(v4)
                    _ = v49
                } else {
                    for fi50 := 0; fi50 < 1; fi50++ {
                        _ = len(v1)
                        _ = fi50
                    }
                    v51 := GenS5{F0: -1, F1: v40}
                    _ = v51
                }
            }
        }
        var v52 int32 = 2147483647
        _ = fi41
        _ = v52
    }
    var v53 uint = 2
    v54 := GenS3{F0: v5, F1: (v53 << (((v53 << ((0) & 63))) & 63)), F2: 436}
    v53 += (v53 ^ v53)
    for fi55 := 0; fi55 < 3; fi55++ {
        fmt.Println(fi55, v53, v40)
        _ = fi55
    }
    for fi56 := 0; fi56 < 12; fi56++ {
        if (-32768 == -1) {
            v57 := GenS6{F0: (v2 * (^(v2)))}
            v40 = v5
            _ = v57
        }
        fmt.Println(v5, v2, v53)
        _ = fi56
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v53
    _ = v54
}
