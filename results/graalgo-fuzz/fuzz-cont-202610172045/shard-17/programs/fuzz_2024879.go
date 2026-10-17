package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() int {
    var v7 uint32 = 990
    var v8 uint32 = v7
    fmt.Println(v7, v8)
    _ = v7
    _ = v8
    return 32767
}

func main() {
    v1 := []string{"hello", "hello", "", ""}
    var v2 uint16 = 42
    var v3 uint32 = 59
    v4 := []uint32{v3, 4294967295, v3, 65535}
    var v5 string = "\n\t"
    var v6 string = v5
    for fi9 := 0; fi9 < 9; fi9++ {
        for fi10 := 0; fi10 < 19; fi10++ {
            switch (fi10 / ((genFn0()) | 1)) {
            case 42:
                switch ("\n\t" + "\n\t") {
                case "foo bar":
                    var v11 float32 = ((-2.25 * -2.25) / (((3.14159 / float32(-2.25)) * ((3.14159 - (0.0 - 0.1)) / 3.14159)) / (float32(float64(float32((0.0 + 0.0)))) - 3.14159)))
                    _ = v11
                case "hello":
                    var v12 int32 = 0
                    _ = v12
                }
            case 4294967295:
                v6 = ((v5 + v5) + (v6 + (v6 + (((("\n\t" + (v5 + v5)) + (v5 + v5)) + ("a" + v6)) + v5))))
            default:
                if ((9223372036854775807 <= -1) || (!false)) {
                    _ = len(v1)
                    if false {
                        if ((v3 - (v3 - v3)) >= (v3 << ((255) & 63))) {
                            v6 = "\n\t"
                            var v13 uint16 = (^(v2))
                            var v14 int8 = 127
                            _ = v13
                            _ = v14
                        }
                        var v15 float64 = -2.25
                        v16 := GenS1{F0: 920}
                        _ = v15
                        _ = v16
                    } else {
                        var v17 int = fi9
                        var v18 string = (v5 + v5)
                        v17 -= (v17 | (fi9 / ((fi9) | 1)))
                        _ = v17
                        _ = v18
                    }
                    if (v2 != (v2 - uint16(v3))) {
                        var v19 int32 = -2147483648
                        v3 = 1
                        var v20 bool = true
                        _ = v19
                        _ = v20
                    }
                } else {
                    if true {
                        fmt.Println(v3, v5, fi9)
                    }
                    if (0 > 131) {
                        if (v3 <= uint32(v2)) {
                            _ = len(v1)
                            v5 = "\n\t"
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v2 += v2
                            var v21 int64 = 784
                            _ = v21
                        }
                    } else {
                        if (!(330 == 1)) {
                            switch "foo bar" {
                            case "foo bar":
                                switch fi10 {
                                case 9223372036854775807:
                                    fmt.Println(v5, fi9)
                                case -1:
                                    v22 := GenS1{F0: 90}
                                    _ = v22
                                default:
                                    v2 -= v2
                                }
                            case "\n\t":
                                v5 = v6
                            case "a":
                                var v23 uint16 = 350
                                _ = v23
                            }
                            var v24 bool = true
                            var v25 int = 128
                            _ = v24
                            _ = v25
                        }
                    }
                    v3 = 65535
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v6 = ((((v5 + ("foo bar" + v5)) + v6) + "") + v6)
            _ = fi10
        }
        v3 -= v3
        _ = fi9
    }
    fmt.Println(v3, v6)
    v6 = v6
    for fi26 := 0; fi26 < 18; fi26++ {
        v27 := GenS1{F0: 2}
        _ = fi26
        _ = v27
    }
    var v28 uint32 = (v3 % (((v3 * 4294967295)) | 1))
    if false {
        fmt.Println(v2, v5, v6)
    }
    for fi29 := 0; fi29 < 1; fi29++ {
        if len(v4) > 0 {
            v4[0] = v3
        }
        var v30 string = "a"
        var v31 uint16 = (v2 >> ((170) & 63))
        _ = fi29
        _ = v30
        _ = v31
    }
    switch v6 {
    case "\n\t":
        v5 = (v5 + (v6 + v6))
    default:
        var v32 float64 = 3.14159
        _ = v32
    }
    fmt.Println(v28, v2)
    var v33 bool = true
    var v34 int = genFn0()
    if len(v4) > 0 {
        v4[0] = v28
    }
    var v35 int8 = 1
    v35 = v35
    var v36 int32 = -2147483648
    v6 = v5
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
