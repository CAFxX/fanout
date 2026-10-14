package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int32; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint; F1 uint; F2 uint16 }
type GenS4 struct { F0 uint8 }
func genFn0() int16 {
    var v6 uint64 = 0
    v6 = 599
    var v7 int8 = 100
    v6 = 2
    var v8 int = 32768
    _ = v6
    _ = v7
    _ = v8
    return -32768
}
func genFn1() uint32 {
    v9 := GenS1{F0: 128, F1: 0}
    var v10 uint64 = 4294967295
    v10 -= (v10 & 2)
    _ = v10
    _ = v9
    return 776
}

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{42, 187, 2}
    var v3 float32 = -2.25
    var v4 int16 = 0
    var v5 string = "hello"
    for fi11 := 0; fi11 < 5; fi11++ {
        if (127 >= 100) {
            for fi12 := 0; fi12 < 5; fi12++ {
                if false {
                    v2 = append(v2, 65535)
                } else {
                    v13 := GenS2{F0: 42, F1: 596, F2: 0}
                    _ = v13
                }
                _ = fi12
            }
            fmt.Println(v4)
        }
        _ = fi11
    }
    if false {
        _ = len(v1)
        if ((0.0 <= ((1.5 / v3) * -2.25)) || ((false || ((2147483647 > -551) || ((("hello" + v5) + v5) > (v5 + v5)))) || true)) {
            switch 42 {
            case 128:
                v1[4294967295] = v5
            case 730:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case -9223372036854775808:
                for fi14 := 0; fi14 < 18; fi14++ {
                    switch fi14 {
                    case -1:
                        _ = len(v1)
                    }
                    v15 := GenS2{F0: -12, F1: 2, F2: 1}
                    var v16 float64 = 0.1
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
            v17 := GenS1{F0: 4294967295, F1: 1}
            fmt.Println(v5, v3)
            _ = v17
        } else {
            if len(v2) > 0 {
                v2[0] = 65535
            }
            v3 = (0.1 + 1.5)
            v2 = append(v2, 42)
        }
    } else {
        v1[1] = ""
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 uint = 430
        _ = fi18
        _ = v19
    }
    if (("foo bar" + v5) <= v5) {
        if (0.0 < 1e10) {
            var v20 int16 = v4
            _ = v20
        }
        fmt.Println(v5, v3)
        v3 += (v3 / v3)
    }
    v21 := GenS2{F0: -2147483648, F1: 1, F2: -35}
    if ((v4 % (((v4 * v4)) | 1)) > 32767) {
        fmt.Println(v4)
    }
    if len(v2) > 0 {
        v2[0] = 65535
    }
    v5 = v5
    v5 = v5
    v4 = (-(v4))
    v4 = genFn0()
    for fi22 := 0; fi22 < 15; fi22++ {
        if true {
            var v23 uint = 843
            _ = v23
        }
        v24 := GenS3{F0: 42, F1: 2, F2: 729}
        v5 = ("foo bar" + (v5 + v5))
        _ = fi22
        _ = v24
    }
    if (((v4 | genFn0()) == 42) && true) {
        v3 = (1.5 / float32(float64(v3)))
        v25 := v1[32767]
        for fi26 := 0; fi26 < 17; fi26++ {
            fmt.Println(v4, v3, fi26)
            for fi27 := 0; fi27 < 12; fi27++ {
                v25 = ((((v5 + "hello") + ((v25 + "foo bar") + "foo bar")) + "foo bar") + (v25 + v5))
                v25 = v25
                fmt.Println(fi27)
                fmt.Println(v25)
                _ = fi27
            }
            v28 := GenS4{F0: 1}
            _ = fi26
            _ = v28
        }
        _ = v25
    }
    if len(v2) > 0 {
        v2[0] = 235
    }
    var v29 float32 = float32(1e10)
    for fi30 := 0; fi30 < 19; fi30++ {
        fmt.Println(v5, v3, v4)
        _ = fi30
    }
    _ = len(v1)
    if len(v2) > 0 {
        v2[0] = 828
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
