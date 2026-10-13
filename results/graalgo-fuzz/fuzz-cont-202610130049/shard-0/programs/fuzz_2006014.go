package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int8 }
func genFn0(p0 uint64, p1 uint8, p2 int64) int {
    if (0 == 2147483647) {
        p1 = 127
    }
    p0 -= p0
    var v4 int16 = -1
    v5 := []int{0, 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 0
}
func genFn1() bool {
    var v6 uint = 2
    switch 662 {
    case 127:
        for fi7 := 0; fi7 < 8; fi7++ {
            v8 := []int{}
            _ = fi7
            _ = v8
        }
    case 259:
        var v9 bool = (-32768 >= -1)
        _ = v9
    case 4294967295:
        for fi10 := 0; fi10 < 3; fi10++ {
            fmt.Println(v6, fi10)
            _ = fi10
        }
    }
    if (!true) {
        v6 = v6
        v11 := []int{-9223372036854775808}
        if true {
            v6 += (v6 | 130)
            var v12 float64 = 1.5
            _ = v12
        }
        _ = v11
    } else {
        v6 = v6
    }
    v13 := []int{0, 143, 127, 32767}
    _ = v13
    _ = v6
    return false
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint16 = 351
    v3 := []uint32{2, 2147483647, 635}
    fmt.Println(v2, v2)
    var v14 float64 = ((-2.25 / (0.1 / 1.5)) + float64(1e10))
    if true {
        var v15 int16 = 0
        switch ("hello" + ("" + "hello")) {
        case "":
            var v16 uint16 = 65535
            _ = v16
        default:
            fmt.Println(v15)
        }
        _ = len(v3)
        _ = v15
    }
    var v17 int = 470
    var v18 float32 = (0.1 - 0.0)
    var v19 uint8 = 155
    v2 -= v2
    var v20 int16 = int16(v1)
    var v21 uint64 = 1
    var v22 int64 = 1
    if false {
        v23 := GenS1{F0: 18446744073709551615, F1: uint8(v21), F2: -1}
        for fi24 := 0; fi24 < 9; fi24++ {
            switch (fi24 | (v17 | 20)) {
            case 4294967295:
                if (v20 > 859) {
                    fmt.Println(v1, v18, v1)
                    v1 -= (-(v1))
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            case 32767:
                switch "" {
                case "":
                    var v25 uint32 = 254
                    _ = v25
                default:
                    v26 := GenS1{F0: (v21 & 1), F1: v19, F2: 42}
                    _ = v26
                }
            default:
                _ = len(v3)
            }
            if ((v21 << ((255) & 63)) >= v21) {
                var v27 string = "hello"
                _ = len(v3)
                v28 := GenS1{F0: (v21 * (v21 << ((49) & 63))), F1: v19, F2: 127}
                _ = v27
                _ = v28
            }
            _ = fi24
        }
        _ = v23
    } else {
        v20 += -819
    }
    var v29 int16 = v20
    v30 := GenS1{F0: v21, F1: (v19 & v19), F2: 42}
    for fi31 := 0; fi31 < 2; fi31++ {
        var v32 uint64 = v21
        var v33 int16 = v20
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 uint8 = uint8(v2)
    if (!true) {
        if (v21 == 18446744073709551615) {
            var v35 uint8 = (v19 << ((18446744073709551615) & 63))
            _ = v35
        }
        v17 = v17
        if (true || false) {
            switch 0 {
            case 32767:
                switch v17 {
                case -9223372036854775808:
                    v36 := GenS1{F0: (v21 << ((0) & 63)), F1: 42, F2: -108}
                    _ = v36
                case 9223372036854775807:
                    var v37 uint64 = 114
                    _ = v37
                }
            case 0:
                var v38 uint16 = v2
                _ = v38
            }
        } else {
            var v39 int16 = (v20 / ((v29) | 1))
            var v40 uint32 = 42
            _ = v39
            _ = v40
        }
    } else {
        fmt.Println(v2, v22, v21)
    }
    if false {
        var v41 bool = true
        fmt.Println(v22)
        _ = v41
    }
    var v42 int64 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v42
}
