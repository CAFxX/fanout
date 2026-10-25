package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int32; F2 uint64 }
func genFn0() uint64 {
    var v5 int16 = 32767
    v6 := []int{-196, -747}
    v5 = v5
    switch "" {
    case "a":
        var v7 uint64 = 0
        _ = v7
    }
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[int]string)
    v3 := []int{755, 32767, 0}
    var v4 uint = 342
    var v8 int = 32767
    v8 -= -9223372036854775808
    var v9 int = v8
    fmt.Println(v1, v4)
    if ((("a" + "\n\t") + ("hello" + "foo bar")) <= "foo bar") {
        var v10 float64 = 1e10
        v4 += (v4 % ((v4) | 1))
        _ = v10
    } else {
        switch (v9 >> ((534) & 63)) {
        case 696:
            v11 := GenS1{F0: v9}
            _ = v11
        case 1:
            var v12 float64 = 0.1
            _ = v12
        }
    }
    v13 := GenS2{F0: false, F1: 2147483647, F2: 685}
    var v14 float32 = 0.0
    fmt.Println(v1, v8)
    fmt.Println(v4, v14, v4)
    v1 = v1
    v14 = ((1e10 - v14) + (v14 * -2.25))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v15 int8 = int8(v1)
    var v16 uint8 = 255
    v17 := GenS2{F0: (false && true), F1: -391, F2: 4294967295}
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        var v18 uint8 = v16
        _ = v18
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 bool = (uint16(v4) == 0)
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
