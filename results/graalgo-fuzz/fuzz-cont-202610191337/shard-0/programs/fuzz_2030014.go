package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int64 }
type GenS2 struct { F0 int32; F1 int8; F2 int32 }
func genFn0() int64 {
    if (0.1 < 1.5) {
        var v4 int16 = -32768
        _ = v4
    } else {
        var v5 int8 = 96
        _ = v5
    }
    for fi6 := 0; fi6 < 13; fi6++ {
        fmt.Println(fi6)
        var v7 int = -2147483648
        _ = fi6
        _ = v7
    }
    if (!true) {
        var v8 uint16 = 819
        var v9 uint8 = 255
        _ = v8
        _ = v9
    }
    var v10 float64 = 0.0
    if (2 <= 803) {
        v11 := []int{739, -72, 4294967295}
        for fi12 := 0; fi12 < 17; fi12++ {
            var v13 int32 = -2147483648
            v11 = append(v11, (^(fi12)))
            v14 := GenS1{F0: ((("" + "") + "hello") + "hello"), F1: 9223372036854775807, F2: -9223372036854775808}
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 int16 = 1
        _ = v11
        _ = v15
    }
    _ = v10
    return -1
}

func main() {
    var v1 uint = 255
    var v2 uint8 = 1
    var v3 int32 = 2147483647
    switch ("hello" + "a") {
    case "foo bar":
        var v16 int64 = 42
        _ = v16
    }
    var v17 float64 = ((0.0 / 1e10) * float64(((0.0 + 1.5) / (1.5 * ((((0.1 * -2.25) + (-2.25 * 3.14159)) + 1e10) - -2.25)))))
    v1 = v1
    var v18 uint64 = 4294967295
    v17 = v17
    v3 -= v3
    v18 += (v18 | 348)
    v1 = (v1 - v1)
    switch "foo bar" {
    case "hello":
        fmt.Println(v2, v1, v1)
    default:
        for fi19 := 0; fi19 < 15; fi19++ {
            var v20 float32 = (3.14159 / 0.1)
            v1 -= v1
            v21 := []int{}
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    fmt.Println(v18, v2)
    v22 := []int{42}
    switch 0 {
    case 128:
        fmt.Println(v17)
    case 895:
        if (4294967295 <= v18) {
            var v23 int8 = -111
            for fi24 := 0; fi24 < 11; fi24++ {
                _ = len(v22)
                for fi25 := 0; fi25 < 3; fi25++ {
                    v23 = v23
                    v2 = 42
                    switch "\n\t" {
                    case "\n\t":
                        v2 -= 1
                        var v26 uint16 = 2
                        _ = len(v22)
                        if len(v22) > 0 {
                            v22[0] = (-(fi25))
                        }
                        _ = v26
                    case "hello":
                        v3 = v3
                    case "foo bar":
                        fmt.Println(v17, fi25)
                    default:
                        v22 = append(v22, fi25)
                    }
                    _ = fi25
                }
                _ = fi24
            }
            _ = v23
        }
    case 0:
        for fi27 := 0; fi27 < 3; fi27++ {
            switch "foo bar" {
            case "hello":
                v28 := GenS1{F0: "a", F1: 42, F2: 9223372036854775807}
                _ = v28
            default:
                for fi29 := 0; fi29 < 19; fi29++ {
                    if true {
                        v30 := GenS2{F0: -1, F1: 127, F2: v3}
                        _ = v30
                    } else {
                        v18 += v18
                        switch (fi29 ^ fi27) {
                        case 32767:
                            v22 = append(v22, -162)
                        case -127:
                            fmt.Println(fi29, fi27)
                        }
                    }
                    var v31 uint8 = v2
                    _ = fi29
                    _ = v31
                }
            }
            v2 = v2
            _ = fi27
        }
    }
    v18 = (^(v18))
    for fi32 := 0; fi32 < 2; fi32++ {
        v3 -= 583
        var v33 uint64 = 322
        var v34 int64 = 988
        _ = fi32
        _ = v33
        _ = v34
    }
    switch -9223372036854775808 {
    case 2147483647:
        if (!(false && (255 == 1))) {
            for fi35 := 0; fi35 < 15; fi35++ {
                var v36 int64 = genFn0()
                _ = fi35
                _ = v36
            }
            var v37 int64 = 9223372036854775807
            _ = v37
        }
    case 426:
        fmt.Println(v17)
    case -1:
        v22 = append(v22, 2147483647)
    }
    if len(v22) > 0 {
        v22[0] = -2147483648
    }
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
}
