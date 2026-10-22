package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int16 }
func genFn0(p0 uint) int64 {
    v7 := GenS1{F0: 659, F1: 128, F2: 1000}
    var v8 uint16 = 1
    switch ("hello" + "\n\t") {
    case "foo bar":
        var v9 string = "a"
        _ = v9
    case "":
        if true {
            v10 := GenS1{F0: 4294967295, F1: 1, F2: -32768}
            for fi11 := 0; fi11 < 11; fi11++ {
                if ((v8 | v8) < v8) {
                    if ("foo bar" < ("a" + "foo bar")) {
                        for fi12 := 0; fi12 < 18; fi12++ {
                            v13 := []int{(^(fi12))}
                            var v14 int16 = int16(v8)
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                    } else {
                        v8 = v8
                        p0 -= 65535
                        v15 := []int{fi11, fi11, fi11}
                        _ = v15
                    }
                    p0 -= 0
                } else {
                    p0 = 255
                }
                fmt.Println(v8, v8)
                var v16 int = (fi11 / ((fi11) | 1))
                _ = fi11
                _ = v16
            }
            _ = v10
        }
    default:
        if true {
            p0 -= p0
        } else {
            v8 += v8
            var v17 uint16 = v8
            _ = v17
        }
    }
    v18 := GenS1{F0: 18446744073709551615, F1: 2, F2: 1000}
    v19 := GenS1{F0: 18446744073709551615, F1: 42, F2: -32768}
    _ = p0
    _ = v18
    _ = v19
    _ = v7
    _ = v8
    return 691
}
func genFn1(p0 uint16, p1 int) uint8 {
    v20 := []int{(-(p1)), (p1 ^ (-(p1)))}
    var v21 string = "a"
    switch p1 {
    case 32768:
        var v22 float64 = 1e10
        _ = v22
    case 32767:
        p0 -= 42
    }
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 4294967295
    var v3 int32 = -478
    v4 := []int{-853}
    var v5 uint8 = 255
    var v6 uint8 = (v5 * 1)
    var v23 int32 = v3
    v23 += -1
    v4 = append(v4, 9223372036854775807)
    var v24 float64 = (0.1 + ((0.1 * 0.0) + float64(1.5)))
    if len(v4) > 0 {
        v4[0] = -9223372036854775808
    }
    switch 42 {
    case -2147483648:
        for fi25 := 0; fi25 < 6; fi25++ {
            switch "" {
            case "":
                v26 := v1[(fi25 + fi25)]
                _ = v26
            case "hello":
                var v27 int16 = -166
                _ = v27
            }
            _ = fi25
        }
    case 4294967295:
        var v28 string = "\n\t"
        _ = v28
    default:
        _ = len(v4)
    }
    fmt.Println(v6, v2)
    for fi29 := 0; fi29 < 2; fi29++ {
        v30 := GenS1{F0: 409, F1: genFn1(2, fi29), F2: -689}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi31 := 0; fi31 < 10; fi31++ {
            v24 += (1.5 / 1.0)
            _ = fi31
        }
        _ = fi29
        _ = v30
    }
    fmt.Println(v2)
    if true {
        var v32 int8 = 121
        if true {
            v3 = (-(v23))
        }
        _ = v32
    }
    v6 -= (^(v5))
    v33 := GenS1{F0: 42, F1: genFn1(692, 4294967295), F2: 42}
    var v34 int64 = -1
    v35 := GenS1{F0: 461, F1: 134, F2: -1}
    v36 := GenS1{F0: 1, F1: v6, F2: -969}
    if (241 >= 4294967295) {
        for fi37 := 0; fi37 < 4; fi37++ {
            if (-128 >= 127) {
                var v38 uint = (v2 + (v2 & uint(v5)))
                v34 += genFn0(701)
                _ = v38
            }
            fmt.Println(v3)
            _ = fi37
        }
        switch 9223372036854775807 {
        case 890:
            var v39 int8 = -1
            _ = v39
        case -874:
            var v40 int64 = (v34 ^ genFn0(v2))
            _ = v40
        default:
            fmt.Println(v24)
        }
        var v41 int = -6
        _ = v41
    } else {
        var v42 int = -702
        var v43 string = "hello"
        _ = v42
        _ = v43
    }
    v44 := GenS1{F0: 2, F1: v5, F2: -32768}
    var v45 int64 = (v34 % (((v34 - int64(v23))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
