package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 int; F2 uint16 }
type GenS4 struct { F0 int32; F1 uint32; F2 uint16 }
func genFn0(p0 uint8, p1 int64, p2 bool) int16 {
    var v4 int16 = -1
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := GenS1{F0: true}
        for fi7 := 0; fi7 < 18; fi7++ {
            switch fi5 {
            case 782:
                switch "\n\t" {
                case "a":
                    switch "a" {
                    case "\n\t":
                        v8 := []int{}
                        _ = v8
                    case "foo bar":
                        v9 := []int{(fi5 + (^(fi5))), fi5}
                        var v10 int32 = -2147483648
                        v9 = append(v9, 4294967295)
                        _ = v10
                        _ = v9
                    }
                case "foo bar":
                    switch "foo bar" {
                    case "a":
                        var v11 uint16 = 2
                        _ = v11
                    case "":
                        if p2 {
                            v4 -= v4
                            if (434 >= v4) {
                                v4 = v4
                            }
                            fmt.Println(v4)
                            var v12 uint64 = 1
                            _ = v12
                        }
                    default:
                        p2 = (-20 <= 0)
                    }
                case "":
                    p1 += (p1 | p1)
                }
            default:
                var v13 int32 = -1
                _ = v13
            }
            _ = fi7
        }
        switch (((((("a" + "\n\t") + "foo bar") + "a") + ("hello" + "foo bar")) + "\n\t") + "hello") {
        case "hello":
            fmt.Println(fi5)
        default:
            v14 := []int{fi5, (fi5 / (((fi5 / ((fi5) | 1))) | 1))}
            _ = v14
        }
        _ = fi5
        _ = v6
    }
    fmt.Println(p0, v4, v4)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return v4
}
func genFn1(p0 int8, p1 uint32, p2 int) uint32 {
    for fi15 := 0; fi15 < 6; fi15++ {
        for fi16 := 0; fi16 < 15; fi16++ {
            p2 += (fi16 ^ 128)
            _ = fi16
        }
        if ("\n\t" == ((("a" + "hello") + "") + "")) {
            for fi17 := 0; fi17 < 1; fi17++ {
                var v18 uint8 = 127
                var v19 int8 = p0
                p1 = 974
                _ = fi17
                _ = v18
                _ = v19
            }
            v20 := GenS2{F0: 0}
            _ = v20
        }
        _ = fi15
    }
    v21 := []int{(^(p2)), (p2 | p2), p2}
    if len(v21) > 0 {
        v21[0] = (p2 << ((2) & 63))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return (p1 * 65535)
}

func main() {
    var v1 int = -1
    var v2 bool = false
    v3 := make(map[string]string)
    fmt.Println(v1, v2)
    for fi22 := 0; fi22 < 18; fi22++ {
        switch v1 {
        case 2147483647:
            v2 = ("" <= ("\n\t" + ("\n\t" + ("hello" + "a"))))
        default:
            fmt.Println(v2, v2, v1)
        }
        _ = fi22
    }
    fmt.Println(v1)
    v23 := v3["\n\t"]
    if (42 > 0) {
        v24 := GenS1{F0: true}
        _ = v24
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        v26 := GenS1{F0: (326 < 0)}
        _ = fi25
        _ = v26
    }
    var v27 int = (v1 + (v1 ^ (-(v1))))
    v28 := GenS3{F0: 42, F1: (v27 ^ v27), F2: 32767}
    v29 := GenS4{F0: -2147483648, F1: 4294967295, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
