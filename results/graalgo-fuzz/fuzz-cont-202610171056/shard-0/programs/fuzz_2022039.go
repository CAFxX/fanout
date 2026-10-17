package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 string }
type GenS2 struct { F0 bool; F1 uint64 }
func genFn0(p0 int8) int32 {
    p0 = 127
    p0 = (p0 & p0)
    var v7 string = ("a" + "a")
    if false {
        if ((("a" + (v7 + v7)) + ("foo bar" + v7)) <= (v7 + v7)) {
            v7 = ""
            var v8 uint64 = 0
            switch -1 {
            case -1:
                v7 = "a"
            case -912:
                var v9 uint = 2
                _ = v9
            }
            _ = v8
        } else {
            v10 := []int{-1, -2147483648}
            switch 42 {
            case 777:
                for fi11 := 0; fi11 < 18; fi11++ {
                    v7 = v7
                    p0 = (^(p0))
                    var v12 uint32 = 2147483647
                    _ = fi11
                    _ = v12
                }
            case 0:
                if (!(!true)) {
                    var v13 string = (("a" + v7) + "foo bar")
                    var v14 float64 = (0.1 / float64(0.1))
                    _ = v13
                    _ = v14
                }
            default:
                var v15 bool = false
                _ = v15
            }
            _ = v10
        }
        v16 := []int{-2147483648, 0}
        v17 := GenS1{F0: 1, F1: 339, F2: (v7 + (v7 + (v7 + "")))}
        _ = v16
        _ = v17
    }
    _ = p0
    _ = v7
    return -675
}

func main() {
    var v1 uint8 = 1
    var v2 uint64 = 187
    v3 := make(map[int]uint32)
    var v4 uint8 = v1
    v5 := make(map[int]bool)
    var v6 uint = 255
    v2 += 2
    var v18 float32 = 0.0
    v6 = (^(v6))
    var v19 float32 = float32(1e10)
    var v20 int16 = 1000
    v21 := GenS1{F0: 4294967295, F1: (^(v20)), F2: ("foo bar" + "\n\t")}
    v22 := GenS2{F0: true, F1: v2}
    var v23 int = 65
    if true {
        v4 += 206
    }
    v20 -= v20
    _ = len(v3)
    v23 += (^(v23))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
