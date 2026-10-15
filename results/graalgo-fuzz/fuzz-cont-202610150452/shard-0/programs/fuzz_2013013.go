package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[string]string)
    var v2 int64 = -1
    var v3 uint8 = 127
    var v4 float32 = ((3.14159 / 1.5) * 3.14159)
    v5 := []string{"foo bar", "\n\t", "hello", (("\n\t" + "foo bar") + (("foo bar" + ("hello" + "")) + ""))}
    _ = len(v1)
    v4 -= (v4 / float32(3.14159))
    for fi6 := 0; fi6 < 13; fi6++ {
        if ((true || false) || (-32768 != 1000)) {
            v3 = (v3 | 0)
        } else {
            v4 = (v4 * v4)
            if (false || true) {
                var v7 bool = (!("a" < "\n\t"))
                _ = v7
            }
            switch fi6 {
            case 127:
                for fi8 := 0; fi8 < 18; fi8++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v4 += float32((-2.25 + 1.5))
                    _ = fi8
                }
            case 42:
                var v9 uint32 = 2147483647
                _ = v9
            default:
                for fi10 := 0; fi10 < 10; fi10++ {
                    _ = len(v5)
                    var v11 uint = 42
                    _ = fi10
                    _ = v11
                }
            }
        }
        if len(v5) > 0 {
            v5[0] = "a"
        }
        var v12 string = "foo bar"
        _ = fi6
        _ = v12
    }
    switch "" {
    case "a":
        switch "hello" {
        case "\n\t":
            if true {
                var v13 float32 = v4
                if true {
                    v1["\n\t"] = "hello"
                    var v14 float32 = (v4 - v4)
                    _ = v14
                }
                switch 159 {
                case 42:
                    for fi15 := 0; fi15 < 19; fi15++ {
                        fmt.Println(v13, v2)
                        if (42 >= 829) {
                            v3 += (v3 * v3)
                            if false {
                                fmt.Println(v2, v13, fi15)
                            } else {
                                var v16 uint16 = 42
                                _ = v16
                            }
                        }
                        if true {
                            switch (fi15 << ((18446744073709551615) & 63)) {
                            case 42:
                                var v17 uint64 = 447
                                _ = v17
                            case -9223372036854775808:
                                v3 = v3
                            case 1:
                                var v18 int32 = 0
                                _ = v18
                            default:
                                if len(v5) > 0 {
                                    v5[0] = "a"
                                }
                            }
                            if false {
                                v4 += v4
                                switch fi15 {
                                case 0:
                                    delete(v1, "\n\t")
                                case 737:
                                    fmt.Println(v2, v2, fi15)
                                default:
                                    for fi19 := 0; fi19 < 19; fi19++ {
                                        switch ("a" + "\n\t") {
                                        case "hello":
                                            var v20 string = (("hello" + "a") + ("foo bar" + "hello"))
                                            _ = v20
                                        case "foo bar":
                                            v13 += float32(0.0)
                                        }
                                        var v21 uint = 1
                                        _ = fi19
                                        _ = v21
                                    }
                                }
                                for fi22 := 0; fi22 < 6; fi22++ {
                                    fmt.Println(fi22, v13)
                                    _ = fi22
                                }
                            } else {
                                var v23 int32 = 0
                                fmt.Println(fi15, v23)
                                var v24 uint16 = 42
                                _ = v23
                                _ = v24
                            }
                        } else {
                            for fi25 := 0; fi25 < 12; fi25++ {
                                _ = len(v5)
                                v26 := GenS1{F0: fi25, F1: 0}
                                v27 := GenS1{F0: fi15, F1: -1}
                                var v28 float32 = (v13 / (v4 / v13))
                                _ = fi25
                                _ = v26
                                _ = v27
                                _ = v28
                            }
                            for fi29 := 0; fi29 < 5; fi29++ {
                                switch ("a" + "foo bar") {
                                case "foo bar":
                                    v30 := GenS1{F0: (fi15 >> ((682) & 63)), F1: 758}
                                    _ = v30
                                case "\n\t":
                                    _ = len(v5)
                                }
                                _ = len(v5)
                                var v31 int64 = -414
                                _ = fi29
                                _ = v31
                            }
                        }
                        _ = fi15
                    }
                case 646:
                    v13 = (0.1 - 1.5)
                case 1:
                    switch -14 {
                    case 2147483647:
                        if false {
                            delete(v1, "\n\t")
                            fmt.Println(v2, v3, v13)
                            v3 -= (v3 << ((18446744073709551615) & 63))
                        }
                    case -321:
                        fmt.Println(v3, v2)
                    case 127:
                        var v32 int16 = 959
                        _ = v32
                    }
                }
                _ = v13
            } else {
                var v33 int16 = -867
                _ = v33
            }
        case "hello":
            var v34 int = 0
            _ = v34
        default:
            v35 := GenS2{F0: uint(v3)}
            _ = v35
        }
    case "\n\t":
        if true {
            if (!false) {
                var v36 uint = 4294967295
                _ = v36
            }
        }
    default:
        v3 = (v3 << ((1) & 63))
    }
    v37 := GenS1{F0: 42, F1: 0}
    var v38 int16 = 42
    fmt.Println(v4, v38)
    var v39 int64 = -9223372036854775808
    v5 = append(v5, "a")
    v4 -= (v4 + v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
