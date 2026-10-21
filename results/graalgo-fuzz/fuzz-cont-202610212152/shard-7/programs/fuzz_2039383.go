package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 uint16; F2 int8 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 int16 }

func main() {
    var v1 float64 = -2.25
    v2 := make(map[uint32]uint32)
    var v3 int = 353
    var v4 float64 = v1
    v5 := GenS1{F0: 0}
    v6 := GenS1{F0: 42}
    var v7 int16 = 1
    if (!true) {
        v8 := GenS2{F0: 2, F1: 852, F2: -64}
        _ = v8
    }
    v4 += v1
    for fi9 := 0; fi9 < 1; fi9++ {
        v7 = -507
        var v10 int64 = 0
        switch "hello" {
        case "foo bar":
            fmt.Println(v7)
        case "hello":
            v11 := GenS1{F0: 42}
            _ = v11
        case "a":
            var v12 uint64 = 4294967295
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    var v13 int32 = -2147483648
    for fi14 := 0; fi14 < 12; fi14++ {
        v7 = 943
        _ = fi14
    }
    v15 := GenS3{F0: "a", F1: (v7 / ((v7) | 1))}
    v2[1] = 4294967295
    v3 -= 932
    fmt.Println(v7, v1, v1)
    fmt.Println(v13)
    if ((!true) || (66 > 127)) {
        fmt.Println(v1, v3)
        if (!true) {
            v3 += 32768
            v16 := GenS4{F0: v7}
            for fi17 := 0; fi17 < 16; fi17++ {
                v4 = v4
                var v18 uint64 = 994
                v19 := []int{(v3 / (((fi17 * -494)) | 1))}
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = v16
        } else {
            fmt.Println(v3, v4, v4)
        }
        var v20 float32 = 1e10
        _ = v20
    }
    var v21 string = ("\n\t" + ("hello" + "a"))
    v13 -= v13
    v13 -= 905
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
