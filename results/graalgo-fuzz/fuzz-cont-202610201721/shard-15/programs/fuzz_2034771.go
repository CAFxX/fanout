package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 string; F1 string; F2 int32 }
type GenS3 struct { F0 uint32; F1 uint; F2 int }
type GenS4 struct { F0 uint32; F1 int8; F2 int }
func genFn0(p0 int64, p1 int) uint8 {
    if (p1 > p1) {
        var v6 int16 = -1
        _ = v6
    } else {
        v7 := []int{}
        _ = v7
    }
    var v8 float32 = 1.5
    v9 := GenS1{F0: -32768, F1: (!(65535 <= 931))}
    p1 = 4294967295
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 38
}
func genFn1() uint8 {
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 uint16 = 32767
        _ = fi10
        _ = v11
    }
    var v12 float64 = (0.0 - -2.25)
    for fi13 := 0; fi13 < 16; fi13++ {
        for fi14 := 0; fi14 < 12; fi14++ {
            if (2 >= 2) {
                if (!(v12 < v12)) {
                    v12 = 0.0
                    var v15 uint32 = 0
                    v12 = 3.14159
                    _ = v15
                }
            }
            var v16 int = fi14
            fmt.Println(fi13)
            if ((484 <= 211) || (!((!((1.5 * v12) <= (1e10 - v12))) || true))) {
                var v17 uint = 255
                v16 += fi14
                _ = v17
            } else {
                v12 = v12
                var v18 int = (v16 / ((v16) | 1))
                var v19 bool = (1 <= 42)
                _ = v18
                _ = v19
            }
            fmt.Println(v12)
            switch "hello" {
            case "hello":
                var v20 uint64 = 42
                _ = v20
            }
            v16 = (v16 + fi13)
            _ = fi14
            _ = v16
        }
        for fi21 := 0; fi21 < 4; fi21++ {
            v12 = ((((float64(-2.25) + v12) - -2.25) / 0.1) * ((1.5 - (v12 + 1e10)) / v12))
            var v22 uint16 = 407
            v23 := []int{fi13}
            _ = fi21
            _ = v22
            _ = v23
        }
        var v24 uint64 = 0
        _ = fi13
        _ = v24
    }
    _ = v12
    return genFn0(42, 9223372036854775807)
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{9223372036854775807, 32768, 191, 1}
    var v3 bool = false
    var v4 bool = (!v3)
    v5 := []uint32{65535}
    v4 = true
    for fi25 := 0; fi25 < 2; fi25++ {
        fmt.Println(v3, v4, v4)
        v4 = false
        _ = fi25
    }
    switch -537 {
    case -486:
        v4 = v3
    case 745:
        var v26 uint = 255
        _ = v26
    case 127:
        v4 = ((((!false) && v4) || v3) && (!v4))
    default:
        v27 := GenS1{F0: -1, F1: v4}
        _ = v27
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        v29 := GenS1{F0: 0, F1: v3}
        for fi30 := 0; fi30 < 1; fi30++ {
            v31 := GenS1{F0: 519, F1: (32767 == (-(fi28)))}
            v32 := GenS2{F0: ("\n\t" + (("a" + (("" + "hello") + "")) + ("a" + ""))), F1: "hello", F2: 42}
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = fi28
        _ = v29
    }
    var v33 uint64 = 0
    fmt.Println(v4, v33)
    _ = len(v2)
    switch 2147483647 {
    case -41:
        var v34 uint16 = 0
        _ = v34
    case -675:
        v35 := GenS3{F0: 2, F1: 42, F2: 32767}
        _ = v35
    case 42:
        var v36 int = 128
        _ = v36
    default:
        v1[860] = (2147483647 == 4294967295)
    }
    if (!(false && v4)) {
        v37 := GenS4{F0: 2147483647, F1: -1, F2: 2147483647}
        v38 := GenS2{F0: "foo bar", F1: "foo bar", F2: 1}
        _ = v37
        _ = v38
    } else {
        for fi39 := 0; fi39 < 14; fi39++ {
            var v40 uint16 = 0
            _ = fi39
            _ = v40
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
