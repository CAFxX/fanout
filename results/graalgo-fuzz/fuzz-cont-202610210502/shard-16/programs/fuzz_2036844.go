package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 bool; F1 int32; F2 int8 }
func genFn0(p0 int16) uint32 {
    v7 := []int{0, -124, -9223372036854775808}
    var v8 bool = true
    p0 -= (^(p0))
    switch 128 {
    case 0:
        v8 = (!(("foo bar" + "") <= (("foo bar" + "") + "\n\t")))
    default:
        v9 := GenS1{F0: 2}
        _ = v9
    }
    v7 = append(v7, -1)
    _ = p0
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 int32, p1 int) string {
    v10 := []int{p1}
    var v11 int = (p1 - 32768)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return "\n\t"
}

func main() {
    var v1 int64 = 42
    var v2 string = "foo bar"
    var v3 uint64 = 1
    v4 := make(map[int]bool)
    var v5 string = ("" + "foo bar")
    v6 := make(map[string]int)
    var v12 uint32 = 2
    for fi13 := 0; fi13 < 2; fi13++ {
        for fi14 := 0; fi14 < 18; fi14++ {
            if (!((42 < 145) || true)) {
                v15 := []int{}
                v16 := GenS1{F0: 504}
                v17 := GenS1{F0: 316}
                v2 = ""
                _ = v15
                _ = v16
                _ = v17
            }
            if false {
                var v18 uint8 = 0
                delete(v6, v5)
                fmt.Println(fi14)
                _ = v18
            }
            switch "foo bar" {
            case "a":
                if true {
                    v19 := GenS2{F0: 154, F1: 204}
                    v5 = v2
                    delete(v4, (fi13 / (((fi13 << ((1) & 63))) | 1)))
                    v20 := []int{(fi14 << ((18446744073709551615) & 63)), fi14}
                    fmt.Println(v3)
                    _ = v19
                    _ = v20
                } else {
                    v2 = v2
                }
            case "hello":
                switch v5 {
                case "":
                    switch "hello" {
                    case "hello":
                        v3 -= v3
                    case "a":
                        v21 := GenS2{F0: 2, F1: (v12 - 65535)}
                        _ = v21
                    case "":
                        var v22 uint8 = 128
                        _ = v22
                    default:
                        v1 = v1
                    }
                }
            case "\n\t":
                var v23 int64 = v1
                _ = v23
            default:
                switch v2 {
                case "foo bar":
                    if (false && true) {
                        var v24 float32 = ((float32(float64(3.14159)) * (0.0 + 3.14159)) - float32(-2.25))
                        if ((0.1 * (float32(3.14159) - v24)) != float32(float64(v24))) {
                            v5 = ""
                            fmt.Println(v12, v5, v3)
                            v25 := []int{(fi13 | 0)}
                            _ = v25
                        }
                        v26 := []int{(fi14 * -90)}
                        _ = v24
                        _ = v26
                    } else {
                        v3 = (^(v3))
                        v27 := []int{}
                        v3 = (^(v3))
                        _ = v27
                    }
                case "a":
                    var v28 string = "hello"
                    _ = v28
                default:
                    fmt.Println(v5, v1)
                }
            }
            v5 = "a"
            _ = fi14
        }
        var v29 int64 = v1
        _ = fi13
        _ = v29
    }
    var v30 uint64 = v3
    for fi31 := 0; fi31 < 4; fi31++ {
        fmt.Println(v2, v1, v3)
        _ = fi31
    }
    if ((^(v1)) == -9223372036854775808) {
        v32 := GenS1{F0: 255}
        v12 += 1
        if (int8(v12) >= 1) {
            var v33 int8 = -94
            v34 := []int{}
            _ = v33
            _ = v34
        }
        _ = v32
    }
    v35 := GenS3{F0: true, F1: 663, F2: 100}
    var v36 uint16 = 1
    fmt.Println(v2, v2, v5)
    if true {
        fmt.Println(v36)
    }
    var v37 float32 = float32(1.5)
    v2 = "\n\t"
    var v38 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
