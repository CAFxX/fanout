package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint8; F1 int64 }
type GenS7 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{"a"}
    v3 := []uint32{}
    v4 := make(map[uint32]string)
    v5 := []int{-424, -666}
    v6 := v4[517]
    fmt.Println(v6)
    if true {
        var v7 float32 = 3.14159
        fmt.Println(v6, v6)
        if false {
            var v8 bool = false
            _ = v8
        }
        _ = v7
    }
    if (0 != 42) {
        fmt.Println(v6)
        if len(v2) > 0 {
            v2[0] = "hello"
        }
        v9 := GenS1{F0: 48}
        _ = v9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        v6 = v6
        delete(v1, 1)
        _ = fi10
    }
    var v11 int32 = 1
    if len(v2) > 0 {
        v2[0] = ((v6 + "\n\t") + ((((("a" + ("\n\t" + ("foo bar" + ("\n\t" + v6)))) + (("hello" + v6) + v6)) + (v6 + "foo bar")) + "foo bar") + (v6 + v6)))
    }
    v2 = append(v2, (((v6 + v6) + v6) + v6))
    v12 := GenS2{F0: 2147483647, F1: 42}
    var v13 uint32 = 59
    var v14 string = ("a" + "")
    fmt.Println(v13)
    v13 = v13
    switch 42 {
    case -9223372036854775808:
        fmt.Println(v13, v11, v11)
    default:
        if (("a" + (v14 + v14)) == (v6 + "hello")) {
            if false {
                for fi15 := 0; fi15 < 4; fi15++ {
                    for fi16 := 0; fi16 < 20; fi16++ {
                        switch 2147483647 {
                        case -437:
                            fmt.Println(fi15, v6)
                        default:
                            v14 = ((v14 + v14) + v6)
                        }
                        if ((1 < 241) || (0 < 42)) {
                            v17 := GenS3{F0: 1, F1: true}
                            v6 = v6
                            delete(v4, v13)
                            _ = v17
                        }
                        v2 = append(v2, v14)
                        _ = fi16
                    }
                    v18 := GenS4{F0: 128}
                    switch "hello" {
                    case "\n\t":
                        fmt.Println(fi15, fi15)
                    case "a":
                        if false {
                            v19 := GenS5{F0: true}
                            _ = v19
                        }
                    case "hello":
                        v13 = v13
                    default:
                        if (("" + v6) > "a") {
                            v6 = (("foo bar" + v14) + v14)
                            v20 := GenS3{F0: 0, F1: true}
                            var v21 int = fi15
                            _ = v20
                            _ = v21
                        } else {
                            fmt.Println(fi15, v13, v11)
                            v22 := GenS1{F0: 43}
                            for fi23 := 0; fi23 < 2; fi23++ {
                                fmt.Println(v6)
                                v11 += (v11 / (((v11 << ((65535) & 63))) | 1))
                                v6 = v6
                                _ = fi23
                            }
                            _ = v22
                        }
                    }
                    _ = fi15
                    _ = v18
                }
                switch -409 {
                case 9223372036854775807:
                    var v24 int8 = 0
                    _ = v24
                case 4294967295:
                    for fi25 := 0; fi25 < 6; fi25++ {
                        v26 := GenS3{F0: 1, F1: true}
                        for fi27 := 0; fi27 < 14; fi27++ {
                            v28 := GenS6{F0: 1, F1: 42}
                            var v29 int = 32767
                            var v30 string = v14
                            _ = fi27
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                        if (127 <= -1) {
                            for fi31 := 0; fi31 < 16; fi31++ {
                                v3 = append(v3, (v13 | 818))
                                v14 = "hello"
                                _ = fi31
                            }
                            if (!(-1 < -616)) {
                                _ = len(v3)
                            } else {
                                switch (("a" + (v6 + "\n\t")) + v6) {
                                case "hello":
                                    if (-2.25 > 0.1) {
                                        delete(v4, (v13 ^ 2147483648))
                                        if true {
                                            v32 := GenS7{F0: -9223372036854775808}
                                            var v33 int8 = -1
                                            _ = v32
                                            _ = v33
                                        }
                                    }
                                default:
                                    if len(v2) > 0 {
                                        v2[0] = ("a" + v6)
                                    }
                                }
                            }
                            if false {
                                var v34 uint64 = 1
                                var v35 string = ""
                                _ = v34
                                _ = v35
                            }
                        }
                        _ = fi25
                        _ = v26
                    }
                }
            }
            v6 = v14
            v13 += (^(v13))
        } else {
            var v36 int8 = 0
            _ = v36
        }
    }
    if (!false) {
        var v37 string = v6
        _ = v37
    }
    switch 4294967295 {
    case -2147483648:
        if (!(true && (!(!false)))) {
            v38 := GenS4{F0: 1}
            _ = v38
        }
    default:
        var v39 float64 = 1e10
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
