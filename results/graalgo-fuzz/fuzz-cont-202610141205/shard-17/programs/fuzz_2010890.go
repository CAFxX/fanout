package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int32 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int }

func main() {
    var v1 uint16 = 2
    var v2 int32 = 42
    v3 := []uint32{uint32(v1), 4294967295}
    var v4 int8 = 8
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 uint8 = 127
        _ = fi5
        _ = v6
    }
    v4 += v4
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v7 := GenS1{F0: (false || (187 != 231)), F1: -1, F2: v2}
        _ = v7
    } else {
        var v8 int64 = -1
        switch "hello" {
        case "hello":
            var v9 int64 = -1
            _ = v9
        case "foo bar":
            for fi10 := 0; fi10 < 9; fi10++ {
                if (4294967295 == 42) {
                    var v11 int8 = v4
                    _ = v11
                }
                _ = fi10
            }
        case "":
            if false {
                v4 = v4
            }
        default:
            fmt.Println(v1, v4)
        }
        if len(v3) > 0 {
            v3[0] = 495
        }
        _ = v8
    }
    var v12 uint64 = 254
    switch (("\n\t" + "a") + "\n\t") {
    case "":
        v13 := GenS1{F0: (428 < (v1 | v1)), F1: -407, F2: (v2 >> ((4294967295) & 63))}
        _ = v13
    case "foo bar":
        switch -9223372036854775808 {
        case 127:
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v14 := GenS1{F0: ((!false) && (!(669 > (^(v1))))), F1: 1000, F2: v2}
                _ = v14
            } else {
                var v15 uint8 = 42
                _ = v15
            }
        case -594:
            v16 := GenS1{F0: true, F1: 32, F2: -70}
            _ = v16
        default:
            for fi17 := 0; fi17 < 7; fi17++ {
                fmt.Println(v2)
                v18 := GenS2{F0: false, F1: ("foo bar" + (("" + "") + "a"))}
                var v19 int8 = v4
                _ = fi17
                _ = v18
                _ = v19
            }
        }
    case "hello":
        if true {
            for fi20 := 0; fi20 < 10; fi20++ {
                for fi21 := 0; fi21 < 14; fi21++ {
                    switch 4294967295 {
                    case 612:
                        var v22 uint16 = uint16(v12)
                        _ = v22
                    default:
                        if true {
                            _ = len(v3)
                            _ = len(v3)
                        } else {
                            v12 = (v12 / ((v12) | 1))
                            fmt.Println(fi20)
                        }
                    }
                    var v23 uint32 = uint32(v1)
                    var v24 float64 = -2.25
                    fmt.Println(v4, v12)
                    _ = fi21
                    _ = v23
                    _ = v24
                }
                for fi25 := 0; fi25 < 6; fi25++ {
                    v1 = (v1 << ((1) & 63))
                    switch fi20 {
                    case -582:
                        switch fi25 {
                        case -2147483648:
                            if len(v3) > 0 {
                                v3[0] = 260
                            }
                        }
                        if ((true || ((0.0 - float64((0.1 * -2.25))) <= 1.5)) || (42 >= 0)) {
                            v2 -= (v2 & v2)
                            if (42 <= 42) {
                                if ((true || false) && true) {
                                    v12 = (v12 << ((767) & 63))
                                }
                                v1 -= 2
                                var v26 int8 = (v4 | v4)
                                v27 := GenS2{F0: (434 >= (^(fi20))), F1: "hello"}
                                var v28 uint = uint(v1)
                                if true {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = 42
                                    }
                                    switch ("hello" + "a") {
                                    case "a":
                                        v1 = 42
                                    case "hello":
                                        v29 := GenS3{F0: 42}
                                        _ = v29
                                    case "":
                                        v30 := GenS4{F0: 0, F1: v1}
                                        _ = v30
                                    }
                                    fmt.Println(v2, v26, v4)
                                    fmt.Println(fi25, v28, v1)
                                    if false {
                                        switch int(v28) {
                                        case 32767:
                                            v12 = v12
                                        case 127:
                                            v26 = (v4 & (v26 - int8(v2)))
                                        }
                                        fmt.Println(v28, v26)
                                    }
                                    v1 = v1
                                }
                                _ = v26
                                _ = v27
                                _ = v28
                            } else {
                                var v31 float64 = (0.0 + 3.14159)
                                if (v31 > v31) {
                                    fmt.Println(v12, v2)
                                    fmt.Println(v4)
                                } else {
                                    fmt.Println(v2, v12)
                                    v1 -= (v1 + v1)
                                    v31 = v31
                                }
                                v32 := GenS5{F0: 2, F1: 417}
                                _ = v31
                                _ = v32
                            }
                            switch "hello" {
                            case "foo bar":
                                var v33 float32 = 0.0
                                _ = v33
                            }
                        } else {
                            if false {
                                switch "a" {
                                case "\n\t":
                                    var v34 int32 = (v2 >> ((4294967295) & 63))
                                    _ = v34
                                case "foo bar":
                                    v2 += v2
                                case "hello":
                                    _ = len(v3)
                                }
                                var v35 int32 = (v2 % (((v2 * int32(v1))) | 1))
                                v35 -= v35
                                _ = v35
                            } else {
                                var v36 int8 = (v4 + (v4 << ((18446744073709551615) & 63)))
                                fmt.Println(v2, fi25)
                                _ = v36
                            }
                            fmt.Println(fi25, v1, v1)
                            var v37 int16 = 32767
                            _ = v37
                        }
                    case 32768:
                        v38 := GenS1{F0: false, F1: 0, F2: (-(v2))}
                        _ = v38
                    default:
                        switch fi20 {
                        case -9223372036854775808:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case -561:
                            var v39 uint32 = 42
                            _ = v39
                        default:
                            v40 := GenS1{F0: ((v4 >> ((255) & 63)) <= (v4 % ((v4) | 1))), F1: -32768, F2: (^(v2))}
                            _ = v40
                        }
                    }
                    _ = fi25
                }
                _ = fi20
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            switch 42 {
            case 283:
                var v41 int64 = 727
                _ = v41
            case 9223372036854775807:
                v42 := GenS6{F0: -1, F1: -1}
                _ = v42
            case -784:
                _ = len(v3)
            }
            fmt.Println(v1)
        }
    default:
        v43 := GenS7{F0: ((!true) || (!true))}
        _ = v43
    }
    var v44 bool = (v2 != (v2 * 0))
    if (2147483647 <= 432) {
        v45 := GenS1{F0: false, F1: 42, F2: -822}
        _ = v45
    } else {
        v2 = -2147483648
        var v46 int32 = v2
        _ = v46
    }
    switch 127 {
    case 9223372036854775807:
        v47 := GenS8{F0: -2147483648}
        _ = v47
    }
    for fi48 := 0; fi48 < 19; fi48++ {
        var v49 float64 = (3.14159 / 3.14159)
        var v50 float64 = (v49 + v49)
        _ = fi48
        _ = v49
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v44
}
