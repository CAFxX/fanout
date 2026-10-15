package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int; F1 int64; F2 int }
func genFn0() uint64 {
    v7 := GenS1{F0: 4294967295, F1: 95}
    var v8 float32 = 0.0
    fmt.Println(v8)
    var v9 uint16 = 0
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint32 = 972
    v2 := make(map[uint32]string)
    var v3 uint = 255
    v4 := make(map[int]bool)
    var v5 uint32 = v1
    var v6 int32 = 0
    v10 := GenS2{F0: 276, F1: 18446744073709551615}
    for fi11 := 0; fi11 < 7; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            v13 := v4[fi11]
            fmt.Println(v6)
            _ = fi12
            _ = v13
        }
        v5 += (v1 * 4294967295)
        var v14 int = 127
        _ = fi11
        _ = v14
    }
    v3 = v3
    var v15 int8 = 1
    fmt.Println(v6)
    v15 -= (v15 + (-(v15)))
    var v16 string = "a"
    v3 = (v3 | v3)
    var v17 int32 = v6
    v18 := v2[v5]
    var v19 uint16 = 901
    switch v18 {
    case "hello":
        fmt.Println(v17)
    default:
        if false {
            if (18446744073709551615 != 682) {
                fmt.Println(v18, v16, v16)
                var v20 int = 127
                var v21 uint64 = 4294967295
                _ = v20
                _ = v21
            }
            v5 = 42
            for fi22 := 0; fi22 < 20; fi22++ {
                for fi23 := 0; fi23 < 2; fi23++ {
                    v16 = ((v16 + (v18 + (v18 + (((v18 + "a") + v18) + (("a" + "a") + v18))))) + "\n\t")
                    if true {
                        v3 -= 2
                    } else {
                        v17 = (-(v6))
                        if (false && ((!false) || (v6 != v6))) {
                            var v24 int32 = int32(v1)
                            v25 := []int{1, 42, (fi22 << ((728) & 63)), 1}
                            if false {
                                v18 = "foo bar"
                            }
                            var v26 int16 = 0
                            v27 := GenS3{F0: fi23, F1: 9223372036854775807, F2: (fi22 ^ 42)}
                            _ = v24
                            _ = v25
                            _ = v26
                            _ = v27
                        } else {
                            v28 := []int{fi22, (-(fi22)), (^(fi22)), fi23}
                            if (!(0.0 <= float64((((3.14159 - ((float32(1.5) - 3.14159) + -2.25)) + -2.25) - (0.0 + 1e10))))) {
                                var v29 bool = true
                                v2[(v5 ^ v1)] = (v16 + v16)
                                _ = v29
                            }
                            _ = v28
                        }
                    }
                    var v30 string = ("\n\t" + v18)
                    _ = fi23
                    _ = v30
                }
                _ = fi22
            }
        }
    }
    if true {
        v31 := []int{32768, 32767}
        v2[v1] = ""
        _ = v31
    }
    fmt.Println(v1)
    var v32 uint32 = 42
    var v33 int64 = 9223372036854775807
    for fi34 := 0; fi34 < 15; fi34++ {
        var v35 uint = 558
        switch (^(fi34)) {
        case 2147483647:
            var v36 int = (fi34 / ((0) | 1))
            _ = v36
        }
        _ = fi34
        _ = v35
    }
    fmt.Println(v16)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
