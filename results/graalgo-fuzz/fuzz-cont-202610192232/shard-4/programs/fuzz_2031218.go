package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int; F1 int8; F2 uint8 }
type GenS3 struct { F0 int32; F1 uint64 }

func main() {
    var v1 int = -2147483648
    var v2 uint64 = 18446744073709551615
    var v3 uint32 = 42
    var v4 float64 = 1.5
    var v5 int = (v1 << ((0) & 63))
    if ((!false) || true) {
        if true {
            var v6 int64 = -9223372036854775808
            if (v3 != v3) {
                v7 := GenS1{F0: 146, F1: "a"}
                v8 := []int{}
                var v9 int64 = (v6 % (((v6 | v6)) | 1))
                _ = v7
                _ = v8
                _ = v9
            }
            if (22 >= 533) {
                v10 := GenS2{F0: (v5 + v5), F1: 127, F2: 2}
                var v11 float32 = ((-2.25 * 1e10) * (1.5 / 1.0))
                if true {
                    var v12 uint32 = 2
                    switch ("foo bar" + ("hello" + "\n\t")) {
                    case "a":
                        v13 := []int{}
                        _ = v13
                    case "":
                        v5 -= 42
                    }
                    v14 := GenS1{F0: 628, F1: ((("hello" + ("foo bar" + "")) + ("hello" + "a")) + ("" + "a"))}
                    _ = v12
                    _ = v14
                } else {
                    if ((v3 * 2147483648) >= v3) {
                        fmt.Println(v1, v6, v6)
                    }
                    fmt.Println(v6, v1, v11)
                }
                _ = v10
                _ = v11
            } else {
                for fi15 := 0; fi15 < 2; fi15++ {
                    var v16 int = int(v6)
                    v17 := []int{fi15, int(v6), -375, v5}
                    if ((v5 / ((int(v2)) | 1)) <= v1) {
                        v17 = append(v17, 128)
                        switch (v5 << ((389) & 63)) {
                        case 128:
                            if true {
                                v3 += v3
                            }
                        case -378:
                            var v18 float32 = (3.14159 * 3.14159)
                            _ = v18
                        default:
                            v19 := GenS2{F0: (v5 << ((18446744073709551615) & 63)), F1: 127, F2: 92}
                            _ = v19
                        }
                        v20 := GenS3{F0: -1, F1: v2}
                        _ = v20
                    } else {
                        v21 := GenS1{F0: 32767, F1: "hello"}
                        for fi22 := 0; fi22 < 5; fi22++ {
                            v4 = v4
                            v3 += 175
                            _ = fi22
                        }
                        _ = v21
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                v5 -= v5
                v23 := []int{v5}
                _ = v23
            }
            _ = v6
        }
    } else {
        var v24 int = (v5 - -46)
        _ = v24
    }
    var v25 string = (("a" + "\n\t") + "hello")
    v5 = 2147483647
    v26 := GenS3{F0: 107, F1: (v2 % ((253) | 1))}
    fmt.Println(v4, v2)
    for fi27 := 0; fi27 < 3; fi27++ {
        var v28 float64 = v4
        var v29 string = ((v25 + "hello") + v25)
        _ = fi27
        _ = v28
        _ = v29
    }
    fmt.Println(v4, v5, v1)
    for fi30 := 0; fi30 < 6; fi30++ {
        v1 += (v1 | 32767)
        v31 := []int{(v1 / ((1) | 1)), v5, -222}
        if true {
            v3 += (v3 * (v3 % ((v3) | 1)))
            var v32 uint32 = (^(v3))
            var v33 float64 = ((3.14159 + v4) + v4)
            _ = v32
            _ = v33
        }
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
