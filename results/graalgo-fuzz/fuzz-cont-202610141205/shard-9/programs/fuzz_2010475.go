package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int64; F1 int }
func genFn0(p0 uint) uint32 {
    var v7 uint8 = 127
    v8 := []int{128}
    var v9 int16 = 1
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 2147483647
}
func genFn1(p0 uint32, p1 uint32) uint8 {
    v10 := []int{32767}
    fmt.Println(p1, p1)
    for fi11 := 0; fi11 < 8; fi11++ {
        p1 += p0
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        var v12 uint = 65535
        _ = fi11
        _ = v12
    }
    if (float32(3.14159) != 0.0) {
        p0 = p1
        var v13 uint = 1
        var v14 uint16 = 723
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v10
    return 128
}

func main() {
    var v1 float64 = float64((float32(float64(float32(float64((0.0 + (float32((3.14159 / ((float64(1e10) * (0.1 - (0.0 - 0.1))) / 0.1))) * 0.0)))))) + 1.5))
    var v2 uint8 = 127
    v3 := make(map[string]bool)
    v4 := []int{4294967295}
    var v5 int64 = -9223372036854775808
    var v6 string = "a"
    if false {
        v1 = 1e10
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        var v15 bool = (!(-545 != 501))
        _ = v15
    }
    v4 = append(v4, -2147483648)
    v3[(v6 + "hello")] = false
    for fi16 := 0; fi16 < 3; fi16++ {
        delete(v3, "\n\t")
        v1 -= 1e10
        _ = fi16
    }
    fmt.Println(v5)
    fmt.Println(v1)
    v6 = (v6 + (v6 + v6))
    var v17 int = 1
    v1 = 1e10
    v1 = (3.14159 / (((v1 + (float64(float32(3.14159)) * v1)) + float64((0.1 + 3.14159))) * 1.5))
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 float32 = ((1e10 - 0.0) * 1.5)
        if (v17 > 2147483647) {
            v20 := GenS1{F0: -1, F1: 2147483647}
            for fi21 := 0; fi21 < 10; fi21++ {
                v19 += v19
                var v22 string = "hello"
                _ = fi21
                _ = v22
            }
            _ = v20
        } else {
            switch (v17 - v17) {
            case 908:
                var v23 bool = (2 == 42)
                _ = v23
            }
            v3[(v6 + v6)] = (false || (!(1000 < -543)))
            var v24 uint16 = 0
            _ = v24
        }
        v25 := GenS2{F0: (v5 ^ v5), F1: -764}
        _ = fi18
        _ = v19
        _ = v25
    }
    if true {
        if (1.5 != 1.5) {
            fmt.Println(v6, v17)
            v3["\n\t"] = (v17 == (v17 * 32768))
        }
    } else {
        switch v17 {
        case -102:
            delete(v3, (v6 + "a"))
        case -1:
            for fi26 := 0; fi26 < 2; fi26++ {
                for fi27 := 0; fi27 < 7; fi27++ {
                    v28 := GenS2{F0: (v5 & v5), F1: (fi26 - (-(v17)))}
                    v4 = append(v4, (fi26 / ((42) | 1)))
                    v6 = "foo bar"
                    _ = fi27
                    _ = v28
                }
                switch "foo bar" {
                case "foo bar":
                    var v29 float32 = 1.5
                    _ = v29
                case "hello":
                    v30 := v3["a"]
                    _ = v30
                case "\n\t":
                    v6 = ("foo bar" + "a")
                default:
                    v6 = v6
                }
                for fi31 := 0; fi31 < 6; fi31++ {
                    fmt.Println(v5, v1)
                    if false {
                        var v32 int8 = -128
                        var v33 uint = 255
                        var v34 int = 32768
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                    _ = fi31
                }
                _ = fi26
            }
        case 127:
            switch 42 {
            case 32767:
                if false {
                    v4 = append(v4, v17)
                    _ = len(v3)
                } else {
                    if false {
                        v17 -= v17
                    } else {
                        v5 = 806
                        fmt.Println(v5, v17, v17)
                    }
                }
            case -9223372036854775808:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case 127:
                fmt.Println(v1, v5, v2)
            default:
                v5 = v5
            }
        }
    }
    var v35 int16 = 32767
    var v36 uint64 = 4294967295
    if (18446744073709551615 < 65535) {
        v35 = (v35 % ((v35) | 1))
        v3["hello"] = true
        if false {
            for fi37 := 0; fi37 < 9; fi37++ {
                v6 = "foo bar"
                var v38 bool = (-128 < int8(v5))
                v36 += 224
                _ = fi37
                _ = v38
            }
        }
    }
    for fi39 := 0; fi39 < 4; fi39++ {
        v36 += v36
        v35 += 42
        _ = fi39
    }
    delete(v3, "")
    v40 := GenS1{F0: 0, F1: -515}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
