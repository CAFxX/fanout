package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 uint; F2 int }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 uint16; F1 uint16; F2 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 845
    v3 := make(map[uint32]uint32)
    for fi4 := 0; fi4 < 16; fi4++ {
        v5 := GenS1{F0: fi4}
        _ = fi4
        _ = v5
    }
    switch 829 {
    case 32767:
        for fi6 := 0; fi6 < 7; fi6++ {
            v2 = v2
            _ = fi6
        }
    case 9223372036854775807:
        var v7 float32 = float32(1e10)
        _ = v7
    case -2147483648:
        v2 -= 9223372036854775807
    default:
        if (!(false && true)) {
            for fi8 := 0; fi8 < 12; fi8++ {
                v9 := v1[(fi8 + (fi8 % ((fi8) | 1)))]
                v10 := GenS1{F0: 335}
                _ = fi8
                _ = v10
                _ = v9
            }
            v2 += 9223372036854775807
            for fi11 := 0; fi11 < 4; fi11++ {
                fmt.Println(v2)
                var v12 string = "\n\t"
                _ = fi11
                _ = v12
            }
        } else {
            v13 := GenS1{F0: -9223372036854775808}
            fmt.Println(v2)
            _ = v13
        }
    }
    var v14 int = 0
    for fi15 := 0; fi15 < 4; fi15++ {
        v2 += v2
        _ = fi15
    }
    v16 := GenS2{F0: (-(v2)), F1: 1, F2: v14}
    fmt.Println(v14)
    var v17 uint64 = 0
    v17 -= (v17 / ((4294967295) | 1))
    v18 := GenS3{F0: "hello", F1: (v2 / ((-1) | 1))}
    v17 += 42
    fmt.Println(v17, v17)
    if (false && true) {
        var v19 float32 = 3.14159
        _ = v19
    } else {
        v20 := GenS1{F0: v14}
        for fi21 := 0; fi21 < 17; fi21++ {
            switch (("hello" + "foo bar") + "foo bar") {
            case "hello":
                v22 := GenS4{F0: 65535, F1: 1, F2: 42}
                _ = v22
            case "":
                fmt.Println(fi21, fi21)
            case "\n\t":
                for fi23 := 0; fi23 < 12; fi23++ {
                    var v24 bool = (true && (false || (-273 >= 1)))
                    v24 = v24
                    var v25 uint8 = 2
                    switch (v14 << ((264) & 63)) {
                    case 42:
                        if (("a" + "") >= "foo bar") {
                            v2 = -1
                            v26 := GenS2{F0: 1, F1: 321, F2: fi23}
                            var v27 int8 = -11
                            _ = v26
                            _ = v27
                        }
                    }
                    v1[1] = "hello"
                    if false {
                        var v28 int16 = 32767
                        var v29 int8 = 42
                        v25 -= uint8(v17)
                        _ = v28
                        _ = v29
                    } else {
                        var v30 float64 = 1.5
                        var v31 uint8 = v25
                        _ = len(v3)
                        _ = v30
                        _ = v31
                    }
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
            v32 := []int{}
            _ = fi21
            _ = v32
        }
        v14 += v14
        _ = v20
    }
    v33 := GenS1{F0: v14}
    var v34 uint8 = 137
    v3[434] = 502
    v17 = 419
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v34
}
