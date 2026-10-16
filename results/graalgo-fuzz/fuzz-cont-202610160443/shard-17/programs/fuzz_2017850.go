package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 int64; F2 uint16 }
func genFn0(p0 int, p1 string, p2 uint) int64 {
    p0 += (p0 | p0)
    p0 = (p0 >> ((p2) & 63))
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 513
    var v3 float32 = (3.14159 + (1.5 / 0.1))
    var v4 uint32 = 314
    var v5 int = 32767
    var v6 int32 = -772
    delete(v1, v4)
    if ((v5 - v5) < v5) {
        for fi7 := 0; fi7 < 16; fi7++ {
            if (128 >= 255) {
                v4 += v4
            } else {
                var v8 int = fi7
                var v9 uint8 = 128
                _ = v8
                _ = v9
            }
            var v10 uint16 = 2
            for fi11 := 0; fi11 < 19; fi11++ {
                var v12 float32 = 0.0
                _ = fi11
                _ = v12
            }
            _ = fi7
            _ = v10
        }
    }
    v2 -= v2
    if false {
        if (-9223372036854775808 >= 42) {
            v5 = v5
            var v13 bool = true
            var v14 bool = v13
            _ = v13
            _ = v14
        }
        v15 := GenS1{F0: genFn0((v5 ^ v5), ("foo bar" + "a"), 1)}
        _ = v15
    }
    if ((v4 * (v4 << ((v2) & 63))) != (v4 & v4)) {
        v16 := GenS2{F0: -67, F1: 2}
        if false {
            var v17 float64 = ((3.14159 + 1.5) - 0.1)
            _ = v17
        } else {
            v3 = v3
        }
        _ = v16
    } else {
        v6 = v6
    }
    var v18 int8 = 47
    for fi19 := 0; fi19 < 20; fi19++ {
        switch (v5 & (^(fi19))) {
        case -498:
            v6 = v6
        case -2147483648:
            v4 = v4
        case 128:
            if (3.14159 < -2.25) {
                v20 := GenS1{F0: 26}
                _ = v20
            }
        }
        if (2 < 346) {
            v21 := GenS1{F0: 0}
            _ = v21
        } else {
            var v22 string = ""
            _ = v22
        }
        _ = fi19
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        var v24 uint8 = 197
        fmt.Println(v18, fi23)
        _ = fi23
        _ = v24
    }
    var v25 string = ("a" + "\n\t")
    var v26 bool = true
    var v27 float32 = v3
    for fi28 := 0; fi28 < 10; fi28++ {
        v29 := GenS3{F0: 1}
        for fi30 := 0; fi30 < 16; fi30++ {
            if (0 <= 127) {
                if (!(v4 <= v4)) {
                    v1[42] = false
                    v31 := GenS4{F0: 42, F1: 9223372036854775807, F2: 1}
                    delete(v1, 2147483647)
                    _ = v31
                }
                var v32 float64 = 0.0
                _ = v32
            }
            _ = fi30
        }
        _ = fi28
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
