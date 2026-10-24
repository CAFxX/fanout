package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 string; F1 uint8; F2 int8 }

func main() {
    var v1 int8 = 127
    var v2 int = 0
    var v3 float32 = ((-2.25 - 0.0) * ((float32(-2.25) / 1.0) + 1e10))
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]string)
    v1 = (^(v1))
    var v6 uint8 = 130
    if false {
        v6 = v6
    } else {
        v7 := GenS1{F0: (^(v2)), F1: 42}
        var v8 uint32 = 65535
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        switch "" {
        case "\n\t":
            for fi10 := 0; fi10 < 6; fi10++ {
                var v11 float32 = float32(((0.0 / (float64(1e10) + (((3.14159 - (1e10 * ((float64(v3) - 0.1) - -2.25))) * float64(3.14159)) - (-2.25 - 0.0)))) * 0.1))
                var v12 int = (fi10 << ((18446744073709551615) & 63))
                if false {
                    fmt.Println(v6)
                }
                _ = fi10
                _ = v11
                _ = v12
            }
        default:
            fmt.Println(fi9)
        }
        v13 := GenS1{F0: int(v6), F1: 0}
        _ = fi9
        _ = v13
    }
    v14 := GenS1{F0: (v2 / ((v2) | 1)), F1: -507}
    var v15 float32 = v3
    var v16 int = v2
    switch -608 {
    case 127:
        if (false || (v1 == (v1 >> ((1) & 63)))) {
            for fi17 := 0; fi17 < 7; fi17++ {
                var v18 bool = (!(42 != 42))
                for fi19 := 0; fi19 < 2; fi19++ {
                    v15 += 3.14159
                    var v20 int64 = -9223372036854775808
                    _ = fi19
                    _ = v20
                }
                v21 := []int{739, (v16 * v2), (fi17 / (((v2 / ((128) | 1))) | 1)), (fi17 ^ (fi17 & (v16 - (v2 | v2))))}
                _ = fi17
                _ = v18
                _ = v21
            }
            if true {
                for fi22 := 0; fi22 < 6; fi22++ {
                    var v23 bool = false
                    v23 = true
                    fmt.Println(v3, v3, v1)
                    _ = fi22
                    _ = v23
                }
                var v24 float32 = 1e10
                v25 := v5[631]
                _ = v24
                _ = v25
            }
            v5[0] = "foo bar"
        } else {
            fmt.Println(v6)
            v26 := GenS1{F0: (-(v2)), F1: -32768}
            v27 := GenS1{F0: int(v6), F1: 42}
            _ = v26
            _ = v27
        }
    default:
        v3 -= (0.1 / float32(-2.25))
    }
    var v28 int8 = (v1 * (v1 & 42))
    var v29 float32 = -2.25
    _ = len(v5)
    v30 := GenS1{F0: -2147483648, F1: 32767}
    if (!(!true)) {
        if true {
            var v31 uint = 27
            _ = v31
        } else {
            var v32 uint32 = 4294967295
            _ = v32
        }
        var v33 int64 = -9223372036854775808
        _ = v33
    }
    switch (("" + "hello") + "hello") {
    case "\n\t":
        v15 = (v15 * 0.0)
    case "a":
        var v34 int16 = 1000
        _ = v34
    case "":
        v29 = float32(0.0)
    }
    v35 := GenS2{F0: "foo bar", F1: 111, F2: v28}
    v36 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
