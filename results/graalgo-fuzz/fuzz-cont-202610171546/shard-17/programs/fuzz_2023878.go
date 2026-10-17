package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 int; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint32 }

func main() {
    var v1 int32 = -2147483648
    v2 := []string{("hello" + "a"), ((("foo bar" + ("" + "hello")) + "") + "foo bar")}
    var v3 uint32 = 197
    var v4 int64 = 42
    v5 := []string{"\n\t", "foo bar", (((("a" + "\n\t") + "") + "\n\t") + "foo bar"), ((("foo bar" + "hello") + "foo bar") + "a")}
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 float32 = 1.5
        var v8 uint16 = 32767
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: true, F1: 1, F2: -1}
    v10 := GenS2{F0: 2, F1: 32767}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v4 = 9223372036854775807
    v1 -= v1
    if (!false) {
        v1 += (v1 % ((v1) | 1))
    } else {
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 float32 = 3.14159
            _ = fi11
            _ = v12
        }
        fmt.Println(v1)
    }
    if true {
        for fi13 := 0; fi13 < 9; fi13++ {
            switch "\n\t" {
            case "":
                for fi14 := 0; fi14 < 8; fi14++ {
                    v3 = (v3 ^ 65535)
                    v15 := GenS3{F0: 2147483647, F1: 255, F2: v3}
                    v1 = 427
                    v4 -= -467
                    switch ("a" + "\n\t") {
                    case "":
                        v16 := GenS4{F0: (fi13 % ((fi13) | 1)), F1: 1}
                        _ = v16
                    case "foo bar":
                        var v17 int = fi13
                        _ = v17
                    }
                    _ = fi14
                    _ = v15
                }
            }
            fmt.Println(fi13, v1, v1)
            _ = fi13
        }
    }
    v1 = (v1 >> ((819) & 63))
    if true {
        if false {
            for fi18 := 0; fi18 < 17; fi18++ {
                if false {
                    switch "" {
                    case "\n\t":
                        for fi19 := 0; fi19 < 10; fi19++ {
                            fmt.Println(fi19)
                            _ = fi19
                        }
                    case "":
                        for fi20 := 0; fi20 < 10; fi20++ {
                            fmt.Println(v3, v1)
                            fmt.Println(v1, v3)
                            _ = fi20
                        }
                    }
                    for fi21 := 0; fi21 < 10; fi21++ {
                        var v22 int8 = 100
                        _ = fi21
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 14; fi23++ {
                        if (3.14159 >= 0.1) {
                            if ("a" != "hello") {
                                fmt.Println(v4)
                            }
                            if (v1 > 0) {
                                if true {
                                    v4 = v4
                                    v24 := GenS5{F0: v1}
                                    if len(v2) > 0 {
                                        v2[0] = ("hello" + "foo bar")
                                    }
                                    _ = v24
                                } else {
                                    var v25 uint = 42
                                    v26 := GenS6{F0: 196}
                                    _ = v25
                                    _ = v26
                                }
                            }
                            var v27 int32 = (v1 % ((v1) | 1))
                            _ = v27
                        } else {
                            if len(v2) > 0 {
                                v2[0] = ("a" + "hello")
                            }
                            if (18446744073709551615 == 18446744073709551615) {
                                var v28 uint = 42
                                fmt.Println(v1, fi23, fi23)
                                _ = v28
                            } else {
                                var v29 int32 = 159
                                _ = v29
                            }
                        }
                        v1 += v1
                        _ = fi23
                    }
                } else {
                    var v30 string = ("hello" + "a")
                    _ = v30
                }
                v31 := GenS5{F0: (-(v1))}
                var v32 int16 = 1
                _ = fi18
                _ = v31
                _ = v32
            }
            var v33 bool = false
            _ = v33
        }
        v4 -= (-(v4))
    }
    for fi34 := 0; fi34 < 8; fi34++ {
        v4 = (-(v4))
        v35 := GenS3{F0: 0, F1: 255, F2: 42}
        _ = fi34
        _ = v35
    }
    var v36 int16 = 0
    for fi37 := 0; fi37 < 18; fi37++ {
        if len(v2) > 0 {
            v2[0] = ("foo bar" + "hello")
        }
        var v38 int = -694
        _ = fi37
        _ = v38
    }
    switch 9223372036854775807 {
    case -201:
        if (18446744073709551615 <= 2) {
            _ = len(v5)
            fmt.Println(v36)
            v3 -= v3
        }
    case 32767:
        for fi39 := 0; fi39 < 11; fi39++ {
            if false {
                var v40 int16 = v36
                for fi41 := 0; fi41 < 11; fi41++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v42 int = fi41
                    v3 += (^(v3))
                    var v43 bool = true
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                _ = v40
            } else {
                v3 += (v3 * (v3 * (v3 % ((v3) | 1))))
            }
            switch (fi39 ^ fi39) {
            case -9223372036854775808:
                for fi44 := 0; fi44 < 2; fi44++ {
                    var v45 int16 = (v36 & 32767)
                    if ((1.5 * 1.5) < ((1e10 + 0.1) + float32(-2.25))) {
                        v2 = append(v2, "")
                        fmt.Println(fi44)
                        fmt.Println(v4)
                    } else {
                        _ = len(v2)
                        v46 := GenS5{F0: v1}
                        var v47 uint16 = 1
                        _ = v46
                        _ = v47
                    }
                    _ = fi44
                    _ = v45
                }
            case -1:
                var v48 uint8 = 42
                _ = v48
            case 32768:
                v49 := GenS2{F0: 586, F1: 0}
                _ = v49
            }
            _ = fi39
        }
    }
    _ = len(v5)
    if false {
        switch -2147483648 {
        case 32767:
            var v50 bool = false
            _ = v50
        case -9223372036854775808:
            switch "foo bar" {
            case "hello":
                for fi51 := 0; fi51 < 13; fi51++ {
                    v52 := GenS7{F0: 65535}
                    _ = fi51
                    _ = v52
                }
            }
        }
        var v53 int8 = 0
        v1 += (v1 & 1)
        _ = v53
    } else {
        v54 := GenS1{F0: (false && false), F1: 2, F2: 81}
        v2 = append(v2, "")
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v9
}
