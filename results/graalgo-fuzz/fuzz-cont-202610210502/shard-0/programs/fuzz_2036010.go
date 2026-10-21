package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
func genFn0() uint64 {
    v4 := []int{-214, -9223372036854775808, 257}
    if len(v4) > 0 {
        v4[0] = -9223372036854775808
    }
    _ = v4
    return 1
}
func genFn1(p0 uint8, p1 uint32) int8 {
    var v5 uint = 2
    var v6 uint32 = p1
    p1 = v6
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -43
}

func main() {
    var v1 int8 = 32
    v2 := []uint32{0, 1, 2, 42}
    v3 := []uint32{}
    if true {
        if (0 >= 587) {
            for fi7 := 0; fi7 < 17; fi7++ {
                v8 := GenS1{F0: (4294967295 < 4294967295), F1: (("foo bar" + ("\n\t" + "hello")) + "hello")}
                v1 = genFn1(0, 434)
                var v9 uint8 = 255
                _ = fi7
                _ = v8
                _ = v9
            }
            var v10 uint = 367
            _ = v10
        } else {
            v1 -= 100
            _ = len(v2)
        }
    } else {
        fmt.Println(v1)
        v3 = append(v3, 1)
        v1 = (v1 & v1)
    }
    switch "a" {
    case "foo bar":
        for fi11 := 0; fi11 < 15; fi11++ {
            for fi12 := 0; fi12 < 5; fi12++ {
                if len(v3) > 0 {
                    v3[0] = 1
                }
                _ = fi12
            }
            var v13 float64 = 3.14159
            fmt.Println(v13)
            _ = fi11
            _ = v13
        }
    case "a":
        fmt.Println(v1)
    default:
        v1 = 0
    }
    var v14 int64 = 9223372036854775807
    fmt.Println(v1, v1)
    if (703 > 2147483647) {
        if (v14 >= 1) {
            fmt.Println(v14)
        }
        if (true && ((-2.25 / (float64(3.14159) + 3.14159)) == float64(1.5))) {
            v3 = append(v3, 474)
            fmt.Println(v14)
        }
        if len(v3) > 0 {
            v3[0] = 1
        }
    } else {
        for fi15 := 0; fi15 < 12; fi15++ {
            switch (fi15 % ((fi15) | 1)) {
            case -9223372036854775808:
                var v16 int16 = 42
                _ = v16
            case 317:
                for fi17 := 0; fi17 < 12; fi17++ {
                    if len(v3) > 0 {
                        v3[0] = 907
                    }
                    _ = fi17
                }
            }
            fmt.Println(v1)
            _ = len(v3)
            _ = fi15
        }
        for fi18 := 0; fi18 < 15; fi18++ {
            var v19 float32 = float32(-2.25)
            var v20 uint8 = 128
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        for fi22 := 0; fi22 < 6; fi22++ {
            var v23 float64 = 0.0
            switch "a" {
            case "hello":
                v23 = 1e10
            case "a":
                v1 -= 42
            case "\n\t":
                if true {
                    v3 = append(v3, 42)
                }
            }
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    var v24 float64 = -2.25
    if len(v2) > 0 {
        v2[0] = 945
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v3
}
