package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int8; F1 uint8; F2 int64 }
type GenS6 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS7 struct { F0 int64; F1 int8 }
func genFn0() int64 {
    var v7 int8 = 100
    for fi8 := 0; fi8 < 12; fi8++ {
        fmt.Println(fi8, fi8)
        var v9 uint8 = 128
        _ = fi8
        _ = v9
    }
    v10 := []int{-694, -341}
    fmt.Println(v7)
    v7 = v7
    _ = v10
    _ = v7
    return -1
}
func genFn1(p0 string) string {
    v11 := []int{118, 9223372036854775807, 42, 259}
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 uint32 = 2147483647
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        p0 = p0
        fmt.Println(p0, p0)
        _ = fi14
    }
    p0 = "\n\t"
    var v15 int16 = -32768
    _ = p0
    _ = v11
    _ = v15
    return (p0 + ("a" + "hello"))
}

func main() {
    var v1 uint16 = 42
    v2 := []string{("\n\t" + "hello"), "", ("" + "\n\t"), ""}
    var v3 int8 = 42
    var v4 int16 = 0
    v5 := make(map[string]int)
    var v6 int8 = -54
    var v16 uint64 = 4294967295
    var v17 int8 = v3
    for fi18 := 0; fi18 < 5; fi18++ {
        switch (((("a" + (("foo bar" + ("a" + "a")) + (("" + "hello") + "\n\t"))) + ((((("a" + "hello") + "") + "a") + "\n\t") + "\n\t")) + ("hello" + ("a" + "\n\t"))) + "a") {
        case "hello":
            _ = len(v2)
        case "a":
            v19 := GenS1{F0: 763, F1: v4}
            _ = v19
        }
        _ = fi18
    }
    if false {
        switch 127 {
        case 42:
            if true {
                for fi20 := 0; fi20 < 8; fi20++ {
                    var v21 bool = true
                    v22 := v5[("\n\t" + ("\n\t" + "hello"))]
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                v23 := v5[("" + ("hello" + "a"))]
                _ = v23
            } else {
                fmt.Println(v4, v4)
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
                v16 -= (v16 / ((222) | 1))
            }
        case -146:
            if true {
                if true {
                    if (!(v16 < (^(v16)))) {
                        v24 := GenS2{F0: 65535}
                        _ = v24
                    } else {
                        v6 = (v6 - v3)
                        switch "hello" {
                        case "hello":
                            for fi25 := 0; fi25 < 9; fi25++ {
                                v26 := GenS2{F0: v1}
                                v6 -= v3
                                var v27 int8 = v17
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                        default:
                            var v28 uint8 = 0
                            _ = v28
                        }
                        v29 := v5[""]
                        _ = v29
                    }
                }
                switch -2147483648 {
                case -2147483648:
                    v30 := GenS1{F0: 34, F1: int16(v16)}
                    _ = v30
                case 0:
                    for fi31 := 0; fi31 < 13; fi31++ {
                        v16 = v16
                        if (2 < 0) {
                            var v32 uint16 = v1
                            if ("a" <= "a") {
                                v16 += 309
                                fmt.Println(v3)
                            } else {
                                for fi33 := 0; fi33 < 2; fi33++ {
                                    switch "foo bar" {
                                    case "foo bar":
                                        v6 -= v17
                                    }
                                    v34 := GenS3{F0: (fi31 >> ((42) & 63))}
                                    v4 -= 0
                                    _ = fi33
                                    _ = v34
                                }
                            }
                            _ = v32
                        }
                        _ = fi31
                    }
                case 127:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                v35 := GenS1{F0: 385, F1: 0}
                _ = v35
            }
        default:
            v17 = 42
        }
        var v36 bool = (3.14159 >= 0.0)
        _ = v36
    } else {
        v1 = (^(v1))
    }
    v37 := GenS4{F0: true}
    v17 = 127
    fmt.Println(v3, v4, v1)
    switch -289 {
    case -768:
        if true {
            fmt.Println(v4, v17)
        }
    default:
        switch 128 {
        case 128:
            for fi38 := 0; fi38 < 4; fi38++ {
                v39 := GenS2{F0: (v1 & 1)}
                for fi40 := 0; fi40 < 6; fi40++ {
                    switch ("" + "foo bar") {
                    case "\n\t":
                        v3 = (-(v6))
                    }
                    _ = fi40
                }
                v1 = (v1 / ((v1) | 1))
                _ = fi38
                _ = v39
            }
        default:
            fmt.Println(v4, v16, v17)
        }
    }
    var v41 int8 = v6
    if (true || false) {
        if false {
            fmt.Println(v17)
        } else {
            fmt.Println(v6, v3, v1)
            if false {
                v42 := GenS5{F0: v3, F1: 42, F2: 1}
                _ = v42
            } else {
                var v43 uint32 = 1
                for fi44 := 0; fi44 < 17; fi44++ {
                    var v45 uint = 18446744073709551615
                    var v46 int16 = v4
                    _ = fi44
                    _ = v45
                    _ = v46
                }
                _ = v43
            }
            v3 += v6
        }
        if false {
            var v47 int = 9223372036854775807
            _ = v47
        } else {
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
            if false {
                var v48 bool = true
                switch 248 {
                case -1:
                    var v49 string = "hello"
                    _ = v49
                case 2147483647:
                    for fi50 := 0; fi50 < 10; fi50++ {
                        fmt.Println(v48, v48)
                        if (4294967295 == 2147483647) {
                            v6 += 30
                            switch "foo bar" {
                            case "hello":
                                var v51 int8 = (v6 >> ((653) & 63))
                                _ = v51
                            case "a":
                                v52 := GenS6{F0: v16, F1: -536, F2: 127}
                                _ = v52
                            }
                        } else {
                            if false {
                                _ = len(v5)
                                fmt.Println(v3, v17, v6)
                            }
                            for fi53 := 0; fi53 < 2; fi53++ {
                                if v48 {
                                    var v54 float64 = 0.1
                                    _ = v54
                                }
                                switch 9223372036854775807 {
                                case 2147483647:
                                    var v55 int64 = 9223372036854775807
                                    _ = v55
                                default:
                                    v48 = ((!false) && true)
                                }
                                _ = fi53
                            }
                        }
                        v5[""] = fi50
                        _ = fi50
                    }
                default:
                    v41 += v17
                }
                _ = v48
            } else {
                fmt.Println(v4)
                switch 846 {
                case 4294967295:
                    v16 -= v16
                case 2147483647:
                    if (-1 <= (v4 / ((1) | 1))) {
                        v56 := GenS7{F0: 0, F1: 1}
                        v4 = v4
                        _ = v56
                    }
                default:
                    v16 -= (^(v16))
                }
            }
        }
        v41 += -1
    }
    v57 := GenS5{F0: (v17 | 1), F1: 19, F2: genFn0()}
    v3 += (v6 >> ((42) & 63))
    if (813 >= 42) {
        for fi58 := 0; fi58 < 10; fi58++ {
            var v59 int16 = (v4 >> ((65535) & 63))
            fmt.Println(v59, v6)
            v60 := GenS5{F0: (v17 & v6), F1: 0, F2: 1}
            _ = fi58
            _ = v59
            _ = v60
        }
        if true {
            delete(v5, ("" + (((("a" + "foo bar") + "") + ((("hello" + ("\n\t" + ("hello" + (("a" + ("a" + ((((("hello" + "hello") + "") + ("foo bar" + ("hello" + "hello"))) + "foo bar") + "\n\t"))) + "foo bar")))) + (("\n\t" + "\n\t") + "hello")) + "a")) + ("" + ""))))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v41
    _ = v5
    _ = v57
    _ = v6
}
