package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 int8 = 0
    var v2 float64 = (((1e10 / 3.14159) - 1.5) + ((1.5 / (0.1 - 1.5)) - 1.5))
    v3 := []string{"a", "foo bar", "", ("" + "foo bar")}
    var v4 int16 = 981
    v5 := make(map[string]bool)
    switch 0 {
    case 564:
        var v6 uint16 = 32767
        _ = v6
    case 4294967295:
        for fi7 := 0; fi7 < 4; fi7++ {
            v8 := GenS1{F0: 65535}
            for fi9 := 0; fi9 < 16; fi9++ {
                var v10 float64 = v2
                v3 = append(v3, "")
                v11 := GenS1{F0: 42}
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi7
            _ = v8
        }
    default:
        v12 := GenS2{F0: 2, F1: 666, F2: 42}
        _ = v12
    }
    v2 += v2
    var v13 float64 = v2
    v2 += (((0.0 - 1.5) * (1.5 + (1e10 - v13))) / 3.14159)
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 int8 = v1
        if false {
            if ((((("a" + ("foo bar" + (("hello" + "\n\t") + "a"))) + "") >= "foo bar") && ((fi14 + (^(fi14))) <= (fi14 / (((fi14 * (fi14 ^ (-(fi14))))) | 1)))) && true) {
                var v16 int = fi14
                if false {
                    var v17 uint = 642
                    for fi18 := 0; fi18 < 5; fi18++ {
                        v19 := GenS2{F0: 0, F1: 418, F2: 94}
                        if true {
                            v1 += (v15 << (((v17 ^ v17)) & 63))
                        }
                        var v20 float32 = (-2.25 - 0.0)
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    if ("hello" <= "a") {
                        if false {
                            v13 -= -2.25
                            var v21 int64 = -1
                            _ = v21
                        }
                        var v22 uint64 = 1
                        _ = v22
                    }
                    _ = v17
                } else {
                    var v23 uint = 569
                    v16 = (-(v16))
                    fmt.Println(v15, v23)
                    _ = v23
                }
                _ = v16
            }
            var v24 uint = 164
            _ = v24
        }
        _ = fi14
        _ = v15
    }
    if (true || (65535 == 586)) {
        var v25 int8 = v1
        v26 := GenS2{F0: 251, F1: 2, F2: 20}
        _ = v25
        _ = v26
    } else {
        switch 0 {
        case 2147483647:
            var v27 bool = false
            _ = v27
        default:
            switch -9223372036854775808 {
            case 9223372036854775807:
                fmt.Println(v4, v4, v1)
            case -2147483648:
                var v28 int16 = (v4 ^ 32767)
                _ = v28
            }
        }
    }
    fmt.Println(v2, v2)
    for fi29 := 0; fi29 < 12; fi29++ {
        v2 = ((float64(float32((v2 + float64(float32((v13 + v13)))))) * v13) + (1e10 + float64(3.14159)))
        v4 = (v4 - v4)
        _ = fi29
    }
    v2 = v13
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi30 := 0; fi30 < 12; fi30++ {
        for fi31 := 0; fi31 < 7; fi31++ {
            delete(v5, "a")
            var v32 int = (fi30 % (((fi30 % (((fi30 + fi30)) | 1))) | 1))
            var v33 bool = ((!(3.14159 < ((-2.25 / (1.5 / 1.5)) + (1e10 / 1e10)))) || (1e10 != 0.1))
            _ = fi31
            _ = v32
            _ = v33
        }
        v2 -= 0.1
        _ = fi30
    }
    v13 -= v13
    var v34 float64 = -2.25
    v4 = 1
    for fi35 := 0; fi35 < 11; fi35++ {
        fmt.Println(v4)
        _ = fi35
    }
    v34 = (1.5 - ((((v13 * v34) - v34) - 0.1) - (-2.25 * v13)))
    v36 := GenS3{F0: -9223372036854775808}
    fmt.Println(v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v34
    _ = v36
    _ = v4
    _ = v5
}
