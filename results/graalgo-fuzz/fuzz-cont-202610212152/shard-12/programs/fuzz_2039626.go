package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint; F1 int; F2 string }
type GenS3 struct { F0 int; F1 int32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    var v3 int64 = -1
    var v4 uint = 65535
    for fi5 := 0; fi5 < 15; fi5++ {
        v3 += 866
        _ = fi5
    }
    var v6 int16 = -79
    v7 := GenS1{F0: 165, F1: 2}
    fmt.Println(v4)
    if true {
        if (1 == 124) {
            v1[2] = false
            if ("\n\t" == ("a" + "foo bar")) {
                if (!true) {
                    var v8 string = "a"
                    if false {
                        v6 = -32768
                    }
                    _ = v8
                } else {
                    delete(v2, "hello")
                    switch ("foo bar" + ("hello" + ("a" + "a"))) {
                    case "foo bar":
                        if ((58 > 1) && (!true)) {
                            v6 += -781
                        }
                    case "\n\t":
                        delete(v1, 53)
                    default:
                        if (false && ((v6 - v6) < v6)) {
                            var v9 uint32 = 259
                            _ = v9
                        }
                    }
                }
                for fi10 := 0; fi10 < 13; fi10++ {
                    fmt.Println(fi10)
                    _ = fi10
                }
            }
            fmt.Println(v3, v4, v3)
        }
        for fi11 := 0; fi11 < 1; fi11++ {
            fmt.Println(v6)
            _ = fi11
        }
        var v12 uint16 = 0
        _ = v12
    } else {
        if ((((3.14159 * 0.1) - 3.14159) * 1e10) > 1.5) {
            v3 += v3
        } else {
            switch 2147483647 {
            case 9223372036854775807:
                if false {
                    v3 -= v3
                }
            case 32768:
                for fi13 := 0; fi13 < 2; fi13++ {
                    for fi14 := 0; fi14 < 3; fi14++ {
                        v3 -= (v3 | int64(fi13))
                        var v15 int64 = int64(fi13)
                        _ = fi14
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 3; fi16++ {
                        switch "foo bar" {
                        case "\n\t":
                            switch "foo bar" {
                            case "":
                                if (77 != 123) {
                                    fmt.Println(v3)
                                    v2[("\n\t" + "")] = fi13
                                }
                            case "a":
                                switch "hello" {
                                case "a":
                                    v3 += v3
                                }
                                v3 += v3
                                var v17 uint8 = 255
                                var v18 float64 = float64(0.1)
                                var v19 uint32 = 841
                                _ = v17
                                _ = v18
                                _ = v19
                            default:
                                switch "hello" {
                                case "foo bar":
                                    if (128 <= 255) {
                                        v20 := GenS2{F0: v4, F1: (fi13 + fi16), F2: ("\n\t" + "")}
                                        _ = v20
                                    }
                                case "a":
                                    v21 := []int{}
                                    _ = v21
                                default:
                                    var v22 int = 0
                                    _ = v22
                                }
                            }
                        case "":
                            var v23 uint = 1
                            _ = v23
                        default:
                            v6 += 880
                        }
                        _ = fi16
                    }
                    _ = fi13
                }
            default:
                v24 := GenS2{F0: v4, F1: 0, F2: "\n\t"}
                _ = v24
            }
        }
    }
    v25 := GenS2{F0: v4, F1: -775, F2: "foo bar"}
    if false {
        v3 += v3
        fmt.Println(v4, v3, v3)
        if ((1 != 0) || (71 < -1)) {
            v3 = 9223372036854775807
            if true {
                v4 = 292
                v26 := GenS2{F0: v4, F1: int(v3), F2: (("\n\t" + "") + (("hello" + "foo bar") + "a"))}
                _ = v26
            }
        }
    }
    var v27 int8 = 42
    for fi28 := 0; fi28 < 9; fi28++ {
        var v29 float32 = 0.1
        var v30 uint64 = 352
        var v31 int = fi28
        _ = fi28
        _ = v29
        _ = v30
        _ = v31
    }
    switch "" {
    case "hello":
        var v32 uint16 = 0
        _ = v32
    default:
        v33 := GenS1{F0: 1, F1: 2}
        _ = v33
    }
    switch (("\n\t" + "foo bar") + "hello") {
    case "":
        fmt.Println(v3, v4, v4)
    case "hello":
        fmt.Println(v27, v6, v6)
    default:
        v27 -= -128
    }
    var v34 int64 = (^(v3))
    for fi35 := 0; fi35 < 7; fi35++ {
        var v36 int16 = (v6 >> (((v4 - (^(v4)))) & 63))
        v37 := []int{fi35, fi35, 42, (fi35 >> ((v4) & 63))}
        if true {
            var v38 int64 = v3
            _ = v38
        } else {
            v39 := GenS3{F0: 9223372036854775807, F1: -860}
            var v40 bool = (!(-2.25 > float64(float32(-2.25))))
            _ = v39
            _ = v40
        }
        _ = fi35
        _ = v36
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v34
    _ = v4
    _ = v6
    _ = v7
}
