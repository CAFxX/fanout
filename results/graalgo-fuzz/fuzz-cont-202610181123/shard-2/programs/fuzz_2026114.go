package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 string; F2 uint16 }

func main() {
    v1 := []uint32{345, 472, 259, 2147483647}
    var v2 int = -913
    var v3 int = -432
    v4 := []string{("" + "hello")}
    var v5 int8 = 1
    v6 := GenS1{F0: 1, F1: (v5 * v5), F2: true}
    var v7 int32 = 572
    if (!true) {
        fmt.Println(v2)
        var v8 float32 = 0.1
        _ = v8
    } else {
        v2 -= (v2 << ((42) & 63))
        fmt.Println(v7, v3)
    }
    v9 := GenS2{F0: 0}
    switch 0 {
    case -9223372036854775808:
        var v10 float64 = 0.0
        _ = v10
    }
    fmt.Println(v3, v2)
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 uint32 = 2147483648
        _ = fi11
        _ = v12
    }
    v3 = 798
    var v13 int = 839
    if true {
        var v14 int64 = 1
        _ = v14
    } else {
        for fi15 := 0; fi15 < 3; fi15++ {
            fmt.Println(v2, v3)
            if true {
                v2 = v3
                var v16 uint = 1
                _ = v16
            } else {
                for fi17 := 0; fi17 < 18; fi17++ {
                    v5 += (v5 | v5)
                    _ = fi17
                }
                var v18 int8 = 34
                _ = v18
            }
            fmt.Println(v13, v3)
            _ = fi15
        }
    }
    var v19 int = v13
    if true {
        if true {
            if ((-348 == 1000) && true) {
                v2 -= 32768
                if len(v4) > 0 {
                    v4[0] = "hello"
                }
                v20 := GenS3{F0: 1000, F1: (("a" + "a") + ("hello" + ("foo bar" + "\n\t"))), F2: 42}
                _ = v20
            } else {
                if (!true) {
                    if true {
                        fmt.Println(v7, v7, v19)
                        var v21 int16 = -247
                        switch (v2 ^ (v19 >> ((586) & 63))) {
                        case -9223372036854775808:
                            v22 := GenS3{F0: (v21 >> ((2) & 63)), F1: "\n\t", F2: 480}
                            _ = v22
                        default:
                            if true {
                                var v23 uint16 = 0
                                _ = v23
                            } else {
                                switch "a" {
                                case "hello":
                                    v19 = 42
                                case "foo bar":
                                    _ = len(v4)
                                case "\n\t":
                                    for fi24 := 0; fi24 < 19; fi24++ {
                                        v25 := GenS3{F0: (v21 << ((2) & 63)), F1: ("\n\t" + (("a" + "foo bar") + "")), F2: 716}
                                        _ = fi24
                                        _ = v25
                                    }
                                }
                                var v26 float32 = -2.25
                                _ = v26
                            }
                        }
                        _ = v21
                    } else {
                        v2 = v19
                        var v27 int8 = (v5 ^ (v5 & (v5 >> ((238) & 63))))
                        var v28 int = (^(v13))
                        _ = v27
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 11; fi29++ {
                        if true {
                            v4 = append(v4, ("\n\t" + "hello"))
                        } else {
                            var v30 int = v19
                            _ = v30
                        }
                        var v31 bool = false
                        _ = fi29
                        _ = v31
                    }
                    if false {
                        var v32 string = "a"
                        var v33 uint64 = 602
                        var v34 bool = true
                        _ = v32
                        _ = v33
                        _ = v34
                    } else {
                        for fi35 := 0; fi35 < 19; fi35++ {
                            v2 = v13
                            _ = fi35
                        }
                    }
                } else {
                    v36 := GenS1{F0: 219, F1: (v5 ^ (v5 / ((v5) | 1))), F2: ((v7 <= v7) && ("a" > ("hello" + "foo bar")))}
                    if (-2.25 > (0.0 - 1e10)) {
                        switch (-(v2)) {
                        case -987:
                            var v37 float32 = -2.25
                            _ = v37
                        case -1:
                            switch "foo bar" {
                            case "hello":
                                fmt.Println(v7, v19, v7)
                            case "a":
                                var v38 int32 = 533
                                _ = v38
                            case "foo bar":
                                var v39 int16 = 1
                                _ = v39
                            }
                        default:
                            var v40 int64 = -724
                            _ = v40
                        }
                        var v41 int8 = v5
                        var v42 float64 = (1.5 - -2.25)
                        _ = v41
                        _ = v42
                    }
                    _ = v36
                }
            }
            v13 -= (v2 + (v13 - 243))
        } else {
            v5 = -104
        }
    }
    if ((v7 << ((2) & 63)) >= v7) {
        v13 = (v3 * -971)
    }
    if len(v4) > 0 {
        v4[0] = ""
    }
    var v43 uint16 = 32767
    v19 = v2
    fmt.Println(v5, v43)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
