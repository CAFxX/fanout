package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 0
    var v3 float64 = ((1.5 + (1e10 / -2.25)) / 3.14159)
    var v4 uint64 = 42
    v5 := []string{("\n\t" + ("\n\t" + "")), "foo bar", "a", "a"}
    if true {
        for fi6 := 0; fi6 < 12; fi6++ {
            var v7 int64 = 42
            switch fi6 {
            case -1:
                var v8 uint8 = 255
                _ = v8
            case 4294967295:
                if len(v5) > 0 {
                    v5[0] = (("a" + "foo bar") + "")
                }
            case 32767:
                for fi9 := 0; fi9 < 10; fi9++ {
                    v7 = 1
                    var v10 float32 = -2.25
                    _ = fi9
                    _ = v10
                }
            }
            _ = fi6
            _ = v7
        }
        v4 = v4
        var v11 float64 = (0.0 + 0.0)
        _ = v11
    }
    switch "hello" {
    case "":
        var v12 float64 = v3
        _ = v12
    case "a":
        fmt.Println(v3)
    }
    v13 := GenS1{F0: 65535}
    for fi14 := 0; fi14 < 2; fi14++ {
        if (-353 == 1) {
            fmt.Println(v4, v3, v3)
            for fi15 := 0; fi15 < 4; fi15++ {
                fmt.Println(v4, v3)
                var v16 float64 = ((float64(float32(0.0)) + float64((3.14159 + float32(0.0)))) - ((0.1 - v3) - v3))
                _ = fi15
                _ = v16
            }
            v3 -= ((v3 / float64((-2.25 - 1e10))) / (float64(1.5) * float64((1e10 / -2.25))))
        }
        var v17 float32 = -2.25
        for fi18 := 0; fi18 < 17; fi18++ {
            fmt.Println(fi18, v17)
            switch (((("foo bar" + "\n\t") + "\n\t") + ("foo bar" + ("a" + ""))) + "foo bar") {
            case "a":
                v19 := GenS2{F0: (-(fi18))}
                _ = v19
            default:
                if (false && true) {
                    var v20 uint64 = 229
                    fmt.Println(fi18)
                    v20 += 11
                    var v21 uint = 42
                    var v22 float32 = v17
                    if false {
                        v20 += 731
                        v23 := GenS1{F0: v21}
                        var v24 uint32 = 632
                        _ = v23
                        _ = v24
                    } else {
                        if (0 == -925) {
                            v25 := GenS1{F0: (v21 * (v21 + (v21 ^ 42)))}
                            var v26 int = fi18
                            _ = v25
                            _ = v26
                        }
                    }
                    _ = v20
                    _ = v21
                    _ = v22
                }
            }
            var v27 float32 = float32(v3)
            if false {
                fmt.Println(fi14)
                if false {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v17 = 1.5
                    fmt.Println(v17)
                }
                v2 = (^(v2))
                v28 := GenS2{F0: (fi18 | fi18)}
                v29 := GenS2{F0: 987}
                _ = v28
                _ = v29
            }
            _ = fi18
            _ = v27
        }
        _ = fi14
        _ = v17
    }
    _ = len(v5)
    var v30 int32 = 42
    var v31 int8 = 100
    v32 := GenS2{F0: 9223372036854775807}
    v4 = 2
    v33 := GenS3{F0: v31, F1: 254}
    var v34 uint16 = 0
    if (true || (2147483647 < 646)) {
        for fi35 := 0; fi35 < 1; fi35++ {
            if true {
                var v36 int8 = -18
                var v37 uint32 = 2147483648
                v38 := v1[2147483647]
                _ = v36
                _ = v37
                _ = v38
            } else {
                v39 := GenS4{F0: 809}
                var v40 uint = 2
                v41 := v1[42]
                _ = v39
                _ = v40
                _ = v41
            }
            _ = fi35
        }
    }
    var v42 bool = (("foo bar" == (("foo bar" + "foo bar") + "foo bar")) && (1e10 > 1e10))
    if v42 {
        _ = len(v1)
        var v43 int64 = -351
        _ = v43
    } else {
        v44 := GenS2{F0: -2147483648}
        _ = v44
    }
    var v45 int8 = -15
    if ((1 != 1) || (((3.14159 + (float32(-2.25) - 1.5)) < 3.14159) && ((4294967295 != 2) && (("\n\t" >= "foo bar") || v42)))) {
        fmt.Println(v2)
        v2 = (^(v2))
    } else {
        for fi46 := 0; fi46 < 18; fi46++ {
            fmt.Println(v34)
            fmt.Println(v3, v42)
            _ = fi46
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v42
    _ = v45
    _ = v5
}
