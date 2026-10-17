package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int16; F1 bool; F2 bool }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint8 }
func genFn0(p0 int32, p1 uint16, p2 int16) int64 {
    var v6 int64 = 381
    v7 := []int{9223372036854775807}
    var v8 uint64 = 608
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return v6
}
func genFn1(p0 int16, p1 uint64) string {
    var v9 int64 = -1
    v9 = (v9 + 42)
    v10 := GenS1{F0: 554, F1: 127}
    for fi11 := 0; fi11 < 13; fi11++ {
        switch 32768 {
        case 128:
            v12 := GenS1{F0: 42, F1: 0}
            _ = v12
        case 9223372036854775807:
            switch (fi11 % ((522) | 1)) {
            case -540:
                v13 := []int{866, fi11}
                _ = v13
            default:
                var v14 uint64 = p1
                _ = v14
            }
        default:
            var v15 string = "hello"
            _ = v15
        }
        v16 := GenS2{F0: -2147483648, F1: (-(p0))}
        _ = fi11
        _ = v16
    }
    var v17 int32 = 643
    _ = p0
    _ = p1
    _ = v10
    _ = v17
    _ = v9
    return ""
}

func main() {
    var v1 uint32 = 4294967295
    v2 := []int{-1, 810}
    var v3 string = "hello"
    var v4 uint32 = 42
    v5 := make(map[uint32]uint32)
    fmt.Println(v3)
    if (true || false) {
        v18 := GenS3{F0: -32768, F1: (!true), F2: (!false)}
        _ = v18
    }
    fmt.Println(v1, v3, v1)
    for fi19 := 0; fi19 < 15; fi19++ {
        v20 := GenS4{F0: -1, F1: 103}
        _ = fi19
        _ = v20
    }
    if false {
        v3 = ("\n\t" + v3)
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        var v22 int = fi21
        v3 = "foo bar"
        _ = fi21
        _ = v22
    }
    var v23 int16 = 726
    var v24 uint = 65
    var v25 uint16 = 711
    for fi26 := 0; fi26 < 4; fi26++ {
        for fi27 := 0; fi27 < 20; fi27++ {
            switch ("foo bar" + ((((v3 + v3) + (v3 + (v3 + (v3 + (("a" + ("foo bar" + v3)) + v3))))) + v3) + ("hello" + "a"))) {
            case "a":
                v23 = (v23 - -1)
                fmt.Println(v3)
                v28 := GenS4{F0: -370, F1: 100}
                _ = v28
            default:
                var v29 uint = 532
                _ = v29
            }
            switch (v3 + "hello") {
            case "\n\t":
                if (213 > 128) {
                    var v30 uint16 = (v25 / ((0) | 1))
                    var v31 int16 = (v23 / (((v23 & (v23 | v23))) | 1))
                    var v32 int8 = -56
                    _ = v30
                    _ = v31
                    _ = v32
                } else {
                    fmt.Println(v23, v23)
                }
            }
            v33 := GenS5{F0: 18446744073709551615}
            _ = fi27
            _ = v33
        }
        var v34 int64 = -1
        _ = fi26
        _ = v34
    }
    for fi35 := 0; fi35 < 11; fi35++ {
        delete(v5, v4)
        switch (fi35 | (fi35 >> (((^(v24))) & 63))) {
        case -16:
            var v36 bool = false
            _ = v36
        case 42:
            for fi37 := 0; fi37 < 16; fi37++ {
                switch fi37 {
                case -1:
                    fmt.Println(v24, fi35)
                    switch (v3 + "hello") {
                    case "a":
                        v25 = v25
                    }
                default:
                    fmt.Println(v4)
                }
                if (((v3 + v3) == (v3 + ("" + ""))) || false) {
                    if false {
                        var v38 int = 0
                        _ = v38
                    }
                    var v39 int8 = 42
                    v24 = v24
                    _ = v39
                }
                if (true && false) {
                    v2 = append(v2, 9223372036854775807)
                    var v40 int8 = -59
                    _ = v40
                }
                var v41 float32 = 3.14159
                switch fi35 {
                case 32768:
                    delete(v5, v1)
                case 342:
                    var v42 uint8 = 0
                    _ = v42
                }
                _ = fi37
                _ = v41
            }
        default:
            var v43 uint64 = 1
            _ = v43
        }
        _ = fi35
    }
    v44 := GenS6{F0: 255}
    var v45 int8 = -1
    fmt.Println(v24)
    var v46 int16 = int16(v25)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
}
