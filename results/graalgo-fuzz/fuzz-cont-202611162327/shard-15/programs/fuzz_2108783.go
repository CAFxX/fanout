package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 int32; F2 uint32 }
func genFn0(p0 bool) bool {
    v4 := []int{-1, 32768, 442}
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 uint = 2
        var v7 float32 = -2.25
        for fi8 := 0; fi8 < 9; fi8++ {
            if p0 {
                var v9 int = fi5
                _ = v9
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v10 int16 = 32767
            _ = fi8
            _ = v10
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = v4
    return (("foo bar" != (("a" + "hello") + ("a" + "\n\t"))) || p0)
}
func genFn1(p0 uint, p1 uint8, p2 int64) int16 {
    p2 += p2
    var v11 uint8 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return -1
}

func main() {
    var v1 int64 = -1
    var v2 int32 = 835
    var v3 bool = ("a" != ((("a" + ("foo bar" + ("hello" + ""))) + "foo bar") + ("" + "\n\t")))
    var v12 int64 = (-(v1))
    v2 = (v2 * (-(v2)))
    for fi13 := 0; fi13 < 10; fi13++ {
        if v3 {
            if ((2 != 1) || ("hello" != "\n\t")) {
                fmt.Println(v1)
                if v3 {
                    v14 := GenS1{F0: 3, F1: v1, F2: 32767}
                    var v15 uint16 = 321
                    _ = v14
                    _ = v15
                }
            }
        }
        switch ((("foo bar" + "\n\t") + "foo bar") + "") {
        case "\n\t":
            switch (^(fi13)) {
            case 1:
                if (127 != 1) {
                    v16 := GenS2{F0: 2}
                    var v17 float64 = (float64(float32(3.14159)) * 1e10)
                    _ = v16
                    _ = v17
                } else {
                    v18 := []int{fi13, fi13}
                    v1 -= v1
                    _ = v18
                }
            case 32768:
                switch "foo bar" {
                case "a":
                    v3 = (100 < -74)
                case "hello":
                    switch "foo bar" {
                    case "hello":
                        var v19 int = fi13
                        _ = v19
                    case "\n\t":
                        fmt.Println(v12, v12)
                    }
                }
            case -508:
                fmt.Println(v3, v12, v3)
            }
        case "foo bar":
            var v20 uint16 = 65535
            _ = v20
        }
        _ = fi13
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        var v22 int32 = -2147483648
        var v23 string = ("hello" + "")
        _ = fi21
        _ = v22
        _ = v23
    }
    v12 = v1
    v24 := GenS1{F0: 42, F1: v1, F2: 505}
    v25 := GenS3{F0: v12, F1: int32(v12), F2: 754}
    v2 = v2
    var v26 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
