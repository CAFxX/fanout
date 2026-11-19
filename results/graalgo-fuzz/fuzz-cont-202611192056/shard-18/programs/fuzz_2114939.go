package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int; F2 int8 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 int32; F1 bool }
func genFn0(p0 int8, p1 int16) uint16 {
    p1 = p1
    v5 := GenS1{F0: 42}
    for fi6 := 0; fi6 < 15; fi6++ {
        p1 -= p1
        if (p1 > (p1 / (((p1 & -972)) | 1))) {
            var v7 float32 = (3.14159 / (float32(float64(3.14159)) / (1.5 + (1e10 * (float32(float64(-2.25)) - (0.0 * (1.5 / float32((float64((-2.25 / 1.0)) * float64(1e10))))))))))
            v7 += v7
            p0 = 127
            _ = v7
        } else {
            var v8 int32 = -362
            switch fi6 {
            case -639:
                var v9 int32 = (-(v8))
                _ = v9
            case 4294967295:
                v10 := GenS1{F0: 1}
                _ = v10
            }
            _ = v8
        }
        _ = fi6
    }
    var v11 uint32 = 992
    for fi12 := 0; fi12 < 20; fi12++ {
        if true {
            var v13 int32 = -2147483648
            var v14 float32 = 3.14159
            switch ("hello" + "a") {
            case "\n\t":
                if (16 <= p0) {
                    var v15 bool = (int(p1) != (fi12 % ((fi12) | 1)))
                    var v16 int64 = -9223372036854775808
                    var v17 int64 = int64(p0)
                    _ = v15
                    _ = v16
                    _ = v17
                } else {
                    var v18 string = "\n\t"
                    if false {
                        v19 := []int{}
                        var v20 float32 = v14
                        _ = v19
                        _ = v20
                    }
                    if (true && false) {
                        if false {
                            fmt.Println(p1, v11, v11)
                            var v21 uint32 = (v11 / ((322) | 1))
                            _ = v21
                        }
                    } else {
                        switch ("\n\t" + (("a" + (v18 + (v18 + (("hello" + "foo bar") + (("\n\t" + "") + ""))))) + v18)) {
                        case "\n\t":
                            p0 += (p0 - (p0 + (p0 >> ((702) & 63))))
                        case "hello":
                            p0 += (p0 * p0)
                        default:
                            var v22 string = v18
                            _ = v22
                        }
                        p1 -= p1
                        v23 := []int{int(v13), fi12, (fi12 - -9223372036854775808)}
                        _ = v23
                    }
                    _ = v18
                }
            case "a":
                var v24 uint8 = 255
                _ = v24
            case "hello":
                for fi25 := 0; fi25 < 5; fi25++ {
                    v14 -= v14
                    v26 := []int{-2147483648}
                    v27 := GenS2{F0: 0, F1: fi25, F2: (^(p0))}
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            default:
                v28 := []int{2147483647, -1, (fi12 * 32768), (fi12 / (((fi12 | (fi12 ^ fi12))) | 1))}
                _ = v28
            }
            _ = v13
            _ = v14
        } else {
            v29 := GenS3{F0: (p0 & 97), F1: (v11 % ((488) | 1))}
            _ = v29
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v5
    return 255
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 0
    var v3 uint32 = 0
    var v4 uint8 = 128
    v30 := []int{0, 32767, -908}
    delete(v1, -895)
    var v31 uint32 = v3
    fmt.Println(v3, v2)
    if (568 > 844) {
        var v32 bool = (false && true)
        v3 += 1
        fmt.Println(v2, v31, v32)
        _ = v32
    } else {
        var v33 float32 = (((1e10 - (0.1 - 0.1)) * (0.0 * 0.0)) / 3.14159)
        v34 := GenS2{F0: 404, F1: 127, F2: (v2 | v2)}
        _ = v33
        _ = v34
    }
    fmt.Println(v3, v3, v31)
    v1[-125] = "foo bar"
    var v35 int16 = 1
    for fi36 := 0; fi36 < 11; fi36++ {
        if false {
            fmt.Println(v31, fi36)
        } else {
            switch (("" + "\n\t") + "") {
            case "foo bar":
                var v37 int = -2147483648
                _ = v37
            default:
                v38 := GenS4{F0: 42, F1: (0 >= 0)}
                _ = v38
            }
        }
        _ = fi36
    }
    v4 = v4
    switch "hello" {
    case "foo bar":
        v4 = v4
    case "a":
        v39 := v1[128]
        _ = v39
    default:
        v40 := GenS4{F0: 1, F1: false}
        _ = v40
    }
    var v41 int8 = (v2 % ((int8(v31)) | 1))
    v2 += 127
    v31 += 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v4
    _ = v41
}
