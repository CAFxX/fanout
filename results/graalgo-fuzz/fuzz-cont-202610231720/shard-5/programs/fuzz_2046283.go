package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string }
func genFn0(p0 int32, p1 uint16, p2 int32) uint8 {
    var v4 uint16 = 2
    switch "hello" {
    case "a":
        v5 := GenS1{F0: 1}
        _ = v5
    case "foo bar":
        p0 = p0
    }
    switch -605 {
    case 555:
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 255
}
func genFn1(p0 int8, p1 uint16) uint64 {
    v6 := GenS1{F0: 0}
    fmt.Println(p1, p0)
    var v7 uint64 = 18446744073709551615
    if true {
        var v8 int8 = (-(p0))
        switch ("a" + "\n\t") {
        case "hello":
            p1 -= uint16(v7)
        }
        var v9 string = "\n\t"
        _ = v8
        _ = v9
    } else {
        var v10 int64 = int64(p0)
        _ = v10
    }
    var v11 float64 = 1.5
    _ = p0
    _ = p1
    _ = v11
    _ = v6
    _ = v7
    return v7
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := []string{"a", "hello", "hello", (((("hello" + "hello") + ("" + ("hello" + "foo bar"))) + "hello") + "foo bar")}
    var v3 int8 = -1
    _ = len(v2)
    switch "foo bar" {
    case "hello":
        var v12 uint8 = 86
        _ = v12
    }
    fmt.Println(v1)
    v13 := GenS1{F0: 1}
    v3 += -112
    v3 = (v3 << (((^(v1))) & 63))
    v14 := GenS1{F0: 2147483647}
    v1 -= 255
    var v15 int64 = 0
    fmt.Println(v15, v1)
    var v16 bool = (!(32767 >= 32768))
    v1 = (v1 | 2)
    if (42 <= 0) {
        v17 := GenS2{F0: -8}
        v1 -= 1
        fmt.Println(v3, v15)
        _ = v17
    } else {
        if ((1 <= 65535) || (0 != (v3 - v3))) {
            v3 = (v3 % ((v3) | 1))
        }
        switch "foo bar" {
        case "hello":
            for fi18 := 0; fi18 < 6; fi18++ {
                for fi19 := 0; fi19 < 10; fi19++ {
                    var v20 uint32 = 2147483648
                    fmt.Println(v15, v20)
                    _ = fi19
                    _ = v20
                }
                for fi21 := 0; fi21 < 9; fi21++ {
                    var v22 uint32 = 4294967295
                    v15 = 42
                    var v23 uint8 = 160
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                v15 += v15
                _ = fi18
            }
        default:
            var v24 int32 = -1
            _ = v24
        }
        switch "hello" {
        case "\n\t":
            if ((!(v3 == v3)) || v16) {
                fmt.Println(v16, v1)
                v25 := GenS2{F0: (v15 / (((v15 % (((v15 & -1)) | 1))) | 1))}
                _ = v25
            } else {
                var v26 uint32 = 4294967295
                switch -9223372036854775808 {
                case 1:
                    if (((v3 >> (((^(v1))) & 63)) <= v3) || v16) {
                        for fi27 := 0; fi27 < 2; fi27++ {
                            for fi28 := 0; fi28 < 9; fi28++ {
                                v29 := GenS3{F0: ""}
                                switch ("a" + "hello") {
                                case "a":
                                    if len(v2) > 0 {
                                        v2[0] = "hello"
                                    }
                                case "":
                                    var v30 uint64 = 484
                                    _ = v30
                                default:
                                    v15 = -1
                                }
                                _ = len(v2)
                                var v31 bool = (uint16(v26) >= 65535)
                                _ = fi28
                                _ = v29
                                _ = v31
                            }
                            if (v16 || (("a" + ("" + ((("\n\t" + (("" + "") + "")) + "\n\t") + ("foo bar" + "\n\t")))) == "\n\t")) {
                                v32 := GenS1{F0: (v26 * v26)}
                                v2 = append(v2, "foo bar")
                                fmt.Println(v1)
                                _ = v32
                            }
                            var v33 int8 = (v3 + v3)
                            _ = fi27
                            _ = v33
                        }
                    } else {
                        var v34 int16 = 32767
                        _ = v34
                    }
                case -334:
                    var v35 int16 = 789
                    _ = v35
                }
                _ = v26
            }
        case "":
            _ = len(v2)
        case "a":
            fmt.Println(v3)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
