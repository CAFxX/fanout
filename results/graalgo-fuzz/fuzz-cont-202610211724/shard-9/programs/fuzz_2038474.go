package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint32 }
type GenS2 struct { F0 string; F1 uint }

func main() {
    v1 := []int{-9223372036854775808, 32768}
    var v2 int64 = 42
    var v3 int64 = v2
    var v4 int64 = v3
    var v5 uint64 = 284
    var v6 uint = 1
    v7 := GenS1{F0: (152 < 427), F1: 0, F2: 4294967295}
    switch 1 {
    case 237:
        fmt.Println(v3, v6)
    default:
        if true {
            var v8 string = (((("a" + "a") + "hello") + "") + "")
            var v9 string = "hello"
            _ = v8
            _ = v9
        } else {
            v6 += v6
            v10 := GenS1{F0: (127 > 100), F1: (v3 ^ (v3 ^ (v3 + v2))), F2: 959}
            for fi11 := 0; fi11 < 9; fi11++ {
                if false {
                    for fi12 := 0; fi12 < 20; fi12++ {
                        var v13 uint = 872
                        _ = fi12
                        _ = v13
                    }
                } else {
                    var v14 uint32 = 2147483648
                    switch "" {
                    case "foo bar":
                        var v15 float64 = 0.1
                        _ = v15
                    default:
                        if (!((!false) || false)) {
                            var v16 uint16 = 944
                            if false {
                                if true {
                                    var v17 int16 = 0
                                    _ = v17
                                }
                                for fi18 := 0; fi18 < 11; fi18++ {
                                    var v19 uint8 = 1
                                    var v20 uint = 42
                                    _ = fi18
                                    _ = v19
                                    _ = v20
                                }
                                for fi21 := 0; fi21 < 9; fi21++ {
                                    var v22 float64 = 3.14159
                                    _ = fi21
                                    _ = v22
                                }
                            } else {
                                v3 = -149
                                var v23 float32 = 1e10
                                if (fi11 < fi11) {
                                    for fi24 := 0; fi24 < 4; fi24++ {
                                        switch 4294967295 {
                                        case -838:
                                            var v25 uint32 = (v14 - (v14 + v14))
                                            _ = v25
                                        }
                                        v26 := GenS1{F0: (128 <= 255), F1: v4, F2: v14}
                                        v6 += (^(v6))
                                        _ = fi24
                                        _ = v26
                                    }
                                }
                                _ = v23
                            }
                            if (106 == 128) {
                                if false {
                                    var v27 uint = v6
                                    if (!false) {
                                        var v28 int64 = v4
                                        if (v14 == (v14 * (v14 % ((795) | 1)))) {
                                            switch ("" + "\n\t") {
                                            case "foo bar":
                                                var v29 uint32 = v14
                                                _ = v29
                                            case "":
                                                var v30 string = "foo bar"
                                                _ = v30
                                            }
                                            switch (fi11 + 9223372036854775807) {
                                            case 0:
                                                var v31 uint64 = 4294967295
                                                _ = v31
                                            default:
                                                if (-38 >= 0) {
                                                    v6 = v27
                                                    v28 += v4
                                                } else {
                                                    if (!((float64(((1e10 / 1.5) - ((0.0 + -2.25) / 1.0))) * 3.14159) < 0.0)) {
                                                        var v32 int64 = (^(v4))
                                                        switch "" {
                                                        case "hello":
                                                            for fi33 := 0; fi33 < 5; fi33++ {
                                                                var v34 bool = true
                                                                _ = fi33
                                                                _ = v34
                                                            }
                                                        case "foo bar":
                                                            v4 = int64(v6)
                                                        case "\n\t":
                                                            switch (fi11 << (((^(v27))) & 63)) {
                                                            case 235:
                                                                var v35 uint8 = 255
                                                                _ = v35
                                                            }
                                                        default:
                                                            if (-1 <= -1) {
                                                                v36 := GenS2{F0: "\n\t", F1: v6}
                                                                _ = v36
                                                            } else {
                                                                v37 := GenS1{F0: (false && false), F1: v32, F2: 0}
                                                                v6 = v27
                                                                _ = v37
                                                            }
                                                        }
                                                        fmt.Println(v28, v14, v27)
                                                        _ = v32
                                                    }
                                                    var v38 string = (("foo bar" + "a") + "hello")
                                                    _ = v38
                                                }
                                            }
                                        } else {
                                            v4 += 0
                                            v3 += int64(v14)
                                        }
                                        fmt.Println(v3)
                                        _ = v28
                                    }
                                    var v39 uint8 = 255
                                    _ = v27
                                    _ = v39
                                }
                                var v40 uint = v6
                                fmt.Println(v5, fi11)
                                _ = v40
                            }
                            _ = v16
                        } else {
                            var v41 uint64 = v5
                            var v42 string = "hello"
                            if len(v1) > 0 {
                                v1[0] = int(v2)
                            }
                            _ = v41
                            _ = v42
                        }
                    }
                    var v43 uint64 = (v5 & v5)
                    _ = v14
                    _ = v43
                }
                fmt.Println(v4, v3)
                _ = fi11
            }
            _ = v10
        }
    }
    var v44 float32 = 3.14159
    _ = len(v1)
    fmt.Println(v2, v2)
    if (3.14159 < 1e10) {
        if len(v1) > 0 {
            v1[0] = 42
        }
        if true {
            var v45 int8 = int8(v5)
            v2 += v3
            var v46 uint64 = (v5 >> (((^(v6))) & 63))
            _ = v45
            _ = v46
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
    _ = v7
}
