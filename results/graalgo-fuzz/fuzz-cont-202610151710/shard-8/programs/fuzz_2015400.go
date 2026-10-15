package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 uint64; F1 uint; F2 uint64 }

func main() {
    v1 := []string{"foo bar", (("" + ((("" + (((("hello" + "foo bar") + "") + "foo bar") + (("a" + "a") + "hello"))) + ("a" + "\n\t")) + ((("\n\t" + "a") + ((("" + (("hello" + ("foo bar" + "a")) + ("a" + ""))) + "foo bar") + ("" + ("a" + "a")))) + "hello"))) + ""), "", "\n\t"}
    var v2 int16 = 42
    var v3 int64 = 495
    v3 -= 0
    _ = len(v1)
    v4 := GenS1{F0: 776}
    if (1.5 >= -2.25) {
        v5 := GenS1{F0: 18446744073709551615}
        _ = v5
    } else {
        switch 32767 {
        case 1:
            var v6 bool = (true || false)
            _ = v6
        default:
            if true {
                if false {
                    v2 += v2
                    var v7 int32 = 0
                    _ = v7
                } else {
                    if true {
                        v8 := GenS1{F0: 2}
                        _ = v8
                    } else {
                        fmt.Println(v2)
                    }
                }
                v2 += (v2 / ((-873) | 1))
            } else {
                if true {
                    switch 1 {
                    case 2147483647:
                        v9 := GenS1{F0: 972}
                        _ = v9
                    case 4294967295:
                        v2 = 29
                    case 1:
                        var v10 uint8 = 0
                        _ = v10
                    default:
                        v3 += v3
                    }
                    fmt.Println(v2)
                }
            }
        }
        if false {
            var v11 string = ((((("\n\t" + "\n\t") + "a") + "foo bar") + ("hello" + "\n\t")) + (("a" + "a") + "foo bar"))
            var v12 float64 = 1e10
            var v13 uint64 = 256
            _ = v11
            _ = v12
            _ = v13
        } else {
            for fi14 := 0; fi14 < 19; fi14++ {
                for fi15 := 0; fi15 < 4; fi15++ {
                    v1 = append(v1, "\n\t")
                    if false {
                        v16 := GenS2{F0: -792, F1: -844}
                        _ = v16
                    } else {
                        if (-42 <= -1) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            switch fi15 {
                            case 32768:
                                var v17 int64 = v3
                                _ = v17
                            case 1:
                                v3 = (v3 & v3)
                            case 128:
                                if true {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v2 = (v2 << ((18446744073709551615) & 63))
                                    v2 -= (v2 * 743)
                                    v2 = (v2 & v2)
                                    fmt.Println(v2)
                                } else {
                                    var v18 uint64 = 0
                                    v19 := GenS3{F0: 0}
                                    v20 := GenS4{F0: -128, F1: (^(v2))}
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                }
                            default:
                                fmt.Println(fi15)
                                _ = len(v1)
                            }
                            switch 32768 {
                            case 4294967295:
                                v21 := GenS1{F0: 18446744073709551615}
                                _ = v21
                            case 32768:
                                v3 -= -1
                            case -2147483648:
                                var v22 int64 = (v3 - v3)
                                _ = v22
                            }
                            var v23 int = -9223372036854775808
                            _ = v23
                        }
                        var v24 string = ("\n\t" + "\n\t")
                        _ = v24
                    }
                    var v25 int16 = v2
                    _ = fi15
                    _ = v25
                }
                _ = fi14
            }
            var v26 int32 = 474
            for fi27 := 0; fi27 < 5; fi27++ {
                v28 := GenS1{F0: 677}
                v26 -= -2147483648
                _ = fi27
                _ = v28
            }
            _ = v26
        }
    }
    v2 -= (v2 + (v2 / ((674) | 1)))
    v3 = 1
    v29 := GenS5{F0: 950, F1: 18446744073709551615, F2: 42}
    fmt.Println(v2)
    var v30 int64 = v3
    fmt.Println(v3, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
