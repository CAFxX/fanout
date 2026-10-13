package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 int64 }
func genFn0() int {
    v6 := []int{}
    var v7 uint8 = 88
    v7 = 1
    var v8 uint32 = 2147483648
    _ = v6
    _ = v7
    _ = v8
    return 32767
}

func main() {
    v1 := []string{}
    var v2 uint = 65535
    v3 := make(map[int]uint32)
    v4 := []int{32767}
    var v5 int8 = -128
    var v9 int32 = 592
    switch "a" {
    case "a":
        if len(v4) > 0 {
            v4[0] = 0
        }
    case "":
        var v10 uint8 = 128
        _ = v10
    }
    v11 := GenS1{F0: 255}
    var v12 uint64 = 1
    if (722 != 1) {
        for fi13 := 0; fi13 < 13; fi13++ {
            v9 = (v9 >> ((v2) & 63))
            var v14 string = (("\n\t" + ("\n\t" + "\n\t")) + "\n\t")
            _ = fi13
            _ = v14
        }
        var v15 uint64 = 1
        v16 := GenS2{F0: 9223372036854775807, F1: (v9 | v9)}
        _ = v15
        _ = v16
    }
    v2 += v2
    if false {
        var v17 uint = (v2 & (^(v2)))
        if false {
            _ = len(v3)
        }
        delete(v3, int(v17))
        _ = v17
    } else {
        v18 := GenS2{F0: genFn0(), F1: v9}
        for fi19 := 0; fi19 < 12; fi19++ {
            var v20 uint = v2
            if len(v4) > 0 {
                v4[0] = fi19
            }
            v21 := GenS3{F0: int8(v12)}
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = v18
    }
    _ = len(v1)
    var v22 string = "hello"
    v9 += (v9 | 42)
    var v23 float64 = ((3.14159 + (((((3.14159 - 1.5) + 1.5) + -2.25) / (((0.1 - (0.1 + float64(-2.25))) * -2.25) - (0.0 - 3.14159))) / 1e10)) + (1.5 * (1.5 + 3.14159)))
    v24 := GenS4{F0: (v12 % (((^(v12))) | 1)), F1: v2}
    for fi25 := 0; fi25 < 12; fi25++ {
        switch "\n\t" {
        case "hello":
            if (2 <= (v2 | (^(v2)))) {
                var v26 string = (("hello" + "foo bar") + "")
                var v27 uint64 = v12
                _ = v26
                _ = v27
            }
        default:
            v28 := GenS5{F0: -9223372036854775808}
            _ = v28
        }
        for fi29 := 0; fi29 < 4; fi29++ {
            var v30 uint8 = 68
            v31 := GenS1{F0: (v30 ^ 96)}
            var v32 bool = ((true || false) || ((42 >= -32768) && (false || true)))
            _ = fi29
            _ = v30
            _ = v31
            _ = v32
        }
        v33 := GenS1{F0: 128}
        _ = fi25
        _ = v33
    }
    for fi34 := 0; fi34 < 14; fi34++ {
        switch v22 {
        case "\n\t":
            var v35 uint8 = 131
            _ = v35
        case "":
            if (1.5 != 0.1) {
                for fi36 := 0; fi36 < 13; fi36++ {
                    var v37 bool = (((fi34 * fi36) < 1) || ((!(("\n\t" + ("a" + v22)) != v22)) && true))
                    _ = fi36
                    _ = v37
                }
                var v38 int16 = 1
                if false {
                    v38 -= (^(v38))
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v9 -= v9
                    delete(v3, (fi34 ^ fi34))
                }
                _ = v38
            } else {
                if true {
                    _ = len(v1)
                    var v39 string = v22
                    switch "\n\t" {
                    case "a":
                        v3[0] = 607
                    case "\n\t":
                        var v40 int64 = 42
                        _ = v40
                    }
                    _ = v39
                }
            }
        case "a":
            for fi41 := 0; fi41 < 7; fi41++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v42 := GenS4{F0: v12, F1: v2}
                switch v22 {
                case "\n\t":
                    v5 -= 0
                case "a":
                    var v43 uint16 = 84
                    _ = v43
                case "foo bar":
                    var v44 int = (fi41 >> ((v2) & 63))
                    _ = v44
                }
                fmt.Println(v22, fi34, v2)
                _ = fi41
                _ = v42
            }
        }
        if false {
            switch -9223372036854775808 {
            case 1:
                delete(v3, (fi34 << (((v2 - v2)) & 63)))
            case 2147483647:
                var v45 int = (fi34 ^ (fi34 % (((fi34 / ((int(v9)) | 1))) | 1)))
                _ = v45
            }
            v9 += -2147483648
        } else {
            var v46 uint = v2
            _ = v46
        }
        _ = fi34
    }
    v22 = v22
    var v47 int16 = 0
    v3[42] = 1
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v9
}
