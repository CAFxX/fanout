package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint64 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int16 }
func genFn0() string {
    if (false && true) {
        v6 := []int{-1}
        _ = v6
    } else {
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 int64 = -507
            _ = fi7
            _ = v8
        }
    }
    if false {
        v9 := []int{401, 127, -9223372036854775808, 42}
        switch ("hello" + "\n\t") {
        case "":
            var v10 bool = (!(-32768 >= 466))
            _ = v10
        }
        _ = v9
    }
    return "hello"
}
func genFn1(p0 uint16, p1 int8, p2 string) uint16 {
    var v11 int8 = 0
    if false {
        for fi12 := 0; fi12 < 5; fi12++ {
            var v13 int = (fi12 << ((7) & 63))
            v14 := GenS1{F0: p2, F1: true, F2: 0}
            _ = fi12
            _ = v13
            _ = v14
        }
        p2 = p2
        for fi15 := 0; fi15 < 18; fi15++ {
            v16 := []int{(fi15 << ((18446744073709551615) & 63)), fi15}
            _ = fi15
            _ = v16
        }
    }
    var v17 uint64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v17
    return uint16(v17)
}

func main() {
    var v1 uint16 = 255
    v2 := []uint32{2147483647, 65535, 520}
    v3 := make(map[string]bool)
    v4 := make(map[int]bool)
    var v5 float64 = 1e10
    var v18 int8 = 42
    if true {
        v18 -= v18
        for fi19 := 0; fi19 < 14; fi19++ {
            v1 -= v1
            var v20 uint64 = 1
            fmt.Println(v1, fi19)
            _ = fi19
            _ = v20
        }
        var v21 int16 = -32768
        _ = v21
    }
    v18 += 100
    fmt.Println(v5, v18, v1)
    v22 := GenS1{F0: "hello", F1: true, F2: 359}
    for fi23 := 0; fi23 < 8; fi23++ {
        v1 -= genFn1(v1, (v18 - v18), "")
        for fi24 := 0; fi24 < 1; fi24++ {
            if false {
                v25 := GenS2{F0: "foo bar", F1: (v18 >> ((4294967295) & 63))}
                v2 = append(v2, 755)
                v5 -= (((v5 - -2.25) + 0.1) + v5)
                if true {
                    v3["a"] = false
                    v26 := GenS3{F0: -543}
                    _ = v26
                } else {
                    var v27 int = 2147483647
                    v28 := GenS2{F0: ("foo bar" + "a"), F1: int8(v1)}
                    var v29 int32 = -590
                    _ = v27
                    _ = v28
                    _ = v29
                }
                _ = v25
            }
            var v30 uint = 65535
            if (false && true) {
                var v31 uint32 = 867
                var v32 int = -1
                _ = v31
                _ = v32
            }
            _ = fi24
            _ = v30
        }
        _ = fi23
    }
    var v33 bool = (803 != 32768)
    var v34 float64 = 1.5
    for fi35 := 0; fi35 < 2; fi35++ {
        for fi36 := 0; fi36 < 20; fi36++ {
            var v37 uint32 = 65535
            _ = fi36
            _ = v37
        }
        _ = fi35
    }
    v34 -= ((float64(3.14159) - v5) / v34)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
