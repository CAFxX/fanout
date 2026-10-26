package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{921, 4294967295, 32767}
    var v3 string = "hello"
    var v4 uint = 407
    v5 := []string{"a"}
    switch v3 {
    case "hello":
        if true {
            for fi6 := 0; fi6 < 1; fi6++ {
                if false {
                    for fi7 := 0; fi7 < 5; fi7++ {
                        var v8 bool = true
                        _ = fi7
                        _ = v8
                    }
                } else {
                    v9 := GenS1{F0: 2147483648}
                    switch 9223372036854775807 {
                    case 168:
                        v10 := v1[0]
                        _ = v10
                    case -112:
                        for fi11 := 0; fi11 < 5; fi11++ {
                            fmt.Println(fi6, v4, fi6)
                            v3 = "a"
                            if len(v5) > 0 {
                                v5[0] = v3
                            }
                            v1[0] = "hello"
                            v4 = (v4 << (((^(v4))) & 63))
                            var v12 int32 = -327
                            _ = fi11
                            _ = v12
                        }
                    default:
                        if (!(!((-2.25 - (((1e10 - (-2.25 * -2.25)) - 0.1) / 1.5)) <= 1.5))) {
                            switch "foo bar" {
                            case "\n\t":
                                var v13 float64 = ((0.0 / -2.25) * -2.25)
                                _ = v13
                            default:
                                v14 := v1[0]
                                _ = v14
                            }
                            v4 = 255
                            if len(v2) > 0 {
                                v2[0] = 32768
                            }
                        } else {
                            var v15 int16 = -325
                            _ = v15
                        }
                    }
                    v1[2147483647] = ("\n\t" + (((((v3 + (((v3 + ("\n\t" + "foo bar")) + v3) + (v3 + (v3 + v3)))) + (v3 + "a")) + (("a" + (v3 + "hello")) + (((v3 + v3) + (v3 + "")) + ("a" + (("foo bar" + v3) + ("hello" + v3)))))) + (v3 + ((v3 + "foo bar") + v3))) + v3))
                    _ = v9
                }
                _ = fi6
            }
            v3 = (v3 + v3)
        }
    }
    v4 += v4
    var v16 uint = v4
    if false {
        v3 = v3
    }
    v3 = v3
    var v17 int8 = 0
    switch (v3 + ("a" + (v3 + "a"))) {
    case "\n\t":
        v1[1] = "\n\t"
    }
    v5 = append(v5, v3)
    for fi18 := 0; fi18 < 18; fi18++ {
        for fi19 := 0; fi19 < 9; fi19++ {
            v20 := GenS1{F0: 65535}
            fmt.Println(fi19, fi19)
            _ = len(v1)
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    v4 -= v16
    fmt.Println(v16)
    for fi21 := 0; fi21 < 9; fi21++ {
        if false {
            var v22 int = (fi21 % (((^(fi21))) | 1))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v3 = "\n\t"
            _ = v22
        }
        v4 = 65535
        _ = fi21
    }
    delete(v1, 1)
    v23 := v1[1]
    v24 := v1[2147483648]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
