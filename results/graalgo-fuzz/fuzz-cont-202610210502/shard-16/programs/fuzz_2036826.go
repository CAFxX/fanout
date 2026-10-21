package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint64; F2 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 uint; F1 string }
func genFn0() int {
    if (-1 == 1000) {
        var v7 uint16 = 90
        _ = v7
    }
    var v8 int16 = 99
    var v9 uint64 = 51
    var v10 string = "\n\t"
    var v11 uint16 = 42
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 128
}
func genFn1() uint16 {
    v12 := []int{}
    v13 := GenS1{F0: "a"}
    var v14 uint = 18446744073709551615
    if len(v12) > 0 {
        v12[0] = 9223372036854775807
    }
    _ = v12
    _ = v13
    _ = v14
    return 0
}

func main() {
    var v1 uint64 = 804
    var v2 bool = (("foo bar" <= "\n\t") && (47 < -1))
    v3 := make(map[uint32]int)
    var v4 uint64 = v1
    var v5 string = "\n\t"
    v6 := []int{}
    var v15 float64 = (float64(-2.25) + 0.0)
    v16 := GenS1{F0: (v5 + "foo bar")}
    var v17 int16 = 0
    if false {
        switch -9223372036854775808 {
        case 32767:
            switch (("hello" + (v5 + v5)) + v5) {
            case "\n\t":
                for fi18 := 0; fi18 < 4; fi18++ {
                    for fi19 := 0; fi19 < 10; fi19++ {
                        switch 645 {
                        case 835:
                            v2 = v2
                        case 32768:
                            v20 := GenS1{F0: "\n\t"}
                            _ = v20
                        }
                        _ = fi19
                    }
                    v21 := GenS1{F0: v5}
                    _ = fi18
                    _ = v21
                }
            case "":
                if false {
                    for fi22 := 0; fi22 < 5; fi22++ {
                        var v23 uint16 = 42
                        _ = len(v6)
                        _ = fi22
                        _ = v23
                    }
                    v15 += v15
                }
            case "foo bar":
                if v2 {
                    for fi24 := 0; fi24 < 19; fi24++ {
                        v5 = (v5 + "foo bar")
                        var v25 int16 = 0
                        for fi26 := 0; fi26 < 5; fi26++ {
                            var v27 uint8 = 51
                            var v28 uint = 1
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        _ = fi24
                        _ = v25
                    }
                    var v29 uint = 961
                    v2 = (592 == -135)
                    _ = v29
                } else {
                    fmt.Println(v5)
                    var v30 uint32 = 2147483648
                    _ = v30
                }
            }
        case -510:
            var v31 float64 = 1.5
            _ = v31
        }
        v32 := GenS2{F0: (v17 & v17), F1: v1, F2: v5}
        _ = v32
    }
    switch 128 {
    case 42:
        if v2 {
            fmt.Println(v17)
        }
    case 32767:
        if (370 > -134) {
            if false {
                v15 -= v15
                if v2 {
                    if (-521 < 0) {
                        for fi33 := 0; fi33 < 7; fi33++ {
                            v15 = v15
                            for fi34 := 0; fi34 < 16; fi34++ {
                                v35 := GenS1{F0: v5}
                                v17 -= 0
                                var v36 int8 = int8(v1)
                                switch "a" {
                                case "a":
                                    v2 = ((^(v36)) != v36)
                                case "":
                                    if (42 == v36) {
                                        if ((v17 & (v17 / ((v17) | 1))) < v17) {
                                            if len(v6) > 0 {
                                                v6[0] = fi34
                                            }
                                        } else {
                                            var v37 int32 = 74
                                            _ = v37
                                        }
                                        var v38 uint = 2
                                        var v39 uint64 = v1
                                        if (908 < (v39 >> (((v38 % ((18446744073709551615) | 1))) & 63))) {
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            if (v2 && (v2 || false)) {
                                                fmt.Println(v39, v39)
                                            } else {
                                                v40 := GenS3{F0: 1}
                                                _ = v40
                                            }
                                            switch "a" {
                                            case "a":
                                                if (4294967295 <= (^(v39))) {
                                                    if false {
                                                        fmt.Println(v5, v1)
                                                        switch (v5 + "foo bar") {
                                                        case "":
                                                            var v41 int64 = 9223372036854775807
                                                            _ = v41
                                                        case "foo bar":
                                                            var v42 uint32 = 2147483647
                                                            _ = v42
                                                        default:
                                                            v43 := GenS4{F0: v2, F1: 65535}
                                                            _ = v43
                                                        }
                                                    }
                                                } else {
                                                    v44 := GenS4{F0: (0 == 42), F1: 255}
                                                    _ = v44
                                                }
                                            case "\n\t":
                                                delete(v3, 65535)
                                            }
                                        }
                                        var v45 uint32 = 65535
                                        _ = v38
                                        _ = v39
                                        _ = v45
                                    }
                                    if v2 {
                                        switch v5 {
                                        case "hello":
                                            fmt.Println(v5)
                                        }
                                        var v46 uint64 = v4
                                        _ = v46
                                    } else {
                                        v47 := GenS5{F0: 152, F1: (v4 + (v4 << ((220) & 63)))}
                                        _ = v47
                                    }
                                }
                                var v48 uint = 42
                                if false {
                                    v15 = v15
                                    v17 += v17
                                }
                                v49 := GenS3{F0: 130}
                                _ = fi34
                                _ = v35
                                _ = v36
                                _ = v48
                                _ = v49
                            }
                            _ = fi33
                        }
                    }
                    if true {
                        if len(v6) > 0 {
                            v6[0] = 4294967295
                        }
                        delete(v3, 2)
                    } else {
                        if (58 >= 989) {
                            delete(v3, 4294967295)
                        }
                    }
                } else {
                    v50 := v3[65535]
                    _ = v50
                }
            }
        } else {
            for fi51 := 0; fi51 < 12; fi51++ {
                v17 = (v17 % ((-96) | 1))
                _ = fi51
            }
            v3[2147483647] = 0
            v52 := GenS3{F0: 127}
            _ = v52
        }
    case 0:
        fmt.Println(v1, v5, v2)
    }
    v15 = v15
    v53 := GenS2{F0: v17, F1: (v1 ^ (v4 & v4)), F2: v5}
    v1 -= (^(v1))
    if len(v6) > 0 {
        v6[0] = 127
    }
    for fi54 := 0; fi54 < 13; fi54++ {
        for fi55 := 0; fi55 < 4; fi55++ {
            var v56 uint = 4294967295
            v4 -= 42
            _ = fi55
            _ = v56
        }
        if true {
            if v2 {
                v1 += (^(v4))
                v57 := GenS3{F0: 230}
                _ = v57
            } else {
                if len(v6) > 0 {
                    v6[0] = (fi54 / (((fi54 / ((136) | 1))) | 1))
                }
                for fi58 := 0; fi58 < 11; fi58++ {
                    v2 = v2
                    var v59 uint8 = 1
                    var v60 float64 = v15
                    _ = fi58
                    _ = v59
                    _ = v60
                }
                v61 := GenS1{F0: "\n\t"}
                _ = v61
            }
            if true {
                for fi62 := 0; fi62 < 16; fi62++ {
                    var v63 uint16 = 333
                    var v64 int32 = -1
                    switch "a" {
                    case "hello":
                        fmt.Println(v1, v64, v63)
                    case "":
                        v1 -= v4
                        fmt.Println(fi62, v1)
                    case "a":
                        v65 := GenS4{F0: false, F1: (v63 + v63)}
                        _ = v65
                    }
                    _ = fi62
                    _ = v63
                    _ = v64
                }
                _ = len(v6)
                switch "a" {
                case "foo bar":
                    v66 := GenS6{F0: 255, F1: (("a" + v5) + ((v5 + v5) + v5))}
                    _ = v66
                case "\n\t":
                    v67 := GenS3{F0: 42}
                    _ = v67
                default:
                    v2 = (0.1 > 1e10)
                }
            }
        } else {
            fmt.Println(v2, fi54)
            var v68 int8 = -128
            _ = v68
        }
        _ = fi54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
}
