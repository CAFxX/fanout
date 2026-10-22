package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 uint; F2 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 string }
func genFn0(p0 uint64, p1 uint64, p2 int16) int64 {
    p0 = (p1 & (p1 >> ((2) & 63)))
    v5 := GenS1{F0: 1}
    var v6 float64 = 1.5
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 int8 = 121
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -111
}
func genFn1(p0 uint16, p1 uint16, p2 int32) uint64 {
    switch 1 {
    case 42:
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := []int{128, (fi9 >> ((505) & 63))}
            switch "foo bar" {
            case "\n\t":
                for fi11 := 0; fi11 < 5; fi11++ {
                    fmt.Println(fi11)
                    v12 := GenS2{F0: 53, F1: 4294967295, F2: int16(p0)}
                    _ = fi11
                    _ = v12
                }
            default:
                var v13 uint16 = 2
                _ = v13
            }
            for fi14 := 0; fi14 < 8; fi14++ {
                var v15 float64 = 0.0
                var v16 string = ("a" + (("\n\t" + (("\n\t" + ("foo bar" + "foo bar")) + "a")) + "\n\t"))
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi9
            _ = v10
        }
    case 649:
        for fi17 := 0; fi17 < 16; fi17++ {
            if false {
                p1 += 1
            } else {
                fmt.Println(p0)
            }
            var v18 int16 = 244
            p1 += p1
            _ = fi17
            _ = v18
        }
    case 1:
        for fi19 := 0; fi19 < 17; fi19++ {
            if true {
                for fi20 := 0; fi20 < 9; fi20++ {
                    v21 := []int{}
                    p1 -= 42
                    _ = fi20
                    _ = v21
                }
                switch 137 {
                case -194:
                    fmt.Println(p0, p0, p1)
                }
                var v22 int = fi19
                _ = v22
            } else {
                v23 := GenS3{F0: 0}
                v24 := GenS4{F0: p0}
                _ = v23
                _ = v24
            }
            p0 -= 378
            v25 := []int{fi19}
            _ = fi19
            _ = v25
        }
    default:
        switch 42 {
        case -1:
            switch ("a" + "hello") {
            case "":
                var v26 uint16 = p1
                _ = v26
            case "a":
                for fi27 := 0; fi27 < 14; fi27++ {
                    v28 := GenS5{F0: -2147483648}
                    _ = fi27
                    _ = v28
                }
            case "hello":
                fmt.Println(p0, p1, p1)
            }
        case 1:
            v29 := []int{-9223372036854775808, 9223372036854775807, 0, 290}
            _ = v29
        case 42:
            v30 := []int{-1, 32768, 127}
            _ = v30
        }
    }
    p0 = p0
    v31 := GenS6{F0: 578}
    _ = p0
    _ = p1
    _ = p2
    _ = v31
    return 0
}

func main() {
    var v1 int32 = -844
    var v2 int = 598
    v3 := make(map[int]bool)
    var v4 uint = 406
    v2 -= v2
    switch (v2 % (((v2 & 32768)) | 1)) {
    case 553:
        v1 = (v1 / ((v1) | 1))
    default:
        for fi32 := 0; fi32 < 18; fi32++ {
            v33 := []int{fi32}
            v1 = (-(v1))
            fmt.Println(v1)
            _ = fi32
            _ = v33
        }
    }
    switch "a" {
    case "foo bar":
        v34 := []int{v2}
        _ = v34
    case "\n\t":
        v35 := GenS7{F0: "\n\t"}
        _ = v35
    case "a":
        v1 = -1
    }
    v2 -= v2
    delete(v3, 635)
    var v36 int32 = v1
    v37 := []int{}
    var v38 uint = (v4 * (v4 / (((v4 >> (((^(v4))) & 63))) | 1)))
    fmt.Println(v1, v1)
    v39 := GenS2{F0: v36, F1: (v4 & (^(v4))), F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
