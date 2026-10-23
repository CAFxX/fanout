package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint64; F1 uint32 }

func main() {
    var v1 uint8 = 2
    var v2 string = (("" + ("foo bar" + "hello")) + ("a" + "a"))
    v3 := make(map[int]string)
    v4 := make(map[int]int)
    v5 := make(map[string]bool)
    if true {
        if (42 >= 222) {
            var v6 int64 = 891
            var v7 string = v2
            v1 -= 255
            _ = v6
            _ = v7
        } else {
            v8 := GenS1{F0: (true || (false || ((v1 & (v1 - 128)) <= (v1 << ((4294967295) & 63)))))}
            _ = v8
        }
        v9 := []int{}
        _ = v9
    }
    var v10 string = (v2 + v2)
    switch (v10 + ("foo bar" + "foo bar")) {
    case "\n\t":
        v11 := GenS2{F0: -2147483648, F1: -395}
        _ = v11
    }
    v12 := []int{-450, 1, -1, 42}
    var v13 bool = true
    var v14 int16 = 42
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 string = (("foo bar" + v2) + (v2 + "\n\t"))
        _ = len(v12)
        v13 = ((4294967295 != 317) || v13)
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        for fi18 := 0; fi18 < 11; fi18++ {
            v19 := GenS1{F0: v13}
            switch (^(fi18)) {
            case 128:
                fmt.Println(fi17, fi17, fi17)
            case 32768:
                v20 := GenS1{F0: ((!true) && (2147483647 > 0))}
                _ = v20
            case -2147483648:
                v21 := GenS2{F0: 0, F1: 42}
                _ = v21
            default:
                v22 := GenS3{F0: -1, F1: (128 >= v1)}
                _ = v22
            }
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    var v23 string = v2
    for fi24 := 0; fi24 < 17; fi24++ {
        v25 := GenS4{F0: 42, F1: 2147483648}
        var v26 int64 = -1
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 uint = 1
    delete(v3, 1)
    var v28 float32 = ((0.0 * 0.0) + (float32((((1.5 + -2.25) * 3.14159) / 1.5)) * (((-2.25 - 1.5) * 3.14159) + (1e10 - (0.1 * (((0.0 + 1e10) * float32(((((0.0 / float64(float32((3.14159 + -2.25)))) * (3.14159 + 1e10)) - float64((1.5 - 1e10))) * -2.25))) + (((-2.25 * (3.14159 - (float32(0.0) - 1e10))) * 1.5) + (1.5 / 1.0))))))))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
