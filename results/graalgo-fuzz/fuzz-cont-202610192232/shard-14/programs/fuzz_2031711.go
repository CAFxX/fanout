package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint16, p1 int8, p2 uint) string {
    if (!((!(255 < 127)) && false)) {
        var v5 uint = p2
        if (!(2147483647 <= -2147483648)) {
            p2 += v5
        } else {
            for fi6 := 0; fi6 < 5; fi6++ {
                v7 := []int{fi6, fi6, (fi6 + fi6), -2147483648}
                fmt.Println(v5, p1)
                _ = fi6
                _ = v7
            }
        }
        if true {
            if (42 > 42) {
                var v8 uint32 = 0
                _ = v8
            }
            p0 += (p0 << (((p2 - (^(p2)))) & 63))
            switch "foo bar" {
            case "foo bar":
                v9 := []int{}
                _ = v9
            case "hello":
                var v10 float64 = (((0.0 + ((-2.25 * ((1.5 + 1e10) + ((3.14159 + float64(1.5)) + 0.0))) * (0.1 + float64((1.5 - 3.14159))))) - (-2.25 - 0.1)) + 3.14159)
                _ = v10
            case "":
                v5 = p2
            default:
                if (1 >= 0) {
                    v11 := []int{-2147483648, 128}
                    if false {
                        fmt.Println(p2, p0, p1)
                    }
                    _ = v11
                } else {
                    v5 -= 18446744073709551615
                    v12 := GenS1{F0: 255, F1: 50}
                    _ = v12
                }
            }
        } else {
            for fi13 := 0; fi13 < 7; fi13++ {
                fmt.Println(p2)
                for fi14 := 0; fi14 < 18; fi14++ {
                    if false {
                        var v15 int16 = 32767
                        v5 = v5
                        p1 = p1
                        _ = v15
                    } else {
                        p0 = p0
                        if true {
                            p0 += (^(p0))
                        } else {
                            v5 -= (p2 % (((p2 * 1)) | 1))
                            p1 = -1
                        }
                        var v16 uint8 = 127
                        _ = v16
                    }
                    _ = fi14
                }
                _ = fi13
            }
        }
        _ = v5
    } else {
        if (-1 < 42) {
            p1 -= p1
            v17 := GenS1{F0: p0, F1: 128}
            _ = v17
        }
        fmt.Println(p0, p1, p0)
        if (false && true) {
            fmt.Println(p0, p0, p1)
        } else {
            var v18 bool = true
            v19 := []int{}
            v20 := GenS1{F0: 0, F1: 1}
            _ = v18
            _ = v19
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        var v22 int64 = -707
        _ = fi21
        _ = v22
    }
    v23 := GenS1{F0: p0, F1: 2}
    var v24 float64 = (-2.25 - float64(((1.5 * float32((1e10 - 0.0))) / 1.0)))
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    _ = v24
    return (((((("a" + ("\n\t" + ("\n\t" + ""))) + ("" + "")) + "a") + ("foo bar" + "a")) + ("hello" + ("a" + ""))) + "\n\t")
}
func genFn1(p0 uint64, p1 int64, p2 bool) string {
    p2 = (!false)
    switch 32767 {
    case 2147483647:
        v25 := []int{4294967295, 4294967295, -9223372036854775808, -1}
        _ = v25
    default:
        var v26 int64 = p1
        _ = v26
    }
    v27 := GenS1{F0: 65535, F1: 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return "foo bar"
}

func main() {
    v1 := []uint32{}
    var v2 uint64 = 0
    var v3 float32 = (0.0 / (float32(-2.25) / 1.5))
    var v4 string = ("" + "foo bar")
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        v29 := GenS1{F0: 981, F1: 42}
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        fmt.Println(fi30, v4, v3)
        if (false || ((-(fi30)) <= (fi30 >> ((65535) & 63)))) {
            v2 = v2
            v3 += float32(((((3.14159 - ((0.0 / 0.1) + 1e10)) - -2.25) + 3.14159) / 3.14159))
            v1 = append(v1, 0)
        }
        _ = fi30
    }
    switch "foo bar" {
    case "hello":
        v31 := GenS1{F0: 663, F1: 1}
        _ = v31
    case "a":
        fmt.Println(v3, v2, v3)
    case "\n\t":
        v32 := GenS2{F0: (false || false)}
        _ = v32
    }
    fmt.Println(v3)
    fmt.Println(v3, v3)
    switch v4 {
    case "\n\t":
        if true {
            v3 = float32(3.14159)
        }
    default:
        v33 := GenS1{F0: 255, F1: 42}
        _ = v33
    }
    var v34 uint16 = 65535
    var v35 uint32 = 2147483648
    fmt.Println(v3, v3, v3)
    v4 = "\n\t"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
