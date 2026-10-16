package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS4 struct { F0 string; F1 uint; F2 uint64 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 uint64, p1 int8, p2 uint8) int64 {
    v4 := GenS1{F0: 2, F1: 2147483648}
    p2 = p2
    p1 -= (p1 & p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 9223372036854775807
}

func main() {
    var v1 int64 = -1
    v2 := []string{"a", (("\n\t" + "\n\t") + (("" + ("foo bar" + "\n\t")) + ("" + "\n\t"))), "\n\t", "a"}
    v3 := make(map[string]int)
    for fi5 := 0; fi5 < 19; fi5++ {
        if true {
            for fi6 := 0; fi6 < 6; fi6++ {
                var v7 uint = 4294967295
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v8 int64 = v1
                _ = fi6
                _ = v7
                _ = v8
            }
        } else {
            v1 -= v1
            fmt.Println(v1)
            fmt.Println(fi5, fi5)
        }
        _ = fi5
    }
    var v9 string = "\n\t"
    if ((18446744073709551615 <= 4294967295) || false) {
        delete(v3, ("\n\t" + v9))
        v10 := GenS2{F0: (v1 & (^(v1))), F1: 32767}
        v11 := GenS1{F0: 128, F1: 2147483648}
        _ = v10
        _ = v11
    }
    var v12 uint = 255
    fmt.Println(v1)
    v13 := GenS3{F0: 1, F1: 269, F2: 4294967295}
    var v14 string = ("" + v9)
    var v15 int16 = 42
    v16 := GenS4{F0: ((v14 + ("hello" + "hello")) + ("\n\t" + v9)), F1: (v12 + v12), F2: 18446744073709551615}
    for fi17 := 0; fi17 < 5; fi17++ {
        v18 := GenS4{F0: ("foo bar" + (((v14 + v9) + (v9 + v9)) + v14)), F1: 14, F2: 302}
        _ = fi17
        _ = v18
    }
    fmt.Println(v15, v1, v9)
    v3[(v14 + "a")] = 1
    var v19 uint32 = 554
    v20 := GenS5{F0: 7}
    var v21 uint64 = 37
    switch 9223372036854775807 {
    case -9223372036854775808:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if (127 < 127) {
        fmt.Println(v14)
    } else {
        fmt.Println(v9, v21, v12)
        if true {
            v22 := GenS2{F0: genFn0(238, -128, 2), F1: 42}
            _ = v22
        } else {
            if len(v2) > 0 {
                v2[0] = v14
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if (v19 < uint32(v21)) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        }
    }
    var v23 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v9
}
