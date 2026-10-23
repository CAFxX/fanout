package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32; F1 uint16 }

func main() {
    var v1 uint32 = 42
    var v2 float32 = 1.5
    v3 := []uint32{2147483648, v1, (v1 * (v1 - v1)), v1}
    var v4 uint32 = (v1 / ((v1) | 1))
    for fi5 := 0; fi5 < 15; fi5++ {
        v1 = v4
        v6 := GenS1{F0: 533}
        _ = fi5
        _ = v6
    }
    var v7 int = 9223372036854775807
    v2 = (1e10 - float32(1.5))
    var v8 uint64 = 42
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v9 uint = 2
    for fi10 := 0; fi10 < 1; fi10++ {
        switch (fi10 >> ((v9) & 63)) {
        case 9223372036854775807:
            var v11 int64 = -1
            _ = v11
        }
        _ = fi10
    }
    var v12 string = (((("\n\t" + ("a" + (("a" + "\n\t") + ("hello" + "a")))) + "hello") + ("hello" + (("a" + ("" + (((((("" + "a") + "hello") + "\n\t") + ((("foo bar" + "") + ("a" + "a")) + "foo bar")) + "hello") + "foo bar"))) + ("" + "foo bar")))) + "")
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 float64 = (3.14159 * 3.14159)
        for fi15 := 0; fi15 < 7; fi15++ {
            if false {
                switch 612 {
                case 1:
                    var v16 int = (v7 / ((fi15) | 1))
                    _ = v16
                case 728:
                    v2 -= v2
                case -754:
                    var v17 uint16 = 65535
                    _ = v17
                default:
                    var v18 bool = ((false && (true && (4294967295 >= v4))) && (false && (-1 >= -829)))
                    _ = v18
                }
            }
            if (-9223372036854775808 >= 77) {
                var v19 bool = true
                var v20 bool = false
                v12 = ((v12 + "") + v12)
                if v19 {
                    v3 = append(v3, (v4 & 65535))
                    switch v12 {
                    case "a":
                        var v21 bool = (!v20)
                        switch 0 {
                        case -450:
                            v22 := GenS2{F0: 2147483647}
                            fmt.Println(v20, v4)
                            if v19 {
                                fmt.Println(v19)
                                var v23 string = ""
                                var v24 uint8 = uint8(v8)
                                _ = v23
                                _ = v24
                            }
                            _ = v22
                        case 659:
                            if (42 >= 42) {
                                v25 := GenS3{F0: (^(v1)), F1: "a"}
                                var v26 int16 = -141
                                if (v19 || v21) {
                                    v27 := GenS4{F0: -477}
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v28 float64 = 1.5
                                    _ = v27
                                    _ = v28
                                }
                                _ = v25
                                _ = v26
                            }
                        default:
                            v8 = 1
                        }
                        var v29 float64 = v14
                        _ = v21
                        _ = v29
                    default:
                        var v30 int64 = 42
                        _ = v30
                    }
                }
                _ = v19
                _ = v20
            } else {
                v1 = 740
                v31 := GenS5{F0: 0, F1: 65535}
                _ = v31
            }
            _ = fi15
        }
        _ = fi13
        _ = v14
    }
    v32 := GenS1{F0: 4294967295}
    for fi33 := 0; fi33 < 18; fi33++ {
        for fi34 := 0; fi34 < 16; fi34++ {
            v12 = ((v12 + "foo bar") + v12)
            switch (((v12 + "foo bar") + "\n\t") + ((v12 + "a") + v12)) {
            case "hello":
                v35 := GenS5{F0: 4294967295, F1: 482}
                _ = v35
            case "\n\t":
                var v36 bool = (711 < 1)
                _ = v36
            case "foo bar":
                v12 = v12
            }
            v9 = (v9 ^ uint(v8))
            _ = fi34
        }
        v12 = (("foo bar" + "\n\t") + v12)
        for fi37 := 0; fi37 < 4; fi37++ {
            var v38 int32 = 1
            fmt.Println(v2, fi33, fi33)
            _ = fi37
            _ = v38
        }
        _ = fi33
    }
    var v39 uint16 = 2
    if true {
        fmt.Println(v1)
    } else {
        for fi40 := 0; fi40 < 16; fi40++ {
            if len(v3) > 0 {
                v3[0] = 542
            }
            _ = fi40
        }
        fmt.Println(v39, v4, v9)
    }
    for fi41 := 0; fi41 < 4; fi41++ {
        fmt.Println(v1, v7)
        _ = fi41
    }
    v4 = uint32(v9)
    for fi42 := 0; fi42 < 18; fi42++ {
        v7 = (fi42 + (fi42 | v7))
        _ = fi42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v32
    _ = v39
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
