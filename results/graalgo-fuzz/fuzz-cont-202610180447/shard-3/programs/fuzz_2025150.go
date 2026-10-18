package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 string; F1 int64; F2 uint }
func genFn0(p0 bool, p1 int) uint16 {
    if p0 {
        if (416 <= 1) {
            var v5 string = ""
            var v6 float64 = (0.0 / 1.0)
            for fi7 := 0; fi7 < 14; fi7++ {
                v8 := GenS1{F0: 0, F1: 238}
                fmt.Println(p1, fi7)
                v9 := GenS1{F0: 2147483647, F1: 1}
                _ = fi7
                _ = v8
                _ = v9
            }
            _ = v5
            _ = v6
        }
        if p0 {
            v10 := []int{p1, (p1 + (-(p1)))}
            fmt.Println(p0, p0)
            if p0 {
                v11 := GenS1{F0: 1, F1: 42}
                v10 = append(v10, (p1 - (p1 << ((2) & 63))))
                _ = v11
            } else {
                p0 = (("a" + "") >= (("foo bar" + ((("hello" + ("foo bar" + (("foo bar" + ("\n\t" + "a")) + ("\n\t" + ("" + "a"))))) + ("foo bar" + "foo bar")) + ("foo bar" + "hello"))) + (("" + "a") + "foo bar")))
                if p0 {
                    p0 = (255 == 838)
                    v12 := GenS1{F0: 462, F1: -1}
                    switch "" {
                    case "hello":
                        if true {
                            var v13 uint64 = 0
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            _ = v13
                        }
                    case "":
                        fmt.Println(p1, p0)
                    case "a":
                        var v14 int16 = -1
                        _ = v14
                    }
                    _ = v12
                } else {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    p0 = p0
                    var v15 int = 1
                    _ = v15
                }
                var v16 int32 = 1
                _ = v16
            }
            _ = v10
        }
    }
    v17 := GenS1{F0: -343, F1: 0}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v17
    return 307
}
func genFn1(p0 string, p1 uint32, p2 uint8) int32 {
    for fi18 := 0; fi18 < 12; fi18++ {
        v19 := GenS1{F0: -2147483648, F1: 542}
        var v20 bool = ((false && (381 == 0)) || true)
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 int64 = 558
    v22 := GenS1{F0: 642, F1: (v21 + v21)}
    v23 := GenS1{F0: int32(p1), F1: (v21 - (v21 >> ((305) & 63)))}
    v24 := GenS1{F0: 2147483647, F1: (v21 % ((v21) | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    return 345
}

func main() {
    var v1 string = "foo bar"
    var v2 uint8 = 53
    v3 := make(map[int]bool)
    v4 := make(map[string]string)
    v1 = "\n\t"
    switch 128 {
    case 0:
        fmt.Println(v1, v2)
    case 32767:
        switch -2147483648 {
        case 32768:
            fmt.Println(v2, v2)
        }
    default:
        var v25 uint16 = 32767
        _ = v25
    }
    for fi26 := 0; fi26 < 10; fi26++ {
        v27 := GenS2{F0: v2, F1: -775}
        _ = fi26
        _ = v27
    }
    delete(v4, (v1 + (v1 + "\n\t")))
    var v28 uint64 = 235
    switch "" {
    case "\n\t":
        fmt.Println(v1, v1)
    case "a":
        fmt.Println(v1)
    case "":
        v29 := []int{-909}
        _ = v29
    }
    switch v1 {
    case "a":
        var v30 int = -886
        _ = v30
    case "\n\t":
        switch "foo bar" {
        case "foo bar":
            for fi31 := 0; fi31 < 19; fi31++ {
                v1 = v1
                for fi32 := 0; fi32 < 16; fi32++ {
                    v33 := []int{(fi32 - fi31), fi31, fi31}
                    v34 := GenS3{F0: v1, F1: 0, F2: 1}
                    v2 += 127
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                _ = fi31
            }
        case "\n\t":
            v3[32768] = true
        case "":
            var v35 int8 = 127
            _ = v35
        }
    case "foo bar":
        for fi36 := 0; fi36 < 12; fi36++ {
            v1 = "foo bar"
            _ = fi36
        }
    }
    var v37 uint = 848
    v38 := v4[v1]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v37
    _ = v38
    _ = v4
}
