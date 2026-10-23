package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint64; F2 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 int32; F2 int8 }
func genFn0(p0 uint64, p1 uint, p2 uint64) uint64 {
    v7 := []int{1}
    v8 := GenS1{F0: 262}
    var v9 bool = (42 <= -1)
    fmt.Println(p1, p1, v9)
    v10 := GenS2{F0: (p0 + (p2 & p0)), F1: 2, F2: ("a" + "hello")}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return (p0 + 1)
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 4294967295
    var v3 float32 = float32(1e10)
    v4 := make(map[uint32]int)
    var v5 float64 = 1e10
    var v6 int64 = 9223372036854775807
    var v11 int8 = 42
    if false {
        if ((false || false) || true) {
            v6 += v6
        }
        v2 = (v2 + (v2 | 28))
    } else {
        for fi12 := 0; fi12 < 1; fi12++ {
            switch ((((("" + "a") + "a") + "\n\t") + "") + "") {
            case "a":
                for fi13 := 0; fi13 < 13; fi13++ {
                    var v14 uint8 = uint8(v2)
                    var v15 uint = 65535
                    var v16 uint64 = 42
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
            }
            v17 := GenS3{F0: ""}
            switch fi12 {
            case 127:
                v6 -= -9223372036854775808
            case 42:
                if (!("" <= ((("foo bar" + (("\n\t" + "hello") + "a")) + "") + ("\n\t" + (("\n\t" + ((("foo bar" + "") + "hello") + "a")) + ""))))) {
                    var v18 float32 = (float32((3.14159 - ((v5 / 1.5) - v5))) + ((v3 + ((v3 - 0.1) / (1.5 * ((v3 / 1.5) - float32(((v5 * 1.5) / 1.0)))))) / v3))
                    var v19 uint16 = 255
                    v20 := []int{fi12, -458, fi12}
                    _ = v18
                    _ = v19
                    _ = v20
                }
            default:
                v21 := GenS4{F0: 566, F1: -1, F2: v11}
                _ = v21
            }
            _ = fi12
            _ = v17
        }
        for fi22 := 0; fi22 < 6; fi22++ {
            v23 := []int{fi22, (fi22 * fi22), 4294967295}
            if false {
                switch -175 {
                case 127:
                    var v24 uint64 = genFn0(42, 0, 18446744073709551615)
                    _ = v24
                default:
                    var v25 bool = (((!false) || (true || (18446744073709551615 > 255))) && true)
                    _ = v25
                }
            } else {
                _ = len(v23)
            }
            _ = fi22
            _ = v23
        }
        v26 := GenS3{F0: (("" + "a") + "\n\t")}
        _ = v26
    }
    var v27 int = 42
    v28 := []int{v27, (v27 & 128)}
    for fi29 := 0; fi29 < 18; fi29++ {
        var v30 uint = 42
        _ = fi29
        _ = v30
    }
    v2 += (^(v2))
    v6 = 1
    _ = len(v28)
    v6 -= (v6 ^ v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
