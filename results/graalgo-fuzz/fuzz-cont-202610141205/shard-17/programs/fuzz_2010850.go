package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }

func main() {
    v1 := []uint32{0, 1, 4294967295}
    var v2 int32 = 2147483647
    var v3 uint = 869
    v4 := make(map[int]uint32)
    var v5 string = "hello"
    v6 := []int{1, -901, -45, 1}
    v7 := GenS1{F0: 255, F1: -32768}
    switch v5 {
    case "hello":
        var v8 uint8 = 2
        _ = v8
    case "\n\t":
        var v9 float64 = 1e10
        _ = v9
    case "":
        if false {
            fmt.Println(v5, v3)
        } else {
            var v10 float32 = float32(((3.14159 + float64(3.14159)) / 1.5))
            _ = v10
        }
    }
    fmt.Println(v5)
    var v11 int32 = (v2 >> ((v3) & 63))
    fmt.Println(v3)
    delete(v4, -187)
    _ = len(v1)
    if ((v5 + "foo bar") <= "hello") {
        for fi12 := 0; fi12 < 14; fi12++ {
            v13 := GenS1{F0: 255, F1: 554}
            _ = fi12
            _ = v13
        }
        v2 = v11
        v14 := GenS1{F0: 32767, F1: -570}
        _ = v14
    } else {
        var v15 int64 = 9223372036854775807
        var v16 string = (v5 + (v5 + "\n\t"))
        for fi17 := 0; fi17 < 9; fi17++ {
            v3 -= 255
            fmt.Println(v2, v16, v3)
            _ = fi17
        }
        _ = v15
        _ = v16
    }
    v3 = (v3 * 65535)
    var v18 uint = v3
    var v19 bool = false
    if len(v6) > 0 {
        v6[0] = 648
    }
    v5 = (v5 + (v5 + v5))
    v3 = v3
    var v20 float32 = 1.5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
