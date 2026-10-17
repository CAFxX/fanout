package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 bool; F2 int8 }
func genFn0() uint32 {
    var v7 uint8 = 1
    for fi8 := 0; fi8 < 16; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 uint64 = 407
            if ((479 != 42) && (float64((1e10 / -2.25)) == 0.0)) {
                var v11 int = fi8
                var v12 int32 = 0
                _ = v11
                _ = v12
            }
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v13 := []int{42, 1, 32768}
    _ = v13
    _ = v7
    return 503
}
func genFn1(p0 int64, p1 string, p2 int64) int64 {
    for fi14 := 0; fi14 < 12; fi14++ {
        switch p1 {
        case "\n\t":
            p0 = -9223372036854775808
        default:
            var v15 uint8 = 2
            _ = v15
        }
        fmt.Println(p2)
        p2 += (p0 + p0)
        _ = fi14
    }
    switch (p1 + (p1 + p1)) {
    case "hello":
        if true {
            var v16 float64 = (0.0 - (0.1 - (((((((-2.25 * 0.0) * ((0.0 * ((-2.25 * 1.5) + 0.1)) / 1.0)) / 3.14159) + 0.1) - (float64(0.0) + -2.25)) + 1e10) * ((0.1 + (float64(0.1) - float64(1.5))) + -2.25))))
            _ = v16
        } else {
            var v17 uint = 1
            for fi18 := 0; fi18 < 14; fi18++ {
                v19 := []int{fi18, (-(fi18))}
                switch fi18 {
                case -836:
                    fmt.Println(v17, v17)
                case 1:
                    v20 := GenS1{F0: 528, F1: p2}
                    _ = v20
                default:
                    for fi21 := 0; fi21 < 18; fi21++ {
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                        _ = fi21
                    }
                }
                var v22 int32 = 1
                _ = fi18
                _ = v19
                _ = v22
            }
            _ = v17
        }
    case "a":
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 bool = (1 < 2)
            _ = fi23
            _ = v24
        }
    case "\n\t":
        var v25 uint8 = 42
        _ = v25
    }
    p1 = ("hello" + p1)
    p1 = (p1 + "foo bar")
    var v26 uint = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    return p0
}

func main() {
    v1 := []string{"foo bar", "hello", ("\n\t" + "\n\t")}
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    var v4 uint8 = 1
    var v5 int64 = -1
    v6 := []uint32{256, 4294967295, 566, 42}
    v4 = (v4 + v4)
    _ = len(v6)
    v4 = 2
    if true {
        v27 := GenS2{F0: -222}
        fmt.Println(v4, v5)
        _ = v27
    }
    v4 -= (v4 % ((0) | 1))
    if (14 > 42) {
        var v28 int = 42
        var v29 uint16 = 792
        _ = v28
        _ = v29
    }
    if false {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    } else {
        var v30 bool = (-2.25 > (0.1 * float64(-2.25)))
        if v30 {
            delete(v2, ("hello" + ("hello" + "\n\t")))
            v31 := GenS1{F0: 0, F1: v5}
            for fi32 := 0; fi32 < 13; fi32++ {
                v33 := GenS3{F0: true, F1: true, F2: 100}
                if (v30 && (!true)) {
                    if len(v6) > 0 {
                        v6[0] = uint32(v4)
                    }
                    var v34 int64 = -292
                    _ = v34
                } else {
                    fmt.Println(v5, v4)
                    v5 -= -183
                }
                _ = fi32
                _ = v33
            }
            _ = v31
        }
        v30 = v30
        _ = v30
    }
    v4 += v4
    v35 := GenS3{F0: (980 >= 0), F1: true, F2: -1}
    var v36 int32 = 42
    if (true || false) {
        for fi37 := 0; fi37 < 9; fi37++ {
            for fi38 := 0; fi38 < 2; fi38++ {
                v39 := v3[fi37]
                var v40 uint16 = 65535
                _ = fi38
                _ = v39
                _ = v40
            }
            v5 -= v5
            var v41 int64 = v5
            _ = fi37
            _ = v41
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println(v36)
    var v42 uint = uint(v4)
    v43 := GenS1{F0: 81, F1: v5}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
