package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 string; F2 int }
func genFn0(p0 string) uint32 {
    var v7 uint32 = 2147483647
    for fi8 := 0; fi8 < 17; fi8++ {
        if (756 <= -417) {
            var v9 float32 = (((float32(1.5) - (((float32(0.1) - 0.1) / -2.25) * 1e10)) * 0.1) / 0.1)
            v10 := GenS1{F0: p0, F1: 101}
            v11 := []int{fi8, fi8}
            _ = v10
            _ = v11
            _ = v9
        }
        var v12 string = ("hello" + (p0 + (("hello" + (p0 + ((((("hello" + ("foo bar" + p0)) + p0) + "a") + p0) + p0))) + (p0 + p0))))
        _ = fi8
        _ = v12
    }
    _ = p0
    _ = v7
    return (^(v7))
}

func main() {
    var v1 int8 = 1
    var v2 uint8 = 127
    var v3 float64 = 0.1
    var v4 uint = 0
    var v5 int = int(v2)
    var v6 int = v5
    switch ("\n\t" + ("foo bar" + "foo bar")) {
    case "a":
        fmt.Println(v2)
    default:
        for fi13 := 0; fi13 < 20; fi13++ {
            var v14 uint16 = 515
            _ = fi13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        if (!((42 != 452) || (((42 == -9223372036854775808) || true) || true))) {
            v16 := []int{(^(v6)), (v6 * v5)}
            _ = v16
        } else {
            fmt.Println(v5)
            v3 -= (1e10 - float64((3.14159 / ((0.0 + (1e10 * 1e10)) + 0.0))))
        }
        var v17 int8 = (-(v1))
        v1 = (v1 / ((v17) | 1))
        _ = fi15
        _ = v17
    }
    if true {
        v18 := []int{(v5 << ((v4) & 63)), v5}
        v19 := GenS1{F0: "", F1: 682}
        if (!((genFn0("hello") == uint32(v2)) && (false && (191 >= 1)))) {
            v5 -= v5
            switch -1 {
            case 42:
                v20 := GenS2{F0: 42}
                _ = v20
            }
            v21 := GenS2{F0: -1}
            _ = v21
        } else {
            for fi22 := 0; fi22 < 14; fi22++ {
                for fi23 := 0; fi23 < 4; fi23++ {
                    v6 = fi22
                    var v24 int32 = -211
                    _ = len(v18)
                    _ = fi23
                    _ = v24
                }
                _ = fi22
            }
            var v25 uint16 = 117
            v1 += v1
            _ = v25
        }
        _ = v18
        _ = v19
    }
    v4 = v4
    v26 := GenS1{F0: "hello", F1: -400}
    fmt.Println(v2)
    v4 += 825
    v27 := []int{}
    var v28 int = (v6 - v5)
    switch (v6 / (((v6 ^ (-(v6)))) | 1)) {
    case 4294967295:
        v29 := GenS3{F0: v1, F1: "hello", F2: 9223372036854775807}
        _ = v29
    case -564:
        v30 := GenS1{F0: ("\n\t" + "hello"), F1: 1000}
        _ = v30
    case 496:
        var v31 uint16 = 32767
        _ = v31
    default:
        switch "a" {
        case "a":
            v32 := GenS3{F0: v1, F1: "hello", F2: -835}
            _ = v32
        case "":
            var v33 uint32 = 2
            _ = v33
        default:
            _ = len(v27)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
