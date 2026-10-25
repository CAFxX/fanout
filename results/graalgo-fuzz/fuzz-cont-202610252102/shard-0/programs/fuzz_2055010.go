package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int8 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 int; F1 int; F2 uint16 }
type GenS4 struct { F0 int32; F1 uint16; F2 int }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int32; F1 bool; F2 int16 }
func genFn0(p0 uint, p1 uint32, p2 uint32) uint8 {
    var v6 uint64 = 261
    switch "a" {
    case "foo bar":
        for fi7 := 0; fi7 < 6; fi7++ {
            switch (fi7 + fi7) {
            case 1:
                for fi8 := 0; fi8 < 2; fi8++ {
                    v6 -= (^(v6))
                    _ = fi8
                }
            case 0:
                v9 := GenS1{F0: -1, F1: -9223372036854775808, F2: 1}
                _ = v9
            }
            p0 += (p0 ^ p0)
            p2 = p1
            _ = fi7
        }
    }
    var v10 int = 32768
    var v11 int16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v6
    return 170
}

func main() {
    var v1 uint = 255
    var v2 uint32 = 809
    var v3 int = 42
    v4 := make(map[string]bool)
    var v5 uint16 = 369
    if (!(0 >= 427)) {
        v3 -= v3
        v1 = 2
    }
    var v12 int16 = -32768
    for fi13 := 0; fi13 < 13; fi13++ {
        v14 := GenS2{F0: 93, F1: "foo bar"}
        for fi15 := 0; fi15 < 1; fi15++ {
            if false {
                if false {
                    v3 -= fi15
                }
                v3 += -26
                fmt.Println(v1, v12)
            }
            if ("a" == "\n\t") {
                v2 = 2147483647
                v1 += v1
                if true {
                    fmt.Println(v12)
                    v16 := GenS3{F0: v3, F1: 4294967295, F2: v5}
                    _ = v16
                } else {
                    var v17 int32 = int32(v1)
                    v18 := GenS4{F0: (v17 & v17), F1: 336, F2: v3}
                    if false {
                        v12 = -805
                    } else {
                        v1 -= 18446744073709551615
                        var v19 uint = (^(v1))
                        v20 := []int{32768}
                        _ = v19
                        _ = v20
                    }
                    v21 := GenS5{F0: "hello"}
                    var v22 uint64 = 999
                    v23 := GenS6{F0: fi15}
                    _ = v17
                    _ = v18
                    _ = v21
                    _ = v22
                    _ = v23
                }
                v24 := GenS1{F0: 127, F1: -9223372036854775808, F2: -90}
                switch 42 {
                case -9223372036854775808:
                    fmt.Println(v5, fi15, fi15)
                case 2147483647:
                    v2 = v2
                case -201:
                    v5 = v5
                }
                _ = v24
            } else {
                var v25 uint8 = 42
                _ = v25
            }
            _ = fi15
        }
        var v26 int8 = -128
        _ = fi13
        _ = v14
        _ = v26
    }
    if (v1 > 0) {
        var v27 int64 = 0
        var v28 int32 = 0
        v5 -= v5
        _ = v27
        _ = v28
    } else {
        fmt.Println(v5, v2, v1)
    }
    var v29 int64 = -9223372036854775808
    v30 := GenS2{F0: (v5 & (v5 - v5)), F1: "hello"}
    v1 = 1
    v29 -= -1
    v31 := GenS3{F0: 906, F1: v3, F2: 42}
    v32 := GenS7{F0: -2147483648, F1: true, F2: v12}
    var v33 int16 = v12
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
