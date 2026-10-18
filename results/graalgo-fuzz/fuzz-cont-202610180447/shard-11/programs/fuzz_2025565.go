package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 int32; F2 uint16 }
func genFn0(p0 int32, p1 uint8, p2 uint64) uint {
    for fi4 := 0; fi4 < 5; fi4++ {
        var v5 uint16 = 32767
        fmt.Println(v5)
        _ = fi4
        _ = v5
    }
    p0 = p0
    if false {
        p1 += (p1 ^ p1)
        var v6 int32 = p0
        _ = v6
    } else {
        p0 = (-(p0))
    }
    _ = p0
    _ = p1
    _ = p2
    return 255
}
func genFn1(p0 int) int {
    switch (p0 << ((1) & 63)) {
    case -9223372036854775808:
        v7 := GenS1{F0: "a"}
        _ = v7
    }
    var v8 int16 = 76
    v9 := GenS1{F0: ""}
    v10 := []int{p0, 42, p0}
    var v11 int16 = 0
    _ = p0
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 127
}

func main() {
    var v1 string = "foo bar"
    v2 := []int{32768, 32768, 128}
    var v3 float32 = (1e10 + float32(3.14159))
    for fi12 := 0; fi12 < 15; fi12++ {
        if false {
            v2 = append(v2, genFn1(genFn1(42)))
            if (true || false) {
                fmt.Println(v1, v3, fi12)
                var v13 int64 = 9223372036854775807
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v13
            } else {
                var v14 string = "a"
                _ = v14
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        var v15 int8 = -126
        _ = fi12
        _ = v15
    }
    if len(v2) > 0 {
        v2[0] = genFn1(127)
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 int16 = 229
        v3 -= (v3 + ((v3 / 1.5) / (v3 * (v3 + 0.0))))
        _ = fi16
        _ = v17
    }
    if ((42 < 1) || (1 < -124)) {
        var v18 int16 = -1
        for fi19 := 0; fi19 < 2; fi19++ {
            v20 := GenS2{F0: 32767, F1: -1, F2: 65535}
            _ = fi19
            _ = v20
        }
        _ = v18
    } else {
        if false {
            v3 = v3
            for fi21 := 0; fi21 < 7; fi21++ {
                var v22 int16 = 344
                _ = fi21
                _ = v22
            }
            if (false && ((v3 == ((0.0 + (v3 - (float32(0.1) * 1e10))) / v3)) && false)) {
                var v23 uint8 = 2
                var v24 float64 = 1e10
                if (((v23 != v23) || (-72 <= 128)) && (-400 != 42)) {
                    var v25 uint32 = 42
                    var v26 uint32 = (v25 | v25)
                    _ = v25
                    _ = v26
                }
                _ = v23
                _ = v24
            }
        }
        if (-1 >= 541) {
            v1 = v1
            var v27 int64 = 42
            _ = v27
        }
        v1 = (v1 + v1)
    }
    for fi28 := 0; fi28 < 2; fi28++ {
        var v29 bool = true
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 6; fi30++ {
        if len(v2) > 0 {
            v2[0] = fi30
        }
        v31 := GenS2{F0: -1, F1: 2147483647, F2: 527}
        v2 = append(v2, fi30)
        _ = fi30
        _ = v31
    }
    switch 128 {
    case 818:
        v2 = append(v2, 128)
    case 127:
        if (!(-9223372036854775808 == 9223372036854775807)) {
            if len(v2) > 0 {
                v2[0] = -1
            }
            var v32 bool = false
            _ = v32
        } else {
            v1 = v1
        }
    case 1:
        fmt.Println(v3, v3)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v33 uint32 = 2147483648
    if len(v2) > 0 {
        v2[0] = 9223372036854775807
    }
    if ((255 <= 296) && false) {
        v34 := GenS2{F0: -160, F1: 0, F2: 587}
        v3 = ((-2.25 * v3) - (v3 - 1.5))
        _ = v34
    } else {
        var v35 int8 = 0
        switch "\n\t" {
        case "foo bar":
            v3 -= float32(0.0)
        case "\n\t":
            fmt.Println(v33, v33, v35)
        case "hello":
            var v36 bool = false
            _ = v36
        }
        if true {
            for fi37 := 0; fi37 < 16; fi37++ {
                for fi38 := 0; fi38 < 10; fi38++ {
                    var v39 float64 = (((-2.25 / (0.0 - (1.5 - (3.14159 - (1e10 * 1e10))))) + ((float64(0.1) * 1.5) * 0.0)) - 0.1)
                    if len(v2) > 0 {
                        v2[0] = -9223372036854775808
                    }
                    _ = fi38
                    _ = v39
                }
                if (!true) {
                    var v40 uint8 = 0
                    _ = v40
                } else {
                    v3 -= v3
                    v33 += 92
                }
                _ = fi37
            }
        }
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
}
