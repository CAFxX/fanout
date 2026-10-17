package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64; F1 int8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int }
func genFn0(p0 uint16, p1 int64) int64 {
    var v6 uint64 = 42
    var v7 uint16 = p0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (p1 + (p1 >> ((738) & 63)))
}
func genFn1(p0 uint8, p1 uint8) int64 {
    var v8 uint = 18446744073709551615
    v8 = (v8 % (((v8 - (v8 & v8))) | 1))
    _ = p0
    _ = p1
    _ = v8
    return genFn0(20, 9223372036854775807)
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 58
    v3 := []string{"a", ((("foo bar" + ((("a" + "hello") + ("" + "foo bar")) + ("foo bar" + ""))) + "\n\t") + "a"), ("" + "a"), "a"}
    var v4 uint32 = v2
    var v5 int32 = -199
    v9 := GenS1{F0: 128, F1: 777}
    v10 := GenS2{F0: v5, F1: ""}
    fmt.Println(v4)
    v11 := GenS3{F0: 980, F1: 0, F2: 248}
    if true {
        switch 162 {
        case -970:
            v12 := GenS4{F0: 2}
            _ = v12
        case 127:
            for fi13 := 0; fi13 < 14; fi13++ {
                for fi14 := 0; fi14 < 10; fi14++ {
                    if (!(true || false)) {
                        v2 += (v4 >> ((42) & 63))
                        v15 := GenS5{F0: fi13}
                        _ = v15
                    } else {
                        if (3.14159 < (1.5 + 1.5)) {
                            var v16 int32 = v5
                            v17 := GenS4{F0: 505}
                            if (fi14 >= fi13) {
                                switch fi13 {
                                case 32767:
                                    var v18 int16 = 32767
                                    _ = v18
                                }
                                v16 = int32(fi13)
                            } else {
                                if (32767 != -32768) {
                                    fmt.Println(v5, v4, v16)
                                }
                                v4 -= v2
                                var v19 uint32 = (v2 ^ v2)
                                _ = v19
                            }
                            var v20 uint8 = 0
                            _ = v16
                            _ = v17
                            _ = v20
                        }
                        switch "foo bar" {
                        case "":
                            fmt.Println(fi14, fi14, fi13)
                        case "foo bar":
                            v4 = v2
                        default:
                            switch 42 {
                            case 965:
                                v21 := GenS5{F0: -259}
                                _ = v21
                            case 692:
                                v2 += (v2 << ((171) & 63))
                            default:
                                if true {
                                    switch "a" {
                                    case "":
                                        var v22 uint32 = (v4 ^ 65535)
                                        _ = v22
                                    case "a":
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    default:
                                        fmt.Println(v5)
                                    }
                                }
                            }
                        }
                    }
                    v23 := GenS6{F0: -75, F1: 1}
                    v5 += (v5 >> ((65535) & 63))
                    _ = fi14
                    _ = v23
                }
                var v24 int = (-(fi13))
                _ = fi13
                _ = v24
            }
        case -901:
            var v25 int64 = -927
            _ = v25
        default:
            v5 -= (^(v5))
        }
        for fi26 := 0; fi26 < 14; fi26++ {
            var v27 int32 = v5
            _ = fi26
            _ = v27
        }
    }
    var v28 int8 = -53
    var v29 float64 = (3.14159 / -2.25)
    if true {
        var v30 int64 = -1
        var v31 string = "foo bar"
        switch -731 {
        case 715:
            var v32 uint64 = 258
            _ = v32
        case 32768:
            v33 := GenS5{F0: -331}
            _ = v33
        case 4294967295:
            if true {
                for fi34 := 0; fi34 < 2; fi34++ {
                    v35 := GenS1{F0: 67, F1: -1}
                    _ = fi34
                    _ = v35
                }
                if ((false || (((-2.25 - (0.0 - 1.5)) - 1e10) <= -2.25)) || true) {
                    v36 := v1[("\n\t" + (v31 + ("foo bar" + v31)))]
                    if len(v3) > 0 {
                        v3[0] = "a"
                    }
                    if true {
                        v2 -= 2147483648
                        v37 := GenS7{F0: (v28 / ((19) | 1))}
                        _ = v37
                    }
                    _ = v36
                } else {
                    for fi38 := 0; fi38 < 14; fi38++ {
                        v39 := GenS1{F0: 1, F1: -352}
                        _ = fi38
                        _ = v39
                    }
                    fmt.Println(v4, v2, v31)
                    var v40 int16 = 908
                    _ = v40
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                var v41 int32 = 0
                _ = v41
            }
        default:
            v42 := GenS8{F0: 42}
            _ = v42
        }
        _ = v30
        _ = v31
    }
    v43 := GenS6{F0: 42, F1: (v28 >> ((4294967295) & 63))}
    v29 = v29
    fmt.Println(v29)
    switch 1 {
    case 127:
        var v44 uint16 = 290
        _ = v44
    case 32768:
        switch 2147483647 {
        case 101:
            var v45 float32 = float32(1e10)
            _ = v45
        case -2147483648:
            v28 = 100
        case 42:
            fmt.Println(v2)
        default:
            fmt.Println(v28)
        }
    case 0:
        var v46 float64 = 1e10
        _ = v46
    default:
        v2 -= v4
    }
    for fi47 := 0; fi47 < 14; fi47++ {
        switch "foo bar" {
        case "":
            switch ("\n\t" + "\n\t") {
            case "hello":
                fmt.Println(v29, v5, v4)
            case "":
                v2 = v2
            default:
                var v48 float32 = 3.14159
                _ = v48
            }
        }
        var v49 bool = (!(!true))
        if (!((1 < v28) && true)) {
            fmt.Println(v29, v49, v4)
            for fi50 := 0; fi50 < 5; fi50++ {
                v4 += (v4 | (v4 | 72))
                _ = fi50
            }
        }
        _ = fi47
        _ = v49
    }
    var v51 bool = false
    v5 -= (v5 << ((0) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v51
    _ = v9
}
