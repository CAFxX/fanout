package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 int8; F1 string }
func genFn0() uint32 {
    var v5 int = 2147483647
    fmt.Println(v5)
    v5 = v5
    var v6 float32 = 0.1
    var v7 bool = false
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int) int64 {
    p0 -= p0
    p0 = p0
    var v8 uint64 = 194
    v9 := GenS1{F0: ((0 > 0) && (0 < 18446744073709551615)), F1: (p0 >> ((0) & 63))}
    var v10 int64 = -1
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return -1
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 286
    var v3 uint8 = 128
    v4 := []string{"\n\t"}
    v2 = (v2 ^ (^(v2)))
    if (!true) {
        fmt.Println(v3, v3)
    } else {
        switch 4294967295 {
        case 370:
            if ("a" <= "hello") {
                switch -493 {
                case 127:
                    var v11 int = 556
                    _ = v11
                case -9223372036854775808:
                    v12 := GenS1{F0: (42 != -9223372036854775808), F1: -9223372036854775808}
                    _ = v12
                case 32767:
                    v3 -= 2
                }
            }
        case -678:
            _ = len(v4)
        case 793:
            switch 128 {
            case 42:
                v13 := GenS2{F0: 42}
                _ = v13
            case 4294967295:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case 962:
                v3 += (v3 - (v3 ^ 1))
            }
        }
        _ = len(v1)
        var v14 uint32 = (v2 % (((^(v2))) | 1))
        _ = v14
    }
    fmt.Println(v3, v2)
    v3 = v3
    v2 = genFn0()
    v3 += 1
    v15 := v1["a"]
    for fi16 := 0; fi16 < 14; fi16++ {
        for fi17 := 0; fi17 < 15; fi17++ {
            if ((-900 > 1) && v15) {
                v2 = (v2 & (v2 / ((4294967295) | 1)))
                v18 := GenS3{F0: -126}
                _ = v18
            }
            v19 := v1["a"]
            _ = fi17
            _ = v19
        }
        v20 := GenS4{F0: -24, F1: -41}
        fmt.Println(v2, v2, fi16)
        _ = fi16
        _ = v20
    }
    v21 := GenS3{F0: 42}
    var v22 uint16 = 535
    v22 = (v22 << ((42) & 63))
    for fi23 := 0; fi23 < 6; fi23++ {
        if true {
            v22 -= (v22 % (((v22 / (((v22 << ((671) & 63))) | 1))) | 1))
            for fi24 := 0; fi24 < 8; fi24++ {
                var v25 uint64 = 2
                _ = fi24
                _ = v25
            }
            v26 := GenS5{F0: -75, F1: ("hello" + "a")}
            _ = v26
        } else {
            for fi27 := 0; fi27 < 5; fi27++ {
                var v28 int = fi23
                v29 := GenS4{F0: 6, F1: -12}
                _ = fi27
                _ = v28
                _ = v29
            }
            var v30 int64 = 1000
            _ = v30
        }
        switch "hello" {
        case "hello":
            for fi31 := 0; fi31 < 2; fi31++ {
                var v32 uint16 = v22
                delete(v1, "")
                _ = fi31
                _ = v32
            }
        }
        _ = fi23
    }
    fmt.Println(v2)
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
