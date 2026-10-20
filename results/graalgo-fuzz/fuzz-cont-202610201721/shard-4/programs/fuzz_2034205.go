package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint32, p1 int, p2 int64) int32 {
    v7 := []int{4294967295, p1}
    p2 -= -703
    var v8 uint16 = 721
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 895
}

func main() {
    var v1 uint16 = 2
    var v2 string = ("\n\t" + ("" + "\n\t"))
    var v3 uint64 = 587
    v4 := []uint32{2147483647, 42}
    var v5 int32 = 386
    v6 := make(map[uint32]bool)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        if (int(v5) != fi9) {
            for fi10 := 0; fi10 < 13; fi10++ {
                v1 -= 934
                v2 = "\n\t"
                fmt.Println(fi9, v1)
                _ = fi10
            }
            v2 = "foo bar"
            v11 := GenS1{F0: 2, F1: v2}
            _ = v11
        } else {
            for fi12 := 0; fi12 < 12; fi12++ {
                v13 := GenS1{F0: 4294967295, F1: ""}
                v14 := GenS1{F0: 529, F1: (v2 + "")}
                v5 -= (^(v5))
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        _ = fi9
    }
    v2 = v2
    for fi15 := 0; fi15 < 7; fi15++ {
        if false {
            var v16 uint32 = 2
            if (18446744073709551615 == 65535) {
                var v17 uint16 = v1
                fmt.Println(v2)
                var v18 int = (fi15 | (fi15 - (fi15 & int(v16))))
                _ = v17
                _ = v18
            } else {
                var v19 uint32 = v16
                delete(v6, v16)
                _ = v19
            }
            _ = v16
        }
        _ = fi15
    }
    v20 := GenS1{F0: 4294967295, F1: "\n\t"}
    switch "a" {
    case "foo bar":
        for fi21 := 0; fi21 < 14; fi21++ {
            if (!true) {
                var v22 int16 = 42
                _ = v22
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v5 += int32(v1)
                if true {
                    if (!(!false)) {
                        if true {
                            switch fi21 {
                            case 32768:
                                if true {
                                    v4 = append(v4, 230)
                                } else {
                                    if true {
                                        var v23 int16 = int16(fi21)
                                        for fi24 := 0; fi24 < 12; fi24++ {
                                            v23 = (v23 ^ (v23 / (((v23 + v23)) | 1)))
                                            if (!(!((v23 & (v23 >> ((uint(v1)) & 63))) == (^(v23))))) {
                                                switch (fi21 | fi24) {
                                                case 128:
                                                    delete(v6, 2147483648)
                                                case -570:
                                                    var v25 uint16 = (^(v1))
                                                    _ = v25
                                                case 4294967295:
                                                    if false {
                                                        var v26 uint = 594
                                                        var v27 int32 = genFn0(2147483648, fi21, -1)
                                                        var v28 float64 = (0.0 * 3.14159)
                                                        _ = v26
                                                        _ = v27
                                                        _ = v28
                                                    }
                                                default:
                                                    switch v2 {
                                                    case "hello":
                                                        _ = len(v6)
                                                    default:
                                                        v3 = (v3 % (((v3 << ((2) & 63))) | 1))
                                                    }
                                                }
                                                v5 = (v5 * v5)
                                                v1 += (v1 * (v1 >> ((29) & 63)))
                                            } else {
                                                v5 -= (v5 >> ((65535) & 63))
                                            }
                                            _ = fi24
                                        }
                                        if (42 >= -39) {
                                            v2 = v2
                                        } else {
                                            switch fi21 {
                                            case -465:
                                                v1 -= v1
                                            case 128:
                                                v29 := GenS2{F0: 42}
                                                _ = v29
                                            case 9223372036854775807:
                                                v5 = (v5 >> ((42) & 63))
                                            default:
                                                if true {
                                                    if false {
                                                        if false {
                                                            var v30 uint = 582
                                                            v31 := GenS1{F0: 530, F1: "a"}
                                                            v32 := GenS3{F0: true}
                                                            _ = v30
                                                            _ = v31
                                                            _ = v32
                                                        }
                                                        if true {
                                                            if false {
                                                                v5 += int32(v1)
                                                            } else {
                                                                if true {
                                                                    var v33 uint8 = 80
                                                                    v34 := GenS4{F0: (true && (!false)), F1: 4294967295, F2: 2147483647}
                                                                    var v35 uint16 = v1
                                                                    _ = v33
                                                                    _ = v34
                                                                    _ = v35
                                                                } else {
                                                                    if true {
                                                                        fmt.Println(v1, v3, v5)
                                                                        if (42 > 172) {
                                                                            v36 := GenS1{F0: 2147483647, F1: "\n\t"}
                                                                            _ = v36
                                                                        } else {
                                                                            v37 := GenS5{F0: -1, F1: 42}
                                                                            var v38 int32 = genFn0(886, (fi21 >> ((18446744073709551615) & 63)), 0)
                                                                            _ = v37
                                                                            _ = v38
                                                                        }
                                                                        if true {
                                                                            v1 = uint16(v3)
                                                                            var v39 int = (fi21 ^ (fi21 >> ((62) & 63)))
                                                                            var v40 uint8 = 42
                                                                            _ = v39
                                                                            _ = v40
                                                                        } else {
                                                                            v5 += -1
                                                                            v6[2] = ((v23 - (v23 | v23)) < 1)
                                                                            var v41 uint32 = 0
                                                                            _ = v41
                                                                        }
                                                                    } else {
                                                                        delete(v6, 2147483647)
                                                                    }
                                                                    fmt.Println(v1)
                                                                }
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    var v42 int16 = v23
                                                    _ = v42
                                                }
                                            }
                                            if true {
                                                for fi43 := 0; fi43 < 7; fi43++ {
                                                    v6[65535] = false
                                                    var v44 int16 = v23
                                                    switch -858 {
                                                    case -444:
                                                        v45 := GenS6{F0: int16(v5), F1: v1}
                                                        _ = v45
                                                    case -9223372036854775808:
                                                        var v46 int32 = int32(fi21)
                                                        _ = v46
                                                    default:
                                                        v47 := GenS3{F0: false}
                                                        _ = v47
                                                    }
                                                    _ = fi43
                                                    _ = v44
                                                }
                                                var v48 uint8 = 123
                                                v49 := GenS2{F0: 18446744073709551615}
                                                _ = v48
                                                _ = v49
                                            }
                                        }
                                        _ = v23
                                    } else {
                                        var v50 uint8 = 128
                                        _ = v50
                                    }
                                    fmt.Println(v3, fi21, v1)
                                }
                            case 9223372036854775807:
                                var v51 uint8 = 42
                                _ = v51
                            }
                        }
                        v3 = (v3 & v3)
                    }
                }
            }
            var v52 bool = ((v3 ^ 2) < 42)
            _ = fi21
            _ = v52
        }
    case "a":
        v53 := GenS3{F0: true}
        _ = v53
    case "":
        v54 := GenS4{F0: (-2.25 == 0.0), F1: v3, F2: 0}
        _ = v54
    default:
        fmt.Println(v5, v5)
    }
    var v55 string = "a"
    v2 = (("a" + "a") + v55)
    for fi56 := 0; fi56 < 11; fi56++ {
        var v57 int8 = -18
        if false {
            var v58 uint32 = 346
            _ = v58
        } else {
            var v59 uint64 = (v3 | v3)
            var v60 uint64 = (v59 ^ v3)
            if (v1 != 32767) {
                var v61 float32 = -2.25
                _ = v61
            }
            _ = v59
            _ = v60
        }
        _ = fi56
        _ = v57
    }
    for fi62 := 0; fi62 < 17; fi62++ {
        var v63 uint64 = 1
        fmt.Println(v55, v55, fi62)
        _ = fi62
        _ = v63
    }
    fmt.Println(v5, v3, v5)
    fmt.Println(v2, v5, v55)
    v5 = v5
    for fi64 := 0; fi64 < 16; fi64++ {
        fmt.Println(v55, v1, v55)
        v6[65535] = true
        v65 := v6[4294967295]
        _ = fi64
        _ = v65
    }
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
}
