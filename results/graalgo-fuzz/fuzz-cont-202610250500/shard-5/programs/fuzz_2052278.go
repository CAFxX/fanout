package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 string; F1 int }

func main() {
    var v1 float32 = 0.1
    var v2 bool = (true || false)
    var v3 uint = 550
    v4 := []int{4294967295, 4294967295, -93, 451}
    if (42 == 42) {
        if ((((0.1 + 3.14159) + ((1e10 - 3.14159) + (-2.25 - 0.0))) + 1e10) != 1.5) {
            v2 = v2
            v3 = 1
            var v5 float32 = v1
            _ = v5
        } else {
            if false {
                _ = len(v4)
                v6 := GenS1{F0: -384, F1: (!false)}
                _ = v6
            } else {
                var v7 int16 = 42
                _ = v7
            }
            var v8 uint32 = 65535
            _ = v8
        }
        for fi9 := 0; fi9 < 15; fi9++ {
            if len(v4) > 0 {
                v4[0] = 0
            }
            v10 := GenS2{F0: 42, F1: ""}
            if true {
                var v11 int32 = 464
                _ = v11
            }
            _ = fi9
            _ = v10
        }
        var v12 int16 = -32768
        _ = v12
    } else {
        for fi13 := 0; fi13 < 4; fi13++ {
            if len(v4) > 0 {
                v4[0] = 607
            }
            var v14 string = ("a" + ("" + ("\n\t" + "\n\t")))
            fmt.Println(v14, v14)
            _ = fi13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        if (0 == 2147483647) {
            v2 = v2
        }
        v16 := GenS2{F0: 2, F1: (("" + "a") + "foo bar")}
        _ = fi15
        _ = v16
    }
    var v17 uint64 = 42
    if v2 {
        fmt.Println(v2, v17, v1)
        switch -2147483648 {
        case -246:
            if v2 {
                var v18 int16 = 42
                v19 := GenS3{F0: 0, F1: 9223372036854775807}
                _ = v18
                _ = v19
            }
        case 0:
            switch 9223372036854775807 {
            case 269:
                switch 32767 {
                case -809:
                    fmt.Println(v2)
                default:
                    if (((!v2) && false) && (576 != 32767)) {
                        var v20 int = 128
                        _ = v20
                    }
                }
            }
        default:
            var v21 uint = (^(v3))
            _ = v21
        }
        v22 := GenS4{F0: v2}
        _ = v22
    } else {
        fmt.Println(v1, v3)
        var v23 uint32 = 250
        fmt.Println(v3, v3)
        _ = v23
    }
    if len(v4) > 0 {
        v4[0] = -21
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        switch "foo bar" {
        case "a":
            fmt.Println(v17, v1)
        default:
            v25 := GenS4{F0: (-9223372036854775808 < fi24)}
            _ = v25
        }
        _ = fi24
    }
    if (42 >= 1) {
        for fi26 := 0; fi26 < 19; fi26++ {
            switch "a" {
            case "":
                for fi27 := 0; fi27 < 8; fi27++ {
                    v4 = append(v4, (fi26 % ((int(v3)) | 1)))
                    v28 := GenS4{F0: (2147483647 != 859)}
                    _ = fi27
                    _ = v28
                }
            case "\n\t":
                var v29 int32 = -1
                _ = v29
            default:
                for fi30 := 0; fi30 < 7; fi30++ {
                    fmt.Println(v2)
                    _ = fi30
                }
            }
            fmt.Println(v3, v3)
            _ = fi26
        }
    } else {
        var v31 uint16 = 0
        _ = v31
    }
    if true {
        fmt.Println(v17)
        v3 = (^(v3))
    } else {
        v32 := GenS3{F0: 84, F1: 44}
        _ = v32
    }
    for fi33 := 0; fi33 < 15; fi33++ {
        if v2 {
            for fi34 := 0; fi34 < 14; fi34++ {
                fmt.Println(v2, v1)
                if len(v4) > 0 {
                    v4[0] = (fi33 | -650)
                }
                v35 := GenS5{F0: v17, F1: -933}
                v36 := GenS6{F0: ("" + "hello"), F1: fi33}
                var v37 uint32 = 22
                var v38 uint8 = 42
                _ = fi34
                _ = v35
                _ = v36
                _ = v37
                _ = v38
            }
        } else {
            v39 := GenS3{F0: 255, F1: 0}
            var v40 int = fi33
            _ = v39
            _ = v40
        }
        _ = fi33
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v4)
    var v41 uint8 = 1
    v41 = v41
    var v42 bool = (!(4294967295 >= (^(v3))))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
}
