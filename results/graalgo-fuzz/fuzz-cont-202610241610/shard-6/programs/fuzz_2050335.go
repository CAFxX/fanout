package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 bool; F1 int16; F2 uint }
type GenS5 struct { F0 uint16; F1 int32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int16; F1 uint8 }
func genFn0(p0 int64) int32 {
    var v7 int64 = p0
    var v8 int16 = -475
    _ = p0
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    var v1 int = -245
    v2 := []uint32{}
    var v3 string = "foo bar"
    v4 := []uint32{0, 1, 770, 4294967295}
    v5 := make(map[string]string)
    v6 := make(map[int]bool)
    for fi9 := 0; fi9 < 18; fi9++ {
        for fi10 := 0; fi10 < 12; fi10++ {
            v11 := GenS1{F0: v3, F1: (v3 + v3)}
            if (true && false) {
                v3 = (v3 + ((((v3 + v3) + "") + v3) + v3))
            }
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    var v12 int8 = 42
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        switch (v3 + "") {
        case "\n\t":
            var v14 uint32 = 42
            _ = v14
        }
        var v15 uint16 = 662
        v1 -= 400
        _ = fi13
        _ = v15
    }
    if (42 < 65535) {
        for fi16 := 0; fi16 < 15; fi16++ {
            v17 := GenS2{F0: 2, F1: -9223372036854775808, F2: 42}
            var v18 uint64 = 4294967295
            _ = fi16
            _ = v17
            _ = v18
        }
        for fi19 := 0; fi19 < 19; fi19++ {
            for fi20 := 0; fi20 < 12; fi20++ {
                if (-32768 <= -1) {
                    var v21 uint32 = 2147483647
                    fmt.Println(v3)
                    fmt.Println(fi19)
                    var v22 uint32 = (v21 % ((0) | 1))
                    var v23 float32 = -2.25
                    _ = v21
                    _ = v22
                    _ = v23
                }
                switch v1 {
                case 32768:
                    v4 = append(v4, 335)
                case -1:
                    switch v1 {
                    case -176:
                        v6[v1] = (false || false)
                    default:
                        if (!((1.5 * (0.0 * 1e10)) > (1.5 + 1e10))) {
                            v24 := v6[fi19]
                            _ = v24
                        } else {
                            if (128 == 2) {
                                fmt.Println(v12, fi19)
                            }
                            if true {
                                switch "a" {
                                case "\n\t":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                case "":
                                    switch fi20 {
                                    case -856:
                                        if (32767 != 830) {
                                            v25 := GenS3{F0: 0, F1: "hello"}
                                            _ = v25
                                        }
                                    case -1:
                                        v26 := GenS4{F0: true, F1: -32768, F2: 18446744073709551615}
                                        _ = v26
                                    default:
                                        v3 = v3
                                    }
                                }
                                var v27 float32 = 1.5
                                v5[v3] = ("\n\t" + (v3 + "hello"))
                                _ = v27
                            } else {
                                v28 := GenS5{F0: 42, F1: 2147483647}
                                _ = v28
                            }
                        }
                        var v29 uint32 = 554
                        var v30 uint16 = 1
                        _ = v29
                        _ = v30
                    }
                case -674:
                    var v31 int8 = v12
                    if (69 != 2) {
                        var v32 uint16 = 1
                        v33 := GenS6{F0: 32768}
                        _ = v32
                        _ = v33
                    }
                    v3 = (v3 + "a")
                    fmt.Println(v1, v3)
                    var v34 uint = 0
                    _ = v31
                    _ = v34
                }
                _ = fi20
            }
            if (!false) {
                if true {
                    var v35 uint = 457
                    var v36 float32 = (float32(((3.14159 * 1.5) / (1.5 - (0.0 / (1e10 + (0.0 * 1e10)))))) / float32(3.14159))
                    _ = v35
                    _ = v36
                }
                for fi37 := 0; fi37 < 1; fi37++ {
                    fmt.Println(v1, fi37, v1)
                    v38 := v5[v3]
                    v39 := GenS4{F0: false, F1: 32767, F2: 455}
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            }
            _ = fi19
        }
        v3 = "hello"
    }
    switch v3 {
    case "\n\t":
        v40 := v6[(v1 % ((-2147483648) | 1))]
        _ = v40
    case "":
        var v41 float64 = -2.25
        _ = v41
    case "hello":
        switch v3 {
        case "a":
            var v42 float32 = -2.25
            _ = v42
        case "":
            v12 += (v12 >> ((42) & 63))
        default:
            switch ((("a" + ((v3 + v3) + v3)) + ("" + ((v3 + "a") + "foo bar"))) + v3) {
            case "\n\t":
                v4 = append(v4, 819)
            case "hello":
                switch v3 {
                case "":
                    var v43 int16 = -808
                    _ = v43
                case "foo bar":
                    v44 := v5[("a" + v3)]
                    _ = v44
                case "\n\t":
                    if true {
                        if true {
                            if false {
                                fmt.Println(v1, v1)
                                v1 = -325
                            } else {
                                v3 = "\n\t"
                            }
                            fmt.Println(v1, v3)
                        } else {
                            var v45 bool = (!true)
                            switch -557 {
                            case -2147483648:
                                fmt.Println(v45, v1, v1)
                            case -721:
                                delete(v5, v3)
                            default:
                                switch v3 {
                                case "a":
                                    var v46 int = v1
                                    _ = v46
                                }
                            }
                            var v47 uint64 = 2
                            _ = v45
                            _ = v47
                        }
                    } else {
                        if false {
                            switch v3 {
                            case "\n\t":
                                var v48 bool = ((!true) && false)
                                _ = v48
                            }
                            v1 = v1
                            switch v1 {
                            case 1:
                                v12 = 100
                            case 492:
                                if (float32(0.0) < 1e10) {
                                    for fi49 := 0; fi49 < 6; fi49++ {
                                        var v50 uint = 2
                                        if len(v2) > 0 {
                                            v2[0] = 2147483647
                                        }
                                        _ = fi49
                                        _ = v50
                                    }
                                    var v51 uint16 = 0
                                    _ = len(v4)
                                    _ = v51
                                }
                            case 2147483647:
                                if (18446744073709551615 >= 2) {
                                    for fi52 := 0; fi52 < 11; fi52++ {
                                        if len(v4) > 0 {
                                            v4[0] = 4294967295
                                        }
                                        _ = fi52
                                    }
                                    v3 = v3
                                }
                            }
                        }
                        for fi53 := 0; fi53 < 10; fi53++ {
                            var v54 int = 42
                            v55 := GenS5{F0: 2, F1: 657}
                            v3 = (v3 + "\n\t")
                            _ = fi53
                            _ = v54
                            _ = v55
                        }
                    }
                default:
                    for fi56 := 0; fi56 < 3; fi56++ {
                        v57 := GenS7{F0: 259, F1: 2}
                        _ = fi56
                        _ = v57
                    }
                }
            }
        }
    }
    var v58 uint64 = 1
    switch "" {
    case "":
        switch v3 {
        case "":
            fmt.Println(v58, v1)
        default:
            _ = len(v4)
        }
    case "hello":
        fmt.Println(v58, v12)
    case "foo bar":
        var v59 uint8 = 255
        _ = v59
    }
    var v60 float64 = (-2.25 - -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
    _ = v60
}
