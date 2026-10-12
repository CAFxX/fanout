package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int }
func genFn0(p0 int32) uint16 {
    for fi4 := 0; fi4 < 5; fi4++ {
        switch "a" {
        case "\n\t":
            p0 += (p0 % ((p0) | 1))
        case "":
            fmt.Println(p0, fi4)
        }
        _ = fi4
    }
    for fi5 := 0; fi5 < 1; fi5++ {
        p0 -= p0
        var v6 float64 = (-2.25 - 1e10)
        _ = fi5
        _ = v6
    }
    p0 += p0
    _ = p0
    return 52
}
func genFn1(p0 int) int16 {
    p0 = 32767
    var v7 float64 = ((((1.5 / (3.14159 / 1e10)) * (float64((1e10 / float32(0.1))) - (((((1.5 * float64(1.5)) / (3.14159 - ((-2.25 - float64((-2.25 - 3.14159))) * ((((0.0 - (0.0 * float64((0.1 + ((0.0 + ((1e10 * (3.14159 + 3.14159)) / float32(0.1))) * ((0.0 - ((-2.25 - (0.0 * 1e10)) / 1.0)) - 1.5)))))) * ((0.1 * 1e10) * 1e10)) / 0.1) - (0.1 + 0.0))))) + 0.1) + float64((1e10 + 0.1))) / 1e10))) + -2.25) + 0.1)
    v8 := GenS1{F0: -1, F1: 1, F2: (p0 | (p0 << ((619) & 63)))}
    _ = p0
    _ = v7
    _ = v8
    return -1
}

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 2
    v3 := make(map[int]string)
    v2 = v2
    for fi9 := 0; fi9 < 3; fi9++ {
        v2 = 2147483648
        _ = fi9
    }
    v10 := []int{127}
    switch 32767 {
    case 0:
        delete(v1, (("" + "\n\t") + ((("foo bar" + ("foo bar" + "")) + (("\n\t" + "a") + "hello")) + "\n\t")))
    case 2147483647:
        var v11 string = "a"
        _ = v11
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        if ((1e10 + (-2.25 / 1.5)) >= 1e10) {
            switch (fi12 * fi12) {
            case 32767:
                if true {
                    if ((!(true && ("" < ""))) || false) {
                        for fi13 := 0; fi13 < 12; fi13++ {
                            switch fi12 {
                            case -9223372036854775808:
                                var v14 bool = (0 >= 239)
                                _ = v14
                            case 989:
                                v2 = v2
                            case 881:
                                _ = len(v1)
                            default:
                                v2 -= v2
                            }
                            v2 -= v2
                            _ = fi13
                        }
                    }
                } else {
                    v15 := GenS1{F0: 1, F1: 945, F2: (fi12 ^ (fi12 >> ((42) & 63)))}
                    fmt.Println(v2)
                    fmt.Println(fi12)
                    _ = v15
                }
            default:
                var v16 bool = (768 <= 42)
                _ = v16
            }
            var v17 string = ("\n\t" + ("\n\t" + ("hello" + (("hello" + ("\n\t" + "")) + "foo bar"))))
            _ = v17
        }
        _ = fi12
    }
    if true {
        switch 32768 {
        case -2147483648:
            v2 = 65535
        default:
            fmt.Println(v2)
        }
    } else {
        for fi18 := 0; fi18 < 14; fi18++ {
            var v19 uint = 661
            _ = fi18
            _ = v19
        }
    }
    switch "a" {
    case "hello":
        var v20 uint = 65535
        _ = v20
    }
    v21 := v1["hello"]
    switch (("a" + "hello") + ("foo bar" + ("foo bar" + ("foo bar" + "hello")))) {
    case "hello":
        var v22 float32 = 3.14159
        _ = v22
    case "foo bar":
        switch ("\n\t" + "hello") {
        case "":
            switch "hello" {
            case "foo bar":
                switch v21 {
                case 4294967295:
                    v2 += v2
                }
            case "a":
                fmt.Println(v2, v21)
            case "hello":
                var v23 uint16 = 975
                _ = v23
            }
        default:
            v2 = 581
        }
    case "a":
        switch v21 {
        case 9223372036854775807:
            switch ("\n\t" + (("foo bar" + "\n\t") + "")) {
            case "\n\t":
                var v24 uint8 = 255
                _ = v24
            default:
                _ = len(v1)
            }
        }
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        if ((!false) || ((float32(3.14159) - 1.5) > ((-2.25 - float32((((((0.0 + 1e10) - (((-2.25 + (((float64(-2.25) - 1.5) * (-2.25 - 1e10)) * 1.5)) * (0.0 + 3.14159)) + (-2.25 * -2.25))) / 1e10) - ((3.14159 * 1e10) / (-2.25 - 1e10))) - 1e10))) - 0.1))) {
            for fi26 := 0; fi26 < 15; fi26++ {
                var v27 float64 = -2.25
                var v28 int64 = int64(fi25)
                _ = fi26
                _ = v27
                _ = v28
            }
        } else {
            v21 += v21
        }
        _ = fi25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
}
