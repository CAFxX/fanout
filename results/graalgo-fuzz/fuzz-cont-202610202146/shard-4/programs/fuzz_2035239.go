package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 string }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int; F1 bool; F2 bool }
type GenS4 struct { F0 uint64; F1 int; F2 int }
func genFn0(p0 int32, p1 int32, p2 bool) int64 {
    var v4 int = 1
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := GenS1{F0: 504, F1: (fi5 ^ fi5), F2: ("hello" + "hello")}
        p2 = true
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        if p2 {
            fmt.Println(v4, p0)
            p0 -= -143
        } else {
            v8 := []int{}
            fmt.Println(p0, p0)
            v9 := GenS1{F0: 398, F1: (v4 * fi7), F2: "foo bar"}
            _ = v8
            _ = v9
        }
        v10 := []int{(v4 << ((42) & 63)), fi7, (v4 & (v4 - v4))}
        var v11 int32 = (p1 & (p1 / ((p1) | 1)))
        _ = fi7
        _ = v10
        _ = v11
    }
    p0 = (p1 + p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 450
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = 1e10
    var v3 int = 127
    var v12 int8 = 100
    v13 := GenS1{F0: 0, F1: v3, F2: "hello"}
    var v14 uint64 = 18446744073709551615
    v15 := GenS1{F0: 313, F1: (v3 ^ v3), F2: "a"}
    v16 := GenS1{F0: 42, F1: (v3 * (v3 >> ((1) & 63))), F2: "\n\t"}
    v3 += (v3 >> ((267) & 63))
    switch (v3 | 32768) {
    case 0:
        v17 := GenS2{F0: v12, F1: ((^(v14)) < v14)}
        _ = v17
    case 32768:
        var v18 float32 = 0.0
        _ = v18
    case 597:
        for fi19 := 0; fi19 < 13; fi19++ {
            v20 := GenS2{F0: (v12 >> ((4294967295) & 63)), F1: (2 > 1)}
            _ = fi19
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        v22 := []int{(fi21 % (((fi21 & fi21)) | 1)), (v3 / ((v3) | 1)), (v3 + 399), v3}
        _ = fi21
        _ = v22
    }
    v14 = v14
    var v23 int16 = 1
    var v24 uint = 2
    for fi25 := 0; fi25 < 8; fi25++ {
        v1[743] = true
        _ = fi25
    }
    if (v3 < v3) {
        var v26 int64 = 42
        _ = v26
    }
    for fi27 := 0; fi27 < 18; fi27++ {
        var v28 int32 = int32(v12)
        var v29 int8 = v12
        fmt.Println(v29, v14, v2)
        _ = fi27
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        fmt.Println(v2, v24)
        if true {
            v31 := GenS3{F0: (^(v3)), F1: (true && true), F2: (true && (1 != 508))}
            v32 := GenS4{F0: v14, F1: 2147483647, F2: -9223372036854775808}
            _ = v31
            _ = v32
        }
        _ = fi30
    }
    v33 := GenS2{F0: (v12 - -59), F1: false}
    for fi34 := 0; fi34 < 11; fi34++ {
        var v35 float64 = -2.25
        if true {
            v36 := []int{fi34, -258, (v3 | fi34), v3}
            var v37 uint16 = 1
            _ = v36
            _ = v37
        }
        v38 := GenS3{F0: (fi34 << ((443) & 63)), F1: (true || true), F2: ((true && (fi34 <= v3)) || true)}
        _ = fi34
        _ = v35
        _ = v38
    }
    for fi39 := 0; fi39 < 14; fi39++ {
        if (9223372036854775807 <= genFn0(42, -727, (168 >= 0))) {
            v2 += ((v2 * v2) - v2)
        }
        v40 := v1[688]
        _ = fi39
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v33
}
