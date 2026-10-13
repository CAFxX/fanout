package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 string; F2 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 uint16; F1 uint16 }
func genFn0(p0 string, p1 int16, p2 bool) int32 {
    p0 = p0
    var v7 int64 = 9223372036854775807
    if p2 {
        for fi8 := 0; fi8 < 20; fi8++ {
            if (538 <= 2) {
                v9 := GenS1{F0: 2}
                _ = v9
            } else {
                var v10 uint16 = 913
                if (false && p2) {
                    v7 = v7
                    v11 := GenS1{F0: 105}
                    _ = v11
                } else {
                    switch fi8 {
                    case -2147483648:
                        v10 -= (v10 >> ((2) & 63))
                    case 0:
                        v12 := GenS1{F0: 312}
                        _ = v12
                    case 1:
                        var v13 int8 = -112
                        _ = v13
                    default:
                        v14 := []int{fi8}
                        _ = v14
                    }
                    fmt.Println(p2)
                }
                _ = v10
            }
            for fi15 := 0; fi15 < 20; fi15++ {
                if (p2 && ((v7 & v7) > 42)) {
                    var v16 string = p0
                    if (128 == 42) {
                        v17 := GenS1{F0: 18446744073709551615}
                        var v18 uint16 = 42
                        v18 = 32767
                        _ = v17
                        _ = v18
                    }
                    _ = v16
                }
                _ = fi15
            }
            _ = fi8
        }
        p2 = (127 == 92)
    } else {
        for fi19 := 0; fi19 < 3; fi19++ {
            if ((42 >= 100) || (0 > 0)) {
                var v20 int = 32768
                for fi21 := 0; fi21 < 14; fi21++ {
                    var v22 float32 = -2.25
                    _ = fi21
                    _ = v22
                }
                switch p0 {
                case "hello":
                    var v23 float64 = (0.1 - float64(float32((3.14159 / 1.0))))
                    _ = v23
                case "\n\t":
                    for fi24 := 0; fi24 < 13; fi24++ {
                        var v25 uint16 = 42
                        v26 := GenS1{F0: 467}
                        if p2 {
                            var v27 uint16 = v25
                            v28 := []int{-1}
                            _ = v27
                            _ = v28
                        } else {
                            v29 := []int{-2147483648, (fi24 % ((fi19) | 1)), (fi19 * fi19), v20}
                            fmt.Println(p0)
                            _ = v29
                        }
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                case "a":
                    v30 := GenS1{F0: 2}
                    _ = v30
                default:
                    for fi31 := 0; fi31 < 12; fi31++ {
                        var v32 uint16 = 0
                        _ = fi31
                        _ = v32
                    }
                }
                _ = v20
            } else {
                for fi33 := 0; fi33 < 12; fi33++ {
                    v34 := GenS1{F0: 395}
                    _ = fi33
                    _ = v34
                }
                p0 = (p0 + "\n\t")
                var v35 int32 = -631
                _ = v35
            }
            for fi36 := 0; fi36 < 6; fi36++ {
                var v37 uint64 = 1
                v37 += 1
                v7 = (v7 >> ((675) & 63))
                _ = fi36
                _ = v37
            }
            v38 := GenS1{F0: 42}
            _ = fi19
            _ = v38
        }
        if (float64(-2.25) <= 0.1) {
            switch (p0 + ("hello" + "a")) {
            case "foo bar":
                if p2 {
                    var v39 int64 = (v7 ^ v7)
                    _ = v39
                } else {
                    for fi40 := 0; fi40 < 5; fi40++ {
                        v41 := GenS2{F0: 42}
                        v42 := GenS3{F0: p0, F1: ((("foo bar" + ("a" + ("foo bar" + p0))) + "hello") + "hello"), F2: 970}
                        v43 := GenS4{F0: p1}
                        _ = fi40
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                    p2 = true
                }
            case "hello":
                for fi44 := 0; fi44 < 7; fi44++ {
                    var v45 string = (p0 + "hello")
                    _ = fi44
                    _ = v45
                }
            case "a":
                v46 := []int{4294967295, 42, 4294967295}
                _ = v46
            default:
                v47 := GenS5{F0: (v7 & v7), F1: 200}
                _ = v47
            }
        }
    }
    v48 := []int{128, 1, -1, 127}
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v48
    _ = v7
    return 2147483647
}
func genFn1() uint {
    var v49 uint16 = 65535
    var v50 uint8 = 128
    _ = v49
    _ = v50
    return 0
}

func main() {
    var v1 string = "hello"
    v2 := []uint32{4294967295}
    var v3 uint64 = 18446744073709551615
    var v4 uint16 = 42
    var v5 uint16 = v4
    var v6 int64 = 1
    v4 = 255
    var v51 uint64 = (v3 % ((627) | 1))
    var v52 uint16 = v4
    for fi53 := 0; fi53 < 12; fi53++ {
        v54 := GenS3{F0: (v1 + "foo bar"), F1: v1, F2: 960}
        switch fi53 {
        case 2147483647:
            v2 = append(v2, 926)
        case -499:
            var v55 string = (v1 + "")
            _ = v55
        case 401:
            v6 += -9223372036854775808
        }
        var v56 uint16 = 32767
        _ = fi53
        _ = v54
        _ = v56
    }
    v52 = (^(v52))
    var v57 uint = 0
    v2 = append(v2, 65535)
    v58 := GenS6{F0: uint16(v3), F1: (v52 << ((4294967295) & 63))}
    _ = len(v2)
    var v59 string = v1
    v4 -= v5
    v60 := GenS2{F0: v4}
    v57 = genFn1()
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v57:", v57)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v60
}
