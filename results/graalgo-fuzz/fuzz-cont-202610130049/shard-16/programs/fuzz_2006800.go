package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]uint32)
    var v3 bool = ((1e10 - 3.14159) != 0.0)
    for fi4 := 0; fi4 < 7; fi4++ {
        if ((v3 && (false && (!false))) || false) {
            var v5 uint = 860
            v6 := GenS1{F0: 127, F1: 41}
            _ = v5
            _ = v6
        } else {
            delete(v2, (-(fi4)))
            for fi7 := 0; fi7 < 13; fi7++ {
                v8 := GenS2{F0: 1}
                var v9 bool = false
                var v10 int16 = 1
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
            v3 = v3
        }
        _ = fi4
    }
    if v3 {
        v11 := GenS2{F0: 42}
        switch 4294967295 {
        case 2147483647:
            var v12 uint = 551
            _ = v12
        }
        v1[9223372036854775807] = 127
        _ = v11
    }
    v13 := GenS1{F0: 97, F1: 0}
    if (227 < 2147483647) {
        var v14 uint = 18446744073709551615
        _ = v14
    } else {
        for fi15 := 0; fi15 < 15; fi15++ {
            if ((128 <= (^(fi15))) || true) {
                switch "\n\t" {
                case "":
                    v3 = (255 > 65535)
                case "a":
                    for fi16 := 0; fi16 < 16; fi16++ {
                        if (v3 || false) {
                            fmt.Println(fi15, fi16)
                            v3 = v3
                            if ((fi15 + (fi16 % (((fi16 << ((42) & 63))) | 1))) < (fi16 ^ fi16)) {
                                v17 := GenS3{F0: 1}
                                if true {
                                    v3 = false
                                }
                                v18 := []int{-654, fi16}
                                _ = v17
                                _ = v18
                            } else {
                                v19 := GenS4{F0: 1}
                                _ = v19
                            }
                            v20 := []int{}
                            v20 = append(v20, 2147483647)
                            var v21 float32 = (1.5 + (((0.0 * float32(1.5)) * 0.1) + 0.1))
                            var v22 int8 = 31
                            var v23 uint32 = 2147483648
                            v22 += 17
                            _ = v20
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                        _ = fi16
                    }
                case "foo bar":
                    v24 := []int{(^(fi15)), fi15}
                    _ = v24
                }
                v25 := []int{fi15}
                _ = v25
            } else {
                var v26 int64 = 9223372036854775807
                v3 = (!((-618 > 825) && v3))
                _ = v26
            }
            _ = fi15
        }
    }
    var v27 bool = true
    if false {
        if (0.1 != 0.0) {
            fmt.Println(v27, v3)
            fmt.Println(v27, v3)
        }
        for fi28 := 0; fi28 < 16; fi28++ {
            var v29 uint8 = 42
            for fi30 := 0; fi30 < 17; fi30++ {
                v27 = v27
                _ = fi30
            }
            var v31 bool = true
            _ = fi28
            _ = v29
            _ = v31
        }
    }
    for fi32 := 0; fi32 < 6; fi32++ {
        var v33 int = 447
        var v34 int8 = 0
        _ = fi32
        _ = v33
        _ = v34
    }
    var v35 uint = 360
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v13
    _ = v2
    _ = v27
    _ = v3
    _ = v35
}
