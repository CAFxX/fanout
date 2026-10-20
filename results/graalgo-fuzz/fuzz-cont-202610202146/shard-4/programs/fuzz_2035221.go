package main

import (
    "fmt"
)

func genFn0(p0 int) uint16 {
    var v5 string = ("" + ("" + (("hello" + "") + "")))
    v5 = ((v5 + "foo bar") + v5)
    _ = p0
    _ = v5
    return 65535
}
func genFn1(p0 int, p1 uint, p2 uint) uint16 {
    v6 := []int{(p0 - (p0 + (^(p0)))), 284}
    p2 -= 346
    for fi7 := 0; fi7 < 2; fi7++ {
        _ = len(v6)
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 65535
}

func main() {
    var v1 int64 = 42
    var v2 int = 578
    v3 := []uint32{49, 559, 2147483647, 2}
    v4 := []uint32{65535}
    fmt.Println(v1, v1)
    var v8 int16 = 32767
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch v2 {
    case -795:
        if len(v4) > 0 {
            v4[0] = 817
        }
    case 128:
        for fi9 := 0; fi9 < 12; fi9++ {
            switch "hello" {
            case "":
                if len(v3) > 0 {
                    v3[0] = 2
                }
            default:
                if len(v3) > 0 {
                    v3[0] = 393
                }
            }
            if false {
                var v10 string = (((("" + (("hello" + "a") + "hello")) + "\n\t") + (("hello" + "foo bar") + "foo bar")) + "hello")
                var v11 int8 = 64
                _ = v10
                _ = v11
            } else {
                if false {
                    for fi12 := 0; fi12 < 7; fi12++ {
                        if false {
                            var v13 string = "a"
                            _ = v13
                        } else {
                            if (!(94 > 100)) {
                                var v14 int32 = -466
                                var v15 uint = 1
                                switch "hello" {
                                case "foo bar":
                                    var v16 string = "hello"
                                    _ = v16
                                default:
                                    var v17 uint64 = 42
                                    _ = v17
                                }
                                _ = v14
                                _ = v15
                            }
                        }
                        _ = fi12
                    }
                    v8 -= v8
                }
            }
            _ = fi9
        }
    case 9223372036854775807:
        for fi18 := 0; fi18 < 12; fi18++ {
            var v19 int64 = (^(v1))
            _ = fi18
            _ = v19
        }
    default:
        fmt.Println(v2, v8)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 += 706
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v2, v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
