package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 bool; F2 string }
type GenS3 struct { F0 uint8; F1 int8; F2 bool }
type GenS4 struct { F0 uint; F1 int; F2 uint32 }
func genFn0() int8 {
    v7 := []int{-108, 32768}
    v8 := GenS1{F0: 361}
    _ = v7
    _ = v8
    return -37
}
func genFn1(p0 int8, p1 uint8) uint8 {
    var v9 float64 = 1e10
    v10 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return (p1 * (^(p1)))
}

func main() {
    var v1 float32 = 1.5
    var v2 uint64 = 0
    var v3 int8 = 1
    var v4 uint32 = 274
    var v5 uint16 = 65535
    v6 := make(map[int]int)
    v6[0] = -9223372036854775808
    v3 -= (-(v3))
    v11 := GenS2{F0: false, F1: false, F2: ("" + ("hello" + ("" + "foo bar")))}
    fmt.Println(v5)
    var v12 int32 = 1
    v13 := GenS1{F0: 0}
    var v14 int8 = (-(v3))
    switch "\n\t" {
    case "":
        if (!false) {
            var v15 int = 127
            _ = v15
        } else {
            v16 := []int{-534, int(v2), 42}
            _ = v16
        }
    case "a":
        v1 += float32((1.5 * ((1e10 * 1.5) - float64(1.5))))
    }
    v4 = v4
    v17 := v6[444]
    fmt.Println(v2, v3, v3)
    v14 = (v3 % (((v3 >> ((65535) & 63))) | 1))
    v3 = (v14 << ((541) & 63))
    v12 += (v12 + (v12 + (v12 + v12)))
    v18 := GenS1{F0: 42}
    v3 += 21
    if false {
        switch (("\n\t" + ("hello" + ("foo bar" + "hello"))) + ("\n\t" + "")) {
        case "\n\t":
            v19 := GenS3{F0: uint8(v5), F1: (v3 | (-(v3))), F2: false}
            _ = v19
        case "foo bar":
            for fi20 := 0; fi20 < 12; fi20++ {
                v21 := v6[v17]
                _ = fi20
                _ = v21
            }
        case "hello":
            v22 := GenS1{F0: 4294967295}
            _ = v22
        }
    }
    v23 := GenS4{F0: 1, F1: (v17 | v17), F2: (v4 + v4)}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
