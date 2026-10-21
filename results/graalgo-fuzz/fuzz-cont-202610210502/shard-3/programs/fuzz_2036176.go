package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int16; F2 uint16 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 float32 = -2.25
    var v2 int = -1
    var v3 uint8 = 1
    var v4 float32 = (-2.25 / 1.5)
    var v5 int64 = -575
    v6 := make(map[uint32]string)
    if false {
        switch (v2 * v2) {
        case 914:
            for fi7 := 0; fi7 < 3; fi7++ {
                var v8 int16 = 0
                var v9 uint64 = 4294967295
                v5 += 572
                _ = fi7
                _ = v8
                _ = v9
            }
        default:
            for fi10 := 0; fi10 < 14; fi10++ {
                switch fi10 {
                case 127:
                    switch (fi10 + (^(v2))) {
                    case 127:
                        var v11 uint64 = 2
                        _ = v11
                    }
                case -9223372036854775808:
                    switch (^(v2)) {
                    case 206:
                        var v12 uint32 = 260
                        _ = v12
                    default:
                        v4 += float32((1.5 / (-2.25 + 1e10)))
                    }
                case 128:
                    var v13 string = "\n\t"
                    _ = v13
                default:
                    for fi14 := 0; fi14 < 12; fi14++ {
                        v15 := v6[2147483648]
                        var v16 bool = (v1 != v1)
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                }
                if false {
                    fmt.Println(v2, v3)
                    v5 = v5
                } else {
                    var v17 uint16 = 844
                    fmt.Println(v17, v3, v5)
                    _ = v17
                }
                _ = fi10
            }
        }
        switch "hello" {
        case "":
            fmt.Println(v1, v4)
        case "foo bar":
            switch (v2 - -2147483648) {
            case 507:
                fmt.Println(v3)
            }
        case "a":
            v18 := GenS1{F0: v2}
            _ = v18
        default:
            var v19 int32 = -2147483648
            _ = v19
        }
        var v20 uint64 = 42
        _ = v20
    }
    fmt.Println(v5, v2)
    var v21 int = (v2 >> ((uint(v3)) & 63))
    v22 := GenS1{F0: (v2 % (((v21 - v2)) | 1))}
    for fi23 := 0; fi23 < 18; fi23++ {
        v1 = (v4 - v4)
        for fi24 := 0; fi24 < 7; fi24++ {
            var v25 uint64 = 0
            if true {
                v3 += (^(v3))
            }
            _ = fi24
            _ = v25
        }
        var v26 uint8 = (v3 & v3)
        _ = fi23
        _ = v26
    }
    v27 := GenS2{F0: "hello", F1: -32768, F2: 850}
    var v28 int64 = (v5 ^ v5)
    var v29 uint32 = 42
    v30 := GenS2{F0: (("hello" + "foo bar") + "hello"), F1: 0, F2: 32767}
    switch 42 {
    case 128:
        v21 -= v2
    case 1:
        switch ("a" + ((("\n\t" + ("foo bar" + "")) + (("hello" + ("hello" + "")) + ("\n\t" + "foo bar"))) + "a")) {
        case "foo bar":
            for fi31 := 0; fi31 < 7; fi31++ {
                v32 := GenS1{F0: 32768}
                _ = fi31
                _ = v32
            }
        case "a":
            fmt.Println(v1)
        default:
            _ = len(v6)
        }
    default:
        v33 := GenS3{F0: -886}
        _ = v33
    }
    var v34 int16 = 63
    switch 32767 {
    case -977:
        var v35 int16 = (v34 & v34)
        _ = v35
    case 303:
        fmt.Println(v29)
    case -2147483648:
        var v36 int32 = 1
        _ = v36
    }
    v37 := []int{(^(v2)), 127, v2, v2}
    switch "hello" {
    case "":
        v38 := v6[(^(v29))]
        _ = v38
    default:
        if (true || ((!true) || false)) {
            v37 = append(v37, (v21 * v21))
            v39 := GenS3{F0: 644}
            var v40 float64 = 1e10
            _ = v39
            _ = v40
        } else {
            if true {
                switch "\n\t" {
                case "":
                    if (false || (65535 >= 65535)) {
                        switch 92 {
                        case 0:
                            var v41 int = v2
                            _ = v41
                        case 588:
                            if len(v37) > 0 {
                                v37[0] = 127
                            }
                        }
                        v2 -= (v2 + v21)
                        var v42 int64 = (v28 & v28)
                        _ = v42
                    } else {
                        fmt.Println(v1, v3, v3)
                        var v43 uint8 = (^(v3))
                        v4 -= (0.0 + -2.25)
                        _ = v43
                    }
                }
                switch "\n\t" {
                case "":
                    var v44 uint = 265
                    _ = v44
                }
            }
        }
    }
    if len(v37) >= 2 {
        v37 = v37[1:len(v37)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
