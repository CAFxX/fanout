package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint; F2 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 int64; F1 uint64; F2 int8 }
type GenS6 struct { F0 uint8; F1 int32 }
func genFn0() string {
    var v4 int64 = 1
    var v5 uint16 = 42
    var v6 int16 = 1000
    var v7 bool = false
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return (("" + ("\n\t" + "hello")) + "a")
}
func genFn1(p0 uint8, p1 uint32, p2 uint64) int32 {
    var v8 uint8 = p0
    var v9 int8 = 28
    var v10 uint8 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return -1
}

func main() {
    v1 := []string{"hello"}
    var v2 int16 = 1000
    var v3 int = 548
    v11 := GenS1{F0: 773}
    switch ("a" + ("a" + (("\n\t" + "foo bar") + ""))) {
    case "a":
        var v12 uint32 = 406
        _ = v12
    case "\n\t":
        switch (((("" + ("" + ("foo bar" + ("hello" + ("hello" + ("\n\t" + "a")))))) + "a") + "\n\t") + ("" + "")) {
        case "hello":
            switch (v3 % ((v3) | 1)) {
            case 32767:
                for fi13 := 0; fi13 < 18; fi13++ {
                    v3 -= (fi13 * (fi13 + (^(fi13))))
                    _ = fi13
                }
            case 0:
                for fi14 := 0; fi14 < 19; fi14++ {
                    v15 := GenS2{F0: 60, F1: 65535, F2: 407}
                    for fi16 := 0; fi16 < 4; fi16++ {
                        var v17 float64 = 1.5
                        v1 = append(v1, "")
                        _ = fi16
                        _ = v17
                    }
                    _ = fi14
                    _ = v15
                }
            case 2147483647:
                v3 = v3
            }
        case "\n\t":
            v3 = (v3 | 1)
        case "foo bar":
            v1 = append(v1, "a")
        default:
            var v18 string = (((((("hello" + "hello") + "a") + "") + "hello") + ((("a" + "hello") + "a") + "a")) + ("foo bar" + ("\n\t" + "a")))
            _ = v18
        }
    case "":
        for fi19 := 0; fi19 < 5; fi19++ {
            fmt.Println(v3, fi19, v2)
            var v20 float32 = 1.5
            var v21 uint16 = 696
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    v22 := GenS3{F0: (!true)}
    v2 = (v2 | (v2 - 0))
    fmt.Println(v3)
    var v23 int32 = 659
    fmt.Println(v23, v23, v23)
    v24 := GenS1{F0: 42}
    for fi25 := 0; fi25 < 5; fi25++ {
        switch ("a" + "") {
        case "a":
            for fi26 := 0; fi26 < 11; fi26++ {
                var v27 int16 = (v2 - 55)
                var v28 string = "hello"
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi26
                _ = v27
                _ = v28
            }
        default:
            if (-371 == -228) {
                for fi29 := 0; fi29 < 4; fi29++ {
                    v1 = append(v1, (("foo bar" + "foo bar") + "a"))
                    v2 = (^(v2))
                    v30 := GenS4{F0: (v23 << ((0) & 63)), F1: 65535}
                    _ = fi29
                    _ = v30
                }
                switch "" {
                case "a":
                    if true {
                        for fi31 := 0; fi31 < 6; fi31++ {
                            if true {
                                v3 -= (fi31 << ((975) & 63))
                            } else {
                                v3 = v3
                                if false {
                                    v32 := GenS3{F0: ((42 <= 2) || (false && false))}
                                    v23 -= v23
                                    switch (fi31 - v3) {
                                    case 9223372036854775807:
                                        if false {
                                            var v33 int64 = 23
                                            if len(v1) > 0 {
                                                v1[0] = "a"
                                            }
                                            var v34 float64 = 1.5
                                            _ = v33
                                            _ = v34
                                        } else {
                                            v3 -= -2147483648
                                        }
                                    case -2147483648:
                                        v35 := GenS2{F0: 0, F1: 65535, F2: 623}
                                        _ = v35
                                    }
                                    fmt.Println(v2)
                                    _ = v32
                                } else {
                                    if false {
                                        v36 := GenS2{F0: -1, F1: 18446744073709551615, F2: 2147483647}
                                        if len(v1) > 0 {
                                            v1[0] = (("hello" + ("foo bar" + "hello")) + "")
                                        }
                                        _ = v36
                                    } else {
                                        fmt.Println(v2)
                                    }
                                    v1 = append(v1, "a")
                                    var v37 float32 = -2.25
                                    _ = v37
                                }
                            }
                            var v38 string = "foo bar"
                            _ = fi31
                            _ = v38
                        }
                        for fi39 := 0; fi39 < 19; fi39++ {
                            var v40 float64 = 1.5
                            if true {
                                v41 := GenS4{F0: v23, F1: 42}
                                _ = v41
                            }
                            _ = fi39
                            _ = v40
                        }
                        v42 := GenS3{F0: true}
                        _ = v42
                    } else {
                        v43 := GenS5{F0: 42, F1: 18446744073709551615, F2: 1}
                        _ = v43
                    }
                default:
                    v2 = v2
                }
                switch ("hello" + "a") {
                case "hello":
                    v23 = -1
                default:
                    if false {
                        v23 += (^(v23))
                        v3 = 653
                    } else {
                        v2 -= (v2 / ((v2) | 1))
                    }
                }
            }
        }
        _ = fi25
    }
    switch "" {
    case "hello":
        var v44 int64 = 42
        _ = v44
    }
    switch "" {
    case "hello":
        switch "\n\t" {
        case "foo bar":
            for fi45 := 0; fi45 < 18; fi45++ {
                v46 := GenS6{F0: 127, F1: v23}
                var v47 uint32 = 65535
                var v48 string = ""
                _ = fi45
                _ = v46
                _ = v47
                _ = v48
            }
        }
    case "foo bar":
        if false {
            fmt.Println(v2, v3, v23)
            if true {
                for fi49 := 0; fi49 < 14; fi49++ {
                    if true {
                        v50 := GenS3{F0: (false || (false && ((true || true) && (1 >= 0))))}
                        switch -1 {
                        case -2147483648:
                            fmt.Println(v23, v23, v23)
                        case 4294967295:
                            var v51 uint8 = 2
                            _ = v51
                        }
                        _ = v50
                    }
                    _ = fi49
                }
                var v52 int16 = (v2 - 0)
                _ = v52
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v23 -= v23
            }
        } else {
            v23 -= v23
            var v53 bool = (!true)
            _ = v53
        }
    case "":
        if len(v1) > 0 {
            v1[0] = (("hello" + "") + "hello")
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
