package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 bool }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int; F1 uint32; F2 uint32 }
func genFn0(p0 uint32) uint16 {
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := []int{(-(fi7)), fi7, (fi7 % ((fi7) | 1))}
        for fi9 := 0; fi9 < 17; fi9++ {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            if len(v8) > 0 {
                v8[0] = -487
            }
            if len(v8) > 0 {
                v8[0] = fi9
            }
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    switch 127 {
    case 32768:
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 int = fi10
            v12 := []int{-458, -561, 32768, v11}
            v13 := GenS1{F0: 0, F1: 42, F2: (!(((127 > 2) && (65535 >= 255)) || true))}
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
    }
    p0 = 0
    fmt.Println(p0)
    var v14 bool = (!(42 < 752))
    _ = p0
    _ = v14
    return 878
}
func genFn1(p0 int8, p1 uint) uint {
    if true {
        switch -757 {
        case -418:
            var v15 int = 2147483647
            _ = v15
        }
    }
    if (0 <= 581) {
        v16 := []int{74, -2147483648, 32767}
        p0 -= p0
        v17 := GenS1{F0: 0, F1: p0, F2: (352 >= 604)}
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        if (!(-9223372036854775808 == 42)) {
            if false {
                p1 = (p1 << ((p1) & 63))
                if ((fi18 >= fi18) && true) {
                    switch fi18 {
                    case 4294967295:
                        for fi19 := 0; fi19 < 12; fi19++ {
                            v20 := GenS2{F0: (p1 & p1), F1: 128}
                            if true {
                                p0 = p0
                            } else {
                                p0 = (p0 << ((p1) & 63))
                                fmt.Println(fi18)
                            }
                            _ = fi19
                            _ = v20
                        }
                    case 0:
                        if (true || false) {
                            var v21 int = fi18
                            switch fi18 {
                            case -2147483648:
                                p1 += p1
                            default:
                                v22 := []int{}
                                _ = v22
                            }
                            var v23 int16 = 734
                            _ = v21
                            _ = v23
                        } else {
                            if false {
                                if false {
                                    v24 := GenS2{F0: (p1 << (((p1 >> (((p1 ^ p1)) & 63))) & 63)), F1: 127}
                                    switch (fi18 & 4294967295) {
                                    case 32767:
                                        if true {
                                            var v25 int8 = p0
                                            v26 := GenS3{F0: 844, F1: 2}
                                            _ = v25
                                            _ = v26
                                        } else {
                                            v27 := GenS3{F0: 0, F1: 1}
                                            var v28 uint32 = 42
                                            _ = v27
                                            _ = v28
                                        }
                                    default:
                                        switch "foo bar" {
                                        case "a":
                                            var v29 string = ""
                                            _ = v29
                                        case "\n\t":
                                            p0 = p0
                                        }
                                    }
                                    v30 := GenS2{F0: (^(p1)), F1: 128}
                                    _ = v24
                                    _ = v30
                                } else {
                                    var v31 uint16 = 65535
                                    _ = v31
                                }
                                var v32 uint64 = 352
                                var v33 bool = ((p0 ^ p0) > (p0 % ((int8(p1)) | 1)))
                                _ = v32
                                _ = v33
                            } else {
                                for fi34 := 0; fi34 < 13; fi34++ {
                                    if false {
                                        p0 = p0
                                        p1 += p1
                                        if false {
                                            var v35 bool = (-9223372036854775808 == -662)
                                            var v36 string = "\n\t"
                                            _ = v35
                                            _ = v36
                                        }
                                        p0 = p0
                                        p1 = (p1 >> (((p1 / (((^(p1))) | 1))) & 63))
                                        var v37 int16 = 32767
                                        var v38 float64 = 3.14159
                                        v37 -= (v37 >> ((42) & 63))
                                        _ = v37
                                        _ = v38
                                    }
                                    _ = fi34
                                }
                            }
                        }
                    default:
                        v39 := []int{-1, fi18, fi18}
                        _ = v39
                    }
                    v40 := GenS3{F0: 255, F1: 0}
                    _ = v40
                } else {
                    v41 := []int{32768, fi18, (fi18 << ((p1) & 63))}
                    _ = v41
                }
                var v42 bool = true
                _ = v42
            }
        } else {
            v43 := GenS4{F0: "a"}
            _ = v43
        }
        var v44 uint16 = 31
        _ = fi18
        _ = v44
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 string = "foo bar"
    var v2 uint8 = 128
    v3 := []string{v1, (v1 + "hello")}
    v4 := []int{9223372036854775807, 32767, 128, 107}
    var v5 uint64 = 297
    var v6 int = 393
    var v45 string = v1
    switch v45 {
    case "":
        for fi46 := 0; fi46 < 15; fi46++ {
            _ = len(v4)
            switch ("\n\t" + "foo bar") {
            case "":
                v1 = v45
            case "a":
                v47 := GenS2{F0: 0, F1: (^(v2))}
                _ = v47
            case "\n\t":
                switch "hello" {
                case "\n\t":
                    v6 = v6
                }
            }
            for fi48 := 0; fi48 < 5; fi48++ {
                v4 = append(v4, fi46)
                v45 = "foo bar"
                _ = fi48
            }
            _ = fi46
        }
    }
    var v49 int32 = -594
    _ = len(v3)
    if (false && false) {
        var v50 int16 = 1000
        if false {
            switch v45 {
            case "foo bar":
                v4 = append(v4, v6)
            case "hello":
                v1 = (v1 + ("a" + v45))
            default:
                var v51 float32 = (0.1 + 1e10)
                _ = v51
            }
            if true {
                if ((-1 <= 9223372036854775807) && false) {
                    var v52 string = ("foo bar" + v1)
                    _ = v52
                } else {
                    fmt.Println(v49, v6)
                }
            }
            if (3.14159 != 3.14159) {
                v50 += v50
                switch ("foo bar" + v45) {
                case "a":
                    var v53 uint = 255
                    _ = v53
                }
            } else {
                var v54 int8 = 100
                v55 := GenS5{F0: -1}
                switch v6 {
                case -9223372036854775808:
                    var v56 uint32 = 42
                    _ = v56
                case 2147483647:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case -513:
                    var v57 uint = genFn1(-128, 0)
                    _ = v57
                default:
                    v6 = (-(v6))
                }
                _ = v54
                _ = v55
            }
        }
        if true {
            v5 = v5
            v58 := GenS6{F0: v6, F1: 1000, F2: 944}
            _ = v58
        }
        _ = v50
    } else {
        if false {
            fmt.Println(v49)
        }
    }
    var v59 string = v45
    v1 = (("\n\t" + "") + ("\n\t" + ((("\n\t" + ("" + "hello")) + v1) + "hello")))
    v49 = v49
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v49
    _ = v5
    _ = v59
    _ = v6
}
