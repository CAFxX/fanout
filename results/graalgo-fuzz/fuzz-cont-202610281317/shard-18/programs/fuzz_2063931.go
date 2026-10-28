package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16; F1 int; F2 string }
func genFn0(p0 string) uint {
    fmt.Println(p0)
    v4 := []int{32768, 83, -280}
    _ = p0
    _ = v4
    return 2
}
func genFn1(p0 bool, p1 int64, p2 int8) int64 {
    for fi5 := 0; fi5 < 4; fi5++ {
        fmt.Println(p2, p2, p1)
        var v6 uint8 = 167
        p1 = 42
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := GenS1{F0: p1, F1: p1}
        v9 := []int{(^(fi7))}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        v11 := GenS2{F0: "", F1: p0}
        v12 := GenS1{F0: 1, F1: (p1 - 0)}
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 string = "a"
    if (!true) {
        v14 := []int{32768, -746, 42, 4294967295}
        v13 = (v13 + "hello")
        _ = v14
    } else {
        var v15 int64 = p1
        for fi16 := 0; fi16 < 13; fi16++ {
            if ((1 != 2) && true) {
                v15 += p1
                var v17 uint64 = 18446744073709551615
                _ = v17
            }
            _ = fi16
        }
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return p1
}

func main() {
    var v1 float64 = 0.0
    var v2 uint64 = 18446744073709551615
    v3 := []int{}
    v18 := GenS3{F0: 118, F1: 0}
    _ = len(v3)
    for fi19 := 0; fi19 < 1; fi19++ {
        if false {
            for fi20 := 0; fi20 < 13; fi20++ {
                if len(v3) > 0 {
                    v3[0] = fi20
                }
                v21 := GenS4{F0: 2}
                v22 := GenS4{F0: 1}
                _ = fi20
                _ = v21
                _ = v22
            }
            v1 -= (v1 - v1)
        } else {
            v1 -= 0.1
        }
        v23 := GenS3{F0: fi19, F1: 0}
        switch "\n\t" {
        case "hello":
            fmt.Println(v2, v2, fi19)
        case "foo bar":
            v3 = append(v3, (fi19 | 417))
        }
        _ = fi19
        _ = v23
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v1 = float64(-2.25)
        var v25 uint64 = (v2 + (v2 % (((^(v2))) | 1)))
        _ = fi24
        _ = v25
    }
    if true {
        for fi26 := 0; fi26 < 6; fi26++ {
            if len(v3) > 0 {
                v3[0] = (fi26 - int(v2))
            }
            if len(v3) > 0 {
                v3[0] = fi26
            }
            if (42 >= 0) {
                v3 = append(v3, 32767)
                v27 := GenS2{F0: ("\n\t" + ("foo bar" + "\n\t")), F1: true}
                _ = v27
            } else {
                if ((-2147483648 != 2147483647) && true) {
                    var v28 int = -254
                    switch (fi26 * (v28 + v28)) {
                    case -1:
                        v1 = (v1 * ((v1 + v1) + -2.25))
                    case 0:
                        switch 2147483647 {
                        case -1:
                            fmt.Println(v28, fi26, v2)
                        default:
                            switch v28 {
                            case -2147483648:
                                fmt.Println(v28, v28, v1)
                            case 13:
                                var v29 uint8 = 2
                                _ = v29
                            case 128:
                                v28 -= (v28 / (((-(fi26))) | 1))
                            }
                        }
                    default:
                        switch v28 {
                        case 32767:
                            v30 := GenS5{F0: 0, F1: fi26, F2: ("hello" + "a")}
                            _ = v30
                        case 0:
                            for fi31 := 0; fi31 < 14; fi31++ {
                                fmt.Println(v1, v2, v28)
                                _ = fi31
                            }
                        case 1:
                            v1 = (v1 + v1)
                        default:
                            if (1 == 100) {
                                v1 += 0.0
                            }
                        }
                    }
                    _ = v28
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                if false {
                    fmt.Println(v2, v2)
                    if false {
                        v32 := GenS5{F0: 973, F1: (fi26 - 964), F2: ("a" + "\n\t")}
                        var v33 int64 = 9223372036854775807
                        _ = v32
                        _ = v33
                    }
                }
                v1 += v1
            }
            _ = fi26
        }
        for fi34 := 0; fi34 < 13; fi34++ {
            switch fi34 {
            case 642:
                var v35 int32 = -1
                _ = v35
            default:
                fmt.Println(v1, fi34)
            }
            _ = fi34
        }
    } else {
        v2 += v2
        v3 = append(v3, 32767)
    }
    v1 = -2.25
    var v36 int = 42
    v2 += 0
    fmt.Println(v1, v1, v36)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v36
}
