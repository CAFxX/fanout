package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 string = ""
    var v2 float32 = float32(float64(0.1))
    v3 := make(map[uint32]bool)
    var v4 float64 = 0.1
    v5 := make(map[int]bool)
    v2 = 3.14159
    v6 := GenS1{F0: 1000, F1: 2}
    v1 = ((("foo bar" + "\n\t") + v1) + "foo bar")
    v2 = v2
    fmt.Println(v1, v1)
    fmt.Println(v2, v1)
    fmt.Println(v1, v2, v1)
    if false {
        switch v1 {
        case "a":
            fmt.Println(v4)
        case "hello":
            for fi7 := 0; fi7 < 2; fi7++ {
                v8 := GenS2{F0: 18446744073709551615}
                for fi9 := 0; fi9 < 3; fi9++ {
                    var v10 uint16 = 0
                    v4 -= v4
                    _ = fi9
                    _ = v10
                }
                fmt.Println(v1)
                _ = fi7
                _ = v8
            }
        case "\n\t":
            v4 = (v4 / ((v4 / 1.5) + v4))
        }
        v11 := []int{0, 32767}
        _ = v11
    } else {
        var v12 float64 = v4
        for fi13 := 0; fi13 < 2; fi13++ {
            if true {
                fmt.Println(v2, fi13)
                fmt.Println(v4)
            } else {
                if (!true) {
                    v14 := GenS1{F0: -1, F1: 65535}
                    var v15 int8 = int8(fi13)
                    _ = v14
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 1; fi16++ {
                        v2 = 3.14159
                        v17 := GenS1{F0: 0, F1: 255}
                        _ = fi16
                        _ = v17
                    }
                }
            }
            var v18 uint32 = 2147483647
            _ = fi13
            _ = v18
        }
        _ = v12
    }
    v2 += 1.5
    _ = len(v5)
    if false {
        var v19 string = v1
        _ = v19
    }
    fmt.Println(v4)
    _ = len(v3)
    for fi20 := 0; fi20 < 11; fi20++ {
        v21 := GenS1{F0: -314, F1: 65535}
        _ = fi20
        _ = v21
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
