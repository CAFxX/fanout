package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint; F1 string; F2 uint }
type GenS4 struct { F0 bool; F1 int8; F2 uint }
type GenS5 struct { F0 int64; F1 bool; F2 uint16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = -8
    v3 := make(map[string]uint32)
    var v4 int32 = -2147483648
    fmt.Println(v2, v4)
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 int64 = 21
        _ = fi5
        _ = v6
    }
    if (65535 != 0) {
        var v7 int16 = 1000
        _ = v7
    }
    fmt.Println(v2)
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS1{F0: ("" + "\n\t"), F1: true}
        _ = fi8
        _ = v9
    }
    v2 = -32768
    v2 -= v2
    for fi10 := 0; fi10 < 11; fi10++ {
        if (0 != 57) {
            _ = len(v1)
            v11 := v3[("\n\t" + "foo bar")]
            for fi12 := 0; fi12 < 11; fi12++ {
                delete(v3, ("\n\t" + (("a" + ((("foo bar" + "\n\t") + "a") + "a")) + "")))
                v2 -= -1
                v4 += (v4 ^ -118)
                _ = fi12
            }
            _ = v11
        } else {
            var v13 float64 = (3.14159 * 1.5)
            if (0.0 != (0.1 / (-2.25 * (((float32((1.5 * v13)) / (3.14159 * (1e10 - -2.25))) - 0.1) + 3.14159)))) {
                fmt.Println(fi10, v2)
            } else {
                if ((!false) && false) {
                    var v14 uint64 = 866
                    switch ("foo bar" + "\n\t") {
                    case "hello":
                        v15 := GenS2{F0: 1, F1: 255}
                        _ = v15
                    case "\n\t":
                        fmt.Println(v4, v14, v4)
                    case "a":
                        var v16 int16 = 1000
                        _ = v16
                    }
                    v1[2147483648] = "a"
                    _ = v14
                }
                v4 -= 1
                var v17 int16 = (^(v2))
                _ = v17
            }
            _ = v13
        }
        fmt.Println(v2, fi10)
        fmt.Println(v2)
        _ = fi10
    }
    switch -929 {
    case 127:
        for fi18 := 0; fi18 < 4; fi18++ {
            v19 := v3["foo bar"]
            var v20 float32 = (1e10 / ((-2.25 + 1.5) / -2.25))
            for fi21 := 0; fi21 < 6; fi21++ {
                v22 := []int{(^(fi21)), int(v2), fi18, fi21}
                fmt.Println(fi18)
                switch (-(fi21)) {
                case 42:
                    v2 -= (v2 & v2)
                default:
                    if (v20 < (-2.25 / v20)) {
                        fmt.Println(v20)
                        fmt.Println(v2, v2)
                    }
                }
                _ = fi21
                _ = v22
            }
            _ = fi18
            _ = v19
            _ = v20
        }
    case 32768:
        v23 := GenS2{F0: 4294967295, F1: 825}
        _ = v23
    case 1:
        v24 := GenS3{F0: 138, F1: "", F2: 88}
        _ = v24
    }
    v25 := GenS4{F0: (9223372036854775807 != 727), F1: -114, F2: 1}
    v26 := GenS5{F0: int64(v2), F1: true, F2: 376}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
