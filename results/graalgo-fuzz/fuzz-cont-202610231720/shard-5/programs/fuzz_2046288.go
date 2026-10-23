package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 bool; F2 int16 }

func main() {
    v1 := []string{"\n\t", "a", ((("a" + (((("hello" + ("a" + "")) + ("" + ((("hello" + ("" + ("" + "hello"))) + "foo bar") + "foo bar"))) + "foo bar") + "\n\t")) + "foo bar") + (("" + "\n\t") + "a"))}
    var v2 int16 = -1
    v3 := make(map[string]string)
    var v4 uint32 = 292
    if (true || (!true)) {
        switch ("\n\t" + "a") {
        case "foo bar":
            v1 = append(v1, "foo bar")
        }
    }
    var v5 string = "a"
    fmt.Println(v4)
    v4 = (v4 << ((255) & 63))
    if false {
        var v6 bool = (!((v5 + ((v5 + v5) + (("\n\t" + v5) + v5))) > (("foo bar" + (("\n\t" + v5) + "hello")) + v5)))
        switch 1 {
        case 42:
            for fi7 := 0; fi7 < 5; fi7++ {
                if false {
                    for fi8 := 0; fi8 < 3; fi8++ {
                        var v9 int32 = 1
                        if (fi8 == (fi8 * fi8)) {
                            switch v5 {
                            case "hello":
                                fmt.Println(v9)
                            case "foo bar":
                                v5 = ("foo bar" + ("foo bar" + v5))
                            default:
                                switch fi7 {
                                case 32767:
                                    v2 += 42
                                case -2147483648:
                                    v10 := GenS1{F0: (!(127 == 1)), F1: -109}
                                    _ = v10
                                default:
                                    if v6 {
                                        v4 -= (v4 | v4)
                                        var v11 float64 = 1e10
                                        _ = v11
                                    } else {
                                        v4 = (v4 + (v4 | (v4 << ((65535) & 63))))
                                    }
                                }
                            }
                        }
                        _ = fi8
                        _ = v9
                    }
                    delete(v3, (("hello" + "foo bar") + v5))
                } else {
                    if true {
                        if ((^(v4)) >= v4) {
                            var v12 uint = 899
                            v5 = v5
                            _ = v12
                        } else {
                            for fi13 := 0; fi13 < 2; fi13++ {
                                v6 = (!(v4 != (^(v4))))
                                _ = fi13
                            }
                            if v6 {
                                v2 -= (^(v2))
                                var v14 int32 = 2147483647
                                _ = v14
                            }
                        }
                        _ = len(v1)
                        v6 = (0 != 1)
                    } else {
                        v15 := GenS2{F0: 9223372036854775807}
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 17; fi16++ {
                        v1 = append(v1, v5)
                        _ = fi16
                    }
                }
                var v17 int8 = 42
                v5 = v5
                _ = fi7
                _ = v17
            }
        case 2147483647:
            _ = len(v1)
        default:
            fmt.Println(v5)
        }
        _ = v6
    } else {
        v4 -= v4
        v5 = (v5 + ("hello" + v5))
    }
    var v18 uint64 = 42
    if (!(255 > 0)) {
        v5 = "hello"
        if false {
            switch 9223372036854775807 {
            case -738:
                var v19 int8 = int8(v2)
                _ = v19
            }
        } else {
            if false {
                for fi20 := 0; fi20 < 15; fi20++ {
                    if false {
                        switch (-(fi20)) {
                        case 32767:
                            v21 := GenS2{F0: -9223372036854775808}
                            _ = v21
                        case -830:
                            var v22 float32 = 1e10
                            _ = v22
                        }
                        switch 9223372036854775807 {
                        case -1:
                            v23 := GenS2{F0: 9223372036854775807}
                            _ = v23
                        }
                        _ = len(v1)
                    } else {
                        for fi24 := 0; fi24 < 5; fi24++ {
                            v5 = (((("hello" + ("foo bar" + v5)) + v5) + "\n\t") + ("foo bar" + v5))
                            v4 -= (v4 % ((v4) | 1))
                            _ = fi24
                        }
                        for fi25 := 0; fi25 < 7; fi25++ {
                            var v26 uint64 = v18
                            _ = fi25
                            _ = v26
                        }
                    }
                    _ = fi20
                }
            }
            fmt.Println(v4, v2)
        }
        for fi27 := 0; fi27 < 18; fi27++ {
            delete(v3, "hello")
            if (!(v2 >= (v2 & v2))) {
                if (289 == 42) {
                    v2 += 534
                    fmt.Println(fi27)
                    var v28 uint32 = v4
                    _ = v28
                }
                var v29 uint64 = 951
                _ = v29
            } else {
                if true {
                    for fi30 := 0; fi30 < 12; fi30++ {
                        v2 += (v2 | v2)
                        if len(v1) > 0 {
                            v1[0] = v5
                        }
                        _ = fi30
                    }
                    var v31 uint = 4294967295
                    _ = v31
                } else {
                    if false {
                        v18 = 4294967295
                    }
                    v32 := GenS1{F0: true, F1: -1}
                    fmt.Println(v2, v5)
                    _ = v32
                }
                var v33 int16 = (v2 ^ (v2 | (v2 + (v2 & (v2 + (v2 ^ v2))))))
                var v34 uint64 = 42
                _ = v33
                _ = v34
            }
            _ = fi27
        }
    } else {
        var v35 int16 = v2
        v36 := GenS1{F0: ((v4 - 2147483648) == v4), F1: 42}
        v37 := GenS3{F0: 1}
        _ = v35
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 8; fi38++ {
        v39 := GenS3{F0: 42}
        _ = fi38
        _ = v39
    }
    fmt.Println(v5, v18)
    switch 42 {
    case 451:
        v40 := GenS1{F0: (966 > 0), F1: -696}
        _ = v40
    case 32768:
        var v41 bool = (v5 == v5)
        _ = v41
    default:
        var v42 float64 = 0.0
        _ = v42
    }
    v2 = v2
    var v43 int = 994
    for fi44 := 0; fi44 < 12; fi44++ {
        var v45 float64 = float64(((0.0 / float32(((((-2.25 * 1e10) * 1.5) - (3.14159 + float64(1.5))) / (1.5 + 1.5)))) * 3.14159))
        if (1 != 42) {
            var v46 uint32 = v4
            if (true && (((127 > 42) || true) && (true || (!false)))) {
                if true {
                    v5 = v5
                    v47 := GenS4{F0: (v4 & v46), F1: (false || ("\n\t" == "")), F2: (v2 << ((18446744073709551615) & 63))}
                    _ = v47
                } else {
                    v46 = v46
                }
            }
            var v48 uint8 = 129
            _ = v46
            _ = v48
        }
        if len(v1) > 0 {
            v1[0] = (v5 + v5)
        }
        _ = fi44
        _ = v45
    }
    var v49 int16 = (v2 << ((0) & 63))
    v5 = v5
    fmt.Println(v4, v4, v4)
    if (v18 >= v18) {
        var v50 uint64 = (v18 % ((v18) | 1))
        _ = v50
    }
    for fi51 := 0; fi51 < 1; fi51++ {
        v4 = v4
        fmt.Println(v2, v18, v2)
        _ = fi51
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v49
    _ = v5
}
