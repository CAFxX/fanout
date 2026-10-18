package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int64; F1 uint }

func main() {
    var v1 int64 = 1
    var v2 uint64 = 18446744073709551615
    v3 := []string{("hello" + ("\n\t" + ("a" + ("\n\t" + "\n\t")))), "foo bar", (("foo bar" + (("foo bar" + "a") + "foo bar")) + (("\n\t" + "foo bar") + "\n\t"))}
    switch 0 {
    case -20:
        fmt.Println(v1, v1)
    case 127:
        switch ("" + ("" + "hello")) {
        case "foo bar":
            fmt.Println(v1, v1)
        case "a":
            var v4 uint = 65535
            _ = v4
        case "\n\t":
            v5 := GenS1{F0: ""}
            _ = v5
        default:
            var v6 int32 = 0
            _ = v6
        }
    case 46:
        v7 := GenS1{F0: "a"}
        _ = v7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        if true {
            v9 := GenS1{F0: "a"}
            v2 -= (v2 * 0)
            _ = v9
        }
        _ = fi8
    }
    v2 -= (v2 >> ((8) & 63))
    if true {
        if (!false) {
            var v10 uint8 = 1
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        if true {
            v12 := GenS1{F0: "a"}
            var v13 int16 = 1000
            var v14 uint16 = 0
            _ = v12
            _ = v13
            _ = v14
        }
        v15 := GenS1{F0: "\n\t"}
        var v16 uint16 = 945
        _ = fi11
        _ = v15
        _ = v16
    }
    var v17 uint16 = 42
    var v18 float64 = float64((0.0 + (3.14159 + 1e10)))
    v17 += v17
    v3 = append(v3, "")
    var v19 uint16 = 0
    if (281 <= (v17 * (v19 - (v19 ^ v17)))) {
        fmt.Println(v18, v1)
    } else {
        switch -59 {
        case 32767:
            if false {
                if true {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                v3 = append(v3, "hello")
                if len(v3) > 0 {
                    v3[0] = "hello"
                }
            } else {
                if len(v3) > 0 {
                    v3[0] = "\n\t"
                }
                for fi20 := 0; fi20 < 12; fi20++ {
                    v21 := GenS1{F0: "hello"}
                    _ = fi20
                    _ = v21
                }
            }
        case -663:
            v17 = 747
        case 207:
            var v22 int8 = 127
            _ = v22
        default:
            var v23 uint8 = 21
            _ = v23
        }
        v24 := GenS1{F0: ""}
        _ = v24
    }
    if (!(!false)) {
        v25 := GenS2{F0: 510, F1: (v1 >> ((102) & 63))}
        fmt.Println(v17)
        _ = v25
    }
    var v26 uint16 = 255
    if len(v3) > 0 {
        v3[0] = ("a" + "foo bar")
    }
    switch 127 {
    case 995:
        switch "hello" {
        case "foo bar":
            v1 += (v1 << ((0) & 63))
        }
    case 42:
        v27 := GenS2{F0: 65535, F1: v1}
        _ = v27
    }
    var v28 int64 = -984
    var v29 uint8 = 2
    if (!((v2 + v2) == (v2 + v2))) {
        switch "hello" {
        case "":
            if (2 >= v19) {
                v30 := GenS1{F0: "\n\t"}
                var v31 uint32 = 42
                _ = v30
                _ = v31
            } else {
                v17 = 42
                v29 += 4
                fmt.Println(v18, v28, v18)
            }
        case "\n\t":
            fmt.Println(v2)
        case "hello":
            for fi32 := 0; fi32 < 14; fi32++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v28 = -1
                v33 := GenS3{F0: (v28 << ((42) & 63)), F1: 830}
                _ = fi32
                _ = v33
            }
        default:
            for fi34 := 0; fi34 < 2; fi34++ {
                v3 = append(v3, ("a" + ("foo bar" + "a")))
                _ = fi34
            }
        }
    } else {
        v18 += v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v28
    _ = v29
    _ = v3
}
