package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 string = ("foo bar" + ((("foo bar" + "") + "hello") + "\n\t"))
    v2 := []string{"\n\t", v1}
    var v3 int = 127
    var v4 int8 = -25
    var v5 int32 = -1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi6 := 0; fi6 < 5; fi6++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if false {
            var v7 int = (fi6 - (v3 - -59))
            var v8 uint16 = 277
            v7 = (v7 * v3)
            _ = v7
            _ = v8
        }
        if false {
            v3 -= fi6
        } else {
            if (true || (4294967295 == 2)) {
                var v9 float32 = 1e10
                if (false || false) {
                    if (66 >= 1) {
                        _ = len(v2)
                        for fi10 := 0; fi10 < 1; fi10++ {
                            var v11 string = v1
                            if false {
                                switch v1 {
                                case "hello":
                                    var v12 uint16 = 1
                                    _ = v12
                                case "\n\t":
                                    v4 = -33
                                case "a":
                                    switch v1 {
                                    case "foo bar":
                                        v13 := GenS1{F0: 100}
                                        _ = v13
                                    case "hello":
                                        var v14 float64 = 1e10
                                        _ = v14
                                    }
                                default:
                                    var v15 float64 = 3.14159
                                    _ = v15
                                }
                            } else {
                                var v16 uint64 = 410
                                _ = v16
                            }
                            if (!(335 != 2)) {
                                var v17 uint8 = 255
                                var v18 bool = false
                                v17 = v17
                                _ = v17
                                _ = v18
                            }
                            _ = fi10
                            _ = v11
                        }
                    }
                    switch ((v1 + "a") + "hello") {
                    case "foo bar":
                        var v19 uint32 = 65535
                        _ = v19
                    default:
                        var v20 uint = 2
                        _ = v20
                    }
                }
                _ = v9
            }
            v3 = 676
            var v21 uint = 4294967295
            _ = v21
        }
        _ = fi6
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    _ = len(v2)
    switch v3 {
    case -2147483648:
        v1 = ((v1 + v1) + "hello")
    case -678:
        for fi22 := 0; fi22 < 8; fi22++ {
            switch (v3 | fi22) {
            case 1:
                v1 = v1
            case -107:
                v23 := GenS2{F0: 127}
                _ = v23
            default:
                var v24 uint = 912
                _ = v24
            }
            v1 = ("hello" + "")
            var v25 float64 = 1.5
            _ = fi22
            _ = v25
        }
    }
    if false {
        var v26 bool = false
        _ = v26
    }
    v2 = append(v2, v1)
    var v27 int8 = v4
    v28 := GenS2{F0: 2}
    if len(v2) > 0 {
        v2[0] = v1
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 -= (v3 / (((v3 << ((65535) & 63))) | 1))
    var v29 uint16 = 359
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
