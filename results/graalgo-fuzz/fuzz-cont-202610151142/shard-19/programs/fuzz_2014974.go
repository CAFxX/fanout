package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 uint64; F1 uint8; F2 int }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint32; F1 uint; F2 string }
type GenS7 struct { F0 int64; F1 uint; F2 bool }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]int)
    v3 := make(map[int]bool)
    var v4 uint64 = 4294967295
    fmt.Println(v4)
    v4 = 0
    v5 := GenS1{F0: -9223372036854775808}
    for fi6 := 0; fi6 < 3; fi6++ {
        fmt.Println(v4)
        for fi7 := 0; fi7 < 15; fi7++ {
            if (false && true) {
                switch "a" {
                case "hello":
                    var v8 uint64 = 0
                    _ = v8
                case "":
                    v9 := GenS2{F0: 42, F1: 1}
                    _ = v9
                case "foo bar":
                    v10 := v1[("\n\t" + "foo bar")]
                    _ = v10
                default:
                    var v11 uint64 = (v4 * v4)
                    _ = v11
                }
                var v12 float32 = 3.14159
                _ = v12
            } else {
                if true {
                    var v13 uint = 42
                    if true {
                        if (170 != -1) {
                            v14 := GenS3{F0: (v4 | 42), F1: 0, F2: -1}
                            _ = v14
                        }
                    } else {
                        v4 += (v4 | (^(v4)))
                        fmt.Println(v13, fi7, v13)
                    }
                    _ = v13
                } else {
                    v15 := GenS1{F0: -1}
                    var v16 string = ""
                    _ = v15
                    _ = v16
                }
                v4 += 2
            }
            _ = fi7
        }
        var v17 int32 = 637
        _ = fi6
        _ = v17
    }
    v18 := GenS4{F0: 32767, F1: 65535}
    _ = len(v3)
    var v19 int8 = 92
    for fi20 := 0; fi20 < 18; fi20++ {
        if false {
            v21 := []int{}
            for fi22 := 0; fi22 < 3; fi22++ {
                v4 = (v4 << ((65535) & 63))
                var v23 uint32 = 778
                v23 -= (v23 >> ((358) & 63))
                _ = fi22
                _ = v23
            }
            v4 += (v4 ^ (v4 - 4294967295))
            _ = v21
        } else {
            switch ("a" + ("foo bar" + "a")) {
            case "a":
                if (!false) {
                    v24 := GenS2{F0: 127, F1: -781}
                    _ = v24
                } else {
                    if true {
                        v19 = 0
                    } else {
                        v25 := GenS1{F0: -9223372036854775808}
                        _ = v25
                    }
                }
            case "hello":
                switch ("" + "hello") {
                case "foo bar":
                    v26 := []int{fi20, fi20, fi20}
                    _ = v26
                case "hello":
                    fmt.Println(fi20)
                default:
                    v27 := GenS5{F0: 42}
                    _ = v27
                }
            case "":
                fmt.Println(v19)
            default:
                fmt.Println(v4, fi20, v4)
            }
            if (fi20 > (fi20 * 1)) {
                delete(v1, ("hello" + "hello"))
                v28 := GenS6{F0: 2147483648, F1: 0, F2: ""}
                v19 += v19
                _ = v28
            } else {
                for fi29 := 0; fi29 < 6; fi29++ {
                    v2[9223372036854775807] = (fi20 & 9223372036854775807)
                    if false {
                        v19 -= 100
                        if false {
                            v4 -= (^(v4))
                            fmt.Println(fi29, v4)
                            var v30 uint = 1
                            _ = v30
                        } else {
                            var v31 int32 = 2147483647
                            _ = v31
                        }
                        v32 := GenS7{F0: -102, F1: 493, F2: (!((-(fi29)) > (-(fi20))))}
                        _ = v32
                    }
                    var v33 int = (fi20 | (fi20 ^ (fi29 * (fi20 >> ((2) & 63)))))
                    _ = fi29
                    _ = v33
                }
                switch "\n\t" {
                case "foo bar":
                    for fi34 := 0; fi34 < 17; fi34++ {
                        if true {
                            fmt.Println(fi34)
                            v4 -= (v4 & (^(v4)))
                        } else {
                            if (974 <= -1) {
                                var v35 uint64 = (v4 / ((v4) | 1))
                                var v36 int = (fi20 & fi34)
                                _ = v35
                                _ = v36
                            }
                            v19 -= (^(v19))
                            v4 = (v4 >> ((233) & 63))
                        }
                        fmt.Println(fi34, v4)
                        _ = fi34
                    }
                case "a":
                    delete(v1, "foo bar")
                case "":
                    v37 := []int{}
                    _ = v37
                default:
                    v38 := GenS4{F0: 32767, F1: 1}
                    _ = v38
                }
                fmt.Println(v4)
            }
            for fi39 := 0; fi39 < 6; fi39++ {
                v4 = v4
                v19 = 0
                v4 -= (v4 + v4)
                _ = fi39
            }
        }
        delete(v3, -9223372036854775808)
        _ = fi20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
