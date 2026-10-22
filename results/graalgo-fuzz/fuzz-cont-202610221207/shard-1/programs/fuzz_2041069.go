package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint8; F1 int64 }
type GenS7 struct { F0 uint16; F1 uint8; F2 uint32 }
func genFn0() int8 {
    for fi4 := 0; fi4 < 14; fi4++ {
        var v5 uint32 = 719
        v6 := []int{}
        v5 = v5
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 int8 = 0
    _ = v7
    return (-(v7))
}

func main() {
    var v1 int8 = 0
    v2 := make(map[int]string)
    v3 := []string{"foo bar"}
    v8 := GenS1{F0: -2147483648, F1: 0, F2: -32768}
    if (!(!true)) {
        v1 += 42
    } else {
        v9 := GenS2{F0: "\n\t"}
        v10 := GenS1{F0: -375, F1: 940, F2: 1000}
        _ = v10
        _ = v9
    }
    var v11 int32 = -998
    v12 := GenS2{F0: ((("hello" + "hello") + "\n\t") + ("hello" + "hello"))}
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 uint = 802
        _ = fi13
        _ = v14
    }
    v11 += v11
    v15 := GenS1{F0: (v11 & (v11 ^ (v11 % ((v11) | 1)))), F1: 2, F2: -1}
    v2[2147483647] = ("foo bar" + ("a" + ("\n\t" + "")))
    for fi16 := 0; fi16 < 15; fi16++ {
        if true {
            v17 := GenS3{F0: 221}
            v11 = v11
            for fi18 := 0; fi18 < 16; fi18++ {
                if true {
                    switch 0 {
                    case 0:
                        _ = len(v2)
                    case 4294967295:
                        _ = len(v3)
                    case 241:
                        v19 := v2[fi16]
                        _ = v19
                    }
                    if false {
                        if false {
                            if (626 > 42) {
                                v11 = v11
                            }
                            v11 = v11
                            v2[(^(fi18))] = ("\n\t" + ("\n\t" + "a"))
                        }
                        delete(v2, 32768)
                    } else {
                        delete(v2, (fi16 * (fi16 * (-(fi18)))))
                        if false {
                            if false {
                                v11 += (^(v11))
                                v20 := GenS1{F0: (-(v11)), F1: 968, F2: 1}
                                switch 373 {
                                case 9223372036854775807:
                                    switch (fi16 - fi16) {
                                    case 9223372036854775807:
                                        v1 = (v1 >> ((556) & 63))
                                    }
                                case 2147483647:
                                    v21 := GenS4{F0: 182, F1: true}
                                    _ = v21
                                default:
                                    var v22 uint8 = 1
                                    _ = v22
                                }
                                v11 = -794
                                _ = v20
                            }
                            if true {
                                v23 := GenS2{F0: ""}
                                if ((!false) || (genFn0() != v1)) {
                                    if (!(!(!(false || false)))) {
                                        v24 := GenS2{F0: ("foo bar" + "foo bar")}
                                        if ("" != "a") {
                                            var v25 uint64 = 0
                                            fmt.Println(fi16)
                                            switch "a" {
                                            case "hello":
                                                var v26 uint64 = (v25 << ((253) & 63))
                                                _ = v26
                                            default:
                                                v27 := GenS5{F0: 0}
                                                _ = v27
                                            }
                                            _ = v25
                                        }
                                        var v28 uint = 255
                                        _ = v24
                                        _ = v28
                                    } else {
                                        v29 := GenS1{F0: v11, F1: 42, F2: 1}
                                        _ = v29
                                    }
                                }
                                v11 += v11
                                v1 = (v1 & v1)
                                _ = v23
                            } else {
                                var v30 int8 = 42
                                _ = v30
                            }
                            v2[fi18] = "a"
                        } else {
                            v2[(fi16 & 9223372036854775807)] = "hello"
                            var v31 int64 = 42
                            _ = v31
                        }
                    }
                    var v32 bool = true
                    _ = v32
                }
                switch "hello" {
                case "":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "\n\t":
                    fmt.Println(v11, fi18, v11)
                case "a":
                    var v33 float32 = 3.14159
                    _ = v33
                default:
                    v34 := GenS6{F0: 0, F1: -9223372036854775808}
                    _ = v34
                }
                _ = fi18
            }
            _ = v17
        }
        var v35 int8 = (v1 | v1)
        fmt.Println(fi16, v1, v35)
        _ = fi16
        _ = v35
    }
    switch "a" {
    case "":
        var v36 uint64 = 0
        _ = v36
    case "a":
        v3 = append(v3, "")
    case "foo bar":
        v37 := GenS7{F0: 1, F1: 0, F2: 0}
        _ = v37
    }
    for fi38 := 0; fi38 < 13; fi38++ {
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        for fi39 := 0; fi39 < 17; fi39++ {
            var v40 uint32 = 2147483648
            _ = fi39
            _ = v40
        }
        v1 -= v1
        _ = fi38
    }
    v11 = v11
    if false {
        v1 += v1
        fmt.Println(v1)
        v11 += v11
    } else {
        v1 = (v1 & v1)
    }
    var v41 int = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v41
    _ = v8
}
