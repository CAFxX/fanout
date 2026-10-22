package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 uint64; F2 string }
type GenS5 struct { F0 int64; F1 uint; F2 int }
func genFn0(p0 uint64, p1 int32, p2 uint8) int16 {
    for fi6 := 0; fi6 < 13; fi6++ {
        switch fi6 {
        case 42:
            v7 := []int{(fi6 >> ((0) & 63))}
            _ = v7
        default:
            p1 -= -1
        }
        _ = fi6
    }
    p0 = p0
    for fi8 := 0; fi8 < 8; fi8++ {
        switch "\n\t" {
        case "hello":
            var v9 uint32 = uint32(p2)
            _ = v9
        case "\n\t":
            p2 = p2
        }
        for fi10 := 0; fi10 < 20; fi10++ {
            switch "hello" {
            case "foo bar":
                if ((p1 / ((p1) | 1)) >= (p1 + -2147483648)) {
                    var v11 uint16 = 32767
                    v12 := []int{fi10, 1, (^(fi10))}
                    _ = v11
                    _ = v12
                } else {
                    v13 := []int{2147483647, -778, 0, 9223372036854775807}
                    var v14 bool = (true || true)
                    p0 = (p0 | (p0 >> ((0) & 63)))
                    _ = v13
                    _ = v14
                }
            case "":
                p0 = p0
            }
            var v15 int = (fi8 * (fi10 ^ fi10))
            fmt.Println(fi8, p1)
            _ = fi10
            _ = v15
        }
        if (!false) {
            fmt.Println(p0, p2)
            if true {
                v16 := []int{fi8, 2147483647, (fi8 ^ (-(fi8))), 9223372036854775807}
                fmt.Println(fi8)
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                _ = v16
            } else {
                p2 = p2
                v17 := GenS1{F0: 127, F1: (p0 & p0), F2: int16(p2)}
                _ = v17
            }
        }
        _ = fi8
    }
    var v18 int64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return 141
}

func main() {
    var v1 int8 = 42
    var v2 int = -2147483648
    v3 := []uint32{2147483647}
    v4 := make(map[uint32]uint32)
    var v5 int16 = 42
    v1 += v1
    v19 := GenS2{F0: 745}
    var v20 string = "\n\t"
    _ = len(v3)
    fmt.Println(v2, v1, v2)
    fmt.Println(v2)
    if (false || true) {
        v21 := GenS1{F0: v1, F1: 42, F2: v5}
        for fi22 := 0; fi22 < 6; fi22++ {
            for fi23 := 0; fi23 < 20; fi23++ {
                var v24 int16 = (v5 * 1000)
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
        for fi25 := 0; fi25 < 10; fi25++ {
            switch "" {
            case "":
                for fi26 := 0; fi26 < 15; fi26++ {
                    var v27 uint16 = 32767
                    _ = fi26
                    _ = v27
                }
            default:
                v28 := GenS3{F0: 1}
                _ = v28
            }
            _ = fi25
        }
        _ = v21
    }
    var v29 int64 = -1
    v5 += genFn0(18446744073709551615, -2147483648, 1)
    var v30 uint8 = 127
    var v31 uint32 = 607
    v20 = v20
    var v32 bool = true
    v33 := GenS4{F0: 0, F1: 986, F2: "a"}
    if (v30 == v30) {
        if v32 {
            v2 = 128
            var v34 uint = 600
            fmt.Println(v20, v20)
            _ = v34
        } else {
            var v35 uint8 = v30
            _ = len(v3)
            _ = v35
        }
        fmt.Println(v30, v29)
    } else {
        fmt.Println(v20, v1)
        var v36 int64 = v29
        _ = v36
    }
    fmt.Println(v5, v32, v30)
    v30 = v30
    if v32 {
        delete(v4, v31)
    } else {
        switch (v2 / ((9223372036854775807) | 1)) {
        case 9223372036854775807:
            var v37 int16 = (v5 * (v5 * (v5 << ((133) & 63))))
            _ = v37
        case 32767:
            switch "foo bar" {
            case "hello":
                var v38 int = 0
                _ = v38
            }
        case 32768:
            for fi39 := 0; fi39 < 18; fi39++ {
                if v32 {
                    v40 := GenS1{F0: v1, F1: 42, F2: (^(v5))}
                    _ = v40
                }
                if (42 > 2147483647) {
                    var v41 uint32 = (v31 * (v31 >> ((82) & 63)))
                    _ = v41
                } else {
                    if false {
                        if ((!false) || (3.14159 <= ((0.0 * ((1e10 * (float64((1.5 - 0.1)) + (3.14159 * (0.0 - (float64(1.5) + (0.1 - 1.5)))))) * float64(-2.25))) * float64(-2.25)))) {
                            var v42 int8 = int8(v2)
                            if false {
                                v43 := GenS5{F0: (v29 ^ 9223372036854775807), F1: 100, F2: v2}
                                fmt.Println(v31, v2)
                                _ = v43
                            } else {
                                var v44 int8 = 0
                                if (v32 && (388 < 255)) {
                                    v20 = (v20 + (v20 + "hello"))
                                }
                                _ = v44
                            }
                            var v45 int16 = genFn0(1, -672, 143)
                            _ = v42
                            _ = v45
                        } else {
                            var v46 int64 = (v29 | -506)
                            if v32 {
                                switch "a" {
                                case "a":
                                    v31 = v31
                                }
                                for fi47 := 0; fi47 < 1; fi47++ {
                                    v20 = v20
                                    _ = fi47
                                }
                            } else {
                                v5 -= v5
                                if v32 {
                                    v48 := GenS2{F0: 42}
                                    if true {
                                        for fi49 := 0; fi49 < 16; fi49++ {
                                            v50 := v4[(v31 & (^(v31)))]
                                            _ = fi49
                                            _ = v50
                                        }
                                        v51 := v4[(v31 - (^(v31)))]
                                        _ = v51
                                    }
                                    var v52 uint64 = 1
                                    _ = v48
                                    _ = v52
                                } else {
                                    if v32 {
                                        if v32 {
                                            switch v20 {
                                            case "\n\t":
                                                var v53 uint32 = v31
                                                _ = v53
                                            case "hello":
                                                v46 = v29
                                            case "":
                                                var v54 uint = 255
                                                _ = v54
                                            }
                                        }
                                    }
                                }
                                var v55 bool = (!(!false))
                                _ = v55
                            }
                            v46 = (v29 << ((0) & 63))
                            _ = v46
                        }
                    } else {
                        if len(v3) > 0 {
                            v3[0] = (v31 ^ v31)
                        }
                    }
                    fmt.Println(v1, v1, v31)
                    if (1 <= 2) {
                        var v56 int64 = (v29 & -159)
                        for fi57 := 0; fi57 < 3; fi57++ {
                            switch v20 {
                            case "hello":
                                fmt.Println(v2, v5, v20)
                            case "foo bar":
                                fmt.Println(fi57)
                            default:
                                if (!false) {
                                    switch (((v20 + (v20 + "hello")) + (v20 + (("\n\t" + ("a" + v20)) + v20))) + "hello") {
                                    case "foo bar":
                                        v32 = (-2147483648 >= 42)
                                    case "hello":
                                        switch v20 {
                                        case "foo bar":
                                            v29 = 42
                                        case "a":
                                            fmt.Println(v20, v20, v32)
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v58 := v4[(v31 + 720)]
                                            _ = v58
                                        default:
                                            v30 = v30
                                        }
                                    case "\n\t":
                                        v5 += (v5 * (v5 * 1000))
                                    }
                                    var v59 int64 = v56
                                    _ = v59
                                }
                            }
                            v32 = v32
                            _ = fi57
                        }
                        fmt.Println(v2, v2, v31)
                        _ = v56
                    } else {
                        v1 = v1
                    }
                }
                _ = fi39
            }
        }
        v4[v31] = 65535
        for fi60 := 0; fi60 < 14; fi60++ {
            switch (fi60 * (fi60 % ((int(v1)) | 1))) {
            case 42:
                v29 = (v29 ^ 9223372036854775807)
            case 617:
                for fi61 := 0; fi61 < 1; fi61++ {
                    fmt.Println(fi61, v2, v32)
                    v5 -= 1000
                    var v62 uint8 = 127
                    _ = fi61
                    _ = v62
                }
            case 127:
                v2 -= fi60
            default:
                v63 := GenS2{F0: 2}
                _ = v63
            }
            if (((!true) || (v1 != (v1 % ((v1) | 1)))) || v32) {
                for fi64 := 0; fi64 < 20; fi64++ {
                    v65 := v4[(^(v31))]
                    v66 := GenS2{F0: 0}
                    _ = fi64
                    _ = v65
                    _ = v66
                }
            } else {
                v2 = 2147483647
            }
            delete(v4, v31)
            _ = fi60
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
