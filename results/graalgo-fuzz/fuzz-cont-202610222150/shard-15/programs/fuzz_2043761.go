package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 string; F1 uint32; F2 int64 }
type GenS5 struct { F0 int64; F1 uint16 }
func genFn0() int8 {
    v7 := []int{2147483647, 629, 42, 32767}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v8 int64 = -1
    if (!false) {
        var v9 uint16 = 0
        _ = v9
    }
    v10 := GenS1{F0: "\n\t"}
    _ = v10
    _ = v7
    _ = v8
    return 127
}

func main() {
    v1 := []int{2147483647}
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 597
    var v4 int = -550
    var v5 uint64 = 2
    v6 := []int{598}
    var v11 string = "foo bar"
    v12 := GenS2{F0: ((("foo bar" + "hello") + v11) + (v11 + v11))}
    v13 := GenS2{F0: (v11 + v11)}
    if ((!((v5 < v5) || (((v5 % (((^(v5))) | 1)) == (^(v5))) || (!(!(-373 < v4)))))) || true) {
        switch 42 {
        case 32768:
            switch v4 {
            case 0:
                var v14 float64 = (-2.25 - float64(-2.25))
                _ = v14
            case -1:
                if (2 >= 42) {
                    var v15 int = v4
                    v16 := GenS3{F0: 4294967295, F1: false}
                    _ = v15
                    _ = v16
                } else {
                    fmt.Println(v5, v4)
                    v3 = 42
                }
            }
        case 32767:
            switch "a" {
            case "foo bar":
                v17 := GenS4{F0: "hello", F1: (v3 ^ 4294967295), F2: 42}
                _ = v17
            }
        case -943:
            v18 := GenS5{F0: -1, F1: 42}
            _ = v18
        }
    }
    if false {
        if true {
            var v19 int8 = -128
            for fi20 := 0; fi20 < 1; fi20++ {
                fmt.Println(v19, v3, v4)
                var v21 string = v11
                var v22 float64 = 0.1
                _ = fi20
                _ = v21
                _ = v22
            }
            v11 = (((("a" + "a") + "foo bar") + v11) + v11)
            _ = v19
        } else {
            var v23 uint8 = 1
            _ = v23
        }
    }
    v3 -= v3
    var v24 uint32 = 1
    v1 = append(v1, v4)
    for fi25 := 0; fi25 < 13; fi25++ {
        if len(v1) > 0 {
            v1[0] = (v4 << ((0) & 63))
        }
        var v26 uint16 = 754
        _ = fi25
        _ = v26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
