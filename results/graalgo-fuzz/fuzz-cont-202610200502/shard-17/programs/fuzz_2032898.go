package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint32; F1 int32; F2 int16 }
type GenS3 struct { F0 int8; F1 int32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{127, -197}
    var v3 uint64 = 934
    var v4 string = "hello"
    var v5 uint16 = 2
    var v6 uint64 = v3
    v4 = v4
    for fi7 := 0; fi7 < 13; fi7++ {
        if len(v2) > 0 {
            v2[0] = -2147483648
        }
        v8 := GenS1{F0: -989, F1: 100}
        _ = fi7
        _ = v8
    }
    var v9 int8 = 100
    var v10 int = 127
    v6 -= v3
    switch (v4 + v4) {
    case "\n\t":
        var v11 int64 = -1
        _ = v11
    case "":
        v9 -= (^(v9))
    }
    if ((v5 / ((v5) | 1)) > 65535) {
        for fi12 := 0; fi12 < 15; fi12++ {
            if (v3 < (v3 * (v3 - (v6 << ((42) & 63))))) {
                var v13 float32 = float32((-2.25 + 0.0))
                v5 = (v5 + 42)
                v14 := GenS2{F0: 2, F1: 2147483647, F2: 42}
                _ = v13
                _ = v14
            }
            if len(v2) > 0 {
                v2[0] = (-(fi12))
            }
            if (-2147483648 >= 42) {
                v4 = (((v4 + (v4 + "\n\t")) + v4) + v4)
                for fi15 := 0; fi15 < 13; fi15++ {
                    var v16 int = (fi12 | (v10 | v10))
                    var v17 int16 = 1000
                    _ = len(v1)
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            _ = fi12
        }
        var v18 string = "a"
        _ = v18
    } else {
        var v19 uint = 42
        _ = v19
    }
    if true {
        if false {
            if (false && true) {
                v4 = (v4 + (v4 + (v4 + v4)))
                var v20 bool = (1000 < -1)
                var v21 uint32 = 656
                _ = v20
                _ = v21
            }
        }
        if (4294967295 > 18446744073709551615) {
            if true {
                for fi22 := 0; fi22 < 11; fi22++ {
                    fmt.Println(v5, v6)
                    var v23 float32 = float32(0.1)
                    if true {
                        v10 -= (fi22 << ((uint(v6)) & 63))
                        v4 = v4
                        v10 += v10
                    }
                    _ = fi22
                    _ = v23
                }
                v1[258] = false
            } else {
                var v24 uint = 970
                v9 += (^(v9))
                _ = v24
            }
            v5 = 32767
        } else {
            var v25 float64 = (1e10 - (1.5 * 0.1))
            v1[2147483648] = (true && true)
            fmt.Println(v6, v9, v4)
            _ = v25
        }
        switch -785 {
        case 9223372036854775807:
            fmt.Println(v3, v9)
        default:
            var v26 uint32 = 1
            _ = v26
        }
    } else {
        v27 := GenS2{F0: 2147483648, F1: 1, F2: 1}
        _ = v27
    }
    v3 = (^(v6))
    if true {
        switch 1 {
        case -210:
            var v28 uint16 = v5
            _ = v28
        default:
            v9 = (v9 % ((-68) | 1))
        }
        if ((!(!true)) && true) {
            var v29 int8 = v9
            v30 := GenS3{F0: int8(v5), F1: -1}
            var v31 uint32 = 776
            _ = v29
            _ = v30
            _ = v31
        } else {
            var v32 uint = 685
            var v33 int32 = 758
            fmt.Println(v33)
            _ = v32
            _ = v33
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
