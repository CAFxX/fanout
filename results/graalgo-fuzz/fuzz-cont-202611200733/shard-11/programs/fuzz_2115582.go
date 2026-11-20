package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() string {
    var v5 int16 = -32768
    v6 := GenS1{F0: 0}
    var v7 string = ("\n\t" + "foo bar")
    var v8 float32 = 1.5
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return v7
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]string)
    v3 := make(map[uint32]uint32)
    var v4 int = -1
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 uint8 = 99
        _ = fi9
        _ = v10
    }
    var v11 uint32 = 246
    var v12 int16 = -1
    v13 := GenS1{F0: 692}
    var v14 uint64 = 4294967295
    v11 = (v11 + v11)
    v11 += (v11 * 2)
    for fi15 := 0; fi15 < 19; fi15++ {
        if ((255 <= 255) && (!true)) {
            for fi16 := 0; fi16 < 2; fi16++ {
                var v17 uint32 = v11
                _ = fi16
                _ = v17
            }
            fmt.Println(v12)
        } else {
            switch (("foo bar" + "a") + (("" + (("a" + ("foo bar" + "foo bar")) + ("hello" + ((("a" + "\n\t") + "a") + "hello")))) + (("\n\t" + ("\n\t" + "")) + "foo bar"))) {
            case "":
                if true {
                    var v18 int64 = 1
                    v12 += -32768
                    _ = v18
                } else {
                    v19 := GenS1{F0: 0}
                    _ = v19
                }
            case "foo bar":
                v11 = v11
            }
            for fi20 := 0; fi20 < 6; fi20++ {
                v12 = (v12 << ((1) & 63))
                fmt.Println(v11, v11)
                v4 = v4
                _ = fi20
            }
            if true {
                v21 := GenS1{F0: 613}
                for fi22 := 0; fi22 < 4; fi22++ {
                    switch ("a" + ("hello" + ("a" + ((("foo bar" + (("foo bar" + ("" + "\n\t")) + "a")) + "foo bar") + ((("" + "") + "foo bar") + "a"))))) {
                    case "foo bar":
                        var v23 uint = 255
                        _ = v23
                    case "\n\t":
                        v11 += (v11 + 2147483647)
                    case "a":
                        fmt.Println(v11)
                    default:
                        var v24 int16 = (v12 | v12)
                        _ = v24
                    }
                    if ((v12 <= (v12 % (((v12 ^ 922)) | 1))) && ((0.1 / 0.1) >= (0.0 + (1e10 - float32(((3.14159 * 1.5) / 3.14159)))))) {
                        fmt.Println(fi22, fi15)
                        v25 := GenS1{F0: int32(v11)}
                        _ = v25
                    } else {
                        fmt.Println(v12, v11, fi15)
                        var v26 int32 = -505
                        _ = v26
                    }
                    v27 := GenS1{F0: -2147483648}
                    _ = fi22
                    _ = v27
                }
                switch "foo bar" {
                case "hello":
                    if true {
                        v11 = 1
                        fmt.Println(v12)
                    }
                default:
                    v4 -= (fi15 ^ -9223372036854775808)
                }
                _ = v21
            } else {
                switch fi15 {
                case 526:
                    v4 = (fi15 / ((fi15) | 1))
                }
                if false {
                    if ((^(v14)) >= (^(v14))) {
                        var v28 uint64 = (v14 / (((v14 + 0)) | 1))
                        _ = v28
                    } else {
                        fmt.Println(v12, v14)
                        var v29 int32 = 694
                        _ = v29
                    }
                } else {
                    fmt.Println(v12, v11, v11)
                    var v30 int8 = 1
                    _ = v30
                }
            }
        }
        fmt.Println(v4, v12, v14)
        _ = fi15
    }
    if false {
        var v31 float64 = (-2.25 + 0.1)
        switch v4 {
        case -2147483648:
            var v32 float32 = 1.5
            _ = v32
        case 32768:
            for fi33 := 0; fi33 < 8; fi33++ {
                v34 := GenS1{F0: -1}
                if (true || (!(!(true && false)))) {
                    v12 -= (v12 * v12)
                }
                _ = fi33
                _ = v34
            }
        case -216:
            fmt.Println(v14, v4, v12)
        default:
            var v35 string = (("hello" + "\n\t") + "")
            _ = v35
        }
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
