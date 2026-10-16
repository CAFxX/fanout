package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 uint64 }

func main() {
    v1 := make(map[int]string)
    var v2 uint = 250
    var v3 int = -1
    v4 := make(map[int]bool)
    var v5 int32 = -2147483648
    var v6 int16 = 1000
    for fi7 := 0; fi7 < 12; fi7++ {
        v2 -= v2
        v8 := GenS1{F0: false, F1: v6}
        var v9 int32 = v5
        _ = fi7
        _ = v8
        _ = v9
    }
    v6 = v6
    var v10 int32 = 1
    var v11 int32 = 1
    if (100 >= int8(v11)) {
        v12 := GenS1{F0: ("" == "a"), F1: -809}
        v2 = (^(v2))
        if true {
            v13 := GenS2{F0: v11, F1: 38, F2: 64}
            for fi14 := 0; fi14 < 13; fi14++ {
                v6 = v6
                v15 := GenS2{F0: (v5 - -1), F1: 0, F2: 58}
                _ = fi14
                _ = v15
            }
            _ = v13
        } else {
            var v16 uint8 = 255
            v17 := v1[202]
            if (1e10 >= -2.25) {
                if (42 == -9223372036854775808) {
                    switch 128 {
                    case 4294967295:
                        fmt.Println(v5, v17)
                    case 127:
                        v2 -= 2
                    case 1:
                        v18 := GenS3{F0: 5}
                        _ = v18
                    default:
                        fmt.Println(v2, v16)
                    }
                    v19 := GenS4{F0: (v2 | (v2 & v2)), F1: 0}
                    _ = v19
                }
                v20 := v1[277]
                _ = v20
            } else {
                var v21 bool = true
                if false {
                    for fi22 := 0; fi22 < 3; fi22++ {
                        var v23 int = 42
                        var v24 uint64 = 1
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                }
                _ = v21
            }
            _ = v16
            _ = v17
        }
        _ = v12
    }
    switch (v3 % ((int(v6)) | 1)) {
    case 32767:
        v25 := GenS4{F0: (v2 >> ((0) & 63)), F1: 42}
        _ = v25
    default:
        var v26 int = 127
        _ = v26
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        var v28 uint8 = 42
        for fi29 := 0; fi29 < 5; fi29++ {
            var v30 int64 = 9223372036854775807
            _ = fi29
            _ = v30
        }
        _ = fi27
        _ = v28
    }
    var v31 string = ("hello" + "a")
    v32 := GenS2{F0: -753, F1: 255, F2: 2}
    for fi33 := 0; fi33 < 10; fi33++ {
        v5 += v10
        _ = fi33
    }
    _ = len(v1)
    fmt.Println(v31, v31, v10)
    v11 = v11
    var v34 float64 = (((float64(0.1) / 1.0) * 1.5) * 1.5)
    v35 := []int{v3, (v3 - v3), (v3 % ((-238) | 1))}
    v2 -= (v2 * v2)
    var v36 uint = (v2 ^ (^(v2)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
