package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int64; F2 uint16 }
func genFn0(p0 bool, p1 string) int16 {
    fmt.Println(p1, p0)
    v7 := []int{}
    p1 = (p1 + (p1 + ((p1 + (p1 + "a")) + "foo bar")))
    switch "foo bar" {
    case "":
        _ = len(v7)
    default:
        v8 := GenS1{F0: 364, F1: 639}
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        v7 = append(v7, -522)
        fmt.Println(p1, p1)
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v7
    return -1
}
func genFn1() int64 {
    v10 := []int{-1, 978, -114, -2147483648}
    var v11 float32 = (-2.25 + 3.14159)
    fmt.Println(v11)
    var v12 uint8 = 128
    _ = v10
    _ = v11
    _ = v12
    return 892
}

func main() {
    var v1 float64 = (-2.25 / float64(0.1))
    var v2 int = -2147483648
    var v3 uint32 = 2147483647
    v4 := make(map[uint32]bool)
    var v5 uint64 = 2
    var v6 float64 = v1
    fmt.Println(v6, v1)
    v3 -= 0
    fmt.Println(v6)
    for fi13 := 0; fi13 < 12; fi13++ {
        v6 = v1
        var v14 float32 = 0.0
        var v15 uint16 = 249
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        fmt.Println(v1)
        var v17 float64 = 0.0
        var v18 uint16 = uint16(v3)
        _ = fi16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        fmt.Println(fi19, v3, v5)
        _ = fi19
    }
    v20 := GenS1{F0: 1, F1: 0}
    v5 += (^(v5))
    v21 := GenS2{F0: 187, F1: -1, F2: 32767}
    v4[v3] = ((true && true) || (v1 > v1))
    v6 -= ((v1 - 3.14159) / 0.1)
    v1 += v6
    v22 := GenS1{F0: 0, F1: 1}
    for fi23 := 0; fi23 < 12; fi23++ {
        for fi24 := 0; fi24 < 3; fi24++ {
            v25 := v4[(v3 >> ((817) & 63))]
            var v26 int64 = -978
            _ = fi24
            _ = v25
            _ = v26
        }
        if false {
            var v27 string = "\n\t"
            for fi28 := 0; fi28 < 19; fi28++ {
                v4[v3] = false
                _ = fi28
            }
            v27 = v27
            _ = v27
        }
        switch "a" {
        case "foo bar":
            if false {
                for fi29 := 0; fi29 < 16; fi29++ {
                    switch "" {
                    case "hello":
                        var v30 bool = (v6 <= -2.25)
                        _ = v30
                    case "\n\t":
                        fmt.Println(v1)
                    default:
                        v2 = v2
                    }
                    _ = fi29
                }
                if true {
                    switch 42 {
                    case 1:
                        switch v2 {
                        case 32768:
                            v31 := []int{}
                            _ = v31
                        case 2147483647:
                            for fi32 := 0; fi32 < 1; fi32++ {
                                v33 := GenS2{F0: 1, F1: 42, F2: 985}
                                if (!(!false)) {
                                    var v34 uint16 = 42
                                    if (float64(float32((0.0 + 1.5))) <= 0.1) {
                                        v35 := GenS2{F0: 255, F1: -751, F2: (v34 / (((^(v34))) | 1))}
                                        var v36 int = v2
                                        var v37 float32 = 0.1
                                        fmt.Println(v3)
                                        var v38 uint64 = 1
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                    } else {
                                        v3 = v3
                                        v39 := []int{32768, fi32}
                                        _ = v39
                                    }
                                    v5 += (v5 ^ v5)
                                    _ = v34
                                }
                                _ = fi32
                                _ = v33
                            }
                        case 853:
                            var v40 float32 = 0.1
                            _ = v40
                        }
                    case 523:
                        var v41 string = (("foo bar" + ("hello" + "\n\t")) + "")
                        _ = v41
                    }
                }
            }
        case "hello":
            v5 = v5
        default:
            if true {
                _ = len(v4)
                delete(v4, v3)
                var v42 int16 = -712
                _ = v42
            } else {
                v43 := GenS2{F0: 0, F1: genFn1(), F2: 788}
                var v44 uint8 = 101
                _ = v43
                _ = v44
            }
        }
        _ = fi23
    }
    if (1000 != 1000) {
        v6 = ((v6 - (v6 / v6)) * (((0.0 * float64(0.0)) * v1) * v6))
    } else {
        if (!(!false)) {
            v4[v3] = (((false || false) || (!true)) || true)
        }
        for fi45 := 0; fi45 < 17; fi45++ {
            fmt.Println(fi45)
            switch (("foo bar" + ("" + ("a" + ""))) + "a") {
            case "\n\t":
                v2 = 1
            case "hello":
                fmt.Println(v5)
            }
            _ = fi45
        }
    }
    var v46 uint64 = (v5 + (v5 << ((1) & 63)))
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
