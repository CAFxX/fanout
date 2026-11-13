package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16; F1 bool }
func genFn0() int64 {
    v4 := []int{}
    var v5 uint32 = 65535
    for fi6 := 0; fi6 < 12; fi6++ {
        fmt.Println(fi6, v5)
        switch "hello" {
        case "a":
            var v7 bool = (true || false)
            _ = v7
        }
        _ = fi6
    }
    v4 = append(v4, -9223372036854775808)
    var v8 string = "hello"
    _ = v4
    _ = v5
    _ = v8
    return -1
}

func main() {
    var v1 float64 = (3.14159 * 1e10)
    var v2 uint8 = 84
    v3 := make(map[string]uint32)
    for fi9 := 0; fi9 < 6; fi9++ {
        fmt.Println(v1, v1, v2)
        var v10 string = ("" + "a")
        v11 := GenS1{F0: false, F1: 127, F2: 42}
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := []int{32767}
    v1 += v1
    for fi13 := 0; fi13 < 3; fi13++ {
        fmt.Println(fi13, v2)
        var v14 int8 = 2
        _ = fi13
        _ = v14
    }
    var v15 int64 = 0
    v16 := GenS1{F0: false, F1: (v2 - (v2 - (v2 | v2))), F2: 1}
    for fi17 := 0; fi17 < 5; fi17++ {
        if (!false) {
            var v18 int16 = 1
            v15 += (v15 >> ((0) & 63))
            if false {
                fmt.Println(fi17, fi17, v2)
                fmt.Println(v15, v15, v2)
            }
            _ = v18
        }
        v3[(("\n\t" + "a") + ("\n\t" + ""))] = 446
        _ = fi17
    }
    switch 367 {
    case -9223372036854775808:
        for fi19 := 0; fi19 < 10; fi19++ {
            for fi20 := 0; fi20 < 9; fi20++ {
                var v21 int32 = 0
                if false {
                    var v22 uint32 = 710
                    _ = v22
                } else {
                    if (42 < 1) {
                        var v23 uint32 = 2147483648
                        if len(v12) > 0 {
                            v12[0] = fi19
                        }
                        if true {
                            v24 := GenS2{F0: "foo bar"}
                            var v25 uint = 561
                            _ = v24
                            _ = v25
                        } else {
                            _ = len(v12)
                        }
                        _ = v23
                    } else {
                        if len(v12) > 0 {
                            v12[0] = -459
                        }
                    }
                    if ((v15 % ((v15) | 1)) >= 9223372036854775807) {
                        var v26 int = fi20
                        _ = v26
                    } else {
                        v27 := GenS2{F0: "hello"}
                        if ((0 < 32767) && (true && (((v21 & 1) == v21) || true))) {
                            v15 = int64(fi20)
                            v28 := GenS3{F0: fi20}
                            _ = v28
                        } else {
                            var v29 uint64 = 0
                            switch "" {
                            case "":
                                var v30 int8 = -1
                                _ = v30
                            case "hello":
                                var v31 uint = 255
                                _ = v31
                            }
                            var v32 float64 = (v1 + v1)
                            _ = v29
                            _ = v32
                        }
                        if true {
                            v2 = v2
                            if true {
                                var v33 uint = 489
                                v34 := GenS4{F0: 1, F1: (true || false)}
                                if (!true) {
                                    v12 = append(v12, -9223372036854775808)
                                    var v35 uint8 = (v2 - 58)
                                    _ = v35
                                }
                                _ = v33
                                _ = v34
                            }
                        }
                        _ = v27
                    }
                }
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
    case -1:
        fmt.Println(v1, v15, v2)
    }
    for fi36 := 0; fi36 < 4; fi36++ {
        v37 := GenS2{F0: "\n\t"}
        switch "\n\t" {
        case "hello":
            var v38 int = (fi36 >> ((65535) & 63))
            _ = v38
        case "a":
            fmt.Println(fi36, v2)
        default:
            for fi39 := 0; fi39 < 16; fi39++ {
                var v40 string = (((((("" + "hello") + "foo bar") + "hello") + "\n\t") + "foo bar") + ("hello" + (((("a" + "") + "\n\t") + ("" + "")) + "hello")))
                var v41 uint16 = 32767
                _ = fi39
                _ = v40
                _ = v41
            }
        }
        v15 += (v15 ^ genFn0())
        _ = fi36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
