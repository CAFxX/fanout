package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS2 struct { F0 int }
func genFn0(p0 uint32) int32 {
    v4 := []int{-9223372036854775808, 693}
    for fi5 := 0; fi5 < 11; fi5++ {
        fmt.Println(fi5, p0)
        switch (("\n\t" + "foo bar") + (("hello" + "hello") + ("hello" + ("\n\t" + (("a" + "\n\t") + "foo bar"))))) {
        case "hello":
            if (!false) {
                switch (fi5 << ((2) & 63)) {
                case 9223372036854775807:
                    var v6 uint8 = 0
                    _ = v6
                case -1:
                    v4 = append(v4, -1)
                default:
                    var v7 int8 = 1
                    _ = v7
                }
                if false {
                    var v8 int8 = -1
                    var v9 int64 = 1
                    _ = v8
                    _ = v9
                } else {
                    if ((!(168 <= 65535)) && (true && true)) {
                        var v10 uint8 = 255
                        _ = v10
                    }
                    for fi11 := 0; fi11 < 20; fi11++ {
                        _ = len(v4)
                        var v12 uint16 = 255
                        p0 += (p0 ^ (p0 ^ p0))
                        _ = fi11
                        _ = v12
                    }
                }
            } else {
                var v13 uint64 = 4294967295
                _ = v13
            }
        case "foo bar":
            v14 := GenS1{F0: 127, F1: uint64(p0), F2: 350}
            _ = v14
        }
        _ = fi5
    }
    var v15 int64 = 237
    if (42 != 0) {
        var v16 int16 = 32767
        _ = v16
    }
    _ = p0
    _ = v15
    _ = v4
    return 858
}
func genFn1(p0 int, p1 uint8, p2 uint32) uint16 {
    v17 := []int{}
    v18 := GenS1{F0: (p1 & (p1 << ((105) & 63))), F1: 18446744073709551615, F2: 9223372036854775807}
    var v19 int32 = 2147483647
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    _ = v19
    return 814
}

func main() {
    var v1 float32 = -2.25
    v2 := []uint32{650}
    v3 := make(map[string]uint32)
    var v20 int = 0
    var v21 uint = 1
    v22 := GenS1{F0: 2, F1: 239, F2: -969}
    var v23 uint32 = 0
    v21 += v21
    if len(v2) > 0 {
        v2[0] = (^(v23))
    }
    v24 := GenS2{F0: (v20 * -537)}
    v20 += v20
    var v25 float64 = -2.25
    v21 = v21
    if (9223372036854775807 == v20) {
        fmt.Println(v1, v1)
    } else {
        var v26 uint16 = 65535
        var v27 bool = false
        fmt.Println(v26)
        _ = v26
        _ = v27
    }
    switch (-(v20)) {
    case 919:
        switch "foo bar" {
        case "hello":
            var v28 int16 = 32767
            _ = v28
        case "\n\t":
            switch v20 {
            case -2147483648:
                v29 := v3[(("foo bar" + "") + ("hello" + ("\n\t" + (((("\n\t" + "foo bar") + "\n\t") + "\n\t") + "hello"))))]
                _ = v29
            case 0:
                fmt.Println(v23, v23)
            }
        }
    case 32767:
        _ = len(v2)
    default:
        var v30 float32 = float32(v25)
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
