package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 bool }
type GenS2 struct { F0 int; F1 int8; F2 bool }
type GenS3 struct { F0 string }

func main() {
    var v1 float64 = -2.25
    var v2 uint8 = 42
    v3 := make(map[uint32]string)
    var v4 float32 = 0.1
    var v5 string = "foo bar"
    var v6 int = 4294967295
    var v7 int8 = 127
    v8 := GenS1{F0: 445, F1: 18446744073709551615, F2: false}
    v9 := GenS1{F0: 4294967295, F1: 4294967295, F2: (true && true)}
    v10 := []int{128, v6, 0}
    if ((v6 << ((255) & 63)) >= v6) {
        delete(v3, 42)
        fmt.Println(v7)
    } else {
        for fi11 := 0; fi11 < 14; fi11++ {
            if (186 < v2) {
                for fi12 := 0; fi12 < 12; fi12++ {
                    v2 -= v2
                    delete(v3, 65535)
                    _ = fi12
                }
                v13 := GenS1{F0: 510, F1: 570, F2: ((367 <= 65535) || true)}
                _ = len(v3)
                _ = v13
            }
            fmt.Println(v5, v5)
            v7 -= -1
            _ = fi11
        }
        if len(v10) > 0 {
            v10[0] = v6
        }
        v1 = 0.1
    }
    var v14 string = ("foo bar" + ("hello" + v5))
    switch (v14 + v5) {
    case "\n\t":
        delete(v3, 4294967295)
    }
    if (42 == 2) {
        switch ((((("" + "") + "hello") + (((((v5 + v14) + v5) + "a") + ("foo bar" + "a")) + v14)) + (("a" + "") + v14)) + (v14 + v14)) {
        case "foo bar":
            v6 -= v6
        case "\n\t":
            v14 = (v5 + v5)
        }
        v4 -= v4
    }
    if (!false) {
        for fi15 := 0; fi15 < 3; fi15++ {
            v7 = 100
            _ = fi15
        }
    }
    switch v5 {
    case "hello":
        for fi16 := 0; fi16 < 5; fi16++ {
            fmt.Println(v2)
            v4 -= 1.5
            var v17 int32 = -26
            _ = fi16
            _ = v17
        }
    case "\n\t":
        v18 := GenS1{F0: 42, F1: 4294967295, F2: (-2.25 >= v1)}
        _ = v18
    case "":
        switch (v6 ^ 32768) {
        case 2147483647:
            if (0 != 167) {
                v5 = ("" + v14)
                v19 := GenS1{F0: 1, F1: 1, F2: (false && ((v6 ^ (v6 | (v6 << ((18446744073709551615) & 63)))) > 9223372036854775807))}
                _ = v19
            }
        case -9223372036854775808:
            v20 := GenS1{F0: 2, F1: 4294967295, F2: (1 < -9223372036854775808)}
            _ = v20
        case 42:
            fmt.Println(v6)
        default:
            for fi21 := 0; fi21 < 18; fi21++ {
                for fi22 := 0; fi22 < 17; fi22++ {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    _ = fi22
                }
                switch "" {
                case "":
                    v23 := GenS1{F0: 18446744073709551615, F1: 18446744073709551615, F2: (!true)}
                    _ = v23
                }
                v14 = v5
                _ = fi21
            }
        }
    default:
        v24 := GenS1{F0: 2, F1: 2, F2: false}
        _ = v24
    }
    v7 -= v7
    switch "hello" {
    case "\n\t":
        var v25 uint = 0
        _ = v25
    default:
        v26 := GenS2{F0: (v6 << ((4294967295) & 63)), F1: v7, F2: (true || (!true))}
        _ = v26
    }
    if (0 <= 2) {
        v14 = "a"
        var v27 int = (v6 >> ((255) & 63))
        if true {
            v6 -= -824
            var v28 uint16 = 555
            v29 := GenS2{F0: (v6 | v27), F1: (v7 ^ -1), F2: (-2.25 <= v1)}
            _ = v28
            _ = v29
        } else {
            v3[111] = "a"
        }
        _ = v27
    } else {
        if true {
            v30 := GenS2{F0: (-(v6)), F1: (-(v7)), F2: (0 == -318)}
            var v31 uint = 65535
            _ = v30
            _ = v31
        } else {
            v32 := GenS3{F0: "hello"}
            _ = v32
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
