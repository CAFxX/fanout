package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int32; F1 uint64; F2 int32 }
type GenS3 struct { F0 int }

func main() {
    var v1 int = 32768
    v2 := make(map[uint32]string)
    var v3 string = "a"
    var v4 float32 = 0.1
    v5 := []string{v3, "\n\t"}
    for fi6 := 0; fi6 < 14; fi6++ {
        if (((4294967295 == 414) && (false && (-414 <= 1))) || false) {
            var v7 string = v3
            var v8 int64 = 42
            _ = v7
            _ = v8
        } else {
            fmt.Println(v1, v4)
        }
        if (127 != (v1 >> ((0) & 63))) {
            v9 := v2[2]
            v10 := GenS1{F0: "foo bar", F1: 65535}
            var v11 uint16 = 65535
            _ = v10
            _ = v11
            _ = v9
        } else {
            v4 += (1e10 + ((-2.25 + 1e10) / -2.25))
            for fi12 := 0; fi12 < 12; fi12++ {
                v1 = (fi12 / ((9223372036854775807) | 1))
                if (v3 != (v3 + "")) {
                    var v13 int32 = 0
                    _ = len(v5)
                    var v14 int8 = 1
                    _ = v13
                    _ = v14
                }
                v15 := GenS1{F0: ((v3 + v3) + "foo bar"), F1: 4294967295}
                _ = fi12
                _ = v15
            }
            var v16 int32 = 0
            _ = v16
        }
        _ = fi6
    }
    var v17 int = -133
    v1 = (v17 | -810)
    var v18 uint32 = 162
    var v19 bool = (-1 < int32(v1))
    var v20 int64 = 0
    v21 := GenS1{F0: "a", F1: 989}
    v22 := GenS1{F0: (v3 + (((("a" + "hello") + "foo bar") + ("a" + "foo bar")) + v3)), F1: 4294967295}
    switch 425 {
    case 1:
        _ = len(v2)
    case 42:
        _ = len(v2)
    }
    var v23 uint16 = 26
    var v24 uint64 = 42
    if (true && (!(!false))) {
        if v19 {
            for fi25 := 0; fi25 < 6; fi25++ {
                var v26 int16 = -951
                fmt.Println(v17)
                _ = fi25
                _ = v26
            }
        }
        if (float32((((((1e10 + (0.0 - (-2.25 / 0.1))) - 3.14159) + (((3.14159 / 1.5) + 0.0) - 0.1)) * 0.0) + 1.5)) != v4) {
            v5 = append(v5, (v3 + "foo bar"))
            v27 := GenS2{F0: 2147483647, F1: (^(v24)), F2: -2147483648}
            _ = v27
        }
    } else {
        if true {
            var v28 int8 = 42
            v23 = (v23 % ((v23) | 1))
            if v19 {
                v28 += 0
                var v29 string = (((v3 + ("hello" + (((v3 + "a") + (("" + "") + "a")) + (v3 + (v3 + "foo bar"))))) + v3) + (v3 + "a"))
                _ = v29
            } else {
                fmt.Println(v19, v17, v3)
            }
            _ = v28
        } else {
            fmt.Println(v1, v20)
            if (false && false) {
                var v30 uint = 526
                v31 := GenS3{F0: (v17 + 42)}
                _ = v30
                _ = v31
            } else {
                v32 := GenS1{F0: "foo bar", F1: 0}
                if (uint64(v23) != (v24 + uint64(v23))) {
                    v33 := v2[808]
                    if (!v19) {
                        for fi34 := 0; fi34 < 4; fi34++ {
                            for fi35 := 0; fi35 < 7; fi35++ {
                                v23 += v23
                                fmt.Println(v19, v17, v3)
                                var v36 string = v33
                                _ = fi35
                                _ = v36
                            }
                            _ = fi34
                        }
                        _ = len(v2)
                    }
                    var v37 int32 = int32(v17)
                    _ = v33
                    _ = v37
                } else {
                    var v38 float64 = float64(v4)
                    v17 = (v1 >> ((666) & 63))
                    v1 = 32768
                    _ = v38
                }
                var v39 uint16 = v23
                _ = v32
                _ = v39
            }
        }
        v40 := GenS3{F0: v1}
        _ = v40
    }
    v41 := v2[(v18 ^ 717)]
    var v42 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
