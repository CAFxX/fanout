package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int8 }
func genFn0(p0 int16) int8 {
    v4 := GenS1{F0: 0, F1: "foo bar"}
    v5 := []int{-9223372036854775808, -967}
    _ = p0
    _ = v4
    _ = v5
    return 127
}

func main() {
    var v1 string = ((("" + ("\n\t" + "a")) + "hello") + ("\n\t" + "hello"))
    var v2 uint32 = 1
    var v3 float64 = float64(float32(float64(0.1)))
    fmt.Println(v2)
    for fi6 := 0; fi6 < 5; fi6++ {
        switch (v1 + ("" + v1)) {
        case "\n\t":
            v7 := GenS1{F0: 0, F1: (v1 + v1)}
            _ = v7
        default:
            v3 = float64(1e10)
        }
        _ = fi6
    }
    switch ("\n\t" + v1) {
    case "hello":
        v8 := []int{32767, 173}
        _ = v8
    case "\n\t":
        v1 = v1
    case "":
        var v9 int64 = 1
        _ = v9
    }
    v3 = v3
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 float64 = v3
        for fi12 := 0; fi12 < 7; fi12++ {
            v13 := GenS2{F0: 18}
            switch ("\n\t" + v1) {
            case "foo bar":
                if ((((fi12 | -1) <= 4294967295) && ((42 > -191) && false)) || false) {
                    var v14 int64 = 1
                    _ = v14
                }
            case "\n\t":
                v15 := []int{(fi12 + fi10), 128}
                _ = v15
            }
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    v2 = (v2 >> ((2) & 63))
    fmt.Println(v2)
    v16 := []int{9223372036854775807, 42}
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
}
