package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 string }
func genFn0(p0 uint8, p1 int64, p2 int) uint64 {
    var v7 float32 = 1e10
    fmt.Println(p0, p0, p1)
    var v8 int32 = 1
    v7 += 1e10
    v9 := []int{(p2 % ((p2) | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 2
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 65535
    var v3 uint8 = 255
    v4 := make(map[string]string)
    var v5 uint32 = 903
    var v6 uint32 = (v5 | (v5 * (v5 >> ((4294967295) & 63))))
    v5 += (^(v6))
    if (!false) {
        for fi10 := 0; fi10 < 14; fi10++ {
            switch fi10 {
            case -210:
                var v11 uint8 = (v3 | 52)
                _ = v11
            case 2147483647:
                v12 := GenS1{F0: 42}
                _ = v12
            case 32768:
                var v13 uint16 = 335
                _ = v13
            }
            _ = fi10
        }
        var v14 uint = 18446744073709551615
        v6 = (v6 ^ (v6 & (v5 + v6)))
        _ = v14
    }
    fmt.Println(v2, v6)
    if true {
        v15 := GenS2{F0: 433, F1: "a"}
        switch "a" {
        case "foo bar":
            v6 -= v5
        case "":
            if (!false) {
                var v16 int64 = 9223372036854775807
                _ = v16
            }
        case "hello":
            switch 32767 {
            case 9223372036854775807:
                var v17 uint16 = 662
                _ = v17
            case 441:
                fmt.Println(v5, v2, v6)
            }
        default:
            for fi18 := 0; fi18 < 16; fi18++ {
                switch "hello" {
                case "a":
                    var v19 int64 = 42
                    _ = v19
                default:
                    var v20 int64 = 21
                    _ = v20
                }
                for fi21 := 0; fi21 < 2; fi21++ {
                    v5 += 42
                    _ = fi21
                }
                if ("foo bar" == "") {
                    _ = len(v1)
                    v2 = 0
                } else {
                    for fi22 := 0; fi22 < 18; fi22++ {
                        v2 = (^(v2))
                        v23 := GenS1{F0: v6}
                        if (fi22 < fi22) {
                            delete(v1, fi22)
                            v24 := GenS3{F0: v3}
                            _ = v24
                        }
                        _ = fi22
                        _ = v23
                    }
                    for fi25 := 0; fi25 < 15; fi25++ {
                        if true {
                            if (((("a" + "") + "\n\t") < "a") && (1 <= 79)) {
                                var v26 uint = 42
                                if false {
                                    fmt.Println(v3, fi18, v3)
                                    v27 := []int{fi25, fi25, (fi18 | fi25)}
                                    _ = v27
                                }
                                var v28 float64 = (-2.25 - (0.1 * -2.25))
                                _ = v26
                                _ = v28
                            } else {
                                v2 = 42
                                v29 := GenS4{F0: ("" + "\n\t")}
                                var v30 int16 = -192
                                _ = v29
                                _ = v30
                            }
                            v31 := []int{(^(fi25)), (fi25 + fi25), 2147483647}
                            _ = v31
                        } else {
                            if (0.1 >= (0.1 - (0.1 / 1.0))) {
                                fmt.Println(v6)
                                v32 := []int{42}
                                var v33 string = "hello"
                                var v34 float64 = (0.0 / 1e10)
                                v3 -= v3
                                v35 := GenS5{F0: -1, F1: -2147483648}
                                _ = v32
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                            if false {
                                switch -1 {
                                case 1:
                                    v36 := []int{}
                                    _ = v36
                                case 2147483647:
                                    if true {
                                        delete(v4, ("a" + ("\n\t" + "a")))
                                        var v37 uint8 = v3
                                        var v38 uint8 = 59
                                        v6 += v6
                                        _ = v37
                                        _ = v38
                                    }
                                }
                                v39 := GenS4{F0: ("foo bar" + "hello")}
                                _ = v39
                            }
                            v40 := []int{}
                            _ = v40
                        }
                        v5 += (v5 % ((v6) | 1))
                        _ = fi25
                    }
                }
                _ = fi18
            }
        }
        _ = v15
    }
    switch 127 {
    case 127:
        v2 += (^(v2))
    case 32768:
        fmt.Println(v2, v3, v6)
    case 32767:
        v5 -= 699
    default:
        var v41 float32 = 0.1
        _ = v41
    }
    var v42 string = "hello"
    switch "" {
    case "foo bar":
        switch v42 {
        case "":
            var v43 uint = 369
            _ = v43
        case "a":
            v2 += v2
        case "foo bar":
            if (!((((v42 + v42) + v42) >= "a") || ((v5 / ((v5) | 1)) > uint32(v2)))) {
                v44 := GenS6{F0: 42}
                var v45 uint8 = (v3 | (v3 % ((v3) | 1)))
                _ = v44
                _ = v45
            } else {
                v6 = (v6 % ((65535) | 1))
            }
        }
    case "\n\t":
        v3 = (v3 - 128)
    }
    if false {
        v1[128] = "foo bar"
        var v46 uint16 = v2
        var v47 bool = (0.0 != (float64(float32(1.5)) / (float64((float32(3.14159) * float32(1e10))) * (1.5 + (float64((float32((-2.25 / 1.0)) / 3.14159)) / -2.25)))))
        _ = v46
        _ = v47
    }
    for fi48 := 0; fi48 < 5; fi48++ {
        switch fi48 {
        case -1:
            fmt.Println(fi48)
        case -993:
            for fi49 := 0; fi49 < 1; fi49++ {
                fmt.Println(fi49, v5)
                _ = fi49
            }
        case -409:
            fmt.Println(v6, fi48, v5)
        default:
            var v50 int16 = 0
            _ = v50
        }
        _ = fi48
    }
    if false {
        for fi51 := 0; fi51 < 8; fi51++ {
            _ = len(v1)
            _ = fi51
        }
        v5 -= (v6 * 2147483648)
    }
    v52 := GenS7{F0: (v42 + (((v42 + "") + "foo bar") + v42))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v52
    _ = v6
}
