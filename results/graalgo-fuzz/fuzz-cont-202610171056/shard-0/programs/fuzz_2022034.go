package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 string }
type GenS2 struct { F0 int16; F1 string; F2 uint8 }
type GenS3 struct { F0 int32; F1 uint }

func main() {
    v1 := []string{"", "", ("hello" + "a")}
    var v2 bool = (!(true || true))
    var v3 int8 = 42
    v4 := make(map[int]bool)
    var v5 bool = v2
    fmt.Println(v3, v5, v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = (815 >= 292)
    var v6 uint32 = 4294967295
    v2 = (((v5 || v5) && v2) || v2)
    var v7 int = 579
    fmt.Println(v7, v2, v7)
    v3 = v3
    v8 := GenS1{F0: 362, F1: 255, F2: "foo bar"}
    if (!v5) {
        v9 := GenS1{F0: 674, F1: 367, F2: "hello"}
        v10 := GenS2{F0: 0, F1: ((("hello" + ("a" + "")) + ("\n\t" + ("hello" + ""))) + ("foo bar" + "hello")), F2: 0}
        v3 = v3
        _ = v10
        _ = v9
    } else {
        v11 := GenS2{F0: 32767, F1: "\n\t", F2: 255}
        _ = v11
    }
    var v12 uint32 = (^(v6))
    v6 -= (v12 << ((42) & 63))
    if (1 >= 18446744073709551615) {
        v13 := GenS1{F0: 18446744073709551615, F1: 1, F2: ((("\n\t" + ("hello" + "\n\t")) + ("\n\t" + "a")) + ("a" + "a"))}
        var v14 uint8 = 72
        fmt.Println(v14, v7, v6)
        _ = v13
        _ = v14
    }
    fmt.Println(v6, v2, v12)
    v12 += 2
    for fi15 := 0; fi15 < 20; fi15++ {
        switch "" {
        case "a":
            if v5 {
                for fi16 := 0; fi16 < 13; fi16++ {
                    var v17 uint32 = 2
                    _ = fi16
                    _ = v17
                }
                for fi18 := 0; fi18 < 2; fi18++ {
                    var v19 uint8 = 96
                    var v20 uint64 = 42
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            } else {
                switch (("a" + ("\n\t" + "foo bar")) + "hello") {
                case "a":
                    switch v7 {
                    case -570:
                        var v21 bool = v2
                        _ = v21
                    case 928:
                        v1 = append(v1, "a")
                    case -1:
                        if (!v5) {
                            v6 = 494
                        } else {
                            for fi22 := 0; fi22 < 13; fi22++ {
                                fmt.Println(v2, v12)
                                var v23 uint64 = 311
                                var v24 int = fi15
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                        }
                    default:
                        var v25 int8 = (v3 >> ((408) & 63))
                        _ = v25
                    }
                case "\n\t":
                    v26 := GenS1{F0: 438, F1: 2, F2: "\n\t"}
                    _ = v26
                default:
                    switch v7 {
                    case 127:
                        switch 32767 {
                        case -601:
                            v3 = v3
                        default:
                            switch "\n\t" {
                            case "a":
                                var v27 int64 = 13
                                _ = v27
                            case "":
                                switch v7 {
                                case -9223372036854775808:
                                    if len(v1) > 0 {
                                        v1[0] = "foo bar"
                                    }
                                }
                            case "\n\t":
                                for fi28 := 0; fi28 < 13; fi28++ {
                                    v29 := GenS3{F0: 42, F1: 42}
                                    _ = fi28
                                    _ = v29
                                }
                            default:
                                switch (v7 - v7) {
                                case 128:
                                    fmt.Println(v7, fi15)
                                default:
                                    var v30 int16 = -415
                                    _ = v30
                                }
                            }
                        }
                    case 2147483647:
                        _ = len(v4)
                    }
                }
            }
        default:
            var v31 float32 = -2.25
            _ = v31
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi15
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
