package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int32; F1 int64; F2 int16 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint32, p1 int, p2 string) int16 {
    fmt.Println(p0, p2, p0)
    for fi7 := 0; fi7 < 2; fi7++ {
        fmt.Println(p2, p2)
        _ = fi7
    }
    if true {
        for fi8 := 0; fi8 < 20; fi8++ {
            p0 = p0
            p2 = ("" + p2)
            _ = fi8
        }
        var v9 uint = 42
        _ = v9
    } else {
        v10 := GenS1{F0: 127}
        v11 := []int{}
        fmt.Println(p2, p1, p0)
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: 1}
    var v13 int64 = 376
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return -934
}

func main() {
    var v1 uint16 = 255
    v2 := []string{"hello", "a", "hello"}
    var v3 bool = (((0.1 * 3.14159) >= (3.14159 / -2.25)) || (-1 <= -759))
    var v4 bool = (-1 > 42)
    v5 := make(map[uint32]string)
    var v6 int = 9223372036854775807
    if (genFn0(925, (v6 & v6), ("" + "hello")) < 837) {
        v3 = v4
        v1 = (v1 ^ 670)
        v6 = v6
    } else {
        v14 := v5[151]
        _ = v14
    }
    v1 = v1
    v4 = (1.5 >= (0.1 * 1e10))
    for fi15 := 0; fi15 < 17; fi15++ {
        v16 := GenS1{F0: 128}
        var v17 float32 = 0.0
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 int8 = 42
    var v19 uint8 = 127
    fmt.Println(v1, v18, v6)
    v6 -= v6
    var v20 uint32 = 4294967295
    v18 = (v18 - int8(v19))
    for fi21 := 0; fi21 < 14; fi21++ {
        for fi22 := 0; fi22 < 10; fi22++ {
            v6 += -1
            fmt.Println(v19, v18)
            var v23 float64 = 1.5
            _ = fi22
            _ = v23
        }
        if v3 {
            switch (fi21 << ((1) & 63)) {
            case 246:
                v5[v20] = "\n\t"
            default:
                if (!(true && (-553 == -1))) {
                    switch "" {
                    case "":
                        v3 = (347 > 1000)
                    case "\n\t":
                        fmt.Println(v18)
                    }
                    var v24 uint64 = 42
                    v25 := GenS1{F0: v19}
                    _ = v24
                    _ = v25
                }
            }
        }
        _ = fi21
    }
    if (v20 >= 434) {
        var v26 uint8 = (v19 | v19)
        for fi27 := 0; fi27 < 4; fi27++ {
            for fi28 := 0; fi28 < 11; fi28++ {
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
                delete(v5, v20)
                _ = fi28
            }
            var v29 int64 = 42
            _ = fi27
            _ = v29
        }
        _ = v26
    } else {
        for fi30 := 0; fi30 < 17; fi30++ {
            switch "foo bar" {
            case "a":
                var v31 uint32 = 0
                _ = v31
            }
            _ = fi30
        }
    }
    switch v6 {
    case 127:
        v20 = 2
    default:
        v3 = v4
    }
    if (true && false) {
        switch (v6 * v6) {
        case -594:
            v6 = (v6 % ((v6) | 1))
        default:
            var v32 int32 = -2147483648
            _ = v32
        }
    } else {
        switch "\n\t" {
        case "a":
            if ((^(v1)) > 0) {
                fmt.Println(v6, v20, v6)
                fmt.Println(v1)
                v5[(v20 % ((419) | 1))] = (("\n\t" + "a") + ((("" + "a") + ("\n\t" + ("foo bar" + "\n\t"))) + (("a" + ("a" + "hello")) + "hello")))
            }
        case "hello":
            for fi33 := 0; fi33 < 12; fi33++ {
                v34 := GenS2{F0: v20, F1: (-(v18))}
                var v35 int16 = 1000
                _ = fi33
                _ = v34
                _ = v35
            }
        case "\n\t":
            if (v6 < (v6 << ((4294967295) & 63))) {
                v5[(v20 >> ((17) & 63))] = "hello"
                var v36 int16 = 32767
                v37 := GenS3{F0: 1, F1: 942, F2: v36}
                _ = v36
                _ = v37
            }
        default:
            v38 := GenS4{F0: (^(v18))}
            _ = v38
        }
        if v3 {
            if ((float64(-2.25) * 0.1) == 3.14159) {
                if v4 {
                    var v39 string = "a"
                    var v40 bool = v4
                    v41 := GenS1{F0: (v19 * (^(v19)))}
                    _ = v39
                    _ = v40
                    _ = v41
                } else {
                    var v42 float64 = ((0.0 / float64(1e10)) + 1e10)
                    v43 := GenS4{F0: 1}
                    _ = v42
                    _ = v43
                }
            } else {
                var v44 float32 = ((0.1 + 1.5) + 1.5)
                v1 -= (v1 / ((v1) | 1))
                _ = v44
            }
        }
        v20 = v20
    }
    v45 := GenS2{F0: 65535, F1: (v18 & v18)}
    for fi46 := 0; fi46 < 5; fi46++ {
        if v4 {
            var v47 bool = v4
            v18 = v18
            _ = v47
        }
        _ = fi46
    }
    var v48 uint = 255
    for fi49 := 0; fi49 < 8; fi49++ {
        for fi50 := 0; fi50 < 7; fi50++ {
            fmt.Println(fi50, v4, fi49)
            _ = fi50
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v51 uint64 = 446
        _ = fi49
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v45
    _ = v48
    _ = v5
    _ = v6
}
