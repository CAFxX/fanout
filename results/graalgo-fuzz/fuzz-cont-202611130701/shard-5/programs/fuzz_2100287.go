package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 uint8; F2 bool }
type GenS3 struct { F0 int64; F1 bool; F2 int }
type GenS4 struct { F0 uint16; F1 int32; F2 uint }
func genFn0(p0 int) int16 {
    switch p0 {
    case 0:
        v5 := GenS1{F0: 32767}
        _ = v5
    }
    var v6 int = (p0 << ((4294967295) & 63))
    if (2 < 2147483647) {
        v7 := GenS2{F0: (("foo bar" + "foo bar") + ("foo bar" + "hello")), F1: 128, F2: false}
        v6 += v6
        v8 := []int{2147483647, v6, v6}
        _ = v7
        _ = v8
    } else {
        if false {
            v9 := []int{p0, p0, 1, (v6 * (v6 / ((v6) | 1)))}
            _ = v9
        } else {
            v6 = 743
            if false {
                for fi10 := 0; fi10 < 6; fi10++ {
                    p0 += (fi10 * -424)
                    v6 = 32767
                    for fi11 := 0; fi11 < 17; fi11++ {
                        v12 := []int{}
                        var v13 float64 = (((3.14159 - (0.1 / 3.14159)) / 0.1) * 0.0)
                        v14 := GenS3{F0: -9223372036854775808, F1: false, F2: fi10}
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    _ = fi10
                }
            } else {
                v15 := []int{p0, p0}
                v15 = append(v15, v6)
                _ = v15
            }
        }
        switch (("hello" + "hello") + "hello") {
        case "\n\t":
            var v16 uint32 = 2147483647
            _ = v16
        }
        v17 := GenS1{F0: 65535}
        _ = v17
    }
    switch 0 {
    case 889:
        if true {
            if (2 >= 2147483647) {
                var v18 string = "\n\t"
                v19 := []int{(-(p0)), p0, p0, (v6 - v6)}
                _ = v18
                _ = v19
            }
        } else {
            v20 := []int{127, p0, 2147483647}
            _ = v20
        }
    default:
        v21 := GenS4{F0: 65535, F1: 845, F2: 231}
        _ = v21
    }
    _ = p0
    _ = v6
    return -32768
}
func genFn1(p0 uint16, p1 int32) uint {
    p0 += p0
    var v22 int = 128
    v22 -= (v22 >> ((861) & 63))
    p0 = p0
    _ = p0
    _ = p1
    _ = v22
    return 65535
}

func main() {
    var v1 string = "foo bar"
    var v2 uint64 = 18446744073709551615
    var v3 int16 = -1
    var v4 uint16 = 32767
    v23 := []int{}
    switch 32768 {
    case 2147483647:
        v3 = 1000
    default:
        for fi24 := 0; fi24 < 5; fi24++ {
            if true {
                if (-2.25 <= 1.5) {
                    v4 = v4
                    var v25 uint8 = 2
                    _ = v25
                } else {
                    v1 = ((v1 + "hello") + v1)
                }
            } else {
                var v26 float64 = ((((3.14159 + (3.14159 + 1e10)) + float64((3.14159 * float32(3.14159)))) / 3.14159) * (1.5 - (float64(((1e10 * 3.14159) * (0.1 / 1.0))) + ((3.14159 - 1.5) + float64(1.5)))))
                _ = v26
            }
            _ = fi24
        }
    }
    var v27 uint32 = 0
    v27 -= (v27 ^ (v27 >> ((255) & 63)))
    if false {
        fmt.Println(v4, v2, v2)
        var v28 uint = genFn1(v4, 0)
        switch 4294967295 {
        case 2147483647:
            switch (v1 + (((("a" + v1) + "foo bar") + v1) + ((v1 + (v1 + "\n\t")) + v1))) {
            case "a":
                v2 = 1
            default:
                v23 = append(v23, -9223372036854775808)
            }
        }
        _ = v28
    } else {
        var v29 string = ("" + "foo bar")
        var v30 int64 = 0
        var v31 uint = 255
        _ = v29
        _ = v30
        _ = v31
    }
    v32 := GenS2{F0: ("foo bar" + v1), F1: 1, F2: ((!true) && true)}
    var v33 float64 = 1.5
    v34 := GenS2{F0: ("hello" + v1), F1: 127, F2: (255 == genFn1(255, -2147483648))}
    if (false || ((!(float64((float32((v33 * v33)) * -2.25)) <= v33)) && ((v1 + (("a" + v1) + (("\n\t" + v1) + (v1 + v1)))) >= "hello"))) {
        fmt.Println(v27, v2)
    }
    v2 += v2
    var v35 int8 = -60
    v36 := GenS2{F0: (v1 + ""), F1: 167, F2: (false && false)}
    v35 = (-(v35))
    v4 = (v4 << ((18446744073709551615) & 63))
    fmt.Println(v4, v27, v27)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
