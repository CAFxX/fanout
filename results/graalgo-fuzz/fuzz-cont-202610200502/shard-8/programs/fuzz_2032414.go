package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 uint16 }

func main() {
    var v1 int32 = 2147483647
    v2 := []string{(("foo bar" + ("hello" + ("" + "\n\t"))) + ("foo bar" + "\n\t")), "\n\t"}
    var v3 int = 926
    var v4 uint16 = 0
    v5 := make(map[string]uint32)
    v1 = v1
    v6 := GenS1{F0: (v3 >> ((255) & 63))}
    fmt.Println(v3, v4)
    if true {
        for fi7 := 0; fi7 < 4; fi7++ {
            if len(v2) > 0 {
                v2[0] = "hello"
            }
            _ = fi7
        }
        fmt.Println(v3)
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 uint = 18446744073709551615
            _ = fi8
            _ = v9
        }
    }
    v3 -= 127
    var v10 float64 = 0.1
    switch "a" {
    case "foo bar":
        fmt.Println(v3, v1)
    case "hello":
        if (v3 <= v3) {
            var v11 float64 = v10
            var v12 int32 = v1
            var v13 bool = (96 < 18446744073709551615)
            _ = v11
            _ = v12
            _ = v13
        }
    case "\n\t":
        v5["a"] = 1
    }
    v2 = append(v2, "hello")
    v10 -= 0.0
    v3 -= v3
    fmt.Println(v1, v10, v10)
    _ = len(v5)
    delete(v5, "a")
    var v14 uint8 = 198
    v15 := GenS2{F0: 2, F1: 0, F2: 124}
    var v16 bool = (462 >= 42)
    switch "" {
    case "a":
        v4 = uint16(v14)
    case "foo bar":
        v17 := GenS1{F0: (-(v3))}
        _ = v17
    case "\n\t":
        v4 -= (v4 & (v4 & v4))
    }
    if (!(0.1 != (1.5 * -2.25))) {
        v1 = int32(v4)
        _ = len(v2)
        if ("" <= "\n\t") {
            if (!(float32(float64(1.5)) != float32((float64((3.14159 + 0.1)) * ((-2.25 / v10) * v10))))) {
                v14 -= (v14 / ((9) | 1))
                v4 += v4
                if ((!(v16 && true)) || (1 <= uint32(v4))) {
                    fmt.Println(v4)
                    v18 := GenS1{F0: v3}
                    switch ("a" + "foo bar") {
                    case "a":
                        v10 -= v10
                    default:
                        var v19 string = ""
                        _ = v19
                    }
                    _ = v18
                } else {
                    v3 += v3
                    v1 -= -12
                }
            } else {
                v20 := GenS2{F0: 4294967295, F1: 558, F2: 2}
                _ = v20
            }
            switch (v3 * (^(v3))) {
            case 128:
                v21 := v5["a"]
                _ = v21
            case -2147483648:
                for fi22 := 0; fi22 < 14; fi22++ {
                    switch "" {
                    case "a":
                        var v23 uint64 = 1
                        _ = v23
                    case "":
                        v24 := GenS3{F0: -9223372036854775808}
                        _ = v24
                    case "hello":
                        var v25 int64 = 0
                        _ = v25
                    default:
                        var v26 int32 = v1
                        _ = v26
                    }
                    switch "" {
                    case "\n\t":
                        v16 = false
                    case "foo bar":
                        v27 := GenS1{F0: 2147483647}
                        _ = v27
                    case "":
                        v10 -= float64(((((0.1 * -2.25) + float32(1.5)) + 0.0) / 1e10))
                    default:
                        if true {
                            switch "hello" {
                            case "\n\t":
                                for fi28 := 0; fi28 < 10; fi28++ {
                                    v1 = v1
                                    if v16 {
                                        var v29 float32 = 3.14159
                                        if ((712 != 4294967295) || (!(((1e10 / (3.14159 * (1e10 / 1.5))) / (3.14159 + (1.5 - v29))) != (v29 * v29)))) {
                                            v30 := GenS4{F0: fi22, F1: v4}
                                            v31 := GenS4{F0: fi22, F1: v4}
                                            _ = v30
                                            _ = v31
                                        }
                                        v10 -= (0.1 - 0.0)
                                        _ = v29
                                    } else {
                                        var v32 bool = (v16 || v16)
                                        _ = v32
                                    }
                                    _ = fi28
                                }
                            case "a":
                                var v33 int64 = -658
                                _ = v33
                            }
                        } else {
                            v34 := GenS4{F0: 32768, F1: (^(v4))}
                            _ = len(v2)
                            _ = v34
                        }
                    }
                    v4 = (v4 | (v4 * 1))
                    _ = fi22
                }
            }
        }
    } else {
        fmt.Println(v14, v4)
        fmt.Println(v14)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
