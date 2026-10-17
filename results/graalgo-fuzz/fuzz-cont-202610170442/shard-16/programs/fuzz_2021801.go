package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint, p1 bool) int32 {
    if (2147483647 <= 0) {
        p1 = false
        for fi7 := 0; fi7 < 16; fi7++ {
            if p1 {
                v8 := GenS1{F0: 33, F1: 32767, F2: 166}
                var v9 int8 = 86
                _ = v8
                _ = v9
            } else {
                var v10 float64 = float64(float32(1e10))
                _ = v10
            }
            _ = fi7
        }
    } else {
        fmt.Println(p1, p1)
        switch 42 {
        case 32768:
            fmt.Println(p1, p1)
        default:
            v11 := GenS1{F0: -532, F1: 649, F2: 2}
            _ = v11
        }
    }
    v12 := []int{}
    if (!true) {
        for fi13 := 0; fi13 < 7; fi13++ {
            _ = len(v12)
            _ = fi13
        }
        fmt.Println(p1, p0)
        var v14 int64 = -1
        _ = v14
    } else {
        var v15 float32 = 1.5
        v16 := GenS1{F0: 0, F1: -32768, F2: 65535}
        switch 0 {
        case 42:
            v15 = (v15 * v15)
        case -212:
            var v17 uint16 = 32767
            _ = v17
        case -416:
            v15 = (0.1 + v15)
        default:
            var v18 uint8 = 2
            _ = v18
        }
        _ = v15
        _ = v16
    }
    p0 -= 18446744073709551615
    _ = p0
    _ = p1
    _ = v12
    return 1
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint = 511
    var v3 int8 = 42
    var v4 float64 = -2.25
    var v5 uint8 = 204
    v6 := make(map[uint32]bool)
    if (false || false) {
        for fi19 := 0; fi19 < 7; fi19++ {
            _ = len(v6)
            var v20 int16 = 42
            for fi21 := 0; fi21 < 4; fi21++ {
                if true {
                    fmt.Println(v3, v1, fi19)
                }
                v22 := []int{(fi21 << ((v2) & 63))}
                _ = fi21
                _ = v22
            }
            _ = fi19
            _ = v20
        }
    }
    v3 = 127
    if (65535 <= 1) {
        if (v4 == float64(float32(v4))) {
            for fi23 := 0; fi23 < 7; fi23++ {
                var v24 string = "hello"
                var v25 int32 = 0
                _ = fi23
                _ = v24
                _ = v25
            }
        }
        v5 = v5
        if false {
            v26 := GenS1{F0: 9223372036854775807, F1: -1, F2: v1}
            _ = v26
        }
    } else {
        var v27 float64 = 1.5
        switch "foo bar" {
        case "":
            fmt.Println(v5, v27, v4)
        case "hello":
            for fi28 := 0; fi28 < 13; fi28++ {
                v4 = v4
                _ = fi28
            }
        case "a":
            fmt.Println(v4, v27)
        }
        _ = v27
    }
    if (0 > 0) {
        if true {
            fmt.Println(v5, v3, v5)
            v29 := GenS1{F0: -439, F1: 32767, F2: 2}
            _ = v29
        }
        if false {
            var v30 float64 = v4
            v30 -= v4
            _ = v30
        } else {
            v31 := GenS1{F0: 0, F1: -32768, F2: 152}
            _ = v31
        }
        var v32 int32 = 0
        _ = v32
    }
    var v33 string = "hello"
    if true {
        v34 := []int{32768, 32767, -91, 32767}
        _ = v34
    }
    var v35 uint64 = 0
    delete(v6, (v1 >> ((234) & 63)))
    if ((int64(v3) >= 42) || true) {
        v36 := GenS1{F0: 42, F1: 42, F2: uint32(v5)}
        _ = v36
    } else {
        if false {
            _ = len(v6)
            if false {
                v37 := []int{}
                v2 = (v2 ^ 820)
                if true {
                    v38 := v6[(v1 * v1)]
                    var v39 int16 = 186
                    _ = v38
                    _ = v39
                } else {
                    var v40 uint = 52
                    _ = v40
                }
                _ = v37
            }
        }
    }
    for fi41 := 0; fi41 < 17; fi41++ {
        switch (fi41 - -2147483648) {
        case 660:
            for fi42 := 0; fi42 < 13; fi42++ {
                var v43 int32 = 0
                v4 -= v4
                _ = fi42
                _ = v43
            }
        case 0:
            v3 = 0
        }
        _ = fi41
    }
    v44 := v6[uint32(v5)]
    v45 := GenS2{F0: (v5 - v5)}
    for fi46 := 0; fi46 < 15; fi46++ {
        delete(v6, (v1 | v1))
        v4 = v4
        _ = fi46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v35
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
