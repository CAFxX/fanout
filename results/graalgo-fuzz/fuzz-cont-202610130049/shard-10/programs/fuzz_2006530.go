package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int64; F1 int; F2 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8 }

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{1}
    v3 := make(map[uint32]int)
    v4 := make(map[string]int)
    v5 := []int{512, 1}
    v6 := make(map[string]bool)
    for fi7 := 0; fi7 < 11; fi7++ {
        if (-1 < 1000) {
            var v8 float32 = 3.14159
            var v9 int16 = 872
            v9 -= (v9 * (v9 >> ((65535) & 63)))
            _ = v8
            _ = v9
        }
        var v10 uint32 = 350
        _ = fi7
        _ = v10
    }
    var v11 uint32 = 150
    var v12 int64 = 42
    v12 -= (-(v12))
    if true {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v13 int8 = 100
        _ = v13
    } else {
        if ((!(-32768 > 0)) || (((0.0 - 3.14159) * (((3.14159 - (0.0 - -2.25)) / ((3.14159 - 1e10) / (((1.5 * (1.5 - 1e10)) + (1.5 / 0.1)) * 1e10))) - (1.5 + 3.14159))) != 1e10)) {
            var v14 string = ("" + "foo bar")
            _ = v14
        }
        v15 := GenS1{F0: (!(!true))}
        _ = v15
    }
    var v16 int32 = -954
    switch 128 {
    case 4294967295:
        fmt.Println(v16)
    case -2147483648:
        v17 := GenS1{F0: (!false)}
        _ = v17
    case 127:
        v18 := GenS2{F0: true, F1: -32768}
        _ = v18
    }
    switch ("a" + ("foo bar" + "a")) {
    case "hello":
        v11 = v11
    case "a":
        for fi19 := 0; fi19 < 1; fi19++ {
            var v20 uint8 = 0
            var v21 string = ("hello" + "foo bar")
            _ = fi19
            _ = v20
            _ = v21
        }
    default:
        if (true || (true || false)) {
            v11 = 2
            v11 += v11
            var v22 string = ""
            _ = v22
        }
    }
    v12 -= (v12 + (v12 | (v12 / (((v12 - (v12 * -1))) | 1))))
    _ = len(v6)
    if (42 != (v16 | -777)) {
        v23 := GenS1{F0: false}
        var v24 float64 = -2.25
        if false {
            fmt.Println(v12)
            v12 = (v12 ^ 42)
            v25 := GenS1{F0: true}
            _ = v25
        } else {
            v26 := GenS2{F0: ((v16 - (v16 - (v16 / ((v16) | 1)))) == 0), F1: 1}
            if true {
                v11 = 2147483648
                for fi27 := 0; fi27 < 16; fi27++ {
                    v11 = 126
                    _ = fi27
                }
                v11 = v11
            } else {
                if ("a" <= ("foo bar" + "")) {
                    fmt.Println(v16, v12)
                } else {
                    v28 := GenS3{F0: (v12 / ((v12) | 1)), F1: 127, F2: (^(v11))}
                    v29 := GenS4{F0: "foo bar"}
                    v30 := GenS5{F0: -128}
                    _ = v28
                    _ = v29
                    _ = v30
                }
            }
            _ = v26
        }
        _ = v23
        _ = v24
    }
    if false {
        v11 = 496
    } else {
        v31 := v6[("foo bar" + "\n\t")]
        _ = v31
    }
    for fi32 := 0; fi32 < 2; fi32++ {
        switch (fi32 | fi32) {
        case 0:
            switch ("foo bar" + ("" + "")) {
            case "hello":
                for fi33 := 0; fi33 < 11; fi33++ {
                    var v34 bool = true
                    if (v34 && v34) {
                        v35 := GenS5{F0: 98}
                        var v36 int32 = (v16 - 2147483647)
                        _ = v35
                        _ = v36
                    } else {
                        v34 = (false || v34)
                        v37 := GenS2{F0: (1.5 == (1e10 / 3.14159)), F1: 1}
                        v38 := GenS4{F0: (((("\n\t" + "foo bar") + ("\n\t" + (("a" + "a") + "foo bar"))) + "\n\t") + (("a" + "") + "\n\t"))}
                        var v39 uint16 = 449
                        v11 = 2147483648
                        if v34 {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        } else {
                            if (v34 && v34) {
                                fmt.Println(v12, fi32)
                                var v40 uint32 = (^(v11))
                                _ = v40
                            } else {
                                switch "foo bar" {
                                case "hello":
                                    var v41 int8 = 127
                                    _ = v41
                                case "\n\t":
                                    v11 -= (v11 / (((v11 / (((v11 + 4294967295)) | 1))) | 1))
                                }
                                var v42 int8 = 100
                                var v43 int = (fi32 >> ((985) & 63))
                                _ = v42
                                _ = v43
                            }
                        }
                        v44 := GenS1{F0: ((1.5 * 0.0) <= 0.1)}
                        v45 := GenS4{F0: "\n\t"}
                        _ = v37
                        _ = v38
                        _ = v39
                        _ = v44
                        _ = v45
                    }
                    fmt.Println(v12)
                    _ = fi33
                    _ = v34
                }
            case "\n\t":
                if len(v5) > 0 {
                    v5[0] = fi32
                }
            default:
                v11 -= (^(v11))
            }
        case 42:
            var v46 uint = 1
            _ = v46
        }
        v12 -= 849
        var v47 int64 = int64(fi32)
        _ = fi32
        _ = v47
    }
    v16 = (v16 + (v16 % ((v16) | 1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
