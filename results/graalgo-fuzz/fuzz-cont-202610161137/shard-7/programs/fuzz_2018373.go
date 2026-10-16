package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint8; F1 string }
type GenS8 struct { F0 int8; F1 string; F2 uint8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS11 struct { F0 uint }

func main() {
    var v1 int8 = 21
    var v2 uint16 = 2
    var v3 uint16 = (v2 << ((255) & 63))
    v4 := []int{128, 912, -257}
    var v5 bool = false
    if false {
        fmt.Println(v2, v1)
        if v5 {
            v2 = (^(v3))
        } else {
            fmt.Println(v2, v2, v3)
            switch "hello" {
            case "\n\t":
                if v5 {
                    fmt.Println(v5, v1)
                    fmt.Println(v5, v3, v5)
                } else {
                    var v6 uint16 = (v2 * v2)
                    v5 = v5
                    _ = v6
                }
            case "foo bar":
                v1 = (v1 / ((v1) | 1))
            }
            fmt.Println(v3, v1)
        }
    } else {
        var v7 int = 415
        v1 = (v1 ^ 100)
        v8 := GenS1{F0: 317, F1: v3, F2: v5}
        _ = v7
        _ = v8
    }
    v2 = v2
    v9 := GenS2{F0: -972}
    if len(v4) > 0 {
        v4[0] = 32767
    }
    if (1 >= 773) {
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 uint32 = 981
            for fi12 := 0; fi12 < 9; fi12++ {
                v13 := GenS2{F0: -9223372036854775808}
                _ = fi12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
        for fi14 := 0; fi14 < 3; fi14++ {
            v1 = (v1 % ((-128) | 1))
            fmt.Println(v1)
            _ = fi14
        }
    } else {
        v1 = v1
        _ = len(v4)
    }
    var v15 float64 = -2.25
    v1 -= v1
    if v5 {
        if v5 {
            var v16 uint8 = 255
            _ = v16
        } else {
            if v5 {
                switch 127 {
                case 42:
                    if true {
                        if ((v1 >> ((314) & 63)) > v1) {
                            if (v5 || (-2.25 < (3.14159 - 0.1))) {
                                for fi17 := 0; fi17 < 9; fi17++ {
                                    v1 += -128
                                    _ = fi17
                                }
                            }
                            _ = len(v4)
                        } else {
                            var v18 float32 = (1e10 * -2.25)
                            _ = v18
                        }
                        v19 := GenS3{F0: 4294967295, F1: false}
                        v15 = 1.5
                        _ = v19
                    }
                case 4294967295:
                    var v20 uint = 0
                    _ = v20
                case 128:
                    if v5 {
                        if (!v5) {
                            if (42 <= 42) {
                                var v21 uint = 118
                                fmt.Println(v15)
                                switch "a" {
                                case "":
                                    var v22 float32 = float32(v15)
                                    _ = v22
                                case "foo bar":
                                    v3 -= v2
                                case "\n\t":
                                    var v23 string = ""
                                    _ = v23
                                default:
                                    switch 755 {
                                    case 32767:
                                        switch "" {
                                        case "a":
                                            v24 := GenS4{F0: 2147483647, F1: 124}
                                            _ = v24
                                        case "":
                                            for fi25 := 0; fi25 < 5; fi25++ {
                                                switch "a" {
                                                case "a":
                                                    switch 42 {
                                                    case 0:
                                                        v26 := GenS5{F0: 0}
                                                        _ = v26
                                                    }
                                                case "":
                                                    v27 := GenS6{F0: v1}
                                                    _ = v27
                                                case "foo bar":
                                                    v21 -= uint(v3)
                                                default:
                                                    v28 := GenS7{F0: 1, F1: ""}
                                                    _ = v28
                                                }
                                                _ = fi25
                                            }
                                        case "\n\t":
                                            v5 = false
                                        }
                                    case -800:
                                        v2 = (v2 - v3)
                                    case 1:
                                        for fi29 := 0; fi29 < 8; fi29++ {
                                            v2 += (v2 ^ v3)
                                            _ = fi29
                                        }
                                    default:
                                        if v5 {
                                            var v30 int16 = 1
                                            switch "" {
                                            case "\n\t":
                                                fmt.Println(v15, v30)
                                            default:
                                                v1 -= (v1 ^ (v1 ^ v1))
                                            }
                                            _ = v30
                                        } else {
                                            v4 = append(v4, 32768)
                                        }
                                    }
                                }
                                _ = v21
                            } else {
                                v1 = (^(v1))
                                for fi31 := 0; fi31 < 6; fi31++ {
                                    v32 := GenS8{F0: v1, F1: "foo bar", F2: 2}
                                    switch -616 {
                                    case 128:
                                        fmt.Println(v1)
                                    case -377:
                                        switch (fi31 | (^(fi31))) {
                                        case 9223372036854775807:
                                            v33 := GenS9{F0: 0}
                                            _ = v33
                                        default:
                                            fmt.Println(v1, v5)
                                        }
                                    case -1:
                                        var v34 int16 = -32768
                                        _ = v34
                                    default:
                                        if true {
                                            if ("hello" > "foo bar") {
                                                for fi35 := 0; fi35 < 20; fi35++ {
                                                    var v36 string = "hello"
                                                    var v37 uint8 = 128
                                                    _ = fi35
                                                    _ = v36
                                                    _ = v37
                                                }
                                            } else {
                                                _ = len(v4)
                                                v38 := GenS10{F0: 0, F1: 42, F2: 18446744073709551615}
                                                v39 := GenS10{F0: -1, F1: 2, F2: 42}
                                                _ = v38
                                                _ = v39
                                            }
                                            if (2 < 255) {
                                                var v40 float64 = float64(1.5)
                                                for fi41 := 0; fi41 < 11; fi41++ {
                                                    v4 = append(v4, fi41)
                                                    fmt.Println(fi41)
                                                    v2 += v3
                                                    _ = fi41
                                                }
                                                _ = v40
                                            } else {
                                                if len(v4) > 0 {
                                                    v4[0] = fi31
                                                }
                                                v3 = (v2 >> ((1000) & 63))
                                                if (v2 >= v3) {
                                                    var v42 int8 = v1
                                                    fmt.Println(v1)
                                                    _ = v42
                                                }
                                            }
                                            for fi43 := 0; fi43 < 9; fi43++ {
                                                v44 := GenS8{F0: v1, F1: "\n\t", F2: 209}
                                                _ = fi43
                                                _ = v44
                                            }
                                        } else {
                                            fmt.Println(v15)
                                            if len(v4) > 0 {
                                                v4[0] = (fi31 << ((255) & 63))
                                            }
                                            v2 = 65535
                                        }
                                    }
                                    _ = fi31
                                    _ = v32
                                }
                                v3 += 746
                            }
                            v45 := GenS11{F0: 2}
                            _ = v45
                        }
                    } else {
                        if (308 <= 427) {
                            fmt.Println(v1, v3, v5)
                            if true {
                                for fi46 := 0; fi46 < 4; fi46++ {
                                    v47 := GenS11{F0: 4294967295}
                                    _ = fi46
                                    _ = v47
                                }
                            } else {
                                var v48 uint64 = 42
                                var v49 int64 = 1
                                v50 := GenS3{F0: 42, F1: v5}
                                _ = v48
                                _ = v49
                                _ = v50
                            }
                            switch ("\n\t" + "a") {
                            case "":
                                for fi51 := 0; fi51 < 10; fi51++ {
                                    fmt.Println(fi51, fi51)
                                    _ = fi51
                                }
                            }
                        }
                        for fi52 := 0; fi52 < 14; fi52++ {
                            v1 -= 1
                            _ = fi52
                        }
                        var v53 bool = v5
                        _ = v53
                    }
                }
            }
        }
    }
    var v54 uint64 = 18446744073709551615
    if (false || v5) {
        v2 -= 255
        v55 := GenS4{F0: 610, F1: 1}
        _ = v55
    }
    v1 = -128
    v3 -= (v2 | v3)
    _ = len(v4)
    v56 := GenS1{F0: 1, F1: v3, F2: (false && v5)}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v56
    _ = v9
}
