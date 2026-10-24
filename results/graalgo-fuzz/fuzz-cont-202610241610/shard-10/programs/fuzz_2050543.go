package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 int8 }

func main() {
    var v1 uint64 = 0
    v2 := make(map[uint32]string)
    v3 := []int{127, -61, 2147483647, 41}
    v4 := []string{"", "foo bar"}
    v5 := []uint32{65535}
    var v6 int16 = 807
    v7 := GenS1{F0: 32768, F1: 0, F2: 209}
    v1 = v1
    var v8 int8 = 1
    if ((1.5 * ((3.14159 * 0.1) / 3.14159)) > 3.14159) {
        switch "\n\t" {
        case "hello":
            v1 -= (v1 % (((v1 * v1)) | 1))
        }
        var v9 float64 = 1e10
        if (1 < 32767) {
            var v10 int = -2147483648
            delete(v2, 85)
            _ = v10
        } else {
            var v11 float64 = (1.5 / ((float64((float32(0.0) * (float32(v9) / 1.0))) * 3.14159) * float64(1.5)))
            var v12 bool = true
            v11 = (3.14159 + 1e10)
            _ = v11
            _ = v12
        }
        _ = v9
    }
    if false {
        v13 := GenS1{F0: 0, F1: 129, F2: 648}
        v8 += (v8 & 71)
        _ = v13
    }
    if false {
        _ = len(v2)
        v6 = 32767
        switch ("\n\t" + ("a" + ((("\n\t" + "hello") + "a") + ""))) {
        case "\n\t":
            var v14 string = "\n\t"
            _ = v14
        default:
            v6 = (v6 & (-(v6)))
        }
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        fmt.Println(v8, v1)
        _ = fi15
    }
    v16 := GenS2{F0: 2}
    for fi17 := 0; fi17 < 14; fi17++ {
        switch int(v8) {
        case 701:
            var v18 uint64 = v1
            _ = v18
        }
        for fi19 := 0; fi19 < 5; fi19++ {
            v20 := v2[692]
            _ = fi19
            _ = v20
        }
        _ = fi17
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        v6 = (-(v6))
        _ = fi21
    }
    var v22 string = "hello"
    var v23 uint16 = 65535
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        v25 := GenS3{F0: 19, F1: (v8 ^ (v8 >> ((919) & 63)))}
        _ = fi24
        _ = v25
    }
    v26 := GenS1{F0: 191, F1: 28, F2: 18446744073709551615}
    for fi27 := 0; fi27 < 1; fi27++ {
        v28 := GenS3{F0: 42, F1: v8}
        if (true && true) {
            fmt.Println(fi27)
        } else {
            v22 = (v22 + v22)
            v1 += (v1 * v1)
        }
        var v29 int32 = 829
        _ = fi27
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
