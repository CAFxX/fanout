package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 int64 }
func genFn0(p0 uint16, p1 int8) string {
    v6 := GenS1{F0: -1}
    var v7 int = 9223372036854775807
    var v8 uint = 42
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return ("a" + "a")
}

func main() {
    var v1 int16 = -32768
    v2 := []uint32{681, 193, 2147483647}
    v3 := make(map[string]uint32)
    v4 := []uint32{0, 42, 1}
    var v5 uint16 = 216
    fmt.Println(v5, v1)
    if true {
        v1 = (v1 / (((v1 | 1000)) | 1))
        if false {
            switch "a" {
            case "":
                v1 -= 0
            case "\n\t":
                if false {
                    if (float64(0.1) != 1e10) {
                        v1 = v1
                        if ("foo bar" == "hello") {
                            if true {
                                for fi9 := 0; fi9 < 10; fi9++ {
                                    if len(v2) > 0 {
                                        v2[0] = 4294967295
                                    }
                                    _ = fi9
                                }
                                v5 = (v5 << ((42) & 63))
                                delete(v3, ("hello" + ("foo bar" + (("a" + "") + "\n\t"))))
                            }
                        } else {
                            v10 := GenS2{F0: 8, F1: 1}
                            for fi11 := 0; fi11 < 20; fi11++ {
                                v1 += 1000
                                _ = fi11
                            }
                            _ = v10
                        }
                    }
                    v12 := GenS1{F0: -2147483648}
                    _ = v12
                }
            default:
                _ = len(v3)
            }
            for fi13 := 0; fi13 < 3; fi13++ {
                for fi14 := 0; fi14 < 10; fi14++ {
                    v1 = v1
                    _ = fi14
                }
                for fi15 := 0; fi15 < 17; fi15++ {
                    var v16 int32 = 2147483647
                    if true {
                        v17 := v3[""]
                        v17 -= (^(v17))
                        _ = v17
                    } else {
                        v16 += (-(v16))
                        switch -946 {
                        case 0:
                            switch -2147483648 {
                            case -9223372036854775808:
                                var v18 int = fi13
                                _ = v18
                            case 932:
                                switch (fi13 >> ((305) & 63)) {
                                case 213:
                                    v16 = 42
                                case 2147483647:
                                    var v19 uint = 65535
                                    _ = v19
                                case 127:
                                    v16 = (v16 >> ((65535) & 63))
                                    fmt.Println(fi13)
                                }
                            default:
                                v3[("foo bar" + "a")] = 2147483648
                            }
                        }
                    }
                    v20 := GenS2{F0: 225, F1: 790}
                    var v21 uint = 481
                    var v22 int8 = 74
                    _ = fi15
                    _ = v16
                    _ = v20
                    _ = v21
                    _ = v22
                }
                _ = fi13
            }
        }
    } else {
        if true {
            var v23 float32 = (1e10 / float32(-2.25))
            v1 += v1
            _ = v23
        } else {
            var v24 uint64 = 2
            v1 -= 1
            var v25 int32 = 166
            _ = v24
            _ = v25
        }
        v1 = (^(v1))
        _ = len(v3)
    }
    v1 -= v1
    fmt.Println(v5, v1)
    v5 += v5
    var v26 uint32 = 2147483647
    v3["a"] = 0
    v1 -= (v1 / (((^(v1))) | 1))
    _ = len(v2)
    v27 := GenS2{F0: 128, F1: 9223372036854775807}
    v5 -= v5
    for fi28 := 0; fi28 < 4; fi28++ {
        for fi29 := 0; fi29 < 11; fi29++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if false {
                var v30 bool = (!true)
                _ = v30
            }
            _ = fi29
        }
        _ = fi28
    }
    var v31 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
