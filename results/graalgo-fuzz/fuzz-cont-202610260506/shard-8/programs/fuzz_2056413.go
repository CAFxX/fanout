package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 int16; F1 int32; F2 int16 }
func genFn0(p0 bool, p1 int8) int16 {
    v5 := []int{127, 4294967295, -9223372036854775808}
    switch (("\n\t" + ("foo bar" + "foo bar")) + (("\n\t" + "hello") + ("" + "\n\t"))) {
    case "":
        v6 := GenS1{F0: (("hello" + ("a" + ("foo bar" + ("a" + "a")))) + "")}
        _ = v6
    case "\n\t":
        if p0 {
            fmt.Println(p0, p1)
            p0 = p0
        }
    }
    fmt.Println(p1)
    v7 := GenS2{F0: -344}
    p1 += p1
    _ = p0
    _ = p1
    _ = v5
    _ = v7
    return 0
}
func genFn1(p0 int) int16 {
    fmt.Println(p0)
    switch (p0 & 127) {
    case 0:
        var v8 float32 = -2.25
        _ = v8
    case -9223372036854775808:
        var v9 uint = 1
        _ = v9
    case -360:
        for fi10 := 0; fi10 < 11; fi10++ {
            switch 729 {
            case 32768:
                v11 := GenS3{F0: 42, F1: 127}
                _ = v11
            default:
                for fi12 := 0; fi12 < 9; fi12++ {
                    switch "a" {
                    case "foo bar":
                        var v13 uint64 = 18446744073709551615
                        _ = v13
                    case "\n\t":
                        var v14 float64 = (0.0 + 1.5)
                        _ = v14
                    default:
                        v15 := []int{fi10, -509}
                        _ = v15
                    }
                    _ = fi12
                }
            }
            _ = fi10
        }
    default:
        fmt.Println(p0)
    }
    _ = p0
    return 32767
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int32 = -2147483648
    v3 := make(map[int]int)
    v4 := make(map[string]uint32)
    for fi16 := 0; fi16 < 9; fi16++ {
        v17 := GenS1{F0: "\n\t"}
        v18 := GenS3{F0: -32768, F1: -128}
        var v19 string = "a"
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    fmt.Println(v2, v2)
    v20 := GenS4{F0: -1, F1: 255}
    for fi21 := 0; fi21 < 10; fi21++ {
        var v22 bool = false
        _ = fi21
        _ = v22
    }
    if true {
        v23 := GenS5{F0: genFn0(true, 0), F1: (v2 >> ((4294967295) & 63)), F2: 32767}
        v2 -= -1
        if true {
            v2 = (v2 / ((v2) | 1))
            if (float64(-2.25) > 0.1) {
                var v24 int8 = 127
                var v25 uint = 4294967295
                _ = v24
                _ = v25
            }
            v26 := GenS5{F0: -1, F1: (v2 / (((v2 ^ int32(v1))) | 1)), F2: 1}
            _ = v26
        }
        _ = v23
    }
    v27 := GenS1{F0: "foo bar"}
    var v28 int64 = v1
    v29 := GenS3{F0: 42, F1: 127}
    v2 -= v2
    if (true || (genFn1(-52) <= -990)) {
        fmt.Println(v1)
        for fi30 := 0; fi30 < 14; fi30++ {
            v31 := []int{32768, (^(fi30))}
            _ = fi30
            _ = v31
        }
    }
    fmt.Println(v1, v2, v1)
    var v32 float64 = float64(float32(1.5))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
}
