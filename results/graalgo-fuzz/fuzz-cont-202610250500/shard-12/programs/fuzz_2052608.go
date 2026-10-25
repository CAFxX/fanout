package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint; F2 uint8 }

func main() {
    v1 := []string{"a", "", "\n\t", (("hello" + "hello") + "")}
    var v2 int32 = 1
    v3 := make(map[uint32]string)
    v4 := make(map[string]uint32)
    var v5 int = 9223372036854775807
    if (2 == 870) {
        if len(v1) > 0 {
            v1[0] = ""
        }
        var v6 int16 = -32768
        _ = v6
    } else {
        if (false && (0.0 <= (0.0 + 3.14159))) {
            v4[("foo bar" + ("foo bar" + "foo bar"))] = 1
            delete(v4, "foo bar")
        } else {
            v5 = (v5 % ((v5) | 1))
            var v7 float64 = ((float64((1e10 / (0.0 + (3.14159 * ((0.0 / 1.5) + ((((float32((float64((-2.25 - ((float32(0.1) / 0.1) + (((-2.25 - 3.14159) * ((0.1 + 1e10) * 0.1)) + 1e10)))) - ((1e10 + 1.5) / 0.1))) - (1e10 + 3.14159)) - -2.25) - 3.14159) / (1e10 + 1.5))))))) / (0.0 + 1.5)) * float64(0.0))
            _ = v7
        }
        var v8 float32 = 3.14159
        _ = v8
    }
    v9 := GenS1{F0: ((3.14159 * 1.5) == -2.25), F1: (false || true)}
    if true {
        v10 := GenS1{F0: (!(!(!true))), F1: (12 <= 1)}
        _ = v10
    } else {
        v3[2147483647] = "a"
    }
    var v11 float64 = 3.14159
    if true {
        fmt.Println(v11, v11, v11)
        v3[65535] = ""
        switch v5 {
        case -347:
            for fi12 := 0; fi12 < 10; fi12++ {
                var v13 int16 = 137
                var v14 uint16 = 255
                switch "\n\t" {
                case "\n\t":
                    var v15 int16 = (v13 % ((v13) | 1))
                    _ = v15
                case "hello":
                    for fi16 := 0; fi16 < 2; fi16++ {
                        v14 = (v14 - 0)
                        v11 += ((1e10 + (v11 + v11)) - 0.0)
                        var v17 int16 = v13
                        _ = fi16
                        _ = v17
                    }
                }
                _ = fi12
                _ = v13
                _ = v14
            }
        case 4294967295:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 1:
            delete(v4, ("hello" + ""))
        }
    } else {
        switch v5 {
        case 4294967295:
            var v18 int64 = 42
            _ = v18
        case -98:
            var v19 int8 = 1
            _ = v19
        case 128:
            v11 -= v11
        }
        for fi20 := 0; fi20 < 15; fi20++ {
            v21 := GenS1{F0: (32767 > (v5 << ((1) & 63))), F1: ((18446744073709551615 != 558) || true)}
            var v22 bool = false
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    fmt.Println(v11)
    for fi23 := 0; fi23 < 9; fi23++ {
        v5 += (fi23 | (fi23 ^ (v5 * v5)))
        v24 := GenS2{F0: (^(fi23))}
        for fi25 := 0; fi25 < 16; fi25++ {
            var v26 float64 = v11
            _ = fi25
            _ = v26
        }
        _ = fi23
        _ = v24
    }
    var v27 uint64 = 4294967295
    v28 := GenS3{F0: 42, F1: 0, F2: 209}
    v29 := GenS2{F0: 42}
    v27 = v27
    v11 += (1.5 / v11)
    var v30 int64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v9
}
