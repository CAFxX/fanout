package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint32 }
func genFn0(p0 uint16, p1 uint8) int32 {
    for fi5 := 0; fi5 < 8; fi5++ {
        switch fi5 {
        case -1:
            for fi6 := 0; fi6 < 11; fi6++ {
                v7 := GenS1{F0: uint32(p1), F1: 701, F2: 127}
                p1 -= p1
                _ = fi6
                _ = v7
            }
        default:
            fmt.Println(fi5, p0)
        }
        var v8 int32 = 2147483647
        _ = fi5
        _ = v8
    }
    v9 := []int{128, 329, 755, 32768}
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 3; fi11++ {
            fmt.Println(fi10, fi10, p0)
            var v12 float32 = ((3.14159 - 0.1) + (((1e10 + 3.14159) / 0.1) + 1.5))
            if len(v9) > 0 {
                v9[0] = (fi11 >> ((967) & 63))
            }
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v9
    return -1
}

func main() {
    var v1 int64 = 42
    v2 := make(map[string]uint32)
    v3 := make(map[string]string)
    v4 := []string{"", ("hello" + "foo bar")}
    switch "foo bar" {
    case "hello":
        v1 -= (v1 % (((v1 | v1)) | 1))
    case "":
        if ((!true) && true) {
            v13 := GenS2{F0: v1, F1: 291}
            fmt.Println(v1)
            v1 = -1
            _ = v13
        }
    case "\n\t":
        v14 := v3["hello"]
        _ = v14
    }
    v2["\n\t"] = 2147483647
    v15 := GenS1{F0: 1, F1: 220, F2: 1}
    var v16 float32 = float32(1.5)
    fmt.Println(v1, v16)
    var v17 string = "hello"
    if true {
        switch v17 {
        case "hello":
            switch -2147483648 {
            case 32767:
                fmt.Println(v16, v1, v16)
            default:
                if false {
                    v18 := GenS1{F0: 2147483647, F1: 0, F2: 127}
                    switch "\n\t" {
                    case "":
                        for fi19 := 0; fi19 < 9; fi19++ {
                            switch (fi19 | fi19) {
                            case -628:
                                var v20 uint16 = 255
                                _ = v20
                            case 9223372036854775807:
                                var v21 bool = (!(v16 > (((v16 + float32(1.5)) + (v16 - v16)) - v16)))
                                _ = v21
                            }
                            var v22 uint32 = 2147483648
                            _ = fi19
                            _ = v22
                        }
                    }
                    switch v17 {
                    case "foo bar":
                        v23 := v3[v17]
                        _ = v23
                    case "a":
                        var v24 int64 = (v1 & (v1 % ((42) | 1)))
                        _ = v24
                    }
                    _ = v18
                }
            }
        case "a":
            for fi25 := 0; fi25 < 11; fi25++ {
                v16 = -2.25
                _ = fi25
            }
        default:
            for fi26 := 0; fi26 < 8; fi26++ {
                var v27 int32 = 1
                fmt.Println(v1)
                switch v17 {
                case "foo bar":
                    switch (v17 + ("foo bar" + (("\n\t" + (v17 + "")) + v17))) {
                    case "":
                        v27 = (v27 + v27)
                    case "hello":
                        v1 += (v1 / ((9223372036854775807) | 1))
                    default:
                        v27 -= v27
                    }
                case "":
                    switch ((v17 + v17) + "foo bar") {
                    case "foo bar":
                        v17 = "a"
                    case "\n\t":
                        v28 := GenS3{F0: 2, F1: 1, F2: 1}
                        _ = v28
                    case "":
                        v1 = v1
                    default:
                        if true {
                            v2["foo bar"] = 2147483648
                        }
                    }
                case "hello":
                    for fi29 := 0; fi29 < 14; fi29++ {
                        v16 += (float32(float64(((1e10 / v16) - (v16 + (0.0 * 3.14159))))) + v16)
                        var v30 uint16 = 697
                        _ = fi29
                        _ = v30
                    }
                }
                _ = fi26
                _ = v27
            }
        }
        switch 805 {
        case 4294967295:
            for fi31 := 0; fi31 < 5; fi31++ {
                var v32 float64 = float64(1e10)
                if true {
                    if (42 == -702) {
                        var v33 float64 = 1.5
                        v16 = (v16 - 3.14159)
                        if false {
                            v33 += 3.14159
                        }
                        _ = v33
                    }
                    v32 = 0.1
                }
                v1 = (v1 - v1)
                _ = fi31
                _ = v32
            }
        default:
            if (22 >= 901) {
                v16 -= (1.5 + v16)
            }
        }
    } else {
        v1 = (^(v1))
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi34 := 0; fi34 < 10; fi34++ {
        var v35 string = ((("" + ("hello" + "\n\t")) + "a") + "foo bar")
        _ = fi34
        _ = v35
    }
    v17 = "hello"
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        if false {
            for fi36 := 0; fi36 < 18; fi36++ {
                v17 = (("" + v17) + "")
                _ = fi36
            }
        } else {
            v16 -= 0.1
            v16 = v16
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
