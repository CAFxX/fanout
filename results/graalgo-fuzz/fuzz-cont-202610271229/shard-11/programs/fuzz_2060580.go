package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint8; F1 int8 }
func genFn0(p0 string, p1 string, p2 bool) uint64 {
    var v5 float64 = 1.5
    var v6 int8 = 127
    v7 := []int{-379, 1, int(v6)}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 string = ""
    var v2 uint32 = 42
    v3 := make(map[int]string)
    var v4 int16 = 0
    switch v1 {
    case "":
        v8 := GenS1{F0: 2}
        _ = v8
    case "\n\t":
        var v9 int8 = -128
        _ = v9
    }
    v10 := GenS1{F0: 1}
    switch ("foo bar" + v1) {
    case "hello":
        v11 := GenS1{F0: 908}
        _ = v11
    case "a":
        v2 += 2
    case "foo bar":
        var v12 int32 = -878
        _ = v12
    }
    v1 = v1
    delete(v3, -38)
    _ = len(v3)
    v1 = ((("a" + v1) + "") + v1)
    v1 = v1
    var v13 string = "\n\t"
    v1 = v13
    v14 := []int{2147483647, 9223372036854775807}
    if true {
        for fi15 := 0; fi15 < 3; fi15++ {
            if true {
                v1 = ("" + "\n\t")
                v16 := GenS2{F0: 1, F1: -786}
                _ = v16
            } else {
                if true {
                    var v17 int64 = -1
                    var v18 bool = (1 < 2)
                    _ = v17
                    _ = v18
                } else {
                    switch "a" {
                    case "a":
                        if false {
                            v14 = append(v14, 4294967295)
                            for fi19 := 0; fi19 < 10; fi19++ {
                                delete(v3, fi19)
                                v1 = "hello"
                                if false {
                                    switch -412 {
                                    case -9223372036854775808:
                                        if ((v4 + (^(v4))) != v4) {
                                            fmt.Println(v1)
                                        }
                                    default:
                                        switch (fi19 * fi15) {
                                        case 683:
                                            if true {
                                                v20 := GenS3{F0: 16, F1: 52}
                                                if (!(true && false)) {
                                                    fmt.Println(v2, v4)
                                                    var v21 uint64 = 1
                                                    v13 = "foo bar"
                                                    _ = v21
                                                } else {
                                                    var v22 float32 = 1e10
                                                    _ = v22
                                                }
                                                var v23 int32 = 2147483647
                                                _ = v20
                                                _ = v23
                                            } else {
                                                _ = len(v3)
                                                switch fi15 {
                                                case 0:
                                                    fmt.Println(v4, fi19, v1)
                                                default:
                                                    fmt.Println(v1)
                                                }
                                            }
                                        default:
                                            v4 = v4
                                        }
                                    }
                                    var v24 uint = 18446744073709551615
                                    _ = v24
                                } else {
                                    var v25 float32 = -2.25
                                    v14 = append(v14, (-(fi15)))
                                    v2 = 338
                                    _ = v25
                                }
                                var v26 int16 = 1000
                                _ = fi19
                                _ = v26
                            }
                            var v27 bool = true
                            _ = v27
                        }
                    }
                }
                for fi28 := 0; fi28 < 6; fi28++ {
                    var v29 int64 = 9223372036854775807
                    switch v13 {
                    case "hello":
                        v14 = append(v14, fi15)
                    case "a":
                        switch v1 {
                        case "foo bar":
                            switch v13 {
                            case "":
                                var v30 int16 = v4
                                _ = v30
                            }
                        case "\n\t":
                            v2 = (v2 - (v2 + (^(v2))))
                        default:
                            var v31 int8 = -79
                            _ = v31
                        }
                    }
                    _ = len(v14)
                    _ = fi28
                    _ = v29
                }
            }
            _ = fi15
        }
        if (1e10 >= ((0.0 - 0.1) * (1e10 - (0.0 * 3.14159)))) {
            v4 = v4
            var v32 int64 = 1
            if true {
                fmt.Println(v32, v13, v32)
                v33 := v3[42]
                v32 -= -969
                _ = v33
            }
            _ = v32
        }
        var v34 uint64 = 152
        _ = v34
    } else {
        if (v2 > (^(v2))) {
            var v35 int64 = 1
            v35 = (v35 * v35)
            _ = v35
        }
    }
    v13 = v1
    var v36 int64 = 536
    var v37 string = v13
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
