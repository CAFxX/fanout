package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS4 struct { F0 uint8; F1 string; F2 uint8 }
type GenS5 struct { F0 int32; F1 int64 }
type GenS6 struct { F0 int32; F1 int32 }
type GenS7 struct { F0 uint16 }
func genFn0(p0 int16, p1 int32, p2 int) uint16 {
    v7 := GenS1{F0: 255}
    switch -9223372036854775808 {
    case 2147483647:
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := GenS1{F0: 876}
            _ = fi8
            _ = v9
        }
    case -345:
        var v10 uint16 = 42
        _ = v10
    case 32767:
        if true {
            for fi11 := 0; fi11 < 15; fi11++ {
                v12 := GenS2{F0: 255, F1: 429}
                p1 -= -1
                _ = fi11
                _ = v12
            }
        }
    }
    v13 := GenS1{F0: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v7
    return 232
}
func genFn1(p0 uint) string {
    v14 := []int{}
    var v15 uint8 = 199
    _ = p0
    _ = v14
    _ = v15
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = -2.25
    var v3 uint8 = 2
    var v4 float64 = 1.5
    var v5 int = 0
    v6 := make(map[uint32]string)
    fmt.Println(v3, v2)
    var v16 int64 = -9223372036854775808
    v17 := []int{988, (v5 / ((v5) | 1)), -977}
    var v18 int64 = (v16 % (((v16 << ((1) & 63))) | 1))
    if (-1 <= 32767) {
        for fi19 := 0; fi19 < 13; fi19++ {
            var v20 string = "\n\t"
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 20; fi21++ {
            switch ("a" + "hello") {
            case "\n\t":
                var v22 int = v5
                _ = v22
            case "foo bar":
                var v23 int = v5
                _ = v23
            default:
                if false {
                    var v24 string = ""
                    for fi25 := 0; fi25 < 9; fi25++ {
                        var v26 int8 = -18
                        v26 -= (v26 >> ((2) & 63))
                        fmt.Println(v16, fi21)
                        _ = fi25
                        _ = v26
                    }
                    _ = v24
                } else {
                    v27 := GenS3{F0: (v3 % ((v3) | 1)), F1: 229, F2: 2}
                    var v28 uint64 = 42
                    var v29 uint32 = uint32(v28)
                    _ = v27
                    _ = v28
                    _ = v29
                }
            }
            v30 := GenS4{F0: v3, F1: (("a" + "") + ""), F2: 48}
            _ = fi21
            _ = v30
        }
    } else {
        _ = len(v6)
        var v31 bool = false
        if ((v3 >> ((2) & 63)) == (v3 >> ((65535) & 63))) {
            if (("foo bar" + ("a" + "foo bar")) < ("" + "\n\t")) {
                for fi32 := 0; fi32 < 20; fi32++ {
                    v33 := GenS5{F0: 1, F1: int64(v5)}
                    _ = fi32
                    _ = v33
                }
                var v34 bool = false
                _ = v34
            } else {
                var v35 int16 = -32768
                fmt.Println(v2, v5, v2)
                _ = v35
            }
            var v36 uint16 = 938
            var v37 float64 = (((v4 / ((float64(((1.5 - -2.25) + (-2.25 - -2.25))) * ((v4 * v2) / v4)) * 0.1)) / v2) - (1.5 - v4))
            _ = v36
            _ = v37
        } else {
            v38 := GenS6{F0: 0, F1: int32(v3)}
            v39 := GenS2{F0: v3, F1: 18446744073709551615}
            v6[2] = "\n\t"
            _ = v38
            _ = v39
        }
        _ = v31
    }
    v17 = append(v17, (v5 + (v5 >> ((84) & 63))))
    for fi40 := 0; fi40 < 19; fi40++ {
        var v41 string = ("" + "foo bar")
        v42 := GenS3{F0: v3, F1: 1, F2: 42}
        v16 = (v16 + (v16 ^ (v18 | (v16 & (v16 % ((v16) | 1))))))
        _ = fi40
        _ = v41
        _ = v42
    }
    v43 := GenS1{F0: 255}
    var v44 uint32 = 2147483648
    v44 += (v44 + (v44 * 2147483647))
    var v45 string = "a"
    v46 := GenS7{F0: 1}
    fmt.Println(v3, v3, v18)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
