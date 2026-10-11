package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS4 struct { F0 uint64; F1 uint8; F2 int32 }
func genFn0() uint {
    var v7 bool = ((false || ((-32768 > -1) || true)) || (9223372036854775807 > 1))
    if false {
        v7 = (!(!v7))
        v8 := []int{32768, 42, -2147483648, 2147483647}
        if (-128 == 42) {
            _ = len(v8)
            for fi9 := 0; fi9 < 7; fi9++ {
                v10 := GenS1{F0: 0}
                v11 := GenS1{F0: 666}
                v7 = (-1 > 1000)
                _ = fi9
                _ = v10
                _ = v11
            }
        } else {
            var v12 int16 = -1
            if (!v7) {
                v12 -= v12
                if true {
                    _ = len(v8)
                }
                if (167 != 2) {
                    fmt.Println(v7, v12)
                    _ = len(v8)
                } else {
                    fmt.Println(v12)
                    fmt.Println(v12)
                }
            } else {
                v7 = ((-2.25 + 0.1) > (1.5 + float64((0.1 * 0.1))))
                var v13 uint64 = 2
                _ = v13
            }
            _ = v12
        }
        _ = v8
    }
    fmt.Println(v7)
    _ = v7
    return 4294967295
}
func genFn1(p0 uint8, p1 int64, p2 uint32) int64 {
    var v14 int32 = -1
    fmt.Println(v14, p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return (p1 + p1)
}

func main() {
    var v1 uint64 = 287
    var v2 string = (("hello" + (((("\n\t" + "") + "a") + "") + "foo bar")) + "a")
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]int)
    var v5 uint = 255
    var v6 uint16 = 1
    v15 := v4[4294967295]
    v16 := GenS2{F0: 887}
    switch (v15 % (((v15 & (v15 % ((1) | 1)))) | 1)) {
    case -2147483648:
        var v17 float64 = 3.14159
        _ = v17
    case -1:
        v1 += (v1 + (v1 ^ (^(v1))))
    }
    var v18 int64 = 0
    fmt.Println(v18, v2, v6)
    fmt.Println(v6, v6, v1)
    v19 := v3[42]
    switch "" {
    case "foo bar":
        var v20 int = 42
        _ = v20
    case "\n\t":
        for fi21 := 0; fi21 < 11; fi21++ {
            v22 := GenS2{F0: v5}
            _ = fi21
            _ = v22
        }
    case "hello":
        var v23 uint32 = (v19 % ((v19) | 1))
        _ = v23
    }
    switch (v15 >> ((genFn0()) & 63)) {
    case 32768:
        for fi24 := 0; fi24 < 4; fi24++ {
            v18 -= (v18 | v18)
            v25 := GenS1{F0: int16(v15)}
            _ = fi24
            _ = v25
        }
    }
    switch v2 {
    case "\n\t":
        v6 = v6
    }
    v26 := GenS3{F0: 429, F1: 21, F2: 42}
    v1 = (v1 | (v1 | 0))
    for fi27 := 0; fi27 < 16; fi27++ {
        switch fi27 {
        case 2147483647:
            var v28 bool = ((v6 - (v6 >> (((^(v5))) & 63))) != v6)
            _ = v28
        case -1:
            v29 := []int{(v15 % (((v15 & fi27)) | 1)), v15, v15}
            _ = v29
        default:
            switch (^(fi27)) {
            case -1:
                switch ("foo bar" + v2) {
                case "foo bar":
                    var v30 float32 = 0.1
                    _ = v30
                case "a":
                    var v31 string = ""
                    _ = v31
                default:
                    if (((!false) && true) || true) {
                        var v32 float64 = (0.0 + 1.5)
                        var v33 uint32 = 1
                        for fi34 := 0; fi34 < 8; fi34++ {
                            v35 := GenS1{F0: -144}
                            var v36 int = fi27
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                        _ = v32
                        _ = v33
                    }
                }
            case 42:
                fmt.Println(v2, v6)
            }
        }
        for fi37 := 0; fi37 < 18; fi37++ {
            fmt.Println(v19, v5)
            _ = fi37
        }
        v38 := GenS4{F0: 0, F1: 34, F2: -307}
        _ = fi27
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
