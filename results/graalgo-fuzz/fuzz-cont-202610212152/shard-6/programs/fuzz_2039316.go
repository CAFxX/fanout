package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 bool }
type GenS2 struct { F0 int16 }
func genFn0() uint16 {
    var v7 uint8 = 255
    if true {
        v8 := []int{}
        fmt.Println(v7)
        for fi9 := 0; fi9 < 11; fi9++ {
            fmt.Println(v7, fi9)
            for fi10 := 0; fi10 < 17; fi10++ {
                v7 = v7
                switch "hello" {
                case "foo bar":
                    if true {
                        v7 -= v7
                        fmt.Println(v7, v7, fi9)
                        v7 = (^(v7))
                    }
                }
                _ = fi10
            }
            v11 := GenS1{F0: true, F1: 250, F2: (65535 != 65535)}
            _ = fi9
            _ = v11
        }
        _ = v8
    }
    if false {
        var v12 uint = 4294967295
        _ = v12
    }
    switch "a" {
    case "a":
        if true {
            var v13 float32 = 1.5
            v14 := GenS1{F0: false, F1: -313, F2: (-2147483648 > 42)}
            _ = v13
            _ = v14
        } else {
            v15 := []int{1, 541}
            var v16 uint64 = 18446744073709551615
            _ = v15
            _ = v16
        }
    case "":
        var v17 float32 = -2.25
        _ = v17
    }
    _ = v7
    return uint16(v7)
}
func genFn1(p0 bool, p1 uint32, p2 int16) int {
    fmt.Println(p2, p0, p0)
    if p0 {
        v18 := []int{2147483647, 42, 354, 9223372036854775807}
        p2 = -32768
        _ = v18
    } else {
        p0 = p0
    }
    fmt.Println(p2, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    return 32768
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = true
    var v3 int8 = 127
    var v4 uint8 = 128
    var v5 float32 = -2.25
    var v6 float32 = v5
    v5 -= 3.14159
    v19 := []int{}
    var v20 string = ""
    v21 := GenS1{F0: ((false || v2) && v2), F1: -1, F2: (v2 || v2)}
    for fi22 := 0; fi22 < 4; fi22++ {
        if v2 {
            var v23 uint32 = 600
            _ = v23
        }
        fmt.Println(v3)
        _ = fi22
    }
    if true {
        for fi24 := 0; fi24 < 15; fi24++ {
            for fi25 := 0; fi25 < 11; fi25++ {
                v2 = (255 >= 241)
                v26 := GenS1{F0: v2, F1: 0, F2: true}
                var v27 uint64 = 4294967295
                _ = fi25
                _ = v26
                _ = v27
            }
            _ = fi24
        }
        v2 = (!(65535 > 0))
        fmt.Println(v5)
    } else {
        for fi28 := 0; fi28 < 17; fi28++ {
            v4 += 22
            v5 -= v6
            v4 += 2
            _ = fi28
        }
        fmt.Println(v4)
    }
    v20 = "hello"
    var v29 uint = 139
    var v30 int8 = v3
    var v31 int32 = -2147483648
    v31 = v31
    switch 9223372036854775807 {
    case -310:
        if v2 {
            for fi32 := 0; fi32 < 11; fi32++ {
                var v33 int = fi32
                _ = fi32
                _ = v33
            }
        }
    case 1:
        var v34 bool = v2
        _ = v34
    default:
        for fi35 := 0; fi35 < 17; fi35++ {
            v36 := GenS1{F0: ((2 < 1) || true), F1: -1, F2: (255 >= v29)}
            v4 -= v4
            if (42 != 0) {
                v37 := v1[uint32(v29)]
                for fi38 := 0; fi38 < 16; fi38++ {
                    var v39 uint16 = 65535
                    switch v20 {
                    case "hello":
                        switch genFn1((((v29 << ((v29) & 63)) <= (^(v29))) && false), 0, 32767) {
                        case -880:
                            v3 = 42
                        case 4294967295:
                            delete(v1, 65535)
                        default:
                            var v40 float64 = float64((v5 - v5))
                            _ = v40
                        }
                        v2 = v2
                        if (true || ((2147483647 > 65535) && v2)) {
                            _ = len(v19)
                            fmt.Println(v5, v2, v30)
                            v39 += 42
                        }
                    case "foo bar":
                        v41 := GenS2{F0: 42}
                        _ = v41
                    case "":
                        fmt.Println(v5, v30, v5)
                    default:
                        var v42 int32 = int32(v29)
                        _ = v42
                    }
                    fmt.Println(v31, v2, v39)
                    _ = fi38
                    _ = v39
                }
                _ = v37
            }
            _ = fi35
            _ = v36
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
