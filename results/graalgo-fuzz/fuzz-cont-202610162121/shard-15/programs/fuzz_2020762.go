package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint16; F1 uint; F2 int16 }

func main() {
    v1 := []string{((("foo bar" + (("" + "hello") + "hello")) + "a") + "a")}
    v2 := make(map[int]bool)
    v3 := make(map[int]string)
    v4 := []string{"", "foo bar"}
    var v5 int = 32767
    if true {
        var v6 uint = 259
        for fi7 := 0; fi7 < 8; fi7++ {
            v6 -= v6
            var v8 int64 = 1
            _ = fi7
            _ = v8
        }
        var v9 int8 = 30
        _ = v6
        _ = v9
    }
    if ((0 != -1) && false) {
        var v10 int32 = -2147483648
        _ = v10
    } else {
        var v11 int32 = 1
        fmt.Println(v11, v5)
        v11 = v11
        _ = v11
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        v13 := GenS1{F0: 42, F1: 127, F2: 78}
        for fi14 := 0; fi14 < 11; fi14++ {
            fmt.Println(fi12, fi14)
            v5 -= fi14
            _ = fi14
        }
        if (0 != 2) {
            if false {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if false {
                    fmt.Println(fi12)
                }
                if false {
                    var v15 int8 = 23
                    v5 -= (fi12 & v5)
                    _ = v15
                }
            }
            v5 = -2147483648
            var v16 uint32 = 65535
            _ = v16
        }
        _ = fi12
        _ = v13
    }
    var v17 int = v5
    var v18 int64 = 0
    v4 = append(v4, "\n\t")
    v17 = 824
    for fi19 := 0; fi19 < 2; fi19++ {
        for fi20 := 0; fi20 < 3; fi20++ {
            switch ("foo bar" + "") {
            case "\n\t":
                v5 = v17
            case "foo bar":
                fmt.Println(v18, fi19, v5)
            case "":
                var v21 uint64 = 18446744073709551615
                _ = v21
            }
            if true {
                v1 = append(v1, ("\n\t" + "hello"))
            }
            _ = fi20
        }
        v17 -= (v5 / ((fi19) | 1))
        _ = fi19
    }
    v22 := v2[127]
    v17 -= 32767
    if ((v5 | v5) != 4294967295) {
        var v23 int8 = 1
        switch "hello" {
        case "a":
            v5 += (v5 | (v5 % (((v17 ^ v5)) | 1)))
        default:
            var v24 uint16 = 0
            _ = v24
        }
        _ = v23
    }
    var v25 int32 = 748
    v26 := GenS1{F0: 18446744073709551615, F1: 0, F2: 0}
    var v27 uint = 18446744073709551615
    fmt.Println(v27)
    for fi28 := 0; fi28 < 20; fi28++ {
        switch "hello" {
        case "\n\t":
            if v22 {
                fmt.Println(v27, v27, v5)
                v22 = (int8(fi28) >= -38)
                v29 := GenS1{F0: v27, F1: 255, F2: 127}
                _ = v29
            }
        case "hello":
            for fi30 := 0; fi30 < 16; fi30++ {
                var v31 float32 = float32(((1e10 / 0.1) - 0.1))
                _ = fi30
                _ = v31
            }
        default:
            for fi32 := 0; fi32 < 8; fi32++ {
                if (0.0 == 1e10) {
                    v18 += v18
                    var v33 int64 = v18
                    v34 := GenS1{F0: (v27 & (v27 << (((^(v27))) & 63))), F1: 2, F2: 1}
                    _ = v33
                    _ = v34
                } else {
                    switch "hello" {
                    case "hello":
                        var v35 float64 = ((-2.25 / float64(float32(((1.5 * -2.25) + 1e10)))) + 1.5)
                        _ = v35
                    }
                }
                fmt.Println(v17, v17)
                if v22 {
                    switch fi32 {
                    case 1:
                        v36 := GenS1{F0: 18446744073709551615, F1: 0, F2: int8(v27)}
                        _ = v36
                    case -2147483648:
                        v37 := GenS1{F0: 65535, F1: 127, F2: -1}
                        _ = v37
                    }
                    v25 += 470
                } else {
                    _ = len(v3)
                    v38 := v2[v17]
                    fmt.Println(v22, v25)
                    _ = v38
                }
                v39 := GenS2{F0: 0, F1: 65535, F2: 1}
                var v40 uint64 = 1
                var v41 bool = (1.5 < 3.14159)
                _ = fi32
                _ = v39
                _ = v40
                _ = v41
            }
        }
        _ = fi28
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
