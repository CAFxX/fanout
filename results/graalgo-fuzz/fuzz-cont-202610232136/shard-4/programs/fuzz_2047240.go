package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint32; F1 string; F2 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := []string{"foo bar", "hello", ("\n\t" + "hello")}
    var v2 int64 = 700
    var v3 uint16 = 934
    v3 -= 65535
    var v4 uint8 = 42
    switch ("\n\t" + "hello") {
    case "hello":
        for fi5 := 0; fi5 < 15; fi5++ {
            if ((fi5 < (fi5 + fi5)) || (!(uint(v4) == 0))) {
                v6 := GenS1{F0: 18446744073709551615, F1: (!((false && false) && (0.0 > (-2.25 - (1.5 - 3.14159)))))}
                _ = v6
            } else {
                if false {
                    v3 += 65535
                } else {
                    for fi7 := 0; fi7 < 1; fi7++ {
                        if false {
                            if (v3 == (v3 * v3)) {
                                fmt.Println(v3)
                                v4 += (v4 + (v4 & 23))
                                var v8 int64 = (v2 & v2)
                                _ = v8
                            } else {
                                v1 = append(v1, ("foo bar" + "a"))
                                switch "foo bar" {
                                case "foo bar":
                                    v9 := GenS1{F0: 18446744073709551615, F1: true}
                                    _ = v9
                                case "\n\t":
                                    fmt.Println(fi7)
                                case "hello":
                                    if (0.1 == 0.1) {
                                        v3 = (v3 ^ (v3 - v3))
                                        fmt.Println(v3, v3, v2)
                                        v4 = (^(v4))
                                    } else {
                                        if (1 >= 127) {
                                            v2 += v2
                                            fmt.Println(fi5, v2)
                                            v3 -= v3
                                        }
                                        v4 += 1
                                    }
                                default:
                                    var v10 uint64 = 226
                                    _ = v10
                                }
                                var v11 float64 = 0.1
                                _ = v11
                            }
                            var v12 int32 = 798
                            var v13 int64 = 802
                            _ = v12
                            _ = v13
                        }
                        _ = fi7
                    }
                    v14 := GenS2{F0: 2147483647, F1: "hello", F2: 2}
                    v15 := GenS2{F0: 70, F1: "a", F2: 42}
                    _ = v14
                    _ = v15
                }
                v4 -= v4
            }
            _ = fi5
        }
    default:
        if false {
            for fi16 := 0; fi16 < 5; fi16++ {
                if (!(1 <= -1)) {
                    var v17 uint = 236
                    if false {
                        switch (fi16 + fi16) {
                        case 247:
                            _ = len(v1)
                        default:
                            v4 = (v4 + v4)
                        }
                    }
                    _ = v17
                }
                v4 = v4
                fmt.Println(v4, v2, v4)
                _ = fi16
            }
            v18 := GenS3{F0: int8(v4)}
            if false {
                v3 = v3
            }
            _ = v18
        } else {
            if (0 >= 214) {
                fmt.Println(v3)
                var v19 float32 = 3.14159
                _ = v19
            }
            v4 = (v4 >> ((18446744073709551615) & 63))
        }
    }
    switch 127 {
    case 32768:
        v20 := GenS2{F0: 65535, F1: "a", F2: (^(v3))}
        _ = v20
    default:
        fmt.Println(v4, v3, v3)
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        switch "a" {
        case "":
            var v22 string = ("" + "")
            _ = v22
        case "foo bar":
            v2 = v2
        case "a":
            v23 := GenS1{F0: 2, F1: (true && true)}
            _ = v23
        default:
            v1 = append(v1, "hello")
        }
        _ = fi21
    }
    v24 := GenS2{F0: 2147483647, F1: (("foo bar" + "hello") + "\n\t"), F2: (v3 / (((v3 + 2)) | 1))}
    _ = len(v1)
    v25 := GenS4{F0: 0}
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
