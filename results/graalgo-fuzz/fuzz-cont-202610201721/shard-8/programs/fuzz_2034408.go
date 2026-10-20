package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint16 }
func genFn0(p0 uint16, p1 uint, p2 int) string {
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := []int{}
        v9 := GenS1{F0: (fi7 * fi7), F1: 1}
        p0 += (p0 << (((^(p1))) & 63))
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{p2, p2, p2, p2}
    if true {
        var v11 uint = p1
        v12 := GenS1{F0: (p2 << (((p1 | v11)) & 63)), F1: 1}
        _ = v11
        _ = v12
    } else {
        for fi13 := 0; fi13 < 8; fi13++ {
            fmt.Println(p0)
            v14 := GenS1{F0: (fi13 ^ p2), F1: 128}
            _ = fi13
            _ = v14
        }
    }
    if true {
        p1 -= p1
        var v15 uint8 = 81
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return "hello"
}

func main() {
    var v1 uint8 = 128
    var v2 int = 4294967295
    var v3 int16 = -32768
    var v4 uint8 = 42
    var v5 bool = (0 == 18446744073709551615)
    v6 := make(map[string]uint32)
    v3 = v3
    v16 := GenS1{F0: (v2 >> ((1) & 63)), F1: 97}
    for fi17 := 0; fi17 < 20; fi17++ {
        v3 -= v3
        v1 -= 255
        _ = fi17
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        var v19 int64 = 822
        _ = fi18
        _ = v19
    }
    v20 := GenS2{F0: 42, F1: 897, F2: 2}
    v21 := v6[("foo bar" + "")]
    for fi22 := 0; fi22 < 13; fi22++ {
        v4 = (v1 ^ v4)
        for fi23 := 0; fi23 < 3; fi23++ {
            var v24 uint8 = v1
            v25 := []int{fi23}
            var v26 float32 = -2.25
            var v27 int16 = -1
            if false {
                switch "foo bar" {
                case "\n\t":
                    var v28 float64 = (((1e10 / 1.0) / 1.0) - 0.0)
                    _ = v28
                }
            } else {
                var v29 float32 = ((-2.25 * (1e10 + 0.0)) * 1e10)
                if false {
                    delete(v6, "hello")
                    var v30 int16 = (v3 * (v27 ^ v27))
                    _ = v30
                } else {
                    var v31 uint32 = 0
                    var v32 int8 = 109
                    _ = v31
                    _ = v32
                }
                v24 += v4
                _ = v29
            }
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
            _ = v27
        }
        var v33 int = v2
        _ = fi22
        _ = v33
    }
    var v34 float32 = 0.1
    v3 -= v3
    for fi35 := 0; fi35 < 3; fi35++ {
        switch -1 {
        case 42:
            var v36 uint16 = 260
            _ = v36
        default:
            v37 := []int{(^(fi35)), (v2 + v2), -719}
            _ = v37
        }
        for fi38 := 0; fi38 < 2; fi38++ {
            delete(v6, "a")
            var v39 uint16 = 42
            if (0 < -797) {
                v5 = v5
            } else {
                v21 = (^(v21))
            }
            _ = fi38
            _ = v39
        }
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
