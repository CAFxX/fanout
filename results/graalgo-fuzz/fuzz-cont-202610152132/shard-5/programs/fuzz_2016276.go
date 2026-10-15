package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 bool }

func main() {
    var v1 bool = (-868 == -32768)
    v2 := []string{"\n\t"}
    var v3 int32 = 2147483647
    var v4 float64 = float64(((1e10 - 1.5) * 1.5))
    var v5 int64 = -983
    var v6 uint8 = 1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi7 := 0; fi7 < 7; fi7++ {
        v5 = v5
        if true {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if (!true) {
                if false {
                    v4 += v4
                    var v8 int8 = 1
                    v4 = ((0.1 + (float64(((float32((v4 - (3.14159 / (3.14159 / v4)))) - 1e10) + 1.5)) / ((v4 * float64((float32(v4) - 0.0))) * v4))) / v4)
                    _ = v8
                }
                for fi9 := 0; fi9 < 3; fi9++ {
                    v4 += (v4 / 3.14159)
                    var v10 uint32 = 0
                    _ = fi9
                    _ = v10
                }
                if v1 {
                    fmt.Println(v5)
                    for fi11 := 0; fi11 < 6; fi11++ {
                        if len(v2) > 0 {
                            v2[0] = "a"
                        }
                        _ = fi11
                    }
                }
            }
            v4 += 0.0
        } else {
            v1 = true
        }
        _ = fi7
    }
    v6 -= v6
    for fi12 := 0; fi12 < 12; fi12++ {
        switch "" {
        case "hello":
            if (42 > -89) {
                if (v1 || true) {
                    switch (fi12 + (fi12 & 32768)) {
                    case 127:
                        if v1 {
                            if (56 > -128) {
                                switch ("a" + "\n\t") {
                                case "hello":
                                    var v13 uint32 = 122
                                    _ = v13
                                default:
                                    if ((v6 & (v6 + 127)) >= (v6 ^ v6)) {
                                        switch fi12 {
                                        case 650:
                                            var v14 string = "a"
                                            _ = v14
                                        }
                                    }
                                }
                                var v15 uint8 = v6
                                _ = v15
                            }
                            v3 = v3
                        }
                    case 128:
                        v16 := GenS1{F0: 194, F1: -32768, F2: v1}
                        _ = v16
                    case 986:
                        fmt.Println(v6, v6, v3)
                    }
                    if true {
                        fmt.Println(fi12, fi12, fi12)
                        switch fi12 {
                        case 32767:
                            fmt.Println(fi12, v5)
                        default:
                            if v1 {
                                var v17 uint64 = 971
                                _ = v17
                            } else {
                                var v18 bool = v1
                                for fi19 := 0; fi19 < 19; fi19++ {
                                    var v20 int16 = -629
                                    _ = fi19
                                    _ = v20
                                }
                                _ = v18
                            }
                        }
                    }
                }
                var v21 uint32 = 42
                if (fi12 != (^(fi12))) {
                    v2 = append(v2, "a")
                    var v22 int32 = -601
                    _ = v22
                }
                _ = v21
            }
        case "foo bar":
            v23 := GenS1{F0: 4294967295, F1: 42, F2: false}
            _ = v23
        case "":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            if false {
                v3 = v3
            } else {
                if ("a" > "foo bar") {
                    v4 = v4
                    v5 += v5
                }
                var v24 float64 = 1e10
                _ = v24
            }
        }
        _ = fi12
    }
    for fi25 := 0; fi25 < 10; fi25++ {
        for fi26 := 0; fi26 < 5; fi26++ {
            var v27 int32 = (v3 >> ((4294967295) & 63))
            fmt.Println(fi25, v27)
            if true {
                var v28 int32 = (v27 * (v27 % (((^(v27))) | 1)))
                _ = v28
            }
            if len(v2) > 0 {
                v2[0] = "hello"
            }
            _ = fi26
            _ = v27
        }
        var v29 int16 = 1000
        var v30 uint32 = 2
        _ = fi25
        _ = v29
        _ = v30
    }
    var v31 int32 = v3
    v1 = v1
    switch ("hello" + (("\n\t" + "") + "a")) {
    case "a":
        for fi32 := 0; fi32 < 15; fi32++ {
            var v33 uint16 = 0
            _ = fi32
            _ = v33
        }
    case "hello":
        for fi34 := 0; fi34 < 13; fi34++ {
            if len(v2) > 0 {
                v2[0] = (((("\n\t" + ("" + ("" + "\n\t"))) + "a") + ("\n\t" + "")) + ("hello" + ""))
            }
            _ = fi34
        }
    case "":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v35 := GenS1{F0: 42, F1: int16(v5), F2: v1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
