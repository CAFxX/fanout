package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 string; F1 int32; F2 int64 }
type GenS6 struct { F0 int; F1 uint; F2 uint }

func main() {
    v1 := []uint32{209, 124, 42, 4294967295}
    var v2 uint64 = 2
    var v3 uint64 = 0
    var v4 float64 = 0.0
    var v5 string = ""
    var v6 uint64 = 1
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 int64 = 9223372036854775807
        switch (fi7 | -332) {
        case 9223372036854775807:
            var v9 bool = false
            _ = v9
        case -40:
            switch "a" {
            case "hello":
                v10 := GenS1{F0: 65535, F1: 0, F2: (v8 & (-(v8)))}
                _ = v10
            default:
                if len(v1) > 0 {
                    v1[0] = 461
                }
            }
        case 4294967295:
            v11 := GenS2{F0: 417, F1: v8, F2: (v6 * (v6 >> ((632) & 63)))}
            _ = v11
        default:
            v2 = 566
        }
        _ = fi7
        _ = v8
    }
    switch 317 {
    case 128:
        var v12 float64 = (float64(0.1) / float64(float32(float64(3.14159))))
        _ = v12
    case 32768:
        if len(v1) > 0 {
            v1[0] = 2
        }
    case 32767:
        if true {
            _ = len(v1)
            var v13 uint32 = 643
            _ = v13
        } else {
            for fi14 := 0; fi14 < 3; fi14++ {
                v15 := GenS3{F0: (v5 + (v5 + "\n\t")), F1: 0}
                fmt.Println(v6)
                if true {
                    v3 -= 18446744073709551615
                    v2 = 593
                }
                _ = fi14
                _ = v15
            }
            v1 = append(v1, 0)
        }
    default:
        if false {
            v2 = (v3 ^ 44)
            for fi16 := 0; fi16 < 6; fi16++ {
                var v17 uint16 = 65535
                _ = fi16
                _ = v17
            }
            var v18 int32 = -1
            _ = v18
        } else {
            for fi19 := 0; fi19 < 7; fi19++ {
                switch v5 {
                case "":
                    v20 := GenS4{F0: v5, F1: 42}
                    _ = v20
                }
                v3 -= 1
                _ = fi19
            }
        }
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        v1 = append(v1, 728)
        var v22 uint32 = 2147483647
        _ = fi21
        _ = v22
    }
    v1 = append(v1, 2147483647)
    v4 = v4
    var v23 string = "a"
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        if true {
            v24 := GenS5{F0: v23, F1: 42, F2: 0}
            v6 = (^(v2))
            var v25 string = v5
            _ = v24
            _ = v25
        }
    default:
        if ((42 != 2) && ("foo bar" != (v5 + v23))) {
            switch v5 {
            case "foo bar":
                for fi26 := 0; fi26 < 11; fi26++ {
                    var v27 float64 = ((v4 - (v4 / float64((float32((v4 - 0.1)) * (((1.5 / (((((1.5 + 0.1) * float32(v4)) * ((0.1 + float32(0.0)) * (1.5 * ((((0.1 - float32(1e10)) * 0.0) - 1.5) + 1.5)))) - 1e10) - 1.5)) - (float32(3.14159) - -2.25)) / float32(1e10)))))) + ((1.5 - 3.14159) + (-2.25 * 0.1)))
                    fmt.Println(v3, v27)
                    for fi28 := 0; fi28 < 7; fi28++ {
                        var v29 uint = 2
                        var v30 uint64 = v2
                        var v31 int = fi28
                        _ = fi28
                        _ = v29
                        _ = v30
                        _ = v31
                    }
                    _ = fi26
                    _ = v27
                }
            }
            v32 := GenS6{F0: 127, F1: 163, F2: 42}
            _ = v32
        } else {
            v1 = append(v1, 909)
            v4 += float64((((0.0 - 1e10) - (0.1 - 0.1)) * (-2.25 / (-2.25 + ((-2.25 + 1.5) / 1.5)))))
        }
    }
    _ = len(v1)
    var v33 uint = 1
    fmt.Println(v33, v33, v2)
    v5 = (("\n\t" + v5) + (v23 + "\n\t"))
    switch ((v5 + "foo bar") + (v23 + v5)) {
    case "hello":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "\n\t":
        for fi34 := 0; fi34 < 4; fi34++ {
            var v35 int = (fi34 & (^(fi34)))
            var v36 int16 = 1
            if len(v1) > 0 {
                v1[0] = uint32(v33)
            }
            _ = fi34
            _ = v35
            _ = v36
        }
    case "foo bar":
        for fi37 := 0; fi37 < 3; fi37++ {
            if len(v1) > 0 {
                v1[0] = 2
            }
            if (-250 >= 0) {
                v2 = v6
                if false {
                    v38 := GenS2{F0: 1, F1: 948, F2: 42}
                    switch v23 {
                    case "\n\t":
                        var v39 bool = false
                        _ = v39
                    case "a":
                        for fi40 := 0; fi40 < 6; fi40++ {
                            v1 = append(v1, 1)
                            v4 = v4
                            _ = fi40
                        }
                    }
                    var v41 uint8 = 0
                    _ = v38
                    _ = v41
                }
                v2 = 225
            }
            if (1 >= 2) {
                v5 = "a"
                fmt.Println(v33, v3, v2)
            } else {
                for fi42 := 0; fi42 < 18; fi42++ {
                    var v43 int32 = 1
                    v3 += (v2 - 45)
                    var v44 string = v23
                    _ = fi42
                    _ = v43
                    _ = v44
                }
                switch (^(fi37)) {
                case 42:
                    v23 = v5
                }
                var v45 string = ("a" + ((("a" + v5) + ((((("hello" + "foo bar") + "a") + "hello") + "\n\t") + "")) + v23))
                _ = v45
            }
            _ = fi37
        }
    }
    v33 -= 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
