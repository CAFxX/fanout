package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int8; F1 string; F2 bool }
func genFn0() uint8 {
    var v7 string = "hello"
    v7 = ("foo bar" + "a")
    v7 = v7
    v8 := []int{4294967295, 845, 168}
    _ = v7
    _ = v8
    return 248
}

func main() {
    var v1 int8 = 127
    var v2 float32 = (1.5 / float32(0.1))
    v3 := make(map[uint32]int)
    var v4 uint = 65535
    var v5 uint = (v4 >> ((v4) & 63))
    var v6 uint16 = 1
    v2 = (v2 * v2)
    for fi9 := 0; fi9 < 11; fi9++ {
        for fi10 := 0; fi10 < 15; fi10++ {
            v11 := GenS1{F0: -386, F1: 0}
            _ = fi10
            _ = v11
        }
        v1 = (v1 % ((0) | 1))
        var v12 bool = (false && (1 <= 2147483647))
        _ = fi9
        _ = v12
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        v5 -= v4
        _ = fi13
    }
    v14 := []int{2147483647, 1}
    var v15 int64 = 42
    fmt.Println(v1, v5)
    v4 = (^(v4))
    var v16 float64 = 0.1
    var v17 float64 = v16
    if ("" > "\n\t") {
        var v18 uint = v5
        switch -371 {
        case -362:
            var v19 uint64 = uint64(v4)
            _ = v19
        default:
            switch ((("a" + "foo bar") + ("foo bar" + ("hello" + (("\n\t" + "hello") + "")))) + "\n\t") {
            case "hello":
                var v20 int16 = 1
                _ = v20
            case "":
                v16 -= ((v16 + v17) * v16)
            case "\n\t":
                _ = len(v3)
            default:
                var v21 string = "hello"
                _ = v21
            }
        }
        var v22 int = int(v6)
        _ = v18
        _ = v22
    }
    v1 -= v1
    fmt.Println(v1)
    if ((float64(v2) / (float64(1.5) - (v16 + (float64((((1e10 + v2) * v2) * -2.25)) + v16)))) != v16) {
        var v23 string = "\n\t"
        if false {
            v2 += (v2 / v2)
            for fi24 := 0; fi24 < 8; fi24++ {
                delete(v3, 1)
                if ((!((true && false) && true)) || true) {
                    switch (fi24 * (^(fi24))) {
                    case -2147483648:
                        delete(v3, 727)
                    default:
                        if true {
                            v17 += (v17 + (v17 + v16))
                            fmt.Println(v1, v6, v5)
                        }
                    }
                    v1 += v1
                }
                _ = fi24
            }
        } else {
            v3[1] = 2147483647
            switch v23 {
            case "\n\t":
                _ = len(v14)
            case "foo bar":
                var v25 int32 = 2147483647
                _ = v25
            case "hello":
                var v26 float64 = 0.1
                _ = v26
            default:
                switch 675 {
                case 1:
                    switch -205 {
                    case 4294967295:
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                    }
                }
            }
            v15 = v15
        }
        switch 2147483647 {
        case 32768:
            switch -868 {
            case 2147483647:
                switch "hello" {
                case "":
                    _ = len(v3)
                case "hello":
                    if false {
                        switch ("" + "foo bar") {
                        case "":
                            var v27 uint = (v4 + v5)
                            _ = v27
                        }
                        v4 -= 0
                        v2 = (v2 / float32(v16))
                    }
                case "a":
                    switch v23 {
                    case "foo bar":
                        v1 = (v1 >> ((v4) & 63))
                    case "\n\t":
                        v15 += v15
                    case "":
                        var v28 int8 = (v1 ^ (-(v1)))
                        _ = v28
                    default:
                        var v29 uint64 = 4294967295
                        _ = v29
                    }
                }
            case 32767:
                switch 9223372036854775807 {
                case 4294967295:
                    v5 = v5
                case 127:
                    switch v23 {
                    case "foo bar":
                        v30 := GenS2{F0: 1, F1: v23}
                        _ = v30
                    }
                case 501:
                    if (1 <= (v15 << ((uint(v6)) & 63))) {
                        var v31 uint32 = 0
                        fmt.Println(v1)
                        var v32 int8 = v1
                        _ = v31
                        _ = v32
                    }
                default:
                    for fi33 := 0; fi33 < 18; fi33++ {
                        switch (^(fi33)) {
                        case -9223372036854775808:
                            v34 := GenS1{F0: (v15 + v15), F1: -477}
                            _ = v34
                        }
                        switch "foo bar" {
                        case "hello":
                            v15 = v15
                        default:
                            v35 := v3[4294967295]
                            _ = v35
                        }
                        v36 := v3[263]
                        _ = fi33
                        _ = v36
                    }
                }
            }
        case -827:
            v4 += 389
        default:
            switch 32767 {
            case 0:
                switch v23 {
                case "\n\t":
                    v37 := GenS1{F0: v15, F1: -1}
                    _ = v37
                default:
                    v2 -= v2
                }
            default:
                for fi38 := 0; fi38 < 9; fi38++ {
                    v17 = -2.25
                    _ = fi38
                }
            }
        }
        _ = v23
    }
    v39 := GenS3{F0: (v1 ^ (v1 & (^(v1)))), F1: "\n\t", F2: true}
    if (42 != 65535) {
        _ = len(v14)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
