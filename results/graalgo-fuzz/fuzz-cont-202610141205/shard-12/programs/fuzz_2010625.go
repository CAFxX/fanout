package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint16, p1 uint16) int {
    if true {
        for fi4 := 0; fi4 < 5; fi4++ {
            p1 -= 32767
            v5 := []int{9223372036854775807, fi4, fi4}
            var v6 int8 = -128
            _ = fi4
            _ = v5
            _ = v6
        }
        v7 := []int{549, -2147483648, 0, -2147483648}
        _ = v7
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 2147483647
}

func main() {
    var v1 float64 = float64((0.1 / (0.0 + 1e10)))
    v2 := []string{"hello", "a", ("" + "")}
    var v3 uint16 = 2
    if false {
        for fi8 := 0; fi8 < 16; fi8++ {
            for fi9 := 0; fi9 < 13; fi9++ {
                v1 += 3.14159
                v10 := GenS1{F0: -984, F1: -1}
                var v11 uint8 = 108
                v3 += (v3 - 0)
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi8
        }
    } else {
        switch 238 {
        case 127:
            if (true && true) {
                fmt.Println(v1, v1)
            }
        case 128:
            if false {
                if len(v2) > 0 {
                    v2[0] = ""
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                for fi12 := 0; fi12 < 18; fi12++ {
                    for fi13 := 0; fi13 < 3; fi13++ {
                        var v14 int8 = 1
                        _ = len(v2)
                        v2 = append(v2, ("\n\t" + ""))
                        _ = fi13
                        _ = v14
                    }
                    switch "hello" {
                    case "":
                        var v15 uint = 18446744073709551615
                        _ = v15
                    case "foo bar":
                        var v16 uint16 = (^(v3))
                        _ = v16
                    }
                    _ = fi12
                }
            }
        case 4294967295:
            if (!(false || (42 == genFn0(857, (v3 % (((^(v3))) | 1)))))) {
                v17 := GenS2{F0: ""}
                _ = len(v2)
                _ = v17
            }
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        v18 := GenS1{F0: 32767, F1: 42}
        _ = v18
    }
    v1 = v1
    var v19 bool = (127 != 210)
    if len(v2) > 0 {
        v2[0] = ((("a" + "foo bar") + "") + (("\n\t" + "a") + ("foo bar" + "\n\t")))
    }
    var v20 float32 = 3.14159
    v21 := GenS1{F0: 42, F1: 0}
    v20 = (v20 * ((v20 * 0.0) + ((float32(v1) * 1.5) * (v20 + v20))))
    switch "a" {
    case "a":
        if len(v2) > 0 {
            v2[0] = "a"
        }
    }
    _ = len(v2)
    fmt.Println(v20, v1)
    fmt.Println(v20)
    switch 4294967295 {
    case 1:
        var v22 int32 = 1
        _ = v22
    }
    if v19 {
        var v23 uint16 = (^(v3))
        fmt.Println(v23, v19, v1)
        var v24 int16 = 0
        _ = v23
        _ = v24
    }
    v20 -= 0.0
    v25 := GenS3{F0: 842}
    v26 := GenS2{F0: ("" + "")}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v3
}
