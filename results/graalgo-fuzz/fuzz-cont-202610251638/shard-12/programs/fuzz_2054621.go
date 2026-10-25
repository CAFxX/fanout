package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int16; F2 uint64 }

func main() {
    v1 := []string{"\n\t", "\n\t", "a"}
    var v2 float64 = 1e10
    var v3 int64 = -9223372036854775808
    v4 := []int{-9223372036854775808}
    var v5 int32 = 793
    var v6 int32 = (v5 * v5)
    if false {
        v6 += v6
    } else {
        var v7 uint32 = 2147483648
        _ = v7
    }
    var v8 float64 = (float64(0.1) * 1e10)
    switch (("a" + "foo bar") + ("a" + "foo bar")) {
    case "foo bar":
        if (!true) {
            v9 := GenS1{F0: 2147483648}
            for fi10 := 0; fi10 < 16; fi10++ {
                if (!true) {
                    switch "" {
                    case "a":
                        var v11 int8 = 14
                        _ = v11
                    case "\n\t":
                        for fi12 := 0; fi12 < 19; fi12++ {
                            if false {
                                var v13 int64 = 1
                                v2 = (1.5 + v8)
                                if len(v4) > 0 {
                                    v4[0] = 1
                                }
                                if (!true) {
                                    v5 -= -1
                                }
                                _ = v13
                            } else {
                                v14 := GenS1{F0: 1}
                                var v15 int32 = (-(v5))
                                var v16 uint16 = 255
                                _ = v14
                                _ = v15
                                _ = v16
                            }
                            if true {
                                if (1 == 1) {
                                    switch ((("a" + (("foo bar" + "") + "")) + "hello") + "") {
                                    case "a":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    default:
                                        var v17 float32 = (0.0 - 1.5)
                                        _ = v17
                                    }
                                    v6 += (v5 - v5)
                                    if true {
                                        v1 = append(v1, ("\n\t" + (("foo bar" + ("\n\t" + "hello")) + "a")))
                                    } else {
                                        var v18 uint32 = 4294967295
                                        var v19 uint16 = 360
                                        _ = v18
                                        _ = v19
                                    }
                                }
                                var v20 int32 = v6
                                var v21 uint = 829
                                _ = v20
                                _ = v21
                            }
                            _ = fi12
                        }
                    }
                    v22 := GenS2{F0: fi10, F1: 0, F2: 124}
                    _ = v22
                }
                if len(v4) > 0 {
                    v4[0] = 4294967295
                }
                _ = fi10
            }
            _ = v9
        } else {
            v3 += (v3 | 0)
            var v23 uint = 4294967295
            _ = v23
        }
    case "\n\t":
        var v24 uint32 = 2
        _ = v24
    case "hello":
        if (v8 > (float64((1.5 + (-2.25 * 0.1))) + v8)) {
            var v25 uint64 = 4294967295
            var v26 uint64 = v25
            var v27 uint32 = 795
            _ = v25
            _ = v26
            _ = v27
        }
    }
    v3 += 42
    v5 = v6
    if true {
        var v28 uint = 65535
        _ = v28
    }
    var v29 int32 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
