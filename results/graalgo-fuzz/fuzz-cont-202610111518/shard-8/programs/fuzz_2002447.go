package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint }
type GenS2 struct { F0 int8; F1 int32; F2 int64 }
func genFn0(p0 int32, p1 uint16) uint32 {
    v7 := []int{4294967295}
    if len(v7) > 0 {
        v7[0] = 42
    }
    if len(v7) > 0 {
        v7[0] = -9223372036854775808
    }
    var v8 int8 = 14
    var v9 int32 = 2147483647
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 822
}
func genFn1(p0 int16, p1 int8, p2 bool) uint64 {
    if (!(!p2)) {
        p1 += 127
    }
    v10 := []int{128}
    var v11 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 1
}

func main() {
    v1 := []uint32{2147483648, 947, 42}
    var v2 int16 = 42
    v3 := []uint32{}
    var v4 uint32 = 65535
    var v5 int32 = -1
    var v6 int32 = (v5 | -1)
    if false {
        _ = len(v3)
    }
    v12 := GenS1{F0: ("foo bar" + "a"), F1: 2, F2: 32}
    var v13 uint = 260
    var v14 float32 = 1.5
    v13 += 0
    _ = len(v3)
    for fi15 := 0; fi15 < 20; fi15++ {
        if false {
            if true {
                for fi16 := 0; fi16 < 8; fi16++ {
                    v13 -= (v13 ^ (v13 % ((v13) | 1)))
                    var v17 uint16 = 32767
                    _ = fi16
                    _ = v17
                }
                for fi18 := 0; fi18 < 4; fi18++ {
                    var v19 string = ("a" + "")
                    _ = fi18
                    _ = v19
                }
                v20 := GenS1{F0: "\n\t", F1: 42, F2: 332}
                _ = v20
            } else {
                v4 = v4
                _ = len(v1)
                v21 := GenS1{F0: "a", F1: 42, F2: v13}
                _ = v21
            }
        } else {
            var v22 uint8 = 42
            for fi23 := 0; fi23 < 2; fi23++ {
                var v24 string = "\n\t"
                var v25 uint64 = 42
                _ = fi23
                _ = v24
                _ = v25
            }
            var v26 float64 = float64(v14)
            _ = v22
            _ = v26
        }
        if (v13 <= v13) {
            switch ("a" + (("foo bar" + "foo bar") + "")) {
            case "a":
                for fi27 := 0; fi27 < 19; fi27++ {
                    var v28 string = "foo bar"
                    v6 -= (v5 - (^(v6)))
                    _ = fi27
                    _ = v28
                }
            case "hello":
                for fi29 := 0; fi29 < 1; fi29++ {
                    var v30 float32 = 3.14159
                    _ = fi29
                    _ = v30
                }
            }
        } else {
            v31 := GenS2{F0: -113, F1: (v5 % (((v6 & (^(v5)))) | 1)), F2: 9223372036854775807}
            var v32 float64 = (0.0 / -2.25)
            var v33 uint64 = genFn1((v2 + v2), 127, true)
            _ = v31
            _ = v32
            _ = v33
        }
        _ = fi15
    }
    var v34 bool = (!((((0.1 * (((1e10 / 3.14159) * (float64((v14 / ((float32(-2.25) / v14) * v14))) - 1.5)) / 1.5)) / 1.5) < (1.5 + ((1.5 / (float64(v14) + 0.0)) * (-2.25 / 3.14159)))) && true))
    v3 = append(v3, v4)
    v5 = v5
    var v35 int16 = (v2 - (v2 - (v2 - (v2 / (((v2 >> (((v13 | v13)) & 63))) | 1)))))
    v14 = v14
    var v36 float64 = 1.5
    v4 = v4
    for fi37 := 0; fi37 < 15; fi37++ {
        fmt.Println(v34, v4)
        _ = fi37
    }
    v13 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
