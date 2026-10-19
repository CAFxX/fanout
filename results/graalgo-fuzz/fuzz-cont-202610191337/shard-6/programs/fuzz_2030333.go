package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS4 struct { F0 int8 }

func main() {
    var v1 bool = (!(!false))
    v2 := make(map[string]int)
    v3 := make(map[string]uint32)
    var v4 bool = (v1 || false)
    v5 := []uint32{0, 289, 186, 42}
    var v6 uint64 = 2
    v7 := v3["\n\t"]
    v4 = v4
    if v4 {
        var v8 int = 394
        var v9 string = "a"
        _ = v8
        _ = v9
    }
    _ = len(v5)
    if true {
        v1 = v1
        if true {
            var v10 int = 125
            if (v4 || false) {
                fmt.Println(v7, v1)
                v4 = false
            }
            switch (v10 + v10) {
            case 1:
                var v11 uint16 = 2
                _ = v11
            case 0:
                for fi12 := 0; fi12 < 2; fi12++ {
                    var v13 uint32 = 116
                    if v4 {
                        var v14 uint32 = (^(v7))
                        _ = v14
                    }
                    _ = fi12
                    _ = v13
                }
            default:
                v2[("hello" + "")] = v10
            }
            _ = v10
        }
    } else {
        var v15 uint = 65535
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v15
    }
    v1 = v4
    var v16 string = "a"
    var v17 float64 = (1e10 - (0.1 / float64((1.5 * float32((1e10 + -2.25))))))
    v3[((v16 + "\n\t") + "a")] = 1
    if ((3.14159 >= 0.1) || (65535 == 42)) {
        for fi18 := 0; fi18 < 5; fi18++ {
            switch fi18 {
            case 127:
                fmt.Println(v4, v6)
            case 4294967295:
                switch 102 {
                case 9223372036854775807:
                    v19 := GenS1{F0: 0}
                    _ = v19
                case -499:
                    var v20 uint = 943
                    _ = v20
                case 4294967295:
                    for fi21 := 0; fi21 < 10; fi21++ {
                        var v22 int64 = -1
                        v23 := GenS2{F0: (^(fi18))}
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                default:
                    v24 := GenS3{F0: 134, F1: -9223372036854775808, F2: (v6 / (((^(v6))) | 1))}
                    _ = v24
                }
            case 242:
                fmt.Println(fi18, v6, v6)
            default:
                if len(v5) > 0 {
                    v5[0] = 929
                }
            }
            for fi25 := 0; fi25 < 2; fi25++ {
                if v4 {
                    if len(v5) > 0 {
                        v5[0] = 65535
                    }
                } else {
                    v26 := GenS1{F0: -1}
                    if (!((!(!true)) && false)) {
                        v4 = (128 <= 232)
                        v27 := GenS2{F0: -762}
                        _ = v27
                    }
                    _ = v26
                }
                _ = fi25
            }
            v16 = (v16 + "foo bar")
            _ = fi18
        }
        var v28 float64 = v17
        _ = v28
    }
    var v29 int = 9223372036854775807
    for fi30 := 0; fi30 < 1; fi30++ {
        v29 = v29
        v31 := GenS3{F0: 2, F1: 650, F2: v6}
        v17 = -2.25
        _ = fi30
        _ = v31
    }
    v17 -= ((-2.25 / v17) + v17)
    for fi32 := 0; fi32 < 7; fi32++ {
        v33 := GenS4{F0: 42}
        switch (v16 + v16) {
        case "hello":
            var v34 float64 = ((0.1 / v17) / ((3.14159 + 1.5) / v17))
            _ = v34
        case "\n\t":
            for fi35 := 0; fi35 < 17; fi35++ {
                v3[("foo bar" + "a")] = 42
                v16 = ""
                var v36 int8 = 30
                _ = len(v5)
                _ = fi35
                _ = v36
            }
        default:
            var v37 int8 = 127
            _ = v37
        }
        _ = fi32
        _ = v33
    }
    var v38 int = 0
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
