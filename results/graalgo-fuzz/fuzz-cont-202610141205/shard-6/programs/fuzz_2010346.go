package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint64; F2 int64 }
type GenS3 struct { F0 int64; F1 int16; F2 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 bool; F2 string }
func genFn0() uint32 {
    v5 := GenS1{F0: 1, F1: 0, F2: 2}
    v6 := []int{4294967295, 227}
    if (!(!true)) {
        v6 = append(v6, 1)
        var v7 uint = 2
        _ = v7
    } else {
        var v8 bool = (508 <= 192)
        var v9 int = 42
        _ = v8
        _ = v9
    }
    var v10 int16 = 1
    _ = v10
    _ = v5
    _ = v6
    return 990
}
func genFn1(p0 uint8, p1 int32) bool {
    fmt.Println(p1, p0)
    for fi11 := 0; fi11 < 7; fi11++ {
        if (-160 > 1000) {
            p1 = (p1 ^ (-(p1)))
        } else {
            v12 := []int{}
            v13 := GenS2{F0: 100, F1: 0, F2: 42}
            switch fi11 {
            case -695:
                var v14 int8 = 46
                _ = v14
            case 194:
                for fi15 := 0; fi15 < 6; fi15++ {
                    var v16 int16 = -32768
                    _ = fi15
                    _ = v16
                }
            default:
                fmt.Println(p1, p1)
            }
            _ = v12
            _ = v13
        }
        v17 := []int{32767, (fi11 | -9223372036854775808)}
        _ = fi11
        _ = v17
    }
    v18 := []int{-845, 42, 1, -1}
    _ = p0
    _ = p1
    _ = v18
    return (((-128 <= 1) && (-9223372036854775808 == 4294967295)) || (!(2 != 1)))
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 1
    var v3 uint64 = 2
    var v4 int64 = 42
    if (498 >= 2) {
        if true {
            var v19 float64 = (0.1 - (-2.25 * 1.5))
            _ = v19
        } else {
            var v20 int = 4294967295
            _ = v20
        }
        v3 = 788
        switch 4294967295 {
        case -1:
            for fi21 := 0; fi21 < 6; fi21++ {
                switch (fi21 - fi21) {
                case 336:
                    v22 := []int{(fi21 * -2147483648), -9223372036854775808, (fi21 / (((fi21 / ((fi21) | 1))) | 1)), fi21}
                    _ = v22
                case -316:
                    v23 := GenS3{F0: -696, F1: 42, F2: 84}
                    _ = v23
                case 9223372036854775807:
                    var v24 float32 = 1e10
                    _ = v24
                }
                var v25 uint16 = 42
                v26 := GenS4{F0: true}
                _ = fi21
                _ = v25
                _ = v26
            }
        case 42:
            v2 += v3
        }
    }
    var v27 int = -2147483648
    if false {
        for fi28 := 0; fi28 < 14; fi28++ {
            fmt.Println(v27, v4, v2)
            var v29 bool = (fi28 >= fi28)
            var v30 int8 = -7
            _ = fi28
            _ = v29
            _ = v30
        }
        var v31 string = (("foo bar" + "hello") + "")
        v32 := GenS2{F0: 100, F1: v2, F2: v4}
        _ = v31
        _ = v32
    } else {
        switch (v27 + v27) {
        case 32768:
            v2 -= (v2 - 916)
        case 127:
            switch "" {
            case "foo bar":
                v1[v27] = 1
            case "\n\t":
                for fi33 := 0; fi33 < 18; fi33++ {
                    for fi34 := 0; fi34 < 1; fi34++ {
                        if true {
                            var v35 int16 = 42
                            _ = v35
                        }
                        _ = fi34
                    }
                    for fi36 := 0; fi36 < 18; fi36++ {
                        if (!(!false)) {
                            v4 = -704
                        }
                        v27 += (fi36 & (fi36 + (fi36 | v27)))
                        var v37 uint32 = 65535
                        _ = fi36
                        _ = v37
                    }
                    for fi38 := 0; fi38 < 18; fi38++ {
                        v2 = 635
                        fmt.Println(v3, fi33, v2)
                        var v39 uint = 126
                        _ = fi38
                        _ = v39
                    }
                    _ = fi33
                }
            case "a":
                v27 -= v27
            default:
                var v40 int16 = -32768
                _ = v40
            }
        case -9223372036854775808:
            v41 := []int{v27, (v27 / ((v27) | 1)), (v27 >> ((255) & 63))}
            _ = v41
        }
    }
    v27 -= v27
    v42 := GenS5{F0: 41, F1: (1 < 42), F2: ("hello" + "")}
    fmt.Println(v4, v27, v2)
    v3 = v2
    var v43 int = v27
    v27 = v27
    var v44 string = "a"
    var v45 uint64 = 4294967295
    switch 42 {
    case -1:
        v46 := GenS2{F0: -27, F1: v3, F2: (v4 / ((v4) | 1))}
        _ = v46
    case 128:
        v4 = 433
    }
    v4 -= (v4 - (v4 >> ((4294967295) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
