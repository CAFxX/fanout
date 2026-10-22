package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 string }

func main() {
    var v1 float32 = (1.5 / 1e10)
    var v2 uint32 = 392
    var v3 int64 = -536
    v4 := make(map[int]bool)
    var v5 uint8 = 127
    v6 := GenS1{F0: 346, F1: -9223372036854775808}
    v5 = 127
    v5 = 76
    for fi7 := 0; fi7 < 19; fi7++ {
        if ((((1 < 0) || ((100 != -90) || (1 < 65535))) && ((0 < 18446744073709551615) && (float32((((0.1 + 1e10) / float64((v1 - 0.0))) * 0.0)) > float32(1.5)))) || ((v2 ^ (v2 - (v2 - 0))) == v2)) {
            var v8 uint8 = (v5 / (((v5 - (v5 % ((v5) | 1)))) | 1))
            _ = v8
        }
        var v9 string = "foo bar"
        v10 := GenS1{F0: 4294967295, F1: fi7}
        _ = fi7
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: (-(v3))}
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 int32 = 790
        switch "hello" {
        case "foo bar":
            for fi14 := 0; fi14 < 19; fi14++ {
                var v15 int = fi14
                if (42 <= v13) {
                    var v16 int32 = -914
                    _ = v16
                }
                switch "\n\t" {
                case "a":
                    if (18446744073709551615 != 993) {
                        _ = len(v4)
                    } else {
                        fmt.Println(v1, fi14, v5)
                        switch "" {
                        case "hello":
                            v17 := []int{fi12, -9223372036854775808, (v15 / ((fi12) | 1))}
                            _ = v17
                        }
                    }
                case "":
                    _ = len(v4)
                }
                _ = fi14
                _ = v15
            }
        }
        if false {
            v18 := []int{}
            _ = v18
        } else {
            if true {
                if false {
                    v19 := []int{}
                    for fi20 := 0; fi20 < 4; fi20++ {
                        v1 -= v1
                        switch (fi12 - 32768) {
                        case 1:
                            _ = len(v19)
                        case -713:
                            v5 = (^(v5))
                            var v21 bool = (fi20 != 128)
                            if (((v13 % (((v13 * (-(v13)))) | 1)) < (v13 | -1)) || true) {
                                var v22 uint32 = (^(v2))
                                var v23 int32 = v13
                                v1 = v1
                                _ = len(v19)
                                _ = v22
                                _ = v23
                            }
                            _ = v21
                        case 127:
                            fmt.Println(v5, v2)
                        }
                        if len(v19) > 0 {
                            v19[0] = fi20
                        }
                        _ = len(v4)
                        _ = fi20
                    }
                    _ = v19
                } else {
                    delete(v4, (-(fi12)))
                }
                v24 := GenS1{F0: 665, F1: fi12}
                _ = v24
            }
        }
        _ = fi12
        _ = v13
    }
    var v25 uint8 = (v5 << ((42) & 63))
    if (true && false) {
        v3 += -234
        for fi26 := 0; fi26 < 7; fi26++ {
            if true {
                fmt.Println(fi26)
                v27 := []int{fi26, fi26, fi26}
                _ = v27
            }
            v25 = v5
            _ = fi26
        }
    } else {
        var v28 bool = false
        var v29 int = 32767
        _ = len(v4)
        _ = v28
        _ = v29
    }
    v2 = (v2 >> ((255) & 63))
    v30 := GenS1{F0: 224, F1: 42}
    v31 := GenS3{F0: 18446744073709551615, F1: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
