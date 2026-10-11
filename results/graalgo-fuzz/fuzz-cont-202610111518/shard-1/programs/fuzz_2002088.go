package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 int32; F2 string }
func genFn0(p0 int, p1 uint32) uint {
    var v4 float64 = ((0.1 * -2.25) * 3.14159)
    var v5 uint = 65535
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return v5
}
func genFn1() int {
    var v6 uint16 = 1
    for fi7 := 0; fi7 < 6; fi7++ {
        if false {
            v8 := []int{349}
            _ = v8
        }
        v6 = (v6 / ((v6) | 1))
        _ = fi7
    }
    _ = v6
    return 32767
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[string]int)
    var v3 int = 42
    v1 = (v1 | v1)
    v9 := []int{v3, (-(v3))}
    v3 = (v3 / ((42) | 1))
    if true {
        switch genFn1() {
        case -2147483648:
            v3 -= 0
        case 684:
            var v10 uint16 = 65535
            _ = v10
        default:
            var v11 uint16 = 317
            _ = v11
        }
        for fi12 := 0; fi12 < 12; fi12++ {
            var v13 uint32 = 42
            _ = fi12
            _ = v13
        }
        var v14 float32 = 1.5
        _ = v14
    }
    var v15 uint32 = 65535
    v1 += v1
    var v16 uint8 = 123
    v15 += v15
    for fi17 := 0; fi17 < 11; fi17++ {
        var v18 uint64 = v1
        for fi19 := 0; fi19 < 15; fi19++ {
            if true {
                switch ((((("foo bar" + "hello") + "foo bar") + (("foo bar" + "") + ("hello" + ("hello" + "hello")))) + "") + "\n\t") {
                case "hello":
                    _ = len(v2)
                case "":
                    var v20 int = (fi17 ^ (fi19 * int(v1)))
                    _ = v20
                }
            } else {
                if false {
                    v1 -= 309
                } else {
                    v9 = append(v9, (fi19 | fi19))
                    v21 := GenS1{F0: 42}
                    v22 := GenS2{F0: 32767, F1: 42, F2: "\n\t"}
                    fmt.Println(fi17, v3)
                    _ = v21
                    _ = v22
                }
                v23 := GenS2{F0: 799, F1: 1, F2: ""}
                _ = v23
            }
            fmt.Println(fi17, fi19)
            _ = fi19
        }
        var v24 float64 = (0.0 - float64(0.1))
        _ = fi17
        _ = v18
        _ = v24
    }
    var v25 int16 = 1
    fmt.Println(v15, v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v9
}
