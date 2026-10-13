package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS2 struct { F0 int64; F1 uint8; F2 string }
type GenS3 struct { F0 int64; F1 string }
func genFn0(p0 string) uint32 {
    var v7 int16 = -1
    v8 := []int{-331}
    _ = p0
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[int]bool)
    v3 := []string{"\n\t", "\n\t", "hello"}
    v4 := make(map[string]int)
    var v5 uint = 1
    v6 := make(map[int]int)
    fmt.Println(v5, v1)
    v5 = 0
    if (false && true) {
        v9 := GenS1{F0: 0, F1: -128, F2: 32767}
        if (-2.25 > ((1e10 - 1e10) * (0.1 - (-2.25 + (0.0 - 0.1))))) {
            for fi10 := 0; fi10 < 11; fi10++ {
                if (!(fi10 > 127)) {
                    fmt.Println(v5, v1)
                    v1 -= 18446744073709551615
                } else {
                    fmt.Println(v1, v1, v5)
                }
                _ = fi10
            }
            var v11 int64 = 1
            fmt.Println(v11, v1, v11)
            _ = v11
        } else {
            delete(v6, 9223372036854775807)
            v12 := GenS1{F0: 0, F1: -95, F2: 32767}
            _ = v12
        }
        v5 -= (v5 - (v5 >> (((v5 | v5)) & 63)))
        _ = v9
    }
    var v13 float32 = (float32(((-2.25 + 3.14159) * -2.25)) + ((0.0 / 3.14159) * ((0.0 + ((float32((0.0 + 3.14159)) - -2.25) * (0.1 / 1.0))) - 0.1)))
    for fi14 := 0; fi14 < 3; fi14++ {
        fmt.Println(v5, fi14)
        v15 := v2[0]
        v1 += v1
        _ = fi14
        _ = v15
    }
    var v16 bool = false
    v17 := GenS1{F0: 498, F1: 100, F2: 65535}
    switch 32767 {
    case -615:
        v18 := GenS1{F0: 2, F1: -51, F2: 234}
        _ = v18
    case 0:
        v13 = ((0.0 - v13) / v13)
    case -2147483648:
        v19 := GenS1{F0: 678, F1: -1, F2: 0}
        _ = v19
    default:
        var v20 int16 = 223
        _ = v20
    }
    v21 := GenS2{F0: 503, F1: 0, F2: "hello"}
    for fi22 := 0; fi22 < 14; fi22++ {
        v13 += v13
        if true {
            v23 := GenS2{F0: 0, F1: 152, F2: ("hello" + ("foo bar" + ("" + (("foo bar" + "foo bar") + ""))))}
            _ = v23
        } else {
            var v24 bool = (v16 || (!true))
            var v25 int64 = 1
            _ = v24
            _ = v25
        }
        v26 := v6[(fi22 | (fi22 << (((v5 & (v5 & v5))) & 63)))]
        _ = fi22
        _ = v26
    }
    if true {
        v27 := GenS3{F0: 1, F1: "hello"}
        v5 += (v5 * (v5 >> ((v5) & 63)))
        _ = v27
    }
    v2[-1] = (v5 == (v5 << ((v5) & 63)))
    var v28 uint = v5
    var v29 uint32 = 2147483647
    if (!v16) {
        v30 := GenS3{F0: 695, F1: "foo bar"}
        _ = v30
    }
    fmt.Println(v13, v5)
    _ = len(v3)
    switch ("" + "\n\t") {
    case "foo bar":
        switch -1 {
        case 1:
            delete(v6, 42)
        }
    case "hello":
        for fi31 := 0; fi31 < 2; fi31++ {
            v13 = v13
            _ = fi31
        }
    default:
        if (-2147483648 == -1) {
            fmt.Println(v29)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
