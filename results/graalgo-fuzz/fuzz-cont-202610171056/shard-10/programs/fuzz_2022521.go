package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 bool; F1 bool }
type GenS6 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 163
    v3 := []int{}
    v4 := make(map[uint32]string)
    var v5 float64 = -2.25
    var v6 uint64 = 4294967295
    switch "" {
    case "\n\t":
        fmt.Println(v6, v2)
    default:
        v3 = append(v3, -97)
    }
    v7 := GenS1{F0: (v6 << ((18446744073709551615) & 63)), F1: (!(false && false))}
    fmt.Println(v5, v5)
    var v8 string = "hello"
    fmt.Println(v6, v6)
    switch v8 {
    case "foo bar":
        v9 := GenS2{F0: 32767, F1: (v6 / (((v6 - (v6 >> ((0) & 63)))) | 1))}
        _ = v9
    default:
        v10 := GenS2{F0: -9223372036854775808, F1: v6}
        _ = v10
    }
    v11 := GenS3{F0: 926}
    var v12 int16 = 1
    if ((-1 < 42) || (1.5 <= v5)) {
        fmt.Println(v12)
    }
    switch 2147483647 {
    case 32768:
        for fi13 := 0; fi13 < 14; fi13++ {
            switch fi13 {
            case 32767:
                switch 32768 {
                case 42:
                    v14 := GenS4{F0: false, F1: -2147483648}
                    _ = v14
                case 128:
                    var v15 int16 = v12
                    _ = v15
                case 112:
                    switch ((v8 + "a") + "a") {
                    case "\n\t":
                        fmt.Println(v6, fi13)
                    default:
                        var v16 string = "\n\t"
                        _ = v16
                    }
                default:
                    if false {
                        var v17 uint16 = 32767
                        var v18 uint8 = v2
                        var v19 uint = 627
                        _ = v17
                        _ = v18
                        _ = v19
                    } else {
                        fmt.Println(v2, v8, v8)
                    }
                }
            case 4294967295:
                if true {
                    var v20 uint = 283
                    var v21 uint8 = (^(v2))
                    var v22 uint64 = 18446744073709551615
                    _ = v20
                    _ = v21
                    _ = v22
                }
            case -372:
                var v23 int8 = 100
                _ = v23
            }
            var v24 uint16 = 255
            v6 += (v6 << ((18446744073709551615) & 63))
            _ = fi13
            _ = v24
        }
    case 4294967295:
        if (-9223372036854775808 > 42) {
            switch v8 {
            case "a":
                fmt.Println(v12, v6)
            case "foo bar":
                v4[65535] = v8
            case "\n\t":
                var v25 int8 = -1
                _ = v25
            default:
                if (32767 == -505) {
                    for fi26 := 0; fi26 < 17; fi26++ {
                        fmt.Println(v2, fi26)
                        _ = fi26
                    }
                    v27 := GenS5{F0: (!("foo bar" >= ("foo bar" + v8))), F1: true}
                    var v28 string = (v8 + v8)
                    _ = v27
                    _ = v28
                } else {
                    var v29 int8 = 1
                    switch "" {
                    case "a":
                        fmt.Println(v6, v12)
                    case "\n\t":
                        if ((1 == (v6 + 0)) || false) {
                            v2 -= v2
                            v1[669] = false
                        }
                    case "":
                        v30 := GenS1{F0: 0, F1: (true && (95 >= -1))}
                        _ = v30
                    }
                    _ = v29
                }
            }
            switch (("" + v8) + (v8 + "foo bar")) {
            case "":
                v4[2] = v8
            }
        } else {
            v8 = ((v8 + "foo bar") + v8)
            v5 -= (float64(float32((v5 - 0.1))) + float64(0.1))
            if true {
                _ = len(v1)
            } else {
                for fi31 := 0; fi31 < 5; fi31++ {
                    v8 = "a"
                    var v32 uint8 = v2
                    _ = fi31
                    _ = v32
                }
                var v33 string = ("hello" + "a")
                _ = v33
            }
        }
    case 2147483647:
        v6 += 605
    default:
        v3 = append(v3, -663)
    }
    var v34 int64 = -1
    if false {
        v6 = 42
        _ = len(v1)
        var v35 int = 4294967295
        _ = v35
    }
    for fi36 := 0; fi36 < 15; fi36++ {
        v37 := GenS6{F0: 232}
        _ = fi36
        _ = v37
    }
    var v38 bool = (!true)
    v38 = (false || v38)
    v2 = (v2 / ((128) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
