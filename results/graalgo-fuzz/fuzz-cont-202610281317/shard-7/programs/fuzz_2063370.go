package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int32 = 112
    v2 := []uint32{1}
    var v3 uint64 = 1
    v4 := make(map[uint32]uint32)
    var v5 int16 = 32767
    switch 4294967295 {
    case 9223372036854775807:
        if (!(-1 <= 400)) {
            if (!false) {
                var v6 uint8 = 141
                _ = v6
            }
            var v7 uint = 255
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v7
        }
    case 128:
        if (-1 >= 42) {
            var v8 uint16 = 425
            _ = v8
        }
    default:
        var v9 int8 = 42
        _ = v9
    }
    if ((v3 * v3) != 2) {
        fmt.Println(v1, v5)
        if false {
            var v10 int16 = v5
            _ = v10
        } else {
            v1 = v1
            fmt.Println(v1, v5, v3)
        }
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        for fi12 := 0; fi12 < 5; fi12++ {
            v13 := v4[0]
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    switch -9223372036854775808 {
    case -604:
        var v14 float64 = float64((1e10 / float32(1e10)))
        _ = v14
    case 32767:
        var v15 int64 = 1
        _ = v15
    }
    v1 += (-(v1))
    if true {
        var v16 uint = 65535
        v4[0] = 277
        _ = v16
    }
    var v17 int16 = (v5 << ((uint(v3)) & 63))
    switch ("\n\t" + "hello") {
    case "a":
        v18 := GenS1{F0: "a"}
        _ = v18
    case "foo bar":
        v17 -= v5
    case "\n\t":
        if (-2.25 > -2.25) {
            v4[823] = 585
        } else {
            for fi19 := 0; fi19 < 8; fi19++ {
                v1 += v1
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi19
            }
            fmt.Println(v1, v3)
            for fi20 := 0; fi20 < 7; fi20++ {
                v17 = v5
                switch "a" {
                case "\n\t":
                    var v21 int32 = v1
                    _ = v21
                case "hello":
                    v17 = v5
                case "foo bar":
                    if ((fi20 + 32768) >= (fi20 + -9223372036854775808)) {
                        if (((fi20 ^ fi20) < fi20) && true) {
                            for fi22 := 0; fi22 < 15; fi22++ {
                                switch (fi22 - fi22) {
                                case 128:
                                    if false {
                                        var v23 uint = 4294967295
                                        v23 = 712
                                        _ = v23
                                    }
                                case 127:
                                    delete(v4, 271)
                                    v5 -= (v5 >> ((1) & 63))
                                    var v24 float32 = 1e10
                                    _ = v24
                                default:
                                    switch (-(fi20)) {
                                    case 287:
                                        v25 := GenS1{F0: (((("" + ("foo bar" + "a")) + "a") + (("foo bar" + "a") + "\n\t")) + ("a" + ""))}
                                        _ = v25
                                    case -10:
                                        if true {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            if (255 == 85) {
                                                var v26 float32 = 0.0
                                                v27 := GenS1{F0: ""}
                                                _ = v26
                                                _ = v27
                                            }
                                            var v28 bool = (1 < 42)
                                            _ = v28
                                        } else {
                                            fmt.Println(v3, v5, v3)
                                            delete(v4, 457)
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        switch (fi20 >> ((0) & 63)) {
                                        case 401:
                                            v5 = v5
                                        case 128:
                                            v29 := GenS1{F0: ""}
                                            _ = v29
                                        case 823:
                                            v3 = v3
                                        }
                                        v2 = append(v2, 2147483647)
                                    }
                                }
                                _ = fi22
                            }
                            v3 = v3
                            v1 = (v1 - v1)
                        } else {
                            if true {
                                var v30 uint16 = 42
                                _ = v30
                            } else {
                                fmt.Println(fi20)
                            }
                            var v31 uint16 = 502
                            _ = v31
                        }
                    } else {
                        v4[4294967295] = 965
                        v1 += (v1 & int32(v3))
                    }
                }
                for fi32 := 0; fi32 < 4; fi32++ {
                    v33 := GenS1{F0: ("a" + "foo bar")}
                    fmt.Println(v1)
                    _ = fi32
                    _ = v33
                }
                _ = fi20
            }
        }
    }
    for fi34 := 0; fi34 < 1; fi34++ {
        fmt.Println(v5)
        _ = fi34
    }
    var v35 uint16 = 255
    switch "a" {
    case "":
        v3 = v3
    }
    var v36 uint8 = 255
    var v37 float64 = 3.14159
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
