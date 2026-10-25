package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint32 }
func genFn0(p0 int, p1 int16, p2 string) uint8 {
    var v7 bool = ((572 <= 42) || (!(0 <= 2)))
    p2 = (p2 + (p2 + p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 199
}
func genFn1(p0 int16, p1 uint16, p2 int) int8 {
    fmt.Println(p1)
    var v8 uint8 = 2
    p0 += (p0 - 1000)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 25
}

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = 42
    var v3 int = 32767
    v4 := make(map[int]int)
    v5 := []string{"\n\t"}
    var v6 string = "hello"
    if (true && ((((128 == 147) && true) && true) || (2 >= 301))) {
        var v9 uint64 = 47
        v10 := GenS1{F0: (v6 + (v6 + ("a" + (v6 + (("a" + "a") + v6)))))}
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: 2}
    fmt.Println(v6, v3)
    fmt.Println(v3, v2, v6)
    var v12 float64 = -2.25
    var v13 uint32 = 1
    v6 = v6
    v4[v3] = -636
    v14 := GenS1{F0: (("" + v6) + v6)}
    switch (v3 >> ((2) & 63)) {
    case 32768:
        fmt.Println(v2, v2)
    }
    var v15 int8 = 46
    switch ((((v6 + (v6 + (v6 + v6))) + "hello") + "\n\t") + v6) {
    case "\n\t":
        v16 := v1[v3]
        _ = v16
    case "hello":
        v17 := GenS3{F0: 128, F1: 4294967295, F2: 42}
        _ = v17
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        var v19 uint8 = 186
        if false {
            switch (v6 + "foo bar") {
            case "":
                var v20 int32 = 42
                _ = v20
            default:
                v3 = (fi18 << ((798) & 63))
            }
        }
        _ = fi18
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
