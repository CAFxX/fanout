package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 bool; F1 string }
type GenS6 struct { F0 uint32; F1 string; F2 int }
func genFn0() uint64 {
    var v7 uint8 = 127
    fmt.Println(v7)
    _ = v7
    return 42
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint8 = 1
    var v3 int8 = -1
    var v4 uint16 = 932
    v5 := []string{"a", "", "", "foo bar"}
    var v6 uint64 = uint64(v2)
    fmt.Println(v6, v6, v3)
    if (42 == -1) {
        var v8 int32 = -1
        if ((v8 % ((42) | 1)) != (^(v8))) {
            var v9 uint32 = 509
            for fi10 := 0; fi10 < 2; fi10++ {
                v11 := GenS1{F0: 94}
                v12 := GenS2{F0: genFn0()}
                fmt.Println(v2, v8)
                _ = fi10
                _ = v11
                _ = v12
            }
            for fi13 := 0; fi13 < 13; fi13++ {
                var v14 uint = 2
                switch 32767 {
                case 388:
                    if true {
                        v8 = v8
                        fmt.Println(v6, v14)
                        switch 42 {
                        case 9223372036854775807:
                            var v15 bool = false
                            _ = v15
                        case 2147483647:
                            switch "" {
                            case "hello":
                                v5 = append(v5, "")
                            case "\n\t":
                                v4 += v4
                            case "foo bar":
                                v16 := GenS1{F0: v9}
                                _ = v16
                            }
                        case 0:
                            if (1.5 == 0.1) {
                                if (254 == v2) {
                                    v8 = (v8 - int32(v9))
                                    var v17 uint32 = 2147483647
                                    _ = v17
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                        }
                    }
                case 32767:
                    if true {
                        v18 := GenS1{F0: v9}
                        var v19 int32 = (^(v8))
                        _ = v18
                        _ = v19
                    } else {
                        v20 := GenS3{F0: -1, F1: v9, F2: genFn0()}
                        v21 := GenS2{F0: (v6 | 585)}
                        for fi22 := 0; fi22 < 15; fi22++ {
                            v23 := GenS4{F0: v3, F1: 32767}
                            _ = fi22
                            _ = v23
                        }
                        _ = v20
                        _ = v21
                    }
                default:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = fi13
                _ = v14
            }
            _ = v9
        } else {
            v24 := GenS1{F0: 155}
            if false {
                fmt.Println(v6)
            } else {
                var v25 int = 32767
                var v26 float64 = (0.1 + 0.0)
                _ = v25
                _ = v26
            }
            for fi27 := 0; fi27 < 3; fi27++ {
                if true {
                    var v28 float32 = 1e10
                    _ = v28
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                for fi29 := 0; fi29 < 10; fi29++ {
                    fmt.Println(v2, v4, v1)
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi29
                }
                fmt.Println(v8, v1)
                _ = fi27
            }
            _ = v24
        }
        var v30 int8 = v3
        _ = v30
        _ = v8
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        _ = len(v5)
        switch "\n\t" {
        case "":
            for fi32 := 0; fi32 < 4; fi32++ {
                switch ("" + "foo bar") {
                case "":
                    switch "foo bar" {
                    case "a":
                        fmt.Println(fi32, v6)
                    case "foo bar":
                        switch ("" + "\n\t") {
                        case "a":
                            v33 := GenS1{F0: 65535}
                            _ = v33
                        case "\n\t":
                            v1 += 4294967295
                        }
                    }
                case "foo bar":
                    v3 = v3
                case "\n\t":
                    if (0 > 0) {
                        if (65535 >= 2) {
                            var v34 int64 = 306
                            v4 -= (v4 << ((255) & 63))
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = v34
                        } else {
                            v6 = 4294967295
                            v4 += (v4 - 255)
                        }
                    }
                }
                _ = fi32
            }
        default:
            fmt.Println(v2, v1, v6)
        }
        _ = fi31
    }
    var v35 string = "\n\t"
    v3 = v3
    v35 = ("a" + v35)
    v36 := GenS1{F0: 2}
    v37 := GenS5{F0: true, F1: "a"}
    if len(v5) > 0 {
        v5[0] = v35
    }
    v3 = v3
    var v38 uint8 = v2
    switch (v35 + v35) {
    case "hello":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "\n\t":
        v5 = append(v5, "a")
    default:
        v39 := GenS5{F0: ((!(!(246 != -957))) || (0.1 > 0.1)), F1: ("a" + "hello")}
        _ = v39
    }
    v4 -= (v4 % ((65535) | 1))
    if false {
        var v40 uint64 = genFn0()
        v41 := GenS6{F0: 0, F1: "foo bar", F2: 42}
        var v42 uint64 = v40
        _ = v40
        _ = v41
        _ = v42
    }
    v1 -= genFn0()
    _ = len(v5)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v43 int = 128
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
