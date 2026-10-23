package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string; F1 string }
func genFn0(p0 int16, p1 int64, p2 int32) bool {
    v7 := []int{578, -9223372036854775808, -9223372036854775808}
    var v8 uint = 18446744073709551615
    if (!(100 <= -37)) {
        if ((p1 ^ -641) != p1) {
            switch "foo bar" {
            case "hello":
                var v9 int16 = (p0 ^ (-(p0)))
                _ = v9
            case "foo bar":
                var v10 int64 = (^(p1))
                _ = v10
            case "":
                if ((p1 >> ((v8) & 63)) < (p1 - (p1 >> (((v8 % ((v8) | 1))) & 63)))) {
                    var v11 float32 = 0.0
                    var v12 int = -2147483648
                    v12 = (v12 / (((^(v12))) | 1))
                    _ = v11
                    _ = v12
                } else {
                    var v13 int = 388
                    v8 = v8
                    _ = v13
                }
            default:
                p2 = p2
            }
            p0 -= p0
        }
        fmt.Println(p0)
    } else {
        p0 += (p0 % (((^(p0))) | 1))
        p0 += -32768
        switch -9223372036854775808 {
        case 717:
            for fi14 := 0; fi14 < 3; fi14++ {
                if len(v7) > 0 {
                    v7[0] = fi14
                }
                _ = fi14
            }
        case -2147483648:
            var v15 int8 = 42
            _ = v15
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return false
}

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[int]string)
    var v3 bool = false
    v4 := make(map[uint32]uint32)
    v5 := []string{}
    var v6 float32 = (-2.25 - 3.14159)
    v16 := GenS1{F0: v3}
    _ = len(v2)
    v1 += 559
    for fi17 := 0; fi17 < 18; fi17++ {
        if v3 {
            var v18 int32 = -2147483648
            fmt.Println(v18)
            _ = v18
        } else {
            if len(v5) > 0 {
                v5[0] = "a"
            }
        }
        _ = len(v4)
        fmt.Println(v6)
        _ = fi17
    }
    switch -2147483648 {
    case -426:
        if v3 {
            for fi19 := 0; fi19 < 4; fi19++ {
                for fi20 := 0; fi20 < 1; fi20++ {
                    if v3 {
                        fmt.Println(fi19)
                        var v21 int = 4294967295
                        fmt.Println(v21)
                        _ = v21
                    }
                    v1 = v1
                    _ = fi20
                }
                var v22 uint = 255
                _ = fi19
                _ = v22
            }
            if true {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v23 string = "a"
                _ = v23
            }
        }
    case 32768:
        for fi24 := 0; fi24 < 6; fi24++ {
            var v25 int8 = 1
            v1 += (v1 + v1)
            _ = fi24
            _ = v25
        }
    default:
        v26 := GenS2{F0: 32767}
        _ = v26
    }
    var v27 float32 = v6
    v2[-889] = ""
    if ((503 > 32767) || true) {
        var v28 uint8 = 42
        _ = v28
    }
    fmt.Println(v27)
    v29 := GenS3{F0: 0, F1: 42}
    var v30 uint = 2
    if v3 {
        v3 = v3
        for fi31 := 0; fi31 < 5; fi31++ {
            for fi32 := 0; fi32 < 14; fi32++ {
                v1 -= (^(v1))
                _ = len(v4)
                if v3 {
                    v30 = (v30 * (v30 * (v30 ^ v30)))
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v33 := GenS2{F0: 42}
                    _ = v33
                }
                _ = fi32
            }
            _ = fi31
        }
        v6 = -2.25
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        fmt.Println(v1, v27, v6)
        if ((42 != 42) && (!v3)) {
            fmt.Println(v3, v1)
            fmt.Println(v6, v3, v1)
            for fi34 := 0; fi34 < 5; fi34++ {
                if (!false) {
                    v35 := GenS4{F0: "\n\t"}
                    v36 := GenS3{F0: (v30 ^ v30), F1: 1}
                    _ = v35
                    _ = v36
                } else {
                    var v37 uint = 369
                    var v38 uint8 = 128
                    _ = v37
                    _ = v38
                }
                _ = fi34
            }
        } else {
            v39 := GenS5{F0: ("foo bar" + "foo bar"), F1: "\n\t"}
            _ = v39
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
