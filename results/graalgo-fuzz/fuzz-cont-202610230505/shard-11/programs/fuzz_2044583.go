package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS3 struct { F0 string; F1 int16; F2 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16 }
func genFn0() int16 {
    v5 := GenS1{F0: 2147483648, F1: 255}
    switch "a" {
    case "\n\t":
        v6 := GenS1{F0: 2147483648, F1: 65535}
        _ = v6
    case "":
        for fi7 := 0; fi7 < 16; fi7++ {
            if (94 > 1) {
                if (true || true) {
                    for fi8 := 0; fi8 < 10; fi8++ {
                        var v9 uint64 = 42
                        v9 = 388
                        _ = fi8
                        _ = v9
                    }
                    for fi10 := 0; fi10 < 12; fi10++ {
                        var v11 int16 = 0
                        var v12 string = "foo bar"
                        v12 = v12
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    var v13 uint16 = 1
                    _ = v13
                }
            } else {
                fmt.Println(fi7)
                var v14 int8 = 64
                _ = v14
            }
            _ = fi7
        }
    case "a":
        var v15 bool = false
        _ = v15
    default:
        v16 := GenS2{F0: 65535, F1: 0, F2: 469}
        _ = v16
    }
    var v17 uint8 = 1
    _ = v17
    _ = v5
    return 1000
}
func genFn1() int32 {
    var v18 string = ("" + "")
    switch "\n\t" {
    case "a":
        switch v18 {
        case "hello":
            v18 = ("" + (v18 + "a"))
        case "foo bar":
            fmt.Println(v18)
        default:
            if false {
                var v19 uint8 = 238
                var v20 float32 = 3.14159
                v18 = (v18 + "foo bar")
                _ = v19
                _ = v20
            } else {
                for fi21 := 0; fi21 < 17; fi21++ {
                    v18 = (v18 + ((v18 + "hello") + v18))
                    _ = fi21
                }
            }
        }
    case "":
        var v22 int = 32768
        _ = v22
    }
    _ = v18
    return 0
}

func main() {
    var v1 bool = (-31 != -2147483648)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]int)
    v4 := make(map[string]bool)
    v23 := GenS1{F0: 4294967295, F1: 2}
    for fi24 := 0; fi24 < 6; fi24++ {
        for fi25 := 0; fi25 < 11; fi25++ {
            var v26 int64 = 1
            _ = fi25
            _ = v26
        }
        var v27 int8 = -1
        var v28 uint = 18446744073709551615
        _ = fi24
        _ = v27
        _ = v28
    }
    v1 = v1
    for fi29 := 0; fi29 < 19; fi29++ {
        v30 := []int{578, (fi29 % (((-(fi29))) | 1))}
        for fi31 := 0; fi31 < 4; fi31++ {
            v1 = (fi29 <= (fi29 % ((fi31) | 1)))
            v32 := GenS3{F0: ("\n\t" + "hello"), F1: -939, F2: (1 <= 2)}
            _ = fi31
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    for fi33 := 0; fi33 < 7; fi33++ {
        v34 := GenS3{F0: "\n\t", F1: -32768, F2: (v1 || true)}
        if true {
            var v35 uint64 = 18446744073709551615
            for fi36 := 0; fi36 < 12; fi36++ {
                v1 = (-144 <= 1)
                _ = fi36
            }
            var v37 uint64 = 18446744073709551615
            _ = v35
            _ = v37
        } else {
            for fi38 := 0; fi38 < 20; fi38++ {
                var v39 float64 = -2.25
                switch (fi33 % ((fi33) | 1)) {
                case 155:
                    var v40 int8 = 0
                    _ = v40
                case -9223372036854775808:
                    switch ("foo bar" + ("" + "hello")) {
                    case "hello":
                        fmt.Println(fi38, v1, v1)
                        switch -579 {
                        case -224:
                            var v41 uint16 = 0
                            _ = v41
                        case -9223372036854775808:
                            if (!(!(((1 != 1) || false) || true))) {
                                fmt.Println(fi38, fi33)
                            } else {
                                v42 := GenS1{F0: 446, F1: 18446744073709551615}
                                switch 127 {
                                case -9223372036854775808:
                                    switch "\n\t" {
                                    case "\n\t":
                                        v1 = (3.14159 != v39)
                                    case "foo bar":
                                        v43 := GenS3{F0: ("hello" + ""), F1: genFn0(), F2: v1}
                                        _ = v43
                                    }
                                }
                                _ = v42
                            }
                        case 386:
                            switch "hello" {
                            case "\n\t":
                                if false {
                                    v1 = true
                                    v1 = v1
                                }
                            default:
                                v39 = v39
                            }
                        }
                        if ((1000 > 1) && ((683 <= 724) && true)) {
                            v39 = v39
                            switch 4294967295 {
                            case -9223372036854775808:
                                if ((91 <= 100) || ("hello" <= (("\n\t" + "a") + "a"))) {
                                    v1 = false
                                    v44 := []int{}
                                    _ = v44
                                } else {
                                    fmt.Println(fi33, fi38, v1)
                                }
                            case 0:
                                var v45 uint = 42
                                _ = v45
                            default:
                                v46 := GenS4{F0: -161}
                                _ = v46
                            }
                            fmt.Println(v1, fi38, v39)
                        }
                        var v47 uint32 = 2147483647
                        fmt.Println(v1, v1)
                        var v48 uint32 = (v47 | (v47 + (v47 ^ 638)))
                        var v49 uint = 255
                        var v50 uint = 598
                        _ = v47
                        _ = v48
                        _ = v49
                        _ = v50
                    }
                }
                var v51 bool = (!v1)
                v52 := GenS5{F0: 255}
                var v53 float32 = 3.14159
                _ = fi38
                _ = v39
                _ = v51
                _ = v52
                _ = v53
            }
        }
        _ = fi33
        _ = v34
    }
    for fi54 := 0; fi54 < 3; fi54++ {
        if (2 < 2) {
            v1 = (42 == 18446744073709551615)
            v1 = v1
            v1 = (v1 || (501 == 325))
        }
        _ = fi54
    }
    var v55 int32 = -2147483648
    for fi56 := 0; fi56 < 14; fi56++ {
        for fi57 := 0; fi57 < 4; fi57++ {
            var v58 string = "a"
            _ = fi57
            _ = v58
        }
        fmt.Println(v55, fi56)
        if false {
            v55 = v55
            var v59 int32 = v55
            _ = v59
        }
        _ = fi56
    }
    v60 := []int{4294967295, 42, -541}
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v60:", len(v60), v60)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v55
    _ = v60
}
