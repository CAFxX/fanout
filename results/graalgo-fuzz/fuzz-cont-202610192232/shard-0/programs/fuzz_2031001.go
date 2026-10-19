package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 bool }
type GenS2 struct { F0 string }
func genFn0() int64 {
    v6 := GenS1{F0: 704, F1: 1, F2: true}
    if ((1e10 * ((1e10 + 1e10) + 3.14159)) >= -2.25) {
        var v7 uint = 42
        _ = v7
    }
    _ = v6
    return -653
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = 32767
    var v3 int = -1
    v4 := []string{("\n\t" + "a"), ((("foo bar" + "") + "hello") + ("a" + "")), "\n\t"}
    var v5 bool = true
    fmt.Println(v2)
    switch (("" + "a") + "hello") {
    case "hello":
        for fi8 := 0; fi8 < 6; fi8++ {
            for fi9 := 0; fi9 < 5; fi9++ {
                if (v5 && (1 <= -27)) {
                    delete(v1, (fi8 | (fi8 & 1)))
                } else {
                    if len(v4) > 0 {
                        v4[0] = "foo bar"
                    }
                    delete(v1, -386)
                }
                var v10 float32 = 1e10
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    case "foo bar":
        var v11 uint32 = 65535
        _ = v11
    default:
        if (v5 || false) {
            v12 := v1[(^(v3))]
            var v13 uint64 = 751
            for fi14 := 0; fi14 < 11; fi14++ {
                var v15 bool = (!v5)
                if v15 {
                    for fi16 := 0; fi16 < 17; fi16++ {
                        if v15 {
                            var v17 bool = (!true)
                            v18 := GenS1{F0: v12, F1: 868, F2: v15}
                            fmt.Println(v13, fi14, fi14)
                            _ = v17
                            _ = v18
                        }
                        if v5 {
                            var v19 uint16 = 203
                            _ = v19
                        } else {
                            v13 = (v13 & v13)
                            _ = len(v4)
                            var v20 int64 = 199
                            _ = v20
                        }
                        v21 := GenS2{F0: "a"}
                        _ = fi16
                        _ = v21
                    }
                } else {
                    switch ("" + "hello") {
                    case "hello":
                        var v22 uint16 = 0
                        _ = v22
                    case "\n\t":
                        v23 := GenS2{F0: "foo bar"}
                        _ = v23
                    default:
                        switch "a" {
                        case "\n\t":
                            fmt.Println(v15)
                        case "foo bar":
                            var v24 uint64 = v13
                            _ = v24
                        }
                    }
                    v25 := GenS2{F0: ""}
                    _ = v25
                }
                _ = fi14
                _ = v15
            }
            _ = v12
            _ = v13
        } else {
            v2 = v2
        }
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        v2 += -32768
        _ = fi26
    }
    v27 := GenS2{F0: "foo bar"}
    fmt.Println(v3, v3, v3)
    var v28 int64 = 941
    v2 = v2
    v2 = -94
    switch 127 {
    case -494:
        switch (v3 - (v3 << ((0) & 63))) {
        case 9223372036854775807:
            var v29 uint16 = 1
            _ = v29
        }
    case 32768:
        var v30 float64 = 0.0
        _ = v30
    }
    fmt.Println(v28, v28)
    if v5 {
        if (2147483648 == 496) {
            v2 = v2
        }
        if v5 {
            var v31 uint8 = 42
            v28 = (v28 >> ((4294967295) & 63))
            _ = v31
        } else {
            v3 += v3
            fmt.Println(v2, v28, v28)
        }
        v32 := GenS2{F0: "hello"}
        _ = v32
    } else {
        for fi33 := 0; fi33 < 9; fi33++ {
            var v34 uint16 = 1
            v2 = (v2 * (v2 << ((946) & 63)))
            _ = fi33
            _ = v34
        }
    }
    if v5 {
        switch ("hello" + ("" + "hello")) {
        case "hello":
            if v5 {
                var v35 float64 = 0.0
                var v36 int32 = 2147483647
                _ = v35
                _ = v36
            }
        case "a":
            v37 := GenS1{F0: 0, F1: -2147483648, F2: (0 < 1)}
            _ = v37
        case "\n\t":
            v5 = false
        default:
            var v38 uint = 1
            _ = v38
        }
        fmt.Println(v5)
        v28 += (v28 << ((701) & 63))
    }
    fmt.Println(v5, v28, v5)
    if true {
        if (!false) {
            v5 = v5
        }
    } else {
        v5 = true
        var v39 uint16 = 127
        _ = v39
    }
    fmt.Println(v3, v5, v3)
    var v40 string = "hello"
    switch v40 {
    case "\n\t":
        var v41 uint8 = 128
        _ = v41
    case "hello":
        var v42 int32 = 42
        _ = v42
    default:
        v3 = v3
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
