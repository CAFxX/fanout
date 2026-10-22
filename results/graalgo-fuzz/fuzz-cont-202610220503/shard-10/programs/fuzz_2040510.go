package main

import (
    "fmt"
)


func main() {
    var v1 int16 = 1000
    var v2 uint8 = 246
    v3 := []string{("foo bar" + "\n\t"), ("a" + "foo bar")}
    v4 := make(map[int]int)
    var v5 int8 = 9
    v6 := make(map[int]bool)
    switch -9223372036854775808 {
    case -648:
        v4[int(v1)] = -1
    default:
        v5 -= v5
    }
    var v7 uint32 = 4294967295
    fmt.Println(v1)
    var v8 uint8 = 128
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 bool = false
        switch ("a" + (("foo bar" + "\n\t") + "a")) {
        case "a":
            var v11 int = 9223372036854775807
            _ = v11
        }
        for fi12 := 0; fi12 < 20; fi12++ {
            switch 32768 {
            case 1:
                v3 = append(v3, "hello")
            default:
                v2 = 0
            }
            _ = fi12
        }
        _ = fi9
        _ = v10
    }
    var v13 bool = (float64(((-2.25 * -2.25) + 1.5)) != (0.0 / ((1e10 - -2.25) + 1.5)))
    var v14 string = "a"
    v13 = false
    var v15 int = 32767
    v8 = v8
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 int32 = -740
        if ((!(!v13)) || false) {
            var v18 uint8 = (^(v8))
            var v19 uint = 2
            _ = v18
            _ = v19
        } else {
            for fi20 := 0; fi20 < 5; fi20++ {
                var v21 string = ((((v14 + "\n\t") + (((v14 + (v14 + "")) + "hello") + v14)) + ((v14 + (v14 + ((v14 + ((("a" + v14) + v14) + (("foo bar" + v14) + v14))) + ""))) + v14)) + v14)
                _ = fi20
                _ = v21
            }
            var v22 uint64 = 18446744073709551615
            var v23 uint16 = 2
            _ = v22
            _ = v23
        }
        if v13 {
            var v24 float64 = 0.1
            _ = v24
        }
        _ = fi16
        _ = v17
    }
    switch (v15 & 2147483647) {
    case 1:
        _ = len(v3)
    case 514:
        v2 += v8
    default:
        v14 = v14
    }
    if len(v3) > 0 {
        v3[0] = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
