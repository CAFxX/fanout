package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint; F1 string }
func genFn0() uint16 {
    var v7 string = "a"
    v7 = (v7 + v7)
    for fi8 := 0; fi8 < 8; fi8++ {
        switch v7 {
        case "\n\t":
            for fi9 := 0; fi9 < 16; fi9++ {
                var v10 float32 = (1.5 - 0.1)
                _ = fi9
                _ = v10
            }
        case "a":
            v11 := GenS1{F0: v7, F1: 65535}
            _ = v11
        default:
            v12 := []int{(fi8 / (((fi8 & fi8)) | 1)), 0, fi8}
            _ = v12
        }
        var v13 int32 = -751
        _ = fi8
        _ = v13
    }
    _ = v7
    return 1
}
func genFn1(p0 uint8, p1 uint, p2 uint32) bool {
    fmt.Println(p2)
    p2 -= 2147483648
    _ = p0
    _ = p1
    _ = p2
    return true
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = 2147483647
    v3 := []string{"foo bar", "foo bar", ("\n\t" + "foo bar"), "hello"}
    var v4 uint8 = 128
    v5 := make(map[uint32]bool)
    v6 := make(map[int]bool)
    var v14 uint64 = 2
    var v15 uint = 65535
    v16 := GenS1{F0: (("hello" + (("" + ("hello" + (("a" + (("hello" + ((("a" + (("" + "hello") + ("a" + "\n\t"))) + ("foo bar" + "\n\t")) + "\n\t")) + "\n\t")) + ("hello" + ("hello" + "foo bar"))))) + "\n\t")) + ("hello" + "hello")), F1: 42}
    if false {
        var v17 int16 = 1
        v2 = 0
        if false {
            var v18 int32 = v2
            _ = v18
        }
        _ = v17
    }
    var v19 uint8 = v4
    _ = len(v3)
    var v20 int16 = -472
    var v21 int64 = -9223372036854775808
    switch "hello" {
    case "\n\t":
        v22 := GenS2{F0: v15, F1: ("\n\t" + "hello")}
        _ = v22
    case "":
        switch 42 {
        case 32768:
            if false {
                v15 += (v15 >> ((v15) & 63))
                fmt.Println(v19)
                v5[4294967295] = (v15 <= (v15 >> (((^(v15))) & 63)))
            }
        case 255:
            if false {
                for fi23 := 0; fi23 < 18; fi23++ {
                    _ = len(v3)
                    if (!(0.0 > 1.5)) {
                        _ = len(v1)
                        if true {
                            delete(v6, (-(fi23)))
                            delete(v5, uint32(v4))
                        }
                        delete(v6, int(v20))
                    }
                    switch ("" + "a") {
                    case "a":
                        _ = len(v5)
                    default:
                        v14 = (v14 | v14)
                    }
                    _ = fi23
                }
            } else {
                var v24 bool = true
                v3 = append(v3, ("foo bar" + ""))
                _ = v24
            }
        default:
            v4 = (v19 ^ uint8(v15))
        }
    case "a":
        v21 = -957
    }
    fmt.Println(v4)
    v4 -= 128
    switch ("" + "a") {
    case "hello":
        v21 = v21
    default:
        if true {
            var v25 uint64 = 0
            fmt.Println(v2, v21)
            _ = v25
        }
    }
    if true {
        switch "\n\t" {
        case "\n\t":
            var v26 int8 = -8
            _ = v26
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
