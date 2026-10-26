package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int8 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int16; F1 int32; F2 uint16 }
func genFn0(p0 uint32, p1 uint64) uint16 {
    for fi4 := 0; fi4 < 7; fi4++ {
        fmt.Println(fi4, fi4, p0)
        _ = fi4
    }
    if (51 > -1) {
        fmt.Println(p0, p1)
    }
    v5 := []int{}
    v6 := GenS1{F0: 569, F1: 4294967295, F2: 127}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 145
}
func genFn1(p0 uint8) uint16 {
    var v7 uint32 = 198
    for fi8 := 0; fi8 < 2; fi8++ {
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 uint8 = (p0 | (p0 ^ 1))
            v11 := GenS1{F0: 4294967295, F1: 907, F2: 0}
            switch fi8 {
            case -23:
                var v12 uint32 = v7
                _ = v12
            case -308:
                if (1.5 < float32(0.1)) {
                    v7 += (v7 | (v7 + (v7 * 2147483647)))
                    var v13 int = (fi9 * fi8)
                    var v14 int = (v13 * fi8)
                    _ = v13
                    _ = v14
                }
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        var v15 uint16 = 255
        for fi16 := 0; fi16 < 2; fi16++ {
            v17 := []int{fi16, fi8, (fi16 + 32768), (-(fi16))}
            v7 += (^(v7))
            v18 := GenS1{F0: 2, F1: 2, F2: 1}
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi8
        _ = v15
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        v7 += 4294967295
        fmt.Println(fi19)
        _ = fi19
    }
    _ = p0
    _ = v7
    return 1
}

func main() {
    v1 := []int{}
    var v2 string = (("a" + "\n\t") + "foo bar")
    v3 := []string{"foo bar", v2, ("\n\t" + v2), v2}
    v2 = ""
    v2 = (v2 + ("a" + v2))
    if true {
        switch ((v2 + (v2 + v2)) + v2) {
        case "":
            var v20 uint8 = 2
            _ = v20
        case "a":
            for fi21 := 0; fi21 < 19; fi21++ {
                switch fi21 {
                case 141:
                    v22 := GenS2{F0: fi21, F1: 954}
                    _ = v22
                case -2147483648:
                    v2 = v2
                case -748:
                    v23 := GenS1{F0: 4294967295, F1: 0, F2: 84}
                    _ = v23
                default:
                    var v24 bool = true
                    _ = v24
                }
                if false {
                    fmt.Println(v2, v2)
                } else {
                    v2 = v2
                }
                _ = fi21
            }
        case "foo bar":
            if false {
                for fi25 := 0; fi25 < 2; fi25++ {
                    for fi26 := 0; fi26 < 12; fi26++ {
                        if (1.5 <= 1e10) {
                            if (42 <= 9223372036854775807) {
                                switch ("hello" + "") {
                                case "\n\t":
                                    switch 272 {
                                    case -100:
                                        var v27 uint = 903
                                        _ = v27
                                    case -9223372036854775808:
                                        var v28 uint64 = 2
                                        _ = v28
                                    case 9223372036854775807:
                                        var v29 uint64 = 1
                                        if ((fi25 >= fi26) && ((^(v29)) == (v29 << ((255) & 63)))) {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            var v30 bool = false
                                            _ = v30
                                        } else {
                                            _ = len(v3)
                                            if false {
                                                fmt.Println(v29, fi26)
                                                v3 = append(v3, (v2 + ("" + ((v2 + (v2 + ((("\n\t" + "foo bar") + (v2 + v2)) + v2))) + v2))))
                                            }
                                        }
                                        _ = v29
                                    }
                                case "foo bar":
                                    if false {
                                        v2 = v2
                                        var v31 uint16 = 0
                                        _ = v31
                                    } else {
                                        v2 = v2
                                        fmt.Println(v2, fi25)
                                        v2 = (v2 + v2)
                                    }
                                case "hello":
                                    switch fi26 {
                                    case -9223372036854775808:
                                        if (!true) {
                                            v2 = ""
                                            v2 = v2
                                            _ = len(v3)
                                        }
                                    case 442:
                                        v2 = "foo bar"
                                    case 2147483647:
                                        v32 := GenS3{F0: 2, F1: 65535}
                                        _ = v32
                                    }
                                }
                                var v33 int16 = -1
                                switch v2 {
                                case "hello":
                                    fmt.Println(fi26, fi26)
                                }
                                _ = v33
                            }
                        } else {
                            var v34 int32 = 42
                            if true {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            } else {
                                switch fi25 {
                                case 32768:
                                    v2 = v2
                                }
                            }
                            var v35 bool = false
                            _ = v34
                            _ = v35
                        }
                        _ = fi26
                    }
                    v36 := GenS2{F0: (fi25 * fi25), F1: 255}
                    _ = fi25
                    _ = v36
                }
            } else {
                switch v2 {
                case "hello":
                    var v37 uint16 = 65535
                    _ = v37
                case "foo bar":
                    for fi38 := 0; fi38 < 9; fi38++ {
                        if (42 <= 149) {
                            var v39 int8 = -11
                            _ = v39
                        }
                        _ = fi38
                    }
                default:
                    v40 := GenS4{F0: -1, F1: -1, F2: 2}
                    _ = v40
                }
                var v41 int = 1
                _ = v41
            }
        }
    }
    switch v2 {
    case "hello":
        if true {
            var v42 uint16 = 2
            _ = v42
        }
    case "foo bar":
        v2 = v2
    default:
        switch 763 {
        case -1:
            v2 = "a"
        case 4294967295:
            if (-114 > -1) {
                switch -9223372036854775808 {
                case 815:
                    if true {
                        for fi43 := 0; fi43 < 15; fi43++ {
                            if (genFn0(2147483647, 1) < 65535) {
                                var v44 int16 = 42
                                v2 = ("foo bar" + "foo bar")
                                _ = v44
                            }
                            v2 = "\n\t"
                            _ = fi43
                        }
                    }
                default:
                    v2 = ((v2 + ("\n\t" + v2)) + v2)
                }
            } else {
                var v45 int64 = -856
                _ = v45
            }
        case -9223372036854775808:
            var v46 uint64 = 4294967295
            _ = v46
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    switch 32767 {
    case -9223372036854775808:
        var v47 int64 = 942
        _ = v47
    case 4294967295:
        for fi48 := 0; fi48 < 12; fi48++ {
            v2 = v2
            v2 = v2
            _ = fi48
        }
    default:
        v2 = "a"
    }
    var v49 uint16 = 255
    if true {
        var v50 int8 = -1
        v2 = "foo bar"
        _ = v50
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v49
}
