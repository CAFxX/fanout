package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 int16; F2 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 string; F2 int }
func genFn0(p0 string, p1 uint8, p2 int32) int8 {
    var v5 string = p0
    switch 127 {
    case -524:
        var v6 uint32 = 2147483648
        _ = v6
    case 0:
        for fi7 := 0; fi7 < 11; fi7++ {
            v8 := []int{}
            _ = fi7
            _ = v8
        }
    case 32767:
        p2 = (p2 % (((p2 << ((2) & 63))) | 1))
    }
    v9 := GenS1{F0: 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v9
    return 42
}
func genFn1(p0 string, p1 uint8) string {
    v10 := GenS2{F0: -2147483648, F1: -1, F2: true}
    var v11 string = (p0 + "\n\t")
    v12 := GenS2{F0: -100, F1: 1000, F2: (2 <= 0)}
    v11 = "\n\t"
    var v13 uint = 42
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return ("\n\t" + ("hello" + p0))
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 904
    var v3 uint = 65535
    var v4 int64 = 982
    if ((true && (745 != -521)) && false) {
        v14 := []int{638}
        v2 -= (^(v2))
        _ = v14
    }
    var v15 uint16 = v2
    v16 := GenS1{F0: int(v3)}
    switch 42 {
    case 4294967295:
        v17 := GenS3{F0: (^(v3))}
        _ = v17
    case 420:
        if (18446744073709551615 < 0) {
            if false {
                v18 := GenS4{F0: (^(v3)), F1: (("a" + "\n\t") + "hello"), F2: -342}
                _ = v18
            } else {
                var v19 int = 42
                _ = v19
            }
            for fi20 := 0; fi20 < 4; fi20++ {
                var v21 float32 = 3.14159
                _ = fi20
                _ = v21
            }
        } else {
            v22 := v1["hello"]
            v23 := GenS2{F0: 32768, F1: -32768, F2: (((^(v22)) > 914) && true)}
            _ = v22
            _ = v23
        }
    }
    var v24 string = ("hello" + "\n\t")
    v2 = (v2 ^ v15)
    switch 301 {
    case -217:
        v25 := []int{}
        _ = v25
    case -1:
        if true {
            var v26 int64 = v4
            _ = v26
        }
    }
    v27 := GenS2{F0: 1, F1: -508, F2: true}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v4
}
