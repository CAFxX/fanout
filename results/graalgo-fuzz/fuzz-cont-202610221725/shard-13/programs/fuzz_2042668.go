package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint8 }
func genFn0() uint16 {
    var v6 uint16 = 847
    v7 := []int{}
    v6 = (v6 & v6)
    var v8 uint8 = 1
    _ = v6
    _ = v7
    _ = v8
    return 904
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]string)
    var v3 int16 = -1
    v4 := make(map[int]bool)
    v5 := []string{"a", "\n\t"}
    v3 += v3
    var v9 uint16 = 0
    v10 := GenS1{F0: false, F1: v3}
    fmt.Println(v3)
    if (true && false) {
        fmt.Println(v3)
        var v11 float64 = (0.1 / 1.5)
        v11 += 0.1
        _ = v11
    } else {
        v12 := GenS2{F0: 2147483647, F1: 374, F2: 42}
        v13 := GenS3{F0: 42, F1: 659, F2: 2}
        v9 -= 65535
        _ = v12
        _ = v13
    }
    var v14 string = ("foo bar" + "hello")
    if (!false) {
        for fi15 := 0; fi15 < 4; fi15++ {
            delete(v1, "")
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 string = (v14 + "foo bar")
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
        fmt.Println(v9)
    }
    v18 := GenS2{F0: -1, F1: 18446744073709551615, F2: 1}
    fmt.Println(v14, v3)
    v9 += (v9 + 32767)
    fmt.Println(v9, v3)
    v14 = ("foo bar" + ("hello" + v14))
    delete(v2, v14)
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 uint32 = 2
        _ = fi19
        _ = v20
    }
    fmt.Println(v3, v14)
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 int = (fi21 - fi21)
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        for fi24 := 0; fi24 < 16; fi24++ {
            v9 = (^(v9))
            var v25 uint64 = 1
            fmt.Println(v9, v3, v9)
            _ = fi24
            _ = v25
        }
        delete(v4, (fi23 % ((-1) | 1)))
        var v26 bool = false
        _ = fi23
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
