package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint16 }
type GenS2 struct { F0 string }
func genFn0(p0 uint64) uint8 {
    if (false || (0 <= (p0 >> ((4294967295) & 63)))) {
        p0 += p0
        p0 = (p0 / (((p0 << ((255) & 63))) | 1))
    }
    if (!(!true)) {
        v5 := []int{-278, 32768}
        p0 += p0
        _ = v5
    } else {
        var v6 int16 = 1
        var v7 uint64 = 968
        _ = v6
        _ = v7
    }
    switch 0 {
    case -499:
        fmt.Println(p0)
    default:
        var v8 uint = 65535
        _ = v8
    }
    switch 2147483647 {
    case 2147483647:
        fmt.Println(p0)
    case -150:
        fmt.Println(p0)
    }
    if (-1 == 0) {
        var v9 uint = 2
        var v10 uint16 = 32767
        _ = v10
        _ = v9
    } else {
        v11 := []int{-1, -87, 42, -2147483648}
        if len(v11) > 0 {
            v11[0] = 127
        }
        _ = v11
    }
    _ = p0
    return 128
}

func main() {
    var v1 float64 = ((1.5 / (float64(-2.25) - 1.5)) / 0.1)
    v2 := []string{}
    v3 := make(map[string]uint32)
    var v4 int32 = -578
    if (2 >= 2) {
        for fi12 := 0; fi12 < 14; fi12++ {
            v13 := GenS1{F0: 0, F1: (fi12 % ((fi12) | 1)), F2: 65535}
            switch -283 {
            case -348:
                delete(v3, ("foo bar" + ("\n\t" + (("foo bar" + ("" + ("a" + "\n\t"))) + ("foo bar" + ("" + "\n\t"))))))
            }
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 6; fi14++ {
            if true {
                var v15 uint8 = 0
                _ = v15
            } else {
                v16 := GenS1{F0: 1, F1: fi14, F2: 555}
                switch ("foo bar" + ("foo bar" + "a")) {
                case "\n\t":
                    var v17 uint8 = 255
                    _ = v17
                case "a":
                    v2 = append(v2, "hello")
                }
                if false {
                    for fi18 := 0; fi18 < 7; fi18++ {
                        fmt.Println(v4, fi14, v1)
                        v1 = (float64(0.1) / (((float64(0.1) + v1) / float64(1e10)) * v1))
                        _ = fi18
                    }
                    for fi19 := 0; fi19 < 2; fi19++ {
                        var v20 uint8 = 128
                        _ = fi19
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 1; fi21++ {
                        v22 := GenS1{F0: 42, F1: 9223372036854775807, F2: 65535}
                        _ = fi21
                        _ = v22
                    }
                }
                _ = v16
            }
            v23 := GenS1{F0: 1, F1: (fi14 | -465), F2: 2}
            var v24 float32 = 1.5
            _ = fi14
            _ = v23
            _ = v24
        }
        v25 := v3["hello"]
        _ = v25
    }
    v4 = v4
    if (1 < -128) {
        v1 = (float64(1.5) / float64(3.14159))
        switch "" {
        case "":
            var v26 uint32 = 1
            _ = v26
        default:
            var v27 uint = 65535
            _ = v27
        }
        var v28 float32 = (1.5 * (0.0 - 1e10))
        _ = v28
    } else {
        if ("a" <= "\n\t") {
            for fi29 := 0; fi29 < 17; fi29++ {
                if true {
                    v1 = v1
                    v30 := GenS1{F0: 42, F1: (fi29 % ((127) | 1)), F2: 32}
                    _ = v30
                } else {
                    v3[""] = 1
                    v31 := GenS1{F0: 1, F1: (fi29 << ((65535) & 63)), F2: 1}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v31
                }
                _ = fi29
            }
        }
    }
    switch "" {
    case "\n\t":
        var v32 string = "foo bar"
        _ = v32
    case "foo bar":
        fmt.Println(v4)
    case "hello":
        _ = len(v2)
    default:
        var v33 float32 = 3.14159
        _ = v33
    }
    for fi34 := 0; fi34 < 12; fi34++ {
        var v35 int8 = 127
        v36 := GenS2{F0: ("foo bar" + "\n\t")}
        v35 -= v35
        _ = fi34
        _ = v35
        _ = v36
    }
    fmt.Println(v1, v1)
    v4 -= v4
    v1 -= (v1 * v1)
    _ = len(v3)
    var v37 uint16 = 0
    var v38 string = "hello"
    if true {
        var v39 uint16 = 42
        _ = v39
    }
    if true {
        v1 = ((0.0 + v1) * (1e10 - -2.25))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
}
