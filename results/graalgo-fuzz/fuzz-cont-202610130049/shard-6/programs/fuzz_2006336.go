package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 uint8; F2 int32 }
func genFn0(p0 int, p1 uint64, p2 int8) uint32 {
    var v5 bool = true
    v5 = v5
    var v6 float32 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 65535
}
func genFn1(p0 int8, p1 int16) uint64 {
    for fi7 := 0; fi7 < 2; fi7++ {
        fmt.Println(fi7)
        _ = fi7
    }
    v8 := GenS1{F0: 42}
    v9 := GenS2{F0: -9223372036854775808}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 177
    var v3 int8 = -1
    v4 := make(map[uint32]bool)
    v3 += -37
    var v10 uint64 = 42
    v10 = genFn1((v3 / ((v3) | 1)), 32767)
    for fi11 := 0; fi11 < 20; fi11++ {
        v10 = (v10 & (v10 % ((v10) | 1)))
        v12 := []int{fi11, (fi11 + fi11)}
        v10 -= v10
        _ = fi11
        _ = v12
    }
    v2 = v2
    if false {
        fmt.Println(v10, v3)
        var v13 bool = true
        _ = v13
    }
    switch int(v2) {
    case 9223372036854775807:
        for fi14 := 0; fi14 < 6; fi14++ {
            fmt.Println(fi14, v2)
            var v15 bool = ((0 != -1) || true)
            _ = fi14
            _ = v15
        }
    default:
        var v16 int = 0
        _ = v16
    }
    v3 += (v3 / (((v3 << ((4294967295) & 63))) | 1))
    v3 += v3
    if (!(v3 < (^(v3)))) {
        var v17 bool = (false && (true && (!false)))
        v18 := GenS3{F0: true, F1: 1, F2: 0}
        for fi19 := 0; fi19 < 5; fi19++ {
            v20 := v1["a"]
            v21 := GenS1{F0: genFn0((^(fi19)), (^(v10)), -128)}
            var v22 uint16 = 2
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
        _ = v17
        _ = v18
    } else {
        if true {
            if (32767 < 322) {
                var v23 int64 = 0
                _ = v23
            }
        } else {
            if (true || true) {
                fmt.Println(v2, v10)
                for fi24 := 0; fi24 < 1; fi24++ {
                    var v25 bool = false
                    _ = fi24
                    _ = v25
                }
            }
        }
        switch "\n\t" {
        case "hello":
            v10 = v10
        case "foo bar":
            fmt.Println(v10, v2, v2)
        case "a":
            v26 := GenS1{F0: 1}
            _ = v26
        default:
            v10 += genFn1(int8(v10), -753)
        }
    }
    if true {
        switch "hello" {
        case "\n\t":
            var v27 uint16 = 255
            _ = v27
        case "hello":
            var v28 string = "a"
            _ = v28
        case "":
            v29 := []int{32767, 410, 2147483647}
            _ = v29
        }
    }
    v10 -= v10
    v3 = (v3 | 65)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
}
