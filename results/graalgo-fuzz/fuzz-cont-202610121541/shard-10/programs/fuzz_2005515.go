package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 string; F2 uint16 }
type GenS6 struct { F0 string; F1 int8; F2 int32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint; F1 bool; F2 string }
func genFn0(p0 int, p1 string, p2 bool) int8 {
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 int = (fi5 % ((32768) | 1))
        _ = fi5
        _ = v6
    }
    p0 += (p0 & (p0 << ((388) & 63)))
    if false {
        for fi7 := 0; fi7 < 20; fi7++ {
            v8 := GenS1{F0: 127, F1: 0}
            _ = fi7
            _ = v8
        }
        var v9 int64 = 481
        _ = v9
    } else {
        var v10 int16 = 1
        var v11 int16 = (v10 | v10)
        fmt.Println(p2, v11, v10)
        _ = v10
        _ = v11
    }
    v12 := GenS2{F0: (p1 < p1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 42
}
func genFn1(p0 int16) uint {
    v13 := GenS3{F0: 2147483648, F1: 18446744073709551615, F2: (p0 >> ((783) & 63))}
    fmt.Println(p0)
    p0 -= (p0 ^ p0)
    if (!((!(717 != 255)) || false)) {
        var v14 uint32 = 2147483647
        if false {
            if true {
                v15 := []int{}
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                _ = v15
            }
            var v16 uint64 = 538
            fmt.Println(v16, p0, p0)
            _ = v16
        }
        _ = v14
    } else {
        v17 := GenS1{F0: 2, F1: -2147483648}
        p0 = (p0 * -774)
        _ = v17
    }
    v18 := []int{127}
    _ = p0
    _ = v13
    _ = v18
    return 1
}

func main() {
    var v1 int32 = -180
    var v2 int8 = 127
    v3 := make(map[string]bool)
    var v4 uint32 = 42
    for fi19 := 0; fi19 < 2; fi19++ {
        if (false || (((false || true) || true) && false)) {
            var v20 bool = (!true)
            if ((v20 || (v1 != v1)) || (42 < 177)) {
                if (v4 >= (v4 - (v4 % (((v4 * 2)) | 1)))) {
                    var v21 uint64 = 2
                    v4 -= (v4 & (v4 + (v4 % ((v4) | 1))))
                    if (!(v20 || (866 <= -9223372036854775808))) {
                        v21 += 2
                    } else {
                        if false {
                            var v22 int32 = (v1 ^ 1)
                            fmt.Println(v22, v1, v1)
                            switch (fi19 + 128) {
                            case 650:
                                v1 = int32(v4)
                            default:
                                v21 = v21
                            }
                            _ = v22
                        } else {
                            v4 = (v4 - 1)
                            fmt.Println(v1, v1)
                            var v23 int16 = -1
                            _ = v23
                        }
                        v24 := GenS4{F0: 1}
                        var v25 int64 = 0
                        _ = v24
                        _ = v25
                    }
                    _ = v21
                }
                var v26 int16 = 32767
                for fi27 := 0; fi27 < 13; fi27++ {
                    fmt.Println(v2, v4)
                    if (18446744073709551615 == genFn1(32767)) {
                        var v28 uint16 = 32767
                        _ = v28
                    }
                    var v29 uint32 = v4
                    _ = fi27
                    _ = v29
                }
                _ = v26
            }
            _ = v20
        }
        var v30 uint64 = 4294967295
        v4 = v4
        _ = fi19
        _ = v30
    }
    v4 -= 42
    var v31 int32 = v1
    v32 := GenS4{F0: -602}
    fmt.Println(v2, v1)
    var v33 int = int(v4)
    if false {
        delete(v3, "\n\t")
    }
    var v34 int32 = -1
    if false {
        v34 -= v31
        v35 := GenS5{F0: v33, F1: "hello", F2: 0}
        if (!((!true) || false)) {
            if false {
                v36 := GenS3{F0: (v4 >> ((585) & 63)), F1: 4294967295, F2: 42}
                _ = v36
            } else {
                switch 1 {
                case -2147483648:
                    switch "foo bar" {
                    case "foo bar":
                        for fi37 := 0; fi37 < 15; fi37++ {
                            if true {
                                v38 := GenS3{F0: 0, F1: 318, F2: 523}
                                for fi39 := 0; fi39 < 18; fi39++ {
                                    switch (fi39 / ((-176) | 1)) {
                                    case -76:
                                        v1 = (v1 << ((uint(v4)) & 63))
                                    case 51:
                                        switch "hello" {
                                        case "a":
                                            var v40 float32 = 1e10
                                            _ = v40
                                        case "\n\t":
                                            var v41 string = ("\n\t" + ("hello" + (("" + ("foo bar" + "")) + "foo bar")))
                                            _ = v41
                                        default:
                                            v34 = (v34 ^ (v34 / ((-2147483648) | 1)))
                                        }
                                    case -1:
                                        if (762 == -117) {
                                            fmt.Println(v2, fi37, v4)
                                            switch 32768 {
                                            case 32768:
                                                v3["\n\t"] = (!(!("foo bar" >= "a")))
                                            default:
                                                fmt.Println(v34)
                                            }
                                            var v42 bool = (fi37 <= 850)
                                            _ = v42
                                        } else {
                                            var v43 int16 = 0
                                            v44 := GenS6{F0: "\n\t", F1: v2, F2: v31}
                                            fmt.Println(v43, v33, v1)
                                            _ = v43
                                            _ = v44
                                        }
                                        v45 := GenS6{F0: "hello", F1: int8(v4), F2: 42}
                                        _ = v45
                                    default:
                                        v46 := []int{(fi39 & fi39), int(v31)}
                                        _ = v46
                                    }
                                    if ("foo bar" <= "a") {
                                        v1 += v1
                                        v47 := GenS7{F0: (v33 - (v33 & (^(v33))))}
                                        _ = len(v3)
                                        _ = v47
                                    } else {
                                        v48 := GenS8{F0: 477, F1: false, F2: ""}
                                        switch (^(v33)) {
                                        case -9223372036854775808:
                                            v49 := GenS7{F0: fi37}
                                            _ = v49
                                        }
                                        _ = v48
                                    }
                                    _ = fi39
                                }
                                _ = v38
                            } else {
                                if ((v4 | 0) <= (^(v4))) {
                                    var v50 int64 = -9223372036854775808
                                    for fi51 := 0; fi51 < 3; fi51++ {
                                        switch "\n\t" {
                                        case "\n\t":
                                            v34 = (v34 >> ((4294967295) & 63))
                                        default:
                                            v52 := GenS4{F0: v50}
                                            _ = v52
                                        }
                                        _ = fi51
                                    }
                                    v31 = v31
                                    _ = v50
                                } else {
                                    switch -1 {
                                    case 4294967295:
                                        if true {
                                            var v53 bool = false
                                            _ = v53
                                        } else {
                                            v54 := v3[((("hello" + "hello") + ((("hello" + ("a" + "foo bar")) + "hello") + "\n\t")) + "hello")]
                                            _ = v54
                                        }
                                    default:
                                        v3["a"] = (-2.25 >= -2.25)
                                    }
                                }
                                delete(v3, "")
                            }
                            var v55 int64 = 9223372036854775807
                            _ = fi37
                            _ = v55
                        }
                    case "a":
                        if true {
                            fmt.Println(v33, v1)
                            v56 := GenS5{F0: (v33 ^ v33), F1: "hello", F2: 65535}
                            _ = v56
                        } else {
                            switch "a" {
                            case "":
                                v33 -= (v33 - 1)
                            }
                        }
                    }
                }
                v34 -= (v1 - v34)
            }
        } else {
            var v57 bool = true
            _ = v57
        }
        _ = v35
    } else {
        v34 += v1
        switch v33 {
        case 1:
            switch ("a" + "\n\t") {
            case "foo bar":
                for fi58 := 0; fi58 < 10; fi58++ {
                    var v59 float32 = ((float32(0.1) - 0.0) - 0.0)
                    v33 = fi58
                    _ = fi58
                    _ = v59
                }
            case "":
                var v60 int64 = -150
                _ = v60
            case "hello":
                var v61 uint16 = 32767
                _ = v61
            }
        case 32768:
            v31 += 1
        }
        fmt.Println(v2)
    }
    v4 += v4
    v33 += 0
    for fi62 := 0; fi62 < 9; fi62++ {
        var v63 uint32 = (v4 >> ((genFn1(1)) & 63))
        _ = fi62
        _ = v63
    }
    for fi64 := 0; fi64 < 14; fi64++ {
        v1 = v1
        switch v33 {
        case 2147483647:
            delete(v3, "a")
        case 35:
            var v65 uint = 65535
            _ = v65
        case -1:
            var v66 uint64 = 42
            _ = v66
        default:
            v67 := []int{}
            _ = v67
        }
        _ = fi64
    }
    v68 := GenS7{F0: 404}
    var v69 int8 = (v2 << ((42) & 63))
    fmt.Println(v69, v31, v34)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v68
    _ = v69
}
