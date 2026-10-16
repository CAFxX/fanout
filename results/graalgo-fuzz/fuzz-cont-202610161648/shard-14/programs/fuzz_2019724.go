package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int64; F1 uint; F2 string }
type GenS3 struct { F0 uint; F1 uint8; F2 bool }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 string; F1 int8; F2 uint8 }
type GenS6 struct { F0 int64; F1 int32; F2 string }
func genFn0(p0 int32, p1 int64, p2 uint16) bool {
    fmt.Println(p0, p1)
    v6 := GenS1{F0: int64(p0), F1: 1, F2: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (((p0 | (p0 | (p0 % ((p0) | 1)))) >= -1) && false)
}

func main() {
    v1 := []int{190, 859, -335, 9223372036854775807}
    var v2 bool = true
    var v3 uint8 = 33
    v4 := make(map[int]int)
    var v5 float64 = (-2.25 * ((1.5 + 0.1) + 1.5))
    _ = len(v1)
    switch 9223372036854775807 {
    case 9223372036854775807:
        v7 := GenS2{F0: 1, F1: 60, F2: ("a" + ("hello" + ""))}
        _ = v7
    case 42:
        v5 = v5
    }
    _ = len(v4)
    for fi8 := 0; fi8 < 3; fi8++ {
        v4[(fi8 & 32767)] = 347
        var v9 uint64 = 334
        switch "hello" {
        case "\n\t":
            v10 := GenS2{F0: 0, F1: 4294967295, F2: ("" + (("foo bar" + "hello") + "a"))}
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    v11 := GenS1{F0: 197, F1: 42, F2: 147}
    for fi12 := 0; fi12 < 9; fi12++ {
        v5 -= 0.0
        v4[fi12] = (fi12 - fi12)
        _ = fi12
    }
    v13 := GenS1{F0: 42, F1: 0, F2: 646}
    if v2 {
        _ = len(v1)
        v2 = v2
        fmt.Println(v2, v2)
    }
    var v14 bool = (v5 < ((v5 / (v5 / 1.5)) - v5))
    if v14 {
        switch (("\n\t" + "foo bar") + "hello") {
        case "":
            v5 = 3.14159
        case "foo bar":
            if len(v1) > 0 {
                v1[0] = 32768
            }
        }
        v2 = v2
    } else {
        if false {
            switch 32768 {
            case -9223372036854775808:
                var v15 uint64 = 42
                _ = v15
            case 42:
                switch ("\n\t" + ("foo bar" + ("a" + "a"))) {
                case "foo bar":
                    var v16 bool = (340 == 65535)
                    _ = v16
                case "":
                    fmt.Println(v3, v3)
                case "hello":
                    fmt.Println(v3, v3, v3)
                }
            default:
                if true {
                    for fi17 := 0; fi17 < 1; fi17++ {
                        v2 = (!v14)
                        v5 -= ((float64(((3.14159 * 1e10) - 1e10)) / 0.1) + ((1e10 + v5) + v5))
                        fmt.Println(fi17, v3, v2)
                        _ = fi17
                    }
                    v3 = 77
                }
            }
            switch 128 {
            case 873:
                var v18 uint8 = 64
                _ = v18
            case 1:
                if v14 {
                    switch 42 {
                    case 2147483647:
                        switch "a" {
                        case "":
                            v19 := GenS2{F0: 0, F1: 223, F2: ("foo bar" + "a")}
                            _ = v19
                        default:
                            switch -9223372036854775808 {
                            case -9223372036854775808:
                                v20 := GenS1{F0: -9223372036854775808, F1: 2, F2: 255}
                                _ = v20
                            case -1:
                                var v21 float32 = 1e10
                                _ = v21
                            case 127:
                                for fi22 := 0; fi22 < 14; fi22++ {
                                    var v23 uint32 = 65535
                                    var v24 bool = v2
                                    _ = fi22
                                    _ = v23
                                    _ = v24
                                }
                            }
                        }
                    case -323:
                        v25 := GenS2{F0: 1, F1: 1, F2: "\n\t"}
                        _ = v25
                    case 1:
                        v14 = true
                    }
                    var v26 uint64 = 917
                    var v27 float64 = v5
                    _ = v26
                    _ = v27
                }
            default:
                v28 := GenS2{F0: 185, F1: 551, F2: ""}
                _ = v28
            }
        } else {
            v14 = (1e10 < 1.5)
        }
        if v2 {
            fmt.Println(v3, v3)
            if len(v1) > 0 {
                v1[0] = 32767
            }
            v29 := GenS3{F0: 4294967295, F1: v3, F2: (490 < -567)}
            _ = v29
        }
        v30 := GenS2{F0: int64(v3), F1: 65535, F2: "hello"}
        _ = v30
    }
    v14 = (v2 && false)
    for fi31 := 0; fi31 < 16; fi31++ {
        _ = len(v1)
        _ = fi31
    }
    var v32 string = "hello"
    v33 := GenS4{F0: "a", F1: 97}
    v34 := GenS4{F0: v32, F1: v3}
    if ((926 >= 104) || (255 == v3)) {
        if v14 {
            v14 = (!v14)
            if v14 {
                v14 = (int16(v3) <= 162)
                if v2 {
                    var v35 float64 = float64(-2.25)
                    _ = v35
                }
                v36 := GenS5{F0: (v32 + v32), F1: -1, F2: v3}
                _ = v36
            }
        }
        if false {
            v37 := v4[1]
            _ = v37
        } else {
            fmt.Println(v32, v3)
            v3 = (v3 | v3)
            if true {
                var v38 uint64 = 223
                var v39 int64 = 9223372036854775807
                var v40 float64 = (float64((0.0 + 1.5)) * float64(1.5))
                _ = v38
                _ = v39
                _ = v40
            } else {
                v41 := GenS6{F0: -236, F1: 693, F2: "a"}
                _ = v41
            }
        }
    }
    v42 := GenS5{F0: (v32 + v32), F1: 0, F2: (v3 << ((18446744073709551615) & 63))}
    fmt.Println(v3, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v42
    _ = v5
}
