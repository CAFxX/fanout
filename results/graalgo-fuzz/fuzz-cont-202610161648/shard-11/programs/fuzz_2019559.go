package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 string; F2 uint64 }
func genFn0(p0 string) bool {
    v7 := []int{4294967295, -344}
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 uint32 = 912
        var v10 int = (fi8 >> ((42) & 63))
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint = 740
    _ = p0
    _ = v11
    _ = v7
    return (false || ((2 >= 42) && (2147483647 == 730)))
}
func genFn1(p0 uint16, p1 int64) int32 {
    if (((!false) && (!true)) && true) {
        for fi12 := 0; fi12 < 5; fi12++ {
            fmt.Println(p0, p0, p1)
            _ = fi12
        }
    }
    var v13 uint8 = 226
    if (!(p1 != p1)) {
        var v14 uint64 = 639
        _ = v14
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v13
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]bool)
    var v3 float32 = (1e10 - 0.1)
    var v4 uint32 = 661
    var v5 int16 = -32768
    var v6 int8 = 42
    v6 -= v6
    switch 32768 {
    case 128:
        fmt.Println(v4)
    default:
        v15 := GenS1{F0: false}
        _ = v15
    }
    fmt.Println(v5, v4, v4)
    var v16 int64 = 1
    v4 += 2147483648
    var v17 string = ("foo bar" + "a")
    if (!true) {
        var v18 uint32 = v4
        v19 := GenS1{F0: (((3.14159 * 0.0) / 3.14159) <= float64(((float32(1.5) + (float32(1e10) * v3)) + 1.5)))}
        _ = v18
        _ = v19
    } else {
        v3 += 1.5
        v20 := GenS1{F0: (false || true)}
        if ((((v3 - v3) - v3) / v3) <= (float32(0.0) - v3)) {
            if false {
                switch 128 {
                case 42:
                    v16 = 9223372036854775807
                case -1:
                    fmt.Println(v16)
                }
                if false {
                    if (false || false) {
                        var v21 int32 = -2147483648
                        _ = v21
                    } else {
                        var v22 float64 = (1e10 / (((((-2.25 + ((0.0 + (3.14159 - 1e10)) + (float64(v3) + (0.0 * 0.1)))) - (0.1 + ((-2.25 - 3.14159) * 1e10))) / (1e10 - 3.14159)) - 0.0) * float64((float32(-2.25) * v3))))
                        _ = v22
                    }
                } else {
                    var v23 uint32 = (v4 / ((672) | 1))
                    switch -766 {
                    case 819:
                        if (true && true) {
                            v24 := GenS1{F0: (398 > (v16 + v16))}
                            for fi25 := 0; fi25 < 13; fi25++ {
                                if false {
                                    var v26 float32 = -2.25
                                    var v27 int64 = (-(v16))
                                    v3 = (v3 * (v3 / float32(3.14159)))
                                    _ = v26
                                    _ = v27
                                } else {
                                    v17 = (((v17 + "foo bar") + ("hello" + (v17 + ("" + "hello")))) + v17)
                                    v23 = (v4 % ((v4) | 1))
                                }
                                var v28 uint32 = (v4 >> ((357) & 63))
                                var v29 int = -1
                                _ = fi25
                                _ = v28
                                _ = v29
                            }
                            v1[(v17 + "")] = (v4 | 518)
                            _ = v24
                        }
                    case 128:
                        switch (v17 + v17) {
                        case "hello":
                            var v30 int8 = v6
                            _ = v30
                        }
                    }
                    _ = v23
                }
            }
            v31 := GenS1{F0: false}
            v5 = (v5 / (((v5 / (((v5 ^ v5)) | 1))) | 1))
            _ = v31
        }
        _ = v20
    }
    var v32 string = v17
    var v33 uint8 = 0
    v2[int(v6)] = true
    v3 = ((0.0 / v3) / v3)
    v33 += (v33 + v33)
    fmt.Println(v32, v16)
    v34 := GenS2{F0: 2147483647, F1: (v32 + "a"), F2: 4294967295}
    v32 = ("a" + (v17 + "hello"))
    v4 = v4
    for fi35 := 0; fi35 < 6; fi35++ {
        fmt.Println(fi35, v16)
        _ = fi35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
