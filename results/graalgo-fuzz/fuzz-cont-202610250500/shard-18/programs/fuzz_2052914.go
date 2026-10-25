package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 string; F1 uint; F2 uint32 }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 uint64; F1 int; F2 int }

func main() {
    v1 := []int{}
    v2 := []uint32{305, 65535, 29, 2}
    var v3 float32 = ((0.1 + 1e10) - 1.5)
    var v4 uint64 = 1
    var v5 uint64 = 1
    v6 := []string{"\n\t", ("" + (("a" + "hello") + "")), ("\n\t" + "hello")}
    fmt.Println(v3, v3, v3)
    var v7 int8 = 100
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch "a" {
    case "a":
        v1 = append(v1, -9223372036854775808)
    case "\n\t":
        v3 = -2.25
    }
    var v8 int64 = -1
    if len(v1) > 0 {
        v1[0] = -1
    }
    var v9 uint16 = 255
    fmt.Println(v7, v5, v5)
    v3 = ((v3 / 1.5) + -2.25)
    fmt.Println(v5, v4)
    if true {
        v10 := GenS1{F0: -623}
        _ = v10
    } else {
        if (false && ("foo bar" > "foo bar")) {
            v4 = (^(v4))
            var v11 uint = 741
            var v12 int8 = 42
            _ = v11
            _ = v12
        }
        var v13 uint = 65535
        _ = v13
    }
    v14 := GenS1{F0: 9223372036854775807}
    switch 42 {
    case 9223372036854775807:
        if true {
            fmt.Println(v4, v5, v8)
            v9 = v9
        }
    case 2147483647:
        v9 = (v9 / ((v9) | 1))
    case -690:
        if (((1.5 * -2.25) + (((1e10 * -2.25) + 1e10) - 0.1)) > 1.5) {
            for fi15 := 0; fi15 < 19; fi15++ {
                var v16 uint32 = 0
                v17 := GenS2{F0: ("hello" + (("foo bar" + ("hello" + "a")) + "")), F1: (v9 >> ((843) & 63))}
                _ = fi15
                _ = v16
                _ = v17
            }
            if (!true) {
                v18 := GenS1{F0: -49}
                _ = len(v1)
                _ = v18
            }
            var v19 uint = 125
            _ = v19
        } else {
            for fi20 := 0; fi20 < 11; fi20++ {
                if (255 >= 0) {
                    switch ("\n\t" + "hello") {
                    case "hello":
                        if len(v2) > 0 {
                            v2[0] = 2147483647
                        }
                    case "\n\t":
                        v21 := GenS3{F0: "", F1: 226, F2: 0}
                        _ = v21
                    case "foo bar":
                        v22 := GenS4{F0: (v7 >> ((65535) & 63)), F1: 0}
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 12; fi23++ {
                        if false {
                            v7 += v7
                            v24 := GenS3{F0: "", F1: 894, F2: 976}
                            switch ("foo bar" + "") {
                            case "a":
                                v25 := GenS1{F0: 128}
                                _ = v25
                            case "foo bar":
                                v8 = -504
                            case "hello":
                                fmt.Println(fi20, fi23)
                            }
                            _ = v24
                        }
                        _ = fi23
                    }
                }
                switch ("hello" + "hello") {
                case "a":
                    fmt.Println(v5)
                default:
                    for fi26 := 0; fi26 < 15; fi26++ {
                        fmt.Println(v7)
                        _ = fi26
                    }
                }
                for fi27 := 0; fi27 < 12; fi27++ {
                    if false {
                        v8 = (v8 % ((v8) | 1))
                        var v28 float32 = v3
                        var v29 uint = 825
                        if true {
                            if ((^(v9)) != (v9 / (((v9 % (((v9 + (^(v9)))) | 1))) | 1))) {
                                v5 -= v5
                                v30 := GenS5{F0: v5, F1: (fi27 % ((fi20) | 1)), F2: fi20}
                                _ = v30
                            }
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                        _ = v28
                        _ = v29
                    } else {
                        v31 := GenS2{F0: "a", F1: v9}
                        switch "\n\t" {
                        case "hello":
                            switch (-(fi20)) {
                            case 32767:
                                var v32 bool = (false || false)
                                _ = v32
                            case 127:
                                if len(v1) > 0 {
                                    v1[0] = -1
                                }
                            case -22:
                                var v33 int16 = 1000
                                _ = v33
                            default:
                                if len(v2) > 0 {
                                    v2[0] = 2
                                }
                            }
                        case "\n\t":
                            _ = len(v1)
                        default:
                            var v34 uint16 = 2
                            _ = v34
                        }
                        v7 = (v7 + v7)
                        _ = v31
                    }
                    _ = fi27
                }
                _ = fi20
            }
            v5 -= v4
            _ = len(v6)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
