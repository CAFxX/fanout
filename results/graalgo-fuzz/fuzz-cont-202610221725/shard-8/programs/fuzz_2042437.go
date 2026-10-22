package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]int)
    var v3 float64 = 1.5
    v4 := make(map[uint32]uint32)
    var v5 float32 = -2.25
    var v6 float64 = float64(0.1)
    switch (("\n\t" + ("foo bar" + ("foo bar" + ""))) + "hello") {
    case "":
        v7 := v2[770]
        _ = v7
    case "\n\t":
        v6 += v3
    }
    fmt.Println(v3, v3)
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := GenS1{F0: -128, F1: ""}
        v10 := GenS2{F0: -1, F1: 0}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint32 = 4294967295
    v11 = (v11 & v11)
    v12 := GenS3{F0: 32767}
    switch 900 {
    case -913:
        var v13 string = "\n\t"
        _ = v13
    case 42:
        if (false || ((float32((v3 + v6)) - float32((v6 / (v3 + v3)))) == (v5 + v5))) {
            for fi14 := 0; fi14 < 17; fi14++ {
                v6 += ((0.1 / float64(v5)) / -2.25)
                _ = fi14
            }
            var v15 float64 = (v3 - v3)
            var v16 uint8 = 255
            _ = v15
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        var v18 float64 = ((float64(v5) + (float64(float32(v6)) - (v3 * float64(float32(0.0))))) * ((v3 / (float64(0.0) + (float64((v5 + v5)) - 1e10))) - v6))
        v19 := v4[v11]
        _ = fi17
        _ = v18
        _ = v19
    }
    v6 = (0.1 * (v3 / (v6 * float64(1.5))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
