package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 int }
func genFn0(p0 uint, p1 int8, p2 int8) int {
    if true {
        for fi4 := 0; fi4 < 5; fi4++ {
            var v5 string = ""
            fmt.Println(fi4, fi4, p1)
            _ = fi4
            _ = v5
        }
        v6 := GenS1{F0: -2147483648}
        _ = v6
    } else {
        p1 -= -98
        fmt.Println(p2)
        var v7 int16 = 1000
        _ = v7
    }
    switch ("foo bar" + "hello") {
    case "":
        var v8 bool = false
        _ = v8
    case "hello":
        var v9 uint16 = 2
        _ = v9
    default:
        for fi10 := 0; fi10 < 15; fi10++ {
            if (false || true) {
                var v11 float64 = (1.5 / 1.5)
                for fi12 := 0; fi12 < 19; fi12++ {
                    v13 := GenS2{F0: 32767, F1: 762}
                    _ = fi12
                    _ = v13
                }
                p2 = (p1 / ((100) | 1))
                _ = v11
            }
            p0 = 890
            if (!(!false)) {
                p2 += p1
                v14 := []int{int(p2), (fi10 / ((fi10) | 1)), (fi10 & (fi10 / ((fi10) | 1)))}
                p1 += int8(p0)
                _ = v14
            } else {
                if false {
                    fmt.Println(p0)
                    var v15 int16 = 0
                    _ = v15
                } else {
                    v16 := []int{651}
                    if true {
                        var v17 uint8 = 42
                        fmt.Println(v17)
                        var v18 uint = p0
                        _ = v17
                        _ = v18
                    }
                    if false {
                        var v19 float32 = -2.25
                        for fi20 := 0; fi20 < 6; fi20++ {
                            v19 = ((v19 / (0.1 / 1.0)) * 3.14159)
                            if len(v16) > 0 {
                                v16[0] = fi20
                            }
                            switch ("foo bar" + "\n\t") {
                            case "":
                                p0 -= p0
                            case "\n\t":
                                v16 = append(v16, -754)
                                var v21 uint = p0
                                var v22 int64 = 0
                                _ = v21
                                _ = v22
                            default:
                                p0 -= p0
                            }
                            _ = fi20
                        }
                        _ = v19
                    }
                    _ = v16
                }
                p1 = p2
            }
            _ = fi10
        }
    }
    v23 := GenS3{F0: 175}
    for fi24 := 0; fi24 < 1; fi24++ {
        var v25 string = "a"
        _ = fi24
        _ = v25
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return 127
}
func genFn1(p0 int8) bool {
    p0 = p0
    var v26 uint64 = 42
    fmt.Println(v26, p0)
    _ = p0
    _ = v26
    return true
}

func main() {
    var v1 int8 = 0
    var v2 uint8 = 94
    v3 := make(map[uint32]int)
    var v27 uint8 = (^(v2))
    if ((!false) || (0 == 1)) {
        for fi28 := 0; fi28 < 17; fi28++ {
            v2 -= (^(v27))
            _ = fi28
        }
    } else {
        fmt.Println(v1, v1, v2)
        fmt.Println(v2, v2)
        v1 -= (v1 / ((77) | 1))
    }
    for fi29 := 0; fi29 < 16; fi29++ {
        for fi30 := 0; fi30 < 8; fi30++ {
            v3[709] = fi30
            var v31 int = fi29
            var v32 float32 = -2.25
            _ = fi30
            _ = v31
            _ = v32
        }
        v2 = v2
        _ = fi29
    }
    v27 = v27
    var v33 int64 = -1
    var v34 uint64 = 4294967295
    v27 = (v2 >> ((748) & 63))
    switch ("hello" + "") {
    case "":
        for fi35 := 0; fi35 < 4; fi35++ {
            for fi36 := 0; fi36 < 18; fi36++ {
                v37 := []int{}
                if len(v37) > 0 {
                    v37[0] = (fi36 >> ((1) & 63))
                }
                var v38 string = "\n\t"
                _ = fi36
                _ = v37
                _ = v38
            }
            fmt.Println(v33)
            var v39 uint64 = v34
            _ = fi35
            _ = v39
        }
    }
    v34 = v34
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v34
}
