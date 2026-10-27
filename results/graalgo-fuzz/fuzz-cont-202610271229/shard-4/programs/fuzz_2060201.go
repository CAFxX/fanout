package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint16; F1 string; F2 int }
type GenS4 struct { F0 int32; F1 int32; F2 int16 }
func genFn0() uint8 {
    v4 := []int{-1, 804, -2147483648, 42}
    for fi5 := 0; fi5 < 7; fi5++ {
        fmt.Println(fi5)
        switch "a" {
        case "hello":
            fmt.Println(fi5)
        case "":
            v6 := GenS1{F0: 1, F1: 255, F2: (fi5 - (fi5 + 32768))}
            _ = v6
        case "foo bar":
            v7 := GenS2{F0: 18446744073709551615, F1: 0}
            _ = v7
        }
        _ = fi5
    }
    _ = v4
    return 0
}

func main() {
    v1 := []uint32{0, 616, 65535, 4294967295}
    v2 := make(map[int]uint32)
    v3 := make(map[string]uint32)
    var v8 uint32 = 902
    var v9 uint32 = v8
    v8 = 0
    fmt.Println(v9, v8)
    v10 := GenS3{F0: 84, F1: "foo bar", F2: 32767}
    _ = len(v1)
    var v11 int8 = 1
    var v12 int8 = v11
    var v13 float64 = ((1e10 * 1.5) + -2.25)
    v1 = append(v1, 2147483648)
    if (!(18446744073709551615 >= 4294967295)) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if (-81 < 1) {
            var v14 int = 94
            switch ("a" + "") {
            case "a":
                if (1 != 1) {
                    fmt.Println(v11)
                    var v15 int32 = 420
                    _ = v15
                } else {
                    v8 = 861
                    var v16 string = "hello"
                    _ = v16
                }
            case "\n\t":
                for fi17 := 0; fi17 < 5; fi17++ {
                    v12 -= v12
                    _ = len(v1)
                    _ = fi17
                }
            }
            _ = len(v3)
            _ = v14
        }
        var v18 float32 = (float32(float64(0.0)) + 1.5)
        _ = v18
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        delete(v3, ("a" + "a"))
        _ = fi19
    }
    v20 := GenS4{F0: 0, F1: 0, F2: -32768}
    v9 = v9
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v8
    _ = v9
}
