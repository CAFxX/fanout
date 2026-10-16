package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
func genFn0() uint {
    var v7 int = 4294967295
    switch "" {
    case "foo bar":
        v7 -= v7
    case "a":
        var v8 uint = 1
        _ = v8
    case "hello":
        v7 = v7
    default:
        switch (v7 >> ((462) & 63)) {
        case -9223372036854775808:
            v9 := []int{v7, v7, v7, (v7 | v7)}
            _ = v9
        case -324:
            var v10 bool = (("\n\t" + "") < (("hello" + (("foo bar" + "hello") + "foo bar")) + (("hello" + "hello") + "\n\t")))
            _ = v10
        case -816:
            if ((1.5 <= (float32((0.0 + 0.1)) - 3.14159)) && ((1 == 990) && true)) {
                fmt.Println(v7)
            }
        }
    }
    if false {
        var v11 uint64 = 2
        for fi12 := 0; fi12 < 18; fi12++ {
            v13 := []int{}
            var v14 string = ""
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = v11
    } else {
        v7 = (v7 >> ((18446744073709551615) & 63))
    }
    _ = v7
    return 42
}

func main() {
    var v1 uint32 = 415
    v2 := []uint32{(^(v1)), 1, (v1 ^ v1)}
    var v3 uint16 = 65535
    v4 := []string{"hello"}
    var v5 float64 = 0.1
    var v6 string = ""
    v2 = append(v2, (v1 / ((403) | 1)))
    v6 = ("foo bar" + v6)
    var v15 uint = 42
    var v16 int8 = 100
    v17 := GenS1{F0: v3}
    var v18 uint32 = v1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v6 = (v6 + (("foo bar" + v6) + v6))
    if (true || false) {
        v1 -= (v18 - (v1 * v1))
    }
    v19 := GenS2{F0: (v1 << (((v15 % ((v15) | 1))) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
