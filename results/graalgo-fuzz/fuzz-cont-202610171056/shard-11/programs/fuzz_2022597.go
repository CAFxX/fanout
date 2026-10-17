package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 int; F1 string; F2 uint }

func main() {
    var v1 int = 0
    var v2 uint16 = 813
    v3 := []string{"\n\t", ("a" + ""), "a", "foo bar"}
    v4 := make(map[uint32]string)
    var v5 uint16 = v2
    var v6 uint = 994
    v6 += 65535
    v2 -= (v5 / ((v2) | 1))
    switch (v1 ^ v1) {
    case -135:
        v5 -= 422
    case 32767:
        switch (^(v1)) {
        case -1:
            v7 := GenS1{F0: v1, F1: "hello"}
            _ = v7
        case 1:
            var v8 int16 = -32768
            _ = v8
        case -412:
            for fi9 := 0; fi9 < 5; fi9++ {
                fmt.Println(v2)
                _ = fi9
            }
        default:
            v1 += -1
        }
    default:
        v10 := GenS2{F0: 2, F1: 2147483648}
        _ = v10
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        if true {
            v5 += 912
            var v12 uint = (^(v6))
            switch (fi11 - fi11) {
            case 89:
                v1 -= (fi11 << ((v12) & 63))
            }
            _ = v12
        } else {
            var v13 int8 = -1
            _ = v13
        }
        if len(v3) > 0 {
            v3[0] = ("foo bar" + "foo bar")
        }
        v14 := GenS3{F0: 107, F1: (100 != 0)}
        _ = fi11
        _ = v14
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        for fi16 := 0; fi16 < 17; fi16++ {
            if (65535 <= 4294967295) {
                v5 += v5
                switch "\n\t" {
                case "\n\t":
                    v17 := GenS4{F0: -64, F1: true}
                    _ = v17
                case "foo bar":
                    v2 += (v2 & (v5 / ((42) | 1)))
                case "hello":
                    switch "" {
                    case "\n\t":
                        v5 = v5
                    default:
                        v18 := GenS3{F0: 2147483647, F1: false}
                        _ = v18
                    }
                default:
                    var v19 int = v1
                    _ = v19
                }
                fmt.Println(v1, fi15, v1)
                switch "a" {
                case "":
                    v6 = (v6 >> ((65535) & 63))
                    if len(v3) > 0 {
                        v3[0] = ""
                    }
                    switch "foo bar" {
                    case "hello":
                        v2 = (^(v2))
                    case "a":
                        var v20 float32 = 3.14159
                        _ = v20
                    case "":
                        v2 -= (v2 + 32767)
                    }
                case "hello":
                    fmt.Println(v6, v2, v6)
                case "\n\t":
                    var v21 bool = true
                    _ = v21
                }
                var v22 uint64 = 591
                v6 -= v6
                _ = v22
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v1 = (^(fi15))
                v5 -= (v2 << ((65535) & 63))
            }
            _ = fi16
        }
        var v23 string = ""
        _ = fi15
        _ = v23
    }
    fmt.Println(v6, v6)
    var v24 float32 = 1.5
    v4[1] = "\n\t"
    if false {
        for fi25 := 0; fi25 < 17; fi25++ {
            var v26 float32 = ((v24 / v24) - (float32(-2.25) / v24))
            fmt.Println(v5)
            _ = fi25
            _ = v26
        }
    } else {
        v27 := GenS3{F0: 9, F1: true}
        var v28 uint = (^(v6))
        v6 = v28
        _ = v27
        _ = v28
    }
    if (false || true) {
        for fi29 := 0; fi29 < 11; fi29++ {
            if (!(v5 != 32767)) {
                _ = len(v3)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                if true {
                    var v30 float32 = float32((1.5 + 0.1))
                    fmt.Println(v24, fi29, v1)
                    _ = v30
                } else {
                    switch "hello" {
                    case "foo bar":
                        if (!false) {
                            if (true && ((-2147483648 == 42) || false)) {
                                if (!false) {
                                    _ = len(v3)
                                } else {
                                    fmt.Println(v6)
                                }
                                switch v1 {
                                case -476:
                                    fmt.Println(fi29)
                                case -325:
                                    if false {
                                        fmt.Println(v2)
                                        var v31 uint32 = 1
                                        _ = v31
                                    }
                                case -2147483648:
                                    v32 := GenS5{F0: (v1 + (-(v1))), F1: ((("\n\t" + "a") + ("\n\t" + "")) + "foo bar"), F2: v6}
                                    _ = v32
                                }
                            } else {
                                if false {
                                    v33 := GenS3{F0: 719, F1: (true && false)}
                                    v3 = append(v3, ("" + "a"))
                                    _ = v33
                                }
                            }
                        } else {
                            var v34 uint32 = 4294967295
                            v34 -= (v34 + 25)
                            if false {
                                var v35 string = ("a" + "")
                                _ = v35
                            }
                            _ = v34
                        }
                    }
                    v4[573] = ("hello" + "\n\t")
                    var v36 float32 = (v24 / (v24 / (v24 / 3.14159)))
                    _ = v36
                }
                v4[107] = ("a" + ("\n\t" + "a"))
                var v37 float32 = v24
                _ = v37
            }
            if ((1 > 0) || (2 > 42)) {
                v38 := GenS1{F0: fi29, F1: "hello"}
                var v39 uint64 = 1
                for fi40 := 0; fi40 < 9; fi40++ {
                    switch ("a" + ("a" + "")) {
                    case "\n\t":
                        var v41 float64 = 0.1
                        _ = v41
                    }
                    _ = fi40
                }
                _ = v38
                _ = v39
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v42 := GenS5{F0: (-(fi29)), F1: "foo bar", F2: v6}
                var v43 bool = true
                _ = v42
                _ = v43
            }
            _ = fi29
        }
        if true {
            var v44 uint64 = uint64(v6)
            var v45 int8 = -33
            _ = v44
            _ = v45
        } else {
            v46 := GenS5{F0: (-(v1)), F1: ("" + ("hello" + "hello")), F2: 255}
            var v47 int16 = 50
            var v48 int32 = -1
            _ = v46
            _ = v47
            _ = v48
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
