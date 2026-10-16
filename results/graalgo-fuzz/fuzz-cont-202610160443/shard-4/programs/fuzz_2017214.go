package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32; F1 int8; F2 int }
type GenS6 struct { F0 int16 }

func main() {
    v1 := make(map[string]string)
    v2 := []string{}
    v3 := []uint32{}
    v4 := make(map[int]string)
    var v5 string = ("hello" + "a")
    var v6 int32 = 42
    for fi7 := 0; fi7 < 12; fi7++ {
        fmt.Println(v6, fi7, fi7)
        v8 := GenS1{F0: 255}
        if (42 < 2) {
            v6 = (v6 % (((-(v6))) | 1))
        } else {
            switch fi7 {
            case 4294967295:
                v9 := GenS1{F0: 43}
                _ = v9
            case -409:
                var v10 uint = 36
                _ = v10
            case 0:
                for fi11 := 0; fi11 < 7; fi11++ {
                    if false {
                        var v12 int64 = -9223372036854775808
                        v13 := GenS2{F0: v6, F1: 102}
                        _ = v12
                        _ = v13
                    } else {
                        v14 := GenS1{F0: 42}
                        _ = v14
                    }
                    fmt.Println(fi11)
                    _ = fi11
                }
            default:
                var v15 uint64 = 1
                _ = v15
            }
            for fi16 := 0; fi16 < 5; fi16++ {
                delete(v4, fi16)
                fmt.Println(v5, fi7)
                _ = fi16
            }
            v17 := GenS3{F0: v5}
            _ = v17
        }
        _ = fi7
        _ = v8
    }
    v3 = append(v3, 2147483647)
    if (!true) {
        v18 := v1["\n\t"]
        switch -430 {
        case 32767:
            if false {
                v18 = ((v5 + v18) + ("a" + v5))
                var v19 string = v5
                for fi20 := 0; fi20 < 6; fi20++ {
                    switch v18 {
                    case "a":
                        switch v19 {
                        case "\n\t":
                            for fi21 := 0; fi21 < 14; fi21++ {
                                if len(v2) > 0 {
                                    v2[0] = (("foo bar" + v18) + v5)
                                }
                                _ = fi21
                            }
                        case "hello":
                            v22 := GenS4{F0: 1}
                            _ = v22
                        case "foo bar":
                            v18 = v19
                        default:
                            switch (fi20 / ((fi20) | 1)) {
                            case 427:
                                switch (v19 + "a") {
                                case "a":
                                    v23 := GenS5{F0: 848, F1: 33, F2: fi20}
                                    _ = v23
                                default:
                                    fmt.Println(fi20)
                                }
                            case 0:
                                switch (v18 + v18) {
                                case "a":
                                    if len(v3) > 0 {
                                        v3[0] = 97
                                    }
                                case "hello":
                                    if len(v2) > 0 {
                                        v2[0] = v18
                                    }
                                case "":
                                    var v24 uint = 65535
                                    _ = v24
                                default:
                                    fmt.Println(v6, v6, fi20)
                                }
                            }
                        }
                    }
                    for fi25 := 0; fi25 < 7; fi25++ {
                        v4[0] = v18
                        _ = fi25
                    }
                    _ = fi20
                }
                _ = v19
            }
        case -2147483648:
            var v26 int64 = 42
            _ = v26
        default:
            var v27 int = -810
            _ = v27
        }
        v5 = (("" + "hello") + "foo bar")
        _ = v18
    }
    var v28 string = v5
    for fi29 := 0; fi29 < 2; fi29++ {
        if true {
            var v30 int = (fi29 % ((-520) | 1))
            v2 = append(v2, "\n\t")
            v5 = "foo bar"
            _ = v30
        } else {
            v28 = v28
        }
        _ = fi29
    }
    v31 := GenS6{F0: 0}
    for fi32 := 0; fi32 < 4; fi32++ {
        v5 = v5
        v3 = append(v3, 4294967295)
        _ = fi32
    }
    var v33 int64 = 0
    for fi34 := 0; fi34 < 9; fi34++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch v28 {
        case "\n\t":
            _ = len(v1)
        }
        _ = fi34
    }
    if (4294967295 <= 2147483648) {
        switch -9223372036854775808 {
        case 0:
            v28 = "\n\t"
        }
        v28 = (("foo bar" + "hello") + v28)
        if (v5 >= "hello") {
            v28 = (v5 + "hello")
        }
    } else {
        v5 = v5
        v3 = append(v3, 2)
    }
    delete(v1, v5)
    if true {
        fmt.Println(v33, v28)
    }
    if false {
        var v35 uint = 18446744073709551615
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
