package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint64 }
type GenS2 struct { F0 int; F1 uint64 }

func main() {
    var v1 bool = (32767 > 461)
    var v2 float32 = float32(-2.25)
    var v3 float32 = 1e10
    var v4 int64 = -87
    var v5 bool = v1
    var v6 string = "foo bar"
    var v7 string = v6
    v8 := []int{2147483647, -2147483648, 1}
    var v9 int8 = 70
    _ = len(v8)
    for fi10 := 0; fi10 < 12; fi10++ {
        for fi11 := 0; fi11 < 7; fi11++ {
            if (100 != v9) {
                v4 = (^(v4))
            }
            if len(v8) > 0 {
                v8[0] = (fi10 + (-(fi11)))
            }
            fmt.Println(fi10, v3)
            _ = fi11
        }
        var v12 int8 = v9
        _ = fi10
        _ = v12
    }
    var v13 uint32 = 183
    if false {
        v14 := GenS1{F0: v7, F1: "", F2: 18446744073709551615}
        _ = v14
    }
    fmt.Println(v7, v2)
    if v1 {
        var v15 int64 = (v4 - v4)
        if true {
            for fi16 := 0; fi16 < 3; fi16++ {
                if (-1 != 1) {
                    var v17 float32 = 1e10
                    for fi18 := 0; fi18 < 4; fi18++ {
                        v19 := GenS1{F0: v6, F1: v7, F2: 18446744073709551615}
                        _ = len(v8)
                        v13 += v13
                        _ = fi18
                        _ = v19
                    }
                    if ((2 > 128) || (2 >= 2)) {
                        switch ((v7 + v7) + (v6 + ("hello" + v6))) {
                        case "hello":
                            switch fi16 {
                            case 32767:
                                var v20 int16 = 42
                                _ = v20
                            case 2147483647:
                                v15 -= v4
                            case 42:
                                fmt.Println(v15)
                            }
                        case "\n\t":
                            for fi21 := 0; fi21 < 9; fi21++ {
                                v1 = v5
                                _ = len(v8)
                                v22 := GenS1{F0: ((v6 + (v6 + (v7 + (v7 + ((v7 + v6) + "\n\t"))))) + v7), F1: "\n\t", F2: uint64(v13)}
                                v23 := GenS1{F0: (("" + ((v7 + (v6 + ((((((v6 + ("foo bar" + (v6 + v6))) + ("a" + "foo bar")) + "a") + (v6 + v6)) + v7) + v7))) + v7)) + "\n\t"), F1: (((("hello" + (v7 + "a")) + v7) + ("\n\t" + "a")) + (v6 + (("" + (((((v7 + v6) + (v7 + v7)) + v6) + ((v6 + (v7 + ("\n\t" + v7))) + "\n\t")) + ("hello" + "\n\t"))) + v6))), F2: 42}
                                var v24 uint32 = 4294967295
                                _ = fi21
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                        default:
                            fmt.Println(v6)
                        }
                        if ("a" >= "hello") {
                            if v1 {
                                if (false && (v5 || (895 >= 595))) {
                                    _ = len(v8)
                                } else {
                                    for fi25 := 0; fi25 < 13; fi25++ {
                                        var v26 uint = 707
                                        var v27 int64 = (v15 << ((v26) & 63))
                                        var v28 int8 = -1
                                        fmt.Println(v28)
                                        if len(v8) >= 2 {
                                            v8 = v8[1:len(v8)]
                                        }
                                        v29 := GenS2{F0: fi16, F1: 600}
                                        _ = fi25
                                        _ = v26
                                        _ = v27
                                        _ = v28
                                        _ = v29
                                    }
                                }
                            } else {
                                var v30 int = fi16
                                v15 -= v15
                                _ = v30
                            }
                            fmt.Println(v1, fi16)
                        }
                    }
                    _ = v17
                } else {
                    var v31 uint64 = 446
                    v32 := GenS2{F0: 1, F1: v31}
                    _ = v31
                    _ = v32
                }
                fmt.Println(v3, v5, v5)
                _ = fi16
            }
        } else {
            var v33 uint64 = 42
            _ = v33
        }
        _ = v15
    } else {
        if v5 {
            switch "a" {
            case "hello":
                v13 -= 42
            case "foo bar":
                if ((0 != 148) || v5) {
                    v5 = true
                    _ = len(v8)
                    v5 = true
                }
            case "":
                for fi34 := 0; fi34 < 12; fi34++ {
                    if (!v1) {
                        v3 = (v3 / ((v2 + v2) * v2))
                    } else {
                        v35 := GenS1{F0: "", F1: ((v7 + "a") + "foo bar"), F2: 2}
                        v36 := GenS2{F0: fi34, F1: 2}
                        v6 = "a"
                        _ = v35
                        _ = v36
                    }
                    _ = fi34
                }
            }
        } else {
            if true {
                _ = len(v8)
            }
        }
        if true {
            if len(v8) > 0 {
                v8[0] = 907
            }
            v7 = v6
            switch 32768 {
            case 128:
                switch "a" {
                case "hello":
                    _ = len(v8)
                default:
                    switch ((v7 + v6) + v7) {
                    case "a":
                        var v37 uint64 = 18446744073709551615
                        _ = v37
                    case "\n\t":
                        var v38 int32 = -2147483648
                        _ = v38
                    }
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
