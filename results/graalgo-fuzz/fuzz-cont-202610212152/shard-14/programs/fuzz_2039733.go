package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
func genFn0(p0 bool, p1 uint64, p2 string) string {
    v4 := []int{}
    var v5 int = -2147483648
    switch ((p2 + (p2 + (p2 + ("foo bar" + ("\n\t" + "a"))))) + "") {
    case "foo bar":
        fmt.Println(v5, v5, p1)
    }
    var v6 uint64 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return p2
}
func genFn1(p0 string, p1 uint16, p2 uint64) int8 {
    for fi7 := 0; fi7 < 6; fi7++ {
        if (!(!(1e10 == 3.14159))) {
            var v8 uint16 = 65535
            _ = v8
        }
        v9 := []int{(fi7 + (fi7 / ((fi7) | 1))), (fi7 ^ 22)}
        _ = fi7
        _ = v9
    }
    p0 = "foo bar"
    if true {
        fmt.Println(p2)
        v10 := []int{-9223372036854775808, 0, 50}
        v11 := GenS1{F0: 0, F1: 0}
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := []int{fi12}
        _ = fi12
        _ = v13
    }
    var v14 int16 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 100
}

func main() {
    var v1 int16 = 0
    var v2 float32 = float32((-2.25 / (1e10 / 1.5)))
    var v3 uint32 = 340
    var v15 string = (("foo bar" + "") + "\n\t")
    switch ("foo bar" + "foo bar") {
    case "hello":
        if (1 < 36) {
            var v16 uint64 = 0
            v1 = (v1 - v1)
            var v17 uint64 = v16
            _ = v16
            _ = v17
        }
    default:
        fmt.Println(v2)
    }
    if true {
        var v18 uint32 = 438
        switch 128 {
        case 32767:
            for fi19 := 0; fi19 < 11; fi19++ {
                var v20 uint64 = 0
                var v21 uint = 4294967295
                _ = fi19
                _ = v20
                _ = v21
            }
        case -1:
            for fi22 := 0; fi22 < 16; fi22++ {
                var v23 int8 = 123
                v2 = ((3.14159 - v2) * (float32((0.1 - (-2.25 - 1e10))) * v2))
                for fi24 := 0; fi24 < 11; fi24++ {
                    fmt.Println(fi24, v18, v2)
                    v25 := []int{(fi24 / ((fi24) | 1)), (fi24 % ((fi22) | 1)), -37, 127}
                    v18 -= v18
                    _ = fi24
                    _ = v25
                }
                _ = fi22
                _ = v23
            }
        default:
            fmt.Println(v1, v18)
        }
        v15 = v15
        _ = v18
    } else {
        var v26 int8 = -1
        _ = v26
    }
    v27 := []int{927, 2147483647, 2147483647}
    var v28 int16 = 747
    var v29 bool = false
    v1 = v28
    for fi30 := 0; fi30 < 19; fi30++ {
        var v31 int8 = 127
        _ = fi30
        _ = v31
    }
    if v29 {
        var v32 bool = false
        if len(v27) >= 2 {
            v27 = v27[1:len(v27)]
        }
        var v33 uint64 = 255
        _ = v32
        _ = v33
    } else {
        var v34 int32 = 42
        _ = v34
    }
    v28 += v28
    var v35 uint64 = 701
    v15 = "foo bar"
    var v36 uint8 = 255
    var v37 uint16 = 2
    switch 42 {
    case -2147483648:
        fmt.Println(v2, v3, v36)
    case 4294967295:
        _ = len(v27)
    }
    for fi38 := 0; fi38 < 20; fi38++ {
        var v39 float32 = ((v2 + -2.25) * v2)
        v27 = append(v27, -2147483648)
        fmt.Println(v29)
        _ = fi38
        _ = v39
    }
    fmt.Println(v2, v37)
    for fi40 := 0; fi40 < 1; fi40++ {
        var v41 float64 = 1.5
        v15 = "a"
        v42 := GenS1{F0: 0, F1: -1}
        _ = fi40
        _ = v41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v37
}
