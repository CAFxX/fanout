package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int }
type GenS2 struct { F0 int32; F1 uint8 }

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 0
    v3 := []int{32767}
    v4 := make(map[string]string)
    v5 := make(map[string]bool)
    var v6 uint16 = 1
    v6 = 89
    v2 = v2
    var v7 bool = ((!(false && ((!(false && false)) && true))) || (!(!(!(!(!(!(false || (-128 == 127)))))))))
    v5["foo bar"] = v7
    v6 = (v6 + v6)
    for fi8 := 0; fi8 < 15; fi8++ {
        v3 = append(v3, fi8)
        if (!true) {
            v6 = 32767
            if ((fi8 - int(v2)) >= (fi8 % (((fi8 | fi8)) | 1))) {
                if false {
                    if ("a" == "a") {
                        var v9 string = "foo bar"
                        v2 -= v2
                        for fi10 := 0; fi10 < 1; fi10++ {
                            fmt.Println(v9, fi8)
                            var v11 float32 = 1.5
                            fmt.Println(v11, v9, v11)
                            v11 = 0.0
                            var v12 uint64 = 172
                            _ = fi10
                            _ = v11
                            _ = v12
                        }
                        _ = v9
                    } else {
                        fmt.Println(fi8, fi8)
                    }
                } else {
                    if v7 {
                        for fi13 := 0; fi13 < 8; fi13++ {
                            if true {
                                var v14 bool = v7
                                fmt.Println(v7, v6)
                                _ = v14
                            }
                            _ = fi13
                        }
                        for fi15 := 0; fi15 < 11; fi15++ {
                            v6 = 410
                            _ = fi15
                        }
                    } else {
                        v16 := GenS1{F0: 0, F1: 232, F2: 409}
                        _ = v16
                    }
                }
                var v17 int16 = -1
                for fi18 := 0; fi18 < 16; fi18++ {
                    v2 += 1
                    _ = fi18
                }
                _ = v17
            }
            for fi19 := 0; fi19 < 16; fi19++ {
                if len(v3) > 0 {
                    v3[0] = fi8
                }
                _ = fi19
            }
        }
        _ = fi8
    }
    v6 = v6
    v20 := GenS1{F0: 1000, F1: 1, F2: 0}
    switch "\n\t" {
    case "a":
        v21 := GenS1{F0: 495, F1: 0, F2: 492}
        _ = v21
    case "hello":
        v6 = v6
    case "foo bar":
        if v7 {
            v22 := GenS2{F0: 2147483647, F1: 1}
            v3 = append(v3, 32767)
            for fi23 := 0; fi23 < 14; fi23++ {
                v24 := GenS1{F0: 389, F1: 255, F2: fi23}
                for fi25 := 0; fi25 < 16; fi25++ {
                    v5[(("" + ("foo bar" + "a")) + "")] = (1 >= 42)
                    _ = fi25
                }
                v2 = v2
                _ = fi23
                _ = v24
            }
            _ = v22
        } else {
            for fi26 := 0; fi26 < 6; fi26++ {
                v2 = (v2 % ((v2) | 1))
                v2 = v2
                for fi27 := 0; fi27 < 10; fi27++ {
                    v7 = v7
                    _ = fi27
                }
                _ = fi26
            }
            var v28 string = ((("hello" + "\n\t") + "a") + "foo bar")
            _ = v28
        }
    }
    v3 = append(v3, 103)
    fmt.Println(v7)
    switch (("foo bar" + "foo bar") + "") {
    case "foo bar":
        if (!false) {
            v7 = (0.1 != (1e10 + float64((-2.25 - 3.14159))))
        }
    default:
        var v29 int = 9223372036854775807
        _ = v29
    }
    v6 -= v6
    switch "foo bar" {
    case "foo bar":
        var v30 uint8 = 0
        _ = v30
    case "\n\t":
        var v31 float32 = 3.14159
        _ = v31
    default:
        v32 := v5[("foo bar" + "a")]
        _ = v32
    }
    delete(v5, "")
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
