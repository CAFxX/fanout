package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 uint8; F1 uint32; F2 int32 }
func genFn0() int16 {
    var v4 uint32 = 310
    v4 = v4
    v4 = v4
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 bool = (!((!false) && true))
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := []int{fi7, fi7, 220}
        if true {
            v4 += 1
        }
        _ = fi7
        _ = v8
    }
    _ = v4
    return 1000
}

func main() {
    v1 := []string{"\n\t", "\n\t"}
    var v2 uint16 = 2
    var v3 uint = 1
    for fi9 := 0; fi9 < 8; fi9++ {
        _ = len(v1)
        v10 := GenS1{F0: (v2 << (((^(v3))) & 63)), F1: false, F2: 143}
        _ = fi9
        _ = v10
    }
    var v11 int64 = 9223372036854775807
    for fi12 := 0; fi12 < 1; fi12++ {
        v2 -= v2
        if (v3 < 65535) {
            if (float64(0.0) < (3.14159 / -2.25)) {
                _ = len(v1)
            }
            v2 = 434
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi12
    }
    if true {
        for fi13 := 0; fi13 < 12; fi13++ {
            if (417 >= 911) {
                var v14 int32 = 928
                if (v14 <= (v14 ^ (v14 ^ (v14 & -241)))) {
                    v15 := GenS1{F0: (v2 >> ((v3) & 63)), F1: (false || ("a" <= "hello")), F2: 2147483648}
                    for fi16 := 0; fi16 < 17; fi16++ {
                        v2 = uint16(v3)
                        _ = len(v1)
                        v17 := GenS1{F0: v2, F1: (-1 < (v11 | (v11 << (((v3 | 65535)) & 63)))), F2: 65535}
                        _ = fi16
                        _ = v17
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v15
                }
                if true {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                _ = v14
            } else {
                v3 = 255
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            v18 := GenS2{F0: (v3 + (^(v3))), F1: v3}
            _ = fi13
            _ = v18
        }
        for fi19 := 0; fi19 < 19; fi19++ {
            fmt.Println(v11)
            _ = fi19
        }
        v11 -= v11
    } else {
        if (1.5 == float64(0.0)) {
            var v20 uint64 = 18446744073709551615
            v2 = v2
            for fi21 := 0; fi21 < 12; fi21++ {
                fmt.Println(v20, v2)
                _ = fi21
            }
            _ = v20
        }
        v22 := GenS1{F0: v2, F1: (!(1 <= 100)), F2: 0}
        if false {
            var v23 string = ("" + "a")
            _ = v23
        } else {
            for fi24 := 0; fi24 < 17; fi24++ {
                for fi25 := 0; fi25 < 12; fi25++ {
                    v26 := GenS3{F0: true, F1: true}
                    _ = fi25
                    _ = v26
                }
                _ = fi24
            }
            v2 = v2
        }
        _ = v22
    }
    var v27 int16 = 42
    var v28 int64 = (v11 >> (((v3 >> ((v3) & 63))) & 63))
    for fi29 := 0; fi29 < 1; fi29++ {
        var v30 int = (fi29 % (((fi29 % ((4294967295) | 1))) | 1))
        v28 -= v28
        _ = fi29
        _ = v30
    }
    v31 := GenS4{F0: 128, F1: 65535, F2: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v31
}
