package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    var v3 uint32 = 1
    var v4 uint16 = 2
    var v5 int16 = -1
    var v6 float64 = float64(float32(-2.25))
    var v7 int16 = v5
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: (v7 * (v7 % (((v5 << ((42) & 63))) | 1))), F1: v5, F2: 338}
        _ = fi8
        _ = v9
    }
    v10 := []int{}
    v6 = v6
    fmt.Println(v7, v6)
    for fi11 := 0; fi11 < 14; fi11++ {
        if false {
            v5 += v7
        }
        v12 := GenS1{F0: v5, F1: 1000, F2: 65535}
        _ = fi11
        _ = v12
    }
    var v13 int16 = (v5 >> ((0) & 63))
    switch 0 {
    case -2147483648:
        v13 = v5
    case -9223372036854775808:
        delete(v1, 0)
    default:
        var v14 float64 = v6
        _ = v14
    }
    if true {
        delete(v2, 1)
        for fi15 := 0; fi15 < 15; fi15++ {
            if true {
                if (!true) {
                    v16 := GenS1{F0: (-(v5)), F1: v13, F2: 255}
                    if (203 < 42) {
                        v13 -= (-(v13))
                        if len(v10) > 0 {
                            v10[0] = fi15
                        }
                        _ = len(v10)
                    } else {
                        var v17 uint32 = 65535
                        _ = v17
                    }
                    _ = v16
                }
                var v18 bool = false
                _ = v18
            }
            _ = fi15
        }
    }
    var v19 uint8 = 2
    v6 = 0.0
    fmt.Println(v5, v19)
    if true {
        if (true || false) {
            _ = len(v10)
            if (false && true) {
                fmt.Println(v4)
                if (1 < 1) {
                    var v20 uint8 = (^(v19))
                    _ = v20
                }
            } else {
                if false {
                    if ((v13 / ((-121) | 1)) >= v7) {
                        v4 = (v4 * 65535)
                    }
                    var v21 uint16 = v4
                    _ = v21
                } else {
                    _ = len(v10)
                    if false {
                        v22 := GenS2{F0: -984}
                        _ = v22
                    } else {
                        if len(v10) > 0 {
                            v10[0] = 9223372036854775807
                        }
                        var v23 int16 = v7
                        for fi24 := 0; fi24 < 18; fi24++ {
                            for fi25 := 0; fi25 < 4; fi25++ {
                                switch int(v4) {
                                case -851:
                                    if (321 != 2) {
                                        var v26 string = "hello"
                                        _ = v26
                                    }
                                case 127:
                                    v27 := GenS2{F0: 862}
                                    _ = v27
                                }
                                var v28 string = ("a" + "a")
                                switch (^(fi24)) {
                                case -1:
                                    var v29 string = (v28 + v28)
                                    var v30 int64 = 9223372036854775807
                                    _ = v29
                                    _ = v30
                                }
                                _ = fi25
                                _ = v28
                            }
                            switch (fi24 ^ (-(fi24))) {
                            case 42:
                                if true {
                                    var v31 int16 = -1
                                    v23 = (^(v23))
                                    _ = v31
                                }
                            case 433:
                                var v32 bool = (((v4 * v4) <= (v4 - v4)) || (!true))
                                _ = v32
                            case -795:
                                var v33 float32 = 3.14159
                                _ = v33
                            }
                            _ = fi24
                        }
                        _ = v23
                    }
                }
            }
        } else {
            v5 -= v5
            fmt.Println(v5)
            v2[1] = (42 > 18446744073709551615)
        }
        fmt.Println(v19, v19)
        v34 := GenS3{F0: "hello"}
        _ = v34
    } else {
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    }
    fmt.Println(v5, v3)
    for fi35 := 0; fi35 < 1; fi35++ {
        v7 -= (v5 << ((2) & 63))
        switch "foo bar" {
        case "hello":
            if false {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                v5 = v7
            }
        case "\n\t":
            var v36 float64 = v6
            _ = v36
        }
        if ((fi35 - fi35) >= fi35) {
            if (4294967295 != 356) {
                for fi37 := 0; fi37 < 15; fi37++ {
                    var v38 int32 = 2147483647
                    if false {
                        v2[32767] = true
                        var v39 uint8 = 255
                        _ = v39
                    }
                    _ = fi37
                    _ = v38
                }
                if (!true) {
                    v40 := GenS2{F0: (fi35 * int(v3))}
                    for fi41 := 0; fi41 < 8; fi41++ {
                        v42 := GenS4{F0: 108}
                        v43 := GenS4{F0: v19}
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                    _ = v40
                } else {
                    v44 := GenS5{F0: (v19 >> ((1) & 63))}
                    v5 -= (v5 - (v13 >> ((4294967295) & 63)))
                    if true {
                        var v45 int64 = 0
                        v46 := GenS4{F0: v19}
                        _ = v45
                        _ = v46
                    }
                    _ = v44
                }
                _ = len(v2)
            } else {
                v2[fi35] = (v4 == (^(v4)))
                var v47 int16 = v7
                _ = v47
            }
            var v48 int = fi35
            v5 = v7
            _ = v48
        } else {
            fmt.Println(v4, v19, v4)
        }
        _ = fi35
    }
    if (v6 == 1.5) {
        fmt.Println(v5)
        var v49 string = (("hello" + ("foo bar" + "")) + (("a" + "a") + ""))
        _ = len(v2)
        _ = v49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
