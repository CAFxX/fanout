package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int16 }

func main() {
    var v1 bool = (820 != 187)
    var v2 uint = 18446744073709551615
    v3 := make(map[string]int)
    if (v1 || (!(83 == 5))) {
        fmt.Println(v2)
    } else {
        v4 := GenS1{F0: 2, F1: 42, F2: -32768}
        for fi5 := 0; fi5 < 17; fi5++ {
            var v6 uint32 = 1
            var v7 float64 = -2.25
            _ = fi5
            _ = v6
            _ = v7
        }
        _ = v4
    }
    _ = len(v3)
    if v1 {
        if v1 {
            v2 -= v2
        } else {
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 uint = (v2 >> ((42) & 63))
                v10 := []int{fi8, fi8}
                _ = fi8
                _ = v10
                _ = v9
            }
            switch 42 {
            case -9223372036854775808:
                switch 42 {
                case -2147483648:
                    v11 := GenS1{F0: 2, F1: 496, F2: -32768}
                    _ = v11
                }
            }
            v12 := GenS1{F0: 32767, F1: 0, F2: 32767}
            _ = v12
        }
        var v13 uint64 = 42
        if v1 {
            var v14 uint16 = 801
            _ = v14
        }
        _ = v13
    }
    v1 = v1
    if (!(!v1)) {
        var v15 uint8 = 1
        v1 = v1
        _ = v15
    } else {
        v16 := GenS1{F0: 255, F1: 1000, F2: -32768}
        _ = v16
    }
    fmt.Println(v1, v1)
    v1 = true
    if v1 {
        v17 := []int{-78, 128}
        v2 += 357
        var v18 bool = ((("" + (("foo bar" + ("a" + "hello")) + "")) + "a") < "")
        _ = v17
        _ = v18
    }
    var v19 string = "hello"
    var v20 float32 = (0.1 - (float32((0.1 + float64(-2.25))) * 1.5))
    var v21 uint32 = 1
    v22 := GenS1{F0: 42, F1: 42, F2: 42}
    fmt.Println(v1)
    switch "hello" {
    case "\n\t":
        v1 = (-1 <= -740)
    case "hello":
        v19 = v19
    case "":
        v20 += 0.0
    }
    delete(v3, v19)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
