package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 int8 }

func main() {
    v1 := []uint32{4294967295, 2147483647}
    v2 := make(map[string]uint32)
    var v3 uint16 = 32767
    var v4 string = (("hello" + ((("" + "") + "a") + "")) + "hello")
    v4 = v4
    v3 = (v3 | (v3 >> ((42) & 63)))
    var v5 float32 = 3.14159
    if (-32768 == -1) {
        var v6 string = v4
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := GenS1{F0: 234}
            if (1 == 42) {
                v9 := GenS2{F0: (^(v3)), F1: 41}
                var v10 uint = 0
                _ = v10
                _ = v9
            }
            v11 := GenS3{F0: 457, F1: 0}
            _ = fi7
            _ = v11
            _ = v8
        }
        _ = v6
    } else {
        for fi12 := 0; fi12 < 4; fi12++ {
            var v13 uint64 = 750
            v13 += (^(v13))
            v2[(((v4 + v4) + "hello") + ("hello" + "foo bar"))] = 42
            _ = fi12
            _ = v13
        }
        var v14 int8 = -1
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        v3 = v3
        var v16 int8 = 22
        var v17 int8 = (v16 | int8(fi15))
        _ = fi15
        _ = v16
        _ = v17
    }
    v4 = (v4 + "a")
    v3 = (^(v3))
    v18 := GenS4{F0: 127}
    v3 = (v3 + (v3 / (((^(v3))) | 1)))
    v19 := GenS3{F0: 32767, F1: 1}
    for fi20 := 0; fi20 < 1; fi20++ {
        _ = len(v1)
        fmt.Println(fi20, v3)
        _ = fi20
    }
    v4 = (v4 + v4)
    v5 += ((v5 / float32(float64(float32(3.14159)))) - ((v5 + (((v5 * -2.25) - 0.1) + v5)) - v5))
    if (!(32768 >= -864)) {
        v4 = ("\n\t" + "a")
    }
    fmt.Println(v3)
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
