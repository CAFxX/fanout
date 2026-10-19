package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 int; F2 int64 }

func main() {
    v1 := []int{-2147483648, 280, -554}
    v2 := []string{"hello", ("\n\t" + "a"), "\n\t"}
    v3 := []string{("\n\t" + ("hello" + "a")), ("" + "\n\t"), ("a" + ("a" + (("hello" + "\n\t") + (("foo bar" + "\n\t") + ("\n\t" + "")))))}
    var v4 int32 = 2147483647
    v5 := make(map[int]int)
    var v6 int16 = -282
    v4 = v4
    switch -1 {
    case -3:
        if false {
            var v7 int16 = v6
            _ = v7
        }
    default:
        v6 -= v6
    }
    v8 := GenS1{F0: -76}
    if false {
        var v9 int8 = 100
        if false {
            v2 = append(v2, ((("foo bar" + "foo bar") + "hello") + "a"))
        }
        _ = v9
    } else {
        switch -848 {
        case -1:
            v10 := GenS1{F0: -128}
            _ = v10
        case 73:
            var v11 float32 = 0.1
            _ = v11
        case -38:
            for fi12 := 0; fi12 < 16; fi12++ {
                for fi13 := 0; fi13 < 8; fi13++ {
                    v4 = -585
                    _ = fi13
                }
                _ = fi12
            }
        default:
            if false {
                v14 := GenS2{F0: 0, F1: 2, F2: 2}
                if len(v1) > 0 {
                    v1[0] = -197
                }
                _ = v14
            }
        }
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 uint32 = 65535
        if (0 > 100) {
            switch "foo bar" {
            case "\n\t":
                if true {
                    if false {
                        v17 := GenS3{F0: 128}
                        fmt.Println(v6)
                        _ = v17
                    } else {
                        if (true || true) {
                            _ = len(v2)
                        }
                        _ = len(v2)
                    }
                    if (false || (65535 == 42)) {
                        fmt.Println(v6, fi15)
                    }
                    v16 -= 42
                } else {
                    fmt.Println(v16, fi15, fi15)
                    fmt.Println(v6)
                }
            default:
                for fi18 := 0; fi18 < 5; fi18++ {
                    fmt.Println(v6, fi15, fi18)
                    v4 = (-(v4))
                    v19 := GenS3{F0: 128}
                    _ = fi18
                    _ = v19
                }
            }
            v16 = (v16 + (v16 | 42))
        } else {
            for fi20 := 0; fi20 < 16; fi20++ {
                v4 = v4
                v6 = 0
                _ = fi20
            }
            for fi21 := 0; fi21 < 10; fi21++ {
                if ((1 < 420) && (42 != v16)) {
                    v3 = append(v3, "\n\t")
                    switch "a" {
                    case "\n\t":
                        switch -2147483648 {
                        case 42:
                            var v22 string = "foo bar"
                            _ = v22
                        case 9223372036854775807:
                            v23 := GenS2{F0: 2, F1: 42, F2: 30}
                            _ = v23
                        case 120:
                            v4 -= 315
                        }
                    }
                    fmt.Println(fi15, fi15)
                } else {
                    _ = len(v1)
                    var v24 uint16 = 2
                    _ = v24
                }
                _ = fi21
            }
        }
        _ = fi15
        _ = v16
    }
    fmt.Println(v4)
    for fi25 := 0; fi25 < 4; fi25++ {
        switch fi25 {
        case -816:
            v26 := GenS3{F0: 218}
            _ = v26
        default:
            for fi27 := 0; fi27 < 14; fi27++ {
                if false {
                    var v28 int = fi25
                    fmt.Println(fi25, fi25)
                    _ = v28
                } else {
                    fmt.Println(v4)
                }
                var v29 uint8 = 125
                var v30 string = "a"
                _ = fi27
                _ = v29
                _ = v30
            }
        }
        v31 := GenS3{F0: 0}
        _ = fi25
        _ = v31
    }
    _ = len(v5)
    var v32 uint = 65535
    fmt.Println(v32)
    var v33 int8 = -34
    switch "\n\t" {
    case "\n\t":
        if (722 != 302) {
            var v34 uint64 = uint64(v32)
            for fi35 := 0; fi35 < 20; fi35++ {
                for fi36 := 0; fi36 < 15; fi36++ {
                    v37 := GenS4{F0: ("" + ("\n\t" + "foo bar")), F1: 2147483647, F2: 334}
                    _ = fi36
                    _ = v37
                }
                for fi38 := 0; fi38 < 13; fi38++ {
                    var v39 uint64 = (v34 - 0)
                    v40 := GenS1{F0: v33}
                    fmt.Println(fi35)
                    _ = fi38
                    _ = v39
                    _ = v40
                }
                delete(v5, fi35)
                _ = fi35
            }
            var v41 int8 = v33
            _ = v34
            _ = v41
        } else {
            v5[127] = 32768
            v1 = append(v1, 9223372036854775807)
            var v42 uint8 = 233
            _ = v42
        }
    case "hello":
        switch -1 {
        case -277:
            v33 -= -95
        case 127:
            fmt.Println(v4, v6)
        default:
            switch 1 {
            case 32768:
                switch -1 {
                case -579:
                    switch ("foo bar" + "foo bar") {
                    case "\n\t":
                        if true {
                            var v43 int = 1
                            var v44 int32 = (-(v4))
                            for fi45 := 0; fi45 < 8; fi45++ {
                                var v46 string = "foo bar"
                                _ = fi45
                                _ = v46
                            }
                            _ = v43
                            _ = v44
                        }
                    default:
                        v32 = (v32 ^ (v32 << ((v32) & 63)))
                    }
                default:
                    v4 = v4
                }
            case -311:
                for fi47 := 0; fi47 < 17; fi47++ {
                    v6 -= 318
                    _ = fi47
                }
            case 4294967295:
                if (false || (true || false)) {
                    if true {
                        v6 = 42
                        v48 := v5[int(v33)]
                        for fi49 := 0; fi49 < 7; fi49++ {
                            var v50 bool = (true || (!(2 >= 2)))
                            delete(v5, v48)
                            _ = fi49
                            _ = v50
                        }
                        _ = v48
                    } else {
                        v32 += 18446744073709551615
                    }
                    var v51 float64 = 0.0
                    var v52 uint64 = 4294967295
                    _ = v51
                    _ = v52
                } else {
                    v53 := GenS1{F0: (v33 & 100)}
                    _ = len(v5)
                    for fi54 := 0; fi54 < 19; fi54++ {
                        if ((v6 - v6) == 1) {
                            for fi55 := 0; fi55 < 9; fi55++ {
                                v33 = v33
                                v56 := GenS3{F0: 2}
                                fmt.Println(v4)
                                _ = fi55
                                _ = v56
                            }
                            v4 -= (v4 | v4)
                        } else {
                            var v57 float64 = 0.0
                            delete(v5, (fi54 >> (((^(v32))) & 63)))
                            var v58 int8 = v33
                            _ = v57
                            _ = v58
                        }
                        _ = fi54
                    }
                    _ = v53
                }
            default:
                v33 = -128
            }
        }
    default:
        if true {
            v33 -= 100
            if (true || (v6 < (v6 * v6))) {
                v4 += (v4 - v4)
                if true {
                    v32 = (v32 >> (((v32 | 2)) & 63))
                    _ = len(v3)
                    v33 += (^(v33))
                }
            }
            var v59 float32 = 0.0
            _ = v59
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
