package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint }
type GenS3 struct { F0 int64; F1 int8; F2 uint16 }
type GenS4 struct { F0 int16 }
func genFn0() int16 {
    v4 := []int{287, -113, 127}
    switch "foo bar" {
    case "a":
        var v5 int16 = 418
        _ = v5
    }
    _ = v4
    return -1
}

func main() {
    var v1 uint8 = 159
    var v2 uint32 = 42
    var v3 bool = false
    v2 = v2
    v6 := []int{}
    v1 = 42
    if (!true) {
        switch ("" + ((((((((("foo bar" + "a") + "hello") + (((("" + ("\n\t" + "foo bar")) + "hello") + "foo bar") + "a")) + "hello") + "") + "foo bar") + "hello") + "\n\t") + "foo bar")) {
        case "hello":
            v7 := GenS1{F0: v3, F1: 0, F2: v1}
            _ = v7
        }
        for fi8 := 0; fi8 < 8; fi8++ {
            for fi9 := 0; fi9 < 5; fi9++ {
                switch (^(fi9)) {
                case 0:
                    v10 := GenS1{F0: (42 > -79), F1: 0, F2: (v1 / (((^(v1))) | 1))}
                    _ = v10
                default:
                    var v11 string = "\n\t"
                    _ = v11
                }
                fmt.Println(v1)
                var v12 uint8 = (v1 & 42)
                _ = fi9
                _ = v12
            }
            _ = fi8
        }
        var v13 uint32 = v2
        _ = v13
    }
    switch 127 {
    case -128:
        v2 = (^(v2))
    case 1:
        var v14 uint8 = v1
        _ = v14
    default:
        v1 -= v1
    }
    _ = len(v6)
    if (v3 || v3) {
        v6 = append(v6, 32768)
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 int16 = genFn0()
        v17 := GenS2{F0: v2, F1: 2147483648, F2: 18446744073709551615}
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 uint64 = 2
    fmt.Println(v2, v2, v3)
    v2 -= (v2 | v2)
    switch 411 {
    case 0:
        if false {
            var v19 uint32 = (v2 >> ((65535) & 63))
            _ = v19
        } else {
            v20 := GenS1{F0: v3, F1: 4294967295, F2: v1}
            var v21 uint32 = (v2 << ((18446744073709551615) & 63))
            _ = v20
            _ = v21
        }
    default:
        v22 := GenS2{F0: (v2 % ((v2) | 1)), F1: (^(v2)), F2: 2}
        _ = v22
    }
    v1 = (v1 >> ((207) & 63))
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v23 := GenS1{F0: v3, F1: 42, F2: v1}
    if (-373 != 0) {
        v24 := GenS3{F0: -1, F1: 1, F2: 65535}
        _ = v24
    } else {
        var v25 string = (("hello" + "\n\t") + ("" + (("a" + "hello") + ("a" + "a"))))
        var v26 float32 = 0.1
        for fi27 := 0; fi27 < 5; fi27++ {
            if v3 {
                for fi28 := 0; fi28 < 1; fi28++ {
                    switch "\n\t" {
                    case "hello":
                        var v29 bool = (!(720 > (v2 / ((v2) | 1))))
                        _ = v29
                    case "a":
                        fmt.Println(v18, v26)
                    default:
                        if (32767 != 38) {
                            switch fi28 {
                            case 127:
                                v30 := GenS1{F0: v3, F1: 310, F2: (v1 ^ (v1 * (^(v1))))}
                                var v31 uint8 = (v1 + (v1 + (v1 + v1)))
                                var v32 int = fi28
                                switch (-(fi28)) {
                                case 32767:
                                    var v33 int64 = -677
                                    v3 = (!(!v3))
                                    _ = v33
                                }
                                _ = v30
                                _ = v31
                                _ = v32
                            case 42:
                                _ = len(v6)
                            case -1:
                                if v3 {
                                    switch fi28 {
                                    case -840:
                                        if v3 {
                                            v2 += v2
                                            if ((-686 == 2147483647) && false) {
                                                var v34 uint = 533
                                                var v35 int64 = 1
                                                _ = v34
                                                _ = v35
                                            }
                                        }
                                    default:
                                        if (0 == 567) {
                                            fmt.Println(fi27, v18, v2)
                                        } else {
                                            fmt.Println(fi27)
                                            var v36 uint8 = v1
                                            _ = v36
                                        }
                                        v37 := GenS1{F0: v3, F1: 42, F2: 42}
                                        _ = v37
                                    }
                                } else {
                                    if v3 {
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        fmt.Println(fi27, fi28)
                                        v38 := GenS1{F0: ((v18 / ((4294967295) | 1)) < (v18 << ((4294967295) & 63))), F1: 42, F2: v1}
                                        _ = v38
                                    }
                                }
                            }
                        } else {
                            v18 -= v18
                            v39 := GenS4{F0: int16(fi28)}
                            _ = v39
                        }
                    }
                    fmt.Println(v26)
                    v40 := GenS2{F0: (v2 * 736), F1: (^(v2)), F2: uint(v18)}
                    _ = fi28
                    _ = v40
                }
                for fi41 := 0; fi41 < 15; fi41++ {
                    v42 := GenS3{F0: -895, F1: 121, F2: 0}
                    _ = fi41
                    _ = v42
                }
            }
            _ = fi27
        }
        _ = v25
        _ = v26
    }
    v2 -= (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v6
}
