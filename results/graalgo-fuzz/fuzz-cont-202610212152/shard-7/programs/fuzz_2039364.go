package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint64; F2 int32 }
type GenS3 struct { F0 int32 }
func genFn0() bool {
    var v6 int16 = -1
    var v7 int32 = -2147483648
    if false {
        var v8 string = "a"
        v6 += (-(v6))
        switch 876 {
        case 32767:
            v9 := GenS1{F0: 1}
            _ = v9
        case 9223372036854775807:
            v6 = (v6 & (v6 | (v6 << ((438) & 63))))
        case -701:
            v8 = (v8 + v8)
        }
        _ = v8
    }
    v6 = v6
    _ = v6
    _ = v7
    return true
}
func genFn1() uint32 {
    switch "a" {
    case "":
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 float64 = (((0.1 / 1e10) - (3.14159 - (-2.25 / 1.0))) * ((((0.1 * ((((float64((((1.5 * 1.5) * (((0.0 - 0.1) + float32((float64(3.14159) - (1.5 - 0.1)))) * float32(0.1))) + float32((0.1 / 3.14159)))) / 0.1) * 0.1) + (0.1 * ((1.5 / (-2.25 / 1e10)) + 1.5))) / float64(3.14159))) + float64(float32(1e10))) / 1.5) / float64((0.1 * 1e10))))
            if (false && (("\n\t" + ((("a" + ("\n\t" + "a")) + ((("" + "hello") + (("hello" + ("a" + ("" + "hello"))) + "")) + "foo bar")) + "foo bar")) > ("hello" + ""))) {
                var v12 bool = true
                v13 := GenS1{F0: 32767}
                _ = v12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
    case "hello":
        v14 := GenS1{F0: -439}
        _ = v14
    }
    var v15 uint8 = 0
    for fi16 := 0; fi16 < 3; fi16++ {
        if ((0.0 + 3.14159) >= (((-2.25 + -2.25) / (1.5 / 3.14159)) / (-2.25 - (1e10 * 1e10)))) {
            if (65535 <= 0) {
                var v17 int16 = -414
                _ = v17
            } else {
                for fi18 := 0; fi18 < 4; fi18++ {
                    v15 = v15
                    _ = fi18
                }
                v19 := GenS1{F0: -766}
                v15 = (^(v15))
                _ = v19
            }
            v20 := GenS2{F0: v15, F1: 42, F2: 1}
            v15 = v15
            _ = v20
        } else {
            v21 := GenS2{F0: 197, F1: 0, F2: 688}
            for fi22 := 0; fi22 < 4; fi22++ {
                fmt.Println(fi22, v15, fi22)
                if false {
                    switch "a" {
                    case "foo bar":
                        v15 -= (v15 ^ 186)
                    case "hello":
                        v23 := []int{(^(fi16)), fi16, 4}
                        _ = v23
                    }
                    var v24 uint = 4294967295
                    var v25 uint64 = 2
                    _ = v24
                    _ = v25
                } else {
                    v26 := GenS1{F0: 42}
                    _ = v26
                }
                fmt.Println(v15)
                _ = fi22
            }
            _ = v21
        }
        var v27 uint64 = 1
        v15 += (v15 + v15)
        _ = fi16
        _ = v27
    }
    if (true || false) {
        v28 := GenS3{F0: 499}
        var v29 uint32 = 2147483648
        _ = v28
        _ = v29
    } else {
        v15 = (v15 * 2)
    }
    if (true && (1000 < -273)) {
        v30 := []int{32768, -9223372036854775808}
        _ = v30
    }
    _ = v15
    return 1
}

func main() {
    var v1 uint64 = 19
    var v2 int32 = 377
    var v3 int8 = -128
    var v4 int32 = -601
    v5 := make(map[int]int)
    fmt.Println(v4)
    if false {
        fmt.Println(v1)
    }
    _ = len(v5)
    var v31 uint16 = 2
    var v32 uint32 = genFn1()
    var v33 string = ("" + "")
    v32 -= (v32 + v32)
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
