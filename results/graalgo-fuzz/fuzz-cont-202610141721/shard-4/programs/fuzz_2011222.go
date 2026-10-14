package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 string; F1 string; F2 uint64 }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 uint64; F1 uint64; F2 uint32 }
func genFn0(p0 string, p1 uint32, p2 uint) int8 {
    var v4 int = -538
    var v5 float32 = 3.14159
    p1 += (p1 * 1)
    switch v4 {
    case 1:
        switch v4 {
        case 42:
            p2 -= (p2 % ((p2) | 1))
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 127
}

func main() {
    var v1 uint32 = 65535
    var v2 bool = true
    var v3 uint16 = 2
    v3 = v3
    if v2 {
        if true {
            v6 := GenS1{F0: 0, F1: 1}
            switch "" {
            case "foo bar":
                v7 := []int{}
                _ = v7
            case "":
                for fi8 := 0; fi8 < 8; fi8++ {
                    var v9 uint16 = (v3 >> ((285) & 63))
                    _ = fi8
                    _ = v9
                }
            }
            _ = v6
        }
    }
    v10 := GenS2{F0: 0, F1: 255}
    v2 = ((v1 >> ((2) & 63)) >= v1)
    v11 := []int{4294967295, 32768, 2147483647, 443}
    var v12 int = -1
    v11 = append(v11, v12)
    fmt.Println(v12, v3)
    var v13 int64 = 9223372036854775807
    if (true && false) {
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 bool = true
            if v15 {
                v16 := GenS3{F0: "foo bar", F1: ("a" + "\n\t"), F2: 18446744073709551615}
                v17 := GenS3{F0: "", F1: "hello", F2: 1}
                _ = v16
                _ = v17
            } else {
                if (2147483647 < 2147483647) {
                    for fi18 := 0; fi18 < 7; fi18++ {
                        v3 = 32767
                        _ = fi18
                    }
                    v2 = true
                    if len(v11) > 0 {
                        v11[0] = fi14
                    }
                } else {
                    var v19 int64 = v13
                    if (v12 < 9223372036854775807) {
                        v19 += (v19 | (v19 >> ((255) & 63)))
                    }
                    _ = v19
                }
                switch (v12 / ((128) | 1)) {
                case 258:
                    v20 := GenS3{F0: "\n\t", F1: "\n\t", F2: 887}
                    _ = v20
                default:
                    v21 := GenS3{F0: "hello", F1: ("hello" + ""), F2: 42}
                    _ = v21
                }
                v22 := GenS3{F0: "\n\t", F1: ((("foo bar" + "hello") + "") + "hello"), F2: 0}
                _ = v22
            }
            _ = fi14
            _ = v15
        }
        v12 -= (^(v12))
    } else {
        fmt.Println(v1, v1, v3)
        v23 := GenS2{F0: 255, F1: 18446744073709551615}
        _ = v23
    }
    fmt.Println(v13, v13)
    switch -522 {
    case 1:
        for fi24 := 0; fi24 < 10; fi24++ {
            switch ("foo bar" + ("hello" + "a")) {
            case "\n\t":
                v1 -= v1
            case "hello":
                v25 := GenS4{F0: v1, F1: 1}
                _ = v25
            case "a":
                var v26 float64 = (((1e10 + 1e10) / (1e10 + -2.25)) - float64(-2.25))
                _ = v26
            }
            v12 += v12
            v3 = v3
            _ = fi24
        }
    }
    switch "" {
    case "":
        v27 := GenS5{F0: 42, F1: 4294967295, F2: v1}
        _ = v27
    }
    _ = len(v11)
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
