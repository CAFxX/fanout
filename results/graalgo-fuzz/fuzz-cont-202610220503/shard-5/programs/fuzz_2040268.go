package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }

func main() {
    v1 := []uint32{1, 4294967295, 1}
    var v2 int16 = -911
    v3 := make(map[int]int)
    v4 := GenS1{F0: ""}
    var v5 int8 = -128
    var v6 int32 = 317
    var v7 float64 = 1.5
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 int16 = (v2 >> ((65535) & 63))
        _ = fi8
        _ = v9
    }
    var v10 float64 = -2.25
    switch -810 {
    case 32767:
        for fi11 := 0; fi11 < 18; fi11++ {
            if false {
                switch 9223372036854775807 {
                case 32767:
                    var v12 int = (-(fi11))
                    _ = v12
                default:
                    var v13 uint32 = 2147483647
                    _ = v13
                }
                switch "foo bar" {
                case "a":
                    for fi14 := 0; fi14 < 4; fi14++ {
                        if (((v7 + v7) == ((3.14159 * (0.0 * 0.0)) + (1.5 + ((((3.14159 + 0.0) / v10) * -2.25) / 1.5)))) && true) {
                            if true {
                                fmt.Println(v7, v5)
                            }
                        } else {
                            delete(v3, 32767)
                            if true {
                                v3[(fi11 >> ((0) & 63))] = fi11
                                var v15 uint8 = 0
                                if (!true) {
                                    v16 := GenS1{F0: ("a" + "hello")}
                                    _ = v16
                                }
                                _ = v15
                            } else {
                                v5 = int8(fi11)
                                v6 = v6
                            }
                        }
                        v7 += (v7 + v7)
                        _ = fi14
                    }
                case "":
                    v17 := GenS1{F0: ("hello" + "a")}
                    _ = v17
                default:
                    if (0 >= (-(v5))) {
                        v18 := GenS2{F0: "hello"}
                        v5 -= v5
                        _ = v18
                    }
                }
            }
            var v19 int8 = 102
            var v20 string = (("hello" + "foo bar") + "")
            _ = fi11
            _ = v19
            _ = v20
        }
    case 127:
        for fi21 := 0; fi21 < 17; fi21++ {
            if ((3.14159 + (-2.25 - 1.5)) == (float32(v7) * 1e10)) {
                var v22 int8 = v5
                var v23 float32 = ((float32(-2.25) - 0.0) + 3.14159)
                v24 := GenS2{F0: ((("foo bar" + "") + "hello") + ("a" + "\n\t"))}
                _ = v22
                _ = v23
                _ = v24
            } else {
                fmt.Println(v5, fi21, v7)
            }
            var v25 string = "foo bar"
            switch fi21 {
            case 4294967295:
                _ = len(v1)
            default:
                var v26 int32 = v6
                _ = v26
            }
            _ = fi21
            _ = v25
        }
    case 0:
        v10 -= 1e10
    default:
        var v27 string = "a"
        _ = v27
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        fmt.Println(v7, v5, v5)
        if len(v1) > 0 {
            v1[0] = 552
        }
        v6 += (v6 * v6)
        _ = fi28
    }
    v7 = (0.1 - 0.1)
    fmt.Println(v5, v7)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v29 := GenS1{F0: "\n\t"}
    v3[4294967295] = 199
    v30 := GenS1{F0: "\n\t"}
    fmt.Println(v7, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
