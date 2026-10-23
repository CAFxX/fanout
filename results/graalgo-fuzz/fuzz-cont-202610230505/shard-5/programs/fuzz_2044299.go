package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64 }

func main() {
    var v1 float64 = ((((1e10 / 1.0) / 0.1) - (float64((1.5 / 1e10)) / 3.14159)) / 1e10)
    var v2 uint = 42
    v3 := make(map[int]string)
    var v4 int8 = -101
    if true {
        for fi5 := 0; fi5 < 12; fi5++ {
            if (32768 > fi5) {
                v4 = v4
                v4 += (v4 - (v4 >> ((431) & 63)))
            } else {
                fmt.Println(fi5, v2, v1)
            }
            _ = fi5
        }
        if false {
            v4 -= v4
        } else {
            fmt.Println(v1)
            v6 := GenS1{F0: 2147483648}
            switch 407 {
            case 630:
                var v7 bool = true
                _ = v7
            }
            _ = v6
        }
    }
    v8 := []int{32768, 32768, 32767, 2147483647}
    v9 := GenS2{F0: -486}
    switch 4294967295 {
    case -410:
        if (true || true) {
            if (!(((((v1 / -2.25) + float64((1e10 + ((-2.25 * -2.25) * float32(float64(1e10)))))) * (v1 + 3.14159)) < 0.0) && (255 <= 42))) {
                v4 += 42
            }
            v10 := GenS2{F0: 214}
            _ = v10
        }
    case 332:
        if (2 > 79) {
            var v11 float64 = v1
            for fi12 := 0; fi12 < 10; fi12++ {
                var v13 int64 = -915
                v4 = -128
                v4 += v4
                _ = fi12
                _ = v13
            }
            if true {
                if len(v8) > 0 {
                    v8[0] = 128
                }
                for fi14 := 0; fi14 < 2; fi14++ {
                    _ = len(v8)
                    _ = fi14
                }
            } else {
                v4 = (v4 - (v4 & (v4 / ((v4) | 1))))
            }
            _ = v11
        }
    }
    if (v4 < v4) {
        var v15 uint8 = 109
        for fi16 := 0; fi16 < 9; fi16++ {
            fmt.Println(v15)
            v4 = -128
            _ = fi16
        }
        _ = v15
    }
    v17 := GenS1{F0: 362}
    switch "foo bar" {
    case "a":
        v2 = v2
    case "":
        _ = len(v3)
    }
    _ = len(v8)
    v4 -= v4
    switch (("hello" + ("hello" + (("\n\t" + ("" + ("foo bar" + "hello"))) + "hello"))) + "hello") {
    case "hello":
        var v18 int8 = v4
        _ = v18
    default:
        switch -1 {
        case 659:
            for fi19 := 0; fi19 < 17; fi19++ {
                v3[fi19] = ("foo bar" + "")
                _ = fi19
            }
        case 127:
            v20 := GenS2{F0: -574}
            _ = v20
        case 2147483647:
            var v21 uint32 = 509
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 12; fi22++ {
        v23 := GenS3{F0: (114 != 1)}
        _ = fi22
        _ = v23
    }
    v8 = append(v8, -2147483648)
    if true {
        v1 += v1
        v24 := GenS2{F0: 32767}
        var v25 uint64 = 1
        _ = v24
        _ = v25
    } else {
        var v26 int16 = 0
        _ = v26
    }
    v4 = v4
    v2 += (v2 << (((^(v2))) & 63))
    switch "\n\t" {
    case "a":
        v27 := GenS2{F0: 4294967295}
        _ = v27
    case "foo bar":
        v28 := GenS4{F0: 11}
        _ = v28
    default:
        if (!(true || false)) {
            v29 := v3[-2147483648]
            _ = v29
        } else {
            var v30 int16 = -32768
            _ = v30
        }
    }
    delete(v3, -699)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
