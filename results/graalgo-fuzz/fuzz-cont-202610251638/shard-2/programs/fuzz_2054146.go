package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint32 }
func genFn0() bool {
    var v5 bool = false
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: 42, F1: -830, F2: 4294967295}
        if (v5 && ((((!v5) && (465 <= -9223372036854775808)) || false) && v5)) {
            if v5 {
                if (1 > 2147483647) {
                    var v8 uint16 = 42
                    _ = v8
                }
            }
        }
        v5 = (1 != 42)
        _ = fi6
        _ = v7
    }
    _ = v5
    return (!(1.5 < -2.25))
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = -2.25
    v3 := make(map[string]uint32)
    var v4 float32 = v2
    var v9 float64 = 0.1
    for fi10 := 0; fi10 < 7; fi10++ {
        fmt.Println(v4, v2, fi10)
        switch 1 {
        case -251:
            var v11 float64 = float64((3.14159 / (v2 - 3.14159)))
            _ = v11
        case 127:
            switch "hello" {
            case "\n\t":
                if true {
                    switch (fi10 ^ fi10) {
                    case 4294967295:
                        v12 := []int{-9223372036854775808}
                        _ = v12
                    case 127:
                        v2 += 3.14159
                    case 272:
                        v13 := GenS1{F0: 32767, F1: 1000, F2: 2}
                        _ = v13
                    }
                } else {
                    for fi14 := 0; fi14 < 18; fi14++ {
                        v2 = v2
                        switch (fi14 & fi10) {
                        case 32768:
                            v2 += ((1.5 * (v4 + v4)) / 3.14159)
                        }
                        if false {
                            if (-914 >= -9223372036854775808) {
                                fmt.Println(v2, v2, fi14)
                            }
                        }
                        _ = fi14
                    }
                    for fi15 := 0; fi15 < 5; fi15++ {
                        v9 += (3.14159 * 3.14159)
                        var v16 bool = false
                        v17 := v1[fi10]
                        var v18 uint64 = 916
                        v19 := []int{(fi10 & 1), fi15}
                        v9 -= 0.0
                        _ = fi15
                        _ = v16
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    v9 += (1e10 / v9)
                }
            case "foo bar":
                v20 := GenS1{F0: 1, F1: 1, F2: 2}
                _ = v20
            default:
                fmt.Println(v9, fi10, v2)
            }
        }
        v21 := []int{(fi10 * fi10), fi10, fi10}
        _ = fi10
        _ = v21
    }
    if (18446744073709551615 >= 2) {
        var v22 bool = (true || ("\n\t" >= ("a" + ("hello" + "a"))))
        _ = v22
    }
    fmt.Println(v4)
    if ((32767 < 32768) || (1.5 >= v2)) {
        var v23 uint = 1
        _ = v23
    }
    if true {
        if (338 < 255) {
            v24 := []int{2147483647}
            switch 950 {
            case 42:
                var v25 float64 = v9
                _ = v25
            case 1:
                if len(v24) > 0 {
                    v24[0] = 0
                }
            case 32767:
                for fi26 := 0; fi26 < 18; fi26++ {
                    fmt.Println(v2)
                    var v27 uint16 = 32767
                    v4 += -2.25
                    _ = fi26
                    _ = v27
                }
            }
            _ = v24
        }
        v28 := GenS1{F0: 42, F1: 0, F2: 2}
        var v29 uint16 = 32767
        _ = v28
        _ = v29
    } else {
        var v30 uint32 = 0
        _ = v30
    }
    v4 = (v4 - (float32(1.5) - ((1.5 * (float32(0.0) / (v2 + 3.14159))) + (((0.1 / -2.25) / v2) * -2.25))))
    for fi31 := 0; fi31 < 4; fi31++ {
        v32 := []int{(fi31 | fi31), fi31, fi31}
        if len(v32) >= 2 {
            v32 = v32[1:len(v32)]
        }
        v9 = float64(v2)
        _ = fi31
        _ = v32
    }
    v9 = v9
    _ = len(v3)
    switch -2147483648 {
    case 32767:
        v2 = float32(v9)
    case 128:
        for fi33 := 0; fi33 < 1; fi33++ {
            for fi34 := 0; fi34 < 17; fi34++ {
                var v35 uint64 = 2
                switch 359 {
                case 4294967295:
                    if false {
                        v35 = v35
                        if ((fi34 / ((fi33) | 1)) < fi33) {
                            var v36 uint8 = 128
                            v1[(fi34 | 571)] = 42
                            _ = v36
                        } else {
                            var v37 float64 = (v9 * 3.14159)
                            fmt.Println(fi34, fi33, v9)
                            _ = v37
                        }
                    } else {
                        v35 = (^(v35))
                    }
                default:
                    var v38 float64 = (v9 + v9)
                    _ = v38
                }
                _ = fi34
                _ = v35
            }
            delete(v3, "\n\t")
            _ = fi33
        }
    }
    var v39 bool = (7 < 0)
    var v40 float64 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v9
}
