package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint64; F1 int32; F2 int8 }
func genFn0(p0 int16, p1 int64, p2 bool) uint {
    v5 := []int{}
    v5 = append(v5, int(p1))
    switch -418 {
    case 882:
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 uint64 = 4294967295
            v7 -= v7
            if (!p2) {
                _ = len(v5)
                fmt.Println(fi6)
            } else {
                p2 = (2 >= (v7 % (((v7 / ((v7) | 1))) | 1)))
            }
            _ = fi6
            _ = v7
        }
    case 679:
        v8 := GenS1{F0: (p1 % ((1) | 1)), F1: p2}
        _ = v8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS1{F0: (p1 - (p1 ^ -9223372036854775808)), F1: false}
        v5 = append(v5, (fi9 ^ (fi9 & (fi9 / ((fi9) | 1)))))
        var v11 bool = p2
        _ = fi9
        _ = v10
        _ = v11
    }
    if p2 {
        if (!(true && p2)) {
            for fi12 := 0; fi12 < 16; fi12++ {
                for fi13 := 0; fi13 < 4; fi13++ {
                    var v14 float32 = -2.25
                    _ = fi13
                    _ = v14
                }
                _ = fi12
            }
            if len(v5) > 0 {
                v5[0] = 9223372036854775807
            }
            p2 = (42 < 100)
        } else {
            for fi15 := 0; fi15 < 9; fi15++ {
                if (!p2) {
                    v5 = append(v5, 9223372036854775807)
                } else {
                    v16 := GenS1{F0: p1, F1: true}
                    _ = len(v5)
                    _ = v16
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v17 := GenS1{F0: (^(p1)), F1: p2}
                _ = fi15
                _ = v17
            }
            if p2 {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            } else {
                p2 = p2
                fmt.Println(p0)
                for fi18 := 0; fi18 < 2; fi18++ {
                    switch fi18 {
                    case 9223372036854775807:
                        v5 = append(v5, fi18)
                    default:
                        p2 = ("a" < (("" + ("\n\t" + ("\n\t" + (("\n\t" + "hello") + ("\n\t" + ("\n\t" + "a")))))) + ""))
                    }
                    var v19 int8 = -128
                    if (false || p2) {
                        fmt.Println(v19, v19, p0)
                    }
                    _ = fi18
                    _ = v19
                }
            }
            v5 = append(v5, 9223372036854775807)
        }
        for fi20 := 0; fi20 < 8; fi20++ {
            for fi21 := 0; fi21 < 2; fi21++ {
                p2 = (float32(0.0) == 0.1)
                _ = fi21
            }
            var v22 uint16 = 807
            _ = fi20
            _ = v22
        }
        var v23 int32 = 0
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 65535
}

func main() {
    var v1 bool = (42 == 1000)
    var v2 int64 = 9223372036854775807
    var v3 int64 = (v2 & v2)
    var v4 uint16 = 844
    v24 := []int{-337}
    if len(v24) > 0 {
        v24[0] = -104
    }
    v2 += (-(v3))
    v2 += v2
    if (879 == 2147483648) {
        v25 := GenS1{F0: v2, F1: ((("" + ("foo bar" + "")) < (("" + (("a" + "foo bar") + "a")) + "hello")) && true)}
        _ = v25
    }
    var v26 uint = 65535
    if v1 {
        for fi27 := 0; fi27 < 11; fi27++ {
            switch (fi27 / (((^(fi27))) | 1)) {
            case -2147483648:
                if len(v24) >= 2 {
                    v24 = v24[1:len(v24)]
                }
            case 0:
                var v28 bool = (-32768 == -1)
                _ = v28
            }
            v2 += 1
            v4 = (v4 << ((genFn0(973, (v2 % ((v2) | 1)), false)) & 63))
            _ = fi27
        }
    }
    if true {
        var v29 int64 = v2
        for fi30 := 0; fi30 < 17; fi30++ {
            for fi31 := 0; fi31 < 6; fi31++ {
                _ = len(v24)
                switch "a" {
                case "a":
                    var v32 int = (^(fi30))
                    _ = v32
                }
                _ = fi31
            }
            v24 = append(v24, fi30)
            fmt.Println(v2, v1)
            _ = fi30
        }
        v24 = append(v24, 4294967295)
        _ = v29
    } else {
        v4 -= v4
        var v33 uint16 = v4
        _ = v33
    }
    v34 := GenS1{F0: (v2 ^ (v3 & (v3 << ((v26) & 63)))), F1: ((^(v2)) >= (v2 & (v3 / (((v3 % (((^(v2))) | 1))) | 1))))}
    v26 = v26
    v35 := GenS2{F0: 18446744073709551615, F1: 2147483647, F2: 127}
    v1 = ((0.1 / float64(1.5)) == float64(float32(-2.25)))
    var v36 int64 = (-(v2))
    var v37 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
