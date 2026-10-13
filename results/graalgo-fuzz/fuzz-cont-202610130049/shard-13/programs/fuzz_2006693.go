package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = 0
    v3 := make(map[string]bool)
    var v4 uint8 = 42
    v5 := GenS1{F0: "hello", F1: ("\n\t" + "a")}
    switch "\n\t" {
    case "\n\t":
        v6 := GenS1{F0: ("a" + "hello"), F1: "\n\t"}
        _ = v6
    }
    v7 := []int{-1, 9223372036854775807}
    v4 -= v4
    v4 += (v4 ^ (^(v4)))
    switch "hello" {
    case "foo bar":
        if false {
            v2 -= v2
            for fi8 := 0; fi8 < 7; fi8++ {
                switch "hello" {
                case "\n\t":
                    if true {
                        for fi9 := 0; fi9 < 17; fi9++ {
                            fmt.Println(fi8, v2)
                            if len(v7) > 0 {
                                v7[0] = fi9
                            }
                            _ = fi9
                        }
                        var v10 uint8 = 0
                        var v11 uint64 = 18446744073709551615
                        _ = v10
                        _ = v11
                    } else {
                        v12 := GenS1{F0: "a", F1: "\n\t"}
                        _ = v12
                    }
                }
                for fi13 := 0; fi13 < 15; fi13++ {
                    fmt.Println(fi13)
                    _ = fi13
                }
                _ = fi8
            }
        } else {
            v3[""] = (18446744073709551615 >= 65535)
            var v14 bool = ((860 != 344) || (42 <= -388))
            if (2 != 42) {
                for fi15 := 0; fi15 < 6; fi15++ {
                    v14 = (!v14)
                    _ = fi15
                }
                if false {
                    for fi16 := 0; fi16 < 18; fi16++ {
                        fmt.Println(v4)
                        _ = fi16
                    }
                    var v17 uint32 = 581
                    for fi18 := 0; fi18 < 5; fi18++ {
                        fmt.Println(v14)
                        var v19 uint8 = v4
                        if ("" <= "hello") {
                            switch (("foo bar" + ("" + ("a" + "foo bar"))) + "foo bar") {
                            case "":
                                v20 := GenS2{F0: "hello"}
                                _ = v20
                            }
                            for fi21 := 0; fi21 < 19; fi21++ {
                                if (!(-9223372036854775808 != 0)) {
                                    switch ("" + "") {
                                    case "hello":
                                        fmt.Println(v14, v14)
                                    case "a":
                                        _ = len(v3)
                                    }
                                    fmt.Println(v14)
                                }
                                var v22 int8 = 100
                                _ = fi21
                                _ = v22
                            }
                        }
                        _ = fi18
                        _ = v19
                    }
                    _ = v17
                } else {
                    var v23 int16 = 522
                    var v24 uint = 255
                    v25 := v3["hello"]
                    _ = v23
                    _ = v24
                    _ = v25
                }
            } else {
                v2 = (v2 >> ((1) & 63))
                v14 = ((!(0.0 == (float32(3.14159) + (-2.25 * 3.14159)))) || (3.14159 <= (float32(1.5) + (1e10 + ((3.14159 / (0.1 * (1e10 - (1e10 / -2.25)))) / -2.25)))))
                v26 := GenS3{F0: 901}
                _ = v26
            }
            _ = v14
        }
    default:
        v4 = v4
    }
    for fi27 := 0; fi27 < 14; fi27++ {
        if true {
            for fi28 := 0; fi28 < 6; fi28++ {
                var v29 int32 = 0
                v29 -= 0
                switch (fi28 >> ((18446744073709551615) & 63)) {
                case 42:
                    var v30 uint8 = v4
                    _ = v30
                case 32768:
                    var v31 float64 = float64(-2.25)
                    _ = v31
                case -914:
                    var v32 int16 = -11
                    var v33 int16 = (v32 & v32)
                    var v34 int16 = v33
                    _ = v32
                    _ = v33
                    _ = v34
                default:
                    fmt.Println(fi27, fi27)
                }
                _ = fi28
                _ = v29
            }
        } else {
            v2 += (v2 << ((424) & 63))
        }
        v4 += (v4 + (v4 ^ 2))
        if true {
            v4 -= 42
        }
        _ = fi27
    }
    fmt.Println(v4)
    if true {
        v2 = (v2 >> ((65535) & 63))
    }
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
