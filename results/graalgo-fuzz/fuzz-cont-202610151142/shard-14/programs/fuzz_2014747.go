package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 string }
type GenS2 struct { F0 string; F1 uint64; F2 string }

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[uint32]uint32)
    var v3 int8 = 127
    v4 := make(map[uint32]string)
    var v5 string = "foo bar"
    v5 = (v5 + v5)
    v6 := GenS1{F0: (v3 << ((65535) & 63)), F1: (506 >= 2147483647), F2: v5}
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := []int{(fi7 / ((fi7) | 1))}
        _ = fi7
        _ = v8
    }
    fmt.Println(v5, v1)
    var v9 bool = false
    fmt.Println(v1, v5, v5)
    for fi10 := 0; fi10 < 2; fi10++ {
        fmt.Println(v1, v5, v3)
        for fi11 := 0; fi11 < 2; fi11++ {
            if (1.5 < (1e10 - 3.14159)) {
                if v9 {
                    var v12 float64 = 1e10
                    fmt.Println(v1, v5, v1)
                    if (42 == 255) {
                        v3 += (^(v3))
                    }
                    var v13 uint64 = (v1 % ((2) | 1))
                    _ = v12
                    _ = v13
                } else {
                    v9 = (v1 != v1)
                    if v9 {
                        v5 = v5
                        v1 = 0
                        v1 -= v1
                    } else {
                        var v14 uint32 = 0
                        fmt.Println(v5, v14)
                        v15 := GenS1{F0: int8(v14), F1: (true || (403 != 579)), F2: v5}
                        _ = v14
                        _ = v15
                    }
                }
                v16 := GenS1{F0: v3, F1: v9, F2: (((v5 + v5) + v5) + ("\n\t" + v5))}
                switch v5 {
                case "":
                    var v17 uint16 = 32767
                    _ = v17
                case "foo bar":
                    v9 = (18446744073709551615 == 65535)
                }
                _ = v16
            } else {
                v18 := []int{fi11}
                v9 = v9
                delete(v4, 42)
                _ = v18
            }
            var v19 uint16 = 785
            _ = fi11
            _ = v19
        }
        v20 := GenS1{F0: v3, F1: (42 <= 692), F2: v5}
        _ = fi10
        _ = v20
    }
    if v9 {
        v21 := []int{32767}
        _ = v21
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        switch -9223372036854775808 {
        case -335:
            v9 = (true && (42 < -1))
        case 4294967295:
            v23 := GenS2{F0: v5, F1: v1, F2: ("\n\t" + v5)}
            _ = v23
        }
        fmt.Println(v9)
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
