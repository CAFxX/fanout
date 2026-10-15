package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 uint32; F2 int }
type GenS4 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 uint32 }
func genFn0() uint16 {
    var v4 int = -2147483648
    var v5 uint32 = 2
    v6 := []int{v4, v4, (^(v4))}
    var v7 int = (v4 % (((v4 | -284)) | 1))
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 float64 = 1e10
    var v2 uint = 351
    v3 := []string{"hello", ("" + ("" + "foo bar")), (("" + ("" + "hello")) + (("hello" + "foo bar") + "\n\t"))}
    if (820 >= 64) {
        var v8 string = (("foo bar" + "foo bar") + "hello")
        for fi9 := 0; fi9 < 10; fi9++ {
            v3 = append(v3, (v8 + "\n\t"))
            _ = fi9
        }
        fmt.Println(v8)
        _ = v8
    }
    v10 := GenS1{F0: true}
    v11 := GenS1{F0: false}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v1)
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 string = ("a" + ("a" + ("foo bar" + "a")))
        _ = fi12
        _ = v13
    }
    switch (("" + "") + "hello") {
    case "a":
        v2 += v2
    case "hello":
        for fi14 := 0; fi14 < 15; fi14++ {
            if ((false && (!(-2.25 <= 0.0))) && true) {
                if (3.14159 != ((float32(v1) * (3.14159 * 3.14159)) + 1.5)) {
                    if (1 == 2147483647) {
                        v1 = v1
                    } else {
                        v1 -= (1e10 * v1)
                        v1 = v1
                    }
                    fmt.Println(fi14)
                }
            } else {
                fmt.Println(v1, fi14, fi14)
                v15 := GenS2{F0: 244}
                v16 := GenS1{F0: (v1 != 0.1)}
                _ = v15
                _ = v16
            }
            _ = fi14
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v17 := GenS2{F0: 369}
    fmt.Println(v2, v1)
    var v18 uint16 = genFn0()
    switch ("" + "") {
    case "a":
        switch -889 {
        case 1:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case -705:
            for fi19 := 0; fi19 < 14; fi19++ {
                switch (("foo bar" + "hello") + "a") {
                case "a":
                    v3 = append(v3, "\n\t")
                }
                _ = fi19
            }
        }
    default:
        v18 += (v18 | (^(v18)))
    }
    if (((v1 * -2.25) / 1.5) <= ((3.14159 - v1) / v1)) {
        v3 = append(v3, "")
        v20 := GenS3{F0: 42, F1: 2, F2: 0}
        _ = v20
    }
    var v21 float64 = 1.5
    if false {
        var v22 string = "a"
        _ = v22
    } else {
        if true {
            v23 := GenS3{F0: -813, F1: 881, F2: 0}
            switch int(v18) {
            case -526:
                v24 := GenS4{F0: 25, F1: 42, F2: 2147483648}
                _ = v24
            case 1:
                var v25 float64 = (v21 * float64(0.1))
                _ = v25
            case 2147483647:
                switch -9223372036854775808 {
                case 808:
                    if len(v3) > 0 {
                        v3[0] = "a"
                    }
                default:
                    v21 = (((-2.25 / v1) / v1) - 0.0)
                }
            default:
                fmt.Println(v1, v2)
            }
            v18 = v18
            _ = v23
        }
    }
    if (!true) {
        for fi26 := 0; fi26 < 1; fi26++ {
            v2 = (^(v2))
            v27 := GenS5{F0: 0, F1: -1}
            v28 := GenS6{F0: 2}
            _ = fi26
            _ = v27
            _ = v28
        }
    } else {
        var v29 bool = ((1e10 - 1.5) >= 0.0)
        v1 += (float64(1e10) / 1e10)
        _ = v29
    }
    var v30 int = -9223372036854775808
    v3 = append(v3, "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v30
}
