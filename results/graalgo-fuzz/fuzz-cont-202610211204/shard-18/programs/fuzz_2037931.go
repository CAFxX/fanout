package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 int; F2 uint8 }
type GenS4 struct { F0 int32; F1 uint32 }

func main() {
    var v1 uint64 = 178
    v2 := []string{"\n\t", "a"}
    v3 := make(map[uint32]int)
    var v4 string = ("a" + ("hello" + ("hello" + "a")))
    delete(v3, 65535)
    v5 := GenS1{F0: -9223372036854775808, F1: 42}
    v6 := GenS1{F0: -732, F1: 225}
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 uint32 = 0
            if false {
                if (852 != 2) {
                    _ = len(v3)
                    var v10 uint8 = uint8(v1)
                    _ = v10
                } else {
                    if (2147483647 == -276) {
                        var v11 uint8 = 0
                        _ = v11
                    } else {
                        switch v4 {
                        case "a":
                            var v12 uint32 = v9
                            _ = v12
                        case "foo bar":
                            fmt.Println(fi7, fi8, fi7)
                        default:
                            if (1 >= 0) {
                                v13 := GenS2{F0: -457}
                                var v14 string = (v4 + (v4 + v4))
                                var v15 int8 = 42
                                _ = v13
                                _ = v14
                                _ = v15
                            } else {
                                v1 = 1
                                fmt.Println(v9, v1, v1)
                                var v16 uint16 = 973
                                if (false || (42 > 2147483647)) {
                                    var v17 float32 = 0.1
                                    _ = v17
                                }
                                _ = v16
                            }
                        }
                    }
                }
                var v18 int32 = 1
                _ = v18
            }
            fmt.Println(v1)
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v19 := GenS3{F0: -1, F1: 176, F2: 0}
    switch 42 {
    case 369:
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
    }
    v1 += 1
    for fi20 := 0; fi20 < 3; fi20++ {
        v4 = ""
        var v21 int16 = 32767
        _ = fi20
        _ = v21
    }
    switch 42 {
    case 32767:
        var v22 int32 = 1
        _ = v22
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        v4 = (v4 + v4)
        _ = fi23
    }
    if false {
        v4 = v4
        if (!(1.5 <= 0.1)) {
            var v24 float64 = 0.1
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v24
        }
        for fi25 := 0; fi25 < 10; fi25++ {
            switch ("foo bar" + v4) {
            case "a":
                if true {
                    var v26 uint64 = 325
                    for fi27 := 0; fi27 < 14; fi27++ {
                        v28 := GenS2{F0: 1000}
                        fmt.Println(fi27, fi27)
                        _ = fi27
                        _ = v28
                    }
                    var v29 float64 = (0.0 / (((float64((3.14159 / 3.14159)) * float64(1.5)) + (3.14159 * -2.25)) * 3.14159))
                    _ = v26
                    _ = v29
                } else {
                    v4 = v4
                    if (0.1 <= 1.5) {
                        var v30 int8 = 42
                        _ = v30
                    } else {
                        var v31 int64 = 9223372036854775807
                        _ = v31
                    }
                }
            case "hello":
                var v32 float64 = float64(float32(0.0))
                _ = v32
            default:
                for fi33 := 0; fi33 < 7; fi33++ {
                    if true {
                        switch ((("a" + ("" + v4)) + "foo bar") + v4) {
                        case "foo bar":
                            var v34 uint64 = (v1 + v1)
                            _ = v34
                        case "":
                            var v35 uint16 = 42
                            _ = v35
                        }
                    }
                    _ = fi33
                }
            }
            _ = fi25
        }
    }
    v36 := GenS3{F0: -1, F1: 32768, F2: 80}
    v37 := GenS3{F0: 0, F1: -540, F2: 128}
    if true {
        _ = len(v2)
        for fi38 := 0; fi38 < 13; fi38++ {
            v4 = ((v4 + ((v4 + v4) + (v4 + (("\n\t" + "hello") + v4)))) + "a")
            _ = fi38
        }
        var v39 int8 = 127
        _ = v39
    } else {
        var v40 uint8 = 255
        switch 42 {
        case 32767:
            switch 1 {
            case 127:
                if (false || (1 < 0)) {
                    var v41 uint32 = 2147483648
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                    _ = v41
                } else {
                    v42 := GenS2{F0: 366}
                    var v43 uint = uint(v1)
                    if ((v43 - (^(v43))) <= (^(v43))) {
                        switch 32768 {
                        case 492:
                            var v44 float64 = 1e10
                            _ = v44
                        default:
                            v43 += (v43 % (((v43 ^ (v43 >> (((v43 - v43)) & 63)))) | 1))
                        }
                        var v45 int64 = -1
                        _ = v45
                    } else {
                        var v46 int8 = int8(v1)
                        _ = v46
                    }
                    _ = v42
                    _ = v43
                }
            case 0:
                v47 := GenS4{F0: 1, F1: uint32(v1)}
                _ = v47
            case 185:
                v40 += (v40 + v40)
            default:
                var v48 uint64 = v1
                _ = v48
            }
        default:
            if len(v2) > 0 {
                v2[0] = ((v4 + (v4 + "")) + "")
            }
        }
        v49 := GenS1{F0: 9223372036854775807, F1: -325}
        _ = v40
        _ = v49
    }
    v1 += v1
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
