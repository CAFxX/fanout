package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint }
type GenS2 struct { F0 uint; F1 int64; F2 int8 }
func genFn0() uint32 {
    var v4 float64 = 0.1
    if true {
        for fi5 := 0; fi5 < 10; fi5++ {
            v6 := GenS1{F0: (!(1 > -1)), F1: 1, F2: 39}
            fmt.Println(fi5, v4)
            v7 := []int{(fi5 + -319), 4294967295, fi5}
            _ = fi5
            _ = v6
            _ = v7
        }
    } else {
        v8 := []int{-723, -1}
        _ = v8
    }
    var v9 uint16 = 765
    switch -736 {
    case -1:
        v10 := GenS1{F0: (-1 < 32767), F1: 84, F2: 65535}
        _ = v10
    case 9223372036854775807:
        fmt.Println(v4, v9)
    }
    _ = v4
    _ = v9
    return 42
}

func main() {
    var v1 uint32 = 42
    var v2 uint32 = (v1 ^ 0)
    var v3 float32 = 0.0
    v2 -= genFn0()
    v11 := []int{32768}
    var v12 int64 = 9223372036854775807
    v3 = float32(0.0)
    v12 += v12
    v2 = (v1 >> ((65535) & 63))
    v13 := GenS2{F0: 4294967295, F1: 42, F2: 123}
    var v14 int16 = 42
    v14 = 32767
    var v15 bool = ((v3 / float32(1.5)) == (1e10 * 3.14159))
    switch -2147483648 {
    case 128:
        var v16 uint = 65535
        _ = v16
    case 32767:
        fmt.Println(v1, v14)
    case -482:
        switch (("foo bar" + "") + "a") {
        case "hello":
            for fi17 := 0; fi17 < 6; fi17++ {
                var v18 bool = (true || v15)
                v19 := GenS1{F0: v15, F1: -64, F2: 757}
                _ = fi17
                _ = v18
                _ = v19
            }
        }
    }
    v11 = append(v11, 127)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
