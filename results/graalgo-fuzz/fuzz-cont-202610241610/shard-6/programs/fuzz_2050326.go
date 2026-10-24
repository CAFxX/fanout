package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int8, p1 int64) uint16 {
    p0 = p0
    v6 := []int{-581, 32768, 127}
    v6 = append(v6, -663)
    _ = p0
    _ = p1
    _ = v6
    return 65535
}
func genFn1(p0 int8, p1 string, p2 uint) uint32 {
    if false {
        fmt.Println(p1, p2)
    } else {
        v7 := []int{}
        var v8 float32 = 3.14159
        _ = v7
        _ = v8
    }
    var v9 uint32 = 0
    if false {
        if (!(!true)) {
            v10 := []int{}
            switch "foo bar" {
            case "hello":
                v10 = append(v10, 4294967295)
            case "a":
                p0 = (^(p0))
            default:
                v9 = (v9 % ((v9) | 1))
            }
            _ = v10
        } else {
            var v11 int32 = 42
            if (!(false || false)) {
                v12 := []int{32767}
                _ = v12
            }
            _ = v11
        }
    }
    p2 -= (p2 / ((p2) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (v9 / ((v9) | 1))
}

func main() {
    v1 := []uint32{42}
    var v2 bool = false
    v3 := []string{"a"}
    var v4 int = 42
    var v5 uint16 = 32767
    var v13 int32 = 1
    var v14 string = "hello"
    v14 = "foo bar"
    v15 := GenS1{F0: 1000}
    v16 := GenS1{F0: 42}
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 uint8 = 127
        v3 = append(v3, v14)
        _ = fi17
        _ = v18
    }
    var v19 uint32 = 0
    v5 += uint16(v19)
    v4 = (v4 ^ v4)
    for fi20 := 0; fi20 < 7; fi20++ {
        for fi21 := 0; fi21 < 11; fi21++ {
            if v2 {
                switch fi20 {
                case 0:
                    v3 = append(v3, "")
                default:
                    _ = len(v1)
                    var v22 int = (^(fi21))
                    _ = v22
                }
                var v23 float64 = (0.0 * 3.14159)
                _ = v23
            }
            _ = fi21
        }
        v13 += v13
        v13 = v13
        _ = fi20
    }
    v19 += (v19 / ((v19) | 1))
    fmt.Println(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
