package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 uint64; F2 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int; F1 string }
type GenS7 struct { F0 bool; F1 int16 }
func genFn0(p0 uint) int {
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 int32 = 0
        for fi7 := 0; fi7 < 8; fi7++ {
            v8 := GenS1{F0: 0}
            _ = fi7
            _ = v8
        }
        var v9 int64 = 9223372036854775807
        _ = fi5
        _ = v6
        _ = v9
    }
    v10 := []int{}
    p0 = 18446744073709551615
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 bool = (fi11 >= fi11)
        for fi13 := 0; fi13 < 7; fi13++ {
            var v14 uint = (p0 & (p0 % (((p0 * (p0 >> (((^(p0))) & 63)))) | 1)))
            v12 = (!true)
            v15 := GenS2{F0: 501}
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi11
        _ = v12
    }
    var v16 int8 = 0
    _ = p0
    _ = v10
    _ = v16
    return 32767
}
func genFn1(p0 uint8, p1 uint16) int16 {
    switch "hello" {
    case "":
        fmt.Println(p1, p0)
    case "a":
        p0 += (p0 + 0)
    case "foo bar":
        if false {
            fmt.Println(p1)
            var v17 uint16 = p1
            p0 = (p0 + 20)
            _ = v17
        }
    }
    v18 := []int{594, genFn0(4294967295)}
    v19 := GenS2{F0: 42}
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return 1000
}

func main() {
    var v1 int = -9223372036854775808
    var v2 float32 = float32(1.5)
    v3 := []uint32{0, 826}
    var v4 int16 = 1
    if (v4 <= (v4 >> ((42) & 63))) {
        v20 := GenS2{F0: 2}
        if ((2147483648 == 158) && false) {
            if false {
                v21 := GenS1{F0: 42}
                v22 := GenS3{F0: 52}
                _ = v21
                _ = v22
            }
            v2 += -2.25
        } else {
            v23 := GenS4{F0: 0, F1: 582, F2: (-(v1))}
            for fi24 := 0; fi24 < 13; fi24++ {
                v3 = append(v3, 0)
                _ = fi24
            }
            if len(v3) > 0 {
                v3[0] = 205
            }
            _ = v23
        }
        v25 := GenS5{F0: 1}
        _ = v20
        _ = v25
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        for fi27 := 0; fi27 < 2; fi27++ {
            var v28 int16 = (v4 * v4)
            fmt.Println(v2, v1, fi26)
            _ = fi27
            _ = v28
        }
        for fi29 := 0; fi29 < 2; fi29++ {
            var v30 uint = 18446744073709551615
            _ = len(v3)
            _ = fi29
            _ = v30
        }
        var v31 uint16 = 213
        _ = fi26
        _ = v31
    }
    switch v1 {
    case -658:
        v4 -= 42
    }
    for fi32 := 0; fi32 < 19; fi32++ {
        for fi33 := 0; fi33 < 3; fi33++ {
            if true {
                var v34 string = ("foo bar" + "foo bar")
                if true {
                    v35 := GenS5{F0: 367}
                    v2 = v2
                    _ = v35
                }
                _ = v34
            } else {
                v1 += fi33
                if len(v3) > 0 {
                    v3[0] = 42
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            switch ("\n\t" + "hello") {
            case "a":
                fmt.Println(fi33)
            case "":
                v36 := GenS6{F0: (fi33 & fi33), F1: ""}
                _ = v36
            }
            v37 := GenS7{F0: true, F1: v4}
            _ = fi33
            _ = v37
        }
        _ = fi32
    }
    v4 -= v4
    if (42 >= 65535) {
        v2 += ((v2 * (float32(3.14159) - float32(1.5))) * (v2 * v2))
        var v38 uint16 = 2
        _ = v38
    } else {
        if false {
            v39 := GenS3{F0: 65535}
            _ = v39
        } else {
            var v40 float64 = (3.14159 - (float64((v2 * (float32(3.14159) / v2))) - -2.25))
            _ = v40
        }
    }
    var v41 string = "\n\t"
    for fi42 := 0; fi42 < 2; fi42++ {
        for fi43 := 0; fi43 < 14; fi43++ {
            v41 = (("hello" + "foo bar") + "hello")
            _ = fi43
        }
        v3 = append(v3, 433)
        _ = fi42
    }
    v3 = append(v3, 0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
}
