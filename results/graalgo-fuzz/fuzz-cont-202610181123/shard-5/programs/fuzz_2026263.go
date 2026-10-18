package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 int }

func main() {
    var v1 int64 = 298
    v2 := make(map[int]bool)
    var v3 uint8 = 128
    var v4 float32 = (0.1 - ((1e10 / 1.0) * 3.14159))
    v5 := make(map[int]uint32)
    var v6 int8 = 120
    v7 := []int{452, 679, -393, 1}
    switch "\n\t" {
    case "a":
        if false {
            v8 := GenS1{F0: (v1 <= 936), F1: uint16(v3), F2: 660}
            _ = v8
        } else {
            fmt.Println(v6, v1, v3)
        }
    default:
        if false {
            for fi9 := 0; fi9 < 17; fi9++ {
                v1 = v1
                var v10 string = "a"
                var v11 uint8 = (v3 | 2)
                _ = fi9
                _ = v10
                _ = v11
            }
            for fi12 := 0; fi12 < 1; fi12++ {
                v1 = v1
                for fi13 := 0; fi13 < 12; fi13++ {
                    if (1.5 != v4) {
                        fmt.Println(fi13, v3, v3)
                        fmt.Println(fi13, fi13)
                    }
                    _ = fi13
                }
                _ = fi12
            }
            v14 := GenS2{F0: false}
            _ = v14
        } else {
            var v15 string = "foo bar"
            for fi16 := 0; fi16 < 15; fi16++ {
                var v17 int8 = v6
                switch ("\n\t" + "a") {
                case "":
                    for fi18 := 0; fi18 < 4; fi18++ {
                        if (787 <= 601) {
                            var v19 int16 = 42
                            v20 := GenS1{F0: true, F1: 65535, F2: 495}
                            _ = v19
                            _ = v20
                        }
                        _ = fi18
                    }
                case "a":
                    switch 1 {
                    case 1:
                        for fi21 := 0; fi21 < 15; fi21++ {
                            fmt.Println(fi21)
                            v15 = (v15 + (v15 + v15))
                            _ = fi21
                        }
                    case 46:
                        if (0 >= -128) {
                            v22 := GenS2{F0: true}
                            switch (^(fi16)) {
                            case 904:
                                v17 = 100
                            case 2147483647:
                                var v23 int64 = 1
                                _ = v23
                            default:
                                var v24 int64 = -778
                                _ = v24
                            }
                            _ = v22
                        }
                    default:
                        v3 += (v3 - 154)
                    }
                default:
                    v25 := GenS2{F0: true}
                    _ = v25
                }
                for fi26 := 0; fi26 < 6; fi26++ {
                    _ = len(v2)
                    v4 = -2.25
                    _ = fi26
                }
                _ = fi16
                _ = v17
            }
            _ = v15
        }
    }
    switch ("a" + "") {
    case "foo bar":
        var v27 int8 = v6
        _ = v27
    }
    v6 = (v6 % ((-128) | 1))
    var v28 string = (("a" + "foo bar") + ("" + "hello"))
    if ((!(!true)) || false) {
        if true {
            switch "hello" {
            case "a":
                var v29 uint32 = 0
                _ = v29
            case "foo bar":
                switch 284 {
                case 128:
                    if (2147483647 == 42) {
                        v30 := GenS1{F0: (42 == 100), F1: 953, F2: 0}
                        if true {
                            v6 -= (v6 / (((-(v6))) | 1))
                            v4 = v4
                        }
                        v31 := GenS2{F0: true}
                        _ = v30
                        _ = v31
                    }
                case 2147483647:
                    var v32 int8 = (v6 & v6)
                    _ = v32
                default:
                    delete(v5, 2)
                }
            default:
                if false {
                    switch 2147483647 {
                    case 758:
                        v6 -= 42
                    default:
                        v33 := GenS2{F0: (v1 < (v1 << ((0) & 63)))}
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 18; fi34++ {
                        v3 = (v3 % ((v3) | 1))
                        v4 += (v4 + (0.1 - v4))
                        v35 := GenS3{F0: (fi34 | fi34)}
                        _ = fi34
                        _ = v35
                    }
                    v36 := GenS4{F0: 2147483648, F1: 128}
                    _ = v36
                } else {
                    _ = len(v7)
                    var v37 string = (v28 + ((v28 + (v28 + v28)) + v28))
                    _ = v37
                }
            }
        }
        v3 += v3
    }
    for fi38 := 0; fi38 < 20; fi38++ {
        var v39 float64 = 3.14159
        v3 = 127
        _ = fi38
        _ = v39
    }
    v1 -= -1
    if true {
        var v40 bool = ((!(((v6 << ((65535) & 63)) <= (v6 - v6)) && false)) && false)
        v41 := GenS3{F0: -9223372036854775808}
        if ((-(v6)) > (v6 - (v6 % ((int8(v3)) | 1)))) {
            if v40 {
                for fi42 := 0; fi42 < 8; fi42++ {
                    var v43 uint64 = 18446744073709551615
                    var v44 int = (fi42 >> ((2) & 63))
                    _ = fi42
                    _ = v43
                    _ = v44
                }
                v6 += 127
                var v45 int32 = 42
                _ = v45
            }
            v46 := v2[-2147483648]
            if (((!v46) || (42 <= 2147483647)) || (!(v28 != ("a" + v28)))) {
                fmt.Println(v40, v6)
            }
            _ = v46
        }
        _ = v40
        _ = v41
    }
    var v47 uint16 = 32767
    v48 := GenS3{F0: 42}
    if true {
        switch 128 {
        case 1:
            v6 = v6
        case 77:
            fmt.Println(v47, v1, v4)
        }
    }
    var v49 uint16 = v47
    switch 32767 {
    case -315:
        var v50 float64 = 3.14159
        _ = v50
    case -965:
        v49 -= 2
    }
    var v51 uint64 = 1
    if false {
        var v52 float64 = ((1e10 - float64(float32(((((((3.14159 / 1.5) - float64(((0.1 / ((0.1 + v4) - 1e10)) + v4))) + -2.25) / ((3.14159 + -2.25) + float64(1e10))) / 1.0) + -2.25)))) / 1.0)
        _ = v52
    }
    if (2147483648 > 0) {
        v53 := GenS4{F0: 777, F1: 689}
        _ = v53
    } else {
        v54 := GenS3{F0: 459}
        v49 = (^(v47))
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v51
    _ = v6
    _ = v7
}
