package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint) int8 {
    for fi4 := 0; fi4 < 11; fi4++ {
        p0 = (p0 << ((p0) & 63))
        v5 := []int{fi4, -620, 32767}
        var v6 int16 = -933
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 uint32 = 1
    p0 += p0
    p0 += (p0 ^ 255)
    _ = p0
    _ = v7
    return -88
}

func main() {
    var v1 int64 = -1
    v2 := []string{((("a" + "\n\t") + "") + ""), "\n\t", "\n\t"}
    v3 := []string{"foo bar"}
    var v8 uint = 18446744073709551615
    fmt.Println(v8)
    fmt.Println(v1, v1)
    v2 = append(v2, ("hello" + "a"))
    for fi9 := 0; fi9 < 4; fi9++ {
        if len(v2) > 0 {
            v2[0] = "a"
        }
        if len(v2) > 0 {
            v2[0] = (("" + "foo bar") + "\n\t")
        }
        _ = fi9
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (!true) {
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 uint8 = 128
            if (18446744073709551615 == 520) {
                for fi12 := 0; fi12 < 16; fi12++ {
                    var v13 int = -86
                    _ = fi12
                    _ = v13
                }
                for fi14 := 0; fi14 < 15; fi14++ {
                    if false {
                        if ("" > "") {
                            switch "a" {
                            case "a":
                                if len(v3) > 0 {
                                    v3[0] = "foo bar"
                                }
                            }
                            v8 = 2
                        }
                    }
                    fmt.Println(v8)
                    _ = fi14
                }
            }
            _ = fi10
            _ = v11
        }
        v3 = append(v3, "\n\t")
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if false {
        switch "hello" {
        case "hello":
            v1 += v1
        case "a":
            if false {
                for fi15 := 0; fi15 < 1; fi15++ {
                    fmt.Println(v8)
                    if len(v3) > 0 {
                        v3[0] = ""
                    }
                    _ = fi15
                }
            }
        }
        v1 += v1
    }
    v16 := GenS1{F0: 4294967295}
    var v17 uint = 2
    v2 = append(v2, (("hello" + "foo bar") + "\n\t"))
    switch 32768 {
    case 127:
        switch "hello" {
        case "a":
            v17 = (v17 << (((^(v17))) & 63))
        case "\n\t":
            switch -961 {
            case -245:
                v3 = append(v3, "\n\t")
            case -9223372036854775808:
                var v18 bool = ((!false) || ((0.1 / ((0.1 * ((3.14159 - float64(0.1)) * ((1e10 / ((1.5 - float64(0.0)) + (1.5 - 1e10))) - (1.5 * 0.0)))) + (((3.14159 - (1e10 + (0.0 * 3.14159))) / 1.0) * ((1.5 * ((-2.25 * (float64(0.0) + 0.1)) / -2.25)) + (0.1 / ((1.5 * 0.1) - 0.0)))))) != (1e10 / 0.1)))
                _ = v18
            }
        case "foo bar":
            switch 525 {
            case 778:
                for fi19 := 0; fi19 < 10; fi19++ {
                    if len(v2) > 0 {
                        v2[0] = "\n\t"
                    }
                    _ = fi19
                }
            case -1:
                fmt.Println(v1, v8, v8)
            case 573:
                var v20 uint64 = 2
                _ = v20
            default:
                for fi21 := 0; fi21 < 15; fi21++ {
                    for fi22 := 0; fi22 < 16; fi22++ {
                        fmt.Println(v1, fi21)
                        var v23 uint = (^(v17))
                        v23 = v17
                        _ = fi22
                        _ = v23
                    }
                    _ = fi21
                }
            }
        default:
            var v24 float64 = (0.0 - 1.5)
            _ = v24
        }
    case -119:
        v17 += v17
    default:
        for fi25 := 0; fi25 < 6; fi25++ {
            fmt.Println(v8, fi25, v1)
            var v26 uint = (v17 / (((v17 - 1)) | 1))
            _ = fi25
            _ = v26
        }
    }
    switch 32767 {
    case 1:
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
    case 128:
        var v27 float32 = 0.1
        _ = v27
    case 42:
        v28 := GenS1{F0: 0}
        _ = v28
    default:
        fmt.Println(v1, v8, v17)
    }
    v8 = (v17 - v8)
    fmt.Println(v17, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v8
}
