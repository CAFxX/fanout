package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 bool }

func main() {
    var v1 uint16 = 1
    v2 := make(map[string]string)
    var v3 bool = (true || (442 >= 2147483648))
    v4 := make(map[int]bool)
    var v5 bool = v3
    v6 := make(map[int]bool)
    v5 = v3
    var v7 int = 128
    if (!true) {
        _ = len(v4)
        if v5 {
            v7 -= v7
            var v8 string = "hello"
            var v9 int8 = 100
            _ = v8
            _ = v9
        } else {
            v10 := []int{(v7 + (v7 | (v7 & (v7 * (v7 ^ 128))))), v7, 32768, v7}
            var v11 uint32 = 2147483647
            var v12 int16 = -11
            _ = v10
            _ = v11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 float64 = 0.1
        v15 := GenS1{F0: 65535, F1: 649}
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        fmt.Println(fi16, fi16)
        _ = fi16
    }
    v3 = v5
    v17 := GenS1{F0: 949, F1: 4294967295}
    for fi18 := 0; fi18 < 17; fi18++ {
        v6[-2147483648] = (v7 <= fi18)
        switch "hello" {
        case "\n\t":
            var v19 float32 = 0.1
            _ = v19
        case "":
            if true {
                v3 = v5
            } else {
                v20 := []int{(fi18 / ((v7) | 1)), (-(v7))}
                _ = v20
            }
        default:
            v4[-2147483648] = ((!false) && v5)
        }
        v21 := GenS1{F0: v1, F1: 255}
        _ = fi18
        _ = v21
    }
    var v22 int = v7
    if (255 <= 2) {
        var v23 int16 = 964
        switch (("a" + ("hello" + ("foo bar" + "\n\t"))) + "\n\t") {
        case "\n\t":
            v24 := GenS2{F0: (!(42 <= 42))}
            _ = v24
        case "a":
            if (v3 || (545 <= 18446744073709551615)) {
                for fi25 := 0; fi25 < 17; fi25++ {
                    for fi26 := 0; fi26 < 1; fi26++ {
                        var v27 uint32 = 482
                        var v28 uint64 = 0
                        fmt.Println(v23, fi26)
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                    _ = len(v2)
                    _ = fi25
                }
                v29 := GenS1{F0: 1, F1: 42}
                if v3 {
                    var v30 float32 = ((1e10 - 3.14159) * (float32(0.0) - float32(-2.25)))
                    var v31 int32 = 89
                    _ = v30
                    _ = v31
                }
                _ = v29
            }
        case "":
            var v32 float64 = ((0.0 * 1.5) / 1.0)
            _ = v32
        }
        if true {
            var v33 uint8 = 54
            v33 += v33
            _ = v33
        } else {
            v3 = false
        }
        _ = v23
    } else {
        fmt.Println(v7)
    }
    switch (v22 * (v7 & v7)) {
    case 4294967295:
        var v34 uint16 = (v1 & v1)
        _ = v34
    case -539:
        v35 := GenS2{F0: (0 == -28)}
        _ = v35
    default:
        v36 := GenS2{F0: (v1 <= (v1 * (v1 % ((65535) | 1))))}
        _ = v36
    }
    v5 = (1 < (v7 & -9223372036854775808))
    var v37 bool = v5
    var v38 float32 = 3.14159
    fmt.Println(v22)
    for fi39 := 0; fi39 < 8; fi39++ {
        for fi40 := 0; fi40 < 20; fi40++ {
            switch "" {
            case "a":
                if (-290 < -934) {
                    v7 += v7
                } else {
                    v41 := []int{fi40, -9223372036854775808, v7, (fi39 % ((v22) | 1))}
                    _ = v41
                }
            }
            fmt.Println(v1, v38)
            _ = fi40
        }
        v42 := []int{(fi39 >> ((155) & 63)), v22, fi39}
        v42 = append(v42, v7)
        _ = fi39
        _ = v42
    }
    var v43 float32 = (3.14159 / v38)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
}
