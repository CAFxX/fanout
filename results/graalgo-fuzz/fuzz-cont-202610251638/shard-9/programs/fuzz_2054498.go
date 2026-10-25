package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 int = 0
    var v2 int16 = -32768
    var v3 uint64 = 522
    v4 := make(map[string]bool)
    var v5 int8 = -1
    fmt.Println(v2, v1, v5)
    v6 := []int{(-(v1)), v1, v1, (v1 % ((v1) | 1))}
    if false {
        v7 := GenS1{F0: v5}
        _ = v7
    } else {
        v2 = 1000
        if (false || false) {
            for fi8 := 0; fi8 < 18; fi8++ {
                v2 -= v2
                v9 := GenS1{F0: (v5 & v5)}
                v2 = (-(v2))
                _ = fi8
                _ = v9
            }
            for fi10 := 0; fi10 < 20; fi10++ {
                switch 32768 {
                case 732:
                    v5 -= (v5 | (v5 % ((26) | 1)))
                case 128:
                    if false {
                        if false {
                            v11 := v4["\n\t"]
                            for fi12 := 0; fi12 < 16; fi12++ {
                                v1 = 32767
                                _ = fi12
                            }
                            _ = v11
                        } else {
                            if true {
                                var v13 int32 = 0
                                _ = v13
                            }
                        }
                    } else {
                        if false {
                            if (!(1e10 > (3.14159 - 0.0))) {
                                fmt.Println(v5, v5, v5)
                            }
                        } else {
                            if ((255 <= 42) && (!true)) {
                                if false {
                                    fmt.Println(v3)
                                    v1 = int(v3)
                                    var v14 int32 = 671
                                    _ = v14
                                }
                                v1 = (fi10 + v1)
                            }
                            fmt.Println(fi10, v5, v3)
                            var v15 uint32 = 2
                            _ = v15
                        }
                    }
                case -1:
                    var v16 float32 = (((0.1 + 3.14159) - 0.1) - ((float32(float64(0.0)) * float32((1e10 + ((0.1 * float64(-2.25)) * (-2.25 * ((((3.14159 + (0.0 + ((1e10 + (0.1 * 0.0)) / -2.25))) / 3.14159) * 1e10) * 0.1)))))) + float32(3.14159)))
                    _ = v16
                }
                _ = fi10
            }
        }
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 uint16 = 65535
            var v19 uint16 = 4
            if ((!false) || (v18 > v19)) {
                v20 := GenS1{F0: (v5 & -38)}
                for fi21 := 0; fi21 < 16; fi21++ {
                    v6 = append(v6, v1)
                    _ = fi21
                }
                v3 = (v3 % ((v3) | 1))
                _ = v20
            }
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    if len(v6) > 0 {
        v6[0] = 127
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        fmt.Println(v5)
        switch "foo bar" {
        case "\n\t":
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        case "":
            v3 -= 0
        case "foo bar":
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        _ = fi22
    }
    v4["hello"] = (2147483647 >= 1)
    v23 := GenS1{F0: v5}
    var v24 int16 = v2
    fmt.Println(v2, v5, v24)
    v1 = -639
    switch (v1 - v1) {
    case -1:
        delete(v4, ("hello" + ("" + (("" + "\n\t") + "a"))))
    }
    switch -2147483648 {
    case 127:
        var v25 bool = false
        _ = v25
    default:
        v6 = append(v6, 32768)
    }
    delete(v4, "")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
