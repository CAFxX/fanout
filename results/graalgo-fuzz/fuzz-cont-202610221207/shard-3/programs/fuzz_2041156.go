package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int; F1 bool; F2 uint16 }
type GenS4 struct { F0 int32 }

func main() {
    var v1 uint16 = 0
    v2 := []string{}
    var v3 float64 = -2.25
    var v4 uint16 = v1
    var v5 float32 = 1.5
    for fi6 := 0; fi6 < 7; fi6++ {
        if false {
            v7 := GenS1{F0: 390, F1: 42, F2: 0}
            v3 = v3
            v1 = v1
            _ = v7
        }
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
        _ = fi6
    }
    if true {
        if (-1 != -9223372036854775808) {
            if ((float64(0.0) < v3) && true) {
                switch -326 {
                case 1:
                    var v8 float32 = (1.5 - 3.14159)
                    _ = v8
                case 32767:
                    switch 127 {
                    case -259:
                        if (false || (255 != 42)) {
                            v1 = (v4 & (v1 + (v4 - (v4 << ((763) & 63)))))
                            switch 1 {
                            case 127:
                                if ((507 < 42) && (false && (false && true))) {
                                    switch "hello" {
                                    case "":
                                        var v9 uint = 137
                                        _ = v9
                                    default:
                                        var v10 int64 = 133
                                        _ = v10
                                    }
                                    v1 += 65535
                                }
                            default:
                                _ = len(v2)
                            }
                        }
                    default:
                        switch 4294967295 {
                        case 1:
                            for fi11 := 0; fi11 < 13; fi11++ {
                                v5 += 1.5
                                _ = fi11
                            }
                        case 127:
                            _ = len(v2)
                        case -9223372036854775808:
                            if (!true) {
                                if (1000 == -32768) {
                                    v2 = append(v2, "foo bar")
                                    if (!true) {
                                        if ((("hello" + "") + "\n\t") > (("" + (("\n\t" + "a") + "a")) + "a")) {
                                            v4 += 32767
                                        }
                                        if false {
                                            v3 = v3
                                            for fi12 := 0; fi12 < 17; fi12++ {
                                                var v13 bool = true
                                                if v13 {
                                                    var v14 bool = (v13 || (v13 && (585 < 680)))
                                                    var v15 bool = true
                                                    var v16 uint16 = v1
                                                    _ = v14
                                                    _ = v15
                                                    _ = v16
                                                } else {
                                                    v17 := GenS1{F0: 189, F1: 255, F2: 42}
                                                    for fi18 := 0; fi18 < 15; fi18++ {
                                                        var v19 uint32 = 2147483648
                                                        v1 += v4
                                                        _ = fi18
                                                        _ = v19
                                                    }
                                                    _ = v17
                                                }
                                                for fi20 := 0; fi20 < 14; fi20++ {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    _ = fi20
                                                }
                                                _ = fi12
                                                _ = v13
                                            }
                                        }
                                    }
                                    for fi21 := 0; fi21 < 8; fi21++ {
                                        fmt.Println(v3)
                                        _ = fi21
                                    }
                                }
                                switch 32767 {
                                case 32767:
                                    var v22 uint16 = v4
                                    _ = v22
                                }
                                v5 += v5
                            } else {
                                v23 := GenS2{F0: 1, F1: int64(v1)}
                                v24 := GenS3{F0: -774, F1: false, F2: v4}
                                _ = v23
                                _ = v24
                            }
                        }
                    }
                }
                v25 := GenS2{F0: 1, F1: 1}
                _ = v25
            } else {
                var v26 uint = 299
                for fi27 := 0; fi27 < 9; fi27++ {
                    v3 = float64(v5)
                    _ = fi27
                }
                _ = v26
            }
        }
        for fi28 := 0; fi28 < 8; fi28++ {
            fmt.Println(v4)
            _ = fi28
        }
        v4 = 65535
    }
    fmt.Println(v5, v1, v4)
    v29 := GenS3{F0: 32768, F1: (-9223372036854775808 <= -9223372036854775808), F2: 1}
    if (true || false) {
        var v30 uint16 = v4
        for fi31 := 0; fi31 < 10; fi31++ {
            v4 += 1
            _ = fi31
        }
        _ = v30
    } else {
        for fi32 := 0; fi32 < 11; fi32++ {
            var v33 uint64 = 2
            v33 -= v33
            _ = len(v2)
            _ = fi32
            _ = v33
        }
        if (1 == 1000) {
            v3 = 1e10
        }
        v34 := GenS4{F0: 239}
        _ = v34
    }
    v3 = v3
    for fi35 := 0; fi35 < 11; fi35++ {
        var v36 int16 = 1000
        _ = fi35
        _ = v36
    }
    switch -320 {
    case 0:
        v37 := GenS3{F0: 2147483647, F1: (4294967295 <= 4294967295), F2: (v1 + 0)}
        _ = v37
    default:
        v4 -= v1
    }
    _ = len(v2)
    _ = len(v2)
    var v38 int64 = -494
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
