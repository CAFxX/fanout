package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 int64; F2 uint32 }
type GenS4 struct { F0 int32; F1 int8; F2 string }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 int32 }

func main() {
    var v1 uint = 65535
    v2 := []string{"a"}
    var v3 int32 = -2147483648
    if (106 > int8(v1)) {
        var v4 int64 = 42
        fmt.Println(v4)
        var v5 uint = v1
        _ = v4
        _ = v5
    } else {
        var v6 int64 = 9223372036854775807
        v7 := GenS1{F0: -84, F1: 42}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v6
        _ = v7
    }
    var v8 float32 = 1e10
    for fi9 := 0; fi9 < 16; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            fmt.Println(v1)
            _ = fi10
        }
        _ = fi9
    }
    fmt.Println(v3, v3)
    v11 := GenS1{F0: 100, F1: 353}
    v1 = v1
    v3 = (v3 & -1)
    var v12 int64 = 1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        if false {
            var v13 int32 = v3
            v12 += (v12 << (((v1 % ((v1) | 1))) & 63))
            _ = v13
        } else {
            var v14 uint64 = 2
            if (2147483648 == 2147483648) {
                var v15 int64 = v12
                for fi16 := 0; fi16 < 10; fi16++ {
                    var v17 int = (-(fi16))
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi16
                    _ = v17
                }
                v18 := GenS2{F0: -123}
                _ = v15
                _ = v18
            } else {
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
                v14 -= v14
            }
            var v19 bool = false
            _ = v14
            _ = v19
        }
        if ((300 < 42) || true) {
            if (-1 >= 9223372036854775807) {
                if (true && false) {
                    switch 42 {
                    case -1:
                        if true {
                            v12 += (v12 ^ v12)
                            v20 := GenS1{F0: -1, F1: 42}
                            _ = v20
                        }
                    case 0:
                        _ = len(v2)
                    default:
                        if (32767 < 1) {
                            v21 := GenS1{F0: 127, F1: 1}
                            var v22 float64 = 1.5
                            var v23 int16 = 0
                            _ = v21
                            _ = v22
                            _ = v23
                        } else {
                            fmt.Println(v8, v3)
                            if false {
                                v24 := GenS3{F0: true, F1: -1, F2: uint32(v1)}
                                _ = v24
                            }
                            _ = len(v2)
                        }
                    }
                    if true {
                        var v25 uint8 = 255
                        _ = len(v2)
                        if (1 > 0) {
                            fmt.Println(v8)
                            switch 2147483647 {
                            case 128:
                                var v26 int16 = -32768
                                _ = v26
                            }
                            for fi27 := 0; fi27 < 11; fi27++ {
                                v12 = v12
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = fi27
                            }
                        } else {
                            switch "\n\t" {
                            case "foo bar":
                                for fi28 := 0; fi28 < 13; fi28++ {
                                    if len(v2) > 0 {
                                        v2[0] = ""
                                    }
                                    var v29 float32 = (1.5 / v8)
                                    _ = fi28
                                    _ = v29
                                }
                            case "hello":
                                var v30 float64 = 0.1
                                _ = v30
                            default:
                                fmt.Println(v8)
                            }
                            v31 := GenS4{F0: 0, F1: 42, F2: ""}
                            v25 += v25
                            _ = v31
                        }
                        _ = v25
                    }
                    v3 -= v3
                }
            } else {
                fmt.Println(v1)
                switch -2147483648 {
                case 4294967295:
                    if (736 <= -838) {
                        v3 = v3
                        v32 := GenS5{F0: ""}
                        var v33 int32 = (-(v3))
                        _ = v32
                        _ = v33
                    }
                case 2147483647:
                    for fi34 := 0; fi34 < 4; fi34++ {
                        v35 := GenS6{F0: fi34, F1: 30}
                        _ = fi34
                        _ = v35
                    }
                case -677:
                    if true {
                        v12 -= v12
                        switch ("" + "") {
                        case "foo bar":
                            for fi36 := 0; fi36 < 10; fi36++ {
                                v37 := GenS7{F0: -2147483648}
                                _ = fi36
                                _ = v37
                            }
                        case "":
                            for fi38 := 0; fi38 < 20; fi38++ {
                                _ = len(v2)
                                _ = fi38
                            }
                        case "a":
                            var v39 int16 = -32768
                            _ = v39
                        default:
                            if false {
                                fmt.Println(v1)
                                v3 += v3
                                v40 := GenS1{F0: 77, F1: 42}
                                _ = v40
                            }
                        }
                    }
                }
            }
            for fi41 := 0; fi41 < 3; fi41++ {
                if true {
                    if true {
                        fmt.Println(v12)
                        for fi42 := 0; fi42 < 18; fi42++ {
                            v1 -= (v1 * (v1 + (v1 | 42)))
                            _ = fi42
                        }
                        var v43 int8 = -128
                        _ = v43
                    }
                    for fi44 := 0; fi44 < 16; fi44++ {
                        var v45 uint16 = 32767
                        _ = fi44
                        _ = v45
                    }
                }
                _ = fi41
            }
        }
        var v46 uint16 = 255
        _ = v46
    }
    fmt.Println(v3, v8)
    switch -549 {
    case 0:
        var v47 bool = false
        _ = v47
    case -9223372036854775808:
        var v48 float32 = ((float32((0.1 + 1e10)) + v8) - v8)
        _ = v48
    }
    var v49 int8 = 0
    v50 := GenS3{F0: false, F1: v12, F2: 4294967295}
    v8 = v8
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v49:", v49)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v49
    _ = v50
    _ = v8
}
