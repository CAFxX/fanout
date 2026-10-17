package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 uint64; F1 uint8 }
func genFn0(p0 uint, p1 int16, p2 int32) int64 {
    v7 := GenS1{F0: 42}
    v8 := GenS2{F0: 0, F1: -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 43
}
func genFn1(p0 int8) int16 {
    p0 = 16
    v9 := []int{-2147483648, -1, 42}
    p0 += p0
    _ = p0
    _ = v9
    return -32768
}

func main() {
    var v1 float64 = ((0.1 / ((-2.25 + 1.5) + float64((0.1 - 0.0)))) * 3.14159)
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]bool)
    var v4 float64 = float64(1.5)
    v5 := make(map[string]bool)
    var v6 bool = false
    v10 := GenS2{F0: 243, F1: -839}
    for fi11 := 0; fi11 < 6; fi11++ {
        v4 -= v1
        if false {
            v12 := GenS2{F0: 255, F1: 2147483647}
            if (32767 >= 1000) {
                if (!(519 >= 2)) {
                    fmt.Println(v4, fi11)
                    fmt.Println(v1, v6)
                    v13 := []int{-1, fi11, (fi11 << ((4294967295) & 63))}
                    _ = v13
                } else {
                    switch fi11 {
                    case -671:
                        for fi14 := 0; fi14 < 6; fi14++ {
                            v1 = 0.1
                            v15 := GenS3{F0: 127, F1: -24}
                            switch (fi11 - 127) {
                            case -9223372036854775808:
                                var v16 int16 = 0
                                _ = v16
                            }
                            _ = fi14
                            _ = v15
                        }
                    }
                    var v17 uint8 = 1
                    _ = v17
                }
            } else {
                for fi18 := 0; fi18 < 2; fi18++ {
                    v6 = (false || v6)
                    _ = fi18
                }
            }
            _ = v12
        }
        v19 := GenS4{F0: 42, F1: 141}
        _ = fi11
        _ = v19
    }
    var v20 int32 = -1
    var v21 uint8 = 2
    var v22 int64 = genFn0(2, genFn1(84), 42)
    fmt.Println(v1, v4)
    v23 := v5["\n\t"]
    if (!v23) {
        var v24 int = 127
        if v6 {
            v25 := v5["a"]
            var v26 int = (v24 * 32768)
            v1 -= float64(3.14159)
            _ = v25
            _ = v26
        } else {
            for fi27 := 0; fi27 < 14; fi27++ {
                v4 = v4
                v28 := GenS2{F0: (v21 << ((2) & 63)), F1: fi27}
                v24 += (fi27 / ((2147483647) | 1))
                _ = fi27
                _ = v28
            }
            if true {
                var v29 bool = (1e10 < (((v4 + v4) * 3.14159) + v4))
                var v30 int16 = -32768
                v31 := []int{(v24 & 589)}
                _ = v29
                _ = v30
                _ = v31
            } else {
                var v32 uint8 = v21
                v33 := v3[(v24 & v24)]
                if ((1.5 - 0.0) == v1) {
                    v34 := GenS1{F0: (v2 + -13)}
                    v35 := GenS1{F0: v2}
                    if v6 {
                        v36 := []int{}
                        v37 := GenS1{F0: 42}
                        var v38 int64 = int64(v32)
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                    _ = v34
                    _ = v35
                } else {
                    v22 = (v22 >> ((0) & 63))
                    var v39 float64 = 1.5
                    var v40 int = 42
                    _ = v39
                    _ = v40
                }
                _ = v32
                _ = v33
            }
            var v41 float64 = (v4 - (((v1 / 1.5) * v1) - v4))
            _ = v41
        }
        _ = v24
    }
    if true {
        v42 := v3[int(v21)]
        fmt.Println(v1)
        var v43 int8 = 42
        _ = v42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
