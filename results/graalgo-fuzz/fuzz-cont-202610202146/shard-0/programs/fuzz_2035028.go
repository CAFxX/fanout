package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 bool; F2 uint32 }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 uint; F1 string }
func genFn0(p0 int16, p1 int64, p2 int8) bool {
    var v5 bool = (65535 == 2)
    v6 := GenS1{F0: 695}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return ((v5 && (!(854 == 238))) || (!(true && (491 != 272))))
}

func main() {
    v1 := make(map[string]string)
    var v2 bool = (false || true)
    var v3 float64 = (0.0 / 3.14159)
    var v4 uint16 = 566
    if (32767 < 1000) {
        if (!v2) {
            delete(v1, "\n\t")
            switch -200 {
            case 32768:
                fmt.Println(v3, v4)
            case -924:
                _ = len(v1)
            }
            var v7 float64 = (-2.25 * (float64((-2.25 / 1.0)) - (float64(-2.25) * v3)))
            _ = v7
        } else {
            v8 := GenS2{F0: v4}
            _ = v8
        }
        fmt.Println(v3, v4)
    }
    v9 := GenS1{F0: -123}
    var v10 bool = v2
    for fi11 := 0; fi11 < 8; fi11++ {
        v10 = (158 == 18446744073709551615)
        _ = fi11
    }
    if false {
        for fi12 := 0; fi12 < 9; fi12++ {
            v3 = float64((((((1.5 / -2.25) - ((1e10 - (0.1 - 1.5)) + (3.14159 / 3.14159))) * 3.14159) - -2.25) + 1e10))
            var v13 uint32 = 4294967295
            switch (fi12 ^ 128) {
            case 32767:
                v14 := GenS3{F0: "a", F1: v2, F2: 2147483648}
                _ = v14
            }
            _ = fi12
            _ = v13
        }
        if (1 == 1000) {
            var v15 int = -9223372036854775808
            v2 = false
            v10 = (!(!(false && (!(false || (v15 != v15))))))
            _ = v15
        } else {
            if (1.5 == 0.0) {
                for fi16 := 0; fi16 < 11; fi16++ {
                    var v17 float64 = 3.14159
                    switch (("\n\t" + "\n\t") + "foo bar") {
                    case "a":
                        var v18 string = "\n\t"
                        _ = v18
                    case "foo bar":
                        v4 -= (v4 & (v4 * v4))
                    case "":
                        var v19 int16 = 1000
                        _ = v19
                    default:
                        var v20 uint8 = 128
                        _ = v20
                    }
                    fmt.Println(v2, v17)
                    _ = fi16
                    _ = v17
                }
                switch "a" {
                case "\n\t":
                    if (42 == 2) {
                        _ = len(v1)
                    }
                }
                v21 := GenS4{F0: "hello", F1: -9223372036854775808}
                _ = v21
            } else {
                var v22 bool = (true || (!(!true)))
                _ = v22
            }
        }
    } else {
        v10 = v10
        delete(v1, (("foo bar" + (((("" + "\n\t") + ("foo bar" + "foo bar")) + "a") + ("\n\t" + "\n\t"))) + "hello"))
    }
    var v23 float64 = 1e10
    v24 := GenS3{F0: ("a" + "\n\t"), F1: false, F2: uint32(v4)}
    if v2 {
        v3 = (1e10 + (float64(float32(0.1)) / (v3 * 1.5)))
        var v25 uint8 = 1
        if ((^(v25)) > (v25 * v25)) {
            v26 := []int{-842, 32768}
            if len(v26) > 0 {
                v26[0] = -565
            }
            switch 2147483647 {
            case 42:
                if false {
                    for fi27 := 0; fi27 < 4; fi27++ {
                        var v28 uint8 = (v25 << ((4294967295) & 63))
                        fmt.Println(v2, v4, v2)
                        if len(v26) >= 2 {
                            v26 = v26[1:len(v26)]
                        }
                        _ = fi27
                        _ = v28
                    }
                }
            }
            _ = v26
        }
        _ = v25
    }
    v29 := GenS1{F0: -2147483648}
    v30 := GenS5{F0: 2, F1: "\n\t"}
    v3 = v3
    for fi31 := 0; fi31 < 9; fi31++ {
        v32 := GenS1{F0: fi31}
        _ = fi31
        _ = v32
    }
    v23 = v3
    v1[""] = (("a" + "") + "a")
    v2 = v2
    var v33 uint32 = 32
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v9
}
