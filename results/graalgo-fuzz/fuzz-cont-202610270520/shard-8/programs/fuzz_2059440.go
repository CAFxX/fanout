package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16 }

func main() {
    v1 := []string{"", "\n\t", "", "foo bar"}
    v2 := []string{}
    v3 := []uint32{2, 65535, 1}
    var v4 uint16 = 32767
    v4 -= 1
    var v5 uint16 = 32767
    if (32767 == (v4 & v4)) {
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 float32 = ((1.5 - ((-2.25 * 1.5) / -2.25)) - 1e10)
            fmt.Println(v4, fi6)
            _ = fi6
            _ = v7
        }
        var v8 int8 = 103
        var v9 int = -265
        _ = v8
        _ = v9
    } else {
        if true {
            _ = len(v2)
            for fi10 := 0; fi10 < 3; fi10++ {
                fmt.Println(v5)
                var v11 bool = true
                var v12 uint = 50
                _ = fi10
                _ = v11
                _ = v12
            }
            var v13 uint = 255
            _ = v13
        }
        v14 := GenS1{F0: -1}
        _ = v14
    }
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    var v15 int32 = 1
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v16 := GenS1{F0: 87}
        if false {
            for fi17 := 0; fi17 < 8; fi17++ {
                fmt.Println(v4, v4)
                var v18 string = ("foo bar" + ("foo bar" + "\n\t"))
                v19 := GenS2{F0: fi17}
                _ = fi17
                _ = v18
                _ = v19
            }
            fmt.Println(v15)
            v5 = 1
        }
        _ = v16
    } else {
        var v20 uint8 = 1
        _ = v20
    }
    if ((!(0.1 <= float64(0.0))) || false) {
        switch -2147483648 {
        case 127:
            if false {
                var v21 uint32 = 2147483648
                v5 += v4
                var v22 bool = (false || true)
                _ = v21
                _ = v22
            }
        case 32767:
            fmt.Println(v4)
        case 42:
            if true {
                if (0 <= (v15 - (v15 - v15))) {
                    var v23 uint8 = 255
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v23
                } else {
                    v5 -= 297
                    var v24 int = -1
                    _ = v24
                }
                v4 += v4
            }
        }
    } else {
        if ((255 == 0) || true) {
            v25 := GenS3{F0: -82}
            switch (("a" + ("\n\t" + "foo bar")) + "foo bar") {
            case "\n\t":
                if len(v3) > 0 {
                    v3[0] = 42
                }
            case "a":
                var v26 int64 = 42
                _ = v26
            default:
                v15 = 0
            }
            if true {
                if false {
                    for fi27 := 0; fi27 < 13; fi27++ {
                        fmt.Println(v4, v5)
                        fmt.Println(v4)
                        _ = fi27
                    }
                    v5 += (^(v5))
                } else {
                    if (!(true || ((true || (-299 != -9223372036854775808)) && false))) {
                        switch "\n\t" {
                        case "foo bar":
                            var v28 uint16 = v5
                            _ = v28
                        case "a":
                            v15 -= -2147483648
                        default:
                            if (((v15 & v15) < v15) && false) {
                                v4 = 32767
                            } else {
                                v29 := GenS1{F0: -42}
                                _ = v29
                            }
                        }
                        v15 -= (v15 | (v15 & (v15 >> ((4294967295) & 63))))
                        _ = len(v3)
                    } else {
                        var v30 string = "\n\t"
                        _ = v30
                    }
                    fmt.Println(v15)
                    v3 = append(v3, 65535)
                }
            }
            _ = v25
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v31 := GenS3{F0: -1}
    v32 := GenS4{F0: 0}
    switch 127 {
    case 0:
        for fi33 := 0; fi33 < 19; fi33++ {
            v5 = v4
            _ = fi33
        }
    }
    if (int64(v15) == -9223372036854775808) {
        v5 = (v5 ^ v4)
        fmt.Println(v5, v5)
        v4 = v5
    }
    v15 = (v15 - -2147483648)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi34 := 0; fi34 < 6; fi34++ {
        _ = len(v2)
        v35 := GenS2{F0: fi34}
        _ = fi34
        _ = v35
    }
    v5 = v4
    for fi36 := 0; fi36 < 17; fi36++ {
        switch ("\n\t" + "") {
        case "\n\t":
            var v37 uint8 = 0
            _ = v37
        case "a":
            fmt.Println(v4, fi36, v4)
        case "foo bar":
            v4 -= (v4 + v5)
        default:
            if false {
                v15 -= 42
                var v38 bool = (((1e10 / -2.25) / 3.14159) > 1e10)
                _ = v38
            }
        }
        var v39 uint = 412
        var v40 int16 = 0
        _ = fi36
        _ = v39
        _ = v40
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v15 -= v15
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
