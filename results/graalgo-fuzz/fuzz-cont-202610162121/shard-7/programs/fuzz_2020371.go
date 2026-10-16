package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 uint64 }
func genFn0(p0 int16) int {
    fmt.Println(p0)
    p0 += (p0 + p0)
    switch "\n\t" {
    case "a":
        for fi4 := 0; fi4 < 10; fi4++ {
            v5 := []int{fi4, -480}
            var v6 uint16 = 547
            p0 = p0
            _ = fi4
            _ = v5
            _ = v6
        }
    case "\n\t":
        p0 = p0
    case "foo bar":
        switch 128 {
        case -340:
            v7 := []int{-490, 127}
            _ = v7
        case 4294967295:
            var v8 float32 = 1e10
            _ = v8
        default:
            p0 -= -1
        }
    default:
        switch -2147483648 {
        case 9223372036854775807:
            var v9 float32 = 1e10
            _ = v9
        case 32767:
            p0 = (p0 * p0)
        case -481:
            p0 = -32768
        default:
            v10 := GenS1{F0: 42, F1: 776, F2: p0}
            _ = v10
        }
    }
    _ = p0
    return 2147483647
}
func genFn1(p0 uint8, p1 uint16, p2 uint) bool {
    v11 := GenS2{F0: 127, F1: ("hello" + ("foo bar" + ("\n\t" + "")))}
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return false
}

func main() {
    v1 := make(map[string]string)
    var v2 int = 0
    var v3 int64 = 0
    v2 += (^(v2))
    v12 := GenS2{F0: -1, F1: (("\n\t" + ("a" + "a")) + ("a" + "a"))}
    v13 := GenS2{F0: 16, F1: "foo bar"}
    var v14 int16 = 42
    var v15 float64 = 0.0
    v16 := GenS3{F0: 0}
    v17 := GenS4{F0: 1, F1: 775}
    v18 := GenS3{F0: (v14 ^ (v14 >> ((2) & 63)))}
    var v19 int8 = -1
    var v20 uint64 = 103
    for fi21 := 0; fi21 < 6; fi21++ {
        if false {
            if (true || (!(true || (("\n\t" <= ("" + (("\n\t" + (("hello" + ((("foo bar" + "a") + "a") + ("foo bar" + "\n\t"))) + "")) + ("foo bar" + ("hello" + "\n\t"))))) || (v19 != v19))))) {
                fmt.Println(fi21, v20)
            }
        } else {
            v19 = (v19 * v19)
        }
        _ = fi21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
