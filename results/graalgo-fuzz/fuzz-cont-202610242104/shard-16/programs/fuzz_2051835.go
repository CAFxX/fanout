package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 bool; F1 int8; F2 int64 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 uint64; F1 uint64; F2 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32; F1 string; F2 uint16 }
func genFn0(p0 int64) uint64 {
    if (int32(p0) >= 2147483647) {
        v7 := GenS1{F0: 32767, F1: (p0 * (p0 ^ 42))}
        _ = v7
    }
    fmt.Println(p0)
    switch -2147483648 {
    case 127:
        var v8 int8 = 0
        _ = v8
    case 1:
        if ("" > ("a" + "foo bar")) {
            p0 = 1
            p0 = (p0 << ((2) & 63))
        }
    }
    _ = p0
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 42
    var v2 int = 32768
    var v3 uint = 93
    v4 := make(map[string]int)
    v5 := make(map[int]string)
    v6 := []int{v2}
    v6 = append(v6, 32768)
    fmt.Println(v2)
    var v9 int64 = -9223372036854775808
    v3 = (^(v3))
    var v10 int32 = -758
    for fi11 := 0; fi11 < 5; fi11++ {
        switch 127 {
        case 128:
            var v12 uint16 = 388
            _ = v12
        case -808:
            if (-1 >= 42) {
                switch 379 {
                case -417:
                    switch "hello" {
                    case "a":
                        var v13 int8 = -1
                        _ = v13
                    }
                default:
                    v14 := GenS2{F0: (v1 != (v1 % ((v1) | 1))), F1: -32, F2: 400}
                    _ = v14
                }
                if len(v6) > 0 {
                    v6[0] = -2147483648
                }
                v1 = (v1 + (v1 << (((v3 >> ((v3) & 63))) & 63)))
            } else {
                v15 := GenS3{F0: int16(v9), F1: v2}
                if (3.14159 >= float32(-2.25)) {
                    fmt.Println(v3, v9, v9)
                    v9 -= (v9 & -1)
                    var v16 int16 = -32768
                    _ = v16
                }
                for fi17 := 0; fi17 < 17; fi17++ {
                    if (!false) {
                        if (true || (!false)) {
                            switch fi11 {
                            case 32767:
                                switch 9223372036854775807 {
                                case -900:
                                    v2 = fi11
                                case 2147483647:
                                    v1 = v1
                                }
                            case 128:
                                if true {
                                    fmt.Println(fi11, v2, v3)
                                    var v18 int16 = -32768
                                    _ = v18
                                }
                            case 42:
                                var v19 uint = v3
                                _ = v19
                            }
                            if false {
                                v2 -= 73
                                v20 := v4["hello"]
                                var v21 float32 = float32((0.0 - 0.1))
                                v22 := v4[""]
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                            var v23 int64 = (-(v9))
                            _ = v23
                        } else {
                            v24 := v4["hello"]
                            v25 := GenS4{F0: 0, F1: 241, F2: (v10 | (v10 * (-(v10))))}
                            var v26 bool = (247 < (v10 * v10))
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        v27 := GenS5{F0: false}
                        _ = v27
                    } else {
                        v1 += 1
                    }
                    v28 := GenS6{F0: v1, F1: (("a" + "hello") + ((("a" + "") + "foo bar") + (("hello" + "foo bar") + "hello"))), F2: 32767}
                    var v29 bool = true
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = fi17
                    _ = v28
                    _ = v29
                }
                _ = v15
            }
        }
        _ = fi11
    }
    v3 = 255
    if true {
        for fi30 := 0; fi30 < 15; fi30++ {
            for fi31 := 0; fi31 < 5; fi31++ {
                v3 -= (v3 / ((v3) | 1))
                switch (v2 & 581) {
                case 127:
                    switch -223 {
                    case 127:
                        var v32 int8 = -1
                        _ = v32
                    case 117:
                        delete(v4, (("foo bar" + "foo bar") + ("" + (("hello" + "hello") + ((("\n\t" + "") + "hello") + ("a" + ""))))))
                    case 9223372036854775807:
                        v33 := GenS4{F0: 2, F1: genFn0((-(v9))), F2: -2147483648}
                        _ = v33
                    }
                    var v34 int = fi30
                    _ = v34
                case -9223372036854775808:
                    switch "foo bar" {
                    case "a":
                        var v35 int32 = (v10 * v10)
                        _ = v35
                    case "foo bar":
                        fmt.Println(v10, v9, v3)
                    case "hello":
                        v2 += -9223372036854775808
                    default:
                        _ = len(v6)
                    }
                }
                v3 += v3
                _ = fi31
            }
            _ = fi30
        }
    }
    for fi36 := 0; fi36 < 5; fi36++ {
        v3 -= 4294967295
        v37 := GenS6{F0: 42, F1: "a", F2: 517}
        var v38 int8 = 43
        _ = fi36
        _ = v37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
