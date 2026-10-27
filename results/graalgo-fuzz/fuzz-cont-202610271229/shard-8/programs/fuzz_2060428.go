package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint; F1 uint32; F2 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 0.1
    v3 := []uint32{461, 531, 2}
    var v4 uint64 = 0
    v5 := make(map[uint32]string)
    switch "hello" {
    case "a":
        switch 0 {
        case -941:
            var v6 float32 = ((v2 + v2) + ((3.14159 + 3.14159) - v2))
            _ = v6
        case 17:
            v7 := GenS1{F0: 1}
            _ = v7
        case 9223372036854775807:
            switch "\n\t" {
            case "\n\t":
                v8 := GenS2{F0: -521}
                _ = v8
            case "":
                fmt.Println(v2, v4)
            }
        }
    }
    if (42 > -1) {
        v9 := GenS3{F0: 127, F1: 42, F2: 4294967295}
        if (((v2 * (v2 / v2)) + v2) == 0.0) {
            if (float64(1.5) > (0.1 + 1.5)) {
                delete(v5, 711)
            } else {
                var v10 uint8 = uint8(v4)
                if ("a" == "\n\t") {
                    for fi11 := 0; fi11 < 20; fi11++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi11
                    }
                    v4 -= (v4 ^ 422)
                } else {
                    v4 = 1
                }
                switch 32768 {
                case 2147483647:
                    v12 := GenS2{F0: -196}
                    _ = v12
                case 1:
                    v2 -= v2
                default:
                    for fi13 := 0; fi13 < 2; fi13++ {
                        v14 := GenS4{F0: 361}
                        v3 = append(v3, 42)
                        _ = fi13
                        _ = v14
                    }
                }
                _ = v10
            }
            for fi15 := 0; fi15 < 1; fi15++ {
                v2 -= v2
                v16 := GenS5{F0: false, F1: 9223372036854775807}
                _ = fi15
                _ = v16
            }
            var v17 uint = 255
            _ = v17
        }
        _ = v9
    }
    switch 4294967295 {
    case 301:
        _ = len(v3)
    case 128:
        v18 := GenS1{F0: 427}
        _ = v18
    case 32767:
        for fi19 := 0; fi19 < 16; fi19++ {
            fmt.Println(v2)
            var v20 uint16 = 65535
            _ = fi19
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        v22 := GenS6{F0: 73}
        switch fi21 {
        case 4294967295:
            if ((!true) && true) {
                if false {
                    switch "foo bar" {
                    case "foo bar":
                        v2 += v2
                    case "a":
                        v4 += 42
                    case "\n\t":
                        switch "" {
                        case "":
                            var v23 int32 = 1
                            _ = v23
                        default:
                            for fi24 := 0; fi24 < 7; fi24++ {
                                v2 += v2
                                var v25 uint8 = 227
                                _ = fi24
                                _ = v25
                            }
                        }
                    }
                    v4 -= 2
                    _ = len(v1)
                }
            } else {
                fmt.Println(fi21, v4)
                v4 -= 138
            }
        case -89:
            if (42 != 238) {
                v26 := GenS5{F0: false, F1: (fi21 + 4294967295)}
                _ = v26
            } else {
                fmt.Println(v4)
            }
        case 2147483647:
            v27 := GenS5{F0: (615 == 2), F1: fi21}
            _ = v27
        default:
            v28 := v1[4294967295]
            _ = v28
        }
        for fi29 := 0; fi29 < 10; fi29++ {
            var v30 float64 = 0.1
            var v31 int16 = -32768
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = fi21
        _ = v22
    }
    if true {
        if (926 >= 313) {
            if (!true) {
                fmt.Println(v2, v4)
            } else {
                v32 := v5[301]
                for fi33 := 0; fi33 < 12; fi33++ {
                    for fi34 := 0; fi34 < 5; fi34++ {
                        if (255 < 255) {
                            if (true || ((0 != 42) && false)) {
                                if false {
                                    v32 = ""
                                    v4 = (v4 >> ((518) & 63))
                                }
                            }
                            var v35 uint8 = 127
                            _ = v35
                        }
                        fmt.Println(fi34)
                        _ = fi34
                    }
                    _ = fi33
                }
                var v36 uint16 = uint16(v4)
                _ = v32
                _ = v36
            }
            v37 := GenS5{F0: false, F1: 32767}
            _ = v37
        }
        v38 := GenS2{F0: 42}
        var v39 int8 = 100
        _ = v38
        _ = v39
    }
    if (1 < 127) {
        if len(v3) > 0 {
            v3[0] = uint32(v4)
        }
        var v40 bool = true
        _ = v40
    } else {
        var v41 float32 = v2
        _ = v41
    }
    _ = len(v3)
    var v42 string = "\n\t"
    if (!(v42 > v42)) {
        var v43 int8 = 42
        var v44 bool = true
        v43 -= (v43 - (v43 & (v43 | v43)))
        _ = v43
        _ = v44
    } else {
        v2 = 3.14159
        v42 = "\n\t"
    }
    if false {
        for fi45 := 0; fi45 < 8; fi45++ {
            v2 += (0.0 + (float32((0.1 * (1.5 / (float64(v2) / (1.5 + -2.25))))) * v2))
            _ = fi45
        }
        fmt.Println(v2, v42, v2)
    } else {
        if ((((v42 >= ("foo bar" + v42)) || false) || ("" >= ("hello" + v42))) && (v2 == (v2 * v2))) {
            switch 32768 {
            case 0:
                v2 = v2
            case 9223372036854775807:
                var v46 uint64 = (v4 >> ((334) & 63))
                _ = v46
            case 32767:
                if false {
                    v42 = v42
                    var v47 bool = false
                    v48 := GenS6{F0: 1}
                    _ = v47
                    _ = v48
                } else {
                    v49 := GenS7{F0: 271, F1: 2147483647, F2: 2}
                    v2 = (v2 * 1e10)
                    _ = v49
                }
            default:
                var v50 int64 = 42
                _ = v50
            }
            switch (("a" + (v42 + v42)) + "") {
            case "hello":
                for fi51 := 0; fi51 < 2; fi51++ {
                    v42 = "hello"
                    var v52 uint8 = 105
                    _ = fi51
                    _ = v52
                }
            case "\n\t":
                v4 = v4
            case "foo bar":
                var v53 uint64 = v4
                _ = v53
            default:
                v42 = (v42 + v42)
            }
        } else {
            for fi54 := 0; fi54 < 19; fi54++ {
                var v55 float32 = ((3.14159 - float32(((float64(float32(0.0)) * -2.25) - 0.0))) + 1e10)
                _ = fi54
                _ = v55
            }
            fmt.Println(v4, v42, v42)
        }
        var v56 float32 = (v2 * float32(float64(((v2 / (1.5 + float32(float64(v2)))) + v2))))
        _ = v56
    }
    var v57 uint64 = 4294967295
    v57 += v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v57
}
