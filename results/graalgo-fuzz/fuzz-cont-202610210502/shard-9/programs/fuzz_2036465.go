package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS6 struct { F0 int8; F1 int64 }
type GenS7 struct { F0 int32 }
func genFn0(p0 uint64, p1 int8, p2 uint16) uint8 {
    p0 -= (p0 / ((p0) | 1))
    var v7 bool = false
    if v7 {
        for fi8 := 0; fi8 < 20; fi8++ {
            if false {
                switch (fi8 >> ((uint(p2)) & 63)) {
                case -9223372036854775808:
                    var v9 uint32 = 1
                    _ = v9
                default:
                    p1 = (p1 - -128)
                }
            }
            _ = fi8
        }
    } else {
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 int = (^(fi10))
            var v12 uint8 = 244
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 int32 = -1
        _ = v13
    }
    switch "hello" {
    case "a":
        switch 375 {
        case 1:
            v14 := GenS1{F0: 4294967295}
            _ = v14
        case 2147483647:
            p0 += 42
        case 32768:
            p2 += p2
        default:
            var v15 uint32 = 2147483647
            _ = v15
        }
    case "\n\t":
        var v16 uint64 = 42
        _ = v16
    case "foo bar":
        fmt.Println(p0, v7, v7)
    default:
        if (0 <= 897) {
            v17 := GenS2{F0: 18446744073709551615, F1: (p2 * p2)}
            v18 := []int{-187, 128, -118}
            var v19 int64 = 0
            _ = v17
            _ = v18
            _ = v19
        } else {
            v20 := GenS1{F0: 0}
            p0 += (p0 << ((uint(p0)) & 63))
            _ = v20
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}
func genFn1(p0 string, p1 uint64, p2 int32) int32 {
    p1 = p1
    v21 := GenS2{F0: 2, F1: 337}
    v22 := GenS1{F0: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v22
    return (p2 >> ((65535) & 63))
}

func main() {
    v1 := []int{}
    var v2 uint32 = 972
    v3 := make(map[string]uint32)
    v4 := make(map[int]bool)
    var v5 uint8 = 144
    var v6 uint8 = v5
    var v23 string = "foo bar"
    for fi24 := 0; fi24 < 17; fi24++ {
        var v25 int8 = 42
        if true {
            if ((3.14159 - 0.0) != (0.0 - (1e10 - 0.0))) {
                v26 := GenS3{F0: 18446744073709551615}
                for fi27 := 0; fi27 < 15; fi27++ {
                    _ = len(v1)
                    _ = fi27
                }
                _ = v26
            }
            var v28 float64 = 3.14159
            _ = v28
        } else {
            fmt.Println(v25, v25)
            v6 = (v5 + v6)
        }
        _ = fi24
        _ = v25
    }
    v29 := GenS4{F0: 4294967295, F1: 0}
    var v30 uint = 928
    var v31 int8 = 1
    v2 = v2
    if true {
        v4[127] = ((v30 & (v30 * 4294967295)) >= (^(v30)))
        var v32 uint64 = 0
        _ = v32
    } else {
        _ = len(v3)
        fmt.Println(v2)
    }
    v31 = v31
    var v33 int32 = -2147483648
    delete(v3, v23)
    v23 = v23
    var v34 uint32 = 1
    if true {
        v35 := GenS1{F0: 0}
        _ = v35
    } else {
        if false {
            var v36 uint16 = uint16(v34)
            var v37 int64 = -9223372036854775808
            if len(v1) > 0 {
                v1[0] = 127
            }
            _ = v36
            _ = v37
        }
    }
    v38 := GenS5{F0: (v34 / ((uint32(v5)) | 1)), F1: v31, F2: genFn0(35, (v31 % ((v31) | 1)), 564)}
    if true {
        switch -1 {
        case 32768:
            v39 := GenS6{F0: 42, F1: 9223372036854775807}
            _ = v39
        default:
            for fi40 := 0; fi40 < 12; fi40++ {
                if false {
                    v30 = (^(v30))
                }
                _ = fi40
            }
        }
    } else {
        if true {
            var v41 uint = 0
            delete(v3, v23)
            var v42 int8 = (v31 >> ((v30) & 63))
            _ = v41
            _ = v42
        }
        v6 += v6
    }
    if true {
        for fi43 := 0; fi43 < 4; fi43++ {
            for fi44 := 0; fi44 < 13; fi44++ {
                var v45 int64 = 0
                _ = fi44
                _ = v45
            }
            v46 := GenS1{F0: v2}
            var v47 uint64 = 524
            _ = fi43
            _ = v46
            _ = v47
        }
        v33 -= (v33 | v33)
        var v48 uint64 = uint64(v5)
        _ = v48
    }
    if true {
        v30 = 18446744073709551615
    } else {
        for fi49 := 0; fi49 < 3; fi49++ {
            v30 -= 1
            _ = fi49
        }
        if true {
            fmt.Println(v6)
            for fi50 := 0; fi50 < 20; fi50++ {
                for fi51 := 0; fi51 < 2; fi51++ {
                    switch (v23 + "foo bar") {
                    case "":
                        v31 -= v31
                    case "hello":
                        if (true || (fi50 >= fi51)) {
                            v52 := GenS2{F0: 531, F1: 541}
                            _ = v52
                        }
                    case "a":
                        var v53 float64 = (-2.25 / 1e10)
                        _ = v53
                    }
                    v23 = ("a" + "")
                    _ = fi51
                }
                _ = fi50
            }
            for fi54 := 0; fi54 < 2; fi54++ {
                if (v30 != v30) {
                    for fi55 := 0; fi55 < 14; fi55++ {
                        v56 := GenS1{F0: v2}
                        v2 = (v34 ^ v34)
                        if false {
                            if true {
                                v5 -= v5
                                var v57 int8 = (v31 * (-(v31)))
                                v58 := GenS7{F0: (v33 / (((v33 / ((v33) | 1))) | 1))}
                                _ = v57
                                _ = v58
                            }
                            v23 = (v23 + "hello")
                            var v59 uint32 = (v34 % ((uint32(v30)) | 1))
                            fmt.Println(v31, v31)
                            _ = v59
                        }
                        _ = fi55
                        _ = v56
                    }
                } else {
                    if ((135 != uint16(v6)) || ((-(fi54)) <= fi54)) {
                        v60 := GenS5{F0: v2, F1: (v31 / ((-1) | 1)), F2: 231}
                        v2 -= (v34 | uint32(v6))
                        _ = v60
                    }
                    var v61 int8 = (v31 >> ((v30) & 63))
                    v23 = v23
                    _ = v61
                }
                if true {
                    var v62 uint64 = uint64(v30)
                    _ = v62
                } else {
                    _ = len(v1)
                }
                v63 := GenS6{F0: 127, F1: -9223372036854775808}
                _ = fi54
                _ = v63
            }
        }
    }
    v1 = append(v1, int(v33))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v34
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
