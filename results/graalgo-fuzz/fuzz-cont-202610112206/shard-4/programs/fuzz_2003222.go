package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 bool }
type GenS2 struct { F0 int64; F1 uint32; F2 int }
type GenS3 struct { F0 int8; F1 uint16; F2 bool }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 int = 127
    var v2 int32 = 2147483647
    v3 := []uint32{0}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 = 584
    v4 := GenS1{F0: (!false), F1: 800, F2: (!(1 < 655))}
    var v5 uint8 = 74
    switch "foo bar" {
    case "\n\t":
        for fi6 := 0; fi6 < 11; fi6++ {
            if ((!(true || true)) && true) {
                var v7 int64 = -1
                _ = v7
            }
            _ = fi6
        }
    case "hello":
        v8 := GenS1{F0: true, F1: 0, F2: false}
        _ = v8
    default:
        if (0 == v1) {
            switch (-(v1)) {
            case -283:
                switch ("\n\t" + "a") {
                case "":
                    if false {
                        v9 := GenS1{F0: (65535 == 1), F1: 32767, F2: (false && (32767 != 482))}
                        if true {
                            v1 = v1
                            for fi10 := 0; fi10 < 5; fi10++ {
                                for fi11 := 0; fi11 < 20; fi11++ {
                                    if ((false && (!(1e10 != -2.25))) || false) {
                                        var v12 uint32 = 2147483647
                                        _ = v12
                                    }
                                    v2 -= (v2 ^ (v2 * v2))
                                    _ = fi11
                                }
                                v5 += (v5 / (((v5 | (v5 & (v5 - v5)))) | 1))
                                v5 = 1
                                _ = fi10
                            }
                        } else {
                            switch "foo bar" {
                            case "\n\t":
                                v13 := GenS1{F0: (true || false), F1: 1, F2: ((v5 & 1) > v5)}
                                _ = v13
                            default:
                                v14 := GenS2{F0: -1, F1: 0, F2: v1}
                                _ = v14
                            }
                            var v15 string = ""
                            var v16 int8 = 40
                            _ = v15
                            _ = v16
                        }
                        v5 = v5
                        _ = v9
                    }
                case "\n\t":
                    v2 = int32(v1)
                default:
                    for fi17 := 0; fi17 < 16; fi17++ {
                        var v18 int = v1
                        _ = fi17
                        _ = v18
                    }
                }
            }
        } else {
            switch (^(v1)) {
            case -787:
                for fi19 := 0; fi19 < 13; fi19++ {
                    _ = len(v3)
                    _ = fi19
                }
            case 9223372036854775807:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case 127:
                for fi20 := 0; fi20 < 11; fi20++ {
                    if (!false) {
                        var v21 string = (((("foo bar" + "hello") + "a") + ("\n\t" + "")) + ("a" + "\n\t"))
                        if false {
                            var v22 float64 = (-2.25 * float64(1.5))
                            v23 := GenS2{F0: 42, F1: 2147483647, F2: (v1 << ((4294967295) & 63))}
                            var v24 uint16 = 32767
                            _ = v22
                            _ = v23
                            _ = v24
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v25 bool = true
                            _ = v25
                        }
                        v21 = "\n\t"
                        _ = v21
                    } else {
                        if (v5 <= (v5 << ((65535) & 63))) {
                            v2 = v2
                            for fi26 := 0; fi26 < 15; fi26++ {
                                if false {
                                    var v27 int64 = -9223372036854775808
                                    v5 += (v5 ^ v5)
                                    _ = v27
                                } else {
                                    v5 += 255
                                    var v28 string = ""
                                    v29 := GenS3{F0: 127, F1: 984, F2: true}
                                    _ = v28
                                    _ = v29
                                }
                                _ = fi26
                            }
                        }
                    }
                    var v30 string = (("hello" + "a") + "foo bar")
                    for fi31 := 0; fi31 < 20; fi31++ {
                        v5 += (v5 / ((v5) | 1))
                        _ = fi31
                    }
                    _ = fi20
                    _ = v30
                }
            }
        }
    }
    if true {
        var v32 uint = 65535
        switch 0 {
        case -364:
            var v33 bool = ((v2 | (v2 << ((303) & 63))) > 446)
            _ = v33
        default:
            var v34 int64 = 42
            _ = v34
        }
        _ = v32
    } else {
        fmt.Println(v5, v1, v2)
        v1 = (v1 / (((v1 & v1)) | 1))
    }
    _ = len(v3)
    v35 := GenS4{F0: 302}
    if len(v3) > 0 {
        v3[0] = 151
    }
    switch "hello" {
    case "":
        v1 -= v1
    default:
        var v36 int32 = 1
        _ = v36
    }
    var v37 int8 = -1
    var v38 int16 = 1000
    var v39 uint16 = 552
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
