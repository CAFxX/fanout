package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 int16; F1 int64 }
func genFn0(p0 int64, p1 bool, p2 int64) int8 {
    var v5 uint32 = 65535
    if true {
        p0 = p0
        v5 -= 2147483648
        v6 := GenS1{F0: 1, F1: v5, F2: 16}
        _ = v6
    }
    v7 := []int{2147483647, 2147483647, 4294967295}
    v8 := GenS1{F0: 1000, F1: (v5 % (((v5 - (v5 >> ((104) & 63)))) | 1)), F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint64) bool {
    var v9 uint16 = 478
    fmt.Println(p0, p0)
    _ = p0
    _ = v9
    return (true || (true || false))
}

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = -128
    v3 := make(map[string]uint32)
    var v4 float64 = 1e10
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 int64 = 1
        fmt.Println(fi10, v2, fi10)
        if (!(18446744073709551615 >= 783)) {
            fmt.Println(v11, fi10, fi10)
            fmt.Println(v11)
            var v12 int32 = -397
            _ = v12
        } else {
            v13 := GenS1{F0: 1000, F1: 0, F2: 2}
            switch -1 {
            case 559:
                var v14 int32 = 1
                _ = v14
            case -214:
                fmt.Println(fi10, fi10)
            default:
                v1[("hello" + "")] = (!(true || false))
            }
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    if true {
        v15 := []int{-1}
        if (2 < 128) {
            var v16 int = 32768
            if (!true) {
                v17 := v1["a"]
                for fi18 := 0; fi18 < 4; fi18++ {
                    if len(v15) > 0 {
                        v15[0] = (v16 << ((4294967295) & 63))
                    }
                    v4 = ((v4 + 1e10) - v4)
                    _ = fi18
                }
                for fi19 := 0; fi19 < 10; fi19++ {
                    for fi20 := 0; fi20 < 13; fi20++ {
                        v4 -= ((float64(1.5) - v4) - v4)
                        if v17 {
                            v21 := GenS2{F0: 272}
                            fmt.Println(v4, fi19, fi20)
                            v22 := GenS1{F0: 0, F1: 2, F2: 1}
                            _ = v21
                            _ = v22
                        } else {
                            delete(v3, "foo bar")
                        }
                        _ = fi20
                    }
                    _ = fi19
                }
                _ = v17
            } else {
                switch ("hello" + "hello") {
                case "hello":
                    delete(v3, ("foo bar" + ""))
                case "\n\t":
                    switch v16 {
                    case -9223372036854775808:
                        var v23 float32 = 3.14159
                        _ = v23
                    case -2147483648:
                        var v24 int8 = 42
                        _ = v24
                    default:
                        var v25 int = (v16 + v16)
                        _ = v25
                    }
                }
            }
            var v26 uint8 = 0
            _ = v16
            _ = v26
        }
        _ = v15
    } else {
        v2 = (v2 >> ((1) & 63))
        v27 := v1[""]
        v27 = v27
        _ = v27
    }
    if false {
        v2 = v2
    } else {
        if false {
            v4 -= (v4 * (v4 * ((1.5 / (0.0 / v4)) + (3.14159 * float64(3.14159)))))
        }
    }
    v28 := GenS3{F0: 0, F1: 1000}
    v29 := []int{-9223372036854775808, 127, 127}
    v30 := GenS1{F0: -32768, F1: 1, F2: 18446744073709551615}
    v4 = (((0.1 * (-2.25 + ((((v4 / v4) + 0.1) * v4) / v4))) / v4) / v4)
    fmt.Println(v4, v4)
    if len(v29) >= 2 {
        v29 = v29[1:len(v29)]
    }
    switch "" {
    case "foo bar":
        for fi31 := 0; fi31 < 5; fi31++ {
            v29 = append(v29, -2147483648)
            if len(v29) > 0 {
                v29[0] = (fi31 >> ((4294967295) & 63))
            }
            if false {
                switch "" {
                case "\n\t":
                    var v32 uint64 = 0
                    _ = v32
                }
            } else {
                switch "\n\t" {
                case "\n\t":
                    for fi33 := 0; fi33 < 11; fi33++ {
                        fmt.Println(fi33)
                        v4 = ((v4 + ((float64((1.5 / 0.1)) + v4) * 1.5)) * (v4 - (0.0 / ((3.14159 * float64(float32(1e10))) / v4))))
                        var v34 int8 = (v2 / ((v2) | 1))
                        _ = fi33
                        _ = v34
                    }
                case "hello":
                    if true {
                        _ = len(v29)
                        v35 := GenS2{F0: 626}
                        _ = v35
                    }
                case "a":
                    v36 := v1[(("a" + ("a" + "\n\t")) + ("" + ("foo bar" + (("hello" + ("a" + "\n\t")) + ("a" + ("a" + ("hello" + ("hello" + (("\n\t" + "") + "\n\t")))))))))]
                    _ = v36
                }
                v2 = 100
                if false {
                    v2 += v2
                    v4 -= 1.5
                }
            }
            _ = fi31
        }
    case "a":
        v4 += v4
    }
    var v37 int8 = (v2 * (-(v2)))
    v2 -= v2
    var v38 uint32 = 4294967295
    v4 -= (3.14159 * v4)
    v39 := GenS4{F0: -17, F1: 9223372036854775807}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
