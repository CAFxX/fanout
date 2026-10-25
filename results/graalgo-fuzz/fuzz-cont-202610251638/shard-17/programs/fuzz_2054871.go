package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint8; F1 uint16; F2 string }
type GenS4 struct { F0 int64; F1 uint64 }

func main() {
    var v1 uint64 = 1
    v2 := make(map[uint32]bool)
    var v3 uint = 2
    v4 := []int{127}
    v5 := make(map[string]bool)
    var v6 uint32 = 42
    if false {
        for fi7 := 0; fi7 < 9; fi7++ {
            var v8 int16 = -609
            if len(v4) > 0 {
                v4[0] = (fi7 ^ fi7)
            }
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 5; fi9++ {
            v4 = append(v4, fi9)
            _ = fi9
        }
    } else {
        if true {
            for fi10 := 0; fi10 < 7; fi10++ {
                v1 += 11
                v5[("foo bar" + "hello")] = false
                _ = fi10
            }
        } else {
            v6 += v6
        }
        if (((1.5 * (1e10 / 1.0)) <= 0.1) && true) {
            v11 := GenS1{F0: -87, F1: -2147483648}
            v12 := GenS1{F0: 1, F1: -1}
            _ = v11
            _ = v12
        } else {
            if true {
                fmt.Println(v3)
                if true {
                    var v13 int8 = 100
                    v13 = (v13 + 1)
                    switch 128 {
                    case -9223372036854775808:
                        _ = len(v4)
                    case 127:
                        v3 = 641
                    case 2147483647:
                        v1 = 4294967295
                    default:
                        v14 := GenS2{F0: 675, F1: false}
                        _ = v14
                    }
                    _ = v13
                }
                for fi15 := 0; fi15 < 18; fi15++ {
                    if (v6 >= 65535) {
                        _ = len(v4)
                        if (988 < 1000) {
                            v4 = append(v4, 733)
                            if (-435 == -182) {
                                delete(v2, (v6 - v6))
                            } else {
                                fmt.Println(v6, fi15)
                                delete(v5, (((("" + "") + "") + "a") + "hello"))
                            }
                            if true {
                                v6 += (^(v6))
                                v4 = append(v4, fi15)
                            } else {
                                v3 += v3
                            }
                        } else {
                            v16 := GenS3{F0: 200, F1: 32767, F2: "foo bar"}
                            _ = v16
                        }
                        v6 = v6
                    } else {
                        for fi17 := 0; fi17 < 7; fi17++ {
                            v1 -= 390
                            delete(v2, v6)
                            _ = fi17
                        }
                    }
                    _ = fi15
                }
            } else {
                v5["foo bar"] = ((1 < 573) && false)
                v4 = append(v4, 555)
            }
            v3 = 2
            var v18 string = "\n\t"
            _ = v18
        }
    }
    delete(v2, 2147483647)
    var v19 float32 = 0.1
    v20 := GenS1{F0: 1, F1: 42}
    v19 += v19
    if (-761 >= 2147483647) {
        v19 += 1.5
    }
    var v21 int64 = -9223372036854775808
    for fi22 := 0; fi22 < 5; fi22++ {
        switch (-(fi22)) {
        case 4294967295:
            switch fi22 {
            case -1:
                fmt.Println(v3, v1)
            default:
                delete(v5, ("\n\t" + ("foo bar" + "hello")))
            }
        default:
            var v23 string = "a"
            _ = v23
        }
        _ = fi22
    }
    v24 := v2[v6]
    switch -2147483648 {
    case -265:
        switch "" {
        case "":
            for fi25 := 0; fi25 < 15; fi25++ {
                var v26 uint64 = (v1 ^ v1)
                for fi27 := 0; fi27 < 19; fi27++ {
                    v4 = append(v4, (fi27 % ((42) | 1)))
                    _ = fi27
                }
                if v24 {
                    v28 := GenS2{F0: 42, F1: (v24 || (-458 <= 32767))}
                    _ = v28
                }
                _ = fi25
                _ = v26
            }
        }
    default:
        switch ("a" + "") {
        case "":
            for fi29 := 0; fi29 < 15; fi29++ {
                v6 -= (^(v6))
                v30 := GenS1{F0: 43, F1: 1}
                _ = fi29
                _ = v30
            }
        default:
            v31 := GenS1{F0: 100, F1: -1}
            _ = v31
        }
    }
    var v32 int64 = (^(v21))
    v33 := GenS3{F0: 1, F1: 56, F2: "hello"}
    if v24 {
        delete(v2, 4294967295)
        if v24 {
            fmt.Println(v24)
            v24 = v24
            for fi34 := 0; fi34 < 7; fi34++ {
                if true {
                    for fi35 := 0; fi35 < 20; fi35++ {
                        var v36 int16 = 1000
                        _ = fi35
                        _ = v36
                    }
                    v1 += v1
                }
                _ = fi34
            }
        } else {
            v37 := GenS3{F0: 201, F1: 255, F2: "hello"}
            v32 = v32
            var v38 uint16 = 32767
            _ = v37
            _ = v38
        }
    }
    delete(v2, 42)
    switch "" {
    case "\n\t":
        v4 = append(v4, -9223372036854775808)
    case "foo bar":
        v39 := GenS4{F0: v21, F1: uint64(v3)}
        _ = v39
    default:
        fmt.Println(v6, v32)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
