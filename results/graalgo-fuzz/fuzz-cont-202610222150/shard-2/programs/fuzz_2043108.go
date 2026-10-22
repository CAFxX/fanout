package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int64 }
type GenS2 struct { F0 int32; F1 uint32; F2 string }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 int; F1 uint; F2 int64 }
type GenS7 struct { F0 uint16; F1 int64; F2 int16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int8; F1 uint16; F2 int64 }
type GenS10 struct { F0 int16; F1 bool }
func genFn0(p0 uint16, p1 bool) uint16 {
    p1 = (4294967295 != 42)
    if p1 {
        v6 := GenS1{F0: 352, F1: 0, F2: 1}
        switch -725 {
        case 0:
            fmt.Println(p1, p1)
        }
        var v7 uint32 = 2147483647
        _ = v6
        _ = v7
    }
    var v8 uint32 = 157
    v9 := GenS2{F0: 42, F1: (^(v8)), F2: "\n\t"}
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := GenS2{F0: 1, F1: (v8 * (v8 / ((v8) | 1))), F2: (("\n\t" + ("" + "\n\t")) + "foo bar")}
        p0 += p0
        if false {
            v12 := GenS3{F0: -9223372036854775808, F1: uint8(v8)}
            _ = v12
        } else {
            v8 += v8
            for fi13 := 0; fi13 < 1; fi13++ {
                var v14 uint64 = 18446744073709551615
                switch (fi10 | fi10) {
                case -9223372036854775808:
                    var v15 int64 = 1
                    _ = v15
                default:
                    v8 = 906
                }
                var v16 string = "\n\t"
                _ = fi13
                _ = v14
                _ = v16
            }
            v17 := []int{(^(fi10)), (fi10 * fi10), 127, (fi10 / ((fi10) | 1))}
            _ = v17
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return p0
}
func genFn1(p0 bool) uint8 {
    fmt.Println(p0)
    if (681 >= 1) {
        var v18 uint8 = 206
        _ = v18
    }
    _ = p0
    return 1
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint32 = 42
    v3 := make(map[int]int)
    var v4 int64 = v1
    v5 := []uint32{v2, v2, 42, v2}
    var v19 int16 = 0
    v4 -= (^(v4))
    switch 32767 {
    case -985:
        var v20 int32 = -996
        _ = v20
    case -790:
        switch ((("foo bar" + "foo bar") + "") + ((("hello" + ("a" + "")) + "hello") + "foo bar")) {
        case "a":
            var v21 bool = false
            _ = v21
        case "foo bar":
            fmt.Println(v4, v19)
        }
    case 841:
        v4 = v4
    default:
        fmt.Println(v4, v1, v2)
    }
    if len(v5) > 0 {
        v5[0] = (v2 + 0)
    }
    switch "a" {
    case "a":
        fmt.Println(v4)
    default:
        if true {
            var v22 float64 = (0.1 + 0.0)
            _ = v22
        } else {
            v3[-9223372036854775808] = 751
            if false {
                var v23 int16 = (v19 & v19)
                v23 = v23
                switch "foo bar" {
                case "foo bar":
                    v23 = v19
                case "a":
                    v24 := GenS3{F0: v4, F1: 128}
                    _ = v24
                }
                _ = v23
            } else {
                v25 := GenS1{F0: 644, F1: 128, F2: -868}
                var v26 float32 = (1.5 + -2.25)
                v27 := GenS4{F0: 255, F1: -1}
                _ = v25
                _ = v26
                _ = v27
            }
        }
    }
    v4 = v4
    var v28 int = 4
    if false {
        for fi29 := 0; fi29 < 2; fi29++ {
            var v30 int = v28
            _ = fi29
            _ = v30
        }
        var v31 uint32 = (v2 & v2)
        if ((42 <= 42) || (false && true)) {
            var v32 int = v28
            fmt.Println(v2, v32)
            _ = v32
        }
        _ = v31
    }
    v1 -= (v4 << ((42) & 63))
    if true {
        if ("hello" >= (("" + "") + ((((("\n\t" + ("a" + "")) + ("" + (("hello" + "") + ""))) + ("foo bar" + "foo bar")) + "\n\t") + ("" + "hello")))) {
            fmt.Println(v19, v28, v19)
        }
    }
    v4 -= (v4 | v4)
    delete(v3, v28)
    _ = len(v3)
    v33 := GenS5{F0: v2, F1: 42}
    for fi34 := 0; fi34 < 12; fi34++ {
        for fi35 := 0; fi35 < 12; fi35++ {
            v36 := GenS6{F0: v28, F1: 42, F2: v4}
            v37 := GenS7{F0: genFn0(32767, (!(0 != 117))), F1: (-(v1)), F2: v19}
            if (false || (!(float32(1e10) == -2.25))) {
                v38 := GenS8{F0: v19}
                v39 := GenS9{F0: -1, F1: uint16(v2), F2: (v1 << ((4294967295) & 63))}
                _ = v38
                _ = v39
            }
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = fi34
    }
    v5 = append(v5, v2)
    v5 = append(v5, 362)
    for fi40 := 0; fi40 < 2; fi40++ {
        if (true || false) {
            v4 += (-(v4))
        }
        switch fi40 {
        case 644:
            var v41 int = (-(v28))
            _ = v41
        case -205:
            fmt.Println(v19, v4, fi40)
        case 62:
            v42 := GenS10{F0: (v19 % ((42) | 1)), F1: true}
            _ = v42
        }
        _ = fi40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
