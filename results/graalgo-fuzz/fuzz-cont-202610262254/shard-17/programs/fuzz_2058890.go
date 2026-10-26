package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() string {
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := GenS1{F0: 32767}
        _ = fi7
        _ = v8
    }
    var v9 float64 = 1.5
    _ = v9
    return ("foo bar" + "hello")
}
func genFn1() uint64 {
    var v10 uint = 828
    v11 := []int{}
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 uint8 = 255
        _ = fi12
        _ = v13
    }
    _ = v10
    _ = v11
    return 580
}

func main() {
    var v1 int32 = 2147483647
    var v2 bool = false
    v3 := make(map[int]string)
    var v4 uint16 = 2
    var v5 int = 4294967295
    v6 := make(map[uint32]string)
    v1 -= (v1 / (((v1 * 42)) | 1))
    fmt.Println(v5, v2, v4)
    v14 := v3[int(v1)]
    for fi15 := 0; fi15 < 5; fi15++ {
        v5 = fi15
        var v16 int16 = 1000
        _ = fi15
        _ = v16
    }
    var v17 uint64 = 2
    var v18 uint64 = v17
    if (18446744073709551615 == 65535) {
        var v19 int = v5
        _ = v19
    }
    v4 += (v4 % (((^(v4))) | 1))
    v3[(^(v5))] = "a"
    for fi20 := 0; fi20 < 18; fi20++ {
        fmt.Println(v1)
        for fi21 := 0; fi21 < 4; fi21++ {
            fmt.Println(fi20, v18)
            _ = fi21
        }
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
