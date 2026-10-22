package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS3 struct { F0 int64; F1 uint64; F2 uint }

func main() {
    var v1 uint64 = 934
    v2 := make(map[int]string)
    var v3 int64 = 1
    var v4 int8 = 100
    v5 := make(map[uint32]int)
    v6 := []string{(("a" + ("foo bar" + "foo bar")) + "")}
    v4 -= v4
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 float64 = 1.5
        fmt.Println(v3)
        var v9 int8 = v4
        _ = fi7
        _ = v8
        _ = v9
    }
    switch 127 {
    case -3:
        for fi10 := 0; fi10 < 16; fi10++ {
            v11 := GenS1{F0: 70}
            var v12 int16 = -138
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    v1 = (v1 - 409)
    if false {
        v13 := GenS1{F0: 0}
        _ = v13
    } else {
        fmt.Println(v1)
        var v14 uint32 = 42
        var v15 int16 = 42
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: -2147483648}
    if false {
        for fi17 := 0; fi17 < 8; fi17++ {
            for fi18 := 0; fi18 < 12; fi18++ {
                if (false && (32 > 219)) {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    switch "\n\t" {
                    case "a":
                        delete(v2, (-(fi18)))
                    case "foo bar":
                        _ = len(v5)
                    case "\n\t":
                        v1 = (^(v1))
                    default:
                        var v19 string = "foo bar"
                        _ = v19
                    }
                }
                _ = fi18
            }
            for fi20 := 0; fi20 < 1; fi20++ {
                var v21 uint64 = v1
                if false {
                    _ = len(v2)
                    var v22 int32 = -2147483648
                    if true {
                        v23 := GenS1{F0: -1}
                        v6 = append(v6, "a")
                        _ = v23
                    } else {
                        switch ("hello" + "\n\t") {
                        case "foo bar":
                            v24 := GenS1{F0: (v22 ^ (v22 | v22))}
                            _ = v24
                        case "\n\t":
                            fmt.Println(fi17, v1, fi20)
                        case "hello":
                            v21 -= 18446744073709551615
                        default:
                            v3 -= v3
                        }
                        switch 127 {
                        case 2147483647:
                            if false {
                                var v25 uint64 = (v1 & 988)
                                _ = v25
                            }
                        case 1:
                            switch "hello" {
                            case "hello":
                                var v26 uint64 = (v1 % (((v1 / ((42) | 1))) | 1))
                                _ = v26
                            }
                        default:
                            v5[540] = fi17
                        }
                    }
                    _ = v22
                }
                fmt.Println(v21, v4)
                _ = fi20
                _ = v21
            }
            _ = fi17
        }
    } else {
        var v27 int64 = v3
        v27 -= (-(v3))
        _ = v27
    }
    v28 := GenS1{F0: 2147483647}
    var v29 uint8 = 220
    v5[2147483648] = 0
    if ((true || (32767 != -1)) || true) {
        if true {
            v1 = v1
            for fi30 := 0; fi30 < 7; fi30++ {
                v31 := GenS1{F0: -1}
                _ = fi30
                _ = v31
            }
            v32 := GenS2{F0: v1, F1: 679, F2: v29}
            _ = v32
        } else {
            v33 := GenS1{F0: 1}
            _ = v33
        }
        var v34 bool = (!(-2147483648 == 0))
        var v35 int = 729
        _ = v34
        _ = v35
    } else {
        var v36 uint = 42
        v37 := GenS3{F0: (-(v3)), F1: v1, F2: 1}
        switch "foo bar" {
        case "\n\t":
            delete(v2, 0)
        default:
            v3 = v3
        }
        _ = v36
        _ = v37
    }
    var v38 uint32 = 20
    var v39 uint64 = v1
    var v40 int8 = int8(v38)
    delete(v2, 765)
    v4 -= v40
    var v41 string = ("foo bar" + ("a" + "a"))
    switch 1 {
    case 128:
        v4 += (v40 / ((42) | 1))
    case 32767:
        fmt.Println(v41, v29, v4)
    case 0:
        var v42 uint16 = 262
        _ = v42
    default:
        var v43 float64 = 1.5
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
