package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 uint8; F2 int8 }
func genFn0(p0 string, p1 uint) int {
    var v5 uint64 = uint64(p1)
    switch (p0 + p0) {
    case "\n\t":
        v5 -= v5
    case "hello":
        var v6 float64 = 3.14159
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v5
    return 128
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[string]bool)
    var v3 uint16 = 65535
    var v4 int32 = -2147483648
    var v7 uint8 = 245
    if true {
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 int16 = -32768
            _ = fi8
            _ = v9
        }
        var v10 uint = uint(v7)
        if false {
            v4 = v4
        }
        _ = v10
    }
    _ = len(v2)
    v11 := GenS1{F0: -32768, F1: v7}
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := GenS2{F0: v7}
        v14 := []int{127, genFn0("a", 18446744073709551615), (fi12 + fi12)}
        delete(v2, "a")
        _ = fi12
        _ = v13
        _ = v14
    }
    v3 -= v3
    for fi15 := 0; fi15 < 14; fi15++ {
        v1 += (^(v1))
        if true {
            if (((v4 / ((int32(v7)) | 1)) != v4) && true) {
                var v16 int8 = 127
                _ = v16
            } else {
                v17 := []int{fi15, -948, (fi15 % ((fi15) | 1))}
                if false {
                    for fi18 := 0; fi18 < 9; fi18++ {
                        var v19 uint64 = 1
                        v20 := GenS3{F0: fi15, F1: (v7 % ((v7) | 1)), F2: 127}
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    v2[""] = ((fi15 != (fi15 << ((2) & 63))) || false)
                }
                _ = v17
            }
        }
        _ = fi15
    }
    v21 := v2["\n\t"]
    if (v3 == v3) {
        v4 = (v4 | v4)
        v4 = (^(v4))
    }
    fmt.Println(v1, v21)
    var v22 int = 724
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v7
}
