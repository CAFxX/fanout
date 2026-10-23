package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0() uint16 {
    if (42 != 2147483648) {
        var v6 bool = (0 < 521)
        if v6 {
            v7 := GenS1{F0: "hello"}
            var v8 float64 = 3.14159
            _ = v7
            _ = v8
        }
        _ = v6
    }
    var v9 uint = 1
    var v10 uint8 = 255
    _ = v10
    _ = v9
    return 2
}

func main() {
    v1 := []string{""}
    var v2 uint64 = 2
    var v3 int16 = -154
    var v4 uint32 = 0
    v5 := []string{(("foo bar" + ("hello" + "foo bar")) + ("hello" + (("\n\t" + "foo bar") + "a")))}
    fmt.Println(v4, v4)
    _ = len(v5)
    fmt.Println(v4)
    switch 128 {
    case 1:
        v4 = (v4 + (^(v4)))
    case 849:
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 string = ""
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi11
            _ = v12
        }
    case -2147483648:
        if false {
            v4 -= 42
            var v13 float64 = 3.14159
            _ = v13
        }
    }
    var v14 uint64 = 18446744073709551615
    if false {
        if (!true) {
            fmt.Println(v4, v3)
        } else {
            var v15 int = 9223372036854775807
            for fi16 := 0; fi16 < 18; fi16++ {
                if true {
                    v4 = 2147483648
                    if false {
                        v1 = append(v1, "hello")
                    } else {
                        fmt.Println(v2, v2)
                        if ((-1 >= -18) && false) {
                            if false {
                                var v17 int64 = -235
                                fmt.Println(v3, fi16)
                                _ = len(v5)
                                _ = v17
                            }
                            var v18 uint32 = v4
                            switch "hello" {
                            case "hello":
                                v4 = (v18 & (v18 + (v4 / ((uint32(v2)) | 1))))
                            }
                            _ = v18
                        } else {
                            var v19 uint32 = v4
                            var v20 int32 = 865
                            var v21 int16 = (-(v3))
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                    }
                }
                _ = fi16
            }
            _ = v15
        }
    } else {
        if false {
            for fi22 := 0; fi22 < 5; fi22++ {
                var v23 int8 = int8(fi22)
                _ = fi22
                _ = v23
            }
            v2 += v14
        }
        var v24 int64 = -968
        if len(v5) > 0 {
            v5[0] = "a"
        }
        _ = v24
    }
    var v25 float32 = (((1.5 - 0.0) - 3.14159) / ((0.0 / 1.0) - 0.1))
    switch 9223372036854775807 {
    case 2147483647:
        if len(v5) > 0 {
            v5[0] = (("a" + "") + "")
        }
    case 9223372036854775807:
        switch "" {
        case "a":
            fmt.Println(v2)
        }
    default:
        var v26 int = 1
        _ = v26
    }
    var v27 int8 = 0
    fmt.Println(v25)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
