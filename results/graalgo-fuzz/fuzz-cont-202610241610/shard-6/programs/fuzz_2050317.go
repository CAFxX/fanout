package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 string; F2 int }
type GenS5 struct { F0 string }
func genFn0(p0 bool, p1 uint16, p2 int) int16 {
    var v5 uint16 = (p1 | (p1 & (p1 | 65535)))
    p1 = p1
    var v6 bool = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 519
}
func genFn1(p0 int64, p1 int16) string {
    if (!(!(true && false))) {
        if true {
            p0 -= 0
            var v7 string = ""
            _ = v7
        }
        switch -2147483648 {
        case -731:
            v8 := []int{-9223372036854775808, 1, 923, int(p1)}
            _ = v8
        case 127:
            for fi9 := 0; fi9 < 15; fi9++ {
                var v10 uint32 = 1
                switch -1 {
                case 42:
                    v11 := GenS1{F0: uint8(v10), F1: true}
                    _ = v11
                default:
                    fmt.Println(p0, p1, v10)
                }
                p1 += (p1 * (-(p1)))
                _ = fi9
                _ = v10
            }
        }
        var v12 int32 = -2147483648
        _ = v12
    } else {
        v13 := GenS1{F0: 255, F1: true}
        if false {
            var v14 string = (("foo bar" + "") + (("" + "a") + (("hello" + (("\n\t" + ("hello" + (("hello" + "a") + ("hello" + ("" + "\n\t"))))) + ("foo bar" + ("" + "foo bar")))) + "a")))
            var v15 uint = 255
            v15 = 821
            _ = v14
            _ = v15
        }
        p0 += 300
        _ = v13
    }
    v16 := GenS2{F0: (65535 == 65535), F1: 255}
    _ = p0
    _ = p1
    _ = v16
    return "foo bar"
}

func main() {
    var v1 uint64 = 0
    var v2 bool = false
    var v3 uint32 = 728
    var v4 int = 9223372036854775807
    var v17 int64 = 364
    if true {
        for fi18 := 0; fi18 < 1; fi18++ {
            switch ("" + ("a" + ("a" + ("foo bar" + (("a" + "a") + ("hello" + "hello")))))) {
            case "\n\t":
                var v19 int8 = 42
                _ = v19
            case "a":
                switch "\n\t" {
                case "\n\t":
                    for fi20 := 0; fi20 < 4; fi20++ {
                        v21 := GenS1{F0: 1, F1: v2}
                        _ = fi20
                        _ = v21
                    }
                }
            case "foo bar":
                v22 := []int{(v4 - v4)}
                _ = v22
            default:
                if v2 {
                    v23 := GenS2{F0: (v17 >= v17), F1: 153}
                    _ = v23
                } else {
                    v17 += -1
                    var v24 bool = (128 < 1)
                    v25 := GenS3{F0: "hello"}
                    _ = v24
                    _ = v25
                }
            }
            var v26 int8 = 1
            _ = fi18
            _ = v26
        }
    }
    v27 := GenS4{F0: "\n\t", F1: ("foo bar" + "hello"), F2: v4}
    v3 = 42
    var v28 uint16 = 202
    if ((v1 % ((42) | 1)) >= (v1 & 1)) {
        v3 = (^(v3))
    }
    var v29 uint = 2
    for fi30 := 0; fi30 < 17; fi30++ {
        v31 := GenS5{F0: "\n\t"}
        v32 := GenS5{F0: ""}
        _ = fi30
        _ = v31
        _ = v32
    }
    if ("hello" != (("hello" + ("foo bar" + "\n\t")) + "a")) {
        v29 = v29
    } else {
        v33 := GenS5{F0: "a"}
        v34 := GenS1{F0: 45, F1: ((0 != 1) && (-2147483648 <= 825))}
        v35 := GenS4{F0: ("\n\t" + "\n\t"), F1: "hello", F2: (v4 ^ (v4 ^ v4))}
        _ = v33
        _ = v34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
