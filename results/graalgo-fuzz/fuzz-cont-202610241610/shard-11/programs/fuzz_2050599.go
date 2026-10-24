package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 bool }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int64) string {
    v5 := GenS1{F0: (32767 > 1000), F1: 255, F2: (float32((float64((3.14159 + 1.5)) * 0.0)) > -2.25)}
    var v6 int32 = 325
    _ = p0
    _ = v5
    _ = v6
    return "a"
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 108
    v3 := []uint32{377, 65535}
    var v4 uint64 = 933
    var v7 int32 = 2147483647
    if (-864 <= -1) {
        delete(v1, "hello")
        for fi8 := 0; fi8 < 10; fi8++ {
            v2 -= fi8
            if (true && true) {
                if false {
                    if (!(true || false)) {
                        v7 += (v7 >> ((2) & 63))
                    } else {
                        if true {
                            fmt.Println(fi8, fi8)
                            _ = len(v1)
                        } else {
                            if false {
                                for fi9 := 0; fi9 < 17; fi9++ {
                                    var v10 float64 = (float64(-2.25) / 1e10)
                                    _ = fi9
                                    _ = v10
                                }
                                _ = len(v1)
                                if false {
                                    fmt.Println(v2, v7, v2)
                                    var v11 int16 = int16(v7)
                                    _ = v11
                                }
                            }
                        }
                    }
                }
                fmt.Println(fi8, fi8, v7)
            }
            if true {
                fmt.Println(v4, v4)
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
            }
            _ = fi8
        }
    } else {
        v4 = v4
    }
    var v12 int32 = (v7 * 1)
    var v13 int16 = -370
    fmt.Println(v7)
    var v14 int16 = (v13 | v13)
    delete(v1, "")
    for fi15 := 0; fi15 < 10; fi15++ {
        v16 := GenS1{F0: false, F1: 4294967295, F2: (!true)}
        fmt.Println(v4, v7, v7)
        _ = fi15
        _ = v16
    }
    switch v2 {
    case -832:
        delete(v1, "\n\t")
    case 42:
        var v17 string = "a"
        _ = v17
    case 1:
        var v18 int64 = 42
        _ = v18
    default:
        for fi19 := 0; fi19 < 3; fi19++ {
            var v20 string = ("hello" + "foo bar")
            var v21 int64 = -1
            var v22 uint64 = 4294967295
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v12, v12)
    switch "" {
    case "hello":
        switch ("foo bar" + "\n\t") {
        case "hello":
            if (!(true || true)) {
                v23 := GenS1{F0: (false || (0.0 != (-2.25 + 1e10))), F1: 65535, F2: (-81 < 55)}
                _ = v23
            }
        default:
            v24 := GenS2{F0: 128}
            _ = v24
        }
    case "\n\t":
        if false {
            v3 = append(v3, 2)
        }
    default:
        for fi25 := 0; fi25 < 17; fi25++ {
            v1[(("a" + "hello") + ("foo bar" + "hello"))] = 942
            var v26 uint16 = uint16(v4)
            _ = fi25
            _ = v26
        }
    }
    if true {
        v2 += v2
        fmt.Println(v2, v12, v13)
        v27 := GenS1{F0: (65535 != 32767), F1: 18446744073709551615, F2: true}
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
