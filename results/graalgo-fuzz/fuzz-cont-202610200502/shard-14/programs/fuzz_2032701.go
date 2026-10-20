package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int64; F2 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint; F1 string; F2 string }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 int16 }

func main() {
    var v1 bool = (303 < 0)
    var v2 int16 = -998
    v3 := []int{}
    var v4 float64 = float64(-2.25)
    v5 := []string{"foo bar", "\n\t"}
    var v6 bool = (!v1)
    for fi7 := 0; fi7 < 6; fi7++ {
        v6 = v6
        for fi8 := 0; fi8 < 4; fi8++ {
            v1 = true
            if (fi8 <= (fi8 & 9223372036854775807)) {
                if v6 {
                    var v9 uint64 = 983
                    _ = v9
                }
                v1 = true
                if (42 == 9223372036854775807) {
                    switch ("a" + "") {
                    case "hello":
                        v3 = append(v3, (fi8 - fi8))
                    default:
                        fmt.Println(fi7, fi8)
                    }
                    if (!v6) {
                        if v6 {
                            var v10 uint64 = 4294967295
                            v1 = true
                            _ = v10
                        }
                        fmt.Println(v4, v2)
                        v11 := GenS1{F0: 620}
                        _ = v11
                    } else {
                        if true {
                            switch fi8 {
                            case 128:
                                if len(v3) > 0 {
                                    v3[0] = 499
                                }
                            }
                            v4 += v4
                        } else {
                            v12 := GenS1{F0: 231}
                            _ = v12
                        }
                        v5 = append(v5, "hello")
                        v5 = append(v5, "hello")
                    }
                    v13 := GenS1{F0: 17}
                    _ = v13
                } else {
                    var v14 uint16 = 255
                    var v15 uint32 = 2
                    v4 = 0.1
                    _ = v14
                    _ = v15
                }
            }
            if true {
                fmt.Println(v1, v1)
            } else {
                fmt.Println(v1, v6, v6)
                if len(v5) > 0 {
                    v5[0] = ("a" + "foo bar")
                }
                var v16 uint16 = 551
                _ = v16
            }
            _ = fi8
        }
        _ = fi7
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        switch ("hello" + "foo bar") {
        case "":
            v18 := GenS1{F0: 42}
            _ = v18
        case "hello":
            v4 -= -2.25
        case "foo bar":
            var v19 uint32 = 70
            _ = v19
        }
        fmt.Println(v2, fi17)
        _ = fi17
    }
    if (!(!(1.5 > 3.14159))) {
        v20 := GenS2{F0: 4294967295, F1: 42, F2: -1}
        var v21 uint64 = 2
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        v23 := GenS3{F0: fi22}
        if v6 {
            if ((1 >= 1) || (0.1 == v4)) {
                v6 = ("a" != "foo bar")
                v1 = (92 == 42)
            }
        }
        _ = fi22
        _ = v23
    }
    var v24 uint16 = 32767
    if v1 {
        var v25 int32 = -1
        switch ("hello" + "") {
        case "\n\t":
            fmt.Println(v4)
        case "":
            v25 -= (v25 | v25)
        }
        _ = v25
    } else {
        v3 = append(v3, -2147483648)
        for fi26 := 0; fi26 < 15; fi26++ {
            for fi27 := 0; fi27 < 6; fi27++ {
                v1 = (74 <= 0)
                switch "foo bar" {
                case "\n\t":
                    v1 = true
                default:
                    if true {
                        v1 = (18446744073709551615 < 4294967295)
                        fmt.Println(fi26, v24, v6)
                    } else {
                        var v28 bool = v6
                        _ = v28
                    }
                }
                _ = fi27
            }
            v29 := GenS2{F0: 94, F1: 892, F2: 0}
            _ = fi26
            _ = v29
        }
        v30 := GenS4{F0: 65535, F1: "foo bar", F2: "\n\t"}
        _ = v30
    }
    v31 := GenS4{F0: 0, F1: ("" + ""), F2: "\n\t"}
    v6 = v6
    fmt.Println(v2, v6)
    switch "a" {
    case "foo bar":
        if (9223372036854775807 != 1) {
            switch int(v24) {
            case -9223372036854775808:
                for fi32 := 0; fi32 < 6; fi32++ {
                    for fi33 := 0; fi33 < 10; fi33++ {
                        var v34 uint64 = 2
                        _ = fi33
                        _ = v34
                    }
                    v24 += v24
                    _ = fi32
                }
            case -47:
                v6 = false
            default:
                switch 42 {
                case -1:
                    if v6 {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        if true {
                            var v35 int8 = 0
                            fmt.Println(v1, v24, v35)
                            _ = v35
                        }
                    } else {
                        var v36 int8 = 0
                        switch ("hello" + "foo bar") {
                        case "":
                            fmt.Println(v2)
                        default:
                            v37 := GenS5{F0: 2147483647, F1: 42}
                            _ = v37
                        }
                        if (292 == 2147483648) {
                            v6 = v1
                            v3 = append(v3, -1)
                        } else {
                            var v38 int8 = (v36 >> ((1) & 63))
                            var v39 uint16 = v24
                            _ = v38
                            _ = v39
                        }
                        _ = v36
                    }
                default:
                    var v40 uint64 = 352
                    _ = v40
                }
            }
            if len(v3) > 0 {
                v3[0] = 32768
            }
            v5 = append(v5, "\n\t")
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v41 uint32 = 614
            fmt.Println(v6, v1)
            _ = v41
        }
    default:
        _ = len(v3)
    }
    var v42 uint = 172
    var v43 string = (("foo bar" + ("" + (("a" + (("a" + "\n\t") + "hello")) + "foo bar"))) + "a")
    v6 = v1
    var v44 int64 = -1
    _ = len(v3)
    var v45 uint64 = 874
    var v46 uint = (v42 << ((260) & 63))
    if (!(42 == 128)) {
        for fi47 := 0; fi47 < 8; fi47++ {
            v48 := GenS6{F0: int16(v42)}
            v44 = v44
            fmt.Println(v46, v45, v4)
            _ = fi47
            _ = v48
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v31
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
