package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 bool }
type GenS2 struct { F0 uint16; F1 int32 }
func genFn0(p0 int) uint16 {
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 float32 = 1.5
        _ = fi5
        _ = v6
    }
    p0 -= p0
    if (842 != 397) {
        v7 := []int{842, p0, p0}
        _ = v7
    }
    p0 = p0
    v8 := []int{p0, 32767, (p0 ^ (p0 ^ (p0 & p0))), 128}
    _ = p0
    _ = v8
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]string)
    v3 := make(map[uint32]string)
    v4 := []int{}
    v9 := v3[42]
    v10 := GenS1{F0: 42, F1: 1, F2: (18446744073709551615 >= 4294967295)}
    v9 = (("a" + v9) + ("foo bar" + (("\n\t" + ("hello" + ("" + v9))) + v9)))
    if ((9223372036854775807 != 9223372036854775807) && true) {
        v9 = "foo bar"
    } else {
        var v11 bool = true
        v12 := v2[9223372036854775807]
        _ = v11
        _ = v12
    }
    v3[2147483647] = ""
    fmt.Println(v9)
    v13 := v3[247]
    v14 := GenS1{F0: 127, F1: 18446744073709551615, F2: true}
    fmt.Println(v13, v9)
    if len(v4) > 0 {
        v4[0] = 127
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        if false {
            v16 := GenS2{F0: 538, F1: 67}
            v17 := GenS1{F0: 25, F1: 4294967295, F2: ((3.14159 * 3.14159) >= (3.14159 * 1.5))}
            v9 = v13
            _ = v16
            _ = v17
        }
        _ = fi15
    }
    v9 = ((v9 + (((v13 + (v13 + "foo bar")) + v13) + v9)) + (("a" + v9) + (((v13 + v9) + "foo bar") + ("foo bar" + v9))))
    switch 128 {
    case 128:
        if ((!(-1 <= -49)) && true) {
            var v18 float32 = 3.14159
            fmt.Println(v9)
            for fi19 := 0; fi19 < 18; fi19++ {
                var v20 int = -936
                fmt.Println(v20, fi19)
                _ = fi19
                _ = v20
            }
            _ = v18
        }
    case 42:
        fmt.Println(v13, v13)
    case 0:
        v9 = "a"
    }
    v13 = v9
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
