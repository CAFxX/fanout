package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 bool; F2 int }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint8, p1 uint16, p2 uint8) string {
    v7 := []int{605, 127}
    p1 = (p1 / ((65535) | 1))
    var v8 uint = 4294967295
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return ("hello" + "foo bar")
}
func genFn1() uint64 {
    var v9 uint64 = 0
    if (3.14159 <= ((0.0 - 1e10) * 1e10)) {
        v10 := GenS1{F0: 117}
        switch ((("\n\t" + "\n\t") + "") + "") {
        case "":
            for fi11 := 0; fi11 < 18; fi11++ {
                v12 := GenS2{F0: (v9 / ((v9) | 1)), F1: true, F2: (fi11 >> ((0) & 63))}
                _ = fi11
                _ = v12
            }
        }
        _ = v10
    } else {
        var v13 uint16 = 1
        var v14 uint32 = 2147483647
        v15 := []int{-1, 0, 2147483647, -751}
        _ = v13
        _ = v14
        _ = v15
    }
    var v16 float32 = 1e10
    _ = v16
    _ = v9
    return v9
}

func main() {
    var v1 uint = 4294967295
    var v2 uint = v1
    var v3 int8 = 127
    var v4 int32 = -1
    v5 := []string{}
    var v6 uint16 = 32767
    v17 := GenS2{F0: 575, F1: ((2 != 18446744073709551615) || false), F2: 127}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = len(v5)
    fmt.Println(v3, v3, v6)
    var v18 bool = (true || true)
    v6 = (v6 & v6)
    var v19 uint32 = uint32(v6)
    var v20 uint32 = v19
    var v21 float64 = 0.1
    v20 -= v19
    v22 := GenS1{F0: v2}
    _ = len(v5)
    for fi23 := 0; fi23 < 12; fi23++ {
        var v24 uint32 = v19
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        switch 9223372036854775807 {
        case 9223372036854775807:
            _ = len(v5)
        case 127:
            for fi26 := 0; fi26 < 19; fi26++ {
                var v27 int = -37
                var v28 uint64 = 0
                _ = fi26
                _ = v27
                _ = v28
            }
        case -2147483648:
            v29 := GenS3{F0: (^(v20))}
            _ = v29
        }
        _ = fi25
    }
    var v30 float32 = (-2.25 - float32((-2.25 / ((v21 + v21) * float64(0.0)))))
    v6 -= v6
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
