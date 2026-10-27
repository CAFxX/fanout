package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 uint8 }
func genFn0(p0 int64) uint16 {
    p0 += (p0 - -719)
    var v6 int16 = 42
    switch 32768 {
    case 195:
        fmt.Println(v6)
    case -751:
        for fi7 := 0; fi7 < 20; fi7++ {
            if true {
                if false {
                    v6 = (v6 ^ (v6 + (v6 | (v6 + v6))))
                } else {
                    if (!(true && (-2147483648 >= 287))) {
                        var v8 uint = 18446744073709551615
                        _ = v8
                    } else {
                        for fi9 := 0; fi9 < 16; fi9++ {
                            if (true && (!(!(!((((false || false) && true) && ("foo bar" == "")) && ((false && (480 < 65535)) || false)))))) {
                                v10 := []int{fi7, -2147483648}
                                var v11 uint32 = 1
                                var v12 bool = (42 <= 42)
                                _ = v10
                                _ = v11
                                _ = v12
                            }
                            var v13 int64 = 42
                            _ = fi9
                            _ = v13
                        }
                        var v14 int = (^(fi7))
                        p0 = (p0 | p0)
                        _ = v14
                    }
                    v15 := GenS1{F0: 255, F1: ""}
                    v16 := []int{}
                    _ = v15
                    _ = v16
                }
            } else {
                var v17 uint16 = 32767
                switch (fi7 + (-(fi7))) {
                case -1:
                    v17 = (v17 * (v17 | 366))
                case -9223372036854775808:
                    p0 = p0
                default:
                    v18 := GenS2{F0: 2147483647}
                    _ = v18
                }
                _ = v17
            }
            v19 := GenS1{F0: 255, F1: "hello"}
            _ = fi7
            _ = v19
        }
    default:
        v6 -= v6
    }
    _ = p0
    _ = v6
    return 32767
}

func main() {
    v1 := []uint32{42}
    v2 := make(map[int]bool)
    var v3 int16 = -661
    v4 := []string{"", "foo bar", (("hello" + "a") + "\n\t")}
    var v5 bool = (0 != -73)
    var v20 string = ("a" + "foo bar")
    for fi21 := 0; fi21 < 4; fi21++ {
        v5 = (-2147483648 >= -9223372036854775808)
        fmt.Println(v20, v20)
        if len(v4) > 0 {
            v4[0] = v20
        }
        _ = fi21
    }
    v22 := GenS2{F0: 1}
    v23 := GenS2{F0: -2147483648}
    v3 -= (v3 * (v3 ^ v3))
    switch 1 {
    case 32767:
        v20 = v20
    }
    var v24 int8 = 127
    var v25 uint32 = 1
    v3 = v3
    for fi26 := 0; fi26 < 13; fi26++ {
        for fi27 := 0; fi27 < 13; fi27++ {
            var v28 string = v20
            fmt.Println(fi26, v28)
            _ = fi27
            _ = v28
        }
        _ = len(v4)
        _ = fi26
    }
    fmt.Println(v25, v20, v3)
    switch "a" {
    case "hello":
        switch 1 {
        case -9223372036854775808:
            var v29 uint64 = 470
            _ = v29
        case 383:
            v30 := GenS3{F0: 0, F1: 0}
            _ = v30
        case 42:
            for fi31 := 0; fi31 < 3; fi31++ {
                for fi32 := 0; fi32 < 7; fi32++ {
                    v33 := v2[1]
                    v34 := v2[4294967295]
                    v35 := GenS2{F0: -1}
                    v36 := GenS1{F0: 2, F1: v20}
                    var v37 int = (fi31 >> ((770) & 63))
                    _ = fi32
                    _ = v33
                    _ = v34
                    _ = v35
                    _ = v36
                    _ = v37
                }
                fmt.Println(v25, fi31)
                _ = fi31
            }
        }
    case "\n\t":
        if v5 {
            v3 = -32768
            if len(v1) > 0 {
                v1[0] = (v25 << ((255) & 63))
            }
        }
    case "":
        v20 = v20
    }
    _ = len(v4)
    v25 += (v25 / ((75) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
