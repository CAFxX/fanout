package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int32 = -464
    v2 := []uint32{65535, 1, 4294967295}
    var v3 int16 = 1000
    v4 := make(map[string]uint32)
    v5 := make(map[int]int)
    var v6 int8 = -70
    var v7 uint8 = 2
    fmt.Println(v7, v6, v3)
    if true {
        var v8 string = "foo bar"
        switch -2147483648 {
        case 42:
            fmt.Println(v6, v1)
        case -9223372036854775808:
            if (2147483647 >= 846) {
                var v9 float32 = 0.0
                v2 = append(v2, 2147483648)
                _ = v9
            } else {
                switch v8 {
                case "foo bar":
                    fmt.Println(v6)
                case "":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case "hello":
                    v3 += 42
                default:
                    if false {
                        v3 = v3
                        if (!(-1 < -347)) {
                            switch 2147483647 {
                            case -826:
                                var v10 float64 = 3.14159
                                _ = v10
                            case 4294967295:
                                var v11 uint16 = 237
                                _ = v11
                            }
                            v2 = append(v2, 65535)
                        } else {
                            if true {
                                v1 += v1
                            }
                            fmt.Println(v6, v7)
                            fmt.Println(v3)
                        }
                        for fi12 := 0; fi12 < 11; fi12++ {
                            fmt.Println(v8, v7, fi12)
                            for fi13 := 0; fi13 < 10; fi13++ {
                                v7 = v7
                                v8 = ""
                                _ = fi13
                            }
                            _ = fi12
                        }
                    }
                }
                var v14 int16 = 0
                fmt.Println(v14)
                _ = v14
            }
        }
        _ = v8
    }
    var v15 int32 = (v1 * (v1 * v1))
    fmt.Println(v1)
    delete(v4, ((("foo bar" + "\n\t") + "foo bar") + "hello"))
    if true {
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 uint = 65535
            v3 = int16(fi16)
            var v18 int32 = -2147483648
            _ = fi16
            _ = v17
            _ = v18
        }
        v7 = v7
    } else {
        for fi19 := 0; fi19 < 16; fi19++ {
            if (false && (v6 < v6)) {
                switch ("foo bar" + "\n\t") {
                case "":
                    if ((^(v3)) != v3) {
                        if true {
                            var v20 uint8 = (v7 + (v7 - (v7 | (^(v7)))))
                            v1 = (v1 ^ v1)
                            _ = v20
                        } else {
                            v3 += 42
                        }
                    } else {
                        v3 -= 32767
                    }
                }
                v21 := GenS1{F0: int32(v6)}
                v7 += (v7 | v7)
                _ = v21
            }
            _ = fi19
        }
        _ = len(v2)
    }
    if true {
        switch 128 {
        case 690:
            var v22 uint32 = 995
            _ = v22
        case 19:
            var v23 int64 = 9223372036854775807
            _ = v23
        }
    }
    v7 = 0
    _ = len(v4)
    v24 := GenS1{F0: int32(v3)}
    var v25 uint8 = (v7 * v7)
    switch -2147483648 {
    case 32767:
        var v26 uint16 = 0
        _ = v26
    default:
        v3 = v3
    }
    v3 -= 1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
