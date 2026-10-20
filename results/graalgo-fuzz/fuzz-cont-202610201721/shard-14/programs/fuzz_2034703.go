package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 bool; F1 uint16 }

func main() {
    var v1 float64 = (-2.25 - ((3.14159 - ((0.0 - 0.0) / 3.14159)) + (3.14159 + (float64(0.0) * 3.14159))))
    var v2 int8 = 10
    var v3 bool = true
    v2 += (v2 ^ v2)
    var v4 bool = v3
    var v5 uint = 42
    v3 = ("\n\t" != ("foo bar" + "hello"))
    if (-673 <= -592) {
        v6 := GenS1{F0: 0}
        v7 := []int{9223372036854775807}
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := GenS2{F0: "hello", F1: v5}
            v1 = 0.0
            _ = fi8
            _ = v9
        }
        _ = v6
        _ = v7
    }
    switch 9223372036854775807 {
    case -447:
        switch 4294967295 {
        case -2147483648:
            v10 := []int{}
            _ = v10
        case 1:
            var v11 int8 = v2
            _ = v11
        case -436:
            v1 -= (v1 - -2.25)
        default:
            var v12 string = "a"
            _ = v12
        }
    case 42:
        v13 := GenS2{F0: ("foo bar" + ("hello" + "foo bar")), F1: (^(v5))}
        _ = v13
    default:
        for fi14 := 0; fi14 < 18; fi14++ {
            v15 := []int{0}
            v15 = append(v15, fi14)
            v15 = append(v15, (fi14 << ((v5) & 63)))
            _ = fi14
            _ = v15
        }
    }
    if (0 <= 2) {
        v16 := GenS2{F0: (("\n\t" + (((((("\n\t" + "foo bar") + "a") + "") + "") + "") + (("hello" + (("foo bar" + "foo bar") + "")) + ("hello" + "\n\t")))) + (("hello" + ((("foo bar" + "a") + "\n\t") + "foo bar")) + ("hello" + (("\n\t" + ("" + "hello")) + "a")))), F1: v5}
        _ = v16
    } else {
        v1 -= v1
        if (-713 < -1) {
            if v4 {
                var v17 int = -1
                _ = v17
            } else {
                var v18 string = ("hello" + "foo bar")
                if true {
                    v19 := []int{127, 0, -981}
                    _ = v19
                } else {
                    for fi20 := 0; fi20 < 7; fi20++ {
                        for fi21 := 0; fi21 < 10; fi21++ {
                            v22 := []int{}
                            _ = fi21
                            _ = v22
                        }
                        _ = fi20
                    }
                    v23 := []int{32768, 128, 229, 32768}
                    _ = v23
                }
                v24 := []int{898}
                _ = v18
                _ = v24
            }
            v2 += (^(v2))
        } else {
            v25 := []int{228, -9223372036854775808, 2147483647}
            v3 = v4
            if v3 {
                v26 := GenS3{F0: (!true), F1: 522}
                if true {
                    var v27 uint32 = 2147483647
                    _ = v27
                }
                var v28 int16 = 1
                _ = v26
                _ = v28
            }
            _ = v25
        }
        fmt.Println(v5, v2, v2)
    }
    var v29 bool = (255 >= 42)
    var v30 uint16 = 369
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
