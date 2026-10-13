package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 uint; F1 int32 }
func genFn0(p0 string, p1 int8) uint16 {
    var v7 float32 = 0.1
    v8 := []int{-561, 4294967295, 32767, -97}
    p1 -= (p1 + (p1 >> ((4294967295) & 63)))
    if len(v8) > 0 {
        v8[0] = -485
    }
    v7 = (float32(0.0) * (0.1 - 1e10))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int32, p1 string, p2 int8) int64 {
    switch p1 {
    case "\n\t":
        p1 = "hello"
    case "hello":
        if true {
            p1 = ("foo bar" + "hello")
            p2 = (p2 & (^(p2)))
            switch (p1 + ("a" + p1)) {
            case "":
                if (!true) {
                    p2 += (-(p2))
                    for fi9 := 0; fi9 < 16; fi9++ {
                        fmt.Println(p1)
                        v10 := GenS1{F0: 724}
                        v11 := []int{-805, (fi9 & (fi9 % ((fi9) | 1)))}
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                }
            case "a":
                for fi12 := 0; fi12 < 5; fi12++ {
                    var v13 uint32 = 42
                    v14 := GenS2{F0: 2147483647}
                    v15 := []int{(fi12 << ((0) & 63)), fi12}
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
            case "hello":
                p1 = (p1 + p1)
            }
        } else {
            v16 := []int{1, -873, -9223372036854775808, -9223372036854775808}
            for fi17 := 0; fi17 < 17; fi17++ {
                v18 := GenS3{F0: 634}
                for fi19 := 0; fi19 < 15; fi19++ {
                    p0 += (-(p0))
                    p1 = ((((((p1 + (p1 + "\n\t")) + p1) + (p1 + "foo bar")) + "\n\t") + p1) + p1)
                    var v20 uint32 = 4294967295
                    _ = fi19
                    _ = v20
                }
                _ = fi17
                _ = v18
            }
            _ = v16
        }
    default:
        var v21 float64 = 3.14159
        _ = v21
    }
    v22 := GenS4{F0: 640, F1: 998}
    v23 := []int{242, 0}
    fmt.Println(p2)
    var v24 float32 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    _ = v24
    return -878
}

func main() {
    var v1 int8 = 0
    v2 := make(map[int]bool)
    var v3 uint16 = 32767
    v4 := make(map[string]bool)
    v5 := make(map[string]string)
    var v6 float64 = 1e10
    var v25 bool = false
    var v26 int32 = 561
    var v27 bool = (42 <= 5)
    switch 4294967295 {
    case -685:
        _ = len(v2)
    default:
        if v27 {
            fmt.Println(v27, v27)
            switch 2147483647 {
            case -864:
                v28 := []int{-121, 9223372036854775807}
                _ = v28
            default:
                var v29 uint16 = (v3 << ((18446744073709551615) & 63))
                _ = v29
            }
        }
    }
    v30 := GenS2{F0: -2147483648}
    var v31 int64 = 1
    v32 := GenS3{F0: genFn0(("\n\t" + ("\n\t" + "a")), v1)}
    if (true || (v3 <= 0)) {
        v33 := GenS1{F0: 1}
        var v34 float64 = v6
        v35 := v2[-336]
        _ = v33
        _ = v34
        _ = v35
    } else {
        v31 = -853
        v36 := GenS5{F0: 0, F1: (v26 ^ v26)}
        _ = v36
    }
    fmt.Println(v6, v3)
    var v37 float32 = 3.14159
    for fi38 := 0; fi38 < 17; fi38++ {
        var v39 bool = (v1 != (v1 ^ v1))
        fmt.Println(v26)
        v40 := v5[""]
        _ = fi38
        _ = v39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
