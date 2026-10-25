package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 int16 }
func genFn0(p0 uint8, p1 int, p2 int) bool {
    p0 = p0
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 uint8 = (^(p0))
        fmt.Println(p1)
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return false
}

func main() {
    v1 := []string{"hello"}
    v2 := make(map[string]bool)
    v3 := make(map[string]string)
    var v4 int8 = 0
    var v5 uint64 = 18446744073709551615
    v4 = (v4 * (v4 / ((v4) | 1)))
    var v8 uint32 = 2147483647
    fmt.Println(v5, v8, v4)
    switch "a" {
    case "\n\t":
        fmt.Println(v4, v8, v5)
    case "":
        if true {
            v9 := GenS1{F0: -627}
            v5 = v5
            var v10 int32 = 1
            _ = v10
            _ = v9
        }
    case "hello":
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 uint8 = 0
            _ = fi11
            _ = v12
        }
    default:
        var v13 uint = 87
        _ = v13
    }
    if false {
        v5 = v5
    }
    v1 = append(v1, "\n\t")
    for fi14 := 0; fi14 < 19; fi14++ {
        fmt.Println(fi14, v4)
        _ = fi14
    }
    fmt.Println(v4)
    switch 4294967295 {
    case -1:
        v4 = (v4 % ((v4) | 1))
    case 0:
        v8 = 2147483648
    case 127:
        switch "hello" {
        case "a":
            v8 = (v8 | (v8 ^ v8))
        default:
            for fi15 := 0; fi15 < 17; fi15++ {
                for fi16 := 0; fi16 < 9; fi16++ {
                    v17 := GenS2{F0: (fi15 / ((fi16) | 1)), F1: -569}
                    _ = fi16
                    _ = v17
                }
                _ = fi15
            }
        }
    default:
        if true {
            switch -395 {
            case -130:
                fmt.Println(v4, v4, v5)
            case -2147483648:
                fmt.Println(v4, v5)
            default:
                if true {
                    var v18 uint16 = 255
                    if (true || (v5 != v5)) {
                        v18 = 2
                        v4 = 42
                        for fi19 := 0; fi19 < 12; fi19++ {
                            var v20 bool = true
                            for fi21 := 0; fi21 < 3; fi21++ {
                                var v22 string = "a"
                                var v23 float32 = (float32(((0.0 * 1.5) * float64(1.5))) - (float32(0.1) + 3.14159))
                                if (0.0 >= (0.0 / 1.0)) {
                                    delete(v3, ((("hello" + (v22 + v22)) + v22) + v22))
                                    _ = len(v2)
                                } else {
                                    if (-613 == 0) {
                                        v1 = append(v1, "foo bar")
                                        switch "\n\t" {
                                        case "hello":
                                            v22 = "hello"
                                        case "a":
                                            v3[(v22 + v22)] = ((("foo bar" + "\n\t") + v22) + "a")
                                        default:
                                            fmt.Println(v18, v4)
                                        }
                                        var v24 uint32 = v8
                                        _ = v24
                                    } else {
                                        if v20 {
                                            if (!v20) {
                                                delete(v2, (v22 + v22))
                                            }
                                            v5 -= 2
                                        }
                                    }
                                }
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                            v1 = append(v1, ("a" + ("\n\t" + ("a" + (("" + "") + "foo bar")))))
                            _ = fi19
                            _ = v20
                        }
                    }
                    _ = v18
                } else {
                    _ = len(v3)
                    if false {
                        v4 = (v4 >> ((18446744073709551615) & 63))
                    } else {
                        v5 = v5
                    }
                }
            }
            v2["a"] = true
        }
    }
    v8 += (v8 ^ v8)
    v25 := GenS1{F0: -343}
    var v26 uint64 = v5
    v4 = v4
    var v27 uint64 = (^(v5))
    v5 = v27
    _ = len(v3)
    fmt.Println(v5, v26)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
