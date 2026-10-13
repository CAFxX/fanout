package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int; F1 int64; F2 uint }
type GenS3 struct { F0 int16 }
func genFn0() bool {
    var v6 float32 = float32((1.5 * (1e10 / 1.0)))
    var v7 string = ("hello" + "")
    var v8 int = 107
    v7 = "hello"
    var v9 uint64 = 4294967295
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (v9 == v9)
}
func genFn1(p0 uint8, p1 int16, p2 uint16) int64 {
    if true {
        fmt.Println(p0)
        var v10 uint64 = 0
        _ = v10
    } else {
        fmt.Println(p2)
        fmt.Println(p2, p2, p2)
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 int16 = 1
        var v13 int8 = 127
        _ = fi11
        _ = v12
        _ = v13
    }
    if (p0 <= 1) {
        v14 := []int{0, 32767}
        p2 += (p2 >> ((2) & 63))
        var v15 int8 = 127
        _ = v14
        _ = v15
    } else {
        for fi16 := 0; fi16 < 19; fi16++ {
            for fi17 := 0; fi17 < 15; fi17++ {
                p1 = (p1 ^ (p1 % ((-32768) | 1)))
                switch (fi17 * -9223372036854775808) {
                case -654:
                    v18 := []int{(-(fi17)), fi17, fi17, -1}
                    _ = v18
                case 42:
                    v19 := GenS1{F0: 30, F1: 908}
                    _ = v19
                default:
                    var v20 uint8 = p0
                    _ = v20
                }
                var v21 uint32 = uint32(p0)
                _ = fi17
                _ = v21
            }
            v22 := GenS1{F0: 0, F1: 42}
            _ = fi16
            _ = v22
        }
        v23 := []int{-135, -2147483648, 9223372036854775807}
        for fi24 := 0; fi24 < 18; fi24++ {
            if len(v23) > 0 {
                v23[0] = (fi24 ^ 718)
            }
            if false {
                if false {
                    v25 := GenS1{F0: 1, F1: 42}
                    _ = v25
                }
            }
            _ = fi24
        }
        _ = v23
    }
    p0 = (^(p0))
    var v26 uint16 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    return 1
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[string]uint32)
    var v3 float32 = 1.5
    v4 := make(map[int]uint32)
    var v5 uint32 = (v1 / ((4294967295) | 1))
    v27 := GenS1{F0: 75, F1: 160}
    fmt.Println(v3, v5, v5)
    v28 := []int{-9223372036854775808}
    switch "\n\t" {
    case "\n\t":
        if len(v28) > 0 {
            v28[0] = 128
        }
    }
    for fi29 := 0; fi29 < 7; fi29++ {
        v5 = (v1 | (^(v5)))
        v1 -= (v5 ^ 890)
        for fi30 := 0; fi30 < 20; fi30++ {
            v31 := GenS1{F0: -47, F1: 4294967295}
            var v32 bool = false
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = fi29
    }
    v3 -= (((((v3 + float32(((((float64((1e10 * float32((((((float64(v3) * 1.5) + 0.1) / 1.0) * (float64((float32(((((1.5 * 1.5) + 1e10) + 0.1) / 1e10)) - (v3 * (0.1 - 0.1)))) * (0.0 * 0.1))) + float64(0.0))))) - -2.25) + (-2.25 - 1e10)) - (1.5 * 1.5)) * 0.1))) * (v3 * (0.0 - v3))) + 1e10) * v3) - (v3 * 0.1))
    var v33 string = "a"
    v3 += 3.14159
    if (65535 != 42) {
        var v34 uint8 = 175
        _ = len(v4)
        _ = v34
    } else {
        delete(v2, "hello")
        v35 := GenS2{F0: -456, F1: 962, F2: 0}
        fmt.Println(v33, v1)
        _ = v35
    }
    if len(v28) > 0 {
        v28[0] = 590
    }
    if false {
        if (1 > 112) {
            switch 1 {
            case 128:
                var v36 uint32 = (^(v5))
                _ = v36
            case -1:
                for fi37 := 0; fi37 < 15; fi37++ {
                    v38 := GenS3{F0: 32767}
                    var v39 bool = (!(-1 < -971))
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            case 510:
                v5 = (v5 / (((v1 ^ v5)) | 1))
            default:
                var v40 int8 = 127
                _ = v40
            }
            v41 := GenS3{F0: 42}
            var v42 int64 = genFn1(255, 1, 2)
            _ = v41
            _ = v42
        }
        v43 := v4[4294967295]
        v5 -= v5
        _ = v43
    } else {
        switch -2147483648 {
        case -1:
            var v44 int8 = -128
            _ = v44
        default:
            if (!(215 <= 42)) {
                var v45 int8 = 100
                v46 := v4[128]
                var v47 uint32 = v5
                _ = v45
                _ = v46
                _ = v47
            }
        }
        _ = len(v4)
        if (v3 <= float32(3.14159)) {
            v1 = (^(v5))
        }
    }
    fmt.Println(v33, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
