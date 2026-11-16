package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint32 }
func genFn0(p0 int64, p1 uint64) uint32 {
    switch -1 {
    case 256:
        if (p0 != -225) {
            var v4 uint32 = 65535
            _ = v4
        }
    default:
        p1 = (p1 * 951)
    }
    v5 := []int{42, -1, 42}
    _ = p0
    _ = p1
    _ = v5
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = -552
    v3 := []string{"a", "a", (("a" + ("hello" + ("a" + "a"))) + "foo bar"), "foo bar"}
    if true {
        v6 := GenS1{F0: 0, F1: (v2 >> ((18446744073709551615) & 63)), F2: 516}
        _ = v6
    }
    var v7 int64 = -1
    v1[225] = "\n\t"
    fmt.Println(v7, v7)
    var v8 int32 = 0
    v8 += 0
    v8 = v8
    v9 := GenS2{F0: 65535}
    var v10 float32 = (1e10 - ((1.5 / 1.5) * (1e10 - 1.5)))
    var v11 uint64 = 2
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        v7 -= (-(v7))
        var v13 string = "a"
        _ = fi12
        _ = v13
    }
    v7 += 9223372036854775807
    v8 -= v8
    var v14 uint8 = 2
    v15 := GenS3{F0: v14, F1: 32767, F2: 477}
    for fi16 := 0; fi16 < 1; fi16++ {
        for fi17 := 0; fi17 < 1; fi17++ {
            var v18 int8 = 56
            if ((^(v2)) <= int16(v18)) {
                var v19 uint32 = 65535
                var v20 uint64 = v11
                _ = v19
                _ = v20
            } else {
                var v21 int = (fi16 * fi16)
                var v22 uint8 = v14
                v3 = append(v3, (("" + "a") + ((((("\n\t" + "hello") + ("" + "a")) + ("" + "")) + "") + "")))
                _ = v21
                _ = v22
            }
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    if (977 <= (v7 | (v7 & (v7 - (v7 << ((42) & 63)))))) {
        switch "hello" {
        case "\n\t":
            for fi23 := 0; fi23 < 4; fi23++ {
                v10 -= v10
                _ = fi23
            }
        case "":
            var v24 int32 = v8
            _ = v24
        case "hello":
            fmt.Println(v2)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
