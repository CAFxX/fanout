package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 int8 }
func genFn0(p0 uint32, p1 int16, p2 uint64) uint {
    var v4 int64 = 0
    for fi5 := 0; fi5 < 12; fi5++ {
        for fi6 := 0; fi6 < 9; fi6++ {
            v7 := GenS1{F0: (false && ((-(p1)) >= (p1 | p1))), F1: p2}
            var v8 bool = true
            var v9 uint8 = 127
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        if (false || false) {
            for fi10 := 0; fi10 < 1; fi10++ {
                switch fi10 {
                case -2147483648:
                    p1 = p1
                case 32768:
                    v11 := GenS1{F0: (!(false || true)), F1: 997}
                    _ = v11
                case 2147483647:
                    var v12 int = fi5
                    var v13 int32 = 42
                    var v14 uint = 744
                    _ = v12
                    _ = v13
                    _ = v14
                }
                v15 := []int{(fi5 + (fi10 << ((65535) & 63))), 4294967295, (^(fi5)), (-(fi10))}
                p2 -= (p2 * (p2 ^ p2))
                _ = fi10
                _ = v15
            }
            v16 := GenS1{F0: (!(!(!(18446744073709551615 >= (p2 % (((p2 ^ (p2 | 2))) | 1)))))), F1: p2}
            for fi17 := 0; fi17 < 19; fi17++ {
                var v18 int = fi17
                var v19 string = "foo bar"
                v20 := []int{fi17, (v18 + v18), 32767, (v18 >> ((2) & 63))}
                switch "a" {
                case "a":
                    var v21 uint16 = 42
                    _ = v21
                case "":
                    v22 := GenS2{F0: 197}
                    _ = v22
                case "foo bar":
                    v23 := GenS2{F0: 211}
                    _ = v23
                }
                if (!(-2.25 <= 1.5)) {
                    switch "\n\t" {
                    case "\n\t":
                        v20 = append(v20, v18)
                    default:
                        fmt.Println(fi5, v19)
                    }
                    v24 := GenS1{F0: (true || (v18 > 4294967295)), F1: p2}
                    if true {
                        var v25 uint32 = 65535
                        var v26 uint32 = 2
                        _ = v25
                        _ = v26
                    }
                    _ = v24
                } else {
                    switch (v18 & fi5) {
                    case 127:
                        if (((false && (3.14159 < 1.5)) || (!true)) && (fi5 < -745)) {
                            var v27 uint32 = (^(p0))
                            _ = v27
                        }
                    case 32768:
                        if len(v20) >= 2 {
                            v20 = v20[1:len(v20)]
                        }
                    case -2147483648:
                        var v28 float64 = 1e10
                        _ = v28
                    default:
                        switch "hello" {
                        case "\n\t":
                            v20 = append(v20, fi5)
                        case "":
                            v18 -= -110
                        case "a":
                            v29 := GenS1{F0: (!true), F1: p2}
                            fmt.Println(p0, p0)
                            var v30 uint8 = 243
                            fmt.Println(fi5)
                            var v31 uint32 = 1
                            if ((^(v30)) == 127) {
                                var v32 uint16 = 1
                                _ = v32
                            } else {
                                fmt.Println(p0, v18, v19)
                                var v33 uint32 = (v31 ^ v31)
                                var v34 int16 = (p1 + 0)
                                _ = v33
                                _ = v34
                            }
                            v31 = uint32(p2)
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                    }
                    v35 := GenS2{F0: 127}
                    v36 := GenS3{F0: 128, F1: 54}
                    _ = v35
                    _ = v36
                }
                switch v18 {
                case 2147483647:
                    p1 = p1
                case 153:
                    var v37 uint8 = 42
                    _ = v37
                case 42:
                    v38 := GenS1{F0: true, F1: (p2 << ((4294967295) & 63))}
                    _ = v38
                default:
                    v4 = (v4 << ((4294967295) & 63))
                }
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
            _ = v16
        }
        _ = fi5
    }
    v39 := []int{4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v39
    _ = v4
    return 18446744073709551615
}
func genFn1(p0 uint, p1 int8, p2 int8) int32 {
    var v40 int16 = 1000
    var v41 uint32 = 2
    v40 -= (v40 % ((v40) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v40
    _ = v41
    return 343
}

func main() {
    v1 := make(map[string]string)
    var v2 float32 = 0.1
    v3 := []int{497}
    v2 = 1.5
    v2 = (1e10 - v2)
    var v42 uint8 = 128
    v2 = v2
    v2 -= (((v2 / 0.1) - (0.1 * 3.14159)) + (v2 - ((((v2 + v2) - (v2 / float32(0.1))) / 1e10) - v2)))
    v43 := GenS2{F0: (^(v42))}
    if true {
        switch 9223372036854775807 {
        case 127:
            var v44 int16 = -32768
            _ = v44
        }
    }
    v2 = v2
    v42 += (v42 % ((1) | 1))
    if (843 == 876) {
        fmt.Println(v42)
        v45 := GenS1{F0: (354 == uint64(v42)), F1: 18446744073709551615}
        _ = v45
    } else {
        v2 += float32((3.14159 + 3.14159))
    }
    for fi46 := 0; fi46 < 9; fi46++ {
        var v47 string = (("a" + "foo bar") + "")
        _ = fi46
        _ = v47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v43
}
