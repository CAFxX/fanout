package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 string; F1 uint16; F2 uint32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS7 struct { F0 uint32; F1 int8; F2 bool }
type GenS8 struct { F0 int64; F1 uint64 }
type GenS9 struct { F0 string; F1 string; F2 uint }
type GenS10 struct { F0 uint8 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]bool)
    var v3 int16 = 1000
    var v4 uint8 = 42
    var v5 uint32 = 0
    v6 := []int{}
    if true {
        for fi7 := 0; fi7 < 18; fi7++ {
            for fi8 := 0; fi8 < 13; fi8++ {
                _ = len(v2)
                _ = fi8
            }
            v9 := v1[((("\n\t" + "") + "") + ("\n\t" + ""))]
            if (595 < v3) {
                var v10 int32 = int32(v4)
                if (9223372036854775807 == -1) {
                    v11 := v1[v9]
                    _ = v11
                }
                for fi12 := 0; fi12 < 8; fi12++ {
                    fmt.Println(v3)
                    switch ("" + (v9 + "")) {
                    case "":
                        if true {
                            fmt.Println(fi7, fi7)
                        } else {
                            if (986 == 0) {
                                var v13 int32 = v10
                                _ = len(v1)
                                v9 = v9
                                _ = v13
                            }
                            v6 = append(v6, 1)
                        }
                    case "hello":
                        var v14 int = (fi12 << ((uint(v4)) & 63))
                        _ = v14
                    case "a":
                        v15 := GenS1{F0: -9223372036854775808, F1: (v10 * (v10 * -715))}
                        _ = v15
                    }
                    _ = fi12
                }
                _ = v10
            } else {
                var v16 int8 = 127
                _ = v16
            }
            _ = fi7
            _ = v9
        }
        switch 761 {
        case 42:
            for fi17 := 0; fi17 < 3; fi17++ {
                v18 := GenS1{F0: 0, F1: 0}
                _ = fi17
                _ = v18
            }
        case 0:
            if (v5 > (v5 << ((0) & 63))) {
                if (!(260 > 18446744073709551615)) {
                    fmt.Println(v4, v3, v5)
                    if ("\n\t" <= ("\n\t" + "")) {
                        v4 += uint8(v5)
                        for fi19 := 0; fi19 < 3; fi19++ {
                            v20 := GenS2{F0: fi19, F1: (^(fi19))}
                            _ = fi19
                            _ = v20
                        }
                    }
                    var v21 uint8 = (^(v4))
                    _ = v21
                }
                delete(v2, v5)
                var v22 uint16 = 2
                _ = v22
            } else {
                switch -936 {
                case -411:
                    v1["a"] = "a"
                default:
                    v3 += -32768
                }
                v23 := GenS3{F0: "hello", F1: 42, F2: v5}
                var v24 uint16 = 255
                _ = v23
                _ = v24
            }
        case -253:
            v3 = (v3 - -32768)
        default:
            v4 = (v4 + 128)
        }
    }
    v25 := GenS4{F0: 65535}
    v26 := GenS3{F0: "\n\t", F1: 0, F2: v5}
    fmt.Println(v3)
    switch "a" {
    case "foo bar":
        switch "\n\t" {
        case "\n\t":
            v27 := GenS1{F0: -1, F1: 2147483647}
            _ = v27
        case "foo bar":
            var v28 float64 = 1e10
            _ = v28
        default:
            fmt.Println(v4, v4)
        }
    case "hello":
        var v29 int = 0
        _ = v29
    case "\n\t":
        v30 := GenS3{F0: "hello", F1: 42, F2: v5}
        _ = v30
    }
    if true {
        if false {
            _ = len(v2)
        } else {
            switch ("foo bar" + "\n\t") {
            case "foo bar":
                v31 := GenS2{F0: -1, F1: -47}
                _ = v31
            case "\n\t":
                v5 = (v5 & v5)
            case "":
                v4 += v4
            default:
                v3 = v3
            }
            for fi32 := 0; fi32 < 1; fi32++ {
                var v33 int32 = -1
                var v34 uint8 = 255
                if false {
                    var v35 uint16 = 42
                    v36 := GenS3{F0: ("hello" + "hello"), F1: 42, F2: (v5 << ((1) & 63))}
                    _ = v35
                    _ = v36
                } else {
                    v34 -= (v34 << ((347) & 63))
                    fmt.Println(v4)
                    fmt.Println(v5, v34)
                }
                _ = fi32
                _ = v33
                _ = v34
            }
        }
    }
    switch 4294967295 {
    case -9223372036854775808:
        if true {
            v4 = v4
            v6 = append(v6, 2147483647)
        } else {
            var v37 uint32 = v5
            v38 := GenS2{F0: 42, F1: 127}
            _ = v37
            _ = v38
        }
    case 4294967295:
        delete(v2, 493)
    }
    v39 := GenS4{F0: 368}
    v40 := GenS5{F0: false, F1: 2}
    if false {
        v41 := GenS6{F0: (v5 * 1), F1: 0, F2: 2}
        v3 = (v3 >> ((358) & 63))
        v42 := GenS5{F0: true, F1: 564}
        _ = v41
        _ = v42
    } else {
        v43 := GenS7{F0: 2147483648, F1: 127, F2: false}
        v44 := GenS8{F0: -9223372036854775808, F1: 2}
        switch 2147483647 {
        case -645:
            var v45 bool = (127 <= -103)
            _ = v45
        default:
            for fi46 := 0; fi46 < 3; fi46++ {
                for fi47 := 0; fi47 < 15; fi47++ {
                    fmt.Println(v5, v4)
                    var v48 int = -2147483648
                    var v49 uint32 = v5
                    _ = fi47
                    _ = v48
                    _ = v49
                }
                _ = fi46
            }
        }
        _ = v43
        _ = v44
    }
    for fi50 := 0; fi50 < 10; fi50++ {
        var v51 int16 = (v3 >> ((18446744073709551615) & 63))
        if (1 > -1) {
            v1["hello"] = ("\n\t" + "")
        } else {
            fmt.Println(fi50, v4, v3)
            v52 := GenS3{F0: "hello", F1: 0, F2: 729}
            if len(v6) > 0 {
                v6[0] = fi50
            }
            _ = v52
        }
        _ = fi50
        _ = v51
    }
    if (!("a" > (("foo bar" + "a") + ""))) {
        switch 2147483647 {
        case 2147483647:
            var v53 int8 = 42
            _ = v53
        case -65:
            v54 := GenS9{F0: "", F1: (("hello" + ("\n\t" + "\n\t")) + ("\n\t" + "a")), F2: 18446744073709551615}
            _ = v54
        case -1:
            v5 += 990
        default:
            if (!((((true && (84 < -907)) && false) && ("hello" >= "")) && true)) {
                switch 2147483647 {
                case -293:
                    v55 := GenS10{F0: 233}
                    _ = v55
                case 943:
                    var v56 float64 = ((((3.14159 / -2.25) / (float64(0.1) - 1.5)) - (1e10 + float64(1.5))) + 3.14159)
                    _ = v56
                case 177:
                    var v57 bool = false
                    _ = v57
                }
                var v58 int = 164
                _ = v58
            }
        }
    } else {
        for fi59 := 0; fi59 < 1; fi59++ {
            v60 := v2[v5]
            _ = fi59
            _ = v60
        }
        v3 = (v3 | v3)
    }
    fmt.Println(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
