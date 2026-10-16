package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 string; F2 uint64 }
func genFn0(p0 uint, p1 uint) uint64 {
    v5 := GenS1{F0: ("foo bar" + "foo bar"), F1: 1, F2: (!((711 == 65535) && false))}
    fmt.Println(p0)
    p0 -= p1
    for fi6 := 0; fi6 < 6; fi6++ {
        fmt.Println(p1, p1, fi6)
        for fi7 := 0; fi7 < 11; fi7++ {
            v8 := []int{fi7}
            var v9 bool = (!false)
            var v10 uint32 = 4294967295
            switch (("foo bar" + "") + "foo bar") {
            case "foo bar":
                v8 = append(v8, (^(fi7)))
            default:
                var v11 uint16 = 1
                _ = v11
            }
            var v12 string = ("\n\t" + "a")
            _ = fi7
            _ = v10
            _ = v12
            _ = v8
            _ = v9
        }
        var v13 int32 = -1
        _ = fi6
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v5
    return 1
}
func genFn1() uint8 {
    v14 := []int{2147483647}
    v15 := GenS1{F0: ("\n\t" + (("" + "foo bar") + "foo bar")), F1: -460, F2: true}
    _ = v14
    _ = v15
    return 42
}

func main() {
    var v1 int16 = 1000
    var v2 float32 = (0.1 + 0.0)
    v3 := []int{42, 489, 0}
    var v4 string = "\n\t"
    if (!true) {
        v16 := GenS2{F0: 65535}
        if true {
            v3 = append(v3, -35)
            var v17 uint = 42
            fmt.Println(v1)
            _ = v17
        } else {
            var v18 int32 = -890
            _ = v18
        }
        _ = v16
    } else {
        v1 += v1
    }
    var v19 uint = 4294967295
    if true {
        var v20 bool = false
        v19 = v19
        switch v4 {
        case "a":
            var v21 float64 = ((0.0 / (1.5 / 1.0)) / 1e10)
            _ = v21
        case "\n\t":
            var v22 uint32 = 206
            _ = v22
        case "hello":
            switch -722 {
            case 32768:
                if len(v3) > 0 {
                    v3[0] = 2147483647
                }
            case -821:
                v19 += v19
            case 714:
                v23 := GenS1{F0: v4, F1: 0, F2: true}
                _ = v23
            default:
                v4 = "foo bar"
            }
        }
        _ = v20
    }
    var v24 int16 = (v1 >> ((1) & 63))
    v2 = (1.5 + v2)
    if true {
        if (-28 <= -7) {
            if false {
                for fi25 := 0; fi25 < 1; fi25++ {
                    switch 51 {
                    case 32767:
                        if true {
                            var v26 uint8 = 47
                            var v27 int64 = 450
                            _ = v26
                            _ = v27
                        }
                    default:
                        v4 = "a"
                    }
                    v24 = (-(v1))
                    var v28 string = "foo bar"
                    _ = fi25
                    _ = v28
                }
                var v29 uint16 = 32767
                _ = len(v3)
                _ = v29
            } else {
                v30 := GenS3{F0: genFn1(), F1: "\n\t", F2: 4294967295}
                v31 := GenS1{F0: ("" + ("a" + "\n\t")), F1: -285, F2: true}
                var v32 int = -2147483648
                _ = v30
                _ = v31
                _ = v32
            }
            for fi33 := 0; fi33 < 1; fi33++ {
                v2 = ((3.14159 * ((v2 * (v2 + 1e10)) - float32(0.0))) + 1e10)
                if (1.5 < -2.25) {
                    fmt.Println(v1)
                    var v34 uint = v19
                    _ = v34
                }
                v19 -= (v19 % (((v19 >> (((v19 % ((v19) | 1))) & 63))) | 1))
                _ = fi33
            }
        } else {
            v35 := GenS2{F0: 0}
            v4 = "foo bar"
            v2 += 1.5
            _ = v35
        }
        switch -9223372036854775808 {
        case 0:
            v2 -= v2
        }
        v24 = 909
    }
    _ = len(v3)
    v36 := GenS2{F0: 65535}
    v1 = 32767
    if true {
        v37 := GenS3{F0: 42, F1: ((v4 + "foo bar") + ((v4 + (v4 + "\n\t")) + v4)), F2: genFn0((^(v19)), (v19 * v19))}
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v36
    _ = v4
}
