package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }

func main() {
    var v1 uint64 = 625
    v2 := make(map[uint32]string)
    var v3 uint64 = (v1 / (((v1 | v1)) | 1))
    var v4 float64 = -2.25
    var v5 int16 = -684
    var v6 float64 = v4
    for fi7 := 0; fi7 < 20; fi7++ {
        v4 += (v4 / v4)
        _ = fi7
    }
    v6 -= (v4 + v4)
    fmt.Println(v6, v4)
    v3 -= 4294967295
    var v8 uint32 = 0
    if false {
        for fi9 := 0; fi9 < 4; fi9++ {
            v5 += 42
            _ = fi9
        }
        var v10 int32 = -137
        _ = v10
    }
    var v11 float32 = 0.1
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 int = (fi12 / (((^(fi12))) | 1))
        _ = fi12
        _ = v13
    }
    if (!((v5 == (v5 - -551)) || ((v5 * int16(v8)) < (v5 * v5)))) {
        for fi14 := 0; fi14 < 17; fi14++ {
            switch fi14 {
            case 4294967295:
                for fi15 := 0; fi15 < 18; fi15++ {
                    var v16 string = "foo bar"
                    fmt.Println(v3, fi14, v5)
                    v17 := GenS1{F0: 4294967295, F1: 276}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            case 9223372036854775807:
                for fi18 := 0; fi18 < 5; fi18++ {
                    delete(v2, v8)
                    if (!false) {
                        v19 := GenS1{F0: 2, F1: 0}
                        if (!("\n\t" > "")) {
                            var v20 int32 = 207
                            _ = v20
                        }
                        _ = v19
                    }
                    v21 := []int{fi18, fi14}
                    _ = fi18
                    _ = v21
                }
            case 293:
                var v22 uint64 = v1
                _ = v22
            }
            _ = fi14
        }
        var v23 bool = false
        _ = v23
    }
    v4 = (v4 / (v6 + (v6 - v4)))
    fmt.Println(v5, v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
