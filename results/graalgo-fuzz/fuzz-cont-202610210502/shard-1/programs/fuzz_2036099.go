package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint64, p1 string, p2 bool) bool {
    fmt.Println(p0, p2)
    p0 -= 0
    for fi6 := 0; fi6 < 20; fi6++ {
        p0 += p0
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    return (!p2)
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]int)
    var v3 int = 42
    var v4 bool = true
    v5 := []uint32{2147483648, 101}
    if (!(!(4294967295 != 1))) {
        v3 -= (v3 + (v3 << ((1) & 63)))
        switch (("hello" + ("foo bar" + ("foo bar" + ""))) + "\n\t") {
        case "foo bar":
            v7 := GenS1{F0: 65535, F1: -1}
            _ = v7
        case "":
            if true {
                var v8 uint64 = 0
                _ = v8
            } else {
                var v9 uint8 = 135
                _ = v9
            }
        case "hello":
            v10 := GenS1{F0: 18446744073709551615, F1: -9223372036854775808}
            _ = v10
        default:
            v11 := GenS2{F0: 0, F1: 42}
            _ = v11
        }
    } else {
        var v12 float64 = (-2.25 - 0.0)
        v12 -= (v12 * v12)
        v13 := GenS1{F0: 0, F1: 42}
        _ = v12
        _ = v13
    }
    v3 += -1
    if (!true) {
        var v14 uint = 1
        v1["a"] = 396
        for fi15 := 0; fi15 < 15; fi15++ {
            v3 += 32767
            if v4 {
                v4 = ("a" == ((("" + "a") + "") + "a"))
            } else {
                var v16 int = fi15
                _ = v16
            }
            if (!v4) {
                for fi17 := 0; fi17 < 6; fi17++ {
                    fmt.Println(v4)
                    var v18 bool = (v4 || (!(-2.25 >= float64(1e10))))
                    v3 -= fi15
                    _ = fi17
                    _ = v18
                }
                v14 += v14
            } else {
                var v19 uint64 = 4294967295
                var v20 int = 1
                _ = v19
                _ = v20
            }
            _ = fi15
        }
        _ = v14
    }
    fmt.Println(v3)
    fmt.Println(v3)
    fmt.Println(v4, v4)
    var v21 uint16 = 1
    var v22 string = "hello"
    for fi23 := 0; fi23 < 1; fi23++ {
        var v24 uint16 = 42
        v5 = append(v5, 42)
        switch ("foo bar" + v22) {
        case "hello":
            v25 := GenS1{F0: 2, F1: 9223372036854775807}
            _ = v25
        default:
            if v4 {
                v24 = v24
                switch v22 {
                case "hello":
                    switch "" {
                    case "a":
                        switch v3 {
                        case 2147483647:
                            v21 = 230
                        }
                    }
                case "":
                    if (v4 && (4294967295 == 0)) {
                        var v26 uint8 = 42
                        for fi27 := 0; fi27 < 1; fi27++ {
                            v28 := GenS2{F0: (^(v21)), F1: 32767}
                            switch (v22 + v22) {
                            case "foo bar":
                                var v29 int64 = 1
                                _ = v29
                            case "":
                                v30 := GenS3{F0: 42}
                                _ = v30
                            default:
                                if v4 {
                                    _ = len(v1)
                                }
                            }
                            _ = fi27
                            _ = v28
                        }
                        v3 -= v3
                        _ = v26
                    }
                case "foo bar":
                    v21 -= v24
                default:
                    var v31 uint = 4294967295
                    _ = v31
                }
            } else {
                var v32 uint64 = 18446744073709551615
                v33 := v1[v22]
                _ = v32
                _ = v33
            }
        }
        _ = fi23
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
