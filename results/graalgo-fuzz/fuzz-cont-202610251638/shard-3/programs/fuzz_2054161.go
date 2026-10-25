package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }

func main() {
    var v1 int32 = 42
    v2 := []uint32{2147483648, 2147483648, 4294967295, 2147483647}
    var v3 uint16 = 0
    var v4 float32 = (-2.25 + (1e10 - 0.0))
    if false {
        if (false || (2147483647 >= v1)) {
            var v5 uint8 = 254
            if false {
                if false {
                    var v6 int8 = 1
                    var v7 bool = true
                    _ = v6
                    _ = v7
                } else {
                    var v8 string = "\n\t"
                    _ = v8
                }
                switch ("foo bar" + ("\n\t" + "hello")) {
                case "":
                    v2 = append(v2, 2147483648)
                case "foo bar":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case "a":
                    for fi9 := 0; fi9 < 14; fi9++ {
                        if ("foo bar" == "") {
                            for fi10 := 0; fi10 < 18; fi10++ {
                                var v11 float32 = float32(-2.25)
                                _ = fi10
                                _ = v11
                            }
                            fmt.Println(fi9, v5, fi9)
                            var v12 uint16 = (v3 >> ((638) & 63))
                            _ = v12
                        }
                        _ = fi9
                    }
                default:
                    for fi13 := 0; fi13 < 10; fi13++ {
                        v4 = -2.25
                        _ = fi13
                    }
                }
                for fi14 := 0; fi14 < 18; fi14++ {
                    var v15 bool = true
                    v16 := GenS1{F0: ((((("hello" + "a") + "a") + (("hello" + ("" + "a")) + "\n\t")) + "hello") + ((("foo bar" + "\n\t") + "") + "a")), F1: "foo bar"}
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            } else {
                if (((128 > 1) && ((v3 == v3) && (!true))) && (!false)) {
                    fmt.Println(v3)
                    fmt.Println(v4, v5)
                    v5 += v5
                }
                for fi17 := 0; fi17 < 18; fi17++ {
                    v4 = (v4 * (v4 - v4))
                    _ = len(v2)
                    _ = fi17
                }
            }
            var v18 uint32 = 42
            _ = v18
            _ = v5
        } else {
            v3 = (^(v3))
        }
        fmt.Println(v1)
    }
    v1 -= 53
    v19 := GenS1{F0: (("hello" + "") + ("foo bar" + ("\n\t" + (("a" + "a") + "\n\t")))), F1: ("a" + (("foo bar" + "\n\t") + "\n\t"))}
    var v20 int8 = 127
    for fi21 := 0; fi21 < 12; fi21++ {
        v1 = (v1 << ((2) & 63))
        _ = fi21
    }
    var v22 uint8 = 127
    var v23 float32 = v4
    if len(v2) > 0 {
        v2[0] = 153
    }
    var v24 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
