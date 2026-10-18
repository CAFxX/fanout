package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 uint; F2 int16 }
type GenS3 struct { F0 int; F1 uint }

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = -1
    var v3 int = 797
    var v4 int32 = v2
    v5 := make(map[int]bool)
    var v6 int = (v3 | (v3 >> ((511) & 63)))
    if (-1 >= 42) {
        if true {
            var v7 float64 = 0.0
            if (true || (!false)) {
                fmt.Println(v3)
            }
            if (-2.25 == (float64(((1e10 / -2.25) + float32(float64(-2.25)))) - 3.14159)) {
                var v8 string = "a"
                var v9 uint = 2
                _ = v8
                _ = v9
            } else {
                for fi10 := 0; fi10 < 7; fi10++ {
                    var v11 uint8 = 178
                    v2 = v2
                    v12 := []int{v3, v6, (v3 * v6), v3}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                var v13 bool = true
                for fi14 := 0; fi14 < 16; fi14++ {
                    v15 := GenS1{F0: -9223372036854775808}
                    _ = fi14
                    _ = v15
                }
                _ = v13
            }
            _ = v7
        }
        v16 := GenS2{F0: -519, F1: 255, F2: 32767}
        _ = v16
    }
    var v17 int = (v6 & -982)
    for fi18 := 0; fi18 < 16; fi18++ {
        v19 := GenS1{F0: 0}
        var v20 float32 = 1.5
        _ = fi18
        _ = v19
        _ = v20
    }
    v4 = int32(v6)
    var v21 uint8 = 2
    v4 = (v2 * (v4 << ((1) & 63)))
    fmt.Println(v2)
    for fi22 := 0; fi22 < 1; fi22++ {
        if false {
            var v23 string = (("a" + "hello") + "foo bar")
            _ = v23
        } else {
            var v24 uint16 = 32767
            v3 -= (v17 & (-(v3)))
            v6 -= v3
            _ = v24
        }
        var v25 int64 = 948
        _ = fi22
        _ = v25
    }
    switch "hello" {
    case "foo bar":
        if false {
            switch v3 {
            case -2147483648:
                v26 := GenS2{F0: -1, F1: 2, F2: 42}
                _ = v26
            default:
                for fi27 := 0; fi27 < 16; fi27++ {
                    switch (^(fi27)) {
                    case 32768:
                        v28 := GenS3{F0: (v3 / (((v6 & v3)) | 1)), F1: 2}
                        _ = v28
                    }
                    fmt.Println(v21, v6)
                    _ = fi27
                }
            }
        }
    case "a":
        if (!("a" >= "a")) {
            for fi29 := 0; fi29 < 17; fi29++ {
                switch "hello" {
                case "hello":
                    v5[fi29] = true
                case "\n\t":
                    v17 += v6
                case "a":
                    v30 := GenS1{F0: 0}
                    _ = v30
                }
                v31 := GenS2{F0: -9223372036854775808, F1: 2, F2: 32767}
                _ = fi29
                _ = v31
            }
        } else {
            v32 := []int{(v17 ^ (v3 / ((711) | 1))), (v6 | v6), -745}
            _ = len(v1)
            _ = v32
        }
    }
    var v33 uint16 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
