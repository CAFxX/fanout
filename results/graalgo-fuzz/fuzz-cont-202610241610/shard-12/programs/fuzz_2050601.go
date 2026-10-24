package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64 }
func genFn0() string {
    var v6 int8 = 100
    fmt.Println(v6)
    var v7 int32 = 2147483647
    v8 := []int{1, 1, 32767, 784}
    var v9 int64 = 1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return ("a" + "foo bar")
}

func main() {
    var v1 uint8 = 142
    var v2 uint = 65535
    var v3 uint32 = 8
    v4 := make(map[int]int)
    var v5 int32 = -95
    if true {
        if false {
            v10 := GenS1{F0: ("\n\t" + "\n\t")}
            for fi11 := 0; fi11 < 8; fi11++ {
                var v12 int32 = (v5 + (v5 % (((-(v5))) | 1)))
                _ = fi11
                _ = v12
            }
            _ = v10
        }
    }
    v3 = (v3 << (((v2 << ((2) & 63))) & 63))
    v13 := []int{42, -1, -715, 127}
    var v14 bool = (uint32(v2) < v3)
    v2 -= v2
    var v15 uint32 = v3
    v3 = (v15 >> ((689) & 63))
    var v16 string = "foo bar"
    if ((0.0 / (0.1 * float32(0.1))) != 1e10) {
        var v17 int = 9223372036854775807
        v3 = (v3 - (v15 - v3))
        v18 := GenS2{F0: 88}
        _ = v17
        _ = v18
    } else {
        v16 = (("hello" + "a") + (("hello" + "a") + "hello"))
        v13 = append(v13, 32767)
    }
    v1 = v1
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    var v19 uint = (v2 / ((v2) | 1))
    var v20 string = ("a" + (("hello" + v16) + (v16 + ("" + ((("" + v16) + "foo bar") + (((v16 + "\n\t") + ("hello" + ("" + "hello"))) + "a"))))))
    var v21 int64 = 42
    var v22 uint8 = v1
    if v14 {
        _ = len(v4)
        switch v20 {
        case "\n\t":
            var v23 uint = v2
            _ = v23
        case "a":
            fmt.Println(v2)
        default:
            if ((1 != 9) || false) {
                if v14 {
                    if ((^(v19)) != 4294967295) {
                        _ = len(v13)
                    } else {
                        switch 9223372036854775807 {
                        case 768:
                            var v24 int32 = -1
                            _ = v24
                        case 128:
                            switch "\n\t" {
                            case "a":
                                for fi25 := 0; fi25 < 5; fi25++ {
                                    var v26 uint32 = (v15 + (v15 * v3))
                                    var v27 int8 = 0
                                    var v28 float64 = 3.14159
                                    _ = fi25
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                            case "foo bar":
                                fmt.Println(v5, v5, v22)
                            case "hello":
                                fmt.Println(v22, v3)
                            default:
                                switch 42 {
                                case 16:
                                    if v14 {
                                        fmt.Println(v3, v19, v2)
                                        var v29 float64 = ((((0.0 - 0.0) * (float64(float32(1e10)) * ((1e10 + ((1.5 - 0.0) + 1.5)) / 1.0))) - 1e10) * 3.14159)
                                        var v30 uint8 = 204
                                        _ = v29
                                        _ = v30
                                    }
                                case 0:
                                    v31 := GenS3{F0: (v21 | v21)}
                                    _ = v31
                                }
                            }
                        case 32767:
                            v32 := GenS1{F0: ("a" + "a")}
                            _ = v32
                        }
                        if true {
                            var v33 uint64 = 18446744073709551615
                            var v34 uint64 = (v33 % ((v33) | 1))
                            var v35 float32 = 0.1
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                    }
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                }
            }
        }
    }
    v16 = v20
    var v36 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
