package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint; F1 int64; F2 uint8 }
type GenS3 struct { F0 bool; F1 string }
func genFn0() int {
    v6 := GenS1{F0: "hello", F1: 32768}
    var v7 uint32 = 343
    _ = v6
    _ = v7
    return 318
}

func main() {
    var v1 int16 = 42
    var v2 int8 = 127
    var v3 int = 32767
    v4 := []string{"hello"}
    v5 := []uint32{310, 2147483648}
    _ = len(v5)
    for fi8 := 0; fi8 < 15; fi8++ {
        v1 = 1000
        _ = fi8
    }
    _ = len(v5)
    if true {
        v9 := GenS1{F0: ((("hello" + "") + "foo bar") + ("hello" + ("" + "a"))), F1: (v3 % ((genFn0()) | 1))}
        var v10 int8 = v2
        _ = v10
        _ = v9
    } else {
        v3 -= v3
    }
    v11 := GenS1{F0: "\n\t", F1: (v3 << ((65535) & 63))}
    v2 = -91
    if ((!(84 < 0)) && false) {
        var v12 int8 = v2
        v13 := GenS1{F0: "a", F1: -9223372036854775808}
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 int8 = (v2 / ((v12) | 1))
            v16 := GenS1{F0: "a", F1: fi14}
            v17 := GenS2{F0: 175, F1: -48, F2: 42}
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        _ = v12
        _ = v13
    } else {
        var v18 bool = false
        _ = v18
    }
    fmt.Println(v1)
    var v19 int16 = v1
    if (!((true || true) && true)) {
        if (((!(-1 >= (v19 >> ((874) & 63)))) || false) || false) {
            v20 := GenS3{F0: true, F1: ""}
            var v21 int64 = 1
            _ = v20
            _ = v21
        }
    }
    fmt.Println(v3, v2)
    if true {
        for fi22 := 0; fi22 < 4; fi22++ {
            for fi23 := 0; fi23 < 9; fi23++ {
                _ = len(v5)
                var v24 int32 = -2147483648
                fmt.Println(fi23, fi22)
                _ = fi23
                _ = v24
            }
            v3 += 283
            v1 = v19
            _ = fi22
        }
        var v25 int8 = (v2 << ((640) & 63))
        _ = v25
    }
    v2 += (v2 / ((100) | 1))
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 float64 = 0.1
        _ = fi26
        _ = v27
    }
    v19 = 1000
    if true {
        var v28 uint32 = 0
        _ = len(v4)
        for fi29 := 0; fi29 < 6; fi29++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v30 int64 = 1
            _ = fi29
            _ = v30
        }
        _ = v28
    } else {
        v31 := GenS2{F0: 65535, F1: 1, F2: 127}
        _ = v31
    }
    fmt.Println(v3, v19)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
