package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 string; F2 uint32 }
type GenS4 struct { F0 int; F1 uint; F2 int64 }
type GenS5 struct { F0 int32 }
func genFn0(p0 int32) int16 {
    var v7 float32 = 0.0
    v8 := GenS1{F0: "", F1: 894, F2: -1}
    v9 := []int{-1, 127, 1}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    var v4 int8 = 72
    var v5 int16 = 42
    var v6 float32 = 0.1
    if (true && (!true)) {
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := GenS1{F0: "", F1: 388, F2: 42}
            for fi12 := 0; fi12 < 15; fi12++ {
                switch ("\n\t" + "hello") {
                case "hello":
                    v13 := GenS1{F0: "a", F1: 848, F2: v4}
                    _ = v13
                default:
                    var v14 uint64 = 2
                    _ = v14
                }
                v6 -= v6
                v6 -= (v6 - v6)
                _ = fi12
            }
            for fi15 := 0; fi15 < 8; fi15++ {
                v1 = (v1 << ((4294967295) & 63))
                v16 := GenS2{F0: ("\n\t" + (("a" + ((((("" + "\n\t") + ("hello" + "hello")) + ("foo bar" + "a")) + "") + "a")) + ("" + "a")))}
                _ = fi15
                _ = v16
            }
            _ = fi10
            _ = v11
        }
        if true {
            var v17 int16 = genFn0(627)
            if false {
                _ = len(v2)
            }
            _ = v17
        }
    }
    v18 := GenS3{F0: false, F1: "\n\t", F2: 660}
    var v19 bool = false
    var v20 float64 = -2.25
    var v21 int8 = (v4 << ((65535) & 63))
    v22 := GenS4{F0: -317, F1: 4294967295, F2: -1}
    v23 := GenS5{F0: 2147483647}
    var v24 uint64 = 262
    v19 = (!false)
    delete(v2, 1)
    if v19 {
        for fi25 := 0; fi25 < 14; fi25++ {
            if v19 {
                switch fi25 {
                case 128:
                    v1 -= 1
                }
            }
            _ = fi25
        }
        v5 += v5
        v6 = v6
    }
    v21 -= (v21 ^ (v4 * v21))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
