package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    var v3 float64 = 1e10
    v4 := make(map[uint32]int)
    var v5 string = "a"
    for fi6 := 0; fi6 < 20; fi6++ {
        fmt.Println(fi6, v3)
        _ = fi6
    }
    if ("a" > "") {
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := GenS1{F0: 0}
            _ = fi7
            _ = v8
        }
        v9 := []int{-13, 4294967295, 1, -9223372036854775808}
        v3 = 1e10
        _ = v9
    } else {
        for fi10 := 0; fi10 < 20; fi10++ {
            for fi11 := 0; fi11 < 6; fi11++ {
                var v12 int16 = 0
                if (0.0 != (v3 / v3)) {
                    if true {
                        v5 = (v5 + v5)
                    }
                    v5 = (v5 + v5)
                } else {
                    fmt.Println(fi11, v12)
                    if (!true) {
                        var v13 int32 = -682
                        var v14 int16 = 1000
                        var v15 uint16 = 1
                        _ = v13
                        _ = v14
                        _ = v15
                    } else {
                        delete(v1, 0)
                        v16 := []int{fi10, fi11, -987, fi10}
                        v16 = append(v16, fi10)
                        _ = v16
                    }
                    v17 := v2[2147483648]
                    switch (("a" + v5) + v5) {
                    case "hello":
                        v18 := GenS2{F0: 302, F1: 2, F2: 0}
                        var v19 int16 = (v12 >> ((65535) & 63))
                        fmt.Println(fi10)
                        var v20 int8 = 1
                        if true {
                            if (true && true) {
                                _ = len(v4)
                            }
                            v21 := []int{int(v19), (fi10 / (((fi10 + fi10)) | 1)), v17, v17}
                            fmt.Println(v3)
                            v22 := GenS1{F0: 187}
                            v19 -= 1000
                            _ = v21
                            _ = v22
                        }
                        _ = v18
                        _ = v19
                        _ = v20
                    case "a":
                        v23 := []int{(-(fi10)), (-(fi10)), fi11}
                        if false {
                            var v24 uint16 = 2
                            v5 = ("foo bar" + (v5 + ((("a" + v5) + v5) + ("" + v5))))
                            _ = v24
                        } else {
                            _ = len(v23)
                            v25 := GenS1{F0: 1}
                            switch (v17 / (((fi10 / (((fi10 - (v17 * fi10))) | 1))) | 1)) {
                            case -2147483648:
                                var v26 uint64 = 1
                                _ = v26
                            }
                            _ = v25
                        }
                        _ = v23
                    }
                    _ = v17
                }
                _ = fi11
                _ = v12
            }
            var v27 uint8 = 0
            _ = fi10
            _ = v27
        }
        v5 = "a"
    }
    var v28 int32 = 1
    var v29 uint8 = 78
    v3 = v3
    var v30 uint16 = 748
    v31 := GenS2{F0: 18446744073709551615, F1: 1, F2: 18446744073709551615}
    var v32 bool = ((-1 < 0) || (true && false))
    if (uint32(v29) > 42) {
        v29 += (v29 ^ v29)
        fmt.Println(v32, v3)
    } else {
        v29 = 143
        v30 = (^(v30))
        delete(v4, 4294967295)
    }
    var v33 int64 = -9223372036854775808
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
