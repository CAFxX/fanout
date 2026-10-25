package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 int8; F2 uint }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 uint16; F1 int64 }
func genFn0(p0 uint, p1 int64, p2 string) string {
    switch p2 {
    case "":
        if false {
            if (36 < 2) {
                if ((!((42 <= 42) && (!((1.5 + 1e10) > 1.5)))) && false) {
                    switch -1 {
                    case 2147483647:
                        p2 = (p2 + "")
                    }
                }
            }
        } else {
            var v7 int = 32768
            _ = v7
        }
    case "a":
        var v8 int32 = int32(p0)
        _ = v8
    case "foo bar":
        v9 := GenS1{F0: 32767, F1: 1, F2: p0}
        _ = v9
    }
    v10 := []int{42}
    for fi11 := 0; fi11 < 7; fi11++ {
        if len(v10) > 0 {
            v10[0] = fi11
        }
        switch fi11 {
        case -1:
            p0 = p0
        case -851:
            fmt.Println(fi11)
        default:
            switch (^(fi11)) {
            case 0:
                for fi12 := 0; fi12 < 1; fi12++ {
                    v13 := GenS1{F0: 255, F1: p1, F2: 18446744073709551615}
                    var v14 uint8 = 128
                    var v15 uint = p0
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
            case -2147483648:
                if len(v10) > 0 {
                    v10[0] = -2147483648
                }
            }
        }
        var v16 float32 = -2.25
        _ = fi11
        _ = v16
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        v18 := GenS2{F0: 42}
        var v19 bool = false
        fmt.Println(fi17)
        _ = fi17
        _ = v18
        _ = v19
    }
    if ((p1 % ((p1) | 1)) < p1) {
        _ = len(v10)
    } else {
        var v20 bool = false
        v21 := GenS2{F0: 2147483647}
        _ = v20
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return "\n\t"
}

func main() {
    var v1 string = ""
    var v2 int8 = 100
    var v3 uint16 = 42
    var v4 string = ((((((((v1 + v1) + "\n\t") + ((v1 + ("foo bar" + "a")) + "a")) + (v1 + (v1 + ((v1 + ("" + "foo bar")) + (("a" + "hello") + v1))))) + (v1 + v1)) + v1) + v1) + v1)
    var v5 float32 = (0.1 / float32(1.5))
    v6 := []int{464, 32768, 32768, 80}
    switch -782 {
    case 791:
        for fi22 := 0; fi22 < 3; fi22++ {
            v5 = ((3.14159 / float32(1e10)) - v5)
            v1 = v4
            fmt.Println(v5)
            _ = fi22
        }
    case -153:
        var v23 int64 = 294
        _ = v23
    case 42:
        var v24 int64 = -9223372036854775808
        _ = v24
    default:
        v25 := GenS2{F0: 679}
        _ = v25
    }
    v6 = append(v6, 32768)
    v26 := GenS3{F0: 1, F1: v2, F2: 42}
    var v27 uint = 949
    v28 := GenS1{F0: v3, F1: 42, F2: v27}
    v4 = ("foo bar" + v4)
    var v29 bool = true
    v30 := GenS4{F0: v29, F1: -9223372036854775808}
    v6 = append(v6, 2147483647)
    v3 = (v3 << (((^(v27))) & 63))
    v3 += 1
    v31 := GenS5{F0: 1, F1: -9223372036854775808}
    if v29 {
        var v32 int8 = -128
        var v33 int8 = (^(v2))
        _ = v32
        _ = v33
    }
    var v34 int32 = 1
    var v35 uint64 = 831
    if len(v6) > 0 {
        v6[0] = -994
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
