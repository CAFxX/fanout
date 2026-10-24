package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint8; F1 string; F2 uint16 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 int32; F1 uint8; F2 int }
type GenS5 struct { F0 int; F1 uint32 }
func genFn0(p0 bool, p1 uint, p2 uint64) bool {
    switch "\n\t" {
    case "":
        for fi4 := 0; fi4 < 7; fi4++ {
            fmt.Println(p1, p1)
            _ = fi4
        }
    case "hello":
        p1 += p1
    case "foo bar":
        fmt.Println(p1, p0, p2)
    default:
        v5 := GenS1{F0: 556, F1: 42}
        _ = v5
    }
    switch 32767 {
    case 128:
        if p0 {
            for fi6 := 0; fi6 < 1; fi6++ {
                var v7 uint8 = 2
                switch -2147483648 {
                case 128:
                    v8 := []int{252, fi6, fi6, fi6}
                    _ = v8
                case -2147483648:
                    var v9 string = ("" + ("a" + (("" + ("foo bar" + "foo bar")) + ("a" + "\n\t"))))
                    _ = v9
                case -9223372036854775808:
                    switch fi6 {
                    case 113:
                        switch int(p1) {
                        case -643:
                            if true {
                                var v10 uint32 = 2147483647
                                _ = v10
                            }
                        case 533:
                            v7 = (^(v7))
                        default:
                            var v11 int = fi6
                            _ = v11
                        }
                    case -2147483648:
                        v12 := GenS2{F0: v7, F1: "\n\t", F2: 324}
                        _ = v12
                    case 171:
                        v13 := GenS1{F0: 32767, F1: -475}
                        _ = v13
                    }
                default:
                    fmt.Println(fi6, v7)
                }
                p0 = false
                _ = fi6
                _ = v7
            }
        } else {
            var v14 int32 = 41
            for fi15 := 0; fi15 < 8; fi15++ {
                switch "\n\t" {
                case "foo bar":
                    for fi16 := 0; fi16 < 5; fi16++ {
                        p2 += (p2 << ((p1) & 63))
                        var v17 uint = p1
                        _ = fi16
                        _ = v17
                    }
                case "hello":
                    fmt.Println(v14, v14, p0)
                case "":
                    v14 += 2147483647
                default:
                    if (1 < 127) {
                        fmt.Println(p1)
                    }
                }
                v18 := []int{fi15, (^(fi15))}
                fmt.Println(p2, p1)
                _ = fi15
                _ = v18
            }
            _ = v14
        }
    case 32768:
        v19 := GenS3{F0: (127 > 495), F1: 65535}
        _ = v19
    case -884:
        p0 = (!(-1 != int32(p2)))
    default:
        for fi20 := 0; fi20 < 5; fi20++ {
            v21 := GenS1{F0: uint16(p2), F1: 1}
            var v22 int32 = 2147483647
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    p2 += (p2 * 985)
    _ = p0
    _ = p1
    _ = p2
    return p0
}

func main() {
    v1 := []uint32{2147483648, 636, 1, 2}
    v2 := make(map[string]uint32)
    var v3 uint16 = 185
    v3 += (^(v3))
    for fi23 := 0; fi23 < 18; fi23++ {
        v24 := GenS4{F0: 1, F1: 0, F2: (fi23 - (fi23 % ((fi23) | 1)))}
        _ = fi23
        _ = v24
    }
    v25 := GenS5{F0: 1, F1: 533}
    var v26 string = "\n\t"
    v3 = 255
    for fi27 := 0; fi27 < 9; fi27++ {
        v28 := v2[("hello" + "\n\t")]
        _ = fi27
        _ = v28
    }
    v29 := GenS3{F0: true, F1: 65535}
    var v30 uint16 = v3
    for fi31 := 0; fi31 < 2; fi31++ {
        for fi32 := 0; fi32 < 13; fi32++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v30 += v3
            var v33 int32 = -1
            fmt.Println(v3)
            var v34 int = 462
            v35 := GenS4{F0: v33, F1: 204, F2: (fi32 - fi32)}
            _ = fi32
            _ = v33
            _ = v34
            _ = v35
        }
        _ = fi31
    }
    v36 := v2["\n\t"]
    if (true || (!((1e10 * (3.14159 - -2.25)) == ((-2.25 / 3.14159) - (1.5 / 3.14159))))) {
        v37 := GenS1{F0: (v3 << ((uint(v36)) & 63)), F1: -1}
        _ = v37
    }
    var v38 string = v26
    if (-676 > 968) {
        if len(v1) > 0 {
            v1[0] = uint32(v30)
        }
    } else {
        for fi39 := 0; fi39 < 1; fi39++ {
            var v40 uint8 = 255
            v41 := GenS1{F0: (v3 + (v30 / ((v30) | 1))), F1: -669}
            _ = fi39
            _ = v40
            _ = v41
        }
    }
    v42 := GenS3{F0: (!((725 <= 1000) && (false || (!(float64((1e10 / float32(float64(-2.25)))) != (float64(0.1) / 1.0)))))), F1: v36}
    v30 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v38
    _ = v42
}
