package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 int64; F2 int32 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int, p1 int8) string {
    fmt.Println(p1, p0)
    v7 := []int{(p0 >> ((2) & 63)), p0, 4294967295, 32768}
    p1 += 0
    _ = p0
    _ = p1
    _ = v7
    return ("hello" + "foo bar")
}
func genFn1() int64 {
    v8 := []int{32768, 32768, -937}
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 uint32 = 42
        _ = fi9
        _ = v10
    }
    var v11 int64 = 1
    switch -903 {
    case 32767:
        for fi12 := 0; fi12 < 5; fi12++ {
            var v13 uint32 = 2147483647
            _ = fi12
            _ = v13
        }
    }
    _ = v11
    _ = v8
    return v11
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = 42
    var v3 float32 = 0.1
    var v4 uint8 = 155
    var v5 int8 = -117
    var v6 uint16 = 255
    var v14 float64 = 1.5
    var v15 int32 = v2
    var v16 int = 127
    v3 += (v3 - v3)
    v15 -= v15
    v17 := GenS1{F0: 4294967295, F1: 994}
    if ("a" >= ("\n\t" + "\n\t")) {
        var v18 string = (("foo bar" + "hello") + "")
        v19 := GenS1{F0: 0, F1: 635}
        v20 := []int{-9223372036854775808, v16, v16}
        _ = v18
        _ = v19
        _ = v20
    } else {
        fmt.Println(v16, v6, v3)
    }
    v5 -= v5
    v6 = v6
    fmt.Println(v2)
    if false {
        v21 := GenS2{F0: 255}
        v22 := GenS2{F0: (^(v6))}
        for fi23 := 0; fi23 < 15; fi23++ {
            switch fi23 {
            case -922:
                if true {
                    v24 := GenS1{F0: 42, F1: 884}
                    _ = v24
                }
            default:
                var v25 string = ""
                _ = v25
            }
            _ = fi23
        }
        _ = v21
        _ = v22
    } else {
        var v26 string = "a"
        _ = v26
    }
    v27 := GenS3{F0: 2, F1: 797, F2: (v2 / (((v2 % ((2147483647) | 1))) | 1))}
    v14 = v14
    if (362 <= genFn1()) {
        v28 := GenS1{F0: 0, F1: 2}
        _ = v28
    }
    switch "hello" {
    case "\n\t":
        v29 := GenS4{F0: 18446744073709551615}
        _ = v29
    default:
        switch ("hello" + (("foo bar" + "a") + "\n\t")) {
        case "hello":
            switch (v16 * 128) {
            case 32767:
                v16 = (v16 ^ v16)
            case 0:
                for fi30 := 0; fi30 < 11; fi30++ {
                    fmt.Println(v16)
                    v31 := []int{fi30}
                    fmt.Println(v6, v4)
                    _ = fi30
                    _ = v31
                }
            case -9223372036854775808:
                v3 = (float32(3.14159) + v3)
            default:
                v32 := GenS3{F0: 2, F1: 42, F2: v15}
                _ = v32
            }
        case "a":
            _ = len(v1)
        case "\n\t":
            v33 := GenS2{F0: (v6 * (v6 >> ((703) & 63)))}
            _ = v33
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
