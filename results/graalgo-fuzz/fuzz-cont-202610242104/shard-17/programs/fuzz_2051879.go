package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 uint8; F2 int16 }
func genFn0(p0 int32, p1 bool, p2 bool) uint {
    var v7 uint16 = 65535
    var v8 uint = 42
    switch -2147483648 {
    case -855:
        fmt.Println(v8, p0, p0)
    case 127:
        v8 += 42
    case 570:
        if false {
            v9 := []int{128, 4294967295, 1}
            v8 -= (v8 / ((v8) | 1))
            p1 = false
            _ = v9
        } else {
            if ((3.14159 < (0.0 + (0.0 - 1.5))) || p2) {
                for fi10 := 0; fi10 < 19; fi10++ {
                    v11 := []int{fi10, 127}
                    v12 := GenS1{F0: 18446744073709551615}
                    v13 := GenS2{F0: v8, F1: 255, F2: -1}
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
                v14 := []int{int(v7), -9223372036854775808}
                if p1 {
                    var v15 int8 = -82
                    var v16 float32 = 1.5
                    _ = v15
                    _ = v16
                }
                _ = v14
            }
            p1 = p2
            switch "foo bar" {
            case "hello":
                p1 = p2
            case "\n\t":
                var v17 int32 = (p0 << ((uint(v7)) & 63))
                _ = v17
            }
        }
    }
    var v18 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v7
    _ = v8
    return v8
}

func main() {
    v1 := make(map[int]int)
    var v2 int = 32767
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]string)
    v5 := []string{(("" + "hello") + "hello"), "foo bar", "\n\t"}
    v6 := make(map[uint32]string)
    v19 := GenS1{F0: 921}
    v20 := GenS1{F0: 4294967295}
    v21 := GenS1{F0: 42}
    v22 := GenS2{F0: 42, F1: 2, F2: 42}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v23 int32 = 454
    var v24 int64 = 0
    var v25 uint8 = 127
    var v26 uint = 207
    switch "\n\t" {
    case "foo bar":
        for fi27 := 0; fi27 < 16; fi27++ {
            v23 = (v23 % ((v23) | 1))
            v2 -= v2
            _ = fi27
        }
    case "a":
        switch ("\n\t" + "hello") {
        case "a":
            v2 = (^(v2))
        case "":
            var v28 float32 = 1e10
            _ = v28
        case "hello":
            if len(v5) > 0 {
                v5[0] = "a"
            }
        default:
            var v29 int8 = -1
            _ = v29
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
