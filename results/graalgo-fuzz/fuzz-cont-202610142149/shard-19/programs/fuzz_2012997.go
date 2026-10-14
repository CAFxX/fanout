package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int16 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int8) uint64 {
    var v6 uint64 = 18446744073709551615
    v7 := GenS1{F0: (p0 + p0), F1: -2147483648, F2: -57}
    _ = p0
    _ = v6
    _ = v7
    return (v6 / (((v6 % ((4294967295) | 1))) | 1))
}
func genFn1(p0 uint32, p1 int, p2 uint16) int {
    for fi8 := 0; fi8 < 14; fi8++ {
        p2 -= p2
        _ = fi8
    }
    v9 := GenS2{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return p1
}

func main() {
    var v1 uint = 538
    var v2 uint32 = 0
    var v3 bool = false
    var v4 uint16 = 65535
    v5 := make(map[string]int)
    var v10 int64 = 42
    fmt.Println(v3)
    var v11 uint32 = (v2 / (((v2 - v2)) | 1))
    if true {
        v5["hello"] = 32768
        var v12 uint16 = (v4 >> (((v1 & v1)) & 63))
        _ = v12
    }
    v1 -= v1
    if v3 {
        for fi13 := 0; fi13 < 8; fi13++ {
            v14 := GenS1{F0: 36, F1: 2147483647, F2: -735}
            fmt.Println(v10, v2, v1)
            fmt.Println(v1, v10, v4)
            _ = fi13
            _ = v14
        }
        if (!((v1 * (v1 * v1)) != (^(v1)))) {
            v15 := []int{160, 32767, -593}
            var v16 bool = v3
            v1 += v1
            _ = v15
            _ = v16
        } else {
            fmt.Println(v10)
            v10 += v10
            if v3 {
                var v17 int16 = 0
                var v18 string = ("foo bar" + "a")
                _ = v17
                _ = v18
            } else {
                var v19 uint32 = 2147483648
                for fi20 := 0; fi20 < 9; fi20++ {
                    if false {
                        switch fi20 {
                        case -2147483648:
                            var v21 float64 = 3.14159
                            _ = v21
                        case 155:
                            for fi22 := 0; fi22 < 10; fi22++ {
                                _ = len(v5)
                                _ = fi22
                            }
                        }
                        v23 := GenS1{F0: 0, F1: 42, F2: 37}
                        for fi24 := 0; fi24 < 16; fi24++ {
                            v3 = false
                            _ = fi24
                        }
                        _ = v23
                    }
                    var v25 uint32 = (v2 << (((v1 % ((561) | 1))) & 63))
                    v26 := GenS1{F0: 23, F1: 2147483647, F2: -32768}
                    _ = fi20
                    _ = v25
                    _ = v26
                }
                if v3 {
                    v27 := GenS2{F0: 2}
                    delete(v5, ("\n\t" + "hello"))
                    _ = v27
                }
                _ = v19
            }
        }
    } else {
        var v28 uint16 = (v4 >> (((v1 - (v1 % ((65535) | 1)))) & 63))
        _ = v28
    }
    var v29 uint = (v1 << ((v1) & 63))
    var v30 bool = (9223372036854775807 >= v10)
    v4 = (v4 | (v4 | 42))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
