package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 uint64; F2 int8 }
func genFn0(p0 int16) int32 {
    var v7 string = "a"
    p0 = -245
    _ = p0
    _ = v7
    return int32(p0)
}
func genFn1(p0 string) int8 {
    if false {
        p0 = p0
        if (42 != -474) {
            if (((false && (true && true)) || true) || (-875 != 520)) {
                if true {
                    for fi8 := 0; fi8 < 20; fi8++ {
                        fmt.Println(fi8, fi8)
                        _ = fi8
                    }
                }
                for fi9 := 0; fi9 < 19; fi9++ {
                    v10 := []int{(-(fi9)), 4294967295}
                    p0 = p0
                    p0 = p0
                    _ = fi9
                    _ = v10
                }
            } else {
                var v11 int16 = 1000
                var v12 uint32 = 1
                v11 = 1
                _ = v11
                _ = v12
            }
            p0 = "a"
        }
        var v13 int32 = 0
        _ = v13
    }
    p0 = "a"
    _ = p0
    return 8
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]string)
    var v3 float64 = (1e10 + 1e10)
    var v4 int32 = -2147483648
    var v5 uint32 = 650
    var v6 bool = false
    var v14 uint = 0
    v3 = (v3 * 0.1)
    for fi15 := 0; fi15 < 3; fi15++ {
        v16 := GenS1{F0: -2147483648}
        _ = fi15
        _ = v16
    }
    if v6 {
        v17 := GenS1{F0: 484}
        v3 = (v3 / (3.14159 * (v3 * (v3 + ((1e10 * v3) + v3)))))
        for fi18 := 0; fi18 < 11; fi18++ {
            if (v6 || (255 < 0)) {
                v19 := GenS2{F0: (fi18 + fi18)}
                _ = v19
            }
            v20 := GenS2{F0: fi18}
            v21 := GenS1{F0: int(v14)}
            _ = fi18
            _ = v20
            _ = v21
        }
        _ = v17
    } else {
        var v22 bool = false
        var v23 uint64 = 1
        v4 = 42
        _ = v22
        _ = v23
    }
    fmt.Println(v14)
    var v24 uint32 = (v5 << ((uint(v5)) & 63))
    for fi25 := 0; fi25 < 9; fi25++ {
        v26 := GenS1{F0: (fi25 >> (((v14 | 4294967295)) & 63))}
        v5 = 1
        _ = fi25
        _ = v26
    }
    v24 = 2
    v27 := GenS3{F0: 2, F1: 4294967295, F2: genFn1((("foo bar" + "\n\t") + ("foo bar" + (("\n\t" + ("\n\t" + "\n\t")) + ("hello" + "")))))}
    v5 = v24
    var v28 uint64 = 42
    if v6 {
        var v29 uint8 = 62
        var v30 float32 = float32((v3 - v3))
        _ = v29
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
