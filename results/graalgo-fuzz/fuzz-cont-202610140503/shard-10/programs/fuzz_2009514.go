package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS6 struct { F0 uint64 }
func genFn0(p0 bool, p1 uint8, p2 int) string {
    var v7 int32 = -371
    switch "\n\t" {
    case "":
        v8 := GenS1{F0: (p1 ^ p1)}
        _ = v8
    }
    var v9 uint64 = 1
    p0 = (-32768 >= 1000)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v9
    return (("\n\t" + ("" + ("foo bar" + "hello"))) + "hello")
}

func main() {
    var v1 int32 = 42
    var v2 float64 = (0.1 + 0.1)
    v3 := make(map[uint32]bool)
    var v4 float64 = (v2 + v2)
    var v5 int64 = 42
    var v6 int32 = 1
    if true {
        v3[2147483648] = true
    } else {
        v4 = float64(1.5)
        var v10 int16 = 42
        var v11 int8 = -1
        _ = v10
        _ = v11
    }
    fmt.Println(v6)
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := GenS2{F0: 409}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        fmt.Println(v4, v4)
        _ = fi14
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 uint = 65535
        _ = fi15
        _ = v16
    }
    var v17 int16 = 42
    var v18 uint16 = 0
    var v19 int = -2147483648
    v6 += int32(v17)
    for fi20 := 0; fi20 < 13; fi20++ {
        v19 += (v19 - v19)
        fmt.Println(v5)
        for fi21 := 0; fi21 < 15; fi21++ {
            v19 += -1
            _ = fi21
        }
        _ = fi20
    }
    if false {
        fmt.Println(v17, v17)
        var v22 float32 = 3.14159
        v18 = v18
        _ = v22
    } else {
        v23 := []int{}
        if false {
            var v24 int = v19
            fmt.Println(v2, v1)
            _ = v24
        } else {
            if true {
                fmt.Println(v19, v2)
            }
            var v25 float64 = ((1.5 + 0.0) * 0.0)
            v26 := GenS3{F0: 879}
            _ = v25
            _ = v26
        }
        if false {
            v6 += (-(v6))
            v17 -= (v17 ^ -805)
            var v27 float32 = 0.1
            _ = v27
        } else {
            v19 = (v19 * 9223372036854775807)
            v28 := GenS2{F0: 463}
            for fi29 := 0; fi29 < 13; fi29++ {
                v17 = v17
                _ = fi29
            }
            _ = v28
        }
        _ = v23
    }
    switch v19 {
    case 4294967295:
        switch (v19 % ((v19) | 1)) {
        case 9223372036854775807:
            var v30 uint = 1
            _ = v30
        case -1:
            switch "foo bar" {
            case "a":
                v17 -= v17
            case "\n\t":
                v5 += -101
            case "hello":
                if (v18 < (^(v18))) {
                    if (!true) {
                        if true {
                            v31 := GenS4{F0: 0, F1: -1}
                            var v32 float64 = (float64(0.1) * v4)
                            v33 := GenS4{F0: -6, F1: 100}
                            _ = v31
                            _ = v32
                            _ = v33
                        } else {
                            v34 := GenS5{F0: 42, F1: 2, F2: 18446744073709551615}
                            v35 := GenS6{F0: 683}
                            var v36 int64 = v5
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                        v37 := GenS1{F0: 16}
                        _ = v37
                    } else {
                        v2 = ((v4 / v2) + (0.1 + v2))
                        v5 = (v5 % (((v5 % ((v5) | 1))) | 1))
                    }
                } else {
                    var v38 int16 = (v17 / ((-1) | 1))
                    _ = v38
                }
            }
        case -9223372036854775808:
            v2 -= v2
        }
    }
    switch "hello" {
    case "foo bar":
        fmt.Println(v1, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
