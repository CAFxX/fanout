package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 uint; F2 bool }
type GenS3 struct { F0 uint16; F1 int32; F2 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool; F1 int64 }
type GenS7 struct { F0 uint32; F1 uint8 }
type GenS8 struct { F0 uint16 }

func main() {
    var v1 int8 = 100
    var v2 bool = true
    v3 := []int{37, 0}
    v4 := make(map[string]string)
    fmt.Println(v1)
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := GenS1{F0: 255}
        if true {
            var v7 uint8 = 124
            if true {
                if (!false) {
                    for fi8 := 0; fi8 < 17; fi8++ {
                        v9 := GenS2{F0: 65535, F1: 497, F2: v2}
                        var v10 int32 = 997
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                    switch "a" {
                    case "foo bar":
                        switch "" {
                        case "foo bar":
                            var v11 uint16 = 42
                            _ = v11
                        }
                    case "\n\t":
                        v2 = true
                    case "hello":
                        for fi12 := 0; fi12 < 11; fi12++ {
                            var v13 float64 = -2.25
                            switch (("foo bar" + "foo bar") + "hello") {
                            case "":
                                if true {
                                    fmt.Println(v7, v7)
                                    var v14 int16 = -281
                                    if v2 {
                                        v7 += (v7 + (v7 & (v7 ^ 22)))
                                    } else {
                                        var v15 uint16 = 139
                                        switch 88 {
                                        case -2147483648:
                                            v13 = v13
                                        default:
                                            fmt.Println(v14)
                                        }
                                        switch (-(fi12)) {
                                        case -2147483648:
                                            v16 := v4["a"]
                                            _ = v16
                                        default:
                                            var v17 int = 127
                                            _ = v17
                                        }
                                        v18 := GenS2{F0: 492, F1: 18446744073709551615, F2: (v15 <= (v15 % (((v15 >> ((896) & 63))) | 1)))}
                                        _ = v15
                                        _ = v18
                                    }
                                    _ = v14
                                } else {
                                    fmt.Println(fi5, v1, v1)
                                    var v19 uint = 65535
                                    v4[(("hello" + "foo bar") + "")] = "foo bar"
                                    _ = v19
                                }
                            case "hello":
                                fmt.Println(fi12)
                            default:
                                var v20 float32 = (-2.25 + (0.1 * 3.14159))
                                _ = v20
                            }
                            _ = fi12
                            _ = v13
                        }
                    default:
                        fmt.Println(fi5, fi5, v1)
                    }
                    var v21 int64 = -9223372036854775808
                    _ = v21
                }
            }
            _ = v7
        } else {
            for fi22 := 0; fi22 < 12; fi22++ {
                v23 := GenS3{F0: 0, F1: 1, F2: 1000}
                _ = fi22
                _ = v23
            }
        }
        _ = fi5
        _ = v6
    }
    if v2 {
        var v24 bool = (-32768 < 1)
        v25 := GenS4{F0: 127}
        _ = v24
        _ = v25
    } else {
        v1 = (v1 & v1)
        for fi26 := 0; fi26 < 16; fi26++ {
            if false {
                switch "hello" {
                case "\n\t":
                    v27 := GenS5{F0: 719}
                    _ = v27
                }
            } else {
                v28 := GenS5{F0: 0}
                for fi29 := 0; fi29 < 20; fi29++ {
                    var v30 int = -851
                    v2 = (v2 || v2)
                    _ = fi29
                    _ = v30
                }
                v2 = (!(1000 <= 1))
                _ = v28
            }
            v3 = append(v3, fi26)
            _ = fi26
        }
        fmt.Println(v2)
    }
    v31 := GenS6{F0: true, F1: 1}
    switch ("a" + "foo bar") {
    case "foo bar":
        for fi32 := 0; fi32 < 1; fi32++ {
            if (true && (true || v2)) {
                var v33 float64 = ((1.5 + (1.5 / -2.25)) + 3.14159)
                _ = v33
            } else {
                for fi34 := 0; fi34 < 8; fi34++ {
                    v2 = (((fi32 / ((fi32) | 1)) >= 2147483647) && (1 < 42))
                    if (261 <= 42) {
                        switch (("foo bar" + ((("hello" + ("a" + "foo bar")) + "\n\t") + ("\n\t" + "hello"))) + "\n\t") {
                        case "\n\t":
                            v4["a"] = ("hello" + (("\n\t" + "hello") + "a"))
                        }
                        v1 = (v1 * -88)
                        var v35 string = ("hello" + "foo bar")
                        v1 = (^(v1))
                        _ = v35
                    } else {
                        var v36 int64 = 9223372036854775807
                        var v37 uint16 = 65535
                        v37 = (v37 - 962)
                        _ = v36
                        _ = v37
                    }
                    _ = fi34
                }
            }
            _ = len(v3)
            v3 = append(v3, (fi32 * fi32))
            _ = fi32
        }
    case "":
        if false {
            _ = len(v4)
            var v38 int64 = 9223372036854775807
            var v39 int8 = (v1 + (v1 / ((v1) | 1)))
            _ = v38
            _ = v39
        } else {
            for fi40 := 0; fi40 < 3; fi40++ {
                v2 = (!("\n\t" < ("\n\t" + "hello")))
                v1 = (v1 * -27)
                var v41 uint8 = 182
                _ = fi40
                _ = v41
            }
            v2 = false
        }
    case "hello":
        var v42 int64 = 42
        _ = v42
    default:
        switch 32768 {
        case 0:
            var v43 uint = 585
            _ = v43
        case -2147483648:
            var v44 uint8 = 34
            _ = v44
        case -1:
            v1 -= (v1 >> ((1) & 63))
        default:
            v3 = append(v3, 547)
        }
    }
    var v45 int32 = 42
    fmt.Println(v1, v2, v2)
    v46 := GenS7{F0: 964, F1: 2}
    for fi47 := 0; fi47 < 17; fi47++ {
        for fi48 := 0; fi48 < 5; fi48++ {
            v4[("\n\t" + "a")] = ""
            if (((!v2) && v2) || false) {
                switch (fi47 ^ (fi48 + fi48)) {
                case 32767:
                    var v49 int16 = -1
                    _ = v49
                }
            } else {
                fmt.Println(fi47, v1, v1)
                v50 := v4[""]
                v2 = v2
                _ = v50
            }
            _ = fi48
        }
        _ = fi47
    }
    v51 := GenS8{F0: 42}
    var v52 int32 = v45
    v2 = (!v2)
    fmt.Println(v45)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v45
    _ = v46
    _ = v51
    _ = v52
}
