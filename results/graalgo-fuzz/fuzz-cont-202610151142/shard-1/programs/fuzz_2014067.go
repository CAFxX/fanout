package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 bool }
func genFn0() string {
    var v6 uint16 = 65535
    v7 := []int{42, 659}
    _ = v6
    _ = v7
    return ("hello" + ((("hello" + "hello") + ("foo bar" + "a")) + ("\n\t" + "\n\t")))
}
func genFn1(p0 int, p1 uint32, p2 int8) bool {
    v8 := []int{}
    v9 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return true
}

func main() {
    var v1 bool = false
    v2 := []int{-2147483648, 32768, 32767}
    var v3 float32 = 0.0
    var v4 uint32 = 42
    var v5 int16 = -192
    fmt.Println(v5)
    var v10 uint8 = 42
    v4 = v4
    v4 -= 1
    v3 += 1e10
    var v11 bool = (-1 >= -997)
    if (32767 == 32767) {
        if len(v2) > 0 {
            v2[0] = 122
        }
    }
    switch ("\n\t" + "hello") {
    case "foo bar":
        var v12 int = 32767
        _ = v12
    case "a":
        v1 = ("" != "hello")
    case "":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        var v13 int16 = (v5 * v5)
        _ = v13
    }
    var v14 int8 = 127
    var v15 float32 = float32(((1.5 / ((-2.25 * -2.25) / 3.14159)) * (3.14159 - -2.25)))
    var v16 int = 105
    var v17 uint8 = 128
    switch (-(v16)) {
    case -1:
        if (v14 < (v14 & -128)) {
            for fi18 := 0; fi18 < 15; fi18++ {
                for fi19 := 0; fi19 < 6; fi19++ {
                    var v20 uint64 = 1
                    var v21 uint32 = v4
                    v22 := GenS2{F0: "hello", F1: 2}
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
                var v23 uint8 = v10
                _ = fi18
                _ = v23
            }
        } else {
            if ((v10 + v10) == 156) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = len(v2)
                fmt.Println(v3)
            }
        }
    case -218:
        if (0 < 2147483647) {
            for fi24 := 0; fi24 < 9; fi24++ {
                v25 := GenS3{F0: false}
                _ = fi24
                _ = v25
            }
        } else {
            fmt.Println(v5)
            switch v16 {
            case -2147483648:
                v1 = ((42 > -66) && (v11 && v1))
            case 42:
                var v26 uint8 = v17
                _ = v26
            case -1:
                fmt.Println(v5)
            default:
                var v27 float64 = -2.25
                _ = v27
            }
        }
    case 637:
        var v28 string = (((("foo bar" + ("a" + ((("hello" + ("hello" + "a")) + "a") + "foo bar"))) + ("" + "")) + (("foo bar" + "foo bar") + ("a" + "foo bar"))) + "a")
        _ = v28
    }
    v1 = false
    fmt.Println(v10, v16, v3)
    var v29 uint8 = (v17 << ((18446744073709551615) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
