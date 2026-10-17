package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 bool }
type GenS2 struct { F0 int64 }
func genFn0() int8 {
    v6 := []int{0, 88, -479}
    var v7 uint8 = 70
    fmt.Println(v7)
    var v8 int64 = 0
    switch -137 {
    case 42:
        if len(v6) > 0 {
            v6[0] = 127
        }
    case 509:
        switch -179 {
        case 696:
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 uint32 = 4294967295
                _ = fi9
                _ = v10
            }
        default:
            v6 = append(v6, -9223372036854775808)
        }
    case 128:
        v11 := GenS1{F0: 1, F1: "", F2: false}
        _ = v11
    default:
        var v12 int = 2147483647
        _ = v12
    }
    _ = v6
    _ = v7
    _ = v8
    return 127
}

func main() {
    var v1 bool = true
    v2 := make(map[int]uint32)
    var v3 float64 = (0.0 / 1.5)
    var v4 float32 = -2.25
    v5 := make(map[string]string)
    switch 9223372036854775807 {
    case 2147483647:
        fmt.Println(v4, v4)
    default:
        v4 -= ((v4 * -2.25) - (float32((1.5 * v3)) * v4))
    }
    var v13 uint = 2
    for fi14 := 0; fi14 < 20; fi14++ {
        var v15 int = (fi14 - 2147483647)
        _ = fi14
        _ = v15
    }
    var v16 uint16 = 42
    var v17 int8 = 127
    fmt.Println(v4)
    v18 := []int{1, 4294967295}
    for fi19 := 0; fi19 < 4; fi19++ {
        v1 = true
        _ = fi19
    }
    var v20 bool = (-941 != -2147483648)
    if false {
        switch ("\n\t" + "") {
        case "hello":
            switch "a" {
            case "hello":
                v21 := GenS1{F0: 0, F1: ("foo bar" + ("foo bar" + "\n\t")), F2: false}
                _ = v21
            case "a":
                for fi22 := 0; fi22 < 7; fi22++ {
                    v4 -= v4
                    _ = fi22
                }
            case "\n\t":
                v17 = v17
            }
        }
        v23 := GenS1{F0: 185, F1: ("foo bar" + ("foo bar" + ((("foo bar" + (("hello" + "hello") + "")) + "") + (("a" + (("a" + "") + "hello")) + ("foo bar" + ("hello" + "hello")))))), F2: v20}
        v4 += (float32(v3) / (((((0.1 - 1.5) - 0.1) / float32(v3)) / (((v4 / float32((v3 / v3))) * v4) - ((3.14159 * (3.14159 / (v4 + v4))) + 3.14159))) + v4))
        _ = v23
    }
    for fi24 := 0; fi24 < 17; fi24++ {
        v4 = 1e10
        _ = fi24
    }
    v25 := GenS2{F0: 401}
    if (!true) {
        for fi26 := 0; fi26 < 3; fi26++ {
            v20 = ((("\n\t" + (((("" + "\n\t") + "") + "foo bar") + ("hello" + "foo bar"))) + ((("\n\t" + "a") + "") + "hello")) == "")
            for fi27 := 0; fi27 < 20; fi27++ {
                switch "" {
                case "hello":
                    switch "hello" {
                    case "":
                        switch "a" {
                        case "\n\t":
                            var v28 uint16 = (v16 >> (((v13 - v13)) & 63))
                            _ = v28
                        default:
                            v13 += 554
                        }
                    case "hello":
                        if true {
                            if len(v18) > 0 {
                                v18[0] = (fi27 << ((2) & 63))
                            }
                            v29 := GenS1{F0: 2, F1: ("\n\t" + (("foo bar" + "hello") + "a")), F2: ((!(0 <= -1)) && (v20 && v20))}
                            v3 += (((3.14159 / (v3 * v3)) / v3) + (-2.25 - v3))
                            var v30 float64 = (1.5 + v3)
                            _ = v29
                            _ = v30
                        } else {
                            var v31 int16 = 0
                            if len(v18) > 0 {
                                v18[0] = -107
                            }
                            _ = v31
                        }
                    }
                }
                if v1 {
                    fmt.Println(v17, v13)
                    var v32 bool = v20
                    _ = v32
                }
                if v20 {
                    v33 := GenS2{F0: 0}
                    _ = len(v2)
                    delete(v2, fi27)
                    _ = v33
                }
                _ = fi27
            }
            v13 -= 2
            _ = fi26
        }
    }
    delete(v2, 9223372036854775807)
    v34 := GenS1{F0: 127, F1: "foo bar", F2: v20}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
