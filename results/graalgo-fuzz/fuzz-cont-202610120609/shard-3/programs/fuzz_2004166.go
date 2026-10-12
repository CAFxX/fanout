package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 string; F1 uint16; F2 int64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{("hello" + ("a" + "")), (("foo bar" + "hello") + "hello"), "\n\t"}
    var v3 string = ""
    var v4 float32 = 0.1
    var v5 bool = (true || false)
    v6 := make(map[int]uint32)
    for fi7 := 0; fi7 < 7; fi7++ {
        _ = len(v2)
        _ = fi7
    }
    switch v3 {
    case "foo bar":
        var v8 int16 = -32768
        _ = v8
    case "":
        var v9 uint16 = 2
        _ = v9
    case "a":
        var v10 uint = 42
        _ = v10
    }
    fmt.Println(v3)
    for fi11 := 0; fi11 < 18; fi11++ {
        fmt.Println(fi11)
        var v12 uint = 857
        _ = fi11
        _ = v12
    }
    v4 = 0.0
    v4 = v4
    switch -9223372036854775808 {
    case -746:
        _ = len(v6)
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        v14 := GenS1{F0: v5, F1: 880, F2: 2147483648}
        if (-9223372036854775808 < -1) {
            _ = len(v6)
            v3 = ((v3 + (v3 + ((v3 + (v3 + v3)) + ("a" + (v3 + ((v3 + (v3 + v3)) + "hello")))))) + v3)
        }
        _ = fi13
        _ = v14
    }
    v3 = v3
    if ((1 >= 243) && v5) {
        v5 = false
    }
    var v15 uint32 = 154
    v4 += (v4 - ((-2.25 + float32(3.14159)) * (v4 * float32((1e10 + ((1e10 - ((0.0 - 3.14159) * 3.14159)) * (1e10 / float64(float32(-2.25)))))))))
    if (-2.25 < 1e10) {
        fmt.Println(v5, v4)
        v16 := GenS2{F0: 32767}
        _ = v16
    } else {
        fmt.Println(v5, v4)
        switch 2147483647 {
        case 4294967295:
            if (42 < 42) {
                var v17 bool = v5
                _ = v17
            } else {
                v5 = true
                var v18 bool = ((-102 > -9223372036854775808) || v5)
                _ = v18
            }
        case 42:
            if (42 <= 18446744073709551615) {
                switch "\n\t" {
                case "\n\t":
                    if v5 {
                        v19 := v1[v15]
                        _ = v19
                    }
                }
            } else {
                var v20 uint8 = 128
                for fi21 := 0; fi21 < 1; fi21++ {
                    v22 := GenS3{F0: (v15 << ((1) & 63)), F1: 0}
                    var v23 int32 = -2147483648
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                _ = v20
            }
        }
        var v24 uint64 = 0
        _ = v24
    }
    v3 = v3
    v3 = (("hello" + "") + v3)
    v25 := GenS4{F0: v3, F1: 803, F2: -834}
    fmt.Println(v4, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
