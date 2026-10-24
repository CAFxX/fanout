package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 bool; F1 int64; F2 int32 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int8; F1 int }
func genFn0(p0 uint32, p1 uint64) uint32 {
    fmt.Println(p0)
    v6 := GenS1{F0: p1}
    var v7 uint8 = 1
    switch 42 {
    case 32767:
        var v8 uint16 = 32767
        _ = v8
    case 1:
        if false {
            v7 += 2
        } else {
            p1 = (p1 - p1)
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 int64 = 42
                v7 = (v7 ^ (^(v7)))
                fmt.Println(v7, v7, fi9)
                _ = fi9
                _ = v10
            }
        }
    }
    v11 := GenS2{F0: -286, F1: -829}
    _ = p0
    _ = p1
    _ = v11
    _ = v6
    _ = v7
    return 512
}
func genFn1(p0 uint, p1 bool) string {
    p1 = (!p1)
    p1 = (937 != -405)
    var v12 int16 = 0
    for fi13 := 0; fi13 < 18; fi13++ {
        p0 -= p0
        v14 := GenS2{F0: -45, F1: 2147483647}
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v12
    return ("" + (("a" + "\n\t") + ""))
}

func main() {
    var v1 int8 = -24
    var v2 uint16 = 32767
    v3 := make(map[string]int)
    var v4 float64 = (3.14159 + 3.14159)
    v5 := []int{0}
    var v15 uint32 = 2147483647
    v2 = v2
    v5 = append(v5, 127)
    fmt.Println(v15)
    var v16 uint16 = (v2 & 821)
    if true {
        var v17 int8 = (v1 | 100)
        switch "a" {
        case "a":
            v2 -= (v16 ^ (v16 | (v2 + v2)))
        case "":
            for fi18 := 0; fi18 < 10; fi18++ {
                v19 := GenS3{F0: true, F1: 9223372036854775807, F2: -361}
                v2 = (v16 << ((372) & 63))
                v4 = ((v4 + float64(((1.5 * -2.25) + (0.0 * 0.0)))) + v4)
                _ = fi18
                _ = v19
            }
        }
        _ = v17
    } else {
        var v20 bool = true
        switch 677 {
        case 4294967295:
            v21 := GenS2{F0: -319, F1: 373}
            _ = v21
        case 2147483647:
            var v22 int16 = 32767
            _ = v22
        default:
            v23 := GenS4{F0: 1, F1: 32767}
            _ = v23
        }
        switch "hello" {
        case "\n\t":
            for fi24 := 0; fi24 < 16; fi24++ {
                v4 = 1.5
                if ((v1 << ((4294967295) & 63)) >= (v1 + v1)) {
                    delete(v3, ((("foo bar" + ("\n\t" + "\n\t")) + "") + "a"))
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v25 int16 = -369
                    _ = v25
                }
                _ = fi24
            }
        case "hello":
            v16 = (^(v2))
        case "":
            for fi26 := 0; fi26 < 17; fi26++ {
                v27 := GenS5{F0: v16}
                _ = fi26
                _ = v27
            }
        default:
            var v28 uint64 = 981
            _ = v28
        }
        _ = v20
    }
    if (0 == -9223372036854775808) {
        var v29 uint64 = 2
        v4 = -2.25
        for fi30 := 0; fi30 < 11; fi30++ {
            if (!((((v1 | v1) <= v1) && false) && (!true))) {
                var v31 int = (fi30 & fi30)
                _ = v31
            } else {
                _ = len(v5)
                var v32 uint = 255
                var v33 bool = true
                _ = v32
                _ = v33
            }
            v29 -= (v29 & 30)
            switch ("foo bar" + "") {
            case "foo bar":
                var v34 uint32 = (v15 + genFn0(v15, v29))
                _ = v34
            case "a":
                for fi35 := 0; fi35 < 9; fi35++ {
                    var v36 string = ("foo bar" + "a")
                    if (v29 != 646) {
                        v29 = 4294967295
                    }
                    v3[(((v36 + v36) + "") + ("hello" + v36))] = (fi30 % ((fi35) | 1))
                    _ = fi35
                    _ = v36
                }
            case "hello":
                for fi37 := 0; fi37 < 1; fi37++ {
                    _ = len(v5)
                    _ = fi37
                }
            }
            _ = fi30
        }
        _ = v29
    }
    v15 -= genFn0((v15 ^ 0), 173)
    for fi38 := 0; fi38 < 2; fi38++ {
        v4 -= v4
        v39 := GenS4{F0: 42, F1: -467}
        _ = fi38
        _ = v39
    }
    fmt.Println(v1, v16, v4)
    for fi40 := 0; fi40 < 4; fi40++ {
        var v41 int = (fi40 ^ -1)
        _ = fi40
        _ = v41
    }
    v42 := GenS6{F0: v1, F1: -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
}
