package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 int8 }
func genFn0(p0 int16, p1 uint64) int {
    var v6 bool = true
    switch ("" + ((("foo bar" + ("hello" + (("a" + ("\n\t" + (("" + ("hello" + "")) + "foo bar"))) + ("hello" + ("a" + "\n\t"))))) + "") + "")) {
    case "hello":
        p1 = 98
    case "\n\t":
        v6 = (174 > 127)
    default:
        switch 42 {
        case 32767:
            switch 32768 {
            case 9223372036854775807:
                v7 := []int{813}
                _ = v7
            case 4294967295:
                for fi8 := 0; fi8 < 19; fi8++ {
                    for fi9 := 0; fi9 < 13; fi9++ {
                        switch "hello" {
                        case "hello":
                            p0 = (p0 ^ (p0 * p0))
                        }
                        v6 = (((p1 / ((18446744073709551615) | 1)) <= 2) || v6)
                        _ = fi9
                    }
                    v6 = (!(!true))
                    _ = fi8
                }
            case 1:
                for fi10 := 0; fi10 < 2; fi10++ {
                    fmt.Println(p0, fi10)
                    var v11 float32 = -2.25
                    v11 += 1.5
                    _ = fi10
                    _ = v11
                }
            default:
                p1 -= 18446744073709551615
            }
        }
    }
    v12 := []int{127, 127, 1}
    for fi13 := 0; fi13 < 7; fi13++ {
        for fi14 := 0; fi14 < 14; fi14++ {
            fmt.Println(fi14)
            p0 = 1
            if ((!(!false)) || false) {
                if false {
                    switch fi13 {
                    case -9223372036854775808:
                        var v15 uint64 = (p1 & 0)
                        _ = v15
                    case 32767:
                        _ = len(v12)
                        fmt.Println(v6, p1, p1)
                        var v16 uint64 = 405
                        switch 666 {
                        case 32768:
                            var v17 float64 = (0.0 - (((1e10 / 1.0) / (((0.1 + 0.1) - ((float64((0.1 - 1e10)) - (3.14159 - 0.1)) - 0.0)) + (3.14159 - float64(3.14159)))) / float64(float32((1e10 + 0.1)))))
                            _ = v17
                        case 127:
                            if true {
                                var v18 uint = 65535
                                v12 = append(v12, fi14)
                                v16 = v16
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                                v19 := GenS1{F0: 2, F1: -96, F2: v18}
                                _ = v18
                                _ = v19
                            } else {
                                var v20 int64 = 9223372036854775807
                                _ = v20
                            }
                        default:
                            v16 -= (p1 >> ((1) & 63))
                        }
                        var v21 uint8 = 255
                        _ = v16
                        _ = v21
                    case -92:
                        var v22 bool = (1.5 > 1.5)
                        _ = v22
                    default:
                        fmt.Println(fi13, fi13, p1)
                    }
                    var v23 float32 = ((0.0 + ((1e10 - 0.1) / 1e10)) * 1.5)
                    switch fi13 {
                    case 367:
                        switch fi13 {
                        case 9223372036854775807:
                            p1 = (p1 / (((p1 - p1)) | 1))
                        case 0:
                            fmt.Println(v6)
                        }
                    case 2147483647:
                        v24 := GenS1{F0: 8, F1: 127, F2: 255}
                        _ = v24
                    case 127:
                        v6 = v6
                    default:
                        if (!v6) {
                            var v25 int8 = -1
                            fmt.Println(v23, v23)
                            v26 := GenS2{F0: 255, F1: (p1 % (((p1 * (^(p1)))) | 1)), F2: 2147483648}
                            _ = v25
                            _ = v26
                        } else {
                            if len(v12) > 0 {
                                v12[0] = (fi14 & fi14)
                            }
                            v23 = float32(0.1)
                        }
                    }
                    _ = v23
                } else {
                    if v6 {
                        p0 = (p0 >> ((65535) & 63))
                        switch (fi14 / ((fi13) | 1)) {
                        case 205:
                            fmt.Println(fi13)
                        case -1:
                            fmt.Println(v6, p0, fi13)
                        case 1:
                            if (true && ("\n\t" < (("foo bar" + "foo bar") + "\n\t"))) {
                                var v27 uint8 = 159
                                _ = v27
                            } else {
                                p0 += -32768
                            }
                        default:
                            p1 += p1
                        }
                    }
                    p0 = p0
                }
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                var v28 int16 = (^(p0))
                _ = v28
            } else {
                switch 128 {
                case -9223372036854775808:
                    v6 = v6
                case 1:
                    var v29 int32 = -1
                    _ = v29
                }
                v30 := GenS1{F0: 255, F1: -53, F2: 309}
                v31 := GenS1{F0: 792, F1: -128, F2: 65535}
                _ = v30
                _ = v31
            }
            _ = fi14
        }
        _ = fi13
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v6
    return -188
}

func main() {
    var v1 float64 = 1.5
    var v2 uint8 = 42
    var v3 bool = false
    var v4 uint = 613
    v5 := make(map[uint32]string)
    var v32 uint8 = v2
    for fi33 := 0; fi33 < 9; fi33++ {
        for fi34 := 0; fi34 < 6; fi34++ {
            v35 := GenS2{F0: (^(v32)), F1: 402, F2: 2147483647}
            v36 := GenS3{F0: true}
            v2 = 152
            v37 := []int{(fi33 << ((v4) & 63))}
            v38 := GenS4{F0: 947, F1: 1}
            v1 = 3.14159
            _ = fi34
            _ = v35
            _ = v36
            _ = v37
            _ = v38
        }
        _ = fi33
    }
    if (397 <= 1) {
        switch genFn0(-32768, 128) {
        case 706:
            fmt.Println(v1, v32, v2)
        case 502:
            delete(v5, 1)
        case -1:
            var v39 int32 = 1
            _ = v39
        }
        for fi40 := 0; fi40 < 3; fi40++ {
            var v41 float32 = 1.5
            _ = fi40
            _ = v41
        }
        v42 := GenS2{F0: (v32 + (v32 + v2)), F1: 0, F2: 97}
        _ = v42
    } else {
        fmt.Println(v4, v4)
    }
    var v43 uint8 = (v32 & 42)
    if true {
        v44 := GenS5{F0: -128}
        v45 := GenS5{F0: 42}
        _ = len(v5)
        _ = v44
        _ = v45
    }
    fmt.Println(v4, v1, v4)
    v3 = v3
    v4 += (^(v4))
    fmt.Println(v1, v1)
    fmt.Println(v4)
    v4 = 0
    if v3 {
        switch "hello" {
        case "\n\t":
            v46 := GenS3{F0: v3}
            _ = v46
        case "foo bar":
            var v47 uint16 = 1
            _ = v47
        case "a":
            fmt.Println(v4, v4, v2)
        default:
            v48 := GenS4{F0: 90, F1: 42}
            _ = v48
        }
        for fi49 := 0; fi49 < 14; fi49++ {
            var v50 uint64 = 42
            var v51 uint8 = (^(v2))
            _ = fi49
            _ = v50
            _ = v51
        }
    }
    v52 := GenS5{F0: -68}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v43
    _ = v5
    _ = v52
}
