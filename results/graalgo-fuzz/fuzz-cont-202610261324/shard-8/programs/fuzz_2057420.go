package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 bool }

func main() {
    v1 := []uint32{808}
    var v2 int = 4294967295
    var v3 int16 = -32768
    _ = len(v1)
    if true {
        fmt.Println(v3, v3)
        fmt.Println(v3, v3)
        if len(v1) > 0 {
            v1[0] = 1
        }
    }
    v3 -= (v3 >> ((940) & 63))
    v2 = v2
    v2 += v2
    if true {
        fmt.Println(v3, v3)
        v3 -= v3
    } else {
        var v4 uint16 = 32767
        var v5 bool = true
        _ = v4
        _ = v5
    }
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 string = ("a" + "foo bar")
        _ = fi6
        _ = v7
    }
    var v8 int16 = v3
    switch (-(v2)) {
    case -523:
        v9 := GenS1{F0: 1}
        _ = v9
    case -530:
        if (!false) {
            for fi10 := 0; fi10 < 17; fi10++ {
                var v11 uint = 2
                v12 := GenS2{F0: 127}
                _ = fi10
                _ = v11
                _ = v12
            }
            var v13 uint16 = 255
            _ = v13
        } else {
            for fi14 := 0; fi14 < 18; fi14++ {
                for fi15 := 0; fi15 < 3; fi15++ {
                    switch ("\n\t" + "\n\t") {
                    case "":
                        switch "a" {
                        case "a":
                            _ = len(v1)
                        case "foo bar":
                            switch "" {
                            case "\n\t":
                                v2 += 1
                            }
                        case "\n\t":
                            v16 := GenS1{F0: 255}
                            _ = v16
                        }
                    case "hello":
                        var v17 int = fi15
                        _ = v17
                    }
                    _ = fi15
                }
                _ = fi14
            }
            switch (v2 | v2) {
            case -929:
                for fi18 := 0; fi18 < 2; fi18++ {
                    v3 += v3
                    _ = fi18
                }
            }
            v3 = v3
        }
    case 1:
        for fi19 := 0; fi19 < 11; fi19++ {
            for fi20 := 0; fi20 < 6; fi20++ {
                v3 += (v8 % ((v3) | 1))
                switch -184 {
                case 42:
                    if true {
                        v21 := GenS3{F0: 13, F1: (true && false)}
                        _ = v21
                    }
                case 4294967295:
                    switch (fi20 << ((566) & 63)) {
                    case 127:
                        var v22 uint = 4294967295
                        _ = v22
                    case 128:
                        if false {
                            if len(v1) > 0 {
                                v1[0] = 1
                            }
                            v8 += (v8 * (v3 * v3))
                        }
                        if (fi20 < 2147483647) {
                            switch ((("hello" + (("" + "\n\t") + "")) + "hello") + ("hello" + "")) {
                            case "":
                                var v23 int8 = 0
                                _ = v23
                            case "\n\t":
                                fmt.Println(v8)
                            case "hello":
                                var v24 int8 = 125
                                _ = v24
                            default:
                                v2 -= fi20
                            }
                            if len(v1) > 0 {
                                v1[0] = 2
                            }
                            v2 = v2
                        }
                        v2 = fi20
                    case 9223372036854775807:
                        v2 = -174
                        var v25 int8 = 42
                        var v26 uint = 0
                        _ = v25
                        _ = v26
                    default:
                        v3 -= 0
                    }
                case -560:
                    var v27 uint32 = 2147483647
                    _ = v27
                }
                _ = fi20
            }
            _ = fi19
        }
    default:
        if len(v1) > 0 {
            v1[0] = 373
        }
    }
    fmt.Println(v3, v3, v8)
    var v28 int32 = 0
    v29 := GenS1{F0: 162}
    v2 = (^(v2))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v8
}
