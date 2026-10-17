package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 uint; F1 bool }
type GenS6 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint16; F1 bool }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 int16; F1 int32; F2 string }
type GenS14 struct { F0 int16; F1 uint8; F2 uint8 }

func main() {
    var v1 uint = 615
    var v2 int8 = 42
    var v3 uint32 = 4294967295
    v4 := make(map[string]uint32)
    var v5 uint16 = 988
    var v6 uint8 = 127
    fmt.Println(v6, v6, v3)
    fmt.Println(v1)
    for fi7 := 0; fi7 < 20; fi7++ {
        fmt.Println(v3, fi7)
        v8 := GenS1{F0: (fi7 >> ((42) & 63)), F1: v3}
        v1 += (v1 << ((v1) & 63))
        _ = fi7
        _ = v8
    }
    v3 = 4294967295
    for fi9 := 0; fi9 < 10; fi9++ {
        v6 = (v6 >> ((v1) & 63))
        var v10 uint16 = (^(v5))
        _ = fi9
        _ = v10
    }
    v4[("foo bar" + "hello")] = (^(v3))
    if ((v1 & (v1 << ((18446744073709551615) & 63))) != (v1 - (v1 << ((4294967295) & 63)))) {
        if (v5 > v5) {
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v2, v6)
            case "a":
                v6 = (v6 ^ (v6 & (v6 | v6)))
            case "foo bar":
                v5 = 678
            default:
                if false {
                    for fi11 := 0; fi11 < 4; fi11++ {
                        for fi12 := 0; fi12 < 10; fi12++ {
                            switch ("foo bar" + "a") {
                            case "\n\t":
                                if false {
                                    v13 := GenS1{F0: -949, F1: uint32(v1)}
                                    if (!(!true)) {
                                        switch "hello" {
                                        case "a":
                                            v1 = (^(v1))
                                        case "":
                                            delete(v4, "hello")
                                        default:
                                            var v14 uint64 = 4294967295
                                            _ = v14
                                        }
                                        v15 := GenS1{F0: (fi11 % ((fi11) | 1)), F1: v3}
                                        _ = v15
                                    }
                                    _ = v13
                                }
                            default:
                                v16 := GenS2{F0: (v2 << (((v1 * (v1 * uint(v5)))) & 63))}
                                _ = v16
                            }
                            v17 := GenS3{F0: uint8(v1), F1: 407, F2: 42}
                            v18 := GenS4{F0: -128, F1: v2}
                            _ = fi12
                            _ = v17
                            _ = v18
                        }
                        v5 = 65535
                        _ = fi11
                    }
                    v1 = v1
                } else {
                    if ((false || (!(false && (!true)))) && true) {
                        v19 := GenS5{F0: 745, F1: false}
                        var v20 float32 = 1.5
                        _ = v19
                        _ = v20
                    } else {
                        if true {
                            var v21 string = "foo bar"
                            fmt.Println(v2, v1)
                            _ = v21
                        }
                        fmt.Println(v3)
                        if true {
                            var v22 uint16 = v5
                            _ = v22
                        } else {
                            v1 += (v1 | 65535)
                        }
                    }
                    var v23 string = "hello"
                    v24 := GenS4{F0: (v2 ^ (v2 + (^(v2)))), F1: 1}
                    _ = v23
                    _ = v24
                }
            }
            for fi25 := 0; fi25 < 19; fi25++ {
                v1 += v1
                if (1e10 <= (3.14159 - 3.14159)) {
                    for fi26 := 0; fi26 < 8; fi26++ {
                        if (((2 > v3) || (1 >= 2)) || false) {
                            v1 += (v1 >> ((v1) & 63))
                            v27 := GenS5{F0: v1, F1: true}
                            var v28 uint32 = 1
                            var v29 uint = v1
                            _ = v27
                            _ = v28
                            _ = v29
                        } else {
                            var v30 string = "a"
                            var v31 uint16 = (v5 % ((v5) | 1))
                            _ = v30
                            _ = v31
                        }
                        v2 -= v2
                        _ = fi26
                    }
                    v32 := v4["a"]
                    _ = v32
                }
                _ = fi25
            }
            for fi33 := 0; fi33 < 11; fi33++ {
                var v34 int8 = 127
                for fi35 := 0; fi35 < 19; fi35++ {
                    switch fi35 {
                    case 0:
                        fmt.Println(fi33, fi35, v3)
                    case 1:
                        fmt.Println(fi33, v34, v1)
                    default:
                        if false {
                            if false {
                                switch (fi35 | 32767) {
                                case -175:
                                    fmt.Println(v5, fi33)
                                case 2147483647:
                                    v4[("hello" + "foo bar")] = 0
                                case 9223372036854775807:
                                    var v36 uint = 247
                                    _ = v36
                                }
                                if (("a" + ("" + "")) == "hello") {
                                    var v37 int64 = -819
                                    var v38 uint = v1
                                    if (v1 > (v1 | (v38 * (v1 - (v38 + (v1 >> ((v38) & 63))))))) {
                                        v4[("" + "\n\t")] = 512
                                        if false {
                                            var v39 int16 = 1
                                            v40 := GenS4{F0: (^(v2)), F1: v2}
                                            _ = v39
                                            _ = v40
                                        }
                                        fmt.Println(v3, v1, v34)
                                    } else {
                                        if false {
                                            v1 -= 65535
                                            v1 = uint(v3)
                                            var v41 int32 = 2
                                            _ = v41
                                        } else {
                                            if (uint8(v1) >= 127) {
                                                var v42 string = ("foo bar" + "")
                                                v43 := v4[((v42 + ("hello" + v42)) + "hello")]
                                                _ = v42
                                                _ = v43
                                            }
                                            v44 := GenS6{F0: v6, F1: 961, F2: v5}
                                            _ = v44
                                        }
                                        v45 := GenS7{F0: v2}
                                        switch fi35 {
                                        case 42:
                                            v46 := GenS5{F0: 746, F1: (true || ("hello" > ""))}
                                            _ = v46
                                        default:
                                            v47 := []int{(fi35 << (((v38 % ((v1) | 1))) & 63)), (^(fi35)), (fi33 + (fi33 + (fi33 ^ fi35))), fi33}
                                            _ = v47
                                        }
                                        _ = v45
                                    }
                                    fmt.Println(v1, v1)
                                    v48 := GenS4{F0: v2, F1: v34}
                                    v6 -= 153
                                    _ = v37
                                    _ = v38
                                    _ = v48
                                } else {
                                    if false {
                                        var v49 int8 = v2
                                        _ = v49
                                    }
                                    v50 := []int{}
                                    v2 += int8(v1)
                                    _ = v50
                                }
                                v51 := GenS8{F0: -1}
                                _ = v51
                            }
                            v52 := GenS4{F0: -128, F1: -62}
                            _ = v52
                        }
                    }
                    switch "a" {
                    case "hello":
                        if false {
                            var v53 uint32 = (v3 << (((v1 + v1)) & 63))
                            _ = v53
                        } else {
                            if (false && false) {
                                if (!(!true)) {
                                    v54 := []int{-2147483648, (^(fi33))}
                                    _ = v54
                                } else {
                                    v3 = 152
                                    fmt.Println(v34)
                                    v1 = v1
                                }
                                v5 = 1
                                v1 -= v1
                            } else {
                                v55 := GenS3{F0: 25, F1: 1, F2: 877}
                                switch 32768 {
                                case 32768:
                                    var v56 int32 = -315
                                    _ = v56
                                }
                                var v57 uint16 = 809
                                _ = v55
                                _ = v57
                            }
                            fmt.Println(v2, fi33, v34)
                        }
                    }
                    switch "" {
                    case "":
                        v5 = v5
                        fmt.Println(v5, v5)
                        v2 = (v34 | v2)
                    default:
                        fmt.Println(v2, fi33, v6)
                    }
                    v58 := GenS9{F0: (v1 & v1)}
                    _ = fi35
                    _ = v58
                }
                _ = fi33
                _ = v34
            }
        } else {
            v59 := GenS10{F0: (^(v5)), F1: true}
            for fi60 := 0; fi60 < 8; fi60++ {
                fmt.Println(v5, v1)
                _ = fi60
            }
            _ = v59
        }
        for fi61 := 0; fi61 < 12; fi61++ {
            v5 = (v5 ^ uint16(v6))
            var v62 int32 = 42
            var v63 float64 = 1.5
            _ = fi61
            _ = v62
            _ = v63
        }
        var v64 bool = false
        _ = v64
    }
    for fi65 := 0; fi65 < 7; fi65++ {
        fmt.Println(v1)
        v66 := GenS11{F0: 66}
        for fi67 := 0; fi67 < 12; fi67++ {
            fmt.Println(v3, v2)
            _ = fi67
        }
        _ = fi65
        _ = v66
    }
    v68 := GenS12{F0: true}
    switch "foo bar" {
    case "hello":
        var v69 int8 = 127
        _ = v69
    default:
        var v70 bool = true
        _ = v70
    }
    fmt.Println(v1, v5, v1)
    v71 := GenS13{F0: -322, F1: -2147483648, F2: "\n\t"}
    var v72 string = ("hello" + (((("" + "\n\t") + "") + "\n\t") + "a"))
    v3 += 545
    for fi73 := 0; fi73 < 20; fi73++ {
        v74 := GenS14{F0: int16(v1), F1: 156, F2: v6}
        for fi75 := 0; fi75 < 10; fi75++ {
            var v76 uint = 230
            _ = fi75
            _ = v76
        }
        _ = fi73
        _ = v74
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v71
    _ = v72
}
