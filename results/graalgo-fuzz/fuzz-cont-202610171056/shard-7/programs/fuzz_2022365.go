package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int; F2 uint32 }

func main() {
    var v1 uint = 1
    v2 := make(map[string]int)
    var v3 string = "a"
    v4 := make(map[string]int)
    v5 := make(map[uint32]string)
    v6 := []int{}
    if ((v3 + (("\n\t" + v3) + (v3 + "\n\t"))) > (v3 + (((v3 + (v3 + v3)) + (v3 + ((v3 + v3) + "foo bar"))) + v3))) {
        switch 128 {
        case 0:
            fmt.Println(v1, v1)
        case 2147483647:
            if len(v6) > 0 {
                v6[0] = 216
            }
        }
        v7 := GenS1{F0: 0}
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v7
    }
    v1 = 883
    var v8 int32 = 692
    switch 4294967295 {
    case 111:
        var v9 int = 2147483647
        _ = v9
    case 128:
        v1 += v1
    default:
        var v10 float64 = 0.1
        _ = v10
    }
    var v11 uint = (v1 % (((v1 % ((v1) | 1))) | 1))
    var v12 int64 = -1
    fmt.Println(v11, v12)
    _ = len(v6)
    fmt.Println(v12)
    if (!(1.5 < 1.5)) {
        v1 -= v11
        for fi13 := 0; fi13 < 4; fi13++ {
            v8 = 1
            var v14 uint16 = 385
            _ = fi13
            _ = v14
        }
        v8 -= (v8 >> ((255) & 63))
    } else {
        v8 += 2147483647
        if false {
            var v15 int = 127
            for fi16 := 0; fi16 < 8; fi16++ {
                if (1 >= v8) {
                    v8 = 1
                } else {
                    if (43 >= 1) {
                        var v17 int32 = (v8 & (v8 >> ((206) & 63)))
                        for fi18 := 0; fi18 < 20; fi18++ {
                            switch "a" {
                            case "":
                                v11 = (^(v11))
                            }
                            if true {
                                fmt.Println(v17)
                            } else {
                                var v19 uint32 = 4294967295
                                _ = v19
                            }
                            switch v15 {
                            case 42:
                                v8 = int32(v15)
                            }
                            _ = fi18
                        }
                        _ = v17
                    } else {
                        v20 := GenS2{F0: -1, F1: 127, F2: 1}
                        v3 = v3
                        switch (v15 % ((-195) | 1)) {
                        case 127:
                            if false {
                                for fi21 := 0; fi21 < 10; fi21++ {
                                    _ = len(v5)
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = fi21
                                }
                            } else {
                                var v22 string = (((v3 + (v3 + ((("hello" + ("hello" + v3)) + (v3 + "\n\t")) + ((v3 + "") + ("hello" + "hello"))))) + v3) + v3)
                                _ = v22
                            }
                        case 32767:
                            v15 += 2147483647
                        }
                        _ = v20
                    }
                    v23 := GenS1{F0: v8}
                    _ = v23
                }
                if false {
                    v12 += (^(v12))
                }
                _ = fi16
            }
            fmt.Println(v3)
            _ = v15
        } else {
            v24 := v2[(v3 + "\n\t")]
            var v25 bool = (-765 < 1000)
            _ = v24
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        v27 := GenS1{F0: v8}
        v11 = v11
        v11 -= 18446744073709551615
        _ = fi26
        _ = v27
    }
    var v28 uint = v1
    v3 = v3
    delete(v4, (v3 + v3))
    var v29 float64 = (1e10 + float64((3.14159 / (3.14159 + (float32(0.1) * 1.5)))))
    v12 = v12
    if false {
        v8 = 42
        var v30 int64 = v12
        if (100 > -62) {
            v12 = 9223372036854775807
            v1 -= (^(v28))
        } else {
            fmt.Println(v8)
            var v31 uint16 = 2
            _ = v31
        }
        _ = v30
    }
    var v32 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
