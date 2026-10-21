package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint; F1 bool }
func genFn0(p0 bool, p1 uint16) int64 {
    switch ("hello" + "hello") {
    case "foo bar":
        p1 -= 1
    case "hello":
        if (("" + "hello") < "\n\t") {
            v5 := GenS1{F0: 103, F1: -254}
            _ = v5
        } else {
            for fi6 := 0; fi6 < 16; fi6++ {
                switch "hello" {
                case "foo bar":
                    p1 = 255
                case "hello":
                    p0 = p0
                default:
                    v7 := GenS1{F0: 127, F1: (fi6 >> ((65535) & 63))}
                    _ = v7
                }
                v8 := []int{361, fi6, -275, -1}
                _ = fi6
                _ = v8
            }
            p0 = true
            if (4294967295 == 0) {
                if (p0 || (128 != 128)) {
                    fmt.Println(p0, p1)
                    var v9 uint8 = 42
                    for fi10 := 0; fi10 < 11; fi10++ {
                        p1 += p1
                        _ = fi10
                    }
                    _ = v9
                } else {
                    switch ("" + "") {
                    case "\n\t":
                        switch 4294967295 {
                        case -9223372036854775808:
                            fmt.Println(p1, p0)
                        }
                    default:
                        switch ("" + "hello") {
                        case "":
                            var v11 uint8 = 128
                            _ = v11
                        }
                    }
                }
            } else {
                switch ("hello" + ("a" + "")) {
                case "foo bar":
                    fmt.Println(p0, p0)
                case "":
                    v12 := []int{128, 640, -2147483648}
                    _ = v12
                default:
                    for fi13 := 0; fi13 < 17; fi13++ {
                        v14 := GenS1{F0: 127, F1: (fi13 / ((fi13) | 1))}
                        p0 = ((p0 && ((42 > 9223372036854775807) || false)) || (0.0 < 0.0))
                        v15 := []int{(fi13 & fi13), (fi13 ^ fi13)}
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                }
                p0 = p0
            }
        }
    }
    var v16 uint32 = 42
    switch -1 {
    case 122:
        v17 := []int{}
        _ = v17
    default:
        v18 := []int{}
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v16
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = "a"
    var v3 uint16 = 0
    v4 := make(map[uint32]bool)
    for fi19 := 0; fi19 < 6; fi19++ {
        for fi20 := 0; fi20 < 16; fi20++ {
            fmt.Println(v2)
            var v21 float32 = 1.5
            var v22 uint16 = v3
            fmt.Println(fi19)
            _ = fi20
            _ = v21
            _ = v22
        }
        if true {
            v3 -= 42
            if (!(!false)) {
                for fi23 := 0; fi23 < 9; fi23++ {
                    switch fi19 {
                    case 2147483647:
                        var v24 uint32 = 65535
                        _ = v24
                    case -2147483648:
                        var v25 uint32 = 65535
                        _ = v25
                    case -1:
                        var v26 uint64 = 303
                        _ = v26
                    }
                    _ = fi23
                }
                switch v2 {
                case "foo bar":
                    var v27 string = "hello"
                    _ = v27
                case "":
                    var v28 bool = (v3 != 42)
                    _ = v28
                case "a":
                    var v29 uint16 = (v3 % ((v3) | 1))
                    _ = v29
                }
                var v30 int32 = -1
                _ = v30
            }
        }
        v31 := GenS2{F0: 255, F1: ((v3 != v3) || false)}
        _ = fi19
        _ = v31
    }
    for fi32 := 0; fi32 < 11; fi32++ {
        v33 := []int{(fi32 % ((-654) | 1)), (fi32 * (fi32 - -39)), (fi32 + (fi32 * 483)), -1}
        _ = fi32
        _ = v33
    }
    fmt.Println(v3, v3)
    fmt.Println(v3, v3)
    v34 := GenS1{F0: -128, F1: 42}
    v35 := GenS1{F0: 67, F1: 127}
    v3 += v3
    v2 = v2
    v3 = (^(v3))
    v3 += 255
    v3 = (v3 << ((0) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
