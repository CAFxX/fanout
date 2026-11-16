package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint; F1 bool; F2 int64 }
func genFn0(p0 string) uint {
    for fi5 := 0; fi5 < 12; fi5++ {
        p0 = p0
        _ = fi5
    }
    p0 = ((p0 + p0) + p0)
    v6 := GenS1{F0: (("a" + p0) != ("a" + ((("foo bar" + p0) + (("a" + (p0 + (((p0 + ("a" + "foo bar")) + p0) + (p0 + p0)))) + p0)) + ("a" + (p0 + (p0 + p0)))))), F1: 27}
    p0 = p0
    var v7 int32 = -1
    _ = p0
    _ = v6
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 uint8 = 177
    var v2 uint32 = 1
    v3 := make(map[int]string)
    var v4 int = -988
    v8 := GenS1{F0: ((1e10 >= ((-2.25 - -2.25) - 3.14159)) || (2 >= 2)), F1: 5}
    if false {
        switch v4 {
        case -824:
            v9 := GenS1{F0: true, F1: 0}
            _ = v9
        case 596:
            switch 32768 {
            case -2147483648:
                for fi10 := 0; fi10 < 10; fi10++ {
                    v11 := []int{fi10, 127}
                    _ = fi10
                    _ = v11
                }
            default:
                for fi12 := 0; fi12 < 7; fi12++ {
                    v2 += (v2 >> ((65535) & 63))
                    v2 = (^(v2))
                    v1 = (v1 / ((0) | 1))
                    _ = fi12
                }
            }
        }
    }
    v13 := GenS1{F0: true, F1: 2}
    if (65535 >= 18446744073709551615) {
        switch v4 {
        case 42:
            v2 = (v2 % ((v2) | 1))
        case 128:
            for fi14 := 0; fi14 < 7; fi14++ {
                var v15 int = (v4 - fi14)
                fmt.Println(v15, fi14, v1)
                _ = fi14
                _ = v15
            }
        }
        switch (v4 << ((uint(v2)) & 63)) {
        case 319:
            v1 -= (v1 * v1)
        case 2147483647:
            fmt.Println(v2, v2, v1)
        }
        var v16 uint64 = 2
        _ = v16
    } else {
        var v17 int16 = 371
        v17 = 1000
        var v18 int8 = -128
        _ = v17
        _ = v18
    }
    v2 -= (v2 | (^(v2)))
    fmt.Println(v4)
    v19 := []int{v4, -46}
    if (255 > 111) {
        var v20 uint16 = 1
        var v21 float64 = float64(1e10)
        _ = v20
        _ = v21
    }
    switch "hello" {
    case "a":
        v22 := GenS2{F0: 4294967295, F1: ("" < ("\n\t" + "a")), F2: 42}
        _ = v22
    }
    var v23 int64 = 0
    if len(v19) > 0 {
        v19[0] = (v4 % ((v4) | 1))
    }
    v2 = (v2 >> ((18446744073709551615) & 63))
    v24 := GenS2{F0: 18446744073709551615, F1: (438 == 0), F2: -9223372036854775808}
    fmt.Println(v4, v1, v23)
    v2 += (^(v2))
    fmt.Println(v1, v4)
    switch ((("a" + "foo bar") + ("a" + ("foo bar" + "\n\t"))) + ("hello" + "foo bar")) {
    case "foo bar":
        v19 = append(v19, 0)
    case "a":
        var v25 int16 = 1000
        _ = v25
    }
    var v26 float32 = (1.5 + float32(3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v8
}
