package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint16; F2 int32 }

func main() {
    var v1 int32 = 1
    var v2 float32 = -2.25
    var v3 int64 = -9223372036854775808
    v4 := make(map[int]int)
    if (42 <= 0) {
        v5 := GenS1{F0: (^(v1))}
        delete(v4, 0)
        _ = v5
    }
    _ = len(v4)
    var v6 uint64 = 241
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := GenS1{F0: v1}
        v9 := []int{-2147483648, fi7}
        _ = fi7
        _ = v8
        _ = v9
    }
    if true {
        v10 := GenS2{F0: -81, F1: 32767, F2: v1}
        _ = v10
    } else {
        fmt.Println(v2, v3)
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        for fi12 := 0; fi12 < 13; fi12++ {
            switch fi12 {
            case 61:
                v13 := []int{fi11, fi12}
                _ = v13
            case 420:
                if ("a" > ("\n\t" + "\n\t")) {
                    var v14 int = (-(fi11))
                    var v15 bool = true
                    switch "" {
                    case "\n\t":
                        var v16 uint32 = 0
                        _ = v16
                    }
                    _ = v14
                    _ = v15
                }
            case 32768:
                if false {
                    v3 = v3
                    v17 := []int{fi11, fi11, (^(fi11))}
                    fmt.Println(v1)
                    _ = v17
                } else {
                    if (!(true && (!true))) {
                        v1 -= (v1 + v1)
                        if (682 >= 65535) {
                            switch (fi12 & fi12) {
                            case 128:
                                if false {
                                    v2 = 0.0
                                } else {
                                    fmt.Println(v3)
                                }
                            case 1:
                                var v18 uint64 = (v6 >> ((1) & 63))
                                _ = v18
                            }
                            v19 := GenS1{F0: (v1 * v1)}
                            _ = v19
                        } else {
                            v20 := GenS1{F0: (^(v1))}
                            _ = v20
                        }
                    } else {
                        if (((((float64(1.5) >= 0.0) || (v2 <= float32(3.14159))) && true) && (-96 <= -128)) || true) {
                            v21 := v4[(fi11 % ((32767) | 1))]
                            _ = v21
                        } else {
                            var v22 float64 = 3.14159
                            _ = v22
                        }
                        switch fi11 {
                        case -2147483648:
                            if true {
                                _ = len(v4)
                            } else {
                                var v23 uint = 4294967295
                                _ = v23
                            }
                        default:
                            v1 = 42
                        }
                    }
                }
            }
            v24 := GenS1{F0: v1}
            _ = fi12
            _ = v24
        }
        for fi25 := 0; fi25 < 17; fi25++ {
            var v26 bool = false
            v1 = (v1 / (((v1 / (((v1 >> ((139) & 63))) | 1))) | 1))
            var v27 float64 = (0.1 / -2.25)
            var v28 uint16 = 42
            _ = fi25
            _ = v26
            _ = v27
            _ = v28
        }
        for fi29 := 0; fi29 < 9; fi29++ {
            var v30 uint16 = 2
            v31 := v4[(fi11 ^ fi29)]
            var v32 bool = (!true)
            _ = fi29
            _ = v30
            _ = v31
            _ = v32
        }
        _ = fi11
    }
    for fi33 := 0; fi33 < 10; fi33++ {
        v34 := []int{(-(fi33))}
        v1 += 2147483647
        _ = fi33
        _ = v34
    }
    v35 := GenS1{F0: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v6
}
