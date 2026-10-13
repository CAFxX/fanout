package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 int; F2 int8 }
type GenS4 struct { F0 uint16; F1 uint64 }

func main() {
    var v1 bool = (!false)
    var v2 float32 = 3.14159
    var v3 int16 = 32767
    v4 := []uint32{318, 65535}
    if v1 {
        var v5 int64 = -771
        var v6 string = "foo bar"
        v1 = v1
        _ = v5
        _ = v6
    } else {
        v2 -= float32((float64(1e10) - 3.14159))
        for fi7 := 0; fi7 < 10; fi7++ {
            v3 += 1000
            for fi8 := 0; fi8 < 5; fi8++ {
                fmt.Println(fi7, fi7)
                if (!(v1 && (0 <= 255))) {
                    switch fi7 {
                    case 0:
                        v9 := GenS1{F0: 1}
                        _ = v9
                    case 32767:
                        v10 := GenS1{F0: 1}
                        _ = v10
                    }
                } else {
                    var v11 bool = true
                    _ = v11
                }
                if true {
                    v2 = (v2 / (1.5 + v2))
                    var v12 uint8 = 1
                    var v13 float32 = (v2 + -2.25)
                    v14 := GenS2{F0: 4294967295}
                    _ = v12
                    _ = v13
                    _ = v14
                } else {
                    v2 += (1e10 - float32(0.0))
                    v15 := GenS1{F0: 1}
                    if (1 < fi7) {
                        var v16 uint16 = 32767
                        var v17 float32 = v2
                        _ = v16
                        _ = v17
                    } else {
                        var v18 float32 = -2.25
                        _ = v18
                    }
                    var v19 int8 = -124
                    switch ("foo bar" + "") {
                    case "foo bar":
                        v20 := GenS1{F0: 128}
                        _ = v20
                    case "\n\t":
                        fmt.Println(v1, v2, v1)
                    default:
                        if ((!true) || v1) {
                            var v21 int16 = v3
                            var v22 uint = 18446744073709551615
                            _ = v21
                            _ = v22
                        } else {
                            v4 = append(v4, 1)
                            fmt.Println(v3)
                            if len(v4) > 0 {
                                v4[0] = 4294967295
                            }
                            var v23 int32 = 0
                            v24 := GenS2{F0: 440}
                            _ = v23
                            _ = v24
                        }
                    }
                    if ((1 >= 2) && ((!(42 >= 861)) || v1)) {
                        fmt.Println(v19, fi8)
                        var v25 uint16 = 428
                        _ = v25
                    }
                    v3 -= 167
                    _ = v15
                    _ = v19
                }
                _ = fi8
            }
            _ = fi7
        }
        v3 = (-(v3))
    }
    if true {
        v3 = v3
    } else {
        v3 = (v3 ^ 363)
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        v27 := GenS3{F0: 137, F1: (-(fi26)), F2: 127}
        fmt.Println(v3)
        _ = fi26
        _ = v27
    }
    var v28 uint8 = 20
    for fi29 := 0; fi29 < 4; fi29++ {
        v30 := GenS4{F0: 1, F1: 18446744073709551615}
        _ = fi29
        _ = v30
    }
    if (true && (!v1)) {
        var v31 uint64 = 198
        _ = v31
    } else {
        v28 += 255
        v32 := GenS1{F0: v28}
        v33 := GenS4{F0: 0, F1: 650}
        _ = v32
        _ = v33
    }
    _ = len(v4)
    switch ((("a" + ("" + (("" + "hello") + ("hello" + ((((("hello" + "a") + ("a" + "\n\t")) + "\n\t") + "a") + "a"))))) + (("" + ("" + ("foo bar" + ("\n\t" + ("foo bar" + "\n\t"))))) + "hello")) + "hello") {
    case "":
        var v34 int8 = 78
        _ = v34
    case "\n\t":
        var v35 uint16 = 402
        _ = v35
    case "a":
        var v36 uint = 721
        _ = v36
    default:
        fmt.Println(v3, v1, v28)
    }
    v37 := GenS4{F0: 92, F1: 42}
    var v38 uint16 = 32767
    var v39 string = "a"
    var v40 float64 = 0.1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v3, v2)
    v41 := GenS4{F0: v38, F1: 4294967295}
    fmt.Println(v28)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
