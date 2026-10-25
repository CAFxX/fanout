package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 string; F2 int64 }

func main() {
    var v1 int32 = 1
    var v2 bool = ((-1 == (v1 * (-(v1)))) && false)
    var v3 int32 = (v1 & v1)
    var v4 uint8 = 42
    var v5 float32 = ((1.5 + float32(0.0)) / -2.25)
    var v6 int32 = (v3 * v3)
    if (42 != 9223372036854775807) {
        var v7 uint = 923
        _ = v7
    } else {
        if (((v5 * -2.25) - v5) < (v5 / v5)) {
            var v8 uint16 = 65535
            switch 1 {
            case -2147483648:
                v5 = 0.1
            case 127:
                var v9 int16 = 42
                _ = v9
            case 2147483647:
                for fi10 := 0; fi10 < 10; fi10++ {
                    v4 -= v4
                    v11 := GenS1{F0: (v3 | v1)}
                    _ = fi10
                    _ = v11
                }
            }
            _ = v8
        } else {
            v12 := GenS1{F0: v1}
            if false {
                v13 := []int{16, 46, -9223372036854775808}
                for fi14 := 0; fi14 < 10; fi14++ {
                    for fi15 := 0; fi15 < 4; fi15++ {
                        v16 := GenS2{F0: "a"}
                        var v17 float32 = v5
                        if (!(v2 && (2147483648 < 750))) {
                            var v18 uint32 = 259
                            v13 = append(v13, (fi15 * int(v4)))
                            switch "foo bar" {
                            case "hello":
                                fmt.Println(v1, fi14)
                            case "\n\t":
                                _ = len(v13)
                            }
                            _ = v18
                        }
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    _ = fi14
                }
                switch ("hello" + "hello") {
                case "foo bar":
                    v2 = (0.0 > -2.25)
                case "a":
                    switch "a" {
                    case "hello":
                        switch "foo bar" {
                        case "\n\t":
                            var v19 int16 = 42
                            _ = v19
                        }
                    case "foo bar":
                        v20 := GenS1{F0: (v3 + (v3 - v1))}
                        _ = v20
                    }
                }
                _ = v13
            } else {
                v21 := GenS1{F0: (v1 ^ v6)}
                _ = v21
            }
            v22 := GenS2{F0: ("" + ((("hello" + "foo bar") + "a") + ""))}
            _ = v12
            _ = v22
        }
    }
    var v23 uint16 = 470
    v24 := []int{918, -1}
    var v25 bool = v2
    fmt.Println(v25, v2, v4)
    var v26 uint = 238
    switch ((("" + ("foo bar" + "")) + "\n\t") + "hello") {
    case "hello":
        v2 = v25
    case "a":
        var v27 int8 = 1
        _ = v27
    default:
        var v28 uint32 = 652
        _ = v28
    }
    if false {
        var v29 uint8 = (v4 * (v4 - (^(v4))))
        v6 -= v3
        _ = v29
    } else {
        if len(v24) >= 2 {
            v24 = v24[1:len(v24)]
        }
        var v30 int64 = 228
        if v25 {
            var v31 int = 665
            var v32 uint32 = 756
            _ = v31
            _ = v32
        } else {
            if (!v2) {
                fmt.Println(v25)
            } else {
                v6 -= (v1 / ((v1) | 1))
                var v33 int64 = 42
                if false {
                    v26 = v26
                    v30 -= (v30 << (((v26 >> (((v26 + (v26 * (v26 + v26)))) & 63))) & 63))
                } else {
                    for fi34 := 0; fi34 < 6; fi34++ {
                        if (536 == 42) {
                            v6 += v6
                            v35 := GenS3{F0: -32768, F1: "foo bar", F2: v33}
                            _ = v35
                        }
                        switch "foo bar" {
                        case "a":
                            var v36 string = ""
                            _ = v36
                        }
                        _ = fi34
                    }
                    var v37 float32 = v5
                    _ = v37
                }
                _ = v33
            }
        }
        _ = v30
    }
    fmt.Println(v2)
    v23 = v23
    v4 += 42
    v38 := GenS2{F0: ""}
    switch 9223372036854775807 {
    case 32767:
        if (!v2) {
            fmt.Println(v2, v2, v5)
        } else {
            fmt.Println(v3, v3, v26)
        }
    }
    var v39 float64 = (1.5 + 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
