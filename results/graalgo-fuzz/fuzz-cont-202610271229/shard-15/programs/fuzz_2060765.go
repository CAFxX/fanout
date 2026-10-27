package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 bool; F1 int16; F2 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 uint32; F2 int8 }
type GenS5 struct { F0 uint8; F1 uint64 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]int)
    v3 := make(map[string]bool)
    v4 := make(map[string]string)
    var v5 uint32 = 415
    fmt.Println(v5)
    var v6 uint32 = (v5 % ((v5) | 1))
    var v7 float32 = (-2.25 * 1.5)
    _ = len(v4)
    fmt.Println(v5)
    for fi8 := 0; fi8 < 12; fi8++ {
        delete(v4, "")
        _ = fi8
    }
    switch -2147483648 {
    case -942:
        v9 := GenS1{F0: true, F1: false}
        _ = v9
    default:
        v5 = 2147483648
    }
    switch 127 {
    case 120:
        delete(v2, "hello")
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        for fi11 := 0; fi11 < 1; fi11++ {
            if (("" + ((("\n\t" + "a") + "foo bar") + ("foo bar" + "hello"))) >= ("\n\t" + ("foo bar" + "\n\t"))) {
                var v12 int16 = -8
                switch ("foo bar" + ("foo bar" + "foo bar")) {
                case "foo bar":
                    v13 := []int{fi11, 0}
                    _ = v13
                case "":
                    v14 := GenS2{F0: true, F1: v12, F2: 322}
                    _ = v14
                case "hello":
                    v15 := []int{(-(fi10)), (fi11 << ((300) & 63))}
                    _ = v15
                default:
                    if (((1e10 * 1.5) == 3.14159) && ((1 >= 127) || true)) {
                        fmt.Println(v6, fi11)
                        v6 += (v5 / ((v6) | 1))
                    }
                }
                v12 = v12
                _ = v12
            }
            v3["\n\t"] = false
            v6 -= v5
            var v16 uint32 = (v5 >> ((2) & 63))
            if true {
                v17 := []int{834, -541, (fi10 ^ fi11), fi10}
                fmt.Println(fi11, v5)
                switch "" {
                case "":
                    v16 = 2
                case "hello":
                    v18 := GenS2{F0: (1 > 18446744073709551615), F1: 0, F2: -9223372036854775808}
                    _ = v18
                }
                _ = v17
            }
            v19 := GenS2{F0: (32767 == 0), F1: 944, F2: 0}
            var v20 int16 = 42
            v21 := GenS3{F0: 42}
            v20 -= -1
            v22 := GenS2{F0: true, F1: (v20 | v20), F2: 0}
            _ = fi11
            _ = v16
            _ = v19
            _ = v20
            _ = v21
            _ = v22
        }
        _ = fi10
    }
    v23 := GenS4{F0: 1, F1: v5, F2: -76}
    for fi24 := 0; fi24 < 1; fi24++ {
        if true {
            fmt.Println(v7)
        }
        if false {
            switch (fi24 + fi24) {
            case -2147483648:
                var v25 int32 = 2147483647
                _ = v25
            }
            fmt.Println(v7, v5, v7)
        }
        v26 := []int{(fi24 + fi24)}
        _ = fi24
        _ = v26
    }
    v27 := GenS5{F0: 127, F1: 0}
    fmt.Println(v5, v6, v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
