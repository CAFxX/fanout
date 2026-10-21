package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 bool; F1 int64; F2 string }
type GenS3 struct { F0 int32; F1 int64 }
func genFn0() uint16 {
    var v5 uint8 = 42
    v5 = (v5 & 127)
    if (-2.25 < 0.1) {
        v6 := GenS1{F0: 42, F1: (("hello" + "a") + ("\n\t" + ""))}
        _ = v6
    } else {
        var v7 uint8 = v5
        v8 := GenS2{F0: true, F1: 479, F2: "a"}
        if false {
            v9 := []int{-614}
            _ = v9
        }
        _ = v7
        _ = v8
    }
    if (99 != 115) {
        fmt.Println(v5)
        fmt.Println(v5)
    } else {
        var v10 int32 = 2147483647
        var v11 int32 = v10
        v12 := GenS2{F0: (!true), F1: 0, F2: "hello"}
        _ = v10
        _ = v11
        _ = v12
    }
    _ = v5
    return 32767
}

func main() {
    var v1 bool = (1 == 128)
    var v2 int = -9223372036854775808
    var v3 float64 = 1e10
    var v4 uint = 42
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 bool = false
        _ = fi13
        _ = v14
    }
    v2 = 42
    v15 := []int{v2}
    var v16 uint32 = 0
    fmt.Println(v3, v3)
    if len(v15) > 0 {
        v15[0] = (v2 % (((v2 / ((-411) | 1))) | 1))
    }
    v17 := GenS2{F0: v1, F1: -88, F2: ""}
    var v18 string = ("\n\t" + ((("foo bar" + "a") + (((((("a" + ("\n\t" + "hello")) + (("hello" + ("\n\t" + "a")) + "foo bar")) + ("foo bar" + "hello")) + "") + ("foo bar" + "\n\t")) + (("foo bar" + "") + ""))) + "hello"))
    v1 = v1
    for fi19 := 0; fi19 < 20; fi19++ {
        v1 = false
        _ = fi19
    }
    if ((!false) || true) {
        switch v2 {
        case 2147483647:
            if true {
                var v20 uint = v4
                v16 = (v16 / ((4294967295) | 1))
                _ = v20
            }
        case -297:
            for fi21 := 0; fi21 < 20; fi21++ {
                var v22 float64 = (v3 * v3)
                if (1 < 0) {
                    v18 = (v18 + v18)
                    if ((v2 | (fi21 % (((fi21 - 127)) | 1))) <= v2) {
                        if v1 {
                            if true {
                                var v23 int8 = 46
                                v24 := GenS3{F0: -1, F1: 9223372036854775807}
                                _ = v23
                                _ = v24
                            } else {
                                v2 = (v2 / ((1) | 1))
                                fmt.Println(v22)
                            }
                        }
                    }
                } else {
                    v15 = append(v15, (fi21 + fi21))
                    v25 := GenS3{F0: 42, F1: 243}
                    _ = v25
                }
                _ = fi21
                _ = v22
            }
        }
        switch "hello" {
        case "":
            v1 = (!(2 >= 722))
        default:
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
        }
    }
    v26 := GenS2{F0: v1, F1: -1, F2: v18}
    if (-82 <= int8(v16)) {
        for fi27 := 0; fi27 < 8; fi27++ {
            var v28 bool = true
            v15 = append(v15, (fi27 >> (((^(v4))) & 63)))
            _ = fi27
            _ = v28
        }
        var v29 uint16 = 32767
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
