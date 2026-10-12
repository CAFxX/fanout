package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 int64; F1 int32 }
func genFn0(p0 uint8, p1 int16) string {
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 uint64 = 18446744073709551615
        _ = fi6
        _ = v7
    }
    var v8 string = "hello"
    switch v8 {
    case "":
        v9 := GenS1{F0: 694}
        _ = v9
    }
    var v10 int64 = -1
    var v11 uint = 753
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v8
    return "hello"
}

func main() {
    v1 := []string{"\n\t", "a"}
    var v2 string = "a"
    var v3 float64 = 1e10
    v4 := make(map[uint32]bool)
    v5 := []string{"foo bar", v2}
    var v12 float32 = 0.0
    v13 := GenS2{F0: 255}
    for fi14 := 0; fi14 < 10; fi14++ {
        fmt.Println(fi14, v12, v3)
        fmt.Println(v2)
        _ = fi14
    }
    v1 = append(v1, v2)
    if false {
        v12 += (0.0 / v12)
    } else {
        fmt.Println(v12, v3)
        switch 916 {
        case 971:
            switch -1 {
            case -364:
                v15 := GenS2{F0: 2}
                _ = v15
            }
        case -845:
            for fi16 := 0; fi16 < 7; fi16++ {
                if ((v2 + "") > v2) {
                    var v17 uint32 = 2147483648
                    if (fi16 >= 2147483647) {
                        var v18 uint8 = 127
                        for fi19 := 0; fi19 < 8; fi19++ {
                            fmt.Println(v12, v18, v17)
                            v18 += (v18 * v18)
                            var v20 float32 = float32(v3)
                            _ = fi19
                            _ = v20
                        }
                        fmt.Println(v18)
                        _ = v18
                    }
                    fmt.Println(fi16, fi16, v3)
                    _ = v17
                } else {
                    switch (fi16 >> ((1) & 63)) {
                    case 2147483647:
                        v21 := GenS3{F0: 2, F1: (65535 < 4294967295)}
                        _ = v21
                    default:
                        fmt.Println(fi16, v2)
                    }
                    v12 = 3.14159
                    switch "foo bar" {
                    case "foo bar":
                        var v22 bool = (("" + v2) >= (("" + v2) + "foo bar"))
                        _ = v22
                    case "":
                        if false {
                            v12 -= v12
                        }
                    case "\n\t":
                        fmt.Println(fi16, v12)
                    }
                }
                var v23 float64 = float64(float32(0.0))
                var v24 int = (fi16 * 2147483647)
                _ = fi16
                _ = v23
                _ = v24
            }
        case 32767:
            v4[12] = (!false)
        default:
            if (!(!(true || ((-9223372036854775808 > 9223372036854775807) || true)))) {
                var v25 uint64 = 200
                _ = v25
            } else {
                for fi26 := 0; fi26 < 7; fi26++ {
                    var v27 uint64 = 1
                    _ = fi26
                    _ = v27
                }
                var v28 int8 = 1
                v2 = (((v2 + "a") + v2) + v2)
                _ = v28
            }
        }
    }
    if ("a" != "\n\t") {
        v1 = append(v1, v2)
        v29 := GenS4{F0: (42 != 2), F1: v2}
        v30 := v4[429]
        _ = v29
        _ = v30
    } else {
        v5 = append(v5, ("a" + v2))
        if len(v1) > 0 {
            v1[0] = v2
        }
        var v31 int16 = 0
        _ = v31
    }
    var v32 int32 = 2147483647
    fmt.Println(v12, v3, v32)
    var v33 int32 = v32
    fmt.Println(v33)
    delete(v4, 1)
    if (v2 >= ("foo bar" + "a")) {
        fmt.Println(v12, v12, v2)
        var v34 int16 = 0
        fmt.Println(v33, v2, v3)
        _ = v34
    } else {
        v35 := GenS5{F0: 616, F1: v32}
        switch 448 {
        case 827:
            v36 := GenS2{F0: 262}
            _ = v36
        case 1:
            var v37 uint16 = 65535
            _ = v37
        case 32767:
            delete(v4, 0)
        }
        _ = v35
    }
    v38 := GenS4{F0: false, F1: ""}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v38
    _ = v4
    _ = v5
}
