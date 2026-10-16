package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 uint; F1 int16; F2 bool }
type GenS5 struct { F0 int32; F1 uint; F2 int16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]string)
    v3 := make(map[int]string)
    var v4 int8 = 1
    if (!false) {
        v4 -= 126
    }
    fmt.Println(v4)
    v5 := []int{1, 395, 1}
    var v6 bool = (true && false)
    v4 = v4
    if true {
        if len(v5) > 0 {
            v5[0] = -1
        }
        for fi7 := 0; fi7 < 16; fi7++ {
            v4 += (v4 % ((127) | 1))
            switch fi7 {
            case 1:
                var v8 int = (fi7 + fi7)
                _ = v8
            }
            if ((((0.0 / (1e10 / 1e10)) / float32((((1.5 / ((1.5 * -2.25) - 1.5)) * (0.0 * ((((0.0 / -2.25) - 3.14159) * 1.5) + 1e10))) + 1.5))) <= float32(-2.25)) || true) {
                var v9 float64 = (0.0 + 1e10)
                _ = v9
            }
            _ = fi7
        }
        v5 = append(v5, 1)
    } else {
        v10 := GenS1{F0: 0, F1: 42}
        switch -515 {
        case -764:
            v6 = false
        }
        _ = v10
    }
    if ((9223372036854775807 == 437) || (0.1 == (3.14159 + ((-2.25 / 1.0) * 1e10)))) {
        v4 += (^(v4))
        v11 := GenS1{F0: 1, F1: 1}
        v4 -= (v4 ^ 100)
        _ = v11
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        if (!(1 == 2147483647)) {
            var v13 uint16 = 0
            if (v4 < -52) {
                var v14 uint16 = v13
                _ = v14
            }
            if v6 {
                if (v4 == v4) {
                    var v15 float32 = (1.5 * float32(float64((((float32(0.0) * 0.1) * ((0.0 + 1e10) * 3.14159)) - 1e10))))
                    _ = v15
                } else {
                    v13 -= 1
                    _ = len(v2)
                    fmt.Println(v4, v4)
                }
                var v16 string = (("\n\t" + "a") + "foo bar")
                _ = v16
            } else {
                v13 += v13
                fmt.Println(v4, v6)
            }
            _ = v13
        } else {
            v17 := v2["\n\t"]
            v18 := v3[32767]
            _ = v17
            _ = v18
        }
        if v6 {
            v19 := GenS2{F0: 2147483647}
            for fi20 := 0; fi20 < 6; fi20++ {
                var v21 int32 = 488
                v21 = (-(v21))
                var v22 bool = (!(!v6))
                _ = fi20
                _ = v21
                _ = v22
            }
            for fi23 := 0; fi23 < 3; fi23++ {
                v24 := GenS3{F0: v6, F1: 1}
                _ = fi23
                _ = v24
            }
            _ = v19
        } else {
            if (fi12 != fi12) {
                v25 := GenS4{F0: 65535, F1: 986, F2: false}
                v4 = v4
                v4 -= (v4 << ((265) & 63))
                _ = v25
            } else {
                v26 := GenS2{F0: 0}
                _ = v26
            }
            var v27 uint32 = 63
            _ = v27
        }
        if v6 {
            switch 60 {
            case 128:
                var v28 int = fi12
                _ = v28
            case -200:
                v5 = append(v5, fi12)
            case 32767:
                v4 = (v4 % ((v4) | 1))
            }
            _ = len(v1)
            v29 := GenS5{F0: 2147483647, F1: 255, F2: -479}
            _ = v29
        } else {
            if v6 {
                fmt.Println(v4, fi12, v6)
            } else {
                v6 = (236 < 42)
                var v30 uint8 = 42
                v31 := GenS2{F0: 2}
                _ = v30
                _ = v31
            }
        }
        _ = fi12
    }
    v32 := GenS4{F0: 42, F1: 1, F2: v6}
    delete(v3, -9223372036854775808)
    v4 += 100
    if v6 {
        v33 := GenS6{F0: 1}
        var v34 bool = true
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 19; fi35++ {
        v36 := GenS4{F0: 255, F1: 32767, F2: v6}
        for fi37 := 0; fi37 < 6; fi37++ {
            var v38 uint16 = 2
            _ = fi37
            _ = v38
        }
        fmt.Println(v4, v6, v6)
        _ = fi35
        _ = v36
    }
    if (1 >= (v4 >> ((42) & 63))) {
        delete(v1, 2147483648)
        var v39 uint32 = 80
        _ = v39
    } else {
        v6 = v6
    }
    v40 := GenS4{F0: 4294967295, F1: 42, F2: ((!v6) || (8 < 2147483648))}
    v41 := GenS7{F0: (v4 % (((v4 - v4)) | 1))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
