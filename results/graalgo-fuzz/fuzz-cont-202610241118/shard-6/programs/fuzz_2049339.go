package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 bool) uint16 {
    var v7 float32 = 0.1
    fmt.Println(v7, v7)
    switch -183 {
    case -2147483648:
        if (9 != 0) {
            v8 := []int{32768}
            var v9 string = "a"
            _ = v8
            _ = v9
        }
    case 9223372036854775807:
        fmt.Println(v7, p0)
    case 2147483647:
        var v10 string = "foo bar"
        _ = v10
    default:
        v11 := []int{}
        _ = v11
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        v7 = v7
        _ = fi12
    }
    _ = p0
    _ = v7
    return 248
}

func main() {
    var v1 string = ("hello" + "hello")
    var v2 string = v1
    var v3 float64 = (0.0 * 0.0)
    v4 := []string{v1, ((v2 + v1) + ("a" + ("foo bar" + v1)))}
    v5 := make(map[string]string)
    v6 := make(map[uint32]string)
    v1 = v1
    switch 32768 {
    case 920:
        var v13 int = 127
        _ = v13
    case 128:
        fmt.Println(v3, v1)
    case -90:
        fmt.Println(v3, v1, v2)
    }
    v1 = "foo bar"
    for fi14 := 0; fi14 < 4; fi14++ {
        fmt.Println(v2, v2, v3)
        _ = fi14
    }
    var v15 uint64 = 42
    v1 = ((("a" + ("" + ((v1 + "\n\t") + v2))) + ("hello" + "\n\t")) + "hello")
    if (!((((v1 + ((v2 + "\n\t") + "\n\t")) + "") + v1) > v1)) {
        v2 = ("a" + "hello")
        v6[708] = v2
    } else {
        fmt.Println(v2, v1)
        v3 = v3
        for fi16 := 0; fi16 < 12; fi16++ {
            v4 = append(v4, v2)
            for fi17 := 0; fi17 < 11; fi17++ {
                var v18 string = v1
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
    }
    _ = len(v5)
    if (false && false) {
        switch "" {
        case "a":
            v19 := GenS1{F0: 100, F1: 42}
            _ = v19
        case "foo bar":
            if true {
                v15 = (v15 << ((1) & 63))
                switch v2 {
                case "":
                    if true {
                        v2 = "a"
                        var v20 int16 = 0
                        v15 = v15
                        _ = v20
                    }
                case "a":
                    v3 = (v3 + v3)
                case "hello":
                    v6[493] = "foo bar"
                }
                switch v1 {
                case "hello":
                    for fi21 := 0; fi21 < 2; fi21++ {
                        var v22 uint16 = 32767
                        switch 128 {
                        case 2147483647:
                            v6[4294967295] = v2
                        case 39:
                            var v23 uint64 = (v15 << ((4294967295) & 63))
                            _ = v23
                        }
                        _ = fi21
                        _ = v22
                    }
                case "":
                    v15 += v15
                default:
                    v24 := GenS2{F0: 875}
                    _ = v24
                }
            }
        case "\n\t":
            v3 += -2.25
        default:
            var v25 uint8 = 243
            _ = v25
        }
        v1 = v2
        fmt.Println(v2, v2)
    } else {
        _ = len(v5)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        if false {
            v3 = (0.1 - v3)
        } else {
            switch "foo bar" {
            case "foo bar":
                var v26 float32 = 0.0
                _ = v26
            }
            v6[874] = ("" + (v1 + ("\n\t" + v2)))
        }
        v15 = (v15 % ((23) | 1))
        fmt.Println(v2, v15)
    }
    var v27 int32 = 828
    var v28 uint = 358
    v15 = v15
    switch 32767 {
    case 2147483647:
        if false {
            var v29 uint = 542
            if true {
                var v30 float32 = 3.14159
                var v31 int = 0
                v5[(v2 + "")] = (((v2 + (v2 + v1)) + ("a" + "a")) + v1)
                _ = v30
                _ = v31
            }
            var v32 int64 = -874
            _ = v29
            _ = v32
        } else {
            var v33 int64 = 9223372036854775807
            var v34 int32 = v27
            if (0 == 255) {
                var v35 uint = 2
                v1 = v2
                var v36 uint = 4294967295
                _ = v35
                _ = v36
            } else {
                var v37 bool = (-105 == 1)
                for fi38 := 0; fi38 < 18; fi38++ {
                    v37 = ((v37 || (1e10 == -2.25)) && v37)
                    v37 = (!v37)
                    v4 = append(v4, v2)
                    _ = fi38
                }
                _ = v37
            }
            _ = v33
            _ = v34
        }
    default:
        v39 := GenS1{F0: 588, F1: 42}
        _ = v39
    }
    var v40 uint32 = 236
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
