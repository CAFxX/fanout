package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 int32; F1 bool }

func main() {
    var v1 int32 = 42
    var v2 int8 = 100
    v3 := []uint32{485}
    var v4 int64 = -1
    var v5 float64 = float64(0.0)
    v6 := make(map[string]bool)
    var v7 float64 = (v5 * (v5 / v5))
    if (!true) {
        fmt.Println(v7, v2, v2)
        fmt.Println(v2, v5, v5)
        v8 := GenS1{F0: 172, F1: (v1 % ((-467) | 1)), F2: 42}
        _ = v8
    } else {
        _ = len(v6)
        if (v4 < -997) {
            _ = len(v6)
            var v9 string = ("hello" + "")
            if (v4 <= (v4 - v4)) {
                v10 := v6[(((((v9 + v9) + v9) + "a") + "foo bar") + v9)]
                _ = v10
            }
            _ = v9
        } else {
            switch (("foo bar" + ("foo bar" + "a")) + "a") {
            case "a":
                if true {
                    v2 += (v2 >> ((88) & 63))
                } else {
                    v11 := GenS1{F0: 99, F1: (v1 ^ (v1 ^ (v1 >> ((264) & 63)))), F2: 65535}
                    _ = v11
                }
            case "":
                if (-1 == -1) {
                    v12 := GenS2{F0: (-(v2))}
                    _ = v12
                }
            default:
                if (false && false) {
                    v13 := GenS2{F0: v2}
                    fmt.Println(v5, v7, v4)
                    _ = v13
                } else {
                    if false {
                        v14 := GenS2{F0: v2}
                        var v15 uint32 = 65535
                        v16 := GenS1{F0: 127, F1: -1, F2: v15}
                        _ = v14
                        _ = v15
                        _ = v16
                    } else {
                        v6["\n\t"] = (!(!false))
                    }
                }
            }
            switch "hello" {
            case "foo bar":
                v4 = v4
            }
            v17 := v6["hello"]
            _ = v17
        }
    }
    var v18 bool = (!(true && false))
    var v19 uint64 = 42
    var v20 float32 = -2.25
    var v21 int64 = (v4 & 9223372036854775807)
    v2 += v2
    var v22 int64 = (v4 / (((v21 >> ((2) & 63))) | 1))
    switch "hello" {
    case "hello":
        var v23 int64 = v21
        _ = v23
    case "":
        v24 := GenS1{F0: 255, F1: v1, F2: 2}
        _ = v24
    }
    v25 := GenS1{F0: 1, F1: 1, F2: 167}
    _ = len(v3)
    v19 += v19
    if (!v18) {
        _ = len(v3)
        var v26 int = 711
        for fi27 := 0; fi27 < 14; fi27++ {
            v18 = ((("" + "a") + "foo bar") >= "foo bar")
            var v28 int32 = 0
            fmt.Println(v21)
            _ = fi27
            _ = v28
        }
        _ = v26
    } else {
        var v29 float64 = (v7 - v7)
        _ = v29
    }
    if v18 {
        var v30 int32 = v1
        _ = v30
    } else {
        var v31 uint8 = 71
        v4 += v21
        switch -463 {
        case 219:
            v3 = append(v3, 2)
        case -2147483648:
            v32 := v6["a"]
            _ = v32
        default:
            v33 := GenS3{F0: -1, F1: 0}
            _ = v33
        }
        _ = v31
    }
    switch "" {
    case "":
        var v34 int16 = 0
        _ = v34
    case "a":
        var v35 int64 = v4
        _ = v35
    }
    var v36 uint64 = (v19 ^ (v19 >> ((33) & 63)))
    if false {
        switch ("\n\t" + "\n\t") {
        case "hello":
            if ((v2 ^ (v2 / ((v2) | 1))) < v2) {
                v36 += (^(v19))
                for fi37 := 0; fi37 < 18; fi37++ {
                    v20 += (1e10 - v20)
                    _ = fi37
                }
            } else {
                for fi38 := 0; fi38 < 12; fi38++ {
                    var v39 uint = 65535
                    _ = fi38
                    _ = v39
                }
                if (32767 > 255) {
                    if v18 {
                        v2 += (^(v2))
                    }
                    _ = len(v3)
                    v40 := GenS2{F0: v2}
                    _ = v40
                }
            }
        case "\n\t":
            var v41 uint32 = 4294967295
            _ = v41
        default:
            for fi42 := 0; fi42 < 5; fi42++ {
                fmt.Println(v36, v36)
                var v43 int8 = v2
                if false {
                    if true {
                        v43 += (v2 << ((2) & 63))
                        var v44 int32 = 581
                        for fi45 := 0; fi45 < 10; fi45++ {
                            v46 := GenS4{F0: v44, F1: true}
                            _ = fi45
                            _ = v46
                        }
                        _ = v44
                    } else {
                        fmt.Println(v36)
                        v47 := GenS4{F0: v1, F1: v18}
                        _ = v47
                    }
                    v1 = (v1 ^ v1)
                }
                _ = fi42
                _ = v43
            }
        }
        v19 = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
