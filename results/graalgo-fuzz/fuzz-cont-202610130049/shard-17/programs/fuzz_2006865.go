package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 int16; F2 int16 }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8; F1 uint64; F2 int }
type GenS7 struct { F0 string }
func genFn0() int {
    var v5 uint = 2
    var v6 float64 = 0.0
    for fi7 := 0; fi7 < 8; fi7++ {
        fmt.Println(v5, v6)
        _ = fi7
    }
    var v8 uint8 = 127
    _ = v5
    _ = v6
    _ = v8
    return 4294967295
}
func genFn1(p0 bool, p1 uint, p2 bool) int {
    fmt.Println(p1, p0)
    fmt.Println(p2, p2, p1)
    var v9 uint8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -577
}

func main() {
    var v1 uint = 502
    var v2 int = -2147483648
    var v3 uint8 = 2
    v4 := []int{(v2 % (((v2 * (^(v2)))) | 1))}
    fmt.Println(v1, v1)
    for fi10 := 0; fi10 < 15; fi10++ {
        v1 = (v1 % ((v1) | 1))
        if false {
            v4 = append(v4, genFn0())
            switch ("\n\t" + "hello") {
            case "\n\t":
                if (int8(v3) > -128) {
                    switch (^(v2)) {
                    case 128:
                        for fi11 := 0; fi11 < 2; fi11++ {
                            v12 := GenS1{F0: (^(v3)), F1: 100}
                            v3 -= v3
                            _ = fi11
                            _ = v12
                        }
                    case -342:
                        v13 := GenS2{F0: fi10}
                        _ = v13
                    }
                    v3 -= (v3 & v3)
                } else {
                    var v14 int = (fi10 << ((uint(v3)) & 63))
                    v15 := GenS2{F0: (^(v2))}
                    _ = v14
                    _ = v15
                }
            case "foo bar":
                v16 := GenS3{F0: 0, F1: -1, F2: 32767}
                _ = v16
            }
            var v17 uint = 65535
            _ = v17
        } else {
            fmt.Println(fi10, v3)
            v2 -= (v2 * 127)
        }
        _ = fi10
    }
    if ((true && true) || (0.1 > ((1e10 / float64(((0.1 * 3.14159) - -2.25))) + 0.0))) {
        if (890 >= 42) {
            v1 -= v1
        }
        v18 := GenS4{F0: (v3 - (v3 >> (((v1 * (v1 / (((v1 % ((v1) | 1))) | 1)))) & 63))), F1: 2}
        _ = v18
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        if len(v4) > 0 {
            v4[0] = v2
        }
        _ = fi19
    }
    v4 = append(v4, (v2 + (v2 & 32767)))
    v3 -= (v3 >> ((v1) & 63))
    if (0.0 > 1e10) {
        var v20 uint32 = 164
        _ = v20
    }
    fmt.Println(v2, v2)
    if len(v4) > 0 {
        v4[0] = 89
    }
    if (65535 == 255) {
        if false {
            var v21 float64 = float64((-2.25 + 0.0))
            _ = v21
        }
        switch 520 {
        case -666:
            for fi22 := 0; fi22 < 19; fi22++ {
                fmt.Println(fi22, v2, v1)
                _ = fi22
            }
        case 1:
            if true {
                var v23 uint32 = 0
                switch 54 {
                case -1:
                    v2 -= -618
                case -2147483648:
                    v24 := GenS2{F0: int(v23)}
                    _ = v24
                default:
                    if true {
                        v23 = 42
                    } else {
                        for fi25 := 0; fi25 < 9; fi25++ {
                            for fi26 := 0; fi26 < 8; fi26++ {
                                fmt.Println(fi25, fi26)
                                _ = len(v4)
                                v27 := GenS1{F0: (v3 ^ 2), F1: 127}
                                switch "\n\t" {
                                case "foo bar":
                                    _ = len(v4)
                                default:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                _ = fi26
                                _ = v27
                            }
                            for fi28 := 0; fi28 < 3; fi28++ {
                                switch ("hello" + "\n\t") {
                                case "foo bar":
                                    if false {
                                        v1 = 4294967295
                                        v29 := GenS5{F0: -802}
                                        _ = v29
                                    } else {
                                        v4 = append(v4, 127)
                                        if false {
                                            switch int(v3) {
                                            case -9223372036854775808:
                                                var v30 string = ""
                                                _ = v30
                                            case 798:
                                                var v31 uint32 = (^(v23))
                                                _ = v31
                                            case -1:
                                                v2 -= fi25
                                            }
                                            if (!true) {
                                                v3 = v3
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            var v32 uint64 = 18446744073709551615
                                            _ = v32
                                        } else {
                                            var v33 bool = (!(42 != 854))
                                            v34 := GenS6{F0: 0, F1: 42, F2: 2147483647}
                                            fmt.Println(v23)
                                            _ = v33
                                            _ = v34
                                        }
                                    }
                                case "\n\t":
                                    v35 := GenS2{F0: 1}
                                    _ = v35
                                case "":
                                    fmt.Println(fi28)
                                default:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                var v36 bool = (((41 < -874) || ((fi25 | 128) < (^(fi28)))) && false)
                                _ = fi28
                                _ = v36
                            }
                            _ = fi25
                        }
                        v37 := GenS7{F0: "\n\t"}
                        if true {
                            var v38 int8 = 0
                            _ = v38
                        } else {
                            if (0 != 0) {
                                for fi39 := 0; fi39 < 12; fi39++ {
                                    var v40 int64 = -886
                                    _ = fi39
                                    _ = v40
                                }
                                for fi41 := 0; fi41 < 10; fi41++ {
                                    var v42 uint32 = v23
                                    switch fi41 {
                                    case 264:
                                        v2 -= 42
                                    case 180:
                                        v42 = uint32(v3)
                                    default:
                                        var v43 float32 = ((1.5 * 0.0) * 3.14159)
                                        _ = v43
                                    }
                                    var v44 uint16 = 42
                                    _ = fi41
                                    _ = v42
                                    _ = v44
                                }
                                var v45 uint32 = uint32(v3)
                                _ = v45
                            } else {
                                var v46 string = "\n\t"
                                for fi47 := 0; fi47 < 12; fi47++ {
                                    var v48 uint16 = 65535
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    _ = fi47
                                    _ = v48
                                }
                                _ = len(v4)
                                _ = v46
                            }
                        }
                        _ = v37
                    }
                }
                _ = v23
            } else {
                if true {
                    v3 -= (v3 + 128)
                    for fi49 := 0; fi49 < 14; fi49++ {
                        for fi50 := 0; fi50 < 17; fi50++ {
                            v2 = v2
                            fmt.Println(v2, fi50)
                            _ = fi50
                        }
                        for fi51 := 0; fi51 < 4; fi51++ {
                            var v52 float64 = (3.14159 + ((float64(0.0) / 0.1) / -2.25))
                            var v53 uint = 2
                            var v54 int16 = 1000
                            _ = fi51
                            _ = v52
                            _ = v53
                            _ = v54
                        }
                        _ = fi49
                    }
                } else {
                    switch v2 {
                    case 1:
                        if (false || false) {
                            v1 += (v1 & v1)
                        }
                    case -102:
                        v55 := GenS5{F0: -985}
                        _ = v55
                    case -9223372036854775808:
                        v3 = v3
                    }
                }
                _ = len(v4)
                if (448 >= 42) {
                    for fi56 := 0; fi56 < 12; fi56++ {
                        switch ("foo bar" + "a") {
                        case "foo bar":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        v2 += (^(v2))
                        _ = fi56
                    }
                    for fi57 := 0; fi57 < 18; fi57++ {
                        v1 = (v1 / (((v1 % (((v1 << ((v1) & 63))) | 1))) | 1))
                        _ = fi57
                    }
                    v2 += int(v1)
                }
            }
        default:
            var v58 bool = (0.1 >= (0.0 / -2.25))
            _ = v58
        }
        var v59 float32 = (3.14159 + (((float32((0.1 * 0.1)) * 0.0) * float32((-2.25 + 1e10))) + -2.25))
        _ = v59
    }
    v2 += 32768
    if len(v4) > 0 {
        v4[0] = v2
    }
    var v60 uint32 = 551
    switch v2 {
    case 1:
        var v61 int8 = -22
        _ = v61
    case -9223372036854775808:
        v62 := GenS3{F0: 127, F1: -32768, F2: 42}
        _ = v62
    default:
        var v63 uint = (v1 % ((v1) | 1))
        _ = v63
    }
    var v64 uint64 = 880
    var v65 uint64 = (v64 >> ((4294967295) & 63))
    v4 = append(v4, (v2 - (^(v2))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v60:", v60)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v60
    _ = v64
    _ = v65
}
