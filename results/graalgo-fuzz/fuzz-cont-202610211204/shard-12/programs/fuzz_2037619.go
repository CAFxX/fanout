package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int8; F1 int32; F2 int64 }
type GenS4 struct { F0 uint; F1 string; F2 uint }
func genFn0(p0 int16, p1 uint32, p2 uint32) int8 {
    switch ("foo bar" + "") {
    case "":
        fmt.Println(p2)
    }
    v6 := []int{1}
    if true {
        switch 1 {
        case 32768:
            v7 := GenS1{F0: -24}
            _ = v7
        case 901:
            var v8 bool = false
            _ = v8
        case 9223372036854775807:
            for fi9 := 0; fi9 < 13; fi9++ {
                for fi10 := 0; fi10 < 16; fi10++ {
                    var v11 float64 = 1e10
                    v6 = append(v6, fi9)
                    _ = fi10
                    _ = v11
                }
                v6 = append(v6, fi9)
                if true {
                    if (104 == 128) {
                        var v12 int16 = p0
                        if ((!(18446744073709551615 < 42)) && ((!(731 < 1)) && false)) {
                            v6 = append(v6, (fi9 << ((0) & 63)))
                            v6 = append(v6, (fi9 | (fi9 & fi9)))
                        }
                        if ((!(true || (!true))) && true) {
                            var v13 int8 = 0
                            fmt.Println(v13, p2)
                            _ = v13
                        } else {
                            if len(v6) > 0 {
                                v6[0] = 127
                            }
                            var v14 float64 = 0.1
                            _ = v14
                        }
                        _ = v12
                    }
                } else {
                    var v15 int8 = int8(p1)
                    _ = v15
                }
                _ = fi9
            }
        }
        v16 := GenS2{F0: (p2 > p1), F1: -1}
        p1 -= (^(p2))
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -1
}

func main() {
    var v1 int32 = 45
    v2 := make(map[uint32]string)
    var v3 float32 = (-2.25 + -2.25)
    v4 := make(map[uint32]bool)
    var v5 uint16 = 2
    var v17 uint = 2
    if false {
        v3 -= (v3 + float32((-2.25 / 1.5)))
    } else {
        v3 += v3
        var v18 uint8 = 127
        for fi19 := 0; fi19 < 10; fi19++ {
            v4[4294967295] = (uint32(v5) < 65535)
            v20 := []int{1}
            _ = fi19
            _ = v20
        }
        _ = v18
    }
    v21 := GenS1{F0: 1}
    v1 = (^(v1))
    v22 := GenS1{F0: int8(v17)}
    v1 = int32(v17)
    v23 := []int{-2147483648, 4294967295}
    for fi24 := 0; fi24 < 12; fi24++ {
        var v25 uint8 = 194
        v26 := GenS3{F0: -29, F1: v1, F2: 42}
        _ = fi24
        _ = v25
        _ = v26
    }
    v27 := GenS4{F0: 655, F1: ("foo bar" + ("a" + "hello")), F2: 2}
    fmt.Println(v5)
    for fi28 := 0; fi28 < 5; fi28++ {
        fmt.Println(v5, v5, v17)
        delete(v2, 2147483647)
        if (-1 <= -1) {
            var v29 float64 = (0.0 / 1.0)
            var v30 string = "foo bar"
            _ = v29
            _ = v30
        }
        _ = fi28
    }
    fmt.Println(v5)
    v31 := GenS3{F0: 42, F1: -57, F2: -9223372036854775808}
    _ = len(v23)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
