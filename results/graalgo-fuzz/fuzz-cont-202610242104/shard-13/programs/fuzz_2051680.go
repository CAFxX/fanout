package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 string; F1 int32; F2 int32 }
type GenS3 struct { F0 int8; F1 string }
func genFn0(p0 int) int16 {
    var v5 int8 = 127
    v5 -= v5
    switch (p0 * p0) {
    case -728:
        if false {
            p0 = (p0 << ((953) & 63))
        }
    case -1:
        var v6 uint64 = 18446744073709551615
        _ = v6
    }
    _ = p0
    _ = v5
    return 0
}
func genFn1() uint64 {
    v7 := []int{-168, 0}
    var v8 int64 = 42
    _ = v7
    _ = v8
    return 18446744073709551615
}

func main() {
    var v1 float32 = 3.14159
    v2 := []string{"", "", ((("a" + "foo bar") + "hello") + "foo bar")}
    var v3 int = -590
    var v4 int64 = -965
    if (false && true) {
        if false {
            fmt.Println(v3, v1)
            var v9 uint8 = 143
            _ = v9
        }
    }
    fmt.Println(v4, v3)
    var v10 int32 = 0
    for fi11 := 0; fi11 < 8; fi11++ {
        v12 := GenS1{F0: 42, F1: 0}
        switch "" {
        case "foo bar":
            v4 -= v4
        case "hello":
            if true {
                v1 = (v1 / v1)
            } else {
                v13 := GenS1{F0: 120, F1: 852}
                if len(v2) > 0 {
                    v2[0] = ("hello" + "a")
                }
                _ = v13
            }
        default:
            switch -1 {
            case 128:
                if true {
                    var v14 string = "hello"
                    v15 := GenS1{F0: 1000, F1: 2147483648}
                    switch "" {
                    case "a":
                        var v16 uint8 = 127
                        _ = v16
                    case "hello":
                        v10 += v10
                    case "foo bar":
                        var v17 uint64 = 359
                        _ = v17
                    }
                    _ = v14
                    _ = v15
                }
            }
        }
        _ = fi11
        _ = v12
    }
    var v18 int64 = 42
    v18 += v18
    var v19 string = "hello"
    switch (v3 | (v3 ^ 208)) {
    case 128:
        v20 := GenS2{F0: v19, F1: v10, F2: (-(v10))}
        _ = v20
    case 127:
        if (!true) {
            fmt.Println(v18, v1)
            var v21 uint16 = 65535
            _ = v21
        }
    default:
        switch v19 {
        case "hello":
            var v22 float64 = ((1.5 - ((1.5 / 1.0) / 1.5)) / 0.1)
            _ = v22
        case "":
            if ((((v1 / v1) > 0.0) || false) && true) {
                switch v3 {
                case 0:
                    if true {
                        var v23 bool = (!(0.0 >= float64(v1)))
                        if false {
                            for fi24 := 0; fi24 < 9; fi24++ {
                                for fi25 := 0; fi25 < 10; fi25++ {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    var v26 uint = 2
                                    fmt.Println(v18)
                                    _ = fi25
                                    _ = v26
                                }
                                var v27 int = int(v10)
                                v28 := GenS1{F0: 1000, F1: 692}
                                _ = fi24
                                _ = v27
                                _ = v28
                            }
                            for fi29 := 0; fi29 < 7; fi29++ {
                                var v30 uint16 = 32767
                                switch (-(fi29)) {
                                case 2147483647:
                                    var v31 int16 = -32768
                                    _ = v31
                                case 1:
                                    for fi32 := 0; fi32 < 4; fi32++ {
                                        if v23 {
                                            var v33 int16 = 572
                                            _ = v33
                                        }
                                        _ = fi32
                                    }
                                default:
                                    v18 = (v4 & (v18 ^ -9223372036854775808))
                                }
                                _ = fi29
                                _ = v30
                            }
                            v10 = v10
                        }
                        v34 := GenS3{F0: 127, F1: v19}
                        _ = v23
                        _ = v34
                    } else {
                        for fi35 := 0; fi35 < 18; fi35++ {
                            v10 = (v10 - v10)
                            v1 = (v1 + 1.5)
                            _ = fi35
                        }
                    }
                case 42:
                    if (true || false) {
                        for fi36 := 0; fi36 < 14; fi36++ {
                            switch (v3 + 977) {
                            case 2147483647:
                                if len(v2) > 0 {
                                    v2[0] = v19
                                }
                            default:
                                v19 = "foo bar"
                            }
                            for fi37 := 0; fi37 < 8; fi37++ {
                                if len(v2) > 0 {
                                    v2[0] = (v19 + v19)
                                }
                                _ = fi37
                            }
                            var v38 uint = 18446744073709551615
                            _ = fi36
                            _ = v38
                        }
                        for fi39 := 0; fi39 < 5; fi39++ {
                            var v40 uint32 = 612
                            fmt.Println(v40)
                            for fi41 := 0; fi41 < 14; fi41++ {
                                if len(v2) > 0 {
                                    v2[0] = v19
                                }
                                if true {
                                    v4 -= (v4 & v18)
                                }
                                fmt.Println(v1, fi41, v40)
                                _ = fi41
                            }
                            _ = fi39
                            _ = v40
                        }
                        _ = len(v2)
                    } else {
                        fmt.Println(v3, v3)
                    }
                default:
                    fmt.Println(v18, v1, v19)
                }
            } else {
                v18 += (v4 % ((int64(v10)) | 1))
            }
        default:
            for fi42 := 0; fi42 < 8; fi42++ {
                var v43 int16 = 1000
                _ = fi42
                _ = v43
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
