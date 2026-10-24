package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }

func main() {
    var v1 uint = 279
    var v2 string = ("\n\t" + (("\n\t" + ("a" + "\n\t")) + ("hello" + ("\n\t" + "hello"))))
    v3 := make(map[string]int)
    var v4 uint8 = 0
    v5 := make(map[string]string)
    v6 := []uint32{42, 4294967295, 115}
    var v7 uint64 = 18446744073709551615
    fmt.Println(v4)
    var v8 float64 = (((0.1 * 1.5) + float64((0.0 + -2.25))) / -2.25)
    if len(v6) > 0 {
        v6[0] = 2
    }
    fmt.Println(v8, v1, v8)
    v7 -= (v7 - (v7 & (v7 | 2)))
    var v9 uint16 = 2
    var v10 float64 = ((-2.25 - float64(1.5)) / (1.5 / 1.5))
    fmt.Println(v2)
    _ = len(v5)
    v2 = "foo bar"
    for fi11 := 0; fi11 < 15; fi11++ {
        switch v2 {
        case "a":
            if false {
                v8 += 3.14159
                v3["\n\t"] = fi11
            }
        case "hello":
            v12 := GenS1{F0: uint16(v1)}
            _ = v12
        case "\n\t":
            switch fi11 {
            case 4294967295:
                if false {
                    if (59 == -15) {
                        v13 := GenS1{F0: (v9 / (((v9 ^ (v9 - (^(v9))))) | 1))}
                        delete(v3, "foo bar")
                        v2 = (v2 + ("" + "hello"))
                        _ = v13
                    }
                    var v14 string = (("a" + ("a" + "")) + (v2 + ("a" + "hello")))
                    fmt.Println(v14, fi11, v10)
                    _ = v14
                }
            }
        default:
            var v15 int64 = 0
            _ = v15
        }
        var v16 uint32 = 4294967295
        _ = fi11
        _ = v16
    }
    var v17 int64 = 9223372036854775807
    switch 383 {
    case 2147483647:
        for fi18 := 0; fi18 < 4; fi18++ {
            v19 := GenS1{F0: 2}
            var v20 float32 = (0.1 + 0.0)
            _ = fi18
            _ = v19
            _ = v20
        }
    case 310:
        switch v2 {
        case "foo bar":
            if (4294967295 == uint(v7)) {
                if false {
                    for fi21 := 0; fi21 < 4; fi21++ {
                        v22 := GenS1{F0: v9}
                        for fi23 := 0; fi23 < 14; fi23++ {
                            v7 += (v7 & v7)
                            _ = fi23
                        }
                        _ = fi21
                        _ = v22
                    }
                } else {
                    for fi24 := 0; fi24 < 16; fi24++ {
                        var v25 bool = false
                        _ = fi24
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 15; fi26++ {
                        for fi27 := 0; fi27 < 10; fi27++ {
                            if true {
                                var v28 float64 = 3.14159
                                _ = v28
                            }
                            if (!(v4 > v4)) {
                                v4 += v4
                                if false {
                                    if ("foo bar" <= "") {
                                        var v29 uint16 = 65535
                                        v30 := GenS2{F0: true}
                                        var v31 int64 = v17
                                        _ = v29
                                        _ = v30
                                        _ = v31
                                    }
                                }
                            }
                            _ = fi27
                        }
                        _ = fi26
                    }
                    v1 += v1
                }
                var v32 string = v2
                _ = v32
            }
        }
    default:
        v4 += (v4 - (^(v4)))
    }
    fmt.Println(v4)
    v17 += -605
    if true {
        fmt.Println(v2, v9)
        v33 := v3[v2]
        v34 := GenS2{F0: (506 == 2147483647)}
        _ = v33
        _ = v34
    } else {
        if true {
            var v35 float64 = (((v10 / v8) / 1e10) * float64(3.14159))
            if (-40 <= 69) {
                v2 = v2
                var v36 string = (("hello" + (("hello" + (v2 + "foo bar")) + ("" + v2))) + v2)
                _ = v36
            }
            _ = v35
        }
        if (v8 != v8) {
            _ = len(v6)
            switch v2 {
            case "a":
                fmt.Println(v2)
            }
        } else {
            v1 = (^(v1))
        }
        v8 += (v8 * (v10 - (-2.25 - float64(((float32((v10 * v10)) - (0.1 * 1e10)) * ((float32(v10) - (0.1 - -2.25)) * (1e10 + (0.0 + 3.14159))))))))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
