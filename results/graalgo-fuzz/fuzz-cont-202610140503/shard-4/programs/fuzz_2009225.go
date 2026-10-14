package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int32 }
type GenS2 struct { F0 string; F1 string; F2 uint64 }
func genFn0(p0 int8, p1 uint32, p2 int) int8 {
    switch "a" {
    case "\n\t":
        switch 127 {
        case -2147483648:
            if true {
                v7 := []int{p2, p2}
                if (198 == 1) {
                    var v8 int16 = 810
                    var v9 uint32 = (^(p1))
                    _ = v8
                    _ = v9
                } else {
                    var v10 string = "a"
                    if true {
                        switch (p2 | p2) {
                        case -143:
                            p2 = 2147483647
                        }
                    }
                    if len(v7) > 0 {
                        v7[0] = 1
                    }
                    _ = v10
                }
                _ = v7
            } else {
                var v11 bool = false
                _ = v11
            }
        case -30:
            switch "\n\t" {
            case "hello":
                fmt.Println(p2, p0, p2)
            case "foo bar":
                v12 := []int{(p2 | 32768), 0, -1000, int(p0)}
                _ = v12
            case "\n\t":
                p2 = (p2 | (p2 | p2))
            default:
                switch (("" + "foo bar") + "a") {
                case "hello":
                    var v13 int16 = -839
                    _ = v13
                }
            }
        }
    }
    var v14 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return (p0 * p0)
}

func main() {
    var v1 int8 = 42
    v2 := []uint32{608, 907, 911}
    var v3 uint32 = 2
    var v4 uint = 105
    var v5 int8 = (v1 * (^(v1)))
    v6 := []uint32{}
    var v15 int32 = 0
    var v16 uint64 = 0
    if ((((v3 ^ 2) != v3) || false) && (128 <= 9223372036854775807)) {
        if (-2147483648 >= (v15 & 479)) {
            var v17 float32 = 3.14159
            v18 := GenS1{F0: 2, F1: v5, F2: 1}
            var v19 int8 = genFn0(v5, (v3 ^ 172), -72)
            _ = v17
            _ = v18
            _ = v19
        } else {
            switch 688 {
            case -530:
                v2 = append(v2, v3)
            case -1:
                switch "\n\t" {
                case "a":
                    var v20 int = int(v15)
                    _ = v20
                }
            }
        }
        var v21 uint8 = 128
        _ = v21
    } else {
        var v22 bool = (0 == 1000)
        v3 = (v3 >> ((v4) & 63))
        _ = v22
    }
    v5 -= v1
    switch 32767 {
    case 556:
        v15 = (v15 * -38)
    case -813:
        for fi23 := 0; fi23 < 13; fi23++ {
            if len(v6) > 0 {
                v6[0] = v3
            }
            switch "" {
            case "foo bar":
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            v6 = append(v6, (v3 & 42))
            _ = fi23
        }
    case 128:
        switch 127 {
        case 32768:
            _ = len(v2)
        }
    default:
        v16 = (^(v16))
    }
    var v24 string = ("\n\t" + "\n\t")
    switch ("a" + v24) {
    case "foo bar":
        var v25 string = ("\n\t" + (v24 + v24))
        _ = v25
    case "":
        var v26 uint32 = (^(v3))
        _ = v26
    case "a":
        v27 := GenS2{F0: (("\n\t" + "hello") + v24), F1: v24, F2: v16}
        _ = v27
    }
    v3 = v3
    var v28 uint16 = 32767
    switch 0 {
    case -9223372036854775808:
        var v29 int16 = 0
        _ = v29
    case 9223372036854775807:
        if ((2147483648 <= (v3 - 1)) || false) {
            v16 = (^(v16))
            v4 += 18446744073709551615
        }
    }
    v16 = v16
    switch 42 {
    case -1:
        var v30 int32 = int32(v16)
        _ = v30
    }
    v31 := GenS1{F0: 605, F1: (v5 ^ v5), F2: (^(v15))}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
