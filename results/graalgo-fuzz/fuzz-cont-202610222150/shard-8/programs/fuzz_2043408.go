package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int16) int64 {
    p0 = (p0 >> ((948) & 63))
    var v4 bool = false
    fmt.Println(v4)
    if true {
        if ((663 >= 42) || (42 != 0)) {
            v5 := GenS1{F0: v4, F1: 0, F2: 51}
            for fi6 := 0; fi6 < 3; fi6++ {
                v7 := GenS1{F0: (!v4), F1: 9, F2: -11}
                _ = fi6
                _ = v7
            }
            switch ("\n\t" + ("a" + "foo bar")) {
            case "\n\t":
                for fi8 := 0; fi8 < 12; fi8++ {
                    var v9 uint8 = 0
                    switch (fi8 % (((fi8 % (((fi8 & 149)) | 1))) | 1)) {
                    case -762:
                        var v10 int64 = 1
                        _ = v10
                    }
                    _ = fi8
                    _ = v9
                }
            }
            _ = v5
        } else {
            var v11 bool = (!v4)
            _ = v11
        }
    } else {
        if (0.1 <= -2.25) {
            var v12 int32 = -1
            _ = v12
        }
    }
    v13 := []int{-671, -224, 9223372036854775807, 544}
    _ = p0
    _ = v13
    _ = v4
    return 31
}
func genFn1(p0 int16, p1 uint8) uint8 {
    var v14 int = -757
    p1 = p1
    var v15 uint = 42
    v16 := GenS2{F0: 2, F1: 42}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v16
    return p1
}

func main() {
    var v1 uint16 = 2
    var v2 bool = ("\n\t" != "\n\t")
    v3 := []int{-2147483648}
    if (255 > 0) {
        var v17 uint64 = 0
        var v18 float64 = ((1e10 - (3.14159 / 0.1)) + (((-2.25 + (0.0 / (((1.5 + ((0.0 / 1.5) - (((3.14159 + 1.5) - float64(1.5)) + (3.14159 / 1.0)))) + 1.5) / 1e10))) - 1.5) * 1e10))
        _ = v17
        _ = v18
    } else {
        var v19 int64 = 0
        switch "" {
        case "foo bar":
            v20 := GenS2{F0: 1, F1: 42}
            _ = v20
        case "\n\t":
            var v21 float32 = float32(-2.25)
            _ = v21
        case "a":
            v3 = append(v3, -9223372036854775808)
        default:
            v2 = v2
        }
        v22 := GenS1{F0: v2, F1: 128, F2: -128}
        _ = v19
        _ = v22
    }
    if (3.14159 == float32(-2.25)) {
        switch 32768 {
        case 32767:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        if v2 {
            fmt.Println(v2, v1)
            var v23 uint = 0
            _ = v23
        } else {
            v1 -= (v1 >> ((702) & 63))
        }
        v2 = (1 > 256)
    }
    switch ("\n\t" + "\n\t") {
    case "hello":
        var v24 int8 = -91
        _ = v24
    }
    if ((float64(float32(-2.25)) / 3.14159) > (1.5 - 1.5)) {
        v1 += (^(v1))
        var v25 uint64 = 893
        _ = v25
    }
    switch "" {
    case "\n\t":
        v1 -= (v1 & v1)
    default:
        var v26 bool = v2
        _ = v26
    }
    fmt.Println(v2, v1)
    var v27 int16 = 1000
    var v28 uint = 1
    var v29 int16 = 32767
    v3 = append(v3, 127)
    for fi30 := 0; fi30 < 20; fi30++ {
        if (v2 && (v2 || false)) {
            v3 = append(v3, 0)
        }
        _ = fi30
    }
    fmt.Println(v29)
    if (("a" + "foo bar") >= "\n\t") {
        for fi31 := 0; fi31 < 18; fi31++ {
            if v2 {
                if (!v2) {
                    _ = len(v3)
                    if len(v3) > 0 {
                        v3[0] = fi31
                    }
                    v32 := GenS3{F0: v27}
                    _ = v32
                } else {
                    if v2 {
                        var v33 uint = (v28 ^ v28)
                        var v34 bool = (v2 || ((v28 + (v28 << ((v33) & 63))) > v33))
                        _ = v33
                        _ = v34
                    } else {
                        v35 := GenS2{F0: (v1 >> ((18446744073709551615) & 63)), F1: 0}
                        _ = v35
                    }
                    fmt.Println(v29, v1)
                }
                var v36 int = fi31
                _ = v36
            } else {
                var v37 uint = (^(v28))
                _ = v37
            }
            _ = fi31
        }
    } else {
        v29 += v27
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    for fi38 := 0; fi38 < 10; fi38++ {
        var v39 float32 = 1e10
        _ = fi38
        _ = v39
    }
    if (32768 <= 127) {
        _ = len(v3)
        v28 = 338
        fmt.Println(v27, v28)
    }
    v2 = v2
    if (v2 || (4294967295 != v28)) {
        var v40 uint8 = 29
        for fi41 := 0; fi41 < 20; fi41++ {
            v42 := GenS2{F0: v1, F1: 42}
            _ = fi41
            _ = v42
        }
        _ = v40
    }
    v27 = (^(v29))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
