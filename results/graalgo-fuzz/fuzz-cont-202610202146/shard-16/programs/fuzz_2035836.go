package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int; F1 bool; F2 string }

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = 1e10
    var v3 int16 = 1000
    var v4 int = 0
    var v5 int16 = v3
    var v6 uint64 = 0
    switch "" {
    case "":
        v7 := GenS1{F0: "foo bar"}
        _ = v7
    case "hello":
        v6 -= 104
    case "a":
        fmt.Println(v6, v3)
    default:
        switch ("\n\t" + "\n\t") {
        case "a":
            if false {
                v8 := GenS1{F0: ("foo bar" + ("" + "a"))}
                _ = v8
            } else {
                v1["foo bar"] = false
                if true {
                    switch (("" + "foo bar") + "") {
                    case "foo bar":
                        v2 -= 1e10
                    case "":
                        if (!(255 >= 128)) {
                            fmt.Println(v2, v6, v5)
                            v6 += (v6 / (((v6 + (v6 * (v6 ^ v6)))) | 1))
                        }
                    default:
                        if false {
                            if true {
                                if (!false) {
                                    switch "foo bar" {
                                    case "\n\t":
                                        v9 := GenS1{F0: ("hello" + "foo bar")}
                                        _ = v9
                                    default:
                                        if (!((409 <= 2147483648) || ((v4 * (v4 ^ 9223372036854775807)) != (v4 | v4)))) {
                                            v10 := GenS1{F0: "\n\t"}
                                            var v11 int32 = -1
                                            _ = v10
                                            _ = v11
                                        }
                                    }
                                } else {
                                    v12 := GenS1{F0: "a"}
                                    _ = v12
                                }
                            } else {
                                for fi13 := 0; fi13 < 15; fi13++ {
                                    var v14 string = ""
                                    _ = fi13
                                    _ = v14
                                }
                            }
                            v4 = 154
                            for fi15 := 0; fi15 < 11; fi15++ {
                                var v16 float64 = v2
                                v17 := GenS1{F0: ("a" + "\n\t")}
                                delete(v1, ("foo bar" + ("hello" + "hello")))
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                        }
                    }
                    v6 -= v6
                } else {
                    fmt.Println(v5, v6)
                    var v18 int = 1
                    _ = v18
                }
            }
        }
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        v20 := []int{}
        var v21 int8 = 72
        _ = fi19
        _ = v20
        _ = v21
    }
    switch (v4 << ((509) & 63)) {
    case 2147483647:
        var v22 float32 = 1e10
        _ = v22
    }
    if false {
        v23 := GenS2{F0: -9223372036854775808, F1: -430}
        _ = v23
    }
    fmt.Println(v4, v4)
    if true {
        v5 -= v5
        v24 := GenS3{F0: v4, F1: (772 < -256), F2: ""}
        var v25 uint8 = 255
        _ = v24
        _ = v25
    } else {
        var v26 uint32 = 4294967295
        _ = v26
    }
    switch "hello" {
    case "foo bar":
        v2 += 3.14159
    }
    var v27 int16 = -116
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
