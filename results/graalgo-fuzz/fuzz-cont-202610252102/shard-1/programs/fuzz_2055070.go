package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int }
type GenS2 struct { F0 int }
func genFn0(p0 int, p1 int32, p2 int16) bool {
    v7 := []int{(p0 & p0), -562, 127}
    if false {
        v8 := GenS1{F0: 18446744073709551615, F1: 1, F2: (^(p0))}
        var v9 int16 = (p2 << ((65535) & 63))
        _ = v8
        _ = v9
    }
    var v10 uint16 = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return true
}

func main() {
    var v1 int8 = 42
    var v2 int8 = (v1 << ((0) & 63))
    var v3 int = 42
    var v4 uint8 = 195
    v5 := []uint32{4294967295}
    v6 := make(map[int]int)
    v5 = append(v5, 4294967295)
    var v11 int = v3
    var v12 float32 = (float32((3.14159 - 1e10)) / 1e10)
    fmt.Println(v11, v4, v3)
    v11 = (v3 ^ v3)
    var v13 int8 = -92
    v14 := GenS1{F0: 18446744073709551615, F1: 963, F2: -93}
    v15 := GenS1{F0: 46, F1: uint(v4), F2: -9223372036854775808}
    if false {
        if (1 < 65535) {
            v4 += (v4 - 1)
        }
        v4 -= 42
        var v16 int = (v11 % (((v11 % ((4294967295) | 1))) | 1))
        _ = v16
    }
    v17 := GenS1{F0: 4294967295, F1: 1, F2: v11}
    if (float32((3.14159 + 0.1)) <= (0.0 / v12)) {
        if (42 != -334) {
            if true {
                for fi18 := 0; fi18 < 1; fi18++ {
                    v2 -= 42
                    v19 := GenS1{F0: 42, F1: 1, F2: v3}
                    _ = fi18
                    _ = v19
                }
                if true {
                    v3 = (v11 / (((v3 / ((-944) | 1))) | 1))
                    var v20 bool = (v3 <= (v11 / ((v11) | 1)))
                    v21 := GenS2{F0: 32767}
                    _ = v20
                    _ = v21
                }
            }
        }
        var v22 uint32 = 386
        for fi23 := 0; fi23 < 5; fi23++ {
            if (v3 > 128) {
                v4 += (v4 ^ 1)
                switch "a" {
                case "hello":
                    v3 = 42
                }
            } else {
                for fi24 := 0; fi24 < 13; fi24++ {
                    if (v4 > (v4 % ((v4) | 1))) {
                        v5 = append(v5, 1)
                    } else {
                        switch (("hello" + "hello") + "") {
                        case "a":
                            v13 -= (-(v13))
                        }
                        v2 = v1
                    }
                    v4 += v4
                    _ = fi24
                }
                v25 := GenS2{F0: -1}
                _ = v25
            }
            var v26 uint = 42
            _ = fi23
            _ = v26
        }
        _ = v22
    } else {
        v2 = 100
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
