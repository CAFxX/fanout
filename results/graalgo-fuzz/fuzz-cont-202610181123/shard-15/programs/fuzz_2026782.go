package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint }
type GenS2 struct { F0 uint; F1 string; F2 int16 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 int = 32767
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    var v4 float64 = 3.14159
    var v5 int32 = 2147483647
    var v6 uint16 = 255
    var v7 int32 = v5
    v8 := []int{(v1 + 115), v1, (v1 | (^(v1))), (v1 + v1)}
    if false {
        v6 -= v6
    } else {
        v9 := v3[("a" + "hello")]
        var v10 string = "hello"
        if false {
            switch "" {
            case "foo bar":
                v11 := GenS1{F0: 4294967295, F1: "\n\t", F2: 65535}
                _ = v11
            }
            for fi12 := 0; fi12 < 1; fi12++ {
                var v13 int32 = (v7 << ((0) & 63))
                v5 = v7
                _ = fi12
                _ = v13
            }
        } else {
            var v14 float32 = (((float32(v4) + ((((1.5 * (1e10 - 1.5)) + (0.1 * 0.1)) - 1.5) * 0.0)) + 1.5) / (((3.14159 + (-2.25 * 0.0)) + 1.5) * 1.5))
            v6 -= v6
            fmt.Println(v7, v7)
            _ = v14
        }
        _ = v10
        _ = v9
    }
    v15 := GenS2{F0: 1, F1: "\n\t", F2: -1}
    v5 -= (v7 + (v5 << ((161) & 63)))
    for fi16 := 0; fi16 < 18; fi16++ {
        switch int(v7) {
        case -2147483648:
            v4 = v4
        case 32767:
            switch "foo bar" {
            case "foo bar":
                switch fi16 {
                case 9223372036854775807:
                    v5 -= 2147483647
                case 32768:
                    var v17 float32 = (-2.25 + 0.0)
                    _ = v17
                case 583:
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                default:
                    for fi18 := 0; fi18 < 7; fi18++ {
                        if (127 != 48) {
                            if true {
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                                switch (v1 % ((9223372036854775807) | 1)) {
                                case -30:
                                    fmt.Println(fi16, v5, v4)
                                case 627:
                                    fmt.Println(v1, v6)
                                default:
                                    if len(v8) > 0 {
                                        v8[0] = fi18
                                    }
                                }
                                v3["a"] = "foo bar"
                                if ((v4 > v4) || (919 <= 9223372036854775807)) {
                                    var v19 int = 9223372036854775807
                                    v4 = 0.0
                                    if (false || (43 != 128)) {
                                        var v20 uint32 = 2147483647
                                        _ = v20
                                    } else {
                                        delete(v3, "hello")
                                    }
                                    _ = v19
                                }
                                if len(v8) > 0 {
                                    v8[0] = fi16
                                }
                            }
                        } else {
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            var v21 bool = (!true)
                            _ = v21
                        }
                        _ = len(v8)
                        _ = fi18
                    }
                }
            }
        default:
            for fi22 := 0; fi22 < 5; fi22++ {
                var v23 int = (fi22 * -2147483648)
                v24 := GenS3{F0: 362}
                fmt.Println(v23)
                _ = fi22
                _ = v23
                _ = v24
            }
        }
        var v25 uint = 972
        v26 := GenS3{F0: 80}
        _ = fi16
        _ = v25
        _ = v26
    }
    switch (v1 >> ((65535) & 63)) {
    case -1:
        delete(v2, 65535)
    case 32767:
        for fi27 := 0; fi27 < 8; fi27++ {
            v28 := GenS3{F0: 42}
            v7 = v5
            _ = fi27
            _ = v28
        }
    default:
        v5 = 563
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
