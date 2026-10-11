package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 uint8 = 241
    var v2 int32 = -1
    v3 := make(map[string]uint32)
    v4 := GenS1{F0: 32767}
    v5 := GenS1{F0: -475}
    v6 := GenS1{F0: 42}
    var v7 uint16 = 355
    if false {
        var v8 int32 = (^(v2))
        v2 += (v2 / ((v2) | 1))
        for fi9 := 0; fi9 < 16; fi9++ {
            v10 := []int{int(v7), fi9}
            v11 := v3["hello"]
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = v8
    } else {
        if (127 == v1) {
            var v12 bool = true
            fmt.Println(v12, v7)
            switch "" {
            case "\n\t":
                var v13 bool = true
                _ = v13
            }
            _ = v12
        }
        var v14 bool = (!true)
        _ = v14
    }
    v15 := []int{-385, 339, -343, 716}
    if (1 < -576) {
        var v16 float64 = 3.14159
        for fi17 := 0; fi17 < 18; fi17++ {
            if false {
                var v18 int8 = 127
                _ = v18
            } else {
                var v19 uint = 1
                _ = v19
            }
            switch int(v7) {
            case 894:
                v2 = v2
            case 874:
                for fi20 := 0; fi20 < 2; fi20++ {
                    if true {
                        switch ("hello" + ("hello" + "")) {
                        case "a":
                            v15 = append(v15, -9223372036854775808)
                        default:
                            v16 = float64(-2.25)
                        }
                        if false {
                            if true {
                                var v21 string = "a"
                                _ = v21
                            }
                            var v22 string = "foo bar"
                            if false {
                                var v23 float32 = 0.0
                                _ = len(v3)
                                _ = v23
                            } else {
                                v16 = v16
                                v24 := GenS1{F0: 1000}
                                if false {
                                    var v25 int16 = 32767
                                    v26 := GenS1{F0: v25}
                                    _ = v25
                                    _ = v26
                                }
                                _ = v24
                            }
                            v27 := GenS1{F0: -998}
                            _ = v22
                            _ = v27
                        } else {
                            switch "\n\t" {
                            case "hello":
                                var v28 uint8 = (^(v1))
                                _ = v28
                            case "\n\t":
                                v29 := GenS1{F0: 502}
                                _ = v29
                            case "":
                                v3[(("a" + "a") + (("a" + "hello") + "hello"))] = 2147483648
                            default:
                                var v30 uint32 = 42
                                _ = v30
                            }
                            fmt.Println(v1)
                        }
                    }
                    var v31 bool = true
                    var v32 float64 = v16
                    _ = fi20
                    _ = v31
                    _ = v32
                }
            case 4294967295:
                for fi33 := 0; fi33 < 6; fi33++ {
                    v1 = v1
                    if true {
                        fmt.Println(fi17, fi17, fi17)
                        fmt.Println(v2)
                    } else {
                        var v34 float32 = -2.25
                        _ = v34
                    }
                    _ = fi33
                }
            }
            _ = fi17
        }
        _ = v16
    } else {
        delete(v3, ("hello" + "\n\t"))
    }
    v1 -= v1
    v35 := GenS1{F0: -1}
    if len(v15) > 0 {
        v15[0] = 1
    }
    if true {
        var v36 uint32 = 0
        _ = v36
    }
    switch 42 {
    case 491:
        v37 := GenS1{F0: -355}
        _ = v37
    default:
        for fi38 := 0; fi38 < 17; fi38++ {
            v3["\n\t"] = 65535
            v7 += (v7 << ((0) & 63))
            _ = fi38
        }
    }
    for fi39 := 0; fi39 < 8; fi39++ {
        var v40 float32 = 0.0
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
