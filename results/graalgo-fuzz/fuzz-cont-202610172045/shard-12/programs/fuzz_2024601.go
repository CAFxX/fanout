package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }

func main() {
    var v1 string = "foo bar"
    var v2 uint32 = 1
    v3 := []uint32{(v2 % ((v2) | 1)), v2, (^(v2)), (^(v2))}
    var v4 int = 4294967295
    var v5 bool = true
    var v6 int32 = 0
    switch v4 {
    case 128:
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    case 32767:
        fmt.Println(v4)
    case 1:
        fmt.Println(v6, v1)
    default:
        var v7 int16 = 1000
        _ = v7
    }
    _ = len(v3)
    for fi8 := 0; fi8 < 17; fi8++ {
        v2 = v2
        _ = fi8
    }
    var v9 string = v1
    var v10 string = ((v1 + v9) + (((v9 + ("" + ("" + ((("" + (v1 + "foo bar")) + v1) + ((v9 + ("" + "foo bar")) + v1))))) + (((("hello" + v9) + v9) + v1) + v1)) + (((v1 + "\n\t") + "foo bar") + (v9 + ("" + ((v1 + v1) + "a"))))))
    v10 = (v1 + v10)
    switch ((("hello" + ("a" + v10)) + v9) + "") {
    case "foo bar":
        if (9223372036854775807 < -185) {
            v3 = append(v3, v2)
            v1 = ""
        }
    case "a":
        v4 = (v4 & v4)
    default:
        fmt.Println(v1)
    }
    var v11 bool = false
    var v12 uint16 = 255
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v2, v12)
    v5 = ((!(v2 == v2)) && (1 == 0))
    fmt.Println(v12)
    if ((4294967295 != 2) || v5) {
        if (v12 >= v12) {
            switch v9 {
            case "":
                var v13 uint32 = (v2 << ((845) & 63))
                _ = v13
            default:
                v5 = v11
            }
            v1 = (("hello" + "hello") + v9)
            for fi14 := 0; fi14 < 10; fi14++ {
                if v5 {
                    v9 = "hello"
                    if true {
                        v15 := GenS1{F0: 42, F1: 0}
                        if true {
                            var v16 uint16 = (^(v12))
                            if len(v3) > 0 {
                                v3[0] = 4294967295
                            }
                            _ = v16
                        }
                        switch v4 {
                        case 2147483647:
                            _ = len(v3)
                        case 9223372036854775807:
                            v9 = (("\n\t" + v9) + "hello")
                        }
                        _ = v15
                    } else {
                        v17 := GenS1{F0: -1, F1: (-(v6))}
                        for fi18 := 0; fi18 < 8; fi18++ {
                            switch v9 {
                            case "\n\t":
                                v9 = v9
                            }
                            _ = fi18
                        }
                        for fi19 := 0; fi19 < 17; fi19++ {
                            v10 = "\n\t"
                            var v20 float32 = -2.25
                            _ = fi19
                            _ = v20
                        }
                        _ = v17
                    }
                    if false {
                        var v21 uint16 = v12
                        v1 = ("hello" + v9)
                        _ = v21
                    }
                }
                _ = fi14
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
