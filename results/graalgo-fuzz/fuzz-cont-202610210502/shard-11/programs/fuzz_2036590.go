package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool; F1 uint32; F2 int8 }
func genFn0(p0 bool, p1 int8) uint8 {
    p0 = p0
    var v5 uint8 = 1
    _ = p0
    _ = p1
    _ = v5
    return v5
}

func main() {
    var v1 bool = false
    v2 := make(map[string]uint32)
    var v3 int16 = -1
    var v4 float32 = 1e10
    for fi6 := 0; fi6 < 20; fi6++ {
        v3 = (v3 * v3)
        _ = len(v2)
        if ((-1 < 42) || (!false)) {
            fmt.Println(v1, fi6)
            fmt.Println(v4, v1)
        } else {
            v1 = v1
        }
        _ = fi6
    }
    if (1 != 2) {
        switch 130 {
        case -588:
            switch 901 {
            case 32767:
                fmt.Println(v3)
            default:
                if (true || false) {
                    var v7 float32 = 1e10
                    _ = v7
                } else {
                    var v8 int16 = (v3 & -32768)
                    v9 := GenS1{F0: 32767, F1: (1 < -9223372036854775808), F2: 112}
                    _ = v8
                    _ = v9
                }
            }
        case 4294967295:
            for fi10 := 0; fi10 < 2; fi10++ {
                var v11 uint32 = 268
                for fi12 := 0; fi12 < 1; fi12++ {
                    v1 = ((!(!false)) && (false || true))
                    v13 := GenS2{F0: 42}
                    if (1 == -674) {
                        var v14 uint16 = 504
                        if (4294967295 != 1) {
                            var v15 int16 = v3
                            _ = v15
                        } else {
                            v16 := GenS1{F0: v14, F1: (132 <= 173), F2: 1}
                            _ = v16
                        }
                        var v17 uint64 = 398
                        switch (fi10 / ((fi10) | 1)) {
                        case -210:
                            var v18 bool = (0 > v11)
                            _ = v18
                        }
                        v19 := GenS1{F0: v14, F1: v1, F2: 0}
                        _ = v14
                        _ = v17
                        _ = v19
                    }
                    _ = fi12
                    _ = v13
                }
                _ = fi10
                _ = v11
            }
        case 127:
            v4 = v4
        }
        for fi20 := 0; fi20 < 15; fi20++ {
            switch (fi20 | 42) {
            case 4294967295:
                for fi21 := 0; fi21 < 5; fi21++ {
                    v22 := GenS3{F0: 108}
                    _ = fi21
                    _ = v22
                }
            default:
                var v23 bool = v1
                _ = v23
            }
            switch -1 {
            case -450:
                fmt.Println(v4, fi20, v1)
            default:
                v24 := []int{fi20, fi20}
                _ = v24
            }
            fmt.Println(v1, v3)
            _ = fi20
        }
        for fi25 := 0; fi25 < 15; fi25++ {
            var v26 float32 = v4
            var v27 int16 = (v3 ^ v3)
            _ = fi25
            _ = v26
            _ = v27
        }
    } else {
        var v28 float64 = (0.0 + (0.0 - 1.5))
        var v29 uint8 = 42
        _ = v28
        _ = v29
    }
    switch "" {
    case "hello":
        for fi30 := 0; fi30 < 6; fi30++ {
            v4 -= 3.14159
            _ = fi30
        }
    case "\n\t":
        v31 := GenS3{F0: 42}
        _ = v31
    }
    v32 := GenS2{F0: 2}
    for fi33 := 0; fi33 < 18; fi33++ {
        var v34 int64 = -516
        var v35 string = "hello"
        if true {
            var v36 int8 = 122
            v2[v35] = 233
            var v37 uint8 = 128
            _ = v36
            _ = v37
        } else {
            if (!v1) {
                v38 := GenS4{F0: 133, F1: 2}
                v39 := GenS2{F0: 65535}
                _ = v38
                _ = v39
            }
            for fi40 := 0; fi40 < 10; fi40++ {
                v41 := GenS3{F0: 42}
                _ = fi40
                _ = v41
            }
        }
        _ = fi33
        _ = v34
        _ = v35
    }
    var v42 bool = (42 > 0)
    switch 32768 {
    case 0:
        v42 = (42 <= 42)
    }
    fmt.Println(v3)
    v42 = ((3.14159 * 1.5) <= -2.25)
    v1 = false
    fmt.Println(v42, v3, v1)
    switch "\n\t" {
    case "":
        v42 = v42
    case "hello":
        v3 -= (v3 / ((v3) | 1))
    case "\n\t":
        var v43 int = 4294967295
        _ = v43
    }
    v44 := GenS5{F0: 4294967295}
    delete(v2, "foo bar")
    var v45 int8 = 1
    v46 := GenS6{F0: v1, F1: 65535, F2: v45}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v42
    _ = v44
    _ = v45
    _ = v46
}
