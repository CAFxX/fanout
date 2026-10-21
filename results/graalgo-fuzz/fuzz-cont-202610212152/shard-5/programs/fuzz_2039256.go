package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint32, p1 uint64, p2 uint16) int8 {
    p0 += p0
    var v5 float64 = 1.5
    switch "" {
    case "a":
        var v6 uint = 255
        _ = v6
    case "hello":
        v7 := GenS1{F0: 42}
        _ = v7
    case "\n\t":
        var v8 float64 = (3.14159 - (((v5 / (float64(0.1) + v5)) / 0.1) - float64(3.14159)))
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 100
}
func genFn1() int8 {
    v9 := []int{217, 128}
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 int32 = 0
        for fi12 := 0; fi12 < 5; fi12++ {
            v11 = v11
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    var v13 uint = 18446744073709551615
    _ = v13
    _ = v9
    return 42
}

func main() {
    var v1 float64 = 1e10
    var v2 bool = (true && (4294967295 >= 42))
    v3 := make(map[int]int)
    var v4 int16 = 0
    var v14 uint8 = 0
    var v15 float64 = ((1e10 - (v1 * -2.25)) * -2.25)
    var v16 float64 = 0.0
    v3[911] = -70
    _ = len(v3)
    switch 849 {
    case -1:
        _ = len(v3)
    case -9223372036854775808:
        v4 += (v4 - (v4 ^ -1))
    case 1:
        for fi17 := 0; fi17 < 4; fi17++ {
            fmt.Println(v1, v4, v14)
            var v18 uint16 = 255
            for fi19 := 0; fi19 < 4; fi19++ {
                var v20 int64 = -1
                _ = fi19
                _ = v20
            }
            _ = fi17
            _ = v18
        }
    }
    var v21 string = ""
    v4 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
