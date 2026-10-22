package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 string; F2 uint32 }
type GenS4 struct { F0 string; F1 int8; F2 uint8 }
type GenS5 struct { F0 int; F1 string; F2 uint32 }

func main() {
    v1 := make(map[int]int)
    var v2 string = ""
    v3 := []int{774, 127}
    v4 := make(map[int]bool)
    var v5 int64 = 34
    v6 := GenS1{F0: (-(v5)), F1: false, F2: 4294967295}
    delete(v1, 42)
    var v7 string = v2
    fmt.Println(v2, v7, v7)
    v7 = (((("\n\t" + (("\n\t" + (v7 + "foo bar")) + "")) + v2) + "hello") + "foo bar")
    v8 := GenS2{F0: 32767}
    var v9 int64 = v5
    var v10 bool = (false && false)
    if v10 {
        v10 = (2147483647 <= 32767)
        if v10 {
            v11 := GenS1{F0: (-(v9)), F1: (v10 && (1 > 1)), F2: 42}
            _ = v11
        } else {
            delete(v1, 42)
            if (v10 && ((v7 + ((v2 + v7) + v2)) != (v2 + v2))) {
                for fi12 := 0; fi12 < 15; fi12++ {
                    for fi13 := 0; fi13 < 3; fi13++ {
                        var v14 uint16 = 0
                        _ = fi13
                        _ = v14
                    }
                    v7 = v2
                    switch ((v2 + v2) + v7) {
                    case "":
                        v5 = v5
                    case "foo bar":
                        for fi15 := 0; fi15 < 5; fi15++ {
                            if (!((!v10) && false)) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v10 = (-1 > 0)
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            _ = fi15
                        }
                    }
                    _ = fi12
                }
            } else {
                fmt.Println(v2)
            }
            switch 0 {
            case -740:
                for fi16 := 0; fi16 < 11; fi16++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v17 float32 = (-2.25 - float32(0.1))
                    _ = fi16
                    _ = v17
                }
            case 127:
                fmt.Println(v9, v5)
            case -1:
                if (v10 || (!((!(!true)) && ((!(42 != 0)) || true)))) {
                    var v18 int16 = -32768
                    if true {
                        switch -485 {
                        case -993:
                            switch 4294967295 {
                            case -2147483648:
                                if (32767 <= 1) {
                                    var v19 uint64 = 42
                                    v5 += v9
                                    var v20 uint32 = 1
                                    _ = v19
                                    _ = v20
                                } else {
                                    v10 = false
                                    v21 := v4[9223372036854775807]
                                    _ = v21
                                }
                            default:
                                for fi22 := 0; fi22 < 6; fi22++ {
                                    v18 += (v18 * 0)
                                    var v23 int64 = (v5 + (v9 & (v5 % ((v5) | 1))))
                                    var v24 int = fi22
                                    _ = fi22
                                    _ = v23
                                    _ = v24
                                }
                            }
                        case 0:
                            v2 = (v7 + v2)
                        case -1:
                            fmt.Println(v5, v5)
                        default:
                            v25 := GenS2{F0: v18}
                            _ = v25
                        }
                    } else {
                        switch -1 {
                        case 9223372036854775807:
                            v2 = (v2 + v7)
                        case 750:
                            if false {
                                if (255 <= 2) {
                                    switch 4294967295 {
                                    case -1:
                                        var v26 int16 = 0
                                        _ = v26
                                    default:
                                        v9 += v5
                                    }
                                    switch -502 {
                                    case -1:
                                        for fi27 := 0; fi27 < 20; fi27++ {
                                            v5 = (v9 ^ 42)
                                            v5 -= v9
                                            var v28 uint = 4294967295
                                            _ = fi27
                                            _ = v28
                                        }
                                    default:
                                        v29 := GenS1{F0: (v5 & (v5 >> ((18446744073709551615) & 63))), F1: ((!v10) && (266 >= 255)), F2: 18446744073709551615}
                                        _ = v29
                                    }
                                } else {
                                    var v30 bool = false
                                    v18 = (v18 + v18)
                                    _ = v30
                                }
                                if (0 >= 42) {
                                    v31 := GenS3{F0: 0, F1: v7, F2: 350}
                                    for fi32 := 0; fi32 < 11; fi32++ {
                                        _ = len(v4)
                                        v5 = (v9 + (v9 / ((v9) | 1)))
                                        _ = fi32
                                    }
                                    if ((!false) && (!v10)) {
                                        v33 := v1[127]
                                        v34 := GenS2{F0: 32767}
                                        if true {
                                            switch v7 {
                                            case "":
                                                _ = len(v3)
                                            case "\n\t":
                                                var v35 uint32 = 0
                                                _ = v35
                                            case "hello":
                                                var v36 int32 = -435
                                                _ = v36
                                            default:
                                                if len(v3) > 0 {
                                                    v3[0] = int(v5)
                                                }
                                            }
                                            for fi37 := 0; fi37 < 16; fi37++ {
                                                if (((-(v33)) > fi37) && v10) {
                                                    var v38 uint8 = 0
                                                    switch v7 {
                                                    case "a":
                                                        for fi39 := 0; fi39 < 15; fi39++ {
                                                            v40 := GenS4{F0: ((v2 + "\n\t") + "hello"), F1: 42, F2: 79}
                                                            _ = fi39
                                                            _ = v40
                                                        }
                                                    case "":
                                                        for fi41 := 0; fi41 < 2; fi41++ {
                                                            var v42 string = "a"
                                                            _ = fi41
                                                            _ = v42
                                                        }
                                                    case "foo bar":
                                                        var v43 float32 = 0.0
                                                        _ = v43
                                                    }
                                                    _ = v38
                                                } else {
                                                    v1[(v33 % ((fi37) | 1))] = v33
                                                }
                                                var v44 string = v2
                                                if v10 {
                                                    v2 = (v44 + v7)
                                                } else {
                                                    var v45 int16 = v18
                                                    _ = v45
                                                }
                                                _ = fi37
                                                _ = v44
                                            }
                                        }
                                        _ = v33
                                        _ = v34
                                    } else {
                                        _ = len(v3)
                                        for fi46 := 0; fi46 < 17; fi46++ {
                                            for fi47 := 0; fi47 < 4; fi47++ {
                                                v7 = (("hello" + "\n\t") + v2)
                                                var v48 uint16 = 42
                                                _ = fi47
                                                _ = v48
                                            }
                                            _ = fi46
                                        }
                                    }
                                    _ = v31
                                } else {
                                    for fi49 := 0; fi49 < 3; fi49++ {
                                        var v50 string = "\n\t"
                                        fmt.Println(v2, v2)
                                        _ = fi49
                                        _ = v50
                                    }
                                    fmt.Println(v5, v5)
                                    var v51 uint64 = 2
                                    _ = v51
                                }
                                var v52 string = ((v2 + v2) + "foo bar")
                                _ = v52
                            } else {
                                switch "a" {
                                case "a":
                                    var v53 int32 = 42
                                    _ = v53
                                }
                                v10 = true
                            }
                        case 128:
                            v9 -= (v5 - 42)
                        }
                        var v54 float32 = 1e10
                        _ = v54
                    }
                    _ = v18
                } else {
                    _ = len(v1)
                    v3 = append(v3, -9223372036854775808)
                }
            }
        }
    }
    for fi55 := 0; fi55 < 8; fi55++ {
        if (2 > 42) {
            fmt.Println(fi55)
            var v56 uint16 = 738
            _ = v56
        } else {
            var v57 uint64 = 0
            if ("" != (v7 + v7)) {
                for fi58 := 0; fi58 < 7; fi58++ {
                    if (!(v2 >= v2)) {
                        var v59 int64 = (-(v9))
                        if v10 {
                            var v60 int8 = 1
                            fmt.Println(v60, v9)
                            _ = v60
                        }
                        v2 = v2
                        _ = v59
                    }
                    v61 := GenS2{F0: -1}
                    v62 := GenS1{F0: v9, F1: (-6 > -128), F2: 42}
                    _ = fi58
                    _ = v61
                    _ = v62
                }
            } else {
                if true {
                    if len(v3) > 0 {
                        v3[0] = fi55
                    }
                    v63 := GenS5{F0: fi55, F1: "foo bar", F2: 4294967295}
                    _ = v63
                } else {
                    fmt.Println(v9)
                }
                var v64 int = 42
                _ = v64
            }
            _ = v57
        }
        fmt.Println(v2, v9)
        _ = fi55
    }
    v2 = ("\n\t" + v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
