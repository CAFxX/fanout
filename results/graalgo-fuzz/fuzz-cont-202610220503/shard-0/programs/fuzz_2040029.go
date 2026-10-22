package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int64 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint16; F1 int; F2 int64 }
type GenS8 struct { F0 int }
func genFn0(p0 uint8, p1 uint8) uint16 {
    v6 := GenS1{F0: (p1 | 63), F1: (("hello" + "foo bar") + "hello"), F2: int64(p0)}
    var v7 int16 = -1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 255
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]uint32)
    var v4 uint64 = (v1 << ((887) & 63))
    var v5 uint8 = 42
    for fi8 := 0; fi8 < 18; fi8++ {
        v4 = (v1 - uint64(v5))
        _ = fi8
    }
    v9 := GenS2{F0: false, F1: 42}
    v5 = v5
    var v10 uint64 = 42
    v11 := GenS2{F0: (!(0 != 9223372036854775807)), F1: 2}
    v10 += v10
    v12 := GenS1{F0: v5, F1: ("foo bar" + "hello"), F2: -1}
    v10 -= (v1 | 1)
    for fi13 := 0; fi13 < 3; fi13++ {
        switch (("foo bar" + "a") + "a") {
        case "a":
            for fi14 := 0; fi14 < 13; fi14++ {
                v15 := v2[1]
                v16 := GenS2{F0: false, F1: 65535}
                if (32767 < 2) {
                    v10 -= v1
                }
                _ = fi14
                _ = v15
                _ = v16
            }
        default:
            switch (fi13 & fi13) {
            case 4294967295:
                for fi17 := 0; fi17 < 7; fi17++ {
                    v10 = uint64(v5)
                    _ = fi17
                }
            case 392:
                var v18 uint16 = 0
                _ = v18
            case 316:
                switch ("foo bar" + (("a" + ("hello" + "")) + "")) {
                case "":
                    for fi19 := 0; fi19 < 5; fi19++ {
                        switch "foo bar" {
                        case "a":
                            v3[2147483647] = 1
                        case "":
                            if true {
                                var v20 string = ""
                                delete(v2, 142)
                                var v21 bool = (true && false)
                                _ = v20
                                _ = v21
                            } else {
                                var v22 float32 = 0.1
                                if (619 < 4294967295) {
                                    var v23 int64 = -194
                                    v2[2147483647] = ("" + ("" + ("hello" + "foo bar")))
                                    _ = v23
                                }
                                _ = v22
                            }
                        case "\n\t":
                            var v24 int16 = 32767
                            _ = v24
                        default:
                            var v25 uint = 2
                            _ = v25
                        }
                        var v26 int16 = 42
                        if true {
                            v26 = 32767
                            _ = len(v3)
                            var v27 int16 = -1
                            _ = v27
                        } else {
                            var v28 float64 = float64((1e10 * 3.14159))
                            _ = v28
                        }
                        v29 := GenS3{F0: 2, F1: 2147483648}
                        v5 = (v5 & v5)
                        _ = fi19
                        _ = v26
                        _ = v29
                    }
                default:
                    var v30 int16 = 1
                    _ = v30
                }
            default:
                var v31 float64 = 1e10
                _ = v31
            }
        }
        v5 = (^(v5))
        _ = fi13
    }
    var v32 int64 = 699
    if (!true) {
        var v33 uint32 = 65535
        v5 = v5
        for fi34 := 0; fi34 < 15; fi34++ {
            for fi35 := 0; fi35 < 2; fi35++ {
                switch "hello" {
                case "foo bar":
                    if (true && (float64(1.5) > -2.25)) {
                        var v36 uint32 = 774
                        _ = v36
                    }
                case "\n\t":
                    v37 := []int{}
                    _ = v37
                case "":
                    v38 := GenS4{F0: (v5 * v5)}
                    _ = v38
                }
                v10 = 18446744073709551615
                _ = fi35
            }
            v33 -= (v33 / ((v33) | 1))
            _ = fi34
        }
        _ = v33
    } else {
        v39 := GenS3{F0: 153, F1: 4294967295}
        if false {
            v32 -= (v32 + 128)
            if ((0.0 / ((0.1 + -2.25) - 1e10)) <= (3.14159 - ((-2.25 * 0.0) - 1.5))) {
                var v40 int = 32767
                v32 -= v32
                _ = v40
            } else {
                v32 = v32
                v41 := GenS3{F0: 81, F1: 326}
                for fi42 := 0; fi42 < 6; fi42++ {
                    v43 := GenS1{F0: 128, F1: ("a" + "a"), F2: 9223372036854775807}
                    switch "\n\t" {
                    case "a":
                        var v44 uint16 = 2
                        _ = v44
                    case "":
                        v45 := GenS5{F0: ((("hello" + "hello") + "foo bar") + "hello")}
                        _ = v45
                    case "foo bar":
                        if true {
                            v1 -= (v4 << ((255) & 63))
                        }
                    default:
                        v46 := GenS6{F0: ("foo bar" + "foo bar")}
                        _ = v46
                    }
                    v47 := GenS7{F0: 65535, F1: (fi42 * (fi42 ^ fi42)), F2: (v32 - (^(v32)))}
                    _ = fi42
                    _ = v43
                    _ = v47
                }
                _ = v41
            }
            v48 := GenS8{F0: 32768}
            _ = v48
        } else {
            fmt.Println(v10, v32)
            v49 := GenS5{F0: "a"}
            _ = v49
        }
        _ = len(v2)
        _ = v39
    }
    var v50 uint8 = (v5 / ((v5) | 1))
    for fi51 := 0; fi51 < 16; fi51++ {
        v50 -= v50
        for fi52 := 0; fi52 < 2; fi52++ {
            if true {
                v10 = v4
                fmt.Println(fi52)
                v4 = (v4 % (((v1 & v4)) | 1))
                var v53 float32 = (1.5 - float32(1e10))
                v54 := []int{}
                _ = v53
                _ = v54
            }
            v50 += 255
            _ = fi52
        }
        switch int(v50) {
        case 128:
            var v55 float64 = -2.25
            _ = v55
        default:
            v32 = (v32 / ((v32) | 1))
        }
        _ = fi51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v50
    _ = v9
}
