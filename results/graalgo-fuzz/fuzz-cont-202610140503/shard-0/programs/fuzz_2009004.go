package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }

func main() {
    var v1 uint16 = 266
    var v2 int32 = -982
    v3 := make(map[int]int)
    v4 := GenS1{F0: -9223372036854775808}
    v1 -= (^(v1))
    v5 := GenS2{F0: (("" + "a") + ("a" + (("hello" + "foo bar") + "hello")))}
    var v6 int64 = -9223372036854775808
    var v7 float64 = 0.1
    var v8 uint64 = 4294967295
    v9 := GenS1{F0: (^(v6))}
    for fi10 := 0; fi10 < 19; fi10++ {
        for fi11 := 0; fi11 < 20; fi11++ {
            switch 4294967295 {
            case 9223372036854775807:
                v2 = int32(v6)
            }
            v12 := []int{127, fi10, -543, 9223372036854775807}
            if (!true) {
                var v13 uint8 = 255
                fmt.Println(fi11)
                _ = v13
            } else {
                if false {
                    fmt.Println(fi10, v1)
                }
                switch "foo bar" {
                case "\n\t":
                    v12 = append(v12, fi11)
                default:
                    if true {
                        if false {
                            var v14 uint64 = v8
                            v15 := GenS2{F0: ("hello" + "")}
                            var v16 float64 = ((((v7 + float64(1.5)) / 0.1) + v7) / 1.0)
                            _ = v14
                            _ = v15
                            _ = v16
                        } else {
                            v17 := GenS2{F0: ((((("hello" + ("" + (("hello" + "\n\t") + "foo bar"))) + "a") + "\n\t") + (("hello" + ("foo bar" + ("\n\t" + "foo bar"))) + "")) + "a")}
                            v6 = (v6 >> ((0) & 63))
                            if true {
                                if len(v12) > 0 {
                                    v12[0] = (fi10 & -526)
                                }
                            } else {
                                var v18 uint64 = 18446744073709551615
                                v19 := GenS1{F0: 42}
                                _ = v18
                                _ = v19
                            }
                            _ = v17
                        }
                        delete(v3, 128)
                        v20 := GenS2{F0: "hello"}
                        _ = v20
                    } else {
                        v21 := v3[int(v8)]
                        v7 += v7
                        var v22 uint = 18446744073709551615
                        _ = v21
                        _ = v22
                    }
                }
            }
            v8 = (^(v8))
            _ = fi11
            _ = v12
        }
        v23 := v3[fi10]
        _ = fi10
        _ = v23
    }
    switch "a" {
    case "":
        for fi24 := 0; fi24 < 15; fi24++ {
            var v25 int16 = 374
            var v26 uint32 = 65535
            switch (fi24 / (((fi24 + 0)) | 1)) {
            case 128:
                v27 := GenS2{F0: "\n\t"}
                _ = v27
            case -9223372036854775808:
                v3[fi24] = 9223372036854775807
            case -2147483648:
                var v28 uint64 = v8
                _ = v28
            default:
                switch 127 {
                case 2147483647:
                    if (1e10 >= 3.14159) {
                        var v29 string = "foo bar"
                        _ = v29
                    }
                case 1:
                    v30 := []int{}
                    _ = v30
                default:
                    v31 := GenS2{F0: (("hello" + "hello") + "a")}
                    _ = v31
                }
            }
            _ = fi24
            _ = v25
            _ = v26
        }
    case "foo bar":
        v32 := GenS2{F0: "a"}
        _ = v32
    case "\n\t":
        var v33 float64 = (v7 - (v7 - v7))
        _ = v33
    }
    var v34 int32 = (v2 ^ v2)
    v7 -= 3.14159
    var v35 float32 = 0.1
    if ((v6 % ((v6) | 1)) < v6) {
        for fi36 := 0; fi36 < 16; fi36++ {
            var v37 int32 = -84
            _ = fi36
            _ = v37
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
