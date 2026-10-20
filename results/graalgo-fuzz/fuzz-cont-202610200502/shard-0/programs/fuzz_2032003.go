package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string; F1 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int64; F1 uint32; F2 uint16 }

func main() {
    var v1 int8 = 0
    v2 := []int{}
    v3 := []uint32{2147483647, 2147483648, 65535}
    var v4 uint32 = 973
    if len(v3) > 0 {
        v3[0] = (v4 + 0)
    }
    var v5 bool = true
    v5 = false
    fmt.Println(v5)
    var v6 uint64 = uint64(v4)
    v6 += (^(v6))
    switch ("hello" + "") {
    case "\n\t":
        switch 42 {
        case -9223372036854775808:
            _ = len(v2)
        case -2147483648:
            switch ("" + "a") {
            case "a":
                var v7 int = 32767
                _ = v7
            case "":
                v4 -= 65535
            }
        default:
            for fi8 := 0; fi8 < 3; fi8++ {
                v2 = append(v2, (^(fi8)))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v9 := GenS1{F0: (v1 | (v1 % ((0) | 1))), F1: 156}
                _ = fi8
                _ = v9
            }
        }
    case "foo bar":
        if (!true) {
            for fi10 := 0; fi10 < 4; fi10++ {
                for fi11 := 0; fi11 < 11; fi11++ {
                    v12 := GenS2{F0: 0, F1: "\n\t"}
                    v3 = append(v3, 228)
                    _ = fi11
                    _ = v12
                }
                var v13 int = (fi10 * (^(fi10)))
                _ = fi10
                _ = v13
            }
            if (!(2 > 63)) {
                v4 = v4
            } else {
                var v14 int8 = (v1 | v1)
                _ = v14
            }
        }
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        switch (fi15 | fi15) {
        case 128:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        v1 += 107
        _ = fi15
    }
    v1 = (v1 * (v1 + (v1 % (((v1 % ((0) | 1))) | 1))))
    fmt.Println(v6, v1, v1)
    v1 -= (v1 + v1)
    v1 = (v1 % ((v1) | 1))
    if v5 {
        if false {
            fmt.Println(v6, v6)
            var v16 int = -589
            var v17 int16 = 32767
            _ = v16
            _ = v17
        } else {
            var v18 uint64 = v6
            v18 = (v6 ^ 42)
            if (9223372036854775807 != 42) {
                v6 = (v18 - v18)
                if v5 {
                    for fi19 := 0; fi19 < 3; fi19++ {
                        var v20 int64 = -1
                        for fi21 := 0; fi21 < 12; fi21++ {
                            if (0 > -153) {
                                if true {
                                    v5 = (true || v5)
                                } else {
                                    var v22 float32 = 0.0
                                    v18 = v6
                                    _ = v22
                                }
                                if len(v2) > 0 {
                                    v2[0] = fi21
                                }
                                var v23 int8 = v1
                                _ = v23
                            } else {
                                var v24 uint16 = uint16(v4)
                                var v25 string = "hello"
                                _ = v24
                                _ = v25
                            }
                            _ = fi21
                        }
                        _ = fi19
                        _ = v20
                    }
                    v26 := GenS1{F0: 100, F1: 42}
                    v6 += (v6 * (v6 >> ((65535) & 63)))
                    _ = v26
                } else {
                    if (3.14159 == -2.25) {
                        if len(v2) > 0 {
                            v2[0] = -611
                        }
                    } else {
                        for fi27 := 0; fi27 < 10; fi27++ {
                            v28 := GenS3{F0: v5, F1: ("\n\t" + "foo bar")}
                            v29 := GenS4{F0: 65535}
                            if false {
                                v30 := GenS2{F0: 740, F1: "\n\t"}
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v30
                            } else {
                                switch (fi27 + fi27) {
                                case 1:
                                    fmt.Println(fi27)
                                case 9223372036854775807:
                                    switch ("a" + "foo bar") {
                                    case "":
                                        var v31 float32 = float32(0.1)
                                        _ = v31
                                    case "a":
                                        v32 := GenS5{F0: "hello", F1: 2147483647}
                                        _ = v32
                                    default:
                                        v33 := GenS6{F0: (v4 & 2147483647)}
                                        _ = v33
                                    }
                                default:
                                    switch (fi27 / ((fi27) | 1)) {
                                    case 128:
                                        if (!(124 != 552)) {
                                            var v34 int32 = -1
                                            v35 := GenS5{F0: (("" + (("foo bar" + ("foo bar" + ("\n\t" + ("foo bar" + "\n\t")))) + "hello")) + ""), F1: fi27}
                                            v5 = false
                                            _ = v34
                                            _ = v35
                                        }
                                    case 504:
                                        var v36 int16 = 0
                                        _ = v36
                                    default:
                                        for fi37 := 0; fi37 < 17; fi37++ {
                                            v18 = (v6 % ((v18) | 1))
                                            v38 := GenS3{F0: (1 > 52), F1: "a"}
                                            _ = fi37
                                            _ = v38
                                        }
                                    }
                                }
                                v39 := GenS7{F0: 42, F1: (v4 + (v4 << ((18446744073709551615) & 63))), F2: 255}
                                _ = v39
                            }
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                        v6 += v6
                        v5 = v5
                    }
                }
                var v40 uint32 = v4
                _ = v40
            } else {
                fmt.Println(v1, v6)
                v3 = append(v3, 2147483647)
            }
            _ = v18
        }
    } else {
        var v41 uint = 255
        v6 = (v6 + v6)
        fmt.Println(v1, v41)
        _ = v41
    }
    for fi42 := 0; fi42 < 1; fi42++ {
        var v43 uint32 = (v4 >> ((42) & 63))
        var v44 string = (("\n\t" + ("\n\t" + "foo bar")) + ("a" + (("hello" + "a") + "foo bar")))
        _ = fi42
        _ = v43
        _ = v44
    }
    v4 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
