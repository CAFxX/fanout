package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int32 }
func genFn0() uint32 {
    var v7 bool = false
    switch "" {
    case "":
        v7 = ("a" <= "")
    case "\n\t":
        if (!(v7 || (((1e10 + ((0.0 * 1e10) / float64((1.5 * 1e10)))) >= float64(3.14159)) && (!v7)))) {
            if ((42 == 5) && v7) {
                if (v7 || false) {
                    if v7 {
                        v7 = true
                        v8 := []int{128, 9223372036854775807, 127}
                        var v9 float64 = 1e10
                        _ = v8
                        _ = v9
                    }
                }
                v10 := GenS1{F0: 898, F1: 42}
                _ = v10
            }
        } else {
            var v11 int = -1
            _ = v11
        }
    case "a":
        if v7 {
            switch "a" {
            case "a":
                v12 := []int{0}
                _ = v12
            case "":
                v7 = (!(-1 > 61))
            case "\n\t":
                v13 := GenS1{F0: 18446744073709551615, F1: 1}
                _ = v13
            }
        } else {
            v14 := GenS1{F0: 18446744073709551615, F1: 127}
            _ = v14
        }
    }
    switch 9223372036854775807 {
    case 1:
        switch "a" {
        case "":
            var v15 int32 = 42
            _ = v15
        default:
            v16 := GenS2{F0: 1}
            _ = v16
        }
    }
    _ = v7
    return 2147483647
}

func main() {
    var v1 uint8 = 191
    var v2 uint8 = v1
    v3 := make(map[int]uint32)
    var v4 uint = 0
    v5 := make(map[uint32]string)
    v6 := make(map[string]string)
    for fi17 := 0; fi17 < 8; fi17++ {
        v4 = (v4 & v4)
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 float32 = -2.25
            fmt.Println(fi18)
            v1 = (v1 << ((65535) & 63))
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    v2 = 167
    v2 = v2
    v4 -= 18446744073709551615
    var v20 int = 332
    v1 -= v1
    _ = len(v5)
    var v21 bool = (((("foo bar" + (("\n\t" + "a") + "hello")) + ("a" + ("\n\t" + "a"))) >= "foo bar") && (!(0.1 > 3.14159)))
    for fi22 := 0; fi22 < 13; fi22++ {
        v23 := GenS2{F0: 1}
        if (v1 < v1) {
            for fi24 := 0; fi24 < 4; fi24++ {
                switch ("foo bar" + "") {
                case "\n\t":
                    v1 = v1
                case "hello":
                    v25 := []int{fi24}
                    _ = v25
                }
                delete(v5, 479)
                if v21 {
                    var v26 int8 = 0
                    _ = v26
                }
                _ = fi24
            }
            v5[65535] = (("" + (((("\n\t" + (("" + ("\n\t" + ("\n\t" + ("" + "foo bar")))) + "hello")) + "a") + "\n\t") + "\n\t")) + "foo bar")
            if v21 {
                var v27 int64 = -122
                var v28 int = fi22
                var v29 uint64 = 833
                _ = v27
                _ = v28
                _ = v29
            } else {
                switch (v20 / ((v20) | 1)) {
                case 42:
                    if v21 {
                        v4 = 28
                        v2 += v1
                    }
                case 836:
                    var v30 uint = 255
                    _ = v30
                }
                v20 += v20
            }
        }
        _ = fi22
        _ = v23
    }
    var v31 uint = (v4 * 42)
    var v32 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
