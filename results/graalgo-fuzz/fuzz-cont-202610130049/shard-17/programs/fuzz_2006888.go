package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 bool; F1 uint; F2 string }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{(("" + "foo bar") + "a")}
    var v3 string = "\n\t"
    var v4 string = ""
    fmt.Println(v3, v3)
    v4 = v3
    v4 = v3
    switch 32767 {
    case 75:
        if (-634 == 505) {
            var v5 uint32 = 685
            var v6 string = (v3 + v3)
            switch "hello" {
            case "foo bar":
                if (0.1 <= (-2.25 / (1e10 / -2.25))) {
                    if true {
                        switch 423 {
                        case 813:
                            var v7 uint64 = 2
                            _ = v7
                        case 9223372036854775807:
                            v4 = v4
                        case 127:
                            for fi8 := 0; fi8 < 14; fi8++ {
                                fmt.Println(fi8, v4, v6)
                                switch (v4 + ("hello" + v6)) {
                                case "foo bar":
                                    v5 = 2147483647
                                case "":
                                    switch (fi8 >> ((4294967295) & 63)) {
                                    case -812:
                                        var v9 int64 = 976
                                        _ = v9
                                    case 341:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case 215:
                                        v5 = 4294967295
                                    }
                                default:
                                    fmt.Println(fi8, v3, fi8)
                                }
                                _ = fi8
                            }
                        }
                    } else {
                        fmt.Println(v3, v4)
                        for fi10 := 0; fi10 < 16; fi10++ {
                            if false {
                                fmt.Println(v3, v3)
                                var v11 string = v3
                                if len(v2) > 0 {
                                    v2[0] = v6
                                }
                                _ = v11
                            } else {
                                var v12 uint64 = 661
                                _ = len(v2)
                                _ = v12
                            }
                            var v13 float64 = 1.5
                            fmt.Println(v4)
                            _ = fi10
                            _ = v13
                        }
                    }
                } else {
                    v14 := GenS1{F0: 247, F1: 967}
                    v5 = v5
                    fmt.Println(v5, v6, v4)
                    _ = v14
                }
            case "hello":
                v4 = ((v4 + v3) + "hello")
            }
            _ = v5
            _ = v6
        } else {
            switch v3 {
            case "":
                var v15 string = ("" + "")
                _ = v15
            case "hello":
                var v16 int32 = 2147483647
                _ = v16
            case "a":
                for fi17 := 0; fi17 < 4; fi17++ {
                    v4 = v4
                    delete(v1, 4294967295)
                    _ = fi17
                }
            default:
                v18 := GenS1{F0: 25, F1: 18446744073709551615}
                _ = v18
            }
            var v19 int64 = 9223372036854775807
            v2 = append(v2, "a")
            _ = v19
        }
    }
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        switch (-(fi20)) {
        case 691:
            for fi21 := 0; fi21 < 6; fi21++ {
                var v22 int64 = 971
                _ = fi21
                _ = v22
            }
        case 632:
            v3 = v3
        }
        switch fi20 {
        case 577:
            v23 := GenS1{F0: 204, F1: 2}
            _ = v23
        case -1:
            if (true && (!(1 >= 255))) {
                v3 = "a"
            }
        case -2147483648:
            if true {
                switch (fi20 - (fi20 << ((1) & 63))) {
                case -1:
                    var v24 float64 = (float64(-2.25) + -2.25)
                    _ = v24
                case 32768:
                    for fi25 := 0; fi25 < 17; fi25++ {
                        var v26 float64 = 1e10
                        v4 = (v3 + v3)
                        _ = fi25
                        _ = v26
                    }
                case 32767:
                    if false {
                        for fi27 := 0; fi27 < 10; fi27++ {
                            var v28 float32 = float32(0.1)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi27
                            _ = v28
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                default:
                    v4 = "\n\t"
                }
                v29 := GenS2{F0: 40, F1: 1}
                _ = v29
            } else {
                var v30 uint8 = 113
                fmt.Println(fi20)
                v31 := GenS3{F0: (-76 != -128), F1: 325, F2: ""}
                _ = v30
                _ = v31
            }
        }
        _ = fi20
    }
    if true {
        v3 = v3
        v32 := v1[2]
        _ = v32
    } else {
        for fi33 := 0; fi33 < 8; fi33++ {
            fmt.Println(v3, v3)
            var v34 uint = 0
            _ = fi33
            _ = v34
        }
    }
    if false {
        v3 = (((v4 + "") + "a") + (v4 + (v3 + "\n\t")))
        v4 = v3
        fmt.Println(v3, v4)
    }
    v35 := GenS1{F0: 0, F1: 78}
    v3 = ((v4 + "") + "foo bar")
    _ = len(v2)
    fmt.Println(v4, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}
