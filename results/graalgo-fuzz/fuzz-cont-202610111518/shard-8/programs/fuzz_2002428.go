package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8 }
func genFn0() uint32 {
    v4 := GenS1{F0: -9223372036854775808}
    var v5 float64 = ((1.5 / 1.0) - float64(0.1))
    _ = v4
    _ = v5
    return 435
}
func genFn1(p0 uint32) int16 {
    if (1 < 2147483647) {
        var v6 uint64 = 42
        fmt.Println(p0, v6)
        _ = v6
    } else {
        fmt.Println(p0)
        var v7 uint64 = 1
        v8 := []int{-1, 128, 32768, 708}
        _ = v7
        _ = v8
    }
    if (-1 > -1) {
        var v9 string = (("" + (("hello" + "hello") + "a")) + ("\n\t" + "\n\t"))
        var v10 bool = (255 == 153)
        _ = v10
        _ = v9
    }
    if ((true || ((2 < 129) && ("" > "hello"))) || ((-2.25 <= 0.1) || false)) {
        if ((!(!true)) || (0 <= -32768)) {
            fmt.Println(p0)
            p0 += 505
        }
        for fi11 := 0; fi11 < 14; fi11++ {
            switch "hello" {
            case "a":
                var v12 uint32 = p0
                _ = v12
            default:
                switch fi11 {
                case 122:
                    v13 := GenS1{F0: int64(p0)}
                    _ = v13
                default:
                    v14 := GenS2{F0: 202}
                    _ = v14
                }
            }
            v15 := GenS1{F0: 42}
            var v16 float64 = (-2.25 - 0.0)
            _ = fi11
            _ = v15
            _ = v16
        }
        v17 := []int{-1, 2147483647, -449}
        _ = v17
    }
    var v18 float64 = 1.5
    _ = p0
    _ = v18
    return -32768
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[int]bool)
    v3 := make(map[uint32]int)
    for fi19 := 0; fi19 < 11; fi19++ {
        v1 = v1
        for fi20 := 0; fi20 < 1; fi20++ {
            v21 := []int{-605}
            v1 = v1
            v1 -= -81
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    if (0.0 == 1.5) {
        switch 630 {
        case 9223372036854775807:
            fmt.Println(v1)
        case -86:
            for fi22 := 0; fi22 < 8; fi22++ {
                v23 := []int{}
                _ = fi22
                _ = v23
            }
        default:
            var v24 uint32 = 1
            _ = v24
        }
    }
    switch "" {
    case "a":
        if (!(false || true)) {
            var v25 float32 = (float32((float64(0.1) - (0.1 - (float64(float32(3.14159)) * (3.14159 * float64(((3.14159 - 0.1) - 0.1))))))) / 1e10)
            _ = v25
        }
    case "\n\t":
        v1 += v1
    case "hello":
        for fi26 := 0; fi26 < 20; fi26++ {
            v1 = genFn1(4294967295)
            v27 := GenS2{F0: 344}
            _ = fi26
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 10; fi28++ {
        var v29 uint64 = 18446744073709551615
        _ = fi28
        _ = v29
    }
    var v30 uint8 = 127
    fmt.Println(v30)
    for fi31 := 0; fi31 < 9; fi31++ {
        v32 := GenS3{F0: 127}
        v1 -= (v1 * v1)
        _ = len(v2)
        _ = fi31
        _ = v32
    }
    _ = len(v3)
    v30 = 128
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
}
