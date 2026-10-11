package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 int32; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 uint64; F1 bool }
type GenS6 struct { F0 uint8; F1 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{2147483647, 2147483648, 4294967295}
    var v3 int64 = 552
    var v4 int32 = 0
    var v5 float64 = 3.14159
    var v6 uint = 164
    if true {
        var v7 uint32 = 645
        _ = len(v2)
        if (v7 >= 2147483647) {
            var v8 int64 = 0
            v9 := GenS1{F0: 4294967295, F1: (v3 ^ 1)}
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 string = "foo bar"
                _ = fi10
                _ = v11
            }
            _ = v8
            _ = v9
        }
        _ = v7
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 uint = v6
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if (!true) {
            for fi14 := 0; fi14 < 17; fi14++ {
                switch ("hello" + "hello") {
                case "\n\t":
                    if (68 == 30) {
                        var v15 bool = true
                        fmt.Println(v15, v3, v15)
                        _ = v15
                    }
                case "a":
                    if ((v4 + (v4 ^ 42)) == 223) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        switch fi14 {
                        case -412:
                            v1[645] = 183
                        case -633:
                            if false {
                                if true {
                                    fmt.Println(v6, v6, v4)
                                    fmt.Println(v13, v6, fi14)
                                }
                                v2 = append(v2, 1)
                            }
                        case 4294967295:
                            v5 = ((v5 / v5) + v5)
                        }
                    } else {
                        v6 += v6
                    }
                default:
                    v4 += (^(v4))
                }
                _ = fi14
            }
        }
        _ = fi12
        _ = v13
    }
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch 127 {
        case 32768:
            fmt.Println(v4, v3)
        case -307:
            v6 = 0
        case -94:
            if false {
                if (810 <= -2147483648) {
                    for fi16 := 0; fi16 < 17; fi16++ {
                        for fi17 := 0; fi17 < 5; fi17++ {
                            if ((fi16 << ((v6) & 63)) > 9223372036854775807) {
                                var v18 uint16 = 838
                                _ = v18
                            } else {
                                v19 := GenS1{F0: 2, F1: 1}
                                v20 := GenS2{F0: -2147483648, F1: 32767, F2: 2}
                                v3 = 0
                                _ = v19
                                _ = v20
                            }
                            switch "a" {
                            case "foo bar":
                                v3 += v3
                            case "a":
                                var v21 float64 = (v5 - (1e10 - float64((0.1 / 1e10))))
                                _ = v21
                            case "\n\t":
                                if true {
                                    var v22 uint64 = 519
                                    _ = v22
                                }
                            default:
                                var v23 int16 = 63
                                _ = v23
                            }
                            _ = fi17
                        }
                        _ = fi16
                    }
                }
                if (true && false) {
                    v3 += (v3 << ((v6) & 63))
                } else {
                    v24 := GenS3{F0: 144, F1: 127}
                    var v25 uint64 = 18446744073709551615
                    var v26 uint32 = uint32(v6)
                    _ = v24
                    _ = v25
                    _ = v26
                }
                v27 := GenS1{F0: 42, F1: 1}
                _ = v27
            }
        default:
            if ((!true) && true) {
                v6 += 0
                if true {
                    delete(v1, 949)
                }
            }
        }
    } else {
        var v28 int16 = 541
        for fi29 := 0; fi29 < 12; fi29++ {
            v4 += v4
            fmt.Println(v3)
            v30 := GenS3{F0: 233, F1: 2}
            _ = fi29
            _ = v30
        }
        v28 = (v28 << (((v6 >> ((v6) & 63))) & 63))
        _ = v28
    }
    _ = len(v2)
    switch -2147483648 {
    case 128:
        if false {
            fmt.Println(v3, v6)
            fmt.Println(v6)
            for fi31 := 0; fi31 < 17; fi31++ {
                var v32 int = (-(fi31))
                _ = fi31
                _ = v32
            }
        } else {
            for fi33 := 0; fi33 < 1; fi33++ {
                v34 := GenS2{F0: (^(v4)), F1: 0, F2: 313}
                _ = len(v2)
                _ = fi33
                _ = v34
            }
            if false {
                v6 -= v6
                switch (("a" + "\n\t") + "foo bar") {
                case "foo bar":
                    for fi35 := 0; fi35 < 20; fi35++ {
                        var v36 uint = (v6 >> ((255) & 63))
                        _ = fi35
                        _ = v36
                    }
                }
                var v37 int16 = 291
                _ = v37
            }
        }
    case 4294967295:
        if ("a" < "") {
            v38 := GenS4{F0: -32768, F1: v4}
            _ = v38
        } else {
            v39 := GenS2{F0: v4, F1: 105, F2: 42}
            fmt.Println(v6, v4, v5)
            _ = v39
        }
    case 223:
        v40 := GenS5{F0: 0, F1: (42 <= 78)}
        _ = v40
    }
    v41 := GenS4{F0: 42, F1: v4}
    fmt.Println(v6, v5)
    v2 = append(v2, 2)
    switch 128 {
    case 4294967295:
        fmt.Println(v4, v3)
    }
    fmt.Println(v4)
    for fi42 := 0; fi42 < 2; fi42++ {
        v3 = (v3 >> ((65535) & 63))
        _ = fi42
    }
    v6 += v6
    for fi43 := 0; fi43 < 1; fi43++ {
        v44 := GenS1{F0: 808, F1: v3}
        v45 := GenS6{F0: 255, F1: 236}
        _ = fi43
        _ = v44
        _ = v45
    }
    _ = len(v1)
    var v46 uint64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v46
    _ = v5
    _ = v6
}
