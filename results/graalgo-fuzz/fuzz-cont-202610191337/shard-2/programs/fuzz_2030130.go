package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int64, p1 string) int16 {
    v7 := []int{0, 128, 127}
    v7 = append(v7, 127)
    p0 -= p0
    _ = p0
    _ = p1
    _ = v7
    return 42
}
func genFn1() uint32 {
    v8 := []int{360}
    var v9 bool = (!(-211 != -869))
    if (v9 && (!((0.1 / -2.25) > -2.25))) {
        v9 = ((float32(0.1) <= ((float32(0.0) + -2.25) + ((-2.25 + 1.5) / -2.25))) && (false || (!v9)))
        fmt.Println(v9)
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 uint32 = 2147483647
            for fi12 := 0; fi12 < 6; fi12++ {
                if true {
                    v11 += (v11 >> ((1) & 63))
                    v13 := GenS1{F0: 4294967295, F1: 647}
                    _ = v13
                } else {
                    if v9 {
                        switch fi10 {
                        case -2147483648:
                            var v14 string = "foo bar"
                            _ = v14
                        case 2147483647:
                            var v15 int32 = -846
                            _ = v15
                        default:
                            v9 = false
                        }
                        fmt.Println(fi10, v9)
                        v8 = append(v8, (fi12 / (((fi12 >> ((65535) & 63))) | 1)))
                    }
                    if v9 {
                        var v16 uint64 = 404
                        v16 -= (v16 & (v16 & 1))
                        v9 = v9
                        if len(v8) > 0 {
                            v8[0] = (^(fi12))
                        }
                        _ = v16
                    } else {
                        var v17 uint = 137
                        v8 = append(v8, -77)
                        v11 = v11
                        _ = v17
                    }
                    fmt.Println(v11)
                }
                _ = fi12
            }
            _ = fi10
            _ = v11
        }
    }
    if len(v8) > 0 {
        v8[0] = 1
    }
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    var v1 uint64 = 40
    v2 := make(map[int]bool)
    var v3 float64 = 0.1
    var v4 uint64 = (v1 | (v1 - 216))
    var v5 float32 = (0.0 + 3.14159)
    v6 := []string{("\n\t" + ("foo bar" + "\n\t")), ("a" + "")}
    switch ("hello" + "foo bar") {
    case "":
        v18 := GenS1{F0: 353, F1: -9223372036854775808}
        _ = v18
    default:
        v3 = (float64((v5 * float32(0.1))) + v3)
    }
    fmt.Println(v5, v4, v3)
    switch 4294967295 {
    case 437:
        for fi19 := 0; fi19 < 13; fi19++ {
            v1 = v4
            _ = fi19
        }
    case 1:
        v4 -= v4
    case 283:
        var v20 float32 = v5
        _ = v20
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        if true {
            _ = len(v6)
            var v22 string = "foo bar"
            _ = v22
        } else {
            switch "hello" {
            case "foo bar":
                var v23 int64 = 1
                _ = v23
            case "a":
                v4 += v4
            }
            v5 = v5
            var v24 int16 = 280
            _ = v24
        }
        var v25 int64 = 640
        fmt.Println(v3, v4, v1)
        _ = fi21
        _ = v25
    }
    var v26 float32 = float32(v3)
    for fi27 := 0; fi27 < 5; fi27++ {
        switch (^(fi27)) {
        case 311:
            for fi28 := 0; fi28 < 15; fi28++ {
                if len(v6) > 0 {
                    v6[0] = ("\n\t" + "a")
                }
                v3 = v3
                v1 -= (v1 % (((v4 + (v4 * v4))) | 1))
                _ = fi28
            }
        case 32768:
            _ = len(v2)
        }
        var v29 int32 = -346
        _ = fi27
        _ = v29
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        _ = len(v2)
        var v31 int8 = 108
        v6 = append(v6, "foo bar")
        _ = fi30
        _ = v31
    }
    v32 := GenS2{F0: 2147483647}
    fmt.Println(v5)
    v26 += float32(float64(v26))
    var v33 int32 = -1
    fmt.Println(v5, v3, v4)
    v34 := GenS2{F0: 2147483647}
    v26 += 0.1
    v35 := v2[1]
    var v36 float64 = (-2.25 + ((0.0 + v3) * -2.25))
    var v37 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
