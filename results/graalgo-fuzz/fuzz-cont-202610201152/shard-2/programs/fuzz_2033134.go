package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint16 }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint16) string {
    var v4 int64 = 42
    v5 := []int{936, 0, -723}
    fmt.Println(v4)
    _ = p0
    _ = v4
    _ = v5
    return (("\n\t" + (("foo bar" + "") + "\n\t")) + ("foo bar" + "\n\t"))
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]int)
    v3 := []int{9223372036854775807, 127}
    if false {
        var v6 uint = 65535
        _ = v6
    }
    var v7 int = 32768
    var v8 uint32 = 498
    v8 = (^(v8))
    v7 = v7
    switch v7 {
    case -391:
        v9 := GenS1{F0: 100, F1: (-(v7)), F2: 255}
        _ = v9
    case 1:
        fmt.Println(v7, v7)
    case -9223372036854775808:
        for fi10 := 0; fi10 < 2; fi10++ {
            v7 = (v7 / ((fi10) | 1))
            for fi11 := 0; fi11 < 2; fi11++ {
                delete(v1, v7)
                v7 += fi10
                if (1 != v7) {
                    fmt.Println(v7, fi10, v8)
                    if false {
                        switch "hello" {
                        case "a":
                            if len(v3) > 0 {
                                v3[0] = (-(v7))
                            }
                            switch ("hello" + "hello") {
                            case "hello":
                                fmt.Println(v8, fi11, fi10)
                            default:
                                var v12 int8 = -40
                                _ = v12
                            }
                        default:
                            var v13 string = ("a" + "foo bar")
                            _ = v13
                        }
                        v8 = v8
                        v3 = append(v3, 240)
                    } else {
                        v8 = 2147483648
                        var v14 int32 = -1
                        _ = v14
                    }
                    v15 := GenS1{F0: -128, F1: 32767, F2: 994}
                    if (((!(true && (true && true))) || ((-2147483648 != 42) && false)) && true) {
                        var v16 int = fi11
                        _ = v16
                    } else {
                        switch (fi10 - (^(v7))) {
                        case 578:
                            var v17 string = ("a" + "foo bar")
                            _ = v17
                        case 9223372036854775807:
                            v7 = 32767
                        case 310:
                            v18 := GenS1{F0: 87, F1: (fi10 >> ((4294967295) & 63)), F2: 65535}
                            _ = v18
                        default:
                            var v19 uint = 65535
                            _ = v19
                        }
                    }
                    _ = v15
                } else {
                    v20 := GenS2{F0: 42}
                    var v21 bool = (!(!false))
                    v21 = ((42 >= 2) || true)
                    if (!(!false)) {
                        fmt.Println(v7, fi11)
                    }
                    if (38 == 42) {
                        var v22 int64 = 0
                        if v21 {
                            v7 = fi10
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        _ = v22
                    }
                    _ = v20
                    _ = v21
                }
                v3 = append(v3, 127)
                _ = fi11
            }
            _ = fi10
        }
    default:
        var v23 int8 = 42
        _ = v23
    }
    v24 := GenS2{F0: -725}
    delete(v1, (v7 % ((v7) | 1)))
    var v25 float32 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v7
    _ = v8
}
