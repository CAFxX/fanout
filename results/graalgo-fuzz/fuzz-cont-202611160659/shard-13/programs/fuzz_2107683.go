package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 uint64; F1 uint16 }
func genFn0() int32 {
    var v7 uint16 = 32767
    var v8 uint16 = 65535
    switch 32767 {
    case -255:
        var v9 bool = (-1 > 0)
        _ = v9
    case 840:
        var v10 uint64 = 461
        _ = v10
    case -212:
        var v11 int8 = 1
        _ = v11
    default:
        switch "foo bar" {
        case "\n\t":
            for fi12 := 0; fi12 < 9; fi12++ {
                for fi13 := 0; fi13 < 1; fi13++ {
                    v8 += 118
                    v14 := GenS1{F0: 293}
                    if false {
                        v8 -= (v7 ^ v7)
                        v8 -= (v8 % ((255) | 1))
                        v15 := []int{(-(fi13)), -2147483648, fi13}
                        _ = v15
                    }
                    _ = fi13
                    _ = v14
                }
                for fi16 := 0; fi16 < 8; fi16++ {
                    if (0 >= -9223372036854775808) {
                        v17 := GenS1{F0: 179}
                        v18 := []int{9223372036854775807, fi16, fi16, (fi12 | 128)}
                        v8 = v8
                        _ = v17
                        _ = v18
                    } else {
                        var v19 float32 = -2.25
                        _ = v19
                    }
                    v20 := GenS1{F0: 299}
                    var v21 int = (fi16 + 2147483647)
                    _ = fi16
                    _ = v20
                    _ = v21
                }
                _ = fi12
            }
        }
    }
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]bool)
    v3 := make(map[string]int)
    var v4 bool = false
    var v5 int64 = 938
    v6 := make(map[uint32]bool)
    var v22 uint64 = 42
    v23 := GenS1{F0: 2}
    if (!((v4 && true) || true)) {
        v24 := []int{}
        if false {
            v24 = append(v24, -398)
        } else {
            v5 = (v5 & v5)
            var v25 int16 = 0
            _ = v25
        }
        v4 = v4
        _ = v24
    }
    v26 := GenS1{F0: 655}
    switch 4294967295 {
    case 127:
        var v27 int = 4294967295
        _ = v27
    case -197:
        _ = len(v6)
    case -147:
        for fi28 := 0; fi28 < 18; fi28++ {
            switch (fi28 << ((65535) & 63)) {
            case 9223372036854775807:
                v29 := GenS1{F0: 42}
                _ = v29
            case -551:
                var v30 float32 = 0.1
                _ = v30
            case 2147483647:
                v4 = (("foo bar" + "foo bar") != "hello")
            }
            var v31 float64 = 0.1
            _ = fi28
            _ = v31
        }
    }
    for fi32 := 0; fi32 < 4; fi32++ {
        if v4 {
            v4 = true
            v33 := GenS1{F0: 65535}
            v22 -= v22
            _ = v33
        }
        for fi34 := 0; fi34 < 3; fi34++ {
            var v35 uint = 193
            if false {
                if v4 {
                    var v36 uint8 = 255
                    _ = v36
                }
                switch ("" + "hello") {
                case "a":
                    v37 := []int{128, fi34, (fi32 * fi34), fi32}
                    _ = v37
                default:
                    switch "foo bar" {
                    case "a":
                        v38 := GenS2{F0: 680, F1: 411}
                        _ = v38
                    default:
                        fmt.Println(v22, v5, fi34)
                    }
                }
                fmt.Println(fi32, v35)
            } else {
                if false {
                    v39 := GenS3{F0: v22, F1: 1}
                    var v40 uint = 0
                    var v41 uint = v40
                    var v42 int64 = v5
                    _ = v39
                    _ = v40
                    _ = v41
                    _ = v42
                }
                switch "a" {
                case "":
                    var v43 float32 = 0.1
                    _ = v43
                case "hello":
                    v35 = v35
                }
            }
            _ = fi34
            _ = v35
        }
        _ = fi32
    }
    var v44 float32 = 1.5
    var v45 int = -113
    var v46 uint16 = 32767
    var v47 uint8 = 0
    switch "hello" {
    case "":
        v4 = (354 < 18446744073709551615)
    default:
        fmt.Println(v45, v47)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
