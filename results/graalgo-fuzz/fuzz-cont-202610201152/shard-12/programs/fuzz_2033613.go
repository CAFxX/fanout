package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 int16 = 1000
    var v2 uint64 = 1
    v3 := []int{127}
    var v4 float64 = float64(-2.25)
    v5 := make(map[string]string)
    v2 = (v2 >> ((436) & 63))
    var v6 int64 = 0
    v7 := GenS1{F0: 83}
    v4 = v4
    for fi8 := 0; fi8 < 3; fi8++ {
        v5[("" + "foo bar")] = "a"
        _ = fi8
    }
    switch 32768 {
    case 2147483647:
        if (!(302 < 2147483648)) {
            for fi9 := 0; fi9 < 4; fi9++ {
                for fi10 := 0; fi10 < 7; fi10++ {
                    fmt.Println(fi9, v6)
                    fmt.Println(v1)
                    v2 += v2
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi10
                }
                fmt.Println(fi9)
                v11 := v5["foo bar"]
                _ = fi9
                _ = v11
            }
            for fi12 := 0; fi12 < 14; fi12++ {
                for fi13 := 0; fi13 < 2; fi13++ {
                    v5["a"] = "\n\t"
                    var v14 int32 = 2147483647
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(fi13)
                    case "hello":
                        v3 = append(v3, (fi13 >> ((4294967295) & 63)))
                    default:
                        v14 -= v14
                    }
                    _ = fi13
                    _ = v14
                }
                _ = fi12
            }
        }
    case 32767:
        var v15 int8 = -128
        _ = v15
    case 9223372036854775807:
        switch ("\n\t" + "a") {
        case "hello":
            fmt.Println(v2, v4)
        case "a":
            fmt.Println(v1, v6)
        }
    default:
        v1 += (-(v1))
    }
    v1 = (-(v1))
    var v16 float32 = (3.14159 * (3.14159 + 1e10))
    _ = len(v3)
    for fi17 := 0; fi17 < 16; fi17++ {
        v16 = (((1.5 * float32(((v4 / 0.1) + float64(((v16 * 0.1) * 1e10))))) + v16) * ((v16 - 0.1) - -2.25))
        _ = fi17
    }
    if false {
        fmt.Println(v6)
    } else {
        var v18 float32 = v16
        switch 32768 {
        case 42:
            switch "a" {
            case "\n\t":
                var v19 uint16 = 32767
                _ = v19
            case "foo bar":
                var v20 uint16 = 42
                _ = v20
            case "":
                var v21 float64 = v4
                _ = v21
            }
        }
        _ = v18
    }
    v2 = 2
    for fi22 := 0; fi22 < 10; fi22++ {
        if (-804 > v1) {
            switch -2147483648 {
            case -2147483648:
                var v23 int8 = -1
                _ = v23
            case -436:
                if true {
                    var v24 bool = (905 > 1)
                    var v25 bool = (2147483647 < fi22)
                    _ = v24
                    _ = v25
                }
            default:
                switch "foo bar" {
                case "a":
                    for fi26 := 0; fi26 < 13; fi26++ {
                        v6 = v6
                        v3 = append(v3, fi26)
                        _ = fi26
                    }
                case "foo bar":
                    fmt.Println(v4)
                case "":
                    if ((v4 <= v4) || true) {
                        delete(v5, "hello")
                    }
                default:
                    for fi27 := 0; fi27 < 8; fi27++ {
                        v28 := v5[""]
                        _ = fi27
                        _ = v28
                    }
                }
            }
            v16 -= v16
            var v29 bool = true
            _ = v29
        }
        fmt.Println(fi22, v6, v2)
        v4 = 0.0
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
