package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 int16; F2 uint16 }
func genFn0() int {
    v7 := GenS1{F0: 42}
    var v8 int = 9223372036854775807
    v9 := []int{}
    switch (("" + "a") + ("hello" + ("a" + "foo bar"))) {
    case "a":
        var v10 uint16 = 42
        _ = v10
    case "":
        for fi11 := 0; fi11 < 16; fi11++ {
            v12 := GenS2{F0: 42, F1: 42}
            v8 = (fi11 / ((128) | 1))
            for fi13 := 0; fi13 < 17; fi13++ {
                v8 += 127
                fmt.Println(fi13)
                _ = fi13
            }
            _ = fi11
            _ = v12
        }
    case "hello":
        var v14 float64 = 1.5
        _ = v14
    }
    _ = v7
    _ = v8
    _ = v9
    return 9223372036854775807
}
func genFn1(p0 int16) uint {
    switch -821 {
    case 32767:
        p0 -= 0
    default:
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 float32 = -2.25
            v16 = 0.1
            v16 = v16
            _ = fi15
            _ = v16
        }
    }
    var v17 float64 = 1e10
    _ = p0
    _ = v17
    return 4294967295
}

func main() {
    v1 := []uint32{}
    var v2 uint32 = 357
    var v3 float32 = 0.0
    v4 := make(map[string]string)
    var v5 string = "a"
    var v6 uint8 = 0
    for fi18 := 0; fi18 < 15; fi18++ {
        switch v5 {
        case "foo bar":
            var v19 string = (v5 + ((v5 + "\n\t") + "\n\t"))
            _ = v19
        case "":
            var v20 uint8 = (v6 - v6)
            _ = v20
        case "\n\t":
            var v21 int32 = 1
            _ = v21
        default:
            if (false || true) {
                var v22 uint8 = 186
                var v23 uint32 = (v2 - (v2 - v2))
                _ = v22
                _ = v23
            }
        }
        var v24 int16 = 0
        _ = fi18
        _ = v24
    }
    if true {
        if ((false && (v5 >= (v5 + "foo bar"))) && true) {
            v25 := GenS3{F0: 0}
            switch genFn0() {
            case -2147483648:
                if (2147483647 == 2147483647) {
                    fmt.Println(v3, v3)
                    switch -9223372036854775808 {
                    case 9223372036854775807:
                        var v26 uint8 = v6
                        _ = v26
                    case 42:
                        var v27 int = -2147483648
                        _ = v27
                    case -1:
                        if (361 <= 1) {
                            if true {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v28 int = -991
                                var v29 int64 = 927
                                _ = v28
                                _ = v29
                            } else {
                                var v30 int64 = 291
                                _ = v30
                            }
                        }
                    }
                } else {
                    switch -9223372036854775808 {
                    case 1:
                        v31 := GenS4{F0: 2, F1: 42, F2: 65535}
                        _ = v31
                    default:
                        switch v5 {
                        case "foo bar":
                            fmt.Println(v5, v3)
                        case "\n\t":
                            var v32 bool = (v5 >= "a")
                            _ = v32
                        }
                    }
                    switch v5 {
                    case "foo bar":
                        if (((-9223372036854775808 < 128) && ((("foo bar" + "a") + ("\n\t" + "hello")) >= v5)) && (!(0 <= 2147483647))) {
                            _ = len(v4)
                        } else {
                            for fi33 := 0; fi33 < 11; fi33++ {
                                if (true || (1000 > -32768)) {
                                    var v34 string = (("" + "a") + v5)
                                    fmt.Println(v3)
                                    switch v5 {
                                    case "\n\t":
                                        delete(v4, v5)
                                    default:
                                        var v35 int16 = 32767
                                        _ = v35
                                    }
                                    _ = v34
                                }
                                _ = fi33
                            }
                            v6 = v6
                            if len(v1) > 0 {
                                v1[0] = v2
                            }
                        }
                    default:
                        if (((-2.25 + -2.25) / 1.0) < 0.0) {
                            v5 = ((v5 + (v5 + v5)) + (v5 + "\n\t"))
                            v3 = v3
                        }
                    }
                }
            case 32767:
                var v36 int8 = -127
                _ = v36
            case 42:
                _ = len(v4)
            default:
                switch v5 {
                case "hello":
                    for fi37 := 0; fi37 < 7; fi37++ {
                        v5 = (v5 + v5)
                        switch ((("a" + v5) + v5) + v5) {
                        case "hello":
                            v4[v5] = v5
                        default:
                            v38 := GenS1{F0: 1}
                            _ = v38
                        }
                        for fi39 := 0; fi39 < 18; fi39++ {
                            fmt.Println(v6)
                            _ = fi39
                        }
                        _ = fi37
                    }
                case "":
                    fmt.Println(v5, v6)
                case "foo bar":
                    if true {
                        fmt.Println(v3, v3)
                        var v40 int64 = 0
                        _ = v40
                    } else {
                        fmt.Println(v6, v2)
                        v6 = v6
                    }
                default:
                    v4[v5] = v5
                }
            }
            fmt.Println(v3, v6, v6)
            _ = v25
        }
        v2 -= v2
        if true {
            v4[""] = v5
        } else {
            for fi41 := 0; fi41 < 8; fi41++ {
                v42 := GenS2{F0: 599, F1: 22}
                _ = fi41
                _ = v42
            }
            v43 := GenS4{F0: 1, F1: 219, F2: 42}
            _ = v43
        }
    }
    v44 := GenS1{F0: 607}
    v3 = float32((0.0 * 1.5))
    var v45 int32 = 0
    var v46 uint8 = v6
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (!false) {
        var v47 int16 = -1
        _ = v47
    } else {
        for fi48 := 0; fi48 < 8; fi48++ {
            v1 = append(v1, 2147483647)
            _ = fi48
        }
        _ = len(v1)
    }
    v45 += v45
    fmt.Println(v46, v2, v6)
    fmt.Println(v6)
    if (-981 != v45) {
        for fi49 := 0; fi49 < 6; fi49++ {
            if (v3 >= (v3 * 0.0)) {
                _ = len(v1)
            } else {
                var v50 string = (v5 + v5)
                _ = v50
            }
            v46 = 128
            _ = fi49
        }
    }
    v45 -= v45
    var v51 int64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v51
    _ = v6
}
