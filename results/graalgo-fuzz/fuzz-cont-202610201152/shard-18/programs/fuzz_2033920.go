package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 string }
func genFn0(p0 bool, p1 uint32, p2 string) uint {
    p2 = "hello"
    v6 := []int{2147483647, -58, -3}
    v7 := GenS1{F0: uint8(p1), F1: -2147483648}
    v6 = append(v6, 32767)
    if len(v6) > 0 {
        v6[0] = 4294967295
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 4294967295
}
func genFn1() uint16 {
    v8 := []int{-1}
    var v9 int = 0
    fmt.Println(v9)
    switch ("hello" + ("hello" + ("hello" + ("" + "foo bar")))) {
    case "":
        if ((42 < -820) || true) {
            for fi10 := 0; fi10 < 18; fi10++ {
                for fi11 := 0; fi11 < 13; fi11++ {
                    var v12 uint = 405
                    switch int(v12) {
                    case 9223372036854775807:
                        _ = len(v8)
                    case 32767:
                        var v13 int16 = 518
                        _ = v13
                    case 128:
                        v14 := GenS1{F0: 128, F1: -2147483648}
                        _ = v14
                    }
                    if (-1 != 1) {
                        var v15 bool = (1 >= 2147483647)
                        var v16 uint16 = 1
                        _ = v15
                        _ = v16
                    } else {
                        v17 := GenS1{F0: 127, F1: 42}
                        v18 := GenS1{F0: 42, F1: 42}
                        if true {
                            v12 = v12
                            v9 += v9
                            var v19 int64 = 76
                            _ = v19
                        }
                        v12 = genFn0((1 <= 566), 65535, "a")
                        var v20 float64 = 1e10
                        _ = v17
                        _ = v18
                        _ = v20
                    }
                    _ = fi11
                    _ = v12
                }
                _ = fi10
            }
        } else {
            if (-32768 > -32768) {
                for fi21 := 0; fi21 < 13; fi21++ {
                    v22 := GenS1{F0: 255, F1: -1}
                    v9 = v9
                    _ = fi21
                    _ = v22
                }
                v23 := GenS2{F0: 65535, F1: false}
                _ = v23
            } else {
                var v24 int = v9
                v25 := GenS1{F0: 42, F1: -1}
                if true {
                    var v26 int16 = 32767
                    _ = v26
                } else {
                    var v27 uint16 = 9
                    switch "" {
                    case "hello":
                        v28 := GenS2{F0: 357, F1: (124 == 1)}
                        _ = v28
                    default:
                        v27 = (v27 + 255)
                    }
                    _ = v27
                }
                _ = v24
                _ = v25
            }
        }
    case "foo bar":
        if true {
            var v29 int32 = 1
            v9 -= 127
            _ = v29
        } else {
            fmt.Println(v9)
        }
    }
    for fi30 := 0; fi30 < 1; fi30++ {
        fmt.Println(v9, fi30)
        _ = fi30
    }
    _ = v8
    _ = v9
    return 298
}

func main() {
    var v1 uint = 0
    v2 := make(map[int]bool)
    v3 := make(map[uint32]string)
    var v4 uint8 = 0
    var v5 float32 = -2.25
    v31 := GenS2{F0: 2, F1: ((("" + "") + "foo bar") <= "hello")}
    if false {
        v1 += (v1 / ((genFn0((128 < int(v4)), 915, ("" + ("hello" + ("\n\t" + ("a" + "")))))) | 1))
        v32 := GenS3{F0: "\n\t"}
        v1 = (v1 | v1)
        _ = v32
    }
    for fi33 := 0; fi33 < 8; fi33++ {
        for fi34 := 0; fi34 < 19; fi34++ {
            v35 := []int{}
            switch int(v4) {
            case -2147483648:
                v5 += (v5 / (((float32(((float64((float32(3.14159) + v5)) + (1.5 - float64(float32(1e10)))) * 0.1)) + v5) / v5) / (float32(3.14159) * v5)))
            }
            _ = fi34
            _ = v35
        }
        var v36 float32 = (v5 - 3.14159)
        var v37 uint = (v1 >> ((genFn0(false, 4294967295, "foo bar")) & 63))
        _ = fi33
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 16; fi38++ {
        v3[uint32(v4)] = ""
        _ = fi38
    }
    fmt.Println(v4)
    v39 := []int{-1, int(v4)}
    v4 += v4
    var v40 int = 32767
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
