package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
func genFn0() bool {
    var v7 float64 = (1.5 + float64(0.0))
    v8 := GenS1{F0: 1, F1: 0}
    var v9 bool = false
    if v9 {
        var v10 uint = 622
        _ = v10
    } else {
        if (42 > 65535) {
            if true {
                if true {
                    v11 := GenS1{F0: 18446744073709551615, F1: -1}
                    _ = v11
                } else {
                    v9 = (65535 >= 975)
                    if v9 {
                        v9 = false
                        for fi12 := 0; fi12 < 7; fi12++ {
                            v9 = v9
                            v13 := []int{}
                            if v9 {
                                switch (fi12 % ((-7) | 1)) {
                                case 2147483647:
                                    var v14 int32 = 1
                                    _ = v14
                                }
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                                var v15 uint32 = 2147483648
                                _ = v15
                            }
                            _ = fi12
                            _ = v13
                        }
                    } else {
                        var v16 int = -536
                        fmt.Println(v16, v9, v7)
                        var v17 uint8 = 207
                        _ = v16
                        _ = v17
                    }
                }
                v9 = (255 <= 42)
            } else {
                fmt.Println(v9)
            }
            v18 := []int{-2147483648, 2147483647, 127, -1}
            _ = v18
        }
        if (v9 || (v9 && false)) {
            var v19 uint64 = 996
            var v20 float32 = 0.1
            _ = v19
            _ = v20
        } else {
            var v21 uint64 = 2
            v21 = v21
            _ = v21
        }
    }
    _ = v7
    _ = v8
    _ = v9
    return v9
}

func main() {
    v1 := make(map[int]string)
    var v2 float64 = 0.1
    var v3 uint8 = 189
    v4 := make(map[string]bool)
    var v5 uint8 = 163
    v6 := []string{"", "a", "hello"}
    var v22 uint16 = 2
    v2 = v2
    if false {
        v1[9223372036854775807] = ("" + "foo bar")
    } else {
        _ = len(v1)
        fmt.Println(v22, v2, v22)
        if (true && (true || false)) {
            var v23 int32 = -1
            _ = v23
        } else {
            v24 := GenS1{F0: 18446744073709551615, F1: 42}
            if (false || (true || true)) {
                for fi25 := 0; fi25 < 1; fi25++ {
                    v5 -= v3
                    v26 := GenS1{F0: 251, F1: 1}
                    _ = fi25
                    _ = v26
                }
                var v27 int32 = -195
                v28 := GenS1{F0: 4294967295, F1: 1}
                _ = v27
                _ = v28
            }
            _ = v24
        }
    }
    v22 = (v22 % (((v22 >> ((65535) & 63))) | 1))
    var v29 int32 = 42
    v2 = v2
    fmt.Println(v22)
    fmt.Println(v2)
    var v30 float64 = 0.1
    delete(v1, 2147483647)
    switch -980 {
    case -416:
        var v31 int8 = 100
        _ = v31
    case 1:
        v3 = v3
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
