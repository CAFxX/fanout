package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 int64; F1 uint; F2 int }
type GenS4 struct { F0 int; F1 int64; F2 int16 }

func main() {
    var v1 uint32 = 583
    v2 := []int{-402, 940}
    var v3 uint = 59
    v4 := make(map[string]int)
    if true {
        fmt.Println(v1, v1)
        for fi5 := 0; fi5 < 20; fi5++ {
            v6 := GenS1{F0: "\n\t", F1: (!(true && ((false || false) && (((("a" + ("\n\t" + "a")) + "") + "a") > "hello")))), F2: fi5}
            _ = fi5
            _ = v6
        }
        var v7 int32 = 1
        _ = v7
    }
    v8 := GenS2{F0: 0, F1: true}
    if true {
        fmt.Println(v3, v1)
        var v9 uint64 = 18446744073709551615
        _ = v9
    }
    if (true || (-1 >= 2147483647)) {
        fmt.Println(v3, v3)
        if (!(508 != 42)) {
            switch 42 {
            case 128:
                v1 = (^(v1))
            case 754:
                v10 := GenS3{F0: 0, F1: 0, F2: 32767}
                _ = v10
            case -1:
                if (!false) {
                    var v11 float64 = 0.0
                    _ = v11
                } else {
                    v12 := GenS1{F0: "foo bar", F1: (1 != 127), F2: 4294967295}
                    var v13 int32 = 2147483647
                    var v14 bool = (v3 <= v3)
                    _ = v12
                    _ = v13
                    _ = v14
                }
            default:
                if (100 != 42) {
                    if false {
                        if (!true) {
                            v1 += (v1 >> ((v3) & 63))
                            var v15 uint = v3
                            _ = v15
                        }
                        var v16 uint64 = 18446744073709551615
                        v1 = v1
                        _ = v16
                    } else {
                        fmt.Println(v3, v3)
                        if ((false && false) || (v1 < (v1 + v1))) {
                            for fi17 := 0; fi17 < 18; fi17++ {
                                for fi18 := 0; fi18 < 11; fi18++ {
                                    var v19 int32 = 0
                                    _ = fi18
                                    _ = v19
                                }
                                switch ("hello" + "foo bar") {
                                case "hello":
                                    fmt.Println(fi17)
                                case "a":
                                    for fi20 := 0; fi20 < 8; fi20++ {
                                        v21 := GenS4{F0: fi20, F1: 42, F2: -32768}
                                        v22 := GenS2{F0: -9223372036854775808, F1: (!(1 != -71))}
                                        _ = fi20
                                        _ = v21
                                        _ = v22
                                    }
                                case "foo bar":
                                    v23 := GenS1{F0: (("\n\t" + ("a" + ((("" + ("\n\t" + "foo bar")) + "foo bar") + "\n\t"))) + "hello"), F1: false, F2: int(v1)}
                                    _ = v23
                                default:
                                    v24 := GenS4{F0: -238, F1: -449, F2: 148}
                                    _ = v24
                                }
                                switch "a" {
                                case "a":
                                    v1 = 42
                                case "\n\t":
                                    fmt.Println(v3, fi17, v1)
                                case "hello":
                                    v25 := GenS2{F0: 1, F1: false}
                                    _ = v25
                                default:
                                    switch "foo bar" {
                                    case "foo bar":
                                        switch fi17 {
                                        case -963:
                                            var v26 int = 228
                                            _ = v26
                                        default:
                                            if false {
                                                var v27 string = "hello"
                                                v27 = "foo bar"
                                                var v28 string = ((("hello" + v27) + v27) + v27)
                                                _ = v27
                                                _ = v28
                                            }
                                        }
                                    case "hello":
                                        v1 = v1
                                    case "":
                                        v1 -= v1
                                    }
                                }
                                _ = fi17
                            }
                        }
                        var v29 float64 = 3.14159
                        _ = v29
                    }
                }
            }
            v1 = v1
        }
        var v30 int = 629
        _ = v30
    } else {
        v3 += v3
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v31 uint32 = v1
        _ = v31
    }
    delete(v4, "\n\t")
    if (false && (0.1 > -2.25)) {
        v32 := GenS2{F0: 9223372036854775807, F1: true}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v32
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi33 := 0; fi33 < 12; fi33++ {
        fmt.Println(v1, fi33)
        _ = fi33
    }
    for fi34 := 0; fi34 < 9; fi34++ {
        fmt.Println(v3, v1, v3)
        _ = fi34
    }
    v35 := GenS1{F0: "a", F1: (true || false), F2: 32768}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v36 bool = false
    var v37 int32 = -2147483648
    if v36 {
        var v38 string = ("\n\t" + (((("a" + "\n\t") + "") + "foo bar") + ("hello" + "hello")))
        v39 := GenS3{F0: 0, F1: v3, F2: int(v1)}
        _ = v38
        _ = v39
    } else {
        _ = len(v4)
    }
    v40 := GenS2{F0: -780, F1: v36}
    for fi41 := 0; fi41 < 2; fi41++ {
        for fi42 := 0; fi42 < 2; fi42++ {
            v3 = (v3 | (v3 / (((v3 | v3)) | 1)))
            v37 = 2147483647
            var v43 int8 = 127
            _ = fi42
            _ = v43
        }
        var v44 int32 = (v37 | v37)
        _ = fi41
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v40
    _ = v8
}
