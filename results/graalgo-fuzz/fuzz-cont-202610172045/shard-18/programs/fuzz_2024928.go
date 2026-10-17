package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint32; F1 int64; F2 int16 }
type GenS3 struct { F0 int8; F1 string; F2 uint }
type GenS4 struct { F0 int16; F1 uint; F2 uint16 }
type GenS5 struct { F0 string; F1 uint8 }
type GenS6 struct { F0 int32; F1 int64; F2 int64 }
type GenS7 struct { F0 uint8; F1 int }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint16; F1 int32; F2 int32 }
type GenS10 struct { F0 int8; F1 bool }
type GenS11 struct { F0 int; F1 int64 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint16; F1 int16 }
type GenS14 struct { F0 uint }
func genFn0(p0 uint64, p1 int16) int16 {
    if true {
        var v7 int16 = (-(p1))
        switch 723 {
        case -9223372036854775808:
            v8 := []int{-633}
            _ = v8
        }
        v9 := GenS1{F0: "a", F1: (false || true)}
        _ = v7
        _ = v9
    }
    switch 222 {
    case 32767:
        v10 := GenS2{F0: 2147483648, F1: 950, F2: p1}
        _ = v10
    case 1:
        var v11 bool = (42 != 4294967295)
        _ = v11
    }
    p0 = p0
    _ = p0
    _ = p1
    return p1
}

func main() {
    var v1 string = ("a" + "")
    v2 := []string{v1, v1, "\n\t"}
    var v3 uint = 4
    var v4 string = "a"
    v5 := []string{v1}
    v6 := make(map[uint32]string)
    var v12 int32 = 1
    fmt.Println(v1, v4)
    for fi13 := 0; fi13 < 5; fi13++ {
        for fi14 := 0; fi14 < 12; fi14++ {
            v15 := GenS3{F0: -119, F1: v1, F2: (v3 ^ v3)}
            var v16 int64 = -9223372036854775808
            v17 := GenS4{F0: -331, F1: (v3 ^ (v3 / ((v3) | 1))), F2: 32767}
            v4 = ("\n\t" + v4)
            v16 -= (v16 & v16)
            v18 := GenS5{F0: v4, F1: 120}
            delete(v6, 0)
            fmt.Println(v1, v3, fi14)
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
            _ = v18
        }
        var v19 int64 = 0
        for fi20 := 0; fi20 < 4; fi20++ {
            var v21 string = ("\n\t" + "\n\t")
            switch v21 {
            case "\n\t":
                v22 := GenS6{F0: 42, F1: (v19 + 0), F2: 42}
                var v23 string = "foo bar"
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(v21, v19)
                case "":
                    if true {
                        v21 = "foo bar"
                        if ((394 < (v3 * v3)) && false) {
                            var v24 uint = (v3 >> ((v3) & 63))
                            if (1.5 == (1.5 / 0.1)) {
                                switch "a" {
                                case "foo bar":
                                    v25 := GenS7{F0: 42, F1: (fi20 << ((1) & 63))}
                                    _ = v25
                                case "hello":
                                    var v26 uint32 = 4294967295
                                    _ = v26
                                }
                            } else {
                                var v27 float32 = (-2.25 * 1e10)
                                v19 -= (v19 + (v19 + v19))
                                delete(v6, 0)
                                _ = v27
                            }
                            _ = v24
                        } else {
                            switch (((v1 + ((v21 + ("a" + v4)) + ((("" + "\n\t") + (v21 + ((v4 + "foo bar") + v1))) + v1))) + v21) + "foo bar") {
                            case "\n\t":
                                var v28 int16 = 32767
                                _ = v28
                            default:
                                _ = len(v6)
                            }
                        }
                    } else {
                        var v29 int32 = 374
                        v30 := GenS8{F0: 32767}
                        _ = v29
                        _ = v30
                    }
                }
                _ = v22
                _ = v23
            case "foo bar":
                fmt.Println(v19)
            case "a":
                v31 := GenS9{F0: 65535, F1: (^(v12)), F2: (v12 / ((277) | 1))}
                if false {
                    if len(v2) > 0 {
                        v2[0] = ("\n\t" + v1)
                    }
                }
                var v32 int = fi20
                switch "" {
                case "foo bar":
                    var v33 uint64 = 908
                    _ = v33
                case "":
                    if true {
                        var v34 int64 = 0
                        var v35 uint32 = 1
                        var v36 uint64 = 0
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                }
                v19 -= (v19 ^ v19)
                _ = v31
                _ = v32
            }
            _ = fi20
            _ = v21
        }
        _ = fi13
        _ = v19
    }
    v1 = ("a" + v4)
    var v37 int32 = 0
    v38 := v6[2]
    var v39 uint16 = 0
    var v40 int32 = v12
    for fi41 := 0; fi41 < 10; fi41++ {
        if ((3.14159 - 0.0) > -2.25) {
            if (!(4294967295 <= 4294967295)) {
                if true {
                    v42 := GenS5{F0: "a", F1: 100}
                    switch fi41 {
                    case 9223372036854775807:
                        _ = len(v5)
                    case -1:
                        v43 := GenS10{F0: 100, F1: true}
                        _ = v43
                    default:
                        v44 := GenS11{F0: (-(fi41)), F1: -51}
                        _ = v44
                    }
                    _ = v42
                }
                var v45 float64 = -2.25
                _ = v45
            } else {
                v40 = v40
                v4 = v4
                if (0 > 2) {
                    var v46 int8 = 42
                    _ = v46
                } else {
                    for fi47 := 0; fi47 < 9; fi47++ {
                        fmt.Println(v40, v39, v39)
                        switch fi47 {
                        case 219:
                            fmt.Println(v38)
                        case 878:
                            v48 := GenS12{F0: v39}
                            fmt.Println(v39, fi47, v12)
                            _ = v48
                        }
                        if (fi41 != (fi41 & fi41)) {
                            var v49 int64 = int64(v39)
                            _ = v49
                        }
                        _ = fi47
                    }
                    v50 := GenS1{F0: v4, F1: (v3 >= (^(v3)))}
                    _ = v50
                }
            }
            v51 := GenS12{F0: v39}
            if len(v5) > 0 {
                v5[0] = v4
            }
            _ = v51
        } else {
            if (!(0.1 > 1.5)) {
                fmt.Println(v3, fi41)
                if ("foo bar" <= ((v38 + v4) + "\n\t")) {
                    v5 = append(v5, ((v38 + v38) + ""))
                    var v52 uint8 = 168
                    _ = v52
                }
            } else {
                var v53 int8 = -30
                var v54 uint32 = 65535
                fmt.Println(v38)
                _ = v53
                _ = v54
            }
            fmt.Println(v37, v12)
        }
        if true {
            for fi55 := 0; fi55 < 3; fi55++ {
                var v56 bool = false
                var v57 int32 = -839
                _ = fi55
                _ = v56
                _ = v57
            }
            v58 := GenS13{F0: (^(v39)), F1: 0}
            var v59 uint16 = v39
            _ = v58
            _ = v59
        } else {
            if (255 <= 1) {
                if len(v5) > 0 {
                    v5[0] = v1
                }
            }
            fmt.Println(v12, v39, v4)
        }
        fmt.Println(v4)
        _ = fi41
    }
    v60 := GenS14{F0: (v3 * 4294967295)}
    fmt.Println(v37, v4)
    v12 = -608
    v61 := GenS6{F0: v12, F1: 0, F2: -948}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v60
    _ = v61
}
