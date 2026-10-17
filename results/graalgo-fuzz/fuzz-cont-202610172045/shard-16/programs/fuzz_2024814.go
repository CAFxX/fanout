package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 bool; F1 string; F2 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 string }
func genFn0(p0 int, p1 uint32, p2 uint16) uint8 {
    v6 := []int{32768}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = len(v6)
    v7 := GenS1{F0: 42, F1: ""}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 255
}
func genFn1(p0 uint64, p1 int8) int8 {
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 int32 = 506
        for fi10 := 0; fi10 < 14; fi10++ {
            v9 -= v9
            if (42 <= 1) {
                var v11 uint64 = (p0 >> ((1) & 63))
                _ = v11
            } else {
                v12 := []int{fi10, (-(fi10))}
                p1 += (p1 >> ((59) & 63))
                v9 += v9
                _ = v12
            }
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    p1 += -102
    p1 -= p1
    if (447 <= 63) {
        var v13 int16 = 1
        _ = v13
    }
    if false {
        if ((0 <= 2) || false) {
            switch 571 {
            case 2147483647:
                var v14 bool = true
                _ = v14
            }
            p0 += p0
        } else {
            var v15 float32 = float32((-2.25 + (-2.25 / 3.14159)))
            var v16 int16 = 337
            var v17 int8 = (p1 << ((18446744073709551615) & 63))
            _ = v15
            _ = v16
            _ = v17
        }
        v18 := []int{4294967295}
        _ = v18
    }
    _ = p0
    _ = p1
    return (-(p1))
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int16 = -1
    var v3 int8 = 100
    v4 := make(map[string]int)
    v5 := []int{9223372036854775807, 0, 32768}
    v5 = append(v5, 296)
    v19 := GenS2{F0: false, F1: "", F2: ("\n\t" + ("foo bar" + ("a" + "foo bar")))}
    fmt.Println(v1)
    fmt.Println(v1, v1)
    switch -2147483648 {
    case 42:
        var v20 string = "\n\t"
        _ = v20
    default:
        if true {
            if false {
                switch 127 {
                case -298:
                    v5 = append(v5, 895)
                case 0:
                    if (18446744073709551615 > 453) {
                        switch "\n\t" {
                        case "hello":
                            var v21 uint64 = 4294967295
                            _ = v21
                        default:
                            if len(v5) > 0 {
                                v5[0] = 4294967295
                            }
                        }
                        if ((v2 * -960) >= v2) {
                            v5 = append(v5, 42)
                            switch 1 {
                            case 0:
                                v1 += v1
                            case -2147483648:
                                v2 += 32767
                            }
                            v2 += (v2 | -32768)
                        }
                        v1 += (v1 ^ (v1 ^ 31))
                    } else {
                        if true {
                            fmt.Println(v1)
                            switch "foo bar" {
                            case "a":
                                if (true || (v3 > v3)) {
                                    for fi22 := 0; fi22 < 10; fi22++ {
                                        var v23 bool = false
                                        fmt.Println(v2, v3, v2)
                                        v23 = v23
                                        _ = fi22
                                        _ = v23
                                    }
                                    v2 = 258
                                } else {
                                    v2 -= (v2 ^ v2)
                                    v24 := GenS1{F0: 481, F1: ""}
                                    _ = v24
                                }
                            case "foo bar":
                                fmt.Println(v1, v3)
                            }
                            v4[((((((("\n\t" + "\n\t") + "a") + "hello") + ("a" + ("a" + ("a" + (("foo bar" + "hello") + ""))))) + "") + ("foo bar" + "hello")) + ("foo bar" + "foo bar"))] = 0
                        }
                    }
                }
                v25 := GenS1{F0: 2, F1: "\n\t"}
                fmt.Println(v2, v3)
                _ = v25
            }
            if (!true) {
                var v26 string = "\n\t"
                v27 := GenS2{F0: ((!true) || false), F1: (v26 + "a"), F2: ""}
                _ = v26
                _ = v27
            }
        } else {
            switch 892 {
            case 195:
                _ = len(v5)
            case 42:
                var v28 string = ("hello" + "")
                _ = v28
            case 9223372036854775807:
                switch 9223372036854775807 {
                case -71:
                    var v29 uint8 = 42
                    _ = v29
                case 127:
                    v1 += (v1 % ((208) | 1))
                case -9223372036854775808:
                    var v30 int16 = 42
                    _ = v30
                }
            default:
                v31 := GenS3{F0: -2147483648}
                _ = v31
            }
        }
    }
    for fi32 := 0; fi32 < 8; fi32++ {
        v3 -= v3
        v33 := GenS4{F0: 1, F1: 218}
        _ = fi32
        _ = v33
    }
    v2 -= 537
    var v34 int64 = (v1 / ((v1) | 1))
    switch -1 {
    case 42:
        fmt.Println(v3)
    case 529:
        v35 := GenS5{F0: "hello"}
        _ = v35
    case -2147483648:
        var v36 string = (("\n\t" + "foo bar") + "")
        _ = v36
    }
    if true {
        for fi37 := 0; fi37 < 8; fi37++ {
            delete(v4, "a")
            _ = fi37
        }
    }
    for fi38 := 0; fi38 < 16; fi38++ {
        fmt.Println(v3)
        _ = fi38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
