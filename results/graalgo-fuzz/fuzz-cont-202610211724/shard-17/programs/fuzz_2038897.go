package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int32; F1 int64; F2 int32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = 32767
    var v3 bool = false
    var v4 string = "hello"
    var v5 int = -2147483648
    v4 = v4
    var v6 int8 = 0
    v6 = v6
    fmt.Println(v2, v2, v5)
    if ((float32(1.5) * (3.14159 - 3.14159)) <= -2.25) {
        var v7 uint = 713
        _ = len(v1)
        switch v5 {
        case -1:
            var v8 uint16 = 2
            _ = v8
        default:
            for fi9 := 0; fi9 < 5; fi9++ {
                v6 -= (v6 << ((18446744073709551615) & 63))
                _ = fi9
            }
        }
        _ = v7
    } else {
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 float32 = (3.14159 - (0.0 - 3.14159))
            fmt.Println(fi10)
            _ = fi10
            _ = v11
        }
        v2 = 1000
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 int64 = -384
        _ = fi12
        _ = v13
    }
    var v14 uint16 = 32767
    v15 := GenS1{F0: 0, F1: v4}
    fmt.Println(v2, v4)
    var v16 int8 = 100
    if false {
        var v17 float32 = (1.5 / (((-2.25 / -2.25) + 0.1) / -2.25))
        fmt.Println(v5, v2)
        switch (v4 + v4) {
        case "foo bar":
            fmt.Println(v5, v5)
        case "a":
            if (3.14159 < (1.5 / ((-2.25 - (float64(1.5) - 3.14159)) + 0.1))) {
                for fi18 := 0; fi18 < 18; fi18++ {
                    for fi19 := 0; fi19 < 5; fi19++ {
                        v1[fi18] = uint32(v14)
                        _ = fi19
                    }
                    v3 = (100 <= v6)
                    var v20 uint16 = v14
                    _ = fi18
                    _ = v20
                }
                v21 := GenS1{F0: 1, F1: ("a" + "hello")}
                v2 += 1000
                _ = v21
            }
        default:
            fmt.Println(v14, v3, v16)
        }
        _ = v17
    } else {
        delete(v1, 0)
        var v22 float32 = 0.1
        switch "foo bar" {
        case "hello":
            var v23 uint = 652
            _ = v23
        case "":
            v24 := v1[42]
            _ = v24
        case "foo bar":
            switch (v5 / ((-9223372036854775808) | 1)) {
            case 2147483647:
                _ = len(v1)
            case 42:
                v6 += v16
            }
        default:
            if (false && true) {
                v22 -= float32(-2.25)
                v25 := GenS2{F0: (v4 + v4), F1: v4}
                _ = v25
            } else {
                if v3 {
                    for fi26 := 0; fi26 < 1; fi26++ {
                        fmt.Println(v2, v22, fi26)
                        v2 -= (v2 + (v2 | (v2 + v2)))
                        _ = fi26
                    }
                    if true {
                        v27 := GenS3{F0: 0, F1: -219, F2: 2147483647}
                        v14 -= v14
                        _ = v27
                    } else {
                        var v28 float32 = 0.0
                        _ = v28
                    }
                } else {
                    v6 -= v16
                    fmt.Println(v3)
                    v29 := []int{4294967295, v5, 128}
                    _ = v29
                }
                v3 = v3
                var v30 uint16 = v14
                _ = v30
            }
        }
        _ = v22
    }
    switch (v5 - (v5 % (((v5 | (v5 + 0))) | 1))) {
    case 810:
        v31 := GenS1{F0: 1, F1: (("\n\t" + v4) + "")}
        _ = v31
    case 9223372036854775807:
        if (0 > 149) {
            v32 := []int{-2147483648}
            _ = v32
        } else {
            fmt.Println(v2, v4)
        }
    default:
        fmt.Println(v3, v5, v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
