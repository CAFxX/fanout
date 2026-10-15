package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 string }
type GenS2 struct { F0 uint; F1 uint32; F2 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 int64; F2 int16 }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 uint32; F1 uint64 }

func main() {
    var v1 int8 = 42
    var v2 bool = (2147483647 < 0)
    var v3 float64 = ((float64(0.1) - (0.1 / 1e10)) * 0.1)
    switch "hello" {
    case "hello":
        if v2 {
            v2 = (!v2)
            var v4 int = 42
            fmt.Println(v2)
            _ = v4
        } else {
            v5 := GenS1{F0: 866, F1: "", F2: ("\n\t" + (("a" + "a") + ("foo bar" + ("hello" + ("foo bar" + ("hello" + "a"))))))}
            fmt.Println(v3, v2, v3)
            v2 = (!v2)
            _ = v5
        }
    }
    if (4294967295 == 32768) {
        var v6 uint = 1
        _ = v6
    }
    var v7 string = "hello"
    v1 = (v1 / ((v1) | 1))
    if v2 {
        v8 := GenS2{F0: 18446744073709551615, F1: 2147483647, F2: 255}
        var v9 int32 = 0
        v10 := []int{9223372036854775807}
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 float32 = 1e10
    if (true && v2) {
        if (1 == 2147483647) {
            v11 -= (0.0 * 0.0)
        } else {
            for fi12 := 0; fi12 < 15; fi12++ {
                fmt.Println(v11, v7, v11)
                v13 := []int{1, (fi12 * (fi12 << ((639) & 63)))}
                if (!(((v7 >= (v7 + "hello")) || (-2147483648 == 164)) || v2)) {
                    var v14 int = 1
                    if v2 {
                        v11 -= 3.14159
                    } else {
                        v15 := GenS3{F0: v2}
                        fmt.Println(v14, v14, v2)
                        _ = v15
                    }
                    var v16 uint = 528
                    _ = v14
                    _ = v16
                }
                _ = fi12
                _ = v13
            }
            if v2 {
                fmt.Println(v3, v1, v1)
            } else {
                v17 := GenS4{F0: -32768, F1: 1, F2: 159}
                _ = v17
            }
        }
        v11 -= (1e10 + ((v11 * v11) * -2.25))
        v11 -= float32(v3)
    } else {
        v18 := []int{457, 1}
        v7 = v7
        _ = v18
    }
    v7 = v7
    v19 := GenS5{F0: -967, F1: -480}
    v20 := GenS6{F0: 645, F1: 2}
    if v2 {
        v3 -= float64(v11)
    } else {
        v21 := []int{1, 910, -1, 744}
        switch "" {
        case "":
            for fi22 := 0; fi22 < 14; fi22++ {
                for fi23 := 0; fi23 < 3; fi23++ {
                    v21 = append(v21, 4294967295)
                    if true {
                        fmt.Println(v3, v1, v2)
                        if (((v3 - (3.14159 + (1.5 + v3))) > (0.0 - 1.5)) && (!v2)) {
                            v21 = append(v21, (fi23 & -2147483648))
                            v24 := GenS5{F0: fi23, F1: -1}
                            fmt.Println(fi23)
                            v1 = v1
                            v2 = ((128 == 127) && ((561 > 42) || ((0 != 1) && v2)))
                            v1 = 42
                            _ = v24
                        }
                    } else {
                        v11 -= v11
                        fmt.Println(fi22, v7)
                        var v25 float32 = 1.5
                        switch (fi22 * (fi23 * fi22)) {
                        case 1:
                            v7 = v7
                        }
                        _ = v25
                    }
                    _ = fi23
                }
                _ = fi22
            }
        case "foo bar":
            v26 := GenS4{F0: 663, F1: 0, F2: -32768}
            _ = v26
        }
        _ = v21
    }
    var v27 int32 = 1
    var v28 string = ""
    for fi29 := 0; fi29 < 12; fi29++ {
        v30 := []int{32767, (fi29 * fi29), -224, (^(fi29))}
        v28 = "a"
        fmt.Println(v7)
        _ = fi29
        _ = v30
    }
    fmt.Println(v11, v11, v27)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v3
    _ = v7
}
