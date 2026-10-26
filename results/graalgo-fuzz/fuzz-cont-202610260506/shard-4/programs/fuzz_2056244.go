package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 bool; F1 int32; F2 uint16 }
type GenS5 struct { F0 bool; F1 uint64; F2 int8 }
func genFn0(p0 uint64, p1 uint64, p2 int16) uint {
    v7 := []int{1}
    fmt.Println(p0)
    switch -503 {
    case -678:
        _ = len(v7)
    }
    var v8 float32 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]string)
    var v2 bool = (-2.25 <= float32((((float64(1e10) / 1e10) + -2.25) / ((0.1 - 0.0) - (-2.25 * (float64(0.0) * (((1.5 - -2.25) * 1.5) + -2.25)))))))
    var v3 bool = (((2147483648 < 2147483648) || v2) && true)
    var v4 uint16 = 32767
    var v5 int64 = 567
    var v6 uint32 = 42
    switch (("foo bar" + "foo bar") + (("hello" + "") + "")) {
    case "\n\t":
        var v9 float32 = (1e10 + (0.1 + (1e10 - (3.14159 - 3.14159))))
        _ = v9
    case "":
        v5 += (v5 % ((v5) | 1))
    }
    v10 := []int{-2147483648, int(v6), 221, -9223372036854775808}
    v11 := GenS1{F0: 32}
    var v12 int64 = (v5 / (((-(v5))) | 1))
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    if v2 {
        v13 := GenS2{F0: "hello"}
        if (!(v3 || (27 > 0))) {
            var v14 int8 = -128
            fmt.Println(v5, v3)
            var v15 float32 = (-2.25 - 1.5)
            _ = v14
            _ = v15
        } else {
            for fi16 := 0; fi16 < 4; fi16++ {
                var v17 uint = 18446744073709551615
                _ = fi16
                _ = v17
            }
        }
        _ = v13
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        if (((-2147483648 <= 194) && v2) && (true && v2)) {
            var v19 int8 = -116
            switch int(v12) {
            case -794:
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            case 731:
                v19 += (v19 % ((101) | 1))
            default:
                var v20 int8 = (v19 - (v19 - (v19 - 13)))
                _ = v20
            }
            _ = v19
        }
        v4 = v4
        _ = fi18
    }
    v6 += (^(v6))
    var v21 uint16 = v4
    if (!true) {
        fmt.Println(v21)
        _ = len(v10)
    } else {
        for fi22 := 0; fi22 < 4; fi22++ {
            v3 = false
            switch "" {
            case "\n\t":
                fmt.Println(v21)
            default:
                for fi23 := 0; fi23 < 19; fi23++ {
                    v4 = (v21 % ((v21) | 1))
                    _ = fi23
                }
            }
            var v24 uint16 = (v21 % (((v21 - v4)) | 1))
            _ = fi22
            _ = v24
        }
        v21 += 826
    }
    if v2 {
        v25 := GenS1{F0: 128}
        v26 := GenS1{F0: 2}
        switch "hello" {
        case "hello":
            fmt.Println(v3, v3, v12)
        }
        _ = v25
        _ = v26
    } else {
        v27 := GenS3{F0: v2, F1: 65535}
        _ = v27
    }
    for fi28 := 0; fi28 < 2; fi28++ {
        var v29 bool = v2
        v30 := GenS4{F0: v3, F1: 2147483647, F2: v21}
        v31 := GenS1{F0: 2}
        _ = fi28
        _ = v29
        _ = v30
        _ = v31
    }
    if true {
        v32 := GenS2{F0: ""}
        v2 = (v2 || (false || v3))
        _ = v32
    } else {
        v33 := GenS5{F0: v3, F1: 108, F2: -31}
        v3 = (true && ((!false) && false))
        _ = v33
    }
    for fi34 := 0; fi34 < 12; fi34++ {
        _ = len(v10)
        _ = fi34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
