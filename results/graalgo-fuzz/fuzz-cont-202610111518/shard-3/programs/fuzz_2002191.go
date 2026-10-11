package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 uint32; F2 int }
type GenS4 struct { F0 uint; F1 int16; F2 int16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32; F1 int }

func main() {
    var v1 int8 = -128
    var v2 uint = 554
    var v3 uint8 = 42
    var v4 uint = (v2 >> ((v2) & 63))
    v5 := []uint32{895, 2, 752}
    v6 := []uint32{65535, 1}
    var v7 int16 = 1000
    fmt.Println(v1, v2)
    v8 := GenS1{F0: v2}
    v9 := GenS1{F0: (v4 & v4)}
    if (235 > 128) {
        fmt.Println(v7, v2, v4)
        for fi10 := 0; fi10 < 13; fi10++ {
            for fi11 := 0; fi11 < 16; fi11++ {
                v5 = append(v5, 2147483647)
                v1 = -128
                v12 := GenS2{F0: (!false)}
                _ = fi11
                _ = v12
            }
            fmt.Println(v3, v7, v7)
            _ = fi10
        }
        var v13 uint64 = uint64(v2)
        _ = v13
    } else {
        fmt.Println(v3, v2, v7)
        if true {
            switch 32768 {
            case 9223372036854775807:
                var v14 string = (("hello" + (("a" + "hello") + "a")) + ("foo bar" + (("\n\t" + "hello") + (("" + (("\n\t" + "hello") + "\n\t")) + ("" + "\n\t")))))
                _ = v14
            }
            var v15 float64 = 1.5
            _ = v15
        } else {
            fmt.Println(v1, v3, v2)
        }
    }
    v7 = (v7 << ((v2) & 63))
    v3 -= v3
    v3 = 2
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 float32 = ((3.14159 - -2.25) / (((1.5 - (-2.25 * -2.25)) - (-2.25 * (1e10 * -2.25))) / (1.5 * 3.14159)))
        if ("hello" > "\n\t") {
            if (v2 == (v4 >> ((42) & 63))) {
                v18 := GenS3{F0: (fi16 >> ((255) & 63)), F1: 0, F2: 9223372036854775807}
                _ = v18
            } else {
                switch fi16 {
                case 2147483647:
                    v19 := GenS1{F0: v4}
                    _ = v19
                }
                switch "hello" {
                case "foo bar":
                    if true {
                        var v20 int = fi16
                        for fi21 := 0; fi21 < 15; fi21++ {
                            switch (v20 + v20) {
                            case -2147483648:
                                v17 = 0.1
                            default:
                                v17 -= v17
                            }
                            v4 += v4
                            _ = fi21
                        }
                        _ = v20
                    } else {
                        var v22 uint8 = v3
                        fmt.Println(v22)
                        _ = v22
                    }
                case "":
                    switch ("foo bar" + ("" + "a")) {
                    case "\n\t":
                        _ = len(v5)
                    case "foo bar":
                        v17 -= -2.25
                    default:
                        if len(v6) > 0 {
                            v6[0] = 42
                        }
                    }
                case "a":
                    if false {
                        for fi23 := 0; fi23 < 12; fi23++ {
                            var v24 string = ""
                            v7 = 1
                            var v25 int64 = 9223372036854775807
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        switch ("hello" + "hello") {
                        case "":
                            v5 = append(v5, 2147483648)
                        default:
                            v26 := GenS2{F0: (1.5 == (3.14159 + 0.1))}
                            _ = v26
                        }
                        var v27 bool = false
                        _ = v27
                    } else {
                        switch 127 {
                        case 1:
                            v28 := GenS1{F0: v2}
                            _ = v28
                        case 32768:
                            var v29 int = 933
                            _ = v29
                        case 0:
                            v30 := GenS4{F0: v4, F1: v7, F2: v7}
                            _ = v30
                        default:
                            var v31 int32 = -1
                            _ = v31
                        }
                    }
                }
            }
        } else {
            v3 += v3
            v3 = 128
        }
        _ = fi16
        _ = v17
    }
    v32 := GenS1{F0: v4}
    _ = len(v5)
    v2 = v4
    for fi33 := 0; fi33 < 20; fi33++ {
        if true {
            _ = len(v5)
        }
        if false {
            fmt.Println(v4)
            v3 = (v3 << ((v2) & 63))
        } else {
            v1 -= 42
            if true {
                if false {
                    var v34 int32 = 2147483647
                    _ = v34
                }
                switch "" {
                case "hello":
                    fmt.Println(v7)
                case "":
                    for fi35 := 0; fi35 < 10; fi35++ {
                        if true {
                            v3 = 128
                        }
                        v36 := GenS2{F0: true}
                        _ = fi35
                        _ = v36
                    }
                }
                fmt.Println(v3)
            } else {
                v4 -= 995
            }
            switch (fi33 * 9223372036854775807) {
            case -1:
                if false {
                    if false {
                        v37 := GenS5{F0: 2147483648}
                        v38 := GenS6{F0: 65535, F1: (fi33 * -9223372036854775808)}
                        _ = v37
                        _ = v38
                    } else {
                        var v39 uint16 = 255
                        _ = v39
                    }
                }
            }
        }
        if (42 <= 0) {
            var v40 int64 = 1
            _ = v40
        }
        _ = fi33
    }
    if false {
        var v41 uint = 1
        _ = v41
    }
    _ = len(v5)
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
