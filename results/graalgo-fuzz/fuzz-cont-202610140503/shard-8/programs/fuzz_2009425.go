package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 uint32; F1 bool; F2 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 2
    v3 := make(map[string]int)
    v4 := make(map[int]string)
    v5 := []int{-2147483648}
    var v6 int = 435
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v5) > 0 {
        v5[0] = v6
    }
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := GenS1{F0: 0, F1: 42}
        _ = fi7
        _ = v8
    }
    v4[v6] = "foo bar"
    for fi9 := 0; fi9 < 14; fi9++ {
        fmt.Println(fi9, v6, v2)
        _ = fi9
    }
    var v10 int8 = 100
    var v11 int8 = 0
    for fi12 := 0; fi12 < 3; fi12++ {
        v11 += v11
        v6 = 416
        if (42 >= 42) {
            var v13 int = 632
            if (255 > 255) {
                delete(v3, (("" + (("foo bar" + "a") + "foo bar")) + ("\n\t" + (("\n\t" + ("\n\t" + "\n\t")) + "hello"))))
                var v14 uint64 = 2
                _ = v14
            } else {
                if true {
                    var v15 uint64 = 0
                    _ = v15
                }
            }
            _ = len(v1)
            _ = v13
        } else {
            v16 := GenS1{F0: 1, F1: 2}
            v6 = (fi12 & 1)
            v10 -= (v10 + v10)
            _ = v16
        }
        _ = fi12
    }
    var v17 uint = 1
    v10 = (v11 * (v11 + (^(v10))))
    v18 := GenS1{F0: 127, F1: 18446744073709551615}
    for fi19 := 0; fi19 < 4; fi19++ {
        fmt.Println(v2, v17)
        _ = fi19
    }
    v20 := GenS1{F0: v11, F1: 2}
    if ((0.0 == (0.1 + (-2.25 / (0.0 + float64(0.1))))) && (v2 <= uint8(v17))) {
        v21 := GenS2{F0: v6, F1: (v11 << ((v17) & 63))}
        var v22 int64 = 547
        fmt.Println(v22)
        _ = v21
        _ = v22
    } else {
        v23 := GenS3{F0: 42, F1: false, F2: (v2 + 127)}
        v2 = (v2 & v2)
        _ = v23
    }
    v24 := v3["\n\t"]
    var v25 uint32 = 2147483648
    v17 -= 2
    v6 += (v24 % (((v24 % ((v6) | 1))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
