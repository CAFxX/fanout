package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 string }
func genFn0(p0 int, p1 uint32) bool {
    var v6 uint8 = 43
    if (int8(p0) <= 114) {
        var v7 int8 = 127
        _ = v7
    } else {
        var v8 float32 = 1e10
        p1 = 2
        var v9 int16 = 342
        _ = v8
        _ = v9
    }
    v10 := []int{p0, (p0 << ((42) & 63)), 9223372036854775807, (p0 % (((p0 ^ (p0 + -186))) | 1))}
    var v11 uint8 = v6
    if ((!true) || ((^(v6)) > 2)) {
        switch int(p1) {
        case -9223372036854775808:
            for fi12 := 0; fi12 < 3; fi12++ {
                var v13 int32 = -528
                if ((1 > 357) && (((-2.25 / (1.5 + float32(((1.5 - (0.1 + (1.5 / ((float64(3.14159) * 1.5) / ((3.14159 - -2.25) + 1e10))))) - 3.14159)))) + (0.1 * ((0.1 / float32((1e10 / 3.14159))) + 0.0))) == 3.14159)) {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    var v14 bool = ("foo bar" != "")
                    v13 = v13
                    _ = v14
                } else {
                    var v15 string = "hello"
                    _ = v15
                }
                _ = fi12
                _ = v13
            }
        case -2147483648:
            var v16 uint = 2
            _ = v16
        default:
            fmt.Println(v11, v11)
        }
        for fi17 := 0; fi17 < 9; fi17++ {
            v11 -= v6
            var v18 float32 = 1.5
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 11; fi19++ {
            if true {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                v20 := GenS1{F0: (v11 | v11), F1: 0}
                var v21 uint16 = 32767
                _ = v20
                _ = v21
            } else {
                var v22 uint16 = 42
                var v23 float32 = 0.1
                var v24 int64 = 0
                _ = v22
                _ = v23
                _ = v24
            }
            _ = fi19
        }
    } else {
        switch (p0 | p0) {
        case 1:
            var v25 int64 = -1
            _ = v25
        case -1:
            v26 := GenS2{F0: "\n\t"}
            _ = v26
        case 2147483647:
            var v27 float32 = (1e10 + 1e10)
            _ = v27
        }
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v6
    return false
}

func main() {
    var v1 float32 = (float32(1.5) - (3.14159 / 3.14159))
    var v2 uint = 255
    v3 := []uint32{uint32(v2)}
    var v4 bool = true
    var v5 string = ("a" + "foo bar")
    for fi28 := 0; fi28 < 1; fi28++ {
        v2 = (v2 | 1000)
        _ = fi28
    }
    if len(v3) > 0 {
        v3[0] = 0
    }
    v5 = v5
    var v29 uint64 = 546
    v30 := GenS2{F0: v5}
    fmt.Println(v29, v2, v5)
    for fi31 := 0; fi31 < 4; fi31++ {
        var v32 string = v5
        _ = fi31
        _ = v32
    }
    var v33 int8 = -97
    var v34 uint32 = 832
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v35 int = 838
    var v36 string = v5
    fmt.Println(v36, v5)
    for fi37 := 0; fi37 < 6; fi37++ {
        var v38 uint64 = v29
        _ = fi37
        _ = v38
    }
    v4 = v4
    v1 += (0.0 * -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
