package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 string }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 string; F1 int }
func genFn0(p0 int8, p1 int32, p2 uint32) int32 {
    if true {
        if true {
            for fi4 := 0; fi4 < 6; fi4++ {
                fmt.Println(fi4)
                for fi5 := 0; fi5 < 5; fi5++ {
                    p0 -= (p0 + p0)
                    v6 := GenS1{F0: 936, F1: (p1 - p1), F2: "a"}
                    v7 := []int{fi5, fi4, fi5, -292}
                    _ = fi5
                    _ = v6
                    _ = v7
                }
                if (true && (32767 >= 65535)) {
                    p0 = -123
                } else {
                    switch fi4 {
                    case 2147483647:
                        v8 := []int{-1, fi4, 9223372036854775807}
                        _ = v8
                    case 311:
                        v9 := GenS1{F0: uint64(p2), F1: (p1 * p1), F2: ("hello" + ("" + ((("foo bar" + "hello") + "") + "foo bar")))}
                        _ = v9
                    }
                }
                _ = fi4
            }
            var v10 int8 = -113
            _ = v10
        } else {
            var v11 uint = 42
            for fi12 := 0; fi12 < 20; fi12++ {
                var v13 bool = false
                v14 := GenS2{F0: 65535, F1: (p0 & -1)}
                var v15 uint8 = 38
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
            _ = v11
        }
        v16 := []int{-1, 42, -391, 0}
        _ = v16
    } else {
        p1 += (p1 << ((65535) & 63))
        var v17 int8 = -128
        if true {
            if (127 > 127) {
                v18 := []int{127, 32767}
                fmt.Println(p2, p1)
                p0 = 0
                _ = v18
            } else {
                if (32768 <= 2147483647) {
                    v19 := []int{}
                    _ = v19
                }
            }
        }
        _ = v17
    }
    var v20 uint64 = 4294967295
    v21 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    return (p1 % ((p1) | 1))
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{}
    var v3 bool = false
    var v22 float32 = ((((-2.25 - 1e10) - 1.5) - 0.0) / ((3.14159 / 0.1) / 1.0))
    var v23 float64 = 0.1
    if false {
        v23 -= -2.25
        if v3 {
            v24 := GenS3{F0: "foo bar", F1: -2147483648}
            _ = v24
        } else {
            v3 = (1 > -2147483648)
        }
    }
    var v25 int = 614
    if ("" == "foo bar") {
        v2 = append(v2, "")
        switch 9223372036854775807 {
        case -9223372036854775808:
            if v3 {
                switch (v25 ^ 1) {
                case 4294967295:
                    switch "foo bar" {
                    case "":
                        switch (-(v25)) {
                        case -277:
                            v3 = (v3 && (162 < -1))
                        case 1:
                            fmt.Println(v22, v22, v25)
                        }
                    default:
                        delete(v1, 507)
                    }
                }
                _ = len(v2)
            }
        case 32767:
            v3 = (!(1 != 42))
        default:
            v1[0] = "hello"
        }
        v22 = v22
    }
    switch v25 {
    case 780:
        var v26 int8 = 1
        _ = v26
    }
    delete(v1, 828)
    v2 = append(v2, ("\n\t" + "\n\t"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v3
}
