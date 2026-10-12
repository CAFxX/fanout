package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 bool }
type GenS2 struct { F0 int8; F1 uint16; F2 int64 }
func genFn0(p0 int, p1 uint32, p2 uint8) int32 {
    var v7 int = p0
    v8 := []int{p0, p0}
    var v9 int64 = -88
    if (true || (-2 > 127)) {
        switch "a" {
        case "foo bar":
            if false {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            } else {
                fmt.Println(p2, v9, p1)
                p1 = 0
                for fi10 := 0; fi10 < 10; fi10++ {
                    var v11 uint16 = 255
                    v12 := GenS1{F0: (386 <= (^(v9))), F1: (v11 >> ((uint(p2)) & 63)), F2: false}
                    v13 := GenS2{F0: -128, F1: (^(v11)), F2: (v9 & 1)}
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
            }
        default:
            p2 -= (p2 ^ p2)
        }
        v9 = v9
    }
    if true {
        switch ("\n\t" + "hello") {
        case "":
            var v14 int64 = 9223372036854775807
            _ = v14
        }
        var v15 int16 = 1
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 int16 = 581
    var v2 uint32 = 570
    v3 := []uint32{v2, (^(v2)), v2, 2147483648}
    v4 := []string{"a", (("hello" + ("\n\t" + ("a" + "hello"))) + ""), "foo bar", ("\n\t" + "a")}
    var v5 string = "foo bar"
    v6 := make(map[int]int)
    v6[102] = 4294967295
    var v16 int16 = -1
    v17 := GenS1{F0: true, F1: 65535, F2: true}
    for fi18 := 0; fi18 < 15; fi18++ {
        v19 := GenS2{F0: 1, F1: 42, F2: 9223372036854775807}
        _ = fi18
        _ = v19
    }
    var v20 uint16 = 840
    var v21 int16 = v1
    v21 = (v1 * (v1 / ((1) | 1)))
    var v22 int = -9223372036854775808
    for fi23 := 0; fi23 < 1; fi23++ {
        v3 = append(v3, v2)
        _ = fi23
    }
    v22 += 32768
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
