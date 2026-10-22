package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int }
type GenS2 struct { F0 uint; F1 int16 }
func genFn0(p0 int64, p1 int32) uint32 {
    var v4 float64 = float64(0.1)
    var v5 int = -1
    var v6 int64 = p0
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 uint64) int8 {
    var v7 uint16 = 2
    p0 -= 730
    _ = p0
    _ = v7
    return -45
}

func main() {
    var v1 uint8 = 253
    var v2 uint64 = 530
    v3 := make(map[int]string)
    if false {
        if (1000 != 1) {
            switch "" {
            case "\n\t":
                var v8 string = (((("a" + "foo bar") + ("\n\t" + "foo bar")) + "a") + ("\n\t" + ""))
                _ = v8
            default:
                v9 := []int{4294967295, 32768}
                _ = v9
            }
        } else {
            var v10 uint = 2
            _ = v10
        }
    }
    v11 := GenS1{F0: -1, F1: (("foo bar" >= "") || (("foo bar" > ((("" + "hello") + "foo bar") + "")) && true)), F2: -9223372036854775808}
    var v12 int64 = 285
    var v13 uint16 = 255
    fmt.Println(v1)
    if false {
        v14 := GenS2{F0: 0, F1: -32768}
        _ = v14
    } else {
        var v15 uint16 = v13
        _ = v15
    }
    v16 := GenS1{F0: -1, F1: false, F2: 32768}
    switch "a" {
    case "\n\t":
        v3[32768] = ""
    case "":
        var v17 int8 = 100
        _ = v17
    case "hello":
        v12 -= v12
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        for fi19 := 0; fi19 < 9; fi19++ {
            var v20 int16 = 0
            v21 := []int{(fi19 + fi18), (^(fi18)), (^(fi18))}
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
}
