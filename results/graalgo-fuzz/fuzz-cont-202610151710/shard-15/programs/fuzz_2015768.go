package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int32; F1 string; F2 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 int16; F2 uint16 }
type GenS5 struct { F0 int; F1 bool; F2 uint32 }
type GenS6 struct { F0 uint; F1 int32; F2 uint32 }
type GenS7 struct { F0 int64; F1 uint; F2 uint64 }
func genFn0(p0 int) uint16 {
    fmt.Println(p0)
    switch p0 {
    case 128:
        if (18446744073709551615 >= 211) {
            if false {
                for fi4 := 0; fi4 < 6; fi4++ {
                    v5 := []int{fi4}
                    p0 += -9223372036854775808
                    _ = fi4
                    _ = v5
                }
                fmt.Println(p0)
            }
        } else {
            var v6 uint8 = 1
            _ = v6
        }
    case 42:
        if false {
            switch p0 {
            case -259:
                var v7 uint16 = 651
                _ = v7
            case 32768:
                var v8 int16 = -651
                _ = v8
            default:
                p0 = 127
            }
            var v9 int64 = 9223372036854775807
            _ = v9
        }
    case 9223372036854775807:
        v10 := []int{p0, -469, p0, p0}
        _ = v10
    }
    p0 += (p0 << ((65535) & 63))
    v11 := GenS1{F0: true, F1: 792}
    _ = p0
    _ = v11
    return 2
}

func main() {
    v1 := []string{"", "", "\n\t"}
    var v2 uint64 = 537
    v3 := make(map[string]string)
    if true {
        if (!("a" >= (("foo bar" + "") + ""))) {
            switch "a" {
            case "hello":
                for fi12 := 0; fi12 < 2; fi12++ {
                    if (v2 <= v2) {
                        if false {
                            fmt.Println(v2, v2)
                            for fi13 := 0; fi13 < 7; fi13++ {
                                if (42 == -544) {
                                    fmt.Println(fi12)
                                } else {
                                    v14 := GenS1{F0: ("\n\t" <= "foo bar"), F1: 2}
                                    var v15 int8 = 71
                                    _ = v14
                                    _ = v15
                                }
                                var v16 int64 = 48
                                v17 := GenS2{F0: 1, F1: "\n\t", F2: 0}
                                _ = fi13
                                _ = v16
                                _ = v17
                            }
                        }
                        switch ("foo bar" + "a") {
                        case "":
                            v18 := GenS3{F0: 0}
                            _ = v18
                        case "a":
                            delete(v3, "")
                        case "foo bar":
                            if (0 == 203) {
                                for fi19 := 0; fi19 < 3; fi19++ {
                                    if true {
                                        var v20 int = fi19
                                        var v21 string = "hello"
                                        v22 := GenS4{F0: "\n\t", F1: 606, F2: 32767}
                                        _ = v20
                                        _ = v21
                                        _ = v22
                                    }
                                    _ = fi19
                                }
                            } else {
                                var v23 int32 = -2147483648
                                for fi24 := 0; fi24 < 6; fi24++ {
                                    delete(v3, (("foo bar" + "foo bar") + ("" + "\n\t")))
                                    v25 := GenS2{F0: -2147483648, F1: (("" + (((("" + "") + "hello") + ("" + "")) + "foo bar")) + ("hello" + "foo bar")), F2: (^(v2))}
                                    _ = fi24
                                    _ = v25
                                }
                                _ = v23
                            }
                        }
                        v2 += v2
                    }
                    var v26 uint64 = 4294967295
                    _ = fi12
                    _ = v26
                }
            case "a":
                var v27 int16 = 32767
                _ = v27
            case "\n\t":
                var v28 int = 32768
                _ = v28
            default:
                v2 += 1
            }
        }
        var v29 int16 = -930
        _ = v29
    }
    delete(v3, "foo bar")
    v30 := GenS5{F0: -741, F1: true, F2: 65535}
    switch "foo bar" {
    case "":
        switch (("foo bar" + "") + "") {
        case "a":
            v2 = (v2 | (v2 % ((v2) | 1)))
        case "foo bar":
            var v31 int = 32767
            _ = v31
        }
    default:
        v2 = (v2 - (v2 + (v2 % ((42) | 1))))
    }
    v2 = 18446744073709551615
    var v32 uint64 = (v2 | v2)
    var v33 uint16 = 0
    if false {
        v2 -= (v2 >> ((18446744073709551615) & 63))
        switch -9223372036854775808 {
        case 32767:
            if (76 >= 42) {
                v33 += v33
            } else {
                var v34 string = ""
                fmt.Println(v32, v33, v34)
                var v35 int64 = 540
                _ = v34
                _ = v35
            }
        case -9223372036854775808:
            fmt.Println(v32, v2)
        case 639:
            fmt.Println(v33)
        }
        for fi36 := 0; fi36 < 2; fi36++ {
            fmt.Println(v32)
            for fi37 := 0; fi37 < 17; fi37++ {
                fmt.Println(fi37, v2)
                v38 := GenS6{F0: 4294967295, F1: 2147483647, F2: 47}
                v39 := GenS7{F0: 0, F1: 1, F2: 2}
                _ = fi37
                _ = v38
                _ = v39
            }
            _ = fi36
        }
    } else {
        switch (("a" + "a") + "hello") {
        case "":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "foo bar":
            var v40 uint32 = 4294967295
            _ = v40
        case "a":
            delete(v3, "hello")
        default:
            delete(v3, "foo bar")
        }
        var v41 int16 = 32767
        if (!true) {
            if true {
                delete(v3, ("foo bar" + ""))
                v1 = append(v1, "hello")
                for fi42 := 0; fi42 < 11; fi42++ {
                    var v43 int64 = 218
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi42
                    _ = v43
                }
            } else {
                if (true && ((127 <= -128) && true)) {
                    for fi44 := 0; fi44 < 9; fi44++ {
                        v1 = append(v1, "a")
                        _ = fi44
                    }
                }
                fmt.Println(v2, v41)
            }
            var v45 int16 = (v41 ^ v41)
            _ = v45
        } else {
            v3[("\n\t" + "foo bar")] = "hello"
        }
        _ = v41
    }
    fmt.Println(v33)
    if (1 != 1) {
        if true {
            var v46 uint16 = v33
            _ = v46
        }
        v1 = append(v1, (("a" + "foo bar") + "a"))
    }
    var v47 uint64 = v2
    var v48 int16 = 32767
    var v49 float32 = 0.1
    var v50 uint64 = v32
    v47 += v50
    if (!(!false)) {
        var v51 uint = 2
        _ = v51
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v47
    _ = v48
    _ = v49
    _ = v50
}
