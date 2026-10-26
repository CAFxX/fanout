package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int64; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 int64; F2 uint }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int8; F1 uint16; F2 int64 }
type GenS8 struct { F0 uint16; F1 int }
func genFn0(p0 int16) string {
    p0 += p0
    v7 := GenS1{F0: 182, F1: 2, F2: 128}
    v8 := GenS2{F0: 127, F1: 42}
    _ = p0
    _ = v7
    _ = v8
    return "\n\t"
}
func genFn1(p0 uint) uint64 {
    for fi9 := 0; fi9 < 11; fi9++ {
        p0 += p0
        _ = fi9
    }
    v10 := GenS2{F0: -860, F1: 0}
    v11 := []int{}
    _ = p0
    _ = v10
    _ = v11
    return 0
}

func main() {
    var v1 uint64 = 517
    var v2 uint = 42
    var v3 int64 = 0
    var v4 int32 = -2147483648
    v5 := []string{"a", (("foo bar" + "foo bar") + "hello")}
    var v6 uint32 = 667
    _ = len(v5)
    if true {
        switch 4294967295 {
        case 118:
            var v12 uint = (v2 >> (((v2 + 888)) & 63))
            _ = v12
        case 976:
            if true {
                if len(v5) > 0 {
                    v5[0] = ("" + "a")
                }
                v13 := GenS2{F0: 357, F1: v1}
                fmt.Println(v4)
                _ = v13
            }
        case 128:
            var v14 uint8 = 127
            _ = v14
        }
    } else {
        if ((false && (v6 < (v6 % (((^(v6))) | 1)))) && false) {
            v15 := GenS2{F0: 32768, F1: genFn1((v2 + 2))}
            var v16 uint32 = 0
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v15
            _ = v16
        }
        _ = len(v5)
    }
    v17 := GenS2{F0: 0, F1: (^(v1))}
    var v18 uint = (v2 % (((v2 % ((uint(v6)) | 1))) | 1))
    fmt.Println(v3, v18, v2)
    fmt.Println(v1)
    for fi19 := 0; fi19 < 1; fi19++ {
        switch "\n\t" {
        case "hello":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        _ = len(v5)
        if len(v5) > 0 {
            v5[0] = ("hello" + "hello")
        }
        _ = fi19
    }
    var v20 float32 = ((1e10 / -2.25) + 0.0)
    switch -9223372036854775808 {
    case 1:
        v21 := GenS3{F0: 0, F1: genFn1(v18), F2: 1}
        _ = v21
    case -725:
        if (!(true && false)) {
            for fi22 := 0; fi22 < 3; fi22++ {
                var v23 uint = v2
                var v24 uint64 = v1
                v6 -= v6
                _ = fi22
                _ = v23
                _ = v24
            }
        } else {
            var v25 uint32 = 65535
            switch -2147483648 {
            case 4294967295:
                if false {
                    if len(v5) > 0 {
                        v5[0] = ""
                    }
                    v26 := GenS4{F0: 42}
                    if false {
                        v1 += 18446744073709551615
                    }
                    _ = v26
                } else {
                    if len(v5) > 0 {
                        v5[0] = ("a" + ("\n\t" + "a"))
                    }
                }
            case 2147483647:
                var v27 uint16 = 255
                _ = v27
            default:
                switch -867 {
                case 941:
                    for fi28 := 0; fi28 < 19; fi28++ {
                        var v29 int = (fi28 % ((fi28) | 1))
                        var v30 bool = false
                        if v30 {
                            fmt.Println(v2, v18)
                        }
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                case -108:
                    fmt.Println(v3, v25)
                default:
                    if false {
                        switch (("a" + ("hello" + "a")) + "\n\t") {
                        case "foo bar":
                            v18 = (v18 + v18)
                        case "a":
                            for fi31 := 0; fi31 < 19; fi31++ {
                                v2 -= (^(v18))
                                if false {
                                    var v32 int16 = -287
                                    _ = v32
                                } else {
                                    var v33 uint64 = v1
                                    _ = v33
                                }
                                _ = fi31
                            }
                        case "\n\t":
                            v34 := GenS1{F0: 554, F1: 255, F2: 192}
                            _ = v34
                        }
                        var v35 string = ((((("" + "") + "hello") + "\n\t") + "hello") + "")
                        _ = v35
                    } else {
                        var v36 int64 = 42
                        v37 := GenS5{F0: v2, F1: (v3 >> ((1) & 63)), F2: (v18 % (((^(v2))) | 1))}
                        _ = v36
                        _ = v37
                    }
                }
            }
            v6 += 2147483647
            _ = v25
        }
    case 2147483647:
        if (-1 != 391) {
            fmt.Println(v18, v2)
            if true {
                switch 9223372036854775807 {
                case 2147483647:
                    var v38 uint16 = 1
                    _ = v38
                case 1:
                    var v39 int64 = v3
                    _ = v39
                default:
                    for fi40 := 0; fi40 < 19; fi40++ {
                        var v41 int16 = -140
                        _ = fi40
                        _ = v41
                    }
                }
            }
            v5 = append(v5, ("\n\t" + ""))
        } else {
            var v42 uint32 = 0
            if (!false) {
                if (-86 > -400) {
                    if (0 > 144) {
                        if (genFn1((^(v18))) >= (v1 ^ v1)) {
                            v2 = (^(v18))
                        } else {
                            for fi43 := 0; fi43 < 16; fi43++ {
                                if true {
                                    var v44 float64 = 0.1
                                    _ = v44
                                }
                                if (true || true) {
                                    if true {
                                        var v45 float64 = 0.1
                                        v46 := GenS6{F0: 32767}
                                        _ = v45
                                        _ = v46
                                    }
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                                if len(v5) > 0 {
                                    v5[0] = "\n\t"
                                }
                                _ = fi43
                            }
                            v2 = v2
                            v47 := GenS4{F0: (v1 & genFn1(v18))}
                            _ = v47
                        }
                    } else {
                        switch 1 {
                        case -774:
                            if len(v5) > 0 {
                                v5[0] = "hello"
                            }
                        case 473:
                            if (false || false) {
                                var v48 float64 = -2.25
                                _ = v48
                            } else {
                                fmt.Println(v18)
                                var v49 float64 = (1.5 * 1.5)
                                v3 += v3
                                _ = v49
                            }
                        case -9223372036854775808:
                            v50 := GenS7{F0: -1, F1: 65535, F2: (v3 << ((v18) & 63))}
                            _ = v50
                        default:
                            for fi51 := 0; fi51 < 4; fi51++ {
                                for fi52 := 0; fi52 < 2; fi52++ {
                                    if len(v5) > 0 {
                                        v5[0] = "a"
                                    }
                                    v4 = 298
                                    _ = fi52
                                }
                                var v53 uint8 = 127
                                if (86 < genFn1(v2)) {
                                    var v54 uint32 = 65535
                                    _ = v54
                                } else {
                                    v55 := GenS8{F0: 924, F1: 2147483647}
                                    _ = v55
                                }
                                _ = fi51
                                _ = v53
                            }
                        }
                        v56 := GenS7{F0: -124, F1: 255, F2: 0}
                        _ = v56
                    }
                }
            } else {
                v4 -= (v4 - v4)
                v4 -= v4
            }
            if len(v5) > 0 {
                v5[0] = ("hello" + "a")
            }
            _ = v42
        }
    }
    v2 -= v2
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
