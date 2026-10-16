package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }

func main() {
    var v1 int8 = -1
    v2 := make(map[string]int)
    var v3 int = -9223372036854775808
    var v4 float32 = (((1e10 - -2.25) / 1.0) / float32(1.5))
    for fi5 := 0; fi5 < 19; fi5++ {
        for fi6 := 0; fi6 < 5; fi6++ {
            if false {
                switch ("foo bar" + "") {
                case "hello":
                    switch "\n\t" {
                    case "\n\t":
                        v7 := []int{int(v1), (fi5 ^ fi6), -252}
                        _ = v7
                    }
                default:
                    v8 := []int{v3, (fi6 << ((4294967295) & 63))}
                    _ = v8
                }
                fmt.Println(v3)
            }
            v3 = 9223372036854775807
            switch "\n\t" {
            case "a":
                v4 -= float32(-2.25)
            case "":
                v4 -= 1.5
            default:
                var v9 uint32 = 2
                var v10 uint8 = 0
                var v11 int64 = -1
                _ = v10
                _ = v11
                _ = v9
            }
            v12 := GenS1{F0: 255, F1: 1}
            var v13 float32 = v4
            if (("foo bar" <= "a") && (!(float64(v4) == 1.5))) {
                var v14 uint16 = 1
                var v15 uint64 = 419
                _ = v14
                _ = v15
            }
            if (v1 > v1) {
                v4 += 3.14159
                v4 = ((1.5 / (float32((-2.25 + ((1.5 - float64((v13 + ((float32(1e10) * -2.25) + v13)))) * 1.5))) * (float32(0.0) / v13))) + -2.25)
                fmt.Println(fi6, v13)
            }
            _ = fi6
            _ = v12
            _ = v13
        }
        v16 := []int{0, (^(fi5)), v3}
        _ = fi5
        _ = v16
    }
    var v17 int32 = 2147483647
    var v18 uint8 = 128
    var v19 int = v3
    if (false && true) {
        if true {
            v1 = -128
            if false {
                switch (v3 ^ 4294967295) {
                case 42:
                    fmt.Println(v4)
                case 2147483647:
                    var v20 uint32 = 2
                    _ = v20
                default:
                    v4 -= 0.0
                }
                fmt.Println(v17)
                v21 := GenS1{F0: 1, F1: 42}
                _ = v21
            } else {
                _ = len(v2)
            }
        }
        var v22 uint16 = 2
        v23 := GenS1{F0: 251, F1: -9223372036854775808}
        _ = v22
        _ = v23
    }
    fmt.Println(v19, v18, v3)
    fmt.Println(v19, v18)
    v1 += 51
    v24 := []int{(v19 / ((4294967295) | 1))}
    v17 -= 1
    var v25 uint32 = 42
    v3 = (-(v19))
    var v26 bool = (4294967295 <= 18446744073709551615)
    if v26 {
        delete(v2, "hello")
        v27 := GenS1{F0: 42, F1: -9223372036854775808}
        switch ("a" + (((((((((("foo bar" + "\n\t") + "") + ("hello" + ((("foo bar" + ("hello" + (("hello" + "foo bar") + "a"))) + "foo bar") + "foo bar"))) + "") + ("a" + "")) + "hello") + "a") + (("foo bar" + "") + "a")) + "a") + "foo bar")) {
        case "\n\t":
            var v28 int32 = v17
            _ = v28
        default:
            var v29 int32 = (-(v17))
            _ = v29
        }
        _ = v27
    }
    v30 := GenS1{F0: 4294967295, F1: 9223372036854775807}
    _ = len(v24)
    if ((!true) && false) {
        fmt.Println(v1, v17)
        switch v19 {
        case 9223372036854775807:
            var v31 int64 = 763
            _ = v31
        }
    } else {
        switch v3 {
        case 32768:
            if ((-2.25 * float64(v4)) < 1.5) {
                v18 = 255
                switch v3 {
                case 32767:
                    if false {
                        for fi32 := 0; fi32 < 11; fi32++ {
                            var v33 uint32 = v25
                            v34 := GenS1{F0: 0, F1: 1}
                            v35 := GenS1{F0: 255, F1: -9223372036854775808}
                            _ = fi32
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                    }
                default:
                    v3 = (v19 >> ((0) & 63))
                }
            } else {
                for fi36 := 0; fi36 < 13; fi36++ {
                    v4 += v4
                    if (!(!v26)) {
                        v25 = (v25 + v25)
                    } else {
                        for fi37 := 0; fi37 < 1; fi37++ {
                            v38 := v2["hello"]
                            _ = fi37
                            _ = v38
                        }
                        var v39 uint64 = 474
                        fmt.Println(v25, v1)
                        _ = v39
                    }
                    _ = fi36
                }
            }
        case 42:
            var v40 uint16 = 42
            _ = v40
        case 128:
            for fi41 := 0; fi41 < 7; fi41++ {
                fmt.Println(v17)
                _ = fi41
            }
        }
        v19 = -9223372036854775808
        for fi42 := 0; fi42 < 6; fi42++ {
            var v43 int32 = (v17 << ((0) & 63))
            _ = fi42
            _ = v43
        }
    }
    if (0 <= (v3 + (v3 / ((248) | 1)))) {
        var v44 int32 = 42
        v25 = (v25 + 65535)
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v4
}
