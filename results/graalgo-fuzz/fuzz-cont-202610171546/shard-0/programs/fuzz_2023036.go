package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 uint32; F2 int8 }
type GenS4 struct { F0 int8 }

func main() {
    var v1 bool = (!((((float32(float64((0.0 * float32(float64(float32(1e10)))))) >= float32((0.1 + 1e10))) && true) && true) || false))
    var v2 uint64 = 2
    v3 := make(map[int]string)
    var v4 uint32 = 0
    var v5 int32 = 42
    var v6 uint8 = 2
    v6 = v6
    v7 := GenS1{F0: 1, F1: -121, F2: -256}
    v8 := GenS2{F0: v4}
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 uint = 255
        if true {
            delete(v3, -9223372036854775808)
            v11 := GenS3{F0: ("foo bar" + "foo bar"), F1: (v4 & (^(v4))), F2: -104}
            v1 = (!v1)
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    if ((v4 & v4) < uint32(v6)) {
        fmt.Println(v1, v4, v5)
        var v12 float64 = float64(3.14159)
        _ = v12
    } else {
        v13 := GenS2{F0: v4}
        _ = v13
    }
    var v14 int8 = 1
    v4 -= (v4 | (^(v4)))
    v3[-587] = ("foo bar" + "hello")
    for fi15 := 0; fi15 < 12; fi15++ {
        fmt.Println(v5, v1, v14)
        var v16 float64 = -2.25
        _ = fi15
        _ = v16
    }
    v5 = v5
    for fi17 := 0; fi17 < 11; fi17++ {
        v18 := []int{fi17, (-(fi17)), 0, (^(fi17))}
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        v3[fi19] = ("a" + ("" + "a"))
        _ = fi19
    }
    v20 := GenS4{F0: v14}
    if v1 {
        var v21 uint32 = v4
        v2 -= 2
        _ = v21
    } else {
        switch "" {
        case "hello":
            var v22 int64 = int64(v2)
            _ = v22
        case "":
            v23 := GenS3{F0: "a", F1: 1, F2: -1}
            _ = v23
        }
        v24 := []int{32767}
        v6 += v6
        _ = v24
    }
    var v25 uint32 = 841
    v26 := GenS1{F0: 32767, F1: (v14 >> ((255) & 63)), F2: 278}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
