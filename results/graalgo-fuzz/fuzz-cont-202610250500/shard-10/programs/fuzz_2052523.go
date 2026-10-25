package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
func genFn0(p0 int8) int8 {
    v6 := []int{1}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    p0 = -53
    _ = p0
    _ = v6
    return 23
}

func main() {
    var v1 float32 = (0.1 / float32((1e10 / 1.0)))
    v2 := []string{"\n\t"}
    var v3 uint32 = 4294967295
    v4 := []uint32{(v3 << ((2) & 63))}
    var v5 float64 = 0.0
    v3 += (v3 ^ v3)
    v7 := GenS1{F0: 42, F1: 1}
    if false {
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := GenS1{F0: 18446744073709551615, F1: 2}
            v3 = (v3 | v3)
            v1 = float32((v5 * (v5 / (3.14159 / v5))))
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 18; fi10++ {
            if (true && (!false)) {
                var v11 string = "a"
                _ = v11
            }
            _ = fi10
        }
    }
    var v12 int8 = 127
    var v13 uint16 = 245
    v1 -= v1
    v12 += v12
    var v14 int32 = 0
    var v15 uint8 = 248
    fmt.Println(v12, v5)
    var v16 uint32 = v3
    switch ("a" + ((("\n\t" + ("foo bar" + "a")) + ("hello" + "a")) + "\n\t")) {
    case "foo bar":
        v12 += 0
    }
    switch -1 {
    case 885:
        switch -9223372036854775808 {
        case -2147483648:
            v1 -= v1
        default:
            v1 += 0.1
        }
    case 2147483647:
        var v17 uint32 = (v3 << ((65535) & 63))
        _ = v17
    }
    v5 = (v5 * v5)
    var v18 int = 0
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
