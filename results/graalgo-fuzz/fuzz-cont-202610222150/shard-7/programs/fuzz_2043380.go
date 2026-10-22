package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 uint64; F2 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32 }

func main() {
    var v1 uint16 = 255
    v2 := []uint32{435, 49, 2}
    var v3 uint32 = 2
    v4 := GenS1{F0: (("hello" + ("\n\t" + "a")) + "")}
    if true {
        v5 := GenS2{F0: "foo bar"}
        _ = v5
    } else {
        switch 32767 {
        case 2147483647:
            for fi6 := 0; fi6 < 14; fi6++ {
                if (v1 == (v1 | v1)) {
                    if (!false) {
                        fmt.Println(v3, v3)
                    } else {
                        for fi7 := 0; fi7 < 17; fi7++ {
                            var v8 int16 = -32768
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v3 -= (v3 ^ (^(v3)))
                            switch "hello" {
                            case "":
                                fmt.Println(v8, v3, v1)
                            case "hello":
                                v3 = v3
                                switch fi7 {
                                case 0:
                                    fmt.Println(v8, fi6)
                                default:
                                    v3 = v3
                                }
                            case "foo bar":
                                v8 -= (v8 + (v8 * v8))
                            default:
                                var v9 int32 = 190
                                _ = v9
                            }
                            fmt.Println(v1, fi6)
                            _ = fi7
                            _ = v8
                        }
                        var v10 int16 = 0
                        _ = v10
                    }
                    v1 = (v1 & v1)
                    v1 += v1
                }
                _ = fi6
            }
        case -282:
            v11 := GenS3{F0: 42, F1: 0, F2: 750}
            _ = v11
        case 9223372036854775807:
            for fi12 := 0; fi12 < 10; fi12++ {
                var v13 int8 = -1
                for fi14 := 0; fi14 < 7; fi14++ {
                    var v15 string = ("\n\t" + ("" + ("foo bar" + "a")))
                    v15 = (("hello" + (("a" + "hello") + ("a" + ("a" + (v15 + (("foo bar" + "a") + v15)))))) + "\n\t")
                    switch v15 {
                    case "a":
                        v13 += (v13 ^ v13)
                    case "":
                        var v16 bool = true
                        v3 -= (v3 * (v3 * 2147483647))
                        var v17 int8 = 42
                        v18 := GenS4{F0: -414}
                        fmt.Println(v16)
                        _ = v16
                        _ = v17
                        _ = v18
                    default:
                        if len(v2) > 0 {
                            v2[0] = (v3 >> ((904) & 63))
                        }
                    }
                    _ = fi14
                    _ = v15
                }
                _ = fi12
                _ = v13
            }
        }
        var v19 bool = false
        _ = v19
    }
    var v20 uint64 = 42
    fmt.Println(v20, v20)
    v21 := GenS5{F0: -1}
    v20 = (v20 ^ v20)
    var v22 string = (("\n\t" + "a") + (("foo bar" + "a") + (("foo bar" + ("\n\t" + "a")) + "")))
    v20 += 1
    v3 -= v3
    v20 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
