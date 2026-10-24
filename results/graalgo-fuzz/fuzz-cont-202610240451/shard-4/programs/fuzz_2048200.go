package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 string; F2 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 uint32 }

func main() {
    var v1 uint = 1
    var v2 uint8 = 0
    v3 := []string{"foo bar", ("hello" + ("" + "hello"))}
    if ((-2.25 * (0.1 - float32(1e10))) <= ((3.14159 * 0.0) * 0.0)) {
        v4 := GenS1{F0: 22}
        v5 := GenS1{F0: 42}
        _ = v4
        _ = v5
    } else {
        v1 = v1
    }
    v6 := GenS2{F0: 42, F1: (("hello" + "foo bar") + "foo bar"), F2: 42}
    v7 := GenS1{F0: 42}
    var v8 int32 = -1
    var v9 string = ""
    var v10 uint16 = 255
    var v11 bool = (v2 >= (v2 + 0))
    var v12 uint16 = (v10 ^ 0)
    if v11 {
        switch v9 {
        case "hello":
            var v13 int8 = 127
            _ = v13
        }
        var v14 uint8 = 119
        var v15 uint64 = 730
        _ = v14
        _ = v15
    }
    var v16 bool = v11
    var v17 uint8 = (v2 - 255)
    switch v9 {
    case "\n\t":
        var v18 uint = (^(v1))
        _ = v18
    case "foo bar":
        if true {
            _ = len(v3)
            if (v9 != v9) {
                fmt.Println(v16, v16)
                v8 = (v8 & v8)
            }
        } else {
            var v19 uint16 = (v12 / (((v10 * (^(v10)))) | 1))
            for fi20 := 0; fi20 < 9; fi20++ {
                if (v16 && v11) {
                    v10 = (v10 << (((v1 << (((^(v1))) & 63))) & 63))
                    for fi21 := 0; fi21 < 11; fi21++ {
                        if v16 {
                            var v22 float32 = 0.0
                            fmt.Println(v8)
                            if ((!(("a" + v9) == v9)) && v11) {
                                v9 = "\n\t"
                            }
                            _ = v22
                        } else {
                            fmt.Println(v8, fi20)
                            if (-2.25 == (1.5 + 0.1)) {
                                v3 = append(v3, (v9 + ""))
                                if (316 != 18446744073709551615) {
                                    v9 = v9
                                    v9 = "a"
                                } else {
                                    v23 := GenS3{F0: uint32(v17)}
                                    _ = v23
                                }
                            } else {
                                switch v9 {
                                case "\n\t":
                                    if v16 {
                                        _ = len(v3)
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v2 -= (^(v2))
                                    }
                                case "hello":
                                    var v24 uint = (^(v1))
                                    _ = v24
                                default:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                v3 = append(v3, "a")
                                fmt.Println(fi20, v2)
                            }
                        }
                        var v25 int32 = (v8 % ((774) | 1))
                        _ = fi21
                        _ = v25
                    }
                    v10 -= (v10 << ((2) & 63))
                }
                v26 := GenS4{F0: ("foo bar" + v9), F1: 9223372036854775807}
                var v27 float32 = ((-2.25 - 0.1) * (0.0 + -2.25))
                _ = fi20
                _ = v26
                _ = v27
            }
            v28 := GenS3{F0: 65535}
            _ = v19
            _ = v28
        }
    case "a":
        if false {
            v29 := GenS2{F0: 0, F1: "hello", F2: 459}
            v2 = (v2 % ((v2) | 1))
            var v30 int64 = -181
            _ = v29
            _ = v30
        } else {
            v16 = ((1e10 <= (0.1 - 1.5)) && ((0.0 / 1.5) < (1.5 - (1.5 * (1e10 + float64(3.14159))))))
        }
    }
    if v11 {
        for fi31 := 0; fi31 < 4; fi31++ {
            v16 = v16
            var v32 int16 = 301
            v33 := GenS3{F0: 2}
            _ = fi31
            _ = v32
            _ = v33
        }
    } else {
        var v34 float32 = ((3.14159 + 0.1) / float32(1e10))
        v35 := GenS5{F0: 443}
        v16 = false
        _ = v34
        _ = v35
    }
    if (!v16) {
        var v36 uint64 = 18446744073709551615
        _ = v36
    } else {
        _ = len(v3)
        v3 = append(v3, ("hello" + v9))
        for fi37 := 0; fi37 < 19; fi37++ {
            v3 = append(v3, (("\n\t" + (v9 + (v9 + ("foo bar" + v9)))) + ((v9 + v9) + ("hello" + v9))))
            var v38 uint16 = (v10 * (v10 & v10))
            v8 += 959
            _ = fi37
            _ = v38
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
