package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 uint8 }
func genFn0(p0 uint8, p1 int8) int8 {
    p1 = -128
    p1 = (-(p1))
    for fi4 := 0; fi4 < 14; fi4++ {
        p0 = (p0 / (((^(p0))) | 1))
        _ = fi4
    }
    _ = p0
    _ = p1
    return (p1 * 0)
}

func main() {
    var v1 string = (("hello" + "\n\t") + "")
    v2 := []int{9223372036854775807, 32767, 32767}
    var v3 int = -276
    var v5 string = "hello"
    var v6 uint64 = 18446744073709551615
    var v7 uint16 = 65535
    if false {
        v8 := GenS1{F0: v6, F1: -347}
        switch 4294967295 {
        case 127:
            for fi9 := 0; fi9 < 2; fi9++ {
                var v10 string = "hello"
                if (!false) {
                    v11 := GenS1{F0: v6, F1: 1}
                    _ = v11
                } else {
                    var v12 uint64 = 4294967295
                    v6 = 18446744073709551615
                    _ = v12
                }
                for fi13 := 0; fi13 < 18; fi13++ {
                    v2 = append(v2, fi9)
                    v14 := GenS1{F0: (v6 ^ (v6 << ((2) & 63))), F1: -1}
                    _ = fi13
                    _ = v14
                }
                _ = fi9
                _ = v10
            }
        case -2147483648:
            fmt.Println(v5, v5, v7)
        }
        if ((false || true) && (!true)) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v15 := GenS2{F0: 65535}
            _ = len(v2)
            _ = v15
        }
        _ = v8
    }
    v2 = append(v2, v3)
    var v16 bool = (false && true)
    if v16 {
        var v17 float64 = 1e10
        _ = v17
    }
    switch v1 {
    case "foo bar":
        if (!true) {
            if false {
                v5 = "a"
                v2 = append(v2, (v3 >> ((876) & 63)))
            } else {
                v7 = v7
                v7 -= 1
                fmt.Println(v3)
            }
            v6 += v6
        }
    case "hello":
        switch "a" {
        case "\n\t":
            v2 = append(v2, 676)
        case "hello":
            switch (v5 + v5) {
            case "hello":
                for fi18 := 0; fi18 < 16; fi18++ {
                    for fi19 := 0; fi19 < 6; fi19++ {
                        switch (^(v3)) {
                        case 32767:
                            v5 = "a"
                        case -9223372036854775808:
                            if v16 {
                                v6 = v6
                                fmt.Println(fi19, fi19)
                                fmt.Println(v1)
                            } else {
                                v5 = ""
                            }
                        case 1:
                            if len(v2) > 0 {
                                v2[0] = fi19
                            }
                        }
                        fmt.Println(v7, fi18)
                        _ = fi19
                    }
                    _ = fi18
                }
            case "a":
                fmt.Println(v3)
            }
        default:
            switch (v3 << ((1) & 63)) {
            case 757:
                switch v3 {
                case 128:
                    v3 = v3
                case 2147483647:
                    _ = len(v2)
                }
            case 128:
                if (!v16) {
                    v16 = (true && false)
                    v20 := GenS1{F0: v6, F1: 770}
                    if ((((!true) || (42 == 450)) && (!v16)) && (1 == 1)) {
                        v16 = (!v16)
                    }
                    _ = v20
                }
            case 32767:
                if v16 {
                    v21 := GenS3{F0: (v3 << ((959) & 63)), F1: 42}
                    _ = v21
                }
            default:
                v22 := GenS4{F0: v16}
                _ = v22
            }
        }
    case "":
        v5 = (v5 + (v5 + ""))
    default:
        v16 = (!(903 < -648))
    }
    v1 = (v1 + v1)
    fmt.Println(v7, v16, v16)
    if v16 {
        if (!v16) {
            v2 = append(v2, v3)
            var v23 int32 = -30
            _ = v23
        } else {
            var v24 int = v3
            fmt.Println(v7, v1)
            var v25 float64 = (3.14159 / 1.0)
            _ = v24
            _ = v25
        }
    }
    if v16 {
        var v26 uint16 = v7
        v6 -= 1
        _ = v26
    } else {
        switch (v3 * (v3 + v3)) {
        case 42:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 4294967295:
            var v27 uint64 = v6
            _ = v27
        }
    }
    if (550 < 194) {
        v28 := GenS5{F0: v7, F1: 133}
        var v29 int8 = 112
        _ = v28
        _ = v29
    } else {
        v30 := GenS5{F0: 32767, F1: 42}
        switch (v1 + "\n\t") {
        case "a":
            var v31 int16 = 1
            _ = v31
        default:
            v6 -= (v6 ^ 823)
        }
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
