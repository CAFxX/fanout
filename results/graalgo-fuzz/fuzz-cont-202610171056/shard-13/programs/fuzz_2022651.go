package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int }
type GenS2 struct { F0 uint64; F1 int8; F2 string }
func genFn0(p0 bool, p1 uint16, p2 uint32) uint8 {
    if p0 {
        switch -1 {
        case 42:
            v7 := GenS1{F0: uint(p2), F1: -2147483648, F2: 9223372036854775807}
            _ = v7
        }
        var v8 int32 = 2147483647
        _ = v8
    }
    if p0 {
        fmt.Println(p2, p1, p1)
        if (255 == 4294967295) {
            for fi9 := 0; fi9 < 16; fi9++ {
                v10 := GenS1{F0: 597, F1: 0, F2: (fi9 + (fi9 | (fi9 * fi9)))}
                _ = fi9
                _ = v10
            }
            if (0.1 == 1e10) {
                var v11 uint16 = (p1 << ((4294967295) & 63))
                if (p1 <= p1) {
                    v11 -= (p1 % (((v11 & v11)) | 1))
                } else {
                    var v12 uint16 = v11
                    _ = v12
                }
                p0 = p0
                _ = v11
            } else {
                var v13 int32 = 0
                _ = v13
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return uint8(p1)
}
func genFn1(p0 uint) string {
    if true {
        p0 -= (p0 * 2)
        fmt.Println(p0)
        fmt.Println(p0)
    } else {
        p0 -= 255
        v14 := GenS1{F0: 18446744073709551615, F1: 0, F2: -1}
        var v15 uint64 = 18446744073709551615
        _ = v14
        _ = v15
    }
    fmt.Println(p0)
    v16 := GenS1{F0: 255, F1: 0, F2: 1}
    for fi17 := 0; fi17 < 7; fi17++ {
        v18 := []int{(^(fi17)), (fi17 << ((p0) & 63)), fi17, (-(fi17))}
        p0 = (p0 & 56)
        _ = fi17
        _ = v18
    }
    var v19 uint8 = 1
    _ = p0
    _ = v16
    _ = v19
    return ("" + (("" + "hello") + "hello"))
}

func main() {
    var v1 bool = true
    v2 := make(map[string]bool)
    v3 := []int{42}
    var v4 bool = (382 == 0)
    var v5 int16 = 1
    var v6 int8 = 0
    if (v1 || false) {
        fmt.Println(v5)
        var v20 uint32 = 65535
        _ = v20
    }
    fmt.Println(v1, v6)
    var v21 uint32 = 2
    v21 = (v21 % (((v21 - v21)) | 1))
    fmt.Println(v1, v21)
    for fi22 := 0; fi22 < 3; fi22++ {
        v23 := GenS1{F0: 65535, F1: 1, F2: fi22}
        _ = fi22
        _ = v23
    }
    var v24 int32 = 1
    if v4 {
        var v25 uint8 = 1
        v6 -= -128
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v25
    }
    var v26 int8 = v6
    if (32767 > -2147483648) {
        fmt.Println(v6, v21)
        var v27 bool = v1
        v3 = append(v3, 128)
        _ = v27
    } else {
        v21 -= (v21 & v21)
    }
    switch "a" {
    case "":
        v1 = v4
    case "\n\t":
        v6 = v26
    }
    v26 -= (v26 % (((v26 << ((2) & 63))) | 1))
    if (v1 || ((32767 != 65535) || ((^(v5)) >= -480))) {
        v28 := GenS2{F0: 432, F1: v6, F2: ""}
        _ = v28
    } else {
        if false {
            switch (("hello" + "") + ("hello" + "a")) {
            case "a":
                var v29 uint64 = 502
                _ = v29
            }
            var v30 float64 = (1e10 * (1e10 * 1e10))
            if v4 {
                v6 = 0
                var v31 string = "hello"
                _ = v31
            }
            _ = v30
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
