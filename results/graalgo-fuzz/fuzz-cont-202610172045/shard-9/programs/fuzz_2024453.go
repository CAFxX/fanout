package main

import (
    "fmt"
)

func genFn0() int64 {
    var v5 int16 = -463
    v6 := []int{}
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 int32 = 0
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v5 = v5
    _ = v5
    _ = v6
    return -1
}

func main() {
    var v1 uint8 = 128
    v2 := []int{461}
    var v3 float32 = 1e10
    v4 := []int{32768}
    var v10 uint32 = 2147483648
    fmt.Println(v10, v1, v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    _ = len(v4)
    var v11 int = 0
    var v12 uint8 = 128
    var v13 int64 = 1
    switch ("foo bar" + ("a" + "hello")) {
    case "a":
        v3 = (v3 - ((v3 / (3.14159 - -2.25)) * ((v3 * float32(0.1)) * v3)))
    case "foo bar":
        if len(v2) > 0 {
            v2[0] = v11
        }
    }
    var v14 int16 = 0
    v14 = (v14 & (^(v14)))
    if false {
        _ = len(v4)
        for fi15 := 0; fi15 < 17; fi15++ {
            var v16 uint16 = 32767
            if (!true) {
                if false {
                    for fi17 := 0; fi17 < 5; fi17++ {
                        switch (v11 ^ (fi15 >> ((255) & 63))) {
                        case 182:
                            v2 = append(v2, 784)
                        }
                        var v18 uint = 42
                        v11 += -91
                        _ = fi17
                        _ = v18
                    }
                    if (("foo bar" + "\n\t") != "") {
                        fmt.Println(v13, v1, v14)
                        if len(v2) > 0 {
                            v2[0] = (fi15 << ((456) & 63))
                        }
                    }
                    var v19 bool = (-2.25 < ((1.5 - 1.5) + -2.25))
                    _ = v19
                }
                var v20 int64 = (v13 ^ (-(v13)))
                for fi21 := 0; fi21 < 16; fi21++ {
                    v4 = append(v4, (-(fi21)))
                    fmt.Println(v11)
                    if true {
                        switch ("" + "hello") {
                        case "":
                            v13 += -346
                        default:
                            if len(v2) > 0 {
                                v2[0] = 42
                            }
                        }
                    }
                    _ = fi21
                }
                _ = v20
            } else {
                v1 -= v1
                if true {
                    v10 += v10
                    var v22 int8 = -1
                    if len(v2) > 0 {
                        v2[0] = (v11 >> ((18446744073709551615) & 63))
                    }
                    _ = v22
                } else {
                    if false {
                        v13 = genFn0()
                        if true {
                            v13 -= v13
                        } else {
                            v14 -= 42
                            v12 = 218
                        }
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                v3 += (0.0 / v3)
            }
            _ = fi15
            _ = v16
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
