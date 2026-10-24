package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 int32 }
func genFn0(p0 bool, p1 int32) bool {
    v5 := []int{32768}
    var v6 int = 32767
    var v7 int = 1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return (913 < 640)
}

func main() {
    var v1 uint32 = 2
    var v2 string = "foo bar"
    v3 := make(map[string]string)
    v4 := []int{567, 477, 9223372036854775807, 0}
    v8 := GenS1{F0: -9223372036854775808, F1: true, F2: ((v2 + (v2 + ("" + (("" + v2) + (("" + "foo bar") + ((((v2 + (v2 + ("" + v2))) + v2) + "") + ((("" + "hello") + v2) + "\n\t"))))))) + v2)}
    var v9 uint32 = (v1 / ((0) | 1))
    var v10 int8 = 16
    v11 := GenS1{F0: 42, F1: true, F2: v2}
    v12 := GenS1{F0: 42, F1: (457 >= 108), F2: ""}
    switch 127 {
    case 128:
        if len(v4) > 0 {
            v4[0] = -2147483648
        }
    }
    v3[(v2 + "hello")] = v2
    var v13 float64 = 0.0
    var v14 bool = (uint16(v1) < 42)
    for fi15 := 0; fi15 < 12; fi15++ {
        v16 := GenS2{F0: true}
        if v14 {
            switch (fi15 + 356) {
            case -2147483648:
                if false {
                    switch (fi15 % ((fi15) | 1)) {
                    case -9223372036854775808:
                        switch fi15 {
                        case 370:
                            v17 := v3[(v2 + v2)]
                            _ = v17
                        case 2147483647:
                            delete(v3, (v2 + ("" + v2)))
                        }
                    case 127:
                        v18 := GenS2{F0: true}
                        _ = v18
                    }
                    v19 := GenS2{F0: true}
                    v3["a"] = v2
                    _ = v19
                } else {
                    for fi20 := 0; fi20 < 19; fi20++ {
                        if ("\n\t" == ((("\n\t" + v2) + v2) + "a")) {
                            delete(v3, v2)
                            v14 = v14
                        } else {
                            if true {
                                if (!(v14 && v14)) {
                                    switch v2 {
                                    case "foo bar":
                                        var v21 uint32 = (v9 * 957)
                                        _ = v21
                                    }
                                    switch fi20 {
                                    case 277:
                                        var v22 uint8 = 55
                                        _ = v22
                                    case 197:
                                        v23 := GenS2{F0: v14}
                                        _ = v23
                                    }
                                }
                            }
                            v24 := GenS3{F0: -9223372036854775808, F1: 1}
                            v25 := GenS3{F0: -1, F1: 634}
                            _ = v24
                            _ = v25
                        }
                        _ = fi20
                    }
                }
            case 9223372036854775807:
                var v26 int32 = int32(fi15)
                _ = v26
            default:
                fmt.Println(v14)
            }
        } else {
            for fi27 := 0; fi27 < 17; fi27++ {
                v28 := GenS3{F0: -968, F1: -2147483648}
                var v29 float64 = ((float64((1.5 - 0.0)) / v13) + float64(0.1))
                _ = fi27
                _ = v28
                _ = v29
            }
        }
        _ = fi15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
