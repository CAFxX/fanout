package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int; F2 string }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    v3 := []string{}
    v4 := GenS1{F0: 0}
    var v5 uint32 = 65535
    var v6 int16 = -558
    v5 -= 2147483647
    if false {
        fmt.Println(v6, v6)
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        fmt.Println(v6, v6, v5)
        v8 := GenS2{F0: 42, F1: 0}
        _ = fi7
        _ = v8
    }
    v9 := GenS3{F0: 236, F1: 32768, F2: (("\n\t" + ((("hello" + "") + ((("\n\t" + ("\n\t" + "a")) + ("hello" + "a")) + "hello")) + "foo bar")) + "")}
    switch -9223372036854775808 {
    case 42:
        var v10 int16 = v6
        _ = v10
    case -9223372036854775808:
        var v11 bool = false
        _ = v11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 int32 = 644
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        if false {
            var v15 uint = 178
            v6 += v6
            _ = v15
        }
        _ = fi14
    }
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v6, v5)
    } else {
        v5 -= 2147483647
    }
    if true {
        switch 9223372036854775807 {
        case 42:
            for fi16 := 0; fi16 < 4; fi16++ {
                if ((!false) || (v5 <= v5)) {
                    for fi17 := 0; fi17 < 13; fi17++ {
                        fmt.Println(v6)
                        _ = fi17
                    }
                }
                fmt.Println(fi16, fi16)
                _ = fi16
            }
        default:
            if (-613 == 42) {
                for fi18 := 0; fi18 < 11; fi18++ {
                    var v19 uint8 = 127
                    _ = fi18
                    _ = v19
                }
                switch "a" {
                case "foo bar":
                    var v20 int32 = 708
                    _ = v20
                case "hello":
                    switch ("hello" + "a") {
                    case "hello":
                        v6 = v6
                    case "":
                        var v21 uint16 = 0
                        _ = v21
                    }
                case "":
                    fmt.Println(v5, v5)
                }
                if ((!(!true)) && (!(523 < 32767))) {
                    var v22 uint8 = 1
                    if (957 != 2) {
                        if false {
                            v5 = 42
                            var v23 float32 = -2.25
                            _ = v23
                        } else {
                            for fi24 := 0; fi24 < 17; fi24++ {
                                if true {
                                    switch ("hello" + (((("hello" + "") + "\n\t") + ("\n\t" + "\n\t")) + ((("" + "\n\t") + ("foo bar" + "\n\t")) + ""))) {
                                    case "foo bar":
                                        v25 := GenS3{F0: (v22 + 83), F1: (fi24 ^ fi24), F2: "\n\t"}
                                        _ = v25
                                    default:
                                        if true {
                                            if (((0 >= (v5 >> ((65535) & 63))) && (((1.5 * 0.0) + float32(3.14159)) < -2.25)) && (9223372036854775807 == int64(v22))) {
                                                v6 += v6
                                                v26 := GenS2{F0: 0, F1: (v5 & 42)}
                                                _ = v26
                                            }
                                            for fi27 := 0; fi27 < 6; fi27++ {
                                                var v28 uint16 = 391
                                                _ = fi27
                                                _ = v28
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        } else {
                                            var v29 int8 = 100
                                            _ = v29
                                        }
                                    }
                                    v22 = 127
                                    _ = len(v2)
                                } else {
                                    v22 -= v22
                                }
                                v30 := GenS2{F0: -128, F1: v5}
                                v22 = 0
                                _ = fi24
                                _ = v30
                            }
                            var v31 uint64 = 0
                            _ = v31
                        }
                        fmt.Println(v5)
                    }
                    if true {
                        delete(v2, -2147483648)
                        var v32 float64 = (1e10 - 0.1)
                        var v33 float32 = float32(0.1)
                        _ = v32
                        _ = v33
                    } else {
                        var v34 uint8 = v22
                        _ = v34
                    }
                    _ = v22
                }
            } else {
                var v35 int8 = -1
                fmt.Println(v6)
                fmt.Println(v35)
                _ = v35
            }
        }
        fmt.Println(v5, v6)
        var v36 bool = ((4294967295 <= 2147483647) || true)
        _ = v36
    }
    v6 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
