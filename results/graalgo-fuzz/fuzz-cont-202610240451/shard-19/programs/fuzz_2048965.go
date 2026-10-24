package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int8; F1 int32 }

func main() {
    var v1 int16 = 1000
    v2 := []uint32{0, 42, 4294967295, 4294967295}
    var v3 uint16 = 65535
    v4 := []string{"foo bar", "hello", ("hello" + "foo bar")}
    var v5 int8 = 1
    v1 = (v1 << ((719) & 63))
    v6 := GenS1{F0: 32768, F1: 0}
    var v7 uint64 = 779
    v3 = v3
    v8 := GenS1{F0: 2147483647, F1: -975}
    if false {
        var v9 string = "\n\t"
        _ = v9
    }
    var v10 int64 = -1
    if false {
        for fi11 := 0; fi11 < 13; fi11++ {
            v12 := GenS1{F0: (fi11 ^ (fi11 >> ((462) & 63))), F1: (fi11 * int(v7))}
            v10 = v10
            if true {
                v10 += (v10 - v10)
                v5 = v5
            }
            _ = fi11
            _ = v12
        }
        if (!false) {
            var v13 uint16 = (v3 | 11)
            v14 := GenS1{F0: 42, F1: 4294967295}
            var v15 string = "hello"
            _ = v13
            _ = v14
            _ = v15
        }
        fmt.Println(v3, v1, v1)
    }
    var v16 int8 = v5
    v17 := GenS2{F0: (v5 + -40), F1: int32(v1)}
    v7 -= 18446744073709551615
    for fi18 := 0; fi18 < 2; fi18++ {
        v4 = append(v4, "a")
        for fi19 := 0; fi19 < 11; fi19++ {
            var v20 uint64 = 18446744073709551615
            v21 := GenS1{F0: 127, F1: -683}
            _ = fi19
            _ = v20
            _ = v21
        }
        var v22 int64 = 0
        _ = fi18
        _ = v22
    }
    v3 = (^(v3))
    if ((v3 - 65535) > (^(v3))) {
        switch 1 {
        case -48:
            if len(v4) > 0 {
                v4[0] = ("a" + "hello")
            }
        default:
            fmt.Println(v10, v3, v5)
        }
        var v23 uint16 = 0
        switch "a" {
        case "hello":
            for fi24 := 0; fi24 < 12; fi24++ {
                fmt.Println(v5, v10, v7)
                _ = fi24
            }
        case "foo bar":
            _ = len(v2)
        case "\n\t":
            var v25 int64 = 42
            _ = v25
        default:
            var v26 int16 = (v1 ^ v1)
            _ = v26
        }
        _ = v23
    } else {
        v27 := GenS1{F0: 384, F1: 784}
        for fi28 := 0; fi28 < 4; fi28++ {
            if (42 >= 2147483647) {
                fmt.Println(v10)
                if false {
                    var v29 uint8 = 86
                    _ = v29
                } else {
                    v30 := GenS1{F0: fi28, F1: fi28}
                    _ = v30
                }
            } else {
                if (4294967295 >= (v7 & (v7 / ((v7) | 1)))) {
                    var v31 string = ("" + "\n\t")
                    switch fi28 {
                    case 9223372036854775807:
                        if (("a" + ((v31 + (("\n\t" + "") + "\n\t")) + v31)) > "") {
                            var v32 uint32 = 2
                            for fi33 := 0; fi33 < 4; fi33++ {
                                if false {
                                    v3 = v3
                                    switch (fi33 + (-(fi28))) {
                                    case -2147483648:
                                        var v34 uint8 = 0
                                        _ = v34
                                    case 127:
                                        v35 := GenS1{F0: fi28, F1: int(v7)}
                                        _ = v35
                                    }
                                    v1 -= -1
                                    var v36 uint64 = 1
                                    _ = v36
                                }
                                var v37 int8 = v16
                                var v38 int64 = int64(v3)
                                _ = fi33
                                _ = v37
                                _ = v38
                            }
                            _ = v32
                        } else {
                            fmt.Println(v3, v10, v1)
                            var v39 int64 = (v10 * (v10 - 0))
                            _ = v39
                        }
                    default:
                        fmt.Println(v5, v7)
                    }
                    _ = v31
                }
                v7 = v7
            }
            var v40 string = ((("a" + (("foo bar" + "\n\t") + "hello")) + "\n\t") + "\n\t")
            v5 += (^(v5))
            _ = fi28
            _ = v40
        }
        _ = v27
    }
    var v41 uint32 = 42
    v4 = append(v4, "foo bar")
    fmt.Println(v3, v16, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
