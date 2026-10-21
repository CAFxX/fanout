package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 bool; F1 int16 }
func genFn0() int16 {
    v6 := []int{1, -1, 0, 42}
    switch ("hello" + "a") {
    case "":
        var v7 string = (("foo bar" + "hello") + "foo bar")
        _ = v7
    }
    var v8 uint64 = 942
    _ = v6
    _ = v8
    return 42
}

func main() {
    var v1 string = "a"
    var v2 int64 = -46
    v3 := []uint32{4294967295}
    var v4 bool = false
    var v5 bool = (395 >= 2)
    v9 := GenS1{F0: -46, F1: 128}
    fmt.Println(v1)
    v1 = "foo bar"
    var v10 string = v1
    fmt.Println(v5)
    _ = len(v3)
    fmt.Println(v2, v4, v5)
    v11 := GenS1{F0: -128, F1: 128}
    v5 = (!v5)
    if false {
        var v12 uint8 = 155
        _ = len(v3)
        _ = v12
    } else {
        v3 = append(v3, 792)
        v2 += v2
        if (-2.25 != (0.1 * -2.25)) {
            switch -567 {
            case -172:
                v13 := GenS2{F0: v5, F1: -1}
                _ = v13
            case -2147483648:
                for fi14 := 0; fi14 < 7; fi14++ {
                    v15 := GenS2{F0: (v5 || v5), F1: genFn0()}
                    v16 := GenS1{F0: -76, F1: 86}
                    switch fi14 {
                    case -53:
                        if v4 {
                            var v17 bool = v4
                            if len(v3) > 0 {
                                v3[0] = 2
                            }
                            v4 = (9223372036854775807 == v2)
                            _ = v17
                        }
                    case -774:
                        var v18 int64 = (v2 + -799)
                        _ = v18
                    case -445:
                        var v19 uint64 = 42
                        _ = v19
                    default:
                        if (4294967295 == 598) {
                            var v20 uint = 42
                            _ = v20
                        }
                    }
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            default:
                var v21 uint16 = 32767
                _ = v21
            }
            v5 = true
            var v22 bool = ((v2 % ((-1) | 1)) < v2)
            _ = v22
        } else {
            var v23 uint8 = 169
            _ = v23
        }
    }
    fmt.Println(v5, v2, v2)
    for fi24 := 0; fi24 < 5; fi24++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi24
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        fmt.Println(v2)
        _ = fi25
    }
    v2 += 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
