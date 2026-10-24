package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int8; F1 uint64; F2 uint32 }

func main() {
    var v1 float64 = (0.1 * 0.1)
    var v2 float32 = float32(v1)
    var v3 int16 = 1000
    var v4 uint32 = 2147483647
    var v5 bool = false
    v6 := make(map[string]string)
    v7 := v6["foo bar"]
    var v8 string = (v7 + (v7 + (v7 + (("foo bar" + v7) + ("" + v7)))))
    v3 -= 0
    v8 = v8
    for fi9 := 0; fi9 < 13; fi9++ {
        if (100 <= 110) {
            var v10 int16 = v3
            var v11 uint32 = v4
            fmt.Println(v3, v3, v3)
            _ = v10
            _ = v11
        }
        fmt.Println(fi9, v1)
        _ = fi9
    }
    var v12 uint8 = 50
    v13 := GenS1{F0: v8, F1: 42}
    if (v1 > (v1 * v1)) {
        v14 := GenS1{F0: "\n\t", F1: 1}
        switch 0 {
        case 1:
            if v5 {
                var v15 bool = (!(v5 && v5))
                v8 = v7
                v16 := []int{147}
                _ = v15
                _ = v16
            }
        case -1:
            if ((v12 + v12) > v12) {
                var v17 int8 = -128
                _ = v17
            } else {
                var v18 float64 = v1
                var v19 uint32 = (^(v4))
                _ = v18
                _ = v19
            }
        case 32768:
            if ((0.1 / (((-2.25 - v1) - 0.0) + (v1 / v1))) <= 1.5) {
                if (!v5) {
                    if v5 {
                        v20 := []int{42, 9223372036854775807, 4294967295, -307}
                        var v21 uint32 = 4294967295
                        _ = v20
                        _ = v21
                    }
                } else {
                    fmt.Println(v2, v7, v3)
                    for fi22 := 0; fi22 < 19; fi22++ {
                        switch v7 {
                        case "":
                            for fi23 := 0; fi23 < 4; fi23++ {
                                v24 := GenS1{F0: "\n\t", F1: 1}
                                var v25 uint16 = 32767
                                switch fi22 {
                                case 9223372036854775807:
                                    delete(v6, (("a" + v7) + ""))
                                case 4294967295:
                                    v26 := GenS2{F0: 102, F1: 42, F2: (v4 * (^(v4)))}
                                    _ = v26
                                case 2147483647:
                                    var v27 uint16 = 65535
                                    _ = v27
                                }
                                v28 := GenS1{F0: ("foo bar" + ("\n\t" + ((v7 + "foo bar") + (v8 + "\n\t")))), F1: 1}
                                fmt.Println(fi23, fi22)
                                _ = fi23
                                _ = v24
                                _ = v25
                                _ = v28
                            }
                        case "\n\t":
                            v29 := GenS1{F0: "foo bar", F1: 479}
                            _ = v29
                        case "a":
                            var v30 bool = ((v2 <= v2) || (1 <= 528))
                            _ = v30
                        default:
                            if (true || false) {
                                fmt.Println(v4, v4)
                                v4 = 1
                            } else {
                                switch (fi22 + int(v4)) {
                                case -679:
                                    var v31 string = "hello"
                                    _ = v31
                                case 423:
                                    for fi32 := 0; fi32 < 4; fi32++ {
                                        v8 = ""
                                        _ = fi32
                                    }
                                case -732:
                                    var v33 int16 = v3
                                    _ = v33
                                }
                                switch ("hello" + "a") {
                                case "":
                                    var v34 uint8 = (v12 >> ((4294967295) & 63))
                                    _ = v34
                                case "\n\t":
                                    var v35 uint16 = 42
                                    _ = v35
                                }
                            }
                        }
                        _ = fi22
                    }
                }
            }
        }
        var v36 int32 = 498
        _ = v14
        _ = v36
    }
    v37 := []int{127, 32767}
    var v38 int16 = v3
    v4 = (v4 & v4)
    v38 = v38
    var v39 int = -475
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
