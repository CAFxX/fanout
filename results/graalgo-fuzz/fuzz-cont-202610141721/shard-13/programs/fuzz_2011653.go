package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS5 struct { F0 string; F1 uint8 }

func main() {
    var v1 int = -803
    var v2 uint16 = 65535
    v3 := make(map[int]int)
    var v4 uint64 = 948
    var v5 int32 = 1
    var v6 int8 = -74
    v2 = v2
    v3[v1] = v1
    var v7 int64 = 148
    if (true || true) {
        if (((0.0 + 0.1) * (0.1 + -2.25)) == 0.0) {
            if false {
                if ((!(false || false)) && true) {
                    fmt.Println(v5, v6)
                } else {
                    if (false && (!false)) {
                        v5 -= v5
                    } else {
                        v8 := GenS1{F0: 1, F1: 65535, F2: (v2 | (v2 | (v2 >> ((255) & 63))))}
                        _ = v8
                    }
                }
            } else {
                if true {
                    for fi9 := 0; fi9 < 8; fi9++ {
                        fmt.Println(v2, v5)
                        var v10 int64 = v7
                        _ = fi9
                        _ = v10
                    }
                    var v11 uint = 312
                    var v12 uint = (v11 + v11)
                    _ = v11
                    _ = v12
                }
                v4 = v4
            }
            switch "a" {
            case "a":
                v4 = 371
            case "foo bar":
                for fi13 := 0; fi13 < 6; fi13++ {
                    var v14 uint = 255
                    var v15 uint = v14
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            case "hello":
                switch 127 {
                case -73:
                    for fi16 := 0; fi16 < 9; fi16++ {
                        var v17 uint8 = 140
                        _ = fi16
                        _ = v17
                    }
                default:
                    switch (v1 << ((1) & 63)) {
                    case -2147483648:
                        fmt.Println(v7, v4, v2)
                    case 2147483647:
                        switch "a" {
                        case "a":
                            var v18 uint = 979
                            _ = v18
                        }
                    }
                }
            default:
                v5 = v5
            }
        } else {
            var v19 int8 = (v6 & (v6 ^ (v6 * (v6 ^ v6))))
            for fi20 := 0; fi20 < 13; fi20++ {
                for fi21 := 0; fi21 < 7; fi21++ {
                    v22 := GenS2{F0: "hello"}
                    fmt.Println(v7)
                    switch "hello" {
                    case "a":
                        v6 = v19
                    case "hello":
                        v23 := GenS3{F0: 0}
                        _ = v23
                    }
                    _ = fi21
                    _ = v22
                }
                _ = fi20
            }
            v24 := []int{(^(v1)), (-(v1)), v1}
            _ = v19
            _ = v24
        }
        fmt.Println(v7)
        v25 := v3[-488]
        _ = v25
    }
    if true {
        for fi26 := 0; fi26 < 16; fi26++ {
            v27 := v3[(v1 << ((651) & 63))]
            _ = fi26
            _ = v27
        }
        _ = len(v3)
    } else {
        var v28 float64 = 0.0
        v29 := GenS4{F0: 1000, F1: 255, F2: 4294967295}
        for fi30 := 0; fi30 < 6; fi30++ {
            var v31 uint64 = 4294967295
            for fi32 := 0; fi32 < 1; fi32++ {
                var v33 uint64 = (v4 << ((65535) & 63))
                _ = fi32
                _ = v33
            }
            _ = fi30
            _ = v31
        }
        _ = v28
        _ = v29
    }
    fmt.Println(v5, v2, v1)
    v2 -= v2
    v7 -= (v7 - 808)
    delete(v3, (v1 - v1))
    var v34 bool = true
    for fi35 := 0; fi35 < 7; fi35++ {
        var v36 int32 = v5
        v37 := GenS5{F0: "", F1: 1}
        _ = fi35
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 4; fi38++ {
        var v39 float32 = 3.14159
        _ = fi38
        _ = v39
    }
    for fi40 := 0; fi40 < 12; fi40++ {
        v41 := GenS3{F0: 42}
        v1 += (fi40 & 1)
        v42 := GenS2{F0: "a"}
        _ = fi40
        _ = v41
        _ = v42
    }
    if v34 {
        v1 = (v1 - 2147483647)
    } else {
        fmt.Println(v34, v6)
        v5 = (v5 - (v5 + (-(v5))))
    }
    var v43 uint8 = 150
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
}
