package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := make(map[string]string)
    var v2 uint = 95
    v3 := []int{9223372036854775807, -9223372036854775808, 32768}
    v4 := []uint32{2147483648, 1, 2147483648}
    v5 := GenS1{F0: 721}
    v6 := GenS1{F0: 568}
    for fi7 := 0; fi7 < 4; fi7++ {
        delete(v1, "foo bar")
        v2 -= v2
        fmt.Println(v2, v2)
        _ = fi7
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        if false {
            switch fi8 {
            case 32767:
                switch ("hello" + "a") {
                case "hello":
                    for fi9 := 0; fi9 < 16; fi9++ {
                        switch (fi9 / ((fi9) | 1)) {
                        case -977:
                            v2 = (v2 << ((v2) & 63))
                        case -569:
                            v2 = (v2 & 251)
                        case 127:
                            v1[""] = "foo bar"
                        }
                        _ = fi9
                    }
                case "":
                    if false {
                        if len(v4) > 0 {
                            v4[0] = uint32(v2)
                        }
                        var v10 float32 = 1.5
                        _ = v10
                    } else {
                        var v11 int8 = 0
                        _ = v11
                    }
                case "a":
                    var v12 int = fi8
                    _ = v12
                default:
                    if false {
                        for fi13 := 0; fi13 < 11; fi13++ {
                            var v14 float32 = 1e10
                            var v15 int16 = 42
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                        fmt.Println(v2, v2)
                        var v16 int8 = 100
                        _ = v16
                    } else {
                        var v17 string = "foo bar"
                        _ = v17
                    }
                }
            case 128:
                if false {
                    v18 := GenS1{F0: 2147483647}
                    switch fi8 {
                    case 128:
                        v2 -= (^(v2))
                    case -883:
                        delete(v1, ("" + ("foo bar" + ((("foo bar" + "hello") + (("\n\t" + "foo bar") + "")) + ("" + ("a" + ("a" + "foo bar")))))))
                    case -9223372036854775808:
                        if len(v4) > 0 {
                            v4[0] = 1
                        }
                    default:
                        v2 = v2
                    }
                    _ = v18
                }
            default:
                v19 := GenS2{F0: 713}
                _ = v19
            }
            v20 := GenS2{F0: -32768}
            _ = v20
        }
        for fi21 := 0; fi21 < 2; fi21++ {
            if ((!true) && (true && (fi21 <= fi21))) {
                var v22 uint32 = 2147483647
                v2 += (v2 ^ (v2 & 1))
                _ = v22
            }
            if true {
                fmt.Println(fi8)
                switch fi8 {
                case 575:
                    v2 = v2
                }
                var v23 bool = (1 == 0)
                _ = v23
            }
            fmt.Println(fi21, fi8, fi21)
            _ = fi21
        }
        var v24 int8 = -26
        _ = fi8
        _ = v24
    }
    v25 := GenS1{F0: 1}
    v26 := GenS2{F0: -32768}
    var v27 float32 = float32(0.0)
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
