package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int64 }

func main() {
    v1 := make(map[string]string)
    var v2 int = -1
    v3 := []int{v2, v2, 1}
    v4 := make(map[string]string)
    var v5 string = ""
    var v6 uint64 = 0
    var v7 int8 = -1
    if true {
        v7 += (v7 * v7)
        if true {
            if true {
                if false {
                    if (1000 <= -32768) {
                        if true {
                            var v8 uint32 = 42
                            v2 = (v2 ^ 64)
                            if (191 < -931) {
                                var v9 int = (v2 << ((2) & 63))
                                _ = v9
                            } else {
                                switch v2 {
                                case 4294967295:
                                    v8 = (v8 - v8)
                                case -873:
                                    v7 += (v7 >> ((134) & 63))
                                default:
                                    var v10 uint8 = 225
                                    _ = v10
                                }
                                if true {
                                    _ = len(v3)
                                    var v11 string = (v5 + "foo bar")
                                    v5 = ((("hello" + v11) + "\n\t") + v11)
                                    _ = v11
                                } else {
                                    v12 := GenS1{F0: (-9223372036854775808 == -1), F1: 284, F2: 1}
                                    _ = v12
                                }
                            }
                            _ = v8
                        }
                    }
                    switch "foo bar" {
                    case "":
                        v13 := GenS1{F0: ((float32(float64((-2.25 + 1.5))) == 1e10) && (32767 > 773)), F1: 211, F2: -9223372036854775808}
                        _ = v13
                    default:
                        var v14 uint32 = 42
                        _ = v14
                    }
                }
                var v15 int = v2
                _ = v15
            }
            for fi16 := 0; fi16 < 7; fi16++ {
                if ((v7 | v7) >= v7) {
                    if (v7 < v7) {
                        if true {
                            v6 += v6
                            v6 = v6
                        } else {
                            fmt.Println(v5)
                            if false {
                                if false {
                                    var v17 uint = 18446744073709551615
                                    switch v5 {
                                    case "\n\t":
                                        v17 += uint(v6)
                                    }
                                    _ = v17
                                }
                            }
                        }
                        var v18 int = v2
                        v18 += 127
                        _ = v18
                    }
                } else {
                    var v19 float32 = -2.25
                    var v20 uint = 42
                    _ = len(v4)
                    _ = v19
                    _ = v20
                }
                _ = fi16
            }
        } else {
            switch v2 {
            case -9223372036854775808:
                for fi21 := 0; fi21 < 11; fi21++ {
                    _ = len(v4)
                    _ = fi21
                }
            case 1:
                var v22 float64 = float64((((-2.25 + -2.25) / 1e10) - -2.25))
                _ = v22
            case 128:
                v7 = (-(v7))
            default:
                v7 = v7
            }
            if true {
                v6 += 0
                if false {
                    var v23 bool = true
                    var v24 uint = 547
                    _ = v23
                    _ = v24
                }
                var v25 int16 = -733
                _ = v25
            }
            v2 += (-(v2))
        }
    }
    if len(v3) > 0 {
        v3[0] = 32767
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        switch (fi26 % (((v2 + (v2 >> ((0) & 63)))) | 1)) {
        case -1:
            for fi27 := 0; fi27 < 19; fi27++ {
                delete(v4, ("" + "a"))
                _ = fi27
            }
        }
        _ = fi26
    }
    delete(v1, (("\n\t" + (v5 + v5)) + v5))
    switch ("foo bar" + "\n\t") {
    case "hello":
        if len(v3) > 0 {
            v3[0] = 42
        }
    }
    var v28 uint16 = 32767
    for fi29 := 0; fi29 < 20; fi29++ {
        switch (v5 + v5) {
        case "foo bar":
            v5 = ""
        }
        _ = fi29
    }
    var v30 int8 = v7
    var v31 float32 = 1.5
    v31 = (v31 * ((-2.25 * 1e10) + ((((v31 * float32((-2.25 - 1e10))) + (1e10 + float32(1e10))) * -2.25) * (((-2.25 * v31) * 1e10) * ((v31 * v31) + v31)))))
    v6 = 0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
