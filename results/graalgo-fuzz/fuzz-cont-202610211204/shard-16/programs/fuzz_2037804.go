package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 bool }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 int; F2 int }
func genFn0(p0 string, p1 int8, p2 int16) int64 {
    if true {
        p0 = ((p0 + ("hello" + ("" + p0))) + "")
    }
    p1 += (p1 - p1)
    v5 := []int{42, -9223372036854775808, 4294967295, -198}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -1
}
func genFn1(p0 uint8) uint16 {
    fmt.Println(p0)
    for fi6 := 0; fi6 < 5; fi6++ {
        if (0.0 == 1.5) {
            p0 += (p0 - (p0 >> ((4294967295) & 63)))
        }
        p0 += 127
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := GenS1{F0: true, F1: fi6, F2: (true || ((42 == 947) || false))}
            v9 := []int{(fi7 + (fi6 / ((184) | 1))), fi6}
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    var v10 uint16 = uint16(p0)
    _ = p0
    _ = v10
    return (v10 >> ((255) & 63))
}

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = false
    var v3 int64 = -9223372036854775808
    var v4 bool = true
    for fi11 := 0; fi11 < 11; fi11++ {
        _ = len(v1)
        for fi12 := 0; fi12 < 10; fi12++ {
            v13 := GenS2{F0: 6, F1: v4}
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    fmt.Println(v4, v3)
    if v2 {
        var v14 float64 = 1.5
        switch 887 {
        case -2147483648:
            fmt.Println(v3)
        case 9223372036854775807:
            fmt.Println(v14)
        case 128:
            v4 = (988 > 2)
        }
        _ = v14
    } else {
        var v15 int8 = -36
        var v16 uint = 1
        _ = v15
        _ = v16
    }
    v17 := []int{-2147483648}
    var v18 bool = v4
    fmt.Println(v18, v4, v2)
    if false {
        v19 := GenS2{F0: 127, F1: (274 <= 42)}
        v18 = ((!(v4 || ((!v4) || v4))) && true)
        _ = v19
    } else {
        if true {
            v4 = (347 <= 1000)
            if (127 != -77) {
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
                switch -852 {
                case 127:
                    var v20 float64 = 0.0
                    _ = v20
                default:
                    for fi21 := 0; fi21 < 17; fi21++ {
                        for fi22 := 0; fi22 < 9; fi22++ {
                            v23 := GenS3{F0: 4294967295}
                            _ = fi22
                            _ = v23
                        }
                        _ = fi21
                    }
                }
                for fi24 := 0; fi24 < 15; fi24++ {
                    var v25 float32 = 1.5
                    v25 += (0.0 * ((1.5 * 0.1) - -2.25))
                    v26 := GenS4{F0: 18446744073709551615, F1: (-(fi24)), F2: (fi24 ^ -1)}
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            } else {
                if ((!(!v4)) && (!((((155 < 962) || false) && false) || (1 <= 1)))) {
                    var v27 uint = 65535
                    switch 4294967295 {
                    case -2147483648:
                        if (v18 && ((-183 != -1) && (255 > 220))) {
                            switch "a" {
                            case "a":
                                fmt.Println(v27, v2, v27)
                            case "hello":
                                v18 = (1 == -2147483648)
                            }
                            v2 = (!(true && (0.0 != 0.0)))
                        }
                    }
                    _ = v27
                } else {
                    if (!(1.5 == 1e10)) {
                        var v28 int64 = v3
                        delete(v1, 2)
                        v17 = append(v17, 127)
                        _ = v28
                    }
                }
                v29 := GenS1{F0: true, F1: 9223372036854775807, F2: v4}
                var v30 string = ""
                _ = v29
                _ = v30
            }
        } else {
            for fi31 := 0; fi31 < 11; fi31++ {
                v4 = (!((0.0 * (1e10 / 1e10)) > 1e10))
                _ = fi31
            }
            var v32 int16 = 1000
            _ = v32
        }
    }
    v18 = ((!v18) || v2)
    v3 = v3
    v17 = append(v17, 4294967295)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
