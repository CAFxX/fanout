package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint32 }
func genFn0(p0 string, p1 uint64) uint8 {
    var v7 uint32 = 194
    v8 := []int{}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint8 = 42
    var v2 uint64 = 1
    var v3 uint8 = (v1 - 42)
    v4 := make(map[int]bool)
    var v5 string = ""
    var v6 int8 = 1
    v9 := GenS1{F0: 0, F1: v3}
    v3 = (v3 >> ((42) & 63))
    v10 := []int{32768, 727, 0, -1}
    v3 = 0
    var v11 int = 32768
    for fi12 := 0; fi12 < 1; fi12++ {
        if true {
            v5 = v5
            v6 += v6
            fmt.Println(fi12, fi12)
        } else {
            v13 := GenS1{F0: 1, F1: genFn0(v5, 1)}
            _ = v13
        }
        _ = fi12
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        fmt.Println(fi14, v3)
        for fi15 := 0; fi15 < 4; fi15++ {
            switch fi14 {
            case -1:
                v4[(fi15 * (v11 + (fi15 % (((v11 >> ((uint(v2)) & 63))) | 1))))] = true
                _ = len(v4)
                var v16 int64 = 42
                _ = v16
            }
            if false {
                fmt.Println(v5, v2)
                var v17 uint8 = (v3 ^ (v1 / ((v3) | 1)))
                v1 = 42
                v11 = v11
                _ = v17
            } else {
                var v18 int64 = 42
                v19 := GenS1{F0: -975, F1: (v3 << ((2) & 63))}
                var v20 int16 = 0
                _ = v18
                _ = v19
                _ = v20
            }
            var v21 uint64 = 359
            _ = fi15
            _ = v21
        }
        var v22 int64 = 0
        _ = fi14
        _ = v22
    }
    v11 += int(v3)
    switch (v11 >> ((830) & 63)) {
    case 127:
        v10 = append(v10, 813)
    }
    if (true && (!false)) {
        v2 += 499
    }
    if ((v1 % ((42) | 1)) > (^(v3))) {
        v23 := GenS2{F0: 948, F1: 880}
        v10 = append(v10, v11)
        _ = v23
    } else {
        if (!true) {
            v24 := GenS1{F0: -1, F1: genFn0("\n\t", v2)}
            if (true || true) {
                for fi25 := 0; fi25 < 16; fi25++ {
                    var v26 uint = 42
                    var v27 int32 = 42
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            } else {
                v28 := GenS3{F0: -946}
                fmt.Println(v6, v11)
                _ = v28
            }
            _ = v24
        } else {
            switch (v5 + "foo bar") {
            case "\n\t":
                v10 = append(v10, (-(v11)))
            default:
                v10 = append(v10, 1)
            }
            v29 := GenS4{F0: (v6 <= (v6 / ((-128) | 1))), F1: v5}
            if true {
                for fi30 := 0; fi30 < 15; fi30++ {
                    var v31 float64 = (0.1 - 3.14159)
                    if (float64(1.5) <= v31) {
                        v32 := GenS3{F0: 679}
                        fmt.Println(v6, v11)
                        var v33 string = v5
                        _ = v32
                        _ = v33
                    }
                    v34 := GenS5{F0: 42}
                    _ = fi30
                    _ = v31
                    _ = v34
                }
            }
            _ = v29
        }
        v2 += 972
        for fi35 := 0; fi35 < 2; fi35++ {
            var v36 uint8 = (v1 + 0)
            switch 127 {
            case 4294967295:
                if (4294967295 <= 42) {
                    v36 -= (v36 + v1)
                    v37 := v4[(fi35 / ((1) | 1))]
                    _ = v37
                } else {
                    v10 = append(v10, (fi35 % ((int(v6)) | 1)))
                }
            case 32768:
                if true {
                    v3 = genFn0(v5, (v2 << ((uint(v1)) & 63)))
                    for fi38 := 0; fi38 < 12; fi38++ {
                        var v39 uint32 = 4294967295
                        fmt.Println(v1, fi35, v5)
                        _ = fi38
                        _ = v39
                    }
                } else {
                    v4[fi35] = (!(2147483647 >= 1))
                    for fi40 := 0; fi40 < 5; fi40++ {
                        v41 := GenS6{F0: 9223372036854775807}
                        _ = fi40
                        _ = v41
                    }
                    var v42 int64 = 42
                    _ = v42
                }
            case 0:
                var v43 uint = 258
                _ = v43
            }
            v6 = v6
            _ = fi35
            _ = v36
        }
    }
    v44 := GenS7{F0: 512}
    delete(v4, (-(v11)))
    v5 = ""
    delete(v4, v11)
    switch (((v5 + v5) + ("foo bar" + ("a" + v5))) + v5) {
    case "hello":
        v45 := v4[(v11 / ((-838) | 1))]
        _ = v45
    case "\n\t":
        var v46 float32 = 0.1
        _ = v46
    }
    v1 -= genFn0(("a" + ""), v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
    _ = v9
}
