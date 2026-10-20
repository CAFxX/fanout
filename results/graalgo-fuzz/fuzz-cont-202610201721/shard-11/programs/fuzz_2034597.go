package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int }
type GenS2 struct { F0 int64; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 string, p1 int16, p2 int16) int32 {
    var v4 float32 = 3.14159
    switch 32768 {
    case 2147483647:
        var v5 uint = 996
        _ = v5
    default:
        p2 = (-(p2))
    }
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := []int{fi6, -186, -1, (fi6 & 0)}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 42
}

func main() {
    var v1 uint = 543
    v2 := make(map[uint32]uint32)
    var v3 string = ("" + ((("foo bar" + "") + ("" + ("" + ((("foo bar" + ("" + ("" + ("" + "\n\t")))) + "hello") + "hello")))) + "foo bar"))
    v3 = (v3 + (v3 + ((v3 + v3) + v3)))
    switch 128 {
    case 2147483647:
        if (!true) {
            var v8 int32 = -615
            var v9 bool = (1e10 > 0.0)
            if (98 < 635) {
                for fi10 := 0; fi10 < 9; fi10++ {
                    v9 = (!v9)
                    v3 = v3
                    _ = fi10
                }
                var v11 bool = (!(!(-837 > -585)))
                _ = v11
            } else {
                if (2 < 42) {
                    v8 = (v8 | -372)
                    v8 += 718
                    v12 := GenS1{F0: v1, F1: 255, F2: 32767}
                    _ = v12
                } else {
                    for fi13 := 0; fi13 < 20; fi13++ {
                        var v14 uint = (^(v1))
                        var v15 int64 = -9223372036854775808
                        var v16 int = fi13
                        _ = fi13
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                    var v17 bool = (2 >= 986)
                    _ = v17
                }
                v8 -= (v8 | (^(v8)))
            }
            _ = v8
            _ = v9
        }
    case 9223372036854775807:
        v1 -= 255
    case 32767:
        if true {
            var v18 bool = ((3.14159 + (0.1 / (0.1 + 0.1))) <= 3.14159)
            _ = v18
        }
    default:
        v19 := []int{135, 0, 127}
        _ = v19
    }
    if true {
        if (true || (!(true || (-32768 != 562)))) {
            if (((-2.25 / (float32(-2.25) + -2.25)) - float32(float64(-2.25))) <= float32(float64((3.14159 * float32((1e10 / (-2.25 * (1e10 * ((1e10 * float64(1e10)) * 3.14159))))))))) {
                var v20 float32 = 3.14159
                _ = v20
            }
            if (!false) {
                fmt.Println(v1, v1)
                var v21 bool = (12 < 127)
                v22 := GenS1{F0: v1, F1: 128, F2: int(v1)}
                _ = v21
                _ = v22
            }
            for fi23 := 0; fi23 < 18; fi23++ {
                for fi24 := 0; fi24 < 7; fi24++ {
                    var v25 uint64 = 2
                    switch (fi23 - fi24) {
                    case 2147483647:
                        v26 := GenS1{F0: v1, F1: 2, F2: 2147483647}
                        _ = v26
                    }
                    _ = fi24
                    _ = v25
                }
                _ = fi23
            }
        }
        for fi27 := 0; fi27 < 5; fi27++ {
            v1 = v1
            _ = fi27
        }
        v3 = (((((v3 + v3) + "hello") + v3) + "") + v3)
    }
    v28 := GenS1{F0: (^(v1)), F1: 123, F2: -1}
    if false {
        if false {
            for fi29 := 0; fi29 < 20; fi29++ {
                var v30 int16 = 0
                _ = fi29
                _ = v30
            }
            var v31 string = "foo bar"
            fmt.Println(v31)
            _ = v31
        } else {
            switch "hello" {
            case "hello":
                var v32 int32 = 1
                _ = v32
            case "":
                v33 := GenS2{F0: 42, F1: -622, F2: 826}
                _ = v33
            }
        }
        v34 := []int{}
        _ = v34
    }
    fmt.Println(v3, v3)
    var v35 int64 = -1
    for fi36 := 0; fi36 < 5; fi36++ {
        v37 := []int{(fi36 / ((fi36) | 1)), (fi36 & (^(fi36)))}
        _ = fi36
        _ = v37
    }
    v38 := GenS3{F0: 542}
    v3 = v3
    var v39 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v35
    _ = v38
    _ = v39
}
