package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 int64 }
func genFn0(p0 uint16) int16 {
    p0 += p0
    fmt.Println(p0)
    for fi7 := 0; fi7 < 18; fi7++ {
        if false {
            for fi8 := 0; fi8 < 7; fi8++ {
                var v9 uint8 = 128
                v10 := []int{fi8, 9223372036854775807, fi8}
                var v11 bool = ((((0.0 * float64((3.14159 * 1e10))) / -2.25) + 1e10) <= 3.14159)
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            var v12 uint32 = 4294967295
            _ = v12
        } else {
            var v13 float32 = 0.1
            p0 = (p0 % ((p0) | 1))
            _ = v13
        }
        _ = fi7
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        v15 := GenS1{F0: 777, F1: ((((p0 >> ((65535) & 63)) > p0) && (("a" + "hello") <= "")) || false), F2: 42}
        _ = fi14
        _ = v15
    }
    if true {
        v16 := GenS2{F0: -539}
        fmt.Println(p0)
        if (9223372036854775807 <= -1) {
            p0 -= (p0 | p0)
            for fi17 := 0; fi17 < 6; fi17++ {
                if (0.0 <= 1e10) {
                    p0 = p0
                    if true {
                        v18 := GenS2{F0: 707}
                        if ((!false) && true) {
                            fmt.Println(p0, fi17)
                        }
                        _ = v18
                    } else {
                        var v19 float64 = 0.0
                        for fi20 := 0; fi20 < 3; fi20++ {
                            p0 = p0
                            _ = fi20
                        }
                        p0 -= (p0 >> ((2) & 63))
                        _ = v19
                    }
                }
                var v21 int = (fi17 >> ((879) & 63))
                v22 := []int{(fi17 >> ((4294967295) & 63)), (v21 << ((65535) & 63)), fi17}
                _ = fi17
                _ = v21
                _ = v22
            }
            v23 := GenS3{F0: int64(p0), F1: -455}
            _ = v23
        }
        _ = v16
    }
    _ = p0
    return 32767
}

func main() {
    v1 := []string{}
    var v2 int64 = -278
    var v3 uint64 = 2
    var v4 uint8 = 0
    v5 := make(map[string]string)
    v6 := make(map[int]string)
    v4 += (v4 << ((973) & 63))
    _ = len(v5)
    if true {
        v2 += v2
        fmt.Println(v3, v3)
        var v24 float32 = (1.5 + float32(-2.25))
        _ = v24
    }
    var v25 uint16 = 0
    v26 := GenS3{F0: v2, F1: v2}
    var v27 uint16 = 0
    for fi28 := 0; fi28 < 3; fi28++ {
        if (true && true) {
            v4 -= 0
            delete(v6, fi28)
        }
        switch ("" + "\n\t") {
        case "":
            var v29 uint32 = 42
            _ = v29
        case "foo bar":
            for fi30 := 0; fi30 < 18; fi30++ {
                var v31 float64 = 1.5
                fmt.Println(v27)
                var v32 int8 = 100
                _ = fi30
                _ = v31
                _ = v32
            }
        }
        v27 = v27
        _ = fi28
    }
    for fi33 := 0; fi33 < 12; fi33++ {
        if false {
            fmt.Println(v3, v25, v4)
        }
        _ = fi33
    }
    var v34 uint32 = uint32(v3)
    v1 = append(v1, ("\n\t" + "a"))
    var v35 uint32 = 42
    v34 -= 4294967295
    v2 -= v2
    fmt.Println(v4, v34)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
