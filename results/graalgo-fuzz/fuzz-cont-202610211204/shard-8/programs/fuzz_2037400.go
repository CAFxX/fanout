package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint32; F2 int32 }
func genFn0(p0 int64) bool {
    fmt.Println(p0)
    v7 := GenS1{F0: ((0 != 799) && (!true))}
    if (2147483647 <= 42) {
        if false {
            v8 := []int{}
            if false {
                p0 = p0
                v8 = append(v8, 2147483647)
                v8 = append(v8, -1)
            } else {
                var v9 int = 32768
                if false {
                    v8 = append(v8, (v9 | v9))
                    v9 = v9
                    fmt.Println(v9)
                } else {
                    if len(v8) > 0 {
                        v8[0] = (v9 % ((v9) | 1))
                    }
                }
                p0 += -893
                _ = v9
            }
            v10 := GenS1{F0: true}
            _ = v10
            _ = v8
        } else {
            v11 := GenS2{F0: 42, F1: 42, F2: -1}
            _ = v11
        }
        fmt.Println(p0)
        p0 += 42
    } else {
        p0 -= 1
        var v12 uint = 11
        fmt.Println(p0)
        _ = v12
    }
    fmt.Println(p0)
    var v13 uint8 = 183
    _ = p0
    _ = v13
    _ = v7
    return (-1 == (p0 + (p0 & (-(p0)))))
}
func genFn1(p0 int, p1 uint32, p2 bool) uint8 {
    var v14 uint8 = 251
    switch ("hello" + "\n\t") {
    case "foo bar":
        for fi15 := 0; fi15 < 6; fi15++ {
            if p2 {
                var v16 string = ("a" + ("a" + ("\n\t" + "")))
                switch -2147483648 {
                case 9223372036854775807:
                    var v17 int8 = -35
                    _ = v17
                case -9223372036854775808:
                    fmt.Println(v14, fi15)
                case 1:
                    var v18 uint = 744
                    _ = v18
                }
                _ = v16
            } else {
                for fi19 := 0; fi19 < 10; fi19++ {
                    v20 := []int{32767, fi15}
                    _ = fi19
                    _ = v20
                }
            }
            v14 -= v14
            _ = fi15
        }
    default:
        p2 = p2
    }
    v14 = v14
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return v14
}

func main() {
    var v1 uint8 = 42
    v2 := []uint32{4294967295, 2147483648, 2147483647, 155}
    var v3 int16 = 686
    var v4 float64 = 0.0
    v5 := make(map[int]int)
    var v6 uint32 = 826
    v2 = append(v2, v6)
    fmt.Println(v1, v1, v1)
    if (v3 >= -697) {
        switch "\n\t" {
        case "":
            if true {
                var v21 uint16 = 255
                for fi22 := 0; fi22 < 1; fi22++ {
                    if true {
                        fmt.Println(fi22, fi22)
                        v1 = (v1 | genFn1((-(fi22)), 200, true))
                    }
                    _ = fi22
                }
                var v23 int = -2147483648
                _ = v21
                _ = v23
            } else {
                var v24 uint = 65535
                for fi25 := 0; fi25 < 4; fi25++ {
                    v26 := GenS1{F0: false}
                    _ = fi25
                    _ = v26
                }
                _ = v24
            }
        }
    } else {
        fmt.Println(v3)
    }
    v27 := GenS1{F0: false}
    v1 = uint8(v6)
    for fi28 := 0; fi28 < 4; fi28++ {
        if (genFn1((fi28 ^ (fi28 >> ((255) & 63))), 157, (!(1 <= 32767))) < (v1 * (v1 & v1))) {
            var v29 uint16 = 703
            if (true && true) {
                v6 = 42
            }
            _ = len(v2)
            _ = v29
        }
        _ = fi28
    }
    fmt.Println(v3, v3, v6)
    v1 -= v1
    switch 0 {
    case 9223372036854775807:
        var v30 int8 = 0
        _ = v30
    case -2147483648:
        for fi31 := 0; fi31 < 2; fi31++ {
            v1 = (v1 * (^(v1)))
            fmt.Println(v1, v3, v6)
            _ = fi31
        }
    case 932:
        var v32 uint64 = uint64(v6)
        _ = v32
    }
    fmt.Println(v1, v4, v3)
    var v33 uint = 438
    v3 += v3
    v6 = (^(v6))
    _ = len(v2)
    for fi34 := 0; fi34 < 5; fi34++ {
        for fi35 := 0; fi35 < 15; fi35++ {
            if false {
                _ = len(v2)
                v36 := GenS1{F0: true}
                _ = v36
            } else {
                var v37 string = "hello"
                _ = v37
            }
            var v38 int64 = -9223372036854775808
            v4 = (-2.25 - v4)
            v1 += genFn1(9223372036854775807, (v6 / ((433) | 1)), (v1 < (v1 * 182)))
            _ = fi35
            _ = v38
        }
        var v39 int16 = -76
        _ = fi34
        _ = v39
    }
    if (v6 <= (v6 ^ 2147483648)) {
        var v40 uint8 = 127
        v41 := GenS2{F0: (-(v3)), F1: 420, F2: 42}
        var v42 uint32 = (v6 * (v6 / ((2) | 1)))
        _ = v40
        _ = v41
        _ = v42
    }
    var v43 int = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
