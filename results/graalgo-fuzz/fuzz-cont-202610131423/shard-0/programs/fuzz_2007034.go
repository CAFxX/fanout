package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8 }
func genFn0() uint32 {
    var v4 uint32 = 2147483647
    v5 := []int{-9223372036854775808}
    var v6 int16 = 1000
    _ = v4
    _ = v5
    _ = v6
    return v4
}

func main() {
    var v1 int32 = 42
    var v2 int8 = 100
    v3 := make(map[uint32]bool)
    var v7 int64 = -714
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 uint16 = 32767
        v10 := GenS1{F0: ("\n\t" + "a"), F1: -100}
        var v11 int64 = (v7 * (^(v7)))
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        v13 := GenS1{F0: "hello", F1: (v2 / (((^(v2))) | 1))}
        if (2147483647 >= 851) {
            var v14 float32 = float32((-2.25 / float64((1.5 + float32((0.0 + (((1.5 + 1.5) + -2.25) - 0.1)))))))
            _ = v14
        } else {
            v15 := GenS1{F0: "hello", F1: (v2 & -1)}
            v3[2] = true
            fmt.Println(v7, v2)
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    v1 -= v1
    v7 += v7
    v16 := GenS2{F0: -334}
    v17 := GenS2{F0: 0}
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 int = (fi18 + (fi18 >> ((4294967295) & 63)))
        v2 -= v2
        var v20 float32 = ((1.5 + 0.1) / float32(0.1))
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 float32 = (((-2.25 * float32(1.5)) / ((1.5 + 3.14159) - ((((((0.0 * (1e10 - float32(float64(0.0)))) + 1.5) / ((0.1 - 0.0) * 3.14159)) - 0.1) / 0.1) + (float32(0.1) - 0.0)))) + 1e10)
    v21 = 0.1
    var v22 float32 = 0.1
    v22 = v22
    v21 = v21
    for fi23 := 0; fi23 < 10; fi23++ {
        if (true && (0.0 <= v21)) {
            switch (fi23 & fi23) {
            case -92:
                _ = len(v3)
            }
            for fi24 := 0; fi24 < 2; fi24++ {
                if true {
                    if (794 == 2) {
                        var v25 int32 = v1
                        _ = v25
                    } else {
                        var v26 int = (fi24 | int(v1))
                        v22 += v21
                        _ = len(v3)
                        _ = v26
                    }
                    delete(v3, 42)
                }
                v27 := GenS3{F0: (-(v2))}
                _ = fi24
                _ = v27
            }
            for fi28 := 0; fi28 < 9; fi28++ {
                v29 := []int{fi28, (-(fi23)), fi28}
                var v30 float64 = -2.25
                v3[0] = false
                _ = fi28
                _ = v29
                _ = v30
            }
        } else {
            fmt.Println(v22, v2, v2)
            v2 += 0
        }
        for fi31 := 0; fi31 < 19; fi31++ {
            var v32 uint64 = 264
            var v33 bool = true
            var v34 string = (("foo bar" + "hello") + "foo bar")
            _ = fi31
            _ = v32
            _ = v33
            _ = v34
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v7
}
