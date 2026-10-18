package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 int32; F1 uint; F2 uint8 }

func main() {
    var v1 bool = (!(192 != 128))
    var v2 bool = (42 >= 9223372036854775807)
    var v3 uint64 = 4294967295
    v4 := []string{("hello" + "hello"), ((("\n\t" + "foo bar") + ("\n\t" + "")) + "hello"), ("hello" + "foo bar")}
    for fi5 := 0; fi5 < 6; fi5++ {
        switch ((("a" + (("hello" + "hello") + ("foo bar" + ("" + "a")))) + ("\n\t" + "foo bar")) + "foo bar") {
        case "\n\t":
            _ = len(v4)
        default:
            for fi6 := 0; fi6 < 19; fi6++ {
                var v7 uint64 = (v3 / (((v3 + v3)) | 1))
                _ = fi6
                _ = v7
            }
        }
        _ = fi5
    }
    var v8 uint16 = 42
    var v9 int8 = 16
    v1 = (!(1e10 > 0.0))
    switch 674 {
    case 127:
        if true {
            v4 = append(v4, "")
            v10 := GenS1{F0: ("foo bar" + "hello")}
            _ = v10
        }
    }
    var v11 int16 = -1
    var v12 int8 = v9
    if (false || false) {
        v13 := GenS2{F0: 839, F1: v11}
        var v14 uint8 = 127
        _ = v13
        _ = v14
    }
    _ = len(v4)
    v11 -= -729
    if (!(0 <= -1)) {
        fmt.Println(v2)
    } else {
        if false {
            if v1 {
                if true {
                    switch "" {
                    case "":
                        switch "" {
                        case "\n\t":
                            if true {
                                v15 := GenS1{F0: "a"}
                                v2 = (!v2)
                                _ = v15
                            }
                        case "foo bar":
                            if (v3 <= 1) {
                                for fi16 := 0; fi16 < 18; fi16++ {
                                    if (!false) {
                                        fmt.Println(v1, v9)
                                        var v17 int8 = (v12 % (((-(v9))) | 1))
                                        v18 := GenS2{F0: (^(v3)), F1: 1}
                                        _ = v17
                                        _ = v18
                                    }
                                    _ = fi16
                                }
                            }
                        default:
                            if (!((false || (!v1)) && (v11 == v11))) {
                                v19 := GenS2{F0: v3, F1: (v11 ^ -1)}
                                _ = v19
                            } else {
                                _ = len(v4)
                                for fi20 := 0; fi20 < 6; fi20++ {
                                    switch "" {
                                    case "a":
                                        for fi21 := 0; fi21 < 18; fi21++ {
                                            var v22 bool = v1
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = fi21
                                            _ = v22
                                        }
                                    case "hello":
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    default:
                                        v23 := GenS2{F0: (v3 | v3), F1: v11}
                                        _ = v23
                                    }
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    _ = fi20
                                }
                                switch (("\n\t" + ("a" + ("a" + "\n\t"))) + "\n\t") {
                                case "":
                                    v3 = 4294967295
                                default:
                                    var v24 int32 = -1
                                    _ = v24
                                }
                            }
                        }
                    case "a":
                        var v25 float32 = 0.1
                        _ = v25
                    case "foo bar":
                        var v26 uint8 = 0
                        _ = v26
                    default:
                        v27 := GenS1{F0: "a"}
                        _ = v27
                    }
                    _ = len(v4)
                    if true {
                        var v28 uint32 = 292
                        _ = v28
                    } else {
                        var v29 uint64 = (v3 / (((v3 << ((382) & 63))) | 1))
                        var v30 uint32 = 2147483648
                        _ = v29
                        _ = v30
                    }
                } else {
                    v31 := GenS3{F0: -2147483648, F1: 255, F2: 240}
                    _ = v31
                }
            }
            v32 := GenS3{F0: 42, F1: 316, F2: 128}
            _ = v32
        }
    }
    v8 -= v8
    for fi33 := 0; fi33 < 6; fi33++ {
        v12 += -128
        if len(v4) > 0 {
            v4[0] = "a"
        }
        _ = fi33
    }
    fmt.Println(v1)
    if ((0.1 / (0.0 + 0.1)) == -2.25) {
        fmt.Println(v12, v12, v2)
        fmt.Println(v2)
    }
    v11 += (v11 >> ((42) & 63))
    v11 = (v11 - 1000)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
