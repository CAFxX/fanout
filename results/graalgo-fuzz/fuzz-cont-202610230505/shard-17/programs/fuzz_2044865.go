package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int8) int {
    var v4 string = ("a" + ("hello" + "a"))
    switch 42 {
    case -216:
        var v5 float32 = -2.25
        _ = v5
    case 1:
        v6 := GenS1{F0: 1}
        _ = v6
    case 4294967295:
        v4 = "a"
    }
    _ = p0
    _ = v4
    return 335
}
func genFn1(p0 uint32) uint {
    v7 := GenS1{F0: (^(p0))}
    fmt.Println(p0)
    p0 -= 65535
    var v8 uint8 = 42
    _ = p0
    _ = v7
    _ = v8
    return 523
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 127
    var v3 int32 = -2147483648
    if true {
        if (!((52 >= 1) || ((!false) && (-9223372036854775808 >= -529)))) {
            if (true && (-2147483648 < 128)) {
                v2 = (v2 | 100)
                var v9 int8 = v2
                v3 += (v3 / (((^(v3))) | 1))
                _ = v9
            } else {
                if (32768 < 1) {
                    var v10 int64 = -507
                    v10 += v10
                    for fi11 := 0; fi11 < 15; fi11++ {
                        switch "a" {
                        case "":
                            if false {
                                for fi12 := 0; fi12 < 12; fi12++ {
                                    v10 = (^(v10))
                                    _ = fi12
                                }
                                v13 := GenS1{F0: 641}
                                for fi14 := 0; fi14 < 14; fi14++ {
                                    fmt.Println(v3, v3)
                                    switch "\n\t" {
                                    case "a":
                                        var v15 int32 = 620
                                        _ = v15
                                    case "\n\t":
                                        v16 := []int{(fi11 * genFn0((v2 - -45))), (fi14 + 0), fi11}
                                        _ = v16
                                    }
                                    v17 := GenS1{F0: 4294967295}
                                    _ = fi14
                                    _ = v17
                                }
                                _ = v13
                            } else {
                                for fi18 := 0; fi18 < 16; fi18++ {
                                    v1[0] = false
                                    fmt.Println(v3)
                                    _ = fi18
                                }
                            }
                        case "foo bar":
                            v3 -= v3
                        case "\n\t":
                            v19 := v1[927]
                            _ = v19
                        default:
                            v3 += v3
                        }
                        v3 += -2147483648
                        _ = fi11
                    }
                    _ = v10
                } else {
                    var v20 uint = 324
                    var v21 float64 = 3.14159
                    _ = v20
                    _ = v21
                }
                v22 := GenS1{F0: 2}
                switch ("foo bar" + "\n\t") {
                case "":
                    fmt.Println(v3)
                }
                _ = v22
            }
            if true {
                v23 := []int{int(v3), 32768, -271, -9223372036854775808}
                _ = v23
            }
        }
    }
    var v24 uint32 = 4294967295
    if true {
        var v25 int = 2147483647
        fmt.Println(v25, v3, v2)
        for fi26 := 0; fi26 < 14; fi26++ {
            var v27 uint8 = 42
            _ = fi26
            _ = v27
        }
        _ = v25
    } else {
        v28 := GenS1{F0: v24}
        var v29 uint8 = 127
        v29 = v29
        _ = v28
        _ = v29
    }
    delete(v1, v24)
    v30 := []int{771, 9223372036854775807, -1, genFn0((v2 ^ -87))}
    var v31 string = "hello"
    fmt.Println(v24, v3)
    fmt.Println(v24, v31, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v31
}
