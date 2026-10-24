package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint64 }
type GenS2 struct { F0 int; F1 int16; F2 int8 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 int32; F1 uint32; F2 int32 }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 string; F1 uint8 }
type GenS7 struct { F0 uint16; F1 uint8 }
type GenS8 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS9 struct { F0 uint; F1 uint; F2 uint64 }
type GenS10 struct { F0 int64; F1 uint32; F2 uint16 }
func genFn0(p0 uint8, p1 int16) int32 {
    switch -438 {
    case 2147483647:
        v5 := []int{1, 4294967295}
        _ = v5
    }
    v6 := GenS1{F0: (255 != 0), F1: -987, F2: 528}
    for fi7 := 0; fi7 < 11; fi7++ {
        if ((1.5 - 0.1) < 3.14159) {
            var v8 int = (fi7 | fi7)
            v9 := []int{-1, v8, -880, (v8 * (fi7 & (-(fi7))))}
            _ = v8
            _ = v9
        } else {
            for fi10 := 0; fi10 < 20; fi10++ {
                switch "hello" {
                case "hello":
                    if true {
                        if (!false) {
                            var v11 string = ("a" + (("\n\t" + "foo bar") + "\n\t"))
                            var v12 uint32 = uint32(p0)
                            _ = v11
                            _ = v12
                        } else {
                            var v13 int = fi7
                            v14 := []int{(fi10 / (((fi7 & 4294967295)) | 1))}
                            _ = v13
                            _ = v14
                        }
                        v15 := GenS1{F0: false, F1: (fi7 << ((2) & 63)), F2: uint64(p0)}
                        _ = v15
                    } else {
                        v16 := GenS2{F0: (-(fi7)), F1: (p1 | 1), F2: -96}
                        var v17 uint16 = 1
                        _ = v16
                        _ = v17
                    }
                case "foo bar":
                    p0 = p0
                    v18 := []int{(fi10 - -181)}
                    var v19 float64 = 1e10
                    _ = v18
                    _ = v19
                }
                _ = fi10
            }
        }
        v20 := GenS1{F0: (42 > 1), F1: (fi7 | (-(fi7))), F2: 324}
        p1 = p1
        _ = fi7
        _ = v20
    }
    switch 4294967295 {
    case 9223372036854775807:
        v21 := GenS3{F0: 194, F1: 2147483647}
        _ = v21
    case 760:
        for fi22 := 0; fi22 < 10; fi22++ {
            p0 = 1
            _ = fi22
        }
    case 2147483647:
        v23 := GenS3{F0: p1, F1: 288}
        _ = v23
    }
    v24 := []int{}
    _ = p0
    _ = p1
    _ = v24
    _ = v6
    return 0
}

func main() {
    v1 := []string{"a", "\n\t", ("hello" + ""), "hello"}
    var v2 int64 = -1
    v3 := make(map[string]string)
    var v4 uint32 = 1
    v2 = (v2 % ((v2) | 1))
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    switch ("hello" + "a") {
    case "foo bar":
        fmt.Println(v4, v2)
    case "\n\t":
        if false {
            for fi25 := 0; fi25 < 15; fi25++ {
                if false {
                    if len(v1) > 0 {
                        v1[0] = ("foo bar" + "foo bar")
                    }
                } else {
                    v26 := GenS4{F0: 586, F1: v4, F2: -646}
                    v4 = (v4 | v4)
                    _ = v26
                }
                _ = fi25
            }
            for fi27 := 0; fi27 < 10; fi27++ {
                var v28 uint = 158
                _ = fi27
                _ = v28
            }
            v4 = v4
        }
    case "":
        switch 127 {
        case -2147483648:
            switch "hello" {
            case "hello":
                fmt.Println(v4, v2)
            case "a":
                switch 2147483647 {
                case 2147483647:
                    fmt.Println(v4)
                case 32768:
                    if false {
                        v29 := GenS5{F0: 137, F1: 1}
                        var v30 bool = true
                        _ = v29
                        _ = v30
                    } else {
                        var v31 int16 = 0
                        if (int32(v4) == -1) {
                            v4 -= v4
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if (false && (!(!false))) {
                                v32 := GenS6{F0: "hello", F1: 0}
                                _ = v32
                            }
                        } else {
                            _ = len(v1)
                            v4 = v4
                            v4 -= 1
                        }
                        fmt.Println(v4, v2, v4)
                        _ = v31
                    }
                }
            }
        case 128:
            var v33 int64 = (v2 - (^(v2)))
            _ = v33
        case -1:
            v34 := GenS4{F0: 211, F1: (v4 | (v4 / (((v4 | 1)) | 1))), F2: 2147483647}
            _ = v34
        default:
            switch -1 {
            case 353:
                if false {
                    var v35 float32 = 1e10
                    _ = v35
                }
            case 128:
                v4 = 42
            case 42:
                switch "hello" {
                case "":
                    for fi36 := 0; fi36 < 14; fi36++ {
                        if len(v1) > 0 {
                            v1[0] = ""
                        }
                        v37 := v3["hello"]
                        var v38 uint8 = 127
                        _ = fi36
                        _ = v37
                        _ = v38
                    }
                }
            }
        }
    default:
        for fi39 := 0; fi39 < 14; fi39++ {
            v2 = v2
            _ = fi39
        }
    }
    fmt.Println(v2, v2)
    fmt.Println(v2, v2)
    var v40 uint16 = 65535
    if true {
        if (true || false) {
            if true {
                switch "hello" {
                case "foo bar":
                    v41 := GenS7{F0: v40, F1: uint8(v4)}
                    _ = v41
                case "\n\t":
                    v40 += (v40 & 399)
                default:
                    if true {
                        fmt.Println(v40, v2)
                        switch 605 {
                        case 128:
                            delete(v3, "a")
                        case 42:
                            v42 := GenS8{F0: (v2 + v2), F1: 197, F2: v4}
                            _ = v42
                        }
                        v4 = v4
                    } else {
                        switch 884 {
                        case -254:
                            var v43 string = "hello"
                            _ = v43
                        case 0:
                            v44 := GenS1{F0: (32767 >= 927), F1: 127, F2: 42}
                            _ = v44
                        case -414:
                            var v45 uint64 = 42
                            _ = v45
                        default:
                            switch 32768 {
                            case -2147483648:
                                v4 = (v4 & (v4 % ((74) | 1)))
                            case -9223372036854775808:
                                v2 += 0
                            case -77:
                                if false {
                                    if true {
                                        v4 -= (v4 * v4)
                                        var v46 uint64 = 2
                                        delete(v3, "a")
                                        _ = v46
                                    }
                                    if false {
                                        for fi47 := 0; fi47 < 16; fi47++ {
                                            v2 -= v2
                                            _ = fi47
                                        }
                                    } else {
                                        v1 = append(v1, "\n\t")
                                        if (954 <= -1) {
                                            var v48 int8 = 127
                                            _ = v48
                                        }
                                    }
                                    if ((3.14159 / 0.1) >= 0.0) {
                                        if true {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                        switch -23 {
                                        case 2147483647:
                                            var v49 float32 = 0.1
                                            _ = v49
                                        }
                                    } else {
                                        delete(v3, "foo bar")
                                    }
                                } else {
                                    var v50 int = 0
                                    v4 = v4
                                    _ = v50
                                }
                            }
                        }
                    }
                }
                var v51 int16 = 1000
                fmt.Println(v51, v4)
                _ = v51
            }
            v52 := GenS9{F0: uint(v40), F1: 18446744073709551615, F2: 18446744073709551615}
            _ = v52
        } else {
            switch "\n\t" {
            case "hello":
                v53 := GenS9{F0: 397, F1: 2, F2: 717}
                _ = v53
            default:
                var v54 string = (("a" + "") + "foo bar")
                _ = v54
            }
            var v55 int16 = 1
            _ = v55
        }
    }
    v2 = v2
    var v56 int32 = 2147483647
    v57 := GenS10{F0: v2, F1: (^(v4)), F2: 32767}
    var v58 int64 = (^(v2))
    var v59 float64 = 0.0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v56
    _ = v57
    _ = v58
    _ = v59
}
