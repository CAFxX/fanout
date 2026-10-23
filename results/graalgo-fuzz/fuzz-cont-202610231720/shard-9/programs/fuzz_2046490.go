package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint32; F1 bool }

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = float32((0.0 - 1e10))
    var v3 float32 = v2
    v4 := make(map[uint32]int)
    var v5 int = 515
    var v6 uint8 = 42
    v7 := GenS1{F0: 4294967295, F1: ((("" + ("a" + "a")) + "\n\t") + "hello")}
    fmt.Println(v5)
    var v8 uint32 = 42
    if true {
        v3 = -2.25
        var v9 int = (v5 ^ (v5 + (v5 | (^(v5)))))
        _ = v9
    }
    v5 = v5
    _ = len(v1)
    v10 := GenS1{F0: uint(v6), F1: "a"}
    if (906 != 611) {
        v8 = (v8 >> ((390) & 63))
        v6 += v6
        if (0 == -128) {
            v1[("a" + ("foo bar" + "\n\t"))] = 42
            for fi11 := 0; fi11 < 16; fi11++ {
                v6 += 1
                for fi12 := 0; fi12 < 19; fi12++ {
                    var v13 float64 = 0.0
                    switch v5 {
                    case 252:
                        fmt.Println(v5, v8)
                    case -2147483648:
                        fmt.Println(v8, v5, v5)
                    case 1:
                        _ = len(v4)
                    }
                    _ = fi12
                    _ = v13
                }
                if false {
                    if (2 <= 2) {
                        fmt.Println(v6, v6, v2)
                    } else {
                        if (-775 == -607) {
                            v14 := []int{(-(v5)), 32767}
                            fmt.Println(v5, v8, v8)
                            v15 := GenS2{F0: (v8 / (((v8 - (v8 >> ((65535) & 63)))) | 1)), F1: ((fi11 % ((fi11) | 1)) < v5)}
                            _ = v14
                            _ = v15
                        } else {
                            if true {
                                var v16 string = "\n\t"
                                _ = v16
                            }
                            if false {
                                v17 := GenS2{F0: (^(v8)), F1: (24 != -122)}
                                if (4294967295 <= 65535) {
                                    for fi18 := 0; fi18 < 15; fi18++ {
                                        v2 = (3.14159 * 0.0)
                                        var v19 int = fi11
                                        var v20 int32 = -2147483648
                                        _ = fi18
                                        _ = v19
                                        _ = v20
                                    }
                                    v21 := []int{v5}
                                    _ = v21
                                } else {
                                    v8 += 0
                                    var v22 uint = 255
                                    _ = v22
                                }
                                _ = v17
                            }
                            if (!(42 < 757)) {
                                fmt.Println(v5, v3, v5)
                            } else {
                                var v23 int32 = 0
                                for fi24 := 0; fi24 < 1; fi24++ {
                                    switch (fi24 >> ((441) & 63)) {
                                    case -9223372036854775808:
                                        v3 -= v2
                                    }
                                    v23 -= (v23 | (v23 / ((v23) | 1)))
                                    _ = fi24
                                }
                                if (true || true) {
                                    fmt.Println(v2)
                                    v8 += v8
                                }
                                _ = v23
                            }
                        }
                        fmt.Println(v5, v8, v2)
                        v25 := []int{32767}
                        _ = v25
                    }
                    v2 -= v2
                    v26 := GenS1{F0: 0, F1: "\n\t"}
                    _ = v26
                } else {
                    var v27 int32 = -1
                    _ = v27
                }
                _ = fi11
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
