package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 bool) uint16 {
    var v4 string = (("hello" + (("" + "") + "hello")) + "a")
    var v5 uint64 = 4294967295
    v6 := []int{}
    p0 = (2 > 4294967295)
    for fi7 := 0; fi7 < 8; fi7++ {
        if (42 != 2) {
            v8 := GenS1{F0: 2147483648, F1: v5}
            for fi9 := 0; fi9 < 10; fi9++ {
                v10 := GenS1{F0: 42, F1: v5}
                _ = fi9
                _ = v10
            }
            p0 = true
            _ = v8
        }
        switch (v4 + ((("\n\t" + "hello") + ((v4 + v4) + "a")) + ((v4 + (("\n\t" + v4) + ("foo bar" + v4))) + "a"))) {
        case "\n\t":
            var v11 uint16 = 152
            _ = v11
        case "foo bar":
            switch fi7 {
            case -590:
                v12 := GenS2{F0: 42}
                _ = v12
            case -845:
                var v13 float32 = 3.14159
                _ = v13
            default:
                fmt.Println(v4, p0, fi7)
            }
        }
        if (int16(v5) >= 32767) {
            var v14 int32 = -1
            var v15 int32 = (v14 ^ int32(fi7))
            v6 = append(v6, fi7)
            _ = v14
            _ = v15
        }
        _ = fi7
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 325
}
func genFn1(p0 uint32) int32 {
    var v16 uint8 = 144
    var v17 uint64 = 0
    v18 := []int{-2147483648, 160}
    var v19 uint16 = 42
    _ = p0
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    return 2147483647
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]bool)
    var v3 bool = true
    if (65535 <= 0) {
        for fi20 := 0; fi20 < 14; fi20++ {
            v21 := []int{}
            _ = fi20
            _ = v21
        }
        v3 = (!v3)
        v3 = v3
    }
    v22 := []int{}
    v23 := GenS1{F0: 111, F1: 369}
    delete(v2, 2147483648)
    v3 = v3
    v24 := GenS2{F0: 42}
    var v25 int64 = -910
    for fi26 := 0; fi26 < 1; fi26++ {
        switch "\n\t" {
        case "a":
            for fi27 := 0; fi27 < 18; fi27++ {
                v28 := GenS3{F0: 0}
                var v29 int64 = v25
                var v30 int64 = (v25 >> ((135) & 63))
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
        case "foo bar":
            var v31 uint = 65535
            _ = v31
        default:
            v3 = v3
        }
        v25 += (-(v25))
        v32 := v2[17]
        _ = fi26
        _ = v32
    }
    if false {
        var v33 int = -1
        _ = v33
    } else {
        if v3 {
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v25)
            case "a":
                for fi34 := 0; fi34 < 16; fi34++ {
                    var v35 int64 = (^(v25))
                    if v3 {
                        var v36 int32 = -538
                        for fi37 := 0; fi37 < 12; fi37++ {
                            if len(v22) >= 2 {
                                v22 = v22[1:len(v22)]
                            }
                            fmt.Println(v35, fi34)
                            var v38 uint64 = 691
                            _ = fi37
                            _ = v38
                        }
                        v3 = (953 <= 0)
                        _ = v36
                    }
                    _ = fi34
                    _ = v35
                }
            default:
                v39 := GenS4{F0: 1, F1: 1}
                _ = v39
            }
            v3 = (-32768 != 32767)
            var v40 string = ("a" + "hello")
            _ = v40
        } else {
            v1[2147483647] = ""
        }
    }
    var v41 int16 = 1000
    v42 := GenS5{F0: 42}
    v3 = ((!v3) || (true && (v3 || v3)))
    var v43 uint8 = 0
    v25 = (v25 ^ v25)
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v41
    _ = v42
    _ = v43
}
