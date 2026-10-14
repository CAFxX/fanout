package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint; F2 int16 }
type GenS3 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS4 struct { F0 uint32; F1 uint64 }
func genFn0(p0 uint16, p1 int32) int {
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := GenS1{F0: 18446744073709551615}
        for fi7 := 0; fi7 < 15; fi7++ {
            v8 := []int{}
            switch "foo bar" {
            case "a":
                v9 := GenS1{F0: 0}
                _ = v9
            }
            p0 = 42
            _ = fi7
            _ = v8
        }
        p1 = (^(p1))
        _ = fi5
        _ = v6
    }
    switch "" {
    case "hello":
        v10 := GenS1{F0: 484}
        _ = v10
    case "a":
        if ((1000 != 946) && (42 != 2147483647)) {
            p1 += (p1 + -2147483648)
            p1 -= 1
            v11 := GenS2{F0: ((p0 + p0) < 972), F1: 521, F2: -32768}
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        p1 = p1
        _ = fi12
    }
    _ = p0
    _ = p1
    return -52
}
func genFn1(p0 int64) string {
    v13 := GenS3{F0: 42, F1: 42, F2: 42}
    p0 += p0
    _ = p0
    _ = v13
    return "hello"
}

func main() {
    v1 := make(map[int]string)
    var v2 bool = true
    var v3 uint = 65535
    var v4 string = "hello"
    if (!(!v2)) {
        var v14 float32 = 1e10
        v15 := []int{-2147483648, 975, 128}
        _ = v14
        _ = v15
    } else {
        switch (v4 + (v4 + v4)) {
        case "foo bar":
            var v16 int = 2147483647
            _ = v16
        default:
            v2 = (0.0 == (float32(1.5) / ((3.14159 - 1.5) / 1e10)))
        }
    }
    var v17 int8 = 100
    v4 = ("foo bar" + (("\n\t" + (((((((v4 + ("foo bar" + v4)) + "a") + (v4 + "a")) + "a") + v4) + ((("a" + v4) + v4) + v4)) + (v4 + ((("foo bar" + (("hello" + v4) + v4)) + v4) + "")))) + v4))
    v17 += -38
    fmt.Println(v4, v2)
    var v18 float32 = 0.0
    for fi19 := 0; fi19 < 2; fi19++ {
        if false {
            v20 := GenS4{F0: 1, F1: 18446744073709551615}
            if (fi19 <= (fi19 ^ (-(fi19)))) {
                var v21 uint = (^(v3))
                _ = v21
            }
            v17 = v17
            _ = v20
        }
        _ = fi19
    }
    for fi22 := 0; fi22 < 8; fi22++ {
        v18 += float32(float64(float32(((float64((1.5 / (v18 + (((v18 / 3.14159) - 3.14159) / (1.5 / ((v18 / v18) / float32(float64(((v18 + v18) + (((((v18 / v18) / (float32(-2.25) + float32(1.5))) * 1e10) + v18) - (1.5 * ((1e10 / ((1.5 / -2.25) / (float32(3.14159) / 1e10))) * float32(0.1))))))))))))) * 0.1) + 0.1))))
        v23 := GenS4{F0: 4294967295, F1: 999}
        _ = fi22
        _ = v23
    }
    var v24 int16 = 739
    for fi25 := 0; fi25 < 2; fi25++ {
        v18 = 0.1
        _ = fi25
    }
    v4 = ""
    var v26 uint16 = 2
    var v27 uint = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
