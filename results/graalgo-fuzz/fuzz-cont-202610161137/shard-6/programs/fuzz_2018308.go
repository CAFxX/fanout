package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 uint8 }

func main() {
    var v1 uint8 = 42
    var v2 float64 = 1.5
    var v3 int8 = 19
    var v4 int64 = 0
    var v5 uint = 65535
    v6 := GenS1{F0: "a"}
    switch 128 {
    case 9223372036854775807:
        switch "a" {
        case "foo bar":
            v1 = 2
        default:
            if ("" != ("\n\t" + "a")) {
                switch 32768 {
                case 32768:
                    var v7 uint = 18446744073709551615
                    _ = v7
                }
                for fi8 := 0; fi8 < 19; fi8++ {
                    var v9 uint32 = 180
                    if (1 == 591) {
                        if true {
                            v4 += v4
                            v10 := []int{(fi8 * 903)}
                            v3 += v3
                            _ = v10
                        }
                        v5 -= v5
                        var v11 int64 = (v4 | v4)
                        _ = v11
                    } else {
                        var v12 bool = false
                        _ = v12
                    }
                    _ = fi8
                    _ = v9
                }
            } else {
                v4 = v4
            }
        }
    case 843:
        v13 := []int{-9223372036854775808, -1, 42}
        _ = v13
    }
    switch "hello" {
    case "":
        for fi14 := 0; fi14 < 15; fi14++ {
            v15 := []int{fi14, -525, fi14, int(v4)}
            for fi16 := 0; fi16 < 20; fi16++ {
                v4 = v4
                _ = fi16
            }
            fmt.Println(v1, v2)
            _ = fi14
            _ = v15
        }
    case "hello":
        fmt.Println(v4)
    case "foo bar":
        if true {
            v17 := []int{1, 42, 128, 32767}
            fmt.Println(v5, v2)
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        v19 := []int{(fi18 | -878), -9223372036854775808, fi18, fi18}
        for fi20 := 0; fi20 < 15; fi20++ {
            v21 := GenS1{F0: ""}
            _ = fi20
            _ = v21
        }
        v22 := GenS1{F0: ("hello" + ("foo bar" + ""))}
        _ = fi18
        _ = v19
        _ = v22
    }
    v23 := []int{433, 839, -1, 32768}
    fmt.Println(v2, v5)
    fmt.Println(v5, v2)
    v24 := GenS1{F0: "foo bar"}
    switch (((("\n\t" + "hello") + "hello") + "\n\t") + "a") {
    case "\n\t":
        if false {
            v25 := GenS1{F0: "foo bar"}
            _ = v25
        }
    case "a":
        v26 := GenS1{F0: ("\n\t" + (("foo bar" + "\n\t") + "a"))}
        _ = v26
    default:
        var v27 int32 = 1
        _ = v27
    }
    v23 = append(v23, -892)
    var v28 int64 = (v4 & (-(v4)))
    for fi29 := 0; fi29 < 8; fi29++ {
        v30 := GenS2{F0: 65535, F1: (v1 >> ((v5) & 63))}
        fmt.Println(fi29, v4)
        _ = fi29
        _ = v30
    }
    var v31 int64 = -1
    if true {
        v32 := GenS1{F0: (((("\n\t" + "\n\t") + "hello") + "foo bar") + "foo bar")}
        _ = v32
    } else {
        for fi33 := 0; fi33 < 13; fi33++ {
            v5 -= v5
            _ = fi33
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
