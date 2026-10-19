package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint) uint64 {
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 uint32 = 2
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint32 = 65535
        switch "\n\t" {
        case "hello":
            var v9 string = "hello"
            _ = v9
        }
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := GenS1{F0: 676}
            var v12 string = ""
            switch fi10 {
            case 2147483647:
                var v13 uint16 = 413
                _ = v13
            case -185:
                v14 := GenS2{F0: true}
                _ = v14
            case -9223372036854775808:
                var v15 int64 = -195
                _ = v15
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    var v16 uint = (p0 | 8)
    p0 += 4294967295
    p0 = (p0 % ((4294967295) | 1))
    _ = p0
    _ = v16
    return 364
}

func main() {
    var v1 int = -662
    var v2 bool = (true && (false || true))
    v3 := make(map[int]string)
    var v4 int64 = 189
    v1 = v1
    for fi17 := 0; fi17 < 18; fi17++ {
        switch ("" + "foo bar") {
        case "a":
            var v18 uint32 = 65535
            _ = v18
        default:
            var v19 float64 = (-2.25 - 1.5)
            _ = v19
        }
        for fi20 := 0; fi20 < 3; fi20++ {
            var v21 int16 = -844
            switch ("foo bar" + (("foo bar" + "hello") + "")) {
            case "":
                v22 := []int{}
                _ = v22
            case "foo bar":
                var v23 int64 = 1
                _ = v23
            default:
                var v24 int16 = (v21 + v21)
                var v25 int = fi17
                v26 := GenS2{F0: (((!(1.5 >= -2.25)) || (false && (!((v24 ^ v21) <= v24)))) && ((genFn0(1) == 869) && (128 == 2)))}
                _ = v24
                _ = v25
                _ = v26
            }
            _ = fi20
            _ = v21
        }
        _ = fi17
    }
    var v27 uint = 0
    fmt.Println(v1)
    fmt.Println(v27)
    v2 = v2
    var v28 int64 = v4
    var v29 float32 = (0.1 * (((-2.25 - (3.14159 / ((1e10 * (0.1 / 0.1)) * 1.5))) + 1e10) + 0.1))
    var v30 bool = v2
    v31 := GenS3{F0: genFn0((v27 - (v27 | 255)))}
    v29 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
