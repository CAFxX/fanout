package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool }
func genFn0(p0 int32, p1 uint32, p2 int8) int16 {
    for fi6 := 0; fi6 < 2; fi6++ {
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 float64 = 0.0
            var v9 uint8 = 68
            v9 = (^(v9))
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    switch "foo bar" {
    case "\n\t":
        switch 127 {
        case -790:
            switch -2147483648 {
            case 128:
                var v10 int8 = p2
                _ = v10
            case 42:
                var v11 uint16 = uint16(p1)
                _ = v11
            default:
                var v12 int = 32768
                _ = v12
            }
        }
    case "hello":
        fmt.Println(p2)
    default:
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    var v1 float64 = 0.0
    var v2 uint32 = 4294967295
    v3 := []uint32{796, v2, 2147483647, (^(v2))}
    var v4 float64 = (0.0 + float64(((0.0 - -2.25) - ((((0.0 * 0.1) * (-2.25 + 0.0)) * 0.0) / 1.5))))
    var v5 float64 = v4
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v4, v5)
    }
    var v13 int16 = 0
    for fi14 := 0; fi14 < 1; fi14++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v13 += (v13 % ((v13) | 1))
        var v15 uint32 = 4294967295
        _ = fi14
        _ = v15
    }
    v4 = 3.14159
    for fi16 := 0; fi16 < 8; fi16++ {
        v1 += v4
        var v17 uint = 586
        var v18 uint8 = 42
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 float64 = ((v5 * v4) * v4)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        for fi21 := 0; fi21 < 9; fi21++ {
            if true {
                if true {
                    if true {
                        var v22 bool = ((v4 * 0.0) != 1e10)
                        _ = v22
                    }
                    _ = len(v3)
                } else {
                    v3 = append(v3, v2)
                }
            } else {
                var v23 uint64 = 301
                var v24 string = "a"
                _ = v23
                _ = v24
            }
            _ = fi21
        }
        _ = fi20
    }
    v25 := GenS1{F0: 65535}
    if (false && true) {
        var v26 uint8 = 0
        if true {
            if (959 == 4294967295) {
                for fi27 := 0; fi27 < 19; fi27++ {
                    switch 0 {
                    case 4294967295:
                        if ((3.14159 > 1.5) && true) {
                            for fi28 := 0; fi28 < 15; fi28++ {
                                v29 := GenS1{F0: 301}
                                switch ("\n\t" + "foo bar") {
                                case "a":
                                    switch -335 {
                                    case 0:
                                        v30 := GenS1{F0: v2}
                                        _ = v30
                                    }
                                case "foo bar":
                                    v3 = append(v3, (v2 ^ (v2 - (v2 | v2))))
                                }
                                _ = fi28
                                _ = v29
                            }
                            v26 -= v26
                        } else {
                            _ = len(v3)
                            v2 -= (v2 << ((255) & 63))
                        }
                    default:
                        for fi31 := 0; fi31 < 4; fi31++ {
                            switch fi31 {
                            case -1:
                                _ = len(v3)
                            case -9223372036854775808:
                                v13 = (-(v13))
                            default:
                                if (1 <= -1) {
                                    v19 += (float64((((1e10 - ((-2.25 / 1.0) + 3.14159)) + (0.1 - 0.0)) / float32(v1))) - (v4 + v4))
                                } else {
                                    fmt.Println(v4, v13)
                                    var v32 int = fi27
                                    if ((^(v26)) < 0) {
                                        var v33 int32 = 1
                                        v34 := GenS2{F0: 0}
                                        if true {
                                            v4 = v1
                                        }
                                        fmt.Println(v33)
                                        var v35 uint = 2
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    } else {
                                        v1 = 0.1
                                        v36 := GenS3{F0: (fi27 < v32)}
                                        fmt.Println(v2, fi27)
                                        _ = v36
                                    }
                                    _ = v32
                                }
                            }
                            _ = fi31
                        }
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi27
                }
                v19 = -2.25
                var v37 uint32 = 1
                _ = v37
            }
            v4 = (((1.5 + (1e10 / (0.0 - (1.5 / (3.14159 / v1))))) + v5) / v1)
        }
        _ = v26
    }
    v3 = append(v3, (v2 >> ((979) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
