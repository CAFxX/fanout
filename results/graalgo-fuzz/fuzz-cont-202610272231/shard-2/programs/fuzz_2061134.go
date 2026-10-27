package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int }
func genFn0(p0 int8) int16 {
    for fi7 := 0; fi7 < 17; fi7++ {
        if false {
            for fi8 := 0; fi8 < 13; fi8++ {
                var v9 int = (^(fi8))
                var v10 int8 = (p0 - (p0 % ((-121) | 1)))
                fmt.Println(fi8)
                _ = fi8
                _ = v10
                _ = v9
            }
            if false {
                var v11 float32 = 1e10
                v11 = 1.5
                _ = v11
            }
        }
        if false {
            v12 := []int{-9223372036854775808, (fi7 / ((32768) | 1))}
            for fi13 := 0; fi13 < 13; fi13++ {
                if (9223372036854775807 != (fi7 << ((531) & 63))) {
                    p0 += (p0 & p0)
                }
                fmt.Println(fi7)
                p0 = 127
                _ = fi13
            }
            _ = v12
        }
        if (!(true && (fi7 >= 859))) {
            fmt.Println(fi7, fi7)
            fmt.Println(p0, fi7)
            v14 := GenS1{F0: (false && (65535 > 1)), F1: 0}
            _ = v14
        }
        _ = fi7
    }
    v15 := GenS2{F0: 4294967295, F1: 1}
    _ = p0
    _ = v15
    return 32767
}
func genFn1() uint16 {
    var v16 uint = 65535
    var v17 uint16 = 213
    for fi18 := 0; fi18 < 9; fi18++ {
        v16 = 18446744073709551615
        fmt.Println(v16, fi18, v17)
        _ = fi18
    }
    _ = v16
    _ = v17
    return 463
}

func main() {
    var v1 uint64 = 520
    v2 := []string{"a"}
    v3 := []uint32{}
    var v4 float64 = 0.1
    var v5 uint8 = 100
    var v6 float32 = 1.5
    for fi19 := 0; fi19 < 20; fi19++ {
        v2 = append(v2, "\n\t")
        v1 += 379
        _ = fi19
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        v4 = ((((v4 / v4) - float64(v6)) * ((v4 + (1.5 * v4)) / (float64((v6 / float32(v4))) - float64(((float32(((float64(v6) - v4) * (v4 * -2.25))) / (float32(-2.25) * 0.1)) - float32(v4)))))) / v4)
        if (127 != (^(v5))) {
            fmt.Println(v4, fi20, v5)
            for fi21 := 0; fi21 < 3; fi21++ {
                if true {
                    var v22 uint32 = 2
                    _ = v22
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi21
            }
            var v23 uint16 = 32767
            _ = v23
        }
        for fi24 := 0; fi24 < 12; fi24++ {
            _ = len(v3)
            var v25 int = fi24
            _ = fi24
            _ = v25
        }
        _ = fi20
    }
    fmt.Println(v5, v6, v1)
    v6 += v6
    var v26 bool = false
    for fi27 := 0; fi27 < 5; fi27++ {
        v5 -= 1
        var v28 bool = v26
        v5 = v5
        _ = fi27
        _ = v28
    }
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    for fi29 := 0; fi29 < 3; fi29++ {
        for fi30 := 0; fi30 < 11; fi30++ {
            switch (fi30 >> ((418) & 63)) {
            case 1:
                var v31 uint = 109
                _ = v31
            case 9223372036854775807:
                v32 := GenS3{F0: fi29}
                _ = v32
            default:
                v33 := GenS3{F0: 2147483647}
                _ = v33
            }
            var v34 float64 = v4
            _ = fi30
            _ = v34
        }
        _ = fi29
    }
    fmt.Println(v4)
    var v35 string = "a"
    v6 = v6
    if v26 {
        v5 += 128
        v36 := GenS3{F0: -2147483648}
        v2 = append(v2, ((v35 + v35) + v35))
        _ = v36
    } else {
        v2 = append(v2, v35)
        var v37 int64 = 1
        for fi38 := 0; fi38 < 11; fi38++ {
            v5 += v5
            _ = fi38
        }
        _ = v37
    }
    if (3.14159 == 1.5) {
        fmt.Println(v5)
    } else {
        _ = len(v2)
        v3 = append(v3, 42)
        for fi39 := 0; fi39 < 18; fi39++ {
            for fi40 := 0; fi40 < 5; fi40++ {
                var v41 bool = (2147483648 > 831)
                _ = fi40
                _ = v41
            }
            _ = fi39
        }
    }
    _ = len(v3)
    if false {
        var v42 bool = v26
        if (v42 && v26) {
            var v43 uint8 = (^(v5))
            var v44 int64 = -9223372036854775808
            var v45 string = ((v35 + v35) + "")
            _ = v43
            _ = v44
            _ = v45
        } else {
            var v46 uint16 = 65535
            if true {
                v6 += float32(v4)
                fmt.Println(v26, v46, v1)
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v46
        }
        _ = len(v3)
        _ = v42
    } else {
        v5 -= 255
    }
    if (-1 > -1) {
        v6 += v6
    } else {
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
        v26 = ((v1 - (v1 % ((4294967295) | 1))) >= (^(v1)))
    }
    v47 := GenS1{F0: (42 < -2147483648), F1: v5}
    switch ("hello" + "hello") {
    case "a":
        switch ("" + v35) {
        case "\n\t":
            fmt.Println(v26)
        }
    case "\n\t":
        for fi48 := 0; fi48 < 12; fi48++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = len(v3)
            switch v35 {
            case "\n\t":
                v26 = v26
            default:
                fmt.Println(v6, v26, v35)
            }
            _ = fi48
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v35
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
