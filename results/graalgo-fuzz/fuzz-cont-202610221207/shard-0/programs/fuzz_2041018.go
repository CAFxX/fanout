package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 uint, p1 int16, p2 int64) bool {
    var v5 float64 = float64(-2.25)
    p0 += (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (false && false)
}
func genFn1() int16 {
    var v6 int64 = 0
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 float32 = 3.14159
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 int16 = 1
            switch fi9 {
            case 560:
                if true {
                    var v11 uint16 = 777
                    _ = v11
                } else {
                    var v12 string = "foo bar"
                    _ = v12
                }
            case 128:
                var v13 float32 = 3.14159
                _ = v13
            default:
                var v14 uint64 = 18446744073709551615
                _ = v14
            }
            _ = fi9
            _ = v10
        }
        var v15 bool = (("\n\t" + "foo bar") <= "foo bar")
        _ = fi7
        _ = v15
        _ = v8
    }
    _ = v6
    return -32768
}

func main() {
    var v1 uint = 0
    v2 := []int{32768}
    var v3 int32 = -414
    v4 := make(map[int]string)
    v1 += v1
    var v16 uint = v1
    v1 -= v1
    switch 1 {
    case -697:
        _ = len(v2)
    }
    v17 := GenS1{F0: (v3 + int32(v1))}
    if ((!true) && (!(968 >= v16))) {
        fmt.Println(v1)
        var v18 string = "\n\t"
        _ = v18
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        fmt.Println(v3, v16, v3)
        var v20 float32 = 0.0
        _ = fi19
        _ = v20
    }
    v16 += 0
    v16 = v1
    for fi21 := 0; fi21 < 6; fi21++ {
        v3 += (v3 * (v3 ^ v3))
        v16 += 1
        for fi22 := 0; fi22 < 8; fi22++ {
            var v23 uint16 = 1
            v23 -= (^(v23))
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    var v24 uint16 = 255
    fmt.Println(v1, v24, v3)
    v3 -= v3
    v16 = (^(v1))
    var v25 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
