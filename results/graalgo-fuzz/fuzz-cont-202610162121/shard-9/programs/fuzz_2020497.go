package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 uint; F1 uint8; F2 uint }
func genFn0(p0 int8, p1 string, p2 int16) int32 {
    if true {
        var v5 int16 = -1
        p1 = ((p1 + (p1 + p1)) + (p1 + p1))
        _ = v5
    }
    var v6 int64 = 0
    var v7 int8 = 127
    switch 9223372036854775807 {
    case -2147483648:
        var v8 int32 = 358
        _ = v8
    default:
        var v9 bool = (-736 != 2147483647)
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}
func genFn1() uint64 {
    if false {
        v10 := GenS1{F0: 2147483647}
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 bool = (0 <= 42)
            v12 = ((fi11 << ((742) & 63)) == (^(fi11)))
            switch ("hello" + (("" + "foo bar") + ("a" + ""))) {
            case "\n\t":
                v12 = (-2147483648 > -2147483648)
            case "hello":
                v12 = ((((0.0 / 1.5) + 1e10) / float32((1.5 * (0.0 - 3.14159)))) == -2.25)
            case "foo bar":
                var v13 uint32 = 0
                _ = v13
            default:
                for fi14 := 0; fi14 < 20; fi14++ {
                    v15 := []int{361, fi14, fi11}
                    if (!("" < ("foo bar" + "foo bar"))) {
                        v16 := GenS1{F0: -71}
                        v17 := GenS1{F0: -213}
                        var v18 uint = 1
                        _ = v16
                        _ = v17
                        _ = v18
                    } else {
                        v12 = ((1.5 / float32((1.5 / 0.1))) <= 0.0)
                    }
                    fmt.Println(fi14, fi11)
                    _ = fi14
                    _ = v15
                }
            }
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    var v19 float32 = float32(0.0)
    var v20 int16 = 0
    v20 -= -32768
    if false {
        v19 -= 1.5
        v19 = v19
    } else {
        v21 := GenS1{F0: 632}
        switch (("\n\t" + "hello") + "") {
        case "hello":
            for fi22 := 0; fi22 < 5; fi22++ {
                v20 = v20
                v23 := []int{(-(fi22)), 127}
                for fi24 := 0; fi24 < 9; fi24++ {
                    var v25 bool = ((fi24 ^ 42) > 127)
                    var v26 uint = 4294967295
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                _ = fi22
                _ = v23
            }
        default:
            v27 := []int{}
            _ = v27
        }
        fmt.Println(v20)
        _ = v21
    }
    _ = v19
    _ = v20
    return 4294967295
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{128}
    v3 := make(map[int]uint32)
    var v4 bool = (!(false && false))
    var v28 bool = v4
    switch 32768 {
    case 32767:
        for fi29 := 0; fi29 < 5; fi29++ {
            v30 := GenS2{F0: 42, F1: ((("hello" + "a") + "foo bar") + "\n\t")}
            v31 := GenS3{F0: 108, F1: -1}
            var v32 uint16 = 2
            _ = fi29
            _ = v30
            _ = v31
            _ = v32
        }
    }
    fmt.Println(v28, v4)
    v28 = v4
    v28 = v4
    v28 = (9223372036854775807 <= 1)
    var v33 float64 = ((1e10 - 3.14159) * 0.0)
    _ = len(v1)
    v34 := GenS4{F0: 46, F1: 207, F2: 0}
    var v35 int = 9223372036854775807
    var v36 int = (v35 ^ v35)
    if false {
        if v4 {
            var v37 string = "\n\t"
            _ = v37
        } else {
            v28 = v4
            v38 := GenS2{F0: 429, F1: "\n\t"}
            _ = v38
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
