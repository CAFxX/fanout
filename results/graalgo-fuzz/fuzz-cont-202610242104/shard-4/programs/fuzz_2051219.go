package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int32; F1 uint32; F2 uint32 }

func main() {
    var v1 uint8 = 0
    v2 := make(map[string]bool)
    var v3 int = 727
    var v4 float64 = 3.14159
    var v5 uint16 = 769
    switch "\n\t" {
    case "hello":
        v1 = v1
    case "":
        for fi6 := 0; fi6 < 12; fi6++ {
            if true {
                if true {
                    switch (v3 + -9223372036854775808) {
                    case -816:
                        v7 := []int{fi6}
                        _ = v7
                    case -613:
                        v8 := GenS1{F0: fi6, F1: (v5 | v5)}
                        _ = v8
                    }
                    switch 127 {
                    case -9223372036854775808:
                        var v9 int32 = 2147483647
                        _ = v9
                    default:
                        var v10 int8 = 1
                        _ = v10
                    }
                    v11 := GenS1{F0: fi6, F1: (v5 / ((1) | 1))}
                    _ = v11
                }
                v1 += v1
            } else {
                if true {
                    v3 = v3
                    v12 := GenS1{F0: (v3 - (v3 & fi6)), F1: v5}
                    _ = v12
                } else {
                    fmt.Println(v5)
                    var v13 uint64 = 4294967295
                    _ = v13
                }
            }
            v1 -= (v1 << ((715) & 63))
            for fi14 := 0; fi14 < 10; fi14++ {
                v2[""] = false
                _ = fi14
            }
            _ = fi6
        }
    }
    if false {
        v5 = (v5 % (((v5 | v5)) | 1))
        v4 = (v4 * (1.5 - v4))
    } else {
        v15 := GenS2{F0: (^(v3)), F1: "foo bar"}
        fmt.Println(v4)
        var v16 int8 = 0
        _ = v15
        _ = v16
    }
    if true {
        var v17 uint64 = 4294967295
        v4 += 1.5
        if (801 >= 42) {
            for fi18 := 0; fi18 < 7; fi18++ {
                v19 := GenS1{F0: (fi18 >> ((285) & 63)), F1: v5}
                for fi20 := 0; fi20 < 5; fi20++ {
                    if (-48 == 100) {
                        v5 -= v5
                    }
                    fmt.Println(fi20)
                    var v21 int = (v3 | 32768)
                    v1 += (v1 * (v1 + v1))
                    v22 := v2["foo bar"]
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                _ = fi18
                _ = v19
            }
        }
        _ = v17
    } else {
        v5 = v5
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        if false {
            if (80 >= 2147483648) {
                fmt.Println(v5, v1)
            }
        } else {
            v5 += (v5 + v5)
            var v24 int = fi23
            var v25 int = -1
            _ = v24
            _ = v25
        }
        delete(v2, "a")
        fmt.Println(fi23, v1, v3)
        _ = fi23
    }
    var v26 float64 = 1.5
    var v27 uint = 65535
    v28 := GenS2{F0: (-(v3)), F1: "hello"}
    v2[("" + "a")] = true
    v29 := v2[""]
    var v30 uint = 65535
    if ((^(v27)) >= v30) {
        var v31 uint8 = 217
        var v32 uint32 = 920
        _ = v31
        _ = v32
    } else {
        for fi33 := 0; fi33 < 1; fi33++ {
            v4 = v26
            var v34 int16 = 0
            var v35 uint32 = 2147483648
            _ = fi33
            _ = v34
            _ = v35
        }
    }
    switch v3 {
    case -772:
        v36 := GenS3{F0: 42, F1: 42, F2: 42}
        _ = v36
    }
    if (v26 > float64(float32((v26 * float64(0.0))))) {
        fmt.Println(v27, v29, v27)
        var v37 int64 = -670
        _ = v37
    } else {
        v1 -= 42
    }
    var v38 int16 = 1000
    var v39 int = v3
    for fi40 := 0; fi40 < 19; fi40++ {
        for fi41 := 0; fi41 < 7; fi41++ {
            delete(v2, ("\n\t" + (("hello" + "\n\t") + (("" + "\n\t") + "a"))))
            switch (fi40 - (fi40 * fi40)) {
            case -500:
                switch "\n\t" {
                case "foo bar":
                    v29 = (255 > (v5 >> ((1) & 63)))
                default:
                    var v42 uint16 = (^(v5))
                    v43 := GenS1{F0: v39, F1: 0}
                    v27 -= v30
                    _ = v42
                    _ = v43
                }
            default:
                v44 := []int{v39}
                _ = v44
            }
            _ = fi41
        }
        v1 = (v1 + (^(v1)))
        v30 -= (^(v27))
        _ = fi40
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
