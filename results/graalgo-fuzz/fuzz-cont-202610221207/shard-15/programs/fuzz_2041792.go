package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 int16; F2 bool }
func genFn0(p0 int32, p1 bool, p2 bool) uint8 {
    var v6 int = 42
    p0 -= p0
    v6 = 925
    switch (v6 >> ((534) & 63)) {
    case 419:
        var v7 uint8 = 255
        _ = v7
    case 4294967295:
        v8 := GenS1{F0: -9223372036854775808, F1: "\n\t"}
        _ = v8
    default:
        if p2 {
            if false {
                for fi9 := 0; fi9 < 13; fi9++ {
                    for fi10 := 0; fi10 < 11; fi10++ {
                        if p2 {
                            if ((v6 - 32767) >= v6) {
                                v11 := GenS1{F0: 1, F1: "a"}
                                if p2 {
                                    p1 = p1
                                    v12 := GenS1{F0: 42, F1: ("hello" + "")}
                                    switch ((("hello" + "\n\t") + "a") + ("foo bar" + "hello")) {
                                    case "a":
                                        v13 := GenS1{F0: -835, F1: ("foo bar" + ((("" + ("hello" + ("" + "foo bar"))) + "hello") + "foo bar"))}
                                        _ = v13
                                    case "hello":
                                        v14 := []int{0, (fi10 ^ fi9), (fi10 / ((v6) | 1)), (fi10 * (fi9 ^ (-(fi10))))}
                                        _ = v14
                                    }
                                    var v15 uint8 = 255
                                    v16 := []int{(fi9 & (fi10 - int(p0))), 2147483647, (^(v6)), (fi10 * fi10)}
                                    _ = v12
                                    _ = v15
                                    _ = v16
                                } else {
                                    if ((((float32(0.1) / (1e10 / 0.1)) * -2.25) < (0.0 / 1.5)) || true) {
                                        if (true || (65 >= 280)) {
                                            var v17 string = ((("\n\t" + ("foo bar" + "a")) + "a") + "a")
                                            _ = v17
                                        }
                                        fmt.Println(v6)
                                    } else {
                                        v18 := GenS1{F0: 9223372036854775807, F1: "hello"}
                                        fmt.Println(fi9, p1)
                                        p0 = p0
                                        _ = v18
                                    }
                                }
                                fmt.Println(p2)
                                _ = v11
                            } else {
                                p1 = p2
                            }
                            var v19 uint8 = 194
                            _ = v19
                        }
                        var v20 int8 = 127
                        v6 = fi10
                        _ = fi10
                        _ = v20
                    }
                    _ = fi9
                }
            }
            fmt.Println(p2, p1, p2)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 255
}
func genFn1() string {
    var v21 int32 = 926
    if (true && true) {
        v22 := GenS2{F0: 32767}
        _ = v22
    } else {
        v21 -= (v21 % ((v21) | 1))
        v23 := []int{32768, 9223372036854775807}
        if len(v23) > 0 {
            v23[0] = -704
        }
        _ = v23
    }
    var v24 int = -2147483648
    _ = v21
    _ = v24
    return ("hello" + "foo bar")
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    var v3 int8 = 6
    var v4 uint = 433
    v5 := []int{270}
    fmt.Println(v4)
    var v25 int64 = -9223372036854775808
    v26 := GenS3{F0: v3, F1: 1000, F2: ((v3 << ((v4) & 63)) > (-(v3)))}
    if (uint64(v4) >= 2) {
        var v27 int64 = 0
        _ = v27
    } else {
        v3 = 1
        var v28 int32 = -1
        _ = v28
    }
    for fi29 := 0; fi29 < 6; fi29++ {
        switch ("hello" + ("a" + (("hello" + "a") + "\n\t"))) {
        case "hello":
            fmt.Println(v3, v25, fi29)
        case "a":
            var v30 int64 = v25
            _ = v30
        }
        var v31 int8 = v3
        _ = fi29
        _ = v31
    }
    v4 = v4
    v25 -= -230
    v25 -= v25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
