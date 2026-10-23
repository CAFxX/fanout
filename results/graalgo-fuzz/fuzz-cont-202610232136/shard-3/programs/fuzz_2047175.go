package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16 }
func genFn0() int32 {
    switch 9223372036854775807 {
    case 291:
        switch (("\n\t" + ((("a" + "") + "a") + "hello")) + "a") {
        case "foo bar":
            v7 := GenS1{F0: -89, F1: 0}
            _ = v7
        case "a":
            if (!false) {
                var v8 float32 = 1.5
                _ = v8
            }
        case "hello":
            var v9 uint32 = 1
            _ = v9
        }
    }
    if true {
        var v10 int64 = 1
        v11 := GenS2{F0: -1}
        fmt.Println(v10)
        _ = v10
        _ = v11
    }
    return 780
}
func genFn1(p0 uint) uint8 {
    fmt.Println(p0)
    switch "\n\t" {
    case "hello":
        switch "foo bar" {
        case "hello":
            var v12 uint16 = 255
            _ = v12
        case "\n\t":
            var v13 int = 0
            _ = v13
        }
    case "a":
        p0 = p0
    default:
        for fi14 := 0; fi14 < 1; fi14++ {
            v15 := GenS1{F0: 54, F1: 75}
            _ = fi14
            _ = v15
        }
    }
    fmt.Println(p0)
    var v16 uint16 = 65535
    _ = p0
    _ = v16
    return 128
}

func main() {
    var v1 bool = (525 != 255)
    var v2 int32 = 744
    var v3 int64 = -9223372036854775808
    var v4 string = ("foo bar" + ("a" + ("hello" + "foo bar")))
    var v5 uint = 230
    var v6 uint = (v5 & v5)
    for fi17 := 0; fi17 < 13; fi17++ {
        v5 += 0
        for fi18 := 0; fi18 < 5; fi18++ {
            var v19 int8 = 0
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    var v20 uint32 = 377
    if v1 {
        for fi21 := 0; fi21 < 12; fi21++ {
            var v22 float32 = 0.1
            _ = fi21
            _ = v22
        }
    }
    switch 22 {
    case 4294967295:
        v1 = (9223372036854775807 > -120)
    default:
        for fi23 := 0; fi23 < 1; fi23++ {
            var v24 uint32 = v20
            v25 := []int{(fi23 ^ (-(fi23))), (fi23 ^ -479), (fi23 ^ fi23)}
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    v2 -= v2
    v26 := GenS3{F0: (v2 * -2147483648)}
    var v27 int32 = (-(v2))
    v3 = v3
    if (v1 || ((1.5 >= ((((0.1 - (3.14159 - -2.25)) + ((float32(((float64(0.1) + 3.14159) - 1e10)) / (-2.25 / (1e10 + 1e10))) - 0.1)) + -2.25) / (3.14159 * (-2.25 + 1.5)))) || (v1 && v1))) {
        for fi28 := 0; fi28 < 3; fi28++ {
            var v29 int = 128
            _ = fi28
            _ = v29
        }
    } else {
        switch 1 {
        case 9223372036854775807:
            v30 := GenS3{F0: v2}
            _ = v30
        }
    }
    v1 = false
    switch v4 {
    case "\n\t":
        v31 := []int{-1, 32767, 42}
        _ = v31
    case "foo bar":
        for fi32 := 0; fi32 < 13; fi32++ {
            v33 := []int{2147483647, fi32}
            _ = fi32
            _ = v33
        }
    default:
        v34 := []int{42, -585, 32767}
        _ = v34
    }
    if ((-2.25 - 1e10) != 0.0) {
        for fi35 := 0; fi35 < 14; fi35++ {
            for fi36 := 0; fi36 < 4; fi36++ {
                v37 := GenS4{F0: 797}
                var v38 int16 = 42
                _ = fi36
                _ = v37
                _ = v38
            }
            v6 = 508
            var v39 int = fi35
            _ = fi35
            _ = v39
        }
    } else {
        for fi40 := 0; fi40 < 16; fi40++ {
            v41 := []int{-550, fi40, fi40}
            for fi42 := 0; fi42 < 4; fi42++ {
                v5 -= (^(v5))
                v43 := GenS3{F0: v2}
                v6 -= (v5 >> ((65535) & 63))
                _ = fi42
                _ = v43
            }
            v27 = (v2 * -1)
            _ = fi40
            _ = v41
        }
        v5 += (v6 ^ v6)
    }
    var v44 float32 = 0.1
    v27 -= 162
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
