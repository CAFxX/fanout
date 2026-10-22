package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int }
type GenS2 struct { F0 int64; F1 uint16; F2 uint64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    v3 := make(map[string]int)
    v4 := []string{(("hello" + "foo bar") + "\n\t"), ((((("" + "") + "a") + "hello") + ("" + "foo bar")) + "hello")}
    v5 := []uint32{583, 65535, 627, 1}
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 int32 = 42
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 42, F1: 69, F2: 952}
    var v9 int32 = 2147483647
    _ = len(v4)
    fmt.Println(v9)
    if (true || false) {
        v4 = append(v4, ("a" + "a"))
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 string = "a"
            var v12 float64 = -2.25
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    v13 := GenS1{F0: 742, F1: 2, F2: 953}
    var v14 uint = 926
    if (1 >= -1) {
        switch 0 {
        case 32767:
            for fi15 := 0; fi15 < 9; fi15++ {
                var v16 int64 = 5
                _ = fi15
                _ = v16
            }
        }
        v14 += 255
        v17 := GenS2{F0: -9223372036854775808, F1: 462, F2: 18446744073709551615}
        _ = v17
    }
    if (!false) {
        for fi18 := 0; fi18 < 4; fi18++ {
            for fi19 := 0; fi19 < 15; fi19++ {
                v20 := GenS1{F0: 1, F1: 127, F2: fi19}
                if true {
                    switch (fi19 + (fi19 - (fi19 + (fi18 / (((fi18 << ((2) & 63))) | 1))))) {
                    case 2147483647:
                        var v21 uint64 = 0
                        _ = v21
                    case 42:
                        v9 += v9
                    }
                    if len(v4) > 0 {
                        v4[0] = "foo bar"
                    }
                    var v22 int32 = v9
                    _ = v22
                }
                if (42 != 0) {
                    v14 = (v14 * (^(v14)))
                    if (3.14159 != float64(0.0)) {
                        v23 := GenS2{F0: 9223372036854775807, F1: 65535, F2: 42}
                        v9 = v9
                        v9 = (v9 + (v9 ^ (v9 >> ((v14) & 63))))
                        _ = v23
                    }
                }
                _ = fi19
                _ = v20
            }
            v24 := GenS2{F0: 42, F1: 32767, F2: uint64(v14)}
            _ = fi18
            _ = v24
        }
        for fi25 := 0; fi25 < 20; fi25++ {
            if false {
                var v26 int64 = -1
                _ = v26
            } else {
                for fi27 := 0; fi27 < 13; fi27++ {
                    if true {
                        v14 -= 542
                        v28 := GenS1{F0: 215, F1: 127, F2: -1}
                        _ = v28
                    }
                    _ = fi27
                }
                v29 := GenS2{F0: 0, F1: 65535, F2: 42}
                for fi30 := 0; fi30 < 16; fi30++ {
                    fmt.Println(v14, v9, fi30)
                    v31 := GenS1{F0: 42, F1: 2, F2: (fi25 / (((fi30 - int(v14))) | 1))}
                    _ = fi30
                    _ = v31
                }
                _ = v29
            }
            _ = fi25
        }
        for fi32 := 0; fi32 < 11; fi32++ {
            var v33 float64 = (3.14159 + (0.0 / (0.1 / 1.0)))
            _ = fi32
            _ = v33
        }
    }
    var v34 uint16 = 974
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
