package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 int8; F1 int32 }
func genFn0(p0 uint8) int {
    var v5 int64 = -1
    p0 += 127
    v6 := GenS1{F0: ("hello" + ("" + (("" + "hello") + ""))), F1: 127, F2: 47}
    _ = p0
    _ = v5
    _ = v6
    return 1
}
func genFn1() int {
    var v7 uint32 = 1
    v7 += 304
    v8 := GenS1{F0: (("" + "a") + "a"), F1: 9223372036854775807, F2: 229}
    _ = v7
    _ = v8
    return -9223372036854775808
}

func main() {
    var v1 uint = 962
    v2 := make(map[uint32]string)
    var v3 int64 = 1
    v4 := []uint32{1, 658, 2}
    v3 = (v3 & -284)
    switch "a" {
    case "foo bar":
        var v9 float32 = -2.25
        _ = v9
    case "a":
        for fi10 := 0; fi10 < 14; fi10++ {
            for fi11 := 0; fi11 < 17; fi11++ {
                if false {
                    switch genFn0(12) {
                    case 42:
                        switch int(v1) {
                        case 128:
                            _ = len(v4)
                        default:
                            fmt.Println(v1, v1, fi10)
                        }
                    case 923:
                        if false {
                            v12 := v2[0]
                            v13 := GenS1{F0: (("hello" + v12) + v12), F1: 128, F2: 128}
                            _ = v12
                            _ = v13
                        }
                    default:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    var v14 uint64 = 137
                    _ = v14
                }
                var v15 uint64 = 2
                if false {
                    var v16 int32 = -2147483648
                    _ = v16
                } else {
                    if true {
                        v1 = (^(v1))
                        v15 = v15
                        var v17 int8 = 0
                        _ = v17
                    }
                    v15 -= v15
                    delete(v2, 438)
                    delete(v2, 42)
                    v1 += (v1 | v1)
                    fmt.Println(fi11, v15)
                }
                v18 := v2[40]
                _ = fi11
                _ = v15
                _ = v18
            }
            var v19 uint = 1
            _ = fi10
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        if false {
            if (237 > 137) {
                fmt.Println(fi20, v3, v1)
                v21 := GenS2{F0: ((-32768 == 1000) || (!true))}
                _ = v21
            } else {
                fmt.Println(v3, v1, v1)
            }
        }
        _ = fi20
    }
    fmt.Println(v1)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v4)
    v22 := GenS1{F0: "", F1: genFn0(137), F2: 42}
    var v23 bool = (2 <= 0)
    if (1 != 32767) {
        var v24 uint64 = 0
        for fi25 := 0; fi25 < 1; fi25++ {
            v26 := GenS2{F0: true}
            switch "a" {
            case "":
                v27 := GenS3{F0: v3, F1: 1000}
                _ = v27
            case "foo bar":
                v28 := GenS1{F0: "", F1: 1, F2: 2}
                _ = v28
            case "\n\t":
                for fi29 := 0; fi29 < 13; fi29++ {
                    v30 := GenS1{F0: "foo bar", F1: genFn0(41), F2: 55}
                    switch -2147483648 {
                    case -757:
                        switch "foo bar" {
                        case "hello":
                            v1 -= v1
                        }
                    case 2147483647:
                        var v31 uint = v1
                        _ = v31
                    case 4294967295:
                        v32 := GenS2{F0: (-120 >= 634)}
                        _ = v32
                    default:
                        var v33 string = "a"
                        _ = v33
                    }
                    v24 = (v24 << ((0) & 63))
                    v3 = v3
                    if (v23 || v23) {
                        fmt.Println(v23, fi25)
                    }
                    _ = fi29
                    _ = v30
                }
            }
            var v34 uint = v1
            _ = fi25
            _ = v26
            _ = v34
        }
        _ = v24
    } else {
        if true {
            var v35 int16 = -32768
            var v36 float64 = 0.0
            _ = v35
            _ = v36
        } else {
            v37 := GenS4{F0: 1, F1: -2147483648}
            var v38 int64 = (v3 ^ 42)
            var v39 bool = v23
            _ = v37
            _ = v38
            _ = v39
        }
        if (!(v23 || (-549 >= -32768))) {
            var v40 int8 = 0
            v23 = (v40 >= (v40 << ((456) & 63)))
            _ = v40
        }
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
