package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint64 }
type GenS2 struct { F0 string }
func genFn0(p0 int64, p1 uint16, p2 uint32) int16 {
    var v7 string = ""
    fmt.Println(p2, p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -652
}
func genFn1(p0 uint) uint {
    v8 := []int{-9223372036854775808, -2147483648, -2147483648}
    switch -867 {
    case -537:
        v9 := GenS1{F0: -9223372036854775808, F1: 32767, F2: 476}
        _ = v9
    case -2147483648:
        p0 += p0
    case 128:
        v10 := GenS1{F0: 1, F1: 156, F2: 0}
        _ = v10
    default:
        var v11 uint32 = 571
        _ = v11
    }
    switch -569 {
    case 42:
        p0 = p0
    case -144:
        p0 = (^(p0))
    default:
        var v12 uint = (p0 - p0)
        _ = v12
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 uint8 = 140
        _ = fi13
        _ = v14
    }
    p0 += 42
    _ = p0
    _ = v8
    return p0
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 1
    v3 := make(map[int]uint32)
    v4 := make(map[int]string)
    var v5 bool = false
    v6 := make(map[uint32]int)
    if false {
        var v15 uint8 = 0
        fmt.Println(v2, v2)
        if v5 {
            fmt.Println(v2, v5)
            v16 := []int{-2147483648, 224}
            v16 = append(v16, 42)
            _ = v16
        }
        _ = v15
    } else {
        var v17 uint64 = 18446744073709551615
        if v5 {
            switch "" {
            case "hello":
                for fi18 := 0; fi18 < 3; fi18++ {
                    fmt.Println(v17, v5)
                    for fi19 := 0; fi19 < 7; fi19++ {
                        v20 := v1[188]
                        v17 = v17
                        v20 = "hello"
                        _ = fi19
                        _ = v20
                    }
                    _ = fi18
                }
            }
        } else {
            for fi21 := 0; fi21 < 7; fi21++ {
                fmt.Println(fi21)
                v17 -= 0
                _ = fi21
            }
            var v22 int8 = 3
            _ = v22
        }
        switch 768 {
        case 820:
            var v23 uint = 926
            _ = v23
        case 545:
            v5 = (-296 != 32768)
        }
        _ = v17
    }
    v24 := GenS1{F0: 1, F1: -1, F2: 42}
    v1[282] = "a"
    if v5 {
        for fi25 := 0; fi25 < 13; fi25++ {
            v5 = false
            fmt.Println(v2)
            v2 += 32767
            _ = fi25
        }
    } else {
        v2 += (v2 | (v2 & 0))
        v2 += 1
    }
    v26 := GenS2{F0: "hello"}
    fmt.Println(v2, v2)
    for fi27 := 0; fi27 < 5; fi27++ {
        v28 := GenS2{F0: ((("foo bar" + ("a" + "")) + (("" + "\n\t") + ("hello" + "foo bar"))) + "foo bar")}
        if v5 {
            v5 = v5
        }
        _ = fi27
        _ = v28
    }
    if true {
        var v29 int8 = 97
        _ = v29
    }
    var v30 int16 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
