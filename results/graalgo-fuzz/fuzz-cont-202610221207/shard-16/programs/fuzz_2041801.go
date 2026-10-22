package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int16; F2 int16 }

func main() {
    var v1 uint16 = 65535
    var v2 int8 = 42
    var v3 int16 = 1
    var v4 uint8 = 155
    var v5 float64 = ((1.5 + 1.5) - (3.14159 + (1e10 * (3.14159 - 0.0))))
    var v6 int32 = -1
    for fi7 := 0; fi7 < 13; fi7++ {
        v3 = v3
        var v8 float32 = 1.5
        _ = fi7
        _ = v8
    }
    switch (((("hello" + (("" + ("\n\t" + "\n\t")) + (("" + "foo bar") + ""))) + "a") + "hello") + ("foo bar" + "a")) {
    case "\n\t":
        if false {
            switch 42 {
            case 1:
                if false {
                    v9 := []int{127, -216}
                    _ = v9
                } else {
                    for fi10 := 0; fi10 < 1; fi10++ {
                        v2 -= (v2 + v2)
                        v6 -= (-(v6))
                        if ((int32(v2) != -2147483648) || false) {
                            var v11 uint32 = uint32(v4)
                            _ = v11
                        }
                        _ = fi10
                    }
                    switch -9223372036854775808 {
                    case 1:
                        v2 = (v2 * v2)
                    case -9223372036854775808:
                        var v12 int64 = 0
                        _ = v12
                    }
                }
            default:
                fmt.Println(v6, v6)
            }
            for fi13 := 0; fi13 < 6; fi13++ {
                var v14 float32 = (3.14159 * (-2.25 + 1.5))
                v15 := []int{}
                for fi16 := 0; fi16 < 7; fi16++ {
                    if false {
                        var v17 uint8 = (v4 << ((18446744073709551615) & 63))
                        _ = v17
                    }
                    v14 -= float32(((0.1 + v5) + 1e10))
                    var v18 string = "a"
                    _ = fi16
                    _ = v18
                }
                _ = fi13
                _ = v14
                _ = v15
            }
            fmt.Println(v5, v4, v4)
        }
    case "":
        for fi19 := 0; fi19 < 10; fi19++ {
            switch "foo bar" {
            case "foo bar":
                if (4294967295 < 150) {
                    var v20 int16 = v3
                    var v21 uint8 = 2
                    _ = v20
                    _ = v21
                }
            default:
                fmt.Println(v6, v2, v6)
            }
            _ = fi19
        }
    case "hello":
        var v22 uint = 4294967295
        _ = v22
    default:
        switch -1 {
        case -1:
            for fi23 := 0; fi23 < 11; fi23++ {
                v1 = v1
                _ = fi23
            }
        case -2147483648:
            if true {
                var v24 uint32 = 42
                var v25 uint16 = v1
                _ = v24
                _ = v25
            } else {
                v4 = v4
            }
        }
    }
    if (v6 == v6) {
        var v26 uint16 = (v1 % (((v1 ^ 255)) | 1))
        v4 = (v4 / ((v4) | 1))
        _ = v26
    } else {
        for fi27 := 0; fi27 < 3; fi27++ {
            var v28 bool = true
            _ = fi27
            _ = v28
        }
        v29 := GenS1{F0: v3}
        var v30 int16 = (v3 << ((4294967295) & 63))
        _ = v29
        _ = v30
    }
    var v31 int32 = (v6 - v6)
    v2 = 1
    fmt.Println(v4)
    switch 32767 {
    case 2147483647:
        fmt.Println(v2, v3)
    default:
        v4 = (v4 * v4)
    }
    if (!true) {
        var v32 string = (("" + (("hello" + "\n\t") + ("\n\t" + (((("foo bar" + "") + "hello") + "foo bar") + ("a" + ""))))) + "\n\t")
        for fi33 := 0; fi33 < 18; fi33++ {
            switch (fi33 >> ((138) & 63)) {
            case 428:
                v32 = v32
            case 128:
                v34 := GenS2{F0: (v31 * (^(v6))), F1: v3, F2: (v3 % ((v3) | 1))}
                _ = v34
            default:
                v1 = (v1 + (v1 + 42))
            }
            for fi35 := 0; fi35 < 9; fi35++ {
                fmt.Println(v5, v2, fi33)
                if (1.5 > (0.0 - v5)) {
                    if false {
                        fmt.Println(v4)
                    }
                } else {
                    var v36 bool = (false && (false && false))
                    v37 := GenS1{F0: v3}
                    v38 := []int{(fi35 + (^(fi35))), fi33, fi35, -48}
                    if (303 == 436) {
                        v3 = 1000
                    }
                    var v39 int16 = -52
                    _ = v36
                    _ = v37
                    _ = v38
                    _ = v39
                }
                _ = fi35
            }
            _ = fi33
        }
        _ = v32
    }
    if (!false) {
        v40 := []int{}
        for fi41 := 0; fi41 < 4; fi41++ {
            fmt.Println(v31)
            _ = fi41
        }
        v40 = append(v40, 32767)
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
