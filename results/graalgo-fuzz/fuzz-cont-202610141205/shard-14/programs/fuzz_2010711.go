package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint32 }
type GenS2 struct { F0 int8; F1 int8; F2 int16 }
type GenS3 struct { F0 int64; F1 string; F2 int64 }
type GenS4 struct { F0 string; F1 int16; F2 uint64 }
func genFn0(p0 string) int64 {
    p0 = ((((p0 + "") + (p0 + p0)) + p0) + ("" + ((("foo bar" + "") + p0) + "hello")))
    var v4 int32 = -1
    switch ("" + p0) {
    case "hello":
        p0 = p0
    case "":
        p0 = (p0 + "")
    case "a":
        var v5 uint8 = 129
        _ = v5
    default:
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 float64 = (1.5 - (((float64((0.0 / 0.1)) * (((3.14159 + 3.14159) - -2.25) / 3.14159)) * float64(1e10)) - (float64(1e10) * float64(float32((((3.14159 * ((3.14159 * 3.14159) / 0.1)) / (1.5 - ((0.0 + 0.0) * 1.5))) * 0.0))))))
            _ = fi6
            _ = v7
        }
    }
    v8 := []int{}
    _ = p0
    _ = v4
    _ = v8
    return -9223372036854775808
}

func main() {
    var v1 uint8 = 230
    v2 := []string{}
    var v3 uint8 = (v1 + (v1 << ((42) & 63)))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v9 := GenS1{F0: 32767, F1: -1, F2: 0}
    switch "" {
    case "":
        var v10 int = 32767
        _ = v10
    case "\n\t":
        if true {
            v11 := GenS2{F0: 42, F1: -128, F2: 1}
            switch 0 {
            case 712:
                for fi12 := 0; fi12 < 8; fi12++ {
                    fmt.Println(v3, v3)
                    v13 := GenS3{F0: -1, F1: "hello", F2: 1}
                    _ = fi12
                    _ = v13
                }
            default:
                for fi14 := 0; fi14 < 19; fi14++ {
                    var v15 bool = (true && true)
                    var v16 uint8 = (^(v3))
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
            _ = v11
        } else {
            var v17 uint16 = 1
            v3 -= (v1 - 128)
            _ = v17
        }
    }
    var v18 uint8 = (v3 * 255)
    var v19 bool = true
    v1 = (v1 & 53)
    var v20 uint = 121
    v18 = 231
    switch 717 {
    case -2147483648:
        _ = len(v2)
    case 1:
        for fi21 := 0; fi21 < 11; fi21++ {
            v1 = v3
            _ = fi21
        }
    case -1:
        v2 = append(v2, ("" + "hello"))
    default:
        if v19 {
            fmt.Println(v3, v18)
            v3 = v3
        } else {
            var v22 int32 = -1
            v23 := GenS1{F0: -801, F1: (-(v22)), F2: 2147483648}
            fmt.Println(v18, v18)
            _ = v22
            _ = v23
        }
    }
    if (!(-108 <= -49)) {
        var v24 int64 = 9223372036854775807
        _ = v24
    } else {
        fmt.Println(v20, v20)
        if v19 {
            v1 = (v18 >> (((v20 << (((^(v20))) & 63))) & 63))
            v25 := GenS4{F0: ("foo bar" + (("a" + ("hello" + (("a" + "a") + ("hello" + "foo bar")))) + "")), F1: -280, F2: 42}
            _ = v25
        } else {
            if v19 {
                v2 = append(v2, "")
                var v26 int64 = -1
                var v27 bool = (127 > 100)
                _ = v26
                _ = v27
            }
            var v28 uint8 = (v3 ^ (v1 + 42))
            _ = v28
        }
        for fi29 := 0; fi29 < 13; fi29++ {
            for fi30 := 0; fi30 < 5; fi30++ {
                v31 := GenS4{F0: "foo bar", F1: 32767, F2: uint64(v18)}
                fmt.Println(v18, fi29)
                _ = fi30
                _ = v31
            }
            v1 += (v1 / ((1) | 1))
            var v32 float32 = (-2.25 + -2.25)
            _ = fi29
            _ = v32
        }
    }
    v19 = v19
    var v33 int = -384
    v2 = append(v2, "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v9
}
